#ifndef __REQ_RX_H
#define __REQ_RX_H
#include "capri.h"
#include "types.h"
#include "req_rx_args.h"
#include "INGRESS_p.h"
#include "ingress.h"
#include "common_phv.h"

#define REQ_RX_CQCB_ADDR_GET(_r, _cqid, _cqcb_base_addr_hi) \
    CQCB_ADDR_GET(_r, _cqid, _cqcb_base_addr_hi);

#define REQ_RX_EQCB_ADDR_GET(_r, _tmp_r, _eqid, _cqcb_base_addr_hi, _log_num_cq_entries) \
    EQCB_ADDR_GET(_r, _tmp_r, _eqid, _cqcb_base_addr_hi, _log_num_cq_entries);

#define REQ_RX_MAX_DMA_CMDS                20
#define REQ_RX_DMA_CMD_START_FLIT_ID       7 // flits 8-11 are used for dma cmds
#define REQ_RX_DMA_CMD_START_FLIT_CMD_ID   1
#define REQ_RX_DMA_CMD_START               1
#define REQ_RX_DMA_CMD_LSN_OR_REXMIT_PSN   1
//#define REQ_RX_DMA_CMD_SQ_DB               2
#define REQ_RX_DMA_CMD_RQ_FLUSH_DB         2
#define REQ_RX_DMA_CMD_RNR_TIMEOUT         2
#define REQ_RX_DMA_CMD_BKTRACK_DB          3
#define REQ_RX_RDMA_PAYLOAD_DMA_CMDS_START 2
#define REQ_RX_RDMA_PAYLOAD_DMA_CMDS_END   14
#define REQ_RX_DMA_CMD_SKIP_TO_EOP         (REQ_RX_MAX_DMA_CMDS - 6)
#define REQ_RX_DMA_CMD_CQ                  (REQ_RX_MAX_DMA_CMDS - 5)
#define REQ_RX_DMA_CMD_EQ                  (REQ_RX_MAX_DMA_CMDS - 4)
//wakeup dpath and EQ are mutually exclusive
#define REQ_RX_DMA_CMD_WAKEUP_DPATH        REQ_RX_DMA_CMD_EQ
#define REQ_RX_DMA_CMD_EQ_INTR             (REQ_RX_MAX_DMA_CMDS - 3)
#define REQ_RX_DMA_CMD_ASYNC_EQ            (REQ_RX_MAX_DMA_CMDS - 2)
#define REQ_RX_DMA_CMD_ASYNC_EQ_INTR       (REQ_RX_MAX_DMA_CMDS - 1)

// phv 
struct req_rx_phv_t {
    // flit 11
    dma_cmd16               : 128;
    dma_cmd17               : 128;
    dma_cmd18               : 128;
    dma_cmd19               : 128;

    // flit 10
    dma_cmd12               : 128;
    dma_cmd13               : 128;
    dma_cmd14               : 128;
    dma_cmd15               : 128;

    // flit 9
    dma_cmd8                : 128;
    dma_cmd9                : 128;
    dma_cmd10               : 128;
    dma_cmd11               : 128;

    // flit 8
    dma_cmd4                : 128;
    dma_cmd5                : 128;
    dma_cmd6                : 128;
    dma_cmd7                : 128;

    //flit 7
    rsvd4                   : 16;
    lsn                     : 24;
    rexmit_psn              : 24;
    ack_timestamp           : 48;
    err_retry_ctr           : 4;
    rnr_retry_ctr           : 4;
    rnr_timeout             : 8;
    dma_cmd1                : 128;
    dma_cmd2                : 128;
    dma_cmd3                : 128;

    //flit 6
    service                 : 4;
    flush_rq                : 1;
    state                   : 3;
    rsvd1                   : 32;
    union {
        struct {
            async_int_assert_data : 32;
            struct eqwqe_t async_eqwqe;
        };
        db_data2                : 64;
    };
    db_data1                : 64;
    union {
        struct {
            int_assert_data : 32;
            struct eqwqe_t eqwqe;
        };
        wakeup_dpath_data   : 64;
    };
    struct cqe_t cqe;
    my_token_id             : 8;

    //flit 0-5
    // common rx
    struct phv_ common;
};

struct req_rx_phv_global_t {
    struct phv_global_common_t common;
};

// stage to stage argument structures

struct req_rx_s0_t {
    lif: 11;
    qtype: 3;
    qid: 24;
    struct p4_2_p4plus_app_hdr_t app_hdr;
};

#endif //__REQ_RX_H
