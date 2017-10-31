// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package collection is a auto generated package.
Input file: protos/collection.proto
*/
package collection

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
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/trace"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta
var _ grpc.ServerStream
var _ fmt.Formatter

// MiddlewareCollectionPolicyV1Client add middleware to the client
type MiddlewareCollectionPolicyV1Client func(ServiceCollectionPolicyV1Client) ServiceCollectionPolicyV1Client

// EndpointsCollectionPolicyV1Client is the endpoints for the client
type EndpointsCollectionPolicyV1Client struct {
	Client CollectionPolicyV1Client

	AutoAddCollectionPolicyEndpoint    endpoint.Endpoint
	AutoDeleteCollectionPolicyEndpoint endpoint.Endpoint
	AutoGetCollectionPolicyEndpoint    endpoint.Endpoint
	AutoListCollectionPolicyEndpoint   endpoint.Endpoint
	AutoUpdateCollectionPolicyEndpoint endpoint.Endpoint
}

// EndpointsCollectionPolicyV1RestClient is the REST client
type EndpointsCollectionPolicyV1RestClient struct {
	logger   log.Logger
	client   *http.Client
	instance string

	AutoAddCollectionPolicyEndpoint    endpoint.Endpoint
	AutoDeleteCollectionPolicyEndpoint endpoint.Endpoint
	AutoGetCollectionPolicyEndpoint    endpoint.Endpoint
	AutoListCollectionPolicyEndpoint   endpoint.Endpoint
	AutoUpdateCollectionPolicyEndpoint endpoint.Endpoint
	AutoWatchCollectionPolicyEndpoint  endpoint.Endpoint
}

// MiddlewareCollectionPolicyV1Server adds middle ware to the server
type MiddlewareCollectionPolicyV1Server func(ServiceCollectionPolicyV1Server) ServiceCollectionPolicyV1Server

// EndpointsCollectionPolicyV1Server is the server endpoints
type EndpointsCollectionPolicyV1Server struct {
	AutoAddCollectionPolicyEndpoint    endpoint.Endpoint
	AutoDeleteCollectionPolicyEndpoint endpoint.Endpoint
	AutoGetCollectionPolicyEndpoint    endpoint.Endpoint
	AutoListCollectionPolicyEndpoint   endpoint.Endpoint
	AutoUpdateCollectionPolicyEndpoint endpoint.Endpoint

	watchHandlerCollectionPolicy func(options *api.ListWatchOptions, stream grpc.ServerStream) error
}

// AutoAddCollectionPolicy is endpoint for AutoAddCollectionPolicy
func (e EndpointsCollectionPolicyV1Client) AutoAddCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	resp, err := e.AutoAddCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return &CollectionPolicy{}, err
	}
	return resp.(*CollectionPolicy), nil
}

type respCollectionPolicyV1AutoAddCollectionPolicy struct {
	V   CollectionPolicy
	Err error
}

// AutoDeleteCollectionPolicy is endpoint for AutoDeleteCollectionPolicy
func (e EndpointsCollectionPolicyV1Client) AutoDeleteCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	resp, err := e.AutoDeleteCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return &CollectionPolicy{}, err
	}
	return resp.(*CollectionPolicy), nil
}

type respCollectionPolicyV1AutoDeleteCollectionPolicy struct {
	V   CollectionPolicy
	Err error
}

// AutoGetCollectionPolicy is endpoint for AutoGetCollectionPolicy
func (e EndpointsCollectionPolicyV1Client) AutoGetCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	resp, err := e.AutoGetCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return &CollectionPolicy{}, err
	}
	return resp.(*CollectionPolicy), nil
}

type respCollectionPolicyV1AutoGetCollectionPolicy struct {
	V   CollectionPolicy
	Err error
}

// AutoListCollectionPolicy is endpoint for AutoListCollectionPolicy
func (e EndpointsCollectionPolicyV1Client) AutoListCollectionPolicy(ctx context.Context, in *api.ListWatchOptions) (*CollectionPolicyList, error) {
	resp, err := e.AutoListCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return &CollectionPolicyList{}, err
	}
	return resp.(*CollectionPolicyList), nil
}

