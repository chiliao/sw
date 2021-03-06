#include "ingress.h"
#include "INGRESS_p.h"
#include "INGRESS_s1_t0_nvme_sesspostdgst_tx_k.h"
#include "capri.h"
#include "nvme_common.h"

struct phv_ p;
struct s1_t0_nvme_sesspostdgst_tx_k_ k;
struct s1_t0_nvme_sesspostdgst_tx_sess_wqe_process_d d;

#define PDU_CTXT1_P   r1

%%
    .param    nvme_sesspostdgst_tx_pdu_ctxt1_process
    .param    nvme_tx_pdu_context_base

.align
nvme_sesspostdgst_tx_sess_wqe_process:

    // calculate pdu_ctxt address
    // nvme_tx_pdu_context_base + (pduid << sizeof(pdu_context))
    addui          PDU_CTXT1_P, r0, hiword(nvme_tx_pdu_context_base)
    addi           PDU_CTXT1_P, PDU_CTXT1_P, loword(nvme_tx_pdu_context_base)
    add            PDU_CTXT1_P, PDU_CTXT1_P, d.pduid, LOG_PDU_CTXT_SIZE
    add            PDU_CTXT1_P, PDU_CTXT1_P, NVME_PDU_CTXT0_SIZE

    // Pass cid to sessdgsttx_cb_writeback to load tso_process
    // starting from first tcp_page
    phvwr          p.to_s5_info_pduid, d.pduid
    phvwr          p.to_s6_info_pduid, d.pduid

    // load cmd_ctx process to get nlb, number of pages
    CAPRI_NEXT_TABLE0_READ_PC_E(CAPRI_TABLE_LOCK_DIS,
                                CAPRI_TABLE_SIZE_512_BITS,
                                nvme_sesspostdgst_tx_pdu_ctxt1_process,
                                PDU_CTXT1_P) // Exit Slot
