//
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file implements upgrade states handling
///
//----------------------------------------------------------------------------

#include "nic/sdk/include/sdk/mem.hpp"
#include "nic/apollo/api/upgrade_state.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/core/mem.hpp"

namespace upg {

upg_state *upg_state::upg_state_ = NULL;

#define PDS_UPG_SHM_NAME        "pds_upgrade"
#define PDS_UPG_SHM_PSTATE_NAME "pds_upgrade_pstate"
// TODO: below size depends on the size of the config hw states to be saved
// and other nicmgr/linkmgr states etc. need to calculate for the maximum
// and adjust its size.
#define PDS_UPG_SHM_SIZE (100 << 10)  // 100KB

sdk_ret_t
upg_state::init_(bool shm_create) {
    sdk::lib::shm_mode_e mode = shm_create ? sdk::lib::SHM_CREATE_ONLY : sdk::lib::SHM_OPEN_ONLY;
    const char *op = shm_create ? "create" : "open";

    try {
        // if create, delete and re-create as previous size and current size may be different
        if (shm_create) {
            shmmgr::remove(PDS_UPG_SHM_NAME);
        }
        shm_mmgr_ = shmmgr::factory(PDS_UPG_SHM_NAME, PDS_UPG_SHM_SIZE, mode, NULL);
        if (shm_mmgr_ == NULL) {
            PDS_TRACE_ERR("Upgrade shared mem %s failed", op);
            return SDK_RET_ERR;
        }
    } catch (...) {
        PDS_TRACE_ERR("Upgrade shared mem %s failed", op);
        return SDK_RET_ERR;
    }

    pstate_ = (upg_pstate_t *)shm_mmgr_->segment_alloc(PDS_UPG_SHM_PSTATE_NAME,
                                                      sizeof(upg_pstate_t), shm_create);
    if (!pstate_) {
        PDS_TRACE_ERR("Upgrade pstate %s failed", op);
        return SDK_RET_ERR;
    }

    PDS_TRACE_DEBUG("Upgrade shared mem %s done", op);
    return SDK_RET_OK;
}

void
upg_state::destroy(upg_state *state) {
    SDK_FREE(api::PDS_MEM_ALLOC_UPG, state);
    shmmgr::remove(PDS_UPG_SHM_NAME);
}

upg_state *
upg_state::factory(bool shm_create) {
    sdk_ret_t ret;
    void *mem;

    mem = SDK_CALLOC(api::PDS_MEM_ALLOC_UPG, sizeof(upg_state));
    if (!mem) {
        PDS_TRACE_ERR("Upgrade state alloc failed");
        return NULL;
    }
    upg_state_ = new (mem) upg_state();
    ret = upg_state_->init_(shm_create);
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Upgrade state init failed");
        goto err_exit;
    }
    return upg_state_;

err_exit:

    SDK_FREE(api::PDS_MEM_ALLOC_UPG, upg_state_);
    upg_state_ = NULL;
    return NULL;
}

upg_state *
upg_state::get_instance(void) {
    SDK_ASSERT(upg_state_);
    return upg_state_;
}

}    // namespace upg