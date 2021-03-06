package base

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/api/generated/monitoring"
	"github.com/pensando/sw/api/generated/security"
	iota "github.com/pensando/sw/iota/protos/gogen"
	Utils "github.com/pensando/sw/iota/svcs/agent/utils"
	constants "github.com/pensando/sw/iota/svcs/common"
	"github.com/pensando/sw/iota/test/venice/iotakit/cfg/enterprise"
	cfgModel "github.com/pensando/sw/iota/test/venice/iotakit/cfg/enterprise"
	"github.com/pensando/sw/iota/test/venice/iotakit/cfg/enterprise/base"
	"github.com/pensando/sw/iota/test/venice/iotakit/cfg/objClient"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/common"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/objects"
	reporter "github.com/pensando/sw/iota/test/venice/iotakit/model/reporter"
	modelUtils "github.com/pensando/sw/iota/test/venice/iotakit/model/utils"
	"github.com/pensando/sw/iota/test/venice/iotakit/testbed"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/willf/bitset"
)

//RunVerifySystemHealth run system health verification
type RunVerifySystemHealth func(collectLogOnErr bool) error

// SysModel represents a objects.of the system under test
type SysModel struct {
	reporter.Reporter
	Type                  common.ModelType
	RandomTrigger         common.RunRandomTrigger
	RunVerifySystemHealth RunVerifySystemHealth
	enterprise.CfgModel
	NaplesHosts                map[string]*objects.Host           // naples tp hosts map
	WorkloadsObjs              map[string]*objects.Workload       // workloads
	ThirdPartyHosts            map[string]*objects.Host           // naples tp hosts map
	Switches                   map[string]*objects.Switch         // switches in test
	SwitchPortsList            []*objects.SwitchPort              // switches in test
	NaplesNodes                map[string]*objects.Naples         // Naples instances
	FakeHosts                  map[string]*objects.Host           // simulated hosts
	FakeNaples                 map[string]*objects.Naples         // simulated Naples instances
	ThirdPartyNodes            map[string]*objects.ThirdPartyNode // Naples instances
	VeniceNodeMap              map[string]*objects.VeniceNode     // Venice nodes
	VeniceNodesMapDisconnected map[string]*objects.VeniceNode     // Venice which are not part of cluster
	InbandNaplesIPAddress      map[string][]string
	AuthToken                  string   // authToken obtained after logging in
	Licenses                   []string //enabled licenses
	NoModeSwitchReboot         bool     // no reboot on mode switch
	NoSetupDataPathAfterSwitch bool     // temp flag to set up datapath post naples
	AutoDiscovery              bool     //whether discovery of venice from naples is auto

	SkipSetup  bool             //to do skip setup or not
	SkipConfig bool             //to do skip reboot or not
	Tb         *testbed.TestBed // testbed
	Scale      bool
	ScaleData  bool
}

// Init init sys model
func (sm *SysModel) Init(tb *testbed.TestBed, cfgType enterprise.CfgType, reinit bool) error {
	sm.Tb = tb
	sm.NaplesHosts = make(map[string]*objects.Host)
	sm.ThirdPartyHosts = make(map[string]*objects.Host)
	sm.Switches = make(map[string]*objects.Switch)
	sm.NaplesNodes = make(map[string]*objects.Naples)
	sm.ThirdPartyNodes = make(map[string]*objects.ThirdPartyNode)
	sm.VeniceNodeMap = make(map[string]*objects.VeniceNode)
	sm.VeniceNodesMapDisconnected = make(map[string]*objects.VeniceNode)
	sm.FakeNaples = make(map[string]*objects.Naples)
	sm.FakeHosts = make(map[string]*objects.Host)
	sm.WorkloadsObjs = make(map[string]*objects.Workload)
	sm.Reporter = reporter.NewReporter(sm)

	sm.SkipSetup = os.Getenv("SKIP_SETUP") != "" || reinit
	sm.RandomTrigger = sm.RunRandomTrigger
	sm.RunVerifySystemHealth = sm.VerifySystemHealth

	return nil
}

const defaultStyle = "\x1b[0m"
const boldStyle = "\x1b[1m"
const redColor = "\x1b[91m"
const greenColor = "\x1b[32m"
const yellowColor = "\x1b[33m"
const cyanColor = "\x1b[36m"
const grayColor = "\x1b[90m"
const lightGrayColor = "\x1b[37m"

func (sm *SysModel) BeforeTestCallback() {

	if os.Getenv("RANDOM_TRIGGER") != "" {
		err := sm.RandomTrigger(100)
		if err != nil {
			log.Errorf("")
			fmt.Printf("%s%sRunning Random trigger failed %s\n", redColor, boldStyle, defaultStyle)
			sm.ModelExit()
		}
	}
}

func (sm *SysModel) AfterTestCallback() {
	sm.AfterTestCommon()
}

func (sm *SysModel) FailTest(bundleName, groupName, tcName string) {
	tcDir := modelLogsDir + "/" + "'" + bundleName + "'" + "/" + "'" + groupName + "'" + "/" + "'" + tcName + "'"
	sm.DownloadTechsupport(tcDir)
}

