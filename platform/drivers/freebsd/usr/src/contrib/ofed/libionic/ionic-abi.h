/* SPDX-License-Identifier: GPL-2.0 OR Linux-OpenIB */
/*
 * Copyright (c) 2018-2020 Pensando Systems, Inc.  All rights reserved.
 *
 * This software is available to you under a choice of one of two
 * licenses.  You may choose to be licensed under the terms of the GNU
 * General Public License (GPL) Version 2, available from the file
 * COPYING in the main directory of this source tree, or the
 * OpenIB.org BSD license below:
 *
 *     Redistribution and use in source and binary forms, with or
 *     without modification, are permitted provided that the following
 *     conditions are met:
 *
 *      - Redistributions of source code must retain the above
 *        copyright notice, this list of conditions and the following
 *        disclaimer.
 *
 *      - Redistributions in binary form must reproduce the above
 *        copyright notice, this list of conditions and the following
 *        disclaimer in the documentation and/or other materials
 *        provided with the distribution.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
 * BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

#ifndef __IONIC_ABI_H__
#define __IONIC_ABI_H__

#include <infiniband/kern-abi.h>

#include "ionic_fw.h"

#define IONIC_ABI_VERSION	2

struct uionic_ctx {
	struct ibv_get_context ibv_cmd;
	__u32 rsvd[2];
};

struct uionic_ctx_resp {
	struct ibv_get_context_resp ibv_resp;
	__u32 rsvd2;
	__u32 page_shift;

	__u64 dbell_offset;

	__u16 version;
	__u8 qp_opcodes;
	__u8 admin_opcodes;

	__u8 sq_qtype;
	__u8 rq_qtype;
	__u8 cq_qtype;
	__u8 admin_qtype;

	__u8 max_stride;
	__u8 max_spec;
	__u8 rsvd[2];
};

struct ionic_qdesc {
	__u64 addr;
	__u32 size;
	__u16 mask;
	__u8 depth_log2;
	__u8 stride_log2;
};

struct uionic_ah_resp {
	struct ibv_create_ah_resp ibv_resp;
	__u32 ahid;
	__u32 pad;
};

struct uionic_cq {
	struct ibv_create_cq ibv_cmd;
	struct ionic_qdesc cq;
};

struct uionic_cq_resp {
	struct ibv_create_cq_resp ibv_resp;
	__u32 cqid;
	__u32 pad;
};

struct uionic_qp {
	struct ibv_create_qp_ex ibv_cmd;
	struct ionic_qdesc sq;
	struct ionic_qdesc rq;
	__u8 sq_spec;
	__u8 rq_spec;
	__u8 rsvd[6];
};

struct uionic_qp_resp {
	struct ibv_create_qp_resp_ex ibv_resp;
	__u32 qpid;
	__u32 rsvd;
	__u64 sq_cmb_offset;
	__u64 rq_cmb_offset;
};

#endif /* __IONIC_ABI_H__ */
