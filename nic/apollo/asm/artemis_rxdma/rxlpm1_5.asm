#include "../../p4/include/lpm_defines.h"

#include "artemis_rxdma.h"
#include "INGRESS_p.h"
#include "ingress.h"

struct phv_                p;
struct rxlpm1_5_k          k;
struct rxlpm1_5_d          d;

// Define Table Name and Action Names
#define table_name         rxlpm1_5
#define action_keys32b     match1_5_32b
#define action_keys128b    match1_5_128b
#define action_data32b     match1_5_32b_retrieve
#define action_data128b    match1_5_128b_retrieve

// Define table field names for the selected actions
#define keys32b(a)         d.u.match1_5_32b_d.key ## a
#define keys128bhi(a)      d.u.match1_5_128b_d.key ## a[127:64]
#define keys128blo(a)      d.u.match1_5_128b_d.key ## a[63:0]
#define keys32b_(a)        d.u.match1_5_32b_retrieve_d.key ## a
#define data32b(a)         d.u.match1_5_32b_retrieve_d.data ## a
#define keys128bhi_(a)     d.u.match1_5_128b_retrieve_d.key ## a[127:64]
#define keys128blo_(a)     d.u.match1_5_128b_retrieve_d.key ## a[63:0]
#define data128b(a)        d.u.match1_5_128b_retrieve_d.data ## a

// Define key field names
#define key                k.lpm_metadata_lpm1_key
#define keylo              k.lpm_metadata_lpm1_key[63:0]
#define keyhi              k.lpm_metadata_lpm1_key[127:64]
#define base_addr          k.lpm_metadata_lpm1_base_addr
#define curr_addr          k.lpm_metadata_lpm1_next_addr

// Define PHV field names
#define next_addr          p.lpm_metadata_lpm1_next_addr

// Define result register and handler function name
#define res_reg            r7
#define result_handler32b  nop_res_handler
#define result_handler128b rxlpm1_res_handler

%%

#include "../include/lpm.h"

nop_res_handler:
    nop.e
    nop

rxlpm1_res_handler:
    /* Is this the second pass? */
    seq          c1, k.capri_p4_intr_recirc_count, 1
    /* If so, write SIP classid result, and Stop */
    phvwr.c1.e   p.rx_to_tx_hdr_sip_classid, res_reg
    /* Is this the third pass? */
    seq          c1, k.capri_p4_intr_recirc_count, 2
    /* If not, go to end */
    bcf          [!c1], exit
    /* Third pass. Write TAG classid result, and Stop */
    seq          c1, k.p4_to_rxdma_direction, TX_FROM_HOST
    phvwr.c1     p.rx_to_tx_hdr_dtag_classid, res_reg
    phvwr.!c1    p.rx_to_tx_hdr_stag_classid, res_reg
exit:
    nop.e
    nop
