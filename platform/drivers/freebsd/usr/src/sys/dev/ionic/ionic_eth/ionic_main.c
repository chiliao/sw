/*
 * Copyright (c) 2017-2019 Pensando Systems, Inc.  All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions
 * are met:
 * 1. Redistributions of source code must retain the above copyright
 *    notice, this list of conditions and the following disclaimer.
 * 2. Redistributions in binary form must reproduce the above copyright
 *    notice, this list of conditions and the following disclaimer in the
 *    documentation and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE AUTHOR AND CONTRIBUTORS ``AS IS'' AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED.  IN NO EVENT SHALL THE AUTHOR OR CONTRIBUTORS BE LIABLE
 * FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
 * DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS
 * OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION)
 * HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT
 * LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY
 * OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF
 * SUCH DAMAGE.
 */

#include <linux/module.h>
#include <linux/netdevice.h>
#include <linux/pci.h>

#include "ionic.h"
#include "ionic_bus.h"
#include "ionic_lif.h"

MODULE_DESCRIPTION(DRV_DESCRIPTION);
MODULE_AUTHOR("Anish Gupta <anish@pensando.io>");
MODULE_VERSION(ionic, 1);

const char *
ionic_port_oper_status_str(enum ionic_port_oper_status status)
{

	switch (status) {
	case IONIC_PORT_OPER_STATUS_NONE:
		return "PORT_OPER_STATUS_NONE";
	case IONIC_PORT_OPER_STATUS_UP:
		return "PORT_OPER_STATUS_UP";
	case IONIC_PORT_OPER_STATUS_DOWN:
		return "PORT_OPER_STATUS_DOWN";
	default:
		return "PORT_OPER_STATUS_UNKNOWN";
	}
}

const char *
ionic_port_admin_state_str(enum ionic_port_admin_state state)
{

	switch (state) {
	case IONIC_PORT_ADMIN_STATE_NONE:
		return "PORT_ADMIN_STATE_NONE";
	case IONIC_PORT_ADMIN_STATE_DOWN:
		return "PORT_ADMIN_STATE_DOWN";
	case IONIC_PORT_ADMIN_STATE_UP:
		return "PORT_ADMIN_STATE_UP";
	default:
		return "PORT_ADMIN_STATE_UNKNOWN";
	}
}

const char *
ionic_port_fec_type_str(enum ionic_port_fec_type type)
{

	switch (type) {
	case IONIC_PORT_FEC_TYPE_NONE:
		return "PORT_FEC_TYPE_NONE";
	case IONIC_PORT_FEC_TYPE_FC:
		return "PORT_FEC_TYPE_FC";
	case IONIC_PORT_FEC_TYPE_RS:
		return "PORT_FEC_TYPE_RS";
	default:
		return "PORT_FEC_TYPE_UNKNOWN";
	}
}

const char *
ionic_port_pause_type_str(enum ionic_port_pause_type type)
{

	switch (type) {
	case IONIC_PORT_PAUSE_TYPE_NONE:
		return "PORT_PAUSE_TYPE_NONE";
	case IONIC_PORT_PAUSE_TYPE_LINK:
		return "PORT_PAUSE_TYPE_LINK";
	case IONIC_PORT_PAUSE_TYPE_PFC:
		return "PORT_PAUSE_TYPE_PFC";
	default:
		return "PORT_PAUSE_TYPE_UNKNOWN";
	}
}

const char *
ionic_port_loopback_mode_str(enum ionic_port_loopback_mode mode)
{

	switch (mode) {
	case IONIC_PORT_LOOPBACK_MODE_NONE:
		return "PORT_LOOPBACK_MODE_NONE";
	case IONIC_PORT_LOOPBACK_MODE_MAC:
		return "PORT_LOOPBACK_MODE_MAC";
	case IONIC_PORT_LOOPBACK_MODE_PHY:
		return "PORT_LOOPBACK_MODE_PHY";
	default:
		return "PORT_LOOPBACK_MODE_UNKNOWN";
	}
}

const char *
ionic_xcvr_state_str(enum ionic_xcvr_state state)
{

	switch (state) {
	case IONIC_XCVR_STATE_REMOVED:
		return "XCVR_STATE_REMOVED";
	case IONIC_XCVR_STATE_INSERTED:
		return "XCVR_STATE_INSERTED";
	case IONIC_XCVR_STATE_PENDING:
		return "XCVR_STATE_PENDING";
	case IONIC_XCVR_STATE_SPROM_READ:
		return "XCVR_STATE_SPROM_READ";
	case IONIC_XCVR_STATE_SPROM_READ_ERR:
		return "XCVR_STATE_SPROM_READ_ERR";
	default:
		return "XCVR_STATE_UNKNOWN";
	}
}

const char *
ionic_phy_type_str(enum ionic_phy_type type)
{

	switch (type) {
	case IONIC_PHY_TYPE_NONE:
		return "PHY_TYPE_NONE";
	case IONIC_PHY_TYPE_COPPER:
		return "PHY_TYPE_COPPER";
	case IONIC_PHY_TYPE_FIBER:
		return "PHY_TYPE_FIBER";
	default:
		return "PHY_TYPE_UNKNOWN";
	}
}

const char *
ionic_error_to_str(enum ionic_status_code code)
{

	switch (code) {
	case IONIC_RC_SUCCESS:
		return "IONIC_RC_SUCCESS";
	case IONIC_RC_EVERSION:
		return "IONIC_RC_EVERSION";
	case IONIC_RC_EOPCODE:
		return "IONIC_RC_EOPCODE";
	case IONIC_RC_EIO:
		return "IONIC_RC_EIO";
	case IONIC_RC_EPERM:
		return "IONIC_RC_EPERM";
	case IONIC_RC_EQID:
		return "IONIC_RC_EQID";
	case IONIC_RC_EQTYPE:
		return "IONIC_RC_EQTYPE";
	case IONIC_RC_ENOENT:
		return "IONIC_RC_ENOENT";
	case IONIC_RC_EINTR:
		return "IONIC_RC_EINTR";
	case IONIC_RC_EAGAIN:
		return "IONIC_RC_EAGAIN";
	case IONIC_RC_ENOMEM:
		return "IONIC_RC_ENOMEM";
	case IONIC_RC_EFAULT:
		return "IONIC_RC_EFAULT";
	case IONIC_RC_EBUSY:
		return "IONIC_RC_EBUSY";
	case IONIC_RC_EEXIST:
		return "IONIC_RC_EEXIST";
	case IONIC_RC_EINVAL:
		return "IONIC_RC_EINVAL";
	case IONIC_RC_ENOSPC:
		return "IONIC_RC_ENOSPC";
	case IONIC_RC_ERANGE:
		return "IONIC_RC_ERANGE";
	case IONIC_RC_BAD_ADDR:
		return "IONIC_RC_BAD_ADDR";
	case IONIC_RC_DEV_CMD:
		return "IONIC_RC_DEV_CMD";
	case IONIC_RC_ERROR:
		return "IONIC_RC_ERROR";
	case IONIC_RC_ERDMA:
		return "IONIC_RC_ERDMA";
	default:
		return "IONIC_RC_UNKNOWN";
	}
}

