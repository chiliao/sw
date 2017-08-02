#ifndef __INTERFACE_SVC_HPP__
#define __INTERFACE_SVC_HPP__

#include <base.h>
#include <grpc++/grpc++.h>
#include <types.pb.h>
#include <interface.grpc.pb.h>

using grpc::ServerContext;
using grpc::Status;

using intf::Interface;
using intf::LifSpec;
using intf::LifKeyHandle;
using intf::LifRequestMsg;
using intf::LifResponse;
using intf::LifResponseMsg;
using intf::LifDeleteRequestMsg;
using intf::LifDeleteResponseMsg;
using intf::LifGetRequestMsg;
using intf::LifGetResponseMsg;
using intf::InterfaceSpec;
using intf::InterfaceStatus;
using intf::InterfaceResponse;
using intf::InterfaceKeyHandle;
using intf::InterfaceRequestMsg;
using intf::InterfaceResponseMsg;
using intf::InterfaceDeleteRequestMsg;
using intf::InterfaceDeleteResponseMsg;
using intf::InterfaceGetRequest;
using intf::InterfaceGetRequestMsg;
using intf::InterfaceGetResponse;
using intf::InterfaceGetResponseMsg;
using intf::InterfaceL2SegmentRequestMsg;
using intf::InterfaceL2SegmentSpec;
using intf::InterfaceL2SegmentResponseMsg;
using intf::InterfaceL2SegmentResponse;

class InterfaceServiceImpl final : public Interface::Service {
public:
    Status LifCreate(ServerContext *context,
                     const LifRequestMsg *req,
                     LifResponseMsg *rsp) override;

    Status LifUpdate(ServerContext *context,
                     const LifRequestMsg *req,
                     LifResponseMsg *rsp) override;

    Status LifDelete(ServerContext *context,
                     const LifDeleteRequestMsg *req,
                     LifDeleteResponseMsg *rsp) override;

    Status LifGet(ServerContext *context,
                  const LifGetRequestMsg *req,
                  LifGetResponseMsg *rsp) override;

    Status InterfaceCreate(ServerContext *context,
                           const InterfaceRequestMsg *req,
                           InterfaceResponseMsg *rsp) override;

    Status InterfaceUpdate(ServerContext *context,
                           const InterfaceRequestMsg *req,
                           InterfaceResponseMsg *rsp) override;

    Status InterfaceDelete(ServerContext *context,
                           const InterfaceDeleteRequestMsg *req,
                           InterfaceDeleteResponseMsg *rsp) override;

    Status InterfaceGet(ServerContext *context,
                        const InterfaceGetRequestMsg *req,
                        InterfaceGetResponseMsg *rsp) override;

    Status AddL2SegmentOnUplink(ServerContext *context,
                                const InterfaceL2SegmentRequestMsg *req,
                                InterfaceL2SegmentResponseMsg *rsp) override;

    Status DelL2SegmentOnUplink(ServerContext *context,
                                const InterfaceL2SegmentRequestMsg *req,
                                InterfaceL2SegmentResponseMsg *rsp) override;
                
};

#endif    // __INTERFACE_SVC_HPP__

