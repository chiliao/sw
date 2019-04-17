// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package rolloutGwService is a auto generated package.
Input file: svc_rollout.proto
*/
package rolloutGwService

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
	rollout "github.com/pensando/sw/api/generated/rollout"
	grpcclient "github.com/pensando/sw/api/generated/rollout/grpc/client"
	"github.com/pensando/sw/api/interfaces"
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

type sRolloutV1GwService struct {
	logger     log.Logger
	defSvcProf apigw.ServiceProfile
	svcProf    map[string]apigw.ServiceProfile
}

type adapterRolloutV1 struct {
	conn    *rpckit.RPCClient
	service rollout.ServiceRolloutV1Client
	gwSvc   *sRolloutV1GwService
	gw      apigw.APIGateway
}

func (a adapterRolloutV1) AutoAddRollout(oldctx oldcontext.Context, t *rollout.Rollout, options ...grpc.CallOption) (*rollout.Rollout, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddRollout")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.CreateOper, "Rollout", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.Rollout)
		return a.service.AutoAddRollout(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.Rollout), err
}

func (a adapterRolloutV1) AutoAddRolloutAction(oldctx oldcontext.Context, t *rollout.RolloutAction, options ...grpc.CallOption) (*rollout.RolloutAction, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddRolloutAction")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.CreateOper, "RolloutAction", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.RolloutAction)
		return a.service.AutoAddRolloutAction(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.RolloutAction), err
}

func (a adapterRolloutV1) AutoDeleteRollout(oldctx oldcontext.Context, t *rollout.Rollout, options ...grpc.CallOption) (*rollout.Rollout, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteRollout")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.DeleteOper, "Rollout", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.Rollout)
		return a.service.AutoDeleteRollout(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.Rollout), err
}

func (a adapterRolloutV1) AutoDeleteRolloutAction(oldctx oldcontext.Context, t *rollout.RolloutAction, options ...grpc.CallOption) (*rollout.RolloutAction, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteRolloutAction")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.DeleteOper, "RolloutAction", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.RolloutAction)
		return a.service.AutoDeleteRolloutAction(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.RolloutAction), err
}

func (a adapterRolloutV1) AutoGetRollout(oldctx oldcontext.Context, t *rollout.Rollout, options ...grpc.CallOption) (*rollout.Rollout, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetRollout")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.GetOper, "Rollout", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.Rollout)
		return a.service.AutoGetRollout(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.Rollout), err
}

func (a adapterRolloutV1) AutoGetRolloutAction(oldctx oldcontext.Context, t *rollout.RolloutAction, options ...grpc.CallOption) (*rollout.RolloutAction, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetRolloutAction")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.GetOper, "RolloutAction", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.RolloutAction)
		return a.service.AutoGetRolloutAction(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.RolloutAction), err
}

func (a adapterRolloutV1) AutoListRollout(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*rollout.RolloutList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListRollout")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	t.Tenant = ""
	t.Namespace = ""
	oper, kind, tenant, namespace, group, name := apiintf.ListOper, "Rollout", t.Tenant, t.Namespace, "rollout", ""

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListRollout(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.RolloutList), err
}

func (a adapterRolloutV1) AutoListRolloutAction(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*rollout.RolloutActionList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListRolloutAction")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	t.Tenant = ""
	t.Namespace = ""
	oper, kind, tenant, namespace, group, name := apiintf.ListOper, "RolloutAction", t.Tenant, t.Namespace, "rollout", ""

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListRolloutAction(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.RolloutActionList), err
}

func (a adapterRolloutV1) AutoUpdateRollout(oldctx oldcontext.Context, t *rollout.Rollout, options ...grpc.CallOption) (*rollout.Rollout, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateRollout")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.UpdateOper, "Rollout", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.Rollout)
		return a.service.AutoUpdateRollout(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.Rollout), err
}

func (a adapterRolloutV1) AutoUpdateRolloutAction(oldctx oldcontext.Context, t *rollout.RolloutAction, options ...grpc.CallOption) (*rollout.RolloutAction, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateRolloutAction")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.UpdateOper, "RolloutAction", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.RolloutAction)
		return a.service.AutoUpdateRolloutAction(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.RolloutAction), err
}

func (a adapterRolloutV1) DoRollout(oldctx oldcontext.Context, t *rollout.Rollout, options ...grpc.CallOption) (*rollout.Rollout, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("DoRollout")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.CreateOper, "Rollout", t.Tenant, t.Namespace, "rollout", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*rollout.Rollout)
		return a.service.DoRollout(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*rollout.Rollout), err
}

