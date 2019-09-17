#include "apulu.h"
#include "ingress.h"
#include "INGRESS_p.h"
#include "INGRESS_ipv4_flow_k.h"

struct ipv4_flow_k_ k;
struct ipv4_flow_d  d;
struct phv_ p;

#define FLOW_HASH_MSB 31:20

%%

ipv4_flow_hash:
    bbne        d.ipv4_flow_hash_d.entry_valid, TRUE, label_flow_miss
    phvwr       p.p4i_i2e_entropy_hash, r1
    bcf         [c1], label_flow_hit
    // Check hash1 and hint1
    seq         c1, r1[FLOW_HASH_MSB], d.ipv4_flow_hash_d.hash1
    sne         c2, d.ipv4_flow_hash_d.hint1, r0
    bcf         [c1&c2], label_flow_hash_hit
    add         r2, r0, d.ipv4_flow_hash_d.hint1
    // Check hash2 and hint2
    seq         c1, r1[FLOW_HASH_MSB], d.ipv4_flow_hash_d.hash2
    sne         c2, d.ipv4_flow_hash_d.hint2, r0
    bcf         [c1&c2], label_flow_hash_hit
    add         r2, r0, d.ipv4_flow_hash_d.hint2
    // Check for more hashes
    seq         c1, d.ipv4_flow_hash_d.more_hashes, TRUE
    sne         c2, d.ipv4_flow_hash_d.more_hints, r0
    bcf         [c1&c2], label_flow_hash_hit
    add         r2, r0, d.ipv4_flow_hash_d.more_hints

label_flow_miss:
    phvwr.e     p.control_metadata_flow_miss, TRUE
    phvwr.f     p.ingress_recirc_flow_done, TRUE

label_flow_hit:
    slt         c1, d.ipv4_flow_hash_d.epoch, k.control_metadata_epoch
    bcf         [c1], label_flow_miss
    phvwr       p.ingress_recirc_flow_done, TRUE
    phvwr.e     p.p4i_i2e_session_id, d.ipv4_flow_hash_d.session_id
    phvwr.f     p.p4i_i2e_flow_role, d.ipv4_flow_hash_d.flow_role

label_flow_hash_hit:
    // Set bit 31 for overflow hash lookup to work
    ori         r2, r2, 0x80000000
    phvwr.e     p.ingress_recirc_flow_ohash, r2
    phvwr.f     p.control_metadata_flow_ohash_lkp, TRUE

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
ipv4_flow_error:
    phvwr.e         p.capri_intrinsic_drop, 1
    nop
