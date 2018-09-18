// Code generated by MockGen. DO NOT EDIT.
// Source: ipsec.pb.go

// Package halproto is a generated GoMock package.
package halproto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// MockisKey_KeyInfo is a mock of isKey_KeyInfo interface
type MockisKey_KeyInfo struct {
	ctrl     *gomock.Controller
	recorder *MockisKey_KeyInfoMockRecorder
}

// MockisKey_KeyInfoMockRecorder is the mock recorder for MockisKey_KeyInfo
type MockisKey_KeyInfoMockRecorder struct {
	mock *MockisKey_KeyInfo
}

// NewMockisKey_KeyInfo creates a new mock instance
func NewMockisKey_KeyInfo(ctrl *gomock.Controller) *MockisKey_KeyInfo {
	mock := &MockisKey_KeyInfo{ctrl: ctrl}
	mock.recorder = &MockisKey_KeyInfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisKey_KeyInfo) EXPECT() *MockisKey_KeyInfoMockRecorder {
	return m.recorder
}

// isKey_KeyInfo mocks base method
func (m *MockisKey_KeyInfo) isKey_KeyInfo() {
	m.ctrl.Call(m, "isKey_KeyInfo")
}

// isKey_KeyInfo indicates an expected call of isKey_KeyInfo
func (mr *MockisKey_KeyInfoMockRecorder) isKey_KeyInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isKey_KeyInfo", reflect.TypeOf((*MockisKey_KeyInfo)(nil).isKey_KeyInfo))
}

// MarshalTo mocks base method
func (m *MockisKey_KeyInfo) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisKey_KeyInfoMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisKey_KeyInfo)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisKey_KeyInfo) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisKey_KeyInfoMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisKey_KeyInfo)(nil).Size))
}

// MockisIpsecSAAction_SaHandle is a mock of isIpsecSAAction_SaHandle interface
type MockisIpsecSAAction_SaHandle struct {
	ctrl     *gomock.Controller
	recorder *MockisIpsecSAAction_SaHandleMockRecorder
}

// MockisIpsecSAAction_SaHandleMockRecorder is the mock recorder for MockisIpsecSAAction_SaHandle
type MockisIpsecSAAction_SaHandleMockRecorder struct {
	mock *MockisIpsecSAAction_SaHandle
}

// NewMockisIpsecSAAction_SaHandle creates a new mock instance
func NewMockisIpsecSAAction_SaHandle(ctrl *gomock.Controller) *MockisIpsecSAAction_SaHandle {
	mock := &MockisIpsecSAAction_SaHandle{ctrl: ctrl}
	mock.recorder = &MockisIpsecSAAction_SaHandleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisIpsecSAAction_SaHandle) EXPECT() *MockisIpsecSAAction_SaHandleMockRecorder {
	return m.recorder
}

// isIpsecSAAction_SaHandle mocks base method
func (m *MockisIpsecSAAction_SaHandle) isIpsecSAAction_SaHandle() {
	m.ctrl.Call(m, "isIpsecSAAction_SaHandle")
}

// isIpsecSAAction_SaHandle indicates an expected call of isIpsecSAAction_SaHandle
func (mr *MockisIpsecSAAction_SaHandleMockRecorder) isIpsecSAAction_SaHandle() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isIpsecSAAction_SaHandle", reflect.TypeOf((*MockisIpsecSAAction_SaHandle)(nil).isIpsecSAAction_SaHandle))
}

// MarshalTo mocks base method
func (m *MockisIpsecSAAction_SaHandle) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisIpsecSAAction_SaHandleMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisIpsecSAAction_SaHandle)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisIpsecSAAction_SaHandle) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisIpsecSAAction_SaHandleMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisIpsecSAAction_SaHandle)(nil).Size))
}

// MockIpsecClient is a mock of IpsecClient interface
type MockIpsecClient struct {
	ctrl     *gomock.Controller
	recorder *MockIpsecClientMockRecorder
}

// MockIpsecClientMockRecorder is the mock recorder for MockIpsecClient
type MockIpsecClientMockRecorder struct {
	mock *MockIpsecClient
}

