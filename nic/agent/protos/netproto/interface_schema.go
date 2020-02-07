// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package netprotoApiServer is a auto generated package.
Input file: interface.proto
*/
package netproto

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapInterface = map[string]*api.Struct{

	"netproto.Interface": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(Interface{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "TypeMeta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.InterfaceSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.InterfaceStatus"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{ID: "kind", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{ID: "api-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{ID: "name", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "name", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{ID: "tenant", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Namespace": api.Field{Name: "Namespace", CLITag: api.CLIInfo{ID: "namespace", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "namespace", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"GenerationID": api.Field{Name: "GenerationID", CLITag: api.CLIInfo{ID: "generation-id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "generation-id", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"ResourceVersion": api.Field{Name: "ResourceVersion", CLITag: api.CLIInfo{ID: "resource-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-version", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"UUID": api.Field{Name: "UUID", CLITag: api.CLIInfo{ID: "uuid", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uuid", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Labels": api.Field{Name: "Labels", CLITag: api.CLIInfo{ID: "labels", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "labels", Pointer: true, Slice: false, Mutable: true, Map: true, Inline: false, FromInline: true, KeyType: "TYPE_STRING", Type: "TYPE_STRING"},

			"CreationTime": api.Field{Name: "CreationTime", CLITag: api.CLIInfo{ID: "creation-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "creation-time", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"ModTime": api.Field{Name: "ModTime", CLITag: api.CLIInfo{ID: "mod-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mod-time", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"SelfLink": api.Field{Name: "SelfLink", CLITag: api.CLIInfo{ID: "self-link", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "self-link", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},
		},

		CLITags: map[string]api.CLIInfo{
			"admin-status":     api.CLIInfo{Path: "Spec.AdminStatus", Skip: false, Insert: "", Help: ""},
			"api-version":      api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"cable-type":       api.CLIInfo{Path: "Status.IFUplinkStatus.TransceiverStatus.TranceiverCableType", Skip: false, Insert: "", Help: ""},
			"dsc":              api.CLIInfo{Path: "Status.Name", Skip: false, Insert: "", Help: ""},
			"generation-id":    api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"host-ifname":      api.CLIInfo{Path: "Status.IFHostStatus.HostIfName", Skip: false, Insert: "", Help: ""},
			"id":               api.CLIInfo{Path: "Status.InterfaceID", Skip: false, Insert: "", Help: ""},
			"if-uuid":          api.CLIInfo{Path: "Status.InterfaceUUID", Skip: false, Insert: "", Help: ""},
			"ip-address":       api.CLIInfo{Path: "Spec.IPAddress", Skip: false, Insert: "", Help: ""},
			"kind":             api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":           api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"link-speed":       api.CLIInfo{Path: "Status.IFUplinkStatus.LinkSpeed", Skip: false, Insert: "", Help: ""},
			"mirror-enabled":   api.CLIInfo{Path: "Status.MirrorEnabled", Skip: false, Insert: "", Help: ""},
			"mtu":              api.CLIInfo{Path: "Spec.MTU", Skip: false, Insert: "", Help: ""},
			"name":             api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":        api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"network":          api.CLIInfo{Path: "Spec.Network", Skip: false, Insert: "", Help: ""},
			"oper-status":      api.CLIInfo{Path: "Status.OperStatus", Skip: false, Insert: "", Help: ""},
			"pid":              api.CLIInfo{Path: "Status.IFUplinkStatus.TransceiverStatus.TranceiverPid", Skip: false, Insert: "", Help: ""},
			"primary-mac":      api.CLIInfo{Path: "Status.PrimaryMac", Skip: false, Insert: "", Help: ""},
			"resource-version": api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"rx-pause-enabled": api.CLIInfo{Path: "Spec.Pause.RxPauseEnabled", Skip: false, Insert: "", Help: ""},
			"self-link":        api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"speed":            api.CLIInfo{Path: "Spec.Speed", Skip: false, Insert: "", Help: ""},
			"state":            api.CLIInfo{Path: "Status.IFUplinkStatus.TransceiverStatus.TransceiverState", Skip: false, Insert: "", Help: ""},
			"tenant":           api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"tx-pause-enabled": api.CLIInfo{Path: "Spec.Pause.TxPauseEnabled", Skip: false, Insert: "", Help: ""},
			"type":             api.CLIInfo{Path: "Spec.Type", Skip: false, Insert: "", Help: ""},
			"uplink-port-id":   api.CLIInfo{Path: "Status.IFUplinkStatus.PortID", Skip: false, Insert: "", Help: ""},
			"uuid":             api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
			"vrf-name":         api.CLIInfo{Path: "Spec.VrfName", Skip: false, Insert: "", Help: ""},
		},
	},
	"netproto.InterfaceEvent": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(InterfaceEvent{}) },
		Fields: map[string]api.Field{
			"EventType": api.Field{Name: "EventType", CLITag: api.CLIInfo{ID: "event-type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "event-type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_ENUM"},

			"Interface": api.Field{Name: "Interface", CLITag: api.CLIInfo{ID: "interface", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "interface", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.Interface"},
		},
	},
	"netproto.InterfaceEventList": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(InterfaceEventList{}) },
		Fields: map[string]api.Field{
			"InterfaceEvents": api.Field{Name: "InterfaceEvents", CLITag: api.CLIInfo{ID: "InterfaceEvents", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.InterfaceEvent"},
		},
	},
	"netproto.InterfaceHostStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(InterfaceHostStatus{}) },
		Fields: map[string]api.Field{
			"HostIfName": api.Field{Name: "HostIfName", CLITag: api.CLIInfo{ID: "host-ifname", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "host-ifname", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"netproto.InterfaceList": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(InterfaceList{}) },
		Fields: map[string]api.Field{
			"interfaces": api.Field{Name: "interfaces", CLITag: api.CLIInfo{ID: "interfaces", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.Interface"},
		},
	},
	"netproto.InterfaceSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(InterfaceSpec{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{ID: "type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"AdminStatus": api.Field{Name: "AdminStatus", CLITag: api.CLIInfo{ID: "admin-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "admin-status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"VrfName": api.Field{Name: "VrfName", CLITag: api.CLIInfo{ID: "vrf-name", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "vrf-name", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Speed": api.Field{Name: "Speed", CLITag: api.CLIInfo{ID: "speed", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "speed", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"MTU": api.Field{Name: "MTU", CLITag: api.CLIInfo{ID: "mtu", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mtu", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"Pause": api.Field{Name: "Pause", CLITag: api.CLIInfo{ID: "pause", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "pause", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.PauseSpec"},

			"IPAddress": api.Field{Name: "IPAddress", CLITag: api.CLIInfo{ID: "ip-address", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ip-address", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Network": api.Field{Name: "Network", CLITag: api.CLIInfo{ID: "network", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "network", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TxMirrorSessions": api.Field{Name: "TxMirrorSessions", CLITag: api.CLIInfo{ID: "TxMirrorSessions", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.MirrorSession"},

			"RxMirrorSessions": api.Field{Name: "RxMirrorSessions", CLITag: api.CLIInfo{ID: "RxMirrorSessions", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.MirrorSession"},
		},
	},
	"netproto.InterfaceStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(InterfaceStatus{}) },
		Fields: map[string]api.Field{
			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{ID: "dsc", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "dsc", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"InterfaceID": api.Field{Name: "InterfaceID", CLITag: api.CLIInfo{ID: "id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "id", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT64"},

			"DSC": api.Field{Name: "DSC", CLITag: api.CLIInfo{ID: "dsc", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "dsc", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"OperStatus": api.Field{Name: "OperStatus", CLITag: api.CLIInfo{ID: "oper-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "oper-status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"PrimaryMac": api.Field{Name: "PrimaryMac", CLITag: api.CLIInfo{ID: "primary-mac", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "primary-mac", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"IFHostStatus": api.Field{Name: "IFHostStatus", CLITag: api.CLIInfo{ID: "if-host-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "if-host-status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.InterfaceHostStatus"},

			"IFUplinkStatus": api.Field{Name: "IFUplinkStatus", CLITag: api.CLIInfo{ID: "if-uplink-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "if-uplink-status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.InterfaceUplinkStatus"},

			"MirrorEnabled": api.Field{Name: "MirrorEnabled", CLITag: api.CLIInfo{ID: "mirror-enabled", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mirror-enabled", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"InterfaceUUID": api.Field{Name: "InterfaceUUID", CLITag: api.CLIInfo{ID: "if-uuid", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "if-uuid", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"netproto.InterfaceUplinkStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(InterfaceUplinkStatus{}) },
		Fields: map[string]api.Field{
			"LinkSpeed": api.Field{Name: "LinkSpeed", CLITag: api.CLIInfo{ID: "link-speed", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "link-speed", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TransceiverStatus": api.Field{Name: "TransceiverStatus", CLITag: api.CLIInfo{ID: "transceiver-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "transceiver-status", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "netproto.TransceiverStatus"},

			"PortID": api.Field{Name: "PortID", CLITag: api.CLIInfo{ID: "uplink-port-id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uplink-port-id", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},
		},
	},
	"netproto.PauseSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(PauseSpec{}) },
		Fields: map[string]api.Field{
			"Type": api.Field{Name: "Type", CLITag: api.CLIInfo{ID: "type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TxPauseEnabled": api.Field{Name: "TxPauseEnabled", CLITag: api.CLIInfo{ID: "tx-pause-enabled", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tx-pause-enabled", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"RxPauseEnabled": api.Field{Name: "RxPauseEnabled", CLITag: api.CLIInfo{ID: "rx-pause-enabled", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "rx-pause-enabled", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},
		},
	},
	"netproto.TransceiverStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TransceiverStatus{}) },
		Fields: map[string]api.Field{
			"TransceiverState": api.Field{Name: "TransceiverState", CLITag: api.CLIInfo{ID: "state", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "state", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TranceiverCableType": api.Field{Name: "TranceiverCableType", CLITag: api.CLIInfo{ID: "cable-type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "cable-type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TranceiverPid": api.Field{Name: "TranceiverPid", CLITag: api.CLIInfo{ID: "pid", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "pid", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
}

var keyMapInterface = map[string][]api.PathsMap{}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapInterface)
	schema.AddPaths(keyMapInterface)
}
