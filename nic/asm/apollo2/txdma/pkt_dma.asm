#include "apollo_txdma.h"
#include "INGRESS_p.h"
#include "ingress.h"

struct phv_             p;
struct pkt_dma_k        k;

%%

pkt_dma:
    phvwr       p.capri_txdma_intr_dma_cmd_ptr, \
                    CAPRI_PHV_START_OFFSET(intrinsic_dma_dma_cmd_phv2pkt_pad)/16
    sub         r1, k.{p4_to_txdma_header_payload_len_sbit0_ebit5, \
                       p4_to_txdma_header_payload_len_sbit6_ebit13}, APOLLO_I2E_HDR_SZ
    phvwr       p.capri_p4_intr_packet_len, r1

    CAPRI_DMA_CMD_PHV2PKT_SETUP3(intrinsic_dma_dma_cmd,
                                 capri_intr_tm_iport, \
                                 capri_txdma_intr_txdma_rsv, \
                                 predicate_header_txdma_drop_event, \
                                 predicate_header_direction, \
                                 txdma_to_p4e_header_p4plus_app_id, \
                                 txdma_to_p4e_header_vcn_id)
    // MEM2PKT macro uses phvwrpair and not able to use here
    // CAPRI_DMA_CMD_MEM2PKT_SETUP(payload_dma_dma_cmd, k.txdma_control_payload_addr, \
    //      k.{p4_to_txdma_header_payload_len_sbit0_ebit5...p4_to_txdma_header_payload_len_sbit6_ebit13})
    phvwr       p.payload_dma_dma_cmd_addr, k.txdma_control_payload_addr

    .assert((offsetof(k, p4_to_txdma_header_payload_len_sbit0_ebit5) - \
            offsetof(k, p4_to_txdma_header_payload_len_sbit6_ebit13)) == 8)
    phvwr       p.payload_dma_dma_cmd_size, \
                    k.{p4_to_txdma_header_payload_len_sbit0_ebit5, \
                       p4_to_txdma_header_payload_len_sbit6_ebit13}
    phvwri      p.{payload_dma_dma_cmd_eop...payload_dma_dma_cmd_type}, CAPRI_DMA_COMMAND_MEM_TO_PKT
    phvwr       p.payload_dma_dma_pkt_eop, 1

    // mem2pkt has an implicit fence. all subsequent dma is blocked
    .assert((offsetof(k, capri_intr_lif_sbit0_ebit2) - \
             offsetof(k, capri_intr_lif_sbit3_ebit10)) == 8)
    CAPRI_RING_DOORBELL_ADDR(0, DB_IDX_UPD_CIDX_SET, DB_SCHED_UPD_EVAL, 0, \
                    k.{capri_intr_lif_sbit0_ebit2...capri_intr_lif_sbit3_ebit10})
    CAPRI_RING_DOORBELL_DATA(0, k.capri_txdma_intr_qid, 0, k.txdma_control_cindex)
    phvwr       p.{doorbell_data_pid...doorbell_data_index}, r3.dx
    CAPRI_DMA_CMD_PHV2MEM_SETUP(doorbell_ci_update_dma_cmd, \
                                r4, \
                                doorbell_data_pid, \
                                doorbell_data_index)
    CAPRI_DMA_CMD_PHV2MEM_SETUP_STOP(rxdma_ci_update_dma_cmd, \
                                    k.{txdma_control_rxdma_cindex_addr_sbit0_ebit31... \
                                       txdma_control_rxdma_cindex_addr_sbit32_ebit33}, \
                                    txdma_control_cindex, \
                                    txdma_control_cindex)
    phvwr       p.capri_intr_tm_iport, TM_PORT_DMA
    phvwr.e     p.capri_intr_tm_oport, TM_PORT_EGRESS
    phvwr       p.txdma_to_p4e_header_vcn_id, k.p4_to_txdma_header_vcn_id

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
pkt_dma_error:
    nop.e
    nop
