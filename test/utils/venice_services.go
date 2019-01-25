// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"time"

	"github.com/pensando/sw/api/client"
	"github.com/pensando/sw/api/generated/auth"
	loginctx "github.com/pensando/sw/api/login/context"
	"github.com/pensando/sw/venice/apigw"
	"github.com/pensando/sw/venice/apigw/pkg"
	"github.com/pensando/sw/venice/ctrler/evtsmgr"
	"github.com/pensando/sw/venice/evtsproxy"
	"github.com/pensando/sw/venice/globals"
	scache "github.com/pensando/sw/venice/spyglass/cache"
	"github.com/pensando/sw/venice/spyglass/finder"
	"github.com/pensando/sw/venice/spyglass/indexer"
	"github.com/pensando/sw/venice/utils/audit"
	auditmgr "github.com/pensando/sw/venice/utils/audit/manager"
	authntestutils "github.com/pensando/sw/venice/utils/authn/testutils"
	"github.com/pensando/sw/venice/utils/elastic"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/resolver"

	// for registering services and hooks
	_ "github.com/pensando/sw/api/generated/exports/apigw"
	_ "github.com/pensando/sw/api/generated/exports/apiserver"
	_ "github.com/pensando/sw/api/hooks/apigw"
	_ "github.com/pensando/sw/api/hooks/apiserver"
	_ "github.com/pensando/sw/venice/apigw/svc"
	_ "github.com/pensando/sw/venice/utils/bootstrapper/auth"
)

var (
	// TestLocalUser test username
	TestLocalUser = "test"
	// TestLocalPassword test password
	TestLocalPassword = "pensando"
	// TestTenant test tenant
	TestTenant = globals.DefaultTenant
)

// SetupAuth setsup the authentication service
func SetupAuth(apiServerAddr string, enableLocalAuth bool, ldapConf *auth.Ldap, radiusConf *auth.Radius, creds *auth.PasswordCredential, logger log.Logger) error {
	// create API server client
	apiClient, err := client.NewGrpcUpstream("venice_integ_test_setupAuth", apiServerAddr, logger)
	if err != nil {
		return fmt.Errorf("failed to create gRPC client, err: %v", err)
	}
	defer apiClient.Close()

	// create cluster
	authntestutils.MustCreateCluster(apiClient)
	// create tenant
	authntestutils.MustCreateTenant(apiClient, creds.GetTenant())
	// create local user
	if enableLocalAuth && creds != nil {
		authntestutils.MustCreateTestUser(apiClient, creds.GetUsername(), creds.GetPassword(), creds.GetTenant())
	}
	// create admin role binding
	authntestutils.MustCreateRoleBinding(apiClient, "AdminRoleBinding", creds.GetTenant(), globals.AdminRole, []string{creds.GetUsername()}, nil)
	// create authentication policy
	authntestutils.MustCreateAuthenticationPolicy(apiClient, &auth.Local{Enabled: enableLocalAuth}, ldapConf, radiusConf)
	// set auth bootstrap flag to true
	authntestutils.MustSetAuthBootstrapFlag(apiClient)
	return nil
}

// CleanupAuth removes user, auth policy and rbac objects
func CleanupAuth(apiServerAddr string, enableLocalAuth, enableLdapAuth bool, creds *auth.PasswordCredential, logger log.Logger) error {
	// create API server client
	apiClient, err := client.NewGrpcUpstream("venice_integ_test", apiServerAddr, logger)
	if err != nil {
		return fmt.Errorf("failed to create gRPC client, err: %v", err)
	}
	defer apiClient.Close()

	// delete tenant
	authntestutils.MustDeleteTenant(apiClient, creds.GetTenant())
	// delete local user
	if enableLocalAuth && creds != nil {
		authntestutils.MustDeleteUser(apiClient, creds.GetUsername(), creds.GetTenant())
	}
	// delete admin role binding
	authntestutils.MustDeleteRoleBinding(apiClient, "AdminRoleBinding", creds.GetTenant())
	// delete authentication policy
	authntestutils.MustDeleteAuthenticationPolicy(apiClient)
	// delete cluster
	authntestutils.MustDeleteCluster(apiClient)
	return nil
}

// GetAuthorizationHeader helper function to login and get the authZ header from login context
func GetAuthorizationHeader(apiGwAddr string, creds *auth.PasswordCredential) (string, error) {
	ctx, err := authntestutils.NewLoggedInContext(context.Background(), fmt.Sprintf("https://%s", apiGwAddr), creds)
	if err != nil {
		return "", err
	}

	// get authz header
	authzHeader, ok := loginctx.AuthzHeaderFromContext(ctx)
	if !ok {
		return "", fmt.Errorf("failed to get authorization header from context")
	}

	return authzHeader, nil
}

// StartAPIGateway helper function to start API gateway.
func StartAPIGateway(serverAddr string, skipAuth bool, backends map[string]string, skipServices []string, resolvers []string, l log.Logger) (apigw.APIGateway, string, error) {
	return StartAPIGatewayWithAuditor(serverAddr, skipAuth, backends, skipServices, resolvers, l, auditmgr.WithAuditors(auditmgr.NewLogAuditor(context.TODO(), l)))
}

