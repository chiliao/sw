/**
 * Copyright (c) 2019 Pensando Systems, Inc.
 *
 * @file    port.hpp
 *
 * @brief   This file handles port operations
 */

#include "nic/sdk/linkmgr/port.hpp"
#include "nic/sdk/platform/drivers/xcvr.hpp"
#include "nic/sdk/include/sdk/if.hpp"
#include "nic/sdk/include/sdk/ip.hpp"
#include "nic/sdk/lib/ipc/ipc.hpp"
#include "nic/sdk/linkmgr/linkmgr.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/core/core.hpp"
#include "nic/apollo/core/event.hpp"
#include "nic/apollo/api/pds_state.hpp"
#include "nic/apollo/api/utils.hpp"
#include "nic/apollo/api/if.hpp"
#include "nic/apollo/api/port.hpp"
#include "nic/apollo/api/internal/metrics.hpp"
#include "nic/operd/alerts/alerts.hpp"

namespace api {

typedef struct port_get_cb_ctxt_s {
    void *ctxt;
    port_get_cb_t port_get_cb;
} port_get_cb_ctxt_t;

typedef struct port_shutdown_walk_cb_ctxt_s {
    bool port_shutdown;
    bool port_pb_shutdown;
    bool err;
} port_shutdown_walk_cb_ctxt_t;

/**
 * @brief        Handle link UP/Down events
 * @param[in]    port_event_info port event information
 */
void
port_event_cb (port_event_info_t *port_event_info)
{
    int phy_port;
    sdk_ret_t ret;
    if_entry *intf;
    pds_obj_key_t key;
    ::core::event_t event;
    pds_event_t pds_event;
    port_event_t port_event = port_event_info->event;
    port_speed_t port_speed = port_event_info->speed;
    port_fec_type_t fec_type = port_event_info->fec_type;
    uint32_t logical_port = port_event_info->logical_port;

    sdk::linkmgr::port_set_leds(logical_port, port_event);

    // broadcast the event to other components of interest
    memset(&event, 0, sizeof(event));
    event.event_id = EVENT_ID_PORT_STATUS;
    event.port.ifindex =
        sdk::lib::catalog::logical_port_to_ifindex(logical_port);
    event.port.event = port_event;
    event.port.speed = port_speed;
    event.port.fec_type = fec_type;
    sdk::ipc::broadcast(EVENT_ID_PORT_STATUS, &event, sizeof(event));

    // notify the agent
    memset(&pds_event, 0, sizeof(pds_event));
    if (port_event_info->event == port_event_t::PORT_EVENT_LINK_DOWN) {
        pds_event.event_id = PDS_EVENT_ID_PORT_DOWN;
    } else if (port_event_info->event == port_event_t::PORT_EVENT_LINK_UP) {
        pds_event.event_id = PDS_EVENT_ID_PORT_UP;
    } else {
        SDK_ASSERT(FALSE);
    }
    // find interface instance corresponding to this port
    intf = if_db()->find(&event.port.ifindex);
    if (intf == NULL)  {
        PDS_TRACE_ERR("Interface instance for ifindex 0x%x not found",
                      event.port.ifindex);
        return;
    }
    key = intf->key();
    ret = sdk::linkmgr::port_get(intf->port_info(), &pds_event.port_info.info);
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Failed to get port 0x%x info, err %u",
                      intf->ifindex(), ret);
        return;
    }
    pds_event.port_info.info.port_num = intf->ifindex();
    phy_port = sdk::lib::catalog::ifindex_to_phy_port(pds_event.port_info.info.port_num);
    if (phy_port != -1) {
        ret = sdk::platform::xcvr_get(phy_port - 1, &pds_event.port_info.info.xcvr_event_info);
        if (ret != SDK_RET_OK) {
            PDS_TRACE_ERR("Failed to get xcvr for port %u, err %u",
                          phy_port, ret);
            return;
        }
    }
    // TODO: @akoradha same as if_walk_port_get_cb(), please fix
    pds_event.port_info.info.port_an_args = (port_an_args_t *)&key;
    // notify the agent
    g_pds_state.event_notify(&pds_event);

    // Raise an event
    operd::alerts::operd_alerts_t alert = operd::alerts::LINK_DOWN;
    if (port_event_info->event == port_event_t::PORT_EVENT_LINK_UP) {
        alert = operd::alerts::LINK_UP;
    }
    operd::alerts::alert_recorder::get()->alert(
        alert, "Port %s, key %s", intf->name().c_str(), key.str());
}

