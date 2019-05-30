// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package diagnosticsApiServer is a auto generated package.
Input file: svc_diagnostics.proto
*/
package diagnosticsApiServer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/pensando/sw/api"
	diagnostics "github.com/pensando/sw/api/generated/diagnostics"
	fieldhooks "github.com/pensando/sw/api/hooks/apiserver/fields"
	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/api/utils"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/apiserver/pkg"
	"github.com/pensando/sw/venice/globals"
	"github.com/pensando/sw/venice/utils/ctxutils"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/rpckit"
	"github.com/pensando/sw/venice/utils/runtime"
)

// dummy vars to suppress unused errors
var _ api.ObjectMeta
var _ listerwatcher.WatcherClient
var _ fmt.Stringer
var _ fieldhooks.Dummy

type sdiagnosticsSvc_diagnosticsBackend struct {
	Services map[string]apiserver.Service
	Messages map[string]apiserver.Message
	logger   log.Logger
	scheme   *runtime.Scheme

	endpointsDiagnosticsV1 *eDiagnosticsV1Endpoints
}

type eDiagnosticsV1Endpoints struct {
	Svc                         sdiagnosticsSvc_diagnosticsBackend
	fnAutoWatchSvcDiagnosticsV1 func(in *api.ListWatchOptions, stream grpc.ServerStream, svcprefix string) error

	fnAutoAddModule    func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoDeleteModule func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoGetModule    func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoListModule   func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoUpdateModule func(ctx context.Context, t interface{}) (interface{}, error)
	fnDebug            func(ctx context.Context, t interface{}) (interface{}, error)

	fnAutoWatchModule func(in *api.ListWatchOptions, stream grpc.ServerStream, svcprefix string) error
}

func (s *sdiagnosticsSvc_diagnosticsBackend) regMsgsFunc(l log.Logger, scheme *runtime.Scheme) {
	l.Infof("registering message for sdiagnosticsSvc_diagnosticsBackend")
	s.Messages = map[string]apiserver.Message{

		"diagnostics.AutoMsgModuleWatchHelper": apisrvpkg.NewMessage("diagnostics.AutoMsgModuleWatchHelper"),
		"diagnostics.ModuleList": apisrvpkg.NewMessage("diagnostics.ModuleList").WithKvListFunc(func(ctx context.Context, kvs kvstore.Interface, options *api.ListWatchOptions, prefix string) (interface{}, error) {

			into := diagnostics.ModuleList{}
			into.Kind = "ModuleList"
			r := diagnostics.Module{}
			r.ObjectMeta = options.ObjectMeta
			key := r.MakeKey(prefix)

			ctx = apiutils.SetVar(ctx, "ObjKind", "diagnostics.Module")
			err := kvs.ListFiltered(ctx, key, &into, *options)
			if err != nil {
				l.ErrorLog("msg", "Object ListFiltered failed", "key", key, "err", err)
				return nil, err
			}
			return into, nil
		}).WithSelfLinkWriter(func(path, ver, prefix string, i interface{}) (interface{}, error) {
			r := i.(diagnostics.ModuleList)
			r.APIVersion = ver
			for i := range r.Items {
				r.Items[i].SelfLink = r.Items[i].MakeURI("configs", ver, prefix)
			}
			return r, nil
		}).WithGetRuntimeObject(func(i interface{}) runtime.Object {
			r := i.(diagnostics.ModuleList)
			return &r
		}),
		// Add a message handler for ListWatch options
		"api.ListWatchOptions": apisrvpkg.NewMessage("api.ListWatchOptions"),
	}

	apisrv.RegisterMessages("diagnostics", s.Messages)
	// add messages to package.
	if pkgMessages == nil {
		pkgMessages = make(map[string]apiserver.Message)
	}
	for k, v := range s.Messages {
		pkgMessages[k] = v
	}
}

