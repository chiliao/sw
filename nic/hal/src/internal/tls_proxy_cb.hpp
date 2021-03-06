//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

#ifndef __TLSCB_HPP__
#define __TLSCB_HPP__

#include "nic/include/base.hpp"
#include "nic/sdk/include/sdk/encap.hpp"
#include "lib/list/list.hpp"
#include "lib/ht/ht.hpp"
#include "gen/proto/internal.pb.h"
#include "nic/include/pd.hpp"
#include "nic/hal/iris/include/hal_state.hpp"

using sdk::lib::ht_ctxt_t;
using sdk::lib::dllist_ctxt_t;

using internal::TlsCbSpec;
using internal::TlsCbStatus;
using internal::TlsCbResponse;
using internal::TlsCbKeyHandle;
using internal::TlsCbRequestMsg;
using internal::TlsCbResponseMsg;
using internal::TlsCbDeleteRequestMsg;
using internal::TlsCbDeleteResponseMsg;
using internal::TlsCbGetRequest;
using internal::TlsCbGetRequestMsg;
using internal::TlsCbGetResponse;
using internal::TlsCbGetResponseMsg;

namespace hal {

typedef struct tlscb_s {
    sdk_spinlock_t        slock;                   // lock to protect this structure
    tlscb_id_t            cb_id;                   // TLS CB id
    uint32_t              command;
    uint64_t              sesq_base;
    uint64_t              serq_base;
    uint16_t              serq_pi;
    uint16_t              serq_ci;
    uint16_t              bsq_pi;
    uint16_t              bsq_ci;
    uint32_t              crypto_key_idx;
    uint32_t              debug_dol;
    uint64_t              tnmdpr_alloc;
    uint64_t              rnmdpr_free;
    uint64_t              enc_requests;
    uint64_t              enc_completions;
    uint64_t              enc_failures;
    uint64_t              dec_requests;
    uint64_t              dec_completions;
    uint64_t              dec_failures;
    uint64_t              mac_requests;
    uint64_t              mac_completions;
    uint64_t              mac_failures;
    uint32_t              salt;
    uint64_t              explicit_iv;
    uint32_t              pre_debug_stage0_7_thread;
    uint32_t              post_debug_stage0_7_thread;
    uint8_t               is_decrypt_flow;
    uint16_t              other_fid;
    types::AppRedirType   l7_proxy_type;
    uint32_t              crypto_hmac_key_idx;
    uint8_t               cpu_id;

    // operational state of TLS Proxy CB
    hal_handle_t          hal_handle;              // HAL allocated handle

    // PD state
    void                  *pd;                     // all PD specific state

    ht_ctxt_t             ht_ctxt;                 // id based hash table ctxt
    ht_ctxt_t             hal_handle_ht_ctxt;      // hal handle based hash table ctxt
} __PACK__ tlscb_t;

// max. number of TLS CBs supported  (TODO: we can take this from cfg file)
#define HAL_MAX_TLSCB                           2048

// allocate a tlscbment instance
static inline tlscb_t *
tlscb_alloc (void)
{
    tlscb_t    *tlscb;

    tlscb = (tlscb_t *)g_hal_state->tlscb_slab()->alloc();
    if (tlscb == NULL) {
        return NULL;
    }
    return tlscb;
}

// initialize a tlscbment instance
static inline tlscb_t *
tlscb_init (tlscb_t *tlscb)
{
    if (!tlscb) {
        return NULL;
    }
    SDK_SPINLOCK_INIT(&tlscb->slock, PTHREAD_PROCESS_PRIVATE);

    // initialize the operational state
    tlscb->pd = NULL;

    // initialize meta information
    tlscb->ht_ctxt.reset();
    tlscb->hal_handle_ht_ctxt.reset();

    return tlscb;
}

// allocate and initialize a TLSCB instance
static inline tlscb_t *
tlscb_alloc_init (void)
{
    return tlscb_init(tlscb_alloc());
}

static inline hal_ret_t
tlscb_free (tlscb_t *tlscb)
{
    SDK_SPINLOCK_DESTROY(&tlscb->slock);
    hal::delay_delete_to_slab(HAL_SLAB_TLSCB, tlscb);
    return HAL_RET_OK;
}

static inline tlscb_t *
find_tlscb_by_id (tlscb_id_t tlscb_id)
{
    return (tlscb_t *)g_hal_state->tlscb_id_ht()->lookup(&tlscb_id);
}

extern void *tlscb_get_key_func(void *entry);
extern uint32_t tlscb_key_size(void);

extern void *tlscb_get_handle_key_func(void *entry);
extern uint32_t tlscb_handle_key_size(void);

}    // namespace hal

#endif    // __TLSCB_HPP__

