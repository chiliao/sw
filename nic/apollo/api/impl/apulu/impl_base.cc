//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// Base object definition for all impl objects
///
//----------------------------------------------------------------------------

#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/framework/impl.hpp"
#include "nic/apollo/framework/impl_base.hpp"
#include "nic/apollo/api/include/pds_device.hpp"
#include "nic/apollo/api/impl/apulu/device_impl.hpp"
#include "nic/apollo/api/impl/apulu/tep_impl.hpp"
#include "nic/apollo/api/impl/apulu/if_impl.hpp"
#include "nic/apollo/api/impl/apulu/vpc_impl.hpp"
#include "nic/apollo/api/impl/apulu/vnic_impl.hpp"
#include "nic/apollo/api/impl/apulu/mapping_impl.hpp"
#include "nic/apollo/api/impl/apulu/route_impl.hpp"
#include "nic/apollo/api/impl/apulu/mirror_impl.hpp"
#include "nic/apollo/api/impl/apulu/security_policy_impl.hpp"
#include "nic/apollo/api/impl/apulu/nexthop_impl.hpp"
#include "nic/apollo/api/impl/apulu/nexthop_group_impl.hpp"
#include "nic/apollo/api/impl/apulu/subnet_impl.hpp"
#include "nic/apollo/api/impl/apulu/policer_impl.hpp"
#include "nic/apollo/api/impl/apulu/service_impl.hpp"
#include "nic/apollo/api/impl/apulu/dhcp_impl.hpp"

