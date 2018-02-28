// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: interface.proto
*/
package restapi

import (
	"testing"

	api "github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/ctrler/npm/rpcserver/netproto"
	"github.com/pensando/sw/venice/utils/netutils"
	. "github.com/pensando/sw/venice/utils/testutils"
)

func TestInterfaceList(t *testing.T) {
	t.Parallel()
	var ok bool
	var interfaceList []*netproto.Interface

	err := netutils.HTTPGet("http://"+agentRestURL+"/api/interfaces/", &interfaceList)

	AssertOk(t, err, "Error getting interfaces from the REST Server")
	for _, o := range interfaceList {
		if o.Name == "preCreatedInterface" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find preCreatedInterface in Response: %v", interfaceList)
	}

}

func TestInterfacePost(t *testing.T) {
	t.Parallel()
	var resp error
	var ok bool
	var interfaceList []*netproto.Interface

	postData := netproto.Interface{
		TypeMeta: api.TypeMeta{Kind: "Interface"},
		ObjectMeta: api.ObjectMeta{
			Tenant: "default",
			Name:   "testPostInterface",
		},
		Spec: netproto.InterfaceSpec{
			Type:        "ENIC",
			AdminStatus: "UP",
		},
		Status: netproto.InterfaceStatus{
			OperStatus: "UP",
		},
	}
	err := netutils.HTTPPost("http://"+agentRestURL+"/api/interfaces/", &postData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/interfaces/", &interfaceList)

	AssertOk(t, err, "Error posting interface to REST Server")
	AssertOk(t, getErr, "Error getting interfaces from the REST Server")
	for _, o := range interfaceList {
		if o.Name == "testPostInterface" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find testPostInterface in Response: %v", interfaceList)
	}

}

func TestInterfaceDelete(t *testing.T) {
	t.Parallel()
	var resp error
	var found bool
	var interfaceList []*netproto.Interface

	deleteData := netproto.Interface{
		TypeMeta: api.TypeMeta{Kind: "Interface"},
		ObjectMeta: api.ObjectMeta{
			Tenant: "default",
			Name:   "testDeleteInterface",
		},
		Spec: netproto.InterfaceSpec{
			Type:        "LIF",
			AdminStatus: "UP",
		},
	}
	postErr := netutils.HTTPPost("http://"+agentRestURL+"/api/interfaces/", &deleteData, &resp)
	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/interfaces/testDeleteInterface", &deleteData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/interfaces/", &interfaceList)

	AssertOk(t, postErr, "Error posting interface to REST Server")
	AssertOk(t, err, "Error deleting interface from REST Server")
	AssertOk(t, getErr, "Error getting interfaces from the REST Server")
	for _, o := range interfaceList {
		if o.Name == "testDeleteInterface" {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Found testDeleteInterface in Response after deleting: %v", interfaceList)
	}

}

func TestInterfaceUpdate(t *testing.T) {
	t.Parallel()
	var resp error
	var interfaceList []*netproto.Interface

	var actualInterfaceSpec netproto.InterfaceSpec
	updatedInterfaceSpec := netproto.InterfaceSpec{
		Type:        "UPLINK",
		AdminStatus: "UP",
	}
	putData := netproto.Interface{
		TypeMeta: api.TypeMeta{Kind: "Interface"},
		ObjectMeta: api.ObjectMeta{
			Tenant: "default",
			Name:   "preCreatedInterface",
		},
		Spec: updatedInterfaceSpec,
	}
	err := netutils.HTTPPut("http://"+agentRestURL+"/api/interfaces/preCreatedInterface", &putData, &resp)
	AssertOk(t, err, "Error updating interface to REST Server")

	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/interfaces/", &interfaceList)
	AssertOk(t, getErr, "Error getting interfaces from the REST Server")

	for _, o := range interfaceList {
		if o.Name == "preCreatedInterface" {
			actualInterfaceSpec = o.Spec
			break
		}
	}
	AssertEquals(t, updatedInterfaceSpec, actualInterfaceSpec, "Could not validate updated spec.")

}
