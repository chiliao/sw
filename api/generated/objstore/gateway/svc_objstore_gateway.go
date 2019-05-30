// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package objstoreGwService is a auto generated package.
Input file: svc_objstore.proto
*/
package objstoreGwService

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
	objstore "github.com/pensando/sw/api/generated/objstore"
	grpcclient "github.com/pensando/sw/api/generated/objstore/grpc/client"
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

type sObjstoreV1GwService struct {
	logger     log.Logger
	defSvcProf apigw.ServiceProfile
	svcProf    map[string]apigw.ServiceProfile
}

type adapterObjstoreV1 struct {
	conn    *rpckit.RPCClient
	service objstore.ServiceObjstoreV1Client
	gwSvc   *sObjstoreV1GwService
	gw      apigw.APIGateway
}

func (a adapterObjstoreV1) AutoAddBucket(oldctx oldcontext.Context, t *objstore.Bucket, options ...grpc.CallOption) (*objstore.Bucket, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddBucket")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.CreateOper, "Bucket", t.Tenant, t.Namespace, "", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Bucket)
		return a.service.AutoAddBucket(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.Bucket), err
}

func (a adapterObjstoreV1) AutoAddObject(oldctx oldcontext.Context, t *objstore.Object, options ...grpc.CallOption) (*objstore.Object, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoAddObject")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.CreateOper, "Object", t.Tenant, t.Namespace, "", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Object)
		return a.service.AutoAddObject(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.Object), err
}

func (a adapterObjstoreV1) AutoDeleteBucket(oldctx oldcontext.Context, t *objstore.Bucket, options ...grpc.CallOption) (*objstore.Bucket, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteBucket")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.DeleteOper, "Bucket", t.Tenant, t.Namespace, "", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Bucket)
		return a.service.AutoDeleteBucket(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.Bucket), err
}

func (a adapterObjstoreV1) AutoDeleteObject(oldctx oldcontext.Context, t *objstore.Object, options ...grpc.CallOption) (*objstore.Object, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoDeleteObject")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.DeleteOper, "Object", t.Tenant, t.Namespace, "", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Object)
		return a.service.AutoDeleteObject(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.Object), err
}

func (a adapterObjstoreV1) AutoGetBucket(oldctx oldcontext.Context, t *objstore.Bucket, options ...grpc.CallOption) (*objstore.Bucket, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetBucket")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.GetOper, "Bucket", t.Tenant, t.Namespace, "", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Bucket)
		return a.service.AutoGetBucket(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.Bucket), err
}

func (a adapterObjstoreV1) AutoGetObject(oldctx oldcontext.Context, t *objstore.Object, options ...grpc.CallOption) (*objstore.Object, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoGetObject")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.GetOper, "Object", t.Tenant, t.Namespace, "", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Object)
		return a.service.AutoGetObject(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.Object), err
}

func (a adapterObjstoreV1) AutoListBucket(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*objstore.BucketList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListBucket")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	t.Tenant = ""
	t.Namespace = ""
	oper, kind, tenant, namespace, group, name := apiintf.ListOper, "Bucket", t.Tenant, t.Namespace, "", ""

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListBucket(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.BucketList), err
}

func (a adapterObjstoreV1) AutoListObject(oldctx oldcontext.Context, t *api.ListWatchOptions, options ...grpc.CallOption) (*objstore.ObjectList, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoListObject")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	if t.Tenant == "" {
		user, ok := apigwpkg.UserFromContext(ctx)
		if !ok {
			return nil, errors.New("could not determine user")
		}
		t.Tenant = user.Tenant
	}
	oper, kind, tenant, namespace, group, name := apiintf.ListOper, "Object", t.Tenant, t.Namespace, "", ""

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*api.ListWatchOptions)
		return a.service.AutoListObject(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.ObjectList), err
}

func (a adapterObjstoreV1) AutoUpdateBucket(oldctx oldcontext.Context, t *objstore.Bucket, options ...grpc.CallOption) (*objstore.Bucket, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateBucket")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.UpdateOper, "Bucket", t.Tenant, t.Namespace, "", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Bucket)
		return a.service.AutoUpdateBucket(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.Bucket), err
}

func (a adapterObjstoreV1) AutoUpdateObject(oldctx oldcontext.Context, t *objstore.Object, options ...grpc.CallOption) (*objstore.Object, error) {
	// Not using options for now. Will be passed through context as needed.
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoUpdateObject")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group, name := apiintf.UpdateOper, "Object", t.Tenant, t.Namespace, "", t.Name

	op := authz.NewAPIServerOperation(authz.NewResource(tenant, group, kind, namespace, name), oper)
	ctx = apigwpkg.NewContextWithOperations(ctx, op)

	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Object)
		return a.service.AutoUpdateObject(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, t, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(*objstore.Object), err
}

