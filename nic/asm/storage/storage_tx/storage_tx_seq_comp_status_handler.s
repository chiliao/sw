/*****************************************************************************
 *  seq_comp_status_handler: Store the compression status in K+I vector. Load
 *                           SGL address for next stage to do the PDMA.
 *****************************************************************************/

#include "storage_asm_defines.h"
#include "ingress.h"
#include "INGRESS_p.h"


struct s2_tbl_k k;
struct s2_tbl_seq_comp_status_handler_d d;
struct phv_ p;

/*
 * SGL rearranged to little-endian layout
 */
struct barco_sgl_le_t {
    rsvd : 64;
    link : 64;
    rsvd2: 32;
    len2 : 32;
    addr2: 64;
    rsvd1: 32;
    len1 : 32;
    addr1: 64;
    rsvd0: 32;
    len0 : 32;
    addr0: 64;
};

/*
 * Registers usage
 */
#define r_comp_data_len             r1  // compression output data length
#define r_total_len                 r2  // above length plus padding
#define r_pad_len                   r3  // padding length
#define r_pad_boundary              r4  // user specified padding boundary
#define r_num_blks                  r5  // number of hash blocks
#define r_last_sgl_p                r6  // pointer to last SGL
#define r_sgl_field_p               r7  // pointer to an SGL field

/*
 * Registers reuse, post padding calculations
 */
#define r_desc_vec_len              r1  // total descriptor vector length
#define r_last_blk_len              r2  // length of last block

#define r_status                    r7  // comp status, briefly used at beginning

%%
    .param storage_tx_seq_barco_chain_action_start
    .param storage_tx_seq_comp_sgl_handler_start

storage_tx_seq_comp_status_handler_start:
   
    // bit 15: valid bit, bits 14-12: error bits
    add         r_status, d.status, r0
    smeqh       c1, r_status, 0xf000, 0x8000
    bcf         [!c1], comp_error
   
    // AOL/SGL padding makes sense only when next_db_en is true, with the
    // assumption that the next db handler would act on the AOL/SGL.
    //
    // Note: output_data_len contains compressed data length plus header length.
    
    seq	        c3, STORAGE_KIVEC5_DATA_LEN_FROM_DESC, 1        // delay slot
    cmov        r_comp_data_len, c3, STORAGE_KIVEC5_DATA_LEN, d.output_data_len
    bbeq        STORAGE_KIVEC5_SGL_PDMA_EN, 1, sgl_pdma_xfer
    phvwr	p.storage_kivec5_data_len, r_comp_data_len      // delay slot
   
    // Preliminary padding calculations:
    // r_num_blks = (r_comp_data_len + r_pad_boundary - 1) / r_pad_boundary)
    // r_total_len = r_num_blks * r_pad_boundary
    // r_pad_len = r_total_len - r_comp_data_len
   
    sll         r_pad_boundary, 1, STORAGE_KIVEC5_PAD_LEN_SHIFT
    add         r_num_blks, r_comp_data_len, r_pad_boundary
    sub         r_num_blks, r_num_blks, 1
    srl         r_num_blks, r_num_blks, STORAGE_KIVEC5_PAD_LEN_SHIFT
    sll         r_total_len, r_num_blks, STORAGE_KIVEC5_PAD_LEN_SHIFT
    sub         r_pad_len, r_total_len, r_comp_data_len

    phvwrpair   p.acc_chain_data_len, r_comp_data_len.wx, \
                p.acc_chain_pad_len, r_pad_len.wx
    bbeq        STORAGE_KIVEC5_SGL_PAD_HASH_EN, 1, sgl_padding_for_hash
    phvwr       p.acc_chain_total_len, r_total_len.wx                // delay slot
    bbeq        STORAGE_KIVEC5_AOL_PAD_EN, 0, possible_barco_push
    nop

aol_padding:
    
    // AOL padding enabled:
    // Note that preliminary DMA setup has been made in storage_tx_seq_comp_status_desc1_handler.
    // We now make adjustment based on pad length result.
    bne         r_pad_len, r0, possible_barco_push
    nop

    // pad length is zero so don't write A1/L1
    DMA_CMD_CANCEL(dma_p2m_3)
    DMA_CMD_CANCEL(dma_p2m_4)
    b           possible_barco_push
    nop

sgl_padding_for_hash:
   
    // SGL padding for hash enabled:
    // Given a vector of SGLs, each prefilled with exactly one block addr and len,
    // i.e., addr0/len0 specifiy one block of data, find the last applicable SGL
    // and apply padding.
   
