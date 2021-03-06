#include "apulu.h"
#include "egress.h"
#include "EGRESS_p.h"
#include "EGRESS_nat2_k.h"

struct nat2_k_  k;
struct nat2_d   d;
struct phv_     p;

%%

nat2_rewrite:
    seq             c1, k.ipv4_1_valid, TRUE
    seq             c2, k.ipv6_1_valid, TRUE
    bcf             [!c1 & !c2], nat2_rewrite_done
    seq             c2, k.udp_1_valid, TRUE
    seq             c3, k.tcp_valid, TRUE
    bbeq            k.control_metadata_rx_packet, TRUE, nat2_rewrite_rx
nat2_rewrite_tx:
    seq             c4, k.rewrite_metadata_flags[P4_REWRITE_DIP_BITS], \
                        P4_REWRITE_DIP_FROM_NAT
    nop.!c4.e
    phvwr.c4        p.control_metadata_update_checksum, TRUE
    phvwr.c1        p.ipv4_1_dstAddr, d.nat2_rewrite_d.ip
    phvwr.!c1       p.ipv6_1_dstAddr, d.nat2_rewrite_d.ip
    seq             c4, k.rewrite_metadata_flags[P4_REWRITE_DPORT_BITS], \
                        P4_REWRITE_DPORT_FROM_NAT
    nop.!c4.e
    nop
    phvwr.c3.e      p.tcp_dstPort, d.nat2_rewrite_d.port
    phvwr.c3        p.control_metadata_update_checksum, TRUE
    phvwr.c2        p.udp_1_dstPort, d.nat2_rewrite_d.port
    nop.e
    phvwr.c2        p.control_metadata_update_checksum, TRUE

nat2_rewrite_rx:
    seq             c4, k.rewrite_metadata_flags[P4_REWRITE_SIP_BITS], \
                        P4_REWRITE_SIP_FROM_NAT
    nop.!c4.e
    phvwr.c4        p.control_metadata_update_checksum, TRUE
    phvwr.c1        p.ipv4_1_srcAddr, d.nat2_rewrite_d.ip
    phvwr.!c1       p.ipv6_1_srcAddr, d.nat2_rewrite_d.ip
    seq             c4, k.rewrite_metadata_flags[P4_REWRITE_SPORT_BITS], \
                        P4_REWRITE_SPORT_FROM_NAT
    nop.!c4.e
    nop
    phvwr.c3.e      p.tcp_srcPort, d.nat2_rewrite_d.port
    phvwr.c3        p.control_metadata_update_checksum, TRUE
    phvwr.c2        p.udp_1_srcPort, d.nat2_rewrite_d.port
    nop.e
    phvwr.c2        p.control_metadata_update_checksum, TRUE

nat2_rewrite_done:
    nop.e
    nop

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
nat2_error:
    phvwr.e         p.capri_intrinsic_drop, 1
    nop
