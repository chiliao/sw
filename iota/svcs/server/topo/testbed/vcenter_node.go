package testbed

import (
	"context"
	"fmt"
	"time"

	iota "github.com/pensando/sw/iota/protos/gogen"
	"github.com/pensando/sw/iota/svcs/agent/workload"
	constants "github.com/pensando/sw/iota/svcs/common"
	vmware "github.com/pensando/sw/iota/svcs/common/vmware"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	//"golang.org/x/sync/errgroup"
)

func (n *VcenterNode) cleanUpVcenter() error {

	//Delete all the VMs associated to the DC/vcenter.
	log.Infof("Cleaning up vcenter node %v %v", n.GetNodeInfo().IPAddress, n.info.License)
	vc, err := vmware.NewVcenter(context.Background(), n.GetNodeInfo().IPAddress, n.GetNodeInfo().Username,
		n.info.Password, n.info.License)
	if err != nil {
		log.Errorf("TOPO SVC | CleanTestBed | Clean Esx node, failed to get host handle  %v %v %v %v %v", err.Error(), n.GetNodeInfo().IPAddress, n.GetNodeInfo().Username,
			n.info.Password, n.info.License)
		return err
	}

	dvsSpec := vmware.DVSwitchSpec{Name: constants.VcenterDCDvs}

	dc, err := vc.SetupDataCenter(constants.VcenterDCName)
	if err == nil {

		//Datacenter exists
		err := dc.DeleteAllVMs()
		if err != nil {
			log.Errorf("TOPO SVC | CleanTestBed | Error disconnecting all hosts from dvs  %v", err.Error())
			return err
		}

		err = dc.DisconnectAllHostFromDvs(dvsSpec)
		if err != nil {
			log.Errorf("TOPO SVC | CleanTestBed | Error disconnecting all hosts from dvs  %v", err.Error())
			return err
		}

		err = dc.DeleteAllHosts()
		if err != nil {
			fmt.Printf("Initing 1 fialed %v.\n", err.Error())
			log.Errorf("TOPO SVC | CleanTestBed | deleting all hosts from datacenter failed  %v", err.Error())
			return err
		}

		err = vc.DestroyDataCenter(constants.VcenterDCName)
		if err != nil {
			fmt.Printf("Initing fialed %v.\n", err.Error())
			log.Errorf("TOPO SVC | CleanTestBed | Destroying data center failed %v", err.Error())
			return err
		}
	}

	return nil
}

