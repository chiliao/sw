// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package diagnosticsGwService is a auto generated package.
Input file: svc_diagnostics.proto
*/
package diagnosticsGwService

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
	diagnostics "github.com/pensando/sw/api/generated/diagnostics"
	grpcclient "github.com/pensando/sw/api/generated/diagnostics/grpc/client"
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

type sDiagnosticsV1GwService struct {
	logger     log.Logger
	defSvcProf apigw.ServiceProfile
	svcProf    map[string]apigw.ServiceProfile
}

type adapterDiagnosticsV1 struct {
	conn    *rpckit.RPCClient
	service diagnostics.ServiceDiagnosticsV1Client
	gwSvc   *sDiagnosticsV1GwService
	gw      apigw.APIGateway
}

func (a adapterDiagnosticsV1) AutoAddModule(oldctx oldcontext.Context, t *diagnostics.Module, options ...grpc.CallOption) (*diagnostics.Module, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.DiagnosticsV1AutoAddModule", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddModule")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.CreateOper, "Module", t.Tenant, t.Namespace, "diagnostics", t.Name, strings.Title(string(apiintf.CreateOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*diagnostics.Module)
		return a.service.AutoAddModule(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*diagnostics.Module), err
}

func (a adapterDiagnosticsV1) AutoDeleteModule(oldctx oldcontext.Context, t *diagnostics.Module, options ...grpc.CallOption) (*diagnostics.Module, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.DiagnosticsV1AutoDeleteModule", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteModule")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.DeleteOper, "Module", t.Tenant, t.Namespace, "diagnostics", t.Name, strings.Title(string(apiintf.DeleteOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*diagnostics.Module)
		return a.service.AutoDeleteModule(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*diagnostics.Module), err
}

func (a adapterDiagnosticsV1) AutoGetModule(oldctx oldcontext.Context, t *diagnostics.Module, options ...grpc.CallOption) (*diagnostics.Module, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.DiagnosticsV1AutoGetModule", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetModule")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.GetOper, "Module", t.Tenant, t.Namespace, "diagnostics", t.Name, strings.Title(string(apiintf.GetOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*diagnostics.Module)
		return a.service.AutoGetModule(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*diagnostics.Module), err
}

func (a adapterDiagnosticsV1) AutoLabelModule(oldctx oldcontext.Context, t *api.Label, options ...grpc.CallOption) (*diagnostics.Module, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.DiagnosticsV1AutoLabelModule", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoLabelModule")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.UpdateOper, "Module", t.Tenant, t.Namespace, "diagnostics", t.Name, strings.Title(string(apiintf.LabelOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.Label)
		return a.service.AutoLabelModule(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*diagnostics.Module), err
}

func (a adapterDiagnosticsV1) AutoListModule(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*diagnostics.ModuleList, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.DiagnosticsV1AutoListModule", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListModule")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	t.Tenant = ""
	t.Namespace = ""
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.ListOper, "Module", t.Tenant, t.Namespace, "diagnostics", "", strings.Title(string(apiintf.ListOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListModule(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*diagnostics.ModuleList), err
}

func (a adapterDiagnosticsV1) AutoUpdateModule(oldctx oldcontext.Context, t *diagnostics.Module, options ...grpc.CallOption) (*diagnostics.Module, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.DiagnosticsV1AutoUpdateModule", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateModule")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.UpdateOper, "Module", t.Tenant, t.Namespace, "diagnostics", t.Name, strings.Title(string(apiintf.UpdateOper))

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*diagnostics.Module)
		return a.service.AutoUpdateModule(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*diagnostics.Module), err
}

func (a adapterDiagnosticsV1) Debug(oldctx oldcontext.Context, t *diagnostics.DiagnosticsRequest, options ...grpc.CallOption) (*diagnostics.DiagnosticsResponse, error) {
	// Not using options for now. Will be passed through context as needed.
	trackTime := time.Now()
	defer func() {
		hdr.Record("apigw.DiagnosticsV1Debug", time.Since(trackTime))
	}()
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("Debug")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name, auditAction := apiintf.CreateOper, "Module", t.Tenant, t.Namespace, "diagnostics", t.Name, "Debug"

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper, auditAction)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*diagnostics.DiagnosticsRequest)
		return a.service.Debug(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*diagnostics.DiagnosticsResponse), err
}

func (a adapterDiagnosticsV1) AutoWatchSvcDiagnosticsV1(oldctx oldcontext.Context, in *api.AggWatchOptions, options ...grpc.CallOption) (diagnostics.DiagnosticsV1_AutoWatchSvcDiagnosticsV1Client, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchSvcDiagnosticsV1")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "", in.Tenant, in.Namespace, "diagnostics"
	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, ""), oper, strings.Title(string(oper)))
	ctx = apigwpkg.NewContextWithOperations(ctx, op)
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.AggWatchOptions)
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
				log.Infof("received close notification on websocket [AutoWatchDiagnosticsV1] (%v/%v)", code, text)
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
		return a.service.AutoWatchSvcDiagnosticsV1(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(diagnostics.DiagnosticsV1_AutoWatchSvcDiagnosticsV1Client), err
}

