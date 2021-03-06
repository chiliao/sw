// {C} Copyright 2017 Pensando Systems Inc. All rights reserved.

package types

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/pensando/sw/nic/agent/dscagent/types/irisproto"

	"github.com/mdlayher/arp"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	clientApi "github.com/pensando/sw/nic/delphi/gosdk/client_api"
	"github.com/pensando/sw/venice/utils/emstore"
)

const (
	// UplinkOffset will ensure that the agent will allocate uplink IDs starting from 128 and this will not clash with any of
	// the internal hal interfaces i.e NETAGENT_IF_ID_UPLINK_MIN
	UplinkOffset = 127

	// SecurityProfileOffset will ensure that agent will allocate security profiles stating from 16 i.e NETAGENT_NWSEC_PROFILE_ID_MIN
	SecurityProfileOffset = 15

	// EnicOffset will ensure that agent will allocate enic IDs starting from UplinkOffset + 10000
	EnicOffset = UplinkOffset + 10000

	// TunnelOffset will ensure that agent will allocate tunnel IDs starting from UplinkOffset + 20000
	TunnelOffset = UplinkOffset + 20000

	// VrfOffset will ensure that agent will allocate VRF IDs starting from 65.i.e. NETAGENT_VRF_ID_MIN
	VrfOffset = 65

	// NetworkOffset will ensure that agent will allocate VRD IDs starting from 16 i.e NETAGENT_L2SEG_ID_MIN
	NetworkOffset = 100
)

// ID types
const (
	NetworkID               = "networkID"
	SecurityGroupID         = "sgID"
	TenantID                = "tenantID"
	NamespaceID             = "namespaceID"
	VrfID                   = "vrfID"
	InterfaceID             = "interfaceID"
	NatPoolID               = "natPoolID"
	NatPolicyID             = "natPolicyID"
	NatRuleID               = "natRuleID"
	RouteID                 = "routeID"
	NatBindingID            = "natBindingID"
	IPSecPolicyID           = "ipSecPolicyID"
	IPSecSAEncryptID        = "ipSecSAEncryptID"
	IPSecSADecryptID        = "ipSecSADencryptID"
	IPSecRuleID             = "ipSecRuleID"
	NetworkSecurityPolicyID = "securityGroupPolicyID"
	TCPProxyPolicyID        = "tcpProxyPolicyID"
	SecurityProfileID       = "securityProfileID"
	AppID                   = "appID"
	UntaggedCollVLAN        = 8191
	OOBManagementInterface  = "oob_mnic0"
	ManagementUplink        = "uplink130"
	InternalUntaggedNetwork = "_internal_untagged_nw"
)

// IntfInfo has the interface names to be plumbed into container
type IntfInfo struct {
	ContainerIntfName string //  Name of container side of the interface
	SwitchIntfName    string // Name of switch side of the interface
}

// ErrCannotDelete is returned when a delete on an object having dependent objects is attempted.
type ErrCannotDelete struct {
	//Message string `json:"message,omitempty"`
	References []string `json:"references,omitempty"`
}

func (e *ErrCannotDelete) Error() (err string) {
	if len(e.References) > 0 {
		//e.Message = fmt.Sprintf("cannot delete an object that has pending references.")
		//b, _ := json.Marshal(e)
		err = fmt.Sprintf("cannot delete an object that has pending references.")
		return
	}
	return
}

// NatPoolRef keeps the mapping between a natpool ID and its corresponding NamespaceID.
// We need this to build a look up table between the natpool name and its refs which can be in non local namespaces
type NatPoolRef struct {
	NamespaceID uint64
	PoolID      uint64
}

// IPSecRuleRef keeps the mapping between a ipsec encrypt/decrypt ID and its corresponding NamespaceID.
// We need this to build a look up table between the ipsec SA encrypt/decrypt rule and its refs which can be in non local namespaces
type IPSecRuleRef struct {
	NamespaceID uint64
	RuleID      uint64
}

// NaplesInfo is a read-only object containing information about Naples
type NaplesInfo struct {
	UUID                 string   `json:"naples-uuid,omitempty"`
	ControllerIPs        []string `json:"controller-ips,omitempty"`
	Mode                 string   `json:"naples-mode,omitempty"`
	IsNpmClientConnected bool     `json:"is-npm-client-connected,omitempty"`
}

