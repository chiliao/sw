package topo

import (
	"context"
	"errors"
	"fmt"
	"net"

	"github.com/pensando/sw/iota/svcs/common"

	"golang.org/x/crypto/ssh"

	"golang.org/x/sync/errgroup"

	"github.com/pensando/sw/venice/utils/log"

	iota "github.com/pensando/sw/iota/protos/gogen"
	"github.com/pensando/sw/iota/svcs/server/topo/testbed"
)

// TopologyService implements topology service API
type TopologyService struct {
	SSHConfig        *ssh.ClientConfig
	TestBedInfo      *iota.TestBedMsg //server
	Nodes            []*testbed.TestNode
	ProvisionedNodes map[string]*testbed.TestNode
}

// NewTopologyServiceHandler Topo service handle
func NewTopologyServiceHandler() *TopologyService {
	var topoServer TopologyService
	return &topoServer
}

// InitTestBed does initiates a test bed
func (ts *TopologyService) InitTestBed(ctx context.Context, req *iota.TestBedMsg) (*iota.TestBedMsg, error) {
	log.Infof("TOPO SVC | DEBUG | InitTestBed. Received Request Msg: %v", req)
	var vlans []uint32
	var err error
	ts.TestBedInfo = req

	// Preflight checks
	if len(req.IpAddress) == 0 {
		log.Errorf("TOPO SVC | InitTestBed | No IP Addresses present. Err: %v", ts.TestBedInfo.SwitchPortId, err)
		req.ApiResponse.ApiStatus = iota.APIResponseType_API_BAD_REQUEST
		req.ApiResponse.ErrorMsg = fmt.Sprintf("request message doesn't have any ip addresses")
		return req, nil
	}

	if len(req.User) == 0 || len(req.Passwd) == 0 {
		log.Errorf("TOPO SVC | InitTestBed | User creds to access the vms are missing.")
		req.ApiResponse.ApiStatus = iota.APIResponseType_API_BAD_REQUEST
		req.ApiResponse.ErrorMsg = "user creds are missing to access the VMs"
		return req, nil
	}

	if len(req.ControlIntf) == 0 {
		log.Errorf("")
	}

	// Allocate VLANs for the test bed
	if vlans, err = testbed.AllocateVLANS(ts.TestBedInfo.SwitchPortId); err != nil {
		log.Errorf("TOPO SVC | InitTestBed | Could not allocate VLANS from the switchport id: %d, Err: %v", ts.TestBedInfo.SwitchPortId, err)
		req.ApiResponse.ErrorMsg = fmt.Sprintf("could not allocate VLANS from the switchport id: %d. Err: %v", ts.TestBedInfo.SwitchPortId, err)
		return req, nil
	}
	ts.TestBedInfo.AllocatedVlans = vlans

	ts.SSHConfig = testbed.InitSSHConfig(ts.TestBedInfo.User, ts.TestBedInfo.Passwd)

	// Run init
	initTestBed := func(ctx context.Context) error {
		pool, ctx := errgroup.WithContext(ctx)

		for idx, ipAddress := range req.IpAddress {
			nodeName := fmt.Sprintf("iota-node-%d", idx)
			ipAddr := net.ParseIP(ipAddress)
			if len(ipAddr) == 0 {
				log.Errorf("TOPO SVC | InitTestBed | Invalid IP Address format. %v", ipAddress)
				return fmt.Errorf("invalid ip address format. %v", ipAddress)
			}

			n := testbed.TestNode{
				Node: &iota.Node{
					IpAddress: ipAddress,
					Name:      nodeName,
				},
			}

			ts.Nodes = append(ts.Nodes, &n)
			copyArtifacts := []string{
				common.IotaAgentBinaryPath,
				ts.TestBedInfo.VeniceImage,
				ts.TestBedInfo.NaplesImage,
			}

			pool.Go(func() error {
				n := n
				return n.InitNode(ts.SSHConfig, copyArtifacts)
			})
		}
		return pool.Wait()
	}
	ts.ProvisionedNodes = make(map[string]*testbed.TestNode)
	err = initTestBed(context.Background())
	if err != nil {
		log.Errorf("TOPO SVC | InitTestBed | Init Test Bed Call Failed. %v", err)
		return nil, err
	}
	return ts.TestBedInfo, err
}