int
ionic_error_to_errno(enum ionic_status_code code)
{
	switch (code) {
	case IONIC_RC_SUCCESS:
		return 0;
	case IONIC_RC_EVERSION:
	case IONIC_RC_EQTYPE:
	case IONIC_RC_EQID:
	case IONIC_RC_EINVAL:
	case IONIC_RC_ENOSUPP:
		return -EINVAL;
	case IONIC_RC_EPERM:
		return -EPERM;
	case IONIC_RC_ENOENT:
		return -ENOENT;
	case IONIC_RC_EAGAIN:
		return -EAGAIN;
	case IONIC_RC_ENOMEM:
		return -ENOMEM;
	case IONIC_RC_EFAULT:
		return -EFAULT;
	case IONIC_RC_EBUSY:
		return -EBUSY;
	case IONIC_RC_EEXIST:
		return -EEXIST;
	case IONIC_RC_ENOSPC:
		return -ENOSPC;
	case IONIC_RC_ERANGE:
		return -ERANGE;
	case IONIC_RC_BAD_ADDR:
		return -EFAULT;
	case IONIC_RC_EOPCODE:
	case IONIC_RC_EINTR:
	case IONIC_RC_DEV_CMD:
	case IONIC_RC_ERROR:
	case IONIC_RC_ERDMA:
	case IONIC_RC_EIO:
	default:
		return -EIO;
	}
}
EXPORT_SYMBOL_GPL(ionic_error_to_errno);

static const char *
ionic_opcode_to_str(enum ionic_cmd_opcode opcode)
{

	switch (opcode) {
	case IONIC_CMD_NOP:
		return "IONIC_CMD_NOP";
	case IONIC_CMD_INIT:
		return "IONIC_CMD_INIT";
	case IONIC_CMD_RESET:
		return "IONIC_CMD_RESET";
	case IONIC_CMD_IDENTIFY:
		return "IONIC_CMD_IDENTIFY";
	case IONIC_CMD_GETATTR:
		return "IONIC_CMD_GETATTR";
	case IONIC_CMD_SETATTR:
		return "IONIC_CMD_SETATTR";
	case IONIC_CMD_PORT_IDENTIFY:
		return "IONIC_CMD_PORT_IDENTIFY";
	case IONIC_CMD_PORT_INIT:
		return "IONIC_CMD_PORT_INIT";
	case IONIC_CMD_PORT_RESET:
		return "IONIC_CMD_PORT_RESET";
	case IONIC_CMD_PORT_GETATTR:
		return "IONIC_CMD_PORT_GETATTR";
	case IONIC_CMD_PORT_SETATTR:
		return "IONIC_CMD_PORT_SETATTR";
	case IONIC_CMD_LIF_INIT:
		return "IONIC_CMD_LIF_INIT";
	case IONIC_CMD_LIF_RESET:
		return "IONIC_CMD_LIF_RESET";
	case IONIC_CMD_LIF_IDENTIFY:
		return "IONIC_CMD_LIF_IDENTIFY";
	case IONIC_CMD_LIF_SETATTR:
		return "IONIC_CMD_LIF_SETATTR";
	case IONIC_CMD_LIF_GETATTR:
		return "IONIC_CMD_LIF_GETATTR";
	case IONIC_CMD_RX_MODE_SET:
		return "IONIC_CMD_RX_MODE_SET";
	case IONIC_CMD_RX_FILTER_ADD:
		return "IONIC_CMD_RX_FILTER_ADD";
	case IONIC_CMD_RX_FILTER_DEL:
		return "IONIC_CMD_RX_FILTER_DEL";
	case IONIC_CMD_Q_INIT:
		return "IONIC_CMD_Q_INIT";
	case IONIC_CMD_Q_CONTROL:
		return "IONIC_CMD_Q_CONTROL";
	case IONIC_CMD_RDMA_RESET_LIF:
		return "IONIC_CMD_RDMA_RESET_LIF";
	case IONIC_CMD_RDMA_CREATE_EQ:
		return "IONIC_CMD_RDMA_CREATE_EQ";
	case IONIC_CMD_RDMA_CREATE_CQ:
		return "IONIC_CMD_RDMA_CREATE_CQ";
	case IONIC_CMD_RDMA_CREATE_ADMINQ:
		return "IONIC_CMD_RDMA_CREATE_ADMINQ";
	case IONIC_CMD_FW_DOWNLOAD:
		return "IONIC_CMD_FW_DOWNLOAD";
	case IONIC_CMD_FW_CONTROL:
		return "IONIC_CMD_FW_CONTROL";
	default:
		return "DEVCMD_UNKNOWN";
	}
}

static void
ionic_adminq_flush(struct ionic_lif *lif)
{
	struct ionic_adminq *adminq = lif->adminq;
	struct ionic_admin_cmd *cmd;
	struct ionic_admin_ctx *ctx;
	int cmd_index;

	IONIC_ADMIN_LOCK(adminq);
	while (!IONIC_Q_EMPTY(adminq)) {
		cmd_index = adminq->tail_index;
		cmd = &adminq->cmd_ring[cmd_index];
		IONIC_QUE_WARN(adminq, "flushing tail: %d cmd %s(%d)\n",
		    adminq->tail_index, ionic_opcode_to_str(cmd->opcode),
		    cmd->opcode);
		ctx = adminq->ctx_ring[cmd_index];
		memset(cmd, 0, sizeof(*cmd));
		adminq->ctx_ring[cmd_index] = NULL;
		adminq->tail_index = IONIC_MOD_INC(adminq, tail_index);
	}
	IONIC_ADMIN_UNLOCK(adminq);

	IONIC_QUE_INFO(adminq, "head: %d tail: %d comp index: %d\n",
	    adminq->head_index, adminq->tail_index, adminq->comp_index);
}

