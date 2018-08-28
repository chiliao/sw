#include <memory>
#include <string>

#include "nic/delphi/sdk/delphi_sdk.hpp"
#include "nic/sysmgr/proto/sysmgr.delphi.hpp"

#include "delphi_messages.hpp"
#include "sysmgr_service.hpp"
#include "sysmgr_service_status_reactor.hpp"
#include "sysmgr_test_complete_req_reactor.hpp"

using namespace std;

SysmgrService::SysmgrService(delphi::SdkPtr sdk, string name, shared_ptr<Pipe<pid_t> > started_pids_pipe, 
    shared_ptr<Pipe<int32_t> > delphi_message_pipe)
{
    this->sdk = sdk;
    this->name = name;
    this->delphi_message_pipe = delphi_message_pipe;

    serviceStatusReactor = make_shared<SysmgrServiceStatusReactor>(started_pids_pipe);
    testCompleteReactor = make_shared<SysmgrTestCompleteReqReactor>(delphi_message_pipe);

    delphi::objects::SysmgrServiceStatus::Mount(this->sdk, delphi::ReadMode);
    delphi::objects::SysmgrTestCompleteReq::Mount(this->sdk, delphi::ReadMode);

    delphi::objects::SysmgrServiceStatus::Watch(this->sdk, serviceStatusReactor);
    delphi::objects::SysmgrTestCompleteReq::Watch(this->sdk, testCompleteReactor);
}

void SysmgrService::OnMountComplete()
{
    delphi_message_pipe->pipe_write(DELPHI_UP);
}