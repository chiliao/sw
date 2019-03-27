#ifndef __CAPRI_BARCO_CRYPTO_HPP__
#define __CAPRI_BARCO_CRYPTO_HPP__

#include "gen/proto/types.pb.h"
#include "nic/include/base.hpp"
#include "platform/capri/capri_cfg.hpp"

namespace hal {
namespace pd {

#define BARCO_CRYPTO_DESC_SZ                128 /* 1024 bits */
#define BARCO_CRYPTO_DESC_ALIGN_BYTES       128

#define BARCO_CRYPTO_KEY_DESC_SZ            16 /* 128 bits */
#define BARCO_CRYPTO_KEY_DESC_ALIGN_BYTES   16

/* FIXME: this needs to be driven from HAL PD, but the includes do not make it to capri */
#define CRYPTO_KEY_COUNT_MAX                (64 * 1024)

#define CAPRI_MAX_TLS_PAD_SIZE              512
#define BARCO_RING_SHADOW_PI_SIZE           2
#define BARCO_RING_SHADOW_CI_SIZE           2
#define BARCO_RING_QSTATS_SIZE              12

#define BARCO_GCM0_PI_HBM_TABLE_OFFSET      CAPRI_MAX_TLS_PAD_SIZE
#define BARCO_GCM0_CI_HBM_TABLE_OFFSET      (BARCO_GCM0_PI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_PI_SIZE)
#define BARCO_GCM0_QSTATS_HBM_TABLE_OFFSET  (BARCO_GCM0_CI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_CI_SIZE)
// GC is just using unused space in the Barco TLS pad region, can
// be moved elsewhere too
#define CAPRI_GC_GLOBAL_TABLE               (CAPRI_MAX_TLS_PAD_SIZE + 128)
#define CAPRI_GC_GLOBAL_RNMDPR_FP_PI        (CAPRI_GC_GLOBAL_TABLE + 0)
#define CAPRI_GC_GLOBAL_TNMDPR_FP_PI        (CAPRI_GC_GLOBAL_TABLE + 4)
#define CAPRI_GC_GLOBAL_OOQ_TX2RX_FP_PI     (CAPRI_GC_GLOBAL_TABLE + 8)
#define CAPRI_IPSEC_ENC_NMDR_ALLOC_PI       (CAPRI_GC_GLOBAL_TABLE + 16)
#define CAPRI_IPSEC_ENC_NMDR_ALLOC_CI       (CAPRI_GC_GLOBAL_TABLE + 20)
#define CAPRI_IPSEC_DEC_NMDR_ALLOC_PI       (CAPRI_GC_GLOBAL_TABLE + 24)
#define CAPRI_IPSEC_DEC_NMDR_ALLOC_CI       (CAPRI_GC_GLOBAL_TABLE + 28)

#define BARCO_GCM1_PI_HBM_TABLE_OFFSET      (BARCO_GCM0_QSTATS_HBM_TABLE_OFFSET + BARCO_RING_QSTATS_SIZE)
#define BARCO_GCM1_CI_HBM_TABLE_OFFSET      (BARCO_GCM1_PI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_PI_SIZE)
#define BARCO_GCM1_QSTATS_HBM_TABLE_OFFSET  (BARCO_GCM1_CI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_CI_SIZE)

#define BARCO_MPP0_PI_HBM_TABLE_OFFSET      (BARCO_GCM1_QSTATS_HBM_TABLE_OFFSET + BARCO_RING_QSTATS_SIZE)
#define BARCO_MPP0_CI_HBM_TABLE_OFFSET      (BARCO_MPP0_PI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_PI_SIZE)
#define BARCO_MPP0_QSTATS_HBM_TABLE_OFFSET  (BARCO_MPP0_CI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_CI_SIZE)

#define BARCO_MPP1_PI_HBM_TABLE_OFFSET      (BARCO_MPP0_QSTATS_HBM_TABLE_OFFSET + BARCO_RING_QSTATS_SIZE)
#define BARCO_MPP1_CI_HBM_TABLE_OFFSET      (BARCO_MPP0_PI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_PI_SIZE)
#define BARCO_MPP1_QSTATS_HBM_TABLE_OFFSET  (BARCO_MPP0_CI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_CI_SIZE)

#define BARCO_MPP2_PI_HBM_TABLE_OFFSET      (BARCO_MPP1_QSTATS_HBM_TABLE_OFFSET + BARCO_RING_QSTATS_SIZE)
#define BARCO_MPP2_CI_HBM_TABLE_OFFSET      (BARCO_MPP1_PI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_PI_SIZE)
#define BARCO_MPP2_QSTATS_HBM_TABLE_OFFSET  (BARCO_MPP1_CI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_CI_SIZE)

#define BARCO_MPP3_PI_HBM_TABLE_OFFSET      (BARCO_MPP2_QSTATS_HBM_TABLE_OFFSET + BARCO_RING_QSTATS_SIZE)
#define BARCO_MPP3_CI_HBM_TABLE_OFFSET      (BARCO_MPP2_PI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_PI_SIZE)
#define BARCO_MPP3_QSTATS_HBM_TABLE_OFFSET  (BARCO_MPP2_CI_HBM_TABLE_OFFSET + BARCO_RING_SHADOW_CI_SIZE)


hal_ret_t capri_barco_rings_init(platform_type_t platform);
hal_ret_t capri_barco_res_allocator_init(void);
hal_ret_t capri_barco_crypto_init(platform_type_t platform);
hal_ret_t capri_barco_init_key(uint32_t key_idx, uint64_t key_addr);
hal_ret_t capri_barco_setup_key(uint32_t key_idx, types::CryptoKeyType key_type, uint8_t *key,
        uint32_t key_size);
hal_ret_t capri_barco_read_key(uint32_t key_idx, types::CryptoKeyType *key_type,
        uint8_t *key, uint32_t *key_size);
hal_ret_t capri_barco_crypto_init_tls_pad_table(void);

/* Barco Crypto specific definitions */
typedef struct capri_barco_key_desc_s {
    uint64_t                key_address;
    uint32_t                key_type;
    uint32_t                reserved;
} __PACK__ capri_barco_key_desc_t;

#define CAPRI_BARCO_KEYTYPE_SHIFT       28
#define CAPRI_BARCO_KEYTYPE_AES         (0x0 << CAPRI_BARCO_KEYTYPE_SHIFT)
#define CAPRI_BARCO_KEYTYPE_AES128      (CAPRI_BARCO_KEYTYPE_AES | 0x0000010)
#define CAPRI_BARCO_KEYTYPE_AES192      (CAPRI_BARCO_KEYTYPE_AES | 0x0000018)
#define CAPRI_BARCO_KEYTYPE_AES256      (CAPRI_BARCO_KEYTYPE_AES | 0x0000020)
#define CAPRI_BARCO_KEYTYPE_DES         ((0x1 << CAPRI_BARCO_KEYTYPE_SHIFT) | (0x0000070))
#define CAPRI_BARCO_KEYTYPE_CHACHA20    ((0x2 << CAPRI_BARCO_KEYTYPE_SHIFT) | (0x0000020))
#define CAPRI_BARCO_KEYTYPE_POY1305     ((0x3 << CAPRI_BARCO_KEYTYPE_SHIFT) | (0x0000020))
#define CAPRI_BARCO_KEYTYPE_HMAC_TYPE   (0x4 << CAPRI_BARCO_KEYTYPE_SHIFT)
#define CAPRI_BARCO_KEYTYPE_HMAC_LEN_MASK   (0xfffffff)
#define CAPRI_BARCO_KEYTYPE_HMAC(len)   ((CAPRI_BARCO_KEYTYPE_HMAC_TYPE) | \
                                        (len & CAPRI_BARCO_KEYTYPE_HMAC_LEN_MASK))

/* Barco Crypto Asym specific definitions */
typedef struct capri_barco_asym_key_desc_s {
    uint64_t                key_param_list;
    uint32_t                command_reg;
    uint32_t                reserved;
} __PACK__ capri_barco_asym_key_desc_t;

/* Asymmetric/PKE Command Definitions */
#define CAPRI_BARCO_ASYM_CMD_CALCR2         0x80000000

#define CAPRI_BARCO_ASYM_CMD_FLAGB          0x40000000

#define CAPRI_BARCO_ASYM_CMD_FLAGA          0x20000000

#define CAPRI_BARCO_ASYM_CMD_SWAP_BYTES     0x10000000

#define CAPRI_BARCO_ASYM_CMD_BUFFER_SEL     0x08000000

#define CAPRI_BARCO_ASYM_CMD_R_AND_PROJ     0x02000000

#define CAPRI_BARCO_ASYM_CMD_R_AND_KE       0x01000000

#define CAPRI_BARCO_ASYM_CMD_SEL_CURVE_NO_ACCEL     0x00000000
#define CAPRI_BARCO_ASYM_CMD_SEL_CURVE_P256         0x00100000
#define CAPRI_BARCO_ASYM_CMD_SEL_CURVE_P384         0x00200000
#define CAPRI_BARCO_ASYM_CMD_SEL_CURVE_P521_E521    0x00300000
#define CAPRI_BARCO_ASYM_CMD_SEL_CURVE_P192         0x00400000
#define CAPRI_BARCO_ASYM_CMD_SEL_CURVE_CURVE25519_ED25519   0x00500000

#define CAPRI_BARCO_ASYM_CMD_SIZE_OF_OP_SHIFT   8
#define CAPRI_BARCO_ASYM_CMD_SIZE_OF_OP(size)   ((size - 1) << CAPRI_BARCO_ASYM_CMD_SIZE_OF_OP_SHIFT)

#define CAPRI_BARCO_ASYM_CMD_FIELD_GFP      0x00000000
#define CAPRI_BARCO_ASYM_CMD_FIELD_GF2M     0x00000080

/* Primitive Arithmetic Operations GF(p) and GF(2^m) */
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_NA             0x00
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_ADD        0x01
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_SUB        0x02
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_MUL_ODD_N  0x03
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_RED_ODD_N  0x04
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_DIV_ODD_N  0x05
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_INV_ODD_N  0x06
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_SQRT       0x07
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MUL            0x08
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_INV_EVEN_N 0x09
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_MOD_RED_EVEN_N 0x0a
#define CAPRI_BARCO_ASYM_CMD_OP_TYPE_CLEAR_DATA_MEM 0x0f

/* High-level RSA, CRT & DSA Operations - GF(p) only */
#define CAPRI_BARCO_ASYM_CMD_MOD_EXPO               0x10
#define CAPRI_BARCO_ASYM_CMD_RSA_PK_GEN             0x11
#define CAPRI_BARCO_ASYM_CMD_RSA_CRT_KEY_PARAM_GEN  0x12
#define CAPRI_BARCO_ASYM_CMD_RSA_CRT_DECRYPT        0x13
#define CAPRI_BARCO_ASYM_CMD_RSA_ENCRYPT            0x14
#define CAPRI_BARCO_ASYM_CMD_RSA_DECRYPT            0x15
#define CAPRI_BARCO_ASYM_CMD_RSA_SIG_GEN            0x16
#define CAPRI_BARCO_ASYM_CMD_RSA_SIG_VERIFY         0x17
#define CAPRI_BARCO_ASYM_CMD_DSA_KEY_GEN            0x18
#define CAPRI_BARCO_ASYM_CMD_DSA_SIG_GEN            0x19
#define CAPRI_BARCO_ASYM_CMD_DSA_SIG_VERIFY         0x1a
#define CAPRI_BARCO_ASYM_CMD_SRP_SERVER_SESS_KEY    0x1b
#define CAPRI_BARCO_ASYM_CMD_SRP_CLIENT_SESS_KEY    0x1c
#define CAPRI_BARCO_ASYM_CMD_RSA_HALF_CRT_RECOMB    0x1d
#define CAPRI_BARCO_ASYM_CMD_SRP_SERVER_PUB_KEY     0x1e
#define CAPRI_BARCO_ASYM_CMD_RSA_HALF_CRT_DECRYPT   0x1f

/* Primitive ECC & Check Point Operations GF(p) & GF(2m) */
#define CAPRI_BARCO_ASYM_CMD_ECC_POINT_DOUBLE       0x20
#define CAPRI_BARCO_ASYM_CMD_ECC_POINT_ADD          0x21
#define CAPRI_BARCO_ASYM_CMD_ECC_POINT_MUL          0x22
#define CAPRI_BARCO_ASYM_CMD_ECC_CHECK_A_AND_B      0x23
#define CAPRI_BARCO_ASYM_CMD_ECC_CHECK_N_NE_Q       0x24
#define CAPRI_BARCO_ASYM_CMD_ECC_CHECK_X_Y_LT_Q     0x25
#define CAPRI_BARCO_ASYM_CMD_ECC_CHECK_POINT_ON_CURVE   0x26
#define CAPRI_BARCO_ASYM_CMD_ECC_POINT_DECOMPRESS   0x27
#define CAPRI_BARCO_ASYM_CMD_ECC_MONT_POINT_MUL     0x28
#define CAPRI_BARCO_ASYM_CMD_SM2_SIG_GEN            0x2d
#define CAPRI_BARCO_ASYM_CMD_SM2_SIG_VERIFY         0x2e
#define CAPRI_BARCO_ASYM_CMD_SM2_KEX                0x2f


/* High-level ECC – ECDSA Operations GF(p) & GF(2m) */
#define CAPRI_BARCO_ASYM_ECDSA_SIG_GEN              0x30
#define CAPRI_BARCO_ASYM_ECDSA_SIG_VERIFY           0x31
#define CAPRI_BARCO_ASYM_ECDSA_DOMAIN_PARAM_VALIDATION  0x32
#define CAPRI_BARCO_ASYM_ECKCDSA_PUB_KEY_GEN        0x33
#define CAPRI_BARCO_ASYM_ECKCDSA_SIG_GEN            0x34
#define CAPRI_BARCO_ASYM_ECKCDSA_SIG_VERIFY         0x35
#define CAPRI_BARCO_ASYM_JPAKE_GEN_ZKP              0x36
#define CAPRI_BARCO_ASYM_JPAKE_VERIFY_ZKP           0x37
#define CAPRI_BARCO_ASYM_JPAKE_2_POINT_ADD          0x38
#define CAPRI_BARCO_ASYM_JPAKE_GEN_SESS_KEY         0x39
#define CAPRI_BARCO_ASYM_JPAKE_GEN_STEP_2           0x3a
#define CAPRI_BARCO_ASYM_EDDSA_POINT_MUL            0x3b
#define CAPRI_BARCO_ASYM_EDDSA_SIG_GEN              0x3c
#define CAPRI_BARCO_ASYM_EDDSA_SIG_VERIFY           0x3d
#define CAPRI_BARCO_ASYM_EDDSA_GEN_SESS_KEY         0x3e

/* Primality Test – Rabin-Miller */
#define CAPRI_BARCO_ASYM_ROUND_RABIN_MILLER         0x40
#define CAPRI_BARCO_ASYM_INIT_RABIN_MILLER          0x41


/*
 * Command encoding for Barco symmetric crypto requests.
 * Based on Table 5. in Pensando_CryptoDMA_1.18.pdf.
 */
#define CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT  28
#define CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT      24
#define CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT      20

#define CAPRI_BARCO_SYM_COMMAND_ID_BA411E_AES       (0x0 << CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT)
#define CAPRI_BARCO_SYM_COMMAND_ID_BA412_DES        (0x1 << CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT)
#define CAPRI_BARCO_SYM_COMMAND_ID_BA413_HASH       (0x2 << CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT)
#define CAPRI_BARCO_SYM_COMMAND_ID_BA415_AES_GCM    (0x3 << CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT)
#define CAPRI_BARCO_SYM_COMMAND_ID_BA416_AES_XTS    (0x4 << CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT)
#define CAPRI_BARCO_SYM_COMMAND_ID_BA417_CHACHAPOLY (0x5 << CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT)
#define CAPRI_BARCO_SYM_COMMAND_ID_BA418_SHA3       (0x6 << CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT)
#define CAPRI_BARCO_SYM_COMMAND_ID_AES_HASH         (0x7 << CAPRI_BARCO_SYM_COMMAND_BARCO_ID_SHIFT)

#define CAPRI_BARCO_SYM_COMMAND_ALGO_AES_CBC     \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA411E_AES       | (0x1 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_AES_CCM     \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA411E_AES       | (0x5 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_DES_ECB     \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA412_DES        | (0x0 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_DES_CBC				\
  (CAPRI_BARCO_SYM_COMMAND_ID_BA412_DES        | (0x1 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA1   \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA413_HASH       | (0x1 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA224 \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA413_HASH       | (0x2 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA256 \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA413_HASH       | (0x3 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA384 \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA413_HASH       | (0x4 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA512 \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA413_HASH       | (0x5 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_AES_GCM     \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA415_AES_GCM    | (0x0 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_AES_XTS     \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA416_AES_XTS    | (0x0 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_CHACHA20    \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA417_CHACHAPOLY | (0x0 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_CHACHA20_POLY1305   \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA417_CHACHAPOLY | (0x1 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_224    \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA418_SHA3       | (0x0 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_256    \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA418_SHA3       | (0x1 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_384    \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA418_SHA3       | (0x2 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_512    \
  (CAPRI_BARCO_SYM_COMMAND_ID_BA418_SHA3       | (0x3 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA256   \
  (CAPRI_BARCO_SYM_COMMAND_ID_AES_HASH         | (0x1 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA384   \
  (CAPRI_BARCO_SYM_COMMAND_ID_AES_HASH         | (0x2 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA256_CBC   \
  (CAPRI_BARCO_SYM_COMMAND_ID_AES_HASH         | (0x3 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA384_CBC   \
  (CAPRI_BARCO_SYM_COMMAND_ID_AES_HASH         | (0x4 << CAPRI_BARCO_SYM_COMMAND_ALGO_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_AES_CBC_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_CBC | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_CBC_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_CBC | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_AES_CCM_NONCE_SHIFT      8
#define CAPRI_BARCO_SYM_COMMAND_AES_CCM_NONCE_SIZE_MASK  (0xf)
#define CAPRI_BARCO_SYM_COMMAND_AES_CCM_TAG_SIZE_MASK   (0xff)
#define CAPRI_BARCO_SYM_COMMAND_AES_CCM_Encrypt(Nonce_size, Tag_size) \
                (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_CCM |                               \
                 (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT) |                   \
		 ((Nonce_size & CAPRI_BARCO_SYM_COMMAND_AES_CCM_NONCE_SIZE_MASK) \
		  << CAPRI_BARCO_SYM_COMMAND_AES_CCM_NONCE_SHIFT) |              \
		 (Tag_size & CAPRI_BARCO_SYM_COMMAND_AES_CCM_TAG_SIZE_MASK))

#define CAPRI_BARCO_SYM_COMMAND_AES_CCM_Decrypt(Nonce_size, Tag_size)                \
                (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_CCM |                               \
                 (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT) |                   \
		 ((Nonce_size & CAPRI_BARCO_SYM_COMMAND_AES_CCM_NONCE_SIZE_MASK) \
		  << CAPRI_BARCO_SYM_COMMAND_AES_CCM_NONCE_SHIFT) |              \
		 (Tag_size & CAPRI_BARCO_SYM_COMMAND_AES_CCM_TAG_SIZE_MASK))

                                                 
#define CAPRI_BARCO_SYM_COMMAND_DES_ECB_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_DES_ECB | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_DES_ECB_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_DES_ECB | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))


#define CAPRI_BARCO_SYM_COMMAND_DES_CBC_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_DES_CBC | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_DES_CBC_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_DES_CBC | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA1_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA1 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA1_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA1 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA1_Generate_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA1 | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA1_Verify_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA1 | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA224_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA224 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA224_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA224 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA224_Generate_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA224 | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA224_Verify_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA224 | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA256_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA256 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA256_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA256 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA256_Generate_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA256 | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA256_Verify_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA256 | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA384_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA384 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA384_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA384 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA384_Generate_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA384 | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA384_Verify_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA384 | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA512_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA512 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA512_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA512 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA512_Generate_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA512 | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA512_Verify_HMAC \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_HASH_SHA512 | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_AES_GCM_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_GCM | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_GCM_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_GCM | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_AES_XTS_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_XTS | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_XTS_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_XTS | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_CHACHA20_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_CHACHA20 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_CHACHA20_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_CHACHA20 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_CHACHA20_POLY1305_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_CHACHA20_POLY1305 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_CHACHA20_POLY1305_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_CHACHA20_POLY1305 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA3_224_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_224 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA3_224_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_224 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA3_256_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_256 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA3_256_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_256 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA3_384_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_384 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA3_384_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_384 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_SHA3_512_Generate_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_512 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_SHA3_512_Verify_Hash \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_SHA3_512 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_CBC_SHA256_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA256 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_CBC_SHA256_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA256 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_CBC_HMAC_SHA256_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA256 | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_CBC_HMAC_SHA256_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA256 | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))


#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_CBC_SHA384_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA384 | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_CBC_SHA384_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA384 | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_CBC_HMAC_SHA384_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA384 | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_CBC_HMAC_SHA384_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_CBC_SHA384 | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_SHA256_CBC_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA256_CBC | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_SHA256_CBC_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA256_CBC | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_HMAC_SHA256_CBC_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA256_CBC | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_HMAC_SHA256_CBC_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA256_CBC | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_SHA384_CBC_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA384_CBC | (0x0 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_SHA384_CBC_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA384_CBC | (0x1 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_HMAC_SHA384_CBC_Encrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA384_CBC | (0x2 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))
#define CAPRI_BARCO_SYM_COMMAND_AES_HASH_HMAC_SHA384_CBC_Decrypt \
  (CAPRI_BARCO_SYM_COMMAND_ALGO_AES_HASH_SHA384_CBC | (0x3 << CAPRI_BARCO_SYM_COMMAND_OPER_SHIFT))

/*
 * Barco Symmetric Crypto Error encoding.
 * Based on Table 8. in Pensando_CryotoDMA_1.18.pdf.
 */
#define CAPRI_BARCO_SYM_ERR_BUS_ERR            (1 << 9)
#define CAPRI_BARCO_SYM_ERR_GEN_PUSH_ERR       (1 << 8)
#define CAPRI_BARCO_SYM_ERR_GEN_FETCH_ERR      (1 << 7)
#define CAPRI_BARCO_SYM_ERR_BUS_UNSUP_MODE     (1 << 6)
#define CAPRI_BARCO_SYM_ERR_BUS_RSVD           (1 << 5)
#define CAPRI_BARCO_SYM_ERR_BUS_BAD_CMD        (1 << 4)
#define CAPRI_BARCO_SYM_ERR_BUS_UNK_STATE      (1 << 3)
#define CAPRI_BARCO_SYM_ERR_BUS_AXI_BUS_RESP   (1 << 2)
#define CAPRI_BARCO_SYM_ERR_BUS_WRONG_KEYTYPE  (1 << 1)
#define CAPRI_BARCO_SYM_ERR_BUS_KEYTYPE_RANGE  (1 << 0)

}    // namespace pd
}    // namespace hal

#endif /*  __CAPRI_BARCO_CRYPTO_HPP__ */
