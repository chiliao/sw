// {C} Copyright 2019 Pensando Systems Inc. All rights reserved

#ifndef __AGENT_SVC_MAPPING_HPP__
#define __AGENT_SVC_MAPPING_HPP__

#include "grpc++/grpc++.h"
#include "gen/proto/types.pb.h"
#include "gen/proto/meta/meta.pb.h"
#include "gen/proto/mapping.grpc.pb.h"

using grpc::Status;
using grpc::ServerContext;

using pds::MappingSvc;
using pds::MappingRequest;
using pds::MappingSpec;
using pds::MappingResponse;

class MappingSvcImpl final : public MappingSvc::Service {
public:
    Status MappingCreate(ServerContext *context, const pds::MappingRequest *req,
                         pds::MappingResponse *rsp) override;
};

#endif    // __AGENT_SVC_MAPPING_HPP__