type respCollectionPolicyV1AutoListCollectionPolicy struct {
	V   CollectionPolicyList
	Err error
}

// AutoUpdateCollectionPolicy is endpoint for AutoUpdateCollectionPolicy
func (e EndpointsCollectionPolicyV1Client) AutoUpdateCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	resp, err := e.AutoUpdateCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return &CollectionPolicy{}, err
	}
	return resp.(*CollectionPolicy), nil
}

type respCollectionPolicyV1AutoUpdateCollectionPolicy struct {
	V   CollectionPolicy
	Err error
}

// AutoWatchCollectionPolicy performs Watch for CollectionPolicy
func (e EndpointsCollectionPolicyV1Client) AutoWatchCollectionPolicy(ctx context.Context, in *api.ListWatchOptions) (CollectionPolicyV1_AutoWatchCollectionPolicyClient, error) {
	return e.Client.AutoWatchCollectionPolicy(ctx, in)
}

// AutoAddCollectionPolicy implementation on server Endpoint
func (e EndpointsCollectionPolicyV1Server) AutoAddCollectionPolicy(ctx context.Context, in CollectionPolicy) (CollectionPolicy, error) {
	resp, err := e.AutoAddCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return CollectionPolicy{}, err
	}
	return *resp.(*CollectionPolicy), nil
}

// MakeCollectionPolicyV1AutoAddCollectionPolicyEndpoint creates  AutoAddCollectionPolicy endpoints for the service
func MakeCollectionPolicyV1AutoAddCollectionPolicyEndpoint(s ServiceCollectionPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CollectionPolicy)
		v, err := s.AutoAddCollectionPolicy(ctx, *req)
		return respCollectionPolicyV1AutoAddCollectionPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("CollectionPolicyV1:AutoAddCollectionPolicy")(f)
}

// AutoDeleteCollectionPolicy implementation on server Endpoint
func (e EndpointsCollectionPolicyV1Server) AutoDeleteCollectionPolicy(ctx context.Context, in CollectionPolicy) (CollectionPolicy, error) {
	resp, err := e.AutoDeleteCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return CollectionPolicy{}, err
	}
	return *resp.(*CollectionPolicy), nil
}

// MakeCollectionPolicyV1AutoDeleteCollectionPolicyEndpoint creates  AutoDeleteCollectionPolicy endpoints for the service
func MakeCollectionPolicyV1AutoDeleteCollectionPolicyEndpoint(s ServiceCollectionPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CollectionPolicy)
		v, err := s.AutoDeleteCollectionPolicy(ctx, *req)
		return respCollectionPolicyV1AutoDeleteCollectionPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("CollectionPolicyV1:AutoDeleteCollectionPolicy")(f)
}

// AutoGetCollectionPolicy implementation on server Endpoint
func (e EndpointsCollectionPolicyV1Server) AutoGetCollectionPolicy(ctx context.Context, in CollectionPolicy) (CollectionPolicy, error) {
	resp, err := e.AutoGetCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return CollectionPolicy{}, err
	}
	return *resp.(*CollectionPolicy), nil
}

// MakeCollectionPolicyV1AutoGetCollectionPolicyEndpoint creates  AutoGetCollectionPolicy endpoints for the service
func MakeCollectionPolicyV1AutoGetCollectionPolicyEndpoint(s ServiceCollectionPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CollectionPolicy)
		v, err := s.AutoGetCollectionPolicy(ctx, *req)
		return respCollectionPolicyV1AutoGetCollectionPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("CollectionPolicyV1:AutoGetCollectionPolicy")(f)
}

// AutoListCollectionPolicy implementation on server Endpoint
func (e EndpointsCollectionPolicyV1Server) AutoListCollectionPolicy(ctx context.Context, in api.ListWatchOptions) (CollectionPolicyList, error) {
	resp, err := e.AutoListCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return CollectionPolicyList{}, err
	}
	return *resp.(*CollectionPolicyList), nil
}

