// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package telemetryApiServer is a auto generated package.
Input file: protos/telemetry.proto
*/
package telemetryApiServer

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
	"google.golang.org/grpc"

	"github.com/pensando/sw/api"
	telemetry "github.com/pensando/sw/api/generated/telemetry"
	"github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/apiserver/pkg"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/rpckit"
	"github.com/pensando/sw/venice/utils/runtime"
)

// dummy vars to suppress unused errors
var _ api.ObjectMeta
var _ listerwatcher.WatcherClient
var _ fmt.Stringer

type stelemetryTelemetryBackend struct {
	Services map[string]apiserver.Service
	Messages map[string]apiserver.Message

	endpointsFlowExportPolicyV1 *eFlowExportPolicyV1Endpoints
	endpointsFwlogPolicyV1      *eFwlogPolicyV1Endpoints
	endpointsStatsPolicyV1      *eStatsPolicyV1Endpoints
}

type eFlowExportPolicyV1Endpoints struct {
	Svc stelemetryTelemetryBackend

	fnAutoAddFlowExportPolicy    func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoDeleteFlowExportPolicy func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoGetFlowExportPolicy    func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoListFlowExportPolicy   func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoUpdateFlowExportPolicy func(ctx context.Context, t interface{}) (interface{}, error)

	fnAutoWatchFlowExportPolicy func(in *api.ListWatchOptions, stream grpc.ServerStream, svcprefix string) error
}
type eFwlogPolicyV1Endpoints struct {
	Svc stelemetryTelemetryBackend

	fnAutoAddFwlogPolicy    func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoDeleteFwlogPolicy func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoGetFwlogPolicy    func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoListFwlogPolicy   func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoUpdateFwlogPolicy func(ctx context.Context, t interface{}) (interface{}, error)

	fnAutoWatchFwlogPolicy func(in *api.ListWatchOptions, stream grpc.ServerStream, svcprefix string) error
}
type eStatsPolicyV1Endpoints struct {
	Svc stelemetryTelemetryBackend

	fnAutoAddStatsPolicy    func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoDeleteStatsPolicy func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoGetStatsPolicy    func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoListStatsPolicy   func(ctx context.Context, t interface{}) (interface{}, error)
	fnAutoUpdateStatsPolicy func(ctx context.Context, t interface{}) (interface{}, error)

	fnAutoWatchStatsPolicy func(in *api.ListWatchOptions, stream grpc.ServerStream, svcprefix string) error
}

