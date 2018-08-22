/*
 * Copyright (c) 2018 Pensando Systems, Inc.  All rights reserved.
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

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>
#include <errno.h>
#include <sys/mman.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>

#include "ionic.h"

static struct ibv_cq *ionic_create_cq(struct ibv_context *ibctx, int ncqe,
				      struct ibv_comp_channel *channel, int vec)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibctx);
	struct ionic_cq *cq;
	struct uionic_cq req = {};
	struct uionic_cq_resp resp = {};

	int rc;

	/* XXX hardcode value */
	if (ncqe > 0xffff) {
		rc = EINVAL;
		goto err;
	}

	cq = calloc(1, sizeof(*cq));
	if (!cq) {
		rc = ENOMEM;
		goto err;
	}

	pthread_spin_init(&cq->lock, PTHREAD_PROCESS_PRIVATE);
	list_head_init(&cq->poll_sq);
	list_head_init(&cq->flush_sq);
	list_head_init(&cq->flush_rq);

	rc = ionic_queue_init(&cq->q, ctx->pg_shift, ncqe,
			      sizeof(struct ionic_v1_cqe));
	if (rc)
		goto err_queue;

	ionic_queue_color_init(&cq->q);

	req.cq.addr = (uintptr_t)cq->q.ptr;
	req.cq.size = cq->q.size;
	req.cq.mask = cq->q.mask;
	req.cq.depth_log2 = cq->q.depth_log2;
	req.cq.stride_log2 = cq->q.stride_log2;

	req.compat = ctx->compat;

	rc = ibv_cmd_create_cq(ibctx, ncqe, channel, vec, &cq->ibcq,
			       &req.ibv_cmd, sizeof(req),
			       &resp.ibv_resp, sizeof(resp));
	if (rc)
		goto err_cmd;

	cq->cqid = resp.cqid;

	ionic_queue_dbell_init(&cq->q, cq->cqid);

	return &cq->ibcq;

err_cmd:
	ionic_queue_destroy(&cq->q);
err_queue:
	pthread_spin_destroy(&cq->lock);
	free(cq);
err:
	errno = rc;
	return NULL;
}

static int ionic_resize_cq(struct ibv_cq *ibcq, int ncqe)
{
	return -ENOSYS;
}

static int ionic_destroy_cq(struct ibv_cq *ibcq)
{
	struct ionic_cq *cq = to_ionic_cq(ibcq);
	int rc;

	rc = ibv_cmd_destroy_cq(ibcq);
	if (rc)
		return rc;

	ionic_queue_destroy(&cq->q);
	pthread_spin_destroy(&cq->lock);
	free(cq);

	return 0;
}

static int ionic_flush_recv(struct ionic_qp *qp, struct ibv_wc *wc)
{
	struct ionic_v1_recv_wqe *wqe;
	struct ionic_rq_meta *meta;

	if (!qp->rq_flush)
		return 0;

	if (ionic_queue_empty(&qp->rq))
		return 0;

	/* This depends on the RQ polled in-order.  It does not work for SRQ,
	 * which can be polled out-of-order.  Driver does not flush SRQ.
	 */
	wqe = ionic_queue_at_cons(&qp->rq);

	/* wqe_id must be a valid queue index */
	if (unlikely(wqe->base.wqe_id >> qp->rq.depth_log2))
		return -EIO;

	/* wqe_id must indicate a request that is outstanding */
	meta = &qp->rq_meta[wqe->base.wqe_id];
	if (unlikely(meta->next != IONIC_META_POSTED))
		return -EIO;

	ionic_queue_consume(&qp->rq);

	memset(wc, 0, sizeof(*wc));

	wc->status = IBV_WC_WR_FLUSH_ERR;
	wc->wr_id = meta->wrid;
	wc->qp_num = qp->qpid;

	meta->next = qp->rq_meta_head;
	qp->rq_meta_head = meta;

	return 1;
}

static int ionic_flush_recv_many(struct ionic_qp *qp, struct ibv_wc *wc, int nwc)
{
	int rc = 0, npolled = 0;

	while(npolled < nwc) {
		rc = ionic_flush_recv(qp, wc + npolled);
		if (rc <= 0)
			break;

		npolled += rc;
	}

	return npolled ?: rc;
}

static int ionic_flush_send(struct ionic_qp *qp, struct ibv_wc *wc)
{
	struct ionic_sq_meta *meta;

	if (!qp->sq_flush)
		return 0;

	if (ionic_queue_empty(&qp->sq))
		return 0;

	meta = &qp->sq_meta[qp->sq.cons];

	ionic_queue_consume(&qp->sq);

	memset(wc, 0, sizeof(*wc));

	wc->status = IBV_WC_WR_FLUSH_ERR;
	wc->wr_id = meta->wrid;
	wc->qp_num = qp->qpid;

	return 1;
}

static int ionic_flush_send_many(struct ionic_qp *qp, struct ibv_wc *wc, int nwc)
{
	int rc = 0, npolled = 0;

	while(npolled < nwc) {
		rc = ionic_flush_send(qp, wc + npolled);
		if (rc <= 0)
			break;

		npolled += rc;
	}

	return npolled ?: rc;
}

static int ionic_poll_recv(struct ionic_ctx *ctx, struct ionic_cq *cq,
			   struct ionic_qp *cqe_qp, struct ionic_v1_cqe *cqe,
			   struct ibv_wc *wc)
{
	struct ionic_qp *qp = NULL;
	struct ionic_rq_meta *meta;
	uint32_t src_qpn;
	uint8_t op;

	if (cqe_qp->rq_flush)
		return 0;

	if (cqe_qp->has_rq) {
		qp = cqe_qp;
	} else {
		if (unlikely(cqe_qp->is_srq))
			return -EIO;

		if (unlikely(!cqe_qp->vqp.qp.srq))
			return -EIO;

		qp = to_ionic_srq(cqe_qp->vqp.qp.srq);
	}

	/* there had better be something in the recv queue to complete */
	if (ionic_queue_empty(&qp->rq))
		return -EIO;

	/* wqe_id must be a valid queue index */
	if (unlikely(cqe->recv.wqe_id >> qp->rq.depth_log2))
		return -EIO;

	/* wqe_id must indicate a request that is outstanding */
	meta = &qp->rq_meta[cqe->recv.wqe_id];
	if (unlikely(meta->next != IONIC_META_POSTED))
		return -EIO;

	meta->next = qp->rq_meta_head;
	qp->rq_meta_head = meta;

	memset(wc, 0, sizeof(*wc));

	wc->wr_id = meta->wrid;
	wc->qp_num = cqe_qp->qpid;

	if (ionic_v1_cqe_error(cqe)) {
		wc->vendor_err = le32toh(cqe->status_length);
		wc->status = ionic_to_ibv_status(wc->vendor_err);

		cqe_qp->rq_flush = !qp->is_srq;
		if (cqe_qp->rq_flush) {
			list_del(&cqe_qp->cq_flush_rq);
			list_add_tail(&cq->flush_rq, &cqe_qp->cq_flush_rq);
		}
		goto out;
	}

	wc->vendor_err = 0;
	wc->status = IBV_WC_SUCCESS;

	src_qpn = be32toh(cqe->recv.src_qpn_op);
	op = src_qpn >> IONIC_V1_CQE_RECV_OP_SHIFT;

	/* XXX fixup op: cqe has recv flags in qtf, not all in srq_qpn_op */
	if (op == OP_TYPE_RDMA_OPER_WITH_IMM) {
		op = IONIC_V1_CQE_RECV_OP_RDMA_IMM;
	} else {
		op = IONIC_V1_CQE_RECV_OP_SEND;
		if (cqe->qid_type_flags & htobe32(IONIC_V1_CQE_RCVD_WITH_IMM))
			op = IONIC_V1_CQE_RECV_OP_SEND_IMM;
		else if (cqe->qid_type_flags & htobe32(IONIC_V1_CQE_RCVD_WITH_INV))
			op = IONIC_V1_CQE_RECV_OP_SEND_INV;
	}

	wc->opcode = IBV_WC_RECV;
	switch (op) {
	case IONIC_V1_CQE_RECV_OP_RDMA_IMM:
		wc->opcode = IBV_WC_RECV_RDMA_WITH_IMM;
		/* fallthrough */
	case IONIC_V1_CQE_RECV_OP_SEND_IMM:
		wc->wc_flags |= IBV_WC_WITH_IMM;
		wc->imm_data = cqe->recv.imm_data_rkey; /* be32 in wc */
		break;
	case IONIC_V1_CQE_RECV_OP_SEND_INV:
		wc->wc_flags |= IBV_WC_WITH_INV;
		wc->invalidated_rkey = be32toh(cqe->recv.imm_data_rkey);
	}

