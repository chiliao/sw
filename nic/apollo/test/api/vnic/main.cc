
//
// {C} Copyright 2019 Pensando Systems Inc. All rights reserved
//
//----------------------------------------------------------------------------
///
/// \file
/// This file contains the all vnic test cases
///
//----------------------------------------------------------------------------

#include "nic/apollo/test/api/utils/device.hpp"
#include "nic/apollo/test/api/utils/subnet.hpp"
#include "nic/apollo/test/api/utils/vnic.hpp"
#include "nic/apollo/test/api/utils/vpc.hpp"
#include "nic/apollo/test/api/utils/policy.hpp"
#include "nic/apollo/test/api/utils/workflow.hpp"

namespace test {
namespace api {

//----------------------------------------------------------------------------
// VNIC test class
//----------------------------------------------------------------------------

class vnic_test : public pds_test_base {
protected:
    vnic_test() {}
    virtual ~vnic_test() {}
    virtual void SetUp() {}
    virtual void TearDown() {}
    static void SetUpTestCase() {
        if (!agent_mode())
            pds_test_base::SetUpTestCase(g_tc_params);
        pds_batch_ctxt_t bctxt = batch_start();
        sample_device_setup(bctxt);
        sample_vpc_setup(bctxt, PDS_VPC_TYPE_TENANT);
        sample_policy_setup(bctxt);
        sample_subnet_setup(bctxt);
        batch_commit(bctxt);
    }
    static void TearDownTestCase() {
        pds_batch_ctxt_t bctxt = batch_start();
        sample_subnet_teardown(bctxt);
        sample_policy_teardown(bctxt);
        sample_vpc_teardown(bctxt, PDS_VPC_TYPE_TENANT);
        sample_device_teardown(bctxt);
        batch_commit(bctxt);
        if (!agent_mode())
            pds_test_base::TearDownTestCase();
    }
};

//----------------------------------------------------------------------------
// VNIC test cases implementation
//----------------------------------------------------------------------------

/// \defgroup VNIC Vnic Tests
/// @{

/// \brief VNIC WF_B1
/// \ref WF_B1
TEST_F(vnic_test, vnic_workflow_b1) {
    vnic_feeder feeder;

    feeder.init(int2pdsobjkey(1), int2pdsobjkey(1), 1, k_feeder_mac,
                PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                true, true, 0, 0);
    workflow_b1<vnic_feeder>(feeder);
}

/// \brief VNIC WF_B2
/// \ref WF_B2
TEST_F(vnic_test, vnic_workflow_b2) {
    if (!apulu()) return;

    vnic_feeder feeder1, feeder1A;

    feeder1.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder1A.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN, FALSE, TRUE,0, 0);
    workflow_b2<vnic_feeder>(feeder1, feeder1A);
}

/// \brief VNIC WF_1
/// \ref WF_1
TEST_F(vnic_test, vnic_workflow_1) {
    if (artemis()) return;

    vnic_feeder feeder;

    feeder.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                true, true, 0, 0);
    workflow_1<vnic_feeder>(feeder);
}

/// \brief VNIC WF_2
/// \ref WF_2
TEST_F(vnic_test, vnic_workflow_2) {
    if (artemis()) return;

    vnic_feeder feeder;

    feeder.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                true, true, 0, 0);
    workflow_2<vnic_feeder>(feeder);
}

/// \brief VNIC WF_3
/// \ref WF_3
TEST_F(vnic_test, vnic_workflow_3) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder2, feeder3;

    feeder1.init(int2pdsobjkey(10), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder2.init(int2pdsobjkey(40), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder3.init(int2pdsobjkey(70), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    workflow_3<vnic_feeder>(feeder1, feeder2, feeder3);
}

/// \brief VNIC WF_4
/// \ref WF_4
TEST_F(vnic_test, vnic_workflow_4) {
    if (artemis()) return;

    vnic_feeder feeder;

    feeder.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                true, true, 0, 0);
    workflow_4<vnic_feeder>(feeder);
}

/// \brief VNIC WF_5
/// \ref WF_5
TEST_F(vnic_test, vnic_workflow_5) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder2, feeder3;

    feeder1.init(int2pdsobjkey(10), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder2.init(int2pdsobjkey(40), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder3.init(int2pdsobjkey(70), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    workflow_5<vnic_feeder>(feeder1, feeder2, feeder3);
}

/// \brief VNIC WF_6
/// \ref WF_6
TEST_F(vnic_test, vnic_workflow_6) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder1A, feeder1B;

    feeder1.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder1A.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                  true, true, 0, 0);
    feeder1B.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  FALSE, TRUE, 0, 0);
    workflow_6<vnic_feeder>(feeder1, feeder1A, feeder1B);
}

/// \brief VNIC WF_7
/// \ref WF_7
TEST_F(vnic_test, vnic_workflow_7) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder1A, feeder1B;

    feeder1.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder1A.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                  true, true, 0, 0);
    feeder1B.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  FALSE, TRUE, 0, 0);
    workflow_7<vnic_feeder>(feeder1, feeder1A, feeder1B);
}

/// \brief VNIC WF_8
/// \ref WF_8
TEST_F(vnic_test, vnic_workflow_8) {
    if (!apulu()) return;
    vnic_feeder feeder1, feeder1A, feeder1B;

    feeder1.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder1A.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  FALSE, TRUE, 0, 0);
    feeder1B.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                  FALSE, TRUE, 0, 0);
    workflow_8<vnic_feeder>(feeder1, feeder1A, feeder1B);
}

