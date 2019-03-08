//-----------------------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------
#include "platform/src/lib/nicmgr/include/logger.hpp"
#include "devapi_iris.hpp"
#include "devapi_mem.hpp"
#include "lif.hpp"
#include "uplink.hpp"
#include "qos.hpp"
#include "port.hpp"
#include "accel.hpp"

namespace iris {

//-----------------------------------------------------------------------------
// Factory method to instantiate the class
//-----------------------------------------------------------------------------
devapi_iris *
devapi_iris::factory(void)
{
    sdk_ret_t ret      = SDK_RET_OK;
    void *mem          = NULL;
    devapi_iris *diris = NULL;

    mem = (devapi_iris *)DEVAPI_CALLOC(DEVAPI_MEM_ALLOC_DEVAPI_IRIS,
                                       sizeof(devapi_iris));
    if (mem) {
        diris = new (mem) devapi_iris();
        ret = diris->init_();
        if (ret != SDK_RET_OK) {
            diris->~devapi_iris();
            DEVAPI_FREE(DEVAPI_MEM_ALLOC_DEVAPI_IRIS, mem);
            diris = NULL;
        }
    }
    return diris;
}

//-----------------------------------------------------------------------------
// Initializing fields in the class
//-----------------------------------------------------------------------------
sdk_ret_t
devapi_iris::init_(void)
{
    // Initialize grpc
    hal_grpc::factory();
    // Populate hal_grpc handler in object
    devapi_object::populate_hal_grpc();
    return SDK_RET_OK;
}

void
devapi_iris::destroy(devapi *dapi)
{
    devapi_iris *dapi_iris = dynamic_cast<devapi_iris*>(dapi);
    devapi_iris::destroy(dapi_iris);
}

void
devapi_iris::destroy(devapi_iris *dapi)
{
    NIC_LOG_DEBUG("Destroy devapi_iris");
    dapi->~devapi_iris();
    DEVAPI_FREE(DEVAPI_MEM_ALLOC_DEVAPI_IRIS, dapi);
}

sdk_ret_t
devapi_iris::lif_create(lif_info_t *info)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;
    lif = devapi_lif::factory(info);
    if (!lif) {
        NIC_LOG_ERR("Failed to create lif. id: {}", info->lif_id);
        ret = SDK_RET_ERR;
    }
    return ret;
}

sdk_ret_t
devapi_iris::lif_destroy(uint32_t lif_id)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (lif) {
        devapi_lif::destroy(lif);
    } else {
        NIC_LOG_ERR("Failed to destroy lif. id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
    }
    return ret;
}

sdk_ret_t
devapi_iris::lif_reset(uint32_t lif_id)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to reset lif. id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->reset();
end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_add_mac(uint32_t lif_id, mac_t mac)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to add mac. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->add_mac(mac);

end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_del_mac(uint32_t lif_id, mac_t mac)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to del mac. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->del_mac(mac);

end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_add_vlan(uint32_t lif_id, vlan_t vlan)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to add vlan. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->add_vlan(vlan);

end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_del_vlan(uint32_t lif_id, vlan_t vlan)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to del vlan. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->del_vlan(vlan);

end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_add_macvlan(uint32_t lif_id, mac_t mac, vlan_t vlan)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to add vlan. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->add_macvlan(mac, vlan);

end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_del_macvlan(uint32_t lif_id, mac_t mac, vlan_t vlan)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to del vlan. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->del_macvlan(mac, vlan);

end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_upd_vlan_offload(uint32_t lif_id, bool vlan_strip,
                                  bool vlan_insert)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to update vlan offl. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->upd_vlanoff(vlan_strip, vlan_insert);

end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_upd_rx_mode(uint32_t lif_id, bool broadcast,
                             bool all_multicast, bool promiscuous)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to update rx mode. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->upd_rxmode(broadcast, all_multicast, promiscuous);

end:
    return ret;
}

sdk_ret_t
devapi_iris::lif_upd_name(uint32_t lif_id, std::string name)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_lif *lif = NULL;

    lif = devapi_lif::lookup(lif_id);
    if (!lif) {
        NIC_LOG_ERR("Failed to update name. lif id: {}. Not found",
                    lif_id);
        ret = SDK_RET_ERR;
        goto end;
    }
    return lif->upd_name(name);

end:
    return ret;
}

sdk_ret_t
devapi_iris::qos_get_txtc_cos(const std::string &group, uint32_t uplink_port,
                              uint8_t *cos)
{
    *cos = devapi_qos::get_txtc_cos(group, uplink_port);
    return SDK_RET_OK;
}

