// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

#ifndef __HAL_MEM_HPP__
#define __HAL_MEM_HPP__

#include "nic/include/base.h"
#include "nic/include/mtrack.hpp"

namespace hal {

// HAL memory slabs
typedef enum hal_slab_e {
    HAL_SLAB_NONE                        = 0,
    HAL_SLAB_PI_MIN                      = 1,
    HAL_SLAB_HANDLE = HAL_SLAB_PI_MIN,
    HAL_SLAB_HANDLE_HT_ENTRY             = 2,
    HAL_SLAB_HANDLE_LIST_ENTRY           = 3,
    HAL_SLAB_HANDLE_ID_HT_ENTRY          = 4,
    HAL_SLAB_HANDLE_ID_LIST_ENTRY        = 5,
    HAL_SLAB_VRF                         = 6,
    HAL_SLAB_NETWORK                     = 7,
    HAL_SLAB_L2SEG                       = 8,
    HAL_SLAB_MC_ENTRY                    = 9,
    HAL_SLAB_LIF                         = 10,
    HAL_SLAB_IF                          = 11,
    HAL_SLAB_ENIC_L2SEG_ENTRY            = 12,
    HAL_SLAB_PORT                        = 13,
    HAL_SLAB_EP                          = 14,
    HAL_SLAB_EP_IP_ENTRY                 = 15,
    HAL_SLAB_EP_L3_ENTRY                 = 16,
    HAL_SLAB_FLOW                        = 17,
    HAL_SLAB_SESSION                     = 18,
    HAL_SLAB_SECURITY_PROFILE            = 19,
    HAL_SLAB_TLSCB                       = 20,
    HAL_SLAB_TCPCB                       = 21,
    HAL_SLAB_QOS_CLASS                   = 22,
    HAL_SLAB_ACL                         = 23,
    HAL_SLAB_WRING                       = 24,
    HAL_SLAB_PROXY                       = 25,
    HAL_SLAB_PROXY_FLOW_INFO             = 26,
    HAL_SLAB_IPSECCB                     = 27,
    HAL_SLAB_L4LB                        = 28,
    HAL_SLAB_CPUCB                       = 29,
    HAL_SLAB_RAWRCB                      = 30,
    HAL_SLAB_RAWCCB                      = 31,
    HAL_SLAB_PROXYRCB                    = 32,
    HAL_SLAB_PROXYCCB                    = 33,
    HAL_SLAB_NWSEC_POLICY                = 34,
    HAL_SLAB_NWSEC_GROUP                 = 35,
    HAL_SLAB_NWSEC_POLICY_RULES          = 36,
    HAL_SLAB_NWSEC_POLICY_CFG            = 37,
    HAL_SLAB_NWSEC_POLICY_SVC            = 38,
    HAL_SLAB_NWSEC_POLICY_APPID          = 39,
    HAL_SLAB_NWSEC_POLICY_EP_INFO        = 40,
    HAL_SLAB_DOS_POLICY                  = 41,
    HAL_SLAB_DOS_POLICY_SG_LIST          = 42,
    HAL_SLAB_LKLSHIM_FLOWDB              = 43,
    HAL_SLAB_LKLSHIM_LSOCKDB             = 44,
    HAL_SLAB_LKLSHIM_LSOCKS              = 45,
    HAL_SLAB_ARP_LEARN                   = 46,
    HAL_SLAB_DHCP_LEARN                  = 47,
    HAL_SLAB_EVENT_MAP                   = 48,
    HAL_SLAB_EVENT_MAP_LISTENER          = 49,
    HAL_SLAB_EVENT_LISTENER              = 50,
    HAL_SLAB_FTP_ALG_APPSESS             = 51,
    HAL_SLAB_FTP_ALG_L4SESS              = 52,
    HAL_SLAB_FTP_ALG_FTPINFO             = 53,
    HAL_SLAB_COPP                        = 54,
    HAL_SLAB_TFTP_ALG_APPSESS            = 55,
    HAL_SLAB_TFTP_ALG_L4SESS             = 56,
    HAL_SLAB_TFTP_ALG_TFTPINFO           = 57,
    HAL_SLAB_RPC_ALG_APPSESS             = 58,
    HAL_SLAB_RPC_ALG_L4SESS              = 59,
    HAL_SLAB_RPC_ALG_RPCINFO             = 60,
    HAL_SLAB_DNS_ALG_APPSESS             = 61,
    HAL_SLAB_DNS_ALG_L4SESS              = 62,
    HAL_SLAB_DNS_ALG_DNSINFO             = 63,
    HAL_SLAB_PI_MAX                      = 64,  // NOTE: MUST be last PI slab id

    // PD Slabs
    HAL_SLAB_PD_MIN                      = 1000,
    HAL_SLAB_VRF_PD = HAL_SLAB_PD_MIN,
    HAL_SLAB_L2SEG_PD                    = 1001,
    HAL_SLAB_MC_ENTRY_PD                 = 1002,
    HAL_SLAB_LIF_PD                      = 1003,
    HAL_SLAB_UPLINKIF_PD                 = 1004,
    HAL_SLAB_UPLINKPC_PD                 = 1005,
    HAL_SLAB_ENICIF_PD                   = 1006,
    HAL_SLAB_IF_L2SEG_PD                 = 1007,
    HAL_SLAB_TUNNELIF_PD                 = 1008,
    HAL_SLAB_CPUIF_PD                    = 1009,
    HAL_SLAB_DOS_POLICY_PD               = 1010,
    HAL_SLAB_SECURITY_PROFILE_PD         = 1011,
    HAL_SLAB_EP_PD                       = 1012,
    HAL_SLAB_EP_IP_ENTRY_PD              = 1013,
    HAL_SLAB_SESSION_PD                  = 1014,
    HAL_SLAB_TLSCB_PD                    = 1015,
    HAL_SLAB_TCPCB_PD                    = 1016,
    HAL_SLAB_QOS_CLASS_PD                = 1017,
    HAL_SLAB_ACL_PD                      = 1018,
    HAL_SLAB_WRING_PD                    = 1019,
    HAL_SLAB_IPSECCB_PD                  = 1020,
    HAL_SLAB_IPSECCB_DECRYPT_PD          = 1021,
    HAL_SLAB_L4LB_PD                     = 1022,
    HAL_SLAB_RW_PD                       = 1023,
    HAL_SLAB_TUNNEL_RW_PD                = 1024,
    HAL_SLAB_CPUCB_PD                    = 1025,
    HAL_SLAB_CPUPKT_PD                   = 1026,
    HAL_SLAB_RAWRCB_PD                   = 1027,
    HAL_SLAB_RAWCCB_PD                   = 1028,
    HAL_SLAB_PROXYRCB_PD                 = 1029,
    HAL_SLAB_PROXYCCB_PD                 = 1030,
    HAL_SLAB_CPUPKT_QINST_INFO_PD        = 1031,
    HAL_SLAB_PORT_PD                     = 1032,
    HAL_SLAB_COPP_PD                     = 1033,
    HAL_SLAB_APP_REDIR_IF_PD             = 1034,
    HAL_SLAB_PD_MAX                      = 1035,   // NOTE: MUST be last PD slab id
    HAL_SLAB_RSVD                        = 1036,  // all non-delay delete slabs can use this
    HAL_SLAB_MAX                         = 1037,
    HAL_SLAB_ALL = 0xFFFFFFFF,     // reserved and shouldn't be used
} hal_slab_t;

enum {
    HAL_MEM_ALLOC_NONE,
    HAL_MEM_ALLOC_INFRA,
    HAL_MEM_ALLOC_LIB_HT,
    HAL_MEM_ALLOC_LIB_SLAB,
    HAL_MEM_ALLOC_LIB_BITMAP,
    HAL_MEM_ALLOC_LIB_TWHEEL,
    HAL_MEM_ALLOC_LIB_THREAD,
    HAL_MEM_ALOC_LIB_PT,
    HAL_MEM_ALLOC_LIB_INDEXER,
    HAL_MEM_ALLOC_IF,
    HAL_MEM_ALLOC_L2,
    HAL_MEM_ALLOC_L3,
    HAL_MEM_ALLOC_EP,
    HAL_MEM_ALLOC_SFW,
    HAL_MEM_ALLOC_L4LB,
    HAL_MEM_ALLOC_FLOW,
    HAL_MEM_ALLOC_PD,
    HAL_MEM_ALLOC_LIB_ACL_TCAM,
    HAL_MEM_ALLOC_DLLIST,
    HAL_MEM_ALLOC_EVENT_MGR,
    HAL_MEM_ALLOC_CATALOG,
    HAL_MEM_ALLOC_DEBUG_CLI,
    HAL_MEM_ALLOC_ALG,
    HAL_MEM_ALLOC_APPID_INFO,
    HAL_MEM_ALLOC_PLUGIN_MANAGER,
    HAL_MEM_ALLOC_FTE,
    HAL_MEM_ALLOC_BLOCK_LIST,
    HAL_MEM_ALLOC_BLOCK_LIST_NODE,
    HAL_MEM_ALLOC_DIRECT_MAP,
    HAL_MEM_ALLOC_DIRECT_MAP_DATA,
    HAL_MEM_ALLOC_DIRECT_MAP_SW_DATA,
    HAL_MEM_ALLOC_DIRECT_MAP_HW_DATA,
    HAL_MEM_ALLOC_DIRECT_MAP_INDEXER,
    HAL_MEM_ALLOC_DIRECT_MAP_STATS,
    HAL_MEM_ALLOC_TCAM_ENTRY,
    HAL_MEM_ALLOC_TCAM,
    HAL_MEM_ALLOC_TCAM_INDEXER,
    HAL_MEM_ALLOC_TCAM_STATS,
    HAL_MEM_ALLOC_ACL_TCAM_STATS,
    HAL_MEM_ALLOC_HASH_ENTRY,
    HAL_MEM_ALLOC_HASH_ENTRY_KEY,
    HAL_MEM_ALLOC_HASH_ENTRY_DATA,
    HAL_MEM_ALLOC_HASH_HW_KEY_INS,
    HAL_MEM_ALLOC_HASH_SW_KEY_MASK_INS,
    HAL_MEM_ALLOC_HASH_HW_KEY_UPD,
    HAL_MEM_ALLOC_HASH_HW_KEY_DEPGM,
    HAL_MEM_ALLOC_HASH,
    HAL_MEM_ALLOC_HASH_STATS,
    HAL_MEM_ALLOC_MET_REPL_ENTRY,
    HAL_MEM_ALLOC_MET_REPL_ENTRY_DATA,
    HAL_MEM_ALLOC_MET_REPL_TABLE_ENTRY,
    HAL_MEM_ALLOC_MET_REPL_LIST,
    HAL_MEM_ALLOC_MET,
    HAL_MEM_ALLOC_MET_REPL_TABLE_INDEXER,
    HAL_MEM_ALLOC_MET_STATS,
    HAL_MEM_ALLOC_FLOW_ENTRY,
    HAL_MEM_ALLOC_FLOW_ENTRY_KEY,
    HAL_MEM_ALLOC_FLOW_ENTRY_DATA,
    HAL_MEM_ALLOC_ENTIRE_FLOW_ENTRY_DATA,
    HAL_MEM_ALLOC_FLOW_ENTRY_HW_KEY,
    HAL_MEM_ALLOC_FLOW_SPINE_ENTRY,
    HAL_MEM_ALLOC_FLOW_SPINE_ENTRY_SW_KEY,
    HAL_MEM_ALLOC_FLOW_SPINE_ENTRY_HW_KEY,
    HAL_MEM_ALLOC_FLOW_HINT_GROUP,
    HAL_MEM_ALLOC_FLOW_TABLE_ENTRY,
    HAL_MEM_ALLOC_FLOW_COLL_INDEXER,
    HAL_MEM_ALLOC_FLOW_ENTRY_INDEXER,
    HAL_MEM_ALLOC_FLOW_STATS,
    HAL_MEM_ALLOC_FLOW_HW_KEY,
    HAL_MEM_ALLOC_INP_PROP_KEY_MASK,
    HAL_MEM_ALLOC_OTHER,
    HAL_MEM_ALLOC_ALL = 0xFFFFFFFF,    // reserved and shouldn't be used
};

//------------------------------------------------------------------------------
// use these memory allocation and free macros when one off allocation is
// needed, otherwise use slabs
//------------------------------------------------------------------------------
#define HAL_MALLOC(alloc_id, size)    (hal::utils::g_hal_mem_mgr.mtrack_alloc(static_cast<uint32_t>(alloc_id), false, size, __FUNCTION__, __LINE__))
#define HAL_CALLOC(alloc_id, size)    (hal::utils::g_hal_mem_mgr.mtrack_alloc(static_cast<uint32_t>(alloc_id), true, size, __FUNCTION__, __LINE__))
#define HAL_FREE(alloc_id, ptr)       (hal::utils::g_hal_mem_mgr.mtrack_free(static_cast<uint32_t>(alloc_id), ptr, __FUNCTION__, __LINE__))

hal_ret_t free_to_slab(hal_slab_t slab_id, void *elem);

namespace pd {

hal_ret_t free_to_slab(hal_slab_t slab_id, void *elem);

}    // namespace pd

}    // namespace hal

#endif    // __HAL_MEM_HPP__