	wc->byte_len = le32toh(cqe->status_length);
	wc->byte_len = meta->len; /* XXX byte_len must come from cqe */
	wc->src_qp = src_qpn & IONIC_V1_CQE_RECV_QPN_MASK;
	wc->pkey_index = be16toh(cqe->recv.pkey_index);

	/* XXX: also need from cqe... slid, sl, dlid_path_bits */
	wc->slid = 0;
	wc->sl = 0;
	wc->dlid_path_bits = 0;

out:
	ionic_queue_consume(&qp->rq);

	return 1;
}

static int ionic_poll_send(struct ionic_cq *cq, struct ionic_qp *qp,
			   struct ibv_wc *wc)
{
	struct ionic_sq_meta *meta;

	if (qp->sq_flush)
		return 0;

	do {
		/* completed all send queue requests? */
		if (ionic_queue_empty(&qp->sq))
			return 0;

		/* waiting for local completion? */
		if (qp->sq.cons == qp->sq_npg_cons)
			return 0;

		meta = &qp->sq_meta[qp->sq.cons];

		/* waiting for remote completion? */
		if (meta->remote && meta->seq == qp->sq_msn_cons)
			return 0;

		ionic_queue_consume(&qp->sq);

		/* produce wc only if signaled or error status */
	} while (!meta->signal && meta->ibsts == IBV_WC_SUCCESS);

	memset(wc, 0, sizeof(*wc));

	wc->status = meta->ibsts;
	wc->wr_id = meta->wrid;
	wc->qp_num = qp->qpid;

	if (meta->ibsts == IBV_WC_SUCCESS) {
		wc->byte_len = meta->len;
		wc->opcode = meta->ibop;
	} else {
		wc->vendor_err = meta->len;

		qp->sq_flush = true;
		list_del(&qp->cq_flush_sq);
		list_add_tail(&cq->flush_sq, &qp->cq_flush_sq);
	}

	return 1;
}

static int ionic_poll_send_many(struct ionic_cq *cq, struct ionic_qp *qp,
				struct ibv_wc *wc, int nwc)
{
	int rc = 0, npolled = 0;

	while(npolled < nwc) {
		rc = ionic_poll_send(cq, qp, wc + npolled);
		if (rc <= 0)
			break;

		npolled += rc;
	}

	return npolled ?: rc;
}

static int ionic_validate_cons(uint16_t prod, uint16_t cons,
			       uint16_t comp, uint16_t mask)
{
	if (((prod - cons) & mask) < ((comp - cons) & mask))
		return -EIO;

	return 0;
}

static int ionic_comp_msn(struct ionic_qp *qp, struct ionic_v1_cqe *cqe)
{
	struct ionic_sq_meta *meta;
	uint16_t cqe_seq, cqe_idx;
	int rc;

	cqe_seq = be32toh(cqe->send.msg_msn) & qp->sq.mask;

	if (ionic_v1_cqe_error(cqe))
		cqe_idx = qp->sq_msn_idx[cqe_seq];
	else
		cqe_idx = qp->sq_msn_idx[(cqe_seq - 1) & qp->sq.mask];

	rc = ionic_validate_cons(qp->sq_msn_prod,
				 qp->sq_msn_cons,
				 cqe_seq, qp->sq.mask);
	if (rc)
		return rc;

	qp->sq_msn_cons = cqe_seq;

	if (ionic_v1_cqe_error(cqe)) {
		meta = &qp->sq_meta[cqe_idx];
		meta->len = le32toh(cqe->status_length);
		meta->ibsts = ionic_to_ibv_status(meta->len);
		meta->remote = false;
	}

	/* remote completion coalesces local requests, too */
	cqe_seq = ionic_queue_next(&qp->sq, cqe_idx);
	if (!ionic_validate_cons(qp->sq.prod,
				 qp->sq_npg_cons,
				 cqe_seq, qp->sq.mask))
		qp->sq_npg_cons = cqe_seq;

	return 0;
}

static int ionic_comp_npg(struct ionic_qp *qp, struct ionic_v1_cqe *cqe)
{
	struct ionic_sq_meta *meta;
	uint16_t cqe_seq, cqe_idx;
	int rc;

	cqe_idx = cqe->send.npg_wqe_id & qp->sq.mask;
	cqe_seq = ionic_queue_next(&qp->sq, cqe_idx);

	rc = ionic_validate_cons(qp->sq.prod,
				 qp->sq_npg_cons,
				 cqe_seq, qp->sq.mask);
	if (rc)
		return rc;

	qp->sq_npg_cons = cqe_seq;

	if (ionic_v1_cqe_error(cqe)) {
		meta = &qp->sq_meta[cqe_idx];
		meta->len = le32toh(cqe->status_length);
		meta->ibsts = ionic_to_ibv_status(meta->len);
		meta->remote = false;
	}

	return 0;
}

static bool ionic_next_cqe(struct ionic_cq *cq, struct ionic_v1_cqe **cqe)
{
	struct ionic_ctx *ctx = to_ionic_ctx(cq->ibcq.context);
	struct ionic_v1_cqe *qcqe = ionic_queue_at_prod(&cq->q);

	if (ionic_queue_color(&cq->q) != ionic_v1_cqe_color(qcqe))
		return false;

	udma_from_device_barrier();

	ionic_dbg(ctx, "poll cq %u prod %u", cq->cqid, cq->q.prod);
	ionic_dbg_xdump(ctx, "cqe", qcqe, 1u << cq->q.stride_log2);

	*cqe = qcqe;

	return true;
}

static void ionic_clean_cq(struct ionic_cq *cq, uint32_t qpid)
{
	struct ionic_v1_cqe *qcqe;
	int prod, qtf, qid;
	bool color;

	prod = cq->q.prod;
	qcqe = ionic_queue_at(&cq->q, prod);
	color = ionic_queue_color(&cq->q);

	while (color == ionic_v1_cqe_color(qcqe)) {
		qtf = ionic_v1_cqe_qtf(qcqe);
		qid = ionic_v1_cqe_qtf_qid(qtf);

		if (qid == qpid)
			ionic_v1_cqe_clean(qcqe);

		prod = ionic_queue_next(&cq->q, prod);
		qcqe = ionic_queue_at(&cq->q, prod);
		color = color != (prod == 0);
	}
}

static int ionic_poll_cq(struct ibv_cq *ibcq, int nwc, struct ibv_wc *wc)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibcq->context);
	struct ionic_cq *cq = to_ionic_cq(ibcq);
	struct ionic_qp *qp, *qp_next;
	struct ionic_v1_cqe *cqe;
	uint32_t qtf, qid;
	uint8_t type;
	uint16_t old_prod;
	int rc = 0, npolled = 0;

	/* Note about rc: (noted here because poll is different)
	 *
	 * Functions without "poll" in the name, if they return an integer,
	 * return zero on success, or a positive error number.  Functions
	 * returning a pointer return NULL on error and set errno to a positve
	 * error number.
	 *
	 * Functions with "poll" in the name return negative error numbers, or
	 * greater or equal to zero number of completions on success.
	 */

	if (nwc < 1)
		return 0;

	pthread_spin_lock(&cq->lock);

	old_prod = cq->q.prod;

	/* poll already indicated work completions for send queue */

	list_for_each_safe(&cq->poll_sq, qp, qp_next, cq_poll_sq) {
		if (npolled == nwc)
			goto out;

		pthread_spin_lock(&qp->sq_lock);
		rc = ionic_poll_send_many(cq, qp, wc + npolled, nwc - npolled);
		pthread_spin_unlock(&qp->sq_lock);

		if (rc > 0)
			npolled += rc;
		else
			list_del_init(&qp->cq_poll_sq);
	}

	/* poll for more work completions */

	while (ionic_next_cqe(cq, &cqe)) {
		if (npolled == nwc)
			goto out;

		qtf = ionic_v1_cqe_qtf(cqe);
		qid = ionic_v1_cqe_qtf_qid(qtf);
		type = ionic_v1_cqe_qtf_type(qtf);

		qp = tbl_lookup(&ctx->qp_tbl, qid);

		if (unlikely(!qp)) {
			ionic_dbg(ctx, "missing qp for qid %#x", qid);
			goto cq_next;
		}

		switch(type) {
		case IONIC_V1_CQE_TYPE_RECV:
			pthread_spin_lock(&qp->rq_lock);
			rc = ionic_poll_recv(ctx, cq, qp, cqe, wc + npolled);
			pthread_spin_unlock(&qp->rq_lock);

			if (rc < 0)
				goto out;

			npolled += rc;
			break;

		case IONIC_V1_CQE_TYPE_SEND_MSN:
			pthread_spin_lock(&qp->sq_lock);
			rc = ionic_comp_msn(qp, cqe);
			pthread_spin_unlock(&qp->sq_lock);

			if (rc < 0)
				goto out;

			list_del(&qp->cq_poll_sq);
			list_add_tail(&cq->poll_sq, &qp->cq_poll_sq);
			break;

		case IONIC_V1_CQE_TYPE_SEND_NPG:
			pthread_spin_lock(&qp->sq_lock);
			rc = ionic_comp_npg(qp, cqe);
			pthread_spin_unlock(&qp->sq_lock);

			if (rc < 0)
				goto out;

			list_del(&qp->cq_poll_sq);
			list_add_tail(&cq->poll_sq, &qp->cq_poll_sq);
			break;

		default:
			ionic_dbg(ctx, "unexpected cqe type %u", type);

			rc = -EIO;
			goto out;
		}

cq_next:
		ionic_queue_produce(&cq->q);
		ionic_queue_color_wrap(&cq->q);
	}

	/* poll newly indicated work completions for send queue */

	list_for_each_safe(&cq->poll_sq, qp, qp_next, cq_poll_sq) {
		if (npolled == nwc)
			goto out;

		pthread_spin_lock(&qp->sq_lock);
		rc = ionic_poll_send_many(cq, qp, wc + npolled, nwc - npolled);
		pthread_spin_unlock(&qp->sq_lock);

		if (rc > 0)
			npolled += rc;
		else
			list_del_init(&qp->cq_poll_sq);
	}

	/* lastly, flush send and recv queues */

	list_for_each_safe(&cq->flush_sq, qp, qp_next, cq_flush_sq) {
		if (npolled == nwc)
			goto out;

		pthread_spin_lock(&qp->sq_lock);
		rc = ionic_flush_send_many(qp, wc + npolled, nwc - npolled);
		pthread_spin_unlock(&qp->sq_lock);

		if (rc > 0)
			npolled += rc;
		else
			list_del_init(&qp->cq_flush_sq);
	}

	list_for_each_safe(&cq->flush_rq, qp, qp_next, cq_flush_rq) {
		if (npolled == nwc)
			goto out;

		pthread_spin_lock(&qp->rq_lock);
		rc = ionic_flush_recv_many(qp, wc + npolled, nwc - npolled);
		pthread_spin_unlock(&qp->rq_lock);

		if (rc > 0)
			npolled += rc;
		else
			list_del_init(&qp->cq_flush_rq);
	}

