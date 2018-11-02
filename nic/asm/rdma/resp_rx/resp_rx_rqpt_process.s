#include "capri.h"
#include "resp_rx.h"
#include "rqcb.h"
#include "common_phv.h"

struct resp_rx_phv_t p;
struct resp_rx_s2_t0_k k;

#define INFO_OUT1_P t0_s2s_rqcb_to_wqe_info

#define GLOBAL_FLAGS r7

#define IN_P t0_s2s_launch_rqpt_to_rqpt_info
#define K_PAGE_OFFSET CAPRI_KEY_RANGE(IN_P, page_offset_sbit0_ebit7, page_offset_sbit8_ebit15)
#define K_PAGE_SEG_OFFSET CAPRI_KEY_FIELD(IN_P, page_seg_offset)
#define K_RQ_IN_HBM CAPRI_KEY_FIELD(IN_P, rq_in_hbm)
#define K_HBM_WQE_PTR CAPRI_KEY_RANGE(IN_P, hbm_wqe_ptr_sbit0_ebit7, hbm_wqe_ptr_sbit56_ebit63)

#define IN_TO_S_P to_s2_ext_hdr_info
#define K_SEND_SGE_OPT CAPRI_KEY_FIELD(IN_TO_S_P, send_sge_opt)

%%
    .param  resp_rx_rqwqe_wrid_process
    .param  resp_rx_rqwqe_process
    .param  resp_rx_rqwqe_opt_process

.align
resp_rx_rqpt_process:
    bbeq    K_RQ_IN_HBM, 1, skip_rqpt_process
    // populate r3 with hbm_wqe_ptr. If q is not in hbm, r3 would be overwritten with right value
    add     r3, r0, K_HBM_WQE_PTR   //BD Slot

    //page_addr_p = (u64 *) (d_p + sizeof(u64) * rqcb_to_pt_info_p->page_seg_offset);

    //big-endian
    sub     r3, (HBM_NUM_PT_ENTRIES_PER_CACHE_LINE-1), K_PAGE_SEG_OFFSET 
    sll     r3, r3, CAPRI_LOG_SIZEOF_U64_BITS
    //big-endian
    tblrdp.dx  r3, r3, 0, CAPRI_SIZEOF_U64_BITS
    or         r3, r3, 1, 63
    or         r3, r3, K_GLOBAL_LIF, 52

    // wqe_p = (void *)(*page_addr_p + rqcb_to_pt_info_p->page_offset);
    add     r3, r3, K_PAGE_OFFSET
    // now r3 has wqe_p to load

skip_rqpt_process:

    CAPRI_RESET_TABLE_0_ARG()

    phvwrpair   CAPRI_PHV_FIELD(INFO_OUT1_P, remaining_payload_bytes), \
                CAPRI_KEY_RANGE(IN_P, remaining_payload_bytes_sbit0_ebit7, remaining_payload_bytes_sbit8_ebit15), \
                CAPRI_PHV_FIELD(INFO_OUT1_P, curr_wqe_ptr), r3

    phvwrpair   CAPRI_PHV_FIELD(INFO_OUT1_P, dma_cmd_index), \
                RESP_RX_DMA_CMD_PYLD_BASE, \
                CAPRI_PHV_FIELD(INFO_OUT1_P, log_pmtu), \
                CAPRI_KEY_FIELD(IN_P, log_pmtu)

    bbeq        K_SEND_SGE_OPT, 1, rqwqe_opt
    add         GLOBAL_FLAGS, r0, K_GLOBAL_FLAGS    //BD Slot

    ARE_ALL_FLAGS_SET(c1, GLOBAL_FLAGS, RESP_RX_FLAG_WRITE|RESP_RX_FLAG_IMMDT) 

    CAPRI_NEXT_TABLE0_READ_PC_CE(CAPRI_TABLE_LOCK_DIS, CAPRI_TABLE_SIZE_512_BITS, \
                                 resp_rx_rqwqe_wrid_process, resp_rx_rqwqe_process, r3, c1)

rqwqe_opt:
    CAPRI_NEXT_TABLE0_READ_PC_E(CAPRI_TABLE_LOCK_DIS, CAPRI_TABLE_SIZE_512_BITS, resp_rx_rqwqe_opt_process, r3)
