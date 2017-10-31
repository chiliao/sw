// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

package grpcclient

import (
	"context"
	"errors"
	oldlog "log"

	"github.com/go-kit/kit/endpoint"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"

	api "github.com/pensando/sw/api"
	cmd "github.com/pensando/sw/api/generated/cmd"
	listerwatcher "github.com/pensando/sw/api/listerwatcher"
	apiserver "github.com/pensando/sw/venice/apiserver"
	"github.com/pensando/sw/venice/utils/kvstore"
	"github.com/pensando/sw/venice/utils/log"
	"github.com/pensando/sw/venice/utils/trace"
)

// Dummy vars to suppress import errors
var _ api.TypeMeta
var _ listerwatcher.WatcherClient
var _ kvstore.Interface

// NewCmdV1 sets up a new client for CmdV1
func NewCmdV1(conn *grpc.ClientConn, logger log.Logger) cmd.ServiceCmdV1Client {

	var lAutoAddClusterEndpoint endpoint.Endpoint
	{
		lAutoAddClusterEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoAddCluster",
			cmd.EncodeGrpcReqCluster,
			cmd.DecodeGrpcRespCluster,
			&cmd.Cluster{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoAddClusterEndpoint = trace.ClientEndPoint("CmdV1:AutoAddCluster")(lAutoAddClusterEndpoint)
	}
	var lAutoAddNodeEndpoint endpoint.Endpoint
	{
		lAutoAddNodeEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoAddNode",
			cmd.EncodeGrpcReqNode,
			cmd.DecodeGrpcRespNode,
			&cmd.Node{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoAddNodeEndpoint = trace.ClientEndPoint("CmdV1:AutoAddNode")(lAutoAddNodeEndpoint)
	}
	var lAutoAddSmartNICEndpoint endpoint.Endpoint
	{
		lAutoAddSmartNICEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoAddSmartNIC",
			cmd.EncodeGrpcReqSmartNIC,
			cmd.DecodeGrpcRespSmartNIC,
			&cmd.SmartNIC{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoAddSmartNICEndpoint = trace.ClientEndPoint("CmdV1:AutoAddSmartNIC")(lAutoAddSmartNICEndpoint)
	}
	var lAutoDeleteClusterEndpoint endpoint.Endpoint
	{
		lAutoDeleteClusterEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoDeleteCluster",
			cmd.EncodeGrpcReqCluster,
			cmd.DecodeGrpcRespCluster,
			&cmd.Cluster{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoDeleteClusterEndpoint = trace.ClientEndPoint("CmdV1:AutoDeleteCluster")(lAutoDeleteClusterEndpoint)
	}
	var lAutoDeleteNodeEndpoint endpoint.Endpoint
	{
		lAutoDeleteNodeEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoDeleteNode",
			cmd.EncodeGrpcReqNode,
			cmd.DecodeGrpcRespNode,
			&cmd.Node{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoDeleteNodeEndpoint = trace.ClientEndPoint("CmdV1:AutoDeleteNode")(lAutoDeleteNodeEndpoint)
	}
	var lAutoDeleteSmartNICEndpoint endpoint.Endpoint
	{
		lAutoDeleteSmartNICEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoDeleteSmartNIC",
			cmd.EncodeGrpcReqSmartNIC,
			cmd.DecodeGrpcRespSmartNIC,
			&cmd.SmartNIC{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoDeleteSmartNICEndpoint = trace.ClientEndPoint("CmdV1:AutoDeleteSmartNIC")(lAutoDeleteSmartNICEndpoint)
	}
	var lAutoGetClusterEndpoint endpoint.Endpoint
	{
		lAutoGetClusterEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoGetCluster",
			cmd.EncodeGrpcReqCluster,
			cmd.DecodeGrpcRespCluster,
			&cmd.Cluster{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoGetClusterEndpoint = trace.ClientEndPoint("CmdV1:AutoGetCluster")(lAutoGetClusterEndpoint)
	}
	var lAutoGetNodeEndpoint endpoint.Endpoint
	{
		lAutoGetNodeEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoGetNode",
			cmd.EncodeGrpcReqNode,
			cmd.DecodeGrpcRespNode,
			&cmd.Node{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoGetNodeEndpoint = trace.ClientEndPoint("CmdV1:AutoGetNode")(lAutoGetNodeEndpoint)
	}
	var lAutoGetSmartNICEndpoint endpoint.Endpoint
	{
		lAutoGetSmartNICEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoGetSmartNIC",
			cmd.EncodeGrpcReqSmartNIC,
			cmd.DecodeGrpcRespSmartNIC,
			&cmd.SmartNIC{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoGetSmartNICEndpoint = trace.ClientEndPoint("CmdV1:AutoGetSmartNIC")(lAutoGetSmartNICEndpoint)
	}
	var lAutoListClusterEndpoint endpoint.Endpoint
	{
		lAutoListClusterEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoListCluster",
			cmd.EncodeGrpcReqListWatchOptions,
			cmd.DecodeGrpcRespClusterList,
			&cmd.ClusterList{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoListClusterEndpoint = trace.ClientEndPoint("CmdV1:AutoListCluster")(lAutoListClusterEndpoint)
	}
	var lAutoListNodeEndpoint endpoint.Endpoint
	{
		lAutoListNodeEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoListNode",
			cmd.EncodeGrpcReqListWatchOptions,
			cmd.DecodeGrpcRespNodeList,
			&cmd.NodeList{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoListNodeEndpoint = trace.ClientEndPoint("CmdV1:AutoListNode")(lAutoListNodeEndpoint)
	}
	var lAutoListSmartNICEndpoint endpoint.Endpoint
	{
		lAutoListSmartNICEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoListSmartNIC",
			cmd.EncodeGrpcReqListWatchOptions,
			cmd.DecodeGrpcRespSmartNICList,
			&cmd.SmartNICList{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoListSmartNICEndpoint = trace.ClientEndPoint("CmdV1:AutoListSmartNIC")(lAutoListSmartNICEndpoint)
	}
	var lAutoUpdateClusterEndpoint endpoint.Endpoint
	{
		lAutoUpdateClusterEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoUpdateCluster",
			cmd.EncodeGrpcReqCluster,
			cmd.DecodeGrpcRespCluster,
			&cmd.Cluster{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoUpdateClusterEndpoint = trace.ClientEndPoint("CmdV1:AutoUpdateCluster")(lAutoUpdateClusterEndpoint)
	}
	var lAutoUpdateNodeEndpoint endpoint.Endpoint
	{
		lAutoUpdateNodeEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoUpdateNode",
			cmd.EncodeGrpcReqNode,
			cmd.DecodeGrpcRespNode,
			&cmd.Node{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoUpdateNodeEndpoint = trace.ClientEndPoint("CmdV1:AutoUpdateNode")(lAutoUpdateNodeEndpoint)
	}
	var lAutoUpdateSmartNICEndpoint endpoint.Endpoint
	{
		lAutoUpdateSmartNICEndpoint = grpctransport.NewClient(
			conn,
			"cmd.CmdV1",
			"AutoUpdateSmartNIC",
			cmd.EncodeGrpcReqSmartNIC,
			cmd.DecodeGrpcRespSmartNIC,
			&cmd.SmartNIC{},
			grpctransport.ClientBefore(trace.ToGRPCRequest(logger)),
			grpctransport.ClientBefore(dummyBefore),
		).Endpoint()
		lAutoUpdateSmartNICEndpoint = trace.ClientEndPoint("CmdV1:AutoUpdateSmartNIC")(lAutoUpdateSmartNICEndpoint)
	}
	return cmd.EndpointsCmdV1Client{
		Client: cmd.NewCmdV1Client(conn),

		AutoAddClusterEndpoint:     lAutoAddClusterEndpoint,
		AutoAddNodeEndpoint:        lAutoAddNodeEndpoint,
		AutoAddSmartNICEndpoint:    lAutoAddSmartNICEndpoint,
		AutoDeleteClusterEndpoint:  lAutoDeleteClusterEndpoint,
		AutoDeleteNodeEndpoint:     lAutoDeleteNodeEndpoint,
		AutoDeleteSmartNICEndpoint: lAutoDeleteSmartNICEndpoint,
		AutoGetClusterEndpoint:     lAutoGetClusterEndpoint,
		AutoGetNodeEndpoint:        lAutoGetNodeEndpoint,
		AutoGetSmartNICEndpoint:    lAutoGetSmartNICEndpoint,
		AutoListClusterEndpoint:    lAutoListClusterEndpoint,
		AutoListNodeEndpoint:       lAutoListNodeEndpoint,
		AutoListSmartNICEndpoint:   lAutoListSmartNICEndpoint,
		AutoUpdateClusterEndpoint:  lAutoUpdateClusterEndpoint,
		AutoUpdateNodeEndpoint:     lAutoUpdateNodeEndpoint,
		AutoUpdateSmartNICEndpoint: lAutoUpdateSmartNICEndpoint,
	}
}

// NewCmdV1Backend creates an instrumented client with middleware
func NewCmdV1Backend(conn *grpc.ClientConn, logger log.Logger) cmd.ServiceCmdV1Client {
	cl := NewCmdV1(conn, logger)
	cl = cmd.LoggingCmdV1MiddlewareClient(logger)(cl)
	return cl
}

type grpcObjCmdV1Cluster struct {
	logger log.Logger
	client cmd.ServiceCmdV1Client
}

func (a *grpcObjCmdV1Cluster) Create(ctx context.Context, in *cmd.Cluster) (*cmd.Cluster, error) {
	a.logger.DebugLog("msg", "recieved call", "object", "Cluster", "oper", "create")
	if in == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	return a.client.AutoAddCluster(nctx, in)
}

func (a *grpcObjCmdV1Cluster) Update(ctx context.Context, in *cmd.Cluster) (*cmd.Cluster, error) {
	a.logger.DebugLog("msg", "received call", "object", "Cluster", "oper", "update")
	if in == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	return a.client.AutoUpdateCluster(nctx, in)
}

func (a *grpcObjCmdV1Cluster) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.Cluster, error) {
	a.logger.DebugLog("msg", "received call", "object", "Cluster", "oper", "get")
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.Cluster{}
	in.ObjectMeta = *objMeta
	nctx := addVersion(ctx, "v1")
	return a.client.AutoGetCluster(nctx, &in)
}

func (a *grpcObjCmdV1Cluster) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.Cluster, error) {
	a.logger.DebugLog("msg", "received call", "object", "Cluster", "oper", "delete")
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.Cluster{}
	in.ObjectMeta = *objMeta
	nctx := addVersion(ctx, "v1")
	return a.client.AutoDeleteCluster(nctx, &in)
}

func (a *grpcObjCmdV1Cluster) List(ctx context.Context, options *api.ListWatchOptions) ([]*cmd.Cluster, error) {
	a.logger.DebugLog("msg", "received call", "object", "Cluster", "oper", "list")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	r, err := a.client.AutoListCluster(nctx, options)
	if err == nil {
		return r.Items, nil
	}
	return nil, err
}

func (a *grpcObjCmdV1Cluster) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	a.logger.DebugLog("msg", "received call", "object", "Cluster", "oper", "WatchOper")
	nctx := addVersion(ctx, "v1")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	stream, err := a.client.AutoWatchCluster(nctx, options)
	if err != nil {
		return nil, err
	}
	wstream := stream.(cmd.CmdV1_AutoWatchClusterClient)
	bridgefn := func(lw *listerwatcher.WatcherClient) {
		for {
			r, err := wstream.Recv()
			if err != nil {
				a.logger.ErrorLog("msg", "error on recieve", "error", err)
				close(lw.OutCh)
				return
			}
			ev := kvstore.WatchEvent{
				Type:   kvstore.WatchEventType(r.Type),
				Object: r.Object,
			}
			select {
			case lw.OutCh <- &ev:
			case <-wstream.Context().Done():
				close(lw.OutCh)
				return
			}
		}
	}
	lw := listerwatcher.NewWatcherClient(wstream, bridgefn)
	lw.Run()
	return lw, nil
}

func (a *grpcObjCmdV1Cluster) Allowed(oper apiserver.APIOperType) bool {
	return true
}

type restObjCmdV1Cluster struct {
	endpoints cmd.EndpointsCmdV1RestClient
	instance  string
}

func (a *restObjCmdV1Cluster) Create(ctx context.Context, in *cmd.Cluster) (*cmd.Cluster, error) {
	if in == nil {
		return nil, errors.New("invalid input")
	}
	return a.endpoints.AutoAddCluster(ctx, in)
}

func (a *restObjCmdV1Cluster) Update(ctx context.Context, in *cmd.Cluster) (*cmd.Cluster, error) {
	if in == nil {
		return nil, errors.New("invalid input")
	}
	return a.endpoints.AutoUpdateCluster(ctx, in)
}

func (a *restObjCmdV1Cluster) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.Cluster, error) {
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.Cluster{}
	in.ObjectMeta = *objMeta
	return a.endpoints.AutoGetCluster(ctx, &in)
}

func (a *restObjCmdV1Cluster) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.Cluster, error) {
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.Cluster{}
	in.ObjectMeta = *objMeta
	return a.endpoints.AutoDeleteCluster(ctx, &in)
}

func (a *restObjCmdV1Cluster) List(ctx context.Context, options *api.ListWatchOptions) ([]*cmd.Cluster, error) {
	if options == nil {
		return nil, errors.New("invalid input")
	}
	r, err := a.endpoints.AutoListCluster(ctx, options)
	if err == nil {
		return r.Items, nil
	}
	return nil, err
}

func (a *restObjCmdV1Cluster) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	return nil, errors.New("not allowed")
}

func (a *restObjCmdV1Cluster) Allowed(oper apiserver.APIOperType) bool {
	switch oper {
	case apiserver.CreateOper:
		return false
	case apiserver.UpdateOper:
		return true
	case apiserver.GetOper:
		return true
	case apiserver.DeleteOper:
		return true
	case apiserver.ListOper:
		return true
	case apiserver.WatchOper:
		return false
	default:
		return false
	}
}

type grpcObjCmdV1Node struct {
	logger log.Logger
	client cmd.ServiceCmdV1Client
}

func (a *grpcObjCmdV1Node) Create(ctx context.Context, in *cmd.Node) (*cmd.Node, error) {
	a.logger.DebugLog("msg", "recieved call", "object", "Node", "oper", "create")
	if in == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	return a.client.AutoAddNode(nctx, in)
}

func (a *grpcObjCmdV1Node) Update(ctx context.Context, in *cmd.Node) (*cmd.Node, error) {
	a.logger.DebugLog("msg", "received call", "object", "Node", "oper", "update")
	if in == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	return a.client.AutoUpdateNode(nctx, in)
}

func (a *grpcObjCmdV1Node) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.Node, error) {
	a.logger.DebugLog("msg", "received call", "object", "Node", "oper", "get")
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.Node{}
	in.ObjectMeta = *objMeta
	nctx := addVersion(ctx, "v1")
	return a.client.AutoGetNode(nctx, &in)
}

func (a *grpcObjCmdV1Node) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.Node, error) {
	a.logger.DebugLog("msg", "received call", "object", "Node", "oper", "delete")
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.Node{}
	in.ObjectMeta = *objMeta
	nctx := addVersion(ctx, "v1")
	return a.client.AutoDeleteNode(nctx, &in)
}

func (a *grpcObjCmdV1Node) List(ctx context.Context, options *api.ListWatchOptions) ([]*cmd.Node, error) {
	a.logger.DebugLog("msg", "received call", "object", "Node", "oper", "list")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	r, err := a.client.AutoListNode(nctx, options)
	if err == nil {
		return r.Items, nil
	}
	return nil, err
}

func (a *grpcObjCmdV1Node) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	a.logger.DebugLog("msg", "received call", "object", "Node", "oper", "WatchOper")
	nctx := addVersion(ctx, "v1")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	stream, err := a.client.AutoWatchNode(nctx, options)
	if err != nil {
		return nil, err
	}
	wstream := stream.(cmd.CmdV1_AutoWatchNodeClient)
	bridgefn := func(lw *listerwatcher.WatcherClient) {
		for {
			r, err := wstream.Recv()
			if err != nil {
				a.logger.ErrorLog("msg", "error on recieve", "error", err)
				close(lw.OutCh)
				return
			}
			ev := kvstore.WatchEvent{
				Type:   kvstore.WatchEventType(r.Type),
				Object: r.Object,
			}
			select {
			case lw.OutCh <- &ev:
			case <-wstream.Context().Done():
				close(lw.OutCh)
				return
			}
		}
	}
	lw := listerwatcher.NewWatcherClient(wstream, bridgefn)
	lw.Run()
	return lw, nil
}

func (a *grpcObjCmdV1Node) Allowed(oper apiserver.APIOperType) bool {
	return true
}

type restObjCmdV1Node struct {
	endpoints cmd.EndpointsCmdV1RestClient
	instance  string
}

func (a *restObjCmdV1Node) Create(ctx context.Context, in *cmd.Node) (*cmd.Node, error) {
	if in == nil {
		return nil, errors.New("invalid input")
	}
	return a.endpoints.AutoAddNode(ctx, in)
}

func (a *restObjCmdV1Node) Update(ctx context.Context, in *cmd.Node) (*cmd.Node, error) {
	if in == nil {
		return nil, errors.New("invalid input")
	}
	return a.endpoints.AutoUpdateNode(ctx, in)
}

func (a *restObjCmdV1Node) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.Node, error) {
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.Node{}
	in.ObjectMeta = *objMeta
	return a.endpoints.AutoGetNode(ctx, &in)
}

func (a *restObjCmdV1Node) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.Node, error) {
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.Node{}
	in.ObjectMeta = *objMeta
	return a.endpoints.AutoDeleteNode(ctx, &in)
}

func (a *restObjCmdV1Node) List(ctx context.Context, options *api.ListWatchOptions) ([]*cmd.Node, error) {
	if options == nil {
		return nil, errors.New("invalid input")
	}
	r, err := a.endpoints.AutoListNode(ctx, options)
	if err == nil {
		return r.Items, nil
	}
	return nil, err
}

func (a *restObjCmdV1Node) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	return nil, errors.New("not allowed")
}

func (a *restObjCmdV1Node) Allowed(oper apiserver.APIOperType) bool {
	switch oper {
	case apiserver.CreateOper:
		return true
	case apiserver.UpdateOper:
		return true
	case apiserver.GetOper:
		return true
	case apiserver.DeleteOper:
		return true
	case apiserver.ListOper:
		return true
	case apiserver.WatchOper:
		return false
	default:
		return false
	}
}

type grpcObjCmdV1SmartNIC struct {
	logger log.Logger
	client cmd.ServiceCmdV1Client
}

func (a *grpcObjCmdV1SmartNIC) Create(ctx context.Context, in *cmd.SmartNIC) (*cmd.SmartNIC, error) {
	a.logger.DebugLog("msg", "recieved call", "object", "SmartNIC", "oper", "create")
	if in == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	return a.client.AutoAddSmartNIC(nctx, in)
}

func (a *grpcObjCmdV1SmartNIC) Update(ctx context.Context, in *cmd.SmartNIC) (*cmd.SmartNIC, error) {
	a.logger.DebugLog("msg", "received call", "object", "SmartNIC", "oper", "update")
	if in == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	return a.client.AutoUpdateSmartNIC(nctx, in)
}

func (a *grpcObjCmdV1SmartNIC) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.SmartNIC, error) {
	a.logger.DebugLog("msg", "received call", "object", "SmartNIC", "oper", "get")
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.SmartNIC{}
	in.ObjectMeta = *objMeta
	nctx := addVersion(ctx, "v1")
	return a.client.AutoGetSmartNIC(nctx, &in)
}

func (a *grpcObjCmdV1SmartNIC) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.SmartNIC, error) {
	a.logger.DebugLog("msg", "received call", "object", "SmartNIC", "oper", "delete")
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.SmartNIC{}
	in.ObjectMeta = *objMeta
	nctx := addVersion(ctx, "v1")
	return a.client.AutoDeleteSmartNIC(nctx, &in)
}

