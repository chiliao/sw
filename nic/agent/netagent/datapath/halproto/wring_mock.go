// Code generated by MockGen. DO NOT EDIT.
// Source: wring.pb.go

// Package halproto is a generated GoMock package.
package halproto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// MockisWRingKeyHandle_KeyOrHandle is a mock of isWRingKeyHandle_KeyOrHandle interface
type MockisWRingKeyHandle_KeyOrHandle struct {
	ctrl     *gomock.Controller
	recorder *MockisWRingKeyHandle_KeyOrHandleMockRecorder
}

// MockisWRingKeyHandle_KeyOrHandleMockRecorder is the mock recorder for MockisWRingKeyHandle_KeyOrHandle
type MockisWRingKeyHandle_KeyOrHandleMockRecorder struct {
	mock *MockisWRingKeyHandle_KeyOrHandle
}

// NewMockisWRingKeyHandle_KeyOrHandle creates a new mock instance
func NewMockisWRingKeyHandle_KeyOrHandle(ctrl *gomock.Controller) *MockisWRingKeyHandle_KeyOrHandle {
	mock := &MockisWRingKeyHandle_KeyOrHandle{ctrl: ctrl}
	mock.recorder = &MockisWRingKeyHandle_KeyOrHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisWRingKeyHandle_KeyOrHandle) EXPECT() *MockisWRingKeyHandle_KeyOrHandleMockRecorder {
	return m.recorder
}

// isWRingKeyHandle_KeyOrHandle mocks base method
func (m *MockisWRingKeyHandle_KeyOrHandle) isWRingKeyHandle_KeyOrHandle() {
	m.ctrl.Call(m, "isWRingKeyHandle_KeyOrHandle")
}

// isWRingKeyHandle_KeyOrHandle indicates an expected call of isWRingKeyHandle_KeyOrHandle
func (mr *MockisWRingKeyHandle_KeyOrHandleMockRecorder) isWRingKeyHandle_KeyOrHandle() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isWRingKeyHandle_KeyOrHandle", reflect.TypeOf((*MockisWRingKeyHandle_KeyOrHandle)(nil).isWRingKeyHandle_KeyOrHandle))
}

// MarshalTo mocks base method
func (m *MockisWRingKeyHandle_KeyOrHandle) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisWRingKeyHandle_KeyOrHandleMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisWRingKeyHandle_KeyOrHandle)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisWRingKeyHandle_KeyOrHandle) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisWRingKeyHandle_KeyOrHandleMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisWRingKeyHandle_KeyOrHandle)(nil).Size))
}

// MockisWRingGetEntriesResponse_WRingSlotInfo is a mock of isWRingGetEntriesResponse_WRingSlotInfo interface
type MockisWRingGetEntriesResponse_WRingSlotInfo struct {
	ctrl     *gomock.Controller
	recorder *MockisWRingGetEntriesResponse_WRingSlotInfoMockRecorder
}

// MockisWRingGetEntriesResponse_WRingSlotInfoMockRecorder is the mock recorder for MockisWRingGetEntriesResponse_WRingSlotInfo
type MockisWRingGetEntriesResponse_WRingSlotInfoMockRecorder struct {
	mock *MockisWRingGetEntriesResponse_WRingSlotInfo
}

// NewMockisWRingGetEntriesResponse_WRingSlotInfo creates a new mock instance
func NewMockisWRingGetEntriesResponse_WRingSlotInfo(ctrl *gomock.Controller) *MockisWRingGetEntriesResponse_WRingSlotInfo {
	mock := &MockisWRingGetEntriesResponse_WRingSlotInfo{ctrl: ctrl}
	mock.recorder = &MockisWRingGetEntriesResponse_WRingSlotInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisWRingGetEntriesResponse_WRingSlotInfo) EXPECT() *MockisWRingGetEntriesResponse_WRingSlotInfoMockRecorder {
	return m.recorder
}

// isWRingGetEntriesResponse_WRingSlotInfo mocks base method
func (m *MockisWRingGetEntriesResponse_WRingSlotInfo) isWRingGetEntriesResponse_WRingSlotInfo() {
	m.ctrl.Call(m, "isWRingGetEntriesResponse_WRingSlotInfo")
}

