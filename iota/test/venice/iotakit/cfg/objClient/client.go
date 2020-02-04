package objClient

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/cluster"
	"github.com/pensando/sw/api/generated/diagnostics"
	"github.com/pensando/sw/api/generated/monitoring"
	"github.com/pensando/sw/api/generated/network"
	"github.com/pensando/sw/api/generated/objstore"
	"github.com/pensando/sw/api/generated/security"
	"github.com/pensando/sw/api/generated/workload"
	loginctx "github.com/pensando/sw/api/login/context"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/netutils"
	"github.com/pensando/sw/venice/utils/workfarm"
)

//ObjClient to do operations for test
type ObjClient interface {
	Context() context.Context
	Urls() []string

	CreateHost(host *cluster.Host) error
	ListHost() (objs []*cluster.Host, err error)
	DeleteHost(wrkld *cluster.Host) error

	CreateNetwork(obj *network.Network) error
	ListNetwork() (objs []*network.Network, err error)

	CreateNetworkSecurityPolicy(sgp *security.NetworkSecurityPolicy) error
	GetNetworkSecurityPolicy(meta *api.ObjectMeta) (sgp *security.NetworkSecurityPolicy, err error)
	ListNetworkSecurityPolicy() (objs []*security.NetworkSecurityPolicy, err error)
	DeleteNetworkSecurityPolicy(sgp *security.NetworkSecurityPolicy) error
	UpdateNetworkSecurityPolicy(sgp *security.NetworkSecurityPolicy) error

	CreateApp(app *security.App) error
	ListApp() (objs []*security.App, err error)
	DeleteApp(app *security.App) error

	CreateWorkloads(wrklds []*workload.Workload) error
	GetWorkload(meta *api.ObjectMeta) (w *workload.Workload, err error)
	ListWorkload() (objs []*workload.Workload, err error)
	DeleteWorkloads(wrklds []*workload.Workload) error

	GetEndpoint(meta *api.ObjectMeta) (ep *workload.Endpoint, err error)
	ListEndpoints(tenant string) (eps []*workload.Endpoint, err error)

	ListFirewallProfile() (objs []*security.FirewallProfile, err error)
	UpdateFirewallProfile(fwp *security.FirewallProfile) error
	DeleteFirewallProfile(fwprofile *security.FirewallProfile) error

	CreateMirrorSession(msp *monitoring.MirrorSession) error
	UpdateMirrorSession(msp *monitoring.MirrorSession) error
	DeleteMirrorSession(msp *monitoring.MirrorSession) error

	GetSmartNIC(name string) (sn *cluster.DistributedServiceCard, err error)
	ListSmartNIC() (snl []*cluster.DistributedServiceCard, err error)
	UpdateSmartNIC(sn *cluster.DistributedServiceCard) error
	GetSmartNICByName(snicName string) (sn *cluster.DistributedServiceCard, err error)

	AddClusterNode(node *cluster.Node) (err error)
	DeleteClusterNode(node *cluster.Node) (err error)
	ListClusterNodes() (snl []*cluster.Node, err error)
	GetCluster() (cl *cluster.Cluster, err error)

	TakeConfigSnapshot(reqname string) (uri string, err error)
	RestoreConfig(filename string) error
	ConfigureSnapshot() error

	PullConfigStatus(configStatus interface{}) error

	ListObjectStoreObjects() (objs []*objstore.Object, err error)
}

type VeniceConfigStatus struct {
	KindObjects struct {
		App                   int `json:"App"`
		Endpoint              int `json:"Endpoint"`
		NetworkSecurityPolicy int `json:"NetworkSecurityPolicy"`
	} `json:"KindObjects"`
	NodesStatus []struct {
		NodeID     string `json:"NodeID"`
		KindStatus struct {
			App struct {
				Status struct {
					Create bool `json:"create-event"`
					Update bool `json:"update-event"`
					Delete bool `json:"delete-event"`
				} `json:"Status"`
			} `json:"App"`
			Endpoint struct {
				Status struct {
					Create bool `json:"create-event"`
					Update bool `json:"update-event"`
					Delete bool `json:"delete-event"`
				} `json:"Status"`
			} `json:"Endpoint"`
			SgPolicy struct {
				Status struct {
					Create bool `json:"create-event"`
					Update bool `json:"update-event"`
					Delete bool `json:"delete-event"`
				} `json:"Status"`
			} `json:"NetworkSecurityPolicy"`
		} `json:"KindStatus"`
	} `json:"NodesStatus"`
}