// Testbed get testbed obj
func (sm *SysModel) Testbed() *testbed.TestBed {
	return sm.Tb
}

// ConfigClient get objclient
func (sm *SysModel) ConfigClient() objClient.ObjClient {
	return sm.CfgModel.ObjClient()
}

//SetReporter set reporter instance
func (sm *SysModel) SetReporter(rep reporter.Reporter) {
	sm.Reporter = rep
}

//GetReporter gets reporter instance
func (sm *SysModel) GetReporter() reporter.Reporter {
	return sm.Reporter
}

func (sm *SysModel) enableSSHDOnNaples(nodes []*testbed.TestNode) error {

	trig := sm.Tb.NewTrigger()
	for _, node := range nodes {
		if node.Personality == iota.PersonalityType_PERSONALITY_NAPLES {
			for _, naples := range node.NaplesConfigs.Configs {
				//enable sshd
				penctlNaplesURL := "http://" + naples.NaplesIpAddress
				if naples.NaplesSecondaryIpAddress != "" {
					penctlNaplesURL = "http://" + naples.NaplesSecondaryIpAddress
				}
				cmd := fmt.Sprintf("NAPLES_URL=%s %s/entities/%s_host/%s/%s -a %s system enable-sshd",
					penctlNaplesURL, hostToolsDir, node.NodeName, penctlPath, penctlLinuxBinary, constants.PenctlAuthTokenFileName)
				trig.AddCommand(cmd, node.NodeName+"_host", node.NodeName)
				cmd = fmt.Sprintf("NAPLES_URL=%s %s/entities/%s_host/%s/%s  -a %s  update ssh-pub-key -f ~/.ssh/id_rsa.pub",
					penctlNaplesURL, hostToolsDir, node.NodeName, penctlPath, penctlLinuxBinary, constants.PenctlAuthTokenFileName)
				trig.AddCommand(cmd, node.NodeName+"_host", node.NodeName)
			}
		}
	}

	resp, err := trig.Run()
	if err != nil {
		return fmt.Errorf("Error update public key on naples. Err: %v", err)
	}

	for _, cmdResp := range resp {
		if cmdResp.ExitCode != 0 {
			log.Errorf("Changing naples mode failed. %+v", cmdResp)
			return fmt.Errorf("Changing naples mode failed. exit code %v, Out: %v, StdErr: %v",
				cmdResp.ExitCode, cmdResp.Stdout, cmdResp.Stderr)

		}
	}

	return nil
}

// SetupConfig sets up the venice cluster and basic config (like auth etc)
func (sm *SysModel) SetupConfig(ctx context.Context) error {

	// build venice nodes
	for _, nr := range sm.Tb.Nodes {
		if nr.Personality == iota.PersonalityType_PERSONALITY_VENICE {
			// create
			sm.VeniceNodeMap[nr.NodeName] = objects.NewVeniceNode(nr)
		}

	}

	//ReadNaples IP address from the json
	if sm.SkipSetup {
		return sm.RestoreClusterDefaults(sm.Tb.Nodes)
	}

	// make venice cluster
	setupVenice := func(done chan error) {
		err := sm.MakeVeniceCluster(ctx)
		if err != nil {
			log.Errorf("Error creating venice cluster. Err: %v", err)
			done <- err
			return
		}

		// setup auth and wait for venice cluster to come up
		err = sm.InitVeniceConfig(ctx)
		if err != nil {
			log.Errorf("Error configuring cluster. Err: %v", err)
			done <- err
			return
		}

		// setup some tooling on venice nodes
		err = sm.SetupVeniceNodes()
		if err != nil {
			log.Errorf("Error setting up venice nodes. Err: %v", err)
			done <- err
			return
		}
		done <- nil
	}

	doModeSwitch := func(done chan error) {
		if !sm.AutoDiscovery {
			// move naples to managed mode
			err := sm.DoModeSwitchOfNaples(sm.Tb.Nodes, sm.NoModeSwitchReboot)
			if err != nil {
				log.Errorf("Setting up naples failed. Err: %v", err)
				done <- err
				return
			}
		} else {
			log.Infof("Skipping mode switch for naples")
			err := sm.SetupPenctl(sm.Tb.Nodes)
			if err != nil {
				log.Errorf("Setting up of penctl failed Err: %v", err)
				done <- err
				return
			}

		}

		done <- nil
	}
	doneChan := make(chan error, 2)
	//Paralleize venice and naples setup.
	go setupVenice(doneChan)
	go doModeSwitch(doneChan)
	for i := 0; i < 2; i++ {
		err := <-doneChan
		if err != nil {
			return err
		}
	}
	//Wait for node to be admitted
	sm.InitCfgModel()
	cfgClient := sm.ConfigClient()

	bkCtx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancelFunc()
	nc := 0
	for _, node := range sm.Tb.Nodes {
		if testbed.IsNaplesHW(node.Personality) {
			nc += len(node.NaplesConfigs.Configs)
		}
		if node.Personality == iota.PersonalityType_PERSONALITY_NAPLES_MULTI_SIM {
			nc = nc + int(node.NaplesMultSimConfig.NumInstances)
		}
		if node.Personality == iota.PersonalityType_PERSONALITY_NAPLES_CONTROL_SIM {
			nc = nc + 1
		}
	}
L:
	for {
		select {
		case <-bkCtx.Done():
			return fmt.Errorf("Less than %v dscs checked into venice", nc)
		default:
			dscObject, err := cfgClient.ListSmartNIC()
			if err == nil {
				log.Infof("Number of DscObjects %v %v", len(dscObject), nc)
				if len(dscObject) != nc {
					time.Sleep(3 * time.Second)
					continue
				}
			} else {
				log.Infof("got error %v", err)
			}
			admittedCount := 0
			for _, dsc := range dscObject {
				if dsc.Status.AdmissionPhase == cluster.DistributedServiceCardStatus_ADMITTED.String() {
					admittedCount++
				} else {
					log.Infof("Status is %s", dsc.Status.AdmissionPhase)
				}
			}
			if admittedCount == nc {
				log.Infof("Nodes admitted %v", admittedCount)
				break L
			}
			time.Sleep(time.Second)
		}

	}

	err := sm.SetUpNaplesPostCluster(sm.Tb.Nodes)
	if err != nil {
		return err
	}

	return sm.RestoreClusterDefaults(sm.Tb.Nodes)

}

