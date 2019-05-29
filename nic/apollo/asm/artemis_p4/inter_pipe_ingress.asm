#include "artemis.h"
#include "ingress.h"
#include "INGRESS_p.h"
#include "INGRESS_inter_pipe_ingress_k.h"

struct inter_pipe_ingress_k_ k;
struct inter_pipe_ingress_d  d;
struct phv_ p;

%%

ingress_to_egress:
    /*
    phvwr           p.p4i_i2e_valid, TRUE
    phvwr           p.txdma_to_p4e_valid, TRUE
    phvwr           p.predicate_header_valid, TRUE
    phvwr           p.capri_p4_intrinsic_valid, TRUE
    bitmap ==> 0 0001 1000 1001
    */
    phvwr           p.{service_header_valid, \
                        p4plus_to_p4_vlan_valid, \
                        p4plus_to_p4_valid, \
                        capri_txdma_intrinsic_valid, \
                        p4i_i2e_valid, \
                        txdma_to_p4e_valid, \
                        p4_to_arm_valid, \
                        p4_to_p4plus_classic_nic_ip_valid, \
                        p4_to_p4plus_classic_nic_valid, \
                        predicate_header_valid, \
                        p4_to_rxdma_valid, \
                        capri_rxdma_intrinsic_valid, \
                        capri_p4_intrinsic_valid}, 0x0189
    phvwr           p.capri_intrinsic_tm_oport, TM_PORT_EGRESS
    seq             c1, k.vxlan_1_valid, TRUE
    phvwr.c1        p.{vxlan_1_valid,udp_1_valid,ipv4_1_valid,ipv6_1_valid, \
                        ctag_1_valid,ethernet_1_valid}, 0
    seq             c1, k.vxlan_2_valid, TRUE
    phvwr.e         p.predicate_header_direction, k.control_metadata_direction
    phvwr.c1        p.{vxlan_2_valid,udp_2_valid,ipv4_2_valid,ipv6_2_valid, \
                        ctag_2_valid,ethernet_2_valid}, 0


.align
ingress_to_cps:
    /*
    phvwr           p.p4_to_rxdma_valid, TRUE
    phvwr           p.capri_rxdma_intrinsic_valid, TRUE
    phvwr           p.capri_p4_intrinsic_valid, TRUE
    bitmap ==> 0 0000 0000 0111
    */
    phvwr           p.{service_header_valid, \
                        p4plus_to_p4_vlan_valid, \
                        p4plus_to_p4_valid, \
                        capri_txdma_intrinsic_valid, \
                        p4i_i2e_valid, \
                        txdma_to_p4e_valid, \
                        p4_to_arm_valid, \
                        p4_to_p4plus_classic_nic_ip_valid, \
                        p4_to_p4plus_classic_nic_valid, \
                        predicate_header_valid, \
                        p4_to_rxdma_valid, \
                        capri_rxdma_intrinsic_valid, \
                        capri_p4_intrinsic_valid}, 0x0007
    phvwr           p.capri_intrinsic_tm_oport, TM_PORT_DMA
    phvwr           p.capri_intrinsic_lif, ARTEMIS_SERVICE_LIF
    phvwr           p.capri_rxdma_intrinsic_rx_splitter_offset, \
                        (CAPRI_GLOBAL_INTRINSIC_HDR_SZ + \
                         CAPRI_RXDMA_INTRINSIC_HDR_SZ + \
                         ARTEMIS_P4_TO_RXDMA_HDR_SZ)
    phvwr.e         p.p4_to_rxdma_table3_valid, TRUE
    phvwr           p.p4_to_rxdma_direction, k.control_metadata_direction

.align
ingress_to_classic_nic:
ingress_to_arm:
    nop.e
    nop

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
inter_pipe_ingress_error:
    phvwr.e         p.capri_intrinsic_drop, 1
    nop