type VeniceRawData struct {
	Diagnostics struct {
		String string `json:"string"`
	} `json:"diagnostics"`
}

//Client rest client handler
type Client struct {
	ctx     context.Context
	restcls []apiclient.Services
	urls    []string
}

//NewClient rest client
func NewClient(ctx context.Context, urls []string) ObjClient {
	client := &Client{ctx: ctx, urls: urls}
	client.init()
	return client
}

//Context for rest operation
func (r *Client) Context() context.Context {
	return r.ctx
}

//Context for rest operation
func (r *Client) Urls() []string {
	return r.urls
}

//Clients clients for operations
func (r *Client) Clients() []apiclient.Services {
	return r.restcls
}

func (r *Client) newRestClient() ([]apiclient.Services, error) {

	var restcls []apiclient.Services
	for _, url := range r.urls {
		restcl, err := apiclient.NewRestAPIClient(url)
		if err != nil {
			log.Errorf("Error connecting to Venice %v. Err: %v", url, err)
			return nil, err
		}
		restcls = append(r.restcls, restcl)
	}

	return restcls, nil
}

func (r *Client) init() error {
	var err error
	r.restcls, err = r.newRestClient()

	return err
}

func (r *Client) CreateHost(host *cluster.Host) error {

	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.ClusterV1().Host().Create(r.ctx, host)
		if err == nil {
			break
		} else if strings.Contains(err.Error(), "already exists") {
			_, err = restcl.ClusterV1().Host().Update(r.ctx, host)
			if err == nil {
				break
			}
		}
	}
	return err

}

// CreateNetwork creates an Network in venice
func (r *Client) CreateNetwork(obj *network.Network) error {

	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.NetworkV1().Network().Create(r.ctx, obj)
		if err == nil {
			break
		}
	}

	return err
}

//CreateNetworkSecurityPolicy create policy
func (r *Client) CreateNetworkSecurityPolicy(sgp *security.NetworkSecurityPolicy) error {
	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.SecurityV1().NetworkSecurityPolicy().Create(r.ctx, sgp)
		if err == nil {
			break
		} else if strings.Contains(err.Error(), "already exists") {
			_, err = restcl.SecurityV1().NetworkSecurityPolicy().Update(r.ctx, sgp)
			if err == nil {
				break
			}
		}
	}

	return err
}

// UpdateNetworkSecurityPolicy updates an SG policy
func (r *Client) UpdateNetworkSecurityPolicy(sgp *security.NetworkSecurityPolicy) error {
	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.SecurityV1().NetworkSecurityPolicy().Update(r.ctx, sgp)
		if err == nil {
			break
		}
	}
	return err
}

//CreateApp create Apps
func (r *Client) CreateApp(app *security.App) error {

	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.SecurityV1().App().Create(r.ctx, app)
		if err == nil {
			break
		} else if strings.Contains(err.Error(), "already exists") {
			_, err = restcl.SecurityV1().App().Update(r.ctx, app)
			if err == nil {
				break
			}
		}
	}
	return err

}

func workloadWork(ctx context.Context, id, iter int, userCtx interface{}) error {

	var err error
	wctx := userCtx.(*workCtx)

	workloads := wctx.objs.([]*workload.Workload)
	if restcls, ok := wctx.restCls[id]; ok {
		for _, restcl := range restcls {
			_, err = restcl.WorkloadV1().Workload().Create(wctx.ctx, workloads[iter])
			if err == nil {
				break
			} else if strings.Contains(err.Error(), "already exists") {
				_, err = restcl.WorkloadV1().Workload().Update(wctx.ctx, workloads[iter])
				if err == nil {
					break
				}
			}
		}
	}

	if err != nil {
		log.Errorf("Workload create %v failed  with error %v", workloads[iter], err.Error())
	}
	return err
}

