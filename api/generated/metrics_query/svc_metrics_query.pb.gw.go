// Code generated by protoc-gen-grpc-gateway
// source: svc_metrics_query.proto
// DO NOT EDIT!

/*
Package metrics_query is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package metrics_query

import (
	"bytes"
	"io"
	"net/http"

	"github.com/gogo/protobuf/proto"
	"github.com/pensando/grpc-gateway/runtime"
	"github.com/pensando/grpc-gateway/utilities"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
)

var _ codes.Code
var _ io.Reader
var _ = runtime.String
var _ = utilities.NewDoubleArray

func request_MetricsV1_Query_0(ctx context.Context, marshaler runtime.Marshaler, client MetricsV1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	protoReq := &QuerySpec{}
	var smetadata runtime.ServerMetadata

	ver := req.Header.Get("Grpc-Metadata-Req-Version")
	if ver == "" {
		ver = "all"
	}
	if req.ContentLength != 0 {
		var buf bytes.Buffer
		tee := io.TeeReader(req.Body, &buf)
		if err := marshaler.NewDecoder(tee).Decode(protoReq); err != nil {
			return nil, smetadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
		}
		changed := protoReq.Defaults(ver)
		if changed {
			if err := marshaler.NewDecoder(&buf).Decode(protoReq); err != nil {
				return nil, smetadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
			}
		}
	} else {
		protoReq.Defaults(ver)
	}

	msg, err := client.Query(ctx, protoReq, grpc.Header(&smetadata.HeaderMD), grpc.Trailer(&smetadata.TrailerMD))
	return msg, smetadata, err

}

var (
	filter_MetricsV1_Query_1 = &utilities.DoubleArray{Encoding: map[string]int{}, Base: []int(nil), Check: []int(nil)}
)

func request_MetricsV1_Query_1(ctx context.Context, marshaler runtime.Marshaler, client MetricsV1Client, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	protoReq := &QuerySpec{}
	var smetadata runtime.ServerMetadata

	ver := req.Header.Get("Grpc-Metadata-Req-Version")
	if ver == "" {
		ver = "all"
	}
	if req.ContentLength != 0 {
		var buf bytes.Buffer
		tee := io.TeeReader(req.Body, &buf)
		if err := marshaler.NewDecoder(tee).Decode(protoReq); err != nil {
			return nil, smetadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
		}
		changed := protoReq.Defaults(ver)
		if changed {
			if err := marshaler.NewDecoder(&buf).Decode(protoReq); err != nil {
				return nil, smetadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
			}
		}
	} else {
		protoReq.Defaults(ver)
	}

	if err := runtime.PopulateQueryParameters(protoReq, req.URL.Query(), filter_MetricsV1_Query_1); err != nil {
		return nil, smetadata, grpc.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.Query(ctx, protoReq, grpc.Header(&smetadata.HeaderMD), grpc.Trailer(&smetadata.TrailerMD))
	return msg, smetadata, err

}

// RegisterMetricsV1HandlerFromEndpoint is same as RegisterMetricsV1Handler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterMetricsV1HandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Printf("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterMetricsV1Handler(ctx, mux, conn)
}

// RegisterMetricsV1Handler registers the http handlers for service MetricsV1 to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterMetricsV1Handler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	client := NewMetricsV1Client(conn)
	return RegisterMetricsV1HandlerWithClient(ctx, mux, client)
}

// RegisterMetricsV1HandlerClient registers the http handlers for service MetricsV1 to "mux".
// The handlers forward requests to the grpc endpoint using client provided.
func RegisterMetricsV1HandlerWithClient(ctx context.Context, mux *runtime.ServeMux, client MetricsV1Client) error {

	mux.Handle("POST", pattern_MetricsV1_Query_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_MetricsV1_Query_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_MetricsV1_Query_0(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	mux.Handle("GET", pattern_MetricsV1_Query_1, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		if cn, ok := w.(http.CloseNotifier); ok {
			go func(done <-chan struct{}, closed <-chan bool) {
				select {
				case <-done:
				case <-closed:
					cancel()
				}
			}(ctx.Done(), cn.CloseNotify())
		}
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, req)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
		}
		resp, md, err := request_MetricsV1_Query_1(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, outboundMarshaler, w, req, err)
			return
		}

		forward_MetricsV1_Query_1(ctx, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_MetricsV1_Query_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"metrics"}, ""))

	pattern_MetricsV1_Query_1 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0}, []string{"metrics"}, ""))
)

var (
	forward_MetricsV1_Query_0 = runtime.ForwardResponseMessage

	forward_MetricsV1_Query_1 = runtime.ForwardResponseMessage
)
