#include "INGRESS_p.h"
#include "ingress.h"
#include "ipsec_defines.h"

struct tx_table_s3_t0_k k;
struct tx_table_s3_t0_ipsec_encap_txdma2_load_in_desc_d d;
struct phv_ p;

%%
        .param ipsec_build_encap_packet 
        .align
ipsec_encap_txdma2_load_in_desc:
    phvwr p.t0_s2s_in_page_addr, d.addr0
    phvwri p.app_header_table0_valid, 1
    phvwri p.common_te0_phv_table_lock_en, 1
    phvwri p.common_te0_phv_table_raw_table_size, 6
    phvwri p.common_te0_phv_table_pc, ipsec_build_encap_packet 
    phvwr  p.common_te0_phv_table_addr, k.txdma2_global_in_desc_addr
    nop.e
 
