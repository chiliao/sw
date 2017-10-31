// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package events is a auto generated package.
Input file: protos/events.proto
*/
package events

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

// MiddlewareEventPolicyV1Client add middleware to the client
type MiddlewareEventPolicyV1Client func(ServiceEventPolicyV1Client) ServiceEventPolicyV1Client

// EndpointsEventPolicyV1Client is the endpoints for the client
type EndpointsEventPolicyV1Client struct {
	Client EventPolicyV1Client

	AutoAddEventPolicyEndpoint    endpoint.Endpoint
	AutoDeleteEventPolicyEndpoint endpoint.Endpoint
	AutoGetEventPolicyEndpoint    endpoint.Endpoint
	AutoListEventPolicyEndpoint   endpoint.Endpoint
	AutoUpdateEventPolicyEndpoint endpoint.Endpoint
}

// EndpointsEventPolicyV1RestClient is the REST client
type EndpointsEventPolicyV1RestClient struct {
	logger   log.Logger
	client   *http.Client
	instance string

	AutoAddEventPolicyEndpoint    endpoint.Endpoint
	AutoDeleteEventPolicyEndpoint endpoint.Endpoint
	AutoGetEventPolicyEndpoint    endpoint.Endpoint
	AutoListEventPolicyEndpoint   endpoint.Endpoint
	AutoUpdateEventPolicyEndpoint endpoint.Endpoint
	AutoWatchEventPolicyEndpoint  endpoint.Endpoint
}

// MiddlewareEventPolicyV1Server adds middle ware to the server
type MiddlewareEventPolicyV1Server func(ServiceEventPolicyV1Server) ServiceEventPolicyV1Server

// EndpointsEventPolicyV1Server is the server endpoints
type EndpointsEventPolicyV1Server struct {
	AutoAddEventPolicyEndpoint    endpoint.Endpoint
	AutoDeleteEventPolicyEndpoint endpoint.Endpoint
	AutoGetEventPolicyEndpoint    endpoint.Endpoint
	AutoListEventPolicyEndpoint   endpoint.Endpoint
	AutoUpdateEventPolicyEndpoint endpoint.Endpoint

	watchHandlerEventPolicy func(options *api.ListWatchOptions, stream grpc.ServerStream) error
}

// AutoAddEventPolicy is endpoint for AutoAddEventPolicy
func (e EndpointsEventPolicyV1Client) AutoAddEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	resp, err := e.AutoAddEventPolicyEndpoint(ctx, in)
	if err != nil {
		return &EventPolicy{}, err
	}
	return resp.(*EventPolicy), nil
}

type respEventPolicyV1AutoAddEventPolicy struct {
	V   EventPolicy
	Err error
}

// AutoDeleteEventPolicy is endpoint for AutoDeleteEventPolicy
func (e EndpointsEventPolicyV1Client) AutoDeleteEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	resp, err := e.AutoDeleteEventPolicyEndpoint(ctx, in)
	if err != nil {
		return &EventPolicy{}, err
	}
	return resp.(*EventPolicy), nil
}

type respEventPolicyV1AutoDeleteEventPolicy struct {
	V   EventPolicy
	Err error
}

// AutoGetEventPolicy is endpoint for AutoGetEventPolicy
func (e EndpointsEventPolicyV1Client) AutoGetEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	resp, err := e.AutoGetEventPolicyEndpoint(ctx, in)
	if err != nil {
		return &EventPolicy{}, err
	}
	return resp.(*EventPolicy), nil
}

type respEventPolicyV1AutoGetEventPolicy struct {
	V   EventPolicy
	Err error
}

// AutoListEventPolicy is endpoint for AutoListEventPolicy
func (e EndpointsEventPolicyV1Client) AutoListEventPolicy(ctx context.Context, in *api.ListWatchOptions) (*EventPolicyList, error) {
	resp, err := e.AutoListEventPolicyEndpoint(ctx, in)
	if err != nil {
		return &EventPolicyList{}, err
	}
	return resp.(*EventPolicyList), nil
}

type respEventPolicyV1AutoListEventPolicy struct {
	V   EventPolicyList
	Err error
}

// AutoUpdateEventPolicy is endpoint for AutoUpdateEventPolicy
func (e EndpointsEventPolicyV1Client) AutoUpdateEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	resp, err := e.AutoUpdateEventPolicyEndpoint(ctx, in)
	if err != nil {
		return &EventPolicy{}, err
	}
	return resp.(*EventPolicy), nil
}