bool
xvcr_event_walk_cb (void *entry, void *ctxt)
{
    int phy_port;
    ::core::event_t event;
    uint32_t logical_port;
    if_index_t ifindex;
    if_entry *intf = (if_entry *)entry;
    xcvr_event_info_t *xcvr_event_info = (xcvr_event_info_t *)ctxt;

    ifindex = intf->ifindex();
    logical_port = sdk::lib::catalog::ifindex_to_logical_port(ifindex);
    phy_port = sdk::lib::catalog::logical_port_to_phy_port(logical_port);
    if ((phy_port == -1) ||
        (phy_port != (int)xcvr_event_info->phy_port)) {
        return false;
    }
    sdk::linkmgr::port_update_xcvr_event(intf->port_info(), xcvr_event_info);

    memset(&event, 0, sizeof(event));
    event.xcvr.ifindex = ifindex;
    event.xcvr.state = xcvr_event_info->state;
    event.xcvr.pid = xcvr_event_info->pid;
    event.xcvr.cable_type = xcvr_event_info->cable_type;
    memcpy(event.xcvr.sprom, xcvr_event_info->xcvr_sprom, XCVR_SPROM_SIZE);
    sdk::ipc::broadcast(EVENT_ID_XCVR_STATUS, &event, sizeof(event));
    return false;
}

/**
 * @brief        Handle transceiver insert/remove events
 * @param[in]    xcvr_event_info    transceiver info filled by linkmgr
 */
void
xcvr_event_cb (xcvr_event_info_t *xcvr_event_info)
{
    /**< ignore xcvr events if xcvr valid check is disabled */
    if (!sdk::platform::xcvr_valid_check_enabled()) {
        return;
    }

    /**
     * if xcvr is removed, bring link down
     * if xcvr sprom read is successful, bring linkup if user admin enabled.
     * ignore all other xcvr states.
     */
    if (xcvr_event_info->state != xcvr_state_t::XCVR_REMOVED &&
        xcvr_event_info->state != xcvr_state_t::XCVR_SPROM_READ) {
        return;
    }
    if_db()->walk(IF_TYPE_ETH, xvcr_event_walk_cb, xcvr_event_info);
}

bool
port_shutdown_walk_cb (void *entry, void *ctxt)
{
    if_entry *intf = (if_entry *)entry;
    port_shutdown_walk_cb_ctxt_t *ctx = (port_shutdown_walk_cb_ctxt_t *)ctxt;
    sdk_ret_t ret;

    if (ctx->port_shutdown) {
        ret = sdk::linkmgr::port_shutdown(intf->port_info());
    } else if (ctx->port_pb_shutdown) {
        ret = sdk::linkmgr::port_pb_shutdown(intf->port_info());
    } else {
        ret = SDK_RET_ERR;
    }

    if (ret != SDK_RET_OK) {
        ctx->err = true;
    }

    return false;
}

