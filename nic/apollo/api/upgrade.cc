//
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------

#include "nic/apollo/api/include/pds_upgrade.hpp"
#include "nic/apollo/api/pds_state.hpp"
#include "nic/apollo/api/upgrade_state.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/core/core.hpp"
#include "nic/apollo/api/port.hpp"
#include "nic/apollo/api/upgrade.hpp"
#include "nic/apollo/nicmgr/nicmgr.hpp"

namespace api {

sdk_ret_t
upg_init (pds_init_params_t *params)
{
    sdk_ret_t ret;
    bool shm_create;
    upg_mode_t mode;

    mode = sdk::upg::upg_init_mode();

    PDS_TRACE_DEBUG("Setting bootup upgrade mode to %s",
                    mode == upg_mode_t::UPGRADE_MODE_NONE ? "NONE" :
                    mode == upg_mode_t::UPGRADE_MODE_GRACEFUL ? "Graceful" :
                    "Hitless");
    shm_create = sdk::platform::upgrade_mode_none(mode);

    // initialize upgrade state and call the upgade compatibitly checks
    // TODO: prefer this to be created by upgrade manager and every process
    // opens it and add segments to it.
    if ((g_upg_state = upg_state::factory(shm_create)) == NULL) {
        PDS_TRACE_ERR("Upgrade state creation failed");
        return SDK_RET_ERR;
    }
    g_upg_state->set_upg_init_mode(mode);
    // offset the memory regions based on the regions in use
    if (sdk::platform::upgrade_mode_hitless(mode)) {
        ret = g_pds_state.mempartition()->upgrade_hitless_offset_regions(
            api::g_pds_state.cfg_path().c_str(), false);
        if (ret != SDK_RET_OK) {
            PDS_TRACE_ERR("Upgrade hitless memory offset failed");
            return ret;
        }
    }
    if (shm_create) {
        g_upg_state->set_memory_profile(params->memory_profile);
    } else if (g_upg_state->memory_profile() != params->memory_profile) {
        PDS_TRACE_ERR("Upgrade scale profile mismatch");
        return SDK_RET_ERR;
    }

    // pds hal registers for upgrade events
    ret = upg_graceful_init(params);
    if (ret == SDK_RET_OK) {
        ret = upg_hitless_init(params);
    }
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Upgrade graceful/hitless init failed");
        return ret;
    }

    // nicmgr registers for upgrade events
    ret = nicmgr_upg_graceful_init();
    if (ret == SDK_RET_OK) {
        ret = nicmgr_upg_hitless_init();
    }
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Upgrade nicmgr graceful/hitless init failed");
        return ret;
    }

    return SDK_RET_OK;
}

}   // namespace api