// InitCfgModel init cfg model
func (sm *SysModel) InitCfgModel() error {

	veniceCtx, err := sm.VeniceLoggedInCtx(context.TODO())
	if err != nil {
		return err
	}
	veniceUrls := sm.GetVeniceURL()

	sm.InitClient(veniceCtx, veniceUrls)
	return nil

}

func convertToVeniceFormatMac(s string) string {
	mac := strings.Replace(s, ":", "", -1)
	var buffer bytes.Buffer
	i := 1
	for _, rune := range mac {
		buffer.WriteRune(rune)
		if i != 0 && i%4 == 0 && i != len(mac) {
			buffer.WriteRune('.')
		}
		i++
	}
	return buffer.String()
}

// CreateNaples creates a naples instance
func (sm *SysModel) CreateNaples(node *testbed.TestNode) error {

	snicInRange := func(macAddr string) (sn *cluster.DistributedServiceCard, err error) {

		const maxMacDiff = 24
		snicList, err := sm.ListSmartNIC()
		if err != nil {
			return nil, err
		}

		// walk all smartnics and see if the mac addr range matches
		for _, snic := range snicList {
			snicMacNum := modelUtils.MacAddrToUint64(snic.Status.PrimaryMAC)
			reqMacNum := modelUtils.MacAddrToUint64(macAddr)
			if (snicMacNum == reqMacNum) || ((reqMacNum - snicMacNum) < maxMacDiff) {
				return snic, nil
			}
		}

		return nil, fmt.Errorf("Could not find smartnic with mac addr %s", macAddr)
	}
	for _, naplesConfig := range node.NaplesConfigs.Configs {
		vmac := convertToVeniceFormatMac(naplesConfig.NodeUuid)
		snic, err := sm.GetSmartNICByName(naplesConfig.Name)
		if sm.Tb.IsMockMode() {
			snic, err = snicInRange(naplesConfig.NodeUuid)
		}
		if err != nil {
			//Try to find by mac at least
			snic, err = sm.GetSmartNICByName(vmac)
		}
		if err != nil {
			err := fmt.Errorf("Failed to get smartnc object for name %v(%v). Err: %+v", node.NodeName, vmac, err)
			log.Errorf("%v", err)
			snic = &cluster.DistributedServiceCard{
				TypeMeta: api.TypeMeta{
					Kind: "DistributedServiceCard",
				},
				ObjectMeta: api.ObjectMeta{
					Name: "dsc-1",
				},
				Spec: cluster.DistributedServiceCardSpec{
					ID: "host-1",
					IPConfig: &cluster.IPConfig{
						IPAddress: "1.2.3.4/32",
					},
					MgmtMode:    "NETWORK",
					NetworkMode: "OOB",
				},
				Status: cluster.DistributedServiceCardStatus{
					AdmissionPhase: "ADMITTED",
					PrimaryMAC:     "502f.9ac7.c246",
					IPConfig: &cluster.IPConfig{
						IPAddress: "1.2.3.4",
					},
				},
			}
			return err
		}

		naplesNode, ok := sm.NaplesNodes[node.NodeName]
		if !ok {
			naplesNode = objects.NewNaplesNode(naplesConfig.Name, node)
			sm.NaplesNodes[node.NodeName] = naplesNode
		}
		naplesNode.AddDSC(naplesConfig.Name, snic)
	}

	return nil
}

// createThirdParty creates a naples instance
func (sm *SysModel) createThirdParty(node *testbed.TestNode) error {

	sm.ThirdPartyNodes[node.NodeName] = objects.NewThirdPartyNode(node.NodeName, node)
	return nil
}