type respEventPolicyV1AutoUpdateEventPolicy struct {
	V   EventPolicy
	Err error
}

// AutoWatchEventPolicy performs Watch for EventPolicy
func (e EndpointsEventPolicyV1Client) AutoWatchEventPolicy(ctx context.Context, in *api.ListWatchOptions) (EventPolicyV1_AutoWatchEventPolicyClient, error) {
	return e.Client.AutoWatchEventPolicy(ctx, in)
}

// AutoAddEventPolicy implementation on server Endpoint
func (e EndpointsEventPolicyV1Server) AutoAddEventPolicy(ctx context.Context, in EventPolicy) (EventPolicy, error) {
	resp, err := e.AutoAddEventPolicyEndpoint(ctx, in)
	if err != nil {
		return EventPolicy{}, err
	}
	return *resp.(*EventPolicy), nil
}

// MakeEventPolicyV1AutoAddEventPolicyEndpoint creates  AutoAddEventPolicy endpoints for the service
func MakeEventPolicyV1AutoAddEventPolicyEndpoint(s ServiceEventPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*EventPolicy)
		v, err := s.AutoAddEventPolicy(ctx, *req)
		return respEventPolicyV1AutoAddEventPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventPolicyV1:AutoAddEventPolicy")(f)
}

// AutoDeleteEventPolicy implementation on server Endpoint
func (e EndpointsEventPolicyV1Server) AutoDeleteEventPolicy(ctx context.Context, in EventPolicy) (EventPolicy, error) {
	resp, err := e.AutoDeleteEventPolicyEndpoint(ctx, in)
	if err != nil {
		return EventPolicy{}, err
	}
	return *resp.(*EventPolicy), nil
}

// MakeEventPolicyV1AutoDeleteEventPolicyEndpoint creates  AutoDeleteEventPolicy endpoints for the service
func MakeEventPolicyV1AutoDeleteEventPolicyEndpoint(s ServiceEventPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*EventPolicy)
		v, err := s.AutoDeleteEventPolicy(ctx, *req)
		return respEventPolicyV1AutoDeleteEventPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventPolicyV1:AutoDeleteEventPolicy")(f)
}

// AutoGetEventPolicy implementation on server Endpoint
func (e EndpointsEventPolicyV1Server) AutoGetEventPolicy(ctx context.Context, in EventPolicy) (EventPolicy, error) {
	resp, err := e.AutoGetEventPolicyEndpoint(ctx, in)
	if err != nil {
		return EventPolicy{}, err
	}
	return *resp.(*EventPolicy), nil
}

// MakeEventPolicyV1AutoGetEventPolicyEndpoint creates  AutoGetEventPolicy endpoints for the service
func MakeEventPolicyV1AutoGetEventPolicyEndpoint(s ServiceEventPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*EventPolicy)
		v, err := s.AutoGetEventPolicy(ctx, *req)
		return respEventPolicyV1AutoGetEventPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventPolicyV1:AutoGetEventPolicy")(f)
}

// AutoListEventPolicy implementation on server Endpoint
func (e EndpointsEventPolicyV1Server) AutoListEventPolicy(ctx context.Context, in api.ListWatchOptions) (EventPolicyList, error) {
	resp, err := e.AutoListEventPolicyEndpoint(ctx, in)
	if err != nil {
		return EventPolicyList{}, err
	}
	return *resp.(*EventPolicyList), nil
}

// MakeEventPolicyV1AutoListEventPolicyEndpoint creates  AutoListEventPolicy endpoints for the service
func MakeEventPolicyV1AutoListEventPolicyEndpoint(s ServiceEventPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.ListWatchOptions)
		v, err := s.AutoListEventPolicy(ctx, *req)
		return respEventPolicyV1AutoListEventPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventPolicyV1:AutoListEventPolicy")(f)
}

// AutoUpdateEventPolicy implementation on server Endpoint
func (e EndpointsEventPolicyV1Server) AutoUpdateEventPolicy(ctx context.Context, in EventPolicy) (EventPolicy, error) {
	resp, err := e.AutoUpdateEventPolicyEndpoint(ctx, in)
	if err != nil {
		return EventPolicy{}, err
	}
	return *resp.(*EventPolicy), nil
}

