/*
 *  Implements the CC stage of the RxDMA P4+ pipeline
 */


#include "tcp-shared-state.h"
#include "tcp-macros.h"
#include "tcp-table.h"
#include "tcp-constants.h"
#include "ingress.h"
#include "INGRESS_p.h"
#include "INGRESS_s5_t0_tcp_rx_k.h"

struct phv_ p;
struct s5_t0_tcp_rx_k_ k;
struct s5_t0_tcp_rx_tcp_fc_d d;

%%
    .param          tcp_rx_dma_serq_stage_start
    .param          tcp_rx_dma_rx2tx_stage_start
    .param          tcp_rx_write_arq_stage_start
    .align

tcp_rx_fc_stage_start:
    seq         c1, k.common_phv_write_arq, 1
    bcf         [c1], tcp_cpu_rx
    nop

    // launch table 1 next stage
    CAPRI_NEXT_TABLE_READ_OFFSET(1, TABLE_LOCK_EN,
                tcp_rx_dma_rx2tx_stage_start, k.common_phv_qstate_addr,
                TCP_TCB_RX_DMA_OFFSET, TABLE_SIZE_512_BITS)

flow_fc_process:
    and         r1, k.common_phv_pending_txdma, \
                    TCP_PENDING_TXDMA_ACK_SEND | TCP_PENDING_TXDMA_DEL_ACK
    seq         c2, r1, 0

    // for dupack, do not recalc window, unless rcv_wnd is 0
    seq         c3, k.common_phv_is_dupack, 1
    sne.c3      c3, d.rcv_wnd, 0
    bcf         [c2 | c3], flow_fc_process_done
    nop

    bbeq        k.common_phv_write_serq, 0, start_window_calc
    nop

    // Calculate average every 16 packets
    tblmincri   d.num_pkts, 4, 1
    tbladd      d.cum_pkt_size, k.s1_s2s_payload_len
    seq         c1, d.num_pkts, 0
    b.!c1       start_window_calc

    srl         r7, d.cum_pkt_size, 4
    tblwr       d.cum_pkt_size, 0

    sle         c1, r7, 128
    tblwr.c1    d.avg_pkt_size_shift, 5
    b.c1        start_window_calc
    sle         c1, r7, 256
    tblwr.c1    d.avg_pkt_size_shift, 6
    b.c1        start_window_calc
    sle         c1, r7, 512
    tblwr.c1    d.avg_pkt_size_shift, 7
    b.c1        start_window_calc
    sle         c1, r7, 768
    tblwr.c1    d.avg_pkt_size_shift, 8
    b.c1        start_window_calc
    sle         c1, r7, 1024
    tblwr.c1    d.avg_pkt_size_shift, 9
    b.c1        start_window_calc

    tblwr.!c1   d.avg_pkt_size_shift, 11

start_window_calc:
    /* Figure out how many entries are free in serq */
    add         r2, k.to_s5_serq_cidx, d.consumer_ring_slots
    sub         r2, r2, k.to_s5_serq_pidx
    and         r2, r2, d.consumer_ring_slots_mask

    seq         c4, k.to_s5_rnmdr_size_valid, 1
    tblwr.c4    d.rnmdr_size, k.to_s5_rnmdr_size
    add.c4      r4, r0, k.to_s5_rnmdr_size
    add.!c4     r4, r0, d.rnmdr_size

    slt         c1, d.high_thresh1, r2
    slt.c1      c1, 0x1800, r4
    sll.c1      r4, r2, d.avg_pkt_size_shift
    b.c1        window_calc_done

    slt         c1, d.high_thresh2, r2
    slt.c1      c1, 0x1000, r4
    sub         r7, d.avg_pkt_size_shift, 1
    sll.c1      r4, r2, r7
    b.c1        window_calc_done

    slt         c1, d.high_thresh3, r2
    slt.c1      c1, 0x800, r4
    sub         r7, d.avg_pkt_size_shift, 2
    sll.c1      r4, r2, r7
    b.c1        window_calc_done

    slt         c1, d.high_thresh4, r2
    slt.c1      c1, 0x400, r4
    sub         r7, d.avg_pkt_size_shift, 3
    sll.c1      r4, r2, r7
    b.c1        window_calc_done
    nop

    // When less than thresh4 slots are free, assume 0 window
    add         r4, r0, 0

