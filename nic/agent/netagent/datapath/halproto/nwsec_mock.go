// Code generated by MockGen. DO NOT EDIT.
// Source: nwsec.pb.go

// Package halproto is a generated GoMock package.
package halproto

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// MockisSecurityProfileStatus_ProfilePdStatus is a mock of isSecurityProfileStatus_ProfilePdStatus interface
type MockisSecurityProfileStatus_ProfilePdStatus struct {
	ctrl     *gomock.Controller
	recorder *MockisSecurityProfileStatus_ProfilePdStatusMockRecorder
}

// MockisSecurityProfileStatus_ProfilePdStatusMockRecorder is the mock recorder for MockisSecurityProfileStatus_ProfilePdStatus
type MockisSecurityProfileStatus_ProfilePdStatusMockRecorder struct {
	mock *MockisSecurityProfileStatus_ProfilePdStatus
}

// NewMockisSecurityProfileStatus_ProfilePdStatus creates a new mock instance
func NewMockisSecurityProfileStatus_ProfilePdStatus(ctrl *gomock.Controller) *MockisSecurityProfileStatus_ProfilePdStatus {
	mock := &MockisSecurityProfileStatus_ProfilePdStatus{ctrl: ctrl}
	mock.recorder = &MockisSecurityProfileStatus_ProfilePdStatusMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisSecurityProfileStatus_ProfilePdStatus) EXPECT() *MockisSecurityProfileStatus_ProfilePdStatusMockRecorder {
	return m.recorder
}

// isSecurityProfileStatus_ProfilePdStatus mocks base method
func (m *MockisSecurityProfileStatus_ProfilePdStatus) isSecurityProfileStatus_ProfilePdStatus() {
	m.ctrl.Call(m, "isSecurityProfileStatus_ProfilePdStatus")
}

// isSecurityProfileStatus_ProfilePdStatus indicates an expected call of isSecurityProfileStatus_ProfilePdStatus
func (mr *MockisSecurityProfileStatus_ProfilePdStatusMockRecorder) isSecurityProfileStatus_ProfilePdStatus() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isSecurityProfileStatus_ProfilePdStatus", reflect.TypeOf((*MockisSecurityProfileStatus_ProfilePdStatus)(nil).isSecurityProfileStatus_ProfilePdStatus))
}

// MarshalTo mocks base method
func (m *MockisSecurityProfileStatus_ProfilePdStatus) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisSecurityProfileStatus_ProfilePdStatusMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisSecurityProfileStatus_ProfilePdStatus)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisSecurityProfileStatus_ProfilePdStatus) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisSecurityProfileStatus_ProfilePdStatusMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisSecurityProfileStatus_ProfilePdStatus)(nil).Size))
}

// MockisService_L4Info is a mock of isService_L4Info interface
type MockisService_L4Info struct {
	ctrl     *gomock.Controller
	recorder *MockisService_L4InfoMockRecorder
}

// MockisService_L4InfoMockRecorder is the mock recorder for MockisService_L4Info
type MockisService_L4InfoMockRecorder struct {
	mock *MockisService_L4Info
}

// NewMockisService_L4Info creates a new mock instance
func NewMockisService_L4Info(ctrl *gomock.Controller) *MockisService_L4Info {
	mock := &MockisService_L4Info{ctrl: ctrl}
	mock.recorder = &MockisService_L4InfoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisService_L4Info) EXPECT() *MockisService_L4InfoMockRecorder {
	return m.recorder
}

// isService_L4Info mocks base method
func (m *MockisService_L4Info) isService_L4Info() {
	m.ctrl.Call(m, "isService_L4Info")
}

// isService_L4Info indicates an expected call of isService_L4Info
func (mr *MockisService_L4InfoMockRecorder) isService_L4Info() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isService_L4Info", reflect.TypeOf((*MockisService_L4Info)(nil).isService_L4Info))
}

// MarshalTo mocks base method
func (m *MockisService_L4Info) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisService_L4InfoMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisService_L4Info)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisService_L4Info) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisService_L4InfoMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisService_L4Info)(nil).Size))
}

// MockisAppData_AppOptions is a mock of isAppData_AppOptions interface
type MockisAppData_AppOptions struct {
	ctrl     *gomock.Controller
	recorder *MockisAppData_AppOptionsMockRecorder
}