sdk_ret_t
port_shutdown_all (void)
{
    port_shutdown_walk_cb_ctxt_t ctxt = { 0 };

    // let first disable all
    ctxt.port_shutdown = true;
    if_db()->walk(IF_TYPE_ETH, port_shutdown_walk_cb, &ctxt);
    if (ctxt.err) {
       return SDK_RET_ERR;
    }

    // TODO : check with team on whether the 2 loop walk is needed or not
    // pb disable
    ctxt.port_shutdown = false;
    ctxt.port_pb_shutdown = true;
    if_db()->walk(IF_TYPE_ETH, port_shutdown_walk_cb, &ctxt);
    if (ctxt.err) {
       return SDK_RET_ERR;
    }
    return SDK_RET_OK;
}

/**
 * @brief        update a port with the given configuration information
 * @param[in]    key           key/uuid of the port
 * @param[in]    api_port_info port info
 * @return       SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
port_update (const pds_obj_key_t *key, port_args_t *api_port_info)
{
    sdk_ret_t ret;
    if_entry *intf;
    port_args_t port_info;

    intf = if_db()->find(key);
    if (intf == NULL) {
        PDS_TRACE_ERR("port %s update failed", key->str());
        return SDK_RET_ENTRY_NOT_FOUND;
    }
    memset(&port_info, 0, sizeof(port_info));

    ret = sdk::linkmgr::port_get(intf->port_info(), &port_info);
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Failed to get port %s info, err %u", key->str(), ret);
        return ret;
    }
    api_port_info->tx_pause_enable = port_info.tx_pause_enable;
    api_port_info->rx_pause_enable = port_info.rx_pause_enable;

    // sdk port_num is logical port
    api_port_info->port_num =
        sdk::lib::catalog::ifindex_to_logical_port(intf->ifindex());

    // update port_args based on the xcvr state
    sdk::linkmgr::port_args_set_by_xcvr_state(api_port_info);

    ret = sdk::linkmgr::port_update(intf->port_info(), api_port_info);
    return ret;
}

/**
 * @brief        create a port with the given configuration information
 * @param[in]    ifindex      interface index
 * @param[in]    port_args    port parameters filled by this API
 * @return       SDK_RET_OK on success, failure status code on error
 */
static sdk_ret_t
create_port (if_index_t ifindex, port_args_t *port_args)
{
    if_entry *intf;
    void *port_info;
    pds_obj_key_t key;

    PDS_TRACE_DEBUG("Creating port %u, ifindex 0x%x",
                    port_args->port_num, ifindex);

    sdk::linkmgr::port_store_user_config(port_args);
    sdk::linkmgr::port_args_set_by_xcvr_state(port_args);
    port_info = sdk::linkmgr::port_create(port_args);
    if (port_info == NULL) {
        PDS_TRACE_ERR("port %u create failed", port_args->port_num);
        return SDK_RET_ERR;
    }
    key = uuid_from_objid(ifindex);
    intf = if_entry::factory(key, ifindex);
    if (intf == NULL) {
        sdk::linkmgr::port_delete(port_info);
        return SDK_RET_ERR;
    }
    intf->set_port_info(port_info);
    if_db()->insert(intf);
    // register the stats region with metrics submodule
    if (port_args->port_type == port_type_t::PORT_TYPE_ETH) {
        sdk::metrics::row_address(g_pds_state.port_metrics_handle(),
                                  *(sdk::metrics::key_t *)key.id,
                                  (void *)sdk::linkmgr::port_stats_addr(ifindex));
    } else if (port_args->port_type == port_type_t::PORT_TYPE_MGMT) {
        sdk::metrics::row_address(g_pds_state.mgmt_port_metrics_handle(),
                                  *(sdk::metrics::key_t *)key.id,
                                  (void *)sdk::linkmgr::port_stats_addr(ifindex));
    }
    return SDK_RET_OK;
}

/**
 * @brief        populate port information based on the catalog
 * @param[in]    ifindex     interface index of this port
 * @param[in]    phy_port    physical port number of this port
 * @param[out]   port_args    port parameters filled by this API
 * @return       SDK_RET_OK on success, failure status code on error
 */