// isWRingGetEntriesResponse_WRingSlotInfo indicates an expected call of isWRingGetEntriesResponse_WRingSlotInfo
func (mr *MockisWRingGetEntriesResponse_WRingSlotInfoMockRecorder) isWRingGetEntriesResponse_WRingSlotInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isWRingGetEntriesResponse_WRingSlotInfo", reflect.TypeOf((*MockisWRingGetEntriesResponse_WRingSlotInfo)(nil).isWRingGetEntriesResponse_WRingSlotInfo))
}

// MarshalTo mocks base method
func (m *MockisWRingGetEntriesResponse_WRingSlotInfo) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisWRingGetEntriesResponse_WRingSlotInfoMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisWRingGetEntriesResponse_WRingSlotInfo)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisWRingGetEntriesResponse_WRingSlotInfo) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisWRingGetEntriesResponse_WRingSlotInfoMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisWRingGetEntriesResponse_WRingSlotInfo)(nil).Size))
}

// MockWRingClient is a mock of WRingClient interface
type MockWRingClient struct {
	ctrl     *gomock.Controller
	recorder *MockWRingClientMockRecorder
}

// MockWRingClientMockRecorder is the mock recorder for MockWRingClient
type MockWRingClientMockRecorder struct {
	mock *MockWRingClient
}

// NewMockWRingClient creates a new mock instance
func NewMockWRingClient(ctrl *gomock.Controller) *MockWRingClient {
	mock := &MockWRingClient{ctrl: ctrl}
	mock.recorder = &MockWRingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWRingClient) EXPECT() *MockWRingClientMockRecorder {
	return m.recorder
}

