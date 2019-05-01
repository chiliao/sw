// {C} Copyright 2018 Pensando Systems Inc. All rights reserved.

package veniceinteg

import (
	"fmt"
	"runtime"
	"strings"

	. "gopkg.in/check.v1"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/auth"
	"github.com/pensando/sw/api/generated/telemetry_query"
	"github.com/pensando/sw/api/login"
	"github.com/pensando/sw/nic/delphi/proto/goproto"
	"github.com/pensando/sw/test/utils"
	"github.com/pensando/sw/venice/globals"
	. "github.com/pensando/sw/venice/utils/authn/testutils"
	"github.com/pensando/sw/venice/utils/authz"
	"github.com/pensando/sw/venice/utils/ntranslate"
	"github.com/pensando/sw/venice/utils/telemetryclient"
	. "github.com/pensando/sw/venice/utils/testutils"
)

type keylookup struct {
	version string
}

func (m *keylookup) KeyToMeta(key interface{}) *api.ObjectMeta {
	return &api.ObjectMeta{
		Name:      fmt.Sprintf("%d", key),
		Tenant:    globals.DefaultTenant,
		Namespace: globals.DefaultNamespace,
	}
}

func (m *keylookup) MetaToKey(meta *api.ObjectMeta) interface{} {
	return nil
}

func (it *veniceIntegSuite) TestMetrics(c *C) {
	// metrics iterators don't work in OSX
	if runtime.GOOS == "darwin" {
		return
	}

	tslt := ntranslate.MustGetTranslator()
	Assert(c, tslt != nil, "failed to get translator")
	tslt.Register("LifMetrics", &keylookup{})
	iter, err := goproto.NewLifMetricsIterator()
	AssertOk(c, err, "Error creating metrics iterator")

	// create an entry
	tmtr, err := iter.Create(3000)
	AssertOk(c, err, "Error creating test metrics entry")

	// set some values
	tmtr.SetRxBroadcastBytes(200)
	tmtr.SetRxBroadcastPackets(300)
	tmtr.SetRxCsumComplete(400.0)
	tmtr.SetRxDescFetchError(500.0)

	// query
	apiGwAddr := "localhost:" + it.config.APIGatewayPort
	tc, err := telemetryclient.NewTelemetryClient(apiGwAddr)
	AssertOk(c, err, "Error creating metrics client")

	ctx, err := it.loggedInCtx()
	AssertOk(c, err, "Error in logged in context")

	AssertEventually(c, func() (bool, interface{}) {
		nodeQuery := &telemetry_query.MetricsQueryList{
			Tenant:    globals.DefaultTenant,
			Namespace: globals.DefaultNamespace,
			Queries: []*telemetry_query.MetricsQuerySpec{
				{
					TypeMeta: api.TypeMeta{
						Kind: "LifMetrics",
					},
				},
			},
		}

		res, err := tc.Metrics(ctx, nodeQuery)
		if err != nil {
			return false, err
		}

		if len(res.Results) == 0 || len(res.Results[0].Series) == 0 {
			return false, res
		}

		return true, res

	}, "failed to query metrics")
}

func (it *veniceIntegSuite) TestMetricsAuthz(c *C) {
	// metrics iterators don't work in OSX
	if runtime.GOOS == "darwin" {
		return
	}

	tslt := ntranslate.MustGetTranslator()
	Assert(c, tslt != nil, "failed to get translator")
	tslt.Register("LifMetrics", &keylookup{})
	iter, err := goproto.NewLifMetricsIterator()
	AssertOk(c, err, "Error creating metrics iterator")

	// create an entry
	tmtr, err := iter.Create(3000)
	AssertOk(c, err, "Error creating test metrics entry")

	// set some values
	tmtr.SetRxBroadcastBytes(200)
	tmtr.SetRxBroadcastPackets(300)
	tmtr.SetRxCsumComplete(400.0)
	tmtr.SetRxDescFetchError(500.0)

	apiGwAddr := "localhost:" + it.config.APIGatewayPort
	tc, err := telemetryclient.NewTelemetryClient(apiGwAddr)
	AssertOk(c, err, "Error creating metrics client")
	const (
		testTenant = "testtenant"
		testUser   = "testuser1"
	)
	MustCreateTenant(it.apisrvClient, testTenant)
	defer MustDeleteTenant(it.apisrvClient, testTenant)
	MustCreateTestUser(it.apisrvClient, testUser, utils.TestLocalPassword, testTenant)
	defer MustDeleteUser(it.apisrvClient, testUser, testTenant)
	MustUpdateRoleBinding(it.apisrvClient, globals.AdminRoleBinding, testTenant, globals.AdminRole, []string{testUser}, nil)
	defer MustUpdateRoleBinding(it.apisrvClient, globals.AdminRoleBinding, testTenant, globals.AdminRole, nil, nil)
	ctx, err := it.loggedInCtxWithCred(testTenant, testUser, utils.TestLocalPassword)
	AssertOk(c, err, "error in logging in testtenant")
	nodeQuery := &telemetry_query.MetricsQueryList{
		Tenant:    globals.DefaultTenant,
		Namespace: globals.DefaultNamespace,
		Queries: []*telemetry_query.MetricsQuerySpec{
			{
				TypeMeta: api.TypeMeta{
					Kind: "LifMetrics",
				},
			},
		},
	}
	// query metrics in another tenant
	_, err = tc.Metrics(ctx, nodeQuery)
	Assert(c, err != nil && strings.Contains(err.Error(), "403"), "expected authorization error while querying metrics in other tenant")

	MustCreateTestUser(it.apisrvClient, testUser, utils.TestLocalPassword, globals.DefaultTenant)
	defer MustDeleteUser(it.apisrvClient, testUser, globals.DefaultTenant)
	ctx, err = it.loggedInCtxWithCred(globals.DefaultTenant, testUser, utils.TestLocalPassword)
	AssertOk(c, err, "error in logging in default tenant")
	// query metrics with no perms in own tenant
	_, err = tc.Metrics(ctx, nodeQuery)
	Assert(c, err != nil && strings.Contains(err.Error(), "403"), "expected authorization error while querying metrics")

	MustCreateRole(it.apisrvClient, "MetricsPerms", globals.DefaultTenant,
		login.NewPermission(globals.DefaultTenant, "", auth.Permission_MetricsQuery.String(), authz.ResourceNamespaceAll, "", auth.Permission_Read.String()),
		login.NewPermission(globals.DefaultTenant, "", "LifMetrics", authz.ResourceNamespaceAll, "", auth.Permission_Read.String()),
	)
	defer MustDeleteRole(it.apisrvClient, "MetricsPerms", globals.DefaultTenant)
	MustCreateRoleBinding(it.apisrvClient, "MetricsPermsRB", globals.DefaultTenant, "MetricsPerms", []string{testUser}, nil)
	defer MustDeleteRoleBinding(it.apisrvClient, "MetricsPermsRB", globals.DefaultTenant)

	// query metrics with valid authorization
	AssertEventually(c, func() (bool, interface{}) {
		nodeQuery := &telemetry_query.MetricsQueryList{
			Tenant:    globals.DefaultTenant,
			Namespace: globals.DefaultNamespace,
			Queries: []*telemetry_query.MetricsQuerySpec{
				{
					TypeMeta: api.TypeMeta{
						Kind: "LifMetrics",
					},
				},
			},
		}

		res, err := tc.Metrics(ctx, nodeQuery)
		if err != nil {
			return false, err
		}

		if len(res.Results) == 0 || len(res.Results[0].Series) == 0 {
			return false, res
		}

		return true, res

	}, "failed to query metrics")
}
