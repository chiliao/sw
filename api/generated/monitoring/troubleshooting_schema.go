// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package monitoringApiServer is a auto generated package.
Input file: troubleshooting.proto
*/
package monitoring

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapTroubleshooting = map[string]*api.Struct{

	"monitoring.PingPktStats": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(PingPktStats{}) },
		Fields: map[string]api.Field{
			"NoResp": api.Field{Name: "NoResp", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "no-response", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"RttMs": api.Field{Name: "RttMs", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "round-trip-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},
		},
	},
	"monitoring.PingStats": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(PingStats{}) },
		Fields: map[string]api.Field{
			"SmartNIC": api.Field{Name: "SmartNIC", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "smart-nic", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"PacketsTx": api.Field{Name: "PacketsTx", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "packets-sent", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"PacketsRx": api.Field{Name: "PacketsRx", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "packets-received", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"PacketLoss": api.Field{Name: "PacketLoss", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "packet-loss", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"MinRttMs": api.Field{Name: "MinRttMs", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "min-round-trip-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_FLOAT"},

			"MaxRttMs": api.Field{Name: "MaxRttMs", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "max-round-trip-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_FLOAT"},

			"AvgRttMs": api.Field{Name: "AvgRttMs", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "avg-round-trip-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_FLOAT"},

			"PktStats": api.Field{Name: "PktStats", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "per-packet-stats", Pointer: true, Slice: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "monitoring.PingPktStats"},
		},
	},
	"monitoring.TimeWindow": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TimeWindow{}) },
		Fields: map[string]api.Field{
			"StartTime": api.Field{Name: "StartTime", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "start-time", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"StopTime": api.Field{Name: "StopTime", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "stop-time", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},
		},
	},
	"monitoring.TraceRouteInfo": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TraceRouteInfo{}) },
		Fields: map[string]api.Field{},
	},
	"monitoring.TroubleshootingSession": &api.Struct{
		Kind: "TroubleshootingSession", APIGroup: "monitoring", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TroubleshootingSession{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "monitoring.TroubleshootingSessionSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TroubleshootingSessionStatus"},

			"Kind": api.Field{Name: "Kind", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "kind", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"APIVersion": api.Field{Name: "APIVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "api-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "name", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Tenant": api.Field{Name: "Tenant", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tenant", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Namespace": api.Field{Name: "Namespace", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "namespace", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"GenerationID": api.Field{Name: "GenerationID", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "generation-id", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"ResourceVersion": api.Field{Name: "ResourceVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "resource-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"UUID": api.Field{Name: "UUID", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "uuid", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},

			"Labels": api.Field{Name: "Labels", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "labels", Pointer: true, Slice: false, Map: true, Inline: false, FromInline: true, KeyType: "TYPE_STRING", Type: "TYPE_STRING"},

			"CreationTime": api.Field{Name: "CreationTime", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "creation-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"ModTime": api.Field{Name: "ModTime", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mod-time", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "api.Timestamp"},

			"SelfLink": api.Field{Name: "SelfLink", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "self-link", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: true, KeyType: "", Type: "TYPE_STRING"},
		},

		CLITags: map[string]api.CLIInfo{
			"api-version":      api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"enable-mirroring": api.CLIInfo{Path: "Spec.EnableMirroring", Skip: false, Insert: "", Help: ""},
			"generation-id":    api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":             api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":           api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"name":             api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":        api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"repeat-every":     api.CLIInfo{Path: "Spec.RepeatEvery", Skip: false, Insert: "", Help: ""},
			"report-url":       api.CLIInfo{Path: "Status.TsResults[].ReportURL", Skip: false, Insert: "", Help: ""},
			"resource-version": api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":        api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"state":            api.CLIInfo{Path: "Status.State", Skip: false, Insert: "", Help: ""},
			"tenant":           api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uuid":             api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"monitoring.TroubleshootingSessionSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TroubleshootingSessionSpec{}) },
		Fields: map[string]api.Field{
			"FlowSelector": api.Field{Name: "FlowSelector", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "flow-selector", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "monitoring.MatchRule"},

			"TimeWindow": api.Field{Name: "TimeWindow", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "time-window", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TimeWindow"},

			"RepeatEvery": api.Field{Name: "RepeatEvery", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "repeat-every", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"EnableMirroring": api.Field{Name: "EnableMirroring", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "enable-mirroring", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},
		},
	},
	"monitoring.TroubleshootingSessionStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TroubleshootingSessionStatus{}) },
		Fields: map[string]api.Field{
			"State": api.Field{Name: "State", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "state", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"TsResults": api.Field{Name: "TsResults", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "troubleshooting-results", Pointer: true, Slice: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "monitoring.TsResult"},
		},
	},
	"monitoring.TsAuditTrail": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TsAuditTrail{}) },
		Fields: map[string]api.Field{},
	},
	"monitoring.TsFlowCounters": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TsFlowCounters{}) },
		Fields: map[string]api.Field{},
	},
	"monitoring.TsFlowLogs": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TsFlowLogs{}) },
		Fields: map[string]api.Field{},
	},
	"monitoring.TsPolicy": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TsPolicy{}) },
		Fields: map[string]api.Field{
			"Sgpolicy": api.Field{Name: "Sgpolicy", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "sg-policy", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectRef"},

			"InRules": api.Field{Name: "InRules", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "in-rules", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGRule"},

			"OutRules": api.Field{Name: "OutRules", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "out-rules", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGRule"},
		},
	},
	"monitoring.TsReport": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TsReport{}) },
		Fields: map[string]api.Field{
			"TimeWindow": api.Field{Name: "TimeWindow", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "time-window", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TimeWindow"},

			"ReportSummary": api.Field{Name: "ReportSummary", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "report-summary", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Events": api.Field{Name: "Events", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "events", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "events.Event"},

			"Alerts": api.Field{Name: "Alerts", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "alerts", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.AlertStatus"},

			"Stats": api.Field{Name: "Stats", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "stats", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TsStats"},

			"FlowCounters": api.Field{Name: "FlowCounters", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "flow-counters", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TsFlowCounters"},

			"FlowLogs": api.Field{Name: "FlowLogs", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "flow-logs", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TsFlowLogs"},

			"AuditTrail": api.Field{Name: "AuditTrail", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "audit-trail", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TsAuditTrail"},

			"Policies": api.Field{Name: "Policies", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "policies", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "monitoring.TsPolicy"},

			"MirrorStatus": api.Field{Name: "MirrorStatus", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "mirror-session-status", Pointer: true, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "monitoring.MirrorSessionStatus"},

			"PingStats": api.Field{Name: "PingStats", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ping-stats", Pointer: true, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "monitoring.PingStats"},

			"TracedRouteInfo": api.Field{Name: "TracedRouteInfo", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "traced-route-info", Pointer: true, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "monitoring.TraceRouteInfo"},
		},
	},
	"monitoring.TsResult": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TsResult{}) },
		Fields: map[string]api.Field{
			"TimeWindow": api.Field{Name: "TimeWindow", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "time-window", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "monitoring.TimeWindow"},

			"ReportURL": api.Field{Name: "ReportURL", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "report-url", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"monitoring.TsStats": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TsStats{}) },
		Fields: map[string]api.Field{},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapTroubleshooting)
}