func workloadDeleteWork(ctx context.Context, id, iter int, userCtx interface{}) error {

	var err error
	wctx := userCtx.(*workCtx)

	workloads := wctx.objs.([]*workload.Workload)
	if restcls, ok := wctx.restCls[id]; ok {
		for _, restcl := range restcls {
			_, err = restcl.WorkloadV1().Workload().Delete(wctx.ctx, &workloads[iter].ObjectMeta)
			if err == nil {
				break
			}
		}
		if err != nil {
			log.Errorf("Workload  delete %v failed  with error %v", workloads[iter], err.Error())
			return err
		}
	}
	return err
}

// GetWorkload returns venice workload by object meta
func (r *Client) GetWorkload(meta *api.ObjectMeta) (w *workload.Workload, err error) {

	for _, restcl := range r.restcls {
		w, err = restcl.WorkloadV1().Workload().Get(r.ctx, meta)
		if err == nil {
			break
		}
	}

	return w, err
}

// CreateWorkloads creates workloads
func (r *Client) CreateWorkloads(wrklds []*workload.Workload) error {
	wCtx := &workCtx{
		objs: wrklds,
		len:  len(wrklds),
		ctx:  r.ctx,
	}
	return r.parallelPush(wCtx, workloadWork)
}

// GetNetworkSecurityPolicy gets NetworkSecurityPolicy from venice cluster
func (r *Client) GetNetworkSecurityPolicy(meta *api.ObjectMeta) (sgp *security.NetworkSecurityPolicy, err error) {
	for _, restcl := range r.restcls {
		sgp, err = restcl.SecurityV1().NetworkSecurityPolicy().Get(r.ctx, meta)
		if err == nil {
			break
		}
	}

	return sgp, err
}

type workCtx struct {
	objs    interface{}
	restCls map[int][]apiclient.Services
	len     int
	ctx     context.Context
}

func (r *Client) parallelPush(wctx *workCtx, opFunc workfarm.WorkFunc) error {

	numOfWorkers := 50

	wctx.restCls = make(map[int][]apiclient.Services)
	for i := 0; i < numOfWorkers; i++ {
		restcls, err := r.newRestClient()
		if err != nil {
			return err
		}
		wctx.restCls[i] = restcls
	}

	defer func() {
		for _, restClients := range wctx.restCls {
			for _, restClient := range restClients {
				go restClient.Close()
			}
		}
	}()

	farm := workfarm.New(int(numOfWorkers), time.Minute*20, opFunc)

	log.Infof("Number of workers %v", numOfWorkers)
	ch, err := farm.Run(context.Background(), wctx.len, 0, math.MaxUint32, wctx)
	if err != nil {
		fmt.Printf("failed to start work (%s)\n", err)
	}

	rslts := <-ch

	if rslts.WorkerErrors != 0 {
		return fmt.Errorf("Workload create failed stats %+v", rslts)
	}
	return nil
}

// ListHost gets all hosts from venice cluster
func (r *Client) ListHost() (objs []*cluster.Host, err error) {
	opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: globals.DefaultTenant}}

	for _, restcl := range r.restcls {
		objs, err = restcl.ClusterV1().Host().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}

	return objs, err
}

// ListNetworkSecurityPolicy gets all SGPolicies from venice cluster
func (r *Client) ListNetworkSecurityPolicy() (objs []*security.NetworkSecurityPolicy, err error) {
	opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: globals.DefaultTenant}}

	for _, restcl := range r.restcls {
		objs, err = restcl.SecurityV1().NetworkSecurityPolicy().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}

	return objs, err
}

// ListNetwork gets all networks from venice cluster
func (r *Client) ListNetwork() (objs []*network.Network, err error) {

	opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: globals.DefaultTenant}}

	for _, restcl := range r.restcls {
		objs, err = restcl.NetworkV1().Network().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}

	return objs, err
}

// ListApp gets all apps from venice cluster
func (r *Client) ListApp() (objs []*security.App, err error) {

	opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: globals.DefaultTenant}}

	for _, restcl := range r.restcls {
		objs, err = restcl.SecurityV1().App().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}

	return objs, err
}

