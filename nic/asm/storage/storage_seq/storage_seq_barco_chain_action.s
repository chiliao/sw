/*****************************************************************************
 *  storage_seq_barco_chain_action: Push to Barco ring by issuing mem2mem DMA
 *                       commands and advancing the p_ndx via ringing the doorbell.
 *                       Note that the number of entries may be greater than 1
 *                       and may cause the ring to wrap around, in which case,
 *                       the end result may become two mem2mem DMA transfers. 
 *****************************************************************************/

#include "storage_asm_defines.h"
#include "ingress.h"
#include "INGRESS_p.h"


struct s4_tbl_k k;
struct s4_tbl_seq_barco_chain_action_d d;
struct phv_ p;

/*
 * Registers usage
 * CAUTION: r1 is also implicitly used by LOAD_TABLEx_FOR_ADDR_PC_IMM()
 */
#define r_ring_addr                 r1  // ring address (dst of xfer)
#define r_avail_slots               r2  // number of ring slots before wrap around
#define r_pi                        r3  // producer index
#define r_xfers0_len                r4  // transfer length before wrap around
#define r_xfers1_len                r5  // transfer length after wrap around
#define r_num_descs                 r6  // # of descs to transfer
#define r_desc_addr                 r7  // starting descriptor address to transfer

%%

storage_seq_barco_chain_action:
    CLEAR_TABLE0

    // Note that d.p_ndx is a shadow copy of the HW ring producer index.
    add         r_pi, d.p_ndx, r0
    
    // Break into two DMA transfers if NUM_DESCS causes
    // p_ndx to wrap around
    sll         r_avail_slots, 1, SEQ_KIVEC4_BARCO_RING_SIZE
    sub         r_avail_slots, r_avail_slots, r_pi
    seq         c3, SEQ_KIVEC10_ALT_DESCS_SELECT, 1
    cmov        r_num_descs, c3, SEQ_KIVEC10_NUM_ALT_DESCS, SEQ_KIVEC4_BARCO_NUM_DESCS
if0:    
    ble         r_num_descs, r_avail_slots, endif0
    sll         r_xfers0_len, r_num_descs, SEQ_KIVEC4_BARCO_DESC_SIZE  // delay slot

    sll         r_xfers0_len, r_avail_slots, SEQ_KIVEC4_BARCO_DESC_SIZE
    sub         r_xfers1_len, r_num_descs, r_avail_slots
    sll         r_xfers1_len, r_xfers1_len, SEQ_KIVEC4_BARCO_DESC_SIZE
    add         r_desc_addr, SEQ_KIVEC4_BARCO_DESC_ADDR, r_xfers0_len

    // MEM2MEM has a max transfer length of 14 bits so ensure we
    // don't exceed it
if1:    
    sle         c1, r_xfers1_len, DMA_CMD_MEM2MEM_SIZE_MAX
    bcf         [c1], endif1
    
    // The extra wrapped around transfer goes to the start of the barco ring
    add         r_ring_addr, SEQ_KIVEC4_BARCO_RING_ADDR, r0     // delay slot
    SEQ_METRICS_SET(hw_batch_errs)
    DMA_CMD_MEM2MEM_SIZE_ERROR_TRAP()
endif1:

    DMA_MEM2MEM_NO_LIF_SETUP_REG_ADDR(CAPRI_DMA_M2M_TYPE_SRC, r_desc_addr,
                                      r_xfers1_len, dma_m2m_16)
    DMA_MEM2MEM_NO_LIF_SETUP_REG_ADDR(CAPRI_DMA_M2M_TYPE_DST, r_ring_addr,
                                      r_xfers1_len, dma_m2m_17)
endif0:

    // The first transfer before the ring wrapped (if it wrapped)
    sll         r_ring_addr, r_pi, SEQ_KIVEC4_BARCO_DESC_SIZE
    add         r_ring_addr, r_ring_addr, SEQ_KIVEC4_BARCO_RING_ADDR
if2:    
    sle         c1, r_xfers0_len, DMA_CMD_MEM2MEM_SIZE_MAX
    bcf         [c1], endif2
    add         r_desc_addr, SEQ_KIVEC4_BARCO_DESC_ADDR, r0     // delay slot
    SEQ_METRICS_SET(hw_batch_errs)
    DMA_CMD_MEM2MEM_SIZE_ERROR_TRAP()
endif2:

    DMA_MEM2MEM_NO_LIF_SETUP_REG_ADDR(CAPRI_DMA_M2M_TYPE_SRC, r_desc_addr,
                                      r_xfers0_len, dma_m2m_14)
    DMA_MEM2MEM_NO_LIF_SETUP_REG_ADDR(CAPRI_DMA_M2M_TYPE_DST, r_ring_addr,
                                      r_xfers0_len, dma_m2m_15)
    SEQ_METRICS_VAL_SET(hw_desc_xfers, r_num_descs)
    mincr       r_pi, SEQ_KIVEC4_BARCO_RING_SIZE, r_num_descs
    
    // Need to word swap before writing back as the p_ndx is little endian
    tblwr.f.e   d.p_ndx, r_pi
    phvwr	p.barco_doorbell_data_p_ndx, r_pi.wx            // delay slot

