// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

#ifndef __UPGRADE_APP_RESP_HDLR_H__
#define __UPGRADE_APP_RESP_HDLR_H__

#include "nic/delphi/sdk/delphi_sdk.hpp"
#include "gen/proto/upgrade.delphi.hpp"
#include "upgrade_handler.hpp"

namespace upgrade {

using namespace std;

//UpgAppRespHdlr class is used to create the object for 
class UpgAppRespHdlr {
    string appName_;
    delphi::SdkPtr sdk_;
public:
    UpgAppRespHdlr(delphi::SdkPtr sk, string name) {
        appName_ = name;
        sdk_ = sk;
    }

    //CreateUpgAppResp creates the object UpgAppResp
    delphi::error CreateUpgAppResp(void);
    //DeleteUpgAppResp deletes the object UpgAppResp
    delphi::error DeleteUpgAppResp(void);

    UpgStateRespType GetUpgAppRespNextPass(UpgReqStateType reqType, UpgType upgType);
    UpgStateRespType GetUpgAppRespNextFail(UpgReqStateType reqType, UpgType upgType);
    UpgStateRespType GetUpgAppRespNext(UpgReqStateType reqType, bool isReqSuccess, UpgType upgType);

    //CreateOrUpdateUpgAppResp creates the response for upgrade_manager
    delphi::error UpdateUpgAppResp(UpgStateRespType type, HdlrResp appHdlrResp, UpgType upgType);

    //CanInvokeHandler is used during OnMountComplete to know if we can InvokeAppHdlr
    bool CanInvokeHandler(UpgReqStateType reqType, UpgType upgType);

    //findUpgAppResp returns the UpgAppResp object for this application
    delphi::objects::UpgAppRespPtr findUpgAppResp(string name);

    //GetUpgStateReqPtr will return the pointer for UpgStateReq object
    delphi::objects::UpgStateReqPtr GetUpgStateReqPtr(void) {
        return delphi::objects::UpgStateReq::FindObject(sdk_);
    }

    string UpgAppRespValToStr(UpgStateRespType type, UpgType upgType);
};
typedef std::shared_ptr<UpgAppRespHdlr> UpgAppRespHdlrPtr;

} // namespace upgrade

#endif // __UPGRADE_APP_RESP_HDLR_H__
