// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

package apiclient

import (
	auth "github.com/pensando/sw/api/generated/auth"
	authClient "github.com/pensando/sw/api/generated/auth/grpc/client"
	bookstore "github.com/pensando/sw/api/generated/bookstore"
	bookstoreClient "github.com/pensando/sw/api/generated/bookstore/grpc/client"
	cluster "github.com/pensando/sw/api/generated/cluster"
	clusterClient "github.com/pensando/sw/api/generated/cluster/grpc/client"
	monitoring "github.com/pensando/sw/api/generated/monitoring"
	monitoringClient "github.com/pensando/sw/api/generated/monitoring/grpc/client"
	network "github.com/pensando/sw/api/generated/network"
	networkClient "github.com/pensando/sw/api/generated/network/grpc/client"
	objstore "github.com/pensando/sw/api/generated/objstore"
	objstoreClient "github.com/pensando/sw/api/generated/objstore/grpc/client"
	rollout "github.com/pensando/sw/api/generated/rollout"
	rolloutClient "github.com/pensando/sw/api/generated/rollout/grpc/client"
	security "github.com/pensando/sw/api/generated/security"
	securityClient "github.com/pensando/sw/api/generated/security/grpc/client"
	staging "github.com/pensando/sw/api/generated/staging"
	stagingClient "github.com/pensando/sw/api/generated/staging/grpc/client"
	workload "github.com/pensando/sw/api/generated/workload"
	workloadClient "github.com/pensando/sw/api/generated/workload/grpc/client"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// APIGroup is an API Group name
type APIGroup string

const (
	GroupAuth       APIGroup = "auth"
	GroupBookstore  APIGroup = "bookstore"
	GroupCluster    APIGroup = "cluster"
	GroupMonitoring APIGroup = "monitoring"
	GroupNetwork    APIGroup = "network"
	GroupObjstore   APIGroup = "objstore"
	GroupRollout    APIGroup = "rollout"
	GroupSecurity   APIGroup = "security"
	GroupStaging    APIGroup = "staging"
	GroupWorkload   APIGroup = "workload"
)

// Services is list of all services exposed by the client ---
type Services interface {
	Close() error

	// Package is auth and len of messages is 4
	AuthV1() auth.AuthV1Interface
	// Package is bookstore and len of messages is 6
	BookstoreV1() bookstore.BookstoreV1Interface
	// Package is cluster and len of messages is 5
	ClusterV1() cluster.ClusterV1Interface
	// Package is monitoring and len of messages is 10
	MonitoringV1() monitoring.MonitoringV1Interface
	// Package is network and len of messages is 3
	NetworkV1() network.NetworkV1Interface
	// Package is objstore and len of messages is 2
	ObjstoreV1() objstore.ObjstoreV1Interface
	// Package is rollout and len of messages is 1
	RolloutV1() rollout.RolloutV1Interface
	// Package is security and len of messages is 6
	SecurityV1() security.SecurityV1Interface
	// Package is staging and len of messages is 1
	StagingV1() staging.StagingV1Interface
	// Package is workload and len of messages is 2
	WorkloadV1() workload.WorkloadV1Interface
}

type apiGrpcServerClient struct {
	url    string
	logger log.Logger
	client *rpckit.RPCClient

	aAuthV1       auth.AuthV1Interface
	aBookstoreV1  bookstore.BookstoreV1Interface
	aClusterV1    cluster.ClusterV1Interface
	aMonitoringV1 monitoring.MonitoringV1Interface
	aNetworkV1    network.NetworkV1Interface
	aObjstoreV1   objstore.ObjstoreV1Interface
	aRolloutV1    rollout.RolloutV1Interface
	aSecurityV1   security.SecurityV1Interface
	aStagingV1    staging.StagingV1Interface
	aWorkloadV1   workload.WorkloadV1Interface
}

// Close closes the client
func (a *apiGrpcServerClient) Close() error {
	return a.client.Close()
}

func (a *apiGrpcServerClient) AuthV1() auth.AuthV1Interface {
	return a.aAuthV1
}

func (a *apiGrpcServerClient) BookstoreV1() bookstore.BookstoreV1Interface {
	return a.aBookstoreV1
}

func (a *apiGrpcServerClient) ClusterV1() cluster.ClusterV1Interface {
	return a.aClusterV1
}

func (a *apiGrpcServerClient) MonitoringV1() monitoring.MonitoringV1Interface {
	return a.aMonitoringV1
}

func (a *apiGrpcServerClient) NetworkV1() network.NetworkV1Interface {
	return a.aNetworkV1
}

func (a *apiGrpcServerClient) ObjstoreV1() objstore.ObjstoreV1Interface {
	return a.aObjstoreV1
}

func (a *apiGrpcServerClient) RolloutV1() rollout.RolloutV1Interface {
	return a.aRolloutV1
}

func (a *apiGrpcServerClient) SecurityV1() security.SecurityV1Interface {
	return a.aSecurityV1
}

func (a *apiGrpcServerClient) StagingV1() staging.StagingV1Interface {
	return a.aStagingV1
}

func (a *apiGrpcServerClient) WorkloadV1() workload.WorkloadV1Interface {
	return a.aWorkloadV1
}

// NewGrpcAPIClient returns a gRPC client
func NewGrpcAPIClient(clientName, url string, logger log.Logger, opts ...rpckit.Option) (Services, error) {
	client, err := rpckit.NewRPCClient(clientName, url, opts...)
	if err != nil {
		logger.ErrorLog("msg", "Failed to connect to gRPC server", "URL", url, "error", err)
		return nil, err
	}
	return &apiGrpcServerClient{
		url:    url,
		client: client,
		logger: logger,

		aAuthV1:       authClient.NewGrpcCrudClientAuthV1(client.ClientConn, logger),
		aBookstoreV1:  bookstoreClient.NewGrpcCrudClientBookstoreV1(client.ClientConn, logger),
		aClusterV1:    clusterClient.NewGrpcCrudClientClusterV1(client.ClientConn, logger),
		aMonitoringV1: monitoringClient.NewGrpcCrudClientMonitoringV1(client.ClientConn, logger),
		aNetworkV1:    networkClient.NewGrpcCrudClientNetworkV1(client.ClientConn, logger),
		aObjstoreV1:   objstoreClient.NewGrpcCrudClientObjstoreV1(client.ClientConn, logger),
		aRolloutV1:    rolloutClient.NewGrpcCrudClientRolloutV1(client.ClientConn, logger),
		aSecurityV1:   securityClient.NewGrpcCrudClientSecurityV1(client.ClientConn, logger),
		aStagingV1:    stagingClient.NewGrpcCrudClientStagingV1(client.ClientConn, logger),
		aWorkloadV1:   workloadClient.NewGrpcCrudClientWorkloadV1(client.ClientConn, logger),
	}, nil
}

type apiRestServerClient struct {
	url    string
	logger log.Logger

	aAuthV1       auth.AuthV1Interface
	aBookstoreV1  bookstore.BookstoreV1Interface
	aClusterV1    cluster.ClusterV1Interface
	aMonitoringV1 monitoring.MonitoringV1Interface
	aNetworkV1    network.NetworkV1Interface
	aObjstoreV1   objstore.ObjstoreV1Interface
	aRolloutV1    rollout.RolloutV1Interface
	aSecurityV1   security.SecurityV1Interface
	aStagingV1    staging.StagingV1Interface
	aWorkloadV1   workload.WorkloadV1Interface
}

// Close closes the client
func (a *apiRestServerClient) Close() error {
	return nil
}

func (a *apiRestServerClient) AuthV1() auth.AuthV1Interface {
	return a.aAuthV1
}

func (a *apiRestServerClient) BookstoreV1() bookstore.BookstoreV1Interface {
	return a.aBookstoreV1
}

func (a *apiRestServerClient) ClusterV1() cluster.ClusterV1Interface {
	return a.aClusterV1
}

func (a *apiRestServerClient) MonitoringV1() monitoring.MonitoringV1Interface {
	return a.aMonitoringV1
}

func (a *apiRestServerClient) NetworkV1() network.NetworkV1Interface {
	return a.aNetworkV1
}

func (a *apiRestServerClient) ObjstoreV1() objstore.ObjstoreV1Interface {
	return a.aObjstoreV1
}

func (a *apiRestServerClient) RolloutV1() rollout.RolloutV1Interface {
	return a.aRolloutV1
}

func (a *apiRestServerClient) SecurityV1() security.SecurityV1Interface {
	return a.aSecurityV1
}

func (a *apiRestServerClient) StagingV1() staging.StagingV1Interface {
	return a.aStagingV1
}

func (a *apiRestServerClient) WorkloadV1() workload.WorkloadV1Interface {
	return a.aWorkloadV1
}

// NewRestAPIClient returns a REST client
func NewRestAPIClient(url string) (Services, error) {
	return &apiRestServerClient{
		url:    url,
		logger: log.WithContext("module", "RestAPIClient"),

		aAuthV1:       authClient.NewRestCrudClientAuthV1(url),
		aBookstoreV1:  bookstoreClient.NewRestCrudClientBookstoreV1(url),
		aClusterV1:    clusterClient.NewRestCrudClientClusterV1(url),
		aMonitoringV1: monitoringClient.NewRestCrudClientMonitoringV1(url),
		aNetworkV1:    networkClient.NewRestCrudClientNetworkV1(url),
		aObjstoreV1:   objstoreClient.NewRestCrudClientObjstoreV1(url),
		aRolloutV1:    rolloutClient.NewRestCrudClientRolloutV1(url),
		aSecurityV1:   securityClient.NewRestCrudClientSecurityV1(url),
		aStagingV1:    stagingClient.NewRestCrudClientStagingV1(url),
		aWorkloadV1:   workloadClient.NewRestCrudClientWorkloadV1(url),
	}, nil
}

// NewStagedRestAPIClient returns a REST client
func NewStagedRestAPIClient(url string, bufferId string) (Services, error) {
	return &apiRestServerClient{
		url:    url,
		logger: log.WithContext("module", "RestAPIClient"),

		aAuthV1:       authClient.NewStagedRestCrudClientAuthV1(url, bufferId),
		aBookstoreV1:  bookstoreClient.NewStagedRestCrudClientBookstoreV1(url, bufferId),
		aClusterV1:    clusterClient.NewStagedRestCrudClientClusterV1(url, bufferId),
		aMonitoringV1: monitoringClient.NewStagedRestCrudClientMonitoringV1(url, bufferId),
		aNetworkV1:    networkClient.NewStagedRestCrudClientNetworkV1(url, bufferId),
		aObjstoreV1:   objstoreClient.NewStagedRestCrudClientObjstoreV1(url, bufferId),
		aRolloutV1:    rolloutClient.NewStagedRestCrudClientRolloutV1(url, bufferId),
		aSecurityV1:   securityClient.NewStagedRestCrudClientSecurityV1(url, bufferId),
		aStagingV1:    stagingClient.NewStagedRestCrudClientStagingV1(url, bufferId),
		aWorkloadV1:   workloadClient.NewStagedRestCrudClientWorkloadV1(url, bufferId),
	}, nil
}
