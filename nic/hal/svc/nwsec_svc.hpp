#ifndef __NWSEC_SVC_HPP__
#define __NWSEC_SVC_HPP__

#include <base.h>
#include <grpc++/grpc++.h>
#include <types.pb.h>
#include <nwsec.grpc.pb.h>

using grpc::ServerContext;
using grpc::Status;

using nwsec::NwSecurity;
using nwsec::SecurityProfileKeyHandle;
using nwsec::SecurityProfileSpec;
using nwsec::SecurityProfileRequestMsg;
using nwsec::SecurityProfileStatus;
using nwsec::SecurityProfileResponse;
using nwsec::SecurityProfileResponseMsg;
using nwsec::SecurityProfileDeleteRequest;
using nwsec::SecurityProfileDeleteRequestMsg;
using nwsec::SecurityProfileDeleteResponseMsg;
using nwsec::SecurityProfileGetRequest;
using nwsec::SecurityProfileGetRequestMsg;
using nwsec::SecurityProfileStats;
using nwsec::SecurityProfileGetResponse;
using nwsec::SecurityProfileGetResponseMsg;
using nwsec::SecurityGroupSpec;
using nwsec::SecurityGroupMsg;
using nwsec::SecurityGroupStatus;
using nwsec::SecurityGroupResponse;
using nwsec::SecurityGroupResponseMsg;
using nwsec::Service;
using nwsec::SecurityPolicyRuleSpec;
using nwsec::SecurityPolicyRuleMsg;
using nwsec::SecurityPolicyRuleStatus;
using nwsec::SecurityPolicyRuleResponse;
using nwsec::SecurityPolicyRuleResponseMsg;

class NwSecurityServiceImpl final : public NwSecurity::Service {
public:
    Status SecurityProfileCreate(ServerContext *context,
                                 const SecurityProfileRequestMsg *req,
                                 SecurityProfileResponseMsg *rsp) override;

    Status SecurityProfileUpdate(ServerContext *context,
                                 const SecurityProfileRequestMsg *req,
                                 SecurityProfileResponseMsg *rsp) override;


    Status SecurityProfileDelete(ServerContext *context,
                                 const SecurityProfileDeleteRequestMsg *req,
                                 SecurityProfileDeleteResponseMsg *rsp) override;

    Status SecurityProfileGet(ServerContext *context,
                              const SecurityProfileGetRequestMsg *req,
                              SecurityProfileGetResponseMsg *rsp) override;
};

#endif    // __NWSEC_SVC_HPP__

