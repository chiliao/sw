
#include "p4pt.h"

struct p4pt_update_read_latency_distribution_k k;
struct p4pt_update_read_latency_distribution_d d;
struct phv_ p;

%%

    .param  p4pt_tcb_iscsi_write_latency_distribution_base
    .param  p4pt_update_write_latency_distribution_start

/*
 * - stage 4: write p4pt_tcb_iscsi_read_latency_distribution, setup lookup p4pt_tcb_iscsi_write_latency_distribution
 *    - k vector: p4pt_global, p4pt_s2s
 *    - d vector: p4pt_tcb_iscsi_read_latency_distribution
 *    - if p4pt_s2s.req:
 *         return
 *      if p4pt_s2s.write:
 *         return
 *      if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE15_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range15++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE14_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range14++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE13_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range13++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE12_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range12++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION1RANGE11_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range11++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE10_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range10++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE9_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range9++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE8_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range8++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE7_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range7++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE6_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range6++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE5_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range5++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE4_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range4++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE3_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range3++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE2_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range2++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE1_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range1++
 *      else if p4pt_global.latency > P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE0_LOW
 *         p4pt_tcb_iscsi_read_latency_distribution.range0++
 *      write p4pt_tcb_iscsi_read_latency_distribution
 *      - lookup launch p4pt_tcb_iscsi_write_latency_distribution
 *         lookup addr = P4PT_TCB_ISCSI_WRITE_LATENCY_DISTRIBUTION_BASE_ADDR + \
 *                       p4pt_global.p4pt_idx * P4PT_TCB_ISCSI_WRITE_LATENCY_DISTRIBUTION_SIZE
 *
 */

p4pt_update_read_latency_distribution_start:
    P4PT_CHECK_EXIT

    seq      c1, k.p4pt_s2s_req, 1
    b.c1     p4pt_update_read_latency_distribution_return

    seq      c1, k.p4pt_s2s_write, 1
    b.c1     p4pt_update_read_latency_distribution_return

    add      r1, r0, k.{p4pt_global_latency_sbit0_ebit3...p4pt_global_latency_sbit28_ebit31}
    addi     r3, r0, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE15_LOW
    sle      c1, r3, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range15
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range15, r2
    add.c1   r2, r0, 1
    b.c1     p4pt_update_read_latency_distribution_return

    addi     r3, r0, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE14_LOW
    sle      c1, r3, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range14
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range14, r2
    b.c1     p4pt_update_read_latency_distribution_return

    addi     r3, r0, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE13_LOW
    sle      c1, r3, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range13
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range13, r2
    b.c1     p4pt_update_read_latency_distribution_return

    addi     r3, r0, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE12_LOW
    sle      c1, r3, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range12
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range12, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE11_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range11
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range11, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE10_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range10
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range10, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE9_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range9
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range9, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE8_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range8
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range8, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE7_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range7
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range7, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE6_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range6
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range6, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE5_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range5
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range5, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE4_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range4
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range4, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE3_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range3
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range3, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE2_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range2
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range2, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE1_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range1
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range1, r2
    b.c1     p4pt_update_read_latency_distribution_return

    sle      c1, P4PT_ISCSI_LATENCY_DISTRIBUTION_RANGE0_LOW, r1
    add      r2, r0, d.u.p4pt_update_read_latency_distribution_d.range0
    add.c1   r2, r0, 1
    tblwr.c1 d.u.p4pt_update_read_latency_distribution_d.range0, r2

p4pt_update_read_latency_distribution_return:
    addi     r1, r0, loword(p4pt_tcb_iscsi_write_latency_distribution_base)
    addui    r1, r1, hiword(p4pt_tcb_iscsi_write_latency_distribution_base)
    add      r1, r1, k.p4pt_global_p4pt_idx, 6
    phvwr    p.common_te0_phv_table_addr, r1
    phvwri   p.common_te0_phv_table_pc, p4pt_update_write_latency_distribution_start[33:6]
    phvwr    p.common_te0_phv_table_raw_table_size, 6
    phvwr.e  p.common_te0_phv_table_lock_en, 0
    nop
