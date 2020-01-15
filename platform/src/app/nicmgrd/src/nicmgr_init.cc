/*
* Copyright (c) 2018, Pensando Systems Inc.
*/

#include <cstdio>
#include <iostream>
#include <iomanip>
#include <algorithm>
#include <cstdlib>
#include <unistd.h>
#include <sys/time.h>

#include "nic/sdk/lib/device/device.hpp"
#include "nic/sdk/lib/event_thread/event_thread.hpp"
#include "nic/sdk/lib/ipc/ipc.hpp"
#include "nic/sdk/lib/utils/port_utils.hpp"
#include "nic/sdk/platform/evutils/include/evutils.h"
#include "nic/hal/core/core.hpp"
#include "platform/src/lib/nicmgr/include/dev.hpp"
#include "platform/src/lib/nicmgr/include/logger.hpp"
#include "nic/hal/core/event_ipc.hpp"
#include "upgrade.hpp"

// for device::MICRO_SEG_ENABLE  : TODO - fix
#include <grpc++/grpc++.h>
#include "gen/proto/device.delphi.hpp"

using namespace std;

DeviceManager *devmgr;
const char* nicmgr_upgrade_state_file = "/update/nicmgr_upgstate";
const char* nicmgr_rollback_state_file = "/update/nicmgr_rollback_state";
static evutil_check log_check;

static void
log_flush (void *arg)
{
    fflush(stdout);
    fflush(stderr);
    if (utils::logger::logger()) {
        utils::logger::logger()->flush();
    }
}

static bool
upgrade_in_progress (void)
{
    return (access(nicmgr_upgrade_state_file, R_OK) == 0);
}

static bool
rollback_in_progress (void)
{
    return (access(nicmgr_rollback_state_file, R_OK) == 0);
}

static void
hal_up_event_handler (sdk::ipc::ipc_msg_ptr msg, const void *ctxt)
{
    NIC_LOG_DEBUG("IPC, HAL UP event handler ...");
    devmgr->HalEventHandler(true);
}

static void
port_event_handler (sdk::ipc::ipc_msg_ptr msg, const void *ctxt)
{
    port_status_t st = { 0 };
    hal::core::event_t *event = (hal::core::event_t *)msg->data();

    st.id = event->port.id;
    st.status =
        (event->port.event == port_event_t::PORT_EVENT_LINK_UP) ? 1 : 0;
    st.speed = sdk::lib::port_speed_enum_to_mbps(event->port.speed);
    NIC_LOG_DEBUG("IPC, Rcvd port event for id {}, speed {}, status {}",
                    st.id, st.speed, st.status);
    devmgr->LinkEventHandler(&st);
}

static void
xcvr_event_handler (sdk::ipc::ipc_msg_ptr msg, const void *ctxt)
{
    port_status_t st = { 0 };
    hal::core::event_t *event = (hal::core::event_t *)msg->data();

    st.id = event->port.id;
    st.xcvr.state = event->xcvr.state;
    st.xcvr.pid = event->xcvr.pid;
    st.xcvr.phy = event->xcvr.cable_type;
    memcpy(st.xcvr.sprom, event->xcvr.sprom, XCVR_SPROM_SIZE);
    NIC_LOG_DEBUG("IPC, Rcvd xcvr event for id {}, state {}, cable type {}, pid {}",
                  st.id, st.xcvr.state, st.xcvr.phy, st.xcvr.pid);
    devmgr->XcvrEventHandler(&st);
#if 0 // why this is required HAREESH - Check with Sarat.  TODO
    devmgr->LinkEventHandler(&st);
#endif
}

static void
micro_seg_event_handler (sdk::ipc::ipc_msg_ptr msg, const void *ctxt)
{
    hal::core::event_t *event = (hal::core::event_t *)msg->data();

    NIC_LOG_DEBUG("System spec update: micro_seg_en: {}", event->mseg.status);
    devmgr->SystemSpecEventHandler(event->mseg.status);
}