func (a *grpcObjCmdV1SmartNIC) List(ctx context.Context, options *api.ListWatchOptions) ([]*cmd.SmartNIC, error) {
	a.logger.DebugLog("msg", "received call", "object", "SmartNIC", "oper", "list")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	nctx := addVersion(ctx, "v1")
	r, err := a.client.AutoListSmartNIC(nctx, options)
	if err == nil {
		return r.Items, nil
	}
	return nil, err
}

func (a *grpcObjCmdV1SmartNIC) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	a.logger.DebugLog("msg", "received call", "object", "SmartNIC", "oper", "WatchOper")
	nctx := addVersion(ctx, "v1")
	if options == nil {
		return nil, errors.New("invalid input")
	}
	stream, err := a.client.AutoWatchSmartNIC(nctx, options)
	if err != nil {
		return nil, err
	}
	wstream := stream.(cmd.CmdV1_AutoWatchSmartNICClient)
	bridgefn := func(lw *listerwatcher.WatcherClient) {
		for {
			r, err := wstream.Recv()
			if err != nil {
				a.logger.ErrorLog("msg", "error on recieve", "error", err)
				close(lw.OutCh)
				return
			}
			ev := kvstore.WatchEvent{
				Type:   kvstore.WatchEventType(r.Type),
				Object: r.Object,
			}
			select {
			case lw.OutCh <- &ev:
			case <-wstream.Context().Done():
				close(lw.OutCh)
				return
			}
		}
	}
	lw := listerwatcher.NewWatcherClient(wstream, bridgefn)
	lw.Run()
	return lw, nil
}

