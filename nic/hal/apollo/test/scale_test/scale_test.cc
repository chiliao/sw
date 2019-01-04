// Copyright (c) 2018 Pensando Systems, Inc.

#include <stdio.h>
#include <getopt.h>
#include <gtest/gtest.h>
#include "boost/foreach.hpp"
#include "boost/optional.hpp"
#include "boost/property_tree/ptree.hpp"
#include "boost/property_tree/json_parser.hpp"
#include "nic/sdk/include/sdk/eth.hpp"
#include "nic/sdk/include/sdk/ip.hpp"
#include "nic/hal/apollo/test/oci_test_base.hpp"
#include "nic/hal/apollo/include/api/oci_batch.hpp"
#include "nic/hal/apollo/include/api/oci_switchport.hpp"
#include "nic/hal/apollo/include/api/oci_tep.hpp"
#include "nic/hal/apollo/include/api/oci_vcn.hpp"
#include "nic/hal/apollo/include/api/oci_subnet.hpp"

using std::string;
char    *g_cfg_file = NULL;
namespace pt = boost::property_tree;

class scale_test : public oci_test_base {
protected:
    scale_test() {}
    virtual ~scale_test() {}
    /**< called immediately after the constructor before each test */
    virtual void SetUp() {}
    /**< called immediately after each test before the destructor */
    virtual void TearDown() {}
    /**< called at the beginning of all test cases in this class */
    static void SetUpTestCase() {
        /**< call base class function */
        oci_test_base::SetUpTestCase(false);
    }
};

static void
create_subnets(uint32_t vcn_id, uint32_t num_subnets, ip_prefix_t *vcn_pfx) {
    sdk_ret_t       rv;
    oci_subnet_t    oci_subnet;

    for (uint32_t i = 1; i <= num_subnets; i++) {
        memset(&oci_subnet, 0, sizeof(oci_subnet));
        oci_subnet.key.vcn_id = vcn_id;
        oci_subnet.key.id = i;
        oci_subnet.pfx = *vcn_pfx;
        oci_subnet.pfx.addr.addr.v4_addr =
            (oci_subnet.pfx.addr.addr.v4_addr & 0xFFFF0000) | ((i - 1) << 8);
        oci_subnet.pfx.len = 24;
        oci_subnet.vr_ip.af = IP_AF_IPV4;
        oci_subnet.vr_ip.addr.v4_addr = oci_subnet.pfx.addr.addr.v4_addr | 0x1;
        MAC_UINT64_TO_ADDR(oci_subnet.vr_mac,
                           (uint64_t)oci_subnet.vr_ip.addr.v4_addr);
        rv = oci_subnet_create(&oci_subnet);
        ASSERT_TRUE(rv == SDK_RET_OK);
    }
}

static void
create_vcns(uint32_t num_vcns, ip_prefix_t *ip_pfx, uint32_t num_subnets) {
    sdk_ret_t    rv;
    oci_vcn_t    oci_vcn;

    for (uint32_t i = 1; i <= num_vcns; i++) {
        memset(&oci_vcn, 0, sizeof(oci_vcn));
        oci_vcn.type = OCI_VCN_TYPE_TENANT;
        oci_vcn.key.id = i;
        oci_vcn.pfx = *ip_pfx;
        oci_vcn.pfx.addr.addr.v4_addr =
            oci_vcn.pfx.addr.addr.v4_addr | ((i - 1) << 16);
        rv = oci_vcn_create(&oci_vcn);
        ASSERT_TRUE(rv == SDK_RET_OK);
        for (uint32_t j = 1; j <= num_subnets; j++) {
            create_subnets(i, j, ip_pfx);
        }
    }
}

static void
create_teps(uint32_t num_teps, ip_prefix_t *ip_pfx) {
    sdk_ret_t    rv;
    oci_tep_t    oci_tep;

    // leave the 1st IP in this prefix for MyTEP
    for (uint32_t i = 1; i <= num_teps; i++) {
        memset(&oci_tep, 0, sizeof(oci_tep));
        oci_tep.key.ip_addr = ip_pfx->addr.addr.v4_addr + 1 + i;
        oci_tep.type = OCI_ENCAP_TYPE_IPINIP_GRE;
        rv = oci_tep_create(&oci_tep);
        ASSERT_TRUE(rv == SDK_RET_OK);
    }
}

