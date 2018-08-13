/*
 * {C} Copyright 2018 Pensando Systems Inc.
 * All rights reserved.
 *
 */

#ifndef OSAL_MEM_H
#define OSAL_MEM_H

#include "osal_stdtypes.h"

#ifndef __KERNEL__
#include <string.h>
#else
#include <linux/string.h>
#endif

#ifdef __cplusplus
extern "C" {
#endif

void *osal_alloc(size_t size);
void *osal_aligned_alloc(size_t alignment, size_t size);
void osal_free(void *ptr);
uint64_t osal_virt_to_phy(void *ptr);
void *osal_phy_to_virt(uint64_t phy);

uint64_t osal_rmem_alloc(size_t size);
uint64_t osal_rmem_aligned_alloc(size_t alignment, size_t size);
void osal_rmem_free(uint64_t ptr, size_t size);

#ifdef __cplusplus
}
#endif

#endif
