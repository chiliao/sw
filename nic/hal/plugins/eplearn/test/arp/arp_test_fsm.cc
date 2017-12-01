#include <arpa/inet.h>
#include <gtest/gtest.h>
#include "nic/include/twheel.hpp"

#include "nic/hal/plugins/eplearn/arp/arp_learn.hpp"
#include "nic/hal/plugins/eplearn/arp/arp_trans.hpp"
#include "nic/hal/test/utils/hal_test_utils.hpp"
#include "nic/hal/test/utils/hal_base_test.hpp"

#include "nic/hal/src/interface.hpp"
#include "nic/hal/src/endpoint.hpp"
#include "nic/hal/src/session.hpp"
#include "nic/hal/src/network.hpp"
#include "nic/gen/proto/hal/interface.pb.h"
#include "nic/gen/proto/hal/l2segment.pb.h"
#include "nic/gen/proto/hal/vrf.pb.h"
#include "nic/gen/proto/hal/endpoint.pb.h"
#include "nic/gen/proto/hal/nw.pb.h"

#include <tins/tins.h>

using namespace Tins;
using namespace hal;
using namespace hal::network;
using hal::utils::twheel;

/*
 * Note: For now timer wheel is supposed to be non-thread safe.
 * Reason: Update tick locks that slice and calls the callback function.
 * During the callback function, we try to add or delete timer which will
 * obviously wait indefinitely as there is not lock available.
 */

namespace hal {
namespace periodic {
extern twheel *g_twheel;
}
}  // namespace hal

vrf_t *dummy_ten;
#define MAX_ENDPOINTS 4
hal_handle_t ep_handles[MAX_ENDPOINTS];
string mac_addr_base = "12345";
#define GET_MAC_ADDR(_ep) ((unsigned char*)((mac_addr_base + std::to_string(_ep)).c_str()))

void fte_ctx_init(fte::ctx_t &ctx, hal::vrf_t *ten, hal::ep_t *ep,
        fte::cpu_rxhdr_t *cpu_rxhdr, uint8_t *pkt, size_t pkt_len,
        fte::flow_t iflow[], fte::flow_t rflow[]);

