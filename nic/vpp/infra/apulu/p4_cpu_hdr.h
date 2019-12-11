//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
// This file contains p4 header interface for rx and tx packets

#ifndef __VPP_INFRA_APULU_P4_CPU_HDR_H__
#define __VPP_INFRA_APULU_P4_CPU_HDR_H__

#include <nic/apollo/p4/include/apulu_defines.h>
#include <vppinfra/clib.h>

#define VPP_CPU_FLAGS_VLAN_VALID           APULU_CPU_FLAGS_VLAN_VALID
#define VPP_CPU_FLAGS_IPV4_1_VALID         APULU_CPU_FLAGS_IPV4_1_VALID
#define VPP_CPU_FLAGS_IPV6_1_VALID         APULU_CPU_FLAGS_IPV6_1_VALID
#define VPP_CPU_FLAGS_ETH_2_VALID          APULU_CPU_FLAGS_ETH_2_VALID
#define VPP_CPU_FLAGS_IPV4_2_VALID         APULU_CPU_FLAGS_IPV4_2_VALID
#define VPP_CPU_FLAGS_IPV6_2_VALID         APULU_CPU_FLAGS_IPV6_2_VALID
#define VPP_ARM_TO_P4_HDR_SZ               APULU_ARM_TO_P4_HDR_SZ
#define VPP_P4_TO_ARM_HDR_SZ               APULU_P4_TO_ARM_HDR_SZ

// Meta received from P4 for rx packet
typedef CLIB_PACKED(struct p4_rx_cpu_hdr_s {
    uint16_t   packet_len;
    uint16_t   flags;
    uint16_t   ingress_bd_id;
    uint32_t   flow_hash;

    // offsets
    uint8_t  l2_offset;
    uint8_t  l3_offset;
    uint8_t  l4_offset;
    uint8_t  l2_inner_offset;
    uint8_t  l3_inner_offset;
    uint8_t  l4_inner_offset;
    uint8_t  payload_offset;

    uint16_t lif;
    uint16_t egress_bd_id;
    uint16_t service_xlate_id;
    uint16_t mapping_xlate_id;
    uint16_t tx_meter_id;
    uint16_t nexthop_id;
    union {
        uint8_t flags_octet;
        struct {
#if __BYTE_ORDER == __BIG_ENDIAN
            uint8_t pad                 : 5;
            uint8_t nexthop_type        : 2;
            uint8_t drop                : 1;
#else
            uint8_t drop                : 1;
            uint8_t nexthop_type        : 2;
            uint8_t pad                 : 5;
#endif
        };
    };
}) p4_rx_cpu_hdr_t;

// Meta sent to P4 for tx packet
typedef CLIB_PACKED(struct p4_tx_cpu_hdr_s {
    union {
        uint16_t lif_pad;
        struct {
#if __BYTE_ORDER == __BIG_ENDIAN
            uint16_t pad : 5;
            uint16_t lif : 11;
#else
            uint16_t lif : 11;
            uint16_t pad : 5;
#endif
        };
    };
}) p4_tx_cpu_hdr_t;

#endif     // __VPP_INFRA_APULU_P4_CPU_HDR_H__
