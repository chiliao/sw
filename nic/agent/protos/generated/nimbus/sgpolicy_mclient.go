// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: sgpolicy.proto
*/

package nimbus

import (
	"context"
	"sync"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/netagent/state"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	"github.com/pensando/sw/venice/utils/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"
)

type NetworkSecurityPolicyReactor interface {
	CreateNetworkSecurityPolicy(networksecuritypolicyObj *netproto.NetworkSecurityPolicy) error // creates an NetworkSecurityPolicy
	FindNetworkSecurityPolicy(meta api.ObjectMeta) (*netproto.NetworkSecurityPolicy, error)     // finds an NetworkSecurityPolicy
	ListNetworkSecurityPolicy() []*netproto.NetworkSecurityPolicy                               // lists all NetworkSecurityPolicys
	UpdateNetworkSecurityPolicy(networksecuritypolicyObj *netproto.NetworkSecurityPolicy) error // updates an NetworkSecurityPolicy
	DeleteNetworkSecurityPolicy(networksecuritypolicyObj, ns, name string) error                // deletes an NetworkSecurityPolicy
	GetWatchOptions(cts context.Context, kind string) api.ObjectMeta
}
type NetworkSecurityPolicyOStream struct {
	sync.Mutex
	stream netproto.NetworkSecurityPolicyApi_NetworkSecurityPolicyOperUpdateClient
}

// WatchNetworkSecurityPolicys runs NetworkSecurityPolicy watcher loop
func (client *NimbusClient) WatchNetworkSecurityPolicys(ctx context.Context, reactor NetworkSecurityPolicyReactor) {
	// setup wait group
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()
	client.debugStats.AddInt("ActiveNetworkSecurityPolicyWatch", 1)

	// make sure rpc client is good
	if client.rpcClient == nil || client.rpcClient.ClientConn == nil || client.rpcClient.ClientConn.GetState() != connectivity.Ready {
		log.Errorf("RPC client is disconnected. Exiting watch")
		return
	}

	// start the watch
	ometa := reactor.GetWatchOptions(ctx, "NetworkSecurityPolicy")
	networksecuritypolicyRPCClient := netproto.NewNetworkSecurityPolicyApiClient(client.rpcClient.ClientConn)
	stream, err := networksecuritypolicyRPCClient.WatchNetworkSecurityPolicys(ctx, &ometa)
	if err != nil {
		log.Errorf("Error watching NetworkSecurityPolicy. Err: %v", err)
		return
	}

	// start oper update stream
	opStream, err := networksecuritypolicyRPCClient.NetworkSecurityPolicyOperUpdate(ctx)
	if err != nil {
		log.Errorf("Error starting NetworkSecurityPolicy oper updates. Err: %v", err)
		return
	}

	ostream := &NetworkSecurityPolicyOStream{stream: opStream}

	// get a list of objects
	objList, err := networksecuritypolicyRPCClient.ListNetworkSecurityPolicys(ctx, &ometa)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok || st.Code() == codes.Unavailable {
			log.Errorf("Error getting NetworkSecurityPolicy list. Err: %v", err)
			return
		}
	} else {
		// perform a diff of the states
		client.diffNetworkSecurityPolicys(objList, reactor, ostream)
	}

	// start grpc stream recv
	recvCh := make(chan *netproto.NetworkSecurityPolicyEvent, evChanLength)
	go client.watchNetworkSecurityPolicyRecvLoop(stream, recvCh)

	// loop till the end
	for {
		evtWork := func(evt *netproto.NetworkSecurityPolicyEvent) {
			client.debugStats.AddInt("NetworkSecurityPolicyWatchEvents", 1)
			log.Infof("Ctrlerif: agent %s got NetworkSecurityPolicy watch event: Type: {%+v} NetworkSecurityPolicy:{%+v}", client.clientName, evt.EventType, evt.NetworkSecurityPolicy.ObjectMeta)
			client.lockObject(evt.NetworkSecurityPolicy.GetObjectKind(), evt.NetworkSecurityPolicy.ObjectMeta)
			go client.processNetworkSecurityPolicyEvent(*evt, reactor, ostream)
			//Give it some time to increment waitgrp
			time.Sleep(100 * time.Microsecond)
		}
		//Give priority to evnt work.
		select {
		case evt, ok := <-recvCh:
			if !ok {
				log.Warnf("NetworkSecurityPolicy Watch channel closed. Exisint NetworkSecurityPolicyWatch")
				return
			}
			evtWork(evt)
			// periodic resync (Disabling as we have aggregate watch support)
			/*case <-time.After(resyncInterval):
			            //Give priority to evt work
			            //Wait for batch interval for inflight work
			            time.Sleep(5 * DefaultWatchHoldInterval)
			            select {
			            case evt, ok := <-recvCh:
			                if !ok {
			                    log.Warnf("NetworkSecurityPolicy Watch channel closed. Exisint NetworkSecurityPolicyWatch")
			                    return
			                }
			                evtWork(evt)
							continue
			            default:
			            }
						// get a list of objects
						objList, err := networksecuritypolicyRPCClient.ListNetworkSecurityPolicys(ctx, &ometa)
						if err != nil {
							st, ok := status.FromError(err)
							if !ok || st.Code() == codes.Unavailable {
								log.Errorf("Error getting NetworkSecurityPolicy list. Err: %v", err)
								return
							}
						} else {
							client.debugStats.AddInt("NetworkSecurityPolicyWatchResyncs", 1)
							// perform a diff of the states
							client.diffNetworkSecurityPolicys(objList, reactor, ostream)
						}
			*/
		}
	}
}