func (a adapterObjstoreV1) AutoWatchSvcObjstoreV1(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (objstore.ObjstoreV1_AutoWatchSvcObjstoreV1Client, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchSvcObjstoreV1")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "", in.Tenant, in.Namespace, ""
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
				log.Infof("received close notification on websocket [AutoWatchObjstoreV1] (%v/%v)", code, text)
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
		return a.service.AutoWatchSvcObjstoreV1(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(objstore.ObjstoreV1_AutoWatchSvcObjstoreV1Client), err
}

func (a adapterObjstoreV1) AutoWatchBucket(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (objstore.ObjstoreV1_AutoWatchBucketClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchBucket")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	in.Tenant = ""
	in.Namespace = ""
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "Bucket", in.Tenant, in.Namespace, ""
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
				log.Infof("received close notification on websocket [AutoWatchBucket] (%v/%v)", code, text)
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
		return a.service.AutoWatchBucket(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(objstore.ObjstoreV1_AutoWatchBucketClient), err
}

func (a adapterObjstoreV1) AutoWatchObject(oldctx oldcontext.Context, in *api.ListWatchOptions, options ...grpc.CallOption) (objstore.ObjstoreV1_AutoWatchObjectClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("AutoWatchObject")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}

	if in.Tenant == "" {
		user, ok := apigwpkg.UserFromContext(ctx)
		if !ok {
			return nil, errors.New("could not determine user")
		}
		in.Tenant = user.Tenant
	}
	oper, kind, tenant, namespace, group := apiintf.WatchOper, "Object", in.Tenant, in.Namespace, ""
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
				log.Infof("received close notification on websocket [AutoWatchObject] (%v/%v)", code, text)
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
		return a.service.AutoWatchObject(ctx, in)
	}
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(objstore.ObjstoreV1_AutoWatchObjectClient), err
}

func (a adapterObjstoreV1) DownloadFile(oldctx oldcontext.Context, in *objstore.Object, options ...grpc.CallOption) (objstore.ObjstoreV1_DownloadFileClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("DownloadFile")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Object)
		return a.service.DownloadFile(ctx, in)
	}
	apiutils.SetVar(ctx, apiutils.CtxKeyAPIGwBinStreamReq, true)
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(objstore.ObjstoreV1_DownloadFileClient), err
}
func (a adapterObjstoreV1) DownloadFileByPrefix(oldctx oldcontext.Context, in *objstore.Object, options ...grpc.CallOption) (objstore.ObjstoreV1_DownloadFileByPrefixClient, error) {
	ctx := context.Context(oldctx)
	prof, err := a.gwSvc.GetServiceProfile("DownloadFileByPrefix")
	if err != nil {
		return nil, errors.New("unknown service profile")
	}
	fn := func(ctx context.Context, i interface{}) (interface{}, error) {
		in := i.(*objstore.Object)
		return a.service.DownloadFileByPrefix(ctx, in)
	}
	apiutils.SetVar(ctx, apiutils.CtxKeyAPIGwBinStreamReq, true)
	ret, err := a.gw.HandleRequest(ctx, in, prof, fn)
	if ret == nil {
		return nil, err
	}
	return ret.(objstore.ObjstoreV1_DownloadFileByPrefixClient), err
}

func (e *sObjstoreV1GwService) setupSvcProfile() {
	e.defSvcProf = apigwpkg.NewServiceProfile(nil, "", "", apiintf.UnknownOper)
	e.defSvcProf.SetDefaults()
	e.svcProf = make(map[string]apigw.ServiceProfile)

	e.svcProf["AutoAddObject"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Object", "", apiintf.CreateOper)

	e.svcProf["AutoDeleteObject"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Object", "", apiintf.DeleteOper)

	e.svcProf["AutoGetObject"] = apigwpkg.NewServiceProfile(e.defSvcProf, "Object", "", apiintf.GetOper)

	e.svcProf["AutoListObject"] = apigwpkg.NewServiceProfile(e.defSvcProf, "ObjectList", "", apiintf.ListOper)

	e.svcProf["AutoWatchObject"] = apigwpkg.NewServiceProfile(e.defSvcProf, "AutoMsgObjectWatchHelper", "", apiintf.WatchOper)

	e.svcProf["DownloadFile"] = apigwpkg.NewServiceProfile(e.defSvcProf, "", "", apiintf.UnknownOper)

	e.svcProf["DownloadFileByPrefix"] = apigwpkg.NewServiceProfile(e.defSvcProf, "", "", apiintf.UnknownOper)
	e.svcProf["_RProxy_"+"/"+"uploads/images"] = apigwpkg.NewServiceProfile(e.defSvcProf, "", "", apiintf.UnknownOper)
}

// GetDefaultServiceProfile returns the default fallback service profile for this service
func (e *sObjstoreV1GwService) GetDefaultServiceProfile() (apigw.ServiceProfile, error) {
	if e.defSvcProf == nil {
		return nil, errors.New("not found")
	}
	return e.defSvcProf, nil
}

// GetServiceProfile returns the service profile for a given method in this service
func (e *sObjstoreV1GwService) GetServiceProfile(method string) (apigw.ServiceProfile, error) {
	if ret, ok := e.svcProf[method]; ok {
		return ret, nil
	}
	return nil, errors.New("not found")
}

// GetCrudServiceProfile returns the service profile for a auto generated crud operation
func (e *sObjstoreV1GwService) GetCrudServiceProfile(obj string, oper apiintf.APIOperType) (apigw.ServiceProfile, error) {
	name := apiserver.GetCrudServiceName(obj, oper)
	if name != "" {
		return e.GetServiceProfile(name)
	}
	return nil, errors.New("not found")
}

// GetProxyServiceProfile returns the service Profile for a reverse proxy path
func (e *sObjstoreV1GwService) GetProxyServiceProfile(path string) (apigw.ServiceProfile, error) {
	name := "_RProxy_" + path
	return e.GetServiceProfile(name)
}

func (e *sObjstoreV1GwService) CompleteRegistration(ctx context.Context,
	logger log.Logger,
	grpcserver *grpc.Server,
	m *http.ServeMux,
	rslvr resolver.Interface,
	wg *sync.WaitGroup) error {
	apigw := apigwpkg.MustGetAPIGateway()
	// IP:port destination or service discovery key.

	grpcaddr := "pen-vos"
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
		logger.ErrorLog("msg", "failed to register swagger spec", "service", "objstore.ObjstoreV1", "err", err)
	}
	{
		name := "_RProxy_" + "/" + "uploads/images"
		svcProf, err := e.GetServiceProfile(name)
		if err != nil {
			logger.Fatalf("failed to get service profile for [%s](%s)", name, err)
		}

		rproxy, err := apigwpkg.NewRProxyHandler("uploads/images", "/objstore/v1/", "/apis/v1", "pen-vos-http", svcProf)
		if err != nil {
			logger.Fatalf("failed to get proxy handler for [%s](%s)", name, err)
		}
		m.Handle("/objstore/v1/uploads/images/", rproxy)
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
				err = objstore.RegisterObjstoreV1HandlerWithClient(ctx, mux, cl)
				muxMutex.Unlock()
				if err == nil {
					logger.InfoLog("msg", "registered service objstore.ObjstoreV1")
					m.Handle("/objstore/v1/", http.StripPrefix("/objstore/v1", mux))
					return
				} else {
					err = errors.Wrap(err, "failed to register")
				}
			} else {
				err = errors.Wrap(err, "failed to create client")
			}
			cancel()
			logger.ErrorLog("msg", "failed to register", "service", "objstore.ObjstoreV1", "err", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Second):
			}
		}
	}()
	return nil
}

func (e *sObjstoreV1GwService) newClient(ctx context.Context, grpcAddr string, rslvr resolver.Interface, devmode bool) (*adapterObjstoreV1, error) {
	var opts []rpckit.Option
	opts = append(opts, rpckit.WithTLSClientIdentity(globals.APIGw))
	if rslvr != nil {
		opts = append(opts, rpckit.WithBalancer(balancer.New(rslvr)))
	} else {

		opts = append(opts, rpckit.WithRemoteServerName("pen-vos"))
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

	cl := &adapterObjstoreV1{conn: client, gw: apigwpkg.MustGetAPIGateway(), gwSvc: e, service: grpcclient.NewObjstoreV1Backend(client.ClientConn, e.logger)}
	return cl, nil
}

func init() {

	apigw := apigwpkg.MustGetAPIGateway()

	svcObjstoreV1 := sObjstoreV1GwService{}
	apigw.Register("objstore.ObjstoreV1", "/", &svcObjstoreV1)
}
