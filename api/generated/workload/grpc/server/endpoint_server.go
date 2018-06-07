// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package workloadApiServer is a auto generated package.
Input file: endpoint.proto
*/
package workloadApiServer

import (
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pensando/sw/api"
	workload "github.com/pensando/sw/api/generated/workload"
	"github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/apiserver/pkg"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/rpckit"
	"github.com/pensando/sw/venice/utils/runtime"
	"github.com/satori/go.uuid"
)

// dummy vars to suppress unused errors
var _ api.ObjectMeta
var _ listerwatcher.WatcherClient
var _ fmt.Stringer

type sworkloadEndpointBackend struct {
	Services map[string]apiserver.Service
	Messages map[string]apiserver.Message
	logger   log.Logger
	scheme   *runtime.Scheme
}

func (s *sworkloadEndpointBackend) regMsgsFunc(l log.Logger, scheme *runtime.Scheme) {
	l.Infof("registering message for sworkloadEndpointBackend")
	s.Messages = map[string]apiserver.Message{

		"workload.Endpoint": apisrvpkg.NewMessage("workload.Endpoint").WithKeyGenerator(func(i interface{}, prefix string) string {
			if i == nil {
				r := workload.Endpoint{}
				return r.MakeKey(prefix)
			}
			r := i.(workload.Endpoint)
			return r.MakeKey(prefix)
		}).WithObjectVersionWriter(func(i interface{}, version string) interface{} {
			r := i.(workload.Endpoint)
			r.APIVersion = version
			return r
		}).WithKvUpdater(func(ctx context.Context, kvs kvstore.Interface, i interface{}, prefix string, create, ignoreStatus bool) (interface{}, error) {
			r := i.(workload.Endpoint)
			key := r.MakeKey(prefix)
			r.Kind = "Endpoint"
			var err error
			if create {
				err = kvs.Create(ctx, key, &r)
				if err != nil {
					l.ErrorLog("msg", "KV create failed", "key", key, "error", err)
				}
			} else {
				if ignoreStatus {
					updateFunc := func(obj runtime.Object) (runtime.Object, error) {
						saved := obj.(*workload.Endpoint)
						if r.ResourceVersion != "" && r.ResourceVersion != saved.ResourceVersion {
							return nil, fmt.Errorf("Resource Version specified does not match Object version")
						}
						r.Status = saved.Status
						return &r, nil
					}
					into := &workload.Endpoint{}
					err = kvs.ConsistentUpdate(ctx, key, into, updateFunc)
				} else {
					if r.ResourceVersion != "" {
						l.Infof("resource version is specified %s\n", r.ResourceVersion)
						err = kvs.Update(ctx, key, &r, kvstore.Compare(kvstore.WithVersion(key), "=", r.ResourceVersion))
					} else {
						err = kvs.Update(ctx, key, &r)
					}
					if err != nil {
						l.ErrorLog("msg", "KV update failed", "key", key, "error", err)
					}
				}

			}
			return r, err
		}).WithKvTxnUpdater(func(ctx context.Context, txn kvstore.Txn, i interface{}, prefix string, create bool) error {
			r := i.(workload.Endpoint)
			key := r.MakeKey(prefix)
			var err error
			if create {
				err = txn.Create(key, &r)
				if err != nil {
					l.ErrorLog("msg", "KV transaction create failed", "key", key, "error", err)
				}
			} else {
				err = txn.Update(key, &r)
				if err != nil {
					l.ErrorLog("msg", "KV transaction update failed", "key", key, "error", err)
				}
			}
			return err
		}).WithUUIDWriter(func(i interface{}) (interface{}, error) {
			r := i.(workload.Endpoint)
			r.UUID = uuid.NewV4().String()
			return r, nil
		}).WithCreationTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(workload.Endpoint)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.CreationTime.Timestamp = *ts
			}
			return r, err
		}).WithModTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(workload.Endpoint)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.ModTime.Timestamp = *ts
			}
			return r, err
		}).WithSelfLinkWriter(func(path, ver, prefix string, i interface{}) (interface{}, error) {
			r := i.(workload.Endpoint)
			r.SelfLink = path
			return r, nil
		}).WithKvGetter(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := workload.Endpoint{}
			err := kvs.Get(ctx, key, &r)
			if err != nil {
				l.ErrorLog("msg", "Object get failed", "key", key, "error", err)
			}
			return r, err
		}).WithKvDelFunc(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := workload.Endpoint{}
			err := kvs.Delete(ctx, key, &r)
			if err != nil {
				l.ErrorLog("msg", "Object delete failed", "key", key, "error", err)
			}
			return r, err
		}).WithKvTxnDelFunc(func(ctx context.Context, txn kvstore.Txn, key string) error {
			err := txn.Delete(key)
			if err != nil {
				l.ErrorLog("msg", "Object Txn delete failed", "key", key, "error", err)
			}
			return err
		}).WithValidate(func(i interface{}, ver string, ignoreStatus bool) []error {
			r := i.(workload.Endpoint)
			return r.Validate(ver, "", ignoreStatus)
		}),

		"workload.EndpointSpec":   apisrvpkg.NewMessage("workload.EndpointSpec"),
		"workload.EndpointStatus": apisrvpkg.NewMessage("workload.EndpointStatus"),
		// Add a message handler for ListWatch options
		"api.ListWatchOptions": apisrvpkg.NewMessage("api.ListWatchOptions"),
	}

	apisrv.RegisterMessages("workload", s.Messages)
	// add messages to package.
	if pkgMessages == nil {
		pkgMessages = make(map[string]apiserver.Message)
	}
	for k, v := range s.Messages {
		pkgMessages[k] = v
	}
}

func (s *sworkloadEndpointBackend) regSvcsFunc(ctx context.Context, logger log.Logger, grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) {

}

func (s *sworkloadEndpointBackend) regWatchersFunc(ctx context.Context, logger log.Logger, grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) {

}

func (s *sworkloadEndpointBackend) CompleteRegistration(ctx context.Context, logger log.Logger,
	grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) error {
	// register all messages in the package if not done already
	s.logger = logger
	s.scheme = scheme
	registerMessages(logger, scheme)
	registerServices(ctx, logger, grpcserver, scheme)
	registerWatchers(ctx, logger, grpcserver, scheme)
	return nil
}

func (s *sworkloadEndpointBackend) Reset() {
	cleanupRegistration()
}

func init() {
	apisrv = apisrvpkg.MustGetAPIServer()

	svc := sworkloadEndpointBackend{}
	addMsgRegFunc(svc.regMsgsFunc)
	addSvcRegFunc(svc.regSvcsFunc)
	addWatcherRegFunc(svc.regWatchersFunc)

	apisrv.Register("workload.endpoint.proto", &svc)
}
