// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

#include "gtest/gtest.h"
#include "upgrade.hpp"
#include "nic/upgrade_manager/proto/upgrade.delphi_utest.hpp"

namespace {

using namespace std;
using namespace upgrade;


// --------------------- Reactor test upgrade ------------------------

UPGRADEREQUESTSPEC_REACTOR_TEST(UpgradeReactorTest, UpgradeMgr);

TEST_F(UpgradeReactorTest, BasicTest) {
    // create an upgrade request spec object
    delphi::objects::UpgradeRequestSpecPtr req = make_shared<delphi::objects::UpgradeRequestSpec>();
    req->mutable_key()->set_id(1);
    req->set_upgreqaction(upgrade::UpgAbort);
    sdk_->QueueUpdate(req);

    // verify spec object is in db
    ASSERT_EQ_EVENTUALLY(sdk_->ListKind("UpgradeRequestSpec").size(), 1) << "Upgrade request spec object was not created";

    // verify corresponding status object got created
    ASSERT_EQ_EVENTUALLY(sdk_->ListKind("UpgReqStatus").size(), 1) << "Upgrade request status object was not created";

    // change the admin status on spec
    req->set_upgreqaction(upgrade::UpgStart);
    sdk_->QueueUpdate(req);

    // verify spec object in db has changed
    ASSERT_EQ_EVENTUALLY(delphi::objects::UpgradeRequestSpec::FindObject(sdk_, req)->upgreqaction(),
                        upgrade::UpgStart) << "Upgrade request spec object has wrong oper state";

    // verify status object has correct operstate too
    delphi::objects::UpgReqStatusPtr upgReqStatusKey = make_shared<delphi::objects::UpgReqStatus>();
    upgReqStatusKey->set_key(1);
    ASSERT_EQ_EVENTUALLY(delphi::objects::UpgReqStatus::FindObject(sdk_, upgReqStatusKey)->upgreqstatus(),
                        upgrade::RcvdUpgReq) << "Upgrade Request status object has wrong oper state";
}


// --------------------- Service test upgrade ------------------------
// init the delphi utest
DELPHI_SERVICE_TEST(UpgradeTest, UpgradeService);

TEST_F(UpgradeTest, BasicTest) {
    usleep(1000);

    // create an interface spec object
    delphi::objects::UpgradeRequestSpecPtr req = make_shared<delphi::objects::UpgradeRequestSpec>();
    req->mutable_key()->set_id(1);
    req->set_upgreqaction(upgrade::UpgAbort);
    sdk_->QueueUpdate(req);

    usleep(1000 * 100);

    // verify spec object is in the db
    ASSERT_EQ(sdk_->ListKind("UpgradeRequestSpec").size(), 1) << "Upgrade Request spec object was not created";

    // verify corresponding status object got created
    ASSERT_EQ(sdk_->ListKind("UpgReqStatus").size(), 1) << "UpgReq status object was not created";

    // change the action status on spec and verify status gets changed too
    auto upgReqSpecList = sdk_->ListKind("UpgradeRequestSpec");
    for (vector<delphi::BaseObjectPtr>::iterator iter=upgReqSpecList.begin(); iter!=upgReqSpecList.end(); ++iter) {
        delphi::objects::UpgradeRequestSpecPtr upgReqSpec = static_pointer_cast<delphi::objects::UpgradeRequestSpec>(*iter);
        upgReqSpec->set_upgreqaction(upgrade::UpgStart);
        sdk_->QueueUpdate(upgReqSpec);
    }

    usleep(1000 * 100);

    auto upgReqStatusList = sdk_->ListKind("UpgReqStatus");
    for (vector<delphi::BaseObjectPtr>::iterator iter=upgReqStatusList.begin(); iter!=upgReqStatusList.end(); ++iter) {
        delphi::objects::UpgReqStatusPtr upgReqStatus = static_pointer_cast<delphi::objects::UpgReqStatus>(*iter);
        ASSERT_EQ(upgReqStatus->upgreqstatus(), upgrade::RcvdUpgReq) << "Upgrade Request status object has wrong oper state";
    }
}
} // namespace

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
