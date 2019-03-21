// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package ctkit is a auto generated package.
Input file: svc_staging.proto
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
	"github.com/pensando/sw/api/generated/staging"
	"github.com/pensando/sw/venice/utils/balancer"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/ref"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// Buffer is a wrapper object that implements additional functionality
type Buffer struct {
	sync.Mutex
	staging.Buffer
	HandlerCtx interface{} // additional state handlers can store
	ctrler     *ctrlerCtx  // reference back to the controller instance
}

func (obj *Buffer) Write() error {
	// if there is no API server to connect to, we are done
	if (obj.ctrler == nil) || (obj.ctrler.resolver == nil) || obj.ctrler.apisrvURL == "" {
		return nil
	}

	apicl, err := obj.ctrler.apiClient()
	if err != nil {
		obj.ctrler.logger.Errorf("Error creating API server clent. Err: %v", err)
		return err
	}

	obj.ctrler.stats.Counter("Buffer_Writes").Inc()

	// write to api server
	if obj.ObjectMeta.ResourceVersion != "" {
		nobj := obj.Buffer
		// FIXME: clear the resource version till we figure out CAS semantics
		// update it
		_, err = apicl.StagingV1().Buffer().Update(context.Background(), &nobj)
	} else {
		//  create
		_, err = apicl.StagingV1().Buffer().Create(context.Background(), &obj.Buffer)
	}

	return nil
}

// BufferHandler is the event handler for Buffer object
type BufferHandler interface {
	OnBufferCreate(obj *Buffer) error
	OnBufferUpdate(obj *Buffer) error
	OnBufferDelete(obj *Buffer) error
}

