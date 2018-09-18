// Code generated by MockGen. DO NOT EDIT.
// Source: ../generated/cluster/svc_cluster_crudinterface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"

	api "github.com/pensando/sw/api"
	cluster "github.com/pensando/sw/api/generated/cluster"
	apiserver "github.com/pensando/sw/venice/apiserver"
	kvstore "github.com/pensando/sw/venice/utils/kvstore"
)

// MockClusterV1ClusterInterface is a mock of (cluster.ClusterV1ClusterInterface)interface
type MockClusterV1ClusterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterV1ClusterInterfaceMockRecorder
}

// MockClusterV1ClusterInterfaceMockRecorder is the mock recorder for MockClusterV1ClusterInterface
type MockClusterV1ClusterInterfaceMockRecorder struct {
	mock *MockClusterV1ClusterInterface
}

// NewMockClusterV1ClusterInterface creates a new mock instance
func NewMockClusterV1ClusterInterface(ctrl *gomock.Controller) *MockClusterV1ClusterInterface {
	mock := &MockClusterV1ClusterInterface{ctrl: ctrl}
	mock.recorder = &MockClusterV1ClusterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterV1ClusterInterface) EXPECT() *MockClusterV1ClusterInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockClusterV1ClusterInterface) Create(ctx context.Context, in *cluster.Cluster) (*cluster.Cluster, error) {
	ret := m.ctrl.Call(m, "Create", ctx, in)
	ret0, _ := ret[0].(*cluster.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockClusterV1ClusterInterfaceMockRecorder) Create(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterV1ClusterInterface)(nil).Create), ctx, in)
}

// Update mocks base method
func (m *MockClusterV1ClusterInterface) Update(ctx context.Context, in *cluster.Cluster) (*cluster.Cluster, error) {
	ret := m.ctrl.Call(m, "Update", ctx, in)
	ret0, _ := ret[0].(*cluster.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockClusterV1ClusterInterfaceMockRecorder) Update(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterV1ClusterInterface)(nil).Update), ctx, in)
}

// Get mocks base method
func (m *MockClusterV1ClusterInterface) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.Cluster, error) {
	ret := m.ctrl.Call(m, "Get", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClusterV1ClusterInterfaceMockRecorder) Get(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterV1ClusterInterface)(nil).Get), ctx, objMeta)
}

// Delete mocks base method
func (m *MockClusterV1ClusterInterface) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.Cluster, error) {
	ret := m.ctrl.Call(m, "Delete", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockClusterV1ClusterInterfaceMockRecorder) Delete(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterV1ClusterInterface)(nil).Delete), ctx, objMeta)
}

// List mocks base method
func (m *MockClusterV1ClusterInterface) List(ctx context.Context, options *api.ListWatchOptions) ([]*cluster.Cluster, error) {
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].([]*cluster.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClusterV1ClusterInterfaceMockRecorder) List(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterV1ClusterInterface)(nil).List), ctx, options)
}

// Watch mocks base method
func (m *MockClusterV1ClusterInterface) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	ret := m.ctrl.Call(m, "Watch", ctx, options)
	ret0, _ := ret[0].(kvstore.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockClusterV1ClusterInterfaceMockRecorder) Watch(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterV1ClusterInterface)(nil).Watch), ctx, options)
}