// MakeEventPolicyV1AutoUpdateEventPolicyEndpoint creates  AutoUpdateEventPolicy endpoints for the service
func MakeEventPolicyV1AutoUpdateEventPolicyEndpoint(s ServiceEventPolicyV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*EventPolicy)
		v, err := s.AutoUpdateEventPolicy(ctx, *req)
		return respEventPolicyV1AutoUpdateEventPolicy{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventPolicyV1:AutoUpdateEventPolicy")(f)
}

// AutoWatchEventPolicy is the watch handler for EventPolicy on the server side.
func (e EndpointsEventPolicyV1Server) AutoWatchEventPolicy(in *api.ListWatchOptions, stream EventPolicyV1_AutoWatchEventPolicyServer) error {
	return e.watchHandlerEventPolicy(in, stream)
}

// MakeAutoWatchEventPolicyEndpoint creates the Watch endpoint
func MakeAutoWatchEventPolicyEndpoint(s ServiceEventPolicyV1Server, logger log.Logger) func(options *api.ListWatchOptions, stream grpc.ServerStream) error {
	return func(options *api.ListWatchOptions, stream grpc.ServerStream) error {
		wstream := stream.(EventPolicyV1_AutoWatchEventPolicyServer)
		return s.AutoWatchEventPolicy(options, wstream)
	}
}

// MakeEventPolicyV1ServerEndpoints creates server endpoints
func MakeEventPolicyV1ServerEndpoints(s ServiceEventPolicyV1Server, logger log.Logger) EndpointsEventPolicyV1Server {
	return EndpointsEventPolicyV1Server{

		AutoAddEventPolicyEndpoint:    MakeEventPolicyV1AutoAddEventPolicyEndpoint(s, logger),
		AutoDeleteEventPolicyEndpoint: MakeEventPolicyV1AutoDeleteEventPolicyEndpoint(s, logger),
		AutoGetEventPolicyEndpoint:    MakeEventPolicyV1AutoGetEventPolicyEndpoint(s, logger),
		AutoListEventPolicyEndpoint:   MakeEventPolicyV1AutoListEventPolicyEndpoint(s, logger),
		AutoUpdateEventPolicyEndpoint: MakeEventPolicyV1AutoUpdateEventPolicyEndpoint(s, logger),

		watchHandlerEventPolicy: MakeAutoWatchEventPolicyEndpoint(s, logger),
	}
}

// LoggingEventPolicyV1MiddlewareClient adds middleware for the client
func LoggingEventPolicyV1MiddlewareClient(logger log.Logger) MiddlewareEventPolicyV1Client {
	return func(next ServiceEventPolicyV1Client) ServiceEventPolicyV1Client {
		return loggingEventPolicyV1MiddlewareClient{
			logger: logger,
			next:   next,
		}
	}
}

type loggingEventPolicyV1MiddlewareClient struct {
	logger log.Logger
	next   ServiceEventPolicyV1Client
}

// LoggingEventPolicyV1MiddlewareServer adds middleware for the client
func LoggingEventPolicyV1MiddlewareServer(logger log.Logger) MiddlewareEventPolicyV1Server {
	return func(next ServiceEventPolicyV1Server) ServiceEventPolicyV1Server {
		return loggingEventPolicyV1MiddlewareServer{
			logger: logger,
			next:   next,
		}
	}
}

type loggingEventPolicyV1MiddlewareServer struct {
	logger log.Logger
	next   ServiceEventPolicyV1Server
}

func (m loggingEventPolicyV1MiddlewareClient) AutoAddEventPolicy(ctx context.Context, in *EventPolicy) (resp *EventPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoAddEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoAddEventPolicy(ctx, in)
	return
}
func (m loggingEventPolicyV1MiddlewareClient) AutoDeleteEventPolicy(ctx context.Context, in *EventPolicy) (resp *EventPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoDeleteEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoDeleteEventPolicy(ctx, in)
	return
}
func (m loggingEventPolicyV1MiddlewareClient) AutoGetEventPolicy(ctx context.Context, in *EventPolicy) (resp *EventPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoGetEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoGetEventPolicy(ctx, in)
	return
}
func (m loggingEventPolicyV1MiddlewareClient) AutoListEventPolicy(ctx context.Context, in *api.ListWatchOptions) (resp *EventPolicyList, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoListEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoListEventPolicy(ctx, in)
	return
}
func (m loggingEventPolicyV1MiddlewareClient) AutoUpdateEventPolicy(ctx context.Context, in *EventPolicy) (resp *EventPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoUpdateEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoUpdateEventPolicy(ctx, in)
	return
}