void arp_topo_setup()
{
   hal_ret_t                   ret;
   VrfSpec                  ten_spec;
   VrfResponse              ten_rsp;
   L2SegmentSpec               l2seg_spec;
   L2SegmentResponse           l2seg_rsp;
   InterfaceSpec               up_spec;
   InterfaceResponse           up_rsp;
   EndpointSpec                ep_spec, ep_spec1;
   EndpointResponse            ep_rsp;
   EndpointUpdateRequest       ep_req, ep_req1;
   NetworkSpec                 nw_spec, nw_spec1;
   NetworkResponse             nw_rsp, nw_rsp1;
   NetworkKeyHandle            *nkh = NULL;

   // Create vrf
   ten_spec.mutable_key_or_handle()->set_vrf_id(1);
   hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
   ret = hal::vrf_create(ten_spec, &ten_rsp);
   hal::hal_cfg_db_close();
   ASSERT_TRUE(ret == HAL_RET_OK);

   // Create network
   nw_spec.mutable_vrf_key_handle()->set_vrf_id(1);
   nw_spec.set_rmac(0x0000DEADBEEE);
   nw_spec.mutable_key_or_handle()->mutable_ip_prefix()->set_prefix_len(24);
   nw_spec.mutable_key_or_handle()->mutable_ip_prefix()->mutable_address()->set_ip_af(types::IP_AF_INET);
   nw_spec.mutable_key_or_handle()->mutable_ip_prefix()->mutable_address()->set_v4_addr(0x0a000000);
   nw_spec.mutable_vrf_key_handle()->set_vrf_id(1);
   hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
   ret = hal::network_create(nw_spec, &nw_rsp);
   hal::hal_cfg_db_close();
   ASSERT_TRUE(ret == HAL_RET_OK);
   uint64_t nw_hdl = nw_rsp.mutable_status()->nw_handle();

   nw_spec1.mutable_vrf_key_handle()->set_vrf_id(1);
   nw_spec1.set_rmac(0x0000DEADBEEF);
   nw_spec1.mutable_key_or_handle()->mutable_ip_prefix()->set_prefix_len(24);
   nw_spec1.mutable_key_or_handle()->mutable_ip_prefix()->mutable_address()->set_ip_af(types::IP_AF_INET);
   nw_spec1.mutable_key_or_handle()->mutable_ip_prefix()->mutable_address()->set_v4_addr(0x0b000000);
   nw_spec1.mutable_vrf_key_handle()->set_vrf_id(1);
   hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
   ret = hal::network_create(nw_spec1, &nw_rsp);
   hal::hal_cfg_db_close();
   ASSERT_TRUE(ret == HAL_RET_OK);
   uint64_t nw_hdl1 = nw_rsp.mutable_status()->nw_handle();

   // Create L2 Segment
   l2seg_spec.mutable_vrf_key_handle()->set_vrf_id(1);
   nkh = l2seg_spec.add_network_key_handle();
   nkh->set_nw_handle(nw_hdl);
   l2seg_spec.mutable_key_or_handle()->set_segment_id(1);
   l2seg_spec.mutable_wire_encap()->set_encap_type(types::ENCAP_TYPE_DOT1Q);
   l2seg_spec.mutable_wire_encap()->set_encap_value(11);
   hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
   ret = hal::l2segment_create(l2seg_spec, &l2seg_rsp);
   hal::hal_cfg_db_close();
   ASSERT_TRUE(ret == HAL_RET_OK);
   uint64_t l2seg_hdl = l2seg_rsp.mutable_l2segment_status()->l2segment_handle();

   l2seg_spec.mutable_vrf_key_handle()->set_vrf_id(1);
   nkh = l2seg_spec.add_network_key_handle();
   nkh->set_nw_handle(nw_hdl1);
   l2seg_spec.mutable_key_or_handle()->set_segment_id(2);
   l2seg_spec.mutable_wire_encap()->set_encap_type(types::ENCAP_TYPE_DOT1Q);
   l2seg_spec.mutable_wire_encap()->set_encap_value(12);
   hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
   ret = hal::l2segment_create(l2seg_spec, &l2seg_rsp);
   hal::hal_cfg_db_close();
   ASSERT_TRUE(ret == HAL_RET_OK);
   // uint64_t l2seg_hdl2 = l2seg_rsp.mutable_l2segment_status()->l2segment_handle();

   // Create an uplink
   up_spec.set_type(intf::IF_TYPE_UPLINK);
   up_spec.mutable_key_or_handle()->set_interface_id(1);
   up_spec.mutable_if_uplink_info()->set_port_num(1);
   hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
   ret = hal::interface_create(up_spec, &up_rsp);
   hal::hal_cfg_db_close();
   ASSERT_TRUE(ret == HAL_RET_OK);
   // ::google::protobuf::uint64 up_hdl = up_rsp.mutable_status()->if_handle();

   up_spec.set_type(intf::IF_TYPE_UPLINK);
   up_spec.mutable_key_or_handle()->set_interface_id(2);
   up_spec.mutable_if_uplink_info()->set_port_num(2);
   hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
   ret = hal::interface_create(up_spec, &up_rsp);
   hal::hal_cfg_db_close();
   ASSERT_TRUE(ret == HAL_RET_OK);
   ::google::protobuf::uint64 up_hdl2 = up_rsp.mutable_status()->if_handle();
   dummy_ten = vrf_lookup_by_id(1);
   ASSERT_TRUE(dummy_ten != NULL);


   ep_t *dummy_ep;
   for (int i = 0; i < MAX_ENDPOINTS; i++) {
       ep_spec.mutable_vrf_key_handle()->set_vrf_id(1);
       ep_spec.mutable_key_or_handle()->mutable_endpoint_key()->mutable_l2_key()->mutable_l2segment_key_handle()->set_l2segment_handle(l2seg_hdl);
       ep_spec.mutable_endpoint_attrs()->mutable_interface_key_handle()->set_if_handle(up_hdl2);
       ep_spec.mutable_key_or_handle()->mutable_endpoint_key()->mutable_l2_key()->set_mac_address(0x00000000ABCD + i);
       //ep_spec.mutable_endpoint_attrs()->add_ip_address();
       hal::hal_cfg_db_open(hal::CFG_OP_WRITE);
       ret = hal::endpoint_create(ep_spec, &ep_rsp);
       hal::hal_cfg_db_close();
       ASSERT_TRUE(ret == HAL_RET_OK);
       ep_handles[i] = ep_rsp.endpoint_status().endpoint_handle();
       dummy_ep = find_ep_by_handle(ep_handles[i]);
       strcpy((char*)(dummy_ep->l2_key.mac_addr),
               (mac_addr_base + std::to_string(i)).c_str());
   }

}

