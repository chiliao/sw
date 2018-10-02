package restapi

import (
	"flag"
	"os"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/netagent/datapath"
	protos "github.com/pensando/sw/nic/agent/netagent/protos"
	"github.com/pensando/sw/nic/agent/netagent/state"
	tpdatapath "github.com/pensando/sw/nic/agent/tpa/datapath"
	tpstate "github.com/pensando/sw/nic/agent/tpa/state"
	tsdatapath "github.com/pensando/sw/nic/agent/troubleshooting/datapath/hal"
	tsstate "github.com/pensando/sw/nic/agent/troubleshooting/state"
	"github.com/pensando/sw/venice/ctrler/npm/rpcserver/netproto"
	"github.com/pensando/sw/venice/utils/log"
)

func TestMain(m *testing.M) {
	srv, err := setup()
	if err != nil {
		log.Fatalf("Test set up failed. Error: %v", err)
	}
	testCode := m.Run()
	srv.Stop()
	os.Exit(testCode)
}

const (
	agentRestURL      = "localhost:1337"
	agentDatapathKind = "mock"
)

var datapathKind = flag.String("datapath", agentDatapathKind, "Specify the datapath type. mock | hal")

func setup() (*RestServer, error) {

	dp, err := datapath.NewHalDatapath(datapath.Kind(*datapathKind))
	if err != nil {
		log.Errorf("Could not create HAL datapath. Kind: %v, Error %v", datapathKind, err)
		return nil, err
	}

	// Set tenant creation expectation
	if dp.Kind.String() == "mock" {
		dp.Hal.MockClients.MockTnclient.EXPECT().VrfCreate(gomock.Any(), gomock.Any()).Return(nil, nil)
	}

	nagent, err := state.NewNetAgent(dp, protos.AgentMode_CLASSIC, "", "dummy-node-uuid")
	if err != nil {
		log.Errorf("Could not create net agent")
		return nil, err
	}

	err = populatePreTestData(nagent)

	if err != nil {
		log.Errorf("Could not meet prerequisites for testing Endpoint CRUD Methods")
		return nil, err
	}

	tsdp, err := tsdatapath.NewHalDatapath(tsdatapath.Kind(*datapathKind))
	if err != nil {
		log.Errorf("Could not create troubleshooting HAL datapath. Kind: %v, Error %v", datapathKind, err)
		return nil, err
	}
	tsagent, err := tsstate.NewTsAgent(tsdp, protos.AgentMode_CLASSIC, "dummy-node-uuid", nagent)
	if err != nil {
		log.Errorf("Could not create ts troubleshooting agent")
		return nil, err
	}

	tpa, err := tpstate.NewTpAgent(nagent, tpdatapath.MockHal())
	if err != nil {
		log.Fatalf("Error creating telemetry policy agent. Err: %v", err)
	}
	log.Printf("telemetry policy agent {%+v} instantiated", tpa)

	return NewRestServer(nagent, tsagent, tpa, agentRestURL)

}

