#ifndef _LKL_ASM_UAPI_LKL_BITSPERLONG_H
#define _LKL_ASM_UAPI_LKL_BITSPERLONG_H

#ifdef LKL_CONFIG_64BIT
#define __LKL__BITS_PER_LONG 64
#else
#define __LKL__BITS_PER_LONG 32
#endif

#define __LKL__ARCH_WANT_STAT64

#endif /* _LKL_ASM_UAPI_LKL_BITSPERLONG_H */
