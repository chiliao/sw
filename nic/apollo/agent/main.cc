//------------------------------------------------------------------------------
// Copyright (c) 2019 Pensando Systems, Inc.
//------------------------------------------------------------------------------

#include <unistd.h>
#include <getopt.h>
#include <limits.h>
#include <string>
#include "nic/apollo/agent/svc/batch.hpp"
#include "nic/apollo/agent/svc/device.hpp"
#include "nic/apollo/agent/svc/vpc.hpp"
#include "nic/apollo/agent/svc/subnet.hpp"
#include "nic/apollo/agent/svc/tunnel.hpp"
#include "nic/apollo/agent/svc/route.hpp"
#include "nic/apollo/agent/svc/vnic.hpp"
#include "nic/apollo/agent/svc/mapping.hpp"
#include "nic/apollo/agent/svc/port.hpp"
#include "nic/apollo/agent/init.hpp"
#include "nic/apollo/agent/trace.hpp"
#include "nic/apollo/test/flow_test/flow_test.hpp"

using std::string;
using grpc::Server;
using grpc::ServerBuilder;
using grpc::ServerContext;
using grpc::Status;
flow_test *g_flow_test_obj;
std::string g_grpc_server_addr;
#define GRPC_API_PORT    9999

static void
svc_reg (void)
{
    ServerBuilder     *server_builder;
    BatchSvcImpl      batch_svc;
    DeviceSvcImpl     device_svc;
    VPCSvcImpl        vpc_svc;
    SubnetSvcImpl     subnet_svc;
    TunnelSvcImpl     tunnel_svc;
    RouteSvcImpl      route_svc;
    VnicSvcImpl       vnic_svc;
    MappingSvcImpl    mapping_svc;
    PortSvcImpl       port_svc;

    // do gRPC initialization
    grpc_init();
    g_grpc_server_addr =
        std::string("0.0.0.0:") + std::to_string(GRPC_API_PORT);
    server_builder = new ServerBuilder();
    server_builder->SetMaxReceiveMessageSize(INT_MAX);
    server_builder->SetMaxSendMessageSize(INT_MAX);
    server_builder->AddListeningPort(g_grpc_server_addr,
                                     grpc::InsecureServerCredentials());

    // register for all the services
    server_builder->RegisterService(&batch_svc);
    server_builder->RegisterService(&device_svc);
    server_builder->RegisterService(&vpc_svc);
    server_builder->RegisterService(&subnet_svc);
    server_builder->RegisterService(&tunnel_svc);
    server_builder->RegisterService(&route_svc);
    server_builder->RegisterService(&vnic_svc);
    server_builder->RegisterService(&mapping_svc);
    server_builder->RegisterService(&port_svc);

    PDS_TRACE_INFO("gRPC server listening on ... {}",
                   g_grpc_server_addr.c_str());
    core::trace_logger()->flush();
    std::unique_ptr<Server> server(server_builder->BuildAndStart());
    server->Wait();
}

static void inline
print_usage (char **argv)
{
    fprintf(stdout, "Usage : %s -c|--config <cfg.json>\n", argv[0]);
}

int
main (int argc, char **argv)
{
    int          oc;
    string       cfg_path, cfg_file, profile, file;
    sdk_ret_t    ret;

    struct option longopts[] = {
       { "config",    required_argument, NULL, 'c' },
       { "profile",   required_argument, NULL, 'p' },
       { "help",      no_argument,       NULL, 'h' },
       { 0,           0,                 0,     0 }
    };

    // parse CLI options
    while ((oc = getopt_long(argc, argv, ":hc:p:W;", longopts, NULL)) != -1) {
        switch (oc) {
        case 'c':
            if (optarg) {
                cfg_file = std::string(optarg);
            } else {
                fprintf(stderr, "config file is not specified\n");
                print_usage(argv);
                exit(1);
            }
            break;

        case 'p':
            if (optarg) {
                profile = std::string(optarg);
            } else {
                fprintf(stderr, "profile is not specified\n");
                print_usage(argv);
                exit(1);
            }
            fprintf(stdout, "Got profile: %s\n", profile.c_str());
            break;

        case 'h':
            print_usage(argv);
            exit(0);
            break;

        case ':':
            fprintf(stderr, "%s: option -%c requires an argument\n",
                    argv[0], optopt);
            print_usage(argv);
            exit(1);
            break;

        case '?':
        default:
            fprintf(stderr, "%s: option -%c is invalid, quitting ...\n",
                    argv[0], optopt);
            print_usage(argv);
            exit(1);
            break;
        }
    }

    // form the full path to the config directory
    cfg_path = std::string(std::getenv("CONFIG_PATH"));
    if (cfg_path.empty()) {
        cfg_path = std::string("./");
    } else {
        cfg_path += "/";
    }

    // make sure the cfg file exists
    file = cfg_path + "apollo/" + std::string(cfg_file);
    if (access(file.c_str(), R_OK) < 0) {
        fprintf(stderr, "Config file %s doesn't exist or not accessible\n",
                file.c_str());
        exit(1);
    }

    // initialize the agent
    if ((ret = core::agent_init(cfg_file, profile)) != SDK_RET_OK) {
        fprintf(stderr, "Agent initialization failed, err %u", ret);
    }

    g_flow_test_obj = new flow_test();
    
    // register for all gRPC services
    svc_reg();

    // wait forver
    printf("Initialization done ...");
    while (1);

    return 0;
}
