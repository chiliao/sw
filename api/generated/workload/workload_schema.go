// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package workloadApiServer is a auto generated package.
Input file: workload.proto
*/
package workload

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapWorkload = map[string]*api.Struct{

	"workload.Workload": &api.Struct{
		Kind: "Workload", APIGroup: "workload", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(Workload{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "workload.WorkloadSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "workload.WorkloadStatus"},

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
			"api-version":       api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"endpoint":          api.CLIInfo{Path: "Status.Interfaces[].Endpoint", Skip: false, Insert: "", Help: ""},
			"external-vlan":     api.CLIInfo{Path: "Spec.Interfaces[].ExternalVlan", Skip: false, Insert: "", Help: "External vlan associated with the workload"},
			"generation-id":     api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"host-name":         api.CLIInfo{Path: "Spec.HostName", Skip: false, Insert: "", Help: "Host name where the workload runs"},
			"ip-addresses":      api.CLIInfo{Path: "Spec.Interfaces[].IpAddresses", Skip: false, Insert: "", Help: ""},
			"kind":              api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":            api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"mac-address":       api.CLIInfo{Path: "Spec.Interfaces[].MACAddress", Skip: false, Insert: "", Help: "MAC address of the interface as seen by the workload"},
			"micro-seg-vlan":    api.CLIInfo{Path: "Spec.Interfaces[].MicroSegVlan", Skip: false, Insert: "", Help: "Vlan identifying host unique vlan id"},
			"migration-timeout": api.CLIInfo{Path: "Spec.MigrationTimeout", Skip: false, Insert: "", Help: ""},
			"name":              api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":         api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"network":           api.CLIInfo{Path: "Spec.Interfaces[].Network", Skip: false, Insert: "", Help: "Network this interface will belong to"},
			"resource-version":  api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":         api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"stage":             api.CLIInfo{Path: "Status.MigrationStatus.Stage", Skip: false, Insert: "", Help: ""},
			"status":            api.CLIInfo{Path: "Status.MigrationStatus.Status", Skip: false, Insert: "", Help: ""},
			"tenant":            api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uuid":              api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"workload.WorkloadIntfSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(WorkloadIntfSpec{}) },
		Fields: map[string]api.Field{
			"MACAddress": api.Field{Name: "MACAddress", CLITag: api.CLIInfo{ID: "mac-address", Path: "", Skip: false, Insert: "", Help: "MAC address of the interface as seen by the workload"}, JSONTag: "mac-address", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"MicroSegVlan": api.Field{Name: "MicroSegVlan", CLITag: api.CLIInfo{ID: "micro-seg-vlan", Path: "", Skip: false, Insert: "", Help: "Vlan identifying host unique vlan id"}, JSONTag: "micro-seg-vlan", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"ExternalVlan": api.Field{Name: "ExternalVlan", CLITag: api.CLIInfo{ID: "external-vlan", Path: "", Skip: false, Insert: "", Help: "External vlan associated with the workload"}, JSONTag: "external-vlan", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"IpAddresses": api.Field{Name: "IpAddresses", CLITag: api.CLIInfo{ID: "ip-addresses", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ip-addresses", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Network": api.Field{Name: "Network", CLITag: api.CLIInfo{ID: "network", Path: "", Skip: false, Insert: "", Help: "Network this interface will belong to"}, JSONTag: "network", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"workload.WorkloadIntfStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(WorkloadIntfStatus{}) },
		Fields: map[string]api.Field{
			"IpAddresses": api.Field{Name: "IpAddresses", CLITag: api.CLIInfo{ID: "ip-addresses", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ip-addresses", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Endpoint": api.Field{Name: "Endpoint", CLITag: api.CLIInfo{ID: "endpoint", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "endpoint", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"MicroSegVlan": api.Field{Name: "MicroSegVlan", CLITag: api.CLIInfo{ID: "micro-seg-vlan", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "micro-seg-vlan", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"MACAddress": api.Field{Name: "MACAddress", CLITag: api.CLIInfo{ID: "mac-address", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mac-address", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ExternalVlan": api.Field{Name: "ExternalVlan", CLITag: api.CLIInfo{ID: "external-vlan", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "external-vlan", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"Network": api.Field{Name: "Network", CLITag: api.CLIInfo{ID: "network", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "network", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"workload.WorkloadMigrationStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(WorkloadMigrationStatus{}) },
		Fields: map[string]api.Field{
			"Stage": api.Field{Name: "Stage", CLITag: api.CLIInfo{ID: "stage", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "stage", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"StartedAt": api.Field{Name: "StartedAt", CLITag: api.CLIInfo{ID: "started-at", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "started-at", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"CompletedAt": api.Field{Name: "CompletedAt", CLITag: api.CLIInfo{ID: "completed-at", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "completed-at", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},
		},
	},
	"workload.WorkloadSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(WorkloadSpec{}) },
		Fields: map[string]api.Field{
			"HostName": api.Field{Name: "HostName", CLITag: api.CLIInfo{ID: "host-name", Path: "", Skip: false, Insert: "", Help: "Host name where the workload runs"}, JSONTag: "host-name", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Interfaces": api.Field{Name: "Interfaces", CLITag: api.CLIInfo{ID: "interfaces", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "interfaces", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "workload.WorkloadIntfSpec"},

			"MigrationTimeout": api.Field{Name: "MigrationTimeout", CLITag: api.CLIInfo{ID: "migration-timeout", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "migration-timeout", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"workload.WorkloadStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(WorkloadStatus{}) },
		Fields: map[string]api.Field{
			"PropagationStatus": api.Field{Name: "PropagationStatus", CLITag: api.CLIInfo{ID: "propagation-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "propagation-status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.PropagationStatus"},

			"Interfaces": api.Field{Name: "Interfaces", CLITag: api.CLIInfo{ID: "interfaces", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "interfaces", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "workload.WorkloadIntfStatus"},

			"HostName": api.Field{Name: "HostName", CLITag: api.CLIInfo{ID: "host-name", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "host-name", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"MigrationStatus": api.Field{Name: "MigrationStatus", CLITag: api.CLIInfo{ID: "migration-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "migration-status", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "workload.WorkloadMigrationStatus"},
		},
	},
}

var keyMapWorkload = map[string][]api.PathsMap{}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapWorkload)
	schema.AddPaths(keyMapWorkload)
}
