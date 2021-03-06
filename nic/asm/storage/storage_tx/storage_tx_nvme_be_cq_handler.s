/*****************************************************************************
 *  nvme_be_cq_handler: Save the NVME status into PHV. Load the saved R2N WQE
 *                      for the command that was sent to the SSD
 *****************************************************************************/

#include "storage_asm_defines.h"
#include "ingress.h"
#include "INGRESS_p.h"

struct s1_tbl_k k;
struct s1_tbl_nvme_be_cq_handler_d d;
struct phv_ p;

%%
   .param storage_tx_nvme_be_wqe_handler_start

storage_tx_nvme_be_cq_handler_start:

   // Save the R2N WQE to PHV
   phvwr	p.{nvme_sta_cspec...nvme_sta_status}, d.{cspec...status}

   // Set the state information for the NVME backend status header
   // TODO: Is there any time where error status needs to be set ? 
   //       Set is_q0, time_us correctly.
   phvwri	p.{nvme_be_sta_hdr_time_us...nvme_be_sta_hdr_be_rsvd}, 0

   // Store the SSD's c_ndx value for DMA to the NVME backend SQ
   phvwr	p.ssd_ci_c_ndx, d.sq_head

   // Setup the DMA command to push the sq_head to the c_ndx of the SSD
   DMA_PHV2MEM_SETUP_ADDR_VEC(ssd_ci_c_ndx, ssd_ci_c_ndx, 
                              STORAGE_KIVEC6_SSD_CI_ADDR, dma_p2m_2)

   // Obtain the saved command index from the command id in the status
   // and save it in the PHV. Store the result in GPR r6 to pass as input
   // to SSD_CMD_ENTRY_ADDR_CALC
   and		r6, d.{cid}.hx, 0xFF
   phvwr	p.r2n_wqe_nvme_cmd_cid, r6
   phvwr	p.storage_kivec0_cmd_index, r6

   // Calculate the table address based on the command index offset into
   // the SSD's list of outstanding commands. Output is stored in GPR r7.
   SSD_CMD_ENTRY_ADDR_CALC

   // Set the table and program address 
   LOAD_TABLE_FOR_ADDR_PC_IMM(r7, STORAGE_DEFAULT_TBL_LOAD_SIZE,
                              storage_tx_nvme_be_wqe_handler_start)
