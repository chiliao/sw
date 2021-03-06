//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

#ifndef __WRING_HPP__
#define __WRING_HPP__

#include "nic/include/base.hpp"
#include "nic/sdk/include/sdk/encap.hpp"
#include "lib/list/list.hpp"
#include "lib/ht/ht.hpp"
#include "gen/proto/internal.pb.h"
#include "nic/include/pd.hpp"
#include "nic/hal/iris/include/hal_state.hpp"

using sdk::lib::ht_ctxt_t;
using sdk::lib::dllist_ctxt_t;

using internal::WRingSpec;
using internal::WRingStatus;
using internal::WRingResponse;
using internal::WRingKeyHandle;
using internal::WRingRequestMsg;
using internal::WRingResponseMsg;
using internal::WRingDeleteRequestMsg;
using internal::WRingDeleteResponseMsg;
using internal::WRingGetEntriesRequest;
using internal::WRingGetEntriesRequestMsg;
using internal::WRingGetEntriesResponse;
using internal::WRingGetEntriesResponseMsg;
using internal::WRingGetMetaResponseMsg;
using internal::WRingGetMetaResponse;
using internal::WRingSetMetaResponseMsg;
using internal::WRingSetMetaResponse;

namespace hal {

typedef struct barco_gcm_desc_s {
    uint64_t                ilist_addr;
    uint64_t                olist_addr;
    uint32_t                command;
    uint32_t                key_desc_index;
    uint64_t                iv_addr;
    uint64_t                status_addr;
    uint64_t                doorbell_addr;
    uint64_t                doorbell_data;
    uint32_t                salt;
    uint64_t                explicit_iv;
    uint32_t                barco_status;
    uint32_t                header_size;
} barco_gcm_desc_t;

typedef union wring_slot_info_u {
    barco_gcm_desc_t        gcm_desc;
} wring_slot_info_t;

typedef struct wring_s {
    sdk_spinlock_t        slock;                   // lock to protect this structure
    wring_id_t            wring_id;                // WRing id
    types::WRingType      wring_type;              // Wring Type
    uint64_t              slot_index;              // PI/CI for the request
    uint64_t              slot_value;              // Slot Value
    // operational state of WRing
    hal_handle_t          hal_handle;              // HAL allocated handle

    // PD state
    void                  *pd;                     // all PD specific state

    ht_ctxt_t             ht_ctxt;                 // id based hash table ctxt
    ht_ctxt_t             hal_handle_ht_ctxt;      // hal handle based hash table ctxt
    uint64_t              phys_base_addr;          // wring base address
    uint32_t              num_entries;
    uint32_t              obj_size;
    bool                  is_global;
    uint32_t              pi;
    uint32_t              ci;
    wring_slot_info_t     slot_info;
} __PACK__ wring_t;

// max. number of WRING supported  (TODO: we can take this from cfg file)
#define HAL_MAX_WRING                           2048

// allocate a wringment instance
static inline wring_t *
wring_alloc (void)
{
    wring_t    *wring;

    wring = (wring_t *)g_hal_state->wring_slab()->alloc();
    if (wring == NULL) {
        return NULL;
    }
    return wring;
}

// initialize a wringment instance
static inline wring_t *
wring_init (wring_t *wring)
{
    if (!wring) {
        return NULL;
    }
    SDK_SPINLOCK_INIT(&wring->slock, PTHREAD_PROCESS_PRIVATE);

    // initialize the operational state
    wring->pd = NULL;

    // initialize meta information
    wring->ht_ctxt.reset();
    wring->hal_handle_ht_ctxt.reset();

    return wring;
}

// allocate and initialize a WRING instance
static inline wring_t *
wring_alloc_init (void)
{
    return wring_init(wring_alloc());
}

static inline hal_ret_t
wring_free (wring_t *wring)
{
    SDK_SPINLOCK_DESTROY(&wring->slock);
    hal::delay_delete_to_slab(HAL_SLAB_WRING, wring);
    return HAL_RET_OK;
}

static inline wring_t *
find_wring_by_id (wring_id_t wring_id)
{
    return (wring_t *)g_hal_state->wring_id_ht()->lookup(&wring_id);
}

extern void *wring_get_key_func(void *entry);
extern uint32_t wring_key_size(void);

extern void *wring_get_handle_key_func(void *entry);
extern uint32_t wring_handle_key_size(void);
extern hal_ret_t wring_get_phys_addr(types::WRingType wring_type,
        wring_id_t wring_id, uint64_t *phys_addr);
extern hal_ret_t wring_get_meta(types::WRingType wring_type,
        wring_id_t wring_id, wring_t *wring);


}    // namespace hal

#endif    // __WRING_HPP__

