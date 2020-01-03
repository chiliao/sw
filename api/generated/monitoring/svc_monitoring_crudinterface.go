// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

package monitoring

import (
	"context"

	api "github.com/pensando/sw/api"
	"github.com/pensando/sw/api/interfaces"
	"github.com/pensando/sw/venice/utils/kvstore"
)

// Dummy vars to suppress unused imports message
var _ context.Context
var _ api.ObjectMeta
var _ kvstore.Interface

// MonitoringV1EventPolicyInterface exposes the CRUD methods for EventPolicy
type MonitoringV1EventPolicyInterface interface {
	Create(ctx context.Context, in *EventPolicy) (*EventPolicy, error)
	Update(ctx context.Context, in *EventPolicy) (*EventPolicy, error)
	UpdateStatus(ctx context.Context, in *EventPolicy) (*EventPolicy, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*EventPolicy, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*EventPolicy, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*EventPolicy, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1StatsPolicyInterface exposes the CRUD methods for StatsPolicy
type MonitoringV1StatsPolicyInterface interface {
	Create(ctx context.Context, in *StatsPolicy) (*StatsPolicy, error)
	Update(ctx context.Context, in *StatsPolicy) (*StatsPolicy, error)
	UpdateStatus(ctx context.Context, in *StatsPolicy) (*StatsPolicy, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*StatsPolicy, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*StatsPolicy, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*StatsPolicy, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1FwlogPolicyInterface exposes the CRUD methods for FwlogPolicy
type MonitoringV1FwlogPolicyInterface interface {
	Create(ctx context.Context, in *FwlogPolicy) (*FwlogPolicy, error)
	Update(ctx context.Context, in *FwlogPolicy) (*FwlogPolicy, error)
	UpdateStatus(ctx context.Context, in *FwlogPolicy) (*FwlogPolicy, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*FwlogPolicy, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*FwlogPolicy, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*FwlogPolicy, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1FlowExportPolicyInterface exposes the CRUD methods for FlowExportPolicy
type MonitoringV1FlowExportPolicyInterface interface {
	Create(ctx context.Context, in *FlowExportPolicy) (*FlowExportPolicy, error)
	Update(ctx context.Context, in *FlowExportPolicy) (*FlowExportPolicy, error)
	UpdateStatus(ctx context.Context, in *FlowExportPolicy) (*FlowExportPolicy, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*FlowExportPolicy, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*FlowExportPolicy, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*FlowExportPolicy, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1AlertInterface exposes the CRUD methods for Alert
type MonitoringV1AlertInterface interface {
	Create(ctx context.Context, in *Alert) (*Alert, error)
	Update(ctx context.Context, in *Alert) (*Alert, error)
	UpdateStatus(ctx context.Context, in *Alert) (*Alert, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*Alert, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*Alert, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*Alert, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1AlertPolicyInterface exposes the CRUD methods for AlertPolicy
type MonitoringV1AlertPolicyInterface interface {
	Create(ctx context.Context, in *AlertPolicy) (*AlertPolicy, error)
	Update(ctx context.Context, in *AlertPolicy) (*AlertPolicy, error)
	UpdateStatus(ctx context.Context, in *AlertPolicy) (*AlertPolicy, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*AlertPolicy, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*AlertPolicy, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*AlertPolicy, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1AlertDestinationInterface exposes the CRUD methods for AlertDestination
type MonitoringV1AlertDestinationInterface interface {
	Create(ctx context.Context, in *AlertDestination) (*AlertDestination, error)
	Update(ctx context.Context, in *AlertDestination) (*AlertDestination, error)
	UpdateStatus(ctx context.Context, in *AlertDestination) (*AlertDestination, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*AlertDestination, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*AlertDestination, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*AlertDestination, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1MirrorSessionInterface exposes the CRUD methods for MirrorSession
type MonitoringV1MirrorSessionInterface interface {
	Create(ctx context.Context, in *MirrorSession) (*MirrorSession, error)
	Update(ctx context.Context, in *MirrorSession) (*MirrorSession, error)
	UpdateStatus(ctx context.Context, in *MirrorSession) (*MirrorSession, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*MirrorSession, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*MirrorSession, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*MirrorSession, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1TroubleshootingSessionInterface exposes the CRUD methods for TroubleshootingSession
type MonitoringV1TroubleshootingSessionInterface interface {
	Create(ctx context.Context, in *TroubleshootingSession) (*TroubleshootingSession, error)
	Update(ctx context.Context, in *TroubleshootingSession) (*TroubleshootingSession, error)
	UpdateStatus(ctx context.Context, in *TroubleshootingSession) (*TroubleshootingSession, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*TroubleshootingSession, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*TroubleshootingSession, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*TroubleshootingSession, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1TechSupportRequestInterface exposes the CRUD methods for TechSupportRequest
type MonitoringV1TechSupportRequestInterface interface {
	Create(ctx context.Context, in *TechSupportRequest) (*TechSupportRequest, error)
	Update(ctx context.Context, in *TechSupportRequest) (*TechSupportRequest, error)
	UpdateStatus(ctx context.Context, in *TechSupportRequest) (*TechSupportRequest, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*TechSupportRequest, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*TechSupportRequest, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*TechSupportRequest, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
}

// MonitoringV1ArchiveRequestInterface exposes the CRUD methods for ArchiveRequest
type MonitoringV1ArchiveRequestInterface interface {
	Create(ctx context.Context, in *ArchiveRequest) (*ArchiveRequest, error)
	Update(ctx context.Context, in *ArchiveRequest) (*ArchiveRequest, error)
	UpdateStatus(ctx context.Context, in *ArchiveRequest) (*ArchiveRequest, error)
	Get(ctx context.Context, objMeta *api.ObjectMeta) (*ArchiveRequest, error)
	Delete(ctx context.Context, objMeta *api.ObjectMeta) (*ArchiveRequest, error)
	List(ctx context.Context, options *api.ListWatchOptions) ([]*ArchiveRequest, error)
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
	Allowed(oper apiintf.APIOperType) bool
	Cancel(ctx context.Context, in *CancelArchiveRequest) (*ArchiveRequest, error)
}

// MonitoringV1Interface exposes objects with CRUD operations allowed by the service
type MonitoringV1Interface interface {
	EventPolicy() MonitoringV1EventPolicyInterface
	StatsPolicy() MonitoringV1StatsPolicyInterface
	FwlogPolicy() MonitoringV1FwlogPolicyInterface
	FlowExportPolicy() MonitoringV1FlowExportPolicyInterface
	Alert() MonitoringV1AlertInterface
	AlertPolicy() MonitoringV1AlertPolicyInterface
	AlertDestination() MonitoringV1AlertDestinationInterface
	MirrorSession() MonitoringV1MirrorSessionInterface
	TroubleshootingSession() MonitoringV1TroubleshootingSessionInterface
	TechSupportRequest() MonitoringV1TechSupportRequestInterface
	ArchiveRequest() MonitoringV1ArchiveRequestInterface
	Watch(ctx context.Context, options *api.ListWatchOptions) (kvstore.Watcher, error)
}
