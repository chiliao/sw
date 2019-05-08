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
#include <linux/dma-mapping.h>
#include <linux/pci.h>

#include "ionic.h"
#include "ionic_bus.h"
#include "ionic_lif.h"

MODULE_DESCRIPTION(DRV_DESCRIPTION);
MODULE_AUTHOR("Anish Gupta <anish@pensando.io>");
MODULE_VERSION(ionic, 1);

static const char *ionic_error_to_str(enum status_code code)
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

static const char *ionic_opcode_to_str(enum cmd_opcode opcode)
{
	switch (opcode) {
	case CMD_OPCODE_NOP:
		return "CMD_OPCODE_NOP";
	case CMD_OPCODE_INIT:
		return "CMD_OPCODE_INIT";
	case CMD_OPCODE_RESET:
		return "CMD_OPCODE_RESET";
	case CMD_OPCODE_IDENTIFY:
		return "CMD_OPCODE_IDENTIFY";
	case CMD_OPCODE_GETATTR:
		return "CMD_OPCODE_GETATTR";
	case CMD_OPCODE_SETATTR:
		return "CMD_OPCODE_SETATTR";
	case CMD_OPCODE_PORT_IDENTIFY:
		return "CMD_OPCODE_PORT_IDENTIFY";
	case CMD_OPCODE_PORT_INIT:
		return "CMD_OPCODE_PORT_INIT";
	case CMD_OPCODE_PORT_RESET:
		return "CMD_OPCODE_PORT_RESET";
	case CMD_OPCODE_PORT_GETATTR:
		return "CMD_OPCODE_PORT_GETATTR";
	case CMD_OPCODE_PORT_SETATTR:
		return "CMD_OPCODE_PORT_SETATTR";
	case CMD_OPCODE_LIF_INIT:
		return "CMD_OPCODE_LIF_INIT";
	case CMD_OPCODE_LIF_RESET:
		return "CMD_OPCODE_LIF_RESET";
	case CMD_OPCODE_LIF_IDENTIFY:
		return "CMD_OPCODE_LIF_IDENTIFY";
	case CMD_OPCODE_LIF_SETATTR:
		return "CMD_OPCODE_LIF_SETATTR";
	case CMD_OPCODE_LIF_GETATTR:
		return "CMD_OPCODE_LIF_GETATTR";
	case CMD_OPCODE_RX_MODE_SET:
		return "CMD_OPCODE_RX_MODE_SET";
	case CMD_OPCODE_RX_FILTER_ADD:
		return "CMD_OPCODE_RX_FILTER_ADD";
	case CMD_OPCODE_RX_FILTER_DEL:
		return "CMD_OPCODE_RX_FILTER_DEL";
	case CMD_OPCODE_Q_INIT:
		return "CMD_OPCODE_Q_INIT";
	case CMD_OPCODE_Q_CONTROL:
		return "CMD_OPCODE_Q_CONTROL";
	case CMD_OPCODE_RDMA_RESET_LIF:
		return "CMD_OPCODE_RDMA_RESET_LIF";
	case CMD_OPCODE_RDMA_CREATE_EQ:
		return "CMD_OPCODE_RDMA_CREATE_EQ";
	case CMD_OPCODE_RDMA_CREATE_CQ:
		return "CMD_OPCODE_RDMA_CREATE_CQ";
	case CMD_OPCODE_RDMA_CREATE_ADMINQ:
		return "CMD_OPCODE_RDMA_CREATE_ADMINQ";
	case CMD_OPCODE_FW_DOWNLOAD:
		return "CMD_OPCODE_FW_DOWNLOAD";
	case CMD_OPCODE_FW_CONTROL:
		return "CMD_OPCODE_FW_CONTROL";
	default:
		return "DEVCMD_UNKNOWN";
	}
}

