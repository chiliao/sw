#include "nic/p4/common/defines.h"
#include "tcp-constants.h"
#include "tcp-shared-state.h"
#include "tcp-macros.h"
#include "tcp-table.h"
#include "tcp-phv.h"
#include "tcp_common.h"
#include "ingress.h"
#include "INGRESS_p.h"
#include "INGRESS_s1_t0_ooq_tcp_tx_k.h"

struct phv_ p;
struct s1_t0_ooq_tcp_tx_k_ k;
struct s1_t0_ooq_tcp_tx_process_next_descr_addr_d d;

%%
    .align
tcp_ooq_txdma_process_next_descr_addr:
    CAPRI_CLEAR_TABLE_VALID(0)
    CAPRI_OPERAND_DEBUG(d.descr_addr)
tcp_ooq_txdma_dma_cmds:
    phvwri          p.{p4_intr_global_tm_iport...p4_intr_global_tm_oport}, \
                        (TM_PORT_DMA << 4) | TM_PORT_DMA
    phvwri          p.p4_txdma_intr_dma_cmd_ptr, TCP_PHV_OOQ_TXDMA_COMMANDS_START

    phvwr           p.intr_rxdma_qid, k.common_phv_fid
    phvwr           p.intr_rxdma_rx_splitter_offset, \
                    (CAPRI_GLOBAL_INTRINSIC_HDR_SZ + CAPRI_RXDMA_INTRINSIC_HDR_SZ + \
                    P4PLUS_TCP_PROXY_BASE_HDR_SZ + P4PLUS_TCP_PROXY_OOQ_HDR_SZ)
    CAPRI_DMA_CMD_PHV2PKT_SETUP2(intrinsic_dma_cmd, p4_intr_global_tm_iport,
                                p4_intr_packet_len,
                                intr_rxdma_qid, intr_rxdma_rxdma_rsv)

    add             r1, r0, d.descr_addr
    CAPRI_DMA_CMD_MEM2PKT_SETUP_PKT_STOP(tcp_app_header_dma_cmd, r1, \
                        P4PLUS_TCP_PROXY_BASE_HDR_SZ + P4PLUS_TCP_PROXY_OOQ_HDR_SZ)
    nop.e
    nop