static int
ionic_adminq_check_err(struct ionic_lif *lif, struct ionic_admin_ctx *ctx,
    bool timeout)
{
	struct net_device *netdev = lif->netdev;
	const char *name;
	const char *status;
	int err = 0;

	if (ctx->comp.comp.status || timeout) {
		name = ionic_opcode_to_str(ctx->cmd.cmd.opcode);
		status = ionic_error_to_str(ctx->comp.comp.status);
		IONIC_NETDEV_ERROR(netdev, "%s (%d) failed: %s (%d)\n",
		    name, ctx->cmd.cmd.opcode, timeout ? "TIMEOUT": status,
		    timeout ? -1 : ctx->comp.comp.status);

		err = timeout ? ETIMEDOUT :
			-ionic_error_to_errno(ctx->comp.comp.status);
		if (timeout)
			ionic_adminq_flush(lif);

	}

	return (err);
}

static void
ionic_adminq_ring_doorbell(struct ionic_adminq *adminq, int index)
{

	IONIC_QUE_INFO(adminq, "ring doorbell for index: %d\n", index);

	ionic_dbell_ring(adminq->lif->kern_dbpage,
			 adminq->hw_type,
			 adminq->dbval | index);
}

static bool
ionic_adminq_avail(struct ionic_adminq *adminq, int want)
{
	int avail;

	avail = ionic_desc_avail(adminq->num_descs,
	    adminq->head_index, adminq->tail_index);
	return (avail > want);
}

static int
ionic_adminq_post(struct ionic_lif *lif, struct ionic_admin_ctx *ctx)
{
	struct ionic_adminq *adminq = lif->adminq;
	struct ionic_admin_cmd *cmd;

	KASSERT(IONIC_LIF_LOCK_OWNED(lif), ("%s lif not locked", lif->name));
	IONIC_ADMIN_LOCK(adminq);

	if (adminq->stop) {
		IONIC_QUE_INFO(adminq, "can't post admin queue command\n");
		IONIC_ADMIN_UNLOCK(adminq);
		return (-ESHUTDOWN);
	}

	if (!ionic_adminq_avail(adminq, 1)) {
		IONIC_QUE_ERROR(adminq, "adminq is hung, head: %d tail: %d\n",
		    adminq->head_index, adminq->tail_index);
		IONIC_ADMIN_UNLOCK(adminq);
		return (-ENOSPC);
	}

	adminq->ctx_ring[adminq->head_index] = ctx;
	cmd = &adminq->cmd_ring[adminq->head_index];
	memcpy(cmd, &ctx->cmd, sizeof(ctx->cmd));

	IONIC_QUE_INFO(adminq, "post admin queue command %d@%d:\n",
	    cmd->opcode, adminq->head_index);
	if (__IONIC_DEBUG)
		print_hex_dump_debug("cmd ", DUMP_PREFIX_OFFSET, 16, 1,
		    &ctx->cmd, sizeof(ctx->cmd), true);

	adminq->head_index = IONIC_MOD_INC(adminq, head_index);
	ionic_adminq_ring_doorbell(adminq, adminq->head_index);

	IONIC_ADMIN_UNLOCK(adminq);

	return (0);
}

int
ionic_adminq_post_wait(struct ionic_lif *lif, struct ionic_admin_ctx *ctx)
{
	struct ionic_adminq *adminq = lif->adminq;
	struct ifnet *ifp = lif->netdev;
	struct ionic_dev *idev = &lif->ionic->idev;
	int err, remaining, processed, timeout;
	const char *name;

	KASSERT(IONIC_LIF_LOCK_OWNED(lif), ("%s lif not locked", lif->name));

	err = ionic_adminq_post(lif, ctx);
	if (err == -ESHUTDOWN) {
		name = ionic_opcode_to_str(ctx->cmd.cmd.opcode);
		IONIC_NETDEV_ERROR(ifp, "%s (%d) failed: adminq stopped\n",
		    name, ctx->cmd.cmd.opcode);
		return (err);
	}
	if (err) {
		IONIC_NETDEV_ERROR(ifp, "adminq_post failed, err: %d\n", err);
		return (err);
	}

	if ((ctx->cmd.cmd.opcode == IONIC_CMD_FW_DOWNLOAD) ||
		(ctx->cmd.cmd.opcode == IONIC_CMD_FW_CONTROL)) {
		timeout = ionic_fw_update_timeout;
		if (timeout < IONIC_FW_MIN_TIMEOUT) {
			if_printf(ifp, "Firmware timeout: %d is low, using default value\n", timeout);
			timeout = IONIC_FW_MIN_TIMEOUT;
			ionic_fw_update_timeout = timeout;
		}
	} else {
		timeout = ionic_devcmd_timeout;
	}

	remaining = wait_for_completion_timeout(&ctx->work,
						timeout * HZ);
	if (remaining == 0) {
		/* Check again in case the interrupt was missed */
		IONIC_ADMIN_LOCK(adminq);
		processed = ionic_adminq_clean(adminq, adminq->num_descs);
		if (processed) {
			ionic_intr_credits(idev->intr_ctrl, adminq->intr.index,
			    processed, IONIC_INTR_CRED_REARM);
			remaining = 1;
			IONIC_NETDEV_INFO(ifp, "adminq timeout avoided\n");
		} else {
			IONIC_NETDEV_ERROR(ifp, "adminq timeout\n");
		}
		IONIC_ADMIN_UNLOCK(adminq);
	}
	return (ionic_adminq_check_err(lif, ctx, remaining == 0));
}