int ionic_adminq_check_err(struct lif *lif, struct ionic_admin_ctx *ctx,
	bool timeout)
{
	struct net_device *netdev = lif->netdev;
	const char *name;
	const char *status;

	if (ctx->comp.comp.status || timeout) {
		name = ionic_opcode_to_str(ctx->cmd.cmd.opcode);
		status = ionic_error_to_str(ctx->comp.comp.status);
		IONIC_NETDEV_ERROR(netdev, "%s (%d) failed: %s (%d)\n",
			name,
			ctx->cmd.cmd.opcode,
			timeout ? "TIMEOUT": status,
			timeout ? -1 : ctx->comp.comp.status);
		return ctx->comp.comp.status;
	}

	return 0;
}

int ionic_adminq_post_wait(struct lif *lif, struct ionic_admin_ctx *ctx)
{
	int err, remaining;

	err = ionic_api_adminq_post(lif, ctx);
	if (err) {
		IONIC_NETDEV_ERROR(lif->netdev, "ionic_api_adminq_post failed, error: %d\n",
			err);
		return err;
	}

	remaining = wait_for_completion_timeout(&ctx->work, ionic_devcmd_timeout * HZ);

	err = ionic_adminq_check_err(lif, ctx, remaining == 0);

	return (err);
}

static int ionic_dev_cmd_wait(struct ionic_dev *idev, unsigned long max_wait)
{
	unsigned long time;
	int done;

	time = jiffies + max_wait;
	do {

		done = ionic_dev_cmd_done(idev);
#ifdef IONIC_DEBUG
		if (done)
			IONIC_INFO("DEVCMD done took %ld secs (%ld jiffies)\n",
			       (jiffies + max_wait - time)/HZ, jiffies + max_wait - time);
#endif
		if (done)
			return 0;

#ifdef __FreeBSD__
		/* XXX: use msleep but need mtx access. */
		DELAY(1000);
#else
		schedule_timeout_uninterruptible(HZ / 10);
#endif

	} while (time_after(time, jiffies));

	IONIC_ERROR("DEVCMD timeout after %ld secs\n", max_wait / HZ);

	return ETIMEDOUT;
}

static int ionic_dev_cmd_check_error(struct ionic_dev *idev)
{
	u8 status;

	status = ionic_dev_cmd_status(idev);

	if (status) {
		IONIC_ERROR("DEVCMD(%d) failed, status: %s\n",
			idev->dev_cmd_regs->cmd.cmd.opcode, ionic_error_to_str(status));
		return (EIO);
	}

	return status;
}

int ionic_dev_cmd_wait_check(struct ionic_dev *idev, unsigned long max_wait)
{
	int err;

	err = ionic_dev_cmd_wait(idev, max_wait);
	if (err)
		return err;
	return ionic_dev_cmd_check_error(idev);
}

int ionic_set_dma_mask(struct ionic *ionic)
{
	struct device *dev = ionic->dev;
	int err;

	/* Set DMA addressing limitations. */
	err = dma_set_mask(dev, DMA_BIT_MASK(IONIC_ADDR_BITS));
	if (err) {
		dev_err(dev, "No usable %d-bit DMA configuration, aborting\n", IONIC_ADDR_BITS);
		return err;
	}

	err = dma_set_coherent_mask(dev, DMA_BIT_MASK(IONIC_ADDR_BITS));
	if (err)
		dev_err(dev, "Unable to obtain %d-bit DMA "
			"for consistent allocations, aborting\n", IONIC_ADDR_BITS);

	dma_set_max_seg_size(dev, 2u * 1024 * 1024 * 1024);

	return err;
}