func (a *grpcObjCmdV1SmartNIC) Allowed(oper apiserver.APIOperType) bool {
	return true
}

type restObjCmdV1SmartNIC struct {
	endpoints cmd.EndpointsCmdV1RestClient
	instance  string
}

func (a *restObjCmdV1SmartNIC) Create(ctx context.Context, in *cmd.SmartNIC) (*cmd.SmartNIC, error) {
	if in == nil {
		return nil, errors.New("invalid input")
	}
	return a.endpoints.AutoAddSmartNIC(ctx, in)
}

func (a *restObjCmdV1SmartNIC) Update(ctx context.Context, in *cmd.SmartNIC) (*cmd.SmartNIC, error) {
	if in == nil {
		return nil, errors.New("invalid input")
	}
	return a.endpoints.AutoUpdateSmartNIC(ctx, in)
}

func (a *restObjCmdV1SmartNIC) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.SmartNIC, error) {
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.SmartNIC{}
	in.ObjectMeta = *objMeta
	return a.endpoints.AutoGetSmartNIC(ctx, &in)
}

func (a *restObjCmdV1SmartNIC) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cmd.SmartNIC, error) {
	if objMeta == nil {
		return nil, errors.New("invalid input")
	}
	in := cmd.SmartNIC{}
	in.ObjectMeta = *objMeta
	return a.endpoints.AutoDeleteSmartNIC(ctx, &in)
}

