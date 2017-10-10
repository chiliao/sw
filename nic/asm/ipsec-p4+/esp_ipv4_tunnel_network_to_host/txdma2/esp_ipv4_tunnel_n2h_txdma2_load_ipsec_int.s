#include "INGRESS_p.h"
#include "ingress.h"
#include "ipsec_asm_defines.h"

struct tx_table_s2_t2_k k;
struct tx_table_s2_t2_esp_v4_tunnel_n2h_txdma2_load_ipsec_int_d d;
struct phv_ p;

%%
        .align
esp_ipv4_tunnel_n2h_txdma2_load_ipsec_int:
    phvwri p.app_header_table2_valid, 0
    phvwr p.txdma2_global_pad_size, d.pad_size
    phvwr p.txdma2_global_l4_protocol, d.l4_protocol
    phvwr p.txdma2_global_payload_size, d.payload_size
    phvwr p.t0_s2s_headroom_offset, d.headroom_offset
    phvwr p.t0_s2s_tailroom_offset, d.tailroom_offset
    phvwr p.ipsec_to_stage3_in_page, d.in_page
    phvwr p.ipsec_to_stage3_headroom, d.headroom
    nop.e
    nop
 
