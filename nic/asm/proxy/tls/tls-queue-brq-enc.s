/*
 * 	Implements the writing of request to BRQ
 */

#include "tls-constants.h"
#include "tls-phv.h"
#include "tls-shared-state.h"
#include "tls-macros.h"
#include "tls-table.h"
#include "ingress.h"
#include "INGRESS_p.h"        

struct phv_ p	;
struct tx_table_s5_t0_k k	;
struct tx_table_s5_t0_tls_queue_brq5_d d ;

	
%%
	.param		BRQ_BASE
tls_queue_brq_enc_process:
	/*   brq.odesc->data_len = brq.idesc->data_len + sizeof(tls_hdr_t); */
dma_cmd_enc_data_len:
	/*   brq.odesc->data_len = tlsp->cur_tls_data_len; */
	add		    r5, r0, k.to_s5_odesc
	addi		r5, r5, NIC_DESC_DATA_LEN_OFFSET
	phvwr		p.dma_cmd1_dma_cmd_addr, r5
	/* Fill the data len */
	add		    r1, k.to_s5_cur_tls_data_len, TLS_HDR_SIZE
	phvwr		p.to_s5_cur_tls_data_len, k.to_s5_cur_tls_data_len

    phvwri      p.dma_cmd1_dma_cmd_phv_start_addr, CAPRI_PHV_START_OFFSET(to_s5_cur_tls_data_len)
	phvwri		p.dma_cmd1_dma_cmd_phv_end_addr, CAPRI_PHV_END_OFFSET(to_s5_cur_tls_data_len)

	phvwri		p.dma_cmd1_dma_cmd_type, CAPRI_DMA_COMMAND_PHV_TO_MEM

	/*
	    SET_DESC_ENTRY(brq.odesc, 0, 
		 md->opage, 
		 NIC_PAGE_HEADROOM + sizeof(tls_hdr_t), 
		 brq.odesc->data_len);
         */
dma_cmd_enc_desc_entry0:
	add		    r5, r0, k.to_s5_odesc
	addi		r5, r5, NIC_DESC_ENTRY0_OFFSET
	phvwr		p.dma_cmd2_dma_cmd_addr, r5

	phvwr		p.aol_A0, k.to_s5_odesc

	addi		r5, r0, NIC_PAGE_HEADROOM
	addi		r5, r5, TLS_HDR_SIZE
	phvwr		p.aol_O0, r5

	/* r1 = d.cur_tls_data_len + TLS_HDR_SIZE */
	phvwr		p.aol_L0, r1

    phvwri      p.dma_cmd2_dma_cmd_phv_start_addr, CAPRI_PHV_START_OFFSET(aol_L0)
	phvwri		p.dma_cmd2_dma_cmd_phv_end_addr, CAPRI_PHV_END_OFFSET(aol_A0)

    phvwri		p.dma_cmd2_dma_cmd_type, CAPRI_DMA_COMMAND_PHV_TO_MEM


dma_cmd_enc_tls_hdr:
	/* tlsh = (tls_hdr_t *)(u64)(md->opage + NIC_PAGE_HEADROOM); */
	add		    r5, r0, k.to_s5_opage
	addi		r5, r5, NIC_PAGE_HEADROOM
	phvwr		p.dma_cmd3_dma_cmd_addr, r5

	/*
	  tlsh->type = NTLS_RECORD_DATA;
          tlsh->version_major = NTLS_TLS_1_2_MAJOR;
          tlsh->version_minor = NTLS_TLS_1_2_MINOR;
          tlsh->len = brq.idesc->data_len;
	 */
	
	

	phvwr		p.tls_global_phv_tls_hdr_len, k.to_s5_cur_tls_data_len

    phvwri      p.dma_cmd3_dma_cmd_phv_start_addr, CAPRI_PHV_START_OFFSET(tls_global_phv_tls_hdr_type)
	phvwri		p.dma_cmd3_dma_cmd_phv_end_addr, CAPRI_PHV_END_OFFSET(tls_global_phv_tls_hdr_len)

    phvwri		p.dma_cmd3_dma_cmd_type, CAPRI_DMA_COMMAND_PHV_TO_MEM

        

dma_cmd_enc_brq_slot:
	add		    r5, r0, d.pi_0
	tbladd		d.pi_0, 1
    sll		    r5, r5, NIC_BRQ_ENTRY_SIZE_SHIFT
	/* Set the DMA_WRITE CMD for BRQ slot */
	addi		r1, r0, BRQ_BASE
	add		    r1, r1, r5

	phvwr		p.dma_cmd4_dma_cmd_addr,r1
	/* Fill the barco request */

    phvwri      p.dma_cmd4_dma_cmd_phv_start_addr, CAPRI_PHV_START_OFFSET(barco_desc_input_list_address)
	phvwri		p.dma_cmd4_dma_cmd_phv_end_addr, CAPRI_PHV_END_OFFSET(barco_desc_status_address)

    phvwri		p.dma_cmd4_dma_cmd_type, CAPRI_DMA_COMMAND_PHV_TO_MEM

    phvwri      p.dma_cmd4_dma_cmd_eop, 1
    phvwri      p.dma_cmd4_dma_cmd_wr_fence, 1

tls_queue_brq_process_done:
	nop.e
	nop