// DepSolver is a netagent state object dependency solver
type DepSolver interface {
	ObjectKey(o api.ObjectMeta, t api.TypeMeta) string        // ObjectKey generates an the lookup key for an object. tenant|name for tenants, namespace|tenant|name for namespaces and kind|tenant|namespace|name for all other objects
	Add(parent interface{}, children ...interface{}) error    // Adds the parent to children dependencies.
	Remove(parent interface{}, children ...interface{}) error // Removes the parent to children dependencies
	Solve(o interface{}) error                                // Checks if an object o has an pending dependencies. Will return ErrCannotDelete if there are pending dependencies
}

// NetAgent is the network agent instance
type NetAgent struct {
	sync.Mutex                                                         // global lock for the agent
	Store                   emstore.Emstore                            // embedded db
	NodeUUID                string                                     // Node's UUID
	Datapath                NetDatapathAPI                             // network datapath
	Ctrlerif                CtrlerAPI                                  // controller object
	DelphiClient            clientApi.Client                           // delphi client
	Solver                  DepSolver                                  // Object dependency resolver
	NetworkDB               map[string]*netproto.Network               // Network object db ToDo Add updating in memory state from persisted DB in case of agent restarts
	EndpointDB              map[string]*netproto.Endpoint              // Endpoint object db
	TenantDB                map[string]*netproto.Tenant                // tenant object db
	NamespaceDB             map[string]*netproto.Namespace             // tenant object db
	NetworkSecurityPolicyDB map[string]*netproto.NetworkSecurityPolicy // Security group policy DB
	TunnelDB                map[string]*netproto.Tunnel                // Tunnel object DB
	HwIfDB                  map[string]*netproto.Interface             // Has all the Uplinks and Lifs
	AppDB                   map[string]*netproto.App                   // App DB
	VrfDB                   map[string]*netproto.Vrf
	PortDB map[string]*netproto.Port
	SecurityProfileDB       map[string]*netproto.SecurityProfile //Security Profile DB
	IPAMPolicyDB            map[string]*netproto.IPAMPolicy      // IPAM Policy DB
	RoutingConfigDB         map[string]*netproto.RoutingConfig   // RoutingConfig DB
	ControllerIPs           []string                             // Controller IPs that NetAgent is using
	Mode                    string                               // Netagent Mode
	NetAgentStartTime       time.Time                            // Time when NetAgent was started
	ArpClient               *arp.Client                          // Arp Client Handler
	ArpCache                ArpCache                             // ARP Cache
	LateralDB               map[string][]string                  // Tracks the dependencies for lateral objects
	MgmtLink                string                               // current management link name
}

// ArpCache maintains a list of ARP entries for the SPAN Dest IPs
type ArpCache struct {
	Cache     sync.Map
	DoneCache map[string]context.CancelFunc
}

// CtrlerAPI is the API provided by controller modules to netagent
type CtrlerAPI interface {
	EndpointCreateReq(epinfo *netproto.Endpoint) (*netproto.Endpoint, error) // sends an endpoint create request
	EndpointAgeoutNotif(epinfo *netproto.Endpoint) error                     // sends an endpoint ageout notification
	EndpointDeleteReq(epinfo *netproto.Endpoint) (*netproto.Endpoint, error) // sends an endpoint delete request
	IsConnected() bool                                                       // returns true if client is connected to controller
}

