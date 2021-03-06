//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

#include "nic/include/base.hpp"
#include "nic/hal/hal.hpp"
#include "nic/sdk/include/sdk/lock.hpp"
#include "nic/hal/iris/include/hal_state.hpp"
#include "nic/hal/src/internal/wring.hpp"
#include "nic/hal/src/internal/internal.hpp"
#include "nic/hal/plugins/cfg/nw/vrf.hpp"
#include "nic/include/pd_api.hpp"

using std::string;
namespace hal {

void *
wring_get_key_func (void *entry)
{
    SDK_ASSERT(entry != NULL);
    return (void *)&(((wring_t *)entry)->wring_id);
}

uint32_t
wring_key_size ()
{
    return sizeof(wring_id_t);
}

void *
wring_get_handle_key_func (void *entry)
{
    SDK_ASSERT(entry != NULL);
    return (void *)&(((wring_t *)entry)->hal_handle);
}

uint32_t
wring_handle_key_size ()
{
    return sizeof(hal_handle_t);
}

//------------------------------------------------------------------------------
// validate an incoming WRING create request
// TODO:
// 1. check if WRING exists already
//------------------------------------------------------------------------------
static hal_ret_t
validate_wring_create (WRingSpec& spec, WRingResponse *rsp)
{
    // must have key-handle set
    if (!spec.has_key_or_handle()) {
        rsp->set_api_status(types::API_STATUS_WRING_ID_INVALID);
        return HAL_RET_INVALID_ARG;
    }

    // must have key in the key-handle
    if (spec.key_or_handle().key_or_handle_case() !=
            WRingKeyHandle::kWringId) {
        rsp->set_api_status(types::API_STATUS_WRING_ID_INVALID);
        return HAL_RET_INVALID_ARG;
    }
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// insert this WRing in all meta data structures
//------------------------------------------------------------------------------
static inline hal_ret_t
add_wring_to_db (wring_t *wring)
{
    g_hal_state->wring_id_ht()->insert(wring, &wring->ht_ctxt);
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// process a Wring create request
// match though)
//------------------------------------------------------------------------------
hal_ret_t
wring_create (WRingSpec& spec, WRingResponse *rsp)
{
    hal_ret_t              ret = HAL_RET_OK;
    wring_t                *wring;
    pd::pd_wring_create_args_t  pd_wring_args;
    pd::pd_func_args_t          pd_func_args = {0};

    // validate the request message
    ret = validate_wring_create(spec, rsp);
    if (ret != HAL_RET_OK) {
        // api_status already set, just return
        HAL_TRACE_ERR("PD Wringvalidate failure, err : {}", ret);
        return ret;
    }

    // instantiate WRing
    wring = wring_alloc_init();
    if (wring == NULL) {
        rsp->set_api_status(types::API_STATUS_OUT_OF_MEM);
        return HAL_RET_OOM;
    }

    wring->wring_id = spec.key_or_handle().wring_id();
    wring->hal_handle = hal_alloc_handle();

    // allocate all PD resources and finish programming
    pd::pd_wring_create_args_init(&pd_wring_args);
    pd_wring_args.wring = wring;
    pd_func_args.pd_wring_create = &pd_wring_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_WRING_CREATE, &pd_func_args);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD Wring create failure, err : {}", ret);
        rsp->set_api_status(types::API_STATUS_HW_PROG_ERR);
        goto cleanup;
    }

    // add this L2 segment to our db
    ret = add_wring_to_db(wring);
    SDK_ASSERT(ret == HAL_RET_OK);

    // prepare the response
    rsp->set_api_status(types::API_STATUS_OK);
    rsp->mutable_wring_status()->set_wring_handle(wring->hal_handle);
    return HAL_RET_OK;

cleanup:

    wring_free(wring);
    return ret;
}

//------------------------------------------------------------------------------
// process a WRing update request
//------------------------------------------------------------------------------
hal_ret_t
wring_update (WRingSpec& spec, WRingResponse *rsp)
{
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// process a WRing get request
//------------------------------------------------------------------------------
hal_ret_t
wring_get_entries(WRingGetEntriesRequest& req, WRingGetEntriesResponseMsg *rsp1)
{
    hal_ret_t               ret = HAL_RET_OK;
    wring_t                 wring;
    pd::pd_wring_get_entry_args_t     pd_wring_args;
    pd::pd_func_args_t          pd_func_args = {0};
    WRingGetEntriesResponse *rsp = rsp1->add_response();

    if(req.type() <= types::WRING_TYPE_NONE) {
        HAL_TRACE_ERR("Invalid wring type");
        rsp->set_api_status(types::API_STATUS_WRING_TYPE_INVALID);
        return HAL_RET_INVALID_ARG;
    }

    wring_init(&wring);
    wring.wring_type = req.type();
    wring.wring_id = req.key_or_handle().wring_id();
    wring.slot_index = req.index();

    pd::pd_wring_get_entry_args_init(&pd_wring_args);
    pd_wring_args.wring = &wring;

    pd_func_args.pd_wring_get_entry = &pd_wring_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_WRING_GET_ENTRY, &pd_func_args);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD Wring: Failed to get, err: {}", ret);
        rsp->set_api_status(types::API_STATUS_NOT_FOUND);
        return HAL_RET_HW_FAIL;
    }

