// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package workloadGwService is a auto generated package.
Input file: svc_workload.proto
*/
package workloadGwService

import (
	"context"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pensando/grpc-gateway/runtime"

	"github.com/pensando/sw/api"
	workload "github.com/pensando/sw/api/generated/workload"
	grpcclient "github.com/pensando/sw/api/generated/workload/grpc/client"
	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/api/utils"
	"github.com/pensando/sw/venice/apigw"
	"github.com/pensando/sw/venice/apigw/pkg"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/authz"
	"github.com/pensando/sw/venice/utils/balancer"
	hdr "github.com/pensando/sw/venice/utils/histogram"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/resolver"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// Dummy vars to suppress import errors
var _ api.TypeMeta
var _ authz.Authorizer

type sWorkloadV1GwService struct {
	logger     log.Logger
	defSvcProf apigw.ServiceProfile
	svcProf    map[string]apigw.ServiceProfile
}

type adapterWorkloadV1 struct {
	conn    *rpckit.RPCClient
	service workload.ServiceWorkloadV1Client
	gwSvc   *sWorkloadV1GwService
	gw      apigw.APIGateway
}

func (a adapterWorkloadV1) AbortMigration(oldctx oldcontext.Context, t *workload.Workload, options ...grpc.CallOption) (*workload.Workload, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AbortMigration", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AbortMigration")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.CreateOper, "Workload", t.Tenant, t.Namespace, "workload", t.Name, "AbortMigration"

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Workload)
		return a.service.AbortMigration(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Workload), err
}

func (a adapterWorkloadV1) AutoAddEndpoint(oldctx oldcontext.Context, t *workload.Endpoint, options ...grpc.CallOption) (*workload.Endpoint, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoAddEndpoint", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.CreateOper, "Endpoint", t.Tenant, t.Namespace, "workload", t.Name, strings.Title(string(apiintf.CreateOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Endpoint)
		return a.service.AutoAddEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Endpoint), err
}

func (a adapterWorkloadV1) AutoAddWorkload(oldctx oldcontext.Context, t *workload.Workload, options ...grpc.CallOption) (*workload.Workload, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoAddWorkload", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddWorkload")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.CreateOper, "Workload", t.Tenant, t.Namespace, "workload", t.Name, strings.Title(string(apiintf.CreateOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Workload)
		return a.service.AutoAddWorkload(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Workload), err
}

func (a adapterWorkloadV1) AutoDeleteEndpoint(oldctx oldcontext.Context, t *workload.Endpoint, options ...grpc.CallOption) (*workload.Endpoint, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoDeleteEndpoint", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.DeleteOper, "Endpoint", t.Tenant, t.Namespace, "workload", t.Name, strings.Title(string(apiintf.DeleteOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Endpoint)
		return a.service.AutoDeleteEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Endpoint), err
}

func (a adapterWorkloadV1) AutoDeleteWorkload(oldctx oldcontext.Context, t *workload.Workload, options ...grpc.CallOption) (*workload.Workload, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoDeleteWorkload", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteWorkload")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.DeleteOper, "Workload", t.Tenant, t.Namespace, "workload", t.Name, strings.Title(string(apiintf.DeleteOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Workload)
		return a.service.AutoDeleteWorkload(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Workload), err
}

func (a adapterWorkloadV1) AutoGetEndpoint(oldctx oldcontext.Context, t *workload.Endpoint, options ...grpc.CallOption) (*workload.Endpoint, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoGetEndpoint", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.GetOper, "Endpoint", t.Tenant, t.Namespace, "workload", t.Name, strings.Title(string(apiintf.GetOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Endpoint)
		return a.service.AutoGetEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Endpoint), err
}

func (a adapterWorkloadV1) AutoGetWorkload(oldctx oldcontext.Context, t *workload.Workload, options ...grpc.CallOption) (*workload.Workload, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoGetWorkload", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetWorkload")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.GetOper, "Workload", t.Tenant, t.Namespace, "workload", t.Name, strings.Title(string(apiintf.GetOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Workload)
		return a.service.AutoGetWorkload(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Workload), err
}

func (a adapterWorkloadV1) AutoListEndpoint(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*workload.EndpointList, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoListEndpoint", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	if t.Tenant == "" {
		t.Tenant = globals.DefaultTenant
	}
	t.Namespace = ""
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.ListOper, "Endpoint", t.Tenant, t.Namespace, "workload", "", strings.Title(string(apiintf.ListOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.EndpointList), err
}