// watchNetworkSecurityPolicyRecvLoop receives from stream and write it to a channel
func (client *NimbusClient) watchNetworkSecurityPolicyRecvLoop(stream netproto.NetworkSecurityPolicyApi_WatchNetworkSecurityPolicysClient, recvch chan<- *netproto.NetworkSecurityPolicyEvent) {
	defer close(recvch)
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// loop till the end
	for {
		// receive from stream
		objList, err := stream.Recv()
		if err != nil {
			log.Errorf("Error receiving from watch channel. Exiting NetworkSecurityPolicy watch. Err: %v", err)
			return
		}
		for _, evt := range objList.NetworkSecurityPolicyEvents {
			recvch <- evt
		}
	}
}

// diffNetworkSecurityPolicy diffs local state with controller state
// FIXME: this is not handling deletes today
func (client *NimbusClient) diffNetworkSecurityPolicys(objList *netproto.NetworkSecurityPolicyList, reactor NetworkSecurityPolicyReactor, ostream *NetworkSecurityPolicyOStream) {
	// build a map of objects
	objmap := make(map[string]*netproto.NetworkSecurityPolicy)
	for _, obj := range objList.NetworkSecurityPolicys {
		key := obj.ObjectMeta.GetKey()
		objmap[key] = obj
	}

	// see if we need to delete any locally found object
	localObjs := reactor.ListNetworkSecurityPolicy()
	for _, lobj := range localObjs {
		ctby, ok := lobj.ObjectMeta.Labels["CreatedBy"]
		if ok && ctby == "Venice" {
			key := lobj.ObjectMeta.GetKey()
			if _, ok := objmap[key]; !ok {
				evt := netproto.NetworkSecurityPolicyEvent{
					EventType:             api.EventType_DeleteEvent,
					NetworkSecurityPolicy: *lobj,
				}
				log.Infof("diffNetworkSecurityPolicys(): Deleting object %+v", lobj.ObjectMeta)
				client.lockObject(evt.NetworkSecurityPolicy.GetObjectKind(), evt.NetworkSecurityPolicy.ObjectMeta)
				client.processNetworkSecurityPolicyEvent(evt, reactor, ostream)
			}
		} else {
			log.Infof("Not deleting non-venice object %+v", lobj.ObjectMeta)
		}
	}

	// add/update all new objects
	for _, obj := range objList.NetworkSecurityPolicys {
		evt := netproto.NetworkSecurityPolicyEvent{
			EventType:             api.EventType_UpdateEvent,
			NetworkSecurityPolicy: *obj,
		}
		client.lockObject(evt.NetworkSecurityPolicy.GetObjectKind(), evt.NetworkSecurityPolicy.ObjectMeta)
		client.processNetworkSecurityPolicyEvent(evt, reactor, ostream)
	}
}