int ionic_identify(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	struct identity *ident = &ionic->ident;
	int err;
	unsigned int i;
	unsigned int nwords;

	ident->drv.os_type = IONIC_OS_TYPE_FREEBSD;
	ident->drv.os_dist = 0;
	strncpy(ident->drv.os_dist_str, "FreeBSD",
		sizeof(ident->drv.os_dist_str) - 1);
	ident->drv.kernel_ver = __FreeBSD_version;
	snprintf(ident->drv.kernel_ver_str, sizeof(ident->drv.kernel_ver_str) - 1,
		"%d", __FreeBSD_version);
	strncpy(ident->drv.driver_ver_str, DRV_VERSION,
		sizeof(ident->drv.driver_ver_str) - 1);

	nwords = min(ARRAY_SIZE(ident->drv.words), ARRAY_SIZE(idev->dev_cmd_regs->data));
	for (i = 0; i < nwords; i++)
		iowrite32(ident->drv.words[i], &idev->dev_cmd_regs->data[i]);

	ionic_dev_cmd_identify(idev, IONIC_IDENTITY_VERSION_1);

	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	if (err)
		goto err_out_unmap;

	nwords = min(ARRAY_SIZE(ident->dev.words), ARRAY_SIZE(idev->dev_cmd_regs->data));
	for (i = 0; i < nwords; i++)
		ident->dev.words[i] = ioread32(&idev->dev_cmd_regs->data[i]);

	return 0;

err_out_unmap:
	return err;
}

int ionic_init(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;

	ionic_dev_cmd_init(idev);
	return ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
}

int ionic_reset(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;

	ionic_dev_cmd_reset(idev);
	return ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
}

int ionic_port_identify(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	struct identity *ident = &ionic->ident;
	int err;
	unsigned int i;
	unsigned int nwords;

	ionic_dev_cmd_port_identify(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	if (!err) {
		nwords = min(ARRAY_SIZE(ident->port.words),
						ARRAY_SIZE(idev->dev_cmd_regs->data));
		for (i = 0; i < nwords; i++)
			ident->port.words[i] = ioread32(&idev->dev_cmd_regs->data[i]);
	}

	return err;
}

int ionic_port_init(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	struct identity *ident = &ionic->ident;
	int err;
	unsigned int i;
	unsigned int nwords;
	union port_config *config;

	if (idev->port_info)
		return 0;

	idev->port_info_sz = ALIGN(sizeof(*idev->port_info), PAGE_SIZE);
	idev->port_info = dma_alloc_coherent(ionic->dev, idev->port_info_sz,
					      &idev->port_info_pa,
					      GFP_KERNEL);
	if (!idev->port_info) {
		dev_err(ionic->dev, "Failed to allocate port info, aborting\n");
		return -ENOMEM;
	}

	nwords = min(ARRAY_SIZE(ident->port.config.words),
					ARRAY_SIZE(idev->dev_cmd_regs->data));
	config = &ident->port.config;
	/* Bring the physical port up. */
	if (!ionic->is_mgmt_nic)
		config->state = PORT_ADMIN_STATE_UP;
	for (i = 0; i < nwords; i++)
		iowrite32(config->words[i], &idev->dev_cmd_regs->data[i]);

	ionic_dev_cmd_port_init(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);

	return err;
}

int ionic_port_reset(struct ionic *ionic)
{
	struct ionic_dev *idev = &ionic->idev;
	int err;

	if (!idev->port_info)
		return 0;

	ionic_dev_cmd_port_reset(idev);
	err = ionic_dev_cmd_wait_check(idev, ionic_devcmd_timeout * HZ);
	if (err) {
		dev_err(ionic->dev, "Failed to reset port\n");
		return err;
	}

	dma_free_coherent(ionic->dev, idev->port_info_sz,
			  idev->port_info, idev->port_info_pa);

	idev->port_info = NULL;
	idev->port_info_pa = 0;

	return 0;
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

static int __init ionic_init_module(void)
{

	ionic_struct_size_checks();
	ionic_validate_params();
	pr_info("%s, ver: %s\n", DRV_DESCRIPTION, DRV_VERSION);

	return ionic_bus_register_driver();
}

static void __exit ionic_cleanup_module(void)
{
	ionic_bus_unregister_driver();
}

module_init(ionic_init_module);
module_exit(ionic_cleanup_module);

MODULE_DEPEND(ionic, linuxkpi, 1, 1, 1);