// CleanUpTestBed cleans up a testbed
func (ts *TopologyService) CleanUpTestBed(ctx context.Context, req *iota.TestBedMsg) (*iota.TestBedMsg, error) {
	log.Infof("TOPO SVC | DEBUG | CleanUpTestBed. Received Request Msg: %v", req)

	resp := iota.TestBedMsg{}
	return &resp, nil
}

// AddNodes adds nodes to the topology
func (ts *TopologyService) AddNodes(ctx context.Context, req *iota.NodeMsg) (*iota.NodeMsg, error) {
	log.Infof("TOPO SVC | DEBUG | AddNodes. Received Request Msg: %v", req)

	if req.NodeOp != iota.Op_ADD {
		log.Errorf("TOPO SVC | AddNodes | AddNodes call failed")
		req.ApiResponse.ApiStatus = iota.APIResponseType_API_BAD_REQUEST
		req.ApiResponse.ErrorMsg = fmt.Sprintf("AddNodes API must specify Add operation. Found: %v", req.NodeOp)
		return req, nil
	}

	// Prep Topo
	newNodes := []*testbed.TestNode{}
	for _, n := range req.Nodes {
		svcName := n.Name
		agentURL := fmt.Sprintf("%s:%d", n.IpAddress, common.IotaAgentPort)
		c, err := common.CreateNewGRPCClient(svcName, agentURL)

		if err != nil {
			log.Errorf("TOPO SVC | AddNodes | AddNodes call failed to establish GRPC Connection to Agent running on Node: %v. Err: %v", n.Name, err)
			return nil, err
		}

		if _, ok := ts.ProvisionedNodes[n.Name]; ok {
			log.Errorf("TOPO SVC | AddNodes | AddNodes call failed as node already provisoned : %v. Err: %v", n.Name, err)
			return nil, err
		}

		ts.ProvisionedNodes[n.Name] = &testbed.TestNode{
			Node:        n,
			AgentClient: iota.NewIotaAgentApiClient(c.Client),
		}
		newNodes = append(newNodes, ts.ProvisionedNodes[n.Name])
	}

	// Add nodes
	addNodes := func(ctx context.Context) error {
		pool, ctx := errgroup.WithContext(ctx)

		for _, node := range newNodes {
			node := node
			pool.Go(func() error {
				return node.AddNode()
			})
		}
		return pool.Wait()
	}
	err := addNodes(context.Background())
	if err != nil {
		log.Errorf("TOPO SVC | AddNodes |AddNodes Call Failed. %v", err)
		req.ApiResponse.ApiStatus = iota.APIResponseType_API_SERVER_ERROR
		req.ApiResponse.ErrorMsg = fmt.Sprintf("AddNodes Returned the error. Err: %v", err)
		return req, err
	}

	for idx, node := range newNodes {
		req.Nodes[idx] = node.Node
	}

	//req.ApiResponse.ApiStatus = iota.APIResponseType_API_STATUS_OK
	return req, nil
}

// DeleteNodes deletes a node from the topology
func (ts *TopologyService) DeleteNodes(ctx context.Context, req *iota.NodeMsg) (*iota.NodeMsg, error) {
	log.Infof("TOPO SVC | DEBUG | DeleteNodes. Received Request Msg: %v", req)

	resp := &iota.NodeMsg{}
	return resp, nil
}

// GetNodes returns the current topology information
func (ts *TopologyService) GetNodes(ctx context.Context, req *iota.NodeMsg) (*iota.NodeMsg, error) {
	log.Infof("TOPO SVC | DEBUG | GetNodes. Received Request Msg: %v", req)

	resp := &iota.NodeMsg{}
	return resp, nil
}

