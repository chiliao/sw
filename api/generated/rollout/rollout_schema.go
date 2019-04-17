// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package rolloutApiServer is a auto generated package.
Input file: rollout.proto
*/
package rollout

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapRollout = map[string]*api.Struct{

	"rollout.Rollout": &api.Struct{
		Kind: "Rollout", APIGroup: "rollout", Scopes: []string{"Cluster"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(Rollout{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "rollout.RolloutSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "rollout.RolloutStatus"},

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
			"api-version":                    api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"completion-percent":             api.CLIInfo{Path: "Status.CompletionPercentage", Skip: false, Insert: "", Help: ""},
			"duration":                       api.CLIInfo{Path: "Spec.Duration", Skip: false, Insert: "", Help: ""},
			"generation-id":                  api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":                           api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":                         api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"max-nic-failures-before-abort":  api.CLIInfo{Path: "Spec.MaxNICFailuresBeforeAbort", Skip: false, Insert: "", Help: ""},
			"max-parallel":                   api.CLIInfo{Path: "Spec.MaxParallel", Skip: false, Insert: "", Help: ""},
			"message":                        api.CLIInfo{Path: "Status.SmartNICsStatus[].Message", Skip: false, Insert: "", Help: ""},
			"name":                           api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":                      api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"phase":                          api.CLIInfo{Path: "Status.SmartNICsStatus[].Phase", Skip: false, Insert: "", Help: ""},
			"prev-version":                   api.CLIInfo{Path: "Status.PreviousVersion", Skip: false, Insert: "", Help: ""},
			"reason":                         api.CLIInfo{Path: "Status.SmartNICsStatus[].Reason", Skip: false, Insert: "", Help: ""},
			"resource-version":               api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":                      api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"smartnic-must-match-constraint": api.CLIInfo{Path: "Spec.SmartNICMustMatchConstraint", Skip: false, Insert: "", Help: ""},
			"smartnics-only":                 api.CLIInfo{Path: "Spec.SmartNICsOnly", Skip: false, Insert: "", Help: ""},
			"state":                          api.CLIInfo{Path: "Status.OperationalState", Skip: false, Insert: "", Help: ""},
			"strategy":                       api.CLIInfo{Path: "Spec.Strategy", Skip: false, Insert: "", Help: ""},
			"suspend":                        api.CLIInfo{Path: "Spec.Suspend", Skip: false, Insert: "", Help: ""},
			"tenant":                         api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"upgrade-type":                   api.CLIInfo{Path: "Spec.UpgradeType", Skip: false, Insert: "", Help: ""},
			"uuid":                           api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
			"version":                        api.CLIInfo{Path: "Spec.Version", Skip: false, Insert: "", Help: ""},
		},
	},
	"rollout.RolloutAction": &api.Struct{
		Kind: "RolloutAction", APIGroup: "rollout", Scopes: []string{"Cluster"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(RolloutAction{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "rollout.RolloutSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "rollout.RolloutActionStatus"},

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
			"api-version":                    api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"completion-percent":             api.CLIInfo{Path: "Status.CompletionPercentage", Skip: false, Insert: "", Help: ""},
			"duration":                       api.CLIInfo{Path: "Spec.Duration", Skip: false, Insert: "", Help: ""},
			"generation-id":                  api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":                           api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":                         api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"max-nic-failures-before-abort":  api.CLIInfo{Path: "Spec.MaxNICFailuresBeforeAbort", Skip: false, Insert: "", Help: ""},
			"max-parallel":                   api.CLIInfo{Path: "Spec.MaxParallel", Skip: false, Insert: "", Help: ""},
			"name":                           api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":                      api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"prev-version":                   api.CLIInfo{Path: "Status.PreviousVersion", Skip: false, Insert: "", Help: ""},
			"resource-version":               api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":                      api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"smartnic-must-match-constraint": api.CLIInfo{Path: "Spec.SmartNICMustMatchConstraint", Skip: false, Insert: "", Help: ""},
			"smartnics-only":                 api.CLIInfo{Path: "Spec.SmartNICsOnly", Skip: false, Insert: "", Help: ""},
			"state":                          api.CLIInfo{Path: "Status.OperationalState", Skip: false, Insert: "", Help: ""},
			"strategy":                       api.CLIInfo{Path: "Spec.Strategy", Skip: false, Insert: "", Help: ""},
			"suspend":                        api.CLIInfo{Path: "Spec.Suspend", Skip: false, Insert: "", Help: ""},
			"tenant":                         api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"upgrade-type":                   api.CLIInfo{Path: "Spec.UpgradeType", Skip: false, Insert: "", Help: ""},
			"uuid":                           api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
			"version":                        api.CLIInfo{Path: "Spec.Version", Skip: false, Insert: "", Help: ""},
		},
	},
	"rollout.RolloutActionStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(RolloutActionStatus{}) },
		Fields: map[string]api.Field{
			"OperationalState": api.Field{Name: "OperationalState", CLITag: api.CLIInfo{ID: "state", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "state", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"CompletionPercentage": api.Field{Name: "CompletionPercentage", CLITag: api.CLIInfo{ID: "completion-percent", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "completion-percent", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"StartTime": api.Field{Name: "StartTime", CLITag: api.CLIInfo{ID: "start-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "start-time", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"EndTime": api.Field{Name: "EndTime", CLITag: api.CLIInfo{ID: "end-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "end-time", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"PreviousVersion": api.Field{Name: "PreviousVersion", CLITag: api.CLIInfo{ID: "prev-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "prev-version", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"rollout.RolloutPhase": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(RolloutPhase{}) },
		Fields: map[string]api.Field{
			"Name": api.Field{Name: "Name", CLITag: api.CLIInfo{ID: "name", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "name", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Phase": api.Field{Name: "Phase", CLITag: api.CLIInfo{ID: "phase", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "phase", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"StartTime": api.Field{Name: "StartTime", CLITag: api.CLIInfo{ID: "start-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "start-time", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"EndTime": api.Field{Name: "EndTime", CLITag: api.CLIInfo{ID: "end-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "end-time", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"Reason": api.Field{Name: "Reason", CLITag: api.CLIInfo{ID: "reason", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "reason", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Message": api.Field{Name: "Message", CLITag: api.CLIInfo{ID: "message", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "message", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"rollout.RolloutSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(RolloutSpec{}) },
		Fields: map[string]api.Field{
			"Version": api.Field{Name: "Version", CLITag: api.CLIInfo{ID: "version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "version", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ScheduledStartTime": api.Field{Name: "ScheduledStartTime", CLITag: api.CLIInfo{ID: "scheduled-start-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "scheduled-start-time", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"Duration": api.Field{Name: "Duration", CLITag: api.CLIInfo{ID: "duration", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "duration", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Strategy": api.Field{Name: "Strategy", CLITag: api.CLIInfo{ID: "strategy", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "strategy", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"MaxParallel": api.Field{Name: "MaxParallel", CLITag: api.CLIInfo{ID: "max-parallel", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "max-parallel", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"MaxNICFailuresBeforeAbort": api.Field{Name: "MaxNICFailuresBeforeAbort", CLITag: api.CLIInfo{ID: "max-nic-failures-before-abort", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "max-nic-failures-before-abort", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"OrderConstraints": api.Field{Name: "OrderConstraints", CLITag: api.CLIInfo{ID: "order-constraints", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "order-constraints", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "labels.Selector"},

			"Suspend": api.Field{Name: "Suspend", CLITag: api.CLIInfo{ID: "suspend", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "suspend", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"SmartNICsOnly": api.Field{Name: "SmartNICsOnly", CLITag: api.CLIInfo{ID: "smartnics-only", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "smartnics-only", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"SmartNICMustMatchConstraint": api.Field{Name: "SmartNICMustMatchConstraint", CLITag: api.CLIInfo{ID: "smartnic-must-match-constraint", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "smartnic-must-match-constraint", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"UpgradeType": api.Field{Name: "UpgradeType", CLITag: api.CLIInfo{ID: "upgrade-type", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "upgrade-type", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"rollout.RolloutStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(RolloutStatus{}) },
		Fields: map[string]api.Field{
			"ControllerNodesStatus": api.Field{Name: "ControllerNodesStatus", CLITag: api.CLIInfo{ID: "controller-nodes-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "controller-nodes-status", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "rollout.RolloutPhase"},

			"ControllerServicesStatus": api.Field{Name: "ControllerServicesStatus", CLITag: api.CLIInfo{ID: "controller-services-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "controller-services-status", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "rollout.RolloutPhase"},

			"SmartNICsStatus": api.Field{Name: "SmartNICsStatus", CLITag: api.CLIInfo{ID: "smartnics-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "smartnics-status", Pointer: true, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "rollout.RolloutPhase"},

			"OperationalState": api.Field{Name: "OperationalState", CLITag: api.CLIInfo{ID: "state", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "state", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"CompletionPercentage": api.Field{Name: "CompletionPercentage", CLITag: api.CLIInfo{ID: "completion-percent", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "completion-percent", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},

			"StartTime": api.Field{Name: "StartTime", CLITag: api.CLIInfo{ID: "start-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "start-time", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"EndTime": api.Field{Name: "EndTime", CLITag: api.CLIInfo{ID: "end-time", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "end-time", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.Timestamp"},

			"PreviousVersion": api.Field{Name: "PreviousVersion", CLITag: api.CLIInfo{ID: "prev-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "prev-version", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
}

var keyMapRollout = map[string][]api.PathsMap{}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapRollout)
	schema.AddPaths(keyMapRollout)
}
