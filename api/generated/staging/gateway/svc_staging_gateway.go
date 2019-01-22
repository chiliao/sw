// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package stagingGwService is a auto generated package.
Input file: svc_staging.proto
*/
package stagingGwService

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/pkg/errors"
	oldcontext "golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/pensando/grpc-gateway/runtime"

	"github.com/pensando/sw/api"
	staging "github.com/pensando/sw/api/generated/staging"
	grpcclient "github.com/pensando/sw/api/generated/staging/grpc/client"
	"github.com/pensando/sw/api/utils"
	"github.com/pensando/sw/venice/apigw"
	"github.com/pensando/sw/venice/apigw/pkg"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/authz"
	"github.com/pensando/sw/venice/utils/balancer"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/resolver"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// Dummy vars to suppress import errors
var _ api.TypeMeta
var _ authz.Authorizer

type sStagingV1GwService struct {
	logger     log.Logger
	defSvcProf apigw.ServiceProfile
	svcProf    map[string]apigw.ServiceProfile
}

type adapterStagingV1 struct {
	conn    *rpckit.RPCClient
	service staging.ServiceStagingV1Client
	gwSvc   *sStagingV1GwService
	gw      apigw.APIGateway
}

func (a adapterStagingV1) AutoAddBuffer(oldctx oldcontext.Context, t *staging.Buffer, options ...grpc.CallOption) (*staging.Buffer, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddBuffer")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiserver.CreateOper, "Buffer", t.Tenant, t.Namespace, "staging", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*staging.Buffer)
		return a.service.AutoAddBuffer(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*staging.Buffer), err
}

func (a adapterStagingV1) AutoDeleteBuffer(oldctx oldcontext.Context, t *staging.Buffer, options ...grpc.CallOption) (*staging.Buffer, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteBuffer")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiserver.DeleteOper, "Buffer", t.Tenant, t.Namespace, "staging", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*staging.Buffer)
		return a.service.AutoDeleteBuffer(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*staging.Buffer), err
}

func (a adapterStagingV1) AutoGetBuffer(oldctx oldcontext.Context, t *staging.Buffer, options ...grpc.CallOption) (*staging.Buffer, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetBuffer")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiserver.GetOper, "Buffer", t.Tenant, t.Namespace, "staging", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*staging.Buffer)
		return a.service.AutoGetBuffer(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*staging.Buffer), err
}

func (a adapterStagingV1) AutoListBuffer(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*staging.BufferList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListBuffer")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	if t.Tenant == "" {
		t.Tenant = globals.DefaultTenant
	}
	t.Namespace = ""
	oper, kind, tenant, namespace, group, name := apiserver.ListOper, "BufferList", t.Tenant, t.Namespace, "staging", ""

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListBuffer(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*staging.BufferList), err
}

func (a adapterStagingV1) AutoUpdateBuffer(oldctx oldcontext.Context, t *staging.Buffer, options ...grpc.CallOption) (*staging.Buffer, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateBuffer")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiserver.UpdateOper, "Buffer", t.Tenant, t.Namespace, "staging", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*staging.Buffer)
		return a.service.AutoUpdateBuffer(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*staging.Buffer), err
}

func (a adapterStagingV1) Clear(oldctx oldcontext.Context, t *staging.ClearAction, options ...grpc.CallOption) (*staging.ClearAction, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("Clear")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiserver.CreateOper, "Buffer", t.Tenant, t.Namespace, "staging", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*staging.ClearAction)
		return a.service.Clear(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*staging.ClearAction), err
}

func (a adapterStagingV1) Commit(oldctx oldcontext.Context, t *staging.CommitAction, options ...grpc.CallOption) (*staging.CommitAction, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("Commit")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiserver.CreateOper, "Buffer", t.Tenant, t.Namespace, "staging", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*staging.CommitAction)
		return a.service.Commit(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*staging.CommitAction), err
}

func (a adapterStagingV1) AutoWatchSvcStagingV1(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (staging.StagingV1_AutoWatchSvcStagingV1Client, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchSvcStagingV1")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group := apiserver.WatchOper, "", in.Tenant, in.Namespace, "staging"
	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, ""), oper)
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
				log.Infof("received close notification on websocket [AutoWatchStagingV1] (%v/%v)", code, text)
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
		return a.service.AutoWatchSvcStagingV1(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(staging.StagingV1_AutoWatchSvcStagingV1Client), err
}

