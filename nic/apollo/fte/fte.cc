//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file contains FTE core functionality
///
//----------------------------------------------------------------------------
#include <cstring>
#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include <stdint.h>
#include <inttypes.h>
#include <sys/types.h>
#include <sys/queue.h>
#include <netinet/in.h>
#include <setjmp.h>
#include <stdarg.h>
#include <ctype.h>
#include <errno.h>
#include <getopt.h>
#include <signal.h>
#include <stdbool.h>

#include <rte_common.h>
#include <rte_log.h>
#include <rte_malloc.h>
#include <rte_memory.h>
#include <rte_memcpy.h>
#include <rte_eal.h>
#include <rte_launch.h>
#include <rte_atomic.h>
#include <rte_cycles.h>
#include <rte_prefetch.h>
#include <rte_lcore.h>
#include <rte_per_lcore.h>
#include <rte_branch_prediction.h>
#include <rte_interrupts.h>
#include <rte_random.h>
#include <rte_debug.h>
#include <rte_ether.h>
#include <rte_ethdev.h>
#include <rte_mempool.h>
#include <rte_mbuf.h>

#include "nic/sdk/lib/thread/thread.hpp"
#include "nic/apollo/core/trace.hpp"
#include "nic/apollo/fte/fte.hpp"

namespace fte {

char const * g_eal_args[] = {"fte", "-l", "2,3", "--vdev=net_ionic0"};
#define RTE_LOGTYPE_FTE RTE_LOGTYPE_USER1
#define NELEMS(_a) (sizeof(_a)/sizeof(_a[0]))

#define MAX_RX_QUEUE_PER_LCORE 16
#define MAX_TX_QUEUE_PER_PORT 16

#define FTE_PID 0 // Default PORTID
#define FTE_QID 0 // Default QueueID
#define FTE_MAX_CORES 1 // Max Cores
#define FTE_MAX_TXDSCR 256 // Max TX Descriptors
#define FTE_MAX_RXDSCR 256 // Max RX Descriptors
#define FTE_PKT_BATCH_SIZE 32 // Packet Batch Size
#define FTE_MEMPOOL_SIZE 256
#define FTE_PREFETCH_NLINES 7

// Global state
typedef struct gls_s {
    struct rte_mempool* mbfpool;
    uint16_t ntxdscr;
    uint16_t nrxdscr;
    uint32_t nmbfs;
    struct rte_eth_dev_tx_buffer* txbf;
    stats_t stats;
    uint64_t ftlstats[sdk::SDK_RET_MAX];
} __rte_cache_aligned gls_t;

gls_t gls = { NULL, FTE_MAX_TXDSCR, FTE_MAX_RXDSCR };

static void
_process (struct rte_mbuf *m)
{
    //gls.ftlstats[ret]++;

    fte_flow_prog(m);

    int numtx = rte_eth_tx_buffer(FTE_PID, FTE_QID, gls.txbf, m);
    if (numtx) {
        gls.stats.tx += numtx;
    }
    // TODO: Free is needed??
    //rte_pktmbuf_free(m);
}

// main processing loop
static void
fte_rx_loop (void)
{
    struct rte_mbuf *pkts_burst[FTE_PKT_BATCH_SIZE];
    int numrx, numtx;

    PDS_TRACE_DEBUG("\nFTE fte_rx_loop.. core:%u\n", rte_lcore_id());
    while (1) {
        numrx = rte_eth_rx_burst(0, 0, pkts_burst, FTE_PKT_BATCH_SIZE);
        if (!numrx) {
            continue;
        }

        PDS_TRACE_DEBUG("\n\nFTE receives %d packets.. \n\n", numrx);
        gls.stats.rx += numrx;
        for (int i = 0; i < numrx; i++) {
            auto m = pkts_burst[i];
            if ((i+1) < numrx) {
                auto m2 = pkts_burst[i+1];
                uint8_t *d2 = rte_pktmbuf_mtod(m2, uint8_t*);
                for (int i = 0; i < FTE_PREFETCH_NLINES; i++) {
                    rte_prefetch0(d2+i*64);
                }
            }
            _process(m);
        }
//#ifndef FTE_REINJECT_DISABLE
        numtx = rte_eth_tx_buffer_flush(0, 0, gls.txbf);
        if (numtx) {
            // TODO: Duplicate increment? Check _process as well.
            gls.stats.tx += numtx;
        }
//#endif
    }
}

static int
fte_launch_one_lcore (__attribute__((unused)) void *dummy)
{
    fte_ftl_set_core_id(rte_lcore_id());
    fte_rx_loop();
    return 0;
}

static void
_init_gls (void) 
{
    uint16_t nports = 0;
    uint32_t nmbfs = 0;

    nports = rte_eth_dev_count_avail();
    if (nports == 0) {
        rte_exit(EXIT_FAILURE, "No Ethernet ports - bye\n");
    }

    nmbfs = nports * (FTE_MAX_RXDSCR + FTE_MAX_TXDSCR + FTE_PKT_BATCH_SIZE +
                      FTE_MAX_CORES * FTE_MEMPOOL_SIZE);
    gls.nmbfs = RTE_MAX(nmbfs, 8192U);

    // create the mbuf pool
    gls.mbfpool = rte_pktmbuf_pool_create("mbuf_pool", gls.nmbfs,
                                          FTE_MEMPOOL_SIZE,
                                          0, RTE_MBUF_DEFAULT_BUF_SIZE,
                                          rte_socket_id());
    if (gls.mbfpool == NULL) {
        rte_exit(EXIT_FAILURE, "Cannot init mbuf pool\n");
    }

    return;
}

static void
_init_port (uint16_t portid)
{
    int ret = 0;
    struct rte_eth_rxconf rxq_conf;
    struct rte_eth_txconf txq_conf;
    struct rte_eth_conf local_port_conf = {0};
    struct rte_eth_dev_info dev_info;

    printf("Initializing port %u... \n", portid);
    rte_eth_dev_info_get(portid, &dev_info);
    if (dev_info.tx_offload_capa & DEV_TX_OFFLOAD_MBUF_FAST_FREE) {
        local_port_conf.txmode.offloads |= DEV_TX_OFFLOAD_MBUF_FAST_FREE;
    }
    ret = rte_eth_dev_configure(portid, 1, 1, &local_port_conf);
    if (ret < 0) {
        rte_exit(EXIT_FAILURE, "Cannot configure device: err=%d, port=%u\n",
                 ret, portid);
    }

    ret = rte_eth_dev_adjust_nb_rx_tx_desc(portid, &gls.nrxdscr, &gls.ntxdscr);
    if (ret < 0) {
        rte_exit(EXIT_FAILURE, "Descriptors updated failed: err=%d, port=%u\n",
                 ret, FTE_PID);
    }

    // init one RX queue
    rxq_conf = dev_info.default_rxconf;
    rxq_conf.offloads = local_port_conf.rxmode.offloads;
    ret = rte_eth_rx_queue_setup(portid, 0, gls.nrxdscr,
                                 rte_eth_dev_socket_id(portid),
                                 &rxq_conf, gls.mbfpool);
    if (ret < 0) {
        rte_exit(EXIT_FAILURE, "rte_eth_rx_queue_setup:err=%d, port=%u\n",
                 ret, portid);
    }

    txq_conf = dev_info.default_txconf;
    txq_conf.offloads = local_port_conf.txmode.offloads;
    ret = rte_eth_tx_queue_setup(portid, 0, gls.ntxdscr,
                                 rte_eth_dev_socket_id(portid),
                                 &txq_conf);
    if (ret < 0) {
        rte_exit(EXIT_FAILURE, "rte_eth_tx_queue_setup:err=%d, port=%u\n",
                 ret, portid);
    }

    return;
}

static void
_init_txbf (uint16_t portid)
{
    int ret = 0;

    // Initialize TX buffers
    gls.txbf = (rte_eth_dev_tx_buffer*)rte_zmalloc_socket("tx_buffer",
                                  RTE_ETH_TX_BUFFER_SIZE(FTE_PKT_BATCH_SIZE), 0,
                                  rte_eth_dev_socket_id(portid));
    if (gls.txbf == NULL) {
        rte_exit(EXIT_FAILURE, "Cannot allocate buffer for tx on port %u\n",
                portid);
    }

    rte_eth_tx_buffer_init(gls.txbf, FTE_PKT_BATCH_SIZE);

    ret = rte_eth_tx_buffer_set_err_callback(gls.txbf,
                                rte_eth_tx_buffer_count_callback,
                                &gls.stats.drop);
    if (ret < 0) {
        rte_exit(EXIT_FAILURE,
                 "Cannot set error callback for tx buffer on port %u\n",
                 portid);
    }

    return;
}

static void
signal_handler (int signum)
{
    if (signum == SIGINT || signum == SIGTERM) {
        PDS_TRACE_DEBUG("\nSIGNAL %d received..core#:%u\n", signum, rte_lcore_id());
        rte_exit(0, "received signal: %d, exiting\n", signum);
    }
}

static int
fte_main (void)
{
    int ret;
 sdk_ret_t sdk_ret;

    // init EAL
    ret = rte_eal_init(NELEMS(g_eal_args), (char**)g_eal_args);
    if (ret < 0) {
        rte_exit(EXIT_FAILURE, "Invalid EAL arguments\n");
    }

    signal(SIGINT, signal_handler);
    signal(SIGTERM, signal_handler);

    // Initialize Global State
    _init_gls();

    // Configure Port
    _init_port(FTE_PID);

    // Initialize TX Buffer
    _init_txbf(FTE_PID);

    // Start device
    ret = rte_eth_dev_start(FTE_PID);
    if (ret < 0) {
        rte_exit(EXIT_FAILURE, "rte_eth_dev_start:err=%d, port=%u\n",
                 ret, FTE_PID);
    }
    rte_eth_promiscuous_enable(FTE_PID);

    // Init FTL
    if ((sdk_ret = fte_ftl_init()) != SDK_RET_OK) {
        rte_exit(EXIT_FAILURE, "fte_ftl_init:err=%d",
                 sdk_ret);
    }
    ret = 0;
    // launch per-lcore init on every lcore
    rte_eal_mp_remote_launch(fte_launch_one_lcore, NULL, SKIP_MASTER);
    return ret;
}

void
fte_init (void)
{
    PDS_TRACE_DEBUG("FTE entering forever loop ...");

    fte_main();

    return;
}

} // namespace fte
