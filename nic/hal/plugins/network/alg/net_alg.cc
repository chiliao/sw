#include "nic/hal/plugins/network/net_plugin.hpp"
#include "nic/hal/plugins/network/alg/alg_tftp.hpp"
#include "nic/hal/plugins/network/alg/alg_rpc.hpp"

namespace hal {
namespace net {

static void
init_alg_entry(fte::ctx_t& ctx)
{
    fte::alg_entry_t entry;

    memset(&entry, 0, sizeof(fte::alg_entry_t));
 
    entry.key = ctx.key();
    entry.role = ctx.role();

    ctx.set_alg_entry(entry);
}

static bool
IsAlgExistingFlow(fte::ctx_t& ctx)
{
    HAL_TRACE_DEBUG("ALG Exec role: {} alg_proto_state: {}", 
                     ctx.role(), ctx.alg_proto_state());

    return (ctx.alg_proto_state() != fte::ALG_PROTO_STATE_NONE);
}

fte::pipeline_action_t
alg_exec(fte::ctx_t& ctx)
{
    hal_ret_t          ret = HAL_RET_OK;
  
    if (ctx.protobuf_request()) {
        return fte::PIPELINE_CONTINUE;
    }
   
    HAL_TRACE_DEBUG("ALG Proto: {}", ctx.alg_proto());
    if (ctx.alg_proto() != nwsec::APP_SVC_NONE) {

        // Init the ALG entry 
        if (ctx.role() == hal::FLOW_ROLE_INITIATOR) 
            init_alg_entry(ctx);

        switch(ctx.alg_proto()) {
            case nwsec::APP_SVC_TFTP:
                ret = process_tftp_first_packet(ctx);
                break;

            case nwsec::APP_SVC_SUN_RPC:
                if (ctx.alg_proto_state() == fte::ALG_PROTO_STATE_RPC_DATA) {
                    ret = process_sunrpc_data_flow(ctx);
                } else {
                    ret = process_sunrpc_control_flow(ctx); 
                }
                break;
  
            case nwsec::APP_SVC_MSFT_RPC:
                if (ctx.alg_proto_state() == fte::ALG_PROTO_STATE_RPC_DATA) {
                    ret = process_msrpc_data_flow(ctx);
                } else {
                    ret = process_msrpc_control_flow(ctx);
                }
                break;
 
            case nwsec::APP_SVC_FTP:
                break;

            case nwsec::APP_SVC_DNS:
                break;

            default:
                break;
        };
    } else if (IsAlgExistingFlow(ctx)) {
        switch(ctx.alg_proto_state()) {
            case fte::ALG_PROTO_STATE_TFTP_RRQ:
            case fte::ALG_PROTO_STATE_TFTP_WRQ:
                ret = process_tftp(ctx);
                break;

            case fte::ALG_PROTO_STATE_RPC_INIT:
            case fte::ALG_PROTO_STATE_RPC_GETPORT:
            case fte::ALG_PROTO_STATE_RPC_DUMP:
                ret = parse_sunrpc_control_flow(ctx);
                return fte::PIPELINE_END;

            default:
                break;
        };
    }

    ctx.set_feature_status(ret);
    
    return fte::PIPELINE_CONTINUE;
}

} // namespace net
} // namespace hal
