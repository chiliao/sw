// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package securityApiServer is a auto generated package.
Input file: app.proto
*/
package securityApiServer

import (
	"context"
	"fmt"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pensando/sw/api"
	security "github.com/pensando/sw/api/generated/security"
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

type ssecurityAppBackend struct {
	Services map[string]apiserver.Service
	Messages map[string]apiserver.Message
	logger   log.Logger
	scheme   *runtime.Scheme
}

func (s *ssecurityAppBackend) regMsgsFunc(l log.Logger, scheme *runtime.Scheme) {
	l.Infof("registering message for ssecurityAppBackend")
	s.Messages = map[string]apiserver.Message{

		"security.App": apisrvpkg.NewMessage("security.App").WithKeyGenerator(func(i interface{}, prefix string) string {
			if i == nil {
				r := security.App{}
				return r.MakeKey(prefix)
			}
			r := i.(security.App)
			return r.MakeKey(prefix)
		}).WithObjectVersionWriter(func(i interface{}, version string) interface{} {
			r := i.(security.App)
			r.APIVersion = version
			return r
		}).WithKvUpdater(func(ctx context.Context, kvs kvstore.Interface, i interface{}, prefix string, create, ignoreStatus bool) (interface{}, error) {
			r := i.(security.App)
			key := r.MakeKey(prefix)
			r.Kind = "App"
			var err error
			if create {
				err = kvs.Create(ctx, key, &r)
				if err != nil {
					l.ErrorLog("msg", "KV create failed", "key", key, "error", err)
				}
			} else {
				if ignoreStatus {
					updateFunc := func(obj runtime.Object) (runtime.Object, error) {
						saved := obj.(*security.App)
						if r.ResourceVersion != "" && r.ResourceVersion != saved.ResourceVersion {
							return nil, fmt.Errorf("Resource Version specified does not match Object version")
						}
						r.Status = saved.Status
						return &r, nil
					}
					into := &security.App{}
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
			r := i.(security.App)
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
			r := i.(security.App)
			r.UUID = uuid.NewV4().String()
			return r, nil
		}).WithCreationTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(security.App)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.CreationTime.Timestamp = *ts
			}
			return r, err
		}).WithModTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(security.App)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.ModTime.Timestamp = *ts
			}
			return r, err
		}).WithSelfLinkWriter(func(path, ver, prefix string, i interface{}) (interface{}, error) {
			r := i.(security.App)
			r.SelfLink = path
			return r, nil
		}).WithKvGetter(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := security.App{}
			err := kvs.Get(ctx, key, &r)
			if err != nil {
				l.ErrorLog("msg", "Object get failed", "key", key, "error", err)
			}
			return r, err
		}).WithKvDelFunc(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := security.App{}
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
			r := i.(security.App)
			return r.Validate(ver, "", ignoreStatus)
		}),

		"security.AppSpec":   apisrvpkg.NewMessage("security.AppSpec"),
		"security.AppStatus": apisrvpkg.NewMessage("security.AppStatus"),
		"security.AppUser": apisrvpkg.NewMessage("security.AppUser").WithKeyGenerator(func(i interface{}, prefix string) string {
			if i == nil {
				r := security.AppUser{}
				return r.MakeKey(prefix)
			}
			r := i.(security.AppUser)
			return r.MakeKey(prefix)
		}).WithObjectVersionWriter(func(i interface{}, version string) interface{} {
			r := i.(security.AppUser)
			r.APIVersion = version
			return r
		}).WithKvUpdater(func(ctx context.Context, kvs kvstore.Interface, i interface{}, prefix string, create, ignoreStatus bool) (interface{}, error) {
			r := i.(security.AppUser)
			key := r.MakeKey(prefix)
			r.Kind = "AppUser"
			var err error
			if create {
				err = kvs.Create(ctx, key, &r)
				if err != nil {
					l.ErrorLog("msg", "KV create failed", "key", key, "error", err)
				}
			} else {
				if ignoreStatus {
					updateFunc := func(obj runtime.Object) (runtime.Object, error) {
						saved := obj.(*security.AppUser)
						if r.ResourceVersion != "" && r.ResourceVersion != saved.ResourceVersion {
							return nil, fmt.Errorf("Resource Version specified does not match Object version")
						}
						r.Status = saved.Status
						return &r, nil
					}
					into := &security.AppUser{}
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
			r := i.(security.AppUser)
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
			r := i.(security.AppUser)
			r.UUID = uuid.NewV4().String()
			return r, nil
		}).WithCreationTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(security.AppUser)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.CreationTime.Timestamp = *ts
			}
			return r, err
		}).WithModTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(security.AppUser)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.ModTime.Timestamp = *ts
			}
			return r, err
		}).WithSelfLinkWriter(func(path, ver, prefix string, i interface{}) (interface{}, error) {
			r := i.(security.AppUser)
			r.SelfLink = path
			return r, nil
		}).WithKvGetter(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := security.AppUser{}
			err := kvs.Get(ctx, key, &r)
			if err != nil {
				l.ErrorLog("msg", "Object get failed", "key", key, "error", err)
			}
			return r, err
		}).WithKvDelFunc(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := security.AppUser{}
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
			r := i.(security.AppUser)
			return r.Validate(ver, "", ignoreStatus)
		}),

		"security.AppUserGrp": apisrvpkg.NewMessage("security.AppUserGrp").WithKeyGenerator(func(i interface{}, prefix string) string {
			if i == nil {
				r := security.AppUserGrp{}
				return r.MakeKey(prefix)
			}
			r := i.(security.AppUserGrp)
			return r.MakeKey(prefix)
		}).WithObjectVersionWriter(func(i interface{}, version string) interface{} {
			r := i.(security.AppUserGrp)
			r.APIVersion = version
			return r
		}).WithKvUpdater(func(ctx context.Context, kvs kvstore.Interface, i interface{}, prefix string, create, ignoreStatus bool) (interface{}, error) {
			r := i.(security.AppUserGrp)
			key := r.MakeKey(prefix)
			r.Kind = "AppUserGrp"
			var err error
			if create {
				err = kvs.Create(ctx, key, &r)
				if err != nil {
					l.ErrorLog("msg", "KV create failed", "key", key, "error", err)
				}
			} else {
				if ignoreStatus {
					updateFunc := func(obj runtime.Object) (runtime.Object, error) {
						saved := obj.(*security.AppUserGrp)
						if r.ResourceVersion != "" && r.ResourceVersion != saved.ResourceVersion {
							return nil, fmt.Errorf("Resource Version specified does not match Object version")
						}
						r.Status = saved.Status
						return &r, nil
					}
					into := &security.AppUserGrp{}
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
			r := i.(security.AppUserGrp)
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
			r := i.(security.AppUserGrp)
			r.UUID = uuid.NewV4().String()
			return r, nil
		}).WithCreationTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(security.AppUserGrp)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.CreationTime.Timestamp = *ts
			}
			return r, err
		}).WithModTimeWriter(func(i interface{}) (interface{}, error) {
			r := i.(security.AppUserGrp)
			var err error
			ts, err := types.TimestampProto(time.Now())
			if err == nil {
				r.ModTime.Timestamp = *ts
			}
			return r, err
		}).WithSelfLinkWriter(func(path, ver, prefix string, i interface{}) (interface{}, error) {
			r := i.(security.AppUserGrp)
			r.SelfLink = path
			return r, nil
		}).WithKvGetter(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := security.AppUserGrp{}
			err := kvs.Get(ctx, key, &r)
			if err != nil {
				l.ErrorLog("msg", "Object get failed", "key", key, "error", err)
			}
			return r, err
		}).WithKvDelFunc(func(ctx context.Context, kvs kvstore.Interface, key string) (interface{}, error) {
			r := security.AppUserGrp{}
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
			r := i.(security.AppUserGrp)
			return r.Validate(ver, "", ignoreStatus)
		}),

		"security.AppUserGrpSpec":   apisrvpkg.NewMessage("security.AppUserGrpSpec"),
		"security.AppUserGrpStatus": apisrvpkg.NewMessage("security.AppUserGrpStatus"),
		"security.AppUserSpec":      apisrvpkg.NewMessage("security.AppUserSpec"),
		"security.AppUserStatus":    apisrvpkg.NewMessage("security.AppUserStatus"),
		// Add a message handler for ListWatch options
		"api.ListWatchOptions": apisrvpkg.NewMessage("api.ListWatchOptions"),
	}

	apisrv.RegisterMessages("security", s.Messages)
	// add messages to package.
	if pkgMessages == nil {
		pkgMessages = make(map[string]apiserver.Message)
	}
	for k, v := range s.Messages {
		pkgMessages[k] = v
	}
}

func (s *ssecurityAppBackend) regSvcsFunc(ctx context.Context, logger log.Logger, grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) {

}

func (s *ssecurityAppBackend) regWatchersFunc(ctx context.Context, logger log.Logger, grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) {

}

func (s *ssecurityAppBackend) CompleteRegistration(ctx context.Context, logger log.Logger,
	grpcserver *rpckit.RPCServer, scheme *runtime.Scheme) error {
	// register all messages in the package if not done already
	s.logger = logger
	s.scheme = scheme
	registerMessages(logger, scheme)
	registerServices(ctx, logger, grpcserver, scheme)
	registerWatchers(ctx, logger, grpcserver, scheme)
	return nil
}

func (s *ssecurityAppBackend) Reset() {
	cleanupRegistration()
}

func init() {
	apisrv = apisrvpkg.MustGetAPIServer()

	svc := ssecurityAppBackend{}
	addMsgRegFunc(svc.regMsgsFunc)
	addSvcRegFunc(svc.regSvcsFunc)
	addWatcherRegFunc(svc.regWatchersFunc)

	apisrv.Register("security.app.proto", &svc)
}
