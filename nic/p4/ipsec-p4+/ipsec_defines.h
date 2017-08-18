#define IPSEC_CB_BASE 0xaaaaaaaa
#define IPSEC_CB_SIZE 4096 

#define IPSEC_CB_SHIFT_SIZE 6

#define IPSEC_CB_TAIL_DESC_ADDR_OFFSET 51
#define IPSEC_CB_IV_OFFSET 14
#define IPSEC_CB_HEAD_DESC_ADDR_OFFSET 46

#define IPSEC_PAD_BYTES_TABLE_BASE 0xbbbbbbbb

#define INDESC_SEMAPHORE_ADDR   0xa0a0a0a0
#define OUTDESC_SEMAPHORE_ADDR  0xa1a1a1a1
#define INPAGE_SEMAPHORE_ADDR   0xa2a2a2a2
#define OUTPAGE_SEMAPHORE_ADDR  0xa3a3a3a3

#define IN_DESC_RING_BASE 0xcccccccc
#define OUT_DESC_RING_BASE 0xcccc0000
#define IN_PAGE_RING_BASE 0xdddddddd
#define OUT_PAGE_RING_BASE 0xdddd0000

#define IN_DESC_ADDR_BASE 0xabababab
#define IN_PAGE_ADDR_BASE 0xbabababa
#define OUT_DESC_ADDR_BASE 0xaaaabbbb
#define OUT_PAGE_ADDR_BASE 0xbbbbaaaa

#define RING_INDEX_WIDTH 16
#define DESC_PTR_SIZE 8
#define PAGE_PTR_SIZE 8


#define ESP_FIXED_HDR_SIZE 8 
#define AOL_OFFSET_WIDTH 32
#define AOL_LENGTH_WIDTH 32

#define DESC_ENTRY_SIZE 128 
#define PAGE_ENTRY_SIZE (10 * 1024)
#define DESC_SHIFT_WIDTH 7

#define IPSEC_RXDMA_HW_SW_INTRINSIC_SIZE 42 
#define IPSEC_TXDMA_HW_INTRINSIC_SIZE 31 


#define IPSEC_INT_START_OFFSET  344
#define IPSEC_INT_END_OFFSET    407
#define IPSEC_IN_DESC_AOL_START 408
#define IPSEC_IN_DESC_AOL_END   471 
#define IPSEC_OUT_DESC_AOL_START 472 
#define IPSEC_OUT_DESC_AOL_END   535  
#define IPSEC_TAIL_DESC_ADDR_PHV_OFFSET_START 128 
#define IPSEC_TAIL_DESC_ADDR_PHV_OFFSET_END   135
#define IPSEC_TXDMA1_BARCO_REQ_PHV_OFFSET_START 128 
#define IPSEC_TXDMA1_BARCO_REQ_PHV_OFFSET_END   135 
#define IPSEC_TXDMA1_HEAD_DESC_PHV_OFFSET_START 73 
#define IPSEC_TXDMA1_HEAD_DESC_PHV_OFFSET_END   80 

#define RXDMA_IPSEC_DMA_COMMANDS_OFFSET 36 
#define TXDMA1_DMA_COMMANDS_OFFSET 36 
#define TXDMA2_DMA_COMMANDS_OFFSET 24
 
#define BRQ_REQ_ENTRY_SIZE 64
#define BRQ_REQ_SEMAPHORE_ADDR 0xcccccccc
#define BRQ_REQ_RING_BASE_ADDR 0xaaaacccc

#define ESP_BASE_OFFSET 18  


#define IPSEC_PAD_BYTES_HBM_TABLE_BASE    0xa0000000


