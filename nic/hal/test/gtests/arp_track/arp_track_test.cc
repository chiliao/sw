//-----------------------------------------------------------------------------
// {C} Copyright 2020 Pensando Systems Inc. All rights reserved
//-----------------------------------------------------------------------------
#include <assert.h>
#include <gtest/gtest.h>

#include "nic/fte/test/fte_base_test.hpp"
#include "nic/hal/iris/include/hal_state.hpp"
#include "nic/hal/pd/iris/hal_state_pd.hpp"
#include "nic/hal/plugins/cfg/telemetry/telemetry.hpp"
#include "nic/hal/plugins/cfg/nw/interface.hpp"
#include "nic/hal/test/hal_calls/hal_calls.hpp"
#include "nic/hal/test/utils/hal_base_test.hpp"

#include "gen/proto/telemetry.grpc.pb.h"

using namespace hal;
using namespace hal::pd;
using namespace sdk::lib;


using telemetry::MirrorSessionSpec;
using telemetry::MirrorSessionStatus;
using telemetry::MirrorSessionResponse;
using telemetry::MirrorSessionDeleteRequest;
using telemetry::MirrorSessionDeleteResponse;
using telemetry::MirrorSessionGetRequest;
using telemetry::MirrorSessionGetResponse;
using telemetry::RuleAction;
using telemetry::ERSpanSpec;

class arp_track_test : public hal_base_test {
protected:
  arp_track_test() {
  }

  virtual ~arp_track_test() {
  }

  // will be called immediately after the constructor before each test
  virtual void SetUp() {
  }

  // will be called immediately after each test before the destructor
  virtual void TearDown() {
  }

  // Will be called at the beginning of all test cases in this class
  static void SetUpTestCase() {
    // hal_base_test::SetUpTestCase();
      hal_base_test::SetUpTestCaseGrpc();
  }
};

