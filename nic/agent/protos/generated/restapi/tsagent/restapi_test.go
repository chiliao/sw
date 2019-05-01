// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package tsproto is a auto generated package.
Input file: techsupport.proto
*/
package restapi_test

import (
	"flag"
	"os"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/pensando/sw/nic/agent/netagent/ctrlerif/restapi"
	"github.com/pensando/sw/nic/agent/netagent/datapath"
	"github.com/pensando/sw/nic/agent/netagent/state"
	tpdatapath "github.com/pensando/sw/nic/agent/tpa/datapath"
	tpstate "github.com/pensando/sw/nic/agent/tpa/state"
	tsdatapath "github.com/pensando/sw/nic/agent/troubleshooting/datapath/hal"
	tsstate "github.com/pensando/sw/nic/agent/troubleshooting/state"
	"github.com/pensando/sw/venice/utils/log"
)

const (
	agentRestURL      = "localhost:1337"
	agentDatapathKind = "mock"
)

// Response captures the HTTP Response sent by Agent REST Server
type Response struct {
	StatusCode int      `json:"status-code,omitempty"`
	Error      string   `json:"error,omitempty"`
	References []string `json:"references,omitempty"`
}

func TestMain(m *testing.M) {
	srv, err := setup()
	if err != nil {
		log.Fatalf("Test set up failed. Error: %v", err)
	}
	testCode := m.Run()
	srv.Stop()
	os.Exit(testCode)
}

var datapathKind = flag.String("datapath", agentDatapathKind, "Specify the datapath type. mock | hal")

func setup() (*restapi.RestServer, error) {

	dp, err := datapath.NewHalDatapath(datapath.Kind(*datapathKind))
	if err != nil {
		log.Errorf("Could not create HAL datapath. Kind: %v, Error %v", datapathKind, err)
		return nil, err
	}

	// Set tenant creation expectation
	if dp.Kind.String() == "mock" {
		dp.Hal.MockClients.MockTnclient.EXPECT().VrfCreate(gomock.Any(), gomock.Any()).Return(nil, nil)
	}

	nagent, err := state.NewNetAgent(dp, "")
	if err != nil {
		log.Errorf("Could not create net agent")
		return nil, err
	}

	tsdp, err := tsdatapath.NewHalDatapath(tsdatapath.Kind("mock"))
	if err != nil {
		log.Errorf("Could not create troubleshooting HAL datapath. Kind: %v, Error %v", datapathKind, err)
		return nil, err
	}
	tsagent, err := tsstate.NewTsAgent(tsdp, "dummy-node-uuid", nagent)
	if err != nil {
		log.Errorf("Could not create ts troubleshooting agent")
		return nil, err
	}

	tpa, err := tpstate.NewTpAgent(nagent, func() string { return "192.168.100.101" }, tpdatapath.MockHal())
	if err != nil {
		log.Fatalf("Error creating telemetry policy agent. Err: %v", err)
	}
	log.Printf("telemetry policy agent {%+v} instantiated", tpa)

	if err != nil {
		log.Errorf("Could not meet prerequisites for testing Endpoint CRUD Methods")
		return nil, err
	}

	return restapi.NewRestServer(nagent, tsagent, tpa, agentRestURL)

}

func TestRestServerStartStop(t *testing.T) {
	t.Parallel()
	// Don't need agent
	restSrv, err := restapi.NewRestServer(nil, nil, nil, ":0")
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
	restSrv, err := restapi.NewRestServer(nil, nil, nil, "")
	if err != nil {
		t.Errorf("Could not start RestServer")
	}
	restSrv.Stop()

	_, err = restapi.NewRestServer(nil, nil, nil, ":65536")
	if err == nil {
		t.Errorf("Should see listener errors for the invalid port: %v", err)
	}
}
