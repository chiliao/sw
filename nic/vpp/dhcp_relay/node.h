//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//

#ifndef __VPP_DHCP_RELAY_NODE_H__
#define __VPP_DHCP_RELAY_NODE_H__

#include <vlib/vlib.h>
#include <vnet/ip/ip.h>
#include <nic/vpp/infra/pkt.h>
#include <nic/vpp/infra/utils.h>

// clfy node related defines
#define foreach_dhcp_relay_clfy_counter                     \
    _(TO_SERVER, "Sent to Server " )                        \
    _(TO_CLIENT, "Sent to Client" )                         \

#define foreach_dhcp_relay_clfy_next                        \
        _(TO_SERVER, "dhcp-proxy-to-server" )               \
        _(TO_CLIENT, "dhcp-proxy-to-client")                \

// server header node related defines
#define foreach_dhcp_relay_svr_hdr_counter                  \
    _(TX, "Sent on Netwrok interface " )                    \

#define foreach_dhcp_relay_svr_hdr_next                     \
    _(INTF_OUT, "interface-tx" )                            \
    _(DROP, "error-drop")                                   \
// client header node related defines
#define foreach_dhcp_relay_client_hdr_counter               \
    _(TX, "Sent on host interface " )                       \

#define foreach_dhcp_relay_client_hdr_next                  \
    _(INTF_OUT, "interface-tx" )                            \
    _(DROP, "error-drop")                                   \


// clfy node related defines
typedef enum
{
#define _(n,s) DHCP_RELAY_CLFY_COUNTER_##n,
    foreach_dhcp_relay_clfy_counter
#undef _
    DHCP_RELAY_CLFY_COUNTER_LAST,
} dhcp_relay_clfy_counter_t;

typedef enum
{
#define _(n,s) PDS_DHCP_RELAY_CLFY_NEXT_##n,
    foreach_dhcp_relay_clfy_next
#undef _
    PDS_DHCP_RELAY_CLFY_N_NEXT,
} dhcp_relay_clfy_next_t;

typedef struct dhcp_relay_clfy_trace_s {
    uint16_t lif;    
} dhcp_relay_clfy_trace_t;

// server header node related defines
typedef enum
{
#define _(n,s) DHCP_RELAY_SVR_HDR_COUNTER_##n,
    foreach_dhcp_relay_svr_hdr_counter
#undef _
    DHCP_RELAY_SVR_HDR_COUNTER_LAST,
} dhcp_relay_svr_hdr_counter_t;

typedef enum
{
#define _(n,s) PDS_DHCP_RELAY_SVR_HDR_NEXT_##n,
    foreach_dhcp_relay_svr_hdr_next
#undef _
    PDS_DHCP_RELAY_SVR_HDR_N_NEXT,
} dhcp_relay_svr_hdr_next_t;

typedef struct dhcp_relay_svr_hdr_trace_s {
    uint16_t next_hop;
} dhcp_relay_svr_hdr_trace_t;

// client header node related defines
typedef enum
{
#define _(n,s) DHCP_RELAY_CLIENT_HDR_COUNTER_##n,
    foreach_dhcp_relay_client_hdr_counter
#undef _
    DHCP_RELAY_CLIENT_HDR_COUNTER_LAST,
} dhcp_relay_client_hdr_counter_t;

typedef enum
{
#define _(n,s) PDS_DHCP_RELAY_CLIENT_HDR_NEXT_##n,
    foreach_dhcp_relay_client_hdr_next
#undef _
    PDS_DHCP_RELAY_CLIENT_HDR_N_NEXT,
} dhcp_relay_client_hdr_next_t;

typedef struct dhcp_relay_client_hdr_trace_s {
    mac_addr_t client_mac;
} dhcp_relay_client_hdr_trace_t;

#endif    // __VPP_DHCP_RELAY_NODE_H__