// MakeCollectionPolicyV1AutoListCollectionPolicyEndpoint creates  AutoListCollectionPolicy endpoints for the service
func MakeCollectionPolicyV1AutoListCollectionPolicyEndpoint(s ServiceCollectionPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.ListWatchOptions)
		v, err := s.AutoListCollectionPolicy(ctx, *req)
		return respCollectionPolicyV1AutoListCollectionPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("CollectionPolicyV1:AutoListCollectionPolicy")(f)
}

// AutoUpdateCollectionPolicy implementation on server Endpoint
func (e EndpointsCollectionPolicyV1Server) AutoUpdateCollectionPolicy(ctx context.Context, in CollectionPolicy) (CollectionPolicy, error) {
	resp, err := e.AutoUpdateCollectionPolicyEndpoint(ctx, in)
	if err != nil {
		return CollectionPolicy{}, err
	}
	return *resp.(*CollectionPolicy), nil
}

// MakeCollectionPolicyV1AutoUpdateCollectionPolicyEndpoint creates  AutoUpdateCollectionPolicy endpoints for the service
func MakeCollectionPolicyV1AutoUpdateCollectionPolicyEndpoint(s ServiceCollectionPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*CollectionPolicy)
		v, err := s.AutoUpdateCollectionPolicy(ctx, *req)
		return respCollectionPolicyV1AutoUpdateCollectionPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("CollectionPolicyV1:AutoUpdateCollectionPolicy")(f)
}

// AutoWatchCollectionPolicy is the watch handler for CollectionPolicy on the server side.
func (e EndpointsCollectionPolicyV1Server) AutoWatchCollectionPolicy(in *api.ListWatchOptions, stream CollectionPolicyV1_AutoWatchCollectionPolicyServer) error {
	return e.watchHandlerCollectionPolicy(in, stream)
}

// MakeAutoWatchCollectionPolicyEndpoint creates the Watch endpoint
func MakeAutoWatchCollectionPolicyEndpoint(s ServiceCollectionPolicyV1Server, logger log.Logger) func(options *api.ListWatchOptions, stream grpc.ServerStream) error {
	return func(options *api.ListWatchOptions, stream grpc.ServerStream) error {
		wstream := stream.(CollectionPolicyV1_AutoWatchCollectionPolicyServer)
		return s.AutoWatchCollectionPolicy(options, wstream)
	}
}

// MakeCollectionPolicyV1ServerEndpoints creates server endpoints
func MakeCollectionPolicyV1ServerEndpoints(s ServiceCollectionPolicyV1Server, logger log.Logger) EndpointsCollectionPolicyV1Server {
	return EndpointsCollectionPolicyV1Server{

		AutoAddCollectionPolicyEndpoint:    MakeCollectionPolicyV1AutoAddCollectionPolicyEndpoint(s, logger),
		AutoDeleteCollectionPolicyEndpoint: MakeCollectionPolicyV1AutoDeleteCollectionPolicyEndpoint(s, logger),
		AutoGetCollectionPolicyEndpoint:    MakeCollectionPolicyV1AutoGetCollectionPolicyEndpoint(s, logger),
		AutoListCollectionPolicyEndpoint:   MakeCollectionPolicyV1AutoListCollectionPolicyEndpoint(s, logger),
		AutoUpdateCollectionPolicyEndpoint: MakeCollectionPolicyV1AutoUpdateCollectionPolicyEndpoint(s, logger),

		watchHandlerCollectionPolicy: MakeAutoWatchCollectionPolicyEndpoint(s, logger),
	}
}

// LoggingCollectionPolicyV1MiddlewareClient adds middleware for the client
func LoggingCollectionPolicyV1MiddlewareClient(logger log.Logger) MiddlewareCollectionPolicyV1Client {
	return func(next ServiceCollectionPolicyV1Client) ServiceCollectionPolicyV1Client {
		return loggingCollectionPolicyV1MiddlewareClient{
			logger: logger,
			next:   next,
		}
	}
}

type loggingCollectionPolicyV1MiddlewareClient struct {
	logger log.Logger
	next   ServiceCollectionPolicyV1Client
}