// CtrlerIntf is the API provided by netagent for the controllers
type CtrlerIntf interface {
	RegisterCtrlerIf(ctrlerif CtrlerAPI) error
	GetAgentID() string

	CreateNetwork(nt *netproto.Network) error                   // create a network
	UpdateNetwork(nt *netproto.Network) error                   // update a network
	DeleteNetwork(tn, ns, name string) error                    // delete a network
	ListNetwork() []*netproto.Network                           // lists all networks
	FindNetwork(meta api.ObjectMeta) (*netproto.Network, error) // finds a network

	CreateEndpoint(ep *netproto.Endpoint) error                   // create an endpoint
	UpdateEndpoint(ep *netproto.Endpoint) error                   // update an endpoint
	DeleteEndpoint(tn, ns, name string) error                     // delete an endpoint
	FindEndpoint(meta api.ObjectMeta) (*netproto.Endpoint, error) // find an endpoint
	ListEndpoint() []*netproto.Endpoint                           // list all endpoints

	CreateTenant(tn *netproto.Tenant) error                   // create a tenant
	DeleteTenant(tn, ns, name string) error                   // delete a tenant
	FindTenant(meta api.ObjectMeta) (*netproto.Tenant, error) // find a tenant
	ListTenant() []*netproto.Tenant                           // lists all tenants
	UpdateTenant(tn *netproto.Tenant) error                   // updates a tenant

	CreateNamespace(ns *netproto.Namespace) error                   // create a namespace
	DeleteNamespace(tn, ns, name string) error                      // delete a namespace
	FindNamespace(meta api.ObjectMeta) (*netproto.Namespace, error) // find a namespace
	ListNamespace() []*netproto.Namespace                           // lists all namespaces
	UpdateNamespace(ns *netproto.Namespace) error                   // updates a namespace

	CreateInterface(intf *netproto.Interface) error // creates an interface
	UpdateInterface(intf *netproto.Interface) error // updates an interface
	DeleteInterface(tn, ns, name string) error      // deletes an interface
	ListInterface() []*netproto.Interface           // lists all interfaces

	CreateNetworkSecurityPolicy(rt *netproto.NetworkSecurityPolicy) error                   // creates an Security Group Policy
	FindNetworkSecurityPolicy(meta api.ObjectMeta) (*netproto.NetworkSecurityPolicy, error) // finds an Security Group Policy
	ListNetworkSecurityPolicy() []*netproto.NetworkSecurityPolicy                           // lists Security Group Policy
	UpdateNetworkSecurityPolicy(rt *netproto.NetworkSecurityPolicy) error                   // updates an Security Group Policy
	DeleteNetworkSecurityPolicy(tn, ns, name string) error                                  // deletes an Security Group Policy

	CreateTunnel(tun *netproto.Tunnel) error                  // creates an Tunnel
	FindTunnel(meta api.ObjectMeta) (*netproto.Tunnel, error) // finds an Tunnel
	ListTunnel() []*netproto.Tunnel                           // lists Tunnel
	UpdateTunnel(tun *netproto.Tunnel) error                  // updates an Tunnel
	DeleteTunnel(tn, ns, name string) error                   // deletes an Tunnel

	CreatePort(tcp *netproto.Port) error                  // creates an Port
	FindPort(meta api.ObjectMeta) (*netproto.Port, error) // finds an Port
	ListPort() []*netproto.Port                           // lists Port
	UpdatePort(port *netproto.Port) error                 // updates an Port
	DeletePort(port, ns, name string) error               // deletes an Port

	CreateSecurityProfile(profile *netproto.SecurityProfile) error              // creates an SecurityProfile
	FindSecurityProfile(meta api.ObjectMeta) (*netproto.SecurityProfile, error) // finds an SecurityProfile
	ListSecurityProfile() []*netproto.SecurityProfile                           // lists SecurityProfile
	UpdateSecurityProfile(profile *netproto.SecurityProfile) error              // updates an SecurityProfile
	DeleteSecurityProfile(profile, ns, name string) error                       // deletes an SecurityProfile

	CreateApp(app *netproto.App) error                  // creates an App
	FindApp(meta api.ObjectMeta) (*netproto.App, error) // finds an App
	ListApp() []*netproto.App                           // lists App
	UpdateApp(app *netproto.App) error                  // updates an App
	DeleteApp(app, ns, name string) error               // deletes an App

	CreateVrf(vrf *netproto.Vrf) error                  // creates an Vrf
	FindVrf(meta api.ObjectMeta) (*netproto.Vrf, error) // finds an Vrf
	ListVrf() []*netproto.Vrf                           // lists Vrf
	UpdateVrf(vrf *netproto.Vrf) error                  // updates an Vrf
	DeleteVrf(tn, namespace, name string) error         // deletes an Vrf

	CreateIPAMPolicy(nt *netproto.IPAMPolicy) error                   // create an IPAM policy
	UpdateIPAMPolicy(nt *netproto.IPAMPolicy) error                   // update an IPAM policy
	DeleteIPAMPolicy(tn, ns, name string) error                       // delete an IPAM policy
	ListIPAMPolicy() []*netproto.IPAMPolicy                           // lists all IPAM policies
	FindIPAMPolicy(meta api.ObjectMeta) (*netproto.IPAMPolicy, error) // finds an IPAM Policy

	CreateRoutingConfig(nt *netproto.RoutingConfig) error                   // create a RoutingConfig
	UpdateRoutingConfig(nt *netproto.RoutingConfig) error                   // update a RoutingConfig
	DeleteRoutingConfig(tn, ns, name string) error                          // delete a RoutingConfig
	ListRoutingConfig() []*netproto.RoutingConfig                           // lists all RoutingConfigs
	FindRoutingConfig(meta api.ObjectMeta) (*netproto.RoutingConfig, error) // finds a RoutingConfig

	GetHwInterfaces() error              // Gets all the uplinks created on the hal by nic mgr
	GetNaplesInfo() (*NaplesInfo, error) // Returns Naples information
	GetNetagentUptime() (string, error)  // Returns NetAgent Uptime

	CreateLateralNetAgentObjects(owner string, mgmtIP, destIP string, tunnelOp bool) error // API for TSAgent and TPAgent to use to create dependent objects
	DeleteLateralNetAgentObjects(owner string, mgmtIP, destIP string, tunnelOp bool) error // API for TSAgent and TPAgent to delete dependent objects
	PurgeConfigs() error                                                                   // Deletes all netagent configs. This is called on decommission workflow/switch to network managed mode
	GetWatchOptions(cts context.Context, kind string) api.ListWatchOptions                 // Allow client to query for options to use for watch
	LifUpdateHandler(lif *halproto.LifGetResponse) error                                   // Handle async lif updates
	GetInterfaceByID(intfID uint64) (*netproto.Interface, error)                           // ID to name mapping API for interfaces
}