    // fill config spec of this WRING
    rsp->mutable_spec()->set_type(wring.wring_type);
    rsp->mutable_spec()->mutable_key_or_handle()->set_wring_id(wring.wring_id);
    rsp->set_index(wring.slot_index);
    switch (wring.wring_type) {
        case types::WRING_TYPE_BRQ:
            rsp->clear_value();
            rsp->mutable_barco_gcm_desc()->set_ilist_addr(wring.slot_info.gcm_desc.ilist_addr);
            rsp->mutable_barco_gcm_desc()->set_olist_addr(wring.slot_info.gcm_desc.olist_addr);
            rsp->mutable_barco_gcm_desc()->set_command(wring.slot_info.gcm_desc.command);
            rsp->mutable_barco_gcm_desc()->set_command(wring.slot_info.gcm_desc.command);
            rsp->mutable_barco_gcm_desc()->set_key_desc_index(wring.slot_info.gcm_desc.key_desc_index);
            rsp->mutable_barco_gcm_desc()->set_iv_addr(wring.slot_info.gcm_desc.iv_addr);
            rsp->mutable_barco_gcm_desc()->set_status_addr(wring.slot_info.gcm_desc.status_addr);
            rsp->mutable_barco_gcm_desc()->set_doorbell_addr(wring.slot_info.gcm_desc.doorbell_addr);
            rsp->mutable_barco_gcm_desc()->set_doorbell_data(wring.slot_info.gcm_desc.doorbell_data);
            rsp->mutable_barco_gcm_desc()->set_salt(wring.slot_info.gcm_desc.salt);
            rsp->mutable_barco_gcm_desc()->set_explicit_iv(wring.slot_info.gcm_desc.explicit_iv);
            rsp->mutable_barco_gcm_desc()->set_header_size(wring.slot_info.gcm_desc.header_size);
            rsp->mutable_barco_gcm_desc()->set_barco_status(wring.slot_info.gcm_desc.barco_status);
            break;
        default:
            rsp->set_value(wring.slot_value);
            break;
    }

    HAL_TRACE_DEBUG("Ring slot_index: {}, type: {}", wring.slot_index, wring.wring_type);
    // fill operational state of this WRING
    //rsp->mutable_status()->set_wring_handle(wring->hal_handle);

    // fill stats of this WRING
    rsp->set_api_status(types::API_STATUS_OK);
    return HAL_RET_OK;
}