func (s *stelemetryTelemetryBackend) CompleteRegistration(ctx context.Context, logger log.Logger,
	grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) error {
	s.Messages = map[string]apiserver.Message{

		"telemetry.AutoMsgFlowExportPolicyWatchHelper": apisrvpkg.NewMessage("telemetry.AutoMsgFlowExportPolicyWatchHelper"),
		"telemetry.AutoMsgFwlogPolicyWatchHelper":      apisrvpkg.NewMessage("telemetry.AutoMsgFwlogPolicyWatchHelper"),
		"telemetry.AutoMsgStatsPolicyWatchHelper":      apisrvpkg.NewMessage("telemetry.AutoMsgStatsPolicyWatchHelper"),
		"telemetry.FlowExportPolicy": apisrvpkg.NewMessage("telemetry.FlowExportPolicy").WithKeyGenerator(func(i interface{}, prefix string) string {
			if i == nil {
				r := telemetry.FlowExportPolicy{}
				return r.MakeKey(prefix)
			}
			r := i.(telemetry.FlowExportPolicy)
			return r.MakeKey(prefix)
		}).WithObjectVersionWriter(func(i interface{}, version string) interface{} {
			r := i.(telemetry.FlowExportPolicy)
			r.APIVersion = version
			return r
		}).WithKvUpdater(func(ctx context.Context, kvs kvstore.Interface, i interface{}, prefix string, create, ignoreStatus bool) (interface{}, error) {
			r := i.(telemetry.FlowExportPolicy)
			key := r.MakeKey(prefix)
			r.Kind = "FlowExportPolicy"
			var err error
			if create {
				err = kvs.Create(ctx, key, &r)
				err = errors.Wrap(err, "KV create failed")
			} else {
				if ignoreStatus {
					updateFunc := func(obj runtime.Object) (runtime.Object, error) {
						saved := obj.(*telemetry.FlowExportPolicy)
						if r.ResourceVersion != "" && r.ResourceVersion != saved.ResourceVersion {
							return nil, fmt.Errorf("Resource Version specified does not match Object version")
						}
						r.Status = saved.Status
						return &r, nil
					}
					into := &telemetry.FlowExportPolicy{}
					err = kvs.ConsistentUpdate(ctx, key, into, updateFunc)
				} else {
					if r.ResourceVersion != "" {
						logger.Infof("resource version is specified %s\n", r.ResourceVersion)
						err = kvs.Update(ctx, key, &r, kvstore.Compare(kvstore.WithVersion(key), "=", r.ResourceVersion))
					} else {
						err = kvs.Update(ctx, key, &r)
					}
					err = errors.Wrap(err, "KV update failed")
				}
			}
			return r, err
		}).WithKvTxnUpdater(func(ctx context.Context, txn kvstore.Txn, i interface{}, prefix string, create bool) error {
			r := i.(telemetry.FlowExportPolicy)
			key := r.MakeKey(prefix)
			var err error
			if create {
				err = txn.Create(key, &r)
				err = errors.Wrap(err, "KV transaction create failed")
			} else {
				err = txn.Update(key, &r)
				err = errors.Wrap(err, "KV transaction update failed")
			}
			return err
		}).WithUUIDWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.FlowExportPolicy)
			r.UUID = uuid.NewV4().String()
			return r, nil
		}).WithCreationTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.FlowExportPolicy)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.CreationTime.Timestamp = *ts
			}
			return r, err
		}).WithModTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.FlowExportPolicy)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.ModTime.Timestamp = *ts
			}
			return r, err
		}).WithSelfLinkWriter(func(path string, i interface{}) (interface{}, error) {
			r := i.(telemetry.FlowExportPolicy)
			r.SelfLink = path
			return r, nil
		}).WithKvGetter(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := telemetry.FlowExportPolicy{}
			err := kvs.Get(ctx, key, &r)
			err = errors.Wrap(err, "KV get failed")
			return r, err
		}).WithKvDelFunc(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := telemetry.FlowExportPolicy{}
			err := kvs.Delete(ctx, key, &r)
			return r, err
		}).WithKvTxnDelFunc(func(ctx context.Context, txn kvstore.Txn, key string) error {
			return txn.Delete(key)
		}).WithValidate(func(i interface{}, ver string, ignoreStatus bool) error {
			r := i.(telemetry.FlowExportPolicy)
			if !r.Validate(ver, ignoreStatus) {
				return fmt.Errorf("Default Validation failed")
			}
			return nil
		}),
		"telemetry.FlowExportPolicyList": apisrvpkg.NewMessage("telemetry.FlowExportPolicyList").WithKvListFunc(func(ctx context.Context, kvs kvstore.Interface, options *api.ListWatchOptions, prefix string) (interface{}, error) {

			into := telemetry.FlowExportPolicyList{}
			r := telemetry.FlowExportPolicy{}
			r.ObjectMeta = options.ObjectMeta
			key := r.MakeKey(prefix)
			err := kvs.List(ctx, key, &into)
			if err != nil {
				return nil, err
			}
			return into, nil
		}),
		"telemetry.FlowExportSpec":   apisrvpkg.NewMessage("telemetry.FlowExportSpec"),
		"telemetry.FlowExportStatus": apisrvpkg.NewMessage("telemetry.FlowExportStatus"),
		"telemetry.FlowExportTarget": apisrvpkg.NewMessage("telemetry.FlowExportTarget"),
		"telemetry.FwlogPolicy": apisrvpkg.NewMessage("telemetry.FwlogPolicy").WithKeyGenerator(func(i interface{}, prefix string) string {
			if i == nil {
				r := telemetry.FwlogPolicy{}
				return r.MakeKey(prefix)
			}
			r := i.(telemetry.FwlogPolicy)
			return r.MakeKey(prefix)
		}).WithObjectVersionWriter(func(i interface{}, version string) interface{} {
			r := i.(telemetry.FwlogPolicy)
			r.APIVersion = version
			return r
		}).WithKvUpdater(func(ctx context.Context, kvs kvstore.Interface, i interface{}, prefix string, create, ignoreStatus bool) (interface{}, error) {
			r := i.(telemetry.FwlogPolicy)
			key := r.MakeKey(prefix)
			r.Kind = "FwlogPolicy"
			var err error
			if create {
				err = kvs.Create(ctx, key, &r)
				err = errors.Wrap(err, "KV create failed")
			} else {
				if ignoreStatus {
					updateFunc := func(obj runtime.Object) (runtime.Object, error) {
						saved := obj.(*telemetry.FwlogPolicy)
						if r.ResourceVersion != "" && r.ResourceVersion != saved.ResourceVersion {
							return nil, fmt.Errorf("Resource Version specified does not match Object version")
						}
						r.Status = saved.Status
						return &r, nil
					}
					into := &telemetry.FwlogPolicy{}
					err = kvs.ConsistentUpdate(ctx, key, into, updateFunc)
				} else {
					if r.ResourceVersion != "" {
						logger.Infof("resource version is specified %s\n", r.ResourceVersion)
						err = kvs.Update(ctx, key, &r, kvstore.Compare(kvstore.WithVersion(key), "=", r.ResourceVersion))
					} else {
						err = kvs.Update(ctx, key, &r)
					}
					err = errors.Wrap(err, "KV update failed")
				}
			}
			return r, err
		}).WithKvTxnUpdater(func(ctx context.Context, txn kvstore.Txn, i interface{}, prefix string, create bool) error {
			r := i.(telemetry.FwlogPolicy)
			key := r.MakeKey(prefix)
			var err error
			if create {
				err = txn.Create(key, &r)
				err = errors.Wrap(err, "KV transaction create failed")
			} else {
				err = txn.Update(key, &r)
				err = errors.Wrap(err, "KV transaction update failed")
			}
			return err
		}).WithUUIDWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.FwlogPolicy)
			r.UUID = uuid.NewV4().String()
			return r, nil
		}).WithCreationTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.FwlogPolicy)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.CreationTime.Timestamp = *ts
			}
			return r, err
		}).WithModTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.FwlogPolicy)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.ModTime.Timestamp = *ts
			}
			return r, err
		}).WithSelfLinkWriter(func(path string, i interface{}) (interface{}, error) {
			r := i.(telemetry.FwlogPolicy)
			r.SelfLink = path
			return r, nil
		}).WithKvGetter(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := telemetry.FwlogPolicy{}
			err := kvs.Get(ctx, key, &r)
			err = errors.Wrap(err, "KV get failed")
			return r, err
		}).WithKvDelFunc(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := telemetry.FwlogPolicy{}
			err := kvs.Delete(ctx, key, &r)
			return r, err
		}).WithKvTxnDelFunc(func(ctx context.Context, txn kvstore.Txn, key string) error {
			return txn.Delete(key)
		}).WithValidate(func(i interface{}, ver string, ignoreStatus bool) error {
			r := i.(telemetry.FwlogPolicy)
			if !r.Validate(ver, ignoreStatus) {
				return fmt.Errorf("Default Validation failed")
			}
			return nil
		}),
		"telemetry.FwlogPolicyList": apisrvpkg.NewMessage("telemetry.FwlogPolicyList").WithKvListFunc(func(ctx context.Context, kvs kvstore.Interface, options *api.ListWatchOptions, prefix string) (interface{}, error) {

			into := telemetry.FwlogPolicyList{}
			r := telemetry.FwlogPolicy{}
			r.ObjectMeta = options.ObjectMeta
			key := r.MakeKey(prefix)
			err := kvs.List(ctx, key, &into)
			if err != nil {
				return nil, err
			}
			return into, nil
		}),
		"telemetry.FwlogSpec":   apisrvpkg.NewMessage("telemetry.FwlogSpec"),
		"telemetry.FwlogStatus": apisrvpkg.NewMessage("telemetry.FwlogStatus"),
		"telemetry.StatsPolicy": apisrvpkg.NewMessage("telemetry.StatsPolicy").WithKeyGenerator(func(i interface{}, prefix string) string {
			if i == nil {
				r := telemetry.StatsPolicy{}
				return r.MakeKey(prefix)
			}
			r := i.(telemetry.StatsPolicy)
			return r.MakeKey(prefix)
		}).WithObjectVersionWriter(func(i interface{}, version string) interface{} {
			r := i.(telemetry.StatsPolicy)
			r.APIVersion = version
			return r
		}).WithKvUpdater(func(ctx context.Context, kvs kvstore.Interface, i interface{}, prefix string, create, ignoreStatus bool) (interface{}, error) {
			r := i.(telemetry.StatsPolicy)
			key := r.MakeKey(prefix)
			r.Kind = "StatsPolicy"
			var err error
			if create {
				err = kvs.Create(ctx, key, &r)
				err = errors.Wrap(err, "KV create failed")
			} else {
				if ignoreStatus {
					updateFunc := func(obj runtime.Object) (runtime.Object, error) {
						saved := obj.(*telemetry.StatsPolicy)
						if r.ResourceVersion != "" && r.ResourceVersion != saved.ResourceVersion {
							return nil, fmt.Errorf("Resource Version specified does not match Object version")
						}
						r.Status = saved.Status
						return &r, nil
					}
					into := &telemetry.StatsPolicy{}
					err = kvs.ConsistentUpdate(ctx, key, into, updateFunc)
				} else {
					if r.ResourceVersion != "" {
						logger.Infof("resource version is specified %s\n", r.ResourceVersion)
						err = kvs.Update(ctx, key, &r, kvstore.Compare(kvstore.WithVersion(key), "=", r.ResourceVersion))
					} else {
						err = kvs.Update(ctx, key, &r)
					}
					err = errors.Wrap(err, "KV update failed")
				}
			}
			return r, err
		}).WithKvTxnUpdater(func(ctx context.Context, txn kvstore.Txn, i interface{}, prefix string, create bool) error {
			r := i.(telemetry.StatsPolicy)
			key := r.MakeKey(prefix)
			var err error
			if create {
				err = txn.Create(key, &r)
				err = errors.Wrap(err, "KV transaction create failed")
			} else {
				err = txn.Update(key, &r)
				err = errors.Wrap(err, "KV transaction update failed")
			}
			return err
		}).WithUUIDWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.StatsPolicy)
			r.UUID = uuid.NewV4().String()
			return r, nil
		}).WithCreationTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.StatsPolicy)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.CreationTime.Timestamp = *ts
			}
			return r, err
		}).WithModTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(telemetry.StatsPolicy)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.ModTime.Timestamp = *ts
			}
			return r, err
		}).WithSelfLinkWriter(func(path string, i interface{}) (interface{}, error) {
			r := i.(telemetry.StatsPolicy)
			r.SelfLink = path
			return r, nil
		}).WithKvGetter(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := telemetry.StatsPolicy{}
			err := kvs.Get(ctx, key, &r)
			err = errors.Wrap(err, "KV get failed")
			return r, err
		}).WithKvDelFunc(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := telemetry.StatsPolicy{}
			err := kvs.Delete(ctx, key, &r)
			return r, err
		}).WithKvTxnDelFunc(func(ctx context.Context, txn kvstore.Txn, key string) error {
			return txn.Delete(key)
		}).WithValidate(func(i interface{}, ver string, ignoreStatus bool) error {
			r := i.(telemetry.StatsPolicy)
			if !r.Validate(ver, ignoreStatus) {
				return fmt.Errorf("Default Validation failed")
			}
			return nil
		}),
		"telemetry.StatsPolicyList": apisrvpkg.NewMessage("telemetry.StatsPolicyList").WithKvListFunc(func(ctx context.Context, kvs kvstore.Interface, options *api.ListWatchOptions, prefix string) (interface{}, error) {

			into := telemetry.StatsPolicyList{}
			r := telemetry.StatsPolicy{}
			r.ObjectMeta = options.ObjectMeta
			key := r.MakeKey(prefix)
			err := kvs.List(ctx, key, &into)
			if err != nil {
				return nil, err
			}
			return into, nil
		}),
		"telemetry.StatsSpec":   apisrvpkg.NewMessage("telemetry.StatsSpec"),
		"telemetry.StatsStatus": apisrvpkg.NewMessage("telemetry.StatsStatus"),
		// Add a message handler for ListWatch options
		"api.ListWatchOptions": apisrvpkg.NewMessage("api.ListWatchOptions"),
	}

	scheme.AddKnownTypes(
		&telemetry.FlowExportPolicy{},
		&telemetry.FwlogPolicy{},
		&telemetry.StatsPolicy{},
	)

	apisrv.RegisterMessages("telemetry", s.Messages)

	{
		srv := apisrvpkg.NewService("FlowExportPolicyV1")

		s.endpointsFlowExportPolicyV1.fnAutoAddFlowExportPolicy = srv.AddMethod("AutoAddFlowExportPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.FlowExportPolicy"], s.Messages["telemetry.FlowExportPolicy"], "flowExportPolicy", "AutoAddFlowExportPolicy")).WithOper(apiserver.CreateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.FlowExportPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "flowExportPolicy/", in.Tenant, "/flowExportPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFlowExportPolicyV1.fnAutoDeleteFlowExportPolicy = srv.AddMethod("AutoDeleteFlowExportPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.FlowExportPolicy"], s.Messages["telemetry.FlowExportPolicy"], "flowExportPolicy", "AutoDeleteFlowExportPolicy")).WithOper(apiserver.DeleteOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.FlowExportPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "flowExportPolicy/", in.Tenant, "/flowExportPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFlowExportPolicyV1.fnAutoGetFlowExportPolicy = srv.AddMethod("AutoGetFlowExportPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.FlowExportPolicy"], s.Messages["telemetry.FlowExportPolicy"], "flowExportPolicy", "AutoGetFlowExportPolicy")).WithOper(apiserver.GetOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.FlowExportPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "flowExportPolicy/", in.Tenant, "/flowExportPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFlowExportPolicyV1.fnAutoListFlowExportPolicy = srv.AddMethod("AutoListFlowExportPolicy",
			apisrvpkg.NewMethod(s.Messages["api.ListWatchOptions"], s.Messages["telemetry.FlowExportPolicyList"], "flowExportPolicy", "AutoListFlowExportPolicy")).WithOper(apiserver.ListOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(api.ListWatchOptions)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "flowExportPolicy/", in.Tenant, "/flowExportPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFlowExportPolicyV1.fnAutoUpdateFlowExportPolicy = srv.AddMethod("AutoUpdateFlowExportPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.FlowExportPolicy"], s.Messages["telemetry.FlowExportPolicy"], "flowExportPolicy", "AutoUpdateFlowExportPolicy")).WithOper(apiserver.UpdateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.FlowExportPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "flowExportPolicy/", in.Tenant, "/flowExportPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFlowExportPolicyV1.fnAutoWatchFlowExportPolicy = s.Messages["telemetry.FlowExportPolicy"].WatchFromKv

		s.Services = map[string]apiserver.Service{
			"telemetry.FlowExportPolicyV1": srv,
		}
		apisrv.RegisterService("telemetry.FlowExportPolicyV1", srv)
		endpoints := telemetry.MakeFlowExportPolicyV1ServerEndpoints(s.endpointsFlowExportPolicyV1, logger)
		server := telemetry.MakeGRPCServerFlowExportPolicyV1(ctx, endpoints, logger)
		telemetry.RegisterFlowExportPolicyV1Server(grpcserver.GrpcServer, server)
	}
	// Add Watchers
	{

		s.Messages["telemetry.FlowExportPolicy"].WithKvWatchFunc(func(l log.Logger, options *api.ListWatchOptions, kvs kvstore.Interface, stream interface{}, txfn func(from, to string, i interface{}) (interface{}, error), version, svcprefix string) error {
			o := telemetry.FlowExportPolicy{}
			key := o.MakeKey(svcprefix)
			if strings.HasSuffix(key, "//") {
				key = strings.TrimSuffix(key, "/")
			}
			wstream := stream.(telemetry.FlowExportPolicyV1_AutoWatchFlowExportPolicyServer)
			nctx, cancel := context.WithCancel(wstream.Context())
			defer cancel()
			if kvs == nil {
				return fmt.Errorf("Nil KVS")
			}
			watcher, err := kvs.PrefixWatch(nctx, key, options.ResourceVersion)
			if err != nil {
				l.ErrorLog("msg", "error starting Watch on KV", "error", err, "object", "FlowExportPolicy")
				return err
			}
			for {
				select {
				case ev, ok := <-watcher.EventChan():
					if !ok {
						l.DebugLog("Channel closed for FlowExportPolicy Watcher")
						return nil
					}
					in, ok := ev.Object.(*telemetry.FlowExportPolicy)
					if !ok {
						status, ok := ev.Object.(*api.Status)
						if !ok {
							return errors.New("unknown error")
						}
						return fmt.Errorf("%v:(%s) %s", status.Code, status.Result, status.Message)
					}
					strEvent := telemetry.AutoMsgFlowExportPolicyWatchHelper{
						Type:   string(ev.Type),
						Object: in,
					}
					l.DebugLog("msg", "received FlowExportPolicy watch event from KV", "type", ev.Type)
					if version != in.APIVersion {
						i, err := txfn(in.APIVersion, version, in)
						if err != nil {
							l.ErrorLog("msg", "Failed to transform message", "type", "FlowExportPolicy", "fromver", in.APIVersion, "tover", version)
							break
						}
						strEvent.Object = i.(*telemetry.FlowExportPolicy)
					}
					l.DebugLog("msg", "writing to stream")
					if err := wstream.Send(&strEvent); err != nil {
						l.DebugLog("msg", "Stream send error'ed for FlowExportPolicy", "error", err)
						return err
					}
				case <-nctx.Done():
					l.DebugLog("msg", "Context cancelled for FlowExportPolicy Watcher")
					return wstream.Context().Err()
				}
			}
		})

	}

	{
		srv := apisrvpkg.NewService("FwlogPolicyV1")

		s.endpointsFwlogPolicyV1.fnAutoAddFwlogPolicy = srv.AddMethod("AutoAddFwlogPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.FwlogPolicy"], s.Messages["telemetry.FwlogPolicy"], "fwlogPolicy", "AutoAddFwlogPolicy")).WithOper(apiserver.CreateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.FwlogPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "fwlogPolicy/", in.Tenant, "/fwlogPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFwlogPolicyV1.fnAutoDeleteFwlogPolicy = srv.AddMethod("AutoDeleteFwlogPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.FwlogPolicy"], s.Messages["telemetry.FwlogPolicy"], "fwlogPolicy", "AutoDeleteFwlogPolicy")).WithOper(apiserver.DeleteOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.FwlogPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "fwlogPolicy/", in.Tenant, "/fwlogPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFwlogPolicyV1.fnAutoGetFwlogPolicy = srv.AddMethod("AutoGetFwlogPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.FwlogPolicy"], s.Messages["telemetry.FwlogPolicy"], "fwlogPolicy", "AutoGetFwlogPolicy")).WithOper(apiserver.GetOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.FwlogPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "fwlogPolicy/", in.Tenant, "/fwlogPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFwlogPolicyV1.fnAutoListFwlogPolicy = srv.AddMethod("AutoListFwlogPolicy",
			apisrvpkg.NewMethod(s.Messages["api.ListWatchOptions"], s.Messages["telemetry.FwlogPolicyList"], "fwlogPolicy", "AutoListFwlogPolicy")).WithOper(apiserver.ListOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(api.ListWatchOptions)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "fwlogPolicy/", in.Tenant, "/fwlogPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFwlogPolicyV1.fnAutoUpdateFwlogPolicy = srv.AddMethod("AutoUpdateFwlogPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.FwlogPolicy"], s.Messages["telemetry.FwlogPolicy"], "fwlogPolicy", "AutoUpdateFwlogPolicy")).WithOper(apiserver.UpdateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.FwlogPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "fwlogPolicy/", in.Tenant, "/fwlogPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsFwlogPolicyV1.fnAutoWatchFwlogPolicy = s.Messages["telemetry.FwlogPolicy"].WatchFromKv

		s.Services = map[string]apiserver.Service{
			"telemetry.FwlogPolicyV1": srv,
		}
		apisrv.RegisterService("telemetry.FwlogPolicyV1", srv)
		endpoints := telemetry.MakeFwlogPolicyV1ServerEndpoints(s.endpointsFwlogPolicyV1, logger)
		server := telemetry.MakeGRPCServerFwlogPolicyV1(ctx, endpoints, logger)
		telemetry.RegisterFwlogPolicyV1Server(grpcserver.GrpcServer, server)
	}
	// Add Watchers
	{

		s.Messages["telemetry.FwlogPolicy"].WithKvWatchFunc(func(l log.Logger, options *api.ListWatchOptions, kvs kvstore.Interface, stream interface{}, txfn func(from, to string, i interface{}) (interface{}, error), version, svcprefix string) error {
			o := telemetry.FwlogPolicy{}
			key := o.MakeKey(svcprefix)
			if strings.HasSuffix(key, "//") {
				key = strings.TrimSuffix(key, "/")
			}
			wstream := stream.(telemetry.FwlogPolicyV1_AutoWatchFwlogPolicyServer)
			nctx, cancel := context.WithCancel(wstream.Context())
			defer cancel()
			if kvs == nil {
				return fmt.Errorf("Nil KVS")
			}
			watcher, err := kvs.PrefixWatch(nctx, key, options.ResourceVersion)
			if err != nil {
				l.ErrorLog("msg", "error starting Watch on KV", "error", err, "object", "FwlogPolicy")
				return err
			}
			for {
				select {
				case ev, ok := <-watcher.EventChan():
					if !ok {
						l.DebugLog("Channel closed for FwlogPolicy Watcher")
						return nil
					}
					in, ok := ev.Object.(*telemetry.FwlogPolicy)
					if !ok {
						status, ok := ev.Object.(*api.Status)
						if !ok {
							return errors.New("unknown error")
						}
						return fmt.Errorf("%v:(%s) %s", status.Code, status.Result, status.Message)
					}
					strEvent := telemetry.AutoMsgFwlogPolicyWatchHelper{
						Type:   string(ev.Type),
						Object: in,
					}
					l.DebugLog("msg", "received FwlogPolicy watch event from KV", "type", ev.Type)
					if version != in.APIVersion {
						i, err := txfn(in.APIVersion, version, in)
						if err != nil {
							l.ErrorLog("msg", "Failed to transform message", "type", "FwlogPolicy", "fromver", in.APIVersion, "tover", version)
							break
						}
						strEvent.Object = i.(*telemetry.FwlogPolicy)
					}
					l.DebugLog("msg", "writing to stream")
					if err := wstream.Send(&strEvent); err != nil {
						l.DebugLog("msg", "Stream send error'ed for FwlogPolicy", "error", err)
						return err
					}
				case <-nctx.Done():
					l.DebugLog("msg", "Context cancelled for FwlogPolicy Watcher")
					return wstream.Context().Err()
				}
			}
		})

	}

	{
		srv := apisrvpkg.NewService("StatsPolicyV1")

		s.endpointsStatsPolicyV1.fnAutoAddStatsPolicy = srv.AddMethod("AutoAddStatsPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.StatsPolicy"], s.Messages["telemetry.StatsPolicy"], "statsPolicy", "AutoAddStatsPolicy")).WithOper(apiserver.CreateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.StatsPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "statsPolicy/", in.Tenant, "/StatsPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsStatsPolicyV1.fnAutoDeleteStatsPolicy = srv.AddMethod("AutoDeleteStatsPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.StatsPolicy"], s.Messages["telemetry.StatsPolicy"], "statsPolicy", "AutoDeleteStatsPolicy")).WithOper(apiserver.DeleteOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.StatsPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "statsPolicy/", in.Tenant, "/StatsPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsStatsPolicyV1.fnAutoGetStatsPolicy = srv.AddMethod("AutoGetStatsPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.StatsPolicy"], s.Messages["telemetry.StatsPolicy"], "statsPolicy", "AutoGetStatsPolicy")).WithOper(apiserver.GetOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.StatsPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "statsPolicy/", in.Tenant, "/StatsPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsStatsPolicyV1.fnAutoListStatsPolicy = srv.AddMethod("AutoListStatsPolicy",
			apisrvpkg.NewMethod(s.Messages["api.ListWatchOptions"], s.Messages["telemetry.StatsPolicyList"], "statsPolicy", "AutoListStatsPolicy")).WithOper(apiserver.ListOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(api.ListWatchOptions)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "statsPolicy/", in.Tenant, "/StatsPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsStatsPolicyV1.fnAutoUpdateStatsPolicy = srv.AddMethod("AutoUpdateStatsPolicy",
			apisrvpkg.NewMethod(s.Messages["telemetry.StatsPolicy"], s.Messages["telemetry.StatsPolicy"], "statsPolicy", "AutoUpdateStatsPolicy")).WithOper(apiserver.UpdateOper).WithVersion("v1").WithMakeURI(func(i interface{}) (string, error) {
			in, ok := i.(telemetry.StatsPolicy)
			if !ok {
				return "", fmt.Errorf("wrong type")
			}
			return fmt.Sprint("/v1/", "statsPolicy/", in.Tenant, "/StatsPolicy/", in.Name), nil
		}).HandleInvocation

		s.endpointsStatsPolicyV1.fnAutoWatchStatsPolicy = s.Messages["telemetry.StatsPolicy"].WatchFromKv

		s.Services = map[string]apiserver.Service{
			"telemetry.StatsPolicyV1": srv,
		}
		apisrv.RegisterService("telemetry.StatsPolicyV1", srv)
		endpoints := telemetry.MakeStatsPolicyV1ServerEndpoints(s.endpointsStatsPolicyV1, logger)
		server := telemetry.MakeGRPCServerStatsPolicyV1(ctx, endpoints, logger)
		telemetry.RegisterStatsPolicyV1Server(grpcserver.GrpcServer, server)
	}
	// Add Watchers
	{

		s.Messages["telemetry.StatsPolicy"].WithKvWatchFunc(func(l log.Logger, options *api.ListWatchOptions, kvs kvstore.Interface, stream interface{}, txfn func(from, to string, i interface{}) (interface{}, error), version, svcprefix string) error {
			o := telemetry.StatsPolicy{}
			key := o.MakeKey(svcprefix)
			if strings.HasSuffix(key, "//") {
				key = strings.TrimSuffix(key, "/")
			}
			wstream := stream.(telemetry.StatsPolicyV1_AutoWatchStatsPolicyServer)
			nctx, cancel := context.WithCancel(wstream.Context())
			defer cancel()
			if kvs == nil {
				return fmt.Errorf("Nil KVS")
			}
			watcher, err := kvs.PrefixWatch(nctx, key, options.ResourceVersion)
			if err != nil {
				l.ErrorLog("msg", "error starting Watch on KV", "error", err, "object", "StatsPolicy")
				return err
			}
			for {
				select {
				case ev, ok := <-watcher.EventChan():
					if !ok {
						l.DebugLog("Channel closed for StatsPolicy Watcher")
						return nil
					}
					in, ok := ev.Object.(*telemetry.StatsPolicy)
					if !ok {
						status, ok := ev.Object.(*api.Status)
						if !ok {
							return errors.New("unknown error")
						}
						return fmt.Errorf("%v:(%s) %s", status.Code, status.Result, status.Message)
					}
					strEvent := telemetry.AutoMsgStatsPolicyWatchHelper{
						Type:   string(ev.Type),
						Object: in,
					}
					l.DebugLog("msg", "received StatsPolicy watch event from KV", "type", ev.Type)
					if version != in.APIVersion {
						i, err := txfn(in.APIVersion, version, in)
						if err != nil {
							l.ErrorLog("msg", "Failed to transform message", "type", "StatsPolicy", "fromver", in.APIVersion, "tover", version)
							break
						}
						strEvent.Object = i.(*telemetry.StatsPolicy)
					}
					l.DebugLog("msg", "writing to stream")
					if err := wstream.Send(&strEvent); err != nil {
						l.DebugLog("msg", "Stream send error'ed for StatsPolicy", "error", err)
						return err
					}
				case <-nctx.Done():
					l.DebugLog("msg", "Context cancelled for StatsPolicy Watcher")
					return wstream.Context().Err()
				}
			}
		})

	}

	return nil
}

