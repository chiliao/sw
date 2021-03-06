//-----------------------------------------------------------------------------
// {C} Copyright 2017 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------

#include "nic/hal/plugins/network/net_plugin.hpp"
#include "nic/hal/plugins/cfg/nw/interface_api.hpp"
#include "nic/include/pd_api.hpp"
#include <net/ethernet.h>

namespace hal {
namespace plugins {
namespace network {

static inline hal_ret_t
update_tunnel_info(fte::ctx_t&ctx)
{
    fte::flow_update_t flowupd = {type: fte::FLOWUPD_HEADER_PUSH};


    if (ctx.dif() == NULL || ctx.dif()->if_type != intf::IF_TYPE_TUNNEL) {
        return HAL_RET_OK;
    }

    // TODO(goli) set appropriate header fields
    HEADER_SET_FLD(flowupd.header_push, ether, dmac, ether_addr{});
    HEADER_SET_FLD(flowupd.header_push, ether, smac, ether_addr{});
    HEADER_SET_FLD(flowupd.header_push, ipv4, sip, ipv4_addr_t{});
    HEADER_SET_FLD(flowupd.header_push, ipv4, dip, ipv4_addr_t{});

    switch (ctx.dl2seg()->wire_encap.type) {
    case types::encapType::ENCAP_TYPE_VXLAN:
        HEADER_SET_FLD(flowupd.header_push, vxlan, vrf_id, ctx.dl2seg()->wire_encap.val);
        break;
    default:
        return HAL_RET_INVALID_ARG;
    }

    return ctx.update_flow(flowupd);
}

fte::pipeline_action_t
tunnel_exec(fte::ctx_t& ctx)
{
    hal_ret_t ret;

    ret = update_tunnel_info(ctx);
    if (ret != HAL_RET_OK) {
        ctx.set_feature_status(ret);
        return fte::PIPELINE_END;
    }

    return fte::PIPELINE_CONTINUE;
}

} // namespace network
} // namespace plugins
} // namespace hal
