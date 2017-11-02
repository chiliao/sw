// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

#ifndef __NWSEC_HPP__
#define __NWSEC_HPP__

#include "nic/include/base.h"
#include "nic/include/hal_state.hpp"
#include "nic/utils/ht/ht.hpp"
#include "nic/gen/proto/hal/nwsec.pb.h"
#include "nic/include/pd.hpp"

using hal::utils::ht_ctxt_t;

using nwsec::SecurityProfileKeyHandle;
using nwsec::SecurityProfileSpec;
using nwsec::SecurityProfileRequestMsg;
using nwsec::SecurityProfileStatus;
using nwsec::SecurityProfileResponse;
using nwsec::SecurityProfileResponseMsg;
using nwsec::SecurityProfileDeleteRequest;
using nwsec::SecurityProfileDeleteRequestMsg;
using nwsec::SecurityProfileDeleteResponseMsg;
using nwsec::SecurityProfileGetRequest;
using nwsec::SecurityProfileGetRequestMsg;
using nwsec::SecurityProfileStats;
using nwsec::SecurityProfileGetResponse;
using nwsec::SecurityProfileGetResponseMsg;
using nwsec::SecurityGroupSpec;
using nwsec::SecurityGroupRequestMsg;
using nwsec::SecurityGroupStatus;
using nwsec::SecurityGroupResponse;
using nwsec::SecurityGroupGetResponse;
using nwsec::SecurityGroupKeyHandle;
using nwsec::SecurityGroupPolicySpec;
using nwsec::SecurityGroupPolicyRequestMsg;
using nwsec::SecurityGroupPolicyStatus;
using nwsec::SecurityGroupPolicyResponse;
using nwsec::SecurityGroupPolicyGetResponse;
using nwsec::SecurityGroupPolicyKeyHandle;
using nwsec::Service;
using types::IPProtocol;
using types::ICMPMsgType;
using nwsec::FirewallAction;
using nwsec::ALGName;


