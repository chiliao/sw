#include <base.h>
#include <hal.hpp>
#include <hal_lock.hpp>
#include <hal_state.hpp>
#include <proxy.hpp>
#include <proxy_svc.hpp>
#include <pd_api.hpp>
#include <lif_manager.hpp>

namespace hal {

void *
proxy_get_key_func (void *entry)
{
    HAL_ASSERT(entry != NULL);
    return (void *)&(((proxy_t *)entry)->type);
}

uint32_t
proxy_compute_hash_func (void *key, uint32_t ht_size)
{
    return utils::hash_algo::fnv_hash(key, sizeof(types::ProxyType)) % ht_size;
}

bool
proxy_compare_key_func (void *key1, void *key2)
{
    HAL_ASSERT((key1 != NULL) && (key2 != NULL));
    if (*(types::ProxyType *)key1 == *(types::ProxyType *)key2) {
        return true;
    }
    return false;
}

void *
proxy_get_handle_key_func (void *entry)
{
    HAL_ASSERT(entry != NULL);
    return (void *)&(((proxy_t *)entry)->hal_handle);
}

uint32_t
proxy_compute_handle_hash_func (void *key, uint32_t ht_size)
{
    return utils::hash_algo::fnv_hash(key, sizeof(hal_handle_t)) % ht_size;
}

bool
proxy_compare_handle_key_func (void *key1, void *key2)
{
    HAL_ASSERT((key1 != NULL) && (key2 != NULL));
    if (*(hal_handle_t *)key1 == *(hal_handle_t *)key2) {
        return true;
    }
    return false;
}

//------------------------------------------------------------------------------
// validate an incoming PROXY enable request
// TODO:
// 1. check if PROXY exists already
//------------------------------------------------------------------------------
static hal_ret_t
validate_proxy_enable (ProxySpec& spec, ProxyResponse *rsp)
{
    // must have key-handle set
    if (spec.proxy_type() == types::PROXY_TYPE_NONE) {
        rsp->set_api_status(types::API_STATUS_PROXY_TYPE_INVALID);
        return HAL_RET_INVALID_ARG;
    }

    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// insert this Proxy in all meta data structures
//------------------------------------------------------------------------------
static inline hal_ret_t
add_proxy_to_db (proxy_t *proxy)
{
    hal_ret_t   ret;
    ret = g_hal_state->proxy_hal_handle_ht()->insert(proxy,
                                               &proxy->hal_handle_ht_ctxt);
    HAL_TRACE_DEBUG("Return1: {} for type: {}", ret, proxy->type);
    ret = g_hal_state->proxy_type_ht()->insert(proxy, &proxy->ht_ctxt);
    HAL_TRACE_DEBUG("Return2: {}, type: {}", ret, proxy->type) ;
    return HAL_RET_OK;
}

//-----------------------------------------------------------------------------
// API to allocate and program QID
//-----------------------------------------------------------------------------

hal_ret_t
proxy_allocate_qid(types::ProxyType type,
                   lif_id_t&        lif,
                   qtype_t&         qtype,
                   qid_t&           qid)
{
    proxy_t             *proxy = NULL;
    indexer::status     rs;

    proxy = find_proxy_by_type(type);
    if(proxy == NULL) {
        HAL_TRACE_ERR("Proxy not found for the type, {}", type);
        return HAL_RET_PROXY_NOT_FOUND;
    }

    qtype = proxy->qtype;
    
    rs = proxy->qid_idxr_->alloc((uint32_t *)&qid);
    if(rs != indexer::SUCCESS) {
        HAL_TRACE_ERR("Error in qid allocation, err: {}", rs);
        return HAL_RET_NO_RESOURCE;
    }

    HAL_TRACE_DEBUG("QID Allocator: qtype: {}, qid: {}", qtype, qid);
    return HAL_RET_OK;
}

hal_ret_t
proxy_program_qid(types::ProxyType type, lif_id_t lif, qtype_t qtype, qid_t qid)
{
    return HAL_RET_OK;
}

hal_ret_t
proxy_allocate_program_qid(types::ProxyType type,
                           lif_id_t    lif,
                           qtype_t& qtype, 
                           qid_t& qid) 
{
    hal_ret_t ret;
    ret = proxy_allocate_qid(type, lif, qtype, qid);
    if(ret!= HAL_RET_OK) {
        return ret;    
    }

    return proxy_program_qid(type, lif, qtype, qid);
}

hal_ret_t 
proxy_program_lif(proxy_t* proxy)
{
    hal_ret_t           ret = HAL_RET_OK;
    intf::LifSpec       lif_spec;
    intf::LifResponse   rsp;
    LIFQStateParams     qstate_params = {0};
    
    // Create LIF 
    lif_spec.mutable_key_or_handle()->set_lif_id(proxy->lif_id); 
    lif_spec.set_admin_status(intf::IF_STATUS_UP);
    ret = lif_create(lif_spec, &rsp);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("lif creation failed for proxy service" );
        return ret;
    }

    // program qstate
    qstate_params.type[proxy->qtype].entries = proxy->qstate_entries;
    qstate_params.type[proxy->qtype].size = proxy->qstate_size;
    
    int32_t rs = g_lif_manager->InitLIFQState(proxy->lif_id, &qstate_params);
    if(rs != 0) {
        HAL_TRACE_ERR("Failed to program lif qstate params: 0x{0:x}", rs);
        return HAL_RET_HW_PROG_ERR;
    }

    // Get the base address based on QID of '0'
    proxy->base_addr = g_lif_manager->GetLIFQStateAddr(proxy->lif_id, proxy->qtype, 0);

    return HAL_RET_OK;
}


//------------------------------------------------------------------------------
// process a Proxy enable request
// TODO: if Proxy exists, treat this as modify (tenant id in the meta must
// match though)
//------------------------------------------------------------------------------

hal_ret_t 
proxy_init_default_params(proxy_t* proxy)
{
    if(NULL == proxy) 
    {
        return HAL_RET_INVALID_ARG;    
    }

    switch(proxy->type) {
        case types::PROXY_TYPE_TCP:
            proxy->lif_id = SERVICE_LIF_TCP_PROXY;
            proxy->qtype = PROXY_TCP_DEF_QTYPE; 
            proxy->qstate_size = PROXY_TCP_DEF_QSTATE_SZ;
            proxy->qstate_entries = PROXY_TCP_DEF_QSTATE_ENTRIES;
            break;

        case types::PROXY_TYPE_TLS:
            proxy->lif_id = SERVICE_LIF_TLS_PROXY;
            proxy->qtype = PROXY_TLS_DEF_QTYPE; 
            proxy->qstate_size = PROXY_TLS_DEF_QSTATE_SZ;
            proxy->qstate_entries = PROXY_TLS_DEF_QSTATE_ENTRIES;
            break;

        default:
            return HAL_RET_INVALID_ARG;
        
    }

    return HAL_RET_OK;
}

proxy_t* 
proxy_factory(types::ProxyType type) 
{
    hal_ret_t       ret = HAL_RET_OK;
    proxy_t         * proxy = NULL;

    // instantiate Proxy
    proxy = proxy_alloc_init();
    if (proxy == NULL) {
        HAL_TRACE_ERR("Failed to allocate proxy service");
        return NULL;
    }

    proxy->type = type;
    proxy->hal_handle = hal_alloc_handle();
    
    // Instantiate QID indexer 
    proxy->qid_idxr_ = new hal::utils::indexer(HAL_MAX_QID);
    HAL_ASSERT(NULL != proxy->qid_idxr_);
   
   // initialize default params
    proxy_init_default_params(proxy);
    
    // program LIF for this proxy
    proxy_program_lif(proxy);
    
    // add this Proxy to our db
    ret = add_proxy_to_db(proxy);
    HAL_ASSERT(ret == HAL_RET_OK);

    return proxy;
}

hal_ret_t
proxy_enable(ProxySpec& spec, ProxyResponse *rsp)
{
    hal_ret_t              ret = HAL_RET_OK;
    proxy_t                *proxy;

    // validate the request message
    ret = validate_proxy_enable(spec, rsp);
    if (ret != HAL_RET_OK) {
        // api_status already set, just return
        HAL_TRACE_ERR("Proxy validate failure, err : {}", ret);
        return ret;
    }

    // instantiate Proxy
    proxy = proxy_factory(spec.proxy_type());
    if (ret != HAL_RET_OK) {
        rsp->set_api_status(types::API_STATUS_OUT_OF_MEM);
        return HAL_RET_OOM;
    }

    // prepare the response
    rsp->set_api_status(types::API_STATUS_OK);
    rsp->mutable_proxy_status()->set_proxy_handle(proxy->hal_handle);
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// Initialize default proxy services
//------------------------------------------------------------------------------

hal_ret_t 
hal_proxy_svc_init(void)
{
    // Reserve Service LIFs
    if(g_lif_manager->LIFRangeAlloc(SERVICE_LIF_START, (SERVICE_LIF_END - SERVICE_LIF_START)) 
            <= 0) 
    {
        HAL_TRACE_ERR("Failed to reserve service LIF");
        return HAL_RET_NO_RESOURCE;
    }
    
    // TODO: This should be enabled by DOL. REMOVE
    hal_init_def_proxy_services();
    
    return HAL_RET_OK;
}

hal_ret_t
hal_init_def_proxy_services(void) 
{
    HAL_ASSERT(NULL != proxy_factory(types::PROXY_TYPE_TCP));
    HAL_TRACE_DEBUG("Initialized TCP Proxy Service...");
    HAL_ASSERT(NULL != proxy_factory(types::PROXY_TYPE_TLS));   
    HAL_TRACE_DEBUG("Initialized TLS Proxy Service...");
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// process a Proxy update request
//------------------------------------------------------------------------------
hal_ret_t
proxy_update (ProxySpec& spec, ProxyResponse *rsp)
{
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// process a Proxy get request
//------------------------------------------------------------------------------
hal_ret_t
proxy_get (ProxyGetRequest& req, ProxyGetResponse *rsp)
{
#if 0
    hal_ret_t              ret = HAL_RET_OK; 
    proxy_t                rproxy;
    proxy_t*               proxy;
    pd::pd_proxy_args_t    pd_proxy_args;

    if (!req.has_meta()) {
        rsp->set_api_status(types::API_STATUS_TENANT_ID_INVALID);
        return HAL_RET_INVALID_ARG;
    }

    if (!req.has_key_or_handle()) {
        rsp->set_api_status(types::API_STATUS_TCP_CB_ID_INVALID);
        return HAL_RET_INVALID_ARG;
    }
    auto kh = req.key_or_handle();

    if (kh.key_or_handle_case() == proxy::ProxyKeyHandle::kProxybId) {
        proxy = find_proxy_by_id(kh.proxy_id());
    } else if (kh.key_or_handle_case() == proxy::ProxyKeyHandle::kProxyHandle) {
        proxy = find_proxy_by_handle(kh.proxy_handle());
    } else {
        rsp->set_api_status(types::API_STATUS_INVALID_ARG);
        return HAL_RET_INVALID_ARG;
    }

    if (proxy == NULL) {
        rsp->set_api_status(types::API_STATUS_TCP_CB_NOT_FOUND);
        return HAL_RET_TCP_CB_NOT_FOUND;
    }
    
    proxy_init(&rproxy);
    rproxy.cb_id = proxy->cb_id;
    pd::pd_proxy_args_init(&pd_proxy_args);
    pd_proxy_args.proxy = &rproxy;
    
    ret = pd::pd_proxy_get(&pd_proxy_args);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD PROXY: Failed to get, err: ", ret);
        rsp->set_api_status(types::API_STATUS_TCP_CB_NOT_FOUND);
        return HAL_RET_HW_FAIL;
    }

    // fill config spec of this TCP CB 
    rsp->mutable_spec()->mutable_key_or_handle()->set_proxy_id(rproxy.cb_id);
    rsp->mutable_spec()->set_rcv_nxt(rproxy.rcv_nxt);

    // fill operational state of this TCP CB
    rsp->mutable_status()->set_proxy_handle(proxy->hal_handle);
#endif

    // fill stats of this TCP CB
    rsp->set_api_status(types::API_STATUS_OK);
    return HAL_RET_OK;
}

}    // namespace hal