window_calc_done:

    /* If the window calculated is smaller than the window
    * advertised previously reset to previous window to avoid
    * shrinking rcv window. See the fix for Bug627496 in Linux
    * kernel.
    *
    * current window = rcv_wup + rcv_wnd - rcv_nxt
    *
    * Ignore the overflow part once we get rcv_wup + rcv_wnd. Need
    * to do 32 bit arithmetic here since all sequence numbers 
    * including rcv_nxt are unsigned 32 bit values.
    */

    add         r3, d.rcv_wup, d.rcv_wnd
    add         r3, r0, r3[31:0]
    sub         r3, r3, k.to_s5_rcv_nxt

    /* In case (rcv_wup + rcv_wnd - rcv_nxt) < 0 reset current
    * window c3 to zero.
    */

    sle.s       c3, r3, 0
    add.c3      r3, r0, r0

    /* r3 is the current window. Add ( 2 ^ wscale ) - 1 to
    * current window to avoid shrinking the window when we
    * apply scale
    */

    sll.!c3     r2, 1, d.rcv_scale
    sub.!c3     r2, r2, 1
    add.!c3     r3, r3, r2

    /* Now compare the calculated window and the previosly advertised
    * window. In order to avoid silly window syndrome open the right
    * edge only if we can move it at least by one MSS (receiver side
    * SWS). */
    add         r2, r3, d.rcv_mss
    sle         c3, r4, r2
    add.c3      r4, r0, r3

    /* r4 is the window we want to advertise */
    tblwr       d.rcv_wnd, r4

flow_fc_process_done:
    /* Apply the scale factor SEG.WND = RCV.WND >> Rcv.Wind.Scale */
    srl         r4, d.rcv_wnd, d.rcv_scale

    /* If scaled widnow is > 0xFFFF we have an issue with the scale
    * factor passed from LKL. Reset to 0xFFFF to avoid overflow.
    */
    slt         c3, 0xffff, r4
    add.c3      r4, r0, 0xffff
    sll.c3      r3, 0xffff, d.rcv_scale
    tblwr.c3    d.rcv_wnd, r3

    /* We have a valid window after applying the scale factor. */
    phvwr       p.rx2tx_extra_rcv_wnd, r4
    tblwr       d.rcv_wup, k.to_s5_rcv_nxt

    /* If we are advertizing zero window, subscribe for application
    * read notification so that we can generate window update. Wait
    * for 5 descriptors to be freed for now. Need to tune this 
    * further.
    */

    add         r2, r0, d.read_notify_addr
    seq         c1, r0, r2
    memwr.h.!c1 r2, 0       // d.read_notify_addr invalid/not set?

    seq.!c1     c2, r0, r4
    memwr.h.c2  r2, 5

    seq         c1, k.common_phv_ooo_rcv, 1
    seq         c2, k.common_phv_ooq_tx2rx_win_upd, 1
    seq.!c2     c2, k.common_phv_ooq_tx2rx_last_ooo_pkt, 1
    phvwri.c2   p.p4_rxdma_intr_dma_cmd_ptr, (CAPRI_PHV_START_OFFSET(rx2tx_extra_dma_dma_cmd_type) / 16)
    bcf         [c1 | c2], flow_fc_skip_serq
    CAPRI_NEXT_TABLE_READ_OFFSET(0, TABLE_LOCK_EN,
                tcp_rx_dma_serq_stage_start, k.common_phv_qstate_addr,
                TCP_TCB_RX_DMA_OFFSET, TABLE_SIZE_512_BITS)

    nop.e
    nop

flow_fc_skip_serq:
    // Skip serq launch for OOO pkt or win_upd packet
    CAPRI_CLEAR_TABLE_VALID(0)

    nop.e
    nop

tcp_cpu_rx:
#ifdef TCP_ACTL_Q
    CPU_TCP_ACTL_Q_SEM_INF_ADDR(d.cpu_id, r3)
#else
    CPU_ARQ_SEM_INF_ADDR(d.cpu_id, r3)
#endif
    phvwr       p.s6_t1_s2s_cpu_id, d.cpu_id

    CAPRI_NEXT_TABLE_READ(1,
                          TABLE_LOCK_DIS,
                          tcp_rx_write_arq_stage_start,
                          r3,
                          TABLE_SIZE_64_BITS)

    b           flow_fc_process
    nop
