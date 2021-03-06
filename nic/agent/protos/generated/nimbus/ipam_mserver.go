// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package nimbus is a auto generated package.
Input file: ipam.proto
*/

package nimbus

import (
	"context"
	"errors"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/nic/agent/protos/netproto"
	hdr "github.com/pensando/sw/venice/utils/histogram"
	"github.com/pensando/sw/venice/utils/log"
	memdb "github.com/pensando/sw/venice/utils/memdb"
	"github.com/pensando/sw/venice/utils/netutils"
	"github.com/pensando/sw/venice/utils/rpckit"
)

// FindIPAMPolicy finds an IPAMPolicy by object meta
func (ms *MbusServer) FindIPAMPolicy(objmeta *api.ObjectMeta) (*netproto.IPAMPolicy, error) {
	// find the object
	obj, err := ms.memDB.FindObject("IPAMPolicy", objmeta)
	if err != nil {
		return nil, err
	}

	return IPAMPolicyFromObj(obj)
}

// ListIPAMPolicys lists all IPAMPolicys in the mbus
func (ms *MbusServer) ListIPAMPolicys(ctx context.Context, nodeID string, filters []memdb.FilterFn) ([]*netproto.IPAMPolicy, error) {
	var objlist []*netproto.IPAMPolicy

	// walk all objects
	objs := ms.memDB.ListObjectsForReceiver("IPAMPolicy", nodeID, filters)
	for _, oo := range objs {
		obj, err := IPAMPolicyFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// ListIPAMPolicysNoFilter lists all IPAMPolicys in the mbus without applying a watch filter
func (ms *MbusServer) ListIPAMPolicysNoFilter(ctx context.Context) ([]*netproto.IPAMPolicy, error) {
	var objlist []*netproto.IPAMPolicy

	// walk all objects
	objs := ms.memDB.ListObjects("IPAMPolicy", nil)
	for _, oo := range objs {
		obj, err := IPAMPolicyFromObj(oo)
		if err == nil {
			objlist = append(objlist, obj)
		}
	}

	return objlist, nil
}

// IPAMPolicyStatusReactor is the reactor interface implemented by controllers
type IPAMPolicyStatusReactor interface {
	OnIPAMPolicyCreateReq(nodeID string, objinfo *netproto.IPAMPolicy) error
	OnIPAMPolicyUpdateReq(nodeID string, objinfo *netproto.IPAMPolicy) error
	OnIPAMPolicyDeleteReq(nodeID string, objinfo *netproto.IPAMPolicy) error
	OnIPAMPolicyOperUpdate(nodeID string, objinfo *netproto.IPAMPolicy) error
	OnIPAMPolicyOperDelete(nodeID string, objinfo *netproto.IPAMPolicy) error
	GetAgentWatchFilter(ctx context.Context, kind string, watchOptions *api.ListWatchOptions) []memdb.FilterFn
}

type IPAMPolicyNodeStatus struct {
	nodeID        string
	watcher       *memdb.Watcher
	opSentStatus  map[api.EventType]*EventStatus
	opAckedStatus map[api.EventType]*EventStatus
}

// IPAMPolicyTopic is the IPAMPolicy topic on message bus
type IPAMPolicyTopic struct {
	sync.Mutex
	grpcServer    *rpckit.RPCServer // gRPC server instance
	server        *MbusServer
	statusReactor IPAMPolicyStatusReactor // status event reactor
	nodeStatus    map[string]*IPAMPolicyNodeStatus
}

// AddIPAMPolicyTopic returns a network RPC server
func AddIPAMPolicyTopic(server *MbusServer, reactor IPAMPolicyStatusReactor) (*IPAMPolicyTopic, error) {
	// RPC handler instance
	handler := IPAMPolicyTopic{
		grpcServer:    server.grpcServer,
		server:        server,
		statusReactor: reactor,
		nodeStatus:    make(map[string]*IPAMPolicyNodeStatus),
	}

	// register the RPC handlers
	if server.grpcServer != nil {
		netproto.RegisterIPAMPolicyApiV1Server(server.grpcServer.GrpcServer, &handler)
	}

	return &handler, nil
}

func (eh *IPAMPolicyTopic) registerWatcher(nodeID string, watcher *memdb.Watcher) {
	eh.Lock()
	defer eh.Unlock()

	eh.nodeStatus[nodeID] = &IPAMPolicyNodeStatus{nodeID: nodeID, watcher: watcher}
	eh.nodeStatus[nodeID].opSentStatus = make(map[api.EventType]*EventStatus)
	eh.nodeStatus[nodeID].opAckedStatus = make(map[api.EventType]*EventStatus)
}

func (eh *IPAMPolicyTopic) unRegisterWatcher(nodeID string) {
	eh.Lock()
	defer eh.Unlock()

	delete(eh.nodeStatus, nodeID)
}

//update recv object status
func (eh *IPAMPolicyTopic) updateAckedObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

	eh.Lock()
	defer eh.Unlock()
	var evStatus *EventStatus

	nodeStatus, ok := eh.nodeStatus[nodeID]
	if !ok {
		//Watcher already unregistered.
		return
	}

	evStatus, ok = nodeStatus.opAckedStatus[event]
	if !ok {
		nodeStatus.opAckedStatus[event] = &EventStatus{}
		evStatus = nodeStatus.opAckedStatus[event]
	}

	if LatencyMeasurementEnabled {
		rcvdTime, _ := objMeta.ModTime.Time()
		sendTime, _ := objMeta.CreationTime.Time()
		delta := rcvdTime.Sub(sendTime)

		hdr.Record(nodeID+"_"+"IPAMPolicy", delta)
		hdr.Record("IPAMPolicy", delta)
		hdr.Record(nodeID, delta)
	}

	new, _ := strconv.Atoi(objMeta.ResourceVersion)
	//for create/delete keep track of last one sent to, this may not be full proof
	//  Create could be processed asynchoronusly by client and can come out of order.
	//  For now should be ok as at least we make sure all messages are processed.
	//For update keep track of only last one as nimbus client periodically pulls
	if evStatus.LastObjectMeta != nil {
		current, _ := strconv.Atoi(evStatus.LastObjectMeta.ResourceVersion)
		if current > new {
			return
		}
	}
	evStatus.LastObjectMeta = objMeta
}

//update recv object status
func (eh *IPAMPolicyTopic) updateSentObjStatus(nodeID string, event api.EventType, objMeta *api.ObjectMeta) {

	eh.Lock()
	defer eh.Unlock()
	var evStatus *EventStatus

	nodeStatus, ok := eh.nodeStatus[nodeID]
	if !ok {
		//Watcher already unregistered.
		return
	}

	evStatus, ok = nodeStatus.opSentStatus[event]
	if !ok {
		nodeStatus.opSentStatus[event] = &EventStatus{}
		evStatus = nodeStatus.opSentStatus[event]
	}

	new, _ := strconv.Atoi(objMeta.ResourceVersion)
	//for create/delete keep track of last one sent to, this may not be full proof
	//  Create could be processed asynchoronusly by client and can come out of order.
	//  For now should be ok as at least we make sure all messages are processed.
	//For update keep track of only last one as nimbus client periodically pulls
	if evStatus.LastObjectMeta != nil {
		current, _ := strconv.Atoi(evStatus.LastObjectMeta.ResourceVersion)
		if current > new {
			return
		}
	}
	evStatus.LastObjectMeta = objMeta
}

//update recv object status
func (eh *IPAMPolicyTopic) WatcherInConfigSync(nodeID string, event api.EventType) bool {

	var ok bool
	var evStatus *EventStatus
	var evAckStatus *EventStatus

	eh.Lock()
	defer eh.Unlock()

	nodeStatus, ok := eh.nodeStatus[nodeID]
	if !ok {
		return true
	}

	evStatus, ok = nodeStatus.opSentStatus[event]
	if !ok {
		//nothing sent, so insync
		return true
	}

	//In-flight object still exists
	if len(nodeStatus.watcher.Channel) != 0 {
		log.Infof("watcher %v still has objects in in-flight %v(%v)", nodeID, "IPAMPolicy", event)
		return false
	}

	evAckStatus, ok = nodeStatus.opAckedStatus[event]
	if !ok {
		//nothing received, failed.
		log.Infof("watcher %v still has not received anything %v(%v)", nodeID, "IPAMPolicy", event)
		return false
	}

	if evAckStatus.LastObjectMeta.ResourceVersion < evStatus.LastObjectMeta.ResourceVersion {
		log.Infof("watcher %v resource version mismatch for %v(%v)  sent %v: recived %v",
			nodeID, "IPAMPolicy", event, evStatus.LastObjectMeta.ResourceVersion,
			evAckStatus.LastObjectMeta.ResourceVersion)
		return false
	}

	return true
}

/*
//GetSentEventStatus
func (eh *IPAMPolicyTopic) GetSentEventStatus(nodeID string, event api.EventType) *EventStatus {

    eh.Lock()
    defer eh.Unlock()
    var evStatus *EventStatus

    objStatus, ok := eh.opSentStatus[nodeID]
    if ok {
        evStatus, ok = objStatus.opStatus[event]
        if ok {
            return evStatus
        }
    }
    return nil
}


//GetAckedEventStatus
func (eh *IPAMPolicyTopic) GetAckedEventStatus(nodeID string, event api.EventType) *EventStatus {

    eh.Lock()
    defer eh.Unlock()
    var evStatus *EventStatus

    objStatus, ok := eh.opAckedStatus[nodeID]
    if ok {
        evStatus, ok = objStatus.opStatus[event]
        if ok {
            return evStatus
        }
    }
    return nil
}

*/

// CreateIPAMPolicy creates IPAMPolicy
func (eh *IPAMPolicyTopic) CreateIPAMPolicy(ctx context.Context, objinfo *netproto.IPAMPolicy) (*netproto.IPAMPolicy, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	// log.Infof("Received CreateIPAMPolicy from node %v: {%+v}", nodeID, objinfo)

	// trigger callbacks. we allow creates to happen before it exists in memdb
	if eh.statusReactor != nil {
		eh.statusReactor.OnIPAMPolicyCreateReq(nodeID, objinfo)
	}

	// increment stats
	eh.server.Stats("IPAMPolicy", "AgentCreate").Inc()

	return objinfo, nil
}

// UpdateIPAMPolicy updates IPAMPolicy
func (eh *IPAMPolicyTopic) UpdateIPAMPolicy(ctx context.Context, objinfo *netproto.IPAMPolicy) (*netproto.IPAMPolicy, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received UpdateIPAMPolicy from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("IPAMPolicy", "AgentUpdate").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnIPAMPolicyUpdateReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// DeleteIPAMPolicy deletes an IPAMPolicy
func (eh *IPAMPolicyTopic) DeleteIPAMPolicy(ctx context.Context, objinfo *netproto.IPAMPolicy) (*netproto.IPAMPolicy, error) {
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	log.Infof("Received DeleteIPAMPolicy from node %v: {%+v}", nodeID, objinfo)

	// incr stats
	eh.server.Stats("IPAMPolicy", "AgentDelete").Inc()

	// trigger callbacks
	if eh.statusReactor != nil {
		eh.statusReactor.OnIPAMPolicyDeleteReq(nodeID, objinfo)
	}

	return objinfo, nil
}

// IPAMPolicyFromObj converts memdb object to IPAMPolicy
func IPAMPolicyFromObj(obj memdb.Object) (*netproto.IPAMPolicy, error) {
	switch obj.(type) {
	case *netproto.IPAMPolicy:
		eobj := obj.(*netproto.IPAMPolicy)
		return eobj, nil
	default:
		return nil, ErrIncorrectObjectType
	}
}

// GetIPAMPolicy returns a specific IPAMPolicy
func (eh *IPAMPolicyTopic) GetIPAMPolicy(ctx context.Context, objmeta *api.ObjectMeta) (*netproto.IPAMPolicy, error) {
	// find the object
	obj, err := eh.server.memDB.FindObject("IPAMPolicy", objmeta)
	if err != nil {
		return nil, err
	}

	return IPAMPolicyFromObj(obj)
}

// ListIPAMPolicys lists all IPAMPolicys matching object selector
func (eh *IPAMPolicyTopic) ListIPAMPolicys(ctx context.Context, objsel *api.ListWatchOptions) (*netproto.IPAMPolicyList, error) {
	var objlist netproto.IPAMPolicyList
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)
	filters := []memdb.FilterFn{}

	filterFn := func(obj, prev memdb.Object) bool {
		return true
	}

	if eh.statusReactor != nil {
		filters = eh.statusReactor.GetAgentWatchFilter(ctx, "netproto.IPAMPolicy", objsel)
	} else {
		filters = append(filters, filterFn)
	}

	// walk all objects
	objs := eh.server.memDB.ListObjectsForReceiver("IPAMPolicy", nodeID, filters)
	//creationTime, _ := types.TimestampProto(time.Now())
	for _, oo := range objs {
		obj, err := IPAMPolicyFromObj(oo)
		if err == nil {
			//obj.CreationTime = api.Timestamp{Timestamp: *creationTime}
			objlist.IPAMPolicys = append(objlist.IPAMPolicys, obj)
			//record the last object sent to check config sync
			eh.updateSentObjStatus(nodeID, api.EventType_UpdateEvent, &obj.ObjectMeta)
		}
	}

	return &objlist, nil
}

// WatchIPAMPolicys watches IPAMPolicys and sends streaming resp
func (eh *IPAMPolicyTopic) WatchIPAMPolicys(watchOptions *api.ListWatchOptions, stream netproto.IPAMPolicyApiV1_WatchIPAMPolicysServer) error {
	// watch for changes
	watcher := memdb.Watcher{}
	watcher.Channel = make(chan memdb.Event, memdb.WatchLen)
	watcher.Filters = make(map[string][]memdb.FilterFn)
	defer close(watcher.Channel)

	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	if eh.statusReactor != nil {
		watcher.Filters["IPAMPolicy"] = eh.statusReactor.GetAgentWatchFilter(ctx, "IPAMPolicy", watchOptions)
	} else {
		filt := func(obj, prev memdb.Object) bool {
			return true
		}
		watcher.Filters["IPAMPolicy"] = append(watcher.Filters["IPAMPolicy"], filt)
	}

	watcher.Name = nodeID
	err := eh.server.memDB.WatchObjects("IPAMPolicy", &watcher)
	if err != nil {
		log.Errorf("Error Starting watch for kind %v Err: %v", "IPAMPolicy", err)
		return err
	}
	defer eh.server.memDB.StopWatchObjects("IPAMPolicy", &watcher)

	// get a list of all IPAMPolicys
	objlist, err := eh.ListIPAMPolicys(context.Background(), watchOptions)
	if err != nil {
		log.Errorf("Error getting a list of objects. Err: %v", err)
		return err
	}

	eh.registerWatcher(nodeID, &watcher)
	defer eh.unRegisterWatcher(nodeID)

	// increment stats
	eh.server.Stats("IPAMPolicy", "ActiveWatch").Inc()
	eh.server.Stats("IPAMPolicy", "WatchConnect").Inc()
	defer eh.server.Stats("IPAMPolicy", "ActiveWatch").Dec()
	defer eh.server.Stats("IPAMPolicy", "WatchDisconnect").Inc()

	// walk all IPAMPolicys and send it out
	watchEvts := netproto.IPAMPolicyEventList{}
	for _, obj := range objlist.IPAMPolicys {
		watchEvt := netproto.IPAMPolicyEvent{
			EventType: api.EventType_CreateEvent,

			IPAMPolicy: *obj,
		}
		watchEvts.IPAMPolicyEvents = append(watchEvts.IPAMPolicyEvents, &watchEvt)
	}
	if len(watchEvts.IPAMPolicyEvents) > 0 {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending IPAMPolicy to stream. Err: %v", err)
			return err
		}
	}
	timer := time.NewTimer(DefaultWatchHoldInterval)
	if !timer.Stop() {
		<-timer.C
	}

	running := false
	watchEvts = netproto.IPAMPolicyEventList{}
	sendToStream := func() error {
		err = stream.Send(&watchEvts)
		if err != nil {
			log.Errorf("Error sending IPAMPolicy to stream. Err: %v", err)
			return err
		}
		watchEvts = netproto.IPAMPolicyEventList{}
		return nil
	}

	// loop forever on watch channel
	for {
		select {
		// read from channel
		case evt, ok := <-watcher.Channel:
			if !ok {
				log.Errorf("Error reading from channel. Closing watch")
				return errors.New("Error reading from channel")
			}

			// convert the events
			var etype api.EventType
			switch evt.EventType {
			case memdb.CreateEvent:
				etype = api.EventType_CreateEvent
			case memdb.UpdateEvent:
				etype = api.EventType_UpdateEvent
			case memdb.DeleteEvent:
				etype = api.EventType_DeleteEvent
			}

			// get the object
			obj, err := IPAMPolicyFromObj(evt.Obj)
			if err != nil {
				return err
			}

			// convert to netproto format
			watchEvt := netproto.IPAMPolicyEvent{
				EventType: etype,

				IPAMPolicy: *obj,
			}
			watchEvts.IPAMPolicyEvents = append(watchEvts.IPAMPolicyEvents, &watchEvt)
			if !running {
				running = true
				timer.Reset(DefaultWatchHoldInterval)
			}
			if len(watchEvts.IPAMPolicyEvents) >= DefaultWatchBatchSize {
				if err = sendToStream(); err != nil {
					return err
				}
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(DefaultWatchHoldInterval)
			}
			eh.updateSentObjStatus(nodeID, etype, &obj.ObjectMeta)
		case <-timer.C:
			running = false
			if err = sendToStream(); err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}

	// done
}

// updateIPAMPolicyOper triggers oper update callbacks
func (eh *IPAMPolicyTopic) updateIPAMPolicyOper(oper *netproto.IPAMPolicyEvent, nodeID string) error {
	eh.updateAckedObjStatus(nodeID, oper.EventType, &oper.IPAMPolicy.ObjectMeta)
	switch oper.EventType {
	case api.EventType_CreateEvent:
		fallthrough
	case api.EventType_UpdateEvent:
		// incr stats
		eh.server.Stats("IPAMPolicy", "AgentUpdate").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {

			return eh.statusReactor.OnIPAMPolicyOperUpdate(nodeID, &oper.IPAMPolicy)

		}
	case api.EventType_DeleteEvent:
		// incr stats
		eh.server.Stats("IPAMPolicy", "AgentDelete").Inc()

		// trigger callbacks
		if eh.statusReactor != nil {

			eh.statusReactor.OnIPAMPolicyOperDelete(nodeID, &oper.IPAMPolicy)

		}
	}

	return nil
}

func (eh *IPAMPolicyTopic) IPAMPolicyOperUpdate(stream netproto.IPAMPolicyApiV1_IPAMPolicyOperUpdateServer) error {
	ctx := stream.Context()
	nodeID := netutils.GetNodeUUIDFromCtx(ctx)

	for {
		oper, err := stream.Recv()
		if err == io.EOF {
			log.Errorf("%v IPAMPolicyOperUpdate stream ended. closing..", nodeID)
			return stream.SendAndClose(&api.TypeMeta{})
		} else if err != nil {
			log.Errorf("Error receiving from %v IPAMPolicyOperUpdate stream. Err: %v", nodeID, err)
			return err
		}

		err = eh.updateIPAMPolicyOper(oper, nodeID)
		if err != nil {
			log.Errorf("Error updating IPAMPolicy oper state. Err: %v", err)
		}
	}
}
