/*	$OpenBSD: xform.c,v 1.16 2001/08/28 12:20:43 ben Exp $	*/
/*-
 * The authors of this code are John Ioannidis (ji@tla.org),
 * Angelos D. Keromytis (kermit@csd.uch.gr),
 * Niels Provos (provos@physnet.uni-hamburg.de) and
 * Damien Miller (djm@mindrot.org).
 *
 * This code was written by John Ioannidis for BSD/OS in Athens, Greece,
 * in November 1995.
 *
 * Ported to OpenBSD and NetBSD, with additional transforms, in December 1996,
 * by Angelos D. Keromytis.
 *
 * Additional transforms and features in 1997 and 1998 by Angelos D. Keromytis
 * and Niels Provos.
 *
 * Additional features in 1999 by Angelos D. Keromytis.
 *
 * AES XTS implementation in 2008 by Damien Miller
 *
 * Copyright (C) 1995, 1996, 1997, 1998, 1999 by John Ioannidis,
 * Angelos D. Keromytis and Niels Provos.
 *
 * Copyright (C) 2001, Angelos D. Keromytis.
 *
 * Copyright (C) 2008, Damien Miller
 * Copyright (c) 2014 The FreeBSD Foundation
 * All rights reserved.
 *
 * Portions of this software were developed by John-Mark Gurney
 * under sponsorship of the FreeBSD Foundation and
 * Rubicon Communications, LLC (Netgate).
 *
 * Permission to use, copy, and modify this software with or without fee
 * is hereby granted, provided that this entire notice is included in
 * all copies of any software which is or includes a copy or
 * modification of this software.
 * You may use this code under the GNU public license if you so wish. Please
 * contribute changes back to the authors under this freer than GPL license
 * so that we may further the use of strong encryption without limitations to
 * all.
 *
 * THIS SOFTWARE IS BEING PROVIDED "AS IS", WITHOUT ANY EXPRESS OR
 * IMPLIED WARRANTY. IN PARTICULAR, NONE OF THE AUTHORS MAKES ANY
 * REPRESENTATION OR WARRANTY OF ANY KIND CONCERNING THE
 * MERCHANTABILITY OF THIS SOFTWARE OR ITS FITNESS FOR ANY PARTICULAR
 * PURPOSE.
 */

#if 0
#include <sys/cdefs.h>
__FBSDID("$FreeBSD$");
#endif

#include "osal_stdtypes.h"
#include "osal_mem.h"
#define KMALLOC(x, ...) osal_alloc(x)
#define KFREE(x, ...) osal_free(x)

#include "xform_enc.h"

#if defined(__FreeBSD_version) && (__FreeBSD_version >= 1200000)
#undef bcopy
#undef bzero
#endif

#define bcopy(a, b, c) memcpy(b, a, c)
#define bzero(a, b) memset(a, 0, b)

static	int aes_xts_setkey(uint8_t **, uint8_t *, int);
static	void aes_xts_encrypt(caddr_t, uint8_t *);
static	void aes_xts_decrypt(caddr_t, uint8_t *);
static	void aes_xts_zerokey(uint8_t **);
static	void aes_xts_reinit(caddr_t, uint8_t *);
static	void aes_xts_reinit2(caddr_t, uint8_t *);

/* Encryption instances */
struct enc_xform enc_xform_aes_xts = {
	CRYPTO_AES_XTS, "AES-XTS",
	AES_BLOCK_LEN, AES_XTS_IV_LEN, AES_XTS_MIN_KEY, AES_XTS_MAX_KEY,
	aes_xts_encrypt,
	aes_xts_decrypt,
	aes_xts_setkey,
	aes_xts_zerokey,
	aes_xts_reinit,
	aes_xts_reinit2,
};

/*
 * Encryption wrapper routines.
 */