out:
	if (likely(cq->q.prod != old_prod)) {
		ionic_dbg(ctx, "dbell qtype %d val %#lx",
			  ctx->cq_qtype, ionic_queue_dbell_val(&cq->q));
		ionic_dbell_ring(&ctx->dbpage[ctx->cq_qtype],
				 ionic_queue_dbell_val(&cq->q));
	}

	pthread_spin_unlock(&cq->lock);

	return npolled ?: rc;
}

static int ionic_req_notify_cq(struct ibv_cq *ibcq, int solicited_only)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibcq->context);
	struct ionic_cq *cq = to_ionic_cq(ibcq);
	uint64_t dbell_val = cq->q.dbell;

	if (solicited_only) {
		cq->arm_sol_prod = ionic_queue_next(&cq->q, cq->arm_sol_prod);
		dbell_val |= cq->arm_sol_prod | IONIC_DBELL_RING_SONLY;
	} else {
		cq->arm_any_prod = ionic_queue_next(&cq->q, cq->arm_any_prod);
		dbell_val |= cq->arm_any_prod | IONIC_DBELL_RING_ARM;
	}

	ionic_dbg(ctx, "dbell qtype %d val %#lx", ctx->cq_qtype, dbell_val);
	ionic_dbell_ring(&ctx->dbpage[ctx->cq_qtype], dbell_val);

	return 0;
}

static int ionic_check_qp_limits(struct ionic_ctx *ctx,
				 struct ibv_qp_cap *cap)
{
	struct ibv_device_attr devattr;
	int rc;

	/* XXX unnecessary system call */
	rc = ibv_query_device(&ctx->vctx.context, &devattr);
	if (rc)
		return rc;

	if (cap->max_send_sge > devattr.max_sge)
		return EINVAL;
	if (cap->max_recv_sge > devattr.max_sge)
		return EINVAL;
	if (cap->max_inline_data > IONIC_MAX_INLINE_SIZE)
		return EINVAL;
	if (cap->max_send_wr > devattr.max_qp_wr)
		cap->max_send_wr = devattr.max_qp_wr;
	if (cap->max_recv_wr > devattr.max_qp_wr)
		cap->max_recv_wr = devattr.max_qp_wr;

	return 0;
}

static int ionic_alloc_queues(struct ionic_ctx *ctx, struct ionic_qp *qp,
			      struct ibv_qp_cap *cap)
{
	uint16_t min_depth, min_stride;
	int rc, i;

	list_node_init(&qp->cq_poll_sq);
	list_node_init(&qp->cq_flush_sq);
	list_node_init(&qp->cq_flush_rq);

	qp->sq_flush = false;
	qp->rq_flush = false;

	if (qp->has_sq) {
		min_depth = cap->max_send_wr;
		min_stride = ionic_v1_send_wqe_min_size(cap->max_send_sge,
							cap->max_inline_data);

		rc = ionic_queue_init(&qp->sq, ctx->pg_shift,
				      min_depth, min_stride);
		if (rc)
			goto err_sq;

		qp->sq_hbm_ptr = NULL;
		qp->sq_hbm_prod = 0;

		qp->sq_meta = calloc((uint32_t)qp->sq.mask + 1,
				     sizeof(*qp->sq_meta));
		if (!qp->sq_meta) {
			rc = ENOMEM;
			goto err_sq_meta;
		}

		qp->sq_msn_idx = calloc((uint32_t)qp->sq.mask + 1,
					sizeof(*qp->sq_msn_idx));
		if (!qp->sq_msn_idx) {
			rc = ENOMEM;
			goto err_sq_msn;
		}

		pthread_spin_init(&qp->sq_lock, PTHREAD_PROCESS_PRIVATE);
	}

	if (qp->has_rq) {
		min_depth = cap->max_recv_wr;
		min_stride = ionic_v1_recv_wqe_min_size(cap->max_recv_sge);

		rc = ionic_queue_init(&qp->rq, ctx->pg_shift,
				      min_depth, min_stride);
		if (rc)
			goto err_rq;

		qp->rq_meta = calloc((uint32_t)qp->rq.mask + 1,
				     sizeof(*qp->rq_meta));
		if (!qp->rq_meta) {
			rc = ENOMEM;
			goto err_rq_meta;
		}

		for (i = 0; i < qp->rq.mask; ++i)
			qp->rq_meta[i].next = &qp->rq_meta[i + 1];
		qp->rq_meta[i].next = IONIC_META_LAST;
		qp->rq_meta_head = &qp->rq_meta[0];

		pthread_spin_init(&qp->rq_lock, PTHREAD_PROCESS_PRIVATE);
	}

	return 0;

err_rq_meta:
	ionic_queue_destroy(&qp->rq);
err_rq:
	if (qp->has_sq) {
		pthread_spin_destroy(&qp->sq_lock);
		free(qp->sq_msn_idx);
err_sq_msn:
		free(qp->sq_meta);
err_sq_meta:
		ionic_queue_destroy(&qp->sq);
	}
err_sq:
	return rc;
}

static void ionic_free_queues(struct ionic_qp *qp)
{
	if (qp->has_rq) {
		pthread_spin_destroy(&qp->rq_lock);
		free(qp->rq_meta);
		ionic_queue_destroy(&qp->rq);
	}

	if (qp->has_sq) {
		pthread_spin_destroy(&qp->sq_lock);
		free(qp->sq_msn_idx);
		free(qp->sq_meta);
		ionic_queue_destroy(&qp->sq);
	}
}