// MockisAppData_AppOptionsMockRecorder is the mock recorder for MockisAppData_AppOptions
type MockisAppData_AppOptionsMockRecorder struct {
	mock *MockisAppData_AppOptions
}

// NewMockisAppData_AppOptions creates a new mock instance
func NewMockisAppData_AppOptions(ctrl *gomock.Controller) *MockisAppData_AppOptions {
	mock := &MockisAppData_AppOptions{ctrl: ctrl}
	mock.recorder = &MockisAppData_AppOptionsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockisAppData_AppOptions) EXPECT() *MockisAppData_AppOptionsMockRecorder {
	return m.recorder
}

// isAppData_AppOptions mocks base method
func (m *MockisAppData_AppOptions) isAppData_AppOptions() {
	m.ctrl.Call(m, "isAppData_AppOptions")
}

// isAppData_AppOptions indicates an expected call of isAppData_AppOptions
func (mr *MockisAppData_AppOptionsMockRecorder) isAppData_AppOptions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "isAppData_AppOptions", reflect.TypeOf((*MockisAppData_AppOptions)(nil).isAppData_AppOptions))
}

// MarshalTo mocks base method
func (m *MockisAppData_AppOptions) MarshalTo(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "MarshalTo", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalTo indicates an expected call of MarshalTo
func (mr *MockisAppData_AppOptionsMockRecorder) MarshalTo(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalTo", reflect.TypeOf((*MockisAppData_AppOptions)(nil).MarshalTo), arg0)
}

// Size mocks base method
func (m *MockisAppData_AppOptions) Size() int {
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size
func (mr *MockisAppData_AppOptionsMockRecorder) Size() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockisAppData_AppOptions)(nil).Size))
}

// MockNwSecurityClient is a mock of NwSecurityClient interface
type MockNwSecurityClient struct {
	ctrl     *gomock.Controller
	recorder *MockNwSecurityClientMockRecorder
}

// MockNwSecurityClientMockRecorder is the mock recorder for MockNwSecurityClient
type MockNwSecurityClientMockRecorder struct {
	mock *MockNwSecurityClient
}

// NewMockNwSecurityClient creates a new mock instance
func NewMockNwSecurityClient(ctrl *gomock.Controller) *MockNwSecurityClient {
	mock := &MockNwSecurityClient{ctrl: ctrl}
	mock.recorder = &MockNwSecurityClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNwSecurityClient) EXPECT() *MockNwSecurityClientMockRecorder {
	return m.recorder
}

// SecurityProfileCreate mocks base method
func (m *MockNwSecurityClient) SecurityProfileCreate(ctx context.Context, in *SecurityProfileRequestMsg, opts ...grpc.CallOption) (*SecurityProfileResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityProfileCreate", varargs...)
	ret0, _ := ret[0].(*SecurityProfileResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileCreate indicates an expected call of SecurityProfileCreate
func (mr *MockNwSecurityClientMockRecorder) SecurityProfileCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProfileCreate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityProfileCreate), varargs...)
}

// SecurityProfileUpdate mocks base method
func (m *MockNwSecurityClient) SecurityProfileUpdate(ctx context.Context, in *SecurityProfileRequestMsg, opts ...grpc.CallOption) (*SecurityProfileResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityProfileUpdate", varargs...)
	ret0, _ := ret[0].(*SecurityProfileResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileUpdate indicates an expected call of SecurityProfileUpdate
func (mr *MockNwSecurityClientMockRecorder) SecurityProfileUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProfileUpdate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityProfileUpdate), varargs...)
}

