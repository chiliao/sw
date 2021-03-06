/*****************************************************************************
 *  nvme_be_sq_handler: Read the NVME backend priority submission queue entry.
 *                      Load the actual NVME command for the next stage.
 *****************************************************************************/

#include "storage_asm_defines.h"
#include "ingress.h"
#include "INGRESS_p.h"

struct s1_tbl_k k;
struct s1_tbl_nvme_be_sq_handler_d d;
struct phv_ p;

%%
   .param storage_tx_nvme_be_cmd_handler_start

storage_tx_nvme_be_sq_handler_start:

   // Save the R2N WQE to PHV
   R2N_WQE_FULL_COPY

   // Set the table and program address 
   add		r7, d.handle, NVME_BE_NVME_CMD_OFFSET
   LOAD_TABLE_FOR_ADDR_PC_IMM(r7, STORAGE_DEFAULT_TBL_LOAD_SIZE,
                              storage_tx_nvme_be_cmd_handler_start)