// handleBufferEvent handles Buffer events from watcher
func (ct *ctrlerCtx) handleBufferEvent(evt *kvstore.WatchEvent) error {
	switch tp := evt.Object.(type) {
	case *staging.Buffer:
		eobj := evt.Object.(*staging.Buffer)
		kind := "Buffer"

		ct.logger.Infof("Watcher: Got %s watch event(%s): {%+v}", kind, evt.Type, eobj)

		handler, ok := ct.handlers[kind]
		if !ok {
			ct.logger.Fatalf("Cant find the handler for %s", kind)
		}
		bufferHandler := handler.(BufferHandler)
		// handle based on event type
		switch evt.Type {
		case kvstore.Created:
			fallthrough
		case kvstore.Updated:
			fobj, err := ct.findObject(kind, eobj.GetKey())
			if err != nil {
				obj := &Buffer{
					Buffer:     *eobj,
					HandlerCtx: nil,
					ctrler:     ct,
				}
				ct.addObject(kind, obj.GetKey(), obj)
				ct.stats.Counter("Buffer_Created_Events").Inc()

				// call the event handler
				obj.Lock()
				err = bufferHandler.OnBufferCreate(obj)
				obj.Unlock()
				if err != nil {
					ct.logger.Errorf("Error creating %s %+v. Err: %v", kind, obj, err)
					ct.delObject(kind, eobj.GetKey())
					return err
				}
			} else {
				obj := fobj.(*Buffer)

				// see if it changed
				_, ok := ref.ObjDiff(obj.Spec, eobj.Spec)
				if ok || obj.ObjectMeta.GenerationID != eobj.ObjectMeta.GenerationID {
					obj.ObjectMeta = eobj.ObjectMeta
					obj.Spec = eobj.Spec

					ct.stats.Counter("Buffer_Updated_Events").Inc()

					// call the event handler
					obj.Lock()
					err = bufferHandler.OnBufferUpdate(obj)
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

			obj := fobj.(*Buffer)

			ct.stats.Counter("Buffer_Deleted_Events").Inc()

			// Call the event reactor
			obj.Lock()
			err = bufferHandler.OnBufferDelete(obj)
			obj.Unlock()
			if err != nil {
				ct.logger.Errorf("Error deleting %s: %+v. Err: %v", kind, obj, err)
			}

			ct.delObject(kind, eobj.GetKey())
		}
	default:
		ct.logger.Fatalf("API watcher Found object of invalid type: %v on Buffer watch channel", tp)
	}

	return nil
}

// diffBuffer does a diff of Buffer objects between local cache and API server
func (ct *ctrlerCtx) diffBuffer(apicl apiclient.Services) {
	opts := api.ListWatchOptions{}

	// get a list of all objects from API server
	objlist, err := apicl.StagingV1().Buffer().List(context.Background(), &opts)
	if err != nil {
		ct.logger.Errorf("Error getting a list of objects. Err: %v", err)
		return
	}

	// build an object map
	objmap := make(map[string]*staging.Buffer)
	for _, obj := range objlist {
		objmap[obj.GetKey()] = obj
	}

	// if an object is in our local cache and not in API server, trigger delete for it
	for _, obj := range ct.Buffer().List() {
		_, ok := objmap[obj.GetKey()]
		if !ok {
			evt := kvstore.WatchEvent{
				Type:   kvstore.Deleted,
				Key:    obj.GetKey(),
				Object: &obj.Buffer,
			}
			ct.handleBufferEvent(&evt)
		}
	}

	// trigger create event for all others
	for _, obj := range objlist {
		evt := kvstore.WatchEvent{
			Type:   kvstore.Created,
			Key:    obj.GetKey(),
			Object: obj,
		}
		ct.handleBufferEvent(&evt)
	}
}

func (ct *ctrlerCtx) runBufferWatcher() {
	kind := "Buffer"

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
	logger := ct.logger.WithContext("submodule", "BufferWatcher")

	ct.stats.Counter("Buffer_Watch").Inc()
	defer ct.stats.Counter("Buffer_Watch").Dec()

	// loop forever
	for {
		// create a grpc client
		apicl, err := apiclient.NewGrpcAPIClient(ct.name, ct.apisrvURL, logger, rpckit.WithBalancer(balancer.New(ct.resolver)))
		if err != nil {
			logger.Warnf("Failed to connect to gRPC server [%s]\n", ct.apisrvURL)
			ct.stats.Counter("Buffer_ApiClientErr").Inc()
		} else {
			logger.Infof("API client connected {%+v}", apicl)

			// Buffer object watcher
			wt, werr := apicl.StagingV1().Buffer().Watch(ctx, &opts)
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
			ct.diffBuffer(apicl)

			// handle api server watch events
		innerLoop:
			for {
				// wait for events
				select {
				case evt, ok := <-wt.EventChan():
					if !ok {
						logger.Error("Error receiving from apisrv watcher")
						ct.stats.Counter("Buffer_WatchErrors").Inc()
						break innerLoop
					}

					// handle event
					ct.handleBufferEvent(evt)
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

// WatchBuffer starts watch on Buffer object
func (ct *ctrlerCtx) WatchBuffer(handler BufferHandler) error {
	kind := "Buffer"

	ct.Lock()
	defer ct.Unlock()

	// see if we already have a watcher
	_, ok := ct.watchers[kind]
	if ok {
		return fmt.Errorf("Buffer watcher already exists")
	}

	// save handler
	ct.handlers[kind] = handler

	// run Buffer watcher in a go routine
	go ct.runBufferWatcher()

	return nil
}

// BufferAPI returns
type BufferAPI interface {
	Create(obj *staging.Buffer) error
	Update(obj *staging.Buffer) error
	Delete(obj *staging.Buffer) error
	List() []*Buffer
	Watch(handler BufferHandler) error
}

// dummy struct that implements BufferAPI
type bufferAPI struct {
	ct *ctrlerCtx
}

// Create creates Buffer object
func (api *bufferAPI) Create(obj *staging.Buffer) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.StagingV1().Buffer().Create(context.Background(), obj)
		if err != nil && strings.Contains(err.Error(), "AlreadyExists") {
			_, err = apicl.StagingV1().Buffer().Update(context.Background(), obj)
		}
		if err != nil {
			return err
		}
	}

	return api.ct.handleBufferEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Created})
}

// Update triggers update on Buffer object
func (api *bufferAPI) Update(obj *staging.Buffer) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		_, err = apicl.StagingV1().Buffer().Update(context.Background(), obj)
		if err != nil {
			return err
		}
	}

	return api.ct.handleBufferEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Updated})
}

// Delete deletes Buffer object
func (api *bufferAPI) Delete(obj *staging.Buffer) error {
	if api.ct.resolver != nil {
		apicl, err := api.ct.apiClient()
		if err != nil {
			api.ct.logger.Errorf("Error creating API server clent. Err: %v", err)
			return err
		}

		apicl.StagingV1().Buffer().Delete(context.Background(), &obj.ObjectMeta)
	}

	return api.ct.handleBufferEvent(&kvstore.WatchEvent{Object: obj, Type: kvstore.Deleted})
}

// List returns a list of all Buffer objects
func (api *bufferAPI) List() []*Buffer {
	var objlist []*Buffer

	objs := api.ct.ListObjects("Buffer")
	for _, obj := range objs {
		switch tp := obj.(type) {
		case *Buffer:
			eobj := obj.(*Buffer)
			objlist = append(objlist, eobj)
		default:
			log.Fatalf("Got invalid object type %v while looking for Buffer", tp)
		}
	}

	return objlist
}

// Watch sets up a event handlers for Buffer object
func (api *bufferAPI) Watch(handler BufferHandler) error {
	return api.ct.WatchBuffer(handler)
}

// Buffer returns BufferAPI
func (ct *ctrlerCtx) Buffer() BufferAPI {
	return &bufferAPI{ct: ct}
}