// ListWorkload gets all workloads from venice cluster
func (r *Client) ListWorkload() (objs []*workload.Workload, err error) {

	opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: globals.DefaultTenant}}

	for _, restcl := range r.restcls {
		objs, err = restcl.WorkloadV1().Workload().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}

	return objs, err
}

// DeleteNetworkSecurityPolicy deletes SG policy
func (r *Client) DeleteNetworkSecurityPolicy(sgp *security.NetworkSecurityPolicy) error {
	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.SecurityV1().NetworkSecurityPolicy().Delete(r.ctx, &sgp.ObjectMeta)
		if err == nil {
			break
		} else {
			log.Errorf("Error deleting object %v", err)
		}
	}

	return err
}

// DeleteApp deletes App object
func (r *Client) DeleteApp(app *security.App) error {

	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.SecurityV1().App().Delete(r.ctx, &app.ObjectMeta)
		if err == nil {
			break
		} else {
			log.Errorf("Error deleting object %v", err)
		}
	}

	return err
}

// DeleteWorkloads creates workloads
func (r *Client) DeleteWorkloads(wrklds []*workload.Workload) error {
	wCtx := &workCtx{
		objs: wrklds,
		len:  len(wrklds),
		ctx:  r.ctx,
	}
	return r.parallelPush(wCtx, workloadDeleteWork)
}

//DeleteHost deletes host object
func (r *Client) DeleteHost(wrkld *cluster.Host) error {

	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.ClusterV1().Host().Delete(r.ctx, &wrkld.ObjectMeta)
		if err == nil {
			break
		} else {
			log.Errorf("Error deleting object %v", err)
		}
	}
	return err
}

// GetCluster gets the venice cluster object
func (r *Client) GetCluster() (cl *cluster.Cluster, err error) {
	for _, restcl := range r.restcls {
		cl, err = restcl.ClusterV1().Cluster().Get(r.ctx, &api.ObjectMeta{Name: "iota-cluster"})
		if err == nil {
			break
		}
	}

	return cl, err
}

// ConfigureSnapshot preforms a snapshot operation
func (r *Client) ConfigureSnapshot() error {
	var err error
	cfg := &cluster.ConfigurationSnapshot{
		ObjectMeta: api.ObjectMeta{
			Name: "GlobalSnapshotConfig",
		},
		Spec: cluster.ConfigurationSnapshotSpec{
			Destination: cluster.SnapshotDestination{
				Type: cluster.SnapshotDestinationType_ObjectStore.String(),
			},
		},
	}

	for _, restcl := range r.restcls {
		_, err = restcl.ClusterV1().ConfigurationSnapshot().Create(r.ctx, cfg)
		if err != nil {
			_, err = restcl.ClusterV1().ConfigurationSnapshot().Update(r.ctx, cfg)
		}

		if err == nil {
			break
		}
	}

	return err
}

// TakeConfigSnapshot preforms a snapshot operation
func (r *Client) TakeConfigSnapshot(reqname string) (uri string, err error) {
	req := &cluster.ConfigurationSnapshotRequest{}
	req.Name = reqname

	for _, restcl := range r.restcls {
		_, err = restcl.ClusterV1().ConfigurationSnapshot().Save(r.ctx, req)
		if err != nil {
			return "", err
		}
		snaps, err := restcl.ClusterV1().ConfigurationSnapshot().Get(r.ctx, &api.ObjectMeta{})
		if err != nil {
			return "", err
		}

		if err == nil {
			return snaps.Status.LastSnapshot.URI, nil
		}
	}

	return "", err
}

// RestoreConfig restores config to snapshot specified in filename
func (r *Client) RestoreConfig(filename string) error {
	req := &cluster.SnapshotRestore{
		ObjectMeta: api.ObjectMeta{
			Name: "IOTARestoreOp",
		},
		Spec: cluster.SnapshotRestoreSpec{
			SnapshotPath: filename,
		},
	}

	var err error
	for _, restcl := range r.restcls {
		resp, err1 := restcl.ClusterV1().SnapshotRestore().Restore(r.ctx, req)
		if err1 != nil {
			continue
		}
		if resp.Status.Status != cluster.SnapshotRestoreStatus_Completed.String() {
			err = errors.New("Restore operation did not complete")
		}
		if err1 == nil {
			return nil
		}
		err = err1
	}

	return err
}