// processNetworkSecurityPolicyEvent handles NetworkSecurityPolicy event
func (client *NimbusClient) processNetworkSecurityPolicyEvent(evt netproto.NetworkSecurityPolicyEvent, reactor NetworkSecurityPolicyReactor, ostream *NetworkSecurityPolicyOStream) error {
	var err error
	client.waitGrp.Add(1)
	defer client.waitGrp.Done()

	// add venice label to the object
	evt.NetworkSecurityPolicy.ObjectMeta.Labels = make(map[string]string)
	evt.NetworkSecurityPolicy.ObjectMeta.Labels["CreatedBy"] = "Venice"

	// unlock the object once we are done
	defer client.unlockObject(evt.NetworkSecurityPolicy.GetObjectKind(), evt.NetworkSecurityPolicy.ObjectMeta)

	// retry till successful
	for iter := 0; iter < maxOpretry; iter++ {
		switch evt.EventType {
		case api.EventType_CreateEvent:
			fallthrough
		case api.EventType_UpdateEvent:
			_, err = reactor.FindNetworkSecurityPolicy(evt.NetworkSecurityPolicy.ObjectMeta)
			if err != nil {
				// create the NetworkSecurityPolicy
				err = reactor.CreateNetworkSecurityPolicy(&evt.NetworkSecurityPolicy)
				if err != nil {
					log.Errorf("Error creating the NetworkSecurityPolicy {%+v}. Err: %v", evt.NetworkSecurityPolicy.ObjectMeta, err)
					client.debugStats.AddInt("CreateNetworkSecurityPolicyError", 1)
				} else {
					client.debugStats.AddInt("CreateNetworkSecurityPolicy", 1)
				}
			} else {
				// update the NetworkSecurityPolicy
				err = reactor.UpdateNetworkSecurityPolicy(&evt.NetworkSecurityPolicy)
				if err != nil {
					log.Errorf("Error updating the NetworkSecurityPolicy {%+v}. Err: %v", evt.NetworkSecurityPolicy.GetKey(), err)
					client.debugStats.AddInt("UpdateNetworkSecurityPolicyError", 1)
				} else {
					client.debugStats.AddInt("UpdateNetworkSecurityPolicy", 1)
				}
			}

		case api.EventType_DeleteEvent:
			// delete the object
			err = reactor.DeleteNetworkSecurityPolicy(evt.NetworkSecurityPolicy.Tenant, evt.NetworkSecurityPolicy.Namespace, evt.NetworkSecurityPolicy.Name)
			if err == state.ErrObjectNotFound { // give idempotency to caller
				log.Debugf("NetworkSecurityPolicy {%+v} not found", evt.NetworkSecurityPolicy.ObjectMeta)
				err = nil
			}
			if err != nil {
				log.Errorf("Error deleting the NetworkSecurityPolicy {%+v}. Err: %v", evt.NetworkSecurityPolicy.ObjectMeta, err)
				client.debugStats.AddInt("DeleteNetworkSecurityPolicyError", 1)
			} else {
				client.debugStats.AddInt("DeleteNetworkSecurityPolicy", 1)
			}
		}

		if ostream == nil {
			return err
		}
		// send oper status and return if there is no error
		if err == nil {
			robj := netproto.NetworkSecurityPolicyEvent{
				EventType: evt.EventType,
				NetworkSecurityPolicy: netproto.NetworkSecurityPolicy{
					TypeMeta:   evt.NetworkSecurityPolicy.TypeMeta,
					ObjectMeta: evt.NetworkSecurityPolicy.ObjectMeta,
					Status:     evt.NetworkSecurityPolicy.Status,
				},
			}

			// send oper status
			ostream.Lock()
			modificationTime, _ := types.TimestampProto(time.Now())
			robj.NetworkSecurityPolicy.ObjectMeta.ModTime = api.Timestamp{Timestamp: *modificationTime}
			err := ostream.stream.Send(&robj)
			if err != nil {
				log.Errorf("failed to send NetworkSecurityPolicy oper Status, %s", err)
				client.debugStats.AddInt("NetworkSecurityPolicyOperSendError", 1)
			} else {
				client.debugStats.AddInt("NetworkSecurityPolicyOperSent", 1)
			}
			ostream.Unlock()

			return err
		}

		// else, retry after some time, with backoff
		time.Sleep(time.Second * time.Duration(2*iter))
	}

	return nil
}

func (client *NimbusClient) processNetworkSecurityPolicyDynamic(evt api.EventType,
	object *netproto.NetworkSecurityPolicy, reactor NetworkSecurityPolicyReactor) error {

	networksecuritypolicyEvt := netproto.NetworkSecurityPolicyEvent{
		EventType:             evt,
		NetworkSecurityPolicy: *object,
	}

	// add venice label to the object
	networksecuritypolicyEvt.NetworkSecurityPolicy.ObjectMeta.Labels = make(map[string]string)
	networksecuritypolicyEvt.NetworkSecurityPolicy.ObjectMeta.Labels["CreatedBy"] = "Venice"

	client.lockObject(networksecuritypolicyEvt.NetworkSecurityPolicy.GetObjectKind(), networksecuritypolicyEvt.NetworkSecurityPolicy.ObjectMeta)

	err := client.processNetworkSecurityPolicyEvent(networksecuritypolicyEvt, reactor, nil)
	modificationTime, _ := types.TimestampProto(time.Now())
	object.ObjectMeta.ModTime = api.Timestamp{Timestamp: *modificationTime}

	return err
}