// Allowed mocks base method
func (m *MockClusterV1ClusterInterface) Allowed(oper apiserver.APIOperType) bool {
	ret := m.ctrl.Call(m, "Allowed", oper)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Allowed indicates an expected call of Allowed
func (mr *MockClusterV1ClusterInterfaceMockRecorder) Allowed(oper interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allowed", reflect.TypeOf((*MockClusterV1ClusterInterface)(nil).Allowed), oper)
}

// AuthBootstrapComplete mocks base method
func (m *MockClusterV1ClusterInterface) AuthBootstrapComplete(ctx context.Context, in *cluster.ClusterAuthBootstrapRequest) (*cluster.Cluster, error) {
	ret := m.ctrl.Call(m, "AuthBootstrapComplete", ctx, in)
	ret0, _ := ret[0].(*cluster.Cluster)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthBootstrapComplete indicates an expected call of AuthBootstrapComplete
func (mr *MockClusterV1ClusterInterfaceMockRecorder) AuthBootstrapComplete(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthBootstrapComplete", reflect.TypeOf((*MockClusterV1ClusterInterface)(nil).AuthBootstrapComplete), ctx, in)
}

// MockClusterV1NodeInterface is a mock of (cluster.ClusterV1NodeInterface)interface
type MockClusterV1NodeInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterV1NodeInterfaceMockRecorder
}

// MockClusterV1NodeInterfaceMockRecorder is the mock recorder for MockClusterV1NodeInterface
type MockClusterV1NodeInterfaceMockRecorder struct {
	mock *MockClusterV1NodeInterface
}

// NewMockClusterV1NodeInterface creates a new mock instance
func NewMockClusterV1NodeInterface(ctrl *gomock.Controller) *MockClusterV1NodeInterface {
	mock := &MockClusterV1NodeInterface{ctrl: ctrl}
	mock.recorder = &MockClusterV1NodeInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterV1NodeInterface) EXPECT() *MockClusterV1NodeInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockClusterV1NodeInterface) Create(ctx context.Context, in *cluster.Node) (*cluster.Node, error) {
	ret := m.ctrl.Call(m, "Create", ctx, in)
	ret0, _ := ret[0].(*cluster.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockClusterV1NodeInterfaceMockRecorder) Create(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterV1NodeInterface)(nil).Create), ctx, in)
}

// Update mocks base method
func (m *MockClusterV1NodeInterface) Update(ctx context.Context, in *cluster.Node) (*cluster.Node, error) {
	ret := m.ctrl.Call(m, "Update", ctx, in)
	ret0, _ := ret[0].(*cluster.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockClusterV1NodeInterfaceMockRecorder) Update(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterV1NodeInterface)(nil).Update), ctx, in)
}

// Get mocks base method
func (m *MockClusterV1NodeInterface) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.Node, error) {
	ret := m.ctrl.Call(m, "Get", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClusterV1NodeInterfaceMockRecorder) Get(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterV1NodeInterface)(nil).Get), ctx, objMeta)
}

// Delete mocks base method
func (m *MockClusterV1NodeInterface) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.Node, error) {
	ret := m.ctrl.Call(m, "Delete", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockClusterV1NodeInterfaceMockRecorder) Delete(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterV1NodeInterface)(nil).Delete), ctx, objMeta)
}

// List mocks base method
func (m *MockClusterV1NodeInterface) List(ctx context.Context, options *api.ListWatchOptions) ([]*cluster.Node, error) {
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].([]*cluster.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClusterV1NodeInterfaceMockRecorder) List(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterV1NodeInterface)(nil).List), ctx, options)
}

// Watch mocks base method
func (m *MockClusterV1NodeInterface) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	ret := m.ctrl.Call(m, "Watch", ctx, options)
	ret0, _ := ret[0].(kvstore.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockClusterV1NodeInterfaceMockRecorder) Watch(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterV1NodeInterface)(nil).Watch), ctx, options)
}

// Allowed mocks base method
func (m *MockClusterV1NodeInterface) Allowed(oper apiserver.APIOperType) bool {
	ret := m.ctrl.Call(m, "Allowed", oper)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Allowed indicates an expected call of Allowed
func (mr *MockClusterV1NodeInterfaceMockRecorder) Allowed(oper interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allowed", reflect.TypeOf((*MockClusterV1NodeInterface)(nil).Allowed), oper)
}

// MockClusterV1HostInterface is a mock of (cluster.ClusterV1HostInterface)interface
type MockClusterV1HostInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterV1HostInterfaceMockRecorder
}

// MockClusterV1HostInterfaceMockRecorder is the mock recorder for MockClusterV1HostInterface
type MockClusterV1HostInterfaceMockRecorder struct {
	mock *MockClusterV1HostInterface
}

// NewMockClusterV1HostInterface creates a new mock instance
func NewMockClusterV1HostInterface(ctrl *gomock.Controller) *MockClusterV1HostInterface {
	mock := &MockClusterV1HostInterface{ctrl: ctrl}
	mock.recorder = &MockClusterV1HostInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterV1HostInterface) EXPECT() *MockClusterV1HostInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockClusterV1HostInterface) Create(ctx context.Context, in *cluster.Host) (*cluster.Host, error) {
	ret := m.ctrl.Call(m, "Create", ctx, in)
	ret0, _ := ret[0].(*cluster.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockClusterV1HostInterfaceMockRecorder) Create(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterV1HostInterface)(nil).Create), ctx, in)
}

