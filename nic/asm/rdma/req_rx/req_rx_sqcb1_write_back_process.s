#include "req_rx.h"
#include "sqcb.h"

struct req_rx_phv_t p;
struct req_rx_sqcb1_write_back_process_k_t k;
struct sqcb1_t d;

%%

.align
req_rx_sqcb1_write_back_process:
    tblwr          d.rrq_in_progress, k.args.rrq_in_progress
    tblwr          d.rrqwqe_cur_sge_id, k.args.cur_sge_id
    tblwr          d.rrqwqe_cur_sge_offset, k.args.cur_sge_offset
    tblwr          d.e_rsp_psn, k.args.e_rsp_psn
    seq            c1, k.args.incr_nxt_to_go_token_id, 1
    tblmincri.c1   d.nxt_to_go_token_id, SIZEOF_TOKEN_ID_BITS, 1
    seq            c1, k.args.last, 1
    tblmincri.c1   RRQ_C_INDEX, d.log_rrq_size, 1 

    bbne           k.args.post_bktrack, 1, end
    nop            // Branch Delay Slot

post_bktrack_ring:
     // get DMA cmd entry based on dma_cmd_index
    DMA_CMD_STATIC_BASE_GET(r6, REQ_RX_DMA_CMD_START_FLIT_ID, REQ_RX_DMA_CMD_BKTRACK_DB)

    // dma_cmd - bktrack_ring db data
    PREPARE_DOORBELL_INC_PINDEX(k.global.lif, k.global.qtype, k.global.qid, SQ_BKTRACK_RING_ID, r1, r2)
    phvwr          p.db_data2, r2.dx
    DMA_HBM_PHV2MEM_SETUP(r6, db_data2, db_data2, r1)
    DMA_SET_WR_FENCE(DMA_CMD_PHV2MEM_T, r6)
    seq            c1, k.args.dma_cmd_eop, 1
    DMA_SET_END_OF_CMDS_C(DMA_CMD_PHV2MEM_T, r6, c1)


end:
     CAPRI_SET_TABLE_3_VALID(0)

     nop.e
     nop

