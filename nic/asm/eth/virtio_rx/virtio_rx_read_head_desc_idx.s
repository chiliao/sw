/*
 * Stage 3, Table 0
 * Implementation of Virtio virtq_avail.ring read for the head descriptor idx
 */
#include "INGRESS_p.h"
#include "ingress.h"
#include "capri-macros.h"
#include "virtio_defines.h"


struct phv_                 p;
struct rx_table_s3_t0_k     k;
struct rx_table_s3_t0_d     d;

#define D(field)    d.{u.read_rx_head_desc_idx_d.##field}
#define K(field)    k.{##field}



%%

    .param      virtio_rx_read_head_desc_start

virtio_rx_read_head_desc_idx_start:

    /* FIXME: Assumes little-endian for now */
    add         r2, r0, D(desc_idx).hx

    /* Relay the descriptor to be written back into the virtq_used
     * ring 
     */
    phvwr       p.virtio_rx_global_phv_head_desc_idx, D(desc_idx).hx


    /* Compute address to the virtq_desc ring offset */
    add         r1, r0, K(to_s3_rx_virtq_desc_addr)
    /* r2 -> K(to_s1_rx_virtq_avail_ci) */
    add         r1, r1, r2, VIRTIO_VIRTQ_DESC_RING_ELEM_SHIFT

	CAPRI_NEXT_TABLE_READ(0, TABLE_LOCK_DIS,
                        virtio_rx_read_head_desc_start,
	                    r1, TABLE_SIZE_128_BITS)

virtio_rx_read_head_desc_idx_done:
	nop.e
    nop
