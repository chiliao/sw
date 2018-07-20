// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

#include <stdio.h>
#include <iostream>

#include "upgrade.hpp"
#include "upgrade_mgr.hpp"
#include "upgrade_app_resp_handlers.hpp"
#include "nic/upgrade_manager/include/c/upgrade_state_machine.hpp"
#include "nic/upgrade_manager/include/c/upgrade_metadata.hpp"
#include "nic/upgrade_manager/utils/upgrade_log.hpp"

namespace upgrade {

using namespace std;
UpgCtx ctx;

void UpgradeMgr::RegNewApp(string name) {
    if (appRegMap_[name] == false) {
        UPG_LOG_DEBUG("App not registered. Registering {} now.", name);
        appRegMap_[name] = true;
    } else {
        UPG_LOG_DEBUG("App {} already registered.", name);
    }
}

UpgReqStateType UpgradeMgr::GetNextState(void) {
    UpgReqStateType  reqType;
    auto reqStatus = findUpgStateReq(10);
    reqType = reqStatus->upgreqstate();
    if (GetAppRespFail() && (reqType != UpgStateFailed) && (reqType != UpgStateCleanup)) {
        UPG_LOG_DEBUG("Some application(s) responded with failure");
        return UpgStateFailed;
    }
    if (reqType == UpgStateSuccess) 
        upgPassed_ = true;
    return StateMachine[reqType].stateNext;
}

bool UpgradeMgr::IsRespTypeFail(UpgStateRespType type) {
    bool ret = false;
    switch (type) {
        case UpgStateCompatCheckRespFail:
        case UpgStateProcessQuiesceRespFail:
        case UpgStatePostBinRestartRespFail:
        case UpgStateDataplaneDowntimePhase1RespFail:
        case UpgStateDataplaneDowntimePhase2RespFail:
        case UpgStateDataplaneDowntimePhase3RespFail:
        case UpgStateDataplaneDowntimePhase4RespFail:
        case UpgStateCleanupRespFail:
        case UpgStateSuccessRespFail:
        case UpgStateFailedRespFail:
        case UpgStateAbortRespFail:
            ret = true;
        default:
            break;
    }
    return ret;
}

UpgStateRespType UpgradeMgr::GetFailRespType(UpgReqStateType type) {
    return StateMachine[type].stateFailResp;
}

UpgStateRespType UpgradeMgr::GetPassRespType(UpgReqStateType type) {
    return StateMachine[type].statePassResp;
}

bool UpgradeMgr::CanMoveStateMachine(void) {
    UpgStateRespType passType, failType;
    UpgReqStateType  reqType;
    bool ret = true;
    UPG_LOG_DEBUG("Checking if state machine can be moved forward");
    //Find UpgStateReq object
    auto reqStatus = findUpgStateReq(10);
    reqType = reqStatus->upgreqstate();
    passType = GetPassRespType(reqType);
    failType = GetFailRespType(reqType);
    UPG_LOG_DEBUG("reqType/passType/failType: {}/{}/{}", reqType, passType, failType);

    //check if all responses have come
    vector<delphi::objects::UpgAppRespPtr> upgAppRespList = delphi::objects::UpgAppResp::List(sdk_);
    if (upgAppRespList.size() != appRegMap_.size()) {
        ret = false;
        UPG_LOG_DEBUG("Number of responses from Applications {} is not same as the number of applications {}", upgAppRespList.size(), appRegMap_.size());
    } else {
        for (vector<delphi::objects::UpgAppRespPtr>::iterator appResp=upgAppRespList.begin(); appResp!=upgAppRespList.end(); ++appResp) {
            if (((*appResp)->upgapprespval() != passType) &&
                ((*appResp)->upgapprespval() != failType)){
                UPG_LOG_DEBUG("Application {} still processing {}", (*appResp)->key(), UpgReqStateTypeToStr(reqType));
                ret = false;
            } else if ((*appResp)->upgapprespval() == passType) {
                UPG_LOG_DEBUG("Got pass from application {}/{}", (*appResp)->key(), ((*appResp))->meta().ShortDebugString());
            } else {
                UPG_LOG_DEBUG("Got fail from application {}", (*appResp)->key());
            }
        }
    }
    if (ret) {
        UPG_LOG_DEBUG("Got pass/fail response from all applications. Can move state machine.");
    }
    return ret;
}

bool UpgradeMgr::InvokePreStateHandler(UpgReqStateType reqType) {
    UpgPreStateFunc preStFunc = StateMachine[reqType].preStateFunc;
    if (preStFunc) {
        UPG_LOG_DEBUG("Going to invoke pre-state handler function");
        if (!(preStateHandlers->*preStFunc)(ctx)) {
            UPG_LOG_DEBUG("pre-state handler function returned false");
            return false;
        }
    }
    return true;
}

bool UpgradeMgr::InvokePostStateHandler(UpgReqStateType reqType) {
    UpgPostStateFunc postStFunc = StateMachine[reqType].postStateFunc;
    if (postStFunc) {
        UPG_LOG_DEBUG("Going to invoke post-state handler function");
        if (!(postStateHandlers->*postStFunc)(ctx)) {
            UPG_LOG_DEBUG("post-state handler function returned false");
            return false;
        }
    }
    return true;
}

bool UpgradeMgr::InvokePrePostStateHandlers(UpgReqStateType reqType) {
    if (!InvokePostStateHandler(reqType)) {
        UPG_LOG_DEBUG("PostState handler returned false");
        return false;
    }
    reqType = GetNextState();
    if (!InvokePreStateHandler(reqType)) {
        UPG_LOG_DEBUG("PreState handler returned false");
        return false;
    }
    return true;
}

string UpgradeMgr::UpgReqStateTypeToStr(UpgReqStateType type) {
    return StateMachine[type].upgReqStateTypeToStr;
}

bool UpgradeMgr::GetAppRespFail(void) {
    return appRespFail_;
}

void UpgradeMgr::ResetAppResp(void) {
    appRespFail_ = false;
}

void UpgradeMgr::SetAppRespFail(void) {
    appRespFail_ = true;
}

void UpgradeMgr::AppendAppRespFailStr (string str) {
    appRespFailStrList_.push_back(str);
}

delphi::error UpgradeMgr::DeleteUpgMgrResp (void) {
    return upgMgrResp_->DeleteUpgMgrResp();
}
delphi::error UpgradeMgr::MoveStateMachine(UpgReqStateType type) {
    //Find UpgStateReq object
    UPG_LOG_DEBUG("UpgradeMgr::MoveStateMachine {}", type);
    auto reqStatus = findUpgStateReq(10);
    reqStatus->set_upgreqstate(type);
    sdk_->SetObject(reqStatus);
    if (type == UpgStateTerminal) {
        UpgRespType respType = UpgRespAbort;
        if (GetAppRespFail())
            respType = UpgRespFail;
        if (upgPassed_ && !upgAborted_)
            respType = UpgRespPass;
        upgMgrResp_->UpgradeFinish(respType, appRespFailStrList_);
        if (appRespFailStrList_.empty()) {
            UPG_LOG_DEBUG("Emptied all the responses from applications to agent");
            ResetAppResp();
            upgPassed_ = false;
            upgAborted_ = false;
        }
    }
    if (type != UpgStateTerminal)
        UPG_LOG_DEBUG("========== Upgrade state moved to {} ==========", UpgReqStateTypeToStr(type));
    return delphi::error::OK();
}

// OnUpgReqCreate gets called when UpgReq object is created
delphi::error UpgradeMgr::OnUpgReqCreate(delphi::objects::UpgReqPtr req) {
    UPG_LOG_DEBUG("UpgReq got created for {}/{}", req, req->meta().ShortDebugString());
    UPG_LOG_INFO("StartUpgrade request received");
    if (appRegMap_.size() == 0) {
        AppendAppRespFailStr("No app registered for upgrade");
        upgMgrResp_->UpgradeFinish(UpgRespFail, appRespFailStrList_);
        return delphi::error("No app registered for upgrade");
    }
    GetUpgCtxFromMeta(ctx);
    UpgReqStateType type = UpgStateCompatCheck;
    // find the status object
    auto upgReqStatus = findUpgStateReq(req->key());
    if (upgReqStatus == NULL) {
        // create it since it doesnt exist
        UpgPreStateFunc preStFunc = StateMachine[UpgStateCompatCheck].preStateFunc;
        if (preStFunc) {
            UPG_LOG_DEBUG("Going to invoke pre-state handler function");
            if (!(preStateHandlers->*preStFunc)(ctx)) {
                UPG_LOG_DEBUG("pre-state handler function returned false");
                type = UpgStateFailed;
                SetAppRespFail();
            }
        }
        RETURN_IF_FAILED(createUpgStateReq(req->key(), type, req->upgreqtype()));
    }

    return delphi::error::OK();
}

// OnUpgReqDelete gets called when UpgReq object is deleted
delphi::error UpgradeMgr::OnUpgReqDelete(delphi::objects::UpgReqPtr req) {
    UPG_LOG_DEBUG("UpgReq got deleted");
    auto upgReqStatus = findUpgStateReq(req->key());
    if (upgReqStatus != NULL) {
        UPG_LOG_DEBUG("Deleting Upgrade Request Status");
        sdk_->DeleteObject(upgReqStatus);
    }
    return delphi::error::OK();
}

delphi::error UpgradeMgr::StartUpgrade(uint32_t key) {
    delphi::objects::UpgStateReqPtr upgReqStatus = findUpgStateReq(key);
    if (upgReqStatus != NULL) {
        upgReqStatus->set_upgreqstate(UpgStateCompatCheck);
        sdk_->SetObject(upgReqStatus);
        UPG_LOG_DEBUG("Updated Upgrade Request Status UpgStateCompatCheck");
        return delphi::error::OK();
    }
    return delphi::error("Did not find UpgStateReqPtr");
}

delphi::error UpgradeMgr::AbortUpgrade(uint32_t key) {
    delphi::objects::UpgStateReqPtr upgReqStatus = findUpgStateReq(key);
    if (upgReqStatus != NULL) {
        upgAborted_ = true;
        upgReqStatus->set_upgreqstate(UpgStateAbort);
        sdk_->SetObject(upgReqStatus);
        UPG_LOG_DEBUG("Updated Upgrade Request Status UpgAborted");
        return delphi::error::OK();
    }
    return delphi::error("Did not find UpgStateReqPtr");
}

// OnUpgReqCmd gets called when UpgReqCmd attribute changes
delphi::error UpgradeMgr::OnUpgReqCmd(delphi::objects::UpgReqPtr req) {
    // start or abort?
    if (req->upgreqcmd() == UpgStart) {
        UPG_LOG_INFO("Start Upgrade");
        return StartUpgrade(req->key());
    } else if (req->upgreqcmd() == UpgAbort) {
        UPG_LOG_INFO("Abort Upgrade");
        return AbortUpgrade(req->key());
    }
    return delphi::error("Cannot decipher the upgreqcmd");
}

// createUpgStateReq creates a upgrade request status object
delphi::error UpgradeMgr::createUpgStateReq(uint32_t id, UpgReqStateType status, UpgType type) {
    // create an object
    delphi::objects::UpgStateReqPtr req = make_shared<delphi::objects::UpgStateReq>();
    req->set_key(id);
    req->set_upgreqstate(status);
    req->set_upgreqtype(type);

    // add it to database
    sdk_->SetObject(req);

    UPG_LOG_DEBUG("Created upgrade request status object for id {} state {} req: {}", id, status, req);

    return delphi::error::OK();
}

//  ffindUpgReqStat::objects::usinds the upgrade request status object
delphi::objects::UpgStateReqPtr UpgradeMgr::findUpgStateReq(uint32_t id) {
    delphi::objects::UpgStateReqPtr req = make_shared<delphi::objects::UpgStateReq>();
    req->set_key(id);

    // find the object
    delphi::BaseObjectPtr obj = sdk_->FindObject(req);

    return static_pointer_cast<delphi::objects::UpgStateReq>(obj);
}

delphi::objects::UpgReqPtr UpgradeMgr::findUpgReq(uint32_t id) {
    delphi::objects::UpgReqPtr req = make_shared<delphi::objects::UpgReq>();
    req->set_key(id);

    // find the object
    delphi::BaseObjectPtr obj = sdk_->FindObject(req);

    return static_pointer_cast<delphi::objects::UpgReq>(obj);
}

} // namespace upgrade
