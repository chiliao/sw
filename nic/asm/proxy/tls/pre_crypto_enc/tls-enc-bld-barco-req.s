/*
 * 	Construct the barco request in this stage
 * Stage 4, Table 0
 */

#include "tls-constants.h"
#include "tls-phv.h"
#include "tls-shared-state.h"
#include "tls-macros.h"
#include "tls-table.h"
#include "tls_common.h"
#include "ingress.h"
#include "INGRESS_p.h"
        
struct tx_table_s4_t0_k k                  ;
struct phv_ p   ;
struct tx_table_s4_t0_d d       ;
        
%%
            .param      tls_enc_queue_brq_process
            .param      tls_enc_queue_brq_mpp_process   
#           .param      BRQ_QPCB_BASE
        
tls_enc_bld_barco_req_process:        

table_read_QUEUE_BRQ:
    CAPRI_SET_DEBUG_STAGE4_7(p.to_s6_debug_stage4_7_thread, CAPRI_MPU_STAGE_4, CAPRI_MPU_TABLE_0)

    /* Fill the barco request in the phv to be DMAed later into BRQ slot */
    addi        r2, r0, PKT_DESC_AOL_OFFSET
    add         r1, r2, k.{to_s4_idesc}

    add         r3, r2, k.{to_s4_odesc}
    phvwrpair   p.barco_desc_input_list_address, r1.dx, \
                p.barco_desc_output_list_address, r3.dx

    phvwri      p.to_s6_enc_requests, 1

    /*
     * When set to use random IV from barco DRBG, the explicit-IV field in the request will be
     * set by another program which reads the random number memory, so we dont set it here.
     */
    smeqb       c1, k.tls_global_phv_debug_dol, TLS_DDOL_EXPLICIT_IV_USE_RANDOM, TLS_DDOL_EXPLICIT_IV_USE_RANDOM

    /*
     * The PHV for IV/nonce for GCM/CCM are different format made as unions. We'll build the phv
     * accordingly.
     */
    bbeq        k.tls_global_phv_flags_do_pre_ccm_enc, 1, tls_enc_bld_ccm_hdr_phv
    nop
        
    phvwr.!c1   p.crypto_iv_explicit_iv, d.u.tls_bld_brq4_d.sequence_no
    CAPRI_OPERAND_DEBUG(d.u.tls_bld_brq4_d.sequence_no)

    /* Setup AAD */
    /* AAD length already setup in Stage 2, Table 3 */
    phvwr.!c1   p.s2_s5_t0_phv_aad_seq_num, d.u.tls_bld_brq4_d.sequence_no
    addi        r1, r0, ((NTLS_RECORD_DATA << 16) | (NTLS_TLS_1_2_MAJOR << 8) | (NTLS_TLS_1_2_MINOR))
    phvwr       p.{s2_s5_t0_phv_aad_type...s2_s5_t0_phv_aad_version_minor}, r1

    tbladd.!c1  d.u.tls_bld_brq4_d.sequence_no, 1

    b           tls_enc_bld_barco_desc
    addi        r1, r0, NTLS_AAD_SIZE

tls_enc_bld_ccm_hdr_phv:
	
    /*
     * Setup PHV for CCM here.
     */
    phvwri      p.ccm_header_with_aad_B_0_flags, TLS_AES_CCM_HDR_B0_FLAGS
    phvwri      p.ccm_header_with_aad_B_1_aad_size, 13
    phvwrpair   p.ccm_header_with_aad_B_0_nonce_explicit_iv, d.u.tls_bld_brq4_d.sequence_no, \
                p.ccm_header_with_aad_B_1_aad_seq_num, d.u.tls_bld_brq4_d.sequence_no
    tbladd      d.u.tls_bld_brq4_d.sequence_no, 1

    addi        r1, r0, ((NTLS_RECORD_DATA << 16) | (NTLS_TLS_1_2_MAJOR << 8) | (NTLS_TLS_1_2_MINOR))
    phvwr       p.{ccm_header_with_aad_B_1_aad_type...ccm_header_with_aad_B_1_aad_version_minor}, r1
    //p.ccm_header_with_aad_aad_length is already updated in tls-enc-read-pkt-descr.s
    phvwri      p.ccm_header_with_aad_B_1_zero_pad, 0

    addi        r1, r0, TLS_AES_CCM_HEADER_SIZE

tls_enc_bld_barco_desc:
    phvwr       p.barco_desc_header_size, r1.wx


    /* address will be in r4 */
    addi        r4, r0, CAPRI_DOORBELL_ADDR(0, DB_IDX_UPD_PIDX_SET, DB_SCHED_UPD_SET, 0, LIF_TLS)

    /*
     * data will be in r3
     *
     * We maintain a shadow-copy of the BSQ PI in qstate CB which we'll increment
     * and use as 'PIDX_SET' instead of using the 'PIDX_INC' auto-increment feature
     * of the doorbell, for better performance.
     */
    /* sw_bsq_pi - does not need to be a masked increment */
    tbladd.f    d.{u.tls_bld_brq4_d.sw_bsq_pi}.hx, 1
    add         r6, r0, d.{u.tls_bld_brq4_d.sw_bsq_pi}.hx
    CAPRI_RING_DOORBELL_DATA(0, k.tls_global_phv_fid, TLS_SCHED_RING_BSQ, r6)
    phvwrpair   p.barco_desc_doorbell_address, r4.dx, \
                p.barco_desc_doorbell_data, r3.dx

    /* The barco-command[31:24] is checked for GCM/CCM/CBC. endian-swapped */
    bbeq        k.tls_global_phv_flags_do_pre_mpp_enc, 1, tls_enc_queue_to_brq_mpp_ring
    nop

    /*
     * NOTE: The next stage program tls_enc_queue_brq_process does NOT need a table-read anymore,
     *	     as we have already read a shadow-copy of Barco PI to be written to.
     */
    CAPRI_NEXT_TABLE_READ_NO_TABLE_LKUP(0, tls_enc_queue_brq_process);
    nop.e
    nop

tls_enc_queue_to_brq_mpp_ring:

    /*
     * If we're doing BYPASS-BARCO, we'll use the 'PIDX_SET' for the BSQ PI using
     * the sw shadow copy in the doorbell.
     */
    smeqb       c4, k.tls_global_phv_debug_dol, TLS_DDOL_BYPASS_BARCO, TLS_DDOL_BYPASS_BARCO
    bcf         [!c4], tls_enc_bsq_doorbell_skip
    nop
    CAPRI_DMA_CMD_RING_DOORBELL_SET_PI(dma_cmd_dbell_dma_cmd, LIF_TLS, 0, k.tls_global_phv_fid, TLS_SCHED_RING_BSQ, r6,
                                       crypto_iv_explicit_iv)
    CAPRI_DMA_CMD_STOP_FENCE(dma_cmd_dbell_dma_cmd)

tls_enc_bsq_doorbell_skip:	

    /*
     * NOTE: The next stage program tls_enc_queue_brq_mpp_process does NOT need a table-read anymore,
     *	     as we have already read a shadow-copy of Barco PI to be written to.
     */
    CAPRI_NEXT_TABLE_READ_NO_TABLE_LKUP(0, tls_enc_queue_brq_mpp_process)
        nop.e
        nop
