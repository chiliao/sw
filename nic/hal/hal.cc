#include <iostream>
#include <stdio.h>
#include <string>
#include <errno.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <hal.hpp>
#include <hal_pd.hpp>
#include <periodic/periodic.hpp>
#include <lif_manager.hpp>
#include <rdma.hpp>
#include "boost/property_tree/ptree.hpp"
#include "boost/property_tree/json_parser.hpp"
#include <interface.hpp>
#include <tcpcb.hpp>
#include <proxy.hpp>
#include <fte.hpp>
#include <plugins/plugins.hpp>

extern "C" void __gcov_flush(void);

#ifdef COVERAGE
#define HAL_GCOV_FLUSH()     { ::__gcov_flush(); }
#else
#define HAL_GCOV_FLUSH()     { }
#endif

namespace hal {

// process globals
thread    *g_hal_threads[HAL_THREAD_ID_MAX];
uint64_t  g_hal_handle = 1;
bool      gl_super_user = false;

// TODO_CLEANUP: THIS DOESN'T BELONG HERE !!
LIFManager *g_lif_manager = nullptr;

// thread local variables
thread_local thread *t_curr_thread;

using boost::property_tree::ptree;
//------------------------------------------------------------------------------
// TODO - dummy for now !!
//------------------------------------------------------------------------------
static void *
fte_pkt_loop (void *ctxt)
{
    t_curr_thread = (thread *)ctxt;
    HAL_TRACE_DEBUG("Thread {} initializing ...", t_curr_thread->name());

    HAL_THREAD_INIT();
    while (1);
    return NULL;
}

//------------------------------------------------------------------------------
// allocate a handle for an object instance
// TODO: if this can be called from FTE, we need atomic increments
//------------------------------------------------------------------------------
hal_handle_t
hal_alloc_handle (void)
{
    return g_hal_handle++;
}

//------------------------------------------------------------------------------
// return a hal handle back so it can be reallocated for another object
//------------------------------------------------------------------------------
void
hal_free_handle (uint64_t handle)
{
    return;
}

//------------------------------------------------------------------------------
// initialize all the signal handlers
//------------------------------------------------------------------------------
static void
hal_sig_handler (int sig, siginfo_t *info, void *ptr)
{
    HAL_TRACE_DEBUG("HAL received signal {}", sig);

    switch (sig) {
    case SIGINT:
    case SIGKILL:
        HAL_GCOV_FLUSH();
        utils::hal_logger().flush();
        exit(0);
        break;

    case SIGUSR1:
    case SIGUSR2:
        HAL_GCOV_FLUSH();
        utils::hal_logger().flush();
        break;

    case SIGHUP:
    case SIGQUIT:
    case SIGSEGV:
    case SIGCHLD:
    case SIGURG:
    case SIGTERM:
    default:
        utils::hal_logger().flush();
        break;
    }
}

//------------------------------------------------------------------------------
// initialize all the signal handlers
// TODO: save old handlers and restore when signal happened
//------------------------------------------------------------------------------
static hal_ret_t
hal_sig_init (void)
{
    struct sigaction    act;

    memset(&act, 0, sizeof(act));
    act.sa_sigaction = hal_sig_handler;
    act.sa_flags = SA_SIGINFO;
    sigaction(SIGHUP, &act, NULL);
    sigaction(SIGQUIT, &act, NULL);
    sigaction(SIGINT, &act, NULL);
    sigaction(SIGUSR1, &act, NULL);
    sigaction(SIGSEGV, &act, NULL);
    sigaction(SIGCHLD, &act, NULL);
    sigaction(SIGURG, &act, NULL);
    sigaction(SIGUSR2, &act, NULL);
    sigaction(SIGTERM, &act, NULL);

    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
//  spawn and setup all the HAL threads - both config and packet loop threads
//------------------------------------------------------------------------------
static hal_ret_t
hal_thread_init (void)
{
    uint32_t              tid, core_id;
    int                   rv, thread_prio;
    char                  thread_name[16];
    struct sched_param    sched_param = { 0 };
    pthread_attr_t        attr;
    cpu_set_t             cpus;

    // spawn data core threads and pin them to their cores
    thread_prio = sched_get_priority_max(SCHED_FIFO);
    assert(thread_prio >= 0);
    for (tid = HAL_THREAD_ID_FTE_MIN, core_id = 1;
         FALSE && tid <= HAL_THREAD_ID_FTE_MAX;       // TODO: fix the env !!
         tid++, core_id++) {
        HAL_TRACE_DEBUG("Spawning FTE thread {}", tid);
        snprintf(thread_name, sizeof(thread_name), "fte-core-%u", core_id);
        g_hal_threads[tid] =
            thread::factory(static_cast<const char *>(thread_name), tid,
                            core_id, fte_pkt_loop,
                            thread_prio, SCHED_FIFO, false);
        HAL_ABORT(g_hal_threads[tid] != NULL);
        g_hal_threads[tid]->start(g_hal_threads[tid]);
    }

    // spawn periodic thread that does background tasks
    g_hal_threads[HAL_THREAD_ID_PERIODIC] =
        thread::factory(std::string("periodic-thread").c_str(),
                        HAL_THREAD_ID_PERIODIC,
                        HAL_CONTROL_CORE_ID,
                        hal::periodic::periodic_thread_start,
                        thread_prio - 1, SCHED_RR, true);
    HAL_ABORT(g_hal_threads[HAL_THREAD_ID_PERIODIC] != NULL);
    g_hal_threads[HAL_THREAD_ID_PERIODIC]->start(g_hal_threads[HAL_THREAD_ID_PERIODIC]);

    // make the current thread, main hal config thread (also a real-time thread)
    rv = pthread_attr_init(&attr);
    if (rv != 0) {
        HAL_TRACE_ERR("pthread_attr_init failure, err : {}", rv);
        return HAL_RET_ERR;
    }

    // set core affinity
    CPU_ZERO(&cpus);
    CPU_SET(HAL_CONTROL_CORE_ID, &cpus);
    rv = pthread_attr_setaffinity_np(&attr, sizeof(cpu_set_t), &cpus);
    if (rv != 0) {
        HAL_TRACE_ERR("pthread_attr_setaffinity_np failure, err : {}", rv);
        return HAL_RET_ERR;
    }

    if (gl_super_user) {
        HAL_TRACE_DEBUG("Started by root, switching to real-time scheduling");
        sched_param.sched_priority = sched_get_priority_max(SCHED_RR);
        rv = sched_setscheduler(0, SCHED_RR, &sched_param);
        if (rv != 0) {
            HAL_TRACE_ERR("sched_setscheduler failure, err : {}", rv);
            return HAL_RET_ERR;
        }
    }

    // create a thread object for this main thread
    g_hal_threads[HAL_THREAD_ID_CFG] =
        thread::factory(std::string("cfg-thread").c_str(),
                        HAL_THREAD_ID_CFG,
                        HAL_CONTROL_CORE_ID,
                        thread::dummy_entry_func,
                        sched_param.sched_priority, SCHED_RR, true);
    g_hal_threads[HAL_THREAD_ID_CFG]->set_ctxt(g_hal_threads[HAL_THREAD_ID_CFG]);
    t_curr_thread = g_hal_threads[HAL_THREAD_ID_CFG];
    g_hal_threads[HAL_THREAD_ID_CFG]->set_pthread_id(pthread_self());
    g_hal_threads[HAL_THREAD_ID_CFG]->set_running(true);

    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// wait for all the HAL threads to be terminated and any other background
// activities
//------------------------------------------------------------------------------
hal_ret_t
hal_wait (void)
{
    int         rv;
    uint32_t    tid;

    for (tid = HAL_THREAD_ID_PERIODIC; tid < HAL_THREAD_ID_MAX; tid++) {
        if (g_hal_threads[tid]) {
            rv = pthread_join(g_hal_threads[tid]->pthread_id(), NULL);
            if (rv != 0) {
                HAL_TRACE_ERR("pthread_join failure, thread {}, err : {}",
                              g_hal_threads[tid]->name(), rv);
                return HAL_RET_ERR;
            }
        }
    }
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// parse HAL configuration
//------------------------------------------------------------------------------
hal_ret_t
hal_parse_cfg (const char *cfgfile, hal_cfg_t *hal_cfg)
{
    ptree             pt;
    std::string       sparam;

    if (!cfgfile || !hal_cfg) {
        return HAL_RET_INVALID_ARG;
    }

    std::ifstream json_cfg(cfgfile);
    read_json(json_cfg, pt);
    try {
		std::string mode = pt.get<std::string>("mode");
        if (mode == "sim") {
            hal_cfg->sim = true;
        } else {
            hal_cfg->sim = false;
        }

        sparam = pt.get<std::string>("asic.name");
        strncpy(hal_cfg->asic_name, sparam.c_str(), HAL_MAX_NAME_STR);
        hal_cfg->grpc_port = pt.get<std::string>("sw.grpc_port");
        if (getenv("HAL_GRPC_PORT")) {
            hal_cfg->grpc_port = getenv("HAL_GRPC_PORT");
            HAL_TRACE_DEBUG("Overriding GRPC Port to : {}", hal_cfg->grpc_port);
        }
        sparam = pt.get<std::string>("sw.feature_set");
        strncpy(hal_cfg->feature_set, sparam.c_str(), HAL_MAX_NAME_STR);
    } catch (std::exception const& e) {
        std::cerr << e.what() << std::endl;
        return HAL_RET_INVALID_ARG;
    }
    return HAL_RET_OK;
}

//------------------------------------------------------------------------------
// init function for HAL
//------------------------------------------------------------------------------
hal_ret_t
hal_init (hal_cfg_t *hal_cfg)
{
    char    *user;

    HAL_TRACE_DEBUG("Initializing HAL ...");

    // check to see if HAL is running with root permissions
    user = getenv("USER");
    if (user && !strcmp(user, "root")) {
        gl_super_user = true;
    }

    // install signal handlers
    hal_sig_init();

    // do memory related initialization
    HAL_ABORT(hal_mem_init() == HAL_RET_OK);

    // init fte and hal plugins
    hal::init_plugins();

    // spawn all necessary PI threads
    HAL_ABORT(hal_thread_init() == HAL_RET_OK);
    HAL_TRACE_DEBUG("Spawned all HAL threads");

    // do platform dependent init
    HAL_ABORT(hal::pd::hal_pd_init(hal_cfg) == HAL_RET_OK);
    HAL_TRACE_DEBUG("Platform initialization done");

    // TODO_CLEANUP: this doesn't belong here, why is this outside
    // hal_state ??? how it this special compared to other global state ??
    g_lif_manager = new LIFManager();
   
    // do rdma init
    HAL_ABORT(rdma_hal_init() == HAL_RET_OK);

    hal_proxy_svc_init();
    
    return HAL_RET_OK;
}

}    // namespace hal