static struct ibv_qp *ionic_create_qp_ex(struct ibv_context *ibctx,
					 struct ibv_qp_init_attr_ex *ex)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibctx);
	struct ionic_qp *qp;
	struct ionic_cq *cq;
	struct uionic_qp req;
	struct uionic_qp_resp resp;
	int rc;

	rc = ionic_check_qp_limits(ctx, &ex->cap);
	if (rc)
		goto err_qp;

	qp = calloc(1, sizeof(*qp));
	if (!qp) {
		rc = ENOMEM;
		goto err_qp;
	}

	qp->vqp.qp.qp_type = ex->qp_type;
	qp->vqp.qp.srq = ex->srq;

	qp->has_sq = ex->qp_type != IBV_QPT_XRC_RECV;

	qp->has_rq = !ex->srq &&
		ex->qp_type != IBV_QPT_XRC_SEND &&
		ex->qp_type != IBV_QPT_XRC_RECV;

	qp->is_srq = false;

	qp->sig_all = ex->sq_sig_all;

	rc = ionic_alloc_queues(ctx, qp, &ex->cap);
	if (rc)
		goto err_queues;

	req.sq.addr = (uintptr_t)qp->sq.ptr;
	req.sq.size = qp->sq.size;
	req.sq.mask = qp->sq.mask;
	req.sq.depth_log2 = qp->sq.depth_log2;
	req.sq.stride_log2 = qp->sq.stride_log2;

	req.rq.addr = (uintptr_t)qp->rq.ptr;
	req.rq.size = qp->rq.size;
	req.rq.mask = qp->rq.mask;
	req.rq.depth_log2 = qp->rq.depth_log2;
	req.rq.stride_log2 = qp->rq.stride_log2;

	req.compat = ctx->compat;

	rc = ibv_cmd_create_qp_ex2(ibctx, &qp->vqp, sizeof(qp->vqp), ex,
				   &req.ibv_cmd, sizeof(req.ibv_cmd), sizeof(req),
				   &resp.ibv_resp, sizeof(resp.ibv_resp), sizeof(resp));
	if (rc)
		goto err_cmd;

	if (resp.sq_hbm_offset) {
		qp->sq_hbm_ptr = ionic_map_device(qp->sq.size,
						  ctx->vctx.context.cmd_fd,
						  resp.sq_hbm_offset);
		if (!qp->sq_hbm_ptr) {
			rc = errno;
			goto err_hbm;
		}
	}

	qp->qpid = resp.qpid;
	ionic_queue_dbell_init(&qp->sq, qp->qpid);
	ionic_queue_dbell_init(&qp->rq, qp->qpid);

	pthread_mutex_lock(&ctx->mut);
	tbl_alloc_node(&ctx->qp_tbl);
	tbl_insert(&ctx->qp_tbl, qp, qp->qpid);
	pthread_mutex_unlock(&ctx->mut);

	if (qp->has_sq) {
		cq = to_ionic_cq(qp->vqp.qp.send_cq);
		pthread_spin_lock(&cq->lock);
		pthread_spin_unlock(&cq->lock);
	}

	if (qp->has_rq) {
		cq = to_ionic_cq(qp->vqp.qp.recv_cq);
		pthread_spin_lock(&cq->lock);
		pthread_spin_unlock(&cq->lock);
	}

	ex->cap.max_send_wr = qp->sq.mask;
	ex->cap.max_recv_wr = qp->rq.mask;
	ex->cap.max_send_sge =
		ionic_v1_send_wqe_max_sge(1u << qp->sq.stride_log2);
	ex->cap.max_recv_sge =
		ionic_v1_recv_wqe_max_sge(1u << qp->rq.stride_log2);
	ex->cap.max_inline_data =
		ionic_v1_send_wqe_max_data(1u << qp->sq.stride_log2);

	return &qp->vqp.qp;

err_hbm:
	ibv_cmd_destroy_qp(&qp->vqp.qp);
err_cmd:
	ionic_free_queues(qp);
err_queues:
	free(qp);
err_qp:
	errno = rc;
	return NULL;
}

static void ionic_reset_qp(struct ionic_qp *qp)
{
	struct ionic_cq *cq;

	if (qp->vqp.qp.send_cq) {
		cq = to_ionic_cq(qp->vqp.qp.send_cq);
		pthread_spin_lock(&cq->lock);
		ionic_clean_cq(cq, qp->qpid);
		pthread_spin_unlock(&cq->lock);
	}

	if (qp->vqp.qp.recv_cq) {
		cq = to_ionic_cq(qp->vqp.qp.recv_cq);
		pthread_spin_lock(&cq->lock);
		ionic_clean_cq(cq, qp->qpid);
		pthread_spin_unlock(&cq->lock);
	}

	if (qp->has_sq) {
		pthread_spin_lock(&qp->sq_lock);
		qp->sq_flush = false;
		qp->sq_msn_prod = 0;
		qp->sq_msn_cons = 0;
		qp->sq_npg_cons = 0;
		qp->sq.prod = 0;
		qp->sq.cons = 0;
		pthread_spin_unlock(&qp->sq_lock);
	}

	if (qp->has_rq) {
		pthread_spin_lock(&qp->rq_lock);
		qp->rq_flush = false;
		qp->rq.prod = 0;
		qp->rq.cons = 0;
		pthread_spin_unlock(&qp->rq_lock);
	}
}

static int ionic_modify_qp(struct ibv_qp *ibqp,
			   struct ibv_qp_attr *attr,
			   int attr_mask)
{
	struct ibv_modify_qp_ex cmd = {};
	struct ib_uverbs_ex_modify_qp_resp resp = {};
	struct ionic_qp *qp = to_ionic_qp(ibqp);
	int rc;

	if (!attr_mask)
		return 0;

	rc = ibv_cmd_modify_qp_ex(ibqp, attr, attr_mask,
				  &cmd, sizeof(cmd), sizeof(cmd),
				  &resp, sizeof(resp), sizeof(resp));
	if (rc)
		goto err_cmd;

	if (attr_mask & IBV_QP_STATE && attr->qp_state == IBV_QPS_RESET)
		ionic_reset_qp(qp);

err_cmd:
	return rc;
}

static int ionic_query_qp(struct ibv_qp *ibqp,
			  struct ibv_qp_attr *attr,
			  int attr_mask,
			  struct ibv_qp_init_attr *init_attr)
{
	struct ibv_query_qp cmd;
	int rc;

	rc = ibv_cmd_query_qp(ibqp, attr, attr_mask, init_attr,
			      &cmd, sizeof(cmd));

	init_attr->cap.max_inline_data = IONIC_MAX_INLINE_SIZE;

	attr->cap = init_attr->cap;

	return rc;
}

static int ionic_destroy_qp(struct ibv_qp *ibqp)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibqp->context);
	struct ionic_qp *qp = to_ionic_qp(ibqp);
	struct ionic_cq *cq;
	int rc;

	rc = ibv_cmd_destroy_qp(ibqp);
	if (rc)
		return rc;

	pthread_mutex_lock(&ctx->mut);
	tbl_free_node(&ctx->qp_tbl);
	tbl_delete(&ctx->qp_tbl, qp->qpid);
	pthread_mutex_unlock(&ctx->mut);

	if (qp->vqp.qp.send_cq) {
		cq = to_ionic_cq(qp->vqp.qp.send_cq);
		pthread_spin_lock(&cq->lock);
		ionic_clean_cq(cq, qp->qpid);
		list_del(&qp->cq_poll_sq);
		list_del(&qp->cq_flush_sq);
		pthread_spin_unlock(&cq->lock);
	}

	if (qp->vqp.qp.recv_cq) {
		cq = to_ionic_cq(qp->vqp.qp.recv_cq);
		pthread_spin_lock(&cq->lock);
		ionic_clean_cq(cq, qp->qpid);
		list_del(&qp->cq_flush_rq);
		pthread_spin_unlock(&cq->lock);
	}

	ionic_unmap(qp->sq_hbm_ptr, qp->sq.size);

	pthread_spin_destroy(&qp->rq_lock);
	pthread_spin_destroy(&qp->sq_lock);
	ionic_queue_destroy(&qp->rq);
	ionic_queue_destroy(&qp->sq);
	free(qp);

	return 0;
}

static int64_t ionic_prep_inline(void *data, uint32_t max_data,
				 struct ibv_sge *ibv_sgl, int num_sge)
{
	static const int64_t bit_31 = 1u << 31;
	int64_t len = 0, sg_len;
	int sg_i;

	for (sg_i = 0; sg_i < num_sge; ++sg_i) {
		sg_len = ibv_sgl[sg_i].length;

		/* sge length zero means 2GB */
		if (unlikely(sg_len == 0))
			sg_len = bit_31;

		/* greater than max inline data is invalid */
		if (unlikely(len + sg_len > max_data))
			return -EINVAL;

		memcpy(data + len, (void *)ibv_sgl[sg_i].addr, sg_len);

		len += sg_len;
	}

	return len;
}

static int64_t ionic_prep_sgl(struct ionic_sge *sgl, uint32_t max_sge,
			      struct ibv_sge *ibv_sgl, int num_sge)
{
	static const int64_t bit_31 = 1l << 31;
	int64_t len = 0, sg_len;
	int sg_i;

	if (unlikely(num_sge < 0 || (uint32_t)num_sge > max_sge))
		return -EINVAL;

	for (sg_i = 0; sg_i < num_sge; ++sg_i) {
		sg_len = ibv_sgl[sg_i].length;

		/* sge length zero means 2GB */
		if (unlikely(sg_len == 0))
			sg_len = bit_31;

		/* greater than 2GB data is invalid */
		if (unlikely(len + sg_len > bit_31))
			return -EINVAL;

		sgl[sg_i].va = htobe64(ibv_sgl[sg_i].addr);
		sgl[sg_i].len = htobe32(sg_len);
		sgl[sg_i].lkey = htobe32(ibv_sgl[sg_i].lkey);

		len += sg_len;
	}

