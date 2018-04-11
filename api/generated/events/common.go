// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package events is a auto generated package.
Input file: protos/events.proto
*/
package events

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"google.golang.org/grpc/metadata"

	"github.com/pensando/sw/api"
	apiserver "github.com/pensando/sw/venice/apiserver"
)

var (
	errInconsistentIDs = errors.New("inconsistent IDs")
	errAlreadyExists   = errors.New("already exists")
	errNotFound        = errors.New("not found")
)

var apisrv apiserver.Server

// FIXME: add HTTP handler here.
func recoverVersion(ctx context.Context, md metadata.MD) context.Context {
	var pairs []string
	xmd := md
	v, ok := xmd[apiserver.RequestParamVersion]
	if ok {
		pairs = append(pairs, apiserver.RequestParamVersion, v[0])
	}
	if v, ok = xmd["req-uri"]; ok {
		pairs = append(pairs, "req-uri", v[0])
	}
	if v, ok = xmd[apiserver.RequestParamMethod]; ok {
		pairs = append(pairs, apiserver.RequestParamMethod, v[0])
	}
	nmd := metadata.Pairs(pairs...)
	nmd = metadata.Join(nmd, md)
	ctx = metadata.NewIncomingContext(ctx, nmd)
	return ctx
}

func encodeHTTPListWatchOptions(ctx context.Context, req *http.Request, request interface{}) error {
	return encodeHTTPRequest(ctx, req, request)
}

func decodeHTTPListWatchOptions(_ context.Context, r *http.Request) (interface{}, error) {
	var req api.ListWatchOptions
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return req, nil
}

// EncodeGrpcReqListWatchOptions encodes ListWatchOptions
func EncodeGrpcReqListWatchOptions(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*api.ListWatchOptions)
	return req, nil
}

// DecodeGrpcReqListWatchOptions encodes ListWatchOptions
func DecodeGrpcReqListWatchOptions(ctx context.Context, request interface{}) (interface{}, error) {
	req := request.(*(api.ListWatchOptions))
	return req, nil
}

// EncodeGrpcRespListWatchOptions encodes response
func EncodeGrpcRespListWatchOptions(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

// DecodeGrpcRespListWatchOptions decodes response
func DecodeGrpcRespListWatchOptions(ctx context.Context, response interface{}) (interface{}, error) {
	return response, nil
}

func encodeHTTPResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeHTTPRequest(_ context.Context, req *http.Request, request interface{}) error {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		return err
	}
	req.Body = ioutil.NopCloser(&buf)
	return nil
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func errorDecoder(r *http.Response) error {
	var w errorWrapper
	if err := json.NewDecoder(r.Body).Decode(&w); err != nil {
		return fmt.Errorf("Status:(%v) Reason:(%s)", r.StatusCode, r.Status)
	}
	return fmt.Errorf("Status:(%v) Reason:(%s)", r.StatusCode, w.Error)
}

type errorWrapper struct {
	Error string `json:"error"`
}

func codeFrom(err error) int {
	switch err {
	case errNotFound:
		return http.StatusNotFound
	case errAlreadyExists, errInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