TEST_F (arp_track_test, mirror_order1) {
    uint32_t       test_id = 1;
    uint32_t       uplinkif_id1 = UPLINK_IF_ID_OFFSET + test_id,
                   uplinkif_id2 = uplinkif_id1 + 1, uplinkif_id3 = uplinkif_id1 + 2;
    uint32_t       enic_wl1 = 256;
    uint32_t       tnnl_if_id1 = 100, tnnl_if_id2 = 101;
    uint32_t       up_port1 = PORT_NUM_1, up_port2 = PORT_NUM_2, up_port3 = PORT_NUM_3;
    uint32_t       vrf_id_hp1 = 65;
    uint32_t       l2seg_id_hp1 = 101, l2seg_id_hp2 = 102;
    uint32_t       src_ip = 0x0a000001;
    uint32_t       ips[2] = {0x0a000002, 0x0a000003};
    uint32_t       sid1 = 1, sid2 = 2, cid1 = 1, cid2 = 2;
    uint32_t       up_ifid[2], ifid_count = 0;
    if_t           *up_if = NULL;
    uint32_t       host_lifid1 = 68;
    uint32_t       wl_encap1 = 100;
    gtest_mirror_t g_mirror;

    // Create uplinks
    ASSERT_EQ(create_uplink(uplinkif_id1, up_port1), HAL_RET_OK);
    ASSERT_EQ(create_uplink(uplinkif_id2, up_port2), HAL_RET_OK);
    ASSERT_EQ(create_uplink(uplinkif_id3, up_port3, 0, true), HAL_RET_OK);

    
    // Create lif
    ASSERT_EQ(create_lif(host_lifid1, uplinkif_id1, types::LIF_TYPE_HOST,
                         "eth0"), HAL_RET_OK);

    up_if = find_if_by_id(uplinkif_id1);

    // Create Inband VRF
    ASSERT_EQ(create_vrf(vrf_id_hp1, types::VRF_TYPE_CUSTOMER, 0), HAL_RET_OK);

    // Create Inband L2seg
    // Create Classic L2segs in each Classic VRF with respective uplink
     up_ifid[0] = uplinkif_id1;
     up_ifid[1] = uplinkif_id2;
     ifid_count = 2;
     ASSERT_EQ(update_l2seg(vrf_id_hp1, l2seg_id_hp1, 8191, up_ifid, ifid_count,
                            l2segment::MulticastFwdPolicy::MULTICAST_FWD_POLICY_FLOOD,
                            l2segment::BroadcastFwdPolicy::BROADCAST_FWD_POLICY_FLOOD,
                            false, true), HAL_RET_OK);

    // Create Collector EP
    ASSERT_EQ(create_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 2), HAL_RET_OK);

    // Create Tunnel If
    ASSERT_EQ(create_tunnel(tnnl_if_id1, vrf_id_hp1, src_ip, ips[0]), HAL_RET_OK);
    ASSERT_EQ(create_tunnel(tnnl_if_id2, vrf_id_hp1, src_ip, ips[1]), HAL_RET_OK);

    g_hal_state->set_inb_bond_active_uplink(HAL_HANDLE_INVALID);

    // Create Mirror session
    ASSERT_EQ(create_mirror(sid1, vrf_id_hp1, src_ip, ips[0], true), HAL_RET_OK);
    ASSERT_EQ(create_mirror(sid2, vrf_id_hp1, src_ip, ips[1]), HAL_RET_OK);

    g_mirror.session_id = sid2;
    g_mirror.span_id = 1000;
    g_mirror.vrf_id = vrf_id_hp1;
    g_mirror.sip = src_ip;
    g_mirror.dip = ips[1];
    g_mirror.vlan_strip_en = true;
    g_mirror.erspan_type = 2;
    ASSERT_EQ(update_mirror(&g_mirror, false), HAL_RET_OK);

    // Create flow monitor
    gtest_flow_mon_t fmon;
    fmon.vrf_id = vrf_id_hp1;
    fmon.rule_id = 1;
    fmon.sess_count = 1;
    fmon.sess_id[0] = sid1;
    ASSERT_EQ(update_flow_monitor(&fmon, true), HAL_RET_OK);

    fmon.sess_count = 2;
    fmon.sess_id[0] = sid1;
    fmon.sess_id[1] = sid2;
    ASSERT_EQ(update_flow_monitor(&fmon, false), HAL_RET_OK);

    g_hal_state->set_inb_bond_active_uplink(up_if->hal_handle); 
    hal::hal_if_inb_bond_active_changed(false);

    // Create collector
    ASSERT_EQ(create_collector(cid1, vrf_id_hp1, l2seg_id_hp1, src_ip, ips[0]), HAL_RET_OK);
    ASSERT_EQ(create_collector(cid2, vrf_id_hp1, l2seg_id_hp1, src_ip, ips[1]), HAL_RET_OK);

    gtest_collector_t coll;
    coll.cid = cid2;
    coll.vrf = vrf_id_hp1;
    coll.l2seg = l2seg_id_hp1;
    coll.src_ip = src_ip;
    coll.dst_ip = ips[1];
    coll.dest_port = 200;
    coll.export_invl = 20;
    ASSERT_EQ(update_collector(&coll, false), HAL_RET_OK);
     
    // Remove IP from EP
    ASSERT_EQ(update_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 1), HAL_RET_OK);

    // Add IP to EP
    ASSERT_EQ(update_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 2), HAL_RET_OK);

    // Delete EP
    ASSERT_EQ(delete_ep(vrf_id_hp1, l2seg_id_hp1, 0x000010002001), HAL_RET_OK);

    // Create Collector EP
    ASSERT_EQ(create_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 2), HAL_RET_OK);

    // Delete tunnel
    ASSERT_EQ(delete_interface(tnnl_if_id1), HAL_RET_OK);

    // Create Tunnel If
    ASSERT_EQ(create_tunnel(tnnl_if_id1, vrf_id_hp1, src_ip, ips[0]), HAL_RET_OK);

    
    // Create customer L2seg with wire encap as 8192 
    up_ifid[0] = uplinkif_id1;
    up_ifid[1] = uplinkif_id2;
    ifid_count = 2;
    ASSERT_EQ(update_l2seg(vrf_id_hp1, l2seg_id_hp2, 102, up_ifid, ifid_count,
                           l2segment::MulticastFwdPolicy::MULTICAST_FWD_POLICY_FLOOD,
                           l2segment::BroadcastFwdPolicy::BROADCAST_FWD_POLICY_FLOOD,
                           false, true), HAL_RET_OK);
    gtest_enic_t enic1 = {0};
    enic1.if_id = enic_wl1;
    enic1.lif_id = host_lifid1;
    enic1.type = intf::IF_ENIC_TYPE_USEG;
    enic1.l2seg_id = l2seg_id_hp2;
    enic1.encap = wl_encap1;
    enic1.native_l2seg_id = 0;
    enic1.tx_mirr[0] = sid1;
    enic1.tx_mirr[1] = sid2;
    enic1.tx_mirr_count = 2;
    enic1.rx_mirr[0] = sid1;
    enic1.rx_mirr[1] = sid2;
    enic1.rx_mirr_count = 2;

    // Create Enic
    ASSERT_EQ(update_enic(&enic1, GTEST_CREATE), HAL_RET_OK);

    // Update rx mirr for enic
    enic1.rx_mirr_count = 1;
    ASSERT_EQ(update_enic(&enic1, GTEST_UPDATE), HAL_RET_OK);

    enic1.tx_mirr_count = 1;
    ASSERT_EQ(update_enic(&enic1, GTEST_UPDATE), HAL_RET_OK);

    enic1.tx_mirr_count = 2;
    enic1.rx_mirr_count = 2;
    ASSERT_EQ(update_enic(&enic1, GTEST_UPDATE), HAL_RET_OK);

    enic1.tx_mirr_count = 0;
    enic1.rx_mirr_count = 0;
    ASSERT_EQ(update_enic(&enic1, GTEST_UPDATE), HAL_RET_OK);

    // Create WL EP
    // ASSERT_EQ(create_ep(vrf_id_hp1, l2seg_id_hp1, enic_wl1, 0x000020002001), HAL_RET_OK);

    // For halctl cli
    // sleep(1000);
}

