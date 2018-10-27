//------------------------------------------------------------------------------
// {C} Copyright 2018 Pensando Systems Inc. All rights reserved
//------------------------------------------------------------------------------

#include <iostream>
#include <stdio.h>
#include "nic/include/base.hpp"
#include "nic/utils/trace/trace.hpp"
#include "nic/hal/iris/delphi/delphi.hpp"
#include "nic/hal/iris/delphi/delphic.hpp"
#include "nic/linkmgr/delphi/linkmgr_delphi.hpp"

namespace hal {
namespace svc {

using namespace std;
using grpc::Status;
using hal::upgrade::upgrade_handler;


delphi::SdkPtr sdk(make_shared<delphi::Sdk>());
shared_ptr<delphi_client> delphic = make_shared<delphi_client>(sdk);

//------------------------------------------------------------------------------
// starting point for delphi thread
//------------------------------------------------------------------------------
void *
delphi_client_start (void *ctxt)
{
    HAL_ASSERT(delphic != NULL);

    HAL_TRACE_DEBUG("HAL delphi thread started...");

    // init linkmgr services
    Status sts = linkmgr::port_svc_init(sdk);
    HAL_ABORT(sts.ok());

    // register delphi client
    sdk->RegisterService(delphic);

    // sit in main loop
    sdk->MainLoop();

    // shouldn't hit
    return NULL;
}

//------------------------------------------------------------------------------
// delphi_client constructor
//------------------------------------------------------------------------------
delphi_client::delphi_client(delphi::SdkPtr &sdk) 
{
    sdk_ = sdk;
    sysmgr_ = ::sysmgr::CreateClient(sdk_, "hal");
    upgsdk_ =
        make_shared<::upgrade::UpgSdk>(sdk_, make_shared<upgrade_handler>(),
                                       "hal", ::upgrade::NON_AGENT, nullptr);

    // create the InterfaceSpec reactor
    if_svc_ = std::make_shared<if_svc>(sdk);

    mount_ok = false;
    init_ok = false;

    // mount InterfaceSpec objects
    delphi::objects::InterfaceSpec::Mount(sdk, delphi::ReadMode);

    // Register InterfaceSpec reactor
    delphi::objects::InterfaceSpec::Watch(sdk, if_svc_);

    // mount interface status objects
    delphi::objects::InterfaceStatus::Mount(sdk, delphi::ReadWriteMode);
}

// OnMountComplete gets called when all the objects are mounted
void
delphi_client::OnMountComplete(void)
{
    HAL_TRACE_DEBUG("OnMountComplete got called..");
    mount_ok = true;
    if (init_ok && this->mount_ok) {
       sysmgr_->init_done();
    }
}

void
delphi_client::init_done(void)
{
   HAL_TRACE_DEBUG("Init done called..");
   init_ok = true;
   if (init_ok && mount_ok) {
      sysmgr_->init_done();
   }
}
   
}    // namespace svc
}    // namespace hal