namespace hal {

typedef uint32_t nwsec_profile_id_t;

typedef struct nwsec_profile_s {
    hal_spinlock_t        slock;                  // lock to protect this structure
    nwsec_profile_id_t    profile_id;             // profile id
    uint32_t              cnxn_tracking_en:1;
    uint32_t              ipsg_en:1;
    uint32_t              tcp_rtt_estimate_en:1;
    uint32_t              session_idle_timeout;
    uint32_t              tcp_cnxn_setup_timeout;
    uint32_t              tcp_close_timeout;
    uint32_t              tcp_close_wait_timeout;

    uint32_t              ip_normalization_en:1;
    uint32_t              tcp_normalization_en:1;
    uint32_t              icmp_normalization_en:1;

    uint32_t              ip_ttl_change_detect_en:1;
    uint32_t              ip_rsvd_flags_action:2;
    uint32_t              ip_df_action:2;
    uint32_t              ip_options_action:2;
    uint32_t              ip_invalid_len_action:2;
    uint32_t              ip_normalize_ttl:8;

    uint32_t              icmp_redirect_msg_drop:1;
    uint32_t              icmp_deprecated_msgs_drop:1;
    uint32_t              icmp_invalid_code_action:2;

    uint32_t              tcp_non_syn_first_pkt_drop:1;
    uint32_t              tcp_syncookie_en:1;
    uint32_t              tcp_split_handshake_detect_en:1;
    uint32_t              tcp_split_handshake_drop:1;
    uint32_t              tcp_rsvd_flags_action:2;
    uint32_t              tcp_unexpected_mss_action:2;
    uint32_t              tcp_unexpected_win_scale_action:2;
    uint32_t              tcp_urg_ptr_not_set_action:2;
    uint32_t              tcp_urg_flag_not_set_action:2;
    uint32_t              tcp_urg_payload_missing_action:2;
    uint32_t              tcp_rst_with_data_action:2;
    uint32_t              tcp_data_len_gt_mss_action:2;
    uint32_t              tcp_data_len_gt_win_size_action:2;
    uint32_t              tcp_unexpected_ts_option_action:2;
    uint32_t              tcp_unexpected_echo_ts_action:2;
    uint32_t              tcp_ts_not_present_drop:1;
    uint32_t              tcp_invalid_flags_drop:1;
    uint32_t              tcp_nonsyn_noack_drop:1;

    hal_handle_t          hal_handle;             // HAL allocated handle

    // operational state of nwsec profile
    dllist_ctxt_t         tenant_list_head;

    // PD state
    void                  *pd;                    // all PD specific state
} __PACK__ nwsec_profile_t;

typedef struct nwsec_create_app_ctxt_s {
} __PACK__ nwsec_create_app_ctxt_t;

typedef struct nwsec_update_app_ctxt_s {
    bool        ipsg_changed;                       // ipsg changed
    bool        nwsec_changed;                      // Any field changed

    // valid for ipsg_changed
    uint32_t    ipsg_en:1;                          // new ipsg_en value

    // valid for any change
    SecurityProfileSpec *spec;
} __PACK__ nwsec_update_app_ctxt_t;

// max. number of security profiles supported  (TODO: we can take this from cfg file)
#define HAL_MAX_NWSEC_PROFILES                       256

static inline void 
nwsec_profile_lock(nwsec_profile_t *nwsec_profile, const char *fname, 
                   int lineno, const char *fxname)
{
    HAL_TRACE_DEBUG("{}:operlock:locking nwsec_profile:{} from {}:{}:{}", 
                    __FUNCTION__, nwsec_profile->profile_id,
                    fname, lineno, fxname);
    HAL_SPINLOCK_LOCK(&nwsec_profile->slock);
}

static inline void 
nwsec_profile_unlock(nwsec_profile_t *nwsec_profile, const char *fname, 
                     int lineno, const char *fxname)
{
    HAL_TRACE_DEBUG("{}:operlock:unlocking nwsec_profile:{} from {}:{}:{}", 
                    __FUNCTION__, nwsec_profile->profile_id,
                    fname, lineno, fxname);
    HAL_SPINLOCK_UNLOCK(&nwsec_profile->slock);
}

// allocate a security profile instance
static inline nwsec_profile_t *
nwsec_profile_alloc (void)
{
    nwsec_profile_t    *sec_prof;

    sec_prof = (nwsec_profile_t *)g_hal_state->nwsec_profile_slab()->alloc();
    if (sec_prof == NULL) {
        return NULL;
    }
    return sec_prof;
}

// initialize a security profile instance
static inline nwsec_profile_t *
nwsec_profile_init (nwsec_profile_t *sec_prof)
{
    if (!sec_prof) {
        return NULL;
    }
    memset(sec_prof, 0, sizeof(nwsec_profile_t));


    HAL_SPINLOCK_INIT(&sec_prof->slock, PTHREAD_PROCESS_PRIVATE);

    // initialize the operational state
    utils::dllist_reset(&sec_prof->tenant_list_head);

    return sec_prof;
}

// allocate and initialize a security profile instance
static inline nwsec_profile_t *
nwsec_profile_alloc_init (void)
{
    return nwsec_profile_init(nwsec_profile_alloc());
}

// free security profile instance
static inline hal_ret_t
nwsec_profile_free (nwsec_profile_t *sec_prof)
{
    HAL_SPINLOCK_DESTROY(&sec_prof->slock);
    g_hal_state->nwsec_profile_slab()->free(sec_prof);
    return HAL_RET_OK;
}

// find a security profile instance by its id
static inline nwsec_profile_t *
find_nwsec_profile_by_id (nwsec_profile_id_t profile_id)
{
    // return (nwsec_profile_t *)g_hal_state->nwsec_profile_id_ht()->lookup(&profile_id);
    hal_handle_id_ht_entry_t    *entry;
    nwsec_profile_t             *sec_prof;

    entry = (hal_handle_id_ht_entry_t *)g_hal_state->nwsec_profile_id_ht()->lookup(&profile_id);
    if (entry) {
        // check for object type
        HAL_ASSERT(hal_handle_get_from_handle_id(entry->handle_id)->obj_id() == 
                   HAL_OBJ_ID_SECURITY_PROFILE);

        sec_prof = (nwsec_profile_t *)hal_handle_get_obj(entry->handle_id);
        return sec_prof;
    }
    return NULL;
}

// find a security profile instance by its handle
static inline nwsec_profile_t *
find_nwsec_profile_by_handle (hal_handle_t handle)
{
    // return (nwsec_profile_t *)g_hal_state->nwsec_profile_hal_handle_ht()->lookup(&handle);
    // check for object type
    auto hal_handle = hal_handle_get_from_handle_id(handle);
    if (!hal_handle) {
        HAL_TRACE_DEBUG("{}:failed to find object with handle:{}",
                        __FUNCTION__, handle);
        return NULL;
    }
    if (hal_handle->obj_id() != HAL_OBJ_ID_SECURITY_PROFILE) {
        HAL_TRACE_DEBUG("{}:failed to find nwsec profile with handle:{}",
                        __FUNCTION__, handle);
        return NULL;
    }
    // HAL_ASSERT(hal_handle->obj_id() == 
    //           HAL_OBJ_ID_SECURITY_PROFILE);
    return (nwsec_profile_t *)hal_handle_get_obj(handle); 
}

extern void *nwsec_profile_id_get_key_func(void *entry);
extern uint32_t nwsec_profile_id_compute_hash_func(void *key, uint32_t ht_size);
extern bool nwsec_profile_id_compare_key_func(void *key1, void *key2);
nwsec_profile_t *nwsec_lookup_key_or_handle (const SecurityProfileKeyHandle& kh);




// TODO: Move nwsec policy to a separate file

// extern void *nwsec_profile_get_handle_key_func(void *entry);
// extern uint32_t nwsec_profile_compute_handle_hash_func(void *key, uint32_t ht_size);
// extern bool nwsec_profile_compare_handle_key_func(void *key1, void *key2);


// Data structure to maintain for security group:
// CFG_DB:
//    nwsec_policy_cfg_t - sg_id (key and stored into nwsec_policy_cfg_ht in cfg_db)
//            maintains linked list of ingres and egress rules.
//            rules are of data type - nwsec_policy_rules_t
//            each rule has list of services defined by - nwsec_policy_svc_t;
//
//
// OPER_DB: 
//     nwsec_policy_entry_t
//        We need to key in (src sg_id, dst sg_id)
//        Define a structure nwsec_policy_entry_t and this will point to the rules list
// To DO:
//      Introduce a indirection and make
//      nwsec_policy_rules_t sharable between cfg and oper db
//
//

typedef struct nwsec_policy_key_s {
    uint32_t            sg_id;
    uint32_t            peer_sg_id;
} __PACK__ nwsec_policy_key_t;
 
typedef struct nwsec_policy_svc_s {
    hal_spinlock_t      slock;              // Lock to protect this structure
    IPProtocol          ipproto;
    union {
        uint32_t        dst_port;
        ICMPMsgType     icmp_msg_type;
    }__PACK__;
    ALGName             alg;
    dllist_ctxt_t       lentry;
} __PACK__ nwsec_policy_svc_t;

typedef  struct nwsec_policy_rules_s {
    hal_spinlock_t         slock;
    dllist_ctxt_t          fw_svc_list_head;
    bool                   log;
    FirewallAction         action;
    dllist_ctxt_t          lentry;
} __PACK__ nwsec_policy_rules_t;


typedef struct nwsec_policy_cfg_s {
    hal_spinlock_t               slock;      // lock to protect this strucuture
    nwsec_policy_key_t           plcy_key;
    dllist_ctxt_t                rules_head; // List of rules - nwsec_policy_rules_cfg_t

    // operational state of security group policy
    hal_handle_t                 hal_handle;         // HAL allocated handle
    ht_ctxt_t                    ht_ctxt;
} __PACK__ nwsec_policy_cfg_t;

typedef struct nwsec_group_s {
    hal_spinlock_t     slock;
    uint32_t           sg_id;
    dllist_ctxt_t      ep_list_head;
    dllist_ctxt_t      nw_list_head;
    // Operational state of Security Group
    hal_handle_t       hal_handle;
    ht_ctxt_t          ht_ctxt;
}__PACK__ nwsec_group_t;


// Empty context for now
typedef struct nwsec_policy_cfg_create_app_ctxt_s {
} __PACK__ nwsec_policy_cfg_create_app_ctxt_t;

typedef struct nwsec_group_create_app_ctxt_s {
} __PACK__ nwsec_group_create_app_ctxt_t;

//-----------------------------------------------------------------------------
//  APIs related to nwsec_policy_svc_t
//_____________________________________________________________________________
static inline nwsec_policy_svc_t *
nwsec_policy_svc_alloc(void)
{
    nwsec_policy_svc_t *nwsec_plcy_svc;
    nwsec_plcy_svc = (nwsec_policy_svc_t *)
                        g_hal_state->nwsec_policy_svc_slab()->alloc();
    if (nwsec_plcy_svc == NULL) {
        return NULL;
    }
    return nwsec_plcy_svc;
}

static inline nwsec_policy_svc_t *
nwsec_policy_svc_init(nwsec_policy_svc_t *nwsec_plcy_svcs)
{
    if (!nwsec_plcy_svcs) {
        return NULL;
    }
    HAL_SPINLOCK_INIT(&nwsec_plcy_svcs->slock, PTHREAD_PROCESS_PRIVATE);
    dllist_reset(&nwsec_plcy_svcs->lentry);
    return nwsec_plcy_svcs;
}

static inline nwsec_policy_svc_t *
nwsec_policy_svc_alloc_and_init(void)
{
    return nwsec_policy_svc_init(nwsec_policy_svc_alloc());
}

static inline hal_ret_t
nwsec_policy_svc_free(nwsec_policy_svc_t *nwsec_plcy_rules)
{
    HAL_SPINLOCK_DESTROY(&nwsec_plcy_rules->slock);
    g_hal_state->nwsec_policy_svc_slab()->free(nwsec_plcy_rules);
    return HAL_RET_OK;
}

//-----------------------------------------------------------------------------
//  APIs related to nwsec_policy_rules_t
//_____________________________________________________________________________
static inline nwsec_policy_rules_t *
nwsec_policy_rules_alloc(void)
{
    nwsec_policy_rules_t *nwsec_plcy_rules;
    nwsec_plcy_rules = (nwsec_policy_rules_t *)
                        g_hal_state->nwsec_policy_rules_slab()->alloc();
    if (nwsec_plcy_rules == NULL) {
        return NULL;
    }
    return nwsec_plcy_rules;
}

static inline nwsec_policy_rules_t *
nwsec_policy_rules_init(nwsec_policy_rules_t *nwsec_plcy_rules)
{
    if (!nwsec_plcy_rules) {
        return NULL;
    }
    dllist_reset(&nwsec_plcy_rules->fw_svc_list_head);
    dllist_reset(&nwsec_plcy_rules->lentry);
    HAL_SPINLOCK_INIT(&nwsec_plcy_rules->slock, PTHREAD_PROCESS_PRIVATE);

    return nwsec_plcy_rules;
}

static inline nwsec_policy_rules_t *
nwsec_policy_rules_alloc_and_init(void)
{
    return nwsec_policy_rules_init(nwsec_policy_rules_alloc());
} 

static inline hal_ret_t 
nwsec_policy_rules_free(nwsec_policy_rules_t *nwsec_plcy_rules) 
{
    HAL_SPINLOCK_DESTROY(&nwsec_plcy_rules->slock);
    g_hal_state->nwsec_policy_rules_slab()->free(nwsec_plcy_rules);
    return HAL_RET_OK;
}

//-----------------------------------------------------------------------------
//  APIs related to nwsec_policy_cfg_t
//_____________________________________________________________________________

extern void *
nwsec_policy_cfg_get_key_func (void *entry);

extern uint32_t
nwsec_policy_cfg_compute_hash_func (void *key, uint32_t ht_size);

extern bool
nwsec_policy_cfg_compare_key_func (void *key1, void *key2);

// max. number of SGs supported  (TODO: we can take this from cfg file)
#define HAL_MAX_NW_SEC_POLICY_CFG  8192

typedef uint32_t nwsec_policy_cfg_id_t;
// nwsec_policy_cfg related functions

static inline nwsec_policy_cfg_t *
nwsec_policy_cfg_alloc(void)
{
    nwsec_policy_cfg_t     *nwsec_plcy_cfg;

    nwsec_plcy_cfg = (nwsec_policy_cfg_t *)
                              g_hal_state->nwsec_policy_cfg_slab()->alloc();
    if (nwsec_plcy_cfg == NULL) {
        return NULL;
    }
    return nwsec_plcy_cfg;
}

// initialize a security policy cfg instance
static inline nwsec_policy_cfg_t *
nwsec_policy_cfg_init (nwsec_policy_cfg_t *nwsec_plcy_cfg)
{
    if (!nwsec_plcy_cfg) {
        return NULL;
    }
    // initialize the operational state

    // initialize the meta information
    nwsec_plcy_cfg->ht_ctxt.reset();
    nwsec_plcy_cfg->hal_handle = HAL_HANDLE_INVALID;
    dllist_reset(&nwsec_plcy_cfg->rules_head);
    return nwsec_plcy_cfg;
}

// allocate and initialize a security policy cfg instance
static inline nwsec_policy_cfg_t *
nwsec_policy_cfg_alloc_init (void)
{
    return nwsec_policy_cfg_init(nwsec_policy_cfg_alloc());
}

// free security policy cfg instance
static inline hal_ret_t
nwsec_policy_cfg_free(nwsec_policy_cfg_t *nwsec_plcy_cfg)
{
    g_hal_state->nwsec_policy_cfg_slab()->free(nwsec_plcy_cfg);
    return HAL_RET_OK;
}

// insert a security policy to meta data strucutre
static inline hal_ret_t
add_nwsec_policy_cfg_to_db (nwsec_policy_cfg_t *nwsec_plcy_cfg)
{
    g_hal_state->nwsec_policy_cfg_ht()->insert(nwsec_plcy_cfg,
                                               &nwsec_plcy_cfg->ht_ctxt);
    return HAL_RET_OK;
}


// find a security policy by sg
static inline nwsec_policy_cfg_t *
nwsec_policy_cfg_lookup_by_key(nwsec_policy_key_t plcy_key)
{
    return (nwsec_policy_cfg_t *)
         g_hal_state->nwsec_policy_cfg_ht()->lookup(&plcy_key);
}

/** 
  * Security Group related config
  *
**/
extern void *
nwsec_group_get_key_func (void *entry);

extern uint32_t
nwsec_group_compute_hash_func (void *key, uint32_t ht_size);

extern bool
nwsec_group_compare_key_func (void *key1, void *key2);

// max. number of SGs supported  (TODO: we can take this from cfg file)
#define HAL_MAX_NW_SEC_GROUP_CFG  8192

typedef uint32_t nwsec_group_id_t;
// nwsec_policy_cfg related functions

static inline nwsec_group_t *
nwsec_group_alloc(void)
{
    nwsec_group_t     *nwsec_grp;

    nwsec_grp = (nwsec_group_t *)
                              g_hal_state->nwsec_group_slab()->alloc();
    if (nwsec_grp == NULL) {
        return NULL;
    }
    return nwsec_grp;
}

// initialize a security policy cfg instance
static inline nwsec_group_t *
nwsec_group_init (nwsec_group_t *nwsec_grp)
{
    if (!nwsec_grp) {
        return NULL;
    }
    // initialize the operational state

    // initialize the meta information
    nwsec_grp->ht_ctxt.reset();
    nwsec_grp->hal_handle = HAL_HANDLE_INVALID;
    dllist_reset(&nwsec_grp->nw_list_head);
    dllist_reset(&nwsec_grp->ep_list_head);
    HAL_SPINLOCK_INIT(&nwsec_grp->slock, PTHREAD_PROCESS_PRIVATE);
    return nwsec_grp;
}

// allocate and initialize a security policy cfg instance
static inline nwsec_group_t *
nwsec_group_alloc_init (void)
{
    return nwsec_group_init(nwsec_group_alloc());
}

// free security policy cfg instance
static inline hal_ret_t
nwsec_group_free(nwsec_group_t *nwsec_grp)
{
    HAL_SPINLOCK_DESTROY(&nwsec_grp->slock);
    g_hal_state->nwsec_group_slab()->free(nwsec_grp);
    return HAL_RET_OK;
}

// insert a security policy to meta data strucutre
static inline hal_ret_t
add_nwsec_group_to_db (nwsec_group_t *nwsec_grp)
{
    hal_ret_t                       ret;
    hal_handle_id_ht_entry_t        *entry;

    HAL_TRACE_DEBUG("Adding to security group hash table sg_id {}", nwsec_grp->sg_id);
    entry = (hal_handle_id_ht_entry_t *)g_hal_state->
            hal_handle_id_ht_entry_slab()->alloc();
    if (entry == NULL) {
        return HAL_RET_OOM;
    }
    // add mapping from security group id to its handle
    entry->handle_id   = nwsec_grp->hal_handle;

    ret = g_hal_state->nwsec_group_ht()->insert_with_key(&nwsec_grp->sg_id, entry,
                                                         &entry->ht_ctxt);
    if (ret != HAL_RET_OK) {
        HAL_TRACE_ERR("Failed to add security group {} to handle mapping, "
                      "err : {}", nwsec_grp->sg_id, ret);
        g_hal_state->hal_handle_id_ht_entry_slab()->free(entry);
    }
    //nwsec_grp->hal_handle = handle_id;
    return HAL_RET_OK;
}

// find a security group by key
static inline nwsec_group_t *
nwsec_group_lookup_by_key(uint32_t sg_id)
{
    hal_handle_id_ht_entry_t     *entry;
    nwsec_group_t                *nwsec_group;

    HAL_TRACE_DEBUG("Lookup sg_id {}", sg_id);
    entry = (hal_handle_id_ht_entry_t *)g_hal_state->nwsec_group_ht()->lookup(&sg_id);
    if (entry) {
        HAL_ASSERT(hal_handle_get_from_handle_id(entry->handle_id)->obj_id() ==
                   HAL_OBJ_ID_SECURITY_GROUP);
        nwsec_group = (nwsec_group_t *) hal_handle_get_obj(entry->handle_id);
        return nwsec_group;
    } else {
        HAL_TRACE_DEBUG("security group id:{} entry not found", sg_id);
    }
     
    return NULL;
}

// find a security group by handle
static inline nwsec_group_t *
nwsec_group_lookup_by_handle (hal_handle_t handle)
{
    auto hal_handle = hal_handle_get_from_handle_id(handle);
    if (!hal_handle) {
        HAL_TRACE_DEBUG("failed to find hal object with handle:{}",
                        handle);
        return NULL;
    }
    if (hal_handle->obj_id() != HAL_OBJ_ID_SECURITY_GROUP) {
        HAL_TRACE_DEBUG("failed to find security_group  with handle:{}",
                        handle);
        return NULL;
    }
    return (nwsec_group_t *)hal_handle->get_obj();
}

//APIs to be consumed by other features

dllist_ctxt_t *
get_ep_list_for_security_group(uint32_t sg_id);
dllist_ctxt_t *
get_nw_list_for_security_group(uint32_t sg_id);
hal_ret_t
add_ep_to_security_group(uint32_t sg_id, hal_handle_t ep_handle);

// Forward declarations
hal_ret_t security_profile_create(nwsec::SecurityProfileSpec& spec,
                                  nwsec::SecurityProfileResponse *rsp);

hal_ret_t security_profile_update(nwsec::SecurityProfileSpec& spec,
                                  nwsec::SecurityProfileResponse *rsp);

hal_ret_t security_profile_delete(nwsec::SecurityProfileDeleteRequest& req,
                                  nwsec::SecurityProfileDeleteResponseMsg *rsp);

hal_ret_t security_profile_get(nwsec::SecurityProfileGetRequest& req,
                               nwsec::SecurityProfileGetResponse *rsp);

hal_ret_t security_group_create(nwsec::SecurityGroupSpec& req,
                               nwsec::SecurityGroupResponse *rsp);

hal_ret_t security_group_update(nwsec::SecurityGroupSpec& req,
                               nwsec::SecurityGroupResponse *rsp);

hal_ret_t security_group_delete(nwsec::SecurityGroupSpec& req,
                               nwsec::SecurityGroupResponse *rsp);

hal_ret_t security_group_get(nwsec::SecurityGroupGetRequest& req,
                               nwsec::SecurityGroupGetResponse *rsp);

hal_ret_t security_group_policy_create(nwsec::SecurityGroupPolicySpec& req,
                                       nwsec::SecurityGroupPolicyResponse *rsp);

hal_ret_t security_group_policy_update(nwsec::SecurityGroupPolicySpec& req,
                                       nwsec::SecurityGroupPolicyResponse *rsp);

hal_ret_t security_group_policy_delete(nwsec::SecurityGroupPolicySpec& req,
                                       nwsec::SecurityGroupPolicyResponse *rsp);

hal_ret_t security_group_policy_get(nwsec::SecurityGroupPolicyGetRequest& req,
                                    nwsec::SecurityGroupPolicyGetResponse *rsp);
}    // namespace hal

#endif    // __NWSEC_HPP__

