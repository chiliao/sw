//----------------------------------------------------------------------------
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//----------------------------------------------------------------------------
#include <iostream>
#include <memory>
#include <string>
#include <grpc++/grpc++.h>
#include <gen/proto/bgp.pb.h>
#include <gen/proto/bgp.grpc.pb.h>
#include <chrono>
#include <stdlib.h>
#include <unistd.h>
#include <arpa/inet.h>
#include "nic/metaswitch/stubs/mgmt/pds_ms_mgmt_utils.hpp"
#include "nic/metaswitch/stubs/test/common/test_config.hpp"
#include "nic/apollo/api/include/pds_if.hpp"
#include "nic/apollo/api/include/pds_device.hpp"
#include "nic/sdk/include/sdk/if.hpp"
#include "gen/proto/device.grpc.pb.h"
#include "gen/proto/interface.grpc.pb.h"
#include <gen/proto/vpc.grpc.pb.h>
#include <gen/proto/subnet.grpc.pb.h>
#include <gen/proto/bgp.grpc.pb.h>
#include <gen/proto/evpn.grpc.pb.h>
#include <gen/proto/cp_route.grpc.pb.h>
#include <gen/proto/cp_interface.grpc.pb.h>
#include "nic/apollo/agent/svc/specs.hpp"

using grpc::Channel;
using grpc::ClientContext;
using grpc::Status;
using namespace pds_ms_test;
using namespace std;
using namespace pds;

static test_config_t g_test_conf_;
static unique_ptr<pds::DeviceSvc::Stub> g_device_stub_;
static unique_ptr<pds::IfSvc::Stub>     g_if_stub_;
static unique_ptr<pds::BGPSvc::Stub>    g_bgp_stub_;
static unique_ptr<pds::EvpnSvc::Stub>    g_evpn_stub_;
static unique_ptr<pds::SubnetSvc::Stub> g_subnet_stub_;
static unique_ptr<pds::VPCSvc::Stub>    g_vpc_stub_;
static unique_ptr<pds::CPInterfaceSvc::Stub>    g_intf_stub_;
static unique_ptr<pds::CPRouteSvc::Stub>        g_route_stub_;