func (a *restObjCmdV1SmartNIC) List(ctx context.Context, options *api.ListWatchOptions) ([]*cmd.SmartNIC, error) {
	if options == nil {
		return nil, errors.New("invalid input")
	}
	r, err := a.endpoints.AutoListSmartNIC(ctx, options)
	if err == nil {
		return r.Items, nil
	}
	return nil, err
}

func (a *restObjCmdV1SmartNIC) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	return nil, errors.New("not allowed")
}

func (a *restObjCmdV1SmartNIC) Allowed(oper apiserver.APIOperType) bool {
	switch oper {
	case apiserver.CreateOper:
		return false
	case apiserver.UpdateOper:
		return true
	case apiserver.GetOper:
		return true
	case apiserver.DeleteOper:
		return true
	case apiserver.ListOper:
		return true
	case apiserver.WatchOper:
		return false
	default:
		return false
	}
}

type crudClientCmdV1 struct {
	grpcCluster  cmd.ClusterInterface
	grpcNode     cmd.NodeInterface
	grpcSmartNIC cmd.SmartNICInterface
}

// NewGrpcCrudClientCmdV1 creates a GRPC client for the service
func NewGrpcCrudClientCmdV1(conn *grpc.ClientConn, logger log.Logger) cmd.CmdV1Interface {
	client := NewCmdV1Backend(conn, logger)
	return &crudClientCmdV1{

		grpcCluster:  &grpcObjCmdV1Cluster{client: client, logger: logger},
		grpcNode:     &grpcObjCmdV1Node{client: client, logger: logger},
		grpcSmartNIC: &grpcObjCmdV1SmartNIC{client: client, logger: logger},
	}
}

