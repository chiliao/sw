// {C} Copyright 2018 Pensando Systems Inc. All rights reserved

/*
 * capri_pxb_pcie.cc
 * Madhava Cheethirala (Pensando Systems)
 */

#include "include/sdk/base.hpp"
#include "include/sdk/platform/capri/capri_pxb_pcie.hpp"

#include "nic/asic/capri/model/utils/cap_blk_reg_model.h"
#include "nic/asic/capri/model/cap_top/cap_top_csr.h"
#include "nic/asic/capri/model/cap_pcie/cap_pxb_csr.h"

namespace sdk {
namespace platform {
namespace capri {

sdk_ret_t
capri_pxb_pcie_init ()
{
    cap_top_csr_t &cap0 = CAP_BLK_REG_MODEL_ACCESS(cap_top_csr_t, 0, 0);
    cap_pxb_csr_t &pxb_csr = cap0.pxb.pxb;

    SDK_TRACE_DEBUG("CAPRI-PXB::%s: Initializing LIF state for all of %d LIFs",
                    __func__, CAPRI_PCIE_MAX_LIFS);

    for (int i = 0; i < CAPRI_PCIE_MAX_LIFS; i++) {
        pxb_csr.dhs_itr_pcihdrt.entry[i].valid(1);
        pxb_csr.dhs_itr_pcihdrt.entry[i].write();
    }
    SDK_TRACE_DEBUG("CAPRI-PXB::%s: Initializing PCIE Atomic Region/Page as 0x%x/0x%x\n",
                    __func__, CAPRI_PCIE_ATOMIC_REGION_ID, CAPRI_PCIE_ATOMIC_PAGE_ID);

    // axi addressing formula :  
    //     {1 const (1bit), region - 4bit , page_id - 19bit , 12bit addr with page };
    // allocate region number 0xf and page 0x3ff within region as atomic
    // above formula will create id_0_addr : 0xf803ff000 to access atomic id[0]
    // id_n_addr = id_0_addr + (64*n)
    pxb_csr.cfg_pcie_local_memaddr_decode.atomic_page_id(CAPRI_PCIE_ATOMIC_PAGE_ID);
    pxb_csr.cfg_pcie_local_memaddr_decode.atomic(CAPRI_PCIE_ATOMIC_REGION_ID);
    pxb_csr.cfg_pcie_local_memaddr_decode.rc_cfg(0);
    pxb_csr.cfg_pcie_local_memaddr_decode.write();
    

    return sdk::SDK_RET_OK;
}

sdk_ret_t
capri_pxb_cfg_lif_bdf (uint32_t lif, uint16_t bdf)
{
    cap_top_csr_t &cap0 = CAP_BLK_REG_MODEL_ACCESS(cap_top_csr_t, 0, 0);
    cap_pxb_csr_t &pxb_csr = cap0.pxb.pxb;

    SDK_TRACE_DEBUG("CAPRI-PXB::%s: Configuring LIF %u with BDF %u",
                    __func__, lif, bdf);

    if (lif >= CAPRI_PCIE_MAX_LIFS) {
      SDK_TRACE_DEBUG("CAPRI-PXB::%s: LIF %u exceeded MAX_LIFS %u",
                      __func__, lif, CAPRI_PCIE_MAX_LIFS);
      return sdk::SDK_RET_ERR;
    }
    pxb_csr.dhs_itr_pcihdrt.entry[lif].bdf(bdf);
    pxb_csr.dhs_itr_pcihdrt.entry[lif].write();

    SDK_TRACE_DEBUG("CAPRI-PXB::%s: Successfully configured LIF %u with BDF %u",
                    __func__, lif, bdf);

    return sdk::SDK_RET_OK;
}

} // namespace capri
} // namespace platform
} // namespace sdk
