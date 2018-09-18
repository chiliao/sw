// Code generated by MockGen. DO NOT EDIT.
// Source: tls_proxy_cb2.pb.go

// Package halproto is a generated GoMock package.
package halproto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// MockisTlsProxyCbKeyHandle_KeyOrHandle is a mock of isTlsProxyCbKeyHandle_KeyOrHandle interface
type MockisTlsProxyCbKeyHandle_KeyOrHandle struct {
	ctrl     *gomock.Controller
	recorder *MockisTlsProxyCbKeyHandle_KeyOrHandleMockRecorder
}

// MockisTlsProxyCbKeyHandle_KeyOrHandleMockRecorder is the mock recorder for MockisTlsProxyCbKeyHandle_KeyOrHandle
type MockisTlsProxyCbKeyHandle_KeyOrHandleMockRecorder struct {
	mock *MockisTlsProxyCbKeyHandle_KeyOrHandle
}

// NewMockisTlsProxyCbKeyHandle_KeyOrHandle creates a new mock instance
func NewMockisTlsProxyCbKeyHandle_KeyOrHandle(ctrl *gomock.Controller) *MockisTlsProxyCbKeyHandle_KeyOrHandle {
	mock := &MockisTlsProxyCbKeyHandle_KeyOrHandle{ctrl: ctrl}
	mock.recorder = &MockisTlsProxyCbKeyHandle_KeyOrHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisTlsProxyCbKeyHandle_KeyOrHandle) EXPECT() *MockisTlsProxyCbKeyHandle_KeyOrHandleMockRecorder {
	return m.recorder
}

// isTlsProxyCbKeyHandle_KeyOrHandle mocks base method
func (m *MockisTlsProxyCbKeyHandle_KeyOrHandle) isTlsProxyCbKeyHandle_KeyOrHandle() {
	m.ctrl.Call(m, "isTlsProxyCbKeyHandle_KeyOrHandle")
}

// isTlsProxyCbKeyHandle_KeyOrHandle indicates an expected call of isTlsProxyCbKeyHandle_KeyOrHandle
func (mr *MockisTlsProxyCbKeyHandle_KeyOrHandleMockRecorder) isTlsProxyCbKeyHandle_KeyOrHandle() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isTlsProxyCbKeyHandle_KeyOrHandle", reflect.TypeOf((*MockisTlsProxyCbKeyHandle_KeyOrHandle)(nil).isTlsProxyCbKeyHandle_KeyOrHandle))
}

// MarshalTo mocks base method
func (m *MockisTlsProxyCbKeyHandle_KeyOrHandle) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisTlsProxyCbKeyHandle_KeyOrHandleMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisTlsProxyCbKeyHandle_KeyOrHandle)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisTlsProxyCbKeyHandle_KeyOrHandle) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisTlsProxyCbKeyHandle_KeyOrHandleMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisTlsProxyCbKeyHandle_KeyOrHandle)(nil).Size))
}

// MockTlsProxyCbClient is a mock of TlsProxyCbClient interface
type MockTlsProxyCbClient struct {
	ctrl     *gomock.Controller
	recorder *MockTlsProxyCbClientMockRecorder
}

// MockTlsProxyCbClientMockRecorder is the mock recorder for MockTlsProxyCbClient
type MockTlsProxyCbClientMockRecorder struct {
	mock *MockTlsProxyCbClient
}

// NewMockTlsProxyCbClient creates a new mock instance
func NewMockTlsProxyCbClient(ctrl *gomock.Controller) *MockTlsProxyCbClient {
	mock := &MockTlsProxyCbClient{ctrl: ctrl}
	mock.recorder = &MockTlsProxyCbClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTlsProxyCbClient) EXPECT() *MockTlsProxyCbClientMockRecorder {
	return m.recorder
}