static sdk_ret_t
populate_port_info (if_index_t ifindex, uint32_t phy_port,
                    port_args_t *port_args)
{
    uint32_t    logical_port;

    logical_port = port_args->port_num =
        sdk::lib::catalog::ifindex_to_logical_port(ifindex);
    port_args->port_type = g_pds_state.catalogue()->port_type_fp(phy_port);
    port_args->port_speed = g_pds_state.catalogue()->port_speed_fp(phy_port);
    port_args->fec_type = g_pds_state.catalogue()->port_fec_type_fp(phy_port);
    if (port_args->port_type == port_type_t::PORT_TYPE_MGMT) {
        port_args->port_speed = port_speed_t::PORT_SPEED_1G;
        port_args->fec_type = port_fec_type_t::PORT_FEC_TYPE_NONE;
    } else {
        port_args->auto_neg_enable = true;
    }
    port_args->admin_state = g_pds_state.catalogue()->admin_state_fp(phy_port);
    port_args->num_lanes = g_pds_state.catalogue()->num_lanes_fp(phy_port);
    port_args->mac_id = g_pds_state.catalogue()->mac_id(logical_port, 0);
    port_args->mac_ch = g_pds_state.catalogue()->mac_ch(logical_port, 0);
    port_args->debounce_time = 0;
    port_args->mtu = 0;    /**< default will be set to max mtu */
    port_args->pause = port_pause_type_t::PORT_PAUSE_TYPE_NONE;
    port_args->loopback_mode = port_loopback_mode_t::PORT_LOOPBACK_MODE_NONE;

    for (uint32_t i = 0; i < port_args->num_lanes; i++) {
        port_args->sbus_addr[i] =
            g_pds_state.catalogue()->sbus_addr(logical_port, i);
    }
    port_args->breakout_modes =
        g_pds_state.catalogue()->breakout_modes(phy_port);

    return SDK_RET_OK;
}

/**
 * @brief     create all ports based on the catalog information
 * @return    SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
create_ports (void)
{
    uint32_t       num_phy_ports;
    port_args_t    port_args;
    if_index_t     ifindex;

    PDS_TRACE_DEBUG("Creating ports ...");
    num_phy_ports = g_pds_state.catalogue()->num_fp_ports();
    for (uint32_t phy_port = 1; phy_port <= num_phy_ports; phy_port++) {
        ifindex = ETH_IFINDEX(g_pds_state.catalogue()->slot(),
                              phy_port, ETH_IF_DEFAULT_CHILD_PORT);
        memset(&port_args, 0, sizeof(port_args));
        populate_port_info(ifindex, phy_port, &port_args);
        create_port(ifindex, &port_args);
    }
    return SDK_RET_OK;
}

bool
if_walk_port_get_cb (void *entry, void *ctxt)
{
    int phy_port;
    sdk_ret_t ret;
    pds_obj_key_t key;
    port_args_t port_info;
    if_entry *intf = (if_entry *)entry;
    uint64_t stats_data[MAX_MAC_STATS];
    port_get_cb_ctxt_t *cb_ctxt = (port_get_cb_ctxt_t *)ctxt;

    key = intf->key();
    memset(&port_info, 0, sizeof(port_info));
    port_info.stats_data = stats_data;
    ret = sdk::linkmgr::port_get(intf->port_info(), &port_info);
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Failed to get port 0x%s info, err %u", intf->key().str(),
                      ret);
        return false;
    }
    port_info.port_num = intf->ifindex();
    phy_port = sdk::lib::catalog::ifindex_to_phy_port(port_info.port_num);
    if (phy_port != -1) {
        ret = sdk::platform::xcvr_get(phy_port - 1, &port_info.xcvr_event_info);
        if (ret != SDK_RET_OK) {
            PDS_TRACE_ERR("Failed to get xcvr for port %u, err %u",
                          phy_port, ret);
        }
    }
    // TODO: @akoradha port_args is exposed all the way to the agent
    //       with the current design, we should create port_spec_t,
    //       port_status_t and port_stats_t like any other object or
    //       better approach is to fold all the port stuff into if_entry
    //       and CLIs etc. will naturally work with current db walks etc.
    //       we have all eth ports in if db already. with port_args_t
    //       going directly upto agent svc layer, there is no way to send uuid
    //       now, so hijacking this pointer field
    port_info.port_an_args = (port_an_args_t *)&key;
    cb_ctxt->port_get_cb(&port_info, cb_ctxt->ctxt);
    return false;
}

/**
 * @brief    get port information based on port number
 * @param[in]    key         key/uuid of the port or k_pds_obj_key_invalid for
 *                           all ports
 * @param[in]    port_get_cb callback invoked per port
 * @param[in]    ctxt        opaque context passed back to the callback
 * @return    SDK_RET_OK on success, failure status code on error
 */
