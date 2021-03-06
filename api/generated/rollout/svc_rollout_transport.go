// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package rollout is a auto generated package.
Input file: svc_rollout.proto
*/
package rollout

import (
	"context"
	"encoding/json"
	"net/http"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/trace"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta

type grpcServerRolloutV1 struct {
	Endpoints EndpointsRolloutV1Server

	AutoAddRolloutHdlr          grpctransport.Handler
	AutoAddRolloutActionHdlr    grpctransport.Handler
	AutoDeleteRolloutHdlr       grpctransport.Handler
	AutoDeleteRolloutActionHdlr grpctransport.Handler
	AutoGetRolloutHdlr          grpctransport.Handler
	AutoGetRolloutActionHdlr    grpctransport.Handler
	AutoLabelRolloutHdlr        grpctransport.Handler
	AutoLabelRolloutActionHdlr  grpctransport.Handler
	AutoListRolloutHdlr         grpctransport.Handler
	AutoListRolloutActionHdlr   grpctransport.Handler
	AutoUpdateRolloutHdlr       grpctransport.Handler
	AutoUpdateRolloutActionHdlr grpctransport.Handler
	CreateRolloutHdlr           grpctransport.Handler
	RemoveRolloutHdlr           grpctransport.Handler
	StopRolloutHdlr             grpctransport.Handler
	UpdateRolloutHdlr           grpctransport.Handler
}

// MakeGRPCServerRolloutV1 creates a GRPC server for RolloutV1 service
func MakeGRPCServerRolloutV1(ctx context.Context, endpoints EndpointsRolloutV1Server, logger log.Logger) RolloutV1Server {
	return &grpcServerRolloutV1{
		Endpoints: endpoints,
		AutoAddRolloutHdlr: grpctransport.NewServer(
			endpoints.AutoAddRolloutEndpoint,
			DecodeGrpcReqRollout,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoAddRollout", logger)))...,
		),

		AutoAddRolloutActionHdlr: grpctransport.NewServer(
			endpoints.AutoAddRolloutActionEndpoint,
			DecodeGrpcReqRolloutAction,
			EncodeGrpcRespRolloutAction,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoAddRolloutAction", logger)))...,
		),

		AutoDeleteRolloutHdlr: grpctransport.NewServer(
			endpoints.AutoDeleteRolloutEndpoint,
			DecodeGrpcReqRollout,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoDeleteRollout", logger)))...,
		),

		AutoDeleteRolloutActionHdlr: grpctransport.NewServer(
			endpoints.AutoDeleteRolloutActionEndpoint,
			DecodeGrpcReqRolloutAction,
			EncodeGrpcRespRolloutAction,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoDeleteRolloutAction", logger)))...,
		),

		AutoGetRolloutHdlr: grpctransport.NewServer(
			endpoints.AutoGetRolloutEndpoint,
			DecodeGrpcReqRollout,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoGetRollout", logger)))...,
		),

		AutoGetRolloutActionHdlr: grpctransport.NewServer(
			endpoints.AutoGetRolloutActionEndpoint,
			DecodeGrpcReqRolloutAction,
			EncodeGrpcRespRolloutAction,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoGetRolloutAction", logger)))...,
		),

		AutoLabelRolloutHdlr: grpctransport.NewServer(
			endpoints.AutoLabelRolloutEndpoint,
			DecodeGrpcReqLabel,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoLabelRollout", logger)))...,
		),

		AutoLabelRolloutActionHdlr: grpctransport.NewServer(
			endpoints.AutoLabelRolloutActionEndpoint,
			DecodeGrpcReqLabel,
			EncodeGrpcRespRolloutAction,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoLabelRolloutAction", logger)))...,
		),

		AutoListRolloutHdlr: grpctransport.NewServer(
			endpoints.AutoListRolloutEndpoint,
			DecodeGrpcReqListWatchOptions,
			EncodeGrpcRespRolloutList,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoListRollout", logger)))...,
		),

		AutoListRolloutActionHdlr: grpctransport.NewServer(
			endpoints.AutoListRolloutActionEndpoint,
			DecodeGrpcReqListWatchOptions,
			EncodeGrpcRespRolloutActionList,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoListRolloutAction", logger)))...,
		),

		AutoUpdateRolloutHdlr: grpctransport.NewServer(
			endpoints.AutoUpdateRolloutEndpoint,
			DecodeGrpcReqRollout,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoUpdateRollout", logger)))...,
		),

		AutoUpdateRolloutActionHdlr: grpctransport.NewServer(
			endpoints.AutoUpdateRolloutActionEndpoint,
			DecodeGrpcReqRolloutAction,
			EncodeGrpcRespRolloutAction,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("AutoUpdateRolloutAction", logger)))...,
		),

		CreateRolloutHdlr: grpctransport.NewServer(
			endpoints.CreateRolloutEndpoint,
			DecodeGrpcReqRollout,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("CreateRollout", logger)))...,
		),

		RemoveRolloutHdlr: grpctransport.NewServer(
			endpoints.RemoveRolloutEndpoint,
			DecodeGrpcReqRollout,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("RemoveRollout", logger)))...,
		),

		StopRolloutHdlr: grpctransport.NewServer(
			endpoints.StopRolloutEndpoint,
			DecodeGrpcReqRollout,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("StopRollout", logger)))...,
		),

		UpdateRolloutHdlr: grpctransport.NewServer(
			endpoints.UpdateRolloutEndpoint,
			DecodeGrpcReqRollout,
			EncodeGrpcRespRollout,
			append([]grpctransport.ServerOption{grpctransport.ServerErrorLogger(logger), grpctransport.ServerBefore(recoverVersion)}, grpctransport.ServerBefore(trace.FromGRPCRequest("UpdateRollout", logger)))...,
		),
	}
}