static void
create_switchport_cfg(ipv4_addr_t ipaddr, uint64_t macaddr, ipv4_addr_t gwip) {
    sdk_ret_t           rv;
    oci_switchport_t    sw_port = { 0 };

    sw_port.switch_ip_addr = ipaddr;
    MAC_UINT64_TO_ADDR(sw_port.switch_mac_addr, macaddr);
    sw_port.gateway_ip_addr = gwip;
    rv = oci_switchport_create(&sw_port);
    ASSERT_TRUE(rv == SDK_RET_OK);
}

static void
create_objects(void) {
    pt::ptree      json_pt;
    uint32_t       count, num_subnets;
    ip_prefix_t    ippfx;
    string         pfxstr;

    // parse the config and create objects
    std::ifstream json_cfg(g_cfg_file);
    read_json(json_cfg, json_pt);
    try {
        BOOST_FOREACH(pt::ptree::value_type& obj, json_pt.get_child("objects")) {
            std::string kind = obj.second.get<std::string>("kind");
            if (kind == "switchport") {
                struct in_addr    ipaddr, gwip;
                uint64_t          macaddr;

                macaddr = std::stol(obj.second.get<std::string>("mac-addr"));
                inet_aton(obj.second.get<std::string>("ip-addr").c_str(),
                          &ipaddr);
                inet_aton(obj.second.get<std::string>("gw-ip-addr").c_str(),
                          &gwip);
                create_switchport_cfg(ntohl(ipaddr.s_addr), macaddr,
                                      ntohl(gwip.s_addr));
            } else if (kind == "tep") {
                count = std::stol(obj.second.get<std::string>("count"));
                pfxstr = obj.second.get<std::string>("prefix");
                ASSERT_TRUE(str2ipv4pfx((char *)pfxstr.c_str(), &ippfx) == 0);
                create_teps(count, &ippfx);
            } else if (kind == "vcn") {
                count = std::stol(obj.second.get<std::string>("count"));
                pfxstr = obj.second.get<std::string>("prefix");
                ASSERT_TRUE(str2ipv4pfx((char *)pfxstr.c_str(), &ippfx) == 0);
                num_subnets = std::stol(obj.second.get<std::string>("subnets"));
                create_vcns(count, &ippfx, num_subnets);
            } else if (kind == "vnic") {
                count = std::stol(obj.second.get<std::string>("count"));
            } else if (kind == "mapping") {
            }
        }
    } catch (std::exception const& e) {
        std::cerr << e.what() << std::endl;
        exit(1);
    }
}

TEST_F(scale_test, scale_test_create) {
    sdk_ret_t             rv;
    oci_batch_params_t    batch_params = { 0 };

    batch_params.epoch = 1;
    rv = oci_batch_start(&batch_params);
    ASSERT_TRUE(rv == SDK_RET_OK);
    create_objects();
    rv = oci_batch_commit();
    ASSERT_TRUE(rv == SDK_RET_OK);
}

// print help message showing usage of HAL
static void inline
print_usage (char **argv)
{
    fprintf(stdout, "Usage : %s -c|--config <cfg.json>\n", argv[0]);
}

int main(int argc, char **argv) {
    int               oc;
    struct option longopts[] = {
       { "config",    required_argument, NULL, 'c' },
       { "help",      no_argument,       NULL, 'h' },
       { 0,           0,                 0,     0 }
    };

    // parse CLI options
    while ((oc = getopt_long(argc, argv, ":hc:W;", longopts, NULL)) != -1) {
        switch (oc) {
        case 'c':
            g_cfg_file = optarg;
            if (!g_cfg_file) {
                fprintf(stderr, "config file is not specified\n");
                print_usage(argv);
                exit(1);
            }
            break;

        default:
            // ignore all other options
            break;
        }
    }

    // make sure cfg file exists
    if (access(g_cfg_file, R_OK) < 0) {
        fprintf(stderr, "Config file %s doesn't exist or not accessible\n",
                g_cfg_file);
        exit(1);
    }

    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
