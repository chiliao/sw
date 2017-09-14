#include "INGRESS_p.h"
#include "ingress.h"
#include "ipsec_asm_defines.h"

struct tx_table_s2_t0_k k;
struct tx_table_s2_t0_ipsec_encap_txdma_load_head_desc_int_header_d d;
struct phv_ p;

%%
        .param esp_ipv4_tunnel_h2n_txdma1_ipsec_write_barco_req
        .align
esp_ipv4_tunnel_h2n_txdma1_ipsec_encap_txdma_load_head_desc_int_header:
    phvwr p.barco_req_brq_in_addr, d.in_desc
    phvwr p.barco_req_brq_out_addr, d.out_desc
    phvwr p.barco_req_brq_iv_addr, d.in_page
    add r1, r0, d.pad_size
    add r1, r1, d.tailroom_offset
    addi r1, r1, 2
    add r1, r1, d.out_page
    phvwr p.barco_req_brq_auth_tag_addr, r1
    phvwr p.barco_req_brq_hdr_size, d.payload_start

    addi r2, r0, esp_ipv4_tunnel_h2n_txdma1_ipsec_write_barco_req
    srl r2, r2, 6
    phvwr p.common_te0_phv_table_pc, r2
    phvwri p.common_te0_phv_table_lock_en, 1
    phvwri p.common_te0_phv_table_raw_table_size, 6
    phvwr  p.common_te0_phv_table_addr, k.ipsec_to_stage2_barco_req_addr 
    phvwri p.app_header_table1_valid, 0

    nop.e
    nop 

