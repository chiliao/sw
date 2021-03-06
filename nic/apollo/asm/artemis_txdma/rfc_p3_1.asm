#include "../../p4/include/artemis_sacl_defines.h"
#include "artemis_txdma.h"
#include "INGRESS_p.h"
#include "INGRESS_rfc_p3_1_k.h"

struct phv_                p;
struct rfc_p3_1_k_         k;

%%

rfc_p3_1:
    /* Compute the index into the classid array */
    mod        r7, k.txdma_control_rfc_index, SACL_P3_ENTRIES_PER_CACHE_LINE
    mul        r7, r7, SACL_P3_ENTRY_WIDTH
    /* Access the classid at the index */
    tblrdp     r7, r7, 0, SACL_P3_ENTRY_WIDTH
    /* Priority = r7 >> SACL_P3_ENTRY_PRIORITY_SHIFT */
    srl        r1, r7, SACL_P3_ENTRY_PRIORITY_SHIFT
    /* Action = r7 & SACL_P3_ENTRY_ACTION_MASK */
    and        r2, r7, SACL_P3_ENTRY_ACTION_MASK
    /* Is Priority higher (numerically less) than that in PHV? */
    slt        c1, r1, k.txdma_control_rule_priority
    /* If so, update PHV with new priority and action */
    phvwr.c1   p.txdma_control_rule_priority, r1
    phvwr.c1   p.session_info_hint_drop, r2

    /* Is this the first pass? */
    seq        c1, k.capri_p4_intr_recirc_count, 0
    /* If not, stop */
    nop.!c1.e

    /* First Pass. Load sacl base addr to r1 */
    add        r1, r0, k.rx_to_tx_hdr_sacl_base_addr
    /* Add SACL_P1_3_TABLE_OFFSET to sacl base address. */
    addi       r1, r1, SACL_P1_3_TABLE_OFFSET
    /* P1 table index = (dtag_classid | (sip_classid << 7)). */
    add        r2, k.rx_to_tx_hdr_dtag_classid, k.rx_to_tx_hdr_sip_classid, \
                                                SACL_TAG_CLASSID_WIDTH
    /* Write P1 table index to PHV */
    phvwr      p.txdma_control_rfc_index, r2
    /* Compute the byte offset for P1 table index */
    div        r2, r2, SACL_P1_ENTRIES_PER_CACHE_LINE
    mul        r2, r2, SACL_CACHE_LINE_SIZE
    /* Add the byte offset to table base */
    add.e      r1, r1, r2
    /* Write the address back to phv for P1 lookup */
    phvwr      p.txdma_control_rfc_table_addr, r1

/*****************************************************************************/
/* error function                                                            */
/*****************************************************************************/
.align
.assert $ < ASM_INSTRUCTION_OFFSET_MAX
rfc_p3_1_error:
    phvwr.e         p.capri_intr_drop, 1
    nop
