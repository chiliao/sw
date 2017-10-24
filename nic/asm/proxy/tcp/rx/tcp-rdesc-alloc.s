/*
 *    Implements the descr allocation stage of the RxDMA P4+ pipeline
 */

#include "tcp-shared-state.h"
#include "tcp-macros.h"
#include "tcp-table.h"
#include "ingress.h"
#include "INGRESS_p.h"

struct phv_ p;
struct tcp_rx_rdesc_alloc_k k;
struct tcp_rx_rdesc_alloc_rdesc_alloc_d d;

%%
    .align
tcp_rx_rdesc_alloc_stage_3_start:
    CAPRI_CLEAR_TABLE1_VALID
    CAPRI_SET_DEBUG_STAGE0_3(p.s6_s2s_debug_stage0_3_thread, CAPRI_MPU_STAGE_3, CAPRI_MPU_TABLE_1)

    CAPRI_OPERAND_DEBUG(d.desc)
    sne             c1, k.common_phv_ooo_rcv, r0
    phvwr.c1        p.to_s5_descr, d.desc
    phvwr.e         p.to_s6_descr, d.desc
    nop

