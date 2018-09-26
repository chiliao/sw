#ifndef SONIC_API_INT_H
#define SONIC_API_INT_H

#include <linux/types.h>
#include "accel_ring.h"


uint64_t sonic_rmem_alloc(size_t size);
uint64_t sonic_rmem_calloc(size_t size);
void sonic_rmem_free(uint64_t pgaddr, size_t size);
void sonic_rmem_set(uint64_t pgaddr, uint8_t val, size_t size);
void sonic_rmem_read(void *dst, uint64_t pgaddr, size_t size);
void sonic_rmem_write(uint64_t pgaddr, const void *src, size_t size);

accel_ring_t *sonic_get_accel_ring(uint32_t accel_ring_id);
uint64_t sonic_hostpa_to_devpa(uint64_t hostpa);
uint64_t sonic_devpa_to_hostpa(uint64_t devpa);

#endif /* SONIC_API_INT_H */
