// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package ctkit is a auto generated package.
Input file: svc_workload.proto
*/
package ctkit

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/api/generated/apiclient"
	"github.com/pensando/sw/api/generated/workload"
	"github.com/pensando/sw/venice/utils/balancer"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// Endpoint is a wrapper object that implements additional functionality
type Endpoint struct {
	sync.Mutex
	workload.Endpoint
	HandlerCtx interface{} // additional state handlers can store
	ctrler     *ctrlerCtx  // reference back to the controller instance
}

func (obj *Endpoint) Write() error {
	// if there is no API server to connect to, we are done
	if (obj.ctrler == nil) || (obj.ctrler.resolver == nil) || obj.ctrler.apisrvURL == "" {
		return nil
	}

	apicl, err := obj.ctrler.apiClient()
	if err != nil {
		obj.ctrler.logger.Errorf("Error creating API server clent. Err: %v", err)
		return err
	}

	obj.ctrler.stats.Counter("Endpoint_Writes").Inc()

	// write to api server
	if obj.ObjectMeta.ResourceVersion != "" {
		nobj := obj.Endpoint
		// FIXME: clear the resource version till we figure out CAS semantics
		// update it
		_, err = apicl.WorkloadV1().Endpoint().Update(context.Background(), &nobj)
	} else {
		//  create
		_, err = apicl.WorkloadV1().Endpoint().Create(context.Background(), &obj.Endpoint)
	}

	return nil
}

// EndpointHandler is the event handler for Endpoint object
type EndpointHandler interface {
	OnEndpointCreate(obj *Endpoint) error
	OnEndpointUpdate(obj *Endpoint) error
	OnEndpointDelete(obj *Endpoint) error
}

// handleEndpointEvent handles Endpoint events from watcher
func (ct *ctrlerCtx) handleEndpointEvent(evt *kvstore.WatchEvent) error {
	switch tp := evt.Object.(type) {
	case *workload.Endpoint:
		eobj := evt.Object.(*workload.Endpoint)
		kind := "Endpoint"

		ct.logger.Infof("Watcher: Got %s watch event(%s): {%+v}", kind, evt.Type, eobj)

		handler, ok := ct.handlers[kind]
		if !ok {
			ct.logger.Fatalf("Cant find the handler for %s", kind)
		}
		endpointHandler := handler.(EndpointHandler)
		// handle based on event type
		switch evt.Type {
		case kvstore.Created:
			fallthrough
		case kvstore.Updated:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				obj := &Endpoint{
					Endpoint:   *eobj,
					HandlerCtx: nil,
					ctrler:     ct,
				}
				ct.addObject(kind, obj.GetKey(), obj)
				ct.stats.Counter("Endpoint_Created_Events").Inc()

				// call the event handler
				obj.Lock()
				err = endpointHandler.OnEndpointCreate(obj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					ct.delObject(kind, eobj.GetKey())
					return err
				}
			} else {
				obj := fobj.(*Endpoint)

				// see if it changed
				_, ok := ref.ObjDiff(obj.Spec, eobj.Spec)
				if ok || obj.ObjectMeta.GenerationID != eobj.ObjectMeta.GenerationID {
					obj.ObjectMeta = eobj.ObjectMeta
					obj.Spec = eobj.Spec

					ct.stats.Counter("Endpoint_Updated_Events").Inc()

					// call the event handler
					obj.Lock()
					err = endpointHandler.OnEndpointUpdate(obj)
					obj.Unlock()
					if err != nil {
						ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
						return err
					}
				}
			}
		case kvstore.Deleted:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				ct.logger.Errorf("Object %s/%s not found durng delete. Err: %v", kind, eobj.GetKey(), err)
				return err
			}

			obj := fobj.(*Endpoint)

			ct.stats.Counter("Endpoint_Deleted_Events").Inc()

			// Call the event reactor
			obj.Lock()
			err = endpointHandler.OnEndpointDelete(obj)
			obj.Unlock()
			if err != nil {
				ct.logger.Errorf("Error deleting %s: %+v. Err: %v", kind, obj, err)
			}

			ct.delObject(kind, eobj.GetKey())
		}
	default:
		ct.logger.Fatalf("API watcher Found object of invalid type: %v on Endpoint watch channel", tp)
	}

	return nil
}