func (a adapterStagingV1) AutoWatchBuffer(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (staging.StagingV1_AutoWatchBufferClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchBuffer")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	in.Namespace = ""
	oper, kind, tenant, namespace, group := apiserver.WatchOper, "Buffer", in.Tenant, in.Namespace, "staging"
	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, ""), oper)
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
				log.Infof("received close notification on websocket [AutoWatchBuffer] (%v/%v)", code, text)
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
		return a.service.AutoWatchBuffer(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(staging.StagingV1_AutoWatchBufferClient), err
}

func (e *sStagingV1GwService) setupSvcProfile() {
	e.defSvcProf = apigwpkg.NewServiceProfile(nil, "", "staging", apiserver.UnknownOper)
	e.defSvcProf.SetDefaults()
	e.svcProf = make(map[string]apigw.ServiceProfile)

	e.svcProf["AutoAddBuffer"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Buffer", "staging", apiserver.CreateOper)

	e.svcProf["AutoDeleteBuffer"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Buffer", "staging", apiserver.DeleteOper)

	e.svcProf["AutoGetBuffer"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Buffer", "staging", apiserver.GetOper)

	e.svcProf["AutoListBuffer"] = apigwpkg.NewServiceProfile(e.defSvcProf, "BufferList", "staging", apiserver.ListOper)

	e.svcProf["Clear"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Buffer", "staging", apiserver.CreateOper)

	e.svcProf["Commit"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Buffer", "staging", apiserver.CreateOper)
}

// GetDefaultServiceProfile returns the default fallback service profile for this service
func (e *sStagingV1GwService) GetDefaultServiceProfile() (apigw.ServiceProfile, error) {
	if e.defSvcProf == nil {
		return nil, errors.New("not found")
	}
	return e.defSvcProf, nil
}

// GetServiceProfile returns the service profile for a given method in this service
func (e *sStagingV1GwService) GetServiceProfile(method string) (apigw.ServiceProfile, error) {
	if ret, ok := e.svcProf[method]; ok {
		return ret, nil
	}
	return nil, errors.New("not found")
}

// GetCrudServiceProfile returns the service profile for a auto generated crud operation
func (e *sStagingV1GwService) GetCrudServiceProfile(obj string, oper apiserver.APIOperType) (apigw.ServiceProfile, error) {
	name := apiserver.GetCrudServiceName(obj, oper)
	if name != "" {
		return e.GetServiceProfile(name)
	}
	return nil, errors.New("not found")
}

// GetProxyServiceProfile returns the service Profile for a reverse proxy path
func (e *sStagingV1GwService) GetProxyServiceProfile(path string) (apigw.ServiceProfile, error) {
	name := "_RProxy_" + path
	return e.GetServiceProfile(name)
}

func (e *sStagingV1GwService) CompleteRegistration(ctx context.Context,
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
		logger.ErrorLog("msg", "failed to register swagger spec", "service", "staging.StagingV1", "error", err)
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
				err = staging.RegisterStagingV1HandlerWithClient(ctx, mux, cl)
				muxMutex.Unlock()
				if err == nil {
					logger.InfoLog("msg", "registered service staging.StagingV1")
					m.Handle("/configs/staging/v1/", http.StripPrefix("/configs/staging/v1", mux))
					return
				} else {
					err = errors.Wrap(err, "failed to register")
				}
			} else {
				err = errors.Wrap(err, "failed to create client")
			}
			cancel()
			logger.ErrorLog("msg", "failed to register", "service", "staging.StagingV1", "error", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()
	return nil
}

func (e *sStagingV1GwService) newClient(ctx context.Context, grpcAddr string, rslvr resolver.Interface, devmode bool) (*adapterStagingV1, error) {
	var opts []rpckit.Option
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
				e.logger.ErrorLog("msg", "Failed to close conn on Done()", "addr", grpcAddr, "error", cerr)
			}
		}()
	}()

	cl := &adapterStagingV1{conn: client, gw: apigwpkg.MustGetAPIGateway(), gwSvc: e, service: grpcclient.NewStagingV1Backend(client.ClientConn, e.logger)}
	return cl, nil
}

func init() {

	apigw := apigwpkg.MustGetAPIGateway()

	svcStagingV1 := sStagingV1GwService{}
	apigw.Register("staging.StagingV1", "staging/", &svcStagingV1)
}
