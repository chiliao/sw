#include "req_rx.h"
#include "sqcb.h"

struct req_rx_phv_t p;
//this routine is invoked on s2_t0 and s2_t1
struct req_rx_s3_t0_k k;
struct key_entry_aligned_t d;

#define IN_P t0_s2s_rrqsge_to_lkey_info

#define K_SGE_VA CAPRI_KEY_RANGE(IN_P, sge_va_sbit0_ebit7, sge_va_sbit56_ebit63)
#define K_SGE_BYTES CAPRI_KEY_RANGE(IN_P, sge_bytes_sbit0_ebit7, sge_bytes_sbit8_ebit15)
#define K_SGE_INDEX CAPRI_KEY_RANGE(IN_P, sge_index_sbit0_ebit1, sge_index_sbit2_ebit7)
#define K_DMA_CMD_START_INDEX CAPRI_KEY_FIELD(IN_P, dma_cmd_start_index)
#define ATOMICAETH_ORIG_DATA common.rdma_bth_aeth_atomicaeth_atomicaeth_orig_data


%%

.align
req_rx_rrqlkey_rsvd_lkey_process:

    // get DMA cmd entry based on dma_cmd_index
    DMA_CMD_I_BASE_GET(r7, r6, REQ_RX_DMA_CMD_START_FLIT_ID, K_DMA_CMD_START_INDEX)

    bbne          CAPRI_KEY_FIELD(IN_P, is_atomic), 1, pkt2mem
    // set cmdeop in the last pkt dma cmd 
    seq        c1, CAPRI_KEY_FIELD(IN_P, dma_cmd_eop), 1  // BD Slot
    // setup phv2mem cmd to transfer data from atomicaeth to host memory
    DMA_HOST_PHV2MEM_SETUP(r7, ATOMICAETH_ORIG_DATA, ATOMICAETH_ORIG_DATA, K_SGE_VA)
    b          continue
    nop

pkt2mem:
    DMA_HOST_PKT2MEM_SETUP(r7, K_SGE_BYTES, K_SGE_VA)

continue:
    DMA_SET_END_OF_CMDS_C(DMA_CMD_PKT2MEM_T, r7, c1)

    DMA_CMD_STATIC_BASE_GET(r7, REQ_RX_DMA_CMD_START_FLIT_ID, REQ_RX_DMA_CMD_LSN_OR_REXMIT_PSN)
    DMA_UNSET_END_OF_CMDS_C(DMA_CMD_PHV2MEM_T, r7, c1)

    add        r1, K_SGE_INDEX, r0
    CAPRI_SET_TABLE_I_VALID(r1, 0)

    nop.e
    nop