// GetSmartNIC returns venice smartnic object
func (r *Client) GetSmartNIC(name string) (sn *cluster.DistributedServiceCard, err error) {

	meta := api.ObjectMeta{
		Name: name,
	}

	for _, restcl := range r.restcls {
		sn, err = restcl.ClusterV1().DistributedServiceCard().Get(r.ctx, &meta)
		if err == nil {
			break
		}
	}

	return sn, err
}

// ListClusterNodes gets a list of nodes
func (r *Client) ListClusterNodes() (snl []*cluster.Node, err error) {

	opts := api.ListWatchOptions{}

	for _, restcl := range r.restcls {
		snl, err = restcl.ClusterV1().Node().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}
	return snl, err
}

// DeleteClusterNode gets a list of nodes
func (r *Client) DeleteClusterNode(node *cluster.Node) (err error) {

	log.Info("Initiating delete..")
	for _, restcl := range r.restcls {
		_, err = restcl.ClusterV1().Node().Delete(r.ctx, &node.ObjectMeta)
		if err == nil {
			break
		}
	}
	log.Info("Initiating deleted competed..")
	if err != nil {
		log.Errorf("Error deleting cluster nodeß %v", err)
	}
	return err
}

// AddClusterNode gets a list of nodes
func (r *Client) AddClusterNode(node *cluster.Node) (err error) {
	for _, restcl := range r.restcls {
		newNode := &cluster.Node{
			ObjectMeta: api.ObjectMeta{
				Name: node.ObjectMeta.Name,
			},
		}
		_, err = restcl.ClusterV1().Node().Create(r.ctx, newNode)
		if err == nil {
			break
		}
	}
	return err
}

// ListSmartNIC gets a list of smartnics
func (r *Client) ListSmartNIC() (snl []*cluster.DistributedServiceCard, err error) {
	opts := api.ListWatchOptions{}

	for _, restcl := range r.restcls {
		snl, err = restcl.ClusterV1().DistributedServiceCard().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}

	return snl, err
}

// UpdateSmartNIC updates an SmartNIC object
func (r *Client) UpdateSmartNIC(sn *cluster.DistributedServiceCard) error {
	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.ClusterV1().DistributedServiceCard().Update(r.ctx, sn)
		if err == nil {
			break
		}
	}
	return err
}

func (r *Client) GetSmartNICByName(snicName string) (sn *cluster.DistributedServiceCard, err error) {
	snicList, err := r.ListSmartNIC()
	if err != nil {
		return nil, err
	}

	// walk all smartnics and see if the mac addr range matches
	for _, snic := range snicList {
		if snic.Spec.ID == snicName {
			return snic, nil
		}
	}

	return nil, fmt.Errorf("Could not find smartnic with name %s", snicName)
}

// GetEndpoint returns the endpoint
func (r *Client) GetEndpoint(meta *api.ObjectMeta) (ep *workload.Endpoint, err error) {

	for _, restcl := range r.restcls {
		ep, err = restcl.WorkloadV1().Endpoint().Get(r.ctx, meta)
		if err == nil {
			break
		}
	}
	return ep, err
}

// ListEndpoints returns list of endpoints known to Venice
func (r *Client) ListEndpoints(tenant string) (eps []*workload.Endpoint, err error) {
	opts := api.ListWatchOptions{}
	opts.Tenant = tenant
	for _, restcl := range r.restcls {
		eps, err = restcl.WorkloadV1().Endpoint().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}
	return eps, err
}

// UpdateFirewallProfile updates firewall profile
func (r *Client) UpdateFirewallProfile(fwp *security.FirewallProfile) error {

	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.SecurityV1().FirewallProfile().Update(r.ctx, fwp)
		if err == nil {
			break
		} else if strings.Contains(err.Error(), "already exists") {
			_, err = restcl.SecurityV1().FirewallProfile().Update(r.ctx, fwp)
			if err == nil {
				break
			}
		}
	}

	return err
}