func (n *VcenterNode) initVcenter() error {

	if err := n.cleanUpVcenter(); err != nil {
		log.Errorf("TOPO SVC | InitTestBed | Clean up ESX node failed :  %v", err.Error())
		return err
	}

	//Delete all the VMs associated to the DC/vcenter.
	log.Infof("Initing up vcenter node %v %v", n.GetNodeInfo().IPAddress, n.info.License)
	vc, err := vmware.NewVcenter(context.Background(), n.info.IPAddress, n.info.Username,
		n.info.Password, n.info.License)
	if err != nil {
		log.Errorf("TOPO SVC | InitTestbed  | Clean Esx node, failed to get vcenter handle  %v", err.Error())
		return err
	}

	n.vc = vc

	//TODO , create based on current User -ID
	dc, err := vc.CreateDataCenter(constants.VcenterDCName)
	if err != nil {
		log.Errorf("TOPO SVC | InitTestbed  | Failed to create datacenter  %v", err.Error())
		return err
	}

	n.dc = dc

	//Create cluster
	cl, err := dc.CreateCluster(constants.VcenterCluster)
	if err != nil {
		log.Errorf("TOPO SVC | InitTestbed  | Failed to create cluster %v", err.Error())
		return err
	}
	n.cl = cl

	hostSpecs := []vmware.DVSwitchHostSpec{}
	//Connect hosts

	for _, node := range n.managedNodes {
		err = cl.AddHost(node.GetNodeInfo().IPAddress, node.GetNodeInfo().Username,
			node.GetNodeInfo().Password)
		if err != nil {
			log.Errorf("TOPO SVC | InitTestbed  | Failed to add hosts to cluster %v", err.Error())
			return err
		}
		intfs, err := node.GetHostInterfaces()
		if err != nil {
			return err
		}
		log.Infof("Adding pnic %v of host %v to dvs", intfs, node.GetNodeInfo().Name)
		hostSpecs = append(hostSpecs, vmware.DVSwitchHostSpec{
			Name:  node.GetNodeInfo().IPAddress,
			Pnics: intfs,
		})

		vNWs := []vmware.NWSpec{
			{Name: constants.IotaVmotionPortgroup},
		}
		vspec := vmware.VswitchSpec{Name: constants.IotaVmotionSwitch}

		err = dc.AddNetworks(constants.VcenterCluster, node.GetNodeInfo().IPAddress, vNWs, vspec)
		if err != nil {
			//Ignore as it may be created already.
			log.Errorf("Error creating vmotion pg %v", err.Error())
		}

		err = dc.AddKernelNic(constants.VcenterCluster, node.GetNodeInfo().IPAddress, constants.IotaVmotionPortgroup, true)

		if err != nil {
			//Ignore as it may be created already.
			log.Errorf("Error creating vmotion pg %v", err.Error())
		}
	}

	dvsSpec := vmware.DVSwitchSpec{Hosts: hostSpecs,
		Name: constants.VcenterDCDvs, Cluster: constants.VcenterCluster,
		MaxPorts: 10,
		Pvlans: []vmware.DvsPvlanPair{vmware.DvsPvlanPair{Primary: constants.VcenterPvlanStart,
			Secondary: constants.VcenterPvlanStart, Type: "promiscuous"}}}
	err = dc.AddDvs(dvsSpec)
	if err != nil {
		log.Errorf("TOPO SVC | InitTestbed  | Error add DVS with host spec %v", err.Error())
		return err
	}
	return nil
}

//CheckHealth  checks health of node
func (n *VcenterNode) CheckHealth(ctx context.Context, health *iota.NodeHealth) (*iota.NodeHealth, error) {
	health.HealthCode = iota.NodeHealth_HEALTH_OK

	if !n.vc.Active() {
		health.HealthCode = iota.NodeHealth_NODE_DOWN
		msg := fmt.Sprintf("Vcenter node is down %v", n.Node.Name)
		log.Info(msg)
		health.NodeName = n.Node.Name
		return health, errors.New(msg)
	}

	for _, mn := range n.managedNodes {
		rhealth, err := mn.CheckHealth(ctx, health)
		if err != nil || rhealth.HealthCode != iota.NodeHealth_HEALTH_OK {
			health.HealthCode = iota.NodeHealth_NODE_DOWN
			health.NodeName = mn.GetNodeInfo().Name
			msg := fmt.Sprintf("Vcenter node is down %v", mn.GetNodeInfo().Name)
			log.Info(msg)
			return health, errors.New(msg)
		}
	}
	return health, nil
}

// InitNode initializes an iota test node. It copies over IOTA Agent binary and starts it on the remote node
func (n *VcenterNode) InitNode(reboot bool, c *ssh.ClientConfig, commonArtifacts []string) error {

	if err := n.initVcenter(); err != nil {
		return errors.Wrap(err, "init vcenter failed")
	}

	return nil
}

// CleanUpNode cleans up the node
func (n *VcenterNode) CleanUpNode(cfg *ssh.ClientConfig, reboot bool) error {
	if err := n.cleanUpVcenter(); err != nil {
		return errors.Wrap(err, "Clean up vcenter failed")
	}

	return nil
}

// AddNode adds a node for vcenter
func (n *VcenterNode) AddNode() error {
	var err error

	n.RespNode = n.Node
	log.Infof("Calling vcenter node ad for %v", n.GetNodeInfo().Name)
	if err = n.initVcenter(); err != nil {
		err = errors.Wrap(err, "Error in initing Vcenter")
		goto out
	}

out:
	if err != nil {
		n.RespNode.NodeStatus = &iota.IotaAPIResponse{ApiStatus: iota.APIResponseType_API_SERVER_ERROR,
			ErrorMsg: fmt.Sprintf("Error adding node %v", err.Error())}
	} else {
		n.RespNode.NodeStatus = &iota.IotaAPIResponse{ApiStatus: iota.APIResponseType_API_STATUS_OK}

	}

	return err

}

