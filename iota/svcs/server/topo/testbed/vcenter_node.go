package testbed

import (
	"context"
	"fmt"
	"net"
	"strings"
	"time"

	iota "github.com/pensando/sw/iota/protos/gogen"
	"github.com/pensando/sw/iota/svcs/agent/workload"
	constants "github.com/pensando/sw/iota/svcs/common"
	vmware "github.com/pensando/sw/iota/svcs/common/vmware"
	modelconsts "github.com/pensando/sw/iota/test/venice/iotakit/model/common"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pkg/errors"
	"golang.org/x/crypto/ssh"
	"golang.org/x/sync/errgroup"
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

	// Disconnect all Host from managedNodes
	for _, node := range n.managedNodes {

		// Handling case where host was PXE-booted
		log.Infof("Attempting to disconnect host from current vCenter %v", node.GetNodeInfo().IPAddress)
		err = vc.DisconnectHost(node.GetNodeInfo().IPAddress)
		if err == nil {
			log.Infof("Removing node from vcenter %s", node.GetNodeInfo().IPAddress)
		} else {
			log.Warnf("Failed to disconnect host from vcenter %v %v", node.GetNodeInfo().IPAddress, err.Error())
			// Continue on failure - assuming host is not connected
		}
		time.Sleep(2 * time.Second) // Give an interval
	}

	dvsSpec := vmware.DVSwitchSpec{Name: n.DistributedSwitch}

	dc, err := vc.SetupDataCenter(n.DCName)
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

		err = vc.DestroyDataCenter(n.DCName)
		if err != nil {
			fmt.Printf("Initing failed %v.\n", err.Error())
			log.Errorf("TOPO SVC | CleanTestBed | Destroying data center failed %v", err.Error())
			return err
		}
	}

	return nil
}