// Update mocks base method
func (m *MockClusterV1HostInterface) Update(ctx context.Context, in *cluster.Host) (*cluster.Host, error) {
	ret := m.ctrl.Call(m, "Update", ctx, in)
	ret0, _ := ret[0].(*cluster.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockClusterV1HostInterfaceMockRecorder) Update(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterV1HostInterface)(nil).Update), ctx, in)
}

// Get mocks base method
func (m *MockClusterV1HostInterface) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.Host, error) {
	ret := m.ctrl.Call(m, "Get", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClusterV1HostInterfaceMockRecorder) Get(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterV1HostInterface)(nil).Get), ctx, objMeta)
}

// Delete mocks base method
func (m *MockClusterV1HostInterface) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.Host, error) {
	ret := m.ctrl.Call(m, "Delete", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockClusterV1HostInterfaceMockRecorder) Delete(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterV1HostInterface)(nil).Delete), ctx, objMeta)
}

// List mocks base method
func (m *MockClusterV1HostInterface) List(ctx context.Context, options *api.ListWatchOptions) ([]*cluster.Host, error) {
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].([]*cluster.Host)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClusterV1HostInterfaceMockRecorder) List(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterV1HostInterface)(nil).List), ctx, options)
}

// Watch mocks base method
func (m *MockClusterV1HostInterface) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	ret := m.ctrl.Call(m, "Watch", ctx, options)
	ret0, _ := ret[0].(kvstore.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockClusterV1HostInterfaceMockRecorder) Watch(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterV1HostInterface)(nil).Watch), ctx, options)
}

// Allowed mocks base method
func (m *MockClusterV1HostInterface) Allowed(oper apiserver.APIOperType) bool {
	ret := m.ctrl.Call(m, "Allowed", oper)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Allowed indicates an expected call of Allowed
func (mr *MockClusterV1HostInterfaceMockRecorder) Allowed(oper interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allowed", reflect.TypeOf((*MockClusterV1HostInterface)(nil).Allowed), oper)
}

// MockClusterV1SmartNICInterface is a mock of (cluster.ClusterV1SmartNICInterface)interface
type MockClusterV1SmartNICInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterV1SmartNICInterfaceMockRecorder
}

// MockClusterV1SmartNICInterfaceMockRecorder is the mock recorder for MockClusterV1SmartNICInterface
type MockClusterV1SmartNICInterfaceMockRecorder struct {
	mock *MockClusterV1SmartNICInterface
}

// NewMockClusterV1SmartNICInterface creates a new mock instance
func NewMockClusterV1SmartNICInterface(ctrl *gomock.Controller) *MockClusterV1SmartNICInterface {
	mock := &MockClusterV1SmartNICInterface{ctrl: ctrl}
	mock.recorder = &MockClusterV1SmartNICInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterV1SmartNICInterface) EXPECT() *MockClusterV1SmartNICInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockClusterV1SmartNICInterface) Create(ctx context.Context, in *cluster.SmartNIC) (*cluster.SmartNIC, error) {
	ret := m.ctrl.Call(m, "Create", ctx, in)
	ret0, _ := ret[0].(*cluster.SmartNIC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockClusterV1SmartNICInterfaceMockRecorder) Create(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterV1SmartNICInterface)(nil).Create), ctx, in)
}

// Update mocks base method
func (m *MockClusterV1SmartNICInterface) Update(ctx context.Context, in *cluster.SmartNIC) (*cluster.SmartNIC, error) {
	ret := m.ctrl.Call(m, "Update", ctx, in)
	ret0, _ := ret[0].(*cluster.SmartNIC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockClusterV1SmartNICInterfaceMockRecorder) Update(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterV1SmartNICInterface)(nil).Update), ctx, in)
}

// Get mocks base method
func (m *MockClusterV1SmartNICInterface) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.SmartNIC, error) {
	ret := m.ctrl.Call(m, "Get", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.SmartNIC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClusterV1SmartNICInterfaceMockRecorder) Get(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterV1SmartNICInterface)(nil).Get), ctx, objMeta)
}

