// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoringApiServer is a auto generated package.
Input file: telemetry.proto
*/
package monitoring

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapTelemetry = map[string]*api.Struct{

	"monitoring.FlowExportPolicy": &api.Struct{
		Kind: "FlowExportPolicy", APIGroup: "monitoring", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(FlowExportPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FlowExportPolicySpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FlowExportPolicyStatus"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{ID: "kind", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{ID: "api-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{ID: "name", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "name", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{ID: "tenant", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Namespace": api.Field{Name: "Namespace", CLITag: api.CLIInfo{ID: "namespace", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "namespace", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"GenerationID": api.Field{Name: "GenerationID", CLITag: api.CLIInfo{ID: "generation-id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "generation-id", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"ResourceVersion": api.Field{Name: "ResourceVersion", CLITag: api.CLIInfo{ID: "resource-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"UUID": api.Field{Name: "UUID", CLITag: api.CLIInfo{ID: "uuid", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uuid", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Labels": api.Field{Name: "Labels", CLITag: api.CLIInfo{ID: "labels", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "labels", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: true, KeyType: "TYPE_STRING", Type: "TYPE_STRING"},

			"CreationTime": api.Field{Name: "CreationTime", CLITag: api.CLIInfo{ID: "creation-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "creation-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"ModTime": api.Field{Name: "ModTime", CLITag: api.CLIInfo{ID: "mod-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mod-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"SelfLink": api.Field{Name: "SelfLink", CLITag: api.CLIInfo{ID: "self-link", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "self-link", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version":      api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"format":           api.CLIInfo{Path: "Spec.Format", Skip: false, Insert: "", Help: ""},
			"generation-id":    api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"interval":         api.CLIInfo{Path: "Spec.Interval", Skip: false, Insert: "", Help: ""},
			"kind":             api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":           api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"name":             api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":        api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"resource-version": api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":        api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"tenant":           api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uuid":             api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"monitoring.FlowExportPolicySpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(FlowExportPolicySpec{}) },
		Fields: map[string]api.Field{
			"Interval": api.Field{Name: "Interval", CLITag: api.CLIInfo{ID: "interval", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "interval", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Format": api.Field{Name: "Format", CLITag: api.CLIInfo{ID: "format", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "format", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"MatchRules": api.Field{Name: "MatchRules", CLITag: api.CLIInfo{ID: "match-rules", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "match-rules", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.MatchRule"},

			"Exports": api.Field{Name: "Exports", CLITag: api.CLIInfo{ID: "exports", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "exports", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.ExportConfig"},
		},
	},
<<<<<<< HEAD
	"monitoring.FlowExportPolicyStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(FlowExportPolicyStatus{}) },
		Fields: map[string]api.Field{},
	},
	"monitoring.FwlogExport": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogExport{}) },
		Fields: map[string]api.Field{
			"Targets": api.Field{Name: "Targets", CLITag: api.CLIInfo{ID: "targets", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "targets", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.ExportConfig"},

			"Format": api.Field{Name: "Format", CLITag: api.CLIInfo{ID: "format", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "format", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Filter": api.Field{Name: "Filter", CLITag: api.CLIInfo{ID: "export-filter", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "export-filter", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"SyslogConfig": api.Field{Name: "SyslogConfig", CLITag: api.CLIInfo{ID: "syslog-config", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "syslog-config", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.SyslogExportConfig"},
		},
	},
=======
>>>>>>> generated changes
	"monitoring.FwlogPolicy": &api.Struct{
		Kind: "FwlogPolicy", APIGroup: "monitoring", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FwlogPolicySpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.FwlogPolicyStatus"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{ID: "kind", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{ID: "api-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{ID: "name", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "name", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{ID: "tenant", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Namespace": api.Field{Name: "Namespace", CLITag: api.CLIInfo{ID: "namespace", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "namespace", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"GenerationID": api.Field{Name: "GenerationID", CLITag: api.CLIInfo{ID: "generation-id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "generation-id", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"ResourceVersion": api.Field{Name: "ResourceVersion", CLITag: api.CLIInfo{ID: "resource-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"UUID": api.Field{Name: "UUID", CLITag: api.CLIInfo{ID: "uuid", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uuid", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Labels": api.Field{Name: "Labels", CLITag: api.CLIInfo{ID: "labels", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "labels", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: true, KeyType: "TYPE_STRING", Type: "TYPE_STRING"},

			"CreationTime": api.Field{Name: "CreationTime", CLITag: api.CLIInfo{ID: "creation-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "creation-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"ModTime": api.Field{Name: "ModTime", CLITag: api.CLIInfo{ID: "mod-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mod-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"SelfLink": api.Field{Name: "SelfLink", CLITag: api.CLIInfo{ID: "self-link", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "self-link", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version":      api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"filter":           api.CLIInfo{Path: "Spec.Filter", Skip: false, Insert: "", Help: ""},
			"format":           api.CLIInfo{Path: "Spec.Format", Skip: false, Insert: "", Help: ""},
			"generation-id":    api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":             api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":           api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"name":             api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":        api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"resource-version": api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":        api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"tenant":           api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uuid":             api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"monitoring.FwlogPolicySpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogPolicySpec{}) },
		Fields: map[string]api.Field{
			"Targets": api.Field{Name: "Targets", CLITag: api.CLIInfo{ID: "targets", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "targets", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.ExportConfig"},

			"Format": api.Field{Name: "Format", CLITag: api.CLIInfo{ID: "format", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "format", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Filter": api.Field{Name: "Filter", CLITag: api.CLIInfo{ID: "filter", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "filter", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Config": api.Field{Name: "Config", CLITag: api.CLIInfo{ID: "config", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "config", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.SyslogExportConfig"},
		},
	},
	"monitoring.FwlogPolicyStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(FwlogPolicyStatus{}) },
		Fields: map[string]api.Field{},
	},
	"monitoring.StatsPolicy": &api.Struct{
		Kind: "StatsPolicy", APIGroup: "monitoring", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(StatsPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.StatsPolicySpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.StatsPolicyStatus"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{ID: "kind", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{ID: "api-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{ID: "name", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "name", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{ID: "tenant", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Namespace": api.Field{Name: "Namespace", CLITag: api.CLIInfo{ID: "namespace", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "namespace", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"GenerationID": api.Field{Name: "GenerationID", CLITag: api.CLIInfo{ID: "generation-id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "generation-id", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"ResourceVersion": api.Field{Name: "ResourceVersion", CLITag: api.CLIInfo{ID: "resource-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"UUID": api.Field{Name: "UUID", CLITag: api.CLIInfo{ID: "uuid", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uuid", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Labels": api.Field{Name: "Labels", CLITag: api.CLIInfo{ID: "labels", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "labels", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: true, KeyType: "TYPE_STRING", Type: "TYPE_STRING"},

			"CreationTime": api.Field{Name: "CreationTime", CLITag: api.CLIInfo{ID: "creation-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "creation-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"ModTime": api.Field{Name: "ModTime", CLITag: api.CLIInfo{ID: "mod-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mod-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"SelfLink": api.Field{Name: "SelfLink", CLITag: api.CLIInfo{ID: "self-link", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "self-link", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version":               api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"downsample-retention-time": api.CLIInfo{Path: "Spec.DownSampleRetentionTime", Skip: false, Insert: "", Help: ""},
			"generation-id":             api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":                      api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":                    api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"name":                      api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":                 api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"resource-version":          api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"retention-time":            api.CLIInfo{Path: "Spec.RetentionTime", Skip: false, Insert: "", Help: ""},
			"self-link":                 api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"tenant":                    api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uuid":                      api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"monitoring.StatsPolicySpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(StatsPolicySpec{}) },
		Fields: map[string]api.Field{
			"RetentionTime": api.Field{Name: "RetentionTime", CLITag: api.CLIInfo{ID: "retention-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "retention-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"DownSampleRetentionTime": api.Field{Name: "DownSampleRetentionTime", CLITag: api.CLIInfo{ID: "downsample-retention-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "downsample-retention-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"monitoring.StatsPolicyStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(StatsPolicyStatus{}) },
		Fields: map[string]api.Field{},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapTelemetry)
}
