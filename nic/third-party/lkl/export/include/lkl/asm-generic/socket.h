#ifndef __LKL__ASM_GENERIC_SOCKET_H
#define __LKL__ASM_GENERIC_SOCKET_H

#include <lkl/asm/sockios.h>

/* For setsockopt(2) */
#define LKL_SOL_SOCKET	1

#define LKL_SO_DEBUG	1
#define LKL_SO_REUSEADDR	2
#define LKL_SO_TYPE		3
#define LKL_SO_ERROR	4
#define LKL_SO_DONTROUTE	5
#define LKL_SO_BROADCAST	6
#define LKL_SO_SNDBUF	7
#define LKL_SO_RCVBUF	8
#define LKL_SO_SNDBUFFORCE	32
#define LKL_SO_RCVBUFFORCE	33
#define LKL_SO_KEEPALIVE	9
#define LKL_SO_OOBINLINE	10
#define LKL_SO_NO_CHECK	11
#define LKL_SO_PRIORITY	12
#define LKL_SO_LINGER	13
#define LKL_SO_BSDCOMPAT	14
#define LKL_SO_REUSEPORT	15
#ifndef LKL_SO_PASSCRED /* powerpc only differs in these */
#define LKL_SO_PASSCRED	16
#define LKL_SO_PEERCRED	17
#define LKL_SO_RCVLOWAT	18
#define LKL_SO_SNDLOWAT	19
#define LKL_SO_RCVTIMEO	20
#define LKL_SO_SNDTIMEO	21
#endif

/* Security levels - as per NRL IPv6 - don't actually do anything */
#define LKL_SO_SECURITY_AUTHENTICATION		22
#define LKL_SO_SECURITY_ENCRYPTION_TRANSPORT	23
#define LKL_SO_SECURITY_ENCRYPTION_NETWORK		24

#define LKL_SO_BINDTODEVICE	25

/* Socket filtering */
#define LKL_SO_ATTACH_FILTER	26
#define LKL_SO_DETACH_FILTER	27
#define LKL_SO_GET_FILTER		LKL_SO_ATTACH_FILTER

#define LKL_SO_PEERNAME		28
#define LKL_SO_TIMESTAMP		29
#define LKL_SCM_TIMESTAMP		LKL_SO_TIMESTAMP

#define LKL_SO_ACCEPTCONN		30

#define LKL_SO_PEERSEC		31
#define LKL_SO_PASSSEC		34
#define LKL_SO_TIMESTAMPNS		35
#define LKL_SCM_TIMESTAMPNS		LKL_SO_TIMESTAMPNS

#define LKL_SO_MARK			36

#define LKL_SO_TIMESTAMPING		37
#define LKL_SCM_TIMESTAMPING	LKL_SO_TIMESTAMPING

#define LKL_SO_PROTOCOL		38
#define LKL_SO_DOMAIN		39

#define LKL_SO_RXQ_OVFL             40

#define LKL_SO_WIFI_STATUS		41
#define LKL_SCM_WIFI_STATUS	LKL_SO_WIFI_STATUS
#define LKL_SO_PEEK_OFF		42

/* Instruct lower device to use last 4-bytes of skb data as FCS */
#define LKL_SO_NOFCS		43

#define LKL_SO_LOCK_FILTER		44

#define LKL_SO_SELECT_ERR_QUEUE	45

#define LKL_SO_BUSY_POLL		46

#define LKL_SO_MAX_PACING_RATE	47

#define LKL_SO_BPF_EXTENSIONS	48

#define LKL_SO_INCOMING_CPU		49

#define LKL_SO_ATTACH_BPF		50
#define LKL_SO_DETACH_BPF		LKL_SO_DETACH_FILTER

#define LKL_SO_ATTACH_REUSEPORT_CBPF	51
#define LKL_SO_ATTACH_REUSEPORT_EBPF	52

#define LKL_SO_CNX_ADVICE		53

#define LKL_SCM_TIMESTAMPING_OPT_STATS	54

#endif /* __LKL__ASM_GENERIC_SOCKET_H */
