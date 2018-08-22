// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: tcp_proxy.proto
*/
package restapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/gorilla/mux"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/httputils"
	agentTypes "github.com/pensando/sw/nic/agent/netagent/state/types"
	"github.com/pensando/sw/venice/ctrler/npm/rpcserver/netproto"
)

// addTCPProxyPolicyAPIRoutes adds TCPProxyPolicy routes
func addTCPProxyPolicyAPIRoutes(r *mux.Router, srv *RestServer) {

	r.Methods("GET").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.listTCPProxyPolicyHandler))

	r.Methods("POST").Subrouter().HandleFunc("/", httputils.MakeHTTPHandler(srv.postTCPProxyPolicyHandler))

	r.Methods("PUT").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.putTCPProxyPolicyHandler))

	r.Methods("DELETE").Subrouter().HandleFunc("/{ObjectMeta.Tenant}/{ObjectMeta.Namespace}/{ObjectMeta.Name}", httputils.MakeHTTPHandler(srv.deleteTCPProxyPolicyHandler))

}

func (s *RestServer) listTCPProxyPolicyHandler(r *http.Request) (interface{}, error) {
	return s.agent.ListTCPProxyPolicy(), nil
}

func (s *RestServer) postTCPProxyPolicyHandler(r *http.Request) (interface{}, error) {
	var res Response

	var o netproto.TCPProxyPolicy
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}
	c, _ := types.TimestampProto(time.Now())
	o.CreationTime = api.Timestamp{
		Timestamp: *c,
	}
	o.ModTime = api.Timestamp{
		Timestamp: *c,
	}

	err = s.agent.CreateTCPProxyPolicy(&o)

	res.References = []string{fmt.Sprintf("%s%s/%s/%s", r.RequestURI, o.Tenant, o.Namespace, o.Name)}

	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()

		return res, err
	}

	res.StatusCode = http.StatusOK
	return res, err
}

func (s *RestServer) putTCPProxyPolicyHandler(r *http.Request) (interface{}, error) {
	var res Response

	var o netproto.TCPProxyPolicy
	b, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(b, &o)
	if err != nil {
		return nil, err
	}
	m, _ := types.TimestampProto(time.Now())
	o.ModTime = api.Timestamp{
		Timestamp: *m,
	}
	err = s.agent.UpdateTCPProxyPolicy(&o)

	res.References = []string{r.RequestURI}

	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()

		return res, err
	}

	res.StatusCode = http.StatusOK
	return res, err
}

func (s *RestServer) deleteTCPProxyPolicyHandler(r *http.Request) (interface{}, error) {
	var res Response

	tenant, _ := mux.Vars(r)["ObjectMeta.Tenant"]
	namespace, _ := mux.Vars(r)["ObjectMeta.Namespace"]
	name, _ := mux.Vars(r)["ObjectMeta.Name"]
	err := s.agent.DeleteTCPProxyPolicy(tenant, namespace, name)

	res.References = []string{r.RequestURI}

	if err != nil {
		res.StatusCode = http.StatusInternalServerError
		res.Error = err.Error()

		// check if its a cannot delete type err
		delErr, ok := err.(*agentTypes.ErrCannotDelete)
		if ok {
			res.References = delErr.References
		}

		return res, err
	}

	res.StatusCode = http.StatusOK
	return res, err
}