func (m loggingEventPolicyV1MiddlewareClient) AutoWatchEventPolicy(ctx context.Context, in *api.ListWatchOptions) (resp EventPolicyV1_AutoWatchEventPolicyClient, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoWatchEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoWatchEventPolicy(ctx, in)
	return
}

func (m loggingEventPolicyV1MiddlewareServer) AutoAddEventPolicy(ctx context.Context, in EventPolicy) (resp EventPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoAddEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoAddEventPolicy(ctx, in)
	return
}
func (m loggingEventPolicyV1MiddlewareServer) AutoDeleteEventPolicy(ctx context.Context, in EventPolicy) (resp EventPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoDeleteEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoDeleteEventPolicy(ctx, in)
	return
}
func (m loggingEventPolicyV1MiddlewareServer) AutoGetEventPolicy(ctx context.Context, in EventPolicy) (resp EventPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoGetEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoGetEventPolicy(ctx, in)
	return
}
func (m loggingEventPolicyV1MiddlewareServer) AutoListEventPolicy(ctx context.Context, in api.ListWatchOptions) (resp EventPolicyList, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoListEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoListEventPolicy(ctx, in)
	return
}
func (m loggingEventPolicyV1MiddlewareServer) AutoUpdateEventPolicy(ctx context.Context, in EventPolicy) (resp EventPolicy, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventPolicyV1", "method", "AutoUpdateEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoUpdateEventPolicy(ctx, in)
	return
}

func (m loggingEventPolicyV1MiddlewareServer) AutoWatchEventPolicy(in *api.ListWatchOptions, stream EventPolicyV1_AutoWatchEventPolicyServer) (err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(stream.Context(), "service", "EventPolicyV1", "method", "AutoWatchEventPolicy", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	err = m.next.AutoWatchEventPolicy(in, stream)
	return
}
func (r *EndpointsEventPolicyV1RestClient) getHTTPRequest(ctx context.Context, in interface{}, method, path string) (*http.Request, error) {
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
func makeURIEventPolicyV1AutoAddEventPolicyCreateOper(in *EventPolicy) string {
	return fmt.Sprint("/v1/eventPolicy", "/", in.Tenant, "/eventPolicy")
}

//
func makeURIEventPolicyV1AutoDeleteEventPolicyDeleteOper(in *EventPolicy) string {
	return fmt.Sprint("/v1/eventPolicy", "/", in.Tenant, "/eventPolicy/", in.Name)
}

//
func makeURIEventPolicyV1AutoGetEventPolicyGetOper(in *EventPolicy) string {
	return fmt.Sprint("/v1/eventPolicy", "/", in.Tenant, "/eventPolicy/", in.Name)
}

//
func makeURIEventPolicyV1AutoUpdateEventPolicyUpdateOper(in *EventPolicy) string {
	return fmt.Sprint("/v1/eventPolicy", "/", in.Tenant, "/eventPolicy/", in.Name)
}

// AutoAddEventPolicy CRUD method for EventPolicy
func (r *EndpointsEventPolicyV1RestClient) AutoAddEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	path := makeURIEventPolicyV1AutoAddEventPolicyCreateOper(in)
	req, err := r.getHTTPRequest(ctx, in, "POST", path)
	if err != nil {
		return nil, err
	}
	httpresp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespEventPolicyV1AutoAddEventPolicy(ctx, httpresp)
	if err != nil {
		return nil, err
	}
	return ret.(*EventPolicy), nil
}

// AutoUpdateEventPolicy CRUD method for EventPolicy
func (r *EndpointsEventPolicyV1RestClient) AutoUpdateEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	path := makeURIEventPolicyV1AutoUpdateEventPolicyUpdateOper(in)
	req, err := r.getHTTPRequest(ctx, in, "PUT", path)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespEventPolicyV1AutoUpdateEventPolicy(ctx, resp)
	if err != nil {
		return nil, err
	}
	return ret.(*EventPolicy), err
}

