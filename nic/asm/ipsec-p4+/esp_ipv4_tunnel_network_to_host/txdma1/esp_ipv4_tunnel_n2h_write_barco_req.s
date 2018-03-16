#include "INGRESS_p.h"
#include "ingress.h"
#include "ipsec_asm_defines.h"
#include "capri-macros.h"

struct tx_table_s3_t0_k k;
struct tx_table_s3_t0_esp_v4_tunnel_n2h_txdma1_write_barco_req_d d;
struct phv_ p;

%%
        .param esp_v4_tunnel_n2h_txdma1_update_cb
        .align
esp_v4_tunnel_n2h_write_barco_req:
    phvwr p.brq_req_write_dma_cmd_addr, k.ipsec_to_stage3_barco_req_addr 
    seq c1, k.ipsec_to_stage3_new_key, 1
    phvwr.c1 p.barco_req_key_desc_index, d.{new_key_index}.wx
    phvwr.!c1 p.barco_req_key_desc_index, d.{key_index}.wx

esp_v4_tunnel_n2h_post_to_barco_ring:
    and r3, d.barco_pindex, IPSEC_BARCO_RING_INDEX_MASK
    sll r3, r3, IPSEC_BARCO_RING_ENTRY_SHIFT_SIZE
    add r3, r3, d.barco_ring_base_addr 
    phvwr p.dma_cmd_post_barco_ring_dma_cmd_addr, r3
    phvwri p.dma_cmd_post_barco_ring_dma_cmd_phv_start_addr, IPSEC_TXDMA1_BARCO_REQ_PHV_OFFSET_START
    phvwri p.dma_cmd_post_barco_ring_dma_cmd_phv_end_addr, IPSEC_TXDMA1_BARCO_REQ_PHV_OFFSET_END

esp_v4_tunnel_n2h_dma_cmd_incr_barco_pindex:
    tbladd d.barco_pindex, 1
    tbland d.barco_pindex, 0x3F

    phvwri p.app_header_table0_valid, 1
    phvwri p.common_te0_phv_table_pc, esp_v4_tunnel_n2h_txdma1_update_cb[33:6]
    phvwri p.{common_te0_phv_table_lock_en...common_te0_phv_table_raw_table_size}, 14 
    phvwr p.common_te0_phv_table_addr, k.txdma1_global_ipsec_cb_addr
    nop.e
    nop
