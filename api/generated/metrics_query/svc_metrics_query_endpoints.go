// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package metrics_query is a auto generated package.
Input file: svc_metrics_query.proto
*/
package metrics_query

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	"google.golang.org/grpc"

	"github.com/pensando/sw/api"
	loginctx "github.com/pensando/sw/api/login/context"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/trace"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta
var _ grpc.ServerStream
var _ fmt.Formatter

// MiddlewareMetricsV1Client add middleware to the client
type MiddlewareMetricsV1Client func(ServiceMetricsV1Client) ServiceMetricsV1Client

// EndpointsMetricsV1Client is the endpoints for the client
type EndpointsMetricsV1Client struct {
	Client                        MetricsV1Client
	AutoWatchSvcMetricsV1Endpoint endpoint.Endpoint

	QueryEndpoint endpoint.Endpoint
}

// EndpointsMetricsV1RestClient is the REST client
type EndpointsMetricsV1RestClient struct {
	logger   log.Logger
	client   *http.Client
	instance string

	AutoWatchSvcMetricsV1Endpoint endpoint.Endpoint
	QueryEndpoint                 endpoint.Endpoint
}

// MiddlewareMetricsV1Server adds middle ware to the server
type MiddlewareMetricsV1Server func(ServiceMetricsV1Server) ServiceMetricsV1Server

// EndpointsMetricsV1Server is the server endpoints
type EndpointsMetricsV1Server struct {
	svcWatchHandlerMetricsV1 func(options *api.ListWatchOptions, stream grpc.ServerStream) error

	QueryEndpoint endpoint.Endpoint
}

// Query is endpoint for Query
func (e EndpointsMetricsV1Client) Query(ctx context.Context, in *QuerySpec) (*QueryResponse, error) {
	resp, err := e.QueryEndpoint(ctx, in)
	if err != nil {
		return &QueryResponse{}, err
	}
	return resp.(*QueryResponse), nil
}

type respMetricsV1Query struct {
	V   QueryResponse
	Err error
}

func (e EndpointsMetricsV1Client) AutoWatchSvcMetricsV1(ctx context.Context, in *api.ListWatchOptions) (MetricsV1_AutoWatchSvcMetricsV1Client, error) {
	return nil, errors.New("not implemented")
}

// Query implementation on server Endpoint
func (e EndpointsMetricsV1Server) Query(ctx context.Context, in QuerySpec) (QueryResponse, error) {
	resp, err := e.QueryEndpoint(ctx, in)
	if err != nil {
		return QueryResponse{}, err
	}
	return *resp.(*QueryResponse), nil
}

// MakeMetricsV1QueryEndpoint creates  Query endpoints for the service
func MakeMetricsV1QueryEndpoint(s ServiceMetricsV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*QuerySpec)
		v, err := s.Query(ctx, *req)
		return respMetricsV1Query{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("MetricsV1:Query")(f)
}

// MakeAutoWatchSvcMetricsV1Endpoint creates the Watch endpoint for the service
func MakeAutoWatchSvcMetricsV1Endpoint(s ServiceMetricsV1Server, logger log.Logger) func(options *api.ListWatchOptions, stream grpc.ServerStream) error {
	return func(options *api.ListWatchOptions, stream grpc.ServerStream) error {
		return errors.New("not implemented")
	}
}

// MakeMetricsV1ServerEndpoints creates server endpoints
func MakeMetricsV1ServerEndpoints(s ServiceMetricsV1Server, logger log.Logger) EndpointsMetricsV1Server {
	return EndpointsMetricsV1Server{
		svcWatchHandlerMetricsV1: MakeAutoWatchSvcMetricsV1Endpoint(s, logger),

		QueryEndpoint: MakeMetricsV1QueryEndpoint(s, logger),
	}
}

// LoggingMetricsV1MiddlewareClient adds middleware for the client
func LoggingMetricsV1MiddlewareClient(logger log.Logger) MiddlewareMetricsV1Client {
	return func(next ServiceMetricsV1Client) ServiceMetricsV1Client {
		return loggingMetricsV1MiddlewareClient{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMetricsV1MiddlewareClient struct {
	logger log.Logger
	next   ServiceMetricsV1Client
}

// LoggingMetricsV1MiddlewareServer adds middleware for the client
func LoggingMetricsV1MiddlewareServer(logger log.Logger) MiddlewareMetricsV1Server {
	return func(next ServiceMetricsV1Server) ServiceMetricsV1Server {
		return loggingMetricsV1MiddlewareServer{
			logger: logger,
			next:   next,
		}
	}
}

type loggingMetricsV1MiddlewareServer struct {
	logger log.Logger
	next   ServiceMetricsV1Server
}

func (m loggingMetricsV1MiddlewareClient) Query(ctx context.Context, in *QuerySpec) (resp *QueryResponse, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "MetricsV1", "method", "Query", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.Query(ctx, in)
	return
}

func (m loggingMetricsV1MiddlewareClient) AutoWatchSvcMetricsV1(ctx context.Context, in *api.ListWatchOptions) (MetricsV1_AutoWatchSvcMetricsV1Client, error) {
	return nil, errors.New("not implemented")
}

func (m loggingMetricsV1MiddlewareServer) Query(ctx context.Context, in QuerySpec) (resp QueryResponse, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "MetricsV1", "method", "Query", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.Query(ctx, in)
	return
}

func (m loggingMetricsV1MiddlewareServer) AutoWatchSvcMetricsV1(in *api.ListWatchOptions, stream MetricsV1_AutoWatchSvcMetricsV1Server) error {
	return errors.New("Not implemented")
}

func (r *EndpointsMetricsV1RestClient) getHTTPRequest(ctx context.Context, in interface{}, method, path string) (*http.Request, error) {
	target, err := url.Parse(r.instance)
	if err != nil {
		return nil, fmt.Errorf("invalid instance %s", r.instance)
	}
	target.Path = path
	req, err := http.NewRequest(method, target.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request (%s)", err)
	}
	val, ok := loginctx.AuthzHeaderFromContext(ctx)
	if ok {
		req.Header.Add("Authorization", val)
	}
	if err = encodeHTTPRequest(ctx, req, in); err != nil {
		return nil, fmt.Errorf("could not encode request (%s)", err)
	}
	return req, nil
}

func (r *EndpointsMetricsV1RestClient) MetricsV1QueryEndpoint(ctx context.Context, in *QuerySpec) (*QueryResponse, error) {
	return nil, errors.New("not allowed")
}

// MakeMetricsV1RestClientEndpoints make REST client endpoints
func MakeMetricsV1RestClientEndpoints(instance string) (EndpointsMetricsV1RestClient, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	return EndpointsMetricsV1RestClient{
		instance: instance,
		client:   http.DefaultClient,
	}, nil

}