class arp_fsm_test : public hal_base_test {
   protected:
    arp_fsm_test() {
    }

    virtual ~arp_fsm_test() {}

    // will be called immediately after the constructor before each test
    virtual void SetUp() {

    }

    // will be called immediately after each test before the destructor
    virtual void TearDown() {}

    // Will be called at the beginning of all test cases in this class
    static void SetUpTestCase() {
        hal_base_test::SetUpTestCase();
        //hal_test_utils_slab_disable_delete();
        /* Override the timer so that we can control  */
        //hal::periodic::g_twheel = twheel::factory(10, 100, false);
        arp_topo_setup();
    }
};

EthernetII *get_default_arp_packet(ARP::Flags type,
                                   const char *sender_ip_address,
                                   unsigned char *sender_hw_addr,
                                   const char *target_ip_addr,
                                   unsigned char *target_hw_addr) {
    EthernetII *eth = new EthernetII();
    ARP *arp = new ARP();

    eth->inner_pdu(arp);

    // Retrieve a pointer to the stored TCP PDU
    arp = eth->find_pdu<ARP>();

    arp->opcode(type);

    HWAddress<6> hw_address(sender_hw_addr);
    arp->sender_hw_addr(hw_address);

    HWAddress<6> hw_address_1(target_hw_addr);
    arp->target_hw_addr(hw_address_1);

    arp->sender_ip_addr(sender_ip_address);
    arp->target_ip_addr(target_ip_addr);

    return eth;
}

hal_ret_t arp_packet_send(hal_handle_t ep_handle,
                     ARP::Flags type, const char *sender_ip_address,
                     unsigned char *sender_hw_addr, const char *target_ip_addr,
                     unsigned char *target_hw_addr = NULL) {
    EthernetII *eth =
        get_default_arp_packet(type, sender_ip_address, sender_hw_addr,
                               target_ip_addr, target_hw_addr);

    std::vector<uint8_t> buffer = eth->serialize();
    fte::ctx_t ctx;
    ep_t *dummy_ep = find_ep_by_handle(ep_handle);
    fte::cpu_rxhdr_t cpu_hdr;
    cpu_hdr.flags = 0;
    fte_ctx_init(ctx, dummy_ten,
            dummy_ep, &cpu_hdr, &buffer[0], buffer.size(), NULL, NULL);
    return arp_process_packet(ctx);
}

TEST_F(arp_fsm_test, arp_request_response) {
    arp_packet_send(ep_handles[0], ARP::Flags::REQUEST, "1.1.1.1",
                    GET_MAC_ADDR(0), "1.1.1.2");
    arp_trans_key_t key;
    ep_t *dummy_ep = find_ep_by_handle(ep_handles[0]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    arp_trans_t *entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);

    ip_addr_t ip = {0};
    inet_pton(AF_INET, "1.1.1.1", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
    arp_packet_send(ep_handles[1], ARP::Flags::REPLY, "1.1.1.2",
                    GET_MAC_ADDR(1), "1.1.1.1", GET_MAC_ADDR(0));
    dummy_ep = find_ep_by_handle(ep_handles[1]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(1), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));

    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    //ASSERT_EQ(arp_trans_t::ep_l3_entry_ht()->num_entries(), 2);
}