// diffEndpoint does a diff of Endpoint objects between local cache and API server
func (ct *ctrlerCtx) diffEndpoint(apicl apiclient.Services) {
	opts := api.ListWatchOptions{}

	// get a list of all objects from API server
	objlist, err := apicl.WorkloadV1().Endpoint().List(context.Background(), &opts)
	if err != nil {
		ct.logger.Errorf("Error getting a list of objects. Err: %v", err)
		return
	}

	// build an object map
	objmap := make(map[string]*workload.Endpoint)
	for _, obj := range objlist {
		objmap[obj.GetKey()] = obj
	}

	// if an object is in our local cache and not in API server, trigger delete for it
	for _, obj := range ct.Endpoint().List() {
		_, ok := objmap[obj.GetKey()]
		if !ok {
			evt := kvstore.WatchEvent{
				Type:   kvstore.Deleted,
				Key:    obj.GetKey(),
				Object: &obj.Endpoint,
			}
			ct.handleEndpointEvent(&evt)
		}
	}

	// trigger create event for all others
	for _, obj := range objlist {
		evt := kvstore.WatchEvent{
			Type:   kvstore.Created,
			Key:    obj.GetKey(),
			Object: obj,
		}
		ct.handleEndpointEvent(&evt)
	}
}

func (ct *ctrlerCtx) runEndpointWatcher() {
	kind := "Endpoint"

	// if there is no API server to connect to, we are done
	if (ct.resolver == nil) || ct.apisrvURL == "" {
		return
	}

	// create context
	ctx, cancel := context.WithCancel(context.Background())
	ct.Lock()
	ct.watchCancel[kind] = cancel
	ct.Unlock()
	opts := api.ListWatchOptions{}

	// setup wait group
	ct.waitGrp.Add(1)
	defer ct.waitGrp.Done()
	logger := ct.logger.WithContext("submodule", "EndpointWatcher")

	ct.stats.Counter("Endpoint_Watch").Inc()
	defer ct.stats.Counter("Endpoint_Watch").Dec()

	// loop forever
	for {
		// create a grpc client
		apicl, err := apiclient.NewGrpcAPIClient(ct.name, ct.apisrvURL, logger, rpckit.WithBalancer(balancer.New(ct.resolver)))
		if err != nil {
			logger.Warnf("Failed to connect to gRPC server [%s]\n", ct.apisrvURL)
			ct.stats.Counter("Endpoint_ApiClientErr").Inc()
		} else {
			logger.Infof("API client connected {%+v}", apicl)

			// Endpoint object watcher
			wt, werr := apicl.WorkloadV1().Endpoint().Watch(ctx, &opts)
			if werr != nil {
				logger.Errorf("Failed to start %s watch (%s)\n", kind, werr)
				// wait for a second and retry connecting to api server
				apicl.Close()
				time.Sleep(time.Second)
				continue
			}
			ct.Lock()
			ct.watchers[kind] = wt
			ct.Unlock()

			// perform a diff with API server and local cache
			time.Sleep(time.Millisecond * 100)
			ct.diffEndpoint(apicl)

			// handle api server watch events
		innerLoop:
			for {
				// wait for events
				select {
				case evt, ok := <-wt.EventChan():
					if !ok {
						logger.Error("Error receiving from apisrv watcher")
						ct.stats.Counter("Endpoint_WatchErrors").Inc()
						break innerLoop
					}

					// handle event
					ct.handleEndpointEvent(evt)
				}
			}
			apicl.Close()
		}

		// if stop flag is set, we are done
		if ct.stoped {
			logger.Infof("Exiting API server watcher")
			return
		}

		// wait for a second and retry connecting to api server
		time.Sleep(time.Second)
	}
}

// WatchEndpoint starts watch on Endpoint object
func (ct *ctrlerCtx) WatchEndpoint(handler EndpointHandler) error {
	kind := "Endpoint"

	ct.Lock()
	defer ct.Unlock()

	// see if we already have a watcher
	_, ok := ct.watchers[kind]
	if ok {
		return fmt.Errorf("Endpoint watcher already exists")
	}

	// save handler
	ct.handlers[kind] = handler

	// run Endpoint watcher in a go routine
	go ct.runEndpointWatcher()

	return nil
}