	return len;
}

static void ionic_v0_prep_base(struct ionic_qp *qp,
			       struct ibv_send_wr *wr,
			       struct ionic_sq_meta *meta,
			       struct sqwqe_t *wqe)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);

	meta->wrid = wr->wr_id;
	meta->ibsts = IBV_WC_SUCCESS;
	meta->signal = false;

	wqe->base.wrid = qp->sq.prod;

	if (wr->send_flags & IBV_SEND_FENCE)
		wqe->base.fence = 1;

	if (wr->send_flags & IBV_SEND_SOLICITED)
		wqe->base.solicited_event = 1;

	if (qp->sig_all || wr->send_flags & IBV_SEND_SIGNALED) {
		wqe->base.complete_notify = 1;
		meta->signal = true;
	}

	meta->seq = qp->sq_msn_prod;
	meta->remote = qp->vqp.qp.qp_type != IBV_QPT_UD &&
		!ionic_ibop_is_local(wr->opcode);

	if (meta->remote) {
		qp->sq_msn_idx[meta->seq] = qp->sq.prod;
		qp->sq_msn_prod = ionic_queue_next(&qp->sq, qp->sq_msn_prod);
	}

	ionic_dbg(ctx, "post send %u prod %u", qp->qpid, qp->sq.prod);
	ionic_dbg_xdump(ctx, "wqe", wqe, 1u << qp->sq.stride_log2);

	ionic_queue_produce(&qp->sq);
}

static void ionic_v1_prep_base(struct ionic_qp *qp,
			       struct ibv_send_wr *wr,
			       struct ionic_sq_meta *meta,
			       struct ionic_v1_send_wqe *wqe)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);

	meta->wrid = wr->wr_id;
	meta->ibsts = IBV_WC_SUCCESS;
	meta->signal = false;

	wqe->base.wqe_id = qp->sq.prod;

	if (wr->send_flags & IBV_SEND_FENCE)
		wqe->base.flags |= htobe32(IONIC_V1_FLAG_FENCE);

	if (wr->send_flags & IBV_SEND_SOLICITED)
		wqe->base.flags |= htobe32(IONIC_V1_FLAG_SOL);

	if (qp->sig_all || wr->send_flags & IBV_SEND_SIGNALED) {
		wqe->base.flags |= htobe32(IONIC_V1_FLAG_SIG);
		meta->signal = true;
	}

	meta->seq = qp->sq_msn_prod;
	meta->remote = qp->vqp.qp.qp_type != IBV_QPT_UD &&
		!ionic_ibop_is_local(wr->opcode);

	if (meta->remote) {
		qp->sq_msn_idx[meta->seq] = qp->sq.prod;
		qp->sq_msn_prod = ionic_queue_next(&qp->sq, qp->sq_msn_prod);
	}

	ionic_dbg(ctx, "post send %u prod %u", qp->qpid, qp->sq.prod);
	ionic_dbg_xdump(ctx, "wqe", wqe, 1u << qp->sq.stride_log2);

	ionic_queue_produce(&qp->sq);
}

static int ionic_v0_prep_common(struct ionic_qp *qp,
				struct ibv_send_wr *wr,
				struct ionic_sq_meta *meta,
				struct sqwqe_t *wqe,
				/* XXX length field offset differs per opcode */
				__u32 *wqe_length_field)
{
	int64_t signed_len;
	uint32_t mval;

	if (wr->send_flags & IBV_SEND_INLINE) {
		wqe->base.num_sges = 0;
		wqe->base.inline_data_vld = 1;
		mval = ionic_v1_send_wqe_max_data(1u << qp->sq.stride_log2);
		signed_len = ionic_prep_inline(wqe->u.non_atomic.sg_arr, mval,
					       wr->sg_list, wr->num_sge);
	} else {
		wqe->base.num_sges = wr->num_sge;
		mval = ionic_v1_send_wqe_max_sge(1u << qp->sq.stride_log2);
		signed_len = ionic_prep_sgl(wqe->u.non_atomic.sg_arr, mval,
					    wr->sg_list, wr->num_sge);
	}

	if (unlikely(signed_len < 0))
		return (int)-signed_len;

	meta->len = (uint32_t)signed_len;
	*wqe_length_field = htobe32((uint32_t)signed_len);

	ionic_v0_prep_base(qp, wr, meta, wqe);

	return 0;
}

static int ionic_v1_prep_common(struct ionic_qp *qp,
				struct ibv_send_wr *wr,
				struct ionic_sq_meta *meta,
				struct ionic_v1_send_wqe *wqe)
{
	int64_t signed_len;
	uint32_t mval;

	if (wr->send_flags & IBV_SEND_INLINE) {
		wqe->base.num_sge_key = 0;
		wqe->base.flags |= htobe32(IONIC_V1_FLAG_INL);
		mval = ionic_v1_send_wqe_max_data(1u << qp->sq.stride_log2);
		signed_len = ionic_prep_inline(wqe->common.data, mval,
					       wr->sg_list, wr->num_sge);
	} else {
		wqe->base.num_sge_key = wr->num_sge;
		mval = ionic_v1_send_wqe_max_sge(1u << qp->sq.stride_log2);
		signed_len = ionic_prep_sgl(wqe->common.sgl, mval,
					    wr->sg_list, wr->num_sge);
	}

	if (unlikely(signed_len < 0))
		return (int)-signed_len;

	meta->len = signed_len;
	wqe->base.length_key = htobe32(signed_len);

	ionic_v1_prep_base(qp, wr, meta, wqe);

	return 0;
}

static int ionic_v0_prep_send(struct ionic_qp *qp,
			      struct ibv_send_wr *wr)
{
	struct ionic_sq_meta *meta;
	struct sqwqe_t *wqe;

	meta = &qp->sq_meta[qp->sq.prod];
	wqe = ionic_queue_at_prod(&qp->sq);

	memset(wqe, 0, 1u << qp->sq.stride_log2);

	meta->ibop = IBV_WC_SEND;

	switch (wr->opcode) {
	case IBV_WR_SEND:
		wqe->base.op_type = IONIC_WR_OPCD_SEND;
		break;
	case IBV_WR_SEND_WITH_IMM:
		wqe->base.op_type = IONIC_WR_OPCD_SEND_IMM;
		wqe->u.non_atomic.wqe.send.imm_data = wr->imm_data;
		break;
	case IBV_WR_SEND_WITH_INV:
		wqe->base.op_type = IONIC_WR_OPCD_SEND_INVAL;
		wqe->u.non_atomic.wqe.send.imm_data =
			htobe32(wr->invalidate_rkey);

		/* XXX just use imm_data field */
		wqe->u.non_atomic.wqe.send.inv_key =
			htobe32(wr->invalidate_rkey);
		break;
	default:
		return EINVAL;
	}

	/* XXX warning: taking address of packed member */
	return ionic_v0_prep_common(qp, wr, meta, wqe,
				 &wqe->u.non_atomic.wqe.send.length);
}

static int ionic_v1_prep_send(struct ionic_qp *qp,
			      struct ibv_send_wr *wr)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);
	struct ionic_sq_meta *meta;
	struct ionic_v1_send_wqe *wqe;

	meta = &qp->sq_meta[qp->sq.prod];
	wqe = ionic_queue_at_prod(&qp->sq);

	memset(wqe, 0, 1u << qp->sq.stride_log2);

	meta->ibop = IBV_WC_SEND;

	switch (wr->opcode) {
	case IBV_WR_SEND:
		wqe->base.op = IONIC_V1_OP_SEND;
		break;
	case IBV_WR_SEND_WITH_IMM:
		wqe->base.op = IONIC_V1_OP_SEND_IMM;
		wqe->common.send.imm_data_rkey = wr->imm_data;
		break;
	case IBV_WR_SEND_WITH_INV:
		wqe->base.op = IONIC_V1_OP_SEND_INV;
		wqe->common.send.imm_data_rkey = htobe32(wr->invalidate_rkey);
		break;
	default:
		return EINVAL;
	}

	/* XXX makeshift will be removed */
	if (ctx->version != 1 || ctx->opcodes <= wqe->base.op)
		return ionic_v0_prep_send(qp, wr);

	return ionic_v1_prep_common(qp, wr, meta, wqe);
}