// AutoGetEventPolicy CRUD method for EventPolicy
func (r *EndpointsEventPolicyV1RestClient) AutoGetEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	path := makeURIEventPolicyV1AutoGetEventPolicyGetOper(in)
	req, err := r.getHTTPRequest(ctx, in, "GET", path)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespEventPolicyV1AutoGetEventPolicy(ctx, resp)
	if err != nil {
		return nil, err
	}
	return ret.(*EventPolicy), err
}

// AutoDeleteEventPolicy CRUD method for EventPolicy
func (r *EndpointsEventPolicyV1RestClient) AutoDeleteEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	path := makeURIEventPolicyV1AutoDeleteEventPolicyDeleteOper(in)
	req, err := r.getHTTPRequest(ctx, in, "DELETE", path)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespEventPolicyV1AutoDeleteEventPolicy(ctx, resp)
	if err != nil {
		return nil, err
	}
	return ret.(*EventPolicy), err
}

// AutoListEventPolicy CRUD method for EventPolicy
func (r *EndpointsEventPolicyV1RestClient) AutoListEventPolicy(ctx context.Context, options *api.ListWatchOptions) (*EventPolicyList, error) {
	return nil, errors.New("not allowed")
}

// AutoWatchEventPolicy CRUD method for EventPolicy
func (r *EndpointsEventPolicyV1RestClient) AutoWatchEventPolicy(ctx context.Context, in *EventPolicy) (*EventPolicy, error) {
	return nil, errors.New("not allowed")
}

// MakeEventPolicyV1RestClientEndpoints make REST client endpoints
func MakeEventPolicyV1RestClientEndpoints(instance string) (EndpointsEventPolicyV1RestClient, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	return EndpointsEventPolicyV1RestClient{
		instance: instance,
		client:   http.DefaultClient,
	}, nil

}

// MiddlewareEventV1Client add middleware to the client
type MiddlewareEventV1Client func(ServiceEventV1Client) ServiceEventV1Client

// EndpointsEventV1Client is the endpoints for the client
type EndpointsEventV1Client struct {
	Client EventV1Client

	AutoAddEventEndpoint    endpoint.Endpoint
	AutoDeleteEventEndpoint endpoint.Endpoint
	AutoGetEventEndpoint    endpoint.Endpoint
	AutoListEventEndpoint   endpoint.Endpoint
	AutoUpdateEventEndpoint endpoint.Endpoint
}

// EndpointsEventV1RestClient is the REST client
type EndpointsEventV1RestClient struct {
	logger   log.Logger
	client   *http.Client
	instance string

	AutoAddEventEndpoint    endpoint.Endpoint
	AutoDeleteEventEndpoint endpoint.Endpoint
	AutoGetEventEndpoint    endpoint.Endpoint
	AutoListEventEndpoint   endpoint.Endpoint
	AutoUpdateEventEndpoint endpoint.Endpoint
	AutoWatchEventEndpoint  endpoint.Endpoint
}

// MiddlewareEventV1Server adds middle ware to the server
type MiddlewareEventV1Server func(ServiceEventV1Server) ServiceEventV1Server

// EndpointsEventV1Server is the server endpoints
type EndpointsEventV1Server struct {
	AutoAddEventEndpoint    endpoint.Endpoint
	AutoDeleteEventEndpoint endpoint.Endpoint
	AutoGetEventEndpoint    endpoint.Endpoint
	AutoListEventEndpoint   endpoint.Endpoint
	AutoUpdateEventEndpoint endpoint.Endpoint

	watchHandlerEvent func(options *api.ListWatchOptions, stream grpc.ServerStream) error
}

// AutoAddEvent is endpoint for AutoAddEvent
func (e EndpointsEventV1Client) AutoAddEvent(ctx context.Context, in *Event) (*Event, error) {
	resp, err := e.AutoAddEventEndpoint(ctx, in)
	if err != nil {
		return &Event{}, err
	}
	return resp.(*Event), nil
}

type respEventV1AutoAddEvent struct {
	V   Event
	Err error
}

// AutoDeleteEvent is endpoint for AutoDeleteEvent
func (e EndpointsEventV1Client) AutoDeleteEvent(ctx context.Context, in *Event) (*Event, error) {
	resp, err := e.AutoDeleteEventEndpoint(ctx, in)
	if err != nil {
		return &Event{}, err
	}
	return resp.(*Event), nil
}

type respEventV1AutoDeleteEvent struct {
	V   Event
	Err error
}