func populatePreTestData(nagent *state.Nagent) (err error) {
	nt := netproto.Network{
		TypeMeta: api.TypeMeta{Kind: "Network"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedNetwork",
		},
		Spec: netproto.NetworkSpec{
			IPv4Subnet:  "10.1.1.0/24",
			IPv4Gateway: "10.1.1.254",
		},
	}
	err = nagent.CreateNetwork(&nt)
	if err != nil {
		log.Errorf("Failed to create network. {%v}", nt)
		return
	}

	ep := netproto.Endpoint{
		TypeMeta: api.TypeMeta{Kind: "Endpoint"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedEndpoint",
		},
		Spec: netproto.EndpointSpec{
			EndpointUUID: "testEndpointUUID",
			WorkloadUUID: "testWorkloadUUID",
			NetworkName:  "preCreatedNetwork",
			NodeUUID:     "dummy-node-uuid",
			IPv4Address:  "10.1.1.0/24",
			MacAddress:   "42:42:42:42:42:42",
		},
	}
	_, err = nagent.CreateEndpoint(&ep)
	if err != nil {
		log.Errorf("Failed to create endpoint. {%v}", ep)
		return
	}

	sg := netproto.SecurityGroup{
		TypeMeta: api.TypeMeta{Kind: "SecurityGroup"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedSecurityGroup",
		},
		Spec: netproto.SecurityGroupSpec{
			SecurityProfile: "unknown",
		},
	}
	err = nagent.CreateSecurityGroup(&sg)
	if err != nil {
		log.Errorf("Failed to create security group. {%v}", sg)
		return
	}

	natPool := netproto.NatPool{
		TypeMeta: api.TypeMeta{Kind: "NatPool"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedNatPool",
		},
		Spec: netproto.NatPoolSpec{
			IPRange: "10.1.2.1-10.1.2.200",
		},
	}
	err = nagent.CreateNatPool(&natPool)
	if err != nil {
		log.Errorf("Failed to create nat pool. {%v}", sg)
		return
	}

	natPolicy := netproto.NatPolicy{
		TypeMeta: api.TypeMeta{Kind: "NatPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedNatPolicy",
		},
		Spec: netproto.NatPolicySpec{
			Rules: []netproto.NatRule{
				{
					Src: &netproto.MatchSelector{
						Addresses: []string{"10.0.0.0 - 10.0.1.0"},
					},
					Dst: &netproto.MatchSelector{
						Addresses: []string{"192.168.0.0 - 192.168.1.1"},
					},
					NatPool: "preCreatedNatPool",
					Action:  "SNAT",
				},
			},
		},
	}

	err = nagent.CreateNatPolicy(&natPolicy)
	if err != nil {
		log.Errorf("Failed to create nat policy. {%v}", sg)
		return
	}

	tn := netproto.Tenant{
		TypeMeta: api.TypeMeta{Kind: "Tenant"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "preCreatedTenant",
			Name:      "preCreatedTenant",
			Namespace: "preCreatedTenant",
		},
	}

	err = nagent.CreateTenant(&tn)
	if err != nil {
		log.Errorf("Failed to create tenant. {%v}", tn)
		return
	}

	ns := netproto.Namespace{
		TypeMeta: api.TypeMeta{Kind: "Namespace"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "preCreatedTenant",
			Name:      "preCreatedNamespace",
			Namespace: "preCreatedNamespace",
		},
	}

	err = nagent.CreateNamespace(&ns)
	if err != nil {
		log.Errorf("Failed to create Namespace. {%v}", ns)
		return
	}

	rt := netproto.Route{
		TypeMeta: api.TypeMeta{Kind: "Route"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "preCreatedTenant",
			Name:      "preCreatedRoute",
			Namespace: "preCreatedNamespace",
		},
		Spec: netproto.RouteSpec{
			IPPrefix:  "192.168.1.0/24",
			GatewayIP: "192.168.1.1",
			Interface: "uplink-2",
		},
	}

	err = nagent.CreateRoute(&rt)
	if err != nil {
		log.Errorf("Failed to create Route. {%v}", rt)
		return

	}

	nb := netproto.NatBinding{
		TypeMeta: api.TypeMeta{Kind: "NatBinding"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Name:      "preCreatedNatBinding",
			Namespace: "default",
		},
		Spec: netproto.NatBindingSpec{
			NatPoolName: "preCreatedNatPool",
			IPAddress:   "10.1.1.1",
		},
	}

	err = nagent.CreateNatBinding(&nb)
	if err != nil {
		log.Errorf("Failed to create Nat Binding. {%v}", ns)
		return
	}

	ipSecEncrypt := netproto.IPSecSAEncrypt{
		TypeMeta: api.TypeMeta{Kind: "IPSecSAEncrypt"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedIPSecSAEncrypt",
		},
		Spec: netproto.IPSecSAEncryptSpec{
			Protocol:      "ESP",
			AuthAlgo:      "AES_GCM",
			AuthKey:       "someRandomString",
			EncryptAlgo:   "AES_GCM_256",
			EncryptionKey: "someRandomKey",
			LocalGwIP:     "10.0.0.1",
			RemoteGwIP:    "192.168.1.1",
			TepNS:         "infra",
		},
	}
	err = nagent.CreateIPSecSAEncrypt(&ipSecEncrypt)
	if err != nil {
		log.Errorf("Failed to create IPSec Encrypt SA. {%v}", ns)
		return
	}

	ipSecDecrypt := netproto.IPSecSADecrypt{
		TypeMeta: api.TypeMeta{Kind: "IPSecSADecrypt"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedIPSecSADecrypt",
		},
		Spec: netproto.IPSecSADecryptSpec{
			Protocol:      "ESP",
			AuthAlgo:      "AES_GCM",
			AuthKey:       "someRandomString",
			DecryptAlgo:   "AES_GCM_256",
			DecryptionKey: "someRandomKey",
			LocalGwIP:     "10.0.0.1",
			RemoteGwIP:    "192.168.1.1",
			TepNS:         "infra",
		},
	}
	err = nagent.CreateIPSecSADecrypt(&ipSecDecrypt)
	if err != nil {
		log.Errorf("Failed to create IPSec Decrypt SA. {%v}", ns)
		return
	}

	ipSecPolicy := netproto.IPSecPolicy{
		TypeMeta: api.TypeMeta{Kind: "IPSecPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedIPSecPolicy",
		},
		Spec: netproto.IPSecPolicySpec{
			Rules: []netproto.IPSecRule{
				{
					Src: &netproto.MatchSelector{
						Addresses: []string{"10.0.0.0 - 10.0.1.0"},
					},
					Dst: &netproto.MatchSelector{
						Addresses: []string{"192.168.0.1 - 192.168.1.0"},
					},
					SAName: "preCreatedIPSecSAEncrypt",
					SAType: "ENCRYPT",
				},
				{
					Src: &netproto.MatchSelector{
						Addresses: []string{"10.0.0.0 - 10.0.1.0"},
					},
					Dst: &netproto.MatchSelector{
						Addresses: []string{"192.168.0.1 - 192.168.1.0"},
					},
					SAName: "preCreatedIPSecSADecrypt",
					SAType: "DECRYPT",
				},
			},
		},
	}
	err = nagent.CreateIPSecPolicy(&ipSecPolicy)
	if err != nil {
		log.Errorf("Failed to create IPSec Policy. {%v}", ns)
		return
	}

	sgPolicy := netproto.SGPolicy{
		TypeMeta: api.TypeMeta{Kind: "SGPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedSGPolicy",
		},
		Spec: netproto.SGPolicySpec{
			AttachGroup:  []string{"preCreatedSecurityGroup"},
			AttachTenant: false,
			Rules: []netproto.PolicyRule{
				{
					Action: "PERMIT",
					Src: &netproto.MatchSelector{
						Addresses: []string{"10.0.0.0 - 10.0.1.0"},
						AppConfigs: []*netproto.AppConfig{
							{
								Port:     "80",
								Protocol: "tcp",
							},
						},
					},
					Dst: &netproto.MatchSelector{
						Addresses: []string{"192.168.0.1 - 192.168.1.0"},
					},
				},
			},
		},
	}

	err = nagent.CreateSGPolicy(&sgPolicy)
	if err != nil {
		log.Errorf("Failed to create SG policy. {%v}", sg)
		return
	}

	tunnel := netproto.Tunnel{
		TypeMeta: api.TypeMeta{Kind: "Tunnel"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedTunnel",
		},
		Spec: netproto.TunnelSpec{
			Type:        "VXLAN",
			AdminStatus: "UP",
			Src:         "10.1.1.1",
			Dst:         "192.168.1.1",
		},
	}

	err = nagent.CreateTunnel(&tunnel)
	if err != nil {
		log.Errorf("Failed to create tunnel. {%v}", tunnel)
		return
	}

	tcpProxy := netproto.TCPProxyPolicy{
		TypeMeta: api.TypeMeta{Kind: "TCPProxyPolicy"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedTCPProxyPolicy",
		},
		Spec: netproto.TCPProxyPolicySpec{
			Rules: []netproto.TCPProxyRule{
				{
					Src: &netproto.MatchSelector{
						Addresses: []string{"64.0.0.1"},
					},

					Dst: &netproto.MatchSelector{
						Addresses: []string{"100.0.0.1"},
					},
					Action: "ENABLE",
				},
			},
		},
	}

	err = nagent.CreateTCPProxyPolicy(&tcpProxy)
	if err != nil {
		log.Errorf("Failed to create tcp proxy policy. {%v}", tcpProxy)
		return
	}

	lif := netproto.Interface{
		TypeMeta: api.TypeMeta{Kind: "Interface"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedInterface",
		},
		Spec: netproto.InterfaceSpec{
			Type:        "LIF",
			AdminStatus: "UP",
		},
		Status: netproto.InterfaceStatus{
			OperStatus: "UP",
		},
	}

	uplink := netproto.Interface{
		TypeMeta: api.TypeMeta{Kind: "Interface"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedUplink",
		},
		Spec: netproto.InterfaceSpec{
			Type:        "UPLINK",
			AdminStatus: "UP",
		},
		Status: netproto.InterfaceStatus{
			OperStatus: "UP",
		},
	}

	enic := netproto.Interface{
		TypeMeta: api.TypeMeta{Kind: "Interface"},
		ObjectMeta: api.ObjectMeta{
			Tenant:    "default",
			Namespace: "default",
			Name:      "preCreatedEnic",
		},
		Spec: netproto.InterfaceSpec{
			Type:        "ENIC",
			AdminStatus: "UP",
		},
		Status: netproto.InterfaceStatus{
			OperStatus: "UP",
		},
	}

	err = nagent.CreateInterface(&lif)
	if err != nil {
		log.Errorf("Failed to create LIF. {%v}", lif)
		return
	}

	err = nagent.CreateInterface(&uplink)
	if err != nil {
		log.Errorf("Failed to create uplink. {%v}", uplink)
		return
	}

	err = nagent.CreateInterface(&enic)
	if err != nil {
		log.Errorf("Failed to create ENIC. {%v}", enic)
		return
	}

	return nil
}

func TestRestServerStartStop(t *testing.T) {
	t.Parallel()
	// Don't need agent
	restSrv, err := NewRestServer(nil, nil, nil, ":0")
	if err != nil {
		t.Errorf("Could not start REST Server. Error: %v", err)
	}

	restURL := restSrv.GetListenURL()
	if len(restURL) == 0 {
		t.Errorf("Could not get the REST URL. URL: %s", restURL)
	}

	err = restSrv.Stop()
	if err != nil {
		t.Errorf("Failed to stop the REST Server. Error: %v", err)
	}
}

func TestRestServerListenFailures(t *testing.T) {
	t.Parallel()
	restSrv, err := NewRestServer(nil, nil, nil, "")
	if err != nil {
		t.Errorf("Could not start RestServer")
	}
	restSrv.Stop()

	_, err = NewRestServer(nil, nil, nil, ":65536")
	if err == nil {
		t.Errorf("Should see listener errors for the invalid port: %v", err)
	}
}