// EndpointAPI returns
type EndpointAPI interface {
	Create(obj *workload.Endpoint) error
	Update(obj *workload.Endpoint) error
	Delete(obj *workload.Endpoint) error
	List() []*Endpoint
	Watch(handler EndpointHandler) error
}

// dummy struct that implements EndpointAPI
type endpointAPI struct {
	ct *ctrlerCtx
}

// Create creates Endpoint object
func (api *endpointAPI) Create(obj *workload.Endpoint) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.WorkloadV1().Endpoint().Create(context.Background(), obj)
		if err != nil && strings.Contains(err.Error(), "AlreadyExists") {
			_, err = apicl.WorkloadV1().Endpoint().Update(context.Background(), obj)
		}
		if err != nil {
			return err
		}
	}

	return api.ct.handleEndpointEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Created})
}

// Update triggers update on Endpoint object
func (api *endpointAPI) Update(obj *workload.Endpoint) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.WorkloadV1().Endpoint().Update(context.Background(), obj)
		if err != nil {
			return err
		}
	}

	return api.ct.handleEndpointEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Updated})
}

// Delete deletes Endpoint object
func (api *endpointAPI) Delete(obj *workload.Endpoint) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		apicl.WorkloadV1().Endpoint().Delete(context.Background(), &obj.ObjectMeta)
	}

	return api.ct.handleEndpointEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Deleted})
}

// List returns a list of all Endpoint objects
func (api *endpointAPI) List() []*Endpoint {
	var objlist []*Endpoint

	objs := api.ct.ListObjects("Endpoint")
	for _, obj := range objs {
		switch tp := obj.(type) {
		case *Endpoint:
			eobj := obj.(*Endpoint)
			objlist = append(objlist, eobj)
		default:
			log.Fatalf("Got invalid object type %v while looking for Endpoint", tp)
		}
	}

	return objlist
}

// Watch sets up a event handlers for Endpoint object
func (api *endpointAPI) Watch(handler EndpointHandler) error {
	return api.ct.WatchEndpoint(handler)
}

// Endpoint returns EndpointAPI
func (ct *ctrlerCtx) Endpoint() EndpointAPI {
	return &endpointAPI{ct: ct}
}

// Workload is a wrapper object that implements additional functionality
type Workload struct {
	sync.Mutex
	workload.Workload
	HandlerCtx interface{} // additional state handlers can store
	ctrler     *ctrlerCtx  // reference back to the controller instance
}

func (obj *Workload) Write() error {
	// if there is no API server to connect to, we are done
	if (obj.ctrler == nil) || (obj.ctrler.resolver == nil) || obj.ctrler.apisrvURL == "" {
		return nil
	}

	apicl, err := obj.ctrler.apiClient()
	if err != nil {
		obj.ctrler.logger.Errorf("Error creating API server clent. Err: %v", err)
		return err
	}

	obj.ctrler.stats.Counter("Workload_Writes").Inc()

	// write to api server
	if obj.ObjectMeta.ResourceVersion != "" {
		nobj := obj.Workload
		// FIXME: clear the resource version till we figure out CAS semantics
		// update it
		_, err = apicl.WorkloadV1().Workload().Update(context.Background(), &nobj)
	} else {
		//  create
		_, err = apicl.WorkloadV1().Workload().Create(context.Background(), &obj.Workload)
	}

	return nil
}

// WorkloadHandler is the event handler for Workload object
type WorkloadHandler interface {
	OnWorkloadCreate(obj *Workload) error
	OnWorkloadUpdate(obj *Workload) error
	OnWorkloadDelete(obj *Workload) error
}