namespace api {
namespace impl {

asic_impl_base *impl_base::asic_impl_  = NULL;
pipeline_impl_base *impl_base::pipeline_impl_  = NULL;

sdk_ret_t
impl_base::init(pds_init_params_t *params, asic_cfg_t *asic_cfg) {
    pipeline_cfg_t pipeline_cfg;

    // instanitiate asic implementaiton object
    asic_impl_ = asic_impl_base::factory(asic_cfg);
    SDK_ASSERT(asic_impl_ != NULL);

    // instanitiate pipeline implementaiton object
    pipeline_cfg.name = params->pipeline;
    pipeline_impl_ = pipeline_impl_base::factory(&pipeline_cfg);
    SDK_ASSERT(pipeline_impl_ != NULL);

    // initialize program and asm specific configs
    pipeline_impl_->program_config_init(params, asic_cfg);
    pipeline_impl_->asm_config_init(params, asic_cfg);
    pipeline_impl_->ring_config_init(asic_cfg);

    // perform asic initialization
    asic_impl_->asic_init(asic_cfg);

    // followed by pipeline initialization
    if (sdk::asic::asic_is_hard_init()) {
        if (sdk::platform::upgrade_mode_none(asic_cfg->upg_init_mode)) {
            pipeline_impl_->pipeline_init();
        } else if (sdk::platform::upgrade_mode_graceful(asic_cfg->upg_init_mode)) {
            pipeline_impl_->pipeline_upgrade_graceful_init();
        } else if (sdk::platform::upgrade_mode_hitless(asic_cfg->upg_init_mode)) {
            pipeline_impl_->pipeline_upgrade_hitless_init();
        }
    } else {
        pipeline_impl_->pipeline_soft_init();
    }
    // dump the MPU program related details
    pipeline_impl_->program_info_dump(asic_cfg->cfg_path);

    return SDK_RET_OK;
}

void
impl_base::destroy(void) {
    pipeline_impl_->destroy(pipeline_impl_);
    asic_impl_->destroy(asic_impl_);
}

impl_base *
impl_base::factory(impl_obj_id_t obj_id, void *spec) {
    switch (obj_id) {
    case IMPL_OBJ_ID_DEVICE:
        return device_impl::factory((pds_device_spec_t *)spec);

    case IMPL_OBJ_ID_TEP:
        return tep_impl::factory((pds_tep_spec_t *)spec);

    case IMPL_OBJ_ID_IF:
        return if_impl::factory((pds_if_spec_t *)spec);

    case IMPL_OBJ_ID_VPC:
        return vpc_impl::factory((pds_vpc_spec_t *)spec);

    case IMPL_OBJ_ID_SUBNET:
        return subnet_impl::factory((pds_subnet_spec_t *)spec);

    case IMPL_OBJ_ID_VNIC:
        return vnic_impl::factory((pds_vnic_spec_t *)spec);

    case IMPL_OBJ_ID_MAPPING:
        return mapping_impl::factory((pds_mapping_spec_t *)spec);

    case IMPL_OBJ_ID_ROUTE_TABLE:
        return route_table_impl::factory((pds_route_table_spec_t *)spec);

    case IMPL_OBJ_ID_SECURITY_POLICY:
        return security_policy_impl::factory((pds_policy_spec_t *)spec);

    case IMPL_OBJ_ID_SVC_MAPPING:
        return svc_mapping_impl::factory((pds_svc_mapping_spec_t *)spec);

    case IMPL_OBJ_ID_MIRROR_SESSION:
        return mirror_impl::factory((pds_mirror_session_spec_t *)spec);

    case IMPL_OBJ_ID_NEXTHOP:
        return nexthop_impl::factory((pds_nexthop_spec_t *)spec);

    case IMPL_OBJ_ID_NEXTHOP_GROUP:
        return nexthop_group_impl::factory((pds_nexthop_group_spec_t *)spec);

    case IMPL_OBJ_ID_POLICER:
        return policer_impl::factory((pds_policer_spec_t *)spec);

    case IMPL_OBJ_ID_DHCP_POLICY:
        return dhcp_policy_impl::factory((pds_dhcp_policy_spec_t *)spec);

    default:
        break;
    }
    return NULL;
}

sdk_ret_t
impl_base::free(impl_obj_id_t obj_id, impl_base *impl) {
    switch (obj_id) {
    case IMPL_OBJ_ID_DEVICE:
        return device_impl::free((device_impl *)impl);

    case IMPL_OBJ_ID_TEP:
        return tep_impl::free((tep_impl *)impl);

    case IMPL_OBJ_ID_IF:
        return if_impl::free((if_impl *)impl);

    case IMPL_OBJ_ID_VPC:
        return vpc_impl::free((vpc_impl *)impl);

    case IMPL_OBJ_ID_SUBNET:
        return subnet_impl::free((subnet_impl *)impl);

    case IMPL_OBJ_ID_VNIC:
        return vnic_impl::free((vnic_impl *)impl);

    case IMPL_OBJ_ID_MAPPING:
        return mapping_impl::free((mapping_impl *)impl);

    case IMPL_OBJ_ID_ROUTE_TABLE:
        return route_table_impl::free((route_table_impl *)impl);

    case IMPL_OBJ_ID_SECURITY_POLICY:
        return security_policy_impl::free((security_policy_impl *)impl);

    case IMPL_OBJ_ID_SVC_MAPPING:
        return svc_mapping_impl::free((svc_mapping_impl *)impl);

    case IMPL_OBJ_ID_MIRROR_SESSION:
        return mirror_impl::free((mirror_impl *)impl);

    case IMPL_OBJ_ID_NEXTHOP:
        return nexthop_impl::free((nexthop_impl *)impl);

    case IMPL_OBJ_ID_NEXTHOP_GROUP:
        return nexthop_group_impl::free((nexthop_group_impl *)impl);

    case IMPL_OBJ_ID_POLICER:
        return policer_impl::free((policer_impl *)impl);

    case IMPL_OBJ_ID_DHCP_POLICY:
        return dhcp_policy_impl::free((dhcp_policy_impl *)impl);

    default:
        break;
    }
    return SDK_RET_INVALID_OP;
}

impl_base *
impl_base::build(impl_obj_id_t obj_id, void *key, api_base *api_obj) {
    switch (obj_id) {
    case IMPL_OBJ_ID_MAPPING:
        return mapping_impl::build((pds_mapping_key_t *)key,
                                   (mapping_entry *)api_obj);

    case IMPL_OBJ_ID_SVC_MAPPING:
        return svc_mapping_impl::build((pds_svc_mapping_key_t *)key,
                                       (svc_mapping *)api_obj);

    default:
        break;
    }
    return NULL;
}

void
impl_base::soft_delete(impl_obj_id_t obj_id, impl_base *impl) {
    switch(obj_id) {
    case IMPL_OBJ_ID_MAPPING:
        mapping_impl::soft_delete((mapping_impl *)impl);
        break;

    case IMPL_OBJ_ID_SVC_MAPPING:
        svc_mapping_impl::soft_delete((svc_mapping_impl *)impl);
        break;

    default:
        PDS_TRACE_ERR("Non-statless obj %u can't be soft deleted\n", obj_id);
        break;
    }
}

void
impl_base::destroy(impl_obj_id_t obj_id, impl_base *impl) {
    switch (obj_id) {
    case IMPL_OBJ_ID_DEVICE:
        return device_impl::destroy((device_impl *)impl);

    case IMPL_OBJ_ID_TEP:
        return tep_impl::destroy((tep_impl *)impl);

    case IMPL_OBJ_ID_IF:
        return if_impl::destroy((if_impl *)impl);

    case IMPL_OBJ_ID_VPC:
        return vpc_impl::destroy((vpc_impl *)impl);

    case IMPL_OBJ_ID_SUBNET:
        return subnet_impl::destroy((subnet_impl *)impl);

    case IMPL_OBJ_ID_VNIC:
        return vnic_impl::destroy((vnic_impl *)impl);

    case IMPL_OBJ_ID_MAPPING:
        return mapping_impl::destroy((mapping_impl *)impl);

    case IMPL_OBJ_ID_ROUTE_TABLE:
        return route_table_impl::destroy((route_table_impl *)impl);

    case IMPL_OBJ_ID_SECURITY_POLICY:
        return security_policy_impl::destroy((security_policy_impl *)impl);

    case IMPL_OBJ_ID_MIRROR_SESSION:
        return mirror_impl::destroy((mirror_impl *)impl);

    case IMPL_OBJ_ID_NEXTHOP:
        return nexthop_impl::destroy((nexthop_impl *)impl);

    case IMPL_OBJ_ID_NEXTHOP_GROUP:
        return nexthop_group_impl::destroy((nexthop_group_impl *)impl);

    case IMPL_OBJ_ID_SVC_MAPPING:
        return svc_mapping_impl::destroy((svc_mapping_impl *)impl);

    case IMPL_OBJ_ID_POLICER:
        return policer_impl::destroy((policer_impl *)impl);

    case IMPL_OBJ_ID_DHCP_POLICY:
        return dhcp_policy_impl::destroy((dhcp_policy_impl *)impl);

    default:
        break;
    }
}

}    // namespace impl
}    // namespace api