func (a adapterDiagnosticsV1) AutoWatchModule(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (diagnostics.DiagnosticsV1_AutoWatchModuleClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchModule")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	in.Tenant = ""
	in.Namespace = ""
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "Module", in.Tenant, in.Namespace, "diagnostics"
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
				log.Infof("received close notification on websocket [AutoWatchModule] (%v/%v)", code, text)
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
		return a.service.AutoWatchModule(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(diagnostics.DiagnosticsV1_AutoWatchModuleClient), err
}

func (e *sDiagnosticsV1GwService) setupSvcProfile() {
	e.defSvcProf = apigwpkg.NewServiceProfile(nil, "", "diagnostics", apiintf.UnknownOper)
	e.defSvcProf.SetDefaults()
	e.svcProf = make(map[string]apigw.ServiceProfile)

	e.svcProf["AutoGetModule"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Module", "diagnostics", apiintf.GetOper)

	e.svcProf["AutoLabelModule"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Module", "diagnostics", apiintf.CreateOper)

	e.svcProf["AutoListModule"] = apigwpkg.NewServiceProfile(e.defSvcProf, "ModuleList", "diagnostics", apiintf.ListOper)

	e.svcProf["AutoUpdateModule"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Module", "diagnostics", apiintf.UpdateOper)

	e.svcProf["AutoWatchModule"] = apigwpkg.NewServiceProfile(e.defSvcProf, "AutoMsgModuleWatchHelper", "diagnostics", apiintf.WatchOper)

	e.svcProf["Debug"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Module", "diagnostics", apiintf.CreateOper)
}

// GetDefaultServiceProfile returns the default fallback service profile for this service
func (e *sDiagnosticsV1GwService) GetDefaultServiceProfile() (apigw.ServiceProfile, error) {
	if e.defSvcProf == nil {
		return nil, errors.New("not found")
	}
	return e.defSvcProf, nil
}

// GetServiceProfile returns the service profile for a given method in this service
func (e *sDiagnosticsV1GwService) GetServiceProfile(method string) (apigw.ServiceProfile, error) {
	if ret, ok := e.svcProf[method]; ok {
		return ret, nil
	}
	return nil, errors.New("not found")
}

// GetCrudServiceProfile returns the service profile for a auto generated crud operation
func (e *sDiagnosticsV1GwService) GetCrudServiceProfile(obj string, oper apiintf.APIOperType) (apigw.ServiceProfile, error) {
	name := apiserver.GetCrudServiceName(obj, oper)
	if name != "" {
		return e.GetServiceProfile(name)
	}
	return nil, errors.New("not found")
}

// GetProxyServiceProfile returns the service Profile for a reverse proxy path
func (e *sDiagnosticsV1GwService) GetProxyServiceProfile(path string) (apigw.ServiceProfile, error) {
	name := "_RProxy_" + path
	return e.GetServiceProfile(name)
}

func (e *sDiagnosticsV1GwService) CompleteRegistration(ctx context.Context,
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
		logger.ErrorLog("msg", "failed to register swagger spec", "service", "diagnostics.DiagnosticsV1", "err", err)
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
				err = diagnostics.RegisterDiagnosticsV1HandlerWithClient(ctx, mux, cl)
				muxMutex.Unlock()
				if err == nil {
					logger.InfoLog("msg", "registered service diagnostics.DiagnosticsV1")
					m.Handle("/configs/diagnostics/v1/", http.StripPrefix("/configs/diagnostics/v1", mux))
					return
				} else {
					err = errors.Wrap(err, "failed to register")
				}
			} else {
				err = errors.Wrap(err, "failed to create client")
			}
			cancel()
			logger.ErrorLog("msg", "failed to register", "service", "diagnostics.DiagnosticsV1", "err", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()
	return nil
}

func (e *sDiagnosticsV1GwService) newClient(ctx context.Context, grpcAddr string, rslvr resolver.Interface, devmode bool) (*adapterDiagnosticsV1, error) {
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

	cl := &adapterDiagnosticsV1{conn: client, gw: apigwpkg.MustGetAPIGateway(), gwSvc: e, service: grpcclient.NewDiagnosticsV1Backend(client.ClientConn, e.logger)}
	return cl, nil
}

func init() {

	apigw := apigwpkg.MustGetAPIGateway()

	svcDiagnosticsV1 := sDiagnosticsV1GwService{}
	apigw.Register("diagnostics.DiagnosticsV1", "diagnostics/", &svcDiagnosticsV1)
}
