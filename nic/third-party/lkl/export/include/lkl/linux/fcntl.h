#ifndef _LKL_LINUX_FCNTL_H
#define _LKL_LINUX_FCNTL_H

#include <lkl/asm/fcntl.h>

#define LKL_F_SETLEASE	(LKL_F_LINUX_SPECIFIC_BASE + 0)
#define LKL_F_GETLEASE	(LKL_F_LINUX_SPECIFIC_BASE + 1)

/*
 * Cancel a blocking posix lock; internal use only until we expose an
 * asynchronous lock api to userspace:
 */
#define LKL_F_CANCELLK	(LKL_F_LINUX_SPECIFIC_BASE + 5)

/* Create a file descriptor with LKL_FD_CLOEXEC set. */
#define LKL_F_DUPFD_CLOEXEC	(LKL_F_LINUX_SPECIFIC_BASE + 6)

/*
 * Request nofications on a directory.
 * See below for events that may be notified.
 */
#define LKL_F_NOTIFY	(LKL_F_LINUX_SPECIFIC_BASE+2)

/*
 * Set and get of pipe page size array
 */
#define LKL_F_SETPIPE_SZ	(LKL_F_LINUX_SPECIFIC_BASE + 7)
#define LKL_F_GETPIPE_SZ	(LKL_F_LINUX_SPECIFIC_BASE + 8)

/*
 * Set/Get seals
 */
#define LKL_F_ADD_SEALS	(LKL_F_LINUX_SPECIFIC_BASE + 9)
#define LKL_F_GET_SEALS	(LKL_F_LINUX_SPECIFIC_BASE + 10)

/*
 * Types of seals
 */
#define LKL_F_SEAL_SEAL	0x0001	/* prevent further seals from being set */
#define LKL_F_SEAL_SHRINK	0x0002	/* prevent file from shrinking */
#define LKL_F_SEAL_GROW	0x0004	/* prevent file from growing */
#define LKL_F_SEAL_WRITE	0x0008	/* prevent writes */
/* (1U << 31) is reserved for signed error codes */

/*
 * Types of directory notifications that may be requested.
 */
#define LKL_DN_ACCESS	0x00000001	/* File accessed */
#define LKL_DN_MODIFY	0x00000002	/* File modified */
#define LKL_DN_CREATE	0x00000004	/* File created */
#define LKL_DN_DELETE	0x00000008	/* File removed */
#define LKL_DN_RENAME	0x00000010	/* File renamed */
#define LKL_DN_ATTRIB	0x00000020	/* File changed attibutes */
#define LKL_DN_MULTISHOT	0x80000000	/* Don't remove notifier */

#define LKL_AT_FDCWD		-100    /* Special value used to indicate
                                           openat should use the current
                                           working directory. */
#define LKL_AT_SYMLINK_NOFOLLOW	0x100   /* Do not follow symbolic links.  */
#define LKL_AT_REMOVEDIR		0x200   /* Remove directory instead of
                                           unlinking file.  */
#define LKL_AT_SYMLINK_FOLLOW	0x400   /* Follow symbolic links.  */
#define LKL_AT_NO_AUTOMOUNT		0x800	/* Suppress terminal automount traversal */
#define LKL_AT_EMPTY_PATH		0x1000	/* Allow empty relative pathname */

#define LKL_AT_STATX_SYNC_TYPE	0x6000	/* Type of synchronisation required from statx() */
#define LKL_AT_STATX_SYNC_AS_STAT	0x0000	/* - Do whatever stat() does */
#define LKL_AT_STATX_FORCE_SYNC	0x2000	/* - Force the attributes to be sync'd with the server */
#define LKL_AT_STATX_DONT_SYNC	0x4000	/* - Don't sync attributes with the server */


#endif /* _LKL_LINUX_FCNTL_H */