// handleWorkloadEvent handles Workload events from watcher
func (ct *ctrlerCtx) handleWorkloadEvent(evt *kvstore.WatchEvent) error {
	switch tp := evt.Object.(type) {
	case *workload.Workload:
		eobj := evt.Object.(*workload.Workload)
		kind := "Workload"

		ct.logger.Infof("Watcher: Got %s watch event(%s): {%+v}", kind, evt.Type, eobj)

		handler, ok := ct.handlers[kind]
		if !ok {
			ct.logger.Fatalf("Cant find the handler for %s", kind)
		}
		workloadHandler := handler.(WorkloadHandler)
		// handle based on event type
		switch evt.Type {
		case kvstore.Created:
			fallthrough
		case kvstore.Updated:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				obj := &Workload{
					Workload:   *eobj,
					HandlerCtx: nil,
					ctrler:     ct,
				}
				ct.addObject(kind, obj.GetKey(), obj)
				ct.stats.Counter("Workload_Created_Events").Inc()

				// call the event handler
				obj.Lock()
				err = workloadHandler.OnWorkloadCreate(obj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					ct.delObject(kind, eobj.GetKey())
					return err
				}
			} else {
				obj := fobj.(*Workload)

				// see if it changed
				_, ok := ref.ObjDiff(obj.Spec, eobj.Spec)
				if ok || obj.ObjectMeta.GenerationID != eobj.ObjectMeta.GenerationID {
					obj.ObjectMeta = eobj.ObjectMeta
					obj.Spec = eobj.Spec

					ct.stats.Counter("Workload_Updated_Events").Inc()

					// call the event handler
					obj.Lock()
					err = workloadHandler.OnWorkloadUpdate(obj)
					obj.Unlock()
					if err != nil {
						ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
						return err
					}
				}
			}
		case kvstore.Deleted:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				ct.logger.Errorf("Object %s/%s not found durng delete. Err: %v", kind, eobj.GetKey(), err)
				return err
			}

			obj := fobj.(*Workload)

			ct.stats.Counter("Workload_Deleted_Events").Inc()

			// Call the event reactor
			obj.Lock()
			err = workloadHandler.OnWorkloadDelete(obj)
			obj.Unlock()
			if err != nil {
				ct.logger.Errorf("Error deleting %s: %+v. Err: %v", kind, obj, err)
			}

			ct.delObject(kind, eobj.GetKey())
		}
	default:
		ct.logger.Fatalf("API watcher Found object of invalid type: %v on Workload watch channel", tp)
	}

	return nil
}

// diffWorkload does a diff of Workload objects between local cache and API server
func (ct *ctrlerCtx) diffWorkload(apicl apiclient.Services) {
	opts := api.ListWatchOptions{}

	// get a list of all objects from API server
	objlist, err := apicl.WorkloadV1().Workload().List(context.Background(), &opts)
	if err != nil {
		ct.logger.Errorf("Error getting a list of objects. Err: %v", err)
		return
	}

	// build an object map
	objmap := make(map[string]*workload.Workload)
	for _, obj := range objlist {
		objmap[obj.GetKey()] = obj
	}

	// if an object is in our local cache and not in API server, trigger delete for it
	for _, obj := range ct.Workload().List() {
		_, ok := objmap[obj.GetKey()]
		if !ok {
			evt := kvstore.WatchEvent{
				Type:   kvstore.Deleted,
				Key:    obj.GetKey(),
				Object: &obj.Workload,
			}
			ct.handleWorkloadEvent(&evt)
		}
	}

	// trigger create event for all others
	for _, obj := range objlist {
		evt := kvstore.WatchEvent{
			Type:   kvstore.Created,
			Key:    obj.GetKey(),
			Object: obj,
		}
		ct.handleWorkloadEvent(&evt)
	}
}

