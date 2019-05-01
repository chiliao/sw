// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package evtsmgrprotos is a auto generated package.
Input file: eventpolicy.proto
*/
package restapi

import (
	"github.com/pensando/sw/nic/agent/nevtsproxy/ctrlerif/types"
)

// this package contains the REST API provided by the agent

// RestServer is the REST api server
type RestServer struct {
	listenURL string // URL where http server is listening
	handler   types.CtrlerIntf
}

// Response captures the HTTP Response sent by Agent REST Server
type Response struct {
	StatusCode int      `json:"status-code,omitempty"`
	Error      string   `json:"error,omitempty"`
	References []string `json:"references,omitempty"`
}

// MakeErrorResponse generates error response for MakeHTTPHandler() API
func MakeErrorResponse(code int, err error) (*Response, error) {
	res := &Response{
		StatusCode: code,
	}

	if err != nil {
		res.Error = err.Error()
	}

	return res, err
}

// NewRestServer creates a new HTTP server servicg REST api
func NewRestServer(handler types.CtrlerIntf, listenURL string) (*RestServer, error) {
	// create server instance
	srv := RestServer{
		listenURL: listenURL,
		handler:   handler,
	}

	return &srv, nil
}

func (s *RestServer) Stop() error {
	return nil
}