func (a *crudClientCmdV1) Cluster() cmd.ClusterInterface {
	return a.grpcCluster
}

func (a *crudClientCmdV1) Node() cmd.NodeInterface {
	return a.grpcNode
}

func (a *crudClientCmdV1) SmartNIC() cmd.SmartNICInterface {
	return a.grpcSmartNIC
}

type crudRestClientCmdV1 struct {
	restCluster  cmd.ClusterInterface
	restNode     cmd.NodeInterface
	restSmartNIC cmd.SmartNICInterface
}

// NewRestCrudClientCmdV1 creates a REST client for the service.
func NewRestCrudClientCmdV1(url string) cmd.CmdV1Interface {
	endpoints, err := cmd.MakeCmdV1RestClientEndpoints(url)
	if err != nil {
		oldlog.Fatal("failed to create client")
	}
	return &crudRestClientCmdV1{

		restCluster:  &restObjCmdV1Cluster{endpoints: endpoints, instance: url},
		restNode:     &restObjCmdV1Node{endpoints: endpoints, instance: url},
		restSmartNIC: &restObjCmdV1SmartNIC{endpoints: endpoints, instance: url},
	}
}

func (a *crudRestClientCmdV1) Cluster() cmd.ClusterInterface {
	return a.restCluster
}

func (a *crudRestClientCmdV1) Node() cmd.NodeInterface {
	return a.restNode
}

func (a *crudRestClientCmdV1) SmartNIC() cmd.SmartNICInterface {
	return a.restSmartNIC
}