func (e *eFlowExportPolicyV1Endpoints) AutoAddFlowExportPolicy(ctx context.Context, t telemetry.FlowExportPolicy) (telemetry.FlowExportPolicy, error) {
	r, err := e.fnAutoAddFlowExportPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FlowExportPolicy), err
	}
	return telemetry.FlowExportPolicy{}, err

}
func (e *eFlowExportPolicyV1Endpoints) AutoDeleteFlowExportPolicy(ctx context.Context, t telemetry.FlowExportPolicy) (telemetry.FlowExportPolicy, error) {
	r, err := e.fnAutoDeleteFlowExportPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FlowExportPolicy), err
	}
	return telemetry.FlowExportPolicy{}, err

}
func (e *eFlowExportPolicyV1Endpoints) AutoGetFlowExportPolicy(ctx context.Context, t telemetry.FlowExportPolicy) (telemetry.FlowExportPolicy, error) {
	r, err := e.fnAutoGetFlowExportPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FlowExportPolicy), err
	}
	return telemetry.FlowExportPolicy{}, err

}
func (e *eFlowExportPolicyV1Endpoints) AutoListFlowExportPolicy(ctx context.Context, t api.ListWatchOptions) (telemetry.FlowExportPolicyList, error) {
	r, err := e.fnAutoListFlowExportPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FlowExportPolicyList), err
	}
	return telemetry.FlowExportPolicyList{}, err

}
func (e *eFlowExportPolicyV1Endpoints) AutoUpdateFlowExportPolicy(ctx context.Context, t telemetry.FlowExportPolicy) (telemetry.FlowExportPolicy, error) {
	r, err := e.fnAutoUpdateFlowExportPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FlowExportPolicy), err
	}
	return telemetry.FlowExportPolicy{}, err

}

