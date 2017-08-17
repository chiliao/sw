/*
 *	Get an index and auto increment it.
 *      This stage will be used to get
	      - RNMDR alloc idx
 */

#include "tcp-shared-state.h"
#include "tcp-macros.h"
#include "tcp-table.h"
#include "ingress.h"
#include "INGRESS_p.h"

struct phv_ p;
struct tcp_rx_read_rnmdr_k k;
struct tcp_rx_read_rnmdr_read_rnmdr_d d;
	
%%
        .param          tcp_rx_rdesc_alloc_stage_3_start
	.align
tcp_rx_read_rnmdr_stage2_start:

        CAPRI_CLEAR_TABLE1_VALID

	phvwr		p.s3_t1_s2s_rnmdr_pidx, d.rnmdr_pidx

table_read_RNMDR_DESC:
	add		r3, r0, k.to_s2_rnmdr_base
	CAPRI_NEXT_TABLE1_READ(d.rnmdr_pidx, TABLE_LOCK_EN,
                            tcp_rx_rdesc_alloc_stage_3_start,
	                    r3, RNMDR_TABLE_ENTRY_SIZE_SHFT,
	                    0, TABLE_SIZE_512_BITS)
	nop.e
	nop