// SecurityProfileDelete mocks base method
func (m *MockNwSecurityClient) SecurityProfileDelete(ctx context.Context, in *SecurityProfileDeleteRequestMsg, opts ...grpc.CallOption) (*SecurityProfileDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityProfileDelete", varargs...)
	ret0, _ := ret[0].(*SecurityProfileDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileDelete indicates an expected call of SecurityProfileDelete
func (mr *MockNwSecurityClientMockRecorder) SecurityProfileDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProfileDelete", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityProfileDelete), varargs...)
}

// SecurityProfileGet mocks base method
func (m *MockNwSecurityClient) SecurityProfileGet(ctx context.Context, in *SecurityProfileGetRequestMsg, opts ...grpc.CallOption) (*SecurityProfileGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityProfileGet", varargs...)
	ret0, _ := ret[0].(*SecurityProfileGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileGet indicates an expected call of SecurityProfileGet
func (mr *MockNwSecurityClientMockRecorder) SecurityProfileGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProfileGet", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityProfileGet), varargs...)
}

// SecurityGroupPolicyCreate mocks base method
func (m *MockNwSecurityClient) SecurityGroupPolicyCreate(ctx context.Context, in *SecurityGroupPolicyRequestMsg, opts ...grpc.CallOption) (*SecurityGroupPolicyResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityGroupPolicyCreate", varargs...)
	ret0, _ := ret[0].(*SecurityGroupPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupPolicyCreate indicates an expected call of SecurityGroupPolicyCreate
func (mr *MockNwSecurityClientMockRecorder) SecurityGroupPolicyCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupPolicyCreate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupPolicyCreate), varargs...)
}

// SecurityGroupPolicyUpdate mocks base method
func (m *MockNwSecurityClient) SecurityGroupPolicyUpdate(ctx context.Context, in *SecurityGroupPolicyRequestMsg, opts ...grpc.CallOption) (*SecurityGroupPolicyResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityGroupPolicyUpdate", varargs...)
	ret0, _ := ret[0].(*SecurityGroupPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupPolicyUpdate indicates an expected call of SecurityGroupPolicyUpdate
func (mr *MockNwSecurityClientMockRecorder) SecurityGroupPolicyUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupPolicyUpdate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupPolicyUpdate), varargs...)
}

// SecurityGroupPolicyDelete mocks base method
func (m *MockNwSecurityClient) SecurityGroupPolicyDelete(ctx context.Context, in *SecurityGroupPolicyDeleteRequestMsg, opts ...grpc.CallOption) (*SecurityGroupPolicyDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityGroupPolicyDelete", varargs...)
	ret0, _ := ret[0].(*SecurityGroupPolicyDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupPolicyDelete indicates an expected call of SecurityGroupPolicyDelete
func (mr *MockNwSecurityClientMockRecorder) SecurityGroupPolicyDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupPolicyDelete", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupPolicyDelete), varargs...)
}

// SecurityGroupPolicyGet mocks base method
func (m *MockNwSecurityClient) SecurityGroupPolicyGet(ctx context.Context, in *SecurityGroupPolicyGetRequestMsg, opts ...grpc.CallOption) (*SecurityGroupPolicyGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityGroupPolicyGet", varargs...)
	ret0, _ := ret[0].(*SecurityGroupPolicyGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupPolicyGet indicates an expected call of SecurityGroupPolicyGet
func (mr *MockNwSecurityClientMockRecorder) SecurityGroupPolicyGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupPolicyGet", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupPolicyGet), varargs...)
}

// SecurityGroupCreate mocks base method
func (m *MockNwSecurityClient) SecurityGroupCreate(ctx context.Context, in *SecurityGroupRequestMsg, opts ...grpc.CallOption) (*SecurityGroupResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityGroupCreate", varargs...)
	ret0, _ := ret[0].(*SecurityGroupResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupCreate indicates an expected call of SecurityGroupCreate
func (mr *MockNwSecurityClientMockRecorder) SecurityGroupCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupCreate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupCreate), varargs...)
}

// SecurityGroupUpdate mocks base method
func (m *MockNwSecurityClient) SecurityGroupUpdate(ctx context.Context, in *SecurityGroupRequestMsg, opts ...grpc.CallOption) (*SecurityGroupResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityGroupUpdate", varargs...)
	ret0, _ := ret[0].(*SecurityGroupResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupUpdate indicates an expected call of SecurityGroupUpdate
func (mr *MockNwSecurityClientMockRecorder) SecurityGroupUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupUpdate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupUpdate), varargs...)
}