// TlsProxyCbCreate mocks base method
func (m *MockTlsProxyCbClient) TlsProxyCbCreate(ctx context.Context, in *TlsProxyCbRequestMsg, opts ...grpc.CallOption) (*TlsProxyCbResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TlsProxyCbCreate", varargs...)
	ret0, _ := ret[0].(*TlsProxyCbResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TlsProxyCbCreate indicates an expected call of TlsProxyCbCreate
func (mr *MockTlsProxyCbClientMockRecorder) TlsProxyCbCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsProxyCbCreate", reflect.TypeOf((*MockTlsProxyCbClient)(nil).TlsProxyCbCreate), varargs...)
}

// TlsProxyCbUpdate mocks base method
func (m *MockTlsProxyCbClient) TlsProxyCbUpdate(ctx context.Context, in *TlsProxyCbRequestMsg, opts ...grpc.CallOption) (*TlsProxyCbResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TlsProxyCbUpdate", varargs...)
	ret0, _ := ret[0].(*TlsProxyCbResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TlsProxyCbUpdate indicates an expected call of TlsProxyCbUpdate
func (mr *MockTlsProxyCbClientMockRecorder) TlsProxyCbUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsProxyCbUpdate", reflect.TypeOf((*MockTlsProxyCbClient)(nil).TlsProxyCbUpdate), varargs...)
}

// TlsProxyCbDelete mocks base method
func (m *MockTlsProxyCbClient) TlsProxyCbDelete(ctx context.Context, in *TlsProxyCbDeleteRequestMsg, opts ...grpc.CallOption) (*TlsProxyCbDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TlsProxyCbDelete", varargs...)
	ret0, _ := ret[0].(*TlsProxyCbDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TlsProxyCbDelete indicates an expected call of TlsProxyCbDelete
func (mr *MockTlsProxyCbClientMockRecorder) TlsProxyCbDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsProxyCbDelete", reflect.TypeOf((*MockTlsProxyCbClient)(nil).TlsProxyCbDelete), varargs...)
}

// TlsProxyCbGet mocks base method
func (m *MockTlsProxyCbClient) TlsProxyCbGet(ctx context.Context, in *TlsProxyCbGetRequestMsg, opts ...grpc.CallOption) (*TlsProxyCbGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TlsProxyCbGet", varargs...)
	ret0, _ := ret[0].(*TlsProxyCbGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TlsProxyCbGet indicates an expected call of TlsProxyCbGet
func (mr *MockTlsProxyCbClientMockRecorder) TlsProxyCbGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsProxyCbGet", reflect.TypeOf((*MockTlsProxyCbClient)(nil).TlsProxyCbGet), varargs...)
}

// MockTlsProxyCbServer is a mock of TlsProxyCbServer interface
type MockTlsProxyCbServer struct {
	ctrl     *gomock.Controller
	recorder *MockTlsProxyCbServerMockRecorder
}

// MockTlsProxyCbServerMockRecorder is the mock recorder for MockTlsProxyCbServer
type MockTlsProxyCbServerMockRecorder struct {
	mock *MockTlsProxyCbServer
}

// NewMockTlsProxyCbServer creates a new mock instance
func NewMockTlsProxyCbServer(ctrl *gomock.Controller) *MockTlsProxyCbServer {
	mock := &MockTlsProxyCbServer{ctrl: ctrl}
	mock.recorder = &MockTlsProxyCbServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTlsProxyCbServer) EXPECT() *MockTlsProxyCbServerMockRecorder {
	return m.recorder
}

// TlsProxyCbCreate mocks base method
func (m *MockTlsProxyCbServer) TlsProxyCbCreate(arg0 context.Context, arg1 *TlsProxyCbRequestMsg) (*TlsProxyCbResponseMsg, error) {
	ret := m.ctrl.Call(m, "TlsProxyCbCreate", arg0, arg1)
	ret0, _ := ret[0].(*TlsProxyCbResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TlsProxyCbCreate indicates an expected call of TlsProxyCbCreate
func (mr *MockTlsProxyCbServerMockRecorder) TlsProxyCbCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsProxyCbCreate", reflect.TypeOf((*MockTlsProxyCbServer)(nil).TlsProxyCbCreate), arg0, arg1)
}

// TlsProxyCbUpdate mocks base method
func (m *MockTlsProxyCbServer) TlsProxyCbUpdate(arg0 context.Context, arg1 *TlsProxyCbRequestMsg) (*TlsProxyCbResponseMsg, error) {
	ret := m.ctrl.Call(m, "TlsProxyCbUpdate", arg0, arg1)
	ret0, _ := ret[0].(*TlsProxyCbResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TlsProxyCbUpdate indicates an expected call of TlsProxyCbUpdate
func (mr *MockTlsProxyCbServerMockRecorder) TlsProxyCbUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsProxyCbUpdate", reflect.TypeOf((*MockTlsProxyCbServer)(nil).TlsProxyCbUpdate), arg0, arg1)
}

// TlsProxyCbDelete mocks base method
func (m *MockTlsProxyCbServer) TlsProxyCbDelete(arg0 context.Context, arg1 *TlsProxyCbDeleteRequestMsg) (*TlsProxyCbDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "TlsProxyCbDelete", arg0, arg1)
	ret0, _ := ret[0].(*TlsProxyCbDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TlsProxyCbDelete indicates an expected call of TlsProxyCbDelete
func (mr *MockTlsProxyCbServerMockRecorder) TlsProxyCbDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsProxyCbDelete", reflect.TypeOf((*MockTlsProxyCbServer)(nil).TlsProxyCbDelete), arg0, arg1)
}

// TlsProxyCbGet mocks base method
func (m *MockTlsProxyCbServer) TlsProxyCbGet(arg0 context.Context, arg1 *TlsProxyCbGetRequestMsg) (*TlsProxyCbGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "TlsProxyCbGet", arg0, arg1)
	ret0, _ := ret[0].(*TlsProxyCbGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TlsProxyCbGet indicates an expected call of TlsProxyCbGet
func (mr *MockTlsProxyCbServerMockRecorder) TlsProxyCbGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TlsProxyCbGet", reflect.TypeOf((*MockTlsProxyCbServer)(nil).TlsProxyCbGet), arg0, arg1)
}