// StartAPIGatewayWithAuditor helper function to start API Gateway with non default auditor
func StartAPIGatewayWithAuditor(serverAddr string, skipAuth bool, backends map[string]string, skipServices []string, resolvers []string, l log.Logger, auditor audit.Auditor) (apigw.APIGateway, string, error) {
	log.Info("starting API gateway ...")

	// Start the API Gateway
	gwConfig := apigw.Config{
		HTTPAddr:        serverAddr,
		DebugMode:       true,
		Logger:          l,
		BackendOverride: backends,
		Resolvers:       resolvers,
		SkipAuth:        skipAuth,
		SkipAuthz:       skipAuth,
		SkipBackends:    skipServices,
		Auditor:         auditor,
	}
	// skip services
	gwConfig.SkipBackends = append(gwConfig.SkipBackends, skipServices...)

	apiGw := apigwpkg.MustGetAPIGateway()
	go apiGw.Run(gwConfig)
	apiGw.WaitRunning()

	apiGwAddr, err := apiGw.GetAddr()
	if err != nil {
		return nil, "", fmt.Errorf("failed to get API gateway addr, err: %v", err)
	}
	port, err := getPortFromAddr(apiGwAddr.String())
	if err != nil {
		return nil, "", err
	}

	localAddr := fmt.Sprintf("localhost:%s", port)
	log.Infof("API gateway running on %v", localAddr)
	return apiGw, localAddr, nil
}

// StartSpyglass helper function to spyglass finder and indexer
func StartSpyglass(service, apiServerAddr string, mr resolver.Interface, cache scache.Interface, logger log.Logger, esClient elastic.ESClient) (interface{}, string, error) {
	var err error
	if esClient == nil {
		esClient, err = CreateElasticClient("", mr, logger.WithContext("submodule", "elastic"), nil, nil)
		if err != nil {
			return nil, "", fmt.Errorf("failed to create Elastic client for %s, err: %v", service, err)
		}
	}

	switch service {
	case "finder": // create finder
		log.Info("starting finder ...")
		ctx := context.Background()
		fdr, err := finder.NewFinder(ctx, "localhost:0", mr, cache, logger, finder.WithElasticClient(esClient))
		if err != nil {
			return nil, "", fmt.Errorf("failed to create finder, err: %v", err)
		}
		err = fdr.Start() // start the finder
		if err != nil {
			return nil, "", fmt.Errorf("failed to start finder, err: %v", err)
		}

		finderAddr := fdr.GetListenURL()
		if err != nil {
			return nil, "", fmt.Errorf("failed to get API server addr, err: %v", err)
		}
		port, err := getPortFromAddr(finderAddr)
		if err != nil {
			return nil, "", err
		}

		localAddr := fmt.Sprintf("localhost:%s", port)
		log.Infof("Spyglass finder running on %v", localAddr)
		return fdr, localAddr, nil
	case "indexer": // create the indexer
		log.Info("starting indexer ...")
		ctx := context.Background()
		idr, err := indexer.NewIndexer(ctx, apiServerAddr, mr, cache, logger, indexer.WithElasticClient(esClient))
		if err != nil {
			return nil, "", fmt.Errorf("failed to create indexer, err: %v", err)
		}
		err = idr.Start() // start the indexer
		if err != nil {
			return nil, "", fmt.Errorf("failed to start indexer, err: %v", err)
		}

		return idr, "", nil
	}

	return nil, "", nil
}

// StartEvtsMgr helper function to start events manager
func StartEvtsMgr(serverAddr string, mr resolver.Interface, logger log.Logger, esClient elastic.ESClient) (*evtsmgr.EventsManager, string, error) {
	log.Infof("starting events manager")

	var err error
	if esClient == nil {
		esClient, err = elastic.NewClient("", mr, logger.WithContext("submodule", "elastic"))
		if err != nil {
			return nil, "", fmt.Errorf("failed to create Elastic client for events manager, err: %v", err)
		}
	}
	evtsMgr, err := evtsmgr.NewEventsManager(globals.EvtsMgr, serverAddr, mr, logger, evtsmgr.WithElasticClient(esClient))
	if err != nil {
		return nil, "", fmt.Errorf("failed start events manager, err: %v", err)
	}

	return evtsMgr, evtsMgr.RPCServer.GetListenURL(), nil
}

// StartEvtsProxy helper function to start events proxy
func StartEvtsProxy(serverAddr string, mr resolver.Interface, logger log.Logger, dedupInterval,
	batchInterval time.Duration) (*evtsproxy.EventsProxy, string, string, error) {
	log.Infof("starting events proxy")

	if len(mr.GetURLs(globals.EvtsMgr)) == 0 {
		return nil, "", "", fmt.Errorf("could not find evtsmgr URL")
	}

	proxyEventsStoreDir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Errorf("failed to create temp events dir, err: %v", err)
		return nil, "", "", err
	}

	evtsProxy, err := evtsproxy.NewEventsProxy(globals.EvtsProxy, serverAddr,
		nil, dedupInterval, batchInterval, proxyEventsStoreDir, logger)
	if err != nil {
		return nil, "", "", fmt.Errorf("failed start events proxy, err: %v", err)
	}
	if err := evtsProxy.RegisterEventsWriter(evtsproxy.Venice, mr.GetURLs(globals.EvtsMgr)[0]); err != nil {
		return nil, "", "", fmt.Errorf("failed to register venice writer with events proxy")
	}
	evtsProxy.StartDispatch()

	return evtsProxy, evtsProxy.RPCServer.GetListenURL(), proxyEventsStoreDir, nil
}

// helper function to parse the port from given address <ip:port>
func getPortFromAddr(addr string) (string, error) {
	_, port, err := net.SplitHostPort(addr)
	if err != nil {
		return "", fmt.Errorf("failed to parse API server addr, err: %v", err)
	}

	return port, nil
}