// AutoGetEvent is endpoint for AutoGetEvent
func (e EndpointsEventV1Client) AutoGetEvent(ctx context.Context, in *Event) (*Event, error) {
	resp, err := e.AutoGetEventEndpoint(ctx, in)
	if err != nil {
		return &Event{}, err
	}
	return resp.(*Event), nil
}

type respEventV1AutoGetEvent struct {
	V   Event
	Err error
}

// AutoListEvent is endpoint for AutoListEvent
func (e EndpointsEventV1Client) AutoListEvent(ctx context.Context, in *api.ListWatchOptions) (*EventList, error) {
	resp, err := e.AutoListEventEndpoint(ctx, in)
	if err != nil {
		return &EventList{}, err
	}
	return resp.(*EventList), nil
}

type respEventV1AutoListEvent struct {
	V   EventList
	Err error
}

// AutoUpdateEvent is endpoint for AutoUpdateEvent
func (e EndpointsEventV1Client) AutoUpdateEvent(ctx context.Context, in *Event) (*Event, error) {
	resp, err := e.AutoUpdateEventEndpoint(ctx, in)
	if err != nil {
		return &Event{}, err
	}
	return resp.(*Event), nil
}

type respEventV1AutoUpdateEvent struct {
	V   Event
	Err error
}

// AutoWatchEvent performs Watch for Event
func (e EndpointsEventV1Client) AutoWatchEvent(ctx context.Context, in *api.ListWatchOptions) (EventV1_AutoWatchEventClient, error) {
	return e.Client.AutoWatchEvent(ctx, in)
}

// AutoAddEvent implementation on server Endpoint
func (e EndpointsEventV1Server) AutoAddEvent(ctx context.Context, in Event) (Event, error) {
	resp, err := e.AutoAddEventEndpoint(ctx, in)
	if err != nil {
		return Event{}, err
	}
	return *resp.(*Event), nil
}

// MakeEventV1AutoAddEventEndpoint creates  AutoAddEvent endpoints for the service
func MakeEventV1AutoAddEventEndpoint(s ServiceEventV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*Event)
		v, err := s.AutoAddEvent(ctx, *req)
		return respEventV1AutoAddEvent{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventV1:AutoAddEvent")(f)
}

// AutoDeleteEvent implementation on server Endpoint
func (e EndpointsEventV1Server) AutoDeleteEvent(ctx context.Context, in Event) (Event, error) {
	resp, err := e.AutoDeleteEventEndpoint(ctx, in)
	if err != nil {
		return Event{}, err
	}
	return *resp.(*Event), nil
}

// MakeEventV1AutoDeleteEventEndpoint creates  AutoDeleteEvent endpoints for the service
func MakeEventV1AutoDeleteEventEndpoint(s ServiceEventV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*Event)
		v, err := s.AutoDeleteEvent(ctx, *req)
		return respEventV1AutoDeleteEvent{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventV1:AutoDeleteEvent")(f)
}

// AutoGetEvent implementation on server Endpoint
func (e EndpointsEventV1Server) AutoGetEvent(ctx context.Context, in Event) (Event, error) {
	resp, err := e.AutoGetEventEndpoint(ctx, in)
	if err != nil {
		return Event{}, err
	}
	return *resp.(*Event), nil
}

// MakeEventV1AutoGetEventEndpoint creates  AutoGetEvent endpoints for the service
func MakeEventV1AutoGetEventEndpoint(s ServiceEventV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*Event)
		v, err := s.AutoGetEvent(ctx, *req)
		return respEventV1AutoGetEvent{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventV1:AutoGetEvent")(f)
}

// AutoListEvent implementation on server Endpoint
func (e EndpointsEventV1Server) AutoListEvent(ctx context.Context, in api.ListWatchOptions) (EventList, error) {
	resp, err := e.AutoListEventEndpoint(ctx, in)
	if err != nil {
		return EventList{}, err
	}
	return *resp.(*EventList), nil
}

// MakeEventV1AutoListEventEndpoint creates  AutoListEvent endpoints for the service
func MakeEventV1AutoListEventEndpoint(s ServiceEventV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.ListWatchOptions)
		v, err := s.AutoListEvent(ctx, *req)
		return respEventV1AutoListEvent{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventV1:AutoListEvent")(f)
}

// AutoUpdateEvent implementation on server Endpoint
func (e EndpointsEventV1Server) AutoUpdateEvent(ctx context.Context, in Event) (Event, error) {
	resp, err := e.AutoUpdateEventEndpoint(ctx, in)
	if err != nil {
		return Event{}, err
	}
	return *resp.(*Event), nil
}