func (s *grpcServerRolloutV1) AutoAddRollout(ctx oldcontext.Context, req *Rollout) (*Rollout, error) {
	_, resp, err := s.AutoAddRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoAddRollout).V
	return &r, resp.(respRolloutV1AutoAddRollout).Err
}

func decodeHTTPrespRolloutV1AutoAddRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoAddRolloutAction(ctx oldcontext.Context, req *RolloutAction) (*RolloutAction, error) {
	_, resp, err := s.AutoAddRolloutActionHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoAddRolloutAction).V
	return &r, resp.(respRolloutV1AutoAddRolloutAction).Err
}

func decodeHTTPrespRolloutV1AutoAddRolloutAction(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp RolloutAction
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoDeleteRollout(ctx oldcontext.Context, req *Rollout) (*Rollout, error) {
	_, resp, err := s.AutoDeleteRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoDeleteRollout).V
	return &r, resp.(respRolloutV1AutoDeleteRollout).Err
}

func decodeHTTPrespRolloutV1AutoDeleteRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoDeleteRolloutAction(ctx oldcontext.Context, req *RolloutAction) (*RolloutAction, error) {
	_, resp, err := s.AutoDeleteRolloutActionHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoDeleteRolloutAction).V
	return &r, resp.(respRolloutV1AutoDeleteRolloutAction).Err
}

func decodeHTTPrespRolloutV1AutoDeleteRolloutAction(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp RolloutAction
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoGetRollout(ctx oldcontext.Context, req *Rollout) (*Rollout, error) {
	_, resp, err := s.AutoGetRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoGetRollout).V
	return &r, resp.(respRolloutV1AutoGetRollout).Err
}

func decodeHTTPrespRolloutV1AutoGetRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoGetRolloutAction(ctx oldcontext.Context, req *RolloutAction) (*RolloutAction, error) {
	_, resp, err := s.AutoGetRolloutActionHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoGetRolloutAction).V
	return &r, resp.(respRolloutV1AutoGetRolloutAction).Err
}

func decodeHTTPrespRolloutV1AutoGetRolloutAction(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp RolloutAction
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoLabelRollout(ctx oldcontext.Context, req *api.Label) (*Rollout, error) {
	_, resp, err := s.AutoLabelRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoLabelRollout).V
	return &r, resp.(respRolloutV1AutoLabelRollout).Err
}

func decodeHTTPrespRolloutV1AutoLabelRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoLabelRolloutAction(ctx oldcontext.Context, req *api.Label) (*RolloutAction, error) {
	_, resp, err := s.AutoLabelRolloutActionHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoLabelRolloutAction).V
	return &r, resp.(respRolloutV1AutoLabelRolloutAction).Err
}

func decodeHTTPrespRolloutV1AutoLabelRolloutAction(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp RolloutAction
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoListRollout(ctx oldcontext.Context, req *api.ListWatchOptions) (*RolloutList, error) {
	_, resp, err := s.AutoListRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoListRollout).V
	return &r, resp.(respRolloutV1AutoListRollout).Err
}

