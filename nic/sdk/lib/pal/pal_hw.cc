// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

#include <sys/mman.h>
#include <dlfcn.h>
#include <unistd.h>
#include "lib/pal/pal.hpp"
#include "lib/pal/pal_internal.hpp"

namespace sdk {
namespace lib {

static void *gl_lib_handle;
extern pal_info_t   gl_pal_info;
typedef void (*hw_init_fn_t)(char *application_name);
typedef void (*reg_read_fn_t)(uint64_t addr, uint32_t *data, uint32_t nw);
typedef void (*reg_write_fn_t)(uint64_t addr, uint32_t *data, uint32_t nw);
typedef int (*mem_read_fn_t)(uint64_t addr, uint8_t * data, uint32_t size, uint32_t flags);
typedef int (*mem_write_fn_t)(uint64_t addr, uint8_t * data, uint32_t size, uint32_t flags);
typedef bool (*ring_doorbell_fn_t)(uint64_t addr, uint64_t data);
typedef uint64_t (*mem_vtop_fn_t)(const void *va);
typedef void *(*mem_ptov_fn_t)(const uint64_t pa);
typedef int (*memset_fn_t)(const uint64_t pa, uint8_t c, const size_t sz, uint32_t flags);

typedef struct pal_hw_vectors_s {
    hw_init_fn_t	    hw_init;
    reg_read_fn_t           reg_read;
    reg_write_fn_t          reg_write;
    mem_read_fn_t           mem_read;
    mem_write_fn_t          mem_write;
    mem_vtop_fn_t           mem_vtop;
    mem_ptov_fn_t           mem_ptov;
    memset_fn_t		    mem_set;
} pal_hw_vectors_t;

static pal_hw_vectors_t   gl_hw_vecs;

static pal_ret_t
pal_init_hw_vectors (void)
{
    gl_hw_vecs.hw_init = (hw_init_fn_t)dlsym(gl_lib_handle, "pal_init");
    SDK_ASSERT(gl_hw_vecs.hw_init);

    gl_hw_vecs.reg_read = (reg_read_fn_t)dlsym(gl_lib_handle, "pal_reg_rd32w");
    SDK_ASSERT(gl_hw_vecs.reg_read);

    gl_hw_vecs.reg_write = (reg_write_fn_t)dlsym(gl_lib_handle, "pal_reg_wr32w");
    SDK_ASSERT(gl_hw_vecs.reg_write);

    gl_hw_vecs.mem_read = (mem_read_fn_t)dlsym(gl_lib_handle, "pal_mem_rd");
    SDK_ASSERT(gl_hw_vecs.mem_read);

    gl_hw_vecs.mem_write = (mem_write_fn_t)dlsym(gl_lib_handle, "pal_mem_wr");
    SDK_ASSERT(gl_hw_vecs.mem_read);

    gl_hw_vecs.mem_vtop = (mem_vtop_fn_t)dlsym(gl_lib_handle, "pal_mem_vtop");
    SDK_ASSERT(gl_hw_vecs.mem_vtop);

    gl_hw_vecs.mem_ptov = (mem_ptov_fn_t)dlsym(gl_lib_handle, "pal_mem_ptov");
    SDK_ASSERT(gl_hw_vecs.mem_ptov);

    gl_hw_vecs.mem_set = (memset_fn_t)dlsym(gl_lib_handle, "pal_memset");
    SDK_ASSERT(gl_hw_vecs.mem_set);

    return PAL_RET_OK;
}

static inline pal_ret_t
pal_hw_physical_addr_to_virtual_addr(uint64_t phy_addr,
                                     uint64_t *virtual_addr)
{
    *virtual_addr = (uint64_t) (*gl_hw_vecs.mem_ptov)(phy_addr);
    return PAL_RET_OK;
}

static inline pal_ret_t
pal_hw_virtual_addr_to_physical_addr(uint64_t virtual_addr,
                                     uint64_t *phy_addr)
{
    *phy_addr = (*gl_hw_vecs.mem_vtop)((void*) virtual_addr);
    return PAL_RET_OK;
}

static pal_ret_t
pal_hw_reg_read (uint64_t addr, uint32_t *data, uint32_t num_words)
{
    (*gl_hw_vecs.reg_read)(addr, data, num_words); 
    
    return PAL_RET_OK;
}

static pal_ret_t
pal_hw_reg_write (uint64_t addr, uint32_t *data, uint32_t num_words)
{
    (*gl_hw_vecs.reg_write)(addr, data, num_words);
    return PAL_RET_OK;
}

static pal_ret_t
pal_hw_mem_read (uint64_t addr, uint8_t *data, uint32_t size, uint32_t flags)
{
    (*gl_hw_vecs.mem_read)(addr, data, size, flags);
    return PAL_RET_OK;
}

static pal_ret_t
pal_hw_mem_write (uint64_t addr, uint8_t *data, uint32_t size, uint32_t flags)
{
    (*gl_hw_vecs.mem_write)(addr, data, size, flags);
    return PAL_RET_OK;
}

static pal_ret_t
pal_hw_ring_doorbell (uint64_t addr, uint64_t data)
{
    uint64_t pa_doorbell = addr + 0x8000000;

    (*gl_hw_vecs.reg_write)(pa_doorbell, (uint32_t*) &data, 2);
    return PAL_RET_OK;
}

static pal_ret_t
pal_hw_memset (uint64_t pa, uint8_t c, uint32_t sz, uint32_t flags)
{
    uint32_t size_written = 0;

    size_written = (*gl_hw_vecs.mem_set)(pa, c, sz, flags);

    if (size_written != sz) {
	return PAL_RET_NOK;
    }
    
    return PAL_RET_OK;
}

static pal_ret_t
pal_hw_init_rwvectors (void)
{
    gl_pal_info.rwvecs.reg_read = pal_hw_reg_read;
    gl_pal_info.rwvecs.reg_write = pal_hw_reg_write;
    gl_pal_info.rwvecs.mem_read = pal_hw_mem_read;
    gl_pal_info.rwvecs.mem_write = pal_hw_mem_write;
    gl_pal_info.rwvecs.ring_doorbell = pal_hw_ring_doorbell;
    gl_pal_info.rwvecs.physical_addr_to_virtual_addr =
                        pal_hw_physical_addr_to_virtual_addr;
    gl_pal_info.rwvecs.virtual_addr_to_physical_addr =
                        pal_hw_virtual_addr_to_physical_addr;
    gl_pal_info.rwvecs.mem_set = pal_hw_memset;

    pal_init_hw_vectors();

    return PAL_RET_OK;
}

static pal_ret_t
pal_hw_dlopen (void)
{
    gl_lib_handle = dlopen("libpal.so", RTLD_NOW | RTLD_GLOBAL);
    SDK_ASSERT(gl_lib_handle);
    return PAL_RET_OK;
}

pal_ret_t
pal_hw_init (void)
{
    pal_ret_t   rv;

    rv = pal_hw_dlopen();
    SDK_ASSERT(IS_PAL_API_SUCCESS(rv));

    rv = pal_hw_init_rwvectors();
    SDK_ASSERT(IS_PAL_API_SUCCESS(rv));


    return PAL_RET_OK;
}

}    // namespace lib
}    // namespace sdk

