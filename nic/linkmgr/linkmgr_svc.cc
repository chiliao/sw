// {C} Copyright 2017 Pensando Systems Inc. All rights reserved

//------------------------------------------------------------------------------
// Port service implementation
//------------------------------------------------------------------------------

#include "nic/include/base.h"
#include "nic/include/trace.hpp"

#include "linkmgr_svc.hpp"
#include "linkmgr_src.hpp"

Status
PortServiceImpl::PortCreate(ServerContext *context,
                            const PortRequestMsg *req,
                            PortResponseMsg *rsp)
{
    uint32_t        i, nreqs = req->request_size();
    PortResponse    *response;

    HAL_TRACE_DEBUG("Rcvd Port Create Request");
    if (nreqs == 0) {
        return Status(grpc::StatusCode::INVALID_ARGUMENT, "Empty Request");
    }

    linkmgr::g_linkmgr_state->cfg_db_open(hal::CFG_OP_WRITE);
    for (i = 0; i < nreqs; i++) {
        response = rsp->add_response();
        auto spec = req->request(i);
        linkmgr::port_create(spec, response);
    }
    linkmgr::g_linkmgr_state->cfg_db_close();
    return Status::OK;
}

Status
PortServiceImpl::PortUpdate(ServerContext *context,
                            const PortRequestMsg *req,
                            PortResponseMsg *rsp)
{
    uint32_t        i, nreqs = req->request_size();
    PortResponse    *response;

    HAL_TRACE_DEBUG("Rcvd Port Update Request");
    if (nreqs == 0) {
        return Status(grpc::StatusCode::INVALID_ARGUMENT, "Empty Request");
    }

    linkmgr::g_linkmgr_state->cfg_db_open(hal::CFG_OP_WRITE);
    for (i = 0; i < nreqs; i++) {
        response = rsp->add_response();
        auto spec = req->request(i);
        linkmgr::port_update(spec, response);
    }
    linkmgr::g_linkmgr_state->cfg_db_close();
    return Status::OK;
}


Status
PortServiceImpl::PortDelete(ServerContext *context,
                            const PortDeleteRequestMsg *req,
                            PortDeleteResponseMsg *rsp)
{
    uint32_t     i, nreqs = req->request_size();

    HAL_TRACE_DEBUG("Rcvd Port Delete Request");
    if (nreqs == 0) {
        return Status(grpc::StatusCode::INVALID_ARGUMENT, "Empty Request");
    }

    linkmgr::g_linkmgr_state->cfg_db_open(hal::CFG_OP_WRITE);
    for (i = 0; i < nreqs; i++) {
        auto spec = req->request(i);
        linkmgr::port_delete(spec, rsp);
    }
    linkmgr::g_linkmgr_state->cfg_db_close();
    return Status::OK;
}

Status
PortServiceImpl::PortGet(ServerContext *context,
                         const PortGetRequestMsg *req,
                         PortGetResponseMsg *rsp)
{
    uint32_t           i, nreqs = req->request_size();
    PortGetResponse    *response;

    HAL_TRACE_DEBUG("Rcvd Port Get Request");
    if (nreqs == 0) {
        return Status(grpc::StatusCode::INVALID_ARGUMENT, "Empty Request");
    }

    linkmgr::g_linkmgr_state->cfg_db_open(hal::CFG_OP_READ);
    for (i = 0; i < nreqs; i++) {
        response = rsp->add_response();
        auto request = req->request(i);
        linkmgr::port_get(request, response);
    }
    linkmgr::g_linkmgr_state->cfg_db_close();
    return Status::OK;
}
