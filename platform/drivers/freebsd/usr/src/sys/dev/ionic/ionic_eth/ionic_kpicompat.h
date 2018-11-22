/*
 * Copyright 2017-2018 Pensando Systems, Inc.  All rights reserved.
 *
 * This program is free software; you may redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; version 2 of the License.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 * EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
 * NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS
 * BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
 * ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */

#ifndef _IONIC_KCOMPAT_H_
#define _IONIC_KCOMPAT_H_

/* XXX : should go to linux/printk.h */

#include <linux/printk.h>

#ifndef print_hex_dump_debug
#define print_hex_dump_debug(...) 			\
		print_hex_dump(NULL, __VA_ARGS__)
#endif

/* XXX : should go to linux/device.h */

#include <linux/device.h>

#undef dev_dbg
#define	dev_dbg(dev, fmt, ...) do {			\
	if (0)						\
		device_printf((dev)->bsddev, fmt, ##__VA_ARGS__); \
} while (0)

/* XXX : should go to linux/netdevice.h */

#include <linux/netdevice.h>

#ifndef netdev_name
#define netdev_name if_name
#endif

#ifndef netdev_emerg
#define netdev_emerg(...) if_printf(__VA_ARGS__)
#define netdev_alert(...) if_printf(__VA_ARGS__)
#define netdev_crit(...) if_printf(__VA_ARGS__)
#define netdev_err(...) if_printf(__VA_ARGS__)
#define netdev_warn(...) if_printf(__VA_ARGS__)
#define netdev_notice(...) if_printf(__VA_ARGS__)
#define netdev_info(...) if_printf(__VA_ARGS__)
#define netdev_dbg(...) do {				\
	if (0)						\
		if_printf(__VA_ARGS__);			\
} while (0)
#endif

/* XXX : should go to linux/addrconf.h */

//#include <linux/addrconf.h>

#define addrconf_ifid_eui48 __addrconf_ifid_eui48
static inline int __addrconf_ifid_eui48(u8 *eui, struct net_device *dev)
{
	const u8 *addr = IF_LLADDR(dev);

	eui[0] = addr[0] ^ 2;
	eui[1] = addr[1];
	eui[2] = addr[2];
	eui[3] = 0xff;
	eui[4] = 0xfe;
	eui[5] = addr[3];
	eui[6] = addr[4];
	eui[7] = addr[5];

	return 0;
}

/* XXX : should go to linux/completion.h */

#include <linux/completion.h>

#ifndef	INIT_COMPLETION_ONSTACK
#define	INIT_COMPLETION_ONSTACK(c) \
	((c).done = 0)
#define	COMPLETION_INITIALIZER(c) \
        { .done = 0 }
#define	COMPLETION_INITIALIZER_ONSTACK(c) \
        { .done = 0 }
#endif

#ifndef wait_for_completion_interruptible
#define wait_for_completion_interruptible(...) \
	wait_for_completion_interuptible(__VA_ARGS__)
#endif

/* XXX: should go to linux/kernel.h */

#include <linux/kernel.h>

#ifndef	IS_ALIGNED
#define	IS_ALIGNED(x, y) (((x) & (y - 1)) == 0)
#endif

/* XXX: should go to linux/mm.h */

#include <linux/mm.h>

#ifndef PAGE_ALIGNED
#define	PAGE_ALIGNED(x)	IS_ALIGNED((unsigned long)x, PAGE_SIZE)
#endif

/* XXX: should go to linux/notifier.h */

#include <linux/notifier.h>

#ifndef	notifier_from_errno
#define notifier_from_errno(v) NOTIFY_DONE
#endif

#endif
