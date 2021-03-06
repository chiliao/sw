#include "capri.h"
#include "req_tx.h"
#include "sqcb.h"
#include "defines.h"

struct req_tx_phv_t p;
struct req_tx_s4_t0_k k;
struct key_entry_aligned_t d;

#define INFO_OUT_T struct req_tx_lkey_to_ptseg_info_t
#define INFO_OUT_P t0_s2s_lkey_to_ptseg

#define IN_P t0_s2s_sge_to_lkey_info
#define IN_TO_S_P to_s4_dcqcn_bind_mw_info
//#define IN_TO_S_P to_s1_dcqcn_bind_mw_info
#define TO_S7_STATS_INFO_P to_s7_stats_info


#define K_SGE_VA CAPRI_KEY_RANGE(IN_P, sge_va_sbit0_ebit7, sge_va_sbit56_ebit63)
#define K_SGE_BYTES CAPRI_KEY_RANGE(IN_P, sge_bytes_sbit0_ebit7, sge_bytes_sbit8_ebit15)
#define K_SGE_INDEX CAPRI_KEY_FIELD(IN_P, sge_index)
#define K_SGE_DMA_CMD_START_INDEX CAPRI_KEY_FIELD(IN_P, dma_cmd_start_index)

#define K_PD CAPRI_KEY_FIELD(IN_TO_S_P, header_template_addr_or_pd)

%%
    .param    req_tx_sqptseg_process

.align
req_tx_sqlkey_process:
     mfspr        r1, spr_mpuid
     seq          c1, r1[4:2], STAGE_4
     bcf          [!c1], bubble_to_next_stage
     seq          c3, K_SGE_INDEX, 0 // Branch Delay Slot

     //If Reserved LKEY is used, but QP doesn't have privileged operations enabled
     bbeq         CAPRI_KEY_FIELD(IN_P, rsvd_key_err), 1, rsvd_lkey_error

     // check if lkey-state is valid.
     seq          c1, d.state, KEY_STATE_VALID  // Branch Delay Slot
     bcf          [!c1], invalid_region

     seq          c2, K_PD, d.pd // Branch Delay Slot
     bcf          [!c2], pd_check_failure

     // If zbva, va = reth_va + MR_base_va else va = reth_va
     IS_ANY_FLAG_SET_B(c1, d.acc_ctrl, ACC_CTRL_ZERO_BASED) // Branch Delay Slot
     add.c1       r1, K_SGE_VA, d.base_va
     add.!c1      r1, K_SGE_VA, r0

     // if ((lkey_info_p->sge_va < lkey_p->base_va) ||
     //     ((lkey_info_p->sge_va + lkey_info_p->sge_bytes) > (lkey_p->base_va + lkey_p->len)))
     slt          c1, r1, d.base_va // Branch Delay Slot
     add          r3, d.len, d.base_va
     sslt         c2, r3, r1, K_SGE_BYTES
     bcf          [c1|c2], access_violation
     seq          c4, d.is_phy_addr, 1  // Branch delay slot

     bcf          [c4], contig_mr
     // my_pt_base_addr = (void *)(hbm_addr_get(PHV_GLOBAL_PT_BASE_ADDR_GET()) +
     //                            (lkey_p->pt_base * sizeof(u64))
     PT_BASE_ADDR_GET2(r4) // Branch delay slot
     add          r3, r4, d.pt_base, CAPRI_LOG_SIZEOF_U64

     // pt_seg_size = HBM_NUM_PT_ENTRIES_PER_CACHE_LINE * lkey_info_p->page_size
     add          r4, d.log_page_size, LOG_HBM_NUM_PT_ENTRIES_PER_CACHELINE

     // lkey_p->base_va % pt_seg_size
     add          r5, d.base_va, r0
     mincr        r5, r4, r0

     // transfer_offset = lkey_info_p->sge_va - lkey_p->base_va + (lkey_p->base_va % pt_seg_size)
     sub          r2, r1, d.base_va
     add          r2, r2, r5

     // if ((transfer_bytes + (transfer_offset % pt_seg_size)) <= pt_seg_size)
     add          r5, r2, 0
     mincr        r5, r4, 0
     add          r5, r5, K_SGE_BYTES
     //  pt_seg_size = 1 << (LOG_PAGE_SIZE + HBM_MUM_PT_ENTRIES_PER_CACHE_LINE)
     sllv         r6, 1, r4
     ble          r5, r6, pt_aligned_access
     add          r5, r2, r0 // Branch Delay Slot

     // Unaligned PT access
pt_unaligned_access:
     // pt_offset = transfer_offset % lkey_info_p->page_size
     mincr        r5, d.log_page_size, r0

     // pt_seg_p = (u64 *)my_pt_base_addr + (transfer_offset / lkey_info_p->log_page_size)
     srlv         r2, r2, d.log_page_size
     b            set_arg
     add          r3, r3, r2, CAPRI_LOG_SIZEOF_U64 // Branch Delay Slot

     // Aligned PT access
pt_aligned_access:
     // pt_offset = transfer_offset % pt_seg_size
     mincr        r5, r4, r0

     // pt_seg_p = (u64 *)my_pt_base_addr + ((transfer_offset / pt_seg_size) * HBM_NUM_PT_ENTRIES_PER_CACHE_LINE)
     srlv         r2, r2, r4
     add          r3, r3, r2, (CAPRI_LOG_SIZEOF_U64 + LOG_HBM_NUM_PT_ENTRIES_PER_CACHELINE)

