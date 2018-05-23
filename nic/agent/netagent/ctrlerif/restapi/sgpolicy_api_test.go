// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netproto is a auto generated package.
Input file: sgpolicy.proto
*/
package restapi

import (
	"testing"

	api "github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/ctrler/npm/rpcserver/netproto"
	"github.com/pensando/sw/venice/utils/netutils"
	. "github.com/pensando/sw/venice/utils/testutils"
)

func TestSGPolicyList(t *testing.T) {
	t.Parallel()
	var ok bool
	var sgpolicyList []*netproto.SGPolicy

	err := netutils.HTTPGet("http://"+agentRestURL+"/api/security/policies/", &sgpolicyList)

	AssertOk(t, err, "Error getting sgpolicys from the REST Server")
	for _, o := range sgpolicyList {
		if o.Name == "preCreatedSGPolicy" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find preCreatedSGPolicy in Response: %v", sgpolicyList)
	}

}

func TestSGPolicyPost(t *testing.T) {
	t.Parallel()
	var resp Response
	var ok bool
	var sgpolicyList []*netproto.SGPolicy

	postData := netproto.SGPolicy{
		TypeMeta: api.TypeMeta{Kind: "SGPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testPostSGPolicy",
		},
		Spec: netproto.SGPolicySpec{
			AttachGroup:  []string{"preCreatedSecurityGroup"},
			AttachTenant: false,
			Rules: []netproto.PolicyRule{
				{
					Action: []string{"PERMIT"},
					Src: &netproto.MatchSelector{
						Address:   "10.0.0.0 - 10.0.1.0",
						App:       "L4PORT",
						AppConfig: "80",
					},
					Dst: &netproto.MatchSelector{
						Address: "192.168.0.1 - 192.168.1.0",
					},
				},
			},
		},
	}
	err := netutils.HTTPPost("http://"+agentRestURL+"/api/security/policies/", &postData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/security/policies/", &sgpolicyList)

	AssertOk(t, err, "Error posting sgpolicy to REST Server")
	AssertOk(t, getErr, "Error getting sgpolicys from the REST Server")
	for _, o := range sgpolicyList {
		if o.Name == "testPostSGPolicy" {
			ok = true
			break
		}
	}
	if !ok {
		t.Errorf("Could not find testPostSGPolicy in Response: %v", sgpolicyList)
	}

}

func TestSGPolicyDelete(t *testing.T) {
	t.Parallel()
	var resp Response
	var found bool
	var sgpolicyList []*netproto.SGPolicy

	deleteData := netproto.SGPolicy{
		TypeMeta: api.TypeMeta{Kind: "SGPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testDeleteSGPolicy",
		},
		Spec: netproto.SGPolicySpec{
			Rules: []netproto.PolicyRule{
				{
					Action: []string{"PERMIT"},
					Src: &netproto.MatchSelector{
						Address:   "10.0.0.0 - 10.0.1.0",
						App:       "L4PORT",
						AppConfig: "80",
					},
					Dst: &netproto.MatchSelector{
						Address: "192.168.0.1 - 192.168.1.0",
					},
				},
			},
		},
	}
	postErr := netutils.HTTPPost("http://"+agentRestURL+"/api/security/policies/", &deleteData, &resp)
	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/security/policies/default/default/testDeleteSGPolicy", &deleteData, &resp)
	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/security/policies/", &sgpolicyList)

	AssertOk(t, postErr, "Error posting sgpolicy to REST Server")
	AssertOk(t, err, "Error deleting sgpolicy from REST Server")
	AssertOk(t, getErr, "Error getting sgpolicys from the REST Server")
	for _, o := range sgpolicyList {
		if o.Name == "testDeleteSGPolicy" {
			found = true
			break
		}
	}
	if found {
		t.Errorf("Found testDeleteSGPolicy in Response after deleting: %v", sgpolicyList)
	}

}

func TestSGPolicyUpdate(t *testing.T) {
	t.Parallel()
	var resp Response
	var sgpolicyList []*netproto.SGPolicy

	var actualSGPolicySpec netproto.SGPolicySpec
	updatedSGPolicySpec := netproto.SGPolicySpec{
		AttachGroup: []string{"preCreatedSecurityGroup"},
		Rules: []netproto.PolicyRule{
			{
				Action: []string{"DENY", "LOG"},
			},
		},
	}
	putData := netproto.SGPolicy{
		TypeMeta: api.TypeMeta{Kind: "SGPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Name:      "preCreatedSGPolicy",
			Namespace: "default",
		},
		Spec: updatedSGPolicySpec,
	}
	err := netutils.HTTPPut("http://"+agentRestURL+"/api/security/policies/default/default/preCreatedSGPolicy", &putData, &resp)
	AssertOk(t, err, "Error updating sgpolicy to REST Server")

	getErr := netutils.HTTPGet("http://"+agentRestURL+"/api/security/policies/", &sgpolicyList)
	AssertOk(t, getErr, "Error getting sgpolicys from the REST Server")
	for _, o := range sgpolicyList {
		if o.Name == "preCreatedSGPolicy" {
			actualSGPolicySpec = o.Spec
			break
		}
	}
	AssertEquals(t, updatedSGPolicySpec, actualSGPolicySpec, "Could not validate updated spec.")

}

func TestSGPolicyCreateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badPostData := netproto.SGPolicy{
		TypeMeta: api.TypeMeta{Kind: "SGPolicy"},
		ObjectMeta: api.ObjectMeta{
			Name: "",
		},
	}

	err := netutils.HTTPPost("http://"+agentRestURL+"/api/security/policies/", &badPostData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestSGPolicyDeleteErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.SGPolicy{
		TypeMeta: api.TypeMeta{Kind: "SGPolicy"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPDelete("http://"+agentRestURL+"/api/security/policies/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}

func TestSGPolicyUpdateErr(t *testing.T) {
	t.Parallel()
	var resp Response
	badDelData := netproto.SGPolicy{
		TypeMeta: api.TypeMeta{Kind: "SGPolicy"},
		ObjectMeta: api.ObjectMeta{Tenant: "default",
			Namespace: "default",
			Name:      "badObject"},
	}

	err := netutils.HTTPPut("http://"+agentRestURL+"/api/security/policies/default/default/badObject", &badDelData, &resp)

	Assert(t, err != nil, "Expected test to error out with 500. It passed instead")
}