TEST_F (arp_track_test, mirror_order2) {
    uint32_t       test_id = 1;
    uint32_t       uplinkif_id1 = UPLINK_IF_ID_OFFSET + test_id,
                   uplinkif_id2 = uplinkif_id1 + 1, uplinkif_id3 = uplinkif_id1 + 2;
    uint32_t       tnnl_if_id1 = 100, tnnl_if_id2 = 101;
    uint32_t       up_port1 = PORT_NUM_1, up_port2 = PORT_NUM_2, up_port3 = PORT_NUM_3;
    uint32_t       vrf_id_hp1 = 65;
    uint32_t       l2seg_id_hp1 = 101;
    uint32_t       src_ip = 0x0a000001;
    uint32_t       ips[2] = {0x0a000002, 0x0a000003};
    uint32_t       sid1 = 1, sid2 = 2;
    uint32_t       up_ifid[2], ifid_count = 0;

    // Create uplinks
    ASSERT_EQ(create_uplink(uplinkif_id1, up_port1), HAL_RET_OK);
    ASSERT_EQ(create_uplink(uplinkif_id2, up_port2), HAL_RET_OK);
    ASSERT_EQ(create_uplink(uplinkif_id3, up_port3, 0, true), HAL_RET_OK);

    // Create Inband VRF
    ASSERT_EQ(create_vrf(vrf_id_hp1, types::VRF_TYPE_CUSTOMER, 0), HAL_RET_OK);

    // Create Inband L2seg
    // Create Classic L2segs in each Classic VRF with respective uplink
     up_ifid[0] = uplinkif_id1;
     up_ifid[1] = uplinkif_id2;
     ifid_count = 2;
     ASSERT_EQ(update_l2seg(vrf_id_hp1, l2seg_id_hp1, 8191, up_ifid, ifid_count,
                            l2segment::MulticastFwdPolicy::MULTICAST_FWD_POLICY_FLOOD,
                            l2segment::BroadcastFwdPolicy::BROADCAST_FWD_POLICY_FLOOD,
                            false, true), HAL_RET_OK);

#if 0
    // Create Collector EP
    ASSERT_EQ(create_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 2), HAL_RET_OK);

    // Create Tunnel If
    ASSERT_EQ(create_tunnel(tnnl_if_id1, vrf_id_hp1, src_ip, ips[0]), HAL_RET_OK);
    ASSERT_EQ(create_tunnel(tnnl_if_id2, vrf_id_hp1, src_ip, ips[1]), HAL_RET_OK);
#endif

    // Create Mirror session
    ASSERT_EQ(create_mirror(sid1, vrf_id_hp1, src_ip, ips[0]), HAL_RET_OK);
    ASSERT_EQ(create_mirror(sid2, vrf_id_hp1, src_ip, ips[1]), HAL_RET_OK);

    // Create Tunnel If
    ASSERT_EQ(create_tunnel(tnnl_if_id1, vrf_id_hp1, src_ip, ips[0]), HAL_RET_OK);
    ASSERT_EQ(create_tunnel(tnnl_if_id2, vrf_id_hp1, src_ip, ips[1]), HAL_RET_OK);

    // Create Collector EP
    ASSERT_EQ(create_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 2), HAL_RET_OK);

#if 0
    uint32_t       cid1 = 1, cid2 = 2;

    // Create collector
    ASSERT_EQ(create_collector(cid1, vrf_id_hp1, l2seg_id_hp1, src_ip, ips[0]), HAL_RET_OK);
    ASSERT_EQ(create_collector(cid2, vrf_id_hp1, l2seg_id_hp1, src_ip, ips[1]), HAL_RET_OK);

    // Remove IP from EP
    ASSERT_EQ(update_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 1), HAL_RET_OK);

    // Add IP to EP
    ASSERT_EQ(update_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 2), HAL_RET_OK);

    // Delete EP
    ASSERT_EQ(delete_ep(vrf_id_hp1, l2seg_id_hp1, 0x000010002001), HAL_RET_OK);

    // Create Collector EP
    ASSERT_EQ(create_ep(vrf_id_hp1, l2seg_id_hp1, uplinkif_id1, 0x000010002001, ips, 2), HAL_RET_OK);

    // Delete tunnel
    ASSERT_EQ(delete_interface(tnnl_if_id1), HAL_RET_OK);

    // Create Tunnel If
    ASSERT_EQ(create_tunnel(tnnl_if_id1, vrf_id_hp1, src_ip, ips[0]), HAL_RET_OK);
#endif
}

int main (int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