func (a adapterWorkloadV1) AutoListWorkload(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*workload.WorkloadList, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoListWorkload", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListWorkload")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	if t.Tenant == "" {
		t.Tenant = globals.DefaultTenant
	}
	t.Namespace = ""
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.ListOper, "Workload", t.Tenant, t.Namespace, "workload", "", strings.Title(string(apiintf.ListOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListWorkload(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.WorkloadList), err
}

func (a adapterWorkloadV1) AutoUpdateEndpoint(oldctx oldcontext.Context, t *workload.Endpoint, options ...grpc.CallOption) (*workload.Endpoint, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoUpdateEndpoint", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.UpdateOper, "Endpoint", t.Tenant, t.Namespace, "workload", t.Name, strings.Title(string(apiintf.UpdateOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Endpoint)
		return a.service.AutoUpdateEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Endpoint), err
}

func (a adapterWorkloadV1) AutoUpdateWorkload(oldctx oldcontext.Context, t *workload.Workload, options ...grpc.CallOption) (*workload.Workload, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1AutoUpdateWorkload", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateWorkload")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.UpdateOper, "Workload", t.Tenant, t.Namespace, "workload", t.Name, strings.Title(string(apiintf.UpdateOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Workload)
		return a.service.AutoUpdateWorkload(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Workload), err
}

func (a adapterWorkloadV1) FinishMigration(oldctx oldcontext.Context, t *workload.Workload, options ...grpc.CallOption) (*workload.Workload, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1FinishMigration", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("FinishMigration")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.CreateOper, "Workload", t.Tenant, t.Namespace, "workload", t.Name, "FinishMigration"

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Workload)
		return a.service.FinishMigration(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Workload), err
}

func (a adapterWorkloadV1) StartMigration(oldctx oldcontext.Context, t *workload.Workload, options ...grpc.CallOption) (*workload.Workload, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.WorkloadV1StartMigration", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("StartMigration")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.CreateOper, "Workload", t.Tenant, t.Namespace, "workload", t.Name, "StartMigration"

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*workload.Workload)
		return a.service.StartMigration(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*workload.Workload), err
}

func (a adapterWorkloadV1) AutoWatchSvcWorkloadV1(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (workload.WorkloadV1_AutoWatchSvcWorkloadV1Client, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchSvcWorkloadV1")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "", in.Tenant, in.Namespace, "workload"
	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, ""), oper, strings.Title(string(oper)))
	ctx = apigwpkg.NewContextWithOperations(ctx, op)
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		iws, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwWebSocketWatch)
		if ok && iws.(bool) {
			nctx, cancel := context.WithCancel(ctx)
			ir, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwHTTPReq)
			if !ok {
				return nil, errors.New("unable to retrieve request")
			}
			iw, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwHTTPWriter)
			if !ok {
				return nil, errors.New("unable to retrieve writer")
			}
			conn, err := wsUpgrader.Upgrade(iw.(http.ResponseWriter), ir.(*http.Request), nil)
			if err != nil {
				log.Errorf("WebSocket Upgrade failed (%s)", err)
				return nil, err
			}
			ctx = apiutils.SetVar(nctx, apiutils.CtxKeyAPIGwWebSocketConn, conn)
			conn.SetCloseHandler(func(code int, text string) error {
				cancel()
				log.Infof("received close notification on websocket [AutoWatchWorkloadV1] (%v/%v)", code, text)
				return nil
			})
			// start a dummy reciever
			go func() {
				for {
					_, _, err := conn.ReadMessage()
					if err != nil {
						log.Errorf("received error on websocket receive (%s)", err)
						cancel()
						return
					}
				}
			}()
		}
		return a.service.AutoWatchSvcWorkloadV1(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(workload.WorkloadV1_AutoWatchSvcWorkloadV1Client), err
}

func (a adapterWorkloadV1) AutoWatchEndpoint(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (workload.WorkloadV1_AutoWatchEndpointClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchEndpoint")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	if in.Tenant == "" {
		in.Tenant = globals.DefaultTenant
	}
	in.Namespace = ""
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "Endpoint", in.Tenant, in.Namespace, "workload"
	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, ""), oper, strings.Title(string(oper)))
	ctx = apigwpkg.NewContextWithOperations(ctx, op)
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		iws, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwWebSocketWatch)
		if ok && iws.(bool) {
			nctx, cancel := context.WithCancel(ctx)
			ir, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwHTTPReq)
			if !ok {
				return nil, errors.New("unable to retrieve request")
			}
			iw, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwHTTPWriter)
			if !ok {
				return nil, errors.New("unable to retrieve writer")
			}
			conn, err := wsUpgrader.Upgrade(iw.(http.ResponseWriter), ir.(*http.Request), nil)
			if err != nil {
				log.Errorf("WebSocket Upgrade failed (%s)", err)
				return nil, err
			}
			ctx = apiutils.SetVar(nctx, apiutils.CtxKeyAPIGwWebSocketConn, conn)
			conn.SetCloseHandler(func(code int, text string) error {
				cancel()
				log.Infof("received close notification on websocket [AutoWatchEndpoint] (%v/%v)", code, text)
				return nil
			})
			// start a dummy reciever
			go func() {
				for {
					_, _, err := conn.ReadMessage()
					if err != nil {
						log.Errorf("received error on websocket receive (%s)", err)
						cancel()
						return
					}
				}
			}()
		}
		return a.service.AutoWatchEndpoint(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(workload.WorkloadV1_AutoWatchEndpointClient), err
}