//SetupNode setup node
func (n *VcenterNode) SetupNode() error {

	if n.Node == nil {
		return nil
	}

	n.connector = n.dc
	n.managedNodes = make(map[string]ManagedNodeInterface)
	n.independentNodes = make(map[string]ManagedNodeInterface)
	n.info.ManagedNodes = make(map[string]NodeInfo)
	for _, mnode := range n.Node.GetVcenterConfig().EsxConfigs {
		inf := NodeInfo{
			Os:        iota.TestBedNodeOs_TESTBED_NODE_OS_ESX,
			IPAddress: mnode.IpAddress,
			Username:  mnode.Username,
			Name:      mnode.Name,
			SSHCfg:    InitSSHConfig(mnode.Username, mnode.Password),
			Password:  mnode.Password}
		tNode := NewTestNode(inf)
		mNode, ok := tNode.(ManagedNodeInterface)
		if !ok {
			return errors.Errorf("Node does not implement managed node interface %v", tNode.GetNodeInfo().Name)
		}
		n.managedNodes[mnode.Name] = mNode
		//F
		tNode.SetNodeController(n)
		tNode.RunTriggerLocally()
		tNode.SetConnector(n.dc)
		n.logger.Infof("Adding %v to managed node", mnode.Name)
		n.info.ManagedNodes[mnode.Name] = inf

	}

	return nil
}

//GetControlledNode gets the node which is being controlled by this node
func (n *VcenterNode) GetControlledNode(name string) TestNodeInterface {

	mn, ok := n.managedNodes[name]
	if ok {
		return mn
	}

	return nil
}

//GetManagedNodes get all managed nodes
func (n *VcenterNode) GetManagedNodes() []TestNodeInterface {
	nodes := []TestNodeInterface{}

	for _, node := range n.managedNodes {
		nodes = append(nodes, node)
	}

	return nodes
}

//AssocaiateIndependentNode additional node
func (n *VcenterNode) AssocaiateIndependentNode(node TestNodeInterface) error {

	mNode, ok := node.(ManagedNodeInterface)
	if !ok {
		return errors.Errorf("Node does not implement managed node interface %v", node.GetNodeInfo().Name)
	}

	n.independentNodes[node.GetNodeInfo().Name] = mNode
	log.Infof("Independent nodes added %+v", node)

	agent := node.GetNodeAgent()
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	resp, err := agent.GetWorkloads(ctx, &iota.WorkloadMsg{})
	if err != nil {
		return errors.Wrap(err, "Get workloads failed")
	}

	for index, mn := range n.managedNodes {
		if mn.GetNodeInfo().Name == node.GetNodeInfo().Name {
			//Set agent for some operations
			mn.SetNodeAgent(agent)
			//Save node info for reload
			mn.(*EsxNode).Node = node.(*EsxNode).Node
			n.managedNodes[index] = mNode

			mNode.SetNodeController(n)
			mNode.RunTriggerLocally()
			mNode.SetConnector(n.dc)
			break
		}
	}

	ip, _ := node.GetNodeIP()
	for _, w := range resp.Workloads {
		n.logger.Infof("Adding workload from independent node %v(%v)%p", w.GetWorkloadName(), ip, mNode)
		iotaWload := iotaWorkload{}
		iotaWload.workload = workload.NewWorkload(workload.WorkloadTypeESX, w.GetWorkloadName(),
			n.info.Name, n.logger)
		iotaWload.workload.SetMgmtIP(ip)
		iotaWload.workload.SetWorkloadAgent(agent)

		//Store that workload as part of the node
		//Hack for now
		mNode.(*EsxNode).workloadMap.Store(w.GetWorkloadName(), iotaWload)
		//n.workloadMap.Store(w.GetWorkloadName(), iotaWload)

	}

	return nil
}