if0:
    beq         r_pad_len, r0, endif0
    sub         r_last_sgl_p, r_num_blks, 1                       // delay slot
    add         r_last_sgl_p, STORAGE_KIVEC2ACC_SGL_VEC_ADDR, \
                r_last_sgl_p, BARCO_SGL_DESC_SIZE_SHIFT

    // Set up DMA for the following:
    // last_sgl_p->len0 = r_pad_boundary - r_pad_len
    // last_sgl_p->addr1 = comp_pad_buf
    // last_sgl_p->len1 = r_pad_len
    
    sub         r_last_blk_len, r_pad_boundary, r_pad_len
    phvwr       p.acc_chain_data_len, r_last_blk_len.wx
    
    add         r_sgl_field_p, r_last_sgl_p, \
                SIZE_IN_BYTES(offsetof(struct barco_sgl_le_t, len0))
    DMA_PHV2MEM_SETUP_ADDR64(acc_chain_data_len, acc_chain_data_len, r_sgl_field_p, dma_p2m_2)
    
    add         r_sgl_field_p, r_last_sgl_p, \
                SIZE_IN_BYTES(offsetof(struct barco_sgl_le_t, addr1))
    DMA_PHV2MEM_SETUP_ADDR64(acc_chain_pad_buf_addr, acc_chain_pad_buf_addr, r_sgl_field_p, dma_p2m_3)
    
    add         r_sgl_field_p, r_last_sgl_p, \
                SIZE_IN_BYTES(offsetof(struct barco_sgl_le_t, len1))
    DMA_PHV2MEM_SETUP_ADDR64(acc_chain_pad_len, acc_chain_pad_len, r_sgl_field_p, dma_p2m_4)
    DMA_PHV2MEM_FENCE(dma_p2m_4)
endif0:

    // In addition, storage_tx_seq_comp_status_desc0_handler has already set up
    // DMA transfer of a hash descriptor for the Barco ring. In the hash case,
    // the descriptor address actually points to the 1st of a vector of descriptors.
    // We now correct the transfer length based on r_num_blks.
    
    sll         r_desc_vec_len, r_num_blks, STORAGE_KIVEC4_BARCO_DESC_SIZE
    DMA_SIZE_UPDATE(r_desc_vec_len, dma_m2m_9)
    DMA_SIZE_UPDATE(r_desc_vec_len, dma_m2m_10)
    phvwr       p.storage_kivec4_barco_num_descs, r_num_blks
    
possible_barco_push:

    // if Barco ring push is applicable, execute table lock read
    // to get the current ring pindex. Note that this must be done
    // in the same stage as storage_tx_seq_barco_entry_handler_start()
    // which is stage 2.
    bbeq        STORAGE_KIVEC5_NEXT_DB_ACTION_BARCO_PUSH, 0, all_dma_complete
    nop

    // Set the table and program address 
    LOAD_TABLE_FOR_ADDR34_PC_IMM(STORAGE_KIVEC4_BARCO_PNDX_SHADOW_ADDR,
                                 STORAGE_KIVEC4_BARCO_PNDX_SIZE,
                                 storage_tx_seq_barco_chain_action_start)
all_dma_complete:

    // Setup the start and end DMA pointers
    DMA_PTR_SETUP(dma_p2m_0_dma_cmd_pad, dma_p2m_11_dma_cmd_eop,
                  p4_txdma_intr_dma_cmd_ptr)

exit:
    LOAD_NO_TABLES
   
comp_error:

    // TODO: if copy_src_dst_on_error wss set, copy header (presumably containing
    // an error header) plus source data to destination.
    //
   
   // Cancel any AOL DMA that might have been set up by storage_tx_seq_comp_status_desc1_handler
    bbeq        STORAGE_KIVEC5_AOL_PAD_EN, 0, possible_stop_chain
    seq         c1, STORAGE_KIVEC5_NEXT_DB_EN, 1 // delay slot
   
    DMA_CMD_CANCEL(dma_p2m_2)
    DMA_CMD_CANCEL(dma_p2m_3)
    DMA_CMD_CANCEL(dma_p2m_4)
    DMA_CMD_CANCEL(dma_p2m_5)

possible_stop_chain:
   
    // if next_db_en and !stop_chain_on_error then ring_db
    bbeq.c1     STORAGE_KIVEC5_STOP_CHAIN_ON_ERROR, 0, possible_barco_push
    nop

    // cancel any barco push prep
    DMA_CMD_CANCEL(dma_m2m_9)
    DMA_CMD_CANCEL(dma_m2m_10)
    DMA_CMD_CANCEL(dma_p2m_11)
    
    // else if intr_en then complete any status DMA and 
    // override doorbell to raising an interrupt
    bbeq        STORAGE_KIVEC5_INTR_EN, 0, exit
    nop

    PCI_SET_INTERRUPT_ADDR_DMA(STORAGE_KIVEC5_INTR_ADDR, dma_p2m_11)
    b           all_dma_complete
    nop
      
sgl_pdma_xfer:
   
    // PDMA compressed data to user buffers specified in SGL
    LOAD_TABLE1_FOR_ADDR_PC_IMM_e(STORAGE_KIVEC2ACC_SGL_PDMA_OUT_ADDR, 
                                  STORAGE_DEFAULT_TBL_LOAD_SIZE,
                                  storage_tx_seq_comp_sgl_handler_start)