// Delete mocks base method
func (m *MockClusterV1SmartNICInterface) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.SmartNIC, error) {
	ret := m.ctrl.Call(m, "Delete", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.SmartNIC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockClusterV1SmartNICInterfaceMockRecorder) Delete(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterV1SmartNICInterface)(nil).Delete), ctx, objMeta)
}

// List mocks base method
func (m *MockClusterV1SmartNICInterface) List(ctx context.Context, options *api.ListWatchOptions) ([]*cluster.SmartNIC, error) {
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].([]*cluster.SmartNIC)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClusterV1SmartNICInterfaceMockRecorder) List(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterV1SmartNICInterface)(nil).List), ctx, options)
}

// Watch mocks base method
func (m *MockClusterV1SmartNICInterface) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	ret := m.ctrl.Call(m, "Watch", ctx, options)
	ret0, _ := ret[0].(kvstore.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockClusterV1SmartNICInterfaceMockRecorder) Watch(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterV1SmartNICInterface)(nil).Watch), ctx, options)
}

// Allowed mocks base method
func (m *MockClusterV1SmartNICInterface) Allowed(oper apiserver.APIOperType) bool {
	ret := m.ctrl.Call(m, "Allowed", oper)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Allowed indicates an expected call of Allowed
func (mr *MockClusterV1SmartNICInterfaceMockRecorder) Allowed(oper interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allowed", reflect.TypeOf((*MockClusterV1SmartNICInterface)(nil).Allowed), oper)
}

// MockClusterV1TenantInterface is a mock of (cluster.ClusterV1TenantInterface)interface
type MockClusterV1TenantInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterV1TenantInterfaceMockRecorder
}

// MockClusterV1TenantInterfaceMockRecorder is the mock recorder for MockClusterV1TenantInterface
type MockClusterV1TenantInterfaceMockRecorder struct {
	mock *MockClusterV1TenantInterface
}

// NewMockClusterV1TenantInterface creates a new mock instance
func NewMockClusterV1TenantInterface(ctrl *gomock.Controller) *MockClusterV1TenantInterface {
	mock := &MockClusterV1TenantInterface{ctrl: ctrl}
	mock.recorder = &MockClusterV1TenantInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterV1TenantInterface) EXPECT() *MockClusterV1TenantInterfaceMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockClusterV1TenantInterface) Create(ctx context.Context, in *cluster.Tenant) (*cluster.Tenant, error) {
	ret := m.ctrl.Call(m, "Create", ctx, in)
	ret0, _ := ret[0].(*cluster.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockClusterV1TenantInterfaceMockRecorder) Create(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClusterV1TenantInterface)(nil).Create), ctx, in)
}

// Update mocks base method
func (m *MockClusterV1TenantInterface) Update(ctx context.Context, in *cluster.Tenant) (*cluster.Tenant, error) {
	ret := m.ctrl.Call(m, "Update", ctx, in)
	ret0, _ := ret[0].(*cluster.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockClusterV1TenantInterfaceMockRecorder) Update(ctx, in interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClusterV1TenantInterface)(nil).Update), ctx, in)
}

// Get mocks base method
func (m *MockClusterV1TenantInterface) Get(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.Tenant, error) {
	ret := m.ctrl.Call(m, "Get", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockClusterV1TenantInterfaceMockRecorder) Get(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockClusterV1TenantInterface)(nil).Get), ctx, objMeta)
}

// Delete mocks base method
func (m *MockClusterV1TenantInterface) Delete(ctx context.Context, objMeta *api.ObjectMeta) (*cluster.Tenant, error) {
	ret := m.ctrl.Call(m, "Delete", ctx, objMeta)
	ret0, _ := ret[0].(*cluster.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockClusterV1TenantInterfaceMockRecorder) Delete(ctx, objMeta interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockClusterV1TenantInterface)(nil).Delete), ctx, objMeta)
}

// List mocks base method
func (m *MockClusterV1TenantInterface) List(ctx context.Context, options *api.ListWatchOptions) ([]*cluster.Tenant, error) {
	ret := m.ctrl.Call(m, "List", ctx, options)
	ret0, _ := ret[0].([]*cluster.Tenant)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockClusterV1TenantInterfaceMockRecorder) List(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockClusterV1TenantInterface)(nil).List), ctx, options)
}