// createNaples creates a naples instance
func (sm *SysModel) createMultiSimNaples(node *testbed.TestNode) error {

	numInstances := node.NaplesMultSimConfig.GetNumInstances()
	if len(node.GetIotaNode().GetNaplesMultiSimConfig().GetSimsInfo()) != int(numInstances) {
		err := fmt.Errorf("Number of instances mismatch in iota node and config expected (%v), actual (%v)",
			numInstances, len(node.GetIotaNode().GetNaplesMultiSimConfig().GetSimsInfo()))
		log.Errorf("%v", err)
		return err

	}
	log.Infof("Adding fake naples : %v", (node.NaplesMultSimConfig.GetNumInstances()))

	success := false
	var err error
	for i := 0; i < 3; i++ {
		var snicList []*cluster.DistributedServiceCard
		snicList, err = sm.ListSmartNIC()
		if err != nil {
			continue
		}
		for _, simInfo := range node.GetIotaNode().GetNaplesMultiSimConfig().GetSimsInfo() {
			//TODO: (iota agent is also following the same format.)
			simName := simInfo.GetName()
			success = false
			for _, snic := range snicList {
				if snic.Spec.ID == simName {
					sm.FakeNaples[simName] = objects.NewNaplesNode(simName, node)
					sm.FakeNaples[simName].AddDSC(simName, snic)
					sm.FakeNaples[simName].SetIP(simInfo.GetIpAddress())
					success = true
				}
			}

			if !success {
				err = fmt.Errorf("Failed to get smartnc object for name %v. Err: %+v", node.NodeName, err)
				log.Errorf("%v", err)
				break
			}
		}
		//All got added, success!
		if success {
			break
		}
	}

	if !success {
		return fmt.Errorf("Errorr adding fake naples  %v", err.Error())
	}

	return nil
}

// createControlSimNaples creates a naples instance
func (sm *SysModel) createControlSimNaples(node *testbed.TestNode) error {

	log.Infof("Adding control sim naples : %v(%v)", (node.NodeName), node.GetIotaNode().GetNaplesControlSimConfig().ControlIp)

	success := false
	var err error
	for i := 0; i < 3; i++ {
		var snicList []*cluster.DistributedServiceCard
		snicList, err = sm.ListSmartNIC()
		if err != nil {
			continue
		}
		//TODO: (iota agent is also following the same format.)
		simName := node.NodeName
		success = false
		for _, snic := range snicList {
			if snic.Spec.ID == simName {
				sm.FakeNaples[simName] = objects.NewNaplesNode(simName, node)
				sm.FakeNaples[simName].AddDSC(simName, snic)
				sm.FakeNaples[simName].SetIP(node.GetIotaNode().GetNaplesControlSimConfig().ControlIp)
				success = true
			}
		}

		if !success {
			err = fmt.Errorf("Failed to get smartnc object for name %v. Err: %+v", node.NodeName, err)
			log.Errorf("%v", err)
			break
		}
		//All got added, success!
		if success {
			break
		}
	}

	if !success {
		return fmt.Errorf("Errorr adding fake naples  %v", err.Error())
	}

	return nil
}

//SetupVeniceNaples setups venice and naples
func (sm *SysModel) SetupVeniceNaples() error {
	ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Minute)
	defer cancel()

	// build venice nodes
	for _, nr := range sm.Tb.Nodes {
		if nr.Personality == iota.PersonalityType_PERSONALITY_VENICE {
			// create
			sm.VeniceNodeMap[nr.NodeName] = objects.NewVeniceNode(nr)
		}
	}

	// make cluster & setup auth
	err := sm.SetupConfig(ctx)
	if err != nil {
		sm.Tb.CollectLogs()
		return err
	}

	return nil
}

