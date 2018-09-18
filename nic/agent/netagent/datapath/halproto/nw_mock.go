// Code generated by MockGen. DO NOT EDIT.
// Source: nw.pb.go

// Package halproto is a generated GoMock package.
package halproto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// MockisNexthopSpec_IfOrEp is a mock of isNexthopSpec_IfOrEp interface
type MockisNexthopSpec_IfOrEp struct {
	ctrl     *gomock.Controller
	recorder *MockisNexthopSpec_IfOrEpMockRecorder
}

// MockisNexthopSpec_IfOrEpMockRecorder is the mock recorder for MockisNexthopSpec_IfOrEp
type MockisNexthopSpec_IfOrEpMockRecorder struct {
	mock *MockisNexthopSpec_IfOrEp
}

// NewMockisNexthopSpec_IfOrEp creates a new mock instance
func NewMockisNexthopSpec_IfOrEp(ctrl *gomock.Controller) *MockisNexthopSpec_IfOrEp {
	mock := &MockisNexthopSpec_IfOrEp{ctrl: ctrl}
	mock.recorder = &MockisNexthopSpec_IfOrEpMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisNexthopSpec_IfOrEp) EXPECT() *MockisNexthopSpec_IfOrEpMockRecorder {
	return m.recorder
}

// isNexthopSpec_IfOrEp mocks base method
func (m *MockisNexthopSpec_IfOrEp) isNexthopSpec_IfOrEp() {
	m.ctrl.Call(m, "isNexthopSpec_IfOrEp")
}

// isNexthopSpec_IfOrEp indicates an expected call of isNexthopSpec_IfOrEp
func (mr *MockisNexthopSpec_IfOrEpMockRecorder) isNexthopSpec_IfOrEp() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isNexthopSpec_IfOrEp", reflect.TypeOf((*MockisNexthopSpec_IfOrEp)(nil).isNexthopSpec_IfOrEp))
}

// MarshalTo mocks base method
func (m *MockisNexthopSpec_IfOrEp) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisNexthopSpec_IfOrEpMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisNexthopSpec_IfOrEp)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisNexthopSpec_IfOrEp) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisNexthopSpec_IfOrEpMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisNexthopSpec_IfOrEp)(nil).Size))
}

// MockNetworkClient is a mock of NetworkClient interface
type MockNetworkClient struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkClientMockRecorder
}

// MockNetworkClientMockRecorder is the mock recorder for MockNetworkClient
type MockNetworkClientMockRecorder struct {
	mock *MockNetworkClient
}

// NewMockNetworkClient creates a new mock instance
func NewMockNetworkClient(ctrl *gomock.Controller) *MockNetworkClient {
	mock := &MockNetworkClient{ctrl: ctrl}
	mock.recorder = &MockNetworkClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkClient) EXPECT() *MockNetworkClientMockRecorder {
	return m.recorder
}

