#include "apollo.h"
#include "ingress.h"
#include "INGRESS_p.h"

struct vnic_tx_stats_k   k;
struct vnic_tx_stats_d   d;
struct phv_     p;

%%

vnic_tx_stats:
    tbladd          d.vnic_tx_stats_d.in_packets, 1
    tbladd          d.vnic_tx_stats_d.in_bytes, \
                    k.{capri_p4_intrinsic_packet_len_sbit0_ebit5, \
                       capri_p4_intrinsic_packet_len_sbit6_ebit13}
    seq             c1, k.service_header_local_ip_mapping_done, FALSE
    seq.!c1         c1, k.service_header_flow_done, FALSE
    bcf             [c1], recirc_packet
    phvwr.!c1       p.capri_intrinsic_tm_oport, TM_PORT_DMA
    phvwr           p.capri_rxdma_intrinsic_rx_splitter_offset, \
                        (CAPRI_GLOBAL_INTRINSIC_HDR_SZ + \
                         CAPRI_RXDMA_INTRINSIC_HDR_SZ + \
                         APOLLO_P4_TO_RXDMA_HDR_SZ)
    phvwr           p.capri_p4_intrinsic_valid, TRUE
    phvwr           p.capri_rxdma_intrinsic_valid, TRUE
    phvwr           p.p4_to_rxdma_header_valid, TRUE
    phvwr           p.p4_to_txdma_header_valid, TRUE
    phvwr.e         p.apollo_i2e_metadata_valid, TRUE
    phvwr           p.service_header_valid, FALSE

recirc_packet:
    phvwr.e         p.capri_intrinsic_tm_oport, TM_PORT_INGRESS
    phvwr           p.service_header_valid, TRUE

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
vnic_tx_stats_error:
    nop.e
    nop
