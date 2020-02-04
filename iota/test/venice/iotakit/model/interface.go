package model

import (
	"context"
	"time"

	"github.com/onsi/gomega"
	"github.com/pensando/sw/api/generated/monitoring"
	"github.com/pensando/sw/api/generated/rollout"

	cfgModel "github.com/pensando/sw/iota/test/venice/iotakit/cfg/enterprise"
	"github.com/pensando/sw/iota/test/venice/iotakit/cfg/objClient"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/common"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/factory"
	"github.com/pensando/sw/iota/test/venice/iotakit/model/objects"
	"github.com/pensando/sw/iota/test/venice/iotakit/testbed"
)

//Type type
type Type int

const (
	//DefaultModel for GS
	DefaultModel Type = 0
	//VcenterModel for Vcenter
	VcenterModel = 1
)

//SysModelInterface interface for sysmodel
type SysModelInterface interface {
	CollectLogs() error
	Cleanup() error

	ActionIntf
	ConfigIntf
	ObjectIntf

	ForEachHost(fn objects.HostIteratorFn) error
	ForEachNaples(fn objects.NaplesIteratorFn) error
	ForEachVeniceNode(fn objects.VeniceNodeIteratorFn) error
	ForEachFakeNaples(fn objects.NaplesIteratorFn) error

	//deleteWorkload(wr *objects.Workload) error
	//findWorkload(name string) *objects.Workload
}

type ConfigIntf interface {
	SetupDefaultConfig(ctx context.Context, scale, scaleData bool) error
	CleanupAllConfig() error
	IsConfigPushComplete() (bool, error)
	ConfigClient() objClient.ObjClient
}

//ActionIntf defines all interfaces
type ActionIntf interface {
	ClusterActionIntf
	VeniceNodeActionIntf
	WorkloadActionIntf
	NodeActionIntf
	NaplesActionIntf
}

type ClusterActionIntf interface {
	FlapDataSwitchPorts(ports *objects.SwitchPortCollection, downTime time.Duration) error

	AllowVeniceNodesFromNaples(vnc *objects.VeniceNodeCollection, naples *objects.NaplesCollection) error

	PerformTechsupport(techsupport *monitoring.TechSupportRequest) error
	VerifyTechsupportStatus(techsupportName string) error

	GetRolloutObject(scaleData bool) (*rollout.Rollout, error)
	PerformRollout(rollout *rollout.Rollout, scaleData bool) error
	VerifyRolloutStatus(rolloutName string) error

	VerifyPolicyStatus(spc *objects.NetworkSecurityPolicyCollection) error
	VerifySystemHealth(collectLogOnErr bool) error
	VerifyClusterStatus() error

	FindFwlogForWorkloadPairs(protocol, fwaction, timestr string, port uint32, wpc *objects.WorkloadPairCollection) error
	VerifyRuleStats(timestr string, spc *objects.NetworkSecurityPolicyCollection, minCounts []map[string]float64) error
	AddNaplesNodes(names []string) error
	DeleteNaplesNodes(names []string) error
	RemoveAddNaples(naples *objects.NaplesCollection) error
	FlapDataSwitchPortsPeriodically(ctx context.Context, ports *objects.SwitchPortCollection,
		downTime time.Duration, flapInterval time.Duration, flapCount int) error
}

type ObjectIntf interface {
	GetOrchestrator() (*objects.Orchestrator, error)
	Hosts() *objects.HostCollection
	HostWorkloads() []*objects.HostWorkloadCollection
	SwitchPorts() *objects.SwitchPortCollection
	VeniceNodes() *objects.VeniceNodeCollection
	Naples() *objects.NaplesCollection
	Networks() *objects.NetworkCollection
	Workloads() *objects.WorkloadCollection
	WorkloadPairs() *objects.WorkloadPairCollection
	NetworkSecurityPolicy(name string) *objects.NetworkSecurityPolicyCollection
	DefaultNetworkSecurityPolicy() *objects.NetworkSecurityPolicyCollection
	NewNetworkSecurityPolicy(name string) *objects.NetworkSecurityPolicyCollection
	NewMirrorSession(name string) *objects.MirrorSessionCollection
	LinkUpEventsSince(since time.Time, npc *objects.NaplesCollection) *objects.EventsCollection
	LinkDownEventsSince(since time.Time, npc *objects.NaplesCollection) *objects.EventsCollection
	ServiceStoppedEvents(since time.Time, npc *objects.NaplesCollection) *objects.EventsCollection
}

//NaplesActionIntf All actions related to naples
type NaplesActionIntf interface {
	StartConsoleLogging() error
	FlapMgmtLinkNaples(naples *objects.NaplesCollection) error
	StartEventsGenOnNaples(naples *objects.NaplesCollection, rate, count string) error
	StopEventsGenOnNaples(naples *objects.NaplesCollection) error
	StartFWLogGenOnNaples(naples *objects.NaplesCollection, rate, count string) error
	StopFWLogGenOnNaples(naples *objects.NaplesCollection) error
	PortFlap(npc *objects.NaplesCollection) error
}