static int ionic_v0_prep_send_ud(struct ionic_qp *qp, struct ibv_send_wr *wr)
{
	struct ionic_sq_meta *meta;
	struct sqwqe_t *wqe;
	struct ionic_ah *ah;

	if (unlikely(!wr->wr.ud.ah))
		return EINVAL;

	ah = to_ionic_ah(wr->wr.ud.ah);

	meta = &qp->sq_meta[qp->sq.prod];
	wqe = ionic_queue_at_prod(&qp->sq);

	memset(wqe, 0, 1u << qp->sq.stride_log2);

	/* XXX endian? */
	wqe->u.non_atomic.wqe.ud_send.q_key = htobe32(wr->wr.ud.remote_qkey);
	wqe->u.non_atomic.wqe.ud_send.ah_size = 0;
	wqe->u.non_atomic.wqe.ud_send.dst_qp = htobe32(wr->wr.ud.remote_qpn) >> 8; /* XXX not portable, get rid of bit fields in wqe! */
	wqe->u.non_atomic.wqe.ud_send.ah_handle = htobe32(ah->ahid);

	meta->ibop = IBV_WC_SEND;

	switch (wr->opcode) {
	case IBV_WR_SEND:
		wqe->base.op_type = IONIC_WR_OPCD_SEND;
		break;
	case IBV_WR_SEND_WITH_IMM:
		wqe->base.op_type = IONIC_WR_OPCD_SEND_IMM;
		wqe->u.non_atomic.wqe.ud_send.imm_data = wr->imm_data;
		break;
	default:
		return EINVAL;
	}

	/* XXX warning: taking address of packed member */
	return ionic_v0_prep_common(qp, wr, meta, wqe,
				    &wqe->u.non_atomic.wqe.ud_send.length);
}

static int ionic_v1_prep_send_ud(struct ionic_qp *qp, struct ibv_send_wr *wr)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);
	struct ionic_sq_meta *meta;
	struct ionic_v1_send_wqe *wqe;
	struct ionic_ah *ah;

	if (unlikely(!wr->wr.ud.ah))
		return EINVAL;

	ah = to_ionic_ah(wr->wr.ud.ah);

	meta = &qp->sq_meta[qp->sq.prod];
	wqe = ionic_queue_at_prod(&qp->sq);

	memset(wqe, 0, 1u << qp->sq.stride_log2);

	wqe->common.send.ah_id = htobe32(ah->ahid);
	wqe->common.send.dest_qpn = htobe32(wr->wr.ud.remote_qpn);
	wqe->common.send.dest_qkey = htobe32(wr->wr.ud.remote_qkey);

	meta->ibop = IBV_WC_SEND;

	switch (wr->opcode) {
	case IBV_WR_SEND:
		wqe->base.op = IONIC_V1_OP_SEND;
		break;
	case IBV_WR_SEND_WITH_IMM:
		wqe->base.op = IONIC_V1_OP_SEND_IMM;
		wqe->common.send.imm_data_rkey = wr->imm_data;
		break;
	default:
		return EINVAL;
	}

	/* XXX makeshift will be removed */
	if (ctx->version != 1 || ctx->opcodes <= wqe->base.op)
		return ionic_v0_prep_send_ud(qp, wr);

	return ionic_v1_prep_common(qp, wr, meta, wqe);
}

static int ionic_v0_prep_rdma(struct ionic_qp *qp,
			      struct ibv_send_wr *wr)
{
	struct ionic_sq_meta *meta;
	struct sqwqe_t *wqe;

	meta = &qp->sq_meta[qp->sq.prod];
	wqe = ionic_queue_at_prod(&qp->sq);

	memset(wqe, 0, 1u << qp->sq.stride_log2);

	meta->ibop = IBV_WC_RDMA_WRITE;

	switch (wr->opcode) {
	case IBV_WR_RDMA_READ:
		if (wr->send_flags & (IBV_SEND_SOLICITED | IBV_SEND_INLINE))
			return EINVAL;
		meta->ibop = IBV_WC_RDMA_READ;
		wqe->base.op_type = IONIC_WR_OPCD_RDMA_READ;
		break;
	case IBV_WR_RDMA_WRITE:
		if (wr->send_flags & IBV_SEND_SOLICITED)
			return EINVAL;
		wqe->base.op_type = IONIC_WR_OPCD_RDMA_WRITE;
		break;
	case IBV_WR_RDMA_WRITE_WITH_IMM:
		wqe->base.op_type = IONIC_WR_OPCD_RDMA_WRITE_IMM;
		wqe->u.non_atomic.wqe.ud_send.imm_data = wr->imm_data;
		break;
	default:
		return EINVAL;
	}

	wqe->u.non_atomic.wqe.rdma.va = htobe64(wr->wr.rdma.remote_addr);
	wqe->u.non_atomic.wqe.rdma.r_key = htobe32(wr->wr.rdma.rkey);

	/* XXX warning: taking address of packed member */
	return ionic_v0_prep_common(qp, wr, meta, wqe,
				    &wqe->u.non_atomic.wqe.rdma.length);
}

static int ionic_v1_prep_rdma(struct ionic_qp *qp,
			      struct ibv_send_wr *wr)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);
	struct ionic_sq_meta *meta;
	struct ionic_v1_send_wqe *wqe;

	meta = &qp->sq_meta[qp->sq.prod];
	wqe = ionic_queue_at_prod(&qp->sq);

	memset(wqe, 0, 1u << qp->sq.stride_log2);

	meta->ibop = IBV_WC_RDMA_WRITE;

	switch (wr->opcode) {
	case IBV_WR_RDMA_READ:
		if (wr->send_flags & (IBV_SEND_SOLICITED | IBV_SEND_INLINE))
			return EINVAL;
		meta->ibop = IBV_WC_RDMA_READ;
		wqe->base.op = IONIC_V1_OP_RDMA_READ;
		break;
	case IBV_WR_RDMA_WRITE:
		if (wr->send_flags & IBV_SEND_SOLICITED)
			return EINVAL;
		wqe->base.op = IONIC_V1_OP_RDMA_WRITE;
		break;
	case IBV_WR_RDMA_WRITE_WITH_IMM:
		wqe->base.op = IONIC_V1_OP_RDMA_WRITE_IMM;
		wqe->common.rdma.imm_data = wr->imm_data;
		break;
	default:
		return EINVAL;
	}

	wqe->common.rdma.remote_va = htobe64(wr->wr.rdma.remote_addr);
	wqe->common.rdma.remote_rkey = htobe32(wr->wr.rdma.rkey);

	/* XXX makeshift will be removed */
	if (ctx->version != 1 || ctx->opcodes <= wqe->base.op)
		return ionic_v0_prep_rdma(qp, wr);

	return ionic_v1_prep_common(qp, wr, meta, wqe);
}

static int ionic_v0_prep_atomic(struct ionic_qp *qp,
				struct ibv_send_wr *wr)
{
	struct ionic_sq_meta *meta;
	struct sqwqe_t *wqe;

	if (wr->num_sge != 1 || wr->sg_list[0].length != 8)
		return EINVAL;

	if (wr->send_flags & (IBV_SEND_SOLICITED | IBV_SEND_INLINE))
		return EINVAL;

	meta = &qp->sq_meta[qp->sq.prod];
	wqe = ionic_queue_at_prod(&qp->sq);

	memset(wqe, 0, 1u << qp->sq.stride_log2);

	switch (wr->opcode) {
	case IBV_WR_ATOMIC_CMP_AND_SWP:
		meta->ibop = IBV_WC_COMP_SWAP;
		wqe->base.op_type = IONIC_WR_OPCD_ATOMIC_CS;
		wqe->u.atomic.swap_or_add_data =
			htobe64(wr->wr.atomic.swap);
		wqe->u.atomic.cmp_data =
			htobe64(wr->wr.atomic.compare_add);
		break;
	case IBV_WR_ATOMIC_FETCH_AND_ADD:
		meta->ibop = IBV_WC_FETCH_ADD;
		wqe->base.op_type = IONIC_WR_OPCD_ATOMIC_FA;
		wqe->u.atomic.swap_or_add_data =
			htobe64(wr->wr.atomic.compare_add);
		break;
	default:
		return EINVAL;
	}

	wqe->u.atomic.r_key = htobe32(wr->wr.atomic.rkey);
	wqe->u.atomic.va = htobe64(wr->wr.atomic.remote_addr);

	wqe->base.num_sges = 1;
	wqe->u.atomic.sge.va = htobe64(wr->sg_list[0].addr);
	wqe->u.atomic.sge.lkey = htobe32(wr->sg_list[0].lkey);
	wqe->u.atomic.sge.len = htobe32(8);

	ionic_v0_prep_base(qp, wr, meta, wqe);

	return 0;
}