func (ct *ctrlerCtx) runWorkloadWatcher() {
	kind := "Workload"

	// if there is no API server to connect to, we are done
	if (ct.resolver == nil) || ct.apisrvURL == "" {
		return
	}

	// create context
	ctx, cancel := context.WithCancel(context.Background())
	ct.Lock()
	ct.watchCancel[kind] = cancel
	ct.Unlock()
	opts := api.ListWatchOptions{}

	// setup wait group
	ct.waitGrp.Add(1)
	defer ct.waitGrp.Done()
	logger := ct.logger.WithContext("submodule", "WorkloadWatcher")

	ct.stats.Counter("Workload_Watch").Inc()
	defer ct.stats.Counter("Workload_Watch").Dec()

	// loop forever
	for {
		// create a grpc client
		apicl, err := apiclient.NewGrpcAPIClient(ct.name, ct.apisrvURL, logger, rpckit.WithBalancer(balancer.New(ct.resolver)))
		if err != nil {
			logger.Warnf("Failed to connect to gRPC server [%s]\n", ct.apisrvURL)
			ct.stats.Counter("Workload_ApiClientErr").Inc()
		} else {
			logger.Infof("API client connected {%+v}", apicl)

			// Workload object watcher
			wt, werr := apicl.WorkloadV1().Workload().Watch(ctx, &opts)
			if werr != nil {
				logger.Errorf("Failed to start %s watch (%s)\n", kind, werr)
				// wait for a second and retry connecting to api server
				apicl.Close()
				time.Sleep(time.Second)
				continue
			}
			ct.Lock()
			ct.watchers[kind] = wt
			ct.Unlock()

			// perform a diff with API server and local cache
			time.Sleep(time.Millisecond * 100)
			ct.diffWorkload(apicl)

			// handle api server watch events
		innerLoop:
			for {
				// wait for events
				select {
				case evt, ok := <-wt.EventChan():
					if !ok {
						logger.Error("Error receiving from apisrv watcher")
						ct.stats.Counter("Workload_WatchErrors").Inc()
						break innerLoop
					}

					// handle event
					ct.handleWorkloadEvent(evt)
				}
			}
			apicl.Close()
		}

		// if stop flag is set, we are done
		if ct.stoped {
			logger.Infof("Exiting API server watcher")
			return
		}

		// wait for a second and retry connecting to api server
		time.Sleep(time.Second)
	}
}

// WatchWorkload starts watch on Workload object
func (ct *ctrlerCtx) WatchWorkload(handler WorkloadHandler) error {
	kind := "Workload"

	ct.Lock()
	defer ct.Unlock()

	// see if we already have a watcher
	_, ok := ct.watchers[kind]
	if ok {
		return fmt.Errorf("Workload watcher already exists")
	}

	// save handler
	ct.handlers[kind] = handler

	// run Workload watcher in a go routine
	go ct.runWorkloadWatcher()

	return nil
}

// WorkloadAPI returns
type WorkloadAPI interface {
	Create(obj *workload.Workload) error
	Update(obj *workload.Workload) error
	Delete(obj *workload.Workload) error
	List() []*Workload
	Watch(handler WorkloadHandler) error
}

// dummy struct that implements WorkloadAPI
type workloadAPI struct {
	ct *ctrlerCtx
}

// Create creates Workload object
func (api *workloadAPI) Create(obj *workload.Workload) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.WorkloadV1().Workload().Create(context.Background(), obj)
		if err != nil && strings.Contains(err.Error(), "AlreadyExists") {
			_, err = apicl.WorkloadV1().Workload().Update(context.Background(), obj)
		}
		if err != nil {
			return err
		}
	}

	return api.ct.handleWorkloadEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Created})
}

// Update triggers update on Workload object
func (api *workloadAPI) Update(obj *workload.Workload) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.WorkloadV1().Workload().Update(context.Background(), obj)
		if err != nil {
			return err
		}
	}

	return api.ct.handleWorkloadEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Updated})
}

// Delete deletes Workload object
func (api *workloadAPI) Delete(obj *workload.Workload) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		apicl.WorkloadV1().Workload().Delete(context.Background(), &obj.ObjectMeta)
	}

	return api.ct.handleWorkloadEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Deleted})
}

// List returns a list of all Workload objects
func (api *workloadAPI) List() []*Workload {
	var objlist []*Workload

	objs := api.ct.ListObjects("Workload")
	for _, obj := range objs {
		switch tp := obj.(type) {
		case *Workload:
			eobj := obj.(*Workload)
			objlist = append(objlist, eobj)
		default:
			log.Fatalf("Got invalid object type %v while looking for Workload", tp)
		}
	}

	return objlist
}

// Watch sets up a event handlers for Workload object
func (api *workloadAPI) Watch(handler WorkloadHandler) error {
	return api.ct.WatchWorkload(handler)
}

// Workload returns WorkloadAPI
func (ct *ctrlerCtx) Workload() WorkloadAPI {
	return &workloadAPI{ct: ct}
}