// NewMockIpsecClient creates a new mock instance
func NewMockIpsecClient(ctrl *gomock.Controller) *MockIpsecClient {
	mock := &MockIpsecClient{ctrl: ctrl}
	mock.recorder = &MockIpsecClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIpsecClient) EXPECT() *MockIpsecClientMockRecorder {
	return m.recorder
}

// IpsecRuleCreate mocks base method
func (m *MockIpsecClient) IpsecRuleCreate(ctx context.Context, in *IpsecRuleRequestMsg, opts ...grpc.CallOption) (*IpsecRuleResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecRuleCreate", varargs...)
	ret0, _ := ret[0].(*IpsecRuleResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecRuleCreate indicates an expected call of IpsecRuleCreate
func (mr *MockIpsecClientMockRecorder) IpsecRuleCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecRuleCreate", reflect.TypeOf((*MockIpsecClient)(nil).IpsecRuleCreate), varargs...)
}

// IpsecRuleUpdate mocks base method
func (m *MockIpsecClient) IpsecRuleUpdate(ctx context.Context, in *IpsecRuleRequestMsg, opts ...grpc.CallOption) (*IpsecRuleResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecRuleUpdate", varargs...)
	ret0, _ := ret[0].(*IpsecRuleResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecRuleUpdate indicates an expected call of IpsecRuleUpdate
func (mr *MockIpsecClientMockRecorder) IpsecRuleUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecRuleUpdate", reflect.TypeOf((*MockIpsecClient)(nil).IpsecRuleUpdate), varargs...)
}

// IpsecRuleDelete mocks base method
func (m *MockIpsecClient) IpsecRuleDelete(ctx context.Context, in *IpsecRuleDeleteRequestMsg, opts ...grpc.CallOption) (*IpsecRuleDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecRuleDelete", varargs...)
	ret0, _ := ret[0].(*IpsecRuleDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecRuleDelete indicates an expected call of IpsecRuleDelete
func (mr *MockIpsecClientMockRecorder) IpsecRuleDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecRuleDelete", reflect.TypeOf((*MockIpsecClient)(nil).IpsecRuleDelete), varargs...)
}

// IpsecRuleGet mocks base method
func (m *MockIpsecClient) IpsecRuleGet(ctx context.Context, in *IpsecRuleGetRequestMsg, opts ...grpc.CallOption) (*IpsecRuleGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecRuleGet", varargs...)
	ret0, _ := ret[0].(*IpsecRuleGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecRuleGet indicates an expected call of IpsecRuleGet
func (mr *MockIpsecClientMockRecorder) IpsecRuleGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecRuleGet", reflect.TypeOf((*MockIpsecClient)(nil).IpsecRuleGet), varargs...)
}

// IpsecSAEncryptCreate mocks base method
func (m *MockIpsecClient) IpsecSAEncryptCreate(ctx context.Context, in *IpsecSAEncryptRequestMsg, opts ...grpc.CallOption) (*IpsecSAEncryptResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecSAEncryptCreate", varargs...)
	ret0, _ := ret[0].(*IpsecSAEncryptResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSAEncryptCreate indicates an expected call of IpsecSAEncryptCreate
func (mr *MockIpsecClientMockRecorder) IpsecSAEncryptCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSAEncryptCreate", reflect.TypeOf((*MockIpsecClient)(nil).IpsecSAEncryptCreate), varargs...)
}

// IpsecSAEncryptUpdate mocks base method
func (m *MockIpsecClient) IpsecSAEncryptUpdate(ctx context.Context, in *IpsecSAEncryptRequestMsg, opts ...grpc.CallOption) (*IpsecSAEncryptResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecSAEncryptUpdate", varargs...)
	ret0, _ := ret[0].(*IpsecSAEncryptResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSAEncryptUpdate indicates an expected call of IpsecSAEncryptUpdate
func (mr *MockIpsecClientMockRecorder) IpsecSAEncryptUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSAEncryptUpdate", reflect.TypeOf((*MockIpsecClient)(nil).IpsecSAEncryptUpdate), varargs...)
}

// IpsecSAEncryptDelete mocks base method
func (m *MockIpsecClient) IpsecSAEncryptDelete(ctx context.Context, in *IpsecSAEncryptDeleteRequestMsg, opts ...grpc.CallOption) (*IpsecSAEncryptDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecSAEncryptDelete", varargs...)
	ret0, _ := ret[0].(*IpsecSAEncryptDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSAEncryptDelete indicates an expected call of IpsecSAEncryptDelete
func (mr *MockIpsecClientMockRecorder) IpsecSAEncryptDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSAEncryptDelete", reflect.TypeOf((*MockIpsecClient)(nil).IpsecSAEncryptDelete), varargs...)
}

// IpsecSAEncryptGet mocks base method
func (m *MockIpsecClient) IpsecSAEncryptGet(ctx context.Context, in *IpsecSAEncryptGetRequestMsg, opts ...grpc.CallOption) (*IpsecSAEncryptGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecSAEncryptGet", varargs...)
	ret0, _ := ret[0].(*IpsecSAEncryptGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSAEncryptGet indicates an expected call of IpsecSAEncryptGet
func (mr *MockIpsecClientMockRecorder) IpsecSAEncryptGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSAEncryptGet", reflect.TypeOf((*MockIpsecClient)(nil).IpsecSAEncryptGet), varargs...)
}

// IpsecSADecryptCreate mocks base method
func (m *MockIpsecClient) IpsecSADecryptCreate(ctx context.Context, in *IpsecSADecryptRequestMsg, opts ...grpc.CallOption) (*IpsecSADecryptResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecSADecryptCreate", varargs...)
	ret0, _ := ret[0].(*IpsecSADecryptResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSADecryptCreate indicates an expected call of IpsecSADecryptCreate
func (mr *MockIpsecClientMockRecorder) IpsecSADecryptCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSADecryptCreate", reflect.TypeOf((*MockIpsecClient)(nil).IpsecSADecryptCreate), varargs...)
}

// IpsecSADecryptUpdate mocks base method
func (m *MockIpsecClient) IpsecSADecryptUpdate(ctx context.Context, in *IpsecSADecryptRequestMsg, opts ...grpc.CallOption) (*IpsecSADecryptResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecSADecryptUpdate", varargs...)
	ret0, _ := ret[0].(*IpsecSADecryptResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSADecryptUpdate indicates an expected call of IpsecSADecryptUpdate
func (mr *MockIpsecClientMockRecorder) IpsecSADecryptUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSADecryptUpdate", reflect.TypeOf((*MockIpsecClient)(nil).IpsecSADecryptUpdate), varargs...)
}

// IpsecSADecryptDelete mocks base method
func (m *MockIpsecClient) IpsecSADecryptDelete(ctx context.Context, in *IpsecSADecryptDeleteRequestMsg, opts ...grpc.CallOption) (*IpsecSADecryptDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecSADecryptDelete", varargs...)
	ret0, _ := ret[0].(*IpsecSADecryptDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSADecryptDelete indicates an expected call of IpsecSADecryptDelete
func (mr *MockIpsecClientMockRecorder) IpsecSADecryptDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSADecryptDelete", reflect.TypeOf((*MockIpsecClient)(nil).IpsecSADecryptDelete), varargs...)
}

// IpsecSADecryptGet mocks base method
func (m *MockIpsecClient) IpsecSADecryptGet(ctx context.Context, in *IpsecSADecryptGetRequestMsg, opts ...grpc.CallOption) (*IpsecSADecryptGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IpsecSADecryptGet", varargs...)
	ret0, _ := ret[0].(*IpsecSADecryptGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSADecryptGet indicates an expected call of IpsecSADecryptGet
func (mr *MockIpsecClientMockRecorder) IpsecSADecryptGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSADecryptGet", reflect.TypeOf((*MockIpsecClient)(nil).IpsecSADecryptGet), varargs...)
}

// MockIpsecServer is a mock of IpsecServer interface
type MockIpsecServer struct {
	ctrl     *gomock.Controller
	recorder *MockIpsecServerMockRecorder
}

// MockIpsecServerMockRecorder is the mock recorder for MockIpsecServer
type MockIpsecServerMockRecorder struct {
	mock *MockIpsecServer
}

// NewMockIpsecServer creates a new mock instance
func NewMockIpsecServer(ctrl *gomock.Controller) *MockIpsecServer {
	mock := &MockIpsecServer{ctrl: ctrl}
	mock.recorder = &MockIpsecServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockIpsecServer) EXPECT() *MockIpsecServerMockRecorder {
	return m.recorder
}

// IpsecRuleCreate mocks base method
func (m *MockIpsecServer) IpsecRuleCreate(arg0 context.Context, arg1 *IpsecRuleRequestMsg) (*IpsecRuleResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecRuleCreate", arg0, arg1)
	ret0, _ := ret[0].(*IpsecRuleResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecRuleCreate indicates an expected call of IpsecRuleCreate
func (mr *MockIpsecServerMockRecorder) IpsecRuleCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecRuleCreate", reflect.TypeOf((*MockIpsecServer)(nil).IpsecRuleCreate), arg0, arg1)
}

// IpsecRuleUpdate mocks base method
func (m *MockIpsecServer) IpsecRuleUpdate(arg0 context.Context, arg1 *IpsecRuleRequestMsg) (*IpsecRuleResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecRuleUpdate", arg0, arg1)
	ret0, _ := ret[0].(*IpsecRuleResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecRuleUpdate indicates an expected call of IpsecRuleUpdate
func (mr *MockIpsecServerMockRecorder) IpsecRuleUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecRuleUpdate", reflect.TypeOf((*MockIpsecServer)(nil).IpsecRuleUpdate), arg0, arg1)
}

// IpsecRuleDelete mocks base method
func (m *MockIpsecServer) IpsecRuleDelete(arg0 context.Context, arg1 *IpsecRuleDeleteRequestMsg) (*IpsecRuleDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecRuleDelete", arg0, arg1)
	ret0, _ := ret[0].(*IpsecRuleDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecRuleDelete indicates an expected call of IpsecRuleDelete
func (mr *MockIpsecServerMockRecorder) IpsecRuleDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecRuleDelete", reflect.TypeOf((*MockIpsecServer)(nil).IpsecRuleDelete), arg0, arg1)
}

// IpsecRuleGet mocks base method
func (m *MockIpsecServer) IpsecRuleGet(arg0 context.Context, arg1 *IpsecRuleGetRequestMsg) (*IpsecRuleGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecRuleGet", arg0, arg1)
	ret0, _ := ret[0].(*IpsecRuleGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecRuleGet indicates an expected call of IpsecRuleGet
func (mr *MockIpsecServerMockRecorder) IpsecRuleGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecRuleGet", reflect.TypeOf((*MockIpsecServer)(nil).IpsecRuleGet), arg0, arg1)
}

// IpsecSAEncryptCreate mocks base method
func (m *MockIpsecServer) IpsecSAEncryptCreate(arg0 context.Context, arg1 *IpsecSAEncryptRequestMsg) (*IpsecSAEncryptResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecSAEncryptCreate", arg0, arg1)
	ret0, _ := ret[0].(*IpsecSAEncryptResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSAEncryptCreate indicates an expected call of IpsecSAEncryptCreate
func (mr *MockIpsecServerMockRecorder) IpsecSAEncryptCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSAEncryptCreate", reflect.TypeOf((*MockIpsecServer)(nil).IpsecSAEncryptCreate), arg0, arg1)
}

// IpsecSAEncryptUpdate mocks base method
func (m *MockIpsecServer) IpsecSAEncryptUpdate(arg0 context.Context, arg1 *IpsecSAEncryptRequestMsg) (*IpsecSAEncryptResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecSAEncryptUpdate", arg0, arg1)
	ret0, _ := ret[0].(*IpsecSAEncryptResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSAEncryptUpdate indicates an expected call of IpsecSAEncryptUpdate
func (mr *MockIpsecServerMockRecorder) IpsecSAEncryptUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSAEncryptUpdate", reflect.TypeOf((*MockIpsecServer)(nil).IpsecSAEncryptUpdate), arg0, arg1)
}

// IpsecSAEncryptDelete mocks base method
func (m *MockIpsecServer) IpsecSAEncryptDelete(arg0 context.Context, arg1 *IpsecSAEncryptDeleteRequestMsg) (*IpsecSAEncryptDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecSAEncryptDelete", arg0, arg1)
	ret0, _ := ret[0].(*IpsecSAEncryptDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSAEncryptDelete indicates an expected call of IpsecSAEncryptDelete
func (mr *MockIpsecServerMockRecorder) IpsecSAEncryptDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSAEncryptDelete", reflect.TypeOf((*MockIpsecServer)(nil).IpsecSAEncryptDelete), arg0, arg1)
}

// IpsecSAEncryptGet mocks base method
func (m *MockIpsecServer) IpsecSAEncryptGet(arg0 context.Context, arg1 *IpsecSAEncryptGetRequestMsg) (*IpsecSAEncryptGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecSAEncryptGet", arg0, arg1)
	ret0, _ := ret[0].(*IpsecSAEncryptGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSAEncryptGet indicates an expected call of IpsecSAEncryptGet
func (mr *MockIpsecServerMockRecorder) IpsecSAEncryptGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSAEncryptGet", reflect.TypeOf((*MockIpsecServer)(nil).IpsecSAEncryptGet), arg0, arg1)
}

// IpsecSADecryptCreate mocks base method
func (m *MockIpsecServer) IpsecSADecryptCreate(arg0 context.Context, arg1 *IpsecSADecryptRequestMsg) (*IpsecSADecryptResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecSADecryptCreate", arg0, arg1)
	ret0, _ := ret[0].(*IpsecSADecryptResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSADecryptCreate indicates an expected call of IpsecSADecryptCreate
func (mr *MockIpsecServerMockRecorder) IpsecSADecryptCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSADecryptCreate", reflect.TypeOf((*MockIpsecServer)(nil).IpsecSADecryptCreate), arg0, arg1)
}

// IpsecSADecryptUpdate mocks base method
func (m *MockIpsecServer) IpsecSADecryptUpdate(arg0 context.Context, arg1 *IpsecSADecryptRequestMsg) (*IpsecSADecryptResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecSADecryptUpdate", arg0, arg1)
	ret0, _ := ret[0].(*IpsecSADecryptResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSADecryptUpdate indicates an expected call of IpsecSADecryptUpdate
func (mr *MockIpsecServerMockRecorder) IpsecSADecryptUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSADecryptUpdate", reflect.TypeOf((*MockIpsecServer)(nil).IpsecSADecryptUpdate), arg0, arg1)
}

// IpsecSADecryptDelete mocks base method
func (m *MockIpsecServer) IpsecSADecryptDelete(arg0 context.Context, arg1 *IpsecSADecryptDeleteRequestMsg) (*IpsecSADecryptDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecSADecryptDelete", arg0, arg1)
	ret0, _ := ret[0].(*IpsecSADecryptDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSADecryptDelete indicates an expected call of IpsecSADecryptDelete
func (mr *MockIpsecServerMockRecorder) IpsecSADecryptDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSADecryptDelete", reflect.TypeOf((*MockIpsecServer)(nil).IpsecSADecryptDelete), arg0, arg1)
}

// IpsecSADecryptGet mocks base method
func (m *MockIpsecServer) IpsecSADecryptGet(arg0 context.Context, arg1 *IpsecSADecryptGetRequestMsg) (*IpsecSADecryptGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "IpsecSADecryptGet", arg0, arg1)
	ret0, _ := ret[0].(*IpsecSADecryptGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IpsecSADecryptGet indicates an expected call of IpsecSADecryptGet
func (mr *MockIpsecServerMockRecorder) IpsecSADecryptGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IpsecSADecryptGet", reflect.TypeOf((*MockIpsecServer)(nil).IpsecSADecryptGet), arg0, arg1)
}
