#ifndef __RESP_RX_H
#define __RESP_RX_H
#include "capri.h"
#include "types.h"
#include "resp_rx_args.h"
#include "INGRESS_p.h"
#include "common_phv.h"

#define RESP_RX_MAX_DMA_CMDS        16
#define RESP_RX_DMA_CMD_PYLD_BASE   3
#define RESP_RX_DMA_CMD_RSQWQE      4
#define RESP_RX_DMA_CMD_RSQ_DB      5

#define RESP_RX_DMA_CMD_START       0
#define RESP_RX_DMA_CMD_ACK         0
#define RESP_RX_DMA_CMD_CQ          (RESP_RX_MAX_DMA_CMDS - 3)
#define RESP_RX_DMA_CMD_IMMDT_AS_DBELL RESP_RX_DMA_CMD_CQ //mutually exclusive

#define RESP_RX_DMA_CMD_START_FLIT_ID   8 // flits 8-11 are used for dma cmds

//TODO: put ack_info.aeth, ack_info.psn adjacent to each other in PHV and also
//      adjacent to each other in rqcb1, in right order. This will eliminate
//      one DMA instruction
#define RESP_RX_POST_ACK_INFO_TO_TXDMA(_dma_base_r, _rqcb1_addr_r, _tmp_r, \
                                       _lif, _qtype, _qid, \
                                       _db_addr_r, _db_data_r) \
    add         _tmp_r, _rqcb1_addr_r, FIELD_OFFSET(rqcb1_t, ack_nak_psn); \
    DMA_HBM_PHV2MEM_SETUP(_dma_base_r, ack_info.psn, ack_info.aeth.msn, _tmp_r); \
    DMA_NEXT_CMD_I_BASE_GET(_dma_base_r, 1); \
    PREPARE_DOORBELL_INC_PINDEX(_lif, _qtype, _qid, ACK_NAK_RING_ID, _db_addr_r, _db_data_r);\
    phvwr       p.db_data1, _db_data_r.dx; \
    DMA_HBM_PHV2MEM_SETUP(_dma_base_r, db_data1, db_data1, _db_addr_r); \
    DMA_SET_WR_FENCE(DMA_CMD_PHV2MEM_T, _dma_base_r); \
    
#define RESP_RX_POST_ACK_INFO_TO_TXDMA_NO_DB(_dma_base_r, _rqcb1_addr_r, _tmp_r) \
    add         _tmp_r, _rqcb1_addr_r, FIELD_OFFSET(rqcb1_t, ack_nak_psn); \
    DMA_HBM_PHV2MEM_SETUP(_dma_base_r, ack_info.psn, ack_info.aeth.msn, _tmp_r); \

#define RESP_RX_POST_IMMDT_AS_DOORBELL(_dma_base_r, _tmp_r, \
                                       _lif, _qtype, _qid, \
                                       _db_addr_r, _db_data_r) \
    PREPARE_DOORBELL_INC_PINDEX(_lif, _qtype, _qid, 0, _db_addr_r, _db_data_r);\
    phvwr       p.immdt_as_dbell_data, _db_data_r.dx; \
    DMA_HBM_PHV2MEM_SETUP(_dma_base_r, immdt_as_dbell_data, immdt_as_dbell_data, _db_addr_r); \
    DMA_SET_WR_FENCE(DMA_CMD_PHV2MEM_T, _dma_base_r); \
       
#define RESP_RX_RING_ACK_NAK_DB(_dma_base_r,  \
                                _lif, _qtype, _qid, \
                                _db_addr_r, _db_data_r) \
    PREPARE_DOORBELL_INC_PINDEX(_lif, _qtype, _qid, ACK_NAK_RING_ID, _db_addr_r, _db_data_r);\
    phvwr       p.db_data1, _db_data_r.dx; \
    DMA_HBM_PHV2MEM_SETUP(_dma_base_r, db_data1, db_data1, _db_addr_r);


#define RSQ_EVAL_MIDDLE     0
#define RSQ_WALK_FORWARD    1
#define RSQ_WALK_BACKWARD   2
    
// phv 
struct resp_rx_phv_t {
    // dma commands (flit 8 - 11)

    // scratch (flit 7):

    // size: 38 =  2 + 8 + 8 + 2 + 5 + 4 + 1 + 8
    eq_int_num: 16;
    db_data1: 64;
    db_data2: 64;
    adjust_rsq_c_index: 16;
    struct ack_info_t ack_info;
    struct eqwqe_t eqwqe;
    my_token_id: 8;
    immdt_as_dbell_data: 64;

    // scratch (flit 6)
    // size: 64  = 32 + 32
    struct rsqwqe_t rsqwqe;
    struct cqwqe_t cqwqe;

    // common rx (flit 0 - 5)
    struct phv_ common;
};

struct resp_rx_phv_global_t {
    struct phv_global_common_t common;
};

struct resp_rx_to_stage_backtrack_info_t {
    va: 64;
    r_key: 32;
    len: 32;
};

struct resp_rx_to_stage_t {
    union {
        struct resp_rx_to_stage_backtrack_info_t backtrack;
        pad: TO_STAGE_DATA_WIDTH;
    };
};

// stage to stage argument structures