sdk_ret_t
port_get (const pds_obj_key_t *key, port_get_cb_t port_get_cb, void *ctxt)
{
    if_entry *intf;
    port_get_cb_ctxt_t cb_ctxt;

    cb_ctxt.ctxt = ctxt;
    cb_ctxt.port_get_cb = port_get_cb;
    if (*key == k_pds_obj_key_invalid) {
        if_db()->walk(IF_TYPE_ETH, if_walk_port_get_cb, &cb_ctxt);
    } else {
        intf = if_db()->find(key);
        if (intf == NULL)  {
            PDS_TRACE_ERR("Port %s not found", key->str());
            return SDK_RET_INVALID_OP;
        }
        if_walk_port_get_cb(intf, &cb_ctxt);
    }
    return SDK_RET_OK;
}

// @brief    get port information based on ifindex
// @param[in]    key         ifindex of the port or 0 for all ports
// @param[in]    port_get_cb callback invoked per port
// @param[in]    ctxt        opaque context passed back to the callback
// @return    SDK_RET_OK on success, failure status code on error
sdk_ret_t
port_get (const if_index_t *key, port_get_cb_t port_get_cb, void *ctxt)
{
    if_entry *intf;
    api::port_get_cb_ctxt_t cb_ctxt;

    cb_ctxt.ctxt = ctxt;
    cb_ctxt.port_get_cb = port_get_cb;
    if (*key == IFINDEX_INVALID) {
        if_db()->walk(IF_TYPE_ETH, if_walk_port_get_cb, &cb_ctxt);
    } else {
        intf = if_db()->find(key);
        if (intf == NULL)  {
            PDS_TRACE_ERR("Port 0x%x not found", *key);
            return SDK_RET_INVALID_OP;
        }
        if_walk_port_get_cb(intf, &cb_ctxt);
    }
    return SDK_RET_OK;
}

static bool
if_walk_port_stats_reset_cb (void *entry, void *ctxt)
{
    sdk_ret_t ret;
    if_entry *intf = (if_entry *)entry;

    ret = sdk::linkmgr::port_stats_reset(intf->port_info());
    if (ret != SDK_RET_OK) {
        PDS_TRACE_ERR("Failed to reset port stats for %s, err %u",
                      eth_ifindex_to_str(intf->ifindex()).c_str(),
                      ret);
    }
    return false;
}

sdk_ret_t
port_stats_reset (const pds_obj_key_t *key)
{
    if_entry *intf;

    if ((key == NULL) || (*key == k_pds_obj_key_invalid)) {
        if_db()->walk(IF_TYPE_ETH, if_walk_port_stats_reset_cb, NULL);
    } else {
        intf = if_db()->find(key);
        if (intf == NULL)  {
            PDS_TRACE_ERR("Port %s not found",
                          eth_ifindex_to_str(intf->ifindex()).c_str());
            return SDK_RET_INVALID_OP;
        }
        if_walk_port_stats_reset_cb(intf, NULL);
    }
    return SDK_RET_OK;
}

}    // namespace api