func decodeHTTPrespRolloutV1AutoListRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp RolloutList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoListRolloutAction(ctx oldcontext.Context, req *api.ListWatchOptions) (*RolloutActionList, error) {
	_, resp, err := s.AutoListRolloutActionHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoListRolloutAction).V
	return &r, resp.(respRolloutV1AutoListRolloutAction).Err
}

func decodeHTTPrespRolloutV1AutoListRolloutAction(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp RolloutActionList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoUpdateRollout(ctx oldcontext.Context, req *Rollout) (*Rollout, error) {
	_, resp, err := s.AutoUpdateRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoUpdateRollout).V
	return &r, resp.(respRolloutV1AutoUpdateRollout).Err
}

func decodeHTTPrespRolloutV1AutoUpdateRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoUpdateRolloutAction(ctx oldcontext.Context, req *RolloutAction) (*RolloutAction, error) {
	_, resp, err := s.AutoUpdateRolloutActionHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1AutoUpdateRolloutAction).V
	return &r, resp.(respRolloutV1AutoUpdateRolloutAction).Err
}

func decodeHTTPrespRolloutV1AutoUpdateRolloutAction(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp RolloutAction
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) CreateRollout(ctx oldcontext.Context, req *Rollout) (*Rollout, error) {
	_, resp, err := s.CreateRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1CreateRollout).V
	return &r, resp.(respRolloutV1CreateRollout).Err
}

func decodeHTTPrespRolloutV1CreateRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) RemoveRollout(ctx oldcontext.Context, req *Rollout) (*Rollout, error) {
	_, resp, err := s.RemoveRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1RemoveRollout).V
	return &r, resp.(respRolloutV1RemoveRollout).Err
}

func decodeHTTPrespRolloutV1RemoveRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) StopRollout(ctx oldcontext.Context, req *Rollout) (*Rollout, error) {
	_, resp, err := s.StopRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1StopRollout).V
	return &r, resp.(respRolloutV1StopRollout).Err
}

func decodeHTTPrespRolloutV1StopRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) UpdateRollout(ctx oldcontext.Context, req *Rollout) (*Rollout, error) {
	_, resp, err := s.UpdateRolloutHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respRolloutV1UpdateRollout).V
	return &r, resp.(respRolloutV1UpdateRollout).Err
}

func decodeHTTPrespRolloutV1UpdateRollout(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp Rollout
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerRolloutV1) AutoWatchSvcRolloutV1(in *api.AggWatchOptions, stream RolloutV1_AutoWatchSvcRolloutV1Server) error {
	return s.Endpoints.AutoWatchSvcRolloutV1(in, stream)
}

func (s *grpcServerRolloutV1) AutoWatchRollout(in *api.ListWatchOptions, stream RolloutV1_AutoWatchRolloutServer) error {
	return s.Endpoints.AutoWatchRollout(in, stream)
}

func (s *grpcServerRolloutV1) AutoWatchRolloutAction(in *api.ListWatchOptions, stream RolloutV1_AutoWatchRolloutActionServer) error {
	return s.Endpoints.AutoWatchRolloutAction(in, stream)
}

func encodeHTTPRolloutActionList(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPRolloutActionList(_ context.Context, r *http.Request) (interface{}, error) {
	var req RolloutActionList
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqRolloutActionList encodes GRPC request
func EncodeGrpcReqRolloutActionList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*RolloutActionList)
	return req, nil
}

// DecodeGrpcReqRolloutActionList decodes GRPC request
func DecodeGrpcReqRolloutActionList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*RolloutActionList)
	return req, nil
}

// EncodeGrpcRespRolloutActionList endodes the GRPC response
func EncodeGrpcRespRolloutActionList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespRolloutActionList decodes the GRPC response
func DecodeGrpcRespRolloutActionList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPRolloutList(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPRolloutList(_ context.Context, r *http.Request) (interface{}, error) {
	var req RolloutList
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqRolloutList encodes GRPC request
func EncodeGrpcReqRolloutList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*RolloutList)
	return req, nil
}

// DecodeGrpcReqRolloutList decodes GRPC request
func DecodeGrpcReqRolloutList(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*RolloutList)
	return req, nil
}

// EncodeGrpcRespRolloutList endodes the GRPC response
func EncodeGrpcRespRolloutList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespRolloutList decodes the GRPC response
func DecodeGrpcRespRolloutList(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