// Watch mocks base method
func (m *MockClusterV1TenantInterface) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	ret := m.ctrl.Call(m, "Watch", ctx, options)
	ret0, _ := ret[0].(kvstore.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockClusterV1TenantInterfaceMockRecorder) Watch(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterV1TenantInterface)(nil).Watch), ctx, options)
}

// Allowed mocks base method
func (m *MockClusterV1TenantInterface) Allowed(oper apiserver.APIOperType) bool {
	ret := m.ctrl.Call(m, "Allowed", oper)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Allowed indicates an expected call of Allowed
func (mr *MockClusterV1TenantInterfaceMockRecorder) Allowed(oper interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Allowed", reflect.TypeOf((*MockClusterV1TenantInterface)(nil).Allowed), oper)
}

// MockClusterV1Interface is a mock of ClusterV1Interface interface
type MockClusterV1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockClusterV1InterfaceMockRecorder
}

// MockClusterV1InterfaceMockRecorder is the mock recorder for MockClusterV1Interface
type MockClusterV1InterfaceMockRecorder struct {
	mock *MockClusterV1Interface
}

// NewMockClusterV1Interface creates a new mock instance
func NewMockClusterV1Interface(ctrl *gomock.Controller) *MockClusterV1Interface {
	mock := &MockClusterV1Interface{ctrl: ctrl}
	mock.recorder = &MockClusterV1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterV1Interface) EXPECT() *MockClusterV1InterfaceMockRecorder {
	return m.recorder
}

// Cluster mocks base method
func (m *MockClusterV1Interface) Cluster() cluster.ClusterV1ClusterInterface {
	ret := m.ctrl.Call(m, "Cluster")
	ret0, _ := ret[0].(cluster.ClusterV1ClusterInterface)
	return ret0
}

// Cluster indicates an expected call of Cluster
func (mr *MockClusterV1InterfaceMockRecorder) Cluster() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cluster", reflect.TypeOf((*MockClusterV1Interface)(nil).Cluster))
}

// Node mocks base method
func (m *MockClusterV1Interface) Node() cluster.ClusterV1NodeInterface {
	ret := m.ctrl.Call(m, "Node")
	ret0, _ := ret[0].(cluster.ClusterV1NodeInterface)
	return ret0
}

// Node indicates an expected call of Node
func (mr *MockClusterV1InterfaceMockRecorder) Node() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Node", reflect.TypeOf((*MockClusterV1Interface)(nil).Node))
}

// Host mocks base method
func (m *MockClusterV1Interface) Host() cluster.ClusterV1HostInterface {
	ret := m.ctrl.Call(m, "Host")
	ret0, _ := ret[0].(cluster.ClusterV1HostInterface)
	return ret0
}

// Host indicates an expected call of Host
func (mr *MockClusterV1InterfaceMockRecorder) Host() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Host", reflect.TypeOf((*MockClusterV1Interface)(nil).Host))
}

// SmartNIC mocks base method
func (m *MockClusterV1Interface) SmartNIC() cluster.ClusterV1SmartNICInterface {
	ret := m.ctrl.Call(m, "SmartNIC")
	ret0, _ := ret[0].(cluster.ClusterV1SmartNICInterface)
	return ret0
}

// SmartNIC indicates an expected call of SmartNIC
func (mr *MockClusterV1InterfaceMockRecorder) SmartNIC() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SmartNIC", reflect.TypeOf((*MockClusterV1Interface)(nil).SmartNIC))
}

// Tenant mocks base method
func (m *MockClusterV1Interface) Tenant() cluster.ClusterV1TenantInterface {
	ret := m.ctrl.Call(m, "Tenant")
	ret0, _ := ret[0].(cluster.ClusterV1TenantInterface)
	return ret0
}

// Tenant indicates an expected call of Tenant
func (mr *MockClusterV1InterfaceMockRecorder) Tenant() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tenant", reflect.TypeOf((*MockClusterV1Interface)(nil).Tenant))
}

// Watch mocks base method
func (m *MockClusterV1Interface) Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error) {
	ret := m.ctrl.Call(m, "Watch", ctx, options)
	ret0, _ := ret[0].(kvstore.Watcher)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Watch indicates an expected call of Watch
func (mr *MockClusterV1InterfaceMockRecorder) Watch(ctx, options interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockClusterV1Interface)(nil).Watch), ctx, options)
}