static void
aes_xts_reinit(caddr_t key, uint8_t *iv)
{
	struct aes_xts_ctx *ctx = (struct aes_xts_ctx *)key;
	uint64_t blocknum;
	unsigned int i;

	/*
	 * Prepare tweak as E_k2(IV). IV is specified as LE representation
	 * of a 64-bit block number which we allow to be passed in directly.
	 */
	bcopy(iv, &blocknum, AES_XTS_IVSIZE);
	for (i = 0; i < AES_XTS_IVSIZE; i++) {
		ctx->tweak[i] = blocknum & 0xff;
		blocknum >>= 8;
	}
	/* Last 64 bits of IV are always zero */
	bzero(ctx->tweak + AES_XTS_IVSIZE, AES_XTS_IVSIZE);

	rijndael_encrypt(&ctx->key2, ctx->tweak, ctx->tweak);
}

/*
 * Encryption wrapper routines.
 */
static void
aes_xts_reinit2(caddr_t key, uint8_t *iv)
{
	struct aes_xts_ctx *ctx = (struct aes_xts_ctx *)key;
	uint64_t blocknum;
	unsigned int i;

	/*
	 * Prepare tweak as E_k2(IV). IV is specified as LE representation
	 * of a 64-bit block number which we allow to be passed in directly.
	 */
	bcopy(iv, &blocknum, AES_XTS_IVSIZE);
	for (i = 0; i < AES_XTS_IVSIZE; i++) {
		ctx->tweak[i] = blocknum & 0xff;
		blocknum >>= 8;
	}

	/* 2nd blocknum for full 16 byte iv */
	bcopy(iv + AES_XTS_IVSIZE, &blocknum, AES_XTS_IVSIZE);
	for (i = 0; i < AES_XTS_IVSIZE; i++) {
		ctx->tweak[AES_XTS_IVSIZE + i] = blocknum & 0xff;
		blocknum >>= 8;
	}

	rijndael_encrypt(&ctx->key2, ctx->tweak, ctx->tweak);
}

static void
aes_xts_crypt(struct aes_xts_ctx *ctx, uint8_t *data, unsigned int do_encrypt)
{
	uint8_t block[AES_XTS_BLOCKSIZE];
	unsigned int i, carry_in, carry_out;

	for (i = 0; i < AES_XTS_BLOCKSIZE; i++)
		block[i] = data[i] ^ ctx->tweak[i];

	if (do_encrypt)
		rijndael_encrypt(&ctx->key1, block, data);
	else
		rijndael_decrypt(&ctx->key1, block, data);

	for (i = 0; i < AES_XTS_BLOCKSIZE; i++)
		data[i] ^= ctx->tweak[i];

	/* Exponentiate tweak */
	carry_in = 0;
	for (i = 0; i < AES_XTS_BLOCKSIZE; i++) {
		carry_out = ctx->tweak[i] & 0x80;
		ctx->tweak[i] = (ctx->tweak[i] << 1) | (carry_in ? 1 : 0);
		carry_in = carry_out;
	}
	if (carry_in)
		ctx->tweak[0] ^= AES_XTS_ALPHA;
	bzero(block, sizeof(block));
}

static void
aes_xts_encrypt(caddr_t key, uint8_t *data)
{
	aes_xts_crypt((struct aes_xts_ctx *)key, data, 1);
}

static void
aes_xts_decrypt(caddr_t key, uint8_t *data)
{
	aes_xts_crypt((struct aes_xts_ctx *)key, data, 0);
}

static int
aes_xts_setkey(uint8_t **sched, uint8_t *key, int len)
{
	struct aes_xts_ctx *ctx;

	if (len != 32 && len != 64)
		return EINVAL;

	*sched = KMALLOC(sizeof(struct aes_xts_ctx), M_CRYPTO_DATA,
	    M_NOWAIT | M_ZERO);
	if (*sched == NULL)
		return ENOMEM;
	ctx = (struct aes_xts_ctx *)*sched;

	rijndael_set_key(&ctx->key1, key, len * 4);
	rijndael_set_key(&ctx->key2, key + (len / 2), len * 4);

	return 0;
}

static void
aes_xts_zerokey(uint8_t **sched)
{
	bzero(*sched, sizeof(struct aes_xts_ctx));
	KFREE(*sched, M_CRYPTO_DATA);
	*sched = NULL;
}