func (a adapterRolloutV1) AutoWatchSvcRolloutV1(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (rollout.RolloutV1_AutoWatchSvcRolloutV1Client, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchSvcRolloutV1")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "", in.Tenant, in.Namespace, "rollout"
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
				log.Infof("received close notification on websocket [AutoWatchRolloutV1] (%v/%v)", code, text)
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
		return a.service.AutoWatchSvcRolloutV1(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(rollout.RolloutV1_AutoWatchSvcRolloutV1Client), err
}

func (a adapterRolloutV1) AutoWatchRollout(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (rollout.RolloutV1_AutoWatchRolloutClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchRollout")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	in.Tenant = ""
	in.Namespace = ""
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "Rollout", in.Tenant, in.Namespace, "rollout"
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
				log.Infof("received close notification on websocket [AutoWatchRollout] (%v/%v)", code, text)
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
		return a.service.AutoWatchRollout(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(rollout.RolloutV1_AutoWatchRolloutClient), err
}

func (a adapterRolloutV1) AutoWatchRolloutAction(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (rollout.RolloutV1_AutoWatchRolloutActionClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchRolloutAction")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	in.Tenant = ""
	in.Namespace = ""
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "RolloutAction", in.Tenant, in.Namespace, "rollout"
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
				log.Infof("received close notification on websocket [AutoWatchRolloutAction] (%v/%v)", code, text)
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
		return a.service.AutoWatchRolloutAction(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(rollout.RolloutV1_AutoWatchRolloutActionClient), err
}

func (e *sRolloutV1GwService) setupSvcProfile() {
	e.defSvcProf = apigwpkg.NewServiceProfile(nil, "", "rollout", apiintf.UnknownOper)
	e.defSvcProf.SetDefaults()
	e.svcProf = make(map[string]apigw.ServiceProfile)

	e.svcProf["AutoDeleteRollout"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Rollout", "rollout", apiintf.DeleteOper)

	e.svcProf["AutoDeleteRolloutAction"] = apigwpkg.NewServiceProfile(e.defSvcProf, "RolloutAction", "rollout", apiintf.DeleteOper)

	e.svcProf["AutoGetRollout"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Rollout", "rollout", apiintf.GetOper)

	e.svcProf["AutoGetRolloutAction"] = apigwpkg.NewServiceProfile(e.defSvcProf, "RolloutAction", "rollout", apiintf.GetOper)

	e.svcProf["AutoListRollout"] = apigwpkg.NewServiceProfile(e.defSvcProf, "RolloutList", "rollout", apiintf.ListOper)

	e.svcProf["AutoWatchRollout"] = apigwpkg.NewServiceProfile(e.defSvcProf, "AutoMsgRolloutWatchHelper", "rollout", apiintf.WatchOper)

	e.svcProf["AutoWatchRolloutAction"] = apigwpkg.NewServiceProfile(e.defSvcProf, "AutoMsgRolloutActionWatchHelper", "rollout", apiintf.WatchOper)

	e.svcProf["DoRollout"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Rollout", "rollout", apiintf.CreateOper)
}

// GetDefaultServiceProfile returns the default fallback service profile for this service
func (e *sRolloutV1GwService) GetDefaultServiceProfile() (apigw.ServiceProfile, error) {
	if e.defSvcProf == nil {
		return nil, errors.New("not found")
	}
	return e.defSvcProf, nil
}

// GetServiceProfile returns the service profile for a given method in this service
func (e *sRolloutV1GwService) GetServiceProfile(method string) (apigw.ServiceProfile, error) {
	if ret, ok := e.svcProf[method]; ok {
		return ret, nil
	}
	return nil, errors.New("not found")
}

// GetCrudServiceProfile returns the service profile for a auto generated crud operation
func (e *sRolloutV1GwService) GetCrudServiceProfile(obj string, oper apiintf.APIOperType) (apigw.ServiceProfile, error) {
	name := apiserver.GetCrudServiceName(obj, oper)
	if name != "" {
		return e.GetServiceProfile(name)
	}
	return nil, errors.New("not found")
}

// GetProxyServiceProfile returns the service Profile for a reverse proxy path
func (e *sRolloutV1GwService) GetProxyServiceProfile(path string) (apigw.ServiceProfile, error) {
	name := "_RProxy_" + path
	return e.GetServiceProfile(name)
}

func (e *sRolloutV1GwService) CompleteRegistration(ctx context.Context,
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
		logger.ErrorLog("msg", "failed to register swagger spec", "service", "rollout.RolloutV1", "error", err)
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
				err = rollout.RegisterRolloutV1HandlerWithClient(ctx, mux, cl)
				muxMutex.Unlock()
				if err == nil {
					logger.InfoLog("msg", "registered service rollout.RolloutV1")
					m.Handle("/configs/rollout/v1/", http.StripPrefix("/configs/rollout/v1", mux))
					return
				} else {
					err = errors.Wrap(err, "failed to register")
				}
			} else {
				err = errors.Wrap(err, "failed to create client")
			}
			cancel()
			logger.ErrorLog("msg", "failed to register", "service", "rollout.RolloutV1", "error", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()
	return nil
}

func (e *sRolloutV1GwService) newClient(ctx context.Context, grpcAddr string, rslvr resolver.Interface, devmode bool) (*adapterRolloutV1, error) {
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

	cl := &adapterRolloutV1{conn: client, gw: apigwpkg.MustGetAPIGateway(), gwSvc: e, service: grpcclient.NewRolloutV1Backend(client.ClientConn, e.logger)}
	return cl, nil
}

func init() {

	apigw := apigwpkg.MustGetAPIGateway()

	svcRolloutV1 := sRolloutV1GwService{}
	apigw.Register("rollout.RolloutV1", "rollout/", &svcRolloutV1)
}