sdk_ret_t
devapi_iris::set_fwd_mode(fwd_mode_t fwd_mode)
{
    fwd_mode_ = fwd_mode;
    hal_grpc::get_hal_grpc()->set_fwd_mode(fwd_mode);
    return SDK_RET_OK;
}

sdk_ret_t
devapi_iris::uplink_create(uint32_t id, uint32_t port, bool is_oob)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_uplink *up = NULL;
    up = devapi_uplink::factory(id, port, is_oob);
    if (!up) {
        NIC_LOG_ERR("Failed to create uplink. id: {}", id);
        ret = SDK_RET_ERR;
    }
    // Create vrf
    up->create_vrf();
    return ret;
}

sdk_ret_t
devapi_iris::uplink_destroy(uint32_t port)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_uplink *up = NULL;

    up = devapi_uplink::get_uplink(port);
    if (!up) {
        NIC_LOG_ERR("Failed to destroy uplink. port: {}", port);
        ret = SDK_RET_ERR;
        goto end;
    }
    // Remove vrf
    up->delete_vrf();
    // Destroy uplink
    devapi_uplink::destroy(up);

end:
    return ret;
}

sdk_ret_t
devapi_iris::port_get_status(uint32_t port_num,
                             port_status_t *status)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_port *port = NULL;

    port = devapi_port::find_or_create(port_num);
    if (!port) {
        NIC_LOG_ERR("Failed to find/create port: {}", port_num);
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = port->port_hal_get_status(status);

end:
    return ret;
}

sdk_ret_t
devapi_iris::port_get_config(uint32_t port_num,
                             port_config_t *config)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_port *port = NULL;

    port = devapi_port::find_or_create(port_num);
    if (!port) {
        NIC_LOG_ERR("Failed to find/create port: {}", port_num);
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = port->port_hal_get_config(config);

end:
    return ret;
}

sdk_ret_t
devapi_iris::port_set_config(uint32_t port_num,
                             port_config_t *config)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_port *port = NULL;

    port = devapi_port::find_or_create(port_num);
    if (!port) {
        NIC_LOG_ERR("Failed to find/create port: {}", port_num);
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = port->port_hal_update_config(config);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_add(string name)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_add(name);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_del(string name)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_del(name);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_ring_add(string name,
                                   std::vector<std::pair<const std::string,uint32_t>>& ring_vec)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_ring_add(name, ring_vec);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_ring_del(string name,
                                   std::vector<std::pair<const std::string,uint32_t>>& ring_vec)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_ring_del(name, ring_vec);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_reset_set(string name, uint32_t sub_ring,
                                    bool reset_sense)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_reset_set(name, sub_ring, reset_sense);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_enable_set(string name, uint32_t sub_ring,
                                     bool enable_sense)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_enable_set(name, sub_ring, enable_sense);

end:
    return ret;
}


sdk_ret_t
devapi_iris::accel_rgroup_pndx_set(string name, uint32_t sub_ring,
                                   uint32_t val, bool conditional)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_pndx_set(name, sub_ring, val, conditional);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_info_get(string name, uint32_t sub_ring,
                                   accel_rgroup_rinfo_rsp_cb_t rsp_cb_func,
                                   void *user_ctx, uint32_t *ret_num_entries)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_info_get(name, sub_ring, rsp_cb_func,
                                       user_ctx, ret_num_entries);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_indices_get(string name,uint32_t sub_ring,
                                      accel_rgroup_rindices_rsp_cb_t rsp_cb_func,
                                      void *user_ctx, uint32_t *ret_num_entries)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_indices_get(name, sub_ring, rsp_cb_func,
                                          user_ctx, ret_num_entries);

end:
    return ret;
}

sdk_ret_t
devapi_iris::accel_rgroup_metrics_get(string name, uint32_t sub_ring,
                                      accel_rgroup_rmetrics_rsp_cb_t rsp_cb_func,
                                      void *user_ctx, uint32_t *ret_num_entries)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->accel_rgroup_metrics_get(name, sub_ring, rsp_cb_func,
                                          user_ctx, ret_num_entries);

end:
    return ret;
}

sdk_ret_t
devapi_iris::crypto_key_index_upd(uint32_t key_index, crypto_key_type_t type,
                                  void *key, uint32_t key_size)
{
    sdk_ret_t ret = SDK_RET_OK;
    devapi_accel *accel = NULL;

    accel = devapi_accel::find_or_create();
    if (!accel) {
        NIC_LOG_ERR("Failed to find/create accel");
        ret = SDK_RET_ERR;
        goto end;
    }
    ret = accel->crypto_key_index_upd(key_index, type, key, key_size);

end:
    return ret;
}


}    // namespace iris