// LoggingCollectionPolicyV1MiddlewareServer adds middleware for the client
func LoggingCollectionPolicyV1MiddlewareServer(logger log.Logger) MiddlewareCollectionPolicyV1Server {
	return func(next ServiceCollectionPolicyV1Server) ServiceCollectionPolicyV1Server {
		return loggingCollectionPolicyV1MiddlewareServer{
			logger: logger,
			next:   next,
		}
	}
}

type loggingCollectionPolicyV1MiddlewareServer struct {
	logger log.Logger
	next   ServiceCollectionPolicyV1Server
}

func (m loggingCollectionPolicyV1MiddlewareClient) AutoAddCollectionPolicy(ctx context.Context, in *CollectionPolicy) (resp *CollectionPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoAddCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoAddCollectionPolicy(ctx, in)
	return
}
func (m loggingCollectionPolicyV1MiddlewareClient) AutoDeleteCollectionPolicy(ctx context.Context, in *CollectionPolicy) (resp *CollectionPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoDeleteCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoDeleteCollectionPolicy(ctx, in)
	return
}
func (m loggingCollectionPolicyV1MiddlewareClient) AutoGetCollectionPolicy(ctx context.Context, in *CollectionPolicy) (resp *CollectionPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoGetCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoGetCollectionPolicy(ctx, in)
	return
}
func (m loggingCollectionPolicyV1MiddlewareClient) AutoListCollectionPolicy(ctx context.Context, in *api.ListWatchOptions) (resp *CollectionPolicyList, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoListCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoListCollectionPolicy(ctx, in)
	return
}
func (m loggingCollectionPolicyV1MiddlewareClient) AutoUpdateCollectionPolicy(ctx context.Context, in *CollectionPolicy) (resp *CollectionPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoUpdateCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoUpdateCollectionPolicy(ctx, in)
	return
}

func (m loggingCollectionPolicyV1MiddlewareClient) AutoWatchCollectionPolicy(ctx context.Context, in *api.ListWatchOptions) (resp CollectionPolicyV1_AutoWatchCollectionPolicyClient, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoWatchCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoWatchCollectionPolicy(ctx, in)
	return
}

func (m loggingCollectionPolicyV1MiddlewareServer) AutoAddCollectionPolicy(ctx context.Context, in CollectionPolicy) (resp CollectionPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoAddCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoAddCollectionPolicy(ctx, in)
	return
}
func (m loggingCollectionPolicyV1MiddlewareServer) AutoDeleteCollectionPolicy(ctx context.Context, in CollectionPolicy) (resp CollectionPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoDeleteCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoDeleteCollectionPolicy(ctx, in)
	return
}
func (m loggingCollectionPolicyV1MiddlewareServer) AutoGetCollectionPolicy(ctx context.Context, in CollectionPolicy) (resp CollectionPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoGetCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoGetCollectionPolicy(ctx, in)
	return
}
func (m loggingCollectionPolicyV1MiddlewareServer) AutoListCollectionPolicy(ctx context.Context, in api.ListWatchOptions) (resp CollectionPolicyList, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoListCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoListCollectionPolicy(ctx, in)
	return
}
func (m loggingCollectionPolicyV1MiddlewareServer) AutoUpdateCollectionPolicy(ctx context.Context, in CollectionPolicy) (resp CollectionPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "CollectionPolicyV1", "method", "AutoUpdateCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoUpdateCollectionPolicy(ctx, in)
	return
}