func (e *eFlowExportPolicyV1Endpoints) AutoWatchFlowExportPolicy(in *api.ListWatchOptions, stream telemetry.FlowExportPolicyV1_AutoWatchFlowExportPolicyServer) error {
	return e.fnAutoWatchFlowExportPolicy(in, stream, "flowExportPolicy")
}
func (e *eFwlogPolicyV1Endpoints) AutoAddFwlogPolicy(ctx context.Context, t telemetry.FwlogPolicy) (telemetry.FwlogPolicy, error) {
	r, err := e.fnAutoAddFwlogPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FwlogPolicy), err
	}
	return telemetry.FwlogPolicy{}, err

}
func (e *eFwlogPolicyV1Endpoints) AutoDeleteFwlogPolicy(ctx context.Context, t telemetry.FwlogPolicy) (telemetry.FwlogPolicy, error) {
	r, err := e.fnAutoDeleteFwlogPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FwlogPolicy), err
	}
	return telemetry.FwlogPolicy{}, err

}
func (e *eFwlogPolicyV1Endpoints) AutoGetFwlogPolicy(ctx context.Context, t telemetry.FwlogPolicy) (telemetry.FwlogPolicy, error) {
	r, err := e.fnAutoGetFwlogPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FwlogPolicy), err
	}
	return telemetry.FwlogPolicy{}, err

}
func (e *eFwlogPolicyV1Endpoints) AutoListFwlogPolicy(ctx context.Context, t api.ListWatchOptions) (telemetry.FwlogPolicyList, error) {
	r, err := e.fnAutoListFwlogPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FwlogPolicyList), err
	}
	return telemetry.FwlogPolicyList{}, err

}
func (e *eFwlogPolicyV1Endpoints) AutoUpdateFwlogPolicy(ctx context.Context, t telemetry.FwlogPolicy) (telemetry.FwlogPolicy, error) {
	r, err := e.fnAutoUpdateFwlogPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.FwlogPolicy), err
	}
	return telemetry.FwlogPolicy{}, err

}