static void
register_for_events (void)
{
    // register for hal up and port events
    sdk::ipc::subscribe(event_id_t::EVENT_ID_PORT_STATUS, port_event_handler, NULL);
    sdk::ipc::subscribe(event_id_t::EVENT_ID_XCVR_STATUS, xcvr_event_handler, NULL);
    sdk::ipc::subscribe(event_id_t::EVENT_ID_HAL_UP, hal_up_event_handler, NULL);
    sdk::ipc::subscribe(event_id_t::EVENT_ID_MICRO_SEG, micro_seg_event_handler, NULL);
}

namespace nicmgr {

void
nicmgr_init (platform_type_t platform,
             struct sdk::event_thread::event_thread *thread)
{

    string profile;
    string device_file;
    bool micro_seg_en = false;
    sdk::lib::device *device = NULL;
    sdk::lib::dev_forwarding_mode_t fwd_mode;
    sdk::lib::dev_feature_profile_t feature_profile;
    UpgradeMode upg_mode;

    // instantiate the logger
    utils::logger::init();

    if (platform == platform_type_t::PLATFORM_TYPE_SIM) {
        profile = std::string(getenv("HAL_CONFIG_PATH")) + "/../../" +
                  "platform/src/app/nicmgrd/etc/eth.json";
        fwd_mode = sdk::lib::FORWARDING_MODE_CLASSIC;
        goto dev_init;
    } else {
        device_file = std::string(SYSCONFIG_PATH) +  "/" + DEVICE_CFG_FNAME;
    }

    if (device_file.empty()) {
        NIC_LOG_ERR("No device file");
        exit(1);
    }

    // Load device configuration
    device = sdk::lib::device::factory(device_file.c_str());
    fwd_mode = device->get_forwarding_mode();
    feature_profile = device->get_feature_profile();
    micro_seg_en = (device->get_micro_seg_en() == device::MICRO_SEG_ENABLE);

#if 0
    // TODO: Profile should be independent of forwarding mode.
    // TODO: No need to figure out the profile while upgrading.
    if (fwd_mode == sdk::lib::FORWARDING_MODE_HOSTPIN ||
        fwd_mode == sdk::lib::FORWARDING_MODE_SWITCH) {
        profile = "/platform/etc/nicmgrd/eth_smart.json";
    } else {
        if (feature_profile == sdk::lib::FEATURE_PROFILE_CLASSIC_ETH_DEV_SCALE) {
            profile = "/platform/etc/nicmgrd/eth_scale.json";
        } else {
            profile = "/platform/etc/nicmgrd/device.json";
        }
    }

    NIC_LOG_INFO("Forwarding Mode {}", fwd_mode);
    NIC_LOG_INFO("Micro-segmentation {}", micro_seg_en);
    NIC_LOG_INFO("Feature Profile {} {}", feature_profile, profile);
#endif

dev_init:
    // Are we in the middle of an upgrade?
    if (rollback_in_progress()) {
        upg_mode = FW_MODE_ROLLBACK;
    } else if (upgrade_in_progress()) {
        upg_mode = FW_MODE_UPGRADE;
    } else {
        upg_mode = FW_MODE_NORMAL_BOOT;
    }

    NIC_LOG_INFO("Upgrade mode: {}", upg_mode);

    register_for_events();

    devmgr = new DeviceManager(platform, device_file, fwd_mode, micro_seg_en,
                               thread->ev_loop());
    devmgr->SetUpgradeMode(upg_mode);
    devmgr->SetThread((sdk::lib::thread *)thread);

    if (upg_mode == FW_MODE_NORMAL_BOOT) {
        devmgr->LoadProfile(profile, true);
    } else if (upg_mode == FW_MODE_ROLLBACK) {
        devmgr->LoadProfile(profile, false);
        unlink(nicmgr_rollback_state_file);
    } else {
        // Restore States will be done
        unlink(nicmgr_upgrade_state_file);
    }

    evutil_add_check(thread->ev_loop(), &log_check, &log_flush, NULL);

    // upgrade init
    nicmgr::nicmgr_upg_init();

    NIC_LOG_INFO("Listening to events");
}

void
nicmgr_exit (void)
{
    if (devmgr) {
        delete devmgr;
    }
    devmgr = NULL;
}

}   // namespace nicmgr