func (a adapterWorkloadV1) AutoWatchWorkload(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (workload.WorkloadV1_AutoWatchWorkloadClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchWorkload")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	if in.Tenant == "" {
		in.Tenant = globals.DefaultTenant
	}
	in.Namespace = ""
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "Workload", in.Tenant, in.Namespace, "workload"
	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, ""), oper, strings.Title(string(oper)))
	ctx = apigwpkg.NewContextWithOperations(ctx, op)
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		iws, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwWebSocketWatch)
		if ok && iws.(bool) {
			nctx, cancel := context.WithCancel(ctx)
			ir, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwHTTPReq)
			if !ok {
				return nil, errors.New("unable to retrieve request")
			}
			iw, ok := apiutils.GetVar(ctx, apiutils.CtxKeyAPIGwHTTPWriter)
			if !ok {
				return nil, errors.New("unable to retrieve writer")
			}
			conn, err := wsUpgrader.Upgrade(iw.(http.ResponseWriter), ir.(*http.Request), nil)
			if err != nil {
				log.Errorf("WebSocket Upgrade failed (%s)", err)
				return nil, err
			}
			ctx = apiutils.SetVar(nctx, apiutils.CtxKeyAPIGwWebSocketConn, conn)
			conn.SetCloseHandler(func(code int, text string) error {
				cancel()
				log.Infof("received close notification on websocket [AutoWatchWorkload] (%v/%v)", code, text)
				return nil
			})
			// start a dummy reciever
			go func() {
				for {
					_, _, err := conn.ReadMessage()
					if err != nil {
						log.Errorf("received error on websocket receive (%s)", err)
						cancel()
						return
					}
				}
			}()
		}
		return a.service.AutoWatchWorkload(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(workload.WorkloadV1_AutoWatchWorkloadClient), err
}