TEST_F(arp_fsm_test, arp_request_response_repeat) {
    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    ASSERT_EQ(g_hal_state->ep_l3_entry_ht()->num_entries(), 2);
    arp_packet_send(ep_handles[0], ARP::Flags::REQUEST, "1.1.1.1",
                    GET_MAC_ADDR(0), "1.1.1.2");
    arp_trans_key_t key;
    ep_t *dummy_ep = find_ep_by_handle(ep_handles[0]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    arp_trans_t *entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);

    ip_addr_t ip = {0};
    inet_pton(AF_INET, "1.1.1.1", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
    arp_packet_send(ep_handles[1], ARP::Flags::REPLY, "1.1.1.2",
                    GET_MAC_ADDR(1), "1.1.1.1", GET_MAC_ADDR(0));
    dummy_ep = find_ep_by_handle(ep_handles[1]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(1), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));

    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    //ASSERT_EQ(arp_trans_t::ep_l3_entry_ht()->num_entries(), 2);
}

TEST_F(arp_fsm_test, arp_spoofing) {

    hal_ret_t ret = arp_packet_send(ep_handles[0],
                    ARP::Flags::REQUEST, "1.1.1.10",
                    GET_MAC_ADDR(1), "1.1.1.2");

    ASSERT_TRUE(ret != HAL_RET_OK);
    ep_t *dummy_ep = find_ep_by_handle(ep_handles[0]);
    ip_addr_t ip = {0};
    inet_pton(AF_INET, "1.1.1.10", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));

    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
}

TEST_F(arp_fsm_test, arp_entry_timeout) {
    arp_packet_send(ep_handles[0], ARP::Flags::REQUEST, "1.1.1.1",
                    GET_MAC_ADDR(0), "1.1.1.2");
    arp_trans_key_t key;
    ep_t *dummy_ep = find_ep_by_handle(ep_handles[0]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    arp_trans_t *entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);

    auto timeout = entry->get_current_state_timeout();
    hal::periodic::g_twheel->tick(timeout + 100);
    sleep(3);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry == NULL);
    dummy_ep = find_ep_by_handle(ep_handles[1]);
    ip_addr_t ip = {0};
    inet_pton(AF_INET, "1.1.1.1", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
}


TEST_F(arp_fsm_test, arp_request_response_again) {
    arp_packet_send(ep_handles[0], ARP::Flags::REQUEST, "1.1.1.1",
                    GET_MAC_ADDR(0), "1.1.1.2");
    arp_trans_key_t key;
    ep_t *dummy_ep = find_ep_by_handle(ep_handles[0]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    arp_trans_t *entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);

    ip_addr_t ip = {0};
    inet_pton(AF_INET, "1.1.1.1", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
    arp_packet_send(ep_handles[1], ARP::Flags::REPLY, "1.1.1.2",
                    GET_MAC_ADDR(1), "1.1.1.1", GET_MAC_ADDR(0));
    dummy_ep = find_ep_by_handle(ep_handles[1]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(1), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));

    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    ASSERT_EQ(g_hal_state->ep_l3_entry_ht()->num_entries(), 2);

    delete entry;
}