// NetworkCreate mocks base method
func (m *MockNetworkClient) NetworkCreate(ctx context.Context, in *NetworkRequestMsg, opts ...grpc.CallOption) (*NetworkResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NetworkCreate", varargs...)
	ret0, _ := ret[0].(*NetworkResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkCreate indicates an expected call of NetworkCreate
func (mr *MockNetworkClientMockRecorder) NetworkCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkCreate", reflect.TypeOf((*MockNetworkClient)(nil).NetworkCreate), varargs...)
}

// NetworkUpdate mocks base method
func (m *MockNetworkClient) NetworkUpdate(ctx context.Context, in *NetworkRequestMsg, opts ...grpc.CallOption) (*NetworkResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NetworkUpdate", varargs...)
	ret0, _ := ret[0].(*NetworkResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkUpdate indicates an expected call of NetworkUpdate
func (mr *MockNetworkClientMockRecorder) NetworkUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkUpdate", reflect.TypeOf((*MockNetworkClient)(nil).NetworkUpdate), varargs...)
}

// NetworkDelete mocks base method
func (m *MockNetworkClient) NetworkDelete(ctx context.Context, in *NetworkDeleteRequestMsg, opts ...grpc.CallOption) (*NetworkDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NetworkDelete", varargs...)
	ret0, _ := ret[0].(*NetworkDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkDelete indicates an expected call of NetworkDelete
func (mr *MockNetworkClientMockRecorder) NetworkDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkDelete", reflect.TypeOf((*MockNetworkClient)(nil).NetworkDelete), varargs...)
}

// NetworkGet mocks base method
func (m *MockNetworkClient) NetworkGet(ctx context.Context, in *NetworkGetRequestMsg, opts ...grpc.CallOption) (*NetworkGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NetworkGet", varargs...)
	ret0, _ := ret[0].(*NetworkGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkGet indicates an expected call of NetworkGet
func (mr *MockNetworkClientMockRecorder) NetworkGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkGet", reflect.TypeOf((*MockNetworkClient)(nil).NetworkGet), varargs...)
}

// NexthopCreate mocks base method
func (m *MockNetworkClient) NexthopCreate(ctx context.Context, in *NexthopRequestMsg, opts ...grpc.CallOption) (*NexthopResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NexthopCreate", varargs...)
	ret0, _ := ret[0].(*NexthopResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NexthopCreate indicates an expected call of NexthopCreate
func (mr *MockNetworkClientMockRecorder) NexthopCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NexthopCreate", reflect.TypeOf((*MockNetworkClient)(nil).NexthopCreate), varargs...)
}

// NexthopUpdate mocks base method
func (m *MockNetworkClient) NexthopUpdate(ctx context.Context, in *NexthopRequestMsg, opts ...grpc.CallOption) (*NexthopResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NexthopUpdate", varargs...)
	ret0, _ := ret[0].(*NexthopResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NexthopUpdate indicates an expected call of NexthopUpdate
func (mr *MockNetworkClientMockRecorder) NexthopUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NexthopUpdate", reflect.TypeOf((*MockNetworkClient)(nil).NexthopUpdate), varargs...)
}

// NexthopDelete mocks base method
func (m *MockNetworkClient) NexthopDelete(ctx context.Context, in *NexthopDeleteRequestMsg, opts ...grpc.CallOption) (*NexthopDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NexthopDelete", varargs...)
	ret0, _ := ret[0].(*NexthopDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NexthopDelete indicates an expected call of NexthopDelete
func (mr *MockNetworkClientMockRecorder) NexthopDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NexthopDelete", reflect.TypeOf((*MockNetworkClient)(nil).NexthopDelete), varargs...)
}

// NexthopGet mocks base method
func (m *MockNetworkClient) NexthopGet(ctx context.Context, in *NexthopGetRequestMsg, opts ...grpc.CallOption) (*NexthopGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "NexthopGet", varargs...)
	ret0, _ := ret[0].(*NexthopGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NexthopGet indicates an expected call of NexthopGet
func (mr *MockNetworkClientMockRecorder) NexthopGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NexthopGet", reflect.TypeOf((*MockNetworkClient)(nil).NexthopGet), varargs...)
}

// RouteCreate mocks base method
func (m *MockNetworkClient) RouteCreate(ctx context.Context, in *RouteRequestMsg, opts ...grpc.CallOption) (*RouteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RouteCreate", varargs...)
	ret0, _ := ret[0].(*RouteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteCreate indicates an expected call of RouteCreate
func (mr *MockNetworkClientMockRecorder) RouteCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteCreate", reflect.TypeOf((*MockNetworkClient)(nil).RouteCreate), varargs...)
}

// RouteUpdate mocks base method
func (m *MockNetworkClient) RouteUpdate(ctx context.Context, in *RouteRequestMsg, opts ...grpc.CallOption) (*RouteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RouteUpdate", varargs...)
	ret0, _ := ret[0].(*RouteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteUpdate indicates an expected call of RouteUpdate
func (mr *MockNetworkClientMockRecorder) RouteUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteUpdate", reflect.TypeOf((*MockNetworkClient)(nil).RouteUpdate), varargs...)
}

// RouteDelete mocks base method
func (m *MockNetworkClient) RouteDelete(ctx context.Context, in *RouteDeleteRequestMsg, opts ...grpc.CallOption) (*RouteDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RouteDelete", varargs...)
	ret0, _ := ret[0].(*RouteDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteDelete indicates an expected call of RouteDelete
func (mr *MockNetworkClientMockRecorder) RouteDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteDelete", reflect.TypeOf((*MockNetworkClient)(nil).RouteDelete), varargs...)
}

// RouteGet mocks base method
func (m *MockNetworkClient) RouteGet(ctx context.Context, in *RouteGetRequestMsg, opts ...grpc.CallOption) (*RouteGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RouteGet", varargs...)
	ret0, _ := ret[0].(*RouteGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteGet indicates an expected call of RouteGet
func (mr *MockNetworkClientMockRecorder) RouteGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteGet", reflect.TypeOf((*MockNetworkClient)(nil).RouteGet), varargs...)
}

// MockNetworkServer is a mock of NetworkServer interface
type MockNetworkServer struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkServerMockRecorder
}

// MockNetworkServerMockRecorder is the mock recorder for MockNetworkServer
type MockNetworkServerMockRecorder struct {
	mock *MockNetworkServer
}

// NewMockNetworkServer creates a new mock instance
func NewMockNetworkServer(ctrl *gomock.Controller) *MockNetworkServer {
	mock := &MockNetworkServer{ctrl: ctrl}
	mock.recorder = &MockNetworkServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNetworkServer) EXPECT() *MockNetworkServerMockRecorder {
	return m.recorder
}

// NetworkCreate mocks base method
func (m *MockNetworkServer) NetworkCreate(arg0 context.Context, arg1 *NetworkRequestMsg) (*NetworkResponseMsg, error) {
	ret := m.ctrl.Call(m, "NetworkCreate", arg0, arg1)
	ret0, _ := ret[0].(*NetworkResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkCreate indicates an expected call of NetworkCreate
func (mr *MockNetworkServerMockRecorder) NetworkCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkCreate", reflect.TypeOf((*MockNetworkServer)(nil).NetworkCreate), arg0, arg1)
}

// NetworkUpdate mocks base method
func (m *MockNetworkServer) NetworkUpdate(arg0 context.Context, arg1 *NetworkRequestMsg) (*NetworkResponseMsg, error) {
	ret := m.ctrl.Call(m, "NetworkUpdate", arg0, arg1)
	ret0, _ := ret[0].(*NetworkResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkUpdate indicates an expected call of NetworkUpdate
func (mr *MockNetworkServerMockRecorder) NetworkUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkUpdate", reflect.TypeOf((*MockNetworkServer)(nil).NetworkUpdate), arg0, arg1)
}

// NetworkDelete mocks base method
func (m *MockNetworkServer) NetworkDelete(arg0 context.Context, arg1 *NetworkDeleteRequestMsg) (*NetworkDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "NetworkDelete", arg0, arg1)
	ret0, _ := ret[0].(*NetworkDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkDelete indicates an expected call of NetworkDelete
func (mr *MockNetworkServerMockRecorder) NetworkDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkDelete", reflect.TypeOf((*MockNetworkServer)(nil).NetworkDelete), arg0, arg1)
}

// NetworkGet mocks base method
func (m *MockNetworkServer) NetworkGet(arg0 context.Context, arg1 *NetworkGetRequestMsg) (*NetworkGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "NetworkGet", arg0, arg1)
	ret0, _ := ret[0].(*NetworkGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NetworkGet indicates an expected call of NetworkGet
func (mr *MockNetworkServerMockRecorder) NetworkGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NetworkGet", reflect.TypeOf((*MockNetworkServer)(nil).NetworkGet), arg0, arg1)
}

// NexthopCreate mocks base method
func (m *MockNetworkServer) NexthopCreate(arg0 context.Context, arg1 *NexthopRequestMsg) (*NexthopResponseMsg, error) {
	ret := m.ctrl.Call(m, "NexthopCreate", arg0, arg1)
	ret0, _ := ret[0].(*NexthopResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NexthopCreate indicates an expected call of NexthopCreate
func (mr *MockNetworkServerMockRecorder) NexthopCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NexthopCreate", reflect.TypeOf((*MockNetworkServer)(nil).NexthopCreate), arg0, arg1)
}

// NexthopUpdate mocks base method
func (m *MockNetworkServer) NexthopUpdate(arg0 context.Context, arg1 *NexthopRequestMsg) (*NexthopResponseMsg, error) {
	ret := m.ctrl.Call(m, "NexthopUpdate", arg0, arg1)
	ret0, _ := ret[0].(*NexthopResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NexthopUpdate indicates an expected call of NexthopUpdate
func (mr *MockNetworkServerMockRecorder) NexthopUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NexthopUpdate", reflect.TypeOf((*MockNetworkServer)(nil).NexthopUpdate), arg0, arg1)
}

// NexthopDelete mocks base method
func (m *MockNetworkServer) NexthopDelete(arg0 context.Context, arg1 *NexthopDeleteRequestMsg) (*NexthopDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "NexthopDelete", arg0, arg1)
	ret0, _ := ret[0].(*NexthopDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NexthopDelete indicates an expected call of NexthopDelete
func (mr *MockNetworkServerMockRecorder) NexthopDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NexthopDelete", reflect.TypeOf((*MockNetworkServer)(nil).NexthopDelete), arg0, arg1)
}

// NexthopGet mocks base method
func (m *MockNetworkServer) NexthopGet(arg0 context.Context, arg1 *NexthopGetRequestMsg) (*NexthopGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "NexthopGet", arg0, arg1)
	ret0, _ := ret[0].(*NexthopGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NexthopGet indicates an expected call of NexthopGet
func (mr *MockNetworkServerMockRecorder) NexthopGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NexthopGet", reflect.TypeOf((*MockNetworkServer)(nil).NexthopGet), arg0, arg1)
}

// RouteCreate mocks base method
func (m *MockNetworkServer) RouteCreate(arg0 context.Context, arg1 *RouteRequestMsg) (*RouteResponseMsg, error) {
	ret := m.ctrl.Call(m, "RouteCreate", arg0, arg1)
	ret0, _ := ret[0].(*RouteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteCreate indicates an expected call of RouteCreate
func (mr *MockNetworkServerMockRecorder) RouteCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteCreate", reflect.TypeOf((*MockNetworkServer)(nil).RouteCreate), arg0, arg1)
}

// RouteUpdate mocks base method
func (m *MockNetworkServer) RouteUpdate(arg0 context.Context, arg1 *RouteRequestMsg) (*RouteResponseMsg, error) {
	ret := m.ctrl.Call(m, "RouteUpdate", arg0, arg1)
	ret0, _ := ret[0].(*RouteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteUpdate indicates an expected call of RouteUpdate
func (mr *MockNetworkServerMockRecorder) RouteUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteUpdate", reflect.TypeOf((*MockNetworkServer)(nil).RouteUpdate), arg0, arg1)
}

// RouteDelete mocks base method
func (m *MockNetworkServer) RouteDelete(arg0 context.Context, arg1 *RouteDeleteRequestMsg) (*RouteDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "RouteDelete", arg0, arg1)
	ret0, _ := ret[0].(*RouteDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteDelete indicates an expected call of RouteDelete
func (mr *MockNetworkServerMockRecorder) RouteDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteDelete", reflect.TypeOf((*MockNetworkServer)(nil).RouteDelete), arg0, arg1)
}

// RouteGet mocks base method
func (m *MockNetworkServer) RouteGet(arg0 context.Context, arg1 *RouteGetRequestMsg) (*RouteGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "RouteGet", arg0, arg1)
	ret0, _ := ret[0].(*RouteGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RouteGet indicates an expected call of RouteGet
func (mr *MockNetworkServerMockRecorder) RouteGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RouteGet", reflect.TypeOf((*MockNetworkServer)(nil).RouteGet), arg0, arg1)
}
