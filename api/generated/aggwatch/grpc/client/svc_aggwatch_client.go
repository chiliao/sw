// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

package grpcclient

import (
	"context"
	"errors"
	"net/http"

	"google.golang.org/grpc"

	api "github.com/pensando/sw/api"
	aggwatch "github.com/pensando/sw/api/generated/aggwatch"
	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
)

// Dummy vars to suppress import errors
var _ api.TypeMeta
var _ listerwatcher.WatcherClient
var _ kvstore.Interface

// NewAggWatchV1 sets up a new client for AggWatchV1
func NewAggWatchV1(conn *grpc.ClientConn, logger log.Logger) aggwatch.ServiceAggWatchV1Client {

	return aggwatch.EndpointsAggWatchV1Client{
		Client: aggwatch.NewAggWatchV1Client(conn),
	}
}

// NewAggWatchV1Backend creates an instrumented client with middleware
func NewAggWatchV1Backend(conn *grpc.ClientConn, logger log.Logger) aggwatch.ServiceAggWatchV1Client {
	cl := NewAggWatchV1(conn, logger)
	cl = aggwatch.LoggingAggWatchV1MiddlewareClient(logger)(cl)
	return cl
}

type crudClientAggWatchV1 struct {
	logger log.Logger
	client aggwatch.ServiceAggWatchV1Client
}

// NewGrpcCrudClientAggWatchV1 creates a GRPC client for the service
func NewGrpcCrudClientAggWatchV1(conn *grpc.ClientConn, logger log.Logger) aggwatch.AggWatchV1Interface {
	client := NewAggWatchV1Backend(conn, logger)
	return &crudClientAggWatchV1{
		logger: logger,
		client: client,
	}
}

func (a *crudClientAggWatchV1) Watch(ctx context.Context, options *api.AggWatchOptions) (kvstore.Watcher, error) {
	a.logger.DebugLog("msg", "received call", "object", "AggWatchV1", "oper", "WatchOper")
	nctx := addVersion(ctx, "v1")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	stream, err := a.client.AutoWatchSvcAggWatchV1(nctx, options)
	if err != nil {
		return nil, err
	}
	wstream := stream.(aggwatch.AggWatchV1_AutoWatchSvcAggWatchV1Client)
	bridgefn := func(lw *listerwatcher.WatcherClient) {
		for {
			r, err := wstream.Recv()
			if err != nil {
				a.logger.ErrorLog("msg", "error on receive", "err", err)
				close(lw.OutCh)
				return
			}
			for _, e := range r.Events {
				ev := kvstore.WatchEvent{Type: kvstore.WatchEventType(e.Type)}
				switch e.Type {
				case string(kvstore.Created), string(kvstore.Updated), string(kvstore.Deleted):
					robj, err := listerwatcher.GetObject(e)
					if err != nil {
						a.logger.ErrorLog("msg", "error on receive unmarshall", "err", err)
						close(lw.OutCh)
						return
					}
					ev.Object = robj
				case string(kvstore.WatcherControl):
					ev.Control = &kvstore.WatchControl{
						Code:    e.Control.Code,
						Message: e.Control.Message,
					}
				}
				select {
				case lw.OutCh <- &ev:
				case <-wstream.Context().Done():
					close(lw.OutCh)
					return
				}
			}
		}
	}
	lw := listerwatcher.NewWatcherClient(wstream, bridgefn)
	lw.Run()
	return lw, nil
}

type crudRestClientAggWatchV1 struct {
}

// NewRestCrudClientAggWatchV1 creates a REST client for the service.
func NewRestCrudClientAggWatchV1(url string, httpClient *http.Client) aggwatch.AggWatchV1Interface {
	return &crudRestClientAggWatchV1{}
}

// NewStagedRestCrudClientAggWatchV1 creates a REST client for the service.
func NewStagedRestCrudClientAggWatchV1(url string, id string, httpClient *http.Client) aggwatch.AggWatchV1Interface {
	return &crudRestClientAggWatchV1{}
}

func (a *crudRestClientAggWatchV1) Watch(ctx context.Context, options *api.AggWatchOptions) (kvstore.Watcher, error) {
	return nil, errors.New("method unimplemented")
}