int
ionic_set_dma_mask(struct ionic *ionic)
{
	struct device *dev = ionic->dev;
	int err;

	/* Set DMA addressing limitations. */
	err = dma_set_mask(dev, DMA_BIT_MASK(IONIC_ADDR_BITS));
	if (err) {
		IONIC_DEV_ERROR(dev, "No usable %d-bit DMA configuration, aborting\n",
		    IONIC_ADDR_BITS);
		return (err);
	}

	err = dma_set_coherent_mask(dev, DMA_BIT_MASK(IONIC_ADDR_BITS));
	if (err)
		IONIC_DEV_ERROR(dev, "Unable to obtain %d-bit DMA "
		    "for consistent allocations, aborting\n", IONIC_ADDR_BITS);

	dma_set_max_seg_size(dev, 2u * 1024 * 1024 * 1024);

	return (err);
}

int
ionic_identify(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	struct ionic_identity *ident = &ionic->ident;
	int i, err, nwords;

	ident->drv.os_type = IONIC_OS_TYPE_FREEBSD;
	ident->drv.os_dist = 0;
	strncpy(ident->drv.os_dist_str, "FreeBSD",
	    sizeof(ident->drv.os_dist_str) - 1);
	ident->drv.kernel_ver = __FreeBSD_version;
	snprintf(ident->drv.kernel_ver_str,
	    sizeof(ident->drv.kernel_ver_str) - 1, "%d", __FreeBSD_version);
	strncpy(ident->drv.driver_ver_str, DRV_VERSION,
	    sizeof(ident->drv.driver_ver_str) - 1);

	nwords = min(ARRAY_SIZE(ident->drv.words), ARRAY_SIZE(idev->dev_cmd_regs->data));
	IONIC_DEV_LOCK(ionic);
	for (i = 0; i < nwords; i++)
		iowrite32(ident->drv.words[i], &idev->dev_cmd_regs->data[i]);

	ionic_dev_cmd_identify(idev, IONIC_IDENTITY_VERSION_1);

	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);

	if (err) {
		IONIC_DEV_UNLOCK(ionic);
		goto err_out_unmap;
	}

	nwords = min(ARRAY_SIZE(ident->dev.words), ARRAY_SIZE(idev->dev_cmd_regs->data));
	for (i = 0; i < nwords; i++)
		ident->dev.words[i] = ioread32(&idev->dev_cmd_regs->data[i]);

	IONIC_DEV_UNLOCK(ionic);
	return (0);

err_out_unmap:
	return (err);
}