//VeniceNodeActionIntf All actions related to venice
type VeniceNodeActionIntf interface {
	//VeniceNodeLoggedInCtx(context.Context) error
	DisconnectVeniceNodesFromCluster(vnc *objects.VeniceNodeCollection, naples *objects.NaplesCollection) error
	ConnectVeniceNodesToCluster(vnc *objects.VeniceNodeCollection, naples *objects.NaplesCollection) error
	VeniceNodeCreateSnapshotConfig(vnc *objects.VeniceNodeCollection) error
	VeniceNodeTakeSnapshot(vnc *objects.VeniceNodeCollection) (string, error)
	VeniceNodeRestoreConfig(vnc *objects.VeniceNodeCollection, name string) error
	ReloadVeniceNodes(vnc *objects.VeniceNodeCollection) error
	RemoveVenice(venice *objects.VeniceNodeCollection) error
	RemoveAddVenice(venice *objects.VeniceNodeCollection) error
	AddVenice(venice *objects.VeniceNodeCollection) error
	GetVeniceServices() (string, error)
	RunCommandOnVeniceNodes(vnc *objects.VeniceNodeCollection, cmd string) error
}

//NodeActionIntf All actions related node intf
type NodeActionIntf interface {
	BringUpNewWorkloads(hc *objects.HostCollection, snc *objects.NetworkCollection, count int) *objects.WorkloadCollection
	TeardownWorkloads(wc *objects.WorkloadCollection) error

	ReloadHosts(hc *objects.HostCollection) error
	DenyVeniceNodesFromNaples(vnc *objects.VeniceNodeCollection, naples *objects.NaplesCollection) error
	DisconnectNaples(npc *objects.NaplesCollection) error
	ConnectNaples(npc *objects.NaplesCollection) error
	RunNaplesCommand(npc *objects.NaplesCollection, cmd string) ([]string, error)
	RunFakeNaplesBackgroundCommand(npc *objects.NaplesCollection, cmd string) (interface{}, error)

	StopCommands(cmdCtx interface{}) ([]string, error)
	GetNaplesEndpoints(npc *objects.NaplesCollection) (map[string]map[string]struct {
		Local bool
		Vlan  int
	}, error)
}

//WorkloadActionIntf  Interface defining all workload actions
type WorkloadActionIntf interface {
	PingAndCapturePackets(wpc *objects.WorkloadPairCollection, wc *objects.WorkloadCollection, wlnum int) error
	TriggerHping(wpc *objects.WorkloadPairCollection, cmd string) error
	PingPairs(wpc *objects.WorkloadPairCollection) error
	PingFails(wpc *objects.WorkloadPairCollection) error
	ConnectionWithOptions(wpc *objects.WorkloadPairCollection, options *objects.ConnectionOptions) error
	TCPSession(wpc *objects.WorkloadPairCollection, port int) error
	UDPSession(wpc *objects.WorkloadPairCollection, port int) error
	TCPSessionFails(wpc *objects.WorkloadPairCollection, port int) error
	UDPSessionFails(wpc *objects.WorkloadPairCollection, port int) error
	WorkloadsSayHelloToDataPath() error
	VerifyWorkloadStatus(wc *objects.WorkloadCollection) error

	FTPGet(wpc *objects.WorkloadPairCollection) error
	FTPGetFails(wpc *objects.WorkloadPairCollection) error
	DropIcmpFlowTTLSession(wpc *objects.WorkloadPairCollection, cmd string) error
	NetcatWrapper(wpc *objects.WorkloadPairCollection, serverOpt, clientOpt string, port int, expFail bool, expClientExitCode int32, expOutput string) error
	FuzIt(wpc *objects.WorkloadPairCollection, numConns int, proto, port string) error

	QueryDropMetricsForWorkloadPairs(wpc *objects.WorkloadPairCollection, timestr string) error

	MoveWorkloads(*objects.WorkloadCollection, *objects.HostCollection) error
}

func NewSysModel(tb *testbed.TestBed, modelType common.ModelType) (SysModelInterface, error) {

	switch modelType {
	case VcenterModel:
		return factory.NewVcenterSysModel(tb, cfgModel.VcenterCfgType)
	}
	return factory.NewDefaultSysModel(tb, cfgModel.GsCfgType)

}

func getModelTypeFromTopo(mtype testbed.ModelType) common.ModelType {
	switch mtype {
	case testbed.VcenterModel:
		return common.VcenterModel
	}

	return common.DefaultModel
}

// InitSuite initializes test suite
func InitSuite(topoName, paramsFile string, scale, scaleData bool) (*testbed.TestBed, SysModelInterface, error) {
	// create testbed

	// setup test params
	if scale {
		gomega.SetDefaultEventuallyTimeout(time.Minute * 30)
		gomega.SetDefaultEventuallyPollingInterval(time.Second * 30)
	} else {
		gomega.SetDefaultEventuallyTimeout(time.Minute * 6)
		gomega.SetDefaultEventuallyPollingInterval(time.Second * 10)
	}

	tb, err := testbed.NewTestBed(topoName, paramsFile)
	if err != nil {
		return nil, nil, err
	}

	// create sysmodel
	model, err := NewSysModel(tb, getModelTypeFromTopo(tb.Topo.Model))
	if err != nil {
		tb.CollectLogs()
		return nil, nil, err
	}

	// setup default config for the sysmodel
	ctx, cancel := context.WithTimeout(context.TODO(), 60*time.Minute)
	defer cancel()
	err = model.SetupDefaultConfig(ctx, scale, scaleData)
	if err != nil {
		model.CollectLogs()
		return nil, nil, err
	}

	return tb, model, nil
}