set_arg:
     CAPRI_GET_TABLE_0_OR_1_ARG(req_tx_phv_t, r7, c3)
     CAPRI_SET_FIELD(r7, INFO_OUT_T, pt_offset, r5)
     CAPRI_SET_FIELD(r7, INFO_OUT_T, log_page_size, d.log_page_size)
     CAPRI_SET_FIELD(r7, INFO_OUT_T, host_addr, d.host_addr)
     //CAPRI_SET_FIELD(r7, INFO_OUT_T, pt_bytes, K_SGE_BYTES)
     //CAPRI_SET_FIELD(r7, INFO_OUT_T, dma_cmd_start_index, K_DMA_CMD_START_INDEX)
     //CAPRI_SET_FIELD(r7, INFO_OUT_T, sge_index, K_SGE_INDEX)
     CAPRI_SET_FIELD_RANGE(r7, INFO_OUT_T, pt_bytes, sge_index, CAPRI_KEY_RANGE(IN_P, sge_bytes_sbit0_ebit7, sge_index))

     CAPRI_GET_TABLE_0_OR_1_K_NO_VALID(req_tx_phv_t, r7, c3)
     CAPRI_NEXT_TABLE_I_READ_PC(r7, CAPRI_TABLE_LOCK_DIS, CAPRI_TABLE_SIZE_512_BITS, req_tx_sqptseg_process, r3)

exit:
     nop.e
     nop

bubble_to_next_stage:
    seq           c1, r1[4:2], STAGE_3
    bcf           [!c1], exit

    CAPRI_GET_TABLE_0_OR_1_K_NO_VALID(req_tx_phv_t, r7, c3) // Branch Delay Slot
    CAPRI_NEXT_TABLE_I_READ_SET_SIZE(r7, CAPRI_TABLE_LOCK_EN, CAPRI_TABLE_SIZE_512_BITS)

    nop.e
    nop

rsvd_lkey_error: 
    phvwrpair    CAPRI_PHV_FIELD(TO_S7_STATS_INFO_P, qp_err_disabled), 1, \
                 CAPRI_PHV_FIELD(TO_S7_STATS_INFO_P, qp_err_dis_lkey_rsvd_lkey), 1 //BD Slot
    b            error_completion
    phvwrpair    p.{rdma_feedback.completion.status, rdma_feedback.completion.error}, (CQ_STATUS_MEM_MGMT_OPER_ERR << 1 | 1), \
                 p.{rdma_feedback.completion.lif_cqe_error_id_vld, rdma_feedback.completion.lif_error_id_vld, rdma_feedback.completion.lif_error_id}, \
                       ((1 << 5) | (1 << 4) | LIF_STATS_RDMA_REQ_STAT(LIF_STATS_REQ_TX_MEMORY_MGMT_ERR_OFFSET)) //BD SLot

pd_check_failure:
    b            local_prot_error
    phvwrpair    CAPRI_PHV_FIELD(TO_S7_STATS_INFO_P, qp_err_disabled), 1, \
                 CAPRI_PHV_FIELD(TO_S7_STATS_INFO_P, qp_err_dis_lkey_inv_pd), 1 //BD Slot
invalid_region:
    b            local_prot_error
    phvwrpair    CAPRI_PHV_FIELD(TO_S7_STATS_INFO_P, qp_err_disabled), 1, \
                 CAPRI_PHV_FIELD(TO_S7_STATS_INFO_P, qp_err_dis_lkey_inv_state), 1 //BD Slot

local_prot_error:

    b            error_completion
    phvwrpair    p.{rdma_feedback.completion.status, rdma_feedback.completion.error}, (CQ_STATUS_LOCAL_PROT_ERR << 1 | 1), \
                 p.{rdma_feedback.completion.lif_cqe_error_id_vld, rdma_feedback.completion.lif_error_id_vld, rdma_feedback.completion.lif_error_id}, \
                       ((1 << 5) | (1 << 4) | LIF_STATS_RDMA_REQ_STAT(LIF_STATS_REQ_TX_LOCAL_ACCESS_ERR_OFFSET)) //BD SLot

access_violation:
    phvwrpair    CAPRI_PHV_FIELD(TO_S7_STATS_INFO_P, qp_err_disabled), 1, \
                 CAPRI_PHV_FIELD(TO_S7_STATS_INFO_P, qp_err_dis_lkey_access_violation), 1 
    phvwrpair    p.{rdma_feedback.completion.status, rdma_feedback.completion.error}, (CQ_STATUS_LOCAL_ACC_ERR << 1 | 1), \
                 p.{rdma_feedback.completion.lif_cqe_error_id_vld, rdma_feedback.completion.lif_error_id_vld, rdma_feedback.completion.lif_error_id}, \
                       ((1 << 5) | (1 << 4) | LIF_STATS_RDMA_REQ_STAT(LIF_STATS_REQ_TX_LOCAL_ACCESS_ERR_OFFSET))
    //fall through

error_completion:
    add          r1, K_SGE_INDEX, r0
    CAPRI_SET_TABLE_I_VALID(r1, 0)

    phvwr.e        CAPRI_PHV_FIELD(phv_global_common, _error_disable_qp),  1
    nop

contig_mr:
    
    // get DMA cmd entry based on dma_cmd_index
    DMA_CMD_I_BASE_GET(r3, r6, REQ_TX_DMA_CMD_START_FLIT_ID, K_SGE_DMA_CMD_START_INDEX)

    sub         r2, K_SGE_VA, d.base_va
    add         r2, r2, d.phy_base_addr
    // setup mem2pkt cmd to transfer data from host memory to pkt payload
    // it is assumed to be host_addr all the time
    DMA_MEM2PKT_SETUP(r3, c0, K_SGE_BYTES, r2)

    seq         c2, K_SGE_INDEX, 0 
    CAPRI_SET_TABLE_0_VALID_CE(c2, 0)
    CAPRI_SET_TABLE_1_VALID_C(!c2, 0)
    