// WRingCreate mocks base method
func (m *MockWRingClient) WRingCreate(ctx context.Context, in *WRingRequestMsg, opts ...grpc.CallOption) (*WRingResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WRingCreate", varargs...)
	ret0, _ := ret[0].(*WRingResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingCreate indicates an expected call of WRingCreate
func (mr *MockWRingClientMockRecorder) WRingCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingCreate", reflect.TypeOf((*MockWRingClient)(nil).WRingCreate), varargs...)
}

// WRingUpdate mocks base method
func (m *MockWRingClient) WRingUpdate(ctx context.Context, in *WRingRequestMsg, opts ...grpc.CallOption) (*WRingResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WRingUpdate", varargs...)
	ret0, _ := ret[0].(*WRingResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingUpdate indicates an expected call of WRingUpdate
func (mr *MockWRingClientMockRecorder) WRingUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingUpdate", reflect.TypeOf((*MockWRingClient)(nil).WRingUpdate), varargs...)
}

// WRingDelete mocks base method
func (m *MockWRingClient) WRingDelete(ctx context.Context, in *WRingDeleteRequestMsg, opts ...grpc.CallOption) (*WRingDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WRingDelete", varargs...)
	ret0, _ := ret[0].(*WRingDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingDelete indicates an expected call of WRingDelete
func (mr *MockWRingClientMockRecorder) WRingDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingDelete", reflect.TypeOf((*MockWRingClient)(nil).WRingDelete), varargs...)
}

// WRingGetEntries mocks base method
func (m *MockWRingClient) WRingGetEntries(ctx context.Context, in *WRingGetEntriesRequestMsg, opts ...grpc.CallOption) (*WRingGetEntriesResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WRingGetEntries", varargs...)
	ret0, _ := ret[0].(*WRingGetEntriesResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingGetEntries indicates an expected call of WRingGetEntries
func (mr *MockWRingClientMockRecorder) WRingGetEntries(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingGetEntries", reflect.TypeOf((*MockWRingClient)(nil).WRingGetEntries), varargs...)
}

// WRingGetMeta mocks base method
func (m *MockWRingClient) WRingGetMeta(ctx context.Context, in *WRingRequestMsg, opts ...grpc.CallOption) (*WRingGetMetaResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WRingGetMeta", varargs...)
	ret0, _ := ret[0].(*WRingGetMetaResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingGetMeta indicates an expected call of WRingGetMeta
func (mr *MockWRingClientMockRecorder) WRingGetMeta(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingGetMeta", reflect.TypeOf((*MockWRingClient)(nil).WRingGetMeta), varargs...)
}

// WRingSetMeta mocks base method
func (m *MockWRingClient) WRingSetMeta(ctx context.Context, in *WRingRequestMsg, opts ...grpc.CallOption) (*WRingSetMetaResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WRingSetMeta", varargs...)
	ret0, _ := ret[0].(*WRingSetMetaResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingSetMeta indicates an expected call of WRingSetMeta
func (mr *MockWRingClientMockRecorder) WRingSetMeta(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingSetMeta", reflect.TypeOf((*MockWRingClient)(nil).WRingSetMeta), varargs...)
}

// MockWRingServer is a mock of WRingServer interface
type MockWRingServer struct {
	ctrl     *gomock.Controller
	recorder *MockWRingServerMockRecorder
}

// MockWRingServerMockRecorder is the mock recorder for MockWRingServer
type MockWRingServerMockRecorder struct {
	mock *MockWRingServer
}

// NewMockWRingServer creates a new mock instance
func NewMockWRingServer(ctrl *gomock.Controller) *MockWRingServer {
	mock := &MockWRingServer{ctrl: ctrl}
	mock.recorder = &MockWRingServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWRingServer) EXPECT() *MockWRingServerMockRecorder {
	return m.recorder
}

// WRingCreate mocks base method
func (m *MockWRingServer) WRingCreate(arg0 context.Context, arg1 *WRingRequestMsg) (*WRingResponseMsg, error) {
	ret := m.ctrl.Call(m, "WRingCreate", arg0, arg1)
	ret0, _ := ret[0].(*WRingResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingCreate indicates an expected call of WRingCreate
func (mr *MockWRingServerMockRecorder) WRingCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingCreate", reflect.TypeOf((*MockWRingServer)(nil).WRingCreate), arg0, arg1)
}

// WRingUpdate mocks base method
func (m *MockWRingServer) WRingUpdate(arg0 context.Context, arg1 *WRingRequestMsg) (*WRingResponseMsg, error) {
	ret := m.ctrl.Call(m, "WRingUpdate", arg0, arg1)
	ret0, _ := ret[0].(*WRingResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingUpdate indicates an expected call of WRingUpdate
func (mr *MockWRingServerMockRecorder) WRingUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingUpdate", reflect.TypeOf((*MockWRingServer)(nil).WRingUpdate), arg0, arg1)
}

// WRingDelete mocks base method
func (m *MockWRingServer) WRingDelete(arg0 context.Context, arg1 *WRingDeleteRequestMsg) (*WRingDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "WRingDelete", arg0, arg1)
	ret0, _ := ret[0].(*WRingDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingDelete indicates an expected call of WRingDelete
func (mr *MockWRingServerMockRecorder) WRingDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingDelete", reflect.TypeOf((*MockWRingServer)(nil).WRingDelete), arg0, arg1)
}

// WRingGetEntries mocks base method
func (m *MockWRingServer) WRingGetEntries(arg0 context.Context, arg1 *WRingGetEntriesRequestMsg) (*WRingGetEntriesResponseMsg, error) {
	ret := m.ctrl.Call(m, "WRingGetEntries", arg0, arg1)
	ret0, _ := ret[0].(*WRingGetEntriesResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingGetEntries indicates an expected call of WRingGetEntries
func (mr *MockWRingServerMockRecorder) WRingGetEntries(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingGetEntries", reflect.TypeOf((*MockWRingServer)(nil).WRingGetEntries), arg0, arg1)
}

// WRingGetMeta mocks base method
func (m *MockWRingServer) WRingGetMeta(arg0 context.Context, arg1 *WRingRequestMsg) (*WRingGetMetaResponseMsg, error) {
	ret := m.ctrl.Call(m, "WRingGetMeta", arg0, arg1)
	ret0, _ := ret[0].(*WRingGetMetaResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingGetMeta indicates an expected call of WRingGetMeta
func (mr *MockWRingServerMockRecorder) WRingGetMeta(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingGetMeta", reflect.TypeOf((*MockWRingServer)(nil).WRingGetMeta), arg0, arg1)
}

// WRingSetMeta mocks base method
func (m *MockWRingServer) WRingSetMeta(arg0 context.Context, arg1 *WRingRequestMsg) (*WRingSetMetaResponseMsg, error) {
	ret := m.ctrl.Call(m, "WRingSetMeta", arg0, arg1)
	ret0, _ := ret[0].(*WRingSetMetaResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WRingSetMeta indicates an expected call of WRingSetMeta
func (mr *MockWRingServerMockRecorder) WRingSetMeta(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WRingSetMeta", reflect.TypeOf((*MockWRingServer)(nil).WRingSetMeta), arg0, arg1)
}