int
ionic_init(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	int err;

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_init(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

int
ionic_reset(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	int err;

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_reset(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

int
ionic_port_identify(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	struct ionic_identity *ident = &ionic->ident;
	int i, err, nwords;

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_port_identify(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	if (!err) {
		nwords = min(ARRAY_SIZE(ident->port.words), ARRAY_SIZE(idev->dev_cmd_regs->data));
		for (i = 0; i < nwords; i++)
			ident->port.words[i] = ioread32(&idev->dev_cmd_regs->data[i]);
	}
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

int
ionic_port_init(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_port_config *config;
	int i, err, nwords;

	if (idev->port_info)
		return (0);

	idev->port_info_sz = ALIGN(sizeof(*idev->port_info), PAGE_SIZE);

	err = ionic_dma_alloc(ionic, idev->port_info_sz, &ionic->port_dma, 0);
	if (err) {
		IONIC_DEV_ERROR(ionic->dev, "failed to allocate memory for port, err: %d\n", err);
		return (-ENOMEM);
	}

	idev->port_info = (struct ionic_port_info *)ionic->port_dma.dma_vaddr;
	if (!idev->port_info) {
		IONIC_DEV_ERROR(ionic->dev, "Failed to allocate port info, aborting\n");
		return (-ENOMEM);
	}

	IONIC_DEV_LOCK(ionic);
	idev->port_info_pa = ionic->port_dma.dma_paddr;
	nwords = min(ARRAY_SIZE(ident->port.config.words), ARRAY_SIZE(idev->dev_cmd_regs->data));
	config = &ident->port.config;

	if (!ionic->is_mgmt_nic)
		config->state = IONIC_PORT_ADMIN_STATE_UP;
	for (i = 0; i < nwords; i++)
		iowrite32(config->words[i], &idev->dev_cmd_regs->data[i]);

	ionic_dev_cmd_port_init(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

int
ionic_port_reset(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	int err;

	if (!idev->port_info)
		return (0);

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_port_reset(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);
	if (err) {
		IONIC_DEV_ERROR(ionic->dev, "Failed to reset port\n");
		return (err);
	}

	ionic_dma_free(ionic, &ionic->port_dma);

	idev->port_info = NULL;
	idev->port_info_pa = 0;

	return (0);
}

void
ionic_set_port_state(struct ionic *ionic, uint8_t state)
{
	struct ionic_dev *idev = &ionic->idev;
	int err;

	if (!idev->port_info)
		return;
	if (ionic->is_mgmt_nic)
		return;

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_port_state(idev, state);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	if (err)
		IONIC_DEV_ERROR(ionic->dev, "Failed to set port state %s, err: %d\n",
		    ionic_port_admin_state_str(state), err);
}

int
ionic_qos_class_identify(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	struct ionic_identity *ident = &ionic->ident;
	int i, err, nwords;

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_qos_class_identify(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	if (!err) {
		nwords = min(ARRAY_SIZE(ident->qos.words), ARRAY_SIZE(idev->dev_cmd_regs->data));
		for (i = 0; i < nwords; i++)
			ident->qos.words[i] = ioread32(&idev->dev_cmd_regs->data[i]);
	}
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

static void
ionic_qos_print(struct ionic *ionic, union ionic_qos_config *qos, int tc, char *ctx)
{

	IONIC_DEV_INFO(ionic->dev, "%s TC: %d pcp: %d weight: %d drop: %s "
	     " class: %d pause_type: %d pfc_cos: %d flags: 0x%x\n",
	     ctx, tc, qos->dot1q_pcp, qos->dwrr_weight,
	     (qos->flags & IONIC_QOS_CONFIG_F_NO_DROP) ? "false" : "true",
	     qos->class_type, qos->pause_type, qos->pfc_cos, qos->flags);
}

int
ionic_qos_init(struct ionic *ionic)
{
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	int i, tc, max_tcs, err;
	uint8_t dscp;

	err = ionic_qos_class_identify(ionic);
	if (err) {
		return (err);
	}

	/* XXX: for now hardcoded, should be read from identify */
	ionic->qos.max_tcs = IONIC_QOS_CLASS_MAX;
	max_tcs = ionic->qos.max_tcs;
	memset(ionic->qos.no_drop, 0, sizeof(ionic->qos.no_drop));
	memset(ionic->qos.pfc_cos, 0, sizeof(ionic->qos.pfc_cos));
	memset(ionic->qos.pcp_to_tc, 0, sizeof(ionic->qos.pcp_to_tc));
	memset(ionic->qos.dscp_to_tc, 0, sizeof(ionic->qos.dscp_to_tc));

	if (ionic->qos.class_type == IONIC_QOS_CLASS_TYPE_NONE)
		ionic->qos.class_type = IONIC_QOS_CLASS_TYPE_PCP;

	for (tc = 0; tc < max_tcs; tc++) {
		qos = &ident->qos.config[tc];
		ionic_qos_print(ionic, qos, tc, "Get");

		ionic->qos.enable_flag[tc] = (qos->flags & IONIC_QOS_CONFIG_F_ENABLE) ? 1 : 0;
		if (ionic->qos.enable_flag[tc] == 0)
			continue;

		ionic->qos.no_drop[tc] = (qos->flags & IONIC_QOS_CONFIG_F_NO_DROP) ? 1 : 0;
		ionic->qos.sched_type[tc] = qos->sched_type;
		ionic->qos.dwrr_bw_perc[tc] = qos->dwrr_weight;
		ionic->qos.pfc_cos[tc] = qos->pfc_cos;

		if ( (qos->class_type == IONIC_QOS_CLASS_TYPE_NONE) ||
		     (qos->class_type == IONIC_QOS_CLASS_TYPE_PCP) ) {
			if (qos->dot1q_pcp >= 0 && qos->dot1q_pcp < IONIC_QOS_PCP_MAX)
				ionic->qos.pcp_to_tc[qos->dot1q_pcp] = tc;
		}
		else if (qos->class_type == IONIC_QOS_CLASS_TYPE_DSCP){
			for (i = 0; i < qos->ndscp; i++) {
				dscp = qos->ip_dscp[i];
				if (dscp >= 0 && dscp < IONIC_QOS_DSCP_MAX)
					ionic->qos.dscp_to_tc[dscp] = tc;
			}
		}
		else {
			/*IONIC_NETDEV_ERROR(ifp,
			    "Invalid class type for TC: %d,", tc);*/
			return (EINVAL);
		}
	}

	return (0);
}

static int
ionic_qos_tc_init(struct ionic *ionic, int tc, union ionic_qos_config *qos)
{
	struct ionic_dev *idev = &ionic->idev;
	int i, err, nwords;

	ionic_qos_print(ionic, qos, tc, "Init");
	nwords = min(ARRAY_SIZE(qos->words), ARRAY_SIZE(idev->dev_cmd_regs->data));
	IONIC_DEV_LOCK(ionic);
	for (i = 0; i < nwords; i++)
		iowrite32(qos->words[i], &idev->dev_cmd_regs->data[i]);

	ionic_dev_cmd_qos_class_init(idev, tc);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

static int
ionic_qos_tc_update(struct ionic_lif *lif, int tc, union ionic_qos_config *qos)
{
	struct ionic *ionic = lif->ionic;
	struct ionic_dev *idev = &ionic->idev;
	int i, err, nwords;

	if (tc == 0) {
		IONIC_NETDEV_WARN(lif->netdev, "TC0 is read only\n");
		return (ENXIO);
	}

	ionic_qos_print(ionic, qos, tc, "Update");
	nwords = min(ARRAY_SIZE(qos->words), ARRAY_SIZE(idev->dev_cmd_regs->data));
	IONIC_DEV_LOCK(ionic);
	for (i = 0; i < nwords; i++)
		iowrite32(qos->words[i], &idev->dev_cmd_regs->data[i]);

	ionic_dev_cmd_qos_class_update(idev, tc);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

int
ionic_qos_class_reset(struct ionic *ionic, uint8_t group)
{
	struct ionic_dev *idev = &ionic->idev;
	int err;

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_qos_class_reset(idev, group);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

int
ionic_qos_clear_stats(struct ionic *ionic, uint8_t qos_group_bitmap)
{
	struct ionic_dev *idev = &ionic->idev;
	int err;

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_qos_clear_stats(idev, qos_group_bitmap);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

static void
ionic_qos_set_default_pcp(struct ionic_lif *lif, int tc, union ionic_qos_config *qos)
 {
 	struct ionic *ionic = lif->ionic;
 	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *search_qos;
 	int i;
	int def_pcp = -1;
	int tc_pcp = tc;
	bool tc_pcp_available = true;
	uint8_t bitmap = 0;

 	/*
 	 * By default, map the PCP to respective TC (TC <-> PCP one-to-one mapped);
 	 * if the PCP is already mapped to some other TC, pick the first available PCP
 	 */
	for (i = 1; i < ionic->qos.max_tcs; i++) {
		search_qos = &ident->qos.config[i];
		if ((search_qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0)
 			continue;
		if (search_qos->dot1q_pcp == IONIC_QOS_ALL_PCP)
 			continue;

		/* if existing config is DSCP, looking for used-up PCPs may not be necessary */
		if(search_qos->class_type == IONIC_QOS_CLASS_TYPE_DSCP)
			continue;

		bitmap |= BIT(search_qos->dot1q_pcp);
 		IONIC_NETDEV_INFO(lif->netdev, " tc: %d pause_type: %d pcp: %d "
 		    "sched_type: %d wt: %d PCP map: 0x%x\n",
		    i, search_qos->pause_type, search_qos->dot1q_pcp,
		    search_qos->sched_type, search_qos->dwrr_weight, bitmap);
 	}

	if ((bitmap & BIT(tc_pcp)) != 0) {
		tc_pcp_available = false;
		def_pcp = -1;
	}

	if (!tc_pcp_available) {
		/* TC PCP not available; pick the first free PCP */
		for (i = IONIC_QOS_PCP_MAX -1;  i >= 0; i--) {
			if ((bitmap & BIT(i)) == 0) {
				def_pcp = i;
				break;
			}
		}
	} else {
		def_pcp = tc_pcp;
	}

	if (def_pcp != -1)
		qos->dot1q_pcp = def_pcp;
}

static void
ionic_qos_set_default_dscp(struct ionic_lif *lif, int tc, union ionic_qos_config *qos)
{
	struct ionic *ionic = lif->ionic;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *search_qos;
	int i, j;
	int tc_dscp[IONIC_DSCP_BLOCK_SIZE] = {0};
	uint64_t bitmap = 0;
	int used_count = 0;

	/*
	 * By default, map the dscp range to respective TC;
	 * 	default DSCP range is TC*8 - TC*8+8
	 */
	for (i = (tc * IONIC_DSCP_BLOCK_SIZE), j = 0;
		i < ((tc * IONIC_DSCP_BLOCK_SIZE) + IONIC_DSCP_BLOCK_SIZE);
		i ++, j ++) {
		tc_dscp[j] = i;
	}

	qos->ndscp = 0;

	/*
	 * Get the bitmap of all the DSCPs mapped to all the TCs.
 	 */
	for (i = 1; i < ionic->qos.max_tcs; i++) {
		if (i == tc)
			continue;

		search_qos = &ident->qos.config[i];
		if ((search_qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0)
			continue;

		/* if existing config is PCP, looking for used-up DSCPs may not be necessary */
		if(search_qos->class_type == IONIC_QOS_CLASS_TYPE_PCP)
			continue;

		for (j = 0; j < search_qos->ndscp; j++)
			bitmap |= BIT(search_qos->ip_dscp[j]);
 	}

	/*
	 * if any of the dscp is already mapped to some other TC, unmap it.
	 * if all the DSCPs are already mapped to other TCs, pick the first avlbl dscp
	 */

	for (i = 0; i < IONIC_DSCP_BLOCK_SIZE; i ++) {
		if ((bitmap & BIT(tc_dscp[i])) != 0) {
			/* dscp is already mapped to some other TC, unmap it. */
			tc_dscp[i] = -1;
			used_count ++;
		}
	}

	if (used_count == IONIC_DSCP_BLOCK_SIZE) {
		/*
		 * all the DSCPs are already mapped to other TCs,
		 * pick the first avlbl dscp
		 */

		/* look at the higher range, first */
		for (i = ((tc * IONIC_DSCP_BLOCK_SIZE) + IONIC_DSCP_BLOCK_SIZE);
			i < IONIC_QOS_DSCP_MAX; i ++) {
			if ((bitmap & BIT(i)) == 0) {
				qos->ip_dscp[0] = i;
				qos->ndscp = 1;
				break;
			}
		}
		if (qos->ndscp == 0) {
			/* no DSCPs avlbl in higher range - look at lower range */
			for (i = ((tc * IONIC_DSCP_BLOCK_SIZE)-1); i >= 0; i --) {
				if ((bitmap & BIT(i)) == 0) {
					qos->ip_dscp[0] = i;
					qos->ndscp = 1;
					break;
				}
			}
		}

		if (qos->ndscp == 0) {
			/* no free DSCPs available */
		}
	} else {
		for (i = 0, j = 0; i < IONIC_DSCP_BLOCK_SIZE; i ++) {
			if (tc_dscp[i] != -1) {
				qos->ip_dscp[j] = (uint8_t) tc_dscp[i];
				qos->ndscp ++;
				j++;
			}
		}
	}
}

static void
ionic_qos_set_default(struct ionic_lif *lif, int tc, union ionic_qos_config *qos)
{
	struct ionic *ionic = lif->ionic;

	qos->pause_type = ionic->idev.port_info->config.pause_type & IONIC_PAUSE_TYPE_MASK;
	qos->sched_type = IONIC_QOS_SCHED_TYPE_DWRR;
	qos->dwrr_weight = 25; /* TODO: set this to 0 when bw_perc is enabled */

	ionic_qos_print(ionic, qos, tc, "Default");
}

#ifdef notyet
int
ionic_qos_bw_update(struct ionic_lif *lif, uint8_t *bw_perc)
{
	struct ionic *ionic = lif->ionic;
	struct ifnet *ifp = lif->netdev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	int tc, error;

	ionic_qos_class_identify(ionic);
	for (tc = 1; tc < ionic->qos.max_tcs; tc++) {
		qos = &ident->qos.config[tc];
		if ((qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0) {
			continue;
		}

		if (qos->sched_type != IONIC_QOS_SCHED_TYPE_DWRR) {
			IONIC_NETDEV_ERROR(ifp,
			    "First set QoS TC: %d policy to DWRR\n", tc);
			return (EINVAL);
		}

		if (qos->dwrr_weight == bw_perc[tc])
			continue;

		ionic_qos_set_default(lif, tc, qos);
		qos->dwrr_weight = bw_perc[tc];

		error = ionic_qos_tc_update(lif, tc, qos);
		if (error) {
			IONIC_NETDEV_ERROR(ifp,
			    "Failed to change bandwidth percentage of TC: %d, error: %d\n",
			    tc, error);
			return (error);
		}
	}

	return (0);
}
#endif
int
ionic_qos_pcp_to_tc_update(struct ionic_lif *lif, uint8_t *pcp)
{
	struct ionic *ionic = lif->ionic;
	struct ifnet *ifp = lif->netdev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	int i, tc, error;
	uint8_t new_pcp;
	bool update;

	ionic_qos_class_identify(ionic);
	for (tc = 1; tc < ionic->qos.max_tcs; tc++) {
		qos = &ident->qos.config[tc];

		if (qos->dot1q_pcp == IONIC_QOS_ALL_PCP)
			continue;
		/* XXX: Allow multiple PCP to TC mapping. */
		update = false;
		for (i = 0; i < IONIC_QOS_PCP_MAX; i++) {
			if (pcp[i] == tc) {
				update = true;
				new_pcp = i;
				break;
			}
		}
		/* FIX ME */
		if (!update && (qos->flags & IONIC_QOS_CONFIG_F_ENABLE)) {
			IONIC_NETDEV_ERROR(ifp, "TC%d PCP value missing\n", tc);
			return (EINVAL);
		}

		if (!update || (qos->dot1q_pcp == new_pcp))
			continue;

		if ((qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0 ||
		    (qos->class_type != IONIC_QOS_CLASS_TYPE_PCP)) {
			continue;
 		}

		IONIC_NETDEV_INFO(ifp, "TC%d new PCP value: %d\n", tc, new_pcp);

		ionic_qos_set_default(lif, tc, qos);
		qos->dot1q_pcp = new_pcp;

		error = ionic_qos_tc_update(lif, tc, qos);
		if (error) {
			IONIC_NETDEV_ERROR(ifp,
			    "Failed to change PCP for TC: %d, error: %d\n",
			    tc, error);
			return (error);
		}
	}

	return (0);
}

int
ionic_qos_pfc_cos_update(struct ionic_lif *lif, uint8_t *pfc_cos)
{
	struct ionic *ionic = lif->ionic;
	struct ifnet *ifp = lif->netdev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	int tc, error;

	ionic_qos_class_identify(ionic);
	for (tc = 1; tc < ionic->qos.max_tcs; tc++) {
		qos = &ident->qos.config[tc];

		if (qos->pfc_cos == pfc_cos[tc])
			continue;

		if ((qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0) {
			IONIC_NETDEV_WARN(ifp, "TC%d is not enabled\n", tc);
			return (EINVAL);
		}

		ionic_qos_set_default(lif, tc, qos);
		qos->pfc_cos = pfc_cos[tc];

		error = ionic_qos_tc_update(lif, tc, qos);
		if (error) {
			IONIC_NETDEV_ERROR(ifp,
			    "Failed to change CoS for TC: %d, error: %d\n",
			    tc, error);
			return (error);
		}
	}

	return (0);
}

/* Given the current and new DSCP configs for a TC, check if an update is required */
bool ionic_qos_is_tc_dscp_updated(union ionic_qos_config *qos, int new_ndscp, uint8_t *new_dscp)
{
	int i;

	/* If the ndscp changed, update is required */
	if (qos->ndscp != new_ndscp) {
		return true;
	}

	/* Check if there is difference in DSCP mapping */
	for (i = 0; i < new_ndscp; i++) {
		if (qos->ip_dscp[i] != new_dscp[i]) {
			return true;
		}
	}

	return false;
}

int
ionic_qos_dscp_to_tc_update(struct ionic_lif *lif, uint8_t *dscp)
{
	struct ionic *ionic = lif->ionic;
	struct ifnet *ifp = lif->netdev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	int i, tc, error, ndscp;
	uint8_t new_dscp[IONIC_QOS_DSCP_MAX];

	ionic_qos_class_identify(ionic);
	for (tc = 1; tc < ionic->qos.max_tcs; tc++) {
		qos = &ident->qos.config[tc];

		if ((qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0 ||
		    (qos->class_type != IONIC_QOS_CLASS_TYPE_DSCP)) {
			continue;
		}

		ndscp = 0;
		for (i = 0; i < IONIC_QOS_DSCP_MAX; i++) {
			if (dscp[i] == tc) {
				new_dscp[ndscp++] = i;
			}
		}

		if (!ndscp) {
			IONIC_NETDEV_ERROR(ifp, "TC%d DSCP value missing\n", tc);
			return (EINVAL);
		}

		/* Check if update is needed */
		if (!ionic_qos_is_tc_dscp_updated(qos, ndscp, new_dscp))
			continue;

		ionic_qos_set_default(lif, tc, qos);
		qos->ndscp = ndscp;
		for (i = 0; i < ndscp; i++) {
			qos->ip_dscp[i] = new_dscp[i];
		}

		error = ionic_qos_tc_update(lif, tc, qos);
		if (error) {
			IONIC_NETDEV_ERROR(ifp,
			    "Failed to change DSCP for TC: %d, error: %d\n",
			    tc, error);
			return (error);
		}
	}

	return (0);
}

int
ionic_qos_enable_update(struct ionic_lif *lif, uint8_t *enable)
{
	struct ionic *ionic = lif->ionic;
	struct ifnet *ifp = lif->netdev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	int tc, error;

	ionic_qos_class_identify(ionic);
	for (tc = 1; tc < ionic->qos.max_tcs; tc++) {
		qos = &ident->qos.config[tc];
		if (enable[tc] == (qos->flags & IONIC_QOS_CONFIG_F_ENABLE))
			continue;

		if (!enable[tc]) {
			/* Delete the class if it already exists, except 0 */
			if (qos->flags & IONIC_QOS_CONFIG_F_ENABLE) {
				IONIC_NETDEV_INFO(ifp, "Resetting QoS TC: %d\n", tc);
				ionic_qos_class_reset(ionic, tc);
			}
			continue;
		}

		qos->flags |= IONIC_QOS_CONFIG_F_ENABLE;
		ionic_qos_set_default(lif, tc, qos);

		if ((lif->ionic->qos.class_type == IONIC_QOS_CLASS_TYPE_NONE) ||
		    (lif->ionic->qos.class_type == IONIC_QOS_CLASS_TYPE_PCP)) {
			qos->ndscp = 0;
			memset(qos->ip_dscp, 0, sizeof(qos->ip_dscp));
			ionic_qos_set_default_pcp(lif, tc, qos);
			qos->class_type = IONIC_QOS_CLASS_TYPE_PCP;
		}
		else if (lif->ionic->qos.class_type == IONIC_QOS_CLASS_TYPE_DSCP) {
			qos->dot1q_pcp = 0;
			ionic_qos_set_default_dscp(lif, tc, qos);
			if(qos->ndscp == 0) {
				IONIC_NETDEV_ERROR(ifp,
					"Failed to find a free DSCP to map "
					"to TC: %d", tc);
				return (EINVAL);
			}
			qos->class_type = IONIC_QOS_CLASS_TYPE_DSCP;
		}
		else {
			IONIC_NETDEV_ERROR(ifp,
			    "Invalid class type for TC: %d,", tc);
			return (EINVAL);
 		}

		error = ionic_qos_tc_init(ionic, tc, qos);
		if (error) {
			IONIC_NETDEV_ERROR(ifp,
			    "Failed to enable QoS TC: %d, error: %d\n",
			    tc, error);
			return (error);
		}
	}

	return (0);
}

int
ionic_qos_no_drop_update(struct ionic_lif *lif, uint8_t *no_drop)
{
	struct ionic *ionic = lif->ionic;
	struct ifnet *ifp = lif->netdev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	int tc, error;

	ionic_qos_class_identify(ionic);
	for (tc = 1; tc < ionic->qos.max_tcs; tc++) {
		qos = &ident->qos.config[tc];

		if (no_drop[tc] == ((qos->flags & IONIC_QOS_CONFIG_F_NO_DROP) >> 1))
			continue;

		if ((qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0) {
			IONIC_NETDEV_WARN(ifp, "TC%d is not enabled\n", tc);
			return (EINVAL);
		}

		ionic_qos_set_default(lif, tc, qos);
		if (no_drop[tc])
			qos->flags |= IONIC_QOS_CONFIG_F_NO_DROP;
		else
			qos->flags &= ~IONIC_QOS_CONFIG_F_NO_DROP;

		qos->pause_type = ionic->idev.port_info->config.pause_type & IONIC_PAUSE_TYPE_MASK;

		error = ionic_qos_tc_update(lif, tc, qos);
		if (error) {
			IONIC_NETDEV_ERROR(ifp,
			    "Failed to set no-drop for TC: %d, error: %d\n",
			    tc, error);
			return (error);
		}
	}

	return (0);
}

int
ionic_qos_sched_type_update(struct ionic_lif *lif, uint8_t *sched)
{
	struct ionic *ionic = lif->ionic;
	struct ifnet *ifp = lif->netdev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	enum ionic_qos_sched_type sched_type;
	int tc, error;

	ionic_qos_class_identify(ionic);
	for (tc = 1; tc < ionic->qos.max_tcs; tc++) {
		qos = &ident->qos.config[tc];

		if (sched[tc])
			sched_type = IONIC_QOS_SCHED_TYPE_DWRR;
		else
			sched_type = IONIC_QOS_SCHED_TYPE_STRICT;
		if (qos->sched_type == sched_type)
			continue;

		if ((qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0) {
			IONIC_NETDEV_WARN(ifp, "TC%d is not enabled\n", tc);
			return (EINVAL);
		}

		ionic_qos_set_default(lif, tc, qos);
		qos->sched_type = sched_type;

		error = ionic_qos_tc_update(lif, tc, qos);
		if (error) {
			IONIC_NETDEV_ERROR(ifp,
			    "Failed to set QoS TC: %d, error: %d\n", tc, error);
			return (error);
		}
	}

	return (0);
}

int
ionic_qos_class_type_update(struct ionic_lif *lif,
    enum ionic_qos_class_type class)
{
	struct ionic *ionic = lif->ionic;
	struct ifnet *ifp = lif->netdev;
	struct ionic_identity *ident = &ionic->ident;
	union ionic_qos_config *qos;
	int tc, error;

	ionic_qos_class_identify(ionic);
	for (tc = 1; tc < ionic->qos.max_tcs; tc++) {
		qos = &ident->qos.config[tc];

		if (qos->class_type == class)
			continue;

		if ((qos->flags & IONIC_QOS_CONFIG_F_ENABLE) == 0) {
			continue;
		}

		ionic_qos_set_default(lif, tc, qos);
		qos->class_type = class;

		if ( (qos->class_type == IONIC_QOS_CLASS_TYPE_NONE) ||
		     (qos->class_type == IONIC_QOS_CLASS_TYPE_PCP) ) {
			qos->ndscp = 0;
			memset(qos->ip_dscp, 0, sizeof(qos->ip_dscp));
			ionic_qos_set_default_pcp(lif, tc, qos);
		}
		else if (qos->class_type == IONIC_QOS_CLASS_TYPE_DSCP) {
			qos->dot1q_pcp = 0;
			ionic_qos_set_default_dscp(lif, tc, qos);
			if(qos->ndscp == 0) {
				IONIC_NETDEV_ERROR(ifp,
					"Failed to find a free DSCP to map "
					"to TC: %d", tc);
				return (EINVAL);
			}
		}
		else {
			IONIC_NETDEV_ERROR(ifp,
			    "Invalid class type %d for TC: %d,", qos->class_type, tc);
			return (EINVAL);
		}

		error = ionic_qos_tc_update(lif, tc, qos);
		if (error) {
			IONIC_NETDEV_ERROR(ifp,
			    "Failed to change class type for TC: %d,"
			    " error: %d\n", tc, error);
			return (error);
		}
	}

	return (0);
}

int
ionic_qos_reset(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	int err;

	IONIC_DEV_LOCK(ionic);
	ionic_dev_cmd_qos_reset(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	IONIC_DEV_UNLOCK(ionic);

	return (err);
}

/*
 * Validate user parameters.
 */
static void
ionic_validate_params(void)
{
	const int div = 4;

	ionic_tx_descs = max(ionic_tx_descs, IONIX_TX_MIN_DESC);
	ionic_tx_descs = min(ionic_tx_descs, IONIX_TX_MAX_DESC);

	ionic_rx_descs = max(ionic_rx_descs, IONIX_RX_MIN_DESC);
	ionic_rx_descs = min(ionic_rx_descs, IONIX_RX_MAX_DESC);
	/* SGL size validation */
	if (ionic_rx_sg_size > MCLBYTES) {
		ionic_rx_sg_size = MJUMPAGESIZE;
	} else if (ionic_rx_sg_size) {
		ionic_rx_sg_size = MCLBYTES;
	}

	/* Doorbell stride has to be between 1 <=  < descs */
	ionic_rx_stride = max(ionic_rx_stride, 1);
	ionic_rx_stride = min(ionic_rx_stride, ionic_rx_descs / div);
	ionic_tx_stride = max(ionic_tx_stride, 1);
	ionic_tx_stride = min(ionic_tx_stride, ionic_tx_descs / div);

	/* Adjust Rx fill threshold. */
	if (ionic_rx_fill_threshold >= ionic_rx_descs / div)
		ionic_rx_fill_threshold /= div;
}

static int __init
ionic_init_module(void)
{

	ionic_validate_params();
	pr_info("%s, ver: %s\n", DRV_DESCRIPTION, DRV_VERSION);

	return (ionic_bus_register_driver());
}

static void __exit
ionic_cleanup_module(void)
{

	ionic_bus_unregister_driver();
}

module_init(ionic_init_module);
module_exit(ionic_cleanup_module);

MODULE_DEPEND(ionic, linuxkpi, 1, 1, 1);
