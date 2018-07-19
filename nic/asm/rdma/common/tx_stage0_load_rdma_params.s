#include "capri.h"
#include "common_phv.h"
#include "common_txdma_actions/asm_out/INGRESS_p.h"
#include "common_txdma_actions/asm_out/ingress.h"
#include "req_tx_args.h"

struct tx_stage0_lif_params_table_k k;
struct tx_stage0_lif_params_table_d d;
struct phv_ p;

#define REQ_TX_TO_S2_T struct req_tx_to_stage_2_t

%%

tx_stage0_load_rdma_params:

    add r4, r0, k.p4_txdma_intr_qtype //BD slot
    sllv r5, 1, r4
    and r5, r5, d.u.tx_stage0_lif_rdma_params_d.rdma_en_qtype_mask
    seq c1, r5, r0
    bcf [c1], done

    add r1, r0, offsetof(struct phv_, common_global_global_data) //BD slot
    CAPRI_SET_FIELD(r1, PHV_GLOBAL_COMMON_T, pt_base_addr_page_id, d.u.tx_stage0_lif_rdma_params_d.pt_base_addr_page_id)
    CAPRI_SET_FIELD(r1, PHV_GLOBAL_COMMON_T, log_num_pt_entries, d.u.tx_stage0_lif_rdma_params_d.log_num_pt_entries)

    add r1, r0, offsetof(struct phv_, to_stage_2_to_stage_data)
    CAPRI_SET_FIELD(r1, REQ_TX_TO_S2_T, ah_base_addr_page_id, d.u.tx_stage0_lif_rdma_params_d.ah_base_addr_page_id)
    #We do not need cqcb base addr in Tx direction, hence commenting for now
    #CAPRI_SET_FIELD(r1, PHV_GLOBAL_COMMON_T, cqcb_base_addr_hi, d.u.tx_stage0_lif_rdma_params_d.cqcb_base_addr_hi)
    #CAPRI_SET_FIELD(r1, PHV_GLOBAL_COMMON_T, log_num_cq_entries, d.u.tx_stage0_lif_rdma_params_d.log_num_cq_entries)
    #CAPRI_SET_FIELD(r1, PHV_GLOBAL_COMMON_T, prefetch_pool_base_addr_page_id, d.u.tx_stage0_lif_rdma_params_d.prefetch_pool_base_addr_page_id)
    #CAPRI_SET_FIELD(r1, PHV_GLOBAL_COMMON_T, log_num_prefetch_pool_entries, d.u.tx_stage0_lif_rdma_params_d.log_num_prefetch_pool_entries)
done:
    nop.e
    nop