// 20
struct resp_rx_rqcb_to_pt_info_t {
    in_progress: 1;
    page_seg_offset: 3;
    tbl_id: 3;
    cache: 1;
    page_offset: 16;
    remaining_payload_bytes: 16;
    inv_r_key: 32;
    pad: 88;
};

struct resp_rx_rqpt_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rqcb_to_pt_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

// 20
struct resp_rx_rqcb_to_wqe_info_t {
    //rqcb1
    in_progress:1;
    cache:1;
    current_sge_id: 6;
    remaining_payload_bytes: 16;
    //rqcb2
    curr_wqe_ptr: 64;
    inv_r_key: 32;
    current_sge_offset: 32;
    //computed
    num_valid_sges: 6;
    tbl_id: 2;
};

struct resp_rx_rqwqe_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rqcb_to_wqe_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_key_info_t {
    va: 64;
    //keep len...tbl_id in the same order
    //aligning with resp_rx_lkey_to_pt_info_t
    len: 16;
    dma_cmd_start_index: 8;
    tbl_id: 8;
    acc_ctrl: 8;
    key_id: 8;
    dma_cmdeop: 1;
    cq_dma_cmd_index: 7;
    inv_r_key: 32;
    nak_code: 8;
    //tightly packed for 160 bits
};

struct resp_rx_key_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_key_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_rqcb0_write_back_info_t {
    in_progress: 8;
    incr_nxt_to_go_token_id: 1;
    incr_c_index: 1;
    tbl_id: 3;
    cache: 1;
    do_not_invalidate_tbl: 1;
    pad: 145;
};

struct resp_rx_rqcb0_write_back_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rqcb0_write_back_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_rqcb1_write_back_info_t {
    curr_wqe_ptr: 64;
    current_sge_offset: 32;
    current_sge_id: 8;
    update_num_sges: 1;
    update_wqe_ptr: 1;
    num_sges: 8;
    pad: 46;
};

struct resp_rx_rqcb1_write_back_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rqcb1_write_back_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_lkey_to_pt_info_t {
    pt_offset: 32;
    //keep pt_bytes...sge_index in the same order
    //aligning with resp_rx_key_info_t
    pt_bytes: 16;
    dma_cmd_start_index: 8;
    sge_index: 8;
    log_page_size: 5;
    dma_cmdeop: 1;
    rsvd: 2;
    pad: 88;
};

struct resp_rx_ptseg_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_lkey_to_pt_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_compl_or_inv_rkey_info_t {
    r_key:  32;
    dma_cmd_index: 8;
    tbl_id: 3;
    rsvd: 5;
    pad: 112;
};

struct resp_rx_compl_or_inv_rkey_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_compl_or_inv_rkey_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_rqcb_to_cq_info_t {
    tbl_id: 3;
    rsvd: 5;
    dma_cmd_index: 8;
    pad: 144;
};

struct resp_rx_cqcb_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rqcb_to_cq_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_cqcb_to_pt_info_t {
    page_offset: 16;
    page_seg_offset: 8;
    dma_cmd_index: 8;
    eq_id: 24;
    cq_id: 24;
    arm: 1;
    tbl_id: 3;
    rsvd: 4;
    pad: 72;
};

struct resp_rx_cqpt_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_cqcb_to_pt_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_cqcb_to_eq_info_t {
    tbl_id: 3;
    rsvd: 5;
    dma_cmd_index: 8;  
    cq_id: 24;
    pad: 120;
};

struct resp_rx_eqcb_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_cqcb_to_eq_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_rqcb_to_rqcb1_info_t {
    in_progress: 1;
    rsvd: 7;
    remaining_payload_bytes: 32;
    inv_r_key: 32;
    pad: 88;
};

struct resp_rx_rqcb1_in_progress_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rqcb_to_rqcb1_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_rqcb_to_write_rkey_info_t {
    va: 64;
    len: 32;
    r_key: 32;
    remaining_payload_bytes: 16;
    load_reth: 1;
    incr_c_index: 1;
    rsvd1: 6;
    rsvd2: 8;
};

struct resp_rx_write_dummy_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rqcb_to_write_rkey_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_inv_rkey_info_t {
    tbl_id: 8;
    key_id: 8;
    pad: 144;
};

struct resp_rx_inv_rkey_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_inv_rkey_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_rsq_backtrack_info_t {                         
    log_pmtu: 5;
    read_or_atomic:1;                                             
    walk:2;
    hi_index: 16;
    lo_index: 16;
    index: 8;
    log_rsq_size: 5;
    rsvd: 3;
    search_psn:24;
    rsq_base_addr: 32;
    pad: 48;
};

struct resp_rx_rsq_backtrack_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rsq_backtrack_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};

//20
struct resp_rx_rsq_backtrack_adjust_info_t {
    adjust_c_index: 1;
    adjust_p_index: 1;
    rsvd: 6;
    index: 8;
    pad: 144;
};

struct resp_rx_rsq_backtrack_adjust_process_k_t {
    struct capri_intrinsic_raw_k_t intrinsic;
    struct resp_rx_rsq_backtrack_adjust_info_t args;
    struct resp_rx_to_stage_t to_stage;
    struct phv_global_common_t global;
};


#endif //__RESP_RX_H