static int ionic_v1_prep_atomic(struct ionic_qp *qp,
				struct ibv_send_wr *wr)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);
	struct ionic_sq_meta *meta;
	struct ionic_v1_send_wqe *wqe;

	if (wr->num_sge != 1 || wr->sg_list[0].length != 8)
		return EINVAL;

	if (wr->send_flags & (IBV_SEND_SOLICITED | IBV_SEND_INLINE))
		return EINVAL;

	meta = &qp->sq_meta[qp->sq.prod];
	wqe = ionic_queue_at_prod(&qp->sq);

	memset(wqe, 0, 1u << qp->sq.stride_log2);

	switch (wr->opcode) {
	case IBV_WR_ATOMIC_CMP_AND_SWP:
		meta->ibop = IBV_WC_COMP_SWAP;
		wqe->base.op = IONIC_V1_OP_ATOMIC_CS;
		wqe->atomic.swap_add_high = htobe32(wr->wr.atomic.swap >> 32);
		wqe->atomic.swap_add_low = htobe32(wr->wr.atomic.swap);
		wqe->atomic.compare_high = htobe32(wr->wr.atomic.compare_add >> 32);
		wqe->atomic.compare_low = htobe32(wr->wr.atomic.compare_add);
		break;
	case IBV_WR_ATOMIC_FETCH_AND_ADD:
		meta->ibop = IBV_WC_FETCH_ADD;
		wqe->base.op = IONIC_V1_OP_ATOMIC_FA;
		wqe->atomic.swap_add_high = htobe32(wr->wr.atomic.compare_add >> 32);
		wqe->atomic.swap_add_low = htobe32(wr->wr.atomic.compare_add);
		break;
	default:
		return EINVAL;
	}

	wqe->atomic.remote_va = htobe64(wr->wr.atomic.remote_addr);
	wqe->atomic.remote_rkey = htobe32(wr->wr.atomic.rkey);

	wqe->base.num_sge_key = 1;
	wqe->atomic.sge.va = htobe64(wr->sg_list[0].addr);
	wqe->atomic.sge.len = htobe32(8);
	wqe->atomic.sge.lkey = htobe32(wr->sg_list[0].lkey);

	/* XXX makeshift will be removed */
	if (ctx->version != 1 || ctx->opcodes <= wqe->base.op)
		return ionic_v0_prep_atomic(qp, wr);

	ionic_v1_prep_base(qp, wr, meta, wqe);

	return 0;
}

static int ionic_prep_one_rc(struct ionic_qp *qp,
			     struct ibv_send_wr *wr)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);
	int rc = 0;

	switch (wr->opcode) {
	case IBV_WR_SEND:
	case IBV_WR_SEND_WITH_IMM:
	case IBV_WR_SEND_WITH_INV:
		rc = ionic_v1_prep_send(qp, wr);
		break;
	case IBV_WR_RDMA_READ:
	case IBV_WR_RDMA_WRITE:
	case IBV_WR_RDMA_WRITE_WITH_IMM:
		rc = ionic_v1_prep_rdma(qp, wr);
		break;
	case IBV_WR_ATOMIC_CMP_AND_SWP:
	case IBV_WR_ATOMIC_FETCH_AND_ADD:
		rc = ionic_v1_prep_atomic(qp, wr);
		break;
	default:
		ionic_dbg(ctx, "invalid opcode %d", wr->opcode);
		rc = EINVAL;
	}

	return rc;
}

static int ionic_prep_one_ud(struct ionic_qp *qp,
			     struct ibv_send_wr *wr)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);
	int rc = 0;

	switch (wr->opcode) {
	case IBV_WR_SEND:
	case IBV_WR_SEND_WITH_IMM:
		rc = ionic_v1_prep_send_ud(qp, wr);
		break;
	default:
		ionic_dbg(ctx, "invalid opcode %d", wr->opcode);
		rc = EINVAL;
	}

	return rc;
}

static void ionic_post_hbm(struct ionic_ctx *ctx, struct ionic_qp *qp)
{
	void *hbm_ptr;
	void *wqe_ptr;
	uint32_t stride;
	uint16_t pos, end;
	uint8_t stride_log2;

	stride_log2 = qp->sq.stride_log2;
	stride = 1u << stride_log2;

	pos = qp->sq_hbm_prod;
	end = qp->sq.prod;

	while (pos != end) {
		hbm_ptr = qp->sq_hbm_ptr + ((size_t)pos << stride_log2);
		wqe_ptr = ionic_queue_at(&qp->sq, pos);

		mmio_wc_start();
		mmio_memcpy_x64(hbm_ptr, wqe_ptr, stride);
		mmio_flush_writes();

		pos = ionic_queue_next(&qp->sq, pos);

		ionic_dbg(ctx, "dbell qtype %d val %#lx",
			  ctx->sq_qtype, qp->sq.dbell | pos);
		ionic_dbell_ring(&ctx->dbpage[ctx->sq_qtype],
				 qp->sq.dbell | pos);
	}

	qp->sq_hbm_prod = end;
}

static int ionic_post_send_common(struct ionic_ctx *ctx,
				  struct ionic_cq *cq,
				  struct ionic_qp *qp,
				  struct ibv_send_wr *wr,
				  struct ibv_send_wr **bad)
{
	uint16_t old_prod;
	bool flush;
	int rc = 0;

	if (unlikely(!bad))
		return EINVAL;

	if (unlikely(!qp->has_sq)) {
		*bad = wr;
		return EINVAL;
	}

	if (unlikely(qp->vqp.qp.state < IBV_QPS_RTS)) {
		*bad = wr;
		return EINVAL;
	}

	pthread_spin_lock(&qp->sq_lock);

	old_prod = qp->sq.prod;

	if (qp->vqp.qp.qp_type == IBV_QPT_UD) {
		while (wr) {
			if (ionic_queue_full(&qp->sq)) {
				ionic_dbg(ctx, "queue full");
				rc = ENOMEM;
				goto out;
			}

			rc = ionic_prep_one_ud(qp, wr);
			if (rc)
				goto out;

			wr = wr->next;
		}
	} else {
		while (wr) {
			if (ionic_queue_full(&qp->sq)) {
				ionic_dbg(ctx, "queue full");
				rc = ENOMEM;
				goto out;
			}

			rc = ionic_prep_one_rc(qp, wr);
			if (rc)
				goto out;

			wr = wr->next;
		}
	}

out:
	if (likely(qp->sq.prod != old_prod)) {
		if (qp->sq_hbm_ptr) {
			ionic_post_hbm(ctx, qp);
		} else {
			udma_to_device_barrier();
			ionic_dbg(ctx, "dbell qtype %d val %#lx",
				  ctx->sq_qtype,
				  ionic_queue_dbell_val(&qp->sq));
			ionic_dbell_ring(&ctx->dbpage[ctx->sq_qtype],
					 ionic_queue_dbell_val(&qp->sq));
		}
	}

	flush = qp->sq_flush;

	pthread_spin_unlock(&qp->sq_lock);

	if (flush) {
		pthread_spin_lock(&cq->lock);
		list_del(&qp->cq_flush_sq);
		list_add_tail(&cq->flush_sq, &qp->cq_flush_sq);
		pthread_spin_unlock(&cq->lock);
	}

	*bad = wr;
	return rc;
}

static int ionic_v1_prep_recv(struct ionic_qp *qp,
			      struct ibv_recv_wr *wr)
{
	struct ionic_ctx *ctx = to_ionic_ctx(qp->vqp.qp.context);
	struct ionic_rq_meta *meta;
	struct ionic_v1_recv_wqe *wqe;
	int64_t signed_len;
	uint32_t mval;

	wqe = ionic_queue_at_prod(&qp->rq);

	/* if wqe is owned by device, caller can try posting again soon */
	if (wqe->base.flags & IONIC_V1_FLAG_FENCE)
		return -EAGAIN;

	meta = qp->rq_meta_head;
	if (unlikely(meta == IONIC_META_LAST) ||
	    unlikely(meta == IONIC_META_POSTED))
		return -EIO;

	memset(wqe, 0, 1u << qp->rq.stride_log2);

	mval = ionic_v1_recv_wqe_max_sge(1u << qp->rq.stride_log2);
	signed_len = ionic_prep_sgl(wqe->sgl, mval,
				    wr->sg_list, wr->num_sge);
	if (signed_len < 0)
		return (int)-signed_len;

	meta->wrid = wr->wr_id;

	/* XXX bytes recvd should come from cqe */
	meta->len = signed_len;

	wqe->base.wqe_id = meta - qp->rq_meta;
	wqe->base.op = wr->num_sge; /* XXX makeshift has num_sge in the opcode position */
	wqe->base.num_sge_key = wr->num_sge;
	wqe->base.length_key = htobe32(signed_len);

	/* if this is a srq, set fence bit to indicate device ownership */
	if (qp->is_srq)
		wqe->base.flags |= htobe16(IONIC_V1_FLAG_FENCE);

	ionic_dbg(ctx, "post recv %u prod %u", qp->qpid, qp->rq.prod);
	ionic_dbg_xdump(ctx, "wqe", wqe, 1u << qp->rq.stride_log2);

	ionic_queue_produce(&qp->rq);

	qp->rq_meta_head = meta->next;
	meta->next = IONIC_META_POSTED;

	return 0;
}