func (e *eFwlogPolicyV1Endpoints) AutoWatchFwlogPolicy(in *api.ListWatchOptions, stream telemetry.FwlogPolicyV1_AutoWatchFwlogPolicyServer) error {
	return e.fnAutoWatchFwlogPolicy(in, stream, "fwlogPolicy")
}
func (e *eStatsPolicyV1Endpoints) AutoAddStatsPolicy(ctx context.Context, t telemetry.StatsPolicy) (telemetry.StatsPolicy, error) {
	r, err := e.fnAutoAddStatsPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.StatsPolicy), err
	}
	return telemetry.StatsPolicy{}, err

}
func (e *eStatsPolicyV1Endpoints) AutoDeleteStatsPolicy(ctx context.Context, t telemetry.StatsPolicy) (telemetry.StatsPolicy, error) {
	r, err := e.fnAutoDeleteStatsPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.StatsPolicy), err
	}
	return telemetry.StatsPolicy{}, err

}
func (e *eStatsPolicyV1Endpoints) AutoGetStatsPolicy(ctx context.Context, t telemetry.StatsPolicy) (telemetry.StatsPolicy, error) {
	r, err := e.fnAutoGetStatsPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.StatsPolicy), err
	}
	return telemetry.StatsPolicy{}, err

}
func (e *eStatsPolicyV1Endpoints) AutoListStatsPolicy(ctx context.Context, t api.ListWatchOptions) (telemetry.StatsPolicyList, error) {
	r, err := e.fnAutoListStatsPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.StatsPolicyList), err
	}
	return telemetry.StatsPolicyList{}, err

}
func (e *eStatsPolicyV1Endpoints) AutoUpdateStatsPolicy(ctx context.Context, t telemetry.StatsPolicy) (telemetry.StatsPolicy, error) {
	r, err := e.fnAutoUpdateStatsPolicy(ctx, t)
	if err == nil {
		return r.(telemetry.StatsPolicy), err
	}
	return telemetry.StatsPolicy{}, err

}

func (e *eStatsPolicyV1Endpoints) AutoWatchStatsPolicy(in *api.ListWatchOptions, stream telemetry.StatsPolicyV1_AutoWatchStatsPolicyServer) error {
	return e.fnAutoWatchStatsPolicy(in, stream, "statsPolicy")
}

func init() {
	apisrv = apisrvpkg.MustGetAPIServer()

	svc := stelemetryTelemetryBackend{}

	{
		e := eFlowExportPolicyV1Endpoints{Svc: svc}
		svc.endpointsFlowExportPolicyV1 = &e
	}
	{
		e := eFwlogPolicyV1Endpoints{Svc: svc}
		svc.endpointsFwlogPolicyV1 = &e
	}
	{
		e := eStatsPolicyV1Endpoints{Svc: svc}
		svc.endpointsStatsPolicyV1 = &e
	}
	apisrv.Register("telemetry.protos/telemetry.proto", &svc)
}