TEST_F(arp_fsm_test, arp_request_response_replace_with_different_ep) {

    arp_packet_send(ep_handles[0], ARP::Flags::REQUEST, "1.1.1.1",
                    GET_MAC_ADDR(0), "1.1.1.2");
    arp_trans_key_t key;
    ep_t *dummy_ep = find_ep_by_handle(ep_handles[0]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    arp_trans_t *entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    arp_packet_send(ep_handles[1], ARP::Flags::REPLY, "1.1.1.2",
                    GET_MAC_ADDR(1), "1.1.1.1", GET_MAC_ADDR(0));
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(1), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));

    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    ASSERT_EQ(g_hal_state->ep_l3_entry_ht()->num_entries(), 2);

    //Lets replace both.
    arp_packet_send(ep_handles[2],
                    ARP::Flags::REQUEST, "1.1.1.1",
                    GET_MAC_ADDR(2), "1.1.1.2");
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(2), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);

    arp_packet_send(ep_handles[3], ARP::Flags::REPLY, "1.1.1.2",
                    GET_MAC_ADDR(3), "1.1.1.1", GET_MAC_ADDR(2));
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(3), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    dummy_ep = find_ep_by_handle(ep_handles[0]);
    ip_addr_t ip = {0};
    inet_pton(AF_INET, "1.1.1.1", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
    dummy_ep = find_ep_by_handle(ep_handles[1]);
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
}

TEST_F(arp_fsm_test, arp_request_response_replace_with_different_ip_same_ep) {

    arp_packet_send(ep_handles[0], ARP::Flags::REQUEST, "1.1.1.1",
                    GET_MAC_ADDR(0), "1.1.1.2");
    arp_trans_key_t key;
    ep_t *dummy_ep = find_ep_by_handle(ep_handles[0]);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    arp_trans_t *entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    arp_packet_send(ep_handles[1], ARP::Flags::REPLY, "1.1.1.2",
                    GET_MAC_ADDR(1), "1.1.1.1", GET_MAC_ADDR(0));
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(1), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));

    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    //ASSERT_EQ(arp_trans_t::ep_l3_entry_ht()->num_entries(), 2);

    //Lets replace both with different IP.
    arp_packet_send(ep_handles[0],
                    ARP::Flags::REQUEST, "1.1.1.3",
                    GET_MAC_ADDR(0), "1.1.1.4");
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);

    arp_packet_send(ep_handles[1], ARP::Flags::REPLY, "1.1.1.4",
                    GET_MAC_ADDR(1), "1.1.1.3", GET_MAC_ADDR(0));
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(1), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    ASSERT_EQ(arp_trans_t::arplearn_ip_entry_ht()->num_entries(), 2);

    dummy_ep = find_ep_by_handle(ep_handles[0]);
    ip_addr_t ip = {0};
    inet_pton(AF_INET, "1.1.1.3", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.1", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.4", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    dummy_ep = find_ep_by_handle(ep_handles[1]);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));

    //Lets replace back both with different IP.
    arp_packet_send(ep_handles[0],
                    ARP::Flags::REQUEST, "1.1.1.1",
                    GET_MAC_ADDR(0), "1.1.1.2");
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(0), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);

    arp_packet_send(ep_handles[1], ARP::Flags::REPLY, "1.1.1.2",
                    GET_MAC_ADDR(1), "1.1.1.1", GET_MAC_ADDR(0));
    arp_trans_t::init_arp_trans_key(GET_MAC_ADDR(1), dummy_ep, &key);
    entry = reinterpret_cast<arp_trans_t *>(
        arp_trans_t::arplearn_key_ht()->lookup(&key));
    ASSERT_TRUE(entry != NULL);
    ASSERT_EQ(arp_trans_t::arplearn_key_ht()->num_entries(), 2);
    ASSERT_EQ(arp_trans_t::arplearn_ip_entry_ht()->num_entries(), 2);

    dummy_ep = find_ep_by_handle(ep_handles[0]);
    ip = {0};
    inet_pton(AF_INET, "1.1.1.1", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.3", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.2", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    dummy_ep = find_ep_by_handle(ep_handles[1]);
    ASSERT_TRUE(ip_in_ep(&ip, dummy_ep, NULL));
    inet_pton(AF_INET, "1.1.1.4", &(ip.addr.v4_addr));
    ip.addr.v4_addr = ntohl(ip.addr.v4_addr);
    ASSERT_FALSE(ip_in_ep(&ip, dummy_ep, NULL));
}