// AddWorkloads adds a workload on a given node
func (ts *TopologyService) AddWorkloads(ctx context.Context, req *iota.WorkloadMsg) (*iota.WorkloadMsg, error) {
	log.Infof("TOPO SVC | DEBUG | AddWorkloads. Received Request Msg: %v", req)

	if req.WorkloadOp != iota.Op_ADD {
		log.Errorf("TOPO SVC | AddWorkloads | AddWorkloads call failed")
		req.ApiResponse.ApiStatus = iota.APIResponseType_API_BAD_REQUEST
		req.ApiResponse.ErrorMsg = fmt.Sprintf("AddWorkloads must specify Op_Add for workload op. Found: %v", req.WorkloadOp)
		return req, nil
	}

	log.Infof("TOPO SVC | DEBUG | STATE | %v", ts.Nodes)

	workloadNodes := []*testbed.TestNode{}
	for _, w := range req.Workloads {

		node, ok := ts.ProvisionedNodes[w.NodeName]
		if !ok {
			req.ApiResponse.ApiStatus = iota.APIResponseType_API_BAD_REQUEST
			req.ApiResponse.ErrorMsg = fmt.Sprintf("AddWorkloads found to unprovisioned node : %v", w.NodeName)
			return req, nil
		}
		node.WorkloadInfo.Workloads = append(node.WorkloadInfo.Workloads, w)
		added := false
		for _, workloadNode := range workloadNodes {
			if workloadNode.Node.Name == node.Node.Name {
				added = true
				break
			}
		}
		if !added {
			workloadNodes = append(workloadNodes, node)
		}
	}

	// Add workloads
	addWorkloads := func(ctx context.Context) error {
		pool, ctx := errgroup.WithContext(ctx)
		for _, node := range workloadNodes {
			node := node
			pool.Go(func() error {
				for _, w := range node.WorkloadInfo.Workloads {
					if err := node.AddWorkload(w); err != nil {
						return err
					}
				}
				return nil

			})

		}
		return pool.Wait()
	}

	resetAddWorkloads := func() {
		for _, node := range workloadNodes {
			node.WorkloadInfo.Workloads = nil
			node.WorkloadResp.Workloads = nil
		}
	}

	defer resetAddWorkloads()
	err := addWorkloads(context.Background())
	if err != nil {
		log.Errorf("TOPO SVC | AddWorkloads |AddWorkloads Call Failed. %v", err)
		return nil, err
	}

	// TODO return fully formed resp here
	return req, nil
}

// DeleteWorkloads deletes a workload
func (ts *TopologyService) DeleteWorkloads(ctx context.Context, req *iota.WorkloadMsg) (*iota.WorkloadMsg, error) {
	log.Infof("TOPO SVC | DEBUG | DeleteWorkloads. Received Request Msg: %v", req)

	resp := &iota.WorkloadMsg{}

	return resp, nil
}

func (ts *TopologyService) runParallelTrigger(ctx context.Context, req *iota.TriggerMsg) (*iota.TriggerMsg, error) {

	triggerNodes := []*testbed.TestNode{}
	for _, cmd := range req.GetCommands() {
		node, _ := ts.ProvisionedNodes[cmd.GetNodeName()]
		triggerMsg := &iota.TriggerMsg{Commands: []*iota.Command{cmd},
			TriggerMode: req.GetTriggerMode(), TriggerOp: req.GetTriggerOp()}
		node.TriggerInfo = append(node.TriggerInfo, triggerMsg)
		node.TriggerResp = append(node.TriggerResp, triggerMsg)
		added := false
		for _, triggerNode := range triggerNodes {
			if triggerNode.Node.Name == node.Node.Name {
				added = true
				break
			}
		}
		if !added {
			triggerNodes = append(triggerNodes, node)
		}
	}
	// Triggers
	triggers := func(ctx context.Context) error {
		pool, ctx := errgroup.WithContext(ctx)
		for _, node := range triggerNodes {
			node := node
			pool.Go(func() error {
				for idx := range node.TriggerInfo {
					if err := node.Trigger(idx); err != nil {
						return err
					}
				}
				return nil
			})

		}
		return pool.Wait()
	}

	resetTriggers := func() {
		for _, node := range triggerNodes {
			node.TriggerInfo = nil
			node.TriggerResp = nil
		}
	}

	err := triggers(context.Background())
	defer resetTriggers()
	if err != nil {
		log.Errorf("TOPO SVC | Trigger | Trigger Call Failed. %v", err)
		return nil, err
	}

	triggerResp := &iota.TriggerMsg{TriggerMode: req.GetTriggerMode(),
		TriggerOp: req.GetTriggerOp()}
	/* Dequeing the commands in same order as it was queued before. */
	for _, cmd := range req.GetCommands() {
		node, _ := ts.ProvisionedNodes[cmd.GetNodeName()]
		cmdResp := node.TriggerInfo[0]
		node.TriggerInfo = node.TriggerInfo[1:]
		triggerResp.Commands = append(triggerResp.Commands, cmdResp.GetCommands()[0])

	}

	return triggerResp, nil
}

