// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package browser is a auto generated package.
Input file: svc_browser.proto
*/
package browser

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	oldcontext "golang.org/x/net/context"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/trace"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta

type grpcServerBrowserV1 struct {
	Endpoints EndpointsBrowserV1Server

	QueryHdlr      grpctransport.Handler
	ReferencesHdlr grpctransport.Handler
	ReferrersHdlr  grpctransport.Handler
}

// MakeGRPCServerBrowserV1 creates a GRPC server for BrowserV1 service
func MakeGRPCServerBrowserV1(ctx context.Context, endpoints EndpointsBrowserV1Server, logger log.Logger) BrowserV1Server {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(recoverVersion),
	}
	return &grpcServerBrowserV1{
		Endpoints: endpoints,
		QueryHdlr: grpctransport.NewServer(
			endpoints.QueryEndpoint,
			DecodeGrpcReqBrowseRequestList,
			EncodeGrpcRespBrowseResponseList,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("Query", logger)))...,
		),

		ReferencesHdlr: grpctransport.NewServer(
			endpoints.ReferencesEndpoint,
			DecodeGrpcReqBrowseRequest,
			EncodeGrpcRespBrowseResponse,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("References", logger)))...,
		),

		ReferrersHdlr: grpctransport.NewServer(
			endpoints.ReferrersEndpoint,
			DecodeGrpcReqBrowseRequest,
			EncodeGrpcRespBrowseResponse,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("Referrers", logger)))...,
		),
	}
}

func (s *grpcServerBrowserV1) Query(ctx oldcontext.Context, req *BrowseRequestList) (*BrowseResponseList, error) {
	_, resp, err := s.QueryHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respBrowserV1Query).V
	return &r, resp.(respBrowserV1Query).Err
}

func decodeHTTPrespBrowserV1Query(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp BrowseResponseList
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerBrowserV1) References(ctx oldcontext.Context, req *BrowseRequest) (*BrowseResponse, error) {
	_, resp, err := s.ReferencesHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respBrowserV1References).V
	return &r, resp.(respBrowserV1References).Err
}

func decodeHTTPrespBrowserV1References(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp BrowseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerBrowserV1) Referrers(ctx oldcontext.Context, req *BrowseRequest) (*BrowseResponse, error) {
	_, resp, err := s.ReferrersHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respBrowserV1Referrers).V
	return &r, resp.(respBrowserV1Referrers).Err
}

func decodeHTTPrespBrowserV1Referrers(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp BrowseResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerBrowserV1) AutoWatchSvcBrowserV1(in *api.ListWatchOptions, stream BrowserV1_AutoWatchSvcBrowserV1Server) error {
	return errors.New("not implemented")
}
