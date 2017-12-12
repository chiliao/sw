// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package events is a auto generated package.
Input file: protos/events.proto
*/
package events

import (
	"context"

	"github.com/pensando/sw/api"
)

// Dummy definitions to suppress nonused warnings
var _ api.ObjectMeta

// ServiceEventPolicyV1Client  is the client interface for the service.
type ServiceEventPolicyV1Client interface {
	AutoAddEventPolicy(ctx context.Context, t *EventPolicy) (*EventPolicy, error)
	AutoDeleteEventPolicy(ctx context.Context, t *EventPolicy) (*EventPolicy, error)
	AutoGetEventPolicy(ctx context.Context, t *EventPolicy) (*EventPolicy, error)
	AutoListEventPolicy(ctx context.Context, t *api.ListWatchOptions) (*EventPolicyList, error)
	AutoUpdateEventPolicy(ctx context.Context, t *EventPolicy) (*EventPolicy, error)

	AutoWatchEventPolicy(ctx context.Context, in *api.ListWatchOptions) (EventPolicyV1_AutoWatchEventPolicyClient, error)
}

// ServiceEventV1Client  is the client interface for the service.
type ServiceEventV1Client interface {
	AutoAddEvent(ctx context.Context, t *Event) (*Event, error)
	AutoDeleteEvent(ctx context.Context, t *Event) (*Event, error)
	AutoGetEvent(ctx context.Context, t *Event) (*Event, error)
	AutoListEvent(ctx context.Context, t *api.ListWatchOptions) (*EventList, error)
	AutoUpdateEvent(ctx context.Context, t *Event) (*Event, error)

	AutoWatchEvent(ctx context.Context, in *api.ListWatchOptions) (EventV1_AutoWatchEventClient, error)
}

// ServiceEventPolicyV1Server is the server interface for the service.
type ServiceEventPolicyV1Server interface {
	AutoAddEventPolicy(ctx context.Context, t EventPolicy) (EventPolicy, error)
	AutoDeleteEventPolicy(ctx context.Context, t EventPolicy) (EventPolicy, error)
	AutoGetEventPolicy(ctx context.Context, t EventPolicy) (EventPolicy, error)
	AutoListEventPolicy(ctx context.Context, t api.ListWatchOptions) (EventPolicyList, error)
	AutoUpdateEventPolicy(ctx context.Context, t EventPolicy) (EventPolicy, error)

	AutoWatchEventPolicy(in *api.ListWatchOptions, stream EventPolicyV1_AutoWatchEventPolicyServer) error
}

// ServiceEventV1Server is the server interface for the service.
type ServiceEventV1Server interface {
	AutoAddEvent(ctx context.Context, t Event) (Event, error)
	AutoDeleteEvent(ctx context.Context, t Event) (Event, error)
	AutoGetEvent(ctx context.Context, t Event) (Event, error)
	AutoListEvent(ctx context.Context, t api.ListWatchOptions) (EventList, error)
	AutoUpdateEvent(ctx context.Context, t Event) (Event, error)

	AutoWatchEvent(in *api.ListWatchOptions, stream EventV1_AutoWatchEventServer) error
}
