//
// {C} Copyright 2018 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// device entry handling
///
//----------------------------------------------------------------------------

#include "sys/utsname.h"
#include "nic/sdk/include/sdk/base.hpp"
#include "nic/sdk/include/sdk/mem.hpp"
#include "nic/sdk/include/sdk/platform.hpp"
#include "nic/sdk/platform/fru/fru.hpp"
#include "nic/sdk/asic/asic.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/core/mem.hpp"
#include "nic/apollo/framework/impl.hpp"
#include "nic/apollo/framework/api_base.hpp"
#include "nic/apollo/framework/api_engine.hpp"
#include "nic/apollo/framework/api_params.hpp"
#include "nic/apollo/api/device.hpp"
#include "nic/apollo/api/pds_state.hpp"
#include "nic/sdk/platform/drivers/ionic_base.h"

namespace api {

/// \defgroup PDS_DEVICE_ENTRY - device entry functionality
/// \ingroup PDS_DEVICE
/// \@{

device_entry *
device_entry::factory(pds_device_spec_t *pds_device) {
    device_entry *device;

    device = device_db()->alloc();
    if (device) {
        new (device) device_entry();
        device->impl_ =
            impl_base::factory(impl::IMPL_OBJ_ID_DEVICE, pds_device);
        if (device->impl_ == NULL) {
            device_entry::destroy(device);
            return NULL;
        }
    }
    return device;
}

void
device_entry::destroy(device_entry *device) {
    device->nuke_resources_();
    if (device->impl_) {
        impl_base::destroy(impl::IMPL_OBJ_ID_DEVICE, device->impl_);
    }
    device->~device_entry();
    device_db()->free(device);
}

api_base *
device_entry::clone(api_ctxt_t *api_ctxt) {
    device_entry *cloned_device;

    cloned_device = device_db()->alloc();
    if (cloned_device) {
        new (cloned_device) device_entry();
        if (cloned_device->init_config(api_ctxt) != SDK_RET_OK) {
            goto error;
        }
        cloned_device->impl_ = impl_->clone();
        if (unlikely(cloned_device->impl_ == NULL)) {
            PDS_TRACE_ERR("Failed to clone device impl");
            goto error;
        }
    }
    return cloned_device;

error:

    cloned_device->~device_entry();
    device_db()->free(cloned_device);
    return NULL;
}

sdk_ret_t
device_entry::free(device_entry *device) {
    if (device->impl_) {
        impl_base::free(impl::IMPL_OBJ_ID_DEVICE, device->impl_);
    }
    device->~device_entry();
    device_db()->free(device);
    return SDK_RET_OK;
}

sdk_ret_t
device_entry::init_config(api_ctxt_t *api_ctxt) {
    pds_device_spec_t *spec = &api_ctxt->api_params->device_spec;

    memcpy(&ip_addr_, &spec->device_ip_addr, sizeof(ip_addr_));
    memcpy(mac_addr_, spec->device_mac_addr, ETH_ADDR_LEN);
    memcpy(&gw_ip_addr_, &spec->gateway_ip_addr, sizeof(gw_ip_addr_));
    oper_mode_ = spec->dev_oper_mode;
    bridging_en_ = spec->bridging_en;
    learning_en_ = spec->learning_en;
    learn_age_timeout_ = spec->learn_age_timeout;
    overlay_routing_en_ = spec->overlay_routing_en;
    return SDK_RET_OK;
}

sdk_ret_t
device_entry::compute_update(api_obj_ctxt_t *obj_ctxt) {
    pds_device_spec_t *spec = &obj_ctxt->api_params->device_spec;

    if ((oper_mode_ != spec->dev_oper_mode) ||
        (bridging_en_ != spec->bridging_en) ||
        (learning_en_ != spec->learning_en) ||
        (overlay_routing_en_ != spec->overlay_routing_en)) {
        PDS_TRACE_WARN("Some of the device obj's attribute changes will take "
                       "affect after next reboot");
    }
    return SDK_RET_OK;
}

sdk_ret_t
device_entry::populate_msg(pds_msg_t *msg, api_obj_ctxt_t *obj_ctxt) {
    msg->cfg_msg.op = obj_ctxt->api_op;
    msg->cfg_msg.obj_id = OBJ_ID_DEVICE;
    if (obj_ctxt->api_op == API_OP_DELETE) {
        msg->cfg_msg.device.key = k_pds_obj_key_invalid;
    } else {
        msg->cfg_msg.device.spec.key = k_pds_obj_key_invalid;
        msg->cfg_msg.device.spec.spec = obj_ctxt->api_params->device_spec;
        if (impl_) {
            impl_->populate_msg(msg, this, obj_ctxt);
        }
    }
    return SDK_RET_OK;
}

sdk_ret_t
device_entry::program_update(api_base *orig_obj, api_obj_ctxt_t *obj_ctxt) {
    return impl_->update_hw(orig_obj, this, obj_ctxt);
}

sdk_ret_t
device_entry::activate_config(pds_epoch_t epoch, api_op_t api_op,
                              api_base *orig_obj, api_obj_ctxt_t *obj_ctxt) {
    return impl_->activate_hw(this, orig_obj, epoch, api_op, obj_ctxt);
}

void
device_entry::fill_spec_(pds_device_spec_t *spec) {
    spec->gateway_ip_addr = gw_ip_addr_;
    spec->dev_oper_mode = oper_mode_;
    spec->learn_age_timeout = learn_age_timeout_;
    spec->device_profile = api::g_pds_state.device_profile();
    spec->memory_profile = api::g_pds_state.memory_profile();
}

void
device_entry::fill_status(pds_device_status_t *status) {
    std::string   mac_str;
    std::string   mem_str;
    std::string   value, svalue;
    struct utsname ubuf;

    // fill fru mac in status
    sdk::platform::readfrukey(BOARD_MACADDRESS_KEY, mac_str);
    mac_str_to_addr((char *)mac_str.c_str(), status->fru_mac);

    mem_str = api::g_pds_state.catalogue()->memory_capacity_str();
    if (mem_str == "4g") {
        status->memory_cap = 4;
    } else if (mem_str == "8g") {
        status->memory_cap = 8;
    }

    sdk::platform::readfrukey(BOARD_MANUFACTURERDATE_KEY, value);
    if (value.empty()) {
        status->mnfg_date = std::string("-");
    } else {
        status->mnfg_date = value;
    }

    sdk::platform::readfrukey(BOARD_PARTNUM_KEY, value);
    if (value.empty() || value == "") {
        status->part_num = std::string("-");
    } else {
        status->part_num = value;
    }

    sdk::platform::readfrukey(BOARD_SERIALNUMBER_KEY, value);
    if (value.empty() || value == "") {
        status->serial_num = std::string("-");
    } else {
        status->serial_num = value;
    }

    sdk::platform::readfrukey(BOARD_PRODUCTNAME_KEY, value);
    if (value.empty() || value == "") {
        status->product_name = std::string("-");
    } else {
        status->product_name = value;
    }

    status->description = DSC_DESCRIPTION;
    status->vendor_name = DSC_VENDOR_NAME;
    status->vendor_id = DSC_VENDOR_ID;
    status->cpu_vendor = DSC_CPU_VENDOR;
    status->cpu_specification = DSC_CPU_SPECIFICATION;
    status->chip_type = api::g_pds_state.catalogue()->asic_type(0);
    sdk::platform::readfrukey(BOARD_PARTNUM_KEY, value);
    if (value.empty() || value == "") {
        status->hardware_revision = std::string("-");
    } else {
        svalue = value.substr(value.length() - 2, 2);
        status->hardware_revision = svalue;
    }
    status->pcie_specification = std::string("-");
    status->pcie_bus_info = std::string("-");
    status->fw_version = api::g_pds_state.firmware_version();
    status->fw_description = api::g_pds_state.firmware_description();
    status->fw_build_time = api::g_pds_state.firmware_build_time();
    if (uname(&ubuf)) {
        status->soc_os_version = std::string("-");
    } else {
        status->soc_os_version = std::string(ubuf.release);
    }
    status->soc_disk_size = std::string("-");
    status->num_pcie_ports = api::g_pds_state.catalogue()->pcie_nportspecs();
    status->num_ports = api::g_pds_state.catalogue()->num_fp_ports();
    status->pxe_version = IPXE_VERSION_CODE;
    status->uefi_version = UEFI_VERSION_CODE;
    status->num_host_if = lif_db()->num_host_lif();
}

sdk_ret_t
device_entry::read(pds_device_info_t *info) {
    fill_spec_(&info->spec);
    api::device_entry::fill_status(&info->status);
    return impl_->read_hw(this, NULL, (impl::obj_info_t *)info);
}

sdk_ret_t
device_entry::add_to_db(void) {
    return device_db()->insert(this);
}

sdk_ret_t
device_entry::del_from_db(void) {
    if (device_db()->remove()) {
        return SDK_RET_OK;
    }
    return SDK_RET_ENTRY_NOT_FOUND;
}

sdk_ret_t
device_entry::update_db(api_base *orig_obj, api_obj_ctxt_t *obj_ctxt) {
    if (device_db()->remove()) {
        return device_db()->insert(this);
    }
    return SDK_RET_ENTRY_NOT_FOUND;
}

sdk_ret_t
device_entry::delay_delete(void) {
    return delay_delete_to_slab(PDS_SLAB_ID_DEVICE, this);
}

/// \@}

}    // namespace api
