// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package securityApiServer is a auto generated package.
Input file: sgpolicy.proto
*/
package security

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapSgpolicy = map[string]*api.Struct{

	"security.PropagationStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(PropagationStatus{}) },
		Fields: map[string]api.Field{
			"GenerationID": api.Field{Name: "GenerationID", CLITag: api.CLIInfo{ID: "generation-id", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "generation-id", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Updated": api.Field{Name: "Updated", CLITag: api.CLIInfo{ID: "updated", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "updated", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT32"},

			"Pending": api.Field{Name: "Pending", CLITag: api.CLIInfo{ID: "pending", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "pending", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT32"},

			"MinVersion": api.Field{Name: "MinVersion", CLITag: api.CLIInfo{ID: "min-version", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "min-version", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"PendingNaples": api.Field{Name: "PendingNaples", CLITag: api.CLIInfo{ID: "pending-smartnics", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "pending-smartnics", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"security.ProtoPort": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(ProtoPort{}) },
		Fields: map[string]api.Field{
			"Protocol": api.Field{Name: "Protocol", CLITag: api.CLIInfo{ID: "protocol", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "protocol", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Ports": api.Field{Name: "Ports", CLITag: api.CLIInfo{ID: "ports", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ports", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"security.SGPolicy": &api.Struct{
		Kind: "SGPolicy", APIGroup: "security", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{ID: "T", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{ID: "meta", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{ID: "spec", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGPolicySpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{ID: "status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGPolicyStatus"},

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
			"action":               api.CLIInfo{Path: "Spec.Rules[].Action", Skip: false, Insert: "", Help: ""},
			"api-version":          api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"apps":                 api.CLIInfo{Path: "Spec.Rules[].Apps", Skip: false, Insert: "", Help: ""},
			"attach-groups":        api.CLIInfo{Path: "Spec.AttachGroups", Skip: false, Insert: "", Help: ""},
			"attach-tenant":        api.CLIInfo{Path: "Spec.AttachTenant", Skip: false, Insert: "", Help: ""},
			"from-ip":              api.CLIInfo{Path: "Spec.Rules[].FromIPAddresses", Skip: false, Insert: "", Help: ""},
			"from-security-groups": api.CLIInfo{Path: "Spec.Rules[].FromSecurityGroups", Skip: false, Insert: "", Help: ""},
			"generation-id":        api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":                 api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":               api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"min-version":          api.CLIInfo{Path: "Status.PropagationStatus.MinVersion", Skip: false, Insert: "", Help: ""},
			"name":                 api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":            api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"pending":              api.CLIInfo{Path: "Status.PropagationStatus.Pending", Skip: false, Insert: "", Help: ""},
			"pending-smartnics":    api.CLIInfo{Path: "Status.PropagationStatus.PendingNaples", Skip: false, Insert: "", Help: ""},
			"ports":                api.CLIInfo{Path: "Spec.Rules[].ProtoPorts[].Ports", Skip: false, Insert: "", Help: ""},
			"protocol":             api.CLIInfo{Path: "Spec.Rules[].ProtoPorts[].Protocol", Skip: false, Insert: "", Help: ""},
			"resource-version":     api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"rule-hash":            api.CLIInfo{Path: "Status.RuleStatus[].RuleHash", Skip: false, Insert: "", Help: ""},
			"self-link":            api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"status":               api.CLIInfo{Path: "Status.PropagationStatus.Status", Skip: false, Insert: "", Help: ""},
			"tenant":               api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"to-ip":                api.CLIInfo{Path: "Spec.Rules[].ToIPAddresses", Skip: false, Insert: "", Help: ""},
			"to-security-groups":   api.CLIInfo{Path: "Spec.Rules[].ToSecurityGroups", Skip: false, Insert: "", Help: ""},
			"updated":              api.CLIInfo{Path: "Status.PropagationStatus.Updated", Skip: false, Insert: "", Help: ""},
			"uuid":                 api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"security.SGPolicySpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGPolicySpec{}) },
		Fields: map[string]api.Field{
			"AttachGroups": api.Field{Name: "AttachGroups", CLITag: api.CLIInfo{ID: "attach-groups", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "attach-groups", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"AttachTenant": api.Field{Name: "AttachTenant", CLITag: api.CLIInfo{ID: "attach-tenant", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "attach-tenant", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"Rules": api.Field{Name: "Rules", CLITag: api.CLIInfo{ID: "rules", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "rules", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGRule"},
		},
	},
	"security.SGPolicyStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGPolicyStatus{}) },
		Fields: map[string]api.Field{
			"PropagationStatus": api.Field{Name: "PropagationStatus", CLITag: api.CLIInfo{ID: "propagation-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "propagation-status", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.PropagationStatus"},

			"RuleStatus": api.Field{Name: "RuleStatus", CLITag: api.CLIInfo{ID: "rule-status", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "rule-status", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGRuleStatus"},
		},
	},
	"security.SGRule": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGRule{}) },
		Fields: map[string]api.Field{
			"Apps": api.Field{Name: "Apps", CLITag: api.CLIInfo{ID: "apps", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "apps", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ProtoPorts": api.Field{Name: "ProtoPorts", CLITag: api.CLIInfo{ID: "proto-ports", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "proto-ports", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.ProtoPort"},

			"Action": api.Field{Name: "Action", CLITag: api.CLIInfo{ID: "action", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "action", Pointer: true, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"FromIPAddresses": api.Field{Name: "FromIPAddresses", CLITag: api.CLIInfo{ID: "from-ip", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "from-ip-addresses", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ToIPAddresses": api.Field{Name: "ToIPAddresses", CLITag: api.CLIInfo{ID: "to-ip", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "to-ip-addresses", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"FromSecurityGroups": api.Field{Name: "FromSecurityGroups", CLITag: api.CLIInfo{ID: "from-security-groups", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "from-security-groups", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ToSecurityGroups": api.Field{Name: "ToSecurityGroups", CLITag: api.CLIInfo{ID: "to-security-groups", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "to-security-groups", Pointer: false, Slice: true, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"security.SGRuleStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGRuleStatus{}) },
		Fields: map[string]api.Field{
			"RuleHash": api.Field{Name: "RuleHash", CLITag: api.CLIInfo{ID: "rule-hash", Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "rule-hash", Pointer: false, Slice: false, Mutable: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
}

var keyMapSgpolicy = map[string][]api.PathsMap{}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapSgpolicy)
	schema.AddPaths(keyMapSgpolicy)
}