// PluginIntf is the API provided by the netagent to plugins
type PluginIntf interface {
	EndpointCreateReq(epinfo *netproto.Endpoint) (*netproto.Endpoint, *IntfInfo, error) // Creates an endpoint
	EndpointDeleteReq(epinfo *netproto.Endpoint) error                                  // deletes an endpoint
}

// NetDatapathAPI is the API provided by datapath modules
type NetDatapathAPI interface {
	SetAgent(ag DatapathIntf) error

	CreateLocalEndpoint(ep *netproto.Endpoint, nw *netproto.Network, lifID, enicID uint64, vrf *netproto.Vrf) (*IntfInfo, error)
	UpdateLocalEndpoint(ep *netproto.Endpoint, nw *netproto.Network, lifID, enicID uint64, vrf *netproto.Vrf) error
	DeleteLocalEndpoint(ep *netproto.Endpoint, nw *netproto.Network, enicID uint64) error
	CreateRemoteEndpoint(ep *netproto.Endpoint, nw *netproto.Network, uplinkID uint64, vrf *netproto.Vrf) error
	UpdateRemoteEndpoint(ep *netproto.Endpoint, nw *netproto.Network) error
	DeleteRemoteEndpoint(ep *netproto.Endpoint, nw *netproto.Network) error

	CreateNetwork(nw *netproto.Network, uplinks []*netproto.Interface, vrf *netproto.Vrf) error
	UpdateNetwork(nw *netproto.Network, uplinks []*netproto.Interface, vrf *netproto.Vrf) error
	DeleteNetwork(nw *netproto.Network, uplinks []*netproto.Interface, vrf *netproto.Vrf) error

	CreateVrf(vrfID uint64, vrfType string) error
	DeleteVrf(vrfID uint64) error
	UpdateVrf(vrfID uint64) error

	CreateInterface(intfs ...*netproto.Interface) error
	UpdateInterface(intf *netproto.Interface) error
	DeleteInterface(intf *netproto.Interface) error
	ListInterfaces() ([]*netproto.Interface, error)

	CreateNetworkSecurityPolicy(sgp *netproto.NetworkSecurityPolicy, vrfID uint64, ruleIDAppLUT *sync.Map) error
	UpdateNetworkSecurityPolicy(sgp *netproto.NetworkSecurityPolicy, vrfID uint64, ruleIDAppLUT *sync.Map) error
	DeleteNetworkSecurityPolicy(sgp *netproto.NetworkSecurityPolicy, vrfID uint64) error

	CreateTunnel(tun *netproto.Tunnel, vrf *netproto.Vrf) error
	UpdateTunnel(tun *netproto.Tunnel, vrf *netproto.Vrf) error
	DeleteTunnel(tun *netproto.Tunnel, vrf *netproto.Vrf) error

	CreatePort(ports ...*netproto.Port) error
	UpdatePort(port *netproto.Port) (*netproto.Port, error)
	DeletePort(port *netproto.Port) error

	CreateSecurityProfile(profile *netproto.SecurityProfile, attachmentVrfs []*netproto.Vrf) error
	UpdateSecurityProfile(profile *netproto.SecurityProfile, attachmentVrfs []*netproto.Vrf) error
	DeleteSecurityProfile(profile *netproto.SecurityProfile, attachmentVrfs []*netproto.Vrf) error

	RegisterStateAPI(stateAPI CtrlerIntf)
	GetUUID() (string, error)
}

// DatapathIntf is the API provided by the netagent to datapaths
type DatapathIntf interface {
	GetAgentID() string
	EndpointCreateReq(epinfo *netproto.Endpoint) (*netproto.Endpoint, *IntfInfo, error) // Creates an endpoint
	EndpointDeleteReq(epinfo *netproto.Endpoint) error                                  // deletes an endpoint
}