func (n *VcenterNode) deployWorkloadTemplates() error {

	images := n.Node.GetVcenterConfig().WorkloadImages
	for _, image := range images {
		err := imageRep.DownloadImage(image)
		if err != nil {
			return err
		}
	}

	pool, _ := errgroup.WithContext(context.Background())
	for _, node := range n.managedNodes {

		node := node
		pool.Go(func() error {
			for _, image := range images {
				imageDir, _ := imageRep.GetImageDir(image)
				tName := templateName(node.GetNodeInfo().IPAddress, image)
				_, err := n.dc.DeployVM(n.ClusterName, node.GetNodeInfo().IPAddress, tName,
					4, 4, constants.EsxDataVMNetworks, imageDir)
				if err != nil {
					return errors.Wrap(err, "Deploy VM failed")
				}
				imageRep.SetImageTemplate(node.GetNodeInfo().IPAddress, image, tName)
			}
			return nil
		})
	}

	err := pool.Wait()
	if err != nil {
		log.Errorf("Error deploying templates.")
		return errors.Wrap(err, "Error deploying templates.")
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
	dc, err := vc.CreateDataCenter(n.DCName)
	if err != nil {
		log.Errorf("TOPO SVC | InitTestbed  | Failed to create datacenter  %v", err.Error())
		return err
	}

	n.dc = dc

	//Create cluster
	cl, err := dc.CreateCluster(n.ClusterName)
	if err != nil {
		log.Errorf("TOPO SVC | InitTestbed  | Failed to create cluster %v", err.Error())
		return err
	}
	n.cl = cl

	hostSpecs := []vmware.DVSwitchHostSpec{}
	//Connect hosts

	for _, node := range n.managedNodes {

		sslThumbprint, err := node.GetSSLThumbprint()
		if err != nil {
			log.Errorf("TOPO SVC | InitTestbed  | Failed to get ssl thumbprint %v", err.Error())
			return err
		}
		err = cl.AddHost(node.GetNodeInfo().IPAddress, node.GetNodeInfo().Username,
			node.GetNodeInfo().Password, sslThumbprint)
		if err != nil {
			log.Errorf("TOPO SVC | InitTestbed  | Failed to add hosts to cluster %v", err.Error())
			return err
		}
		intfs, err := node.GetHostInterfaces()
		if err != nil {
			return err
		}
		log.Infof("Adding pnic %v of host %v (%v) to dvs", intfs, node.GetNodeInfo().Name, node.GetNodeInfo().IPAddress)
		hostSpec := vmware.DVSwitchHostSpec{
			Name: node.GetNodeInfo().IPAddress,
		}
		for _, intf := range intfs {
			hostSpec.Pnics = append(hostSpec.Pnics, intf)
		}
		hostSpecs = append(hostSpecs, hostSpec)

		if n.Node.GetVcenterConfig().EnableVmotionOverMgmt {
			vNWs := []vmware.NWSpec{
				{Name: constants.IotaVmotionPortgroup},
			}
			vspec := vmware.VswitchSpec{Name: constants.IotaVmotionSwitch}

			err = dc.AddNetworks(n.ClusterName, node.GetNodeInfo().IPAddress, vNWs, vspec)
			if err != nil {
				//Ignore as it may be created already.
				log.Errorf("Error creating vmotion pg %v", err.Error())
			}

			nwSpec := vmware.KernelNetworkSpec{
				EnableVmotion: true,
				Portgroup:     constants.IotaVmotionPortgroup,
			}
			err = dc.AddKernelNic(n.ClusterName, node.GetNodeInfo().IPAddress, nwSpec)

			if err != nil {
				//Ignore as it may be created already.
				log.Errorf("Error creating vmotion pg %v", err.Error())
			}
		}
	}

	dvsSpec := vmware.DVSwitchSpec{Hosts: hostSpecs,
		Name: n.DistributedSwitch, Cluster: n.ClusterName,
		Version:  constants.DvsVersion,
		MaxPorts: 100}

	start := n.Node.GetVcenterConfig().PvlanStart
	end := n.Node.GetVcenterConfig().PvlanEnd
	if start == 0 {
		start = 2
		end = constants.ReservedPGVlanCount
	}
	for i := start; i < end; i += 2 {
		pvlanSpecProm := vmware.DvsPvlanPair{
			Primary:   int32(i),
			Type:      "promiscuous",
			Secondary: int32(i),
		}
		pvlanSpecIsolated := vmware.DvsPvlanPair{
			Primary:   int32(i),
			Type:      "isolated",
			Secondary: int32(i + 1),
		}
		dvsSpec.Pvlans = append(dvsSpec.Pvlans, pvlanSpecProm)
		dvsSpec.Pvlans = append(dvsSpec.Pvlans, pvlanSpecIsolated)
	}

	err = dc.AddDvs(dvsSpec)
	if err != nil {
		log.Errorf("TOPO SVC | InitTestbed  | Error add DVS with host spec %v", err.Error())
		return err
	}

	err = n.deployWorkloadTemplates()
	if err != nil {
		log.Errorf("TOPO SVC | InitTestbed  | Failed to deploy VM templates %v", err.Error())
		return err
	}

	//Make sure all nodes now go through controller for operations
	for _, mnode := range n.managedNodes {
		mnode.SetNodeController(n)
		mnode.RunTriggerLocally()
		mnode.SetConnector(n.dc)
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
func (n *VcenterNode) InitNode(reboot, restoreAgentFiles bool, c *ssh.ClientConfig, commonArtifacts []string) error {

	if err := n.cleanUpVcenter(); err != nil {
		return errors.Wrap(err, "Clean up venter failed")
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
	n.DCName = n.Node.GetVcenterConfig().DcName
	n.ClusterName = n.Node.GetVcenterConfig().ClusterName
	n.DistributedSwitch = n.Node.GetVcenterConfig().DistributedSwitch
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
		log.Infof("Checking managed node %v %v", mn.GetNodeInfo().Name, node.GetNodeInfo().Name)
		if mn.GetNodeInfo().Name == node.GetNodeInfo().Name {
			//Set agent for some operations
			mn.SetNodeAgent(agent)
			//Save node info for reload
			mn.(*EsxNode).Node = node.(*EsxNode).Node
			n.managedNodes[index] = mNode

			mNode.SetNodeController(n)
			mNode.RunTriggerLocally()
			mNode.SetConnector(n.dc)
			mNode.SetDC(n.DCName)
			mNode.SetCluster(n.ClusterName)
			log.Infof("Added managed node %v %v", node.GetNodeInfo().Name, mNode.NodeConnector())
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

//AddNetworks add network
func (n *VcenterNode) AddNetworks(ctx context.Context, networkMsg *iota.NetworksMsg) (*iota.NetworksMsg, error) {

	networkMsg.ApiResponse = &iota.IotaAPIResponse{
		ApiStatus: iota.APIResponseType_API_STATUS_OK,
	}

	managedNodes := n.GetManagedNodes()
	for _, nw := range networkMsg.Network {
		if nw.Type == iota.NetworkType_NETWORK_TYPE_VMK_VMOTION {
			for _, mn := range managedNodes {
				if mn.GetNodeInfo().Name == nw.Node {
					nwSpec := vmware.KernelNetworkSpec{
						EnableVmotion: true,
						Portgroup:     nw.Name,
					}

					if !nw.Dhcp {

						ipAddr := ""
						if ip := net.ParseIP(mn.GetNodeInfo().IPAddress); ip == nil {
							addresses, err := net.LookupIP(mn.GetNodeInfo().IPAddress)
							if err != nil {
								networkMsg.ApiResponse.ErrorMsg = errors.Wrap(err, "lookup IP failed").Error()
								networkMsg.ApiResponse.ApiStatus = iota.APIResponseType_API_SERVER_ERROR
								return networkMsg, nil
							}
							for _, addr := range addresses {
								if err := net.ParseIP(addr.String()); err == nil {
									ipAddr = addr.String()
									break
								}
							}
						} else {
							ipAddr = mn.GetNodeInfo().IPAddress
						}
						if ipAddr == "" {
							networkMsg.ApiResponse.ErrorMsg = "Not able to retrieve ip addresss for " + mn.GetNodeInfo().IPAddress
							networkMsg.ApiResponse.ApiStatus = iota.APIResponseType_API_SERVER_ERROR
							return networkMsg, nil
						}
						nwSpec.IPAddress = modelconsts.VmotionSubnet + "." + strings.Split(ipAddr, ".")[3]
						nwSpec.Subnet = "255.255.255.0"
					}

					if nw.MacAddress != "" {
						nwSpec.MacAddress = nw.MacAddress
					}
					log.Infof("Add vmk IP addr %v, %v on node %v", nwSpec.IPAddress, nwSpec.Subnet, nw.Node)
					err := n.dc.AddKernelNic(nw.Cluster, mn.GetNodeInfo().IPAddress, nwSpec)
					if err != nil {
						networkMsg.ApiResponse.ErrorMsg = errors.Wrap(err, "Error adding vmotion pg").Error()
						networkMsg.ApiResponse.ApiStatus = iota.APIResponseType_API_SERVER_ERROR
						return networkMsg, nil
					}
				}
			}
		}
	}
	return networkMsg, nil
}

//RemoveNetworks not supported for other nodes
func (n *VcenterNode) RemoveNetworks(ctx context.Context, req *iota.NetworksMsg) (*iota.NetworksMsg, error) {

	pgs, err := n.dc.FetchDVPortGroupsNames(req.Switch)
	if err != nil {
		req.ApiResponse.ErrorMsg = errors.Wrap(err, "Error removing networks - no PGs found").Error()
		req.ApiResponse.ApiStatus = iota.APIResponseType_API_SERVER_ERROR
		return req, nil
	}
	if req.Network[0] == nil {
		req.ApiResponse.ErrorMsg = errors.Wrap(err, "Error removing networks - no cluster info").Error()
		req.ApiResponse.ApiStatus = iota.APIResponseType_API_SERVER_ERROR
		return req, nil
	}
	managedNodes := n.GetManagedNodes()
	for _, pg := range pgs {
		for _, mn := range managedNodes {
			log.Infof("Remove vmks on PG %v on host %v", pg, mn.GetNodeInfo().Name)
			err := n.dc.RemoveKernelNic(req.Network[0].Cluster, mn.GetNodeInfo().IPAddress, pg)
			if err != nil {
				req.ApiResponse.ErrorMsg = errors.Wrap(err, "Error removing vmknic").Error()
				req.ApiResponse.ApiStatus = iota.APIResponseType_API_SERVER_ERROR
				return req, nil
			}
		}
	}

	err = n.dc.RemoveAllPortGroupsFromDvs(req.Switch)
	if err != nil {
		req.ApiResponse = &iota.IotaAPIResponse{
			ApiStatus: iota.APIResponseType_API_SERVER_ERROR,
			ErrorMsg:  fmt.Sprintf("Error removing networks from %v %v", req.Switch, err.Error()),
		}

		return req, nil
	}

	req.ApiResponse = &iota.IotaAPIResponse{
		ApiStatus: iota.APIResponseType_API_STATUS_OK,
	}

	return req, nil
}

func init() {

}