static void create_device_proto_grpc () {
    ClientContext   context;
    DeviceRequest   request;
    DeviceResponse  response;
    Status          ret_status;
    pds_device_spec_t device_spec = {0};

    device_spec.overlay_routing_en = TRUE;
    device_spec.device_ip_addr.af  = types::IP_AF_INET;
    device_spec.device_ip_addr.addr.v4_addr = (g_test_conf_.local_lo_ip_addr);
    device_spec.gateway_ip_addr.af          = types::IP_AF_INET;
    device_spec.gateway_ip_addr.addr.v4_addr = g_test_conf_.local_gwip_addr;

    pds_device_api_spec_to_proto (request.mutable_request(), &device_spec);

    printf ("Pushing Device Proto...\n");
    ret_status = g_device_stub_->DeviceCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_l3_intf_proto_grpc () {
    InterfaceRequest    request;
    InterfaceResponse   response;
    ClientContext       context;
    Status              ret_status;
    pds_if_spec_t       pds_if = {0};

    request.mutable_batchctxt()->set_batchcookie(1);

    pds_if.key.id = g_test_conf_.eth_if_index;
    pds_if.type = PDS_IF_TYPE_L3;
    pds_if.admin_state = PDS_IF_STATE_UP;
    pds_if.l3_if_info.vpc.id = 1;
    pds_if.l3_if_info.ip_prefix.addr.af = IP_AF_IPV4;
    pds_if.l3_if_info.ip_prefix.addr.addr.v4_addr = g_test_conf_.local_ip_addr;
    pds_if.l3_if_info.ip_prefix.len = 16;
    pds_if.l3_if_info.eth_ifindex = ETH_IFINDEX(ETH_IF_DEFAULT_SLOT, 1,
                                                ETH_IF_DEFAULT_CHILD_PORT);

    pds_if_api_spec_to_proto (request.add_request(), &pds_if);

    printf ("Pushing Interface Proto...\n");
    ret_status = g_if_stub_->InterfaceCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_bgp_global_proto_grpc () {
    BGPRequest      request;
    BGPResponse     response;
    ClientContext   context;
    Status          ret_status;

    auto proto_spec = request.add_request();
    proto_spec->set_vrfid (PDS_MS_BGP_RM_ENT_INDEX);
    proto_spec->set_localasn (g_test_conf_.local_asn);
    proto_spec->set_routerid(ntohl(g_test_conf_.local_lo_ip_addr));

    printf ("Pushing BGP Global proto...\n");
    ret_status = g_bgp_stub_->BGPGlobalSpecCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_evpn_evi_proto_grpc () {
    EvpnEviRequest  request;
    EvpnResponse    response;
    ClientContext   context;
    Status          ret_status;

    auto proto_spec = request.add_request ();
    proto_spec->set_eviid (1);
    if (g_test_conf_.manual_rd) {
        proto_spec->set_autord (pds::EVPN_CFG_MANUAL);
        proto_spec->set_rd((const char *)g_test_conf_.rd, 8);
    } else {
        proto_spec->set_autord (pds::EVPN_CFG_AUTO);
    }
    if (g_test_conf_.manual_rt) {
        proto_spec->set_autort (pds::EVPN_CFG_MANUAL);
    } else {
        proto_spec->set_autort (pds::EVPN_CFG_AUTO);
    }
    proto_spec->set_rttype (pds::EVPN_RT_IMPORT_EXPORT);
    proto_spec->set_encap (pds::EVPN_ENCAP_VXLAN);

    printf ("Pushing EVPN Evi proto...\n");
    ret_status = g_evpn_stub_->EvpnEviSpecCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_evpn_evi_rt_proto_grpc () {
    EvpnEviRtRequest    request;
    EvpnResponse        response;
    ClientContext       context;
    Status              ret_status;

    auto proto_spec = request.add_request ();
    proto_spec->set_eviid (1);
    proto_spec->set_rt((const char *)g_test_conf_.rt, 8);
    proto_spec->set_rttype (pds::EVPN_RT_IMPORT_EXPORT);

    printf ("Pushing EVPN Evi RT proto...\n");
    ret_status = g_evpn_stub_->EvpnEviRtSpecCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_route_proto_grpc () {
    CPStaticRouteRequest  request;
    CPStaticRouteResponse response;
    ClientContext         context;
    Status                ret_status;

    auto proto_spec = request.add_request ();
    auto dest_addr  = proto_spec->mutable_destaddr();
    dest_addr->set_af (types::IP_AF_INET);
    dest_addr->set_v4addr (0);
    proto_spec->set_prefixlen (0);
    auto next_hop   = proto_spec->mutable_nexthopaddr();
    next_hop->set_af (types::IP_AF_INET);
    next_hop->set_v4addr (g_test_conf_.remote_ip_addr);
    proto_spec->set_adminstatus (ADMIN_UP);
    proto_spec->set_override (BOOL_TRUE);
    proto_spec->set_admindist (250);
    proto_spec->set_action (STRT_ACTION_FWD);

    printf ("Pushing Default (0/0) Static Route proto...\n");
    ret_status = g_route_stub_->CPStaticRouteSpecCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
    }
}

static void create_bgp_peer_proto_grpc (bool lo=false) {
    BGPPeerRequest  request;
    BGPResponse     response;
    ClientContext   context;
    Status          ret_status;

    auto proto_spec = request.add_request();
    auto peeraddr = proto_spec->mutable_peeraddr();
    peeraddr->set_af(types::IP_AF_INET);
    if (lo) {
        peeraddr->set_v4addr(g_test_conf_.remote_lo_ip_addr);
    } else {
        peeraddr->set_v4addr(g_test_conf_.remote_ip_addr);
    }
    proto_spec->set_vrfid(PDS_MS_BGP_RM_ENT_INDEX);
    proto_spec->set_adminen(pds::ADMIN_UP);
    proto_spec->set_peerport(0);
    auto localaddr = proto_spec->mutable_localaddr();
    localaddr->set_af(types::IP_AF_INET);
    if (lo) {
        localaddr->set_v4addr(g_test_conf_.local_lo_ip_addr);
    } else {
        localaddr->set_v4addr(0);
    }
    proto_spec->set_localport(0);
    proto_spec->set_ifid(0);
    proto_spec->set_remoteasn(g_test_conf_.remote_asn);
    proto_spec->set_localasn(g_test_conf_.local_asn);
    proto_spec->set_connectretry(5);
    proto_spec->set_sendcomm(pds::BOOL_TRUE);
    proto_spec->set_sendextcomm(pds::BOOL_TRUE);

    printf ("Pushing BGP %s Peer proto...\n", (lo) ? "Overlay" : "Underlay" );
    ret_status = g_bgp_stub_->BGPPeerSpecCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_bgp_peer_af_proto_grpc (bool lo=false) {
    BGPPeerAfRequest  request;
    BGPResponse     response;
    ClientContext   context;
    Status          ret_status;

    auto proto_spec = request.add_request();
    auto peeraddr = proto_spec->mutable_peeraddr();
    peeraddr->set_af(types::IP_AF_INET);

    if (lo) {
        peeraddr->set_v4addr(g_test_conf_.remote_lo_ip_addr);
    } else {
        peeraddr->set_v4addr(g_test_conf_.remote_ip_addr);
    }
    proto_spec->set_vrfid(PDS_MS_BGP_RM_ENT_INDEX);
    proto_spec->set_peerport(0);
    auto localaddr = proto_spec->mutable_localaddr();
    localaddr->set_af(types::IP_AF_INET);
    if (lo) {
        localaddr->set_v4addr(g_test_conf_.local_lo_ip_addr);
    } else {
        localaddr->set_v4addr(0);
    }

    proto_spec->set_localport(0);
    proto_spec->set_ifid(0);
    if (lo) {
        // Disable IP
        proto_spec->set_afi(pds::BGP_AFI_IPV4);
        proto_spec->set_safi(pds::BGP_SAFI_UNICAST);
    } else {
        // Disable EVPN
        proto_spec->set_afi(pds::BGP_AFI_L2VPN);
        proto_spec->set_safi(pds::BGP_SAFI_EVPN);
    }
    proto_spec->set_disable(pds::BOOL_TRUE);
    proto_spec->set_nhself(pds::BOOL_FALSE);
    proto_spec->set_defaultorig(pds::BOOL_FALSE);

    printf ("Pushing BGP %s Peer AF proto...\n", (lo) ? "Overlay" : "Underlay" );
    ret_status = g_bgp_stub_->BGPPeerAfCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_subnet_proto_grpc () {
    SubnetRequest   request;
    SubnetResponse  response;
    ClientContext   context;
    Status          ret_status;

    request.mutable_batchctxt()->set_batchcookie(1);

    auto proto_spec = request.add_request();
    proto_spec->set_id(1);
    proto_spec->set_vpcid(1);
    auto proto_encap = proto_spec->mutable_fabricencap();
    proto_encap->set_type(types::ENCAP_TYPE_VXLAN);
    proto_encap->mutable_value()->set_vnid(g_test_conf_.vni);
    proto_spec->set_hostifindex(g_test_conf_.lif_if_index);
    proto_spec->set_ipv4virtualrouterip(g_test_conf_.local_gwip_addr);
    proto_spec->set_virtualroutermac((uint64_t)0x001122334455);
    auto v4_prefix = proto_spec->mutable_v4prefix();
    v4_prefix->set_len(24);
    v4_prefix->set_addr (g_test_conf_.local_gwip_addr);

    printf ("Pushing Subnet proto...\n");
    ret_status = g_subnet_stub_->SubnetCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_vpc_proto_grpc () {
    VPCRequest      request;
    VPCResponse     response;
    ClientContext   context;
    Status          ret_status;

    request.mutable_batchctxt()->set_batchcookie(1);

    auto proto_spec = request.add_request();
    proto_spec->set_id(1);
    proto_spec->set_type(pds::VPC_TYPE_TENANT);
    auto proto_encap = proto_spec->mutable_fabricencap();
    proto_encap->set_type(types::ENCAP_TYPE_VXLAN);
    proto_encap->mutable_value()->set_vnid(g_test_conf_.vni);

    printf ("Pushing VPC proto...\n");
    ret_status = g_vpc_stub_->VPCCreate(&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_loopback_proto_grpc () {
    CPInterfaceRequest  request;
    CPInterfaceResponse   response;
    ClientContext         context;
    Status                ret_status;

    auto proto_spec = request.add_request();
    proto_spec->set_iftype (pds::CPIntfType::CP_IF_TYPE_LOOPBACK);
    proto_spec->set_ifid (10);

    printf ("Pushing Loopback Interface proto...\n");
    ret_status = g_intf_stub_->CPInterfaceSpecCreate (&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}

static void create_loopback_addr_proto_grpc () {
    CPInterfaceAddrRequest  request;
    CPInterfaceResponse     response;
    ClientContext           context;
    Status                  ret_status;

    auto proto_spec = request.add_request();
    proto_spec->set_iftype (pds::CPIntfType::CP_IF_TYPE_LOOPBACK);
    proto_spec->set_ifid (10);
    auto ipaddr = proto_spec->mutable_ipaddr();
    ipaddr->set_af (types::IP_AF_INET);
    ipaddr->set_v4addr (g_test_conf_.local_lo_ip_addr);
    proto_spec->set_prefixlen (32);

    printf ("Pushing Loopback Interface IP Address proto...0x%x\n", g_test_conf_.local_lo_ip_addr);
    ret_status = g_intf_stub_->CPInterfaceAddrSpecCreate (&context, request, &response);
    if (!ret_status.ok() || (response.apistatus() != types::API_STATUS_OK)) {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
        exit(1);
    }
}
static void get_peer_status_all() {
    BGPPeerRequest       request;
    BGPPeerSpecResponse  response;
    ClientContext        context;
    Status               ret_status;

    auto proto_spec = request.add_request();
    proto_spec->set_vrfid (PDS_MS_BGP_RM_ENT_INDEX);

    ret_status = g_bgp_stub_->BGPPeerSpecGetAll (&context, request, &response);
    if (ret_status.ok()) {
        printf ("No of BGP Peer Status Table Entries: %d\n", response.response_size());
        for (int i=0; i<response.response_size(); i++) {
            auto resp = response.response(i);
            printf (" Entry :: %d\n", i+1);
            printf (" ===========\n");
            printf ("  VRF Id               : %d\n", resp.vrfid());
            printf ("  Local Port           : %d\n",resp.localport());
            auto paddr = resp.localaddr().v4addr();
            struct in_addr ip_addr;
            ip_addr.s_addr = paddr;
            printf ("  Local Address        : %s\n", inet_ntoa(ip_addr));
            printf ("  Peer Port            : %d\n",resp.peerport());
            paddr = resp.peeraddr().v4addr();
            ip_addr.s_addr = paddr;
            printf ("  Peer Address         : %s\n", inet_ntoa(ip_addr));
            printf ("  If Id                : %d\n", resp.ifid());
            printf ("  Local ASN            : %d\n", resp.localasn());
            printf ("  Remote ASN           : %d\n", resp.remoteasn());
            printf ("  Status               : %d\n", resp.status());
            printf ("  Previous Status      : %d\n", resp.prevstatus());
            uint8_t ler[2];
            memcpy(ler, resp.lasterrorrcvd().c_str(), 2);
            printf  ("  last-error-rcvd::error      : %d\n", unsigned(ler[0]));;
            printf  ("  last-error-rcvd::suberror   : %d\n", unsigned(ler[1]));;
            uint8_t les[2];
            memcpy(les, resp.lasterrorsent().c_str(), 2);
            printf  ("  last-error-sent::error      : %d\n", unsigned(les[0]));;
            printf  ("  last-error-sent::suberror   : %d\n\n", unsigned(les[1]));;
        }
    } else {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
    }
}

static void get_evpn_mac_ip_all () {
    EvpnMacIpSpecRequest    request;
    EvpnMacIpSpecResponse   response;
    ClientContext           context;
    Status                  ret_status;

    auto proto_spec = request.add_request();
    proto_spec->set_eviid(1);
    ret_status = g_evpn_stub_->EvpnMacIpSpecGetAll (&context, request, &response);
    if (ret_status.ok()) {
        printf ("No of EVPN MAC IP Table Entries: %d\n", response.response_size());
        for (int i=0; i<response.response_size(); i++) {
            auto resp = response.response(i);
            printf (" Entry :: %d\n", i+1);
            printf (" ===========\n");
            printf ("  Source       : %s\n", (resp.source() == 2) ? "Remote" : "Local");
            printf ("  EVI Id       : %d\n", resp.eviid());
            uint8_t mac[6];
            memcpy(mac, resp.macaddress().c_str(),6);
            printf ("  MAC Address  : 0x%X 0x%X 0x%X 0x%X 0x%X 0x%X\n", 
                    mac[0], mac[1], mac[2], mac[3], mac[4], mac[5]);
            auto ipaddr = resp.ipaddress().v4addr();
            struct in_addr ip_addr;
            ip_addr.s_addr = ipaddr;
            printf ("  IP Address   : %s\n\n", inet_ntoa(ip_addr));
        }
    } else {
        printf("%s failed! ret_status=%d (%s) response.status=%d\n",
                __FUNCTION__, ret_status.error_code(), ret_status.error_message().c_str(),
                response.apistatus());
    }
}

int main(int argc, char** argv)
{
    // parse json config file
    if (parse_json_config(&g_test_conf_, 1) < 0) {
        cout << "Config file not found! Check CONFIG_PATH env var\n";
        exit(1);
    }

    std::shared_ptr<Channel> channel = grpc::CreateChannel("localhost:9999",
            grpc::InsecureChannelCredentials());
    g_device_stub_  = DeviceSvc::NewStub (channel);
    g_if_stub_      = IfSvc::NewStub (channel);
    g_bgp_stub_     = BGPSvc::NewStub (channel);
    g_evpn_stub_    = EvpnSvc::NewStub (channel);
    g_vpc_stub_     = VPCSvc::NewStub (channel);
    g_subnet_stub_  = SubnetSvc::NewStub (channel);
    g_route_stub_   = CPRouteSvc::NewStub (channel);
    g_intf_stub_    = CPInterfaceSvc::NewStub (channel);

    if (argc == 1) 
    {
        // Send protos to grpc server
        create_device_proto_grpc();
        create_l3_intf_proto_grpc();
        create_loopback_proto_grpc();
        create_loopback_addr_proto_grpc();
        create_route_proto_grpc();
        create_bgp_global_proto_grpc();
        create_bgp_peer_proto_grpc();
        create_bgp_peer_af_proto_grpc();
        sleep(40);
        create_bgp_peer_proto_grpc(true /* loopback */);
        create_bgp_peer_af_proto_grpc(true /* loopback */);
        create_vpc_proto_grpc();
        create_subnet_proto_grpc();
        create_evpn_evi_proto_grpc();
        if (g_test_conf_.manual_rt) {
            create_evpn_evi_rt_proto_grpc();
        }
        printf ("Testapp Config Init is successful!\n");
        return 0;
    } else if (argc == 2) {
        if (!strcmp(argv[1], "peer_status")) {
            get_peer_status_all();
            return 0;
        } else if (!strcmp (argv[1], "evpn_mac_ip")) {
            get_evpn_mac_ip_all();
            return 0;
        }
    }

    printf ("Invalid CLI Arguments!  Usage: \n"
            "no arguments : run test init config\n"
            "peer_status  : display Peer Status Table\n"
            "evpn_mac_ip  : display EVPN MAC IP Table\n");
    return 0;

}
