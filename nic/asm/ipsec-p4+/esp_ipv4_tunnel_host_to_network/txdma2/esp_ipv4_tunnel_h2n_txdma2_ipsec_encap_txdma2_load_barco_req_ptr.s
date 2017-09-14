#include "INGRESS_p.h"
#include "ingress.h"
#include "ipsec_asm_defines.h"

struct tx_table_s1_t0_k k;
struct tx_table_s1_t0_ipsec_encap_txdma2_load_barco_req_ptr_d d;
struct phv_ p;

%%
        .param esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_barco_req
        .align
esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_barco_req_ptr:
    phvwri p.app_header_table0_valid, 1
    phvwri p.common_te0_phv_table_lock_en, 1
    phvwri p.common_te0_phv_table_raw_table_size, 6 
    addi r2, r0, esp_ipv4_tunnel_h2n_txdma2_ipsec_encap_txdma2_load_barco_req
    srl r2, r2, 6
    phvwr p.common_te0_phv_table_pc, r2 
    phvwr  p.common_te0_phv_table_addr, d.barco_req_address 
    nop.e
    nop