func (m loggingCollectionPolicyV1MiddlewareServer) AutoWatchCollectionPolicy(in *api.ListWatchOptions, stream CollectionPolicyV1_AutoWatchCollectionPolicyServer) (err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(stream.Context(), "service", "CollectionPolicyV1", "method", "AutoWatchCollectionPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	err = m.next.AutoWatchCollectionPolicy(in, stream)
	return
}
func (r *EndpointsCollectionPolicyV1RestClient) getHTTPRequest(ctx context.Context, in interface{}, method, path string) (*http.Request, error) {
	target, err := url.Parse(r.instance)
	if err != nil {
		return nil, fmt.Errorf("invalid instance %s", r.instance)
	}
	target.Path = path
	req, err := http.NewRequest(method, target.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request (%s)", err)
	}
	if err = encodeHTTPRequest(ctx, req, in); err != nil {
		return nil, fmt.Errorf("could not encode request (%s)", err)
	}
	return req, nil
}

//
func makeURICollectionPolicyV1AutoAddCollectionPolicyCreateOper(in *CollectionPolicy) string {
	return fmt.Sprint("/v1/collectionPolicy", "/", in.Tenant, "/collectionPolicy")
}

//
func makeURICollectionPolicyV1AutoDeleteCollectionPolicyDeleteOper(in *CollectionPolicy) string {
	return fmt.Sprint("/v1/collectionPolicy", "/", in.Tenant, "/collectionPolicy/", in.Name)
}

//
func makeURICollectionPolicyV1AutoGetCollectionPolicyGetOper(in *CollectionPolicy) string {
	return fmt.Sprint("/v1/collectionPolicy", "/", in.Tenant, "/collectionPolicy/", in.Name)
}

//
func makeURICollectionPolicyV1AutoUpdateCollectionPolicyUpdateOper(in *CollectionPolicy) string {
	return fmt.Sprint("/v1/collectionPolicy", "/", in.Tenant, "/collectionPolicy/", in.Name)
}

// AutoAddCollectionPolicy CRUD method for CollectionPolicy
func (r *EndpointsCollectionPolicyV1RestClient) AutoAddCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	path := makeURICollectionPolicyV1AutoAddCollectionPolicyCreateOper(in)
	req, err := r.getHTTPRequest(ctx, in, "POST", path)
	if err != nil {
		return nil, err
	}
	httpresp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespCollectionPolicyV1AutoAddCollectionPolicy(ctx, httpresp)
	if err != nil {
		return nil, err
	}
	return ret.(*CollectionPolicy), nil
}

// AutoUpdateCollectionPolicy CRUD method for CollectionPolicy
func (r *EndpointsCollectionPolicyV1RestClient) AutoUpdateCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	path := makeURICollectionPolicyV1AutoUpdateCollectionPolicyUpdateOper(in)
	req, err := r.getHTTPRequest(ctx, in, "PUT", path)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespCollectionPolicyV1AutoUpdateCollectionPolicy(ctx, resp)
	if err != nil {
		return nil, err
	}
	return ret.(*CollectionPolicy), err
}

// AutoGetCollectionPolicy CRUD method for CollectionPolicy
func (r *EndpointsCollectionPolicyV1RestClient) AutoGetCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	path := makeURICollectionPolicyV1AutoGetCollectionPolicyGetOper(in)
	req, err := r.getHTTPRequest(ctx, in, "GET", path)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespCollectionPolicyV1AutoGetCollectionPolicy(ctx, resp)
	if err != nil {
		return nil, err
	}
	return ret.(*CollectionPolicy), err
}

// AutoDeleteCollectionPolicy CRUD method for CollectionPolicy
func (r *EndpointsCollectionPolicyV1RestClient) AutoDeleteCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	path := makeURICollectionPolicyV1AutoDeleteCollectionPolicyDeleteOper(in)
	req, err := r.getHTTPRequest(ctx, in, "DELETE", path)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespCollectionPolicyV1AutoDeleteCollectionPolicy(ctx, resp)
	if err != nil {
		return nil, err
	}
	return ret.(*CollectionPolicy), err
}

// AutoListCollectionPolicy CRUD method for CollectionPolicy
func (r *EndpointsCollectionPolicyV1RestClient) AutoListCollectionPolicy(ctx context.Context, options *api.ListWatchOptions) (*CollectionPolicyList, error) {
	return nil, errors.New("not allowed")
}

// AutoWatchCollectionPolicy CRUD method for CollectionPolicy
func (r *EndpointsCollectionPolicyV1RestClient) AutoWatchCollectionPolicy(ctx context.Context, in *CollectionPolicy) (*CollectionPolicy, error) {
	return nil, errors.New("not allowed")
}

// MakeCollectionPolicyV1RestClientEndpoints make REST client endpoints
func MakeCollectionPolicyV1RestClientEndpoints(instance string) (EndpointsCollectionPolicyV1RestClient, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	return EndpointsCollectionPolicyV1RestClient{
		instance: instance,
		client:   http.DefaultClient,
	}, nil

}