func (ts *TopologyService) runSerialTrigger(ctx context.Context, req *iota.TriggerMsg) (*iota.TriggerMsg, error) {

	for cidx, cmd := range req.GetCommands() {
		node, _ := ts.ProvisionedNodes[cmd.GetNodeName()]
		triggerMsg := &iota.TriggerMsg{Commands: []*iota.Command{cmd},
			TriggerMode: req.GetTriggerMode(), TriggerOp: req.GetTriggerOp()}
		node.TriggerInfo = append(node.TriggerInfo, triggerMsg)
		node.TriggerResp = append(node.TriggerResp, triggerMsg)
		if err := node.Trigger(0); err != nil {
			req.ApiResponse = &iota.IotaAPIResponse{ApiStatus: iota.APIResponseType_API_BAD_REQUEST}
			return req, err
		}
		/* Only one command sent anyway */
		req.Commands[cidx] = node.TriggerResp[0].GetCommands()[0]
		node.TriggerInfo = nil
		node.TriggerResp = nil
	}

	return req, nil
}

// Trigger triggers a workload
func (ts *TopologyService) Trigger(ctx context.Context, req *iota.TriggerMsg) (*iota.TriggerMsg, error) {
	log.Infof("TOPO SVC | DEBUG | Trigger. Received Request Msg: %v", req)

	if req.TriggerOp == iota.TriggerOp_TYPE_NONE {
		log.Errorf("TOPO SVC | Trigger | Trigger call failed")
		req.ApiResponse.ApiStatus = iota.APIResponseType_API_BAD_REQUEST
		req.ApiResponse.ErrorMsg = fmt.Sprintf("Trigger must specify TriggerOp for workload op. Found: %v", req.TriggerOp)
		return req, nil
	}

	for _, cmd := range req.GetCommands() {
		if _, ok := ts.ProvisionedNodes[cmd.NodeName]; !ok {
			req.ApiResponse.ApiStatus = iota.APIResponseType_API_BAD_REQUEST
			req.ApiResponse.ErrorMsg = fmt.Sprintf("Trigger command found to unprovisioned node : %v", cmd.NodeName)
			return req, nil
		}

	}

	if req.GetTriggerMode() == iota.TriggerMode_TRIGGER_PARALLEL {
		return ts.runParallelTrigger(ctx, req)
	}

	return ts.runSerialTrigger(ctx, req)
}

// CheckClusterHealth checks the e2e cluster health
func (ts *TopologyService) CheckClusterHealth(ctx context.Context, req *iota.NodeMsg) (*iota.ClusterHealthMsg, error) {

	resp := &iota.ClusterHealthMsg{}
	for _, reqNode := range req.Nodes {
		if node, ok := ts.ProvisionedNodes[reqNode.Name]; ok {
			nodeHealth := &iota.NodeHealth{NodeName: reqNode.Name}
			nodeResp, err := node.AgentClient.CheckHealth(ctx, nodeHealth)
			if err == nil {
				resp.Health = append(resp.Health, nodeResp)
			}
		} else {
			resp.Health = append(resp.Health, &iota.NodeHealth{NodeName: reqNode.Name, HealthCode: iota.NodeHealth_NOT_PROVISIONED})
		}
	}

	return resp, nil
}

// WorkloadCopy does copy of items to/from workload.
func (ts *TopologyService) WorkloadCopy(ctx context.Context, req *iota.WorkloadCopyMsg) (*iota.WorkloadCopyMsg, error) {

	_, ok := ts.ProvisionedNodes[req.NodeName]
	if !ok {
		errMsg := fmt.Sprintf("Node %s  not provisioned", req.NodeName)
		req.ApiResponse = &iota.IotaAPIResponse{ApiStatus: iota.APIResponseType_API_BAD_REQUEST,
			ErrorMsg: errMsg}
		return nil, errors.New(errMsg)
	}

	if req.Direction == iota.WorkloadCopyDirection_DIR_IN {
		/*if err := node.InitNode(ts, constants.DstIotaAgentDir, artifacts); err != nil {
			log.Errorf("TOPO SVC | InitTestBed | Failed to copy agent binary: %v, to TestNode: %v, at IPAddress: %v", constants.IotaAgentBinaryPath, n.Node.Name, n.Node.IpAddress)
			return err
		}*/
	} else if req.Direction == iota.WorkloadCopyDirection_DIR_OUT {

	} else {
		errMsg := fmt.Sprintf("No direction specified for workload copy")
		req.ApiResponse = &iota.IotaAPIResponse{ApiStatus: iota.APIResponseType_API_BAD_REQUEST,
			ErrorMsg: errMsg}
		return nil, errors.New(errMsg)
	}

	return nil, nil
}