// MakeEventV1AutoUpdateEventEndpoint creates  AutoUpdateEvent endpoints for the service
func MakeEventV1AutoUpdateEventEndpoint(s ServiceEventV1Server, logger log.Logger) endpoint.Endpoint {
	f := func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*Event)
		v, err := s.AutoUpdateEvent(ctx, *req)
		return respEventV1AutoUpdateEvent{
			V:   v,
			Err: err,
		}, nil
	}
	return trace.ServerEndpoint("EventV1:AutoUpdateEvent")(f)
}

// AutoWatchEvent is the watch handler for Event on the server side.
func (e EndpointsEventV1Server) AutoWatchEvent(in *api.ListWatchOptions, stream EventV1_AutoWatchEventServer) error {
	return e.watchHandlerEvent(in, stream)
}

// MakeAutoWatchEventEndpoint creates the Watch endpoint
func MakeAutoWatchEventEndpoint(s ServiceEventV1Server, logger log.Logger) func(options *api.ListWatchOptions, stream grpc.ServerStream) error {
	return func(options *api.ListWatchOptions, stream grpc.ServerStream) error {
		wstream := stream.(EventV1_AutoWatchEventServer)
		return s.AutoWatchEvent(options, wstream)
	}
}

// MakeEventV1ServerEndpoints creates server endpoints
func MakeEventV1ServerEndpoints(s ServiceEventV1Server, logger log.Logger) EndpointsEventV1Server {
	return EndpointsEventV1Server{

		AutoAddEventEndpoint:    MakeEventV1AutoAddEventEndpoint(s, logger),
		AutoDeleteEventEndpoint: MakeEventV1AutoDeleteEventEndpoint(s, logger),
		AutoGetEventEndpoint:    MakeEventV1AutoGetEventEndpoint(s, logger),
		AutoListEventEndpoint:   MakeEventV1AutoListEventEndpoint(s, logger),
		AutoUpdateEventEndpoint: MakeEventV1AutoUpdateEventEndpoint(s, logger),

		watchHandlerEvent: MakeAutoWatchEventEndpoint(s, logger),
	}
}

// LoggingEventV1MiddlewareClient adds middleware for the client
func LoggingEventV1MiddlewareClient(logger log.Logger) MiddlewareEventV1Client {
	return func(next ServiceEventV1Client) ServiceEventV1Client {
		return loggingEventV1MiddlewareClient{
			logger: logger,
			next:   next,
		}
	}
}

type loggingEventV1MiddlewareClient struct {
	logger log.Logger
	next   ServiceEventV1Client
}

// LoggingEventV1MiddlewareServer adds middleware for the client
func LoggingEventV1MiddlewareServer(logger log.Logger) MiddlewareEventV1Server {
	return func(next ServiceEventV1Server) ServiceEventV1Server {
		return loggingEventV1MiddlewareServer{
			logger: logger,
			next:   next,
		}
	}
}

type loggingEventV1MiddlewareServer struct {
	logger log.Logger
	next   ServiceEventV1Server
}

func (m loggingEventV1MiddlewareClient) AutoAddEvent(ctx context.Context, in *Event) (resp *Event, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoAddEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoAddEvent(ctx, in)
	return
}
func (m loggingEventV1MiddlewareClient) AutoDeleteEvent(ctx context.Context, in *Event) (resp *Event, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoDeleteEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoDeleteEvent(ctx, in)
	return
}
func (m loggingEventV1MiddlewareClient) AutoGetEvent(ctx context.Context, in *Event) (resp *Event, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoGetEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoGetEvent(ctx, in)
	return
}
func (m loggingEventV1MiddlewareClient) AutoListEvent(ctx context.Context, in *api.ListWatchOptions) (resp *EventList, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoListEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoListEvent(ctx, in)
	return
}
func (m loggingEventV1MiddlewareClient) AutoUpdateEvent(ctx context.Context, in *Event) (resp *Event, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoUpdateEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoUpdateEvent(ctx, in)
	return
}

func (m loggingEventV1MiddlewareClient) AutoWatchEvent(ctx context.Context, in *api.ListWatchOptions) (resp EventV1_AutoWatchEventClient, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoWatchEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoWatchEvent(ctx, in)
	return
}

