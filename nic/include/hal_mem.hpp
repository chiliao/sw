#ifndef __HAL_MEM_HPP__
#define __HAL_MEM_HPP__

#include <base.h>

namespace hal {

// HAL memory slabs
typedef enum hal_slab_e {
    HAL_SLAB_NONE,
    HAL_SLAB_TENANT,
    HAL_SLAB_L2SEG,
    HAL_SLAB_LIF,
    HAL_SLAB_IF,
    HAL_SLAB_EP,
    HAL_SLAB_EP_IP_ENTRY,
    HAL_SLAB_EP_L3_ENTRY,
    HAL_SLAB_FLOW,
    HAL_SLAB_SESSION,
    HAL_SLAB_TCP_STATE,
    HAL_SLAB_SECURITY_PROFILE,
    HAL_SLAB_PI_MAX,                 // NOTE: MUST be last PI slab id

    // PD Slabs
    HAL_SLAB_TENANT_PD,
    HAL_SLAB_L2SEG_PD,
    HAL_SLAB_LIF_PD,
    HAL_SLAB_UPLINKIF_PD,
    HAL_SLAB_ENICIF_PD,
    HAL_SLAB_SECURITY_PROFILE_PD,
    HAL_SLAB_EP_PD,
    HAL_SLAB_EP_IP_ENTRY_PD,
    HAL_SLAB_SESSION_PD,
    HAL_SLAB_PD_MAX,                 // NOTE: MUST be last PD slab id

    HAL_SLAB_RSVD,    // all non-delay delete slabs can use this

    HAL_SLAB_MAX,
} hal_slab_t;

typedef enum hal_mem_alloc_e {
    HAL_MEM_ALLOC_NONE,
    HAL_MEM_ALLOC_LIB_HT,
    HAL_MEM_ALLOC_LIB_SLAB,
    HAL_MEM_ALLOC_LIB_BITMAP,
    HAL_MEM_ALLOC_LIB_TWHEEL,
    HAL_MEM_ALLOC_IF,
    HAL_MEM_ALLOC_L2,
    HAL_MEM_ALLOC_L3,
    HAL_MEM_ALLOC_EP,
    HAL_MEM_ALLOC_SFW,
    HAL_MEM_ALLOC_L4LB,
    HAL_MEM_ALLOC_FLOW,
    HAL_MEM_ALLOC_PD,
    HAL_MEM_ALLOC_OTHER,
} hal_mem_alloc_t;

hal_ret_t free_to_slab(hal_slab_t slab_id, void *elem);

namespace pd {

hal_ret_t free_to_slab(hal_slab_t slab_id, void *elem);

}    // namespace pd

}    // namespace hal

#endif    // __HAL_MEM_HPP__

