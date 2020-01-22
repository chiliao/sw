// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: flowexport.proto
*/
package restapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	protoTypes "github.com/gogo/protobuf/types"
	"github.com/gorilla/mux"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/dscagent/types"
	"github.com/pensando/sw/nic/agent/httputils"
	"github.com/pensando/sw/nic/agent/protos/netproto"
)

// AddFlowExportPolicyAPIRoutes adds FlowExportPolicy routes
func (s *RestServer) AddFlowExportPolicyAPIRoutes(r *mux.Router) {

	r.Methods("GET").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(s.getFlowExportPolicyHandler))

	r.Methods("GET").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(s.listFlowExportPolicyHandler))

	r.Methods("POST").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(s.postFlowExportPolicyHandler))

	r.Methods("DELETE").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(s.deleteFlowExportPolicyHandler))

	r.Methods("PUT").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(s.putFlowExportPolicyHandler))

}

func (s *RestServer) getFlowExportPolicyHandler(r *http.Request) (interface{}, error) {
	tenant, _ := mux.Vars(r)["ObjectMeta.Tenant"]
	namespace, _ := mux.Vars(r)["ObjectMeta.Namespace"]
	name, _ := mux.Vars(r)["ObjectMeta.Name"]
	o := netproto.FlowExportPolicy{
		TypeMeta: api.TypeMeta{Kind: "FlowExportPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    tenant,
			Namespace: namespace,
			Name:      name,
		},
	}

	data, err := s.pipelineAPI.HandleFlowExportPolicy(types.Get, o)
	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return data, nil

}

func (s *RestServer) listFlowExportPolicyHandler(r *http.Request) (interface{}, error) {
	o := netproto.FlowExportPolicy{
		TypeMeta: api.TypeMeta{Kind: "FlowExportPolicy"},
	}

	return s.pipelineAPI.HandleFlowExportPolicy(types.List, o)
}

func (s *RestServer) postFlowExportPolicyHandler(r *http.Request) (interface{}, error) {
	var o netproto.FlowExportPolicy
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}
	c, _ := protoTypes.TimestampProto(time.Now())
	o.CreationTime = api.Timestamp{
		Timestamp: *c,
	}
	o.ModTime = api.Timestamp{
		Timestamp: *c,
	}

	_, err = s.pipelineAPI.HandleFlowExportPolicy(types.Create, o)

	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return Response{
		StatusCode: http.StatusOK,
	}, nil
}

func (s *RestServer) deleteFlowExportPolicyHandler(r *http.Request) (interface{}, error) {
	tenant, _ := mux.Vars(r)["ObjectMeta.Tenant"]
	namespace, _ := mux.Vars(r)["ObjectMeta.Namespace"]
	name, _ := mux.Vars(r)["ObjectMeta.Name"]
	o := netproto.FlowExportPolicy{
		TypeMeta: api.TypeMeta{Kind: "FlowExportPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    tenant,
			Namespace: namespace,
			Name:      name,
		},
	}

	_, err := s.pipelineAPI.HandleFlowExportPolicy(types.Delete, o)
	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return Response{
		StatusCode: http.StatusOK,
	}, nil
}

func (s *RestServer) putFlowExportPolicyHandler(r *http.Request) (interface{}, error) {
	var o netproto.FlowExportPolicy
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}
	c, _ := protoTypes.TimestampProto(time.Now())
	o.CreationTime = api.Timestamp{
		Timestamp: *c,
	}
	o.ModTime = api.Timestamp{
		Timestamp: *c,
	}

	_, err = s.pipelineAPI.HandleFlowExportPolicy(types.Update, o)
	if err != nil {
		return Response{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return Response{
		StatusCode: http.StatusOK,
	}, nil
}