func (m loggingEventV1MiddlewareServer) AutoAddEvent(ctx context.Context, in Event) (resp Event, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoAddEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoAddEvent(ctx, in)
	return
}
func (m loggingEventV1MiddlewareServer) AutoDeleteEvent(ctx context.Context, in Event) (resp Event, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoDeleteEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoDeleteEvent(ctx, in)
	return
}
func (m loggingEventV1MiddlewareServer) AutoGetEvent(ctx context.Context, in Event) (resp Event, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoGetEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoGetEvent(ctx, in)
	return
}
func (m loggingEventV1MiddlewareServer) AutoListEvent(ctx context.Context, in api.ListWatchOptions) (resp EventList, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoListEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoListEvent(ctx, in)
	return
}
func (m loggingEventV1MiddlewareServer) AutoUpdateEvent(ctx context.Context, in Event) (resp Event, err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(ctx, "service", "EventV1", "method", "AutoUpdateEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	resp, err = m.next.AutoUpdateEvent(ctx, in)
	return
}

func (m loggingEventV1MiddlewareServer) AutoWatchEvent(in *api.ListWatchOptions, stream EventV1_AutoWatchEventServer) (err error) {
	defer func(begin time.Time) {
		var rslt string
		if err == nil {
			rslt = "Success"
		} else {
			rslt = err.Error()
		}
		m.logger.Audit(stream.Context(), "service", "EventV1", "method", "AutoWatchEvent", "result", rslt, "duration", time.Since(begin))
	}(time.Now())
	err = m.next.AutoWatchEvent(in, stream)
	return
}
func (r *EndpointsEventV1RestClient) getHTTPRequest(ctx context.Context, in interface{}, method, path string) (*http.Request, error) {
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
func makeURIEventV1AutoGetEventGetOper(in *Event) string {
	return fmt.Sprint("/v1/event", "/", in.Tenant, "/events/", in.Name)
}

//
func makeURIEventV1AutoListEventListOper(in *api.ListWatchOptions) string {
	return fmt.Sprint("/v1/event", "/", in.Tenant, "/events")
}

// AutoAddEvent CRUD method for Event
func (r *EndpointsEventV1RestClient) AutoAddEvent(ctx context.Context, in *Event) (*Event, error) {
	return nil, errors.New("not allowed")
}

// AutoUpdateEvent CRUD method for Event
func (r *EndpointsEventV1RestClient) AutoUpdateEvent(ctx context.Context, in *Event) (*Event, error) {
	return nil, errors.New("not allowed")
}

// AutoGetEvent CRUD method for Event
func (r *EndpointsEventV1RestClient) AutoGetEvent(ctx context.Context, in *Event) (*Event, error) {
	path := makeURIEventV1AutoGetEventGetOper(in)
	req, err := r.getHTTPRequest(ctx, in, "GET", path)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespEventV1AutoGetEvent(ctx, resp)
	if err != nil {
		return nil, err
	}
	return ret.(*Event), err
}

// AutoDeleteEvent CRUD method for Event
func (r *EndpointsEventV1RestClient) AutoDeleteEvent(ctx context.Context, in *Event) (*Event, error) {
	return nil, errors.New("not allowed")
}

// AutoListEvent CRUD method for Event
func (r *EndpointsEventV1RestClient) AutoListEvent(ctx context.Context, options *api.ListWatchOptions) (*EventList, error) {
	path := makeURIEventV1AutoListEventListOper(options)
	req, err := r.getHTTPRequest(ctx, options, "GET", path)
	if err != nil {
		return nil, err
	}
	resp, err := r.client.Do(req.WithContext(ctx))
	if err != nil {
		return nil, fmt.Errorf("request failed (%s)", err)
	}
	ret, err := decodeHTTPrespEventV1AutoListEvent(ctx, resp)
	if err != nil {
		return nil, err
	}
	return ret.(*EventList), err
}

// AutoWatchEvent CRUD method for Event
func (r *EndpointsEventV1RestClient) AutoWatchEvent(ctx context.Context, in *Event) (*Event, error) {
	return nil, errors.New("not allowed")
}

// MakeEventV1RestClientEndpoints make REST client endpoints
func MakeEventV1RestClientEndpoints(instance string) (EndpointsEventV1RestClient, error) {
	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	return EndpointsEventV1RestClient{
		instance: instance,
		client:   http.DefaultClient,
	}, nil

}