// SecurityGroupDelete mocks base method
func (m *MockNwSecurityClient) SecurityGroupDelete(ctx context.Context, in *SecurityGroupDeleteRequestMsg, opts ...grpc.CallOption) (*SecurityGroupDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityGroupDelete", varargs...)
	ret0, _ := ret[0].(*SecurityGroupDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupDelete indicates an expected call of SecurityGroupDelete
func (mr *MockNwSecurityClientMockRecorder) SecurityGroupDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupDelete", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupDelete), varargs...)
}

// SecurityGroupGet mocks base method
func (m *MockNwSecurityClient) SecurityGroupGet(ctx context.Context, in *SecurityGroupGetRequestMsg, opts ...grpc.CallOption) (*SecurityGroupGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityGroupGet", varargs...)
	ret0, _ := ret[0].(*SecurityGroupGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupGet indicates an expected call of SecurityGroupGet
func (mr *MockNwSecurityClientMockRecorder) SecurityGroupGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupGet", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityGroupGet), varargs...)
}

// SecurityPolicyCreate mocks base method
func (m *MockNwSecurityClient) SecurityPolicyCreate(ctx context.Context, in *SecurityPolicyRequestMsg, opts ...grpc.CallOption) (*SecurityPolicyResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityPolicyCreate", varargs...)
	ret0, _ := ret[0].(*SecurityPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityPolicyCreate indicates an expected call of SecurityPolicyCreate
func (mr *MockNwSecurityClientMockRecorder) SecurityPolicyCreate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityPolicyCreate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityPolicyCreate), varargs...)
}

// SecurityPolicyUpdate mocks base method
func (m *MockNwSecurityClient) SecurityPolicyUpdate(ctx context.Context, in *SecurityPolicyRequestMsg, opts ...grpc.CallOption) (*SecurityPolicyResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityPolicyUpdate", varargs...)
	ret0, _ := ret[0].(*SecurityPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityPolicyUpdate indicates an expected call of SecurityPolicyUpdate
func (mr *MockNwSecurityClientMockRecorder) SecurityPolicyUpdate(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityPolicyUpdate", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityPolicyUpdate), varargs...)
}

// SecurityPolicyDelete mocks base method
func (m *MockNwSecurityClient) SecurityPolicyDelete(ctx context.Context, in *SecurityPolicyDeleteRequestMsg, opts ...grpc.CallOption) (*SecurityPolicyDeleteResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityPolicyDelete", varargs...)
	ret0, _ := ret[0].(*SecurityPolicyDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityPolicyDelete indicates an expected call of SecurityPolicyDelete
func (mr *MockNwSecurityClientMockRecorder) SecurityPolicyDelete(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityPolicyDelete", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityPolicyDelete), varargs...)
}

// SecurityPolicyGet mocks base method
func (m *MockNwSecurityClient) SecurityPolicyGet(ctx context.Context, in *SecurityPolicyGetRequestMsg, opts ...grpc.CallOption) (*SecurityPolicyGetResponseMsg, error) {
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SecurityPolicyGet", varargs...)
	ret0, _ := ret[0].(*SecurityPolicyGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityPolicyGet indicates an expected call of SecurityPolicyGet
func (mr *MockNwSecurityClientMockRecorder) SecurityPolicyGet(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityPolicyGet", reflect.TypeOf((*MockNwSecurityClient)(nil).SecurityPolicyGet), varargs...)
}

// MockNwSecurityServer is a mock of NwSecurityServer interface
type MockNwSecurityServer struct {
	ctrl     *gomock.Controller
	recorder *MockNwSecurityServerMockRecorder
}

// MockNwSecurityServerMockRecorder is the mock recorder for MockNwSecurityServer
type MockNwSecurityServerMockRecorder struct {
	mock *MockNwSecurityServer
}

// NewMockNwSecurityServer creates a new mock instance
func NewMockNwSecurityServer(ctrl *gomock.Controller) *MockNwSecurityServer {
	mock := &MockNwSecurityServer{ctrl: ctrl}
	mock.recorder = &MockNwSecurityServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNwSecurityServer) EXPECT() *MockNwSecurityServerMockRecorder {
	return m.recorder
}

// SecurityProfileCreate mocks base method
func (m *MockNwSecurityServer) SecurityProfileCreate(arg0 context.Context, arg1 *SecurityProfileRequestMsg) (*SecurityProfileResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityProfileCreate", arg0, arg1)
	ret0, _ := ret[0].(*SecurityProfileResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileCreate indicates an expected call of SecurityProfileCreate
func (mr *MockNwSecurityServerMockRecorder) SecurityProfileCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProfileCreate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityProfileCreate), arg0, arg1)
}

// SecurityProfileUpdate mocks base method
func (m *MockNwSecurityServer) SecurityProfileUpdate(arg0 context.Context, arg1 *SecurityProfileRequestMsg) (*SecurityProfileResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityProfileUpdate", arg0, arg1)
	ret0, _ := ret[0].(*SecurityProfileResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileUpdate indicates an expected call of SecurityProfileUpdate
func (mr *MockNwSecurityServerMockRecorder) SecurityProfileUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProfileUpdate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityProfileUpdate), arg0, arg1)
}

// SecurityProfileDelete mocks base method
func (m *MockNwSecurityServer) SecurityProfileDelete(arg0 context.Context, arg1 *SecurityProfileDeleteRequestMsg) (*SecurityProfileDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityProfileDelete", arg0, arg1)
	ret0, _ := ret[0].(*SecurityProfileDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileDelete indicates an expected call of SecurityProfileDelete
func (mr *MockNwSecurityServerMockRecorder) SecurityProfileDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProfileDelete", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityProfileDelete), arg0, arg1)
}

// SecurityProfileGet mocks base method
func (m *MockNwSecurityServer) SecurityProfileGet(arg0 context.Context, arg1 *SecurityProfileGetRequestMsg) (*SecurityProfileGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityProfileGet", arg0, arg1)
	ret0, _ := ret[0].(*SecurityProfileGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityProfileGet indicates an expected call of SecurityProfileGet
func (mr *MockNwSecurityServerMockRecorder) SecurityProfileGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityProfileGet", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityProfileGet), arg0, arg1)
}

// SecurityGroupPolicyCreate mocks base method
func (m *MockNwSecurityServer) SecurityGroupPolicyCreate(arg0 context.Context, arg1 *SecurityGroupPolicyRequestMsg) (*SecurityGroupPolicyResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityGroupPolicyCreate", arg0, arg1)
	ret0, _ := ret[0].(*SecurityGroupPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupPolicyCreate indicates an expected call of SecurityGroupPolicyCreate
func (mr *MockNwSecurityServerMockRecorder) SecurityGroupPolicyCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupPolicyCreate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupPolicyCreate), arg0, arg1)
}

// SecurityGroupPolicyUpdate mocks base method
func (m *MockNwSecurityServer) SecurityGroupPolicyUpdate(arg0 context.Context, arg1 *SecurityGroupPolicyRequestMsg) (*SecurityGroupPolicyResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityGroupPolicyUpdate", arg0, arg1)
	ret0, _ := ret[0].(*SecurityGroupPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupPolicyUpdate indicates an expected call of SecurityGroupPolicyUpdate
func (mr *MockNwSecurityServerMockRecorder) SecurityGroupPolicyUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupPolicyUpdate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupPolicyUpdate), arg0, arg1)
}

// SecurityGroupPolicyDelete mocks base method
func (m *MockNwSecurityServer) SecurityGroupPolicyDelete(arg0 context.Context, arg1 *SecurityGroupPolicyDeleteRequestMsg) (*SecurityGroupPolicyDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityGroupPolicyDelete", arg0, arg1)
	ret0, _ := ret[0].(*SecurityGroupPolicyDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupPolicyDelete indicates an expected call of SecurityGroupPolicyDelete
func (mr *MockNwSecurityServerMockRecorder) SecurityGroupPolicyDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupPolicyDelete", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupPolicyDelete), arg0, arg1)
}

// SecurityGroupPolicyGet mocks base method
func (m *MockNwSecurityServer) SecurityGroupPolicyGet(arg0 context.Context, arg1 *SecurityGroupPolicyGetRequestMsg) (*SecurityGroupPolicyGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityGroupPolicyGet", arg0, arg1)
	ret0, _ := ret[0].(*SecurityGroupPolicyGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupPolicyGet indicates an expected call of SecurityGroupPolicyGet
func (mr *MockNwSecurityServerMockRecorder) SecurityGroupPolicyGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupPolicyGet", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupPolicyGet), arg0, arg1)
}

// SecurityGroupCreate mocks base method
func (m *MockNwSecurityServer) SecurityGroupCreate(arg0 context.Context, arg1 *SecurityGroupRequestMsg) (*SecurityGroupResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityGroupCreate", arg0, arg1)
	ret0, _ := ret[0].(*SecurityGroupResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupCreate indicates an expected call of SecurityGroupCreate
func (mr *MockNwSecurityServerMockRecorder) SecurityGroupCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupCreate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupCreate), arg0, arg1)
}

// SecurityGroupUpdate mocks base method
func (m *MockNwSecurityServer) SecurityGroupUpdate(arg0 context.Context, arg1 *SecurityGroupRequestMsg) (*SecurityGroupResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityGroupUpdate", arg0, arg1)
	ret0, _ := ret[0].(*SecurityGroupResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupUpdate indicates an expected call of SecurityGroupUpdate
func (mr *MockNwSecurityServerMockRecorder) SecurityGroupUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupUpdate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupUpdate), arg0, arg1)
}

// SecurityGroupDelete mocks base method
func (m *MockNwSecurityServer) SecurityGroupDelete(arg0 context.Context, arg1 *SecurityGroupDeleteRequestMsg) (*SecurityGroupDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityGroupDelete", arg0, arg1)
	ret0, _ := ret[0].(*SecurityGroupDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupDelete indicates an expected call of SecurityGroupDelete
func (mr *MockNwSecurityServerMockRecorder) SecurityGroupDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupDelete", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupDelete), arg0, arg1)
}

// SecurityGroupGet mocks base method
func (m *MockNwSecurityServer) SecurityGroupGet(arg0 context.Context, arg1 *SecurityGroupGetRequestMsg) (*SecurityGroupGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityGroupGet", arg0, arg1)
	ret0, _ := ret[0].(*SecurityGroupGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityGroupGet indicates an expected call of SecurityGroupGet
func (mr *MockNwSecurityServerMockRecorder) SecurityGroupGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityGroupGet", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityGroupGet), arg0, arg1)
}

// SecurityPolicyCreate mocks base method
func (m *MockNwSecurityServer) SecurityPolicyCreate(arg0 context.Context, arg1 *SecurityPolicyRequestMsg) (*SecurityPolicyResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityPolicyCreate", arg0, arg1)
	ret0, _ := ret[0].(*SecurityPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityPolicyCreate indicates an expected call of SecurityPolicyCreate
func (mr *MockNwSecurityServerMockRecorder) SecurityPolicyCreate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityPolicyCreate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityPolicyCreate), arg0, arg1)
}

// SecurityPolicyUpdate mocks base method
func (m *MockNwSecurityServer) SecurityPolicyUpdate(arg0 context.Context, arg1 *SecurityPolicyRequestMsg) (*SecurityPolicyResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityPolicyUpdate", arg0, arg1)
	ret0, _ := ret[0].(*SecurityPolicyResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityPolicyUpdate indicates an expected call of SecurityPolicyUpdate
func (mr *MockNwSecurityServerMockRecorder) SecurityPolicyUpdate(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityPolicyUpdate", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityPolicyUpdate), arg0, arg1)
}

// SecurityPolicyDelete mocks base method
func (m *MockNwSecurityServer) SecurityPolicyDelete(arg0 context.Context, arg1 *SecurityPolicyDeleteRequestMsg) (*SecurityPolicyDeleteResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityPolicyDelete", arg0, arg1)
	ret0, _ := ret[0].(*SecurityPolicyDeleteResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityPolicyDelete indicates an expected call of SecurityPolicyDelete
func (mr *MockNwSecurityServerMockRecorder) SecurityPolicyDelete(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityPolicyDelete", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityPolicyDelete), arg0, arg1)
}

// SecurityPolicyGet mocks base method
func (m *MockNwSecurityServer) SecurityPolicyGet(arg0 context.Context, arg1 *SecurityPolicyGetRequestMsg) (*SecurityPolicyGetResponseMsg, error) {
	ret := m.ctrl.Call(m, "SecurityPolicyGet", arg0, arg1)
	ret0, _ := ret[0].(*SecurityPolicyGetResponseMsg)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecurityPolicyGet indicates an expected call of SecurityPolicyGet
func (mr *MockNwSecurityServerMockRecorder) SecurityPolicyGet(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecurityPolicyGet", reflect.TypeOf((*MockNwSecurityServer)(nil).SecurityPolicyGet), arg0, arg1)
}