//------------------------------------------------------------------------------
// process a WRing get Meta request
//------------------------------------------------------------------------------
hal_ret_t
wring_get_meta(WRingSpec& spec, WRingGetMetaResponseMsg *rsp1)
{
    hal_ret_t               ret = HAL_RET_OK;
    wring_t                 wring;
    pd::pd_wring_get_meta_args_t     pd_wring_args;
    pd::pd_func_args_t          pd_func_args = {0};
    WRingGetMetaResponse *rsp = rsp1->add_response();


    if(spec.type() <= types::WRING_TYPE_NONE) {
        HAL_TRACE_ERR("Invalid wring type");
        rsp->set_api_status(types::API_STATUS_WRING_TYPE_INVALID);
        return HAL_RET_INVALID_ARG;
    }

    wring_init(&wring);
    wring.wring_type = spec.type();
    wring.wring_id = spec.key_or_handle().wring_id();

    pd::pd_wring_get_meta_args_init(&pd_wring_args);
    pd_wring_args.wring = &wring;

    pd_func_args.pd_wring_get_meta = &pd_wring_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_WRING_GET_META, &pd_func_args);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD Wring: Failed to get, err: {}", ret);
        rsp->set_api_status(types::API_STATUS_NOT_FOUND);
        return HAL_RET_HW_FAIL;
    }

    // fill config spec of this WRING
    rsp->mutable_spec()->set_type(wring.wring_type);
    rsp->mutable_spec()->set_pi(wring.pi);
    rsp->mutable_spec()->set_ci(wring.ci);

    HAL_TRACE_DEBUG("Ring pi: {} ci: {}", wring.pi, wring.ci);

    // fill stats of this WRING
    rsp->set_api_status(types::API_STATUS_OK);
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// process a WRing set Meta request
//------------------------------------------------------------------------------
hal_ret_t
wring_set_meta (WRingSpec& spec, WRingSetMetaResponse *rsp)
{
    hal_ret_t               ret = HAL_RET_OK;
    wring_t                 wring;
    pd::pd_wring_set_meta_args_t     pd_wring_args;
    pd::pd_func_args_t          pd_func_args = {0};

    if(spec.type() <= types::WRING_TYPE_NONE) {
        HAL_TRACE_ERR("Invalid wring type");
        rsp->set_api_status(types::API_STATUS_WRING_TYPE_INVALID);
        return HAL_RET_INVALID_ARG;
    }

    wring_init(&wring);
    wring.wring_type = spec.type();
    wring.pi = spec.pi();
    wring.ci = spec.ci();

    pd::pd_wring_set_meta_args_init(&pd_wring_args);
    pd_wring_args.wring = &wring;

    pd_func_args.pd_wring_set_meta = &pd_wring_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_WRING_SET_META, &pd_func_args);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD Wring: Failed to get, err: {}", ret);
        rsp->set_api_status(types::API_STATUS_NOT_FOUND);
        return HAL_RET_HW_FAIL;
    }

    // fill config spec of this WRING
    rsp->mutable_spec()->set_type(wring.wring_type);
    rsp->mutable_spec()->set_pi(wring.pi);
    rsp->mutable_spec()->set_ci(wring.ci);

    HAL_TRACE_DEBUG("Ring pi: {} ci: {}", wring.pi, wring.ci);

    // fill stats of this WRING
    rsp->set_api_status(types::API_STATUS_OK);
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// process a WRing get base address
//------------------------------------------------------------------------------
hal_ret_t
wring_get_phys_addr(types::WRingType wring_type, wring_id_t wring_id,
        uint64_t *phys_addr)
{
    hal_ret_t ret = HAL_RET_OK;
    wring_t wring;
    pd::pd_wring_get_meta_args_t pd_wring_args;
    pd::pd_func_args_t pd_func_args = {0};

    if (wring_type <= types::WRING_TYPE_NONE) {
        HAL_TRACE_ERR("Invalid wring type");
        return HAL_RET_INVALID_ARG;
    }

    wring_init(&wring);
    wring.wring_type = wring_type;
    wring.wring_id = wring_id;

    pd::pd_wring_get_meta_args_init(&pd_wring_args);
    pd_wring_args.wring = &wring;

    pd_func_args.pd_wring_get_meta = &pd_wring_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_WRING_GET_BASE_ADDR, &pd_func_args);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD Wring: Failed to get, err: {}", ret);
        return HAL_RET_HW_FAIL;
    }

    HAL_TRACE_DEBUG("Ring base_addr {}", wring.phys_base_addr);
    *phys_addr = wring.phys_base_addr;

    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// Get ring meta info
//------------------------------------------------------------------------------
hal_ret_t
wring_get_meta(types::WRingType wring_type, wring_id_t wring_id,
        wring_t *wring)
{
    hal_ret_t               ret = HAL_RET_OK;
    pd::pd_wring_get_meta_args_t     pd_wring_args;
    pd::pd_func_args_t          pd_func_args = {0};

    wring_init(wring);
    wring->wring_type = wring_type;
    wring->wring_id = wring_id;

    pd::pd_wring_get_meta_args_init(&pd_wring_args);
    pd_wring_args.wring = wring;

    pd_func_args.pd_wring_get_meta = &pd_wring_args;
    ret = pd::hal_pd_call(pd::PD_FUNC_ID_WRING_GET_META, &pd_func_args);
    if(ret != HAL_RET_OK) {
        HAL_TRACE_ERR("PD Wring: Failed to get, err: {}", ret);
        return HAL_RET_HW_FAIL;
    }

    HAL_TRACE_DEBUG("Ring {} phys_addr {:#x} num_entries {}", wring_type,
            wring->phys_base_addr, wring->num_entries);

    return HAL_RET_OK;
}

}    // namespace hal
