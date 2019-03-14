// {C} Copyright 2019 Pensando Systems Inc. All rights reserved

#ifndef __AGENT_SVC_ROUTE_HPP__
#define __AGENT_SVC_ROUTE_HPP__

#include "grpc++/grpc++.h"
#include "gen/proto/types.pb.h"
#include "gen/proto/meta/meta.pb.h"
#include "gen/proto/route.grpc.pb.h"

using grpc::Status;
using grpc::ServerContext;

using pds::RouteSvc;
using pds::RouteTableRequest;
using pds::RouteTableSpec;
using pds::RouteTableResponse;

class RouteSvcImpl final : public RouteSvc::Service {
public:
    Status RouteTableCreate(ServerContext *context,
                            const pds::RouteTableRequest *req,
                            pds::RouteTableResponse *rsp) override;
};

#endif    // __AGENT_SVC_ROUTE_HPP__
