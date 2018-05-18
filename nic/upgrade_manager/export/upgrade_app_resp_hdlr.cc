// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

#include <stdio.h>
#include <iostream>

#include "upgrade_app_resp_hdlr.hpp"

namespace upgrade {

using namespace std;

delphi::objects::UpgAppRespPtr UpgAppRespHdlr::findUpgAppResp(string name) {
    delphi::objects::UpgAppRespPtr req = make_shared<delphi::objects::UpgAppResp>();
    req->set_key(name);

    // find the object
    delphi::BaseObjectPtr obj = sdk_->FindObject(req);

    return static_pointer_cast<delphi::objects::UpgAppResp>(obj);
}

delphi::error UpgAppRespHdlr::CreateUpgAppResp(delphi::objects::UpgStateReqPtr ptr) {
    LogInfo("UpgAppRespHdlr::CreateUpgAppResp called");
    auto upgAppResp = this->findUpgAppResp(this->appName_);
    if (upgAppResp == NULL) {
        upgAppResp = make_shared<delphi::objects::UpgAppResp>();
        upgAppResp->set_key(this->appName_);
        if (upgAppResp == NULL)
            return delphi::error("application unable to create response object");
    }
    this->sdk_->SetObject(upgAppResp);
    return delphi::error::OK();
}

delphi::error UpgAppRespHdlr::UpdateUpgAppResp(UpgRespStateType type) {
    LogInfo("UpgAppRespHdlr::UpdateUpgAppResp called for {} with type {}", this->appName_, type);
    auto upgAppResp = this->findUpgAppResp(this->appName_);
    if (upgAppResp == NULL) {
        LogInfo("UpgAppRespHdlr::UpdateUpgAppResp returning error for {}", this->appName_);
        return delphi::error("application unable to find response object");
    }
    LogInfo("Setting upgAppRespVal to {}", type);
    upgAppResp->set_upgapprespval(type);
    this->sdk_->SetObject(upgAppResp);
    return delphi::error::OK();
}

UpgRespStateType
UpgAppRespHdlr::GetUpgAppRespNextPass(UpgReqStateType reqType) {
    LogInfo("UpgAppRespHdlr::GetUpgAppRespNextPass got called for reqType {}", reqType);
    switch (reqType) {
        case UpgReqRcvd:
            return UpgReqRcvdPass;
        case PreUpgState:
            return PreUpgStatePass;
        case PostBinRestart:
            return PostBinRestartPass;
        case ProcessesQuiesced:
            return ProcessesQuiescedPass;
        case DataplaneDowntimeStart:
            return DataplaneDowntimeStartPass;
        case Cleanup:
            return CleanupPass;
        case UpgSuccess:
            return UpgSuccessPass;
        case UpgFailed:
            return UpgFailedPass;
        case InvalidUpgState:
            return InvalidUpgStatePass;
        default:
            return InvalidUpgStatePass;
     }
}

UpgRespStateType
UpgAppRespHdlr::GetUpgAppRespNextFail(UpgReqStateType reqType) {
    switch (reqType) {
        case UpgReqRcvd:
            return UpgReqRcvdFail;
        case PreUpgState:
            return PreUpgStateFail;
        case PostBinRestart:
            return PostBinRestartFail;
        case ProcessesQuiesced:
            return ProcessesQuiescedFail;
        case DataplaneDowntimeStart:
            return DataplaneDowntimeStartFail;
        case Cleanup:
            return CleanupFail;
        case UpgSuccess:
            return UpgSuccessFail;
        case UpgFailed:
            return UpgFailedFail;
        case InvalidUpgState:
            return InvalidUpgStateFail;
        default:
            return InvalidUpgStateFail;
     }
}

UpgRespStateType
UpgAppRespHdlr::GetUpgAppRespNext(UpgReqStateType reqType, bool isReqSuccess) {
    if (isReqSuccess) {
        return GetUpgAppRespNextPass(reqType);
    } else {
        return GetUpgAppRespNextFail(reqType);
    }
}

} // namespace upgrade