func (e *sWorkloadV1GwService) setupSvcProfile() {
	e.defSvcProf = apigwpkg.NewServiceProfile(nil, "", "workload", apiintf.UnknownOper)
	e.defSvcProf.SetDefaults()
	e.svcProf = make(map[string]apigw.ServiceProfile)

	e.svcProf["AbortMigration"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Workload", "workload", apiintf.CreateOper)

	e.svcProf["AutoAddWorkload"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Workload", "workload", apiintf.CreateOper)

	e.svcProf["AutoDeleteWorkload"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Workload", "workload", apiintf.DeleteOper)

	e.svcProf["AutoGetEndpoint"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Endpoint", "workload", apiintf.GetOper)

	e.svcProf["AutoGetWorkload"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Workload", "workload", apiintf.GetOper)

	e.svcProf["AutoListEndpoint"] = apigwpkg.NewServiceProfile(e.defSvcProf, "EndpointList", "workload", apiintf.ListOper)

	e.svcProf["AutoListWorkload"] = apigwpkg.NewServiceProfile(e.defSvcProf, "WorkloadList", "workload", apiintf.ListOper)

	e.svcProf["AutoUpdateWorkload"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Workload", "workload", apiintf.UpdateOper)

	e.svcProf["AutoWatchEndpoint"] = apigwpkg.NewServiceProfile(e.defSvcProf, "AutoMsgEndpointWatchHelper", "workload", apiintf.WatchOper)

	e.svcProf["AutoWatchWorkload"] = apigwpkg.NewServiceProfile(e.defSvcProf, "AutoMsgWorkloadWatchHelper", "workload", apiintf.WatchOper)

	e.svcProf["FinishMigration"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Workload", "workload", apiintf.CreateOper)

	e.svcProf["StartMigration"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Workload", "workload", apiintf.CreateOper)
}

// GetDefaultServiceProfile returns the default fallback service profile for this service
func (e *sWorkloadV1GwService) GetDefaultServiceProfile() (apigw.ServiceProfile, error) {
	if e.defSvcProf == nil {
		return nil, errors.New("not found")
	}
	return e.defSvcProf, nil
}

// GetServiceProfile returns the service profile for a given method in this service
func (e *sWorkloadV1GwService) GetServiceProfile(method string) (apigw.ServiceProfile, error) {
	if ret, ok := e.svcProf[method]; ok {
		return ret, nil
	}
	return nil, errors.New("not found")
}

// GetCrudServiceProfile returns the service profile for a auto generated crud operation
func (e *sWorkloadV1GwService) GetCrudServiceProfile(obj string, oper apiintf.APIOperType) (apigw.ServiceProfile, error) {
	name := apiserver.GetCrudServiceName(obj, oper)
	if name != "" {
		return e.GetServiceProfile(name)
	}
	return nil, errors.New("not found")
}

// GetProxyServiceProfile returns the service Profile for a reverse proxy path
func (e *sWorkloadV1GwService) GetProxyServiceProfile(path string) (apigw.ServiceProfile, error) {
	name := "_RProxy_" + path
	return e.GetServiceProfile(name)
}

func (e *sWorkloadV1GwService) CompleteRegistration(ctx context.Context,
	logger log.Logger,
	grpcserver *grpc.Server,
	m *http.ServeMux,
	rslvr resolver.Interface,
	wg *sync.WaitGroup) error {
	apigw := apigwpkg.MustGetAPIGateway()
	// IP:port destination or service discovery key.
	grpcaddr := "pen-apiserver"
	grpcaddr = apigw.GetAPIServerAddr(grpcaddr)
	e.logger = logger

	marshaller := runtime.JSONBuiltin{}
	opts := runtime.WithMarshalerOption("*", &marshaller)
	muxMutex.Lock()
	if mux == nil {
		mux = runtime.NewServeMux(opts)
	}
	muxMutex.Unlock()
	e.setupSvcProfile()

	err := registerSwaggerDef(m, logger)
	if err != nil {
		logger.ErrorLog("msg", "failed to register swagger spec", "service", "workload.WorkloadV1", "err", err)
	}
	wg.Add(1)
	go func() {
		defer func() {
			muxMutex.Lock()
			mux = nil
			muxMutex.Unlock()
		}()
		defer wg.Done()
		for {
			nctx, cancel := context.WithCancel(ctx)
			cl, err := e.newClient(nctx, grpcaddr, rslvr, apigw.GetDevMode())
			if err == nil {
				muxMutex.Lock()
				err = workload.RegisterWorkloadV1HandlerWithClient(ctx, mux, cl)
				muxMutex.Unlock()
				if err == nil {
					logger.InfoLog("msg", "registered service workload.WorkloadV1")
					m.Handle("/configs/workload/v1/", http.StripPrefix("/configs/workload/v1", mux))
					return
				} else {
					err = errors.Wrap(err, "failed to register")
				}
			} else {
				err = errors.Wrap(err, "failed to create client")
			}
			cancel()
			logger.ErrorLog("msg", "failed to register", "service", "workload.WorkloadV1", "err", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()
	return nil
}

func (e *sWorkloadV1GwService) newClient(ctx context.Context, grpcAddr string, rslvr resolver.Interface, devmode bool) (*adapterWorkloadV1, error) {
	var opts []rpckit.Option
	opts = append(opts, rpckit.WithTLSClientIdentity(globals.APIGw))
	if rslvr != nil {
		opts = append(opts, rpckit.WithBalancer(balancer.New(rslvr)))
	} else {
		opts = append(opts, rpckit.WithRemoteServerName("pen-apiserver"))
	}

	if !devmode {
		opts = append(opts, rpckit.WithTracerEnabled(false))
		opts = append(opts, rpckit.WithLoggerEnabled(false))
		opts = append(opts, rpckit.WithStatsEnabled(false))
	}

	client, err := rpckit.NewRPCClient(globals.APIGw, grpcAddr, opts...)
	if err != nil {
		return nil, errors.Wrap(err, "create rpc client")
	}

	e.logger.Infof("Connected to GRPC Server %s", grpcAddr)
	defer func() {
		go func() {
			<-ctx.Done()
			if cerr := client.Close(); cerr != nil {
				e.logger.ErrorLog("msg", "Failed to close conn on Done()", "addr", grpcAddr, "err", cerr)
			}
		}()
	}()

	cl := &adapterWorkloadV1{conn: client, gw: apigwpkg.MustGetAPIGateway(), gwSvc: e, service: grpcclient.NewWorkloadV1Backend(client.ClientConn, e.logger)}
	return cl, nil
}

func init() {

	apigw := apigwpkg.MustGetAPIGateway()

	svcWorkloadV1 := sWorkloadV1GwService{}
	apigw.Register("workload.WorkloadV1", "workload/", &svcWorkloadV1)
}