// ListFirewallProfile gets all fw profile apps from venice cluster
func (r *Client) ListFirewallProfile() (objs []*security.FirewallProfile, err error) {
	opts := api.ListWatchOptions{ObjectMeta: api.ObjectMeta{Tenant: globals.DefaultTenant}}

	for _, restcl := range r.restcls {
		objs, err = restcl.SecurityV1().FirewallProfile().List(r.ctx, &opts)
		if err == nil {
			break
		}
	}

	return objs, err
}

// DeleteFirewallProfile deletes FirewallProfile object
func (r *Client) DeleteFirewallProfile(fwprofile *security.FirewallProfile) error {
	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.SecurityV1().FirewallProfile().Delete(r.ctx, &fwprofile.ObjectMeta)
		if err == nil {
			break
		}
	}

	return err
}

// CreateMirrorSession creates Mirror policy
func (r *Client) CreateMirrorSession(msp *monitoring.MirrorSession) error {
	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.MonitoringV1().MirrorSession().Create(r.ctx, msp)
		if err == nil {
			break
		} else if strings.Contains(err.Error(), "already exists") {
			_, err = restcl.MonitoringV1().MirrorSession().Update(r.ctx, msp)
			if err == nil {
				break
			}
		}
	}

	return err
}

// UpdateMirrorSession updates an Mirror policy
func (r *Client) UpdateMirrorSession(msp *monitoring.MirrorSession) error {
	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.MonitoringV1().MirrorSession().Update(r.ctx, msp)
		if err == nil {
			break
		}
	}
	return err
}

// DeleteMirrorSession deletes Mirror policy
func (r *Client) DeleteMirrorSession(msp *monitoring.MirrorSession) error {

	var err error
	for _, restcl := range r.restcls {
		_, err = restcl.MonitoringV1().MirrorSession().Delete(r.ctx, &msp.ObjectMeta)
		if err == nil {
			break
		}
	}

	return err
}

// ListObjectStoreObjects list object store objects
func (r *Client) ListObjectStoreObjects() (objs []*objstore.Object, err error) {

	for _, restcl := range r.restcls {
		objs, err = restcl.ObjstoreV1().Object().List(r.ctx, &api.ListWatchOptions{})
		if err == nil {
			break
		}
	}

	return objs, err
}

//GetNpmDebugModuleURLs gets npm debug module
func (r *Client) GetNpmDebugModuleURLs() (urls []string, err error) {
	for _, restcl := range r.restcls {
		data, err := restcl.DiagnosticsV1().Module().List(r.ctx, &api.ListWatchOptions{})
		if err == nil {
			for _, module := range data {
				if strings.Contains(module.ObjectMeta.Name, "pen-npm") {
					for _, veniceURL := range r.urls {
						urls = append(urls, "https://"+veniceURL+module.GetSelfLink()+"/Debug")
					}
				}
			}
		}
	}

	if len(urls) == 0 {
		return nil, errors.New("Could not find NPM debug URL")
	}
	return urls, nil
}

func (r *Client) doConfigPostAction(action string, configStatus interface{}) error {

	npmURLs, err := r.GetNpmDebugModuleURLs()
	if err != nil {
		return errors.New("Npm debug URL not found")
	}

	req := &diagnostics.DiagnosticsRequest{
		Query:      "action",
		Parameters: map[string]string{"action": action}}

	restcl := netutils.NewHTTPClient()
	restcl.WithTLSConfig(&tls.Config{InsecureSkipVerify: true})
	restcl.DisableKeepAlives()
	defer restcl.CloseIdleConnections()

	// get authz header
	authzHeader, ok := loginctx.AuthzHeaderFromContext(r.ctx)
	if !ok {
		return fmt.Errorf("no authorization header in context")
	}
	restcl.SetHeader("Authorization", authzHeader)
	for _, url := range npmURLs {
		_, err = restcl.Req("POST", url, req, configStatus)
		if err == nil {
			return nil
		}
		fmt.Printf("Error in request %+v\n", err)

	}

	return fmt.Errorf("Failed Request for config push : %v", err)
}

//PullConfigStatus pulls config status
func (r *Client) PullConfigStatus(configStatus interface{}) error {

	return r.doConfigPostAction("config-status", configStatus)
}