func (s *sdiagnosticsSvc_diagnosticsBackend) regSvcsFunc(ctx context.Context, logger log.Logger, grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) {

	{
		srv := apisrvpkg.NewService("diagnostics.DiagnosticsV1")
		s.endpointsDiagnosticsV1.fnAutoWatchSvcDiagnosticsV1 = srv.WatchFromKv

		s.endpointsDiagnosticsV1.fnAutoAddModule = srv.AddMethod("AutoAddModule",
			apisrvpkg.NewMethod(srv, pkgMessages["diagnostics.Module"], pkgMessages["diagnostics.Module"], "diagnostics", "AutoAddModule")).WithOper(apiintf.CreateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			return "", fmt.Errorf("not rest endpoint")
		}).HandleInvocation

		s.endpointsDiagnosticsV1.fnAutoDeleteModule = srv.AddMethod("AutoDeleteModule",
			apisrvpkg.NewMethod(srv, pkgMessages["diagnostics.Module"], pkgMessages["diagnostics.Module"], "diagnostics", "AutoDeleteModule")).WithOper(apiintf.DeleteOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			return "", fmt.Errorf("not rest endpoint")
		}).HandleInvocation

		s.endpointsDiagnosticsV1.fnAutoGetModule = srv.AddMethod("AutoGetModule",
			apisrvpkg.NewMethod(srv, pkgMessages["diagnostics.Module"], pkgMessages["diagnostics.Module"], "diagnostics", "AutoGetModule")).WithOper(apiintf.GetOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(diagnostics.Module)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/", globals.ConfigURIPrefix, "/", "diagnostics/v1/modules/", in.Name), nil
		}).HandleInvocation

		s.endpointsDiagnosticsV1.fnAutoListModule = srv.AddMethod("AutoListModule",
			apisrvpkg.NewMethod(srv, pkgMessages["api.ListWatchOptions"], pkgMessages["diagnostics.ModuleList"], "diagnostics", "AutoListModule")).WithOper(apiintf.ListOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(api.ListWatchOptions)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/", globals.ConfigURIPrefix, "/", "diagnostics/v1/modules/", in.Name), nil
		}).HandleInvocation

		s.endpointsDiagnosticsV1.fnAutoUpdateModule = srv.AddMethod("AutoUpdateModule",
			apisrvpkg.NewMethod(srv, pkgMessages["diagnostics.Module"], pkgMessages["diagnostics.Module"], "diagnostics", "AutoUpdateModule")).WithOper(apiintf.UpdateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(diagnostics.Module)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/", globals.ConfigURIPrefix, "/", "diagnostics/v1/modules/", in.Name), nil
		}).HandleInvocation

		s.endpointsDiagnosticsV1.fnDebug = srv.AddMethod("Debug",
			apisrvpkg.NewMethod(srv, pkgMessages["diagnostics.DiagnosticsRequest"], pkgMessages["diagnostics.DiagnosticsResponse"], "diagnostics", "Debug")).WithOper(apiintf.CreateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(diagnostics.DiagnosticsRequest)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/", globals.ConfigURIPrefix, "/", "diagnostics/v1/tenant/", in.Tenant, "/modules/", in.Name), nil
		}).HandleInvocation

		s.endpointsDiagnosticsV1.fnAutoWatchModule = pkgMessages["diagnostics.Module"].WatchFromKv

		s.Services = map[string]apiserver.Service{
			"diagnostics.DiagnosticsV1": srv,
		}
		apisrv.RegisterService("diagnostics.DiagnosticsV1", srv)
		endpoints := diagnostics.MakeDiagnosticsV1ServerEndpoints(s.endpointsDiagnosticsV1, logger)
		server := diagnostics.MakeGRPCServerDiagnosticsV1(ctx, endpoints, logger)
		diagnostics.RegisterDiagnosticsV1Server(grpcserver.GrpcServer, server)
		svcObjs := []string{"Module"}
		fieldhooks.RegisterImmutableFieldsServiceHooks("diagnostics", "DiagnosticsV1", svcObjs)
	}
}

func (s *sdiagnosticsSvc_diagnosticsBackend) regWatchersFunc(ctx context.Context, logger log.Logger, grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) {

	// Add Watchers
	{

		// Service watcher
		svc := s.Services["diagnostics.DiagnosticsV1"]
		if svc != nil {
			svc.WithKvWatchFunc(func(l log.Logger, options *api.ListWatchOptions, kvs kvstore.Interface, stream interface{}, txfnMap map[string]func(from, to string, i interface{}) (interface{}, error), version, svcprefix string) error {
				key := globals.ConfigRootPrefix + "/diagnostics"
				wstream := stream.(grpc.ServerStream)
				nctx, cancel := context.WithCancel(wstream.Context())
				defer cancel()
				watcher, err := kvs.WatchFiltered(nctx, key, *options)
				if err != nil {
					l.ErrorLog("msg", "error starting Watch for service", "err", err, "service", "DiagnosticsV1")
					return err
				}
				return listerwatcher.SvcWatch(nctx, watcher, wstream, txfnMap, version, l)
			})
		}

		pkgMessages["diagnostics.Module"].WithKvWatchFunc(func(l log.Logger, options *api.ListWatchOptions, kvs kvstore.Interface, stream interface{}, txfn func(from, to string, i interface{}) (interface{}, error), version, svcprefix string) error {
			o := diagnostics.Module{}
			key := o.MakeKey(svcprefix)
			if strings.HasSuffix(key, "//") {
				key = strings.TrimSuffix(key, "/")
			}
			wstream := stream.(diagnostics.DiagnosticsV1_AutoWatchModuleServer)
			nctx, cancel := context.WithCancel(wstream.Context())
			defer cancel()
			id := fmt.Sprintf("%s-%x", ctxutils.GetPeerID(nctx), &key)

			nctx = ctxutils.SetContextID(nctx, id)
			if kvs == nil {
				return fmt.Errorf("Nil KVS")
			}
			nctx = apiutils.SetVar(nctx, "ObjKind", "diagnostics.Module")
			l.InfoLog("msg", "KVWatcher starting watch", "WatcherID", id, "object", "diagnostics.Module")
			watcher, err := kvs.WatchFiltered(nctx, key, *options)
			if err != nil {
				l.ErrorLog("msg", "error starting Watch on KV", "err", err, "WatcherID", id, "bbject", "diagnostics.Module")
				return err
			}
			timer := time.NewTimer(apiserver.DefaultWatchHoldInterval)
			if !timer.Stop() {
				<-timer.C
			}
			running := false
			events := &diagnostics.AutoMsgModuleWatchHelper{}
			sendToStream := func() error {
				l.DebugLog("msg", "writing to stream", "len", len(events.Events))
				if err := wstream.Send(events); err != nil {
					l.ErrorLog("msg", "Stream send error'ed for Order", "err", err, "WatcherID", id, "bbject", "diagnostics.Module")
					return err
				}
				events = &diagnostics.AutoMsgModuleWatchHelper{}
				return nil
			}
			defer l.InfoLog("msg", "exiting watcher", "service", "diagnostics.Module")
			for {
				select {
				case ev, ok := <-watcher.EventChan():
					if !ok {
						l.ErrorLog("msg", "Channel closed for Watcher", "WatcherID", id, "bbject", "diagnostics.Module")
						return nil
					}
					evin, ok := ev.Object.(*diagnostics.Module)
					if !ok {
						status, ok := ev.Object.(*api.Status)
						if !ok {
							return errors.New("unknown error")
						}
						return fmt.Errorf("%v:(%s) %s", status.Code, status.Result, status.Message)
					}
					// XXX-TODO(sanjayt): Avoid a copy and update selflink at enqueue.
					cin, err := evin.Clone(nil)
					if err != nil {
						return fmt.Errorf("unable to clone object (%s)", err)
					}
					in := cin.(*diagnostics.Module)
					in.SelfLink = in.MakeURI(globals.ConfigURIPrefix, "v1", "diagnostics")

					strEvent := &diagnostics.AutoMsgModuleWatchHelper_WatchEvent{
						Type:   string(ev.Type),
						Object: in,
					}
					l.DebugLog("msg", "received Module watch event from KV", "type", ev.Type)
					if version != in.APIVersion {
						i, err := txfn(in.APIVersion, version, in)
						if err != nil {
							l.ErrorLog("msg", "Failed to transform message", "type", "Module", "fromver", in.APIVersion, "tover", version, "WatcherID", id, "bbject", "diagnostics.Module")
							break
						}
						strEvent.Object = i.(*diagnostics.Module)
					}
					events.Events = append(events.Events, strEvent)
					if !running {
						running = true
						timer.Reset(apiserver.DefaultWatchHoldInterval)
					}
					if len(events.Events) >= apiserver.DefaultWatchBatchSize {
						if err = sendToStream(); err != nil {
							return err
						}
						if !timer.Stop() {
							<-timer.C
						}
						timer.Reset(apiserver.DefaultWatchHoldInterval)
					}
				case <-timer.C:
					running = false
					if err = sendToStream(); err != nil {
						return err
					}
				case <-nctx.Done():
					l.DebugLog("msg", "Context cancelled for Watcher", "WatcherID", id, "bbject", "diagnostics.Module")
					return wstream.Context().Err()
				}
			}
		})

	}

}

func (s *sdiagnosticsSvc_diagnosticsBackend) CompleteRegistration(ctx context.Context, logger log.Logger,
	grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) error {
	// register all messages in the package if not done already
	s.logger = logger
	s.scheme = scheme
	registerMessages(logger, scheme)
	registerServices(ctx, logger, grpcserver, scheme)
	registerWatchers(ctx, logger, grpcserver, scheme)
	return nil
}

func (s *sdiagnosticsSvc_diagnosticsBackend) Reset() {
	cleanupRegistration()
}

func (e *eDiagnosticsV1Endpoints) AutoAddModule(ctx context.Context, t diagnostics.Module) (diagnostics.Module, error) {
	r, err := e.fnAutoAddModule(ctx, t)
	if err == nil {
		return r.(diagnostics.Module), err
	}
	return diagnostics.Module{}, err

}
func (e *eDiagnosticsV1Endpoints) AutoDeleteModule(ctx context.Context, t diagnostics.Module) (diagnostics.Module, error) {
	r, err := e.fnAutoDeleteModule(ctx, t)
	if err == nil {
		return r.(diagnostics.Module), err
	}
	return diagnostics.Module{}, err

}
func (e *eDiagnosticsV1Endpoints) AutoGetModule(ctx context.Context, t diagnostics.Module) (diagnostics.Module, error) {
	r, err := e.fnAutoGetModule(ctx, t)
	if err == nil {
		return r.(diagnostics.Module), err
	}
	return diagnostics.Module{}, err

}
func (e *eDiagnosticsV1Endpoints) AutoListModule(ctx context.Context, t api.ListWatchOptions) (diagnostics.ModuleList, error) {
	r, err := e.fnAutoListModule(ctx, t)
	if err == nil {
		return r.(diagnostics.ModuleList), err
	}
	return diagnostics.ModuleList{}, err

}
func (e *eDiagnosticsV1Endpoints) AutoUpdateModule(ctx context.Context, t diagnostics.Module) (diagnostics.Module, error) {
	r, err := e.fnAutoUpdateModule(ctx, t)
	if err == nil {
		return r.(diagnostics.Module), err
	}
	return diagnostics.Module{}, err

}
func (e *eDiagnosticsV1Endpoints) Debug(ctx context.Context, t diagnostics.DiagnosticsRequest) (diagnostics.DiagnosticsResponse, error) {
	r, err := e.fnDebug(ctx, t)
	if err == nil {
		return r.(diagnostics.DiagnosticsResponse), err
	}
	return diagnostics.DiagnosticsResponse{}, err

}

func (e *eDiagnosticsV1Endpoints) AutoWatchModule(in *api.ListWatchOptions, stream diagnostics.DiagnosticsV1_AutoWatchModuleServer) error {
	return e.fnAutoWatchModule(in, stream, "diagnostics")
}
func (e *eDiagnosticsV1Endpoints) AutoWatchSvcDiagnosticsV1(in *api.ListWatchOptions, stream diagnostics.DiagnosticsV1_AutoWatchSvcDiagnosticsV1Server) error {
	return e.fnAutoWatchSvcDiagnosticsV1(in, stream, "")
}

func init() {
	apisrv = apisrvpkg.MustGetAPIServer()

	svc := sdiagnosticsSvc_diagnosticsBackend{}
	addMsgRegFunc(svc.regMsgsFunc)
	addSvcRegFunc(svc.regSvcsFunc)
	addWatcherRegFunc(svc.regWatchersFunc)

	{
		e := eDiagnosticsV1Endpoints{Svc: svc}
		svc.endpointsDiagnosticsV1 = &e
	}
	apisrv.Register("diagnostics.svc_diagnostics.proto", &svc)
}
