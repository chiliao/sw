package agent

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	iota "github.com/pensando/sw/iota/protos/gogen"
	Workload "github.com/pensando/sw/iota/svcs/agent/workload"
	"github.com/pensando/sw/iota/svcs/common"
)

const (
	nodeAddTimeout       = 300
	naplesHealthyTimeout = 120
)

var (
	agentService *common.GRPCServer
)

//IotaNode interface
type IotaNode interface {
	Init(*iota.Node) (*iota.Node, error)

	// DeleteNode, remove the personaltiy set
	Destroy(*iota.Node) (*iota.Node, error)

	// AddWorkloads brings up a workload type on a given node
	AddWorkloads(*iota.WorkloadMsg) (*iota.WorkloadMsg, error)

	// DeleteWorkloads deletes workloads specified
	DeleteWorkloads(*iota.WorkloadMsg) (*iota.WorkloadMsg, error)

	// Trigger invokes the workload's trigger. It could be ping, start client/server etc..
	Trigger(*iota.TriggerMsg) (*iota.TriggerMsg, error)

	// CheckClusterHealth returns the cluster health
	CheckHealth(*iota.NodeHealth) (*iota.NodeHealth, error)

	//Type
	NodeType() iota.PersonalityType

	//Type
	NodeName() string

	//Set log
	SetLogger(*log.Logger)

	// GetMsg node msg
	GetMsg() *iota.Node

	//GetWorkloadMsgs
	GetWorkloadMsgs() []*iota.Workload
}

//Base implementations for all node
type iotaNode struct {
	name    string
	nodeMsg *iota.Node
	logger  *log.Logger
}

//Base implementations for all workload node
type iotaWorkload struct {
	name        string
	workloadMsg *iota.Workload
	workload    Workload.Workload
}

func (s *iotaNode) SetLogger(logger *log.Logger) {
	s.logger = logger
}

func (s *iotaNode) NodeName() string {
	return s.name
}

func (s *iotaNode) Destroy(iotanode *iota.Node) (*iota.Node, error) {
	s.logger.Printf("Doing Node clean up.")
	return &iota.Node{NodeStatus: &iota.IotaAPIResponse{ApiStatus: iota.APIResponseType_API_STATUS_OK}}, nil
}

// IOTAAgentListenURL is the default URL for IOTA Agent
var IOTAAgentListenURL = fmt.Sprintf(":%d", common.IotaAgentPort)

// StartIOTAAgent starts IOTA Agent
func StartIOTAAgent(stubMode *bool) {
	agentSvc, err := common.CreateNewGRPCServer("IOTA Agent", IOTAAgentListenURL, common.GrpcMaxMsgSize)
	if err != nil {
		log.Errorf("Could not start IOTA Agent. Err: %v", err)
	}
	// Change this to NewAgentService when ready to integrate
	if !*stubMode {
		agentHandler := NewAgentService()
		iota.RegisterIotaAgentApiServer(agentSvc.Srv, agentHandler)
	} else {
		agentHandler := NewAgentStubService()
		iota.RegisterIotaAgentApiServer(agentSvc.Srv, agentHandler)
	}

	agentService = agentSvc
	agentSvc.Start()
}

// StopIOTAAgent stops the agent
func StopIOTAAgent() {
	agentService.Stop()
}
