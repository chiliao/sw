// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package telemetry_query is a auto generated package.
Input file: svc_telemetry_query.proto
*/
package telemetry_query

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

type grpcServerTelemetryV1 struct {
	Endpoints EndpointsTelemetryV1Server

	MetricsHdlr grpctransport.Handler
}

// MakeGRPCServerTelemetryV1 creates a GRPC server for TelemetryV1 service
func MakeGRPCServerTelemetryV1(ctx context.Context, endpoints EndpointsTelemetryV1Server, logger log.Logger) TelemetryV1Server {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
		grpctransport.ServerBefore(recoverVersion),
	}
	return &grpcServerTelemetryV1{
		Endpoints: endpoints,
		MetricsHdlr: grpctransport.NewServer(
			endpoints.MetricsEndpoint,
			DecodeGrpcReqMetricsQueryList,
			EncodeGrpcRespMetricsQueryResponse,
			append(options, grpctransport.ServerBefore(trace.FromGRPCRequest("Metrics", logger)))...,
		),
	}
}

func (s *grpcServerTelemetryV1) Metrics(ctx oldcontext.Context, req *MetricsQueryList) (*MetricsQueryResponse, error) {
	_, resp, err := s.MetricsHdlr.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	r := resp.(respTelemetryV1Metrics).V
	return &r, resp.(respTelemetryV1Metrics).Err
}

func decodeHTTPrespTelemetryV1Metrics(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errorDecoder(r)
	}
	var resp MetricsQueryResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return &resp, err
}

func (s *grpcServerTelemetryV1) AutoWatchSvcTelemetryV1(in *api.ListWatchOptions, stream TelemetryV1_AutoWatchSvcTelemetryV1Server) error {
	return errors.New("not implemented")
}
