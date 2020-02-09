// +build apulu

package apulu

import (
	"testing"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/dscagent/types"
	"github.com/pensando/sw/nic/agent/protos/netproto"
)

func TestHandleInterface(t *testing.T) {
	t.Parallel()
	intf := netproto.Interface{
		TypeMeta: api.TypeMeta{Kind: "Interface"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testInterface",
		},
		Spec: netproto.InterfaceSpec{
			Type:        "L3",
			AdminStatus: "UP",
		},
		Status: netproto.InterfaceStatus{
			InterfaceID: 42,
			IFUplinkStatus: netproto.InterfaceUplinkStatus{
				PortID: 42,
			},
		},
	}

	err := HandleInterface(infraAPI, intfClient, types.Create, intf)
	if err != nil {
		t.Fatal(err)
	}

	err = HandleInterface(infraAPI, intfClient, types.Update, intf)
	if err != nil {
		t.Fatal(err)
	}

	err = HandleInterface(infraAPI, intfClient, types.Delete, intf)
	if err != nil {
		t.Fatal(err)
	}

	err = HandleInterface(infraAPI, intfClient, 42, intf)
	if err == nil {
		t.Fatal("Invalid op must return a valid error.")
	}
}

func TestHandleInterfaceInfraFailures(t *testing.T) {
	t.Parallel()
	intf := netproto.Interface{
		TypeMeta: api.TypeMeta{Kind: "Interface"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "testInterface",
		},
		Spec: netproto.InterfaceSpec{
			Type:        "L3",
			AdminStatus: "UP",
		},
		Status: netproto.InterfaceStatus{
			InterfaceID: 42,
		},
	}
	i := newBadInfraAPI()
	err := HandleInterface(i, intfClient, types.Create, intf)
	if err == nil {
		t.Fatalf("Must return a valid error. Err: %v", err)
	}

	err = HandleInterface(i, intfClient, types.Update, intf)
	if err == nil {
		t.Fatalf("Must return a valid error. Err: %v", err)
	}

	err = HandleInterface(i, intfClient, types.Delete, intf)
	if err == nil {
		t.Fatalf("Must return a valid error. Err: %v", err)
	}
}