/*****************************************************************************
 *  push_cq : Push the NVME status after setting the phase bit to the CQ.
 *            Alter the phase bit in the CQ context in event wrapping around
 *            the CQ.
 *****************************************************************************/

#include "storage_asm_defines.h"
#include "ingress.h"
#include "INGRESS_p.h"


struct s6_tbl0_k k;
struct s6_tbl0_push_cq_d d;
struct phv_ p;

%%
   .param storage_nvme_push_roce_rq_start
   .param storage_nvme_send_sta_free_iob_start

storage_nvme_push_cq_start:
   // Check queue full condition and exit
   // TODO: Push error handling
   //QUEUE_FULL(d.p_ndx, d.c_ndx, d.num_entries, exit)

   // Queue full condition based on Qstate does not work for PCI queue state push
   // as the c_ndx is in host memory in software. Need to read the c_ndx correctly
   // via another stage to perform this computation. 
   // For now in the SSD case, the throttling of max_cmds takes care of never 
   // overflowing this buffer.  
   // TODO: Need to come up with a good strategy for this.

   // Calculate the address to which the entry to be pushed has to be 
   // written to in the destination queue. Output will be stored in GPR r7.
   QUEUE_PUSH_ADDR(d.base_addr, d.p_ndx, d.entry_size)

   // DMA the NVME status entry to the CQ ring buffer
   DMA_PHV2MEM_SETUP_ADDR64(nvme_sta_cspec, nvme_sta_status, 
                            r7, dma_p2m_5)

   // Push the entry to the queue (this increments p_ndx and writes to table)
   QUEUE_PUSH(d.p_ndx, d.num_entries)

   // Store the ROCR RQ information in the K+I vector for releasing the
   // IO status buffer
   phvwrpair	p.nvme_kivec_t1_s2s_dst_lif, d.rrq_lif,	\
		p.nvme_kivec_t1_s2s_dst_qtype, d.rrq_qtype
   phvwrpair	p.nvme_kivec_t1_s2s_dst_qid, d.rrq_qid,	\
		p.nvme_kivec_t1_s2s_dst_qaddr, d.rrq_base

   // Delay slot to check the interrupt enable bit (since reading p_ndx 
   // straight after writing takes a one cycle stall)
   seq		c2, d.intr_en, 1

   // If new p_ndx has wrapped around, flip the phase bit
   seq		c1, d.p_ndx, 0
   tblmincri.c1	d.phase, 2, 1

   // Check if interrupt is enabled and branch
   bcf		![c2], skip_intr
   nop

   // Raise the interrupt with a DMA update
   PCI_RAISE_INTERRUPT(dma_p2m_6)

skip_intr:
   // Load table 1 and program for pushing the status buffer back to the RQ.
   LOAD_TABLE1_FOR_ADDR34_PC_IMM(d.rrq_qaddr,
                                 STORAGE_DEFAULT_TBL_LOAD_SIZE,
                                 storage_nvme_push_roce_rq_start)

   // Load table 0 and program for reading the destination arm queue
   // state in the next stage for sending command to free the IOB.
   LOAD_TABLE_FOR_ADDR34_PC_IMM(NVME_KIVEC_ARM_DST6_ARM_QADDR,
                                STORAGE_DEFAULT_TBL_LOAD_SIZE,
                                storage_nvme_send_sta_free_iob_start)

exit:
   // Exit pipeline
   LOAD_NO_TABLES