static int ionic_post_recv_common(struct ionic_ctx *ctx,
				  struct ionic_cq *cq,
				  struct ionic_qp *qp,
				  struct ibv_recv_wr *wr,
				  struct ibv_recv_wr **bad)
{
	uint16_t old_prod;
	bool flush;
	int rc = 0;

	if (unlikely(!bad))
		return EINVAL;

	if (unlikely(!qp->has_rq)) {
		*bad = wr;
		return EINVAL;
	}

	if (unlikely(qp->vqp.qp.state < IBV_QPS_INIT)) {
		*bad = wr;
		return EINVAL;
	}

	pthread_spin_lock(&qp->rq_lock);

	old_prod = qp->rq.prod;

	while (wr) {
		if (ionic_queue_full(&qp->rq)) {
			ionic_dbg(ctx, "queue full");
			rc = ENOMEM;
			goto out;
		}

		rc = ionic_v1_prep_recv(qp, wr);
		if (rc)
			goto out;

		wr = wr->next;
	}

out:
	if (likely(qp->rq.prod != old_prod)) {
		udma_to_device_barrier();
		ionic_dbg(ctx, "dbell qtype %d val %#lx",
			  ctx->rq_qtype, ionic_queue_dbell_val(&qp->rq));
		ionic_dbell_ring(&ctx->dbpage[ctx->rq_qtype],
				 ionic_queue_dbell_val(&qp->rq));
	}

	flush = qp->rq_flush;

	pthread_spin_unlock(&qp->rq_lock);

	if (flush) {
		pthread_spin_lock(&cq->lock);
		list_del(&qp->cq_flush_rq);
		list_add_tail(&cq->flush_rq, &qp->cq_flush_rq);
		pthread_spin_unlock(&cq->lock);
	}

	*bad = wr;
	return rc;
}

static int ionic_post_send(struct ibv_qp *ibqp,
			   struct ibv_send_wr *wr,
			   struct ibv_send_wr **bad)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibqp->context);
	struct ionic_qp *qp = to_ionic_qp(ibqp);
	struct ionic_cq *cq = to_ionic_cq(ibqp->send_cq);

	return ionic_post_send_common(ctx, cq, qp, wr, bad);
}

static int ionic_post_recv(struct ibv_qp *ibqp,
			   struct ibv_recv_wr *wr,
			   struct ibv_recv_wr **bad)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibqp->context);
	struct ionic_qp *qp = to_ionic_qp(ibqp);
	struct ionic_cq *cq = to_ionic_cq(ibqp->recv_cq);

	return ionic_post_recv_common(ctx, cq, qp, wr, bad);
}

static struct ibv_srq *ionic_create_srq_ex(struct ibv_context *ibctx,
					   struct ibv_srq_init_attr_ex *ex)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibctx);
	struct ionic_qp *qp;
	struct ionic_cq *cq;
	struct ibv_qp_cap cap = {
		.max_recv_wr = ex->attr.max_wr,
		.max_recv_sge = ex->attr.max_sge,
	};
	struct uionic_srq req;
	struct uionic_srq_resp resp;
	int rc;

	qp = calloc(1, sizeof(*qp));
	if (!qp) {
		rc = ENOMEM;
		goto err_qp;
	}

	qp->has_sq = false;
	qp->has_rq = true;
	qp->is_srq = true;

	qp->sig_all = false;

	rc = ionic_alloc_queues(ctx, qp, &cap);
	if (rc)
		goto err_queues;

	req.rq.addr = (uintptr_t)qp->rq.ptr;
	req.rq.size = qp->rq.size;
	req.rq.mask = qp->rq.mask;
	req.rq.depth_log2 = qp->rq.depth_log2;
	req.rq.stride_log2 = qp->rq.stride_log2;

	req.compat = ctx->compat;

	rc = ibv_cmd_create_srq_ex(ibctx, &qp->vsrq, sizeof(qp->vsrq), ex,
				   &req.ibv_cmd, sizeof(req),
				   &resp.ibv_resp, sizeof(resp));
	if (rc)
		goto err_cmd;

	qp->qpid = resp.qpid;

	if (ex->cq) {
		pthread_mutex_lock(&ctx->mut);
		tbl_alloc_node(&ctx->qp_tbl);
		tbl_insert(&ctx->qp_tbl, qp, qp->qpid);
		pthread_mutex_unlock(&ctx->mut);

		cq = to_ionic_cq(ex->cq);
		pthread_spin_lock(&cq->lock);
		list_del(&qp->cq_flush_rq);
		pthread_spin_unlock(&cq->lock);
	}

	ex->attr.max_wr = qp->rq.mask;
	ex->attr.max_sge = ionic_v1_recv_wqe_max_sge(1u << qp->rq.stride_log2);

	return &qp->vsrq.srq;

err_cmd:
	ionic_free_queues(qp);
err_queues:
	free(qp);
err_qp:
	errno = rc;
	return NULL;
}

static int ionic_get_srq_num(struct ibv_srq *ibsrq, uint32_t *srq_num)
{
	struct ionic_qp *qp = to_ionic_srq(ibsrq);

	*srq_num = qp->qpid;

	return 0;
}

static int ionic_modify_srq(struct ibv_srq *ibsrq, struct ibv_srq_attr *attr,
			    int init_attr)
{
	return -ENOSYS;
}

static int ionic_destroy_srq(struct ibv_srq *ibsrq)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibsrq->context);
	struct ionic_qp *qp = to_ionic_srq(ibsrq);
	struct ionic_cq *cq;
	int rc;

	rc = ibv_cmd_destroy_srq(ibsrq);
	if (rc)
		return rc;

	pthread_mutex_lock(&ctx->mut);
	tbl_free_node(&ctx->qp_tbl);
	tbl_delete(&ctx->qp_tbl, qp->qpid);
	pthread_mutex_unlock(&ctx->mut);

	if (qp->vsrq.cq) {
		cq = to_ionic_cq(qp->vsrq.cq);
		pthread_spin_lock(&cq->lock);
		ionic_clean_cq(cq, qp->qpid);
		list_del(&qp->cq_flush_rq);
		pthread_spin_unlock(&cq->lock);
	}

	pthread_spin_destroy(&qp->rq_lock);
	ionic_queue_destroy(&qp->rq);
	free(qp);

	return 0;
}

static int ionic_query_srq(struct ibv_srq *ibsrq, struct ibv_srq_attr *attr)
{
	return -ENOSYS;
}

static int ionic_post_srq_recv(struct ibv_srq *ibsrq, struct ibv_recv_wr *wr,
			       struct ibv_recv_wr **bad)
{
	struct ionic_ctx *ctx = to_ionic_ctx(ibsrq->context);
	struct ionic_qp *qp = to_ionic_srq(ibsrq);
	struct ionic_cq *cq = NULL;

	if (qp->vsrq.cq)
		cq = to_ionic_cq(qp->vsrq.cq);

	return ionic_post_recv_common(ctx, cq, qp, wr, bad);
}

static struct ibv_ah *ionic_create_ah(struct ibv_pd *ibpd,
				      struct ibv_ah_attr *attr)
{
	struct ionic_ah *ah;
	struct uionic_ah_resp resp;
	int rc;

	ah = calloc(1, sizeof(*ah));
	if (!ah) {
		rc = errno;
		goto err_ah;
	}

	rc = ibv_cmd_create_ah(ibpd, &ah->ibah, attr,
			       &resp.ibv_resp, sizeof(resp));
	if (rc)
		goto err_cmd;

	ah->ahid = resp.ahid;

	return &ah->ibah;

err_cmd:
	free(ah);
err_ah:
	errno = rc;
	return NULL;
}

static int ionic_destroy_ah(struct ibv_ah *ibah)
{
	struct ionic_ah *ah = to_ionic_ah(ibah);
	int rc;

	rc = ibv_cmd_destroy_ah(ibah);
	if (rc)
		return rc;

	free(ah);

	return 0;
}

const struct verbs_context_ops ionic_ctx_ops = {
	.create_cq		= ionic_create_cq,
	.poll_cq		= ionic_poll_cq,
	.req_notify_cq		= ionic_req_notify_cq,
	.resize_cq		= ionic_resize_cq,
	.destroy_cq		= ionic_destroy_cq,
	.create_srq_ex		= ionic_create_srq_ex,
	.get_srq_num		= ionic_get_srq_num,
	.modify_srq		= ionic_modify_srq,
	.query_srq		= ionic_query_srq,
	.destroy_srq		= ionic_destroy_srq,
	.post_srq_recv		= ionic_post_srq_recv,
	.create_qp_ex		= ionic_create_qp_ex,
	.query_qp		= ionic_query_qp,
	.modify_qp		= ionic_modify_qp,
	.destroy_qp		= ionic_destroy_qp,
	.post_send		= ionic_post_send,
	.post_recv		= ionic_post_recv,
	.create_ah		= ionic_create_ah,
	.destroy_ah		= ionic_destroy_ah
};