/// \brief VNIC WF_9
/// \ref WF_9
TEST_F(vnic_test, vnic_workflow_9) {
    if (!apulu()) return;

    vnic_feeder feeder1, feeder1A;

    feeder1.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder1A.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  true, true, 0, 0);
    workflow_9<vnic_feeder>(feeder1, feeder1A);
}

/// \brief VNIC WF_10
/// \ref WF_10
TEST_F(vnic_test, vnic_workflow_10) {
    if (!apulu()) return;

    vnic_feeder feeder1, feeder2, feeder3, feeder4;
    vnic_feeder feeder2A, feeder3A;

    feeder1.init(int2pdsobjkey(10), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder2.init(int2pdsobjkey(40), int2pdsobjkey(1), 20, 0x202020000000,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder2A.init(int2pdsobjkey(40), int2pdsobjkey(1), 20, 0x202020000000,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  FALSE, TRUE, 0, 0);
    feeder3.init(int2pdsobjkey(70), int2pdsobjkey(1), 20, 0x303030000000,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder3A.init(int2pdsobjkey(70), int2pdsobjkey(1), 20, 0x303030000000,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  FALSE, TRUE, 0, 0);
    feeder4.init(int2pdsobjkey(100), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    workflow_10<vnic_feeder>(feeder1, feeder2, feeder2A, feeder3,
                             feeder3A, feeder4);
}

/// \brief VNIC WF_N_1
/// \ref WF_N_1
TEST_F(vnic_test, vnic_workflow_neg_1) {
    if (artemis()) return;

    vnic_feeder feeder;

    feeder.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                true, true, 0, 0);
    workflow_neg_1<vnic_feeder>(feeder);
}

/// \brief VNIC WF_N_2
/// \ref WF_N_2
TEST_F(vnic_test, DISABLED_vnic_workflow_neg_2) {
    if (artemis()) return;

    vnic_feeder feeder;

    feeder.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic+1,
                k_feeder_mac, PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                true, true, 0, 0);
    workflow_neg_2<vnic_feeder>(feeder);
}

/// \brief VNIC WF_N_3
/// \ref WF_N_3
TEST_F(vnic_test, vnic_workflow_neg_3) {
    if (artemis()) return;

    vnic_feeder feeder;

    feeder.init(int2pdsobjkey(1), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                true, true, 0, 0);
    workflow_neg_3<vnic_feeder>(feeder);
}

/// \brief VNIC WF_N_4
/// \ref WF_N_4
TEST_F(vnic_test, vnic_workflow_neg_4) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder2;

    feeder1.init(int2pdsobjkey(10), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder2.init(int2pdsobjkey(40), int2pdsobjkey(1), 20, 0x202020000000,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    workflow_neg_4<vnic_feeder>(feeder1, feeder2);
}

/// \brief VNIC WF_N_5
/// \ref WF_N_5
TEST_F(vnic_test, vnic_workflow_neg_5) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder1A;

    feeder1.init(int2pdsobjkey(10), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder1A.init(int2pdsobjkey(10), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  FALSE, TRUE, 0, 0);
    workflow_neg_5<vnic_feeder>(feeder1, feeder1A);
}

/// \brief VNIC WF_N_6
/// \ref WF_N_6
TEST_F(vnic_test, vnic_workflow_neg_6) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder1A;

    feeder1.init(int2pdsobjkey(10), int2pdsobjkey(1), k_max_vnic, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder1A.init(int2pdsobjkey(10), int2pdsobjkey(1), k_max_vnic+1,
                  k_feeder_mac, PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  FALSE, TRUE, 0, 0);
    workflow_neg_6<vnic_feeder>(feeder1, feeder1A);
}

/// \brief VNIC WF_N_7
/// \ref WF_N_7
TEST_F(vnic_test, vnic_workflow_neg_7) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder1A, feeder2;

    feeder1.init(int2pdsobjkey(10), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder1A.init(int2pdsobjkey(10), int2pdsobjkey(1), 20, k_feeder_mac,
                  PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                  FALSE, TRUE, 0, 0);
    feeder2.init(int2pdsobjkey(40), int2pdsobjkey(1), 20, 0x202020000000,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    workflow_neg_7<vnic_feeder>(feeder1, feeder1A, feeder2);
}

/// \brief VNIC WF_N_8
/// \ref WF_N_8
TEST_F(vnic_test, vnic_workflow_neg_8) {
    if (artemis()) return;

    vnic_feeder feeder1, feeder2;

    feeder1.init(int2pdsobjkey(10), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_MPLSoUDP,
                 true, true, 0, 0);
    feeder2.init(int2pdsobjkey(40), int2pdsobjkey(1), 20, k_feeder_mac,
                 PDS_ENCAP_TYPE_DOT1Q, PDS_ENCAP_TYPE_VXLAN,
                 FALSE, TRUE, 0, 0);
    workflow_neg_8<vnic_feeder>(feeder1, feeder2);
}

/// @}

}    // namespace api
}    // namespace test

//----------------------------------------------------------------------------
// Entry point
//----------------------------------------------------------------------------

/// @private
int
main (int argc, char **argv)
{
    return api_test_program_run(argc, argv);
}