//SetupNodes setups up nodes
func (sm *SysModel) SetupNodes() error {

	clusterNodes, err := sm.ListClusterNodes()
	if err != nil {
		return err
	}

	// setup nodes
	for _, nr := range sm.Tb.Nodes {
		if testbed.IsNaples(nr.Personality) {
			err := sm.CreateNaples(nr)
			if err != nil {
				return err
			}
		} else if nr.Personality == iota.PersonalityType_PERSONALITY_NAPLES_MULTI_SIM {
			err := sm.createMultiSimNaples(nr)
			if err != nil {
				return err
			}
		} else if nr.Personality == iota.PersonalityType_PERSONALITY_NAPLES_CONTROL_SIM {
			err := sm.createControlSimNaples(nr)
			if err != nil {
				return err
			}
		} else if nr.Personality == iota.PersonalityType_PERSONALITY_VENICE {
			for _, cnode := range clusterNodes {
				if cnode.Name == nr.NodeMgmtIP || nr.SecondaryIP == cnode.Name {
					log.Infof("Setting up cluster node : %v", cnode.Name)
					vnode := sm.VeniceNodeMap[nr.NodeName]
					vnode.ClusterNode = cnode
				}
			}
		} else if testbed.IsThirdParty(nr.Personality) {
			err := sm.createThirdParty(nr)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

// AssociateHosts gets all real hosts from venice cluster
func (sm *SysModel) AssociateHosts() error {
	objs, err := sm.ListHost()
	if err != nil {
		return err
	}

	if sm.Tb.IsMockMode() {
		/*
			//One on One association if mock mode
			convertMac := func(s string) string {
				mac := strings.Replace(s, ".", "", -1)
				var buffer bytes.Buffer
				var l1 = len(mac) - 1
				for i, rune := range mac {
					buffer.WriteRune(rune)
					if i%2 == 1 && i != l1 {
						buffer.WriteRune(':')
					}
				}
				return buffer.String()
			}
			//In mockmode we retrieve mac from the snicList, so we should convert to format that venice likes.
			for _, naples := range sm.NaplesNodes {
				//naples.Nodeuuid = convertMac(naples.Nodeuuid)
			}
		*/
	}

	for k := range sm.NaplesHosts {
		delete(sm.NaplesHosts, k)
	}

	for _, n := range sm.NaplesNodes {
		for _, inst := range n.Instances {
			dsc := inst.Dsc
			dsc.Labels = make(map[string]string)
			nodeMac := strings.Replace(dsc.Status.PrimaryMAC, ":", "", -1)
			nodeMac = strings.Replace(nodeMac, ".", "", -1)
			dscAssociated := false
		L:
			for _, obj := range objs {
				for _, odsc := range obj.Spec.DSCs {
					objMac := strings.Replace(odsc.MACAddress, ".", "", -1)
					if objMac == nodeMac {
						dscAssociated = true
						log.Infof("Associating host %v(ip:%v) with %v(ip:%v)\n", obj.GetName(),
							n.IP(), n.Name(),
							dsc.GetStatus().IPConfig.IPAddress)
						bs := bitset.New(uint(4096))
						bs.Set(0).Set(1).Set(4095)

						dsc, err = sm.GetSmartNIC(dsc.Name)
						if err != nil {
							log.Infof("Error reading smart nic object %v", err)
							return err
						}

						//Add BM type to support upgrade
						dsc.Labels = make(map[string]string)
						dsc.Labels["type"] = "bm"
						if err := sm.UpdateSmartNIC(dsc); err != nil {
							log.Infof("Error updating smart nic object %v", err)
							return err
						}

						h := objects.NewHost(obj, n.GetIotaNode(), n)
						sm.NaplesHosts[n.GetIotaNode().Name] = h
						break L
					}
				}
			}
			if !dscAssociated {
				msg := fmt.Sprintf("Error associating DSC with  %v", dsc.Status.PrimaryMAC)
				log.Infof(msg)
				return fmt.Errorf(msg)
			}
		}
	}

	for k := range sm.FakeNaples {
		delete(sm.FakeHosts, k)
	}
	for simName, n := range sm.FakeNaples {
		for _, inst := range n.Instances {
			dsc := inst.Dsc
			dsc.Labels = make(map[string]string)
			for _, simNaples := range n.GetIotaNode().GetNaplesMultiSimConfig().GetSimsInfo() {
				if simNaples.GetName() == simName {
					nodeMac := strings.Replace(simNaples.GetNodeUuid(), ":", "", -1)
					nodeMac = strings.Replace(nodeMac, ".", "", -1)
					for _, obj := range objs {
						objMac := strings.Replace(obj.GetSpec().DSCs[0].MACAddress, ".", "", -1)
						if objMac == nodeMac {
							log.Infof("Associating host %v(ip:%v) with %v(%v on %v)\n", obj.GetName(),
								n.GetIotaNode().GetIpAddress(), simName, simNaples.GetIpAddress(), n.GetIotaNode().Name)
							bs := bitset.New(uint(4096))
							bs.Set(0).Set(1).Set(4095)

							// add it to database
							h := objects.NewHost(obj, n.GetIotaNode(), n)
							sm.FakeHosts[obj.GetName()] = h
							dsc, err = sm.GetSmartNIC(dsc.Name)
							if err != nil {
								log.Infof("Error reading smart nic object %v", err)
								return err
							}
							//Add BM type to support upgrade
							dsc.Labels = make(map[string]string)
							dsc.Labels["type"] = "sim"
							if err := sm.UpdateSmartNIC(dsc); err != nil {
								log.Infof("Error updating smart nic object %v", err)
								return err
							}
						}
					}
				}
			}

			ctrlSim := n.GetIotaNode().GetNaplesControlSimConfig()
			if ctrlSim != nil {
				nodeMac := strings.Replace(ctrlSim.GetNodeUuid(), ":", "", -1)
				nodeMac = strings.Replace(nodeMac, ".", "", -1)
				for _, obj := range objs {
					objMac := strings.Replace(obj.GetSpec().DSCs[0].MACAddress, ".", "", -1)
					if objMac == nodeMac {
						log.Infof("Associating host %v(ip:%v) with %v(%v on %v)\n", obj.GetName(),
							n.GetIotaNode().GetIpAddress(), simName, ctrlSim.ControlIp, n.GetIotaNode().Name)
						bs := bitset.New(uint(4096))
						bs.Set(0).Set(1).Set(4095)

						// add it to database
						h := objects.NewHost(obj, n.GetIotaNode(), n)
						sm.FakeHosts[obj.GetName()] = h
						dsc, err = sm.GetSmartNIC(dsc.Name)
						if err != nil {
							log.Infof("Error reading smart nic object %v", err)
							return err
						}
						//Add BM type to support upgrade
						dsc.Labels = make(map[string]string)
						dsc.Labels["type"] = "sim"
						if err := sm.UpdateSmartNIC(dsc); err != nil {
							log.Infof("Error updating smart nic object %v", err)
							return err
						}
					}
				}
			}

		}
	}

	for k := range sm.ThirdPartyHosts {
		delete(sm.ThirdPartyHosts, k)
	}
	for _, n := range sm.ThirdPartyNodes {
		for _, obj := range objs {
			for _, odsc := range obj.Spec.DSCs {
				if odsc.MACAddress == n.UUID {
					h := objects.NewHost(obj, n.GetIotaNode(), nil)
					log.Infof("Associating third party host %v(ip:%v) with %v(\n", obj.GetName(),
						n.IP(), n.Name())
					sm.ThirdPartyHosts[n.Name()] = h
				}
			}
		}

	}

	log.Infof("Total number of hosts associated %v\n", len(sm.NaplesHosts)+len(sm.FakeHosts))
	log.Infof("Total number of 3rd Party hosts associated %v\n", len(sm.ThirdPartyHosts))
	return nil
}

func getThirdPartyNic(name, mac string) *cluster.DistributedServiceCard {

	return &cluster.DistributedServiceCard{
		TypeMeta: api.TypeMeta{
			Kind: "DistributedServiceCard",
		},
		ObjectMeta: api.ObjectMeta{
			Name: name,
		},
		Spec: cluster.DistributedServiceCardSpec{
			ID: "host-1",
			IPConfig: &cluster.IPConfig{
				IPAddress: "1.2.3.4/32",
			},
			MgmtMode:    "NETWORK",
			NetworkMode: "OOB",
		},
		Status: cluster.DistributedServiceCardStatus{
			AdmissionPhase: "ADMITTED",
			PrimaryMAC:     mac,
			IPConfig: &cluster.IPConfig{
				IPAddress: "1.2.3.4",
			},
		},
	}
}

func (sm *SysModel) modifyConfig() error {

	log.Infof("Modifying config as per model spec")
	cfgObjects := sm.GetCfgObjects()

	//For base model, network not set
	for i := range cfgObjects.Workloads {
		for j := range cfgObjects.Workloads[i].Spec.Interfaces {
			cfgObjects.Workloads[i].Spec.Interfaces[j].Network = ""
		}
	}
	return nil
}

// InitConfig sets up a default config for the system
func (sm *SysModel) InitConfig(scale, scaleData bool) error {
	skipSetup := os.Getenv("SKIP_SETUP")
	skipConfig := os.Getenv("SKIP_CONFIG")
	cfgParams := &base.ConfigParams{
		Scale:                         scale,
		Regenerate:                    skipSetup == "",
		Vlans:                         sm.Tb.AllocatedVlans(),
		NumberOfInterfacesPerWorkload: 1,
	}
	cfgParams.NaplesLoopBackIPs = make(map[string]string)
	for _, naples := range sm.NaplesNodes {
		dscs := []*cluster.DistributedServiceCard{}
		for _, inst := range naples.Instances {
			dscs = append(dscs, inst.Dsc)
		}
		cfgParams.Dscs = append(cfgParams.Dscs, dscs)
	}

	index := 0
	for name, node := range sm.ThirdPartyNodes {
		node.UUID = "50df.9ac7.c24" + fmt.Sprintf("%v", index)
		n := getThirdPartyNic(name, node.UUID)
		cfgParams.Dscs = append(cfgParams.Dscs, []*cluster.DistributedServiceCard{n})
		cfgParams.ThirdPartyDscs = append(cfgParams.ThirdPartyDscs, n)
		index++
	}

	for _, naples := range sm.FakeNaples {
		//cfgParams.Dscs = append(cfgParams.Dscs, naples.SmartNic)
		cfgParams.FakeDscs = append(cfgParams.FakeDscs, naples.Instances[0].Dsc)
	}

	for _, node := range sm.VeniceNodeMap {
		cfgParams.VeniceNodes = append(cfgParams.VeniceNodes, node.ClusterNode)
	}

	err := sm.PopulateConfig(cfgParams)
	if err != nil {
		return err
	}

	if skipConfig == "" {
		err = sm.CleanupAllConfig()
		if err != nil {
			return err
		}

		err = sm.modifyConfig()
		if err != nil {
			return err
		}

		err = sm.PushConfig()
		if err != nil {
			return err
		}

		ok, err := sm.IsConfigPushComplete()
		if !ok || err != nil {
			return err
		}

	} else {
		log.Info("Skipping config")
	}

	return nil
}

// CreateSwitch creates a new switch
func (sm *SysModel) CreateSwitch(sw *iota.DataSwitch) (*objects.Switch, error) {

	smSw := objects.NewSwitch(sw)

	getHostName := func(ip, port string) (string, error) {
		for _, node := range sm.Tb.Nodes {
			for _, nic := range node.InstanceParams().Nics {
				for _, nport := range nic.Ports {
					portName := "e1/" + strconv.Itoa(nport.SwitchPort)
					if portName == port && nport.SwitchIP == ip {
						return node.NodeName, nil
					}
				}
			}
		}
		return "", fmt.Errorf("Node name not found for switch %v %v", port, ip)
	}
	for _, p := range sw.GetPorts() {

		hostName, err := getHostName(smSw.IP(), p)
		if err != nil {
			return nil, err
		}
		swPort := objects.NewSwitchPort(hostName, smSw, p)

		sm.SwitchPortsList = append(sm.SwitchPortsList, swPort)
	}
	sm.Switches[sw.GetIp()] = smSw

	return smSw, nil
}

// Cleanup TODO
func (sm *SysModel) Cleanup() error {
	// collect all log files
	sm.CollectLogs()
	return nil
}

//GetOrchestrator Default objects.has no orchestrator
func (sm *SysModel) GetOrchestrator() (*objects.Orchestrator, error) {
	return nil, nil
}

// NetworkSecurityPolicy finds an SG policy by name
func (sm *SysModel) NetworkSecurityPolicy(name string) *objects.NetworkSecurityPolicyCollection {
	return nil
}

// DefaultNetworkSecurityPolicy resturns default-policy that prevails across tests cases in the system
func (sm *SysModel) DefaultNetworkSecurityPolicy() *objects.NetworkSecurityPolicyCollection {
	return nil
}

// NewMirrorSession creates a new mirrorsession
func (sm *SysModel) NewMirrorSession(name string) *objects.MirrorSessionCollection {
	return objects.NewMirrorSession(name, sm.ObjClient(), sm.Tb)
}

// MirrorSessions returns all Mirror Sessions in the model
func (sm *SysModel) MirrorSessions() *objects.MirrorSessionCollection {
	msc := objects.MirrorSessionCollection{}
	for _, sess := range msc.Sessions {
		msc.Sessions = append(msc.Sessions, sess)
	}

	return &msc
}

// NewNetworkSecurityPolicy TODO
func (sm *SysModel) NewNetworkSecurityPolicy(name string) *objects.NetworkSecurityPolicyCollection {
	policy := &objects.NetworkSecurityPolicy{
		VenicePolicy: &security.NetworkSecurityPolicy{
			TypeMeta: api.TypeMeta{Kind: "NetworkSecurityPolicy"},
			ObjectMeta: api.ObjectMeta{
				Tenant:    "default",
				Namespace: "default",
				Name:      name,
			},
			Spec: security.NetworkSecurityPolicySpec{
				AttachTenant: true,
			},
		},
	}
	//sm.sgpolicies[name] = policy
	return objects.NewNetworkSecurityPolicyCollection(policy, sm.ObjClient(), sm.Tb)
}

// VerifyRuleStats TODO
func (sm *SysModel) VerifyRuleStats(timestr string, spc *objects.NetworkSecurityPolicyCollection, minCounts []map[string]float64) error {
	return errors.New("Not implemented")
}

func getCfgModelType(mtype testbed.ModelType) cfgModel.CfgType {
	switch mtype {
	case testbed.VcenterModel:
		return cfgModel.VcenterCfgType
	case testbed.CloudModel:
		return cfgModel.VcenterCfgType
	case testbed.BaseNetModel:
		return cfgModel.VcenterCfgType
	}

	return cfgModel.GsCfgType
}

//SetConfigModel changes config model
func (sm *SysModel) SetConfigModel(mType testbed.ModelType) error {

	sm.CfgModel = enterprise.NewCfgModel(getCfgModelType(mType))
	log.Infof("Setting config Model to %v ", mType)
	if sm.CfgModel == nil {
		return errors.New("could not initialize config objects")
	}
	err := sm.InitCfgModel()
	if err != nil {
		return err
	}

	return nil
}

//SetupDefaultConfig setup default config
func (sm *SysModel) SetupDefaultConfig(ctx context.Context, scale, scaleData bool) error {
	panic("Base model does not implement default config")
	return nil
}

// AfterTestCommon common handling after each test
func (sm *SysModel) AfterTestCommon() error {

	retries := 30
	if err := sm.Tb.CheckIotaClusterHealth(); err == nil {
		for i := 0; i < retries; i++ {
			if err = sm.RunVerifySystemHealth(true); err == nil {
				return nil
			}
			time.Sleep(1 * time.Second)
		}
	}

	fmt.Printf("%s%sSystem not in good health, stopping running tests %s\n", redColor, boldStyle, defaultStyle)
	sm.PrintResult()
	fmt.Printf("%s%sStopped running tests as system not in good state%s\n", redColor, boldStyle, defaultStyle)
	os.Exit(1)
	return nil
}

func (sm *SysModel) combineLogs() {

	// create a tar.gz from all log files
	cmdStr := fmt.Sprintf("pushd %s/src/github.com/pensando/sw/iota && tar cvzf venice-iota.tgz *.log logs && mv venice-iota.tgz logs/  && popd", os.Getenv("GOPATH"))
	cmd := exec.Command("bash", "-c", cmdStr)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("tar command out:\n%s\n", string(out))
		log.Errorf("Collecting log files failed with: %s. trying to collect server logs\n", err)
		cmdStr = fmt.Sprintf("pushd %s/src/github.com/pensando/sw/iota/logs && tar cvzf venice-iota.tgz ../*.log && popd", os.Getenv("GOPATH"))
		cmd = exec.Command("bash", "-c", cmdStr)
		out, err = cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("tar command out:\n%s\n", string(out))
			log.Errorf("Collecting server log files failed with: %s.\n", err)
		}
	}
}

// PrintResult prints test result summary
func (sm *SysModel) PrintResult() {
	sm.combineLogs()
	sm.Reporter.PrintReport()
}

func (sm *SysModel) ModelExit() {

	fmt.Printf("%s%sSystem not in good health, stopping running tests %s\n", redColor, boldStyle, defaultStyle)
	sm.DownloadTechsupport(modelLogsDir)
	sm.PrintResult()
	fmt.Printf("%s%sStopped running tests as system not in good state%s\n", redColor, boldStyle, defaultStyle)
	os.Exit(1)
}

func getTechSupportRequest(name string, nodeNames []string) monitoring.TechSupportRequest {
	return monitoring.TechSupportRequest{
		TypeMeta: api.TypeMeta{
			Kind: string(monitoring.KindTechSupportRequest),
		},
		ObjectMeta: api.ObjectMeta{
			Name: name,
		},
		Spec: monitoring.TechSupportRequestSpec{
			Verbosity: 1,
			NodeSelector: &monitoring.TechSupportRequestSpec_NodeSelectorSpec{
				Names: nodeNames,
			},
		},
	}
}

//DownloadTechsupport download techsuport
func (sm *SysModel) DownloadTechsupport(dir string) error {

	log.Infof("Performing techsuppprt....")
	var nodeNames []string

	if dir == "" {
		dir = modelLogsDir
	}
	techsupportName := "iota-techsupport"
	techsupportFileName := techsupportName
	naples := sm.Naples().Names()
	nodeNames = append(nodeNames, naples...)
	for _, vn := range sm.VeniceNodes().Nodes {
		nodeNames = append(nodeNames, vn.IP())
	}
	techsupport := getTechSupportRequest(techsupportName, nodeNames)

	mkdirCmd := []string{"mkdir", "-p", dir}

	if retCode, stdout, err := Utils.RunCmd(mkdirCmd, 0, false, true, nil); err != nil || retCode != 0 {
		log.Errorf("Failed to create directory %v", stdout)
		return fmt.Errorf("Failed to create directory  %v", stdout)
	}

	rmOldTechsupport := []string{"rm", "-f", dir + "/" + techsupportName + "*"}

	if retCode, stdout, err := Utils.RunCmd(rmOldTechsupport, 0, false, true, nil); err != nil || retCode != 0 {
		log.Errorf("Failed to create directory %v", stdout)
		return fmt.Errorf("Failed to create directory  %v", stdout)
	}

	sm.DeleteTechsupport(techsupportName)
	err := sm.PerformTechsupport(&techsupport)
	if err != nil {
		log.Errorf("Error performing tech support %v", err.Error())
		return fmt.Errorf("Error performing tech support %v", err.Error())
	}

	bkCtx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancelFunc()
L:
	for {

		select {
		case <-bkCtx.Done():
			log.Errorf("Error pulling techsupport, timeout!")
			return fmt.Errorf("Error pulling techsupport %v, timeout", techsupportName)
		default:
			if err := sm.VerifyTechsupport(techsupportName); err == nil {
				break L
			}
			time.Sleep(10 * time.Second)
		}

	}
	instanceID, err := sm.GetTechSupportInstanceID(techsupportName)
	if err != nil {
		log.Errorf("Error getting techsupport status  %v", err.Error())
		return fmt.Errorf("Error getting techsupport status %v", err.Error())
	}
	sessionIDFile := cookiefile
	veniceIP := sm.VeniceNodes().Any(1).GenVeniceIPs()[0]
	loginCmd := `curl -k -d '{"username":"admin", "password":"%s", "tenant":"default"}' -c %s -H "Content-Type: application/json" -X POST 'https://%s/v1/login'`
	loginCmd = fmt.Sprintf(loginCmd, constants.UserPassword, sessionIDFile, veniceIP)

	if retCode, stdout, err := Utils.RunCmd(strings.Split(loginCmd, " "), 0, false, true, nil); err != nil || retCode != 0 {
		log.Errorf("Error logging into venice node %v", stdout)
		return fmt.Errorf("Error logging into venice node %v", stdout)
	}
	downloadCmd := `curl -b %s -s -k https://%s/objstore/v1/downloads/all/tenant/default/techsupport/%s --output %s/%s`

	techsupportFileName = techsupportFileName + "-" + strings.Split(instanceID, "-")[0] + ".zip"
	downloadCmd = fmt.Sprintf(downloadCmd, sessionIDFile, veniceIP, techsupportFileName, dir, techsupportFileName)

	if retCode, stdout, err := Utils.RunCmd(strings.Split(downloadCmd, " "), 0, false, true, nil); err != nil || retCode != 0 {
		log.Errorf("Error downloading techsupport %v", stdout)
		return fmt.Errorf("Error logging into venice node %v", stdout)
	}
	return nil
}

var modelLogsDir = fmt.Sprintf("%s/src/github.com/pensando/sw/iota/logs", os.Getenv("GOPATH"))
var cookiefile = fmt.Sprintf("%s/src/github.com/pensando/sw/iota/logs/cookie.jar", os.Getenv("GOPATH"))
