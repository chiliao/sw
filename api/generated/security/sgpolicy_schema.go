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

	"security.SGPolicy": &api.Struct{
		Kind: "SGPolicy", APIGroup: "security", Scopes: []string{"Tenant"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGPolicy{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGPolicySpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGPolicyStatus"},

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
			"action":               api.CLIInfo{Path: "Spec.Rules[].Action", Skip: false, Insert: "", Help: ""},
			"api-version":          api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"apps":                 api.CLIInfo{Path: "Spec.Rules[].Apps", Skip: false, Insert: "", Help: ""},
			"attach-groups":        api.CLIInfo{Path: "Spec.AttachGroups", Skip: false, Insert: "", Help: ""},
			"attach-tenant":        api.CLIInfo{Path: "Spec.AttachTenant", Skip: false, Insert: "", Help: ""},
			"from-ip-addresses":    api.CLIInfo{Path: "Spec.Rules[].FromIPAddresses", Skip: false, Insert: "", Help: ""},
			"from-security-groups": api.CLIInfo{Path: "Spec.Rules[].FromSecurityGroups", Skip: false, Insert: "", Help: ""},
			"generation-id":        api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":                 api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":               api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"min-version":          api.CLIInfo{Path: "Status.PropagationStatus.MinVersion", Skip: false, Insert: "", Help: ""},
			"name":                 api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":            api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"pending":              api.CLIInfo{Path: "Status.PropagationStatus.Pending", Skip: false, Insert: "", Help: ""},
			"resource-version":     api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":            api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"tenant":               api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"to-ip-addresses":      api.CLIInfo{Path: "Spec.Rules[].ToIPAddresses", Skip: false, Insert: "", Help: ""},
			"to-security-groups":   api.CLIInfo{Path: "Spec.Rules[].ToSecurityGroups", Skip: false, Insert: "", Help: ""},
			"updated":              api.CLIInfo{Path: "Status.PropagationStatus.Updated", Skip: false, Insert: "", Help: ""},
			"uuid":                 api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
			"workloads":            api.CLIInfo{Path: "Status.Workloads", Skip: false, Insert: "", Help: ""},
		},
	},
	"security.SGPolicyPropagationStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGPolicyPropagationStatus{}) },
		Fields: map[string]api.Field{
			"GenerationID": api.Field{Name: "GenerationID", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "generation-id", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Updated": api.Field{Name: "Updated", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "updated", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT32"},

			"Pending": api.Field{Name: "Pending", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "pending", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_INT32"},

			"MinVersion": api.Field{Name: "MinVersion", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "min-version", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"security.SGPolicySpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGPolicySpec{}) },
		Fields: map[string]api.Field{
			"AttachGroups": api.Field{Name: "AttachGroups", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "attach-groups", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"AttachTenant": api.Field{Name: "AttachTenant", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "attach-tenant", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},

			"Rules": api.Field{Name: "Rules", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "rules", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGRule"},
		},
	},
	"security.SGPolicyStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGPolicyStatus{}) },
		Fields: map[string]api.Field{
			"Workloads": api.Field{Name: "Workloads", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "workloads", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"PropagationStatus": api.Field{Name: "PropagationStatus", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "propagation-status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SGPolicyPropagationStatus"},
		},
	},
	"security.SGRule": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SGRule{}) },
		Fields: map[string]api.Field{
			"Apps": api.Field{Name: "Apps", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "apps", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"Action": api.Field{Name: "Action", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "action", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"FromIPAddresses": api.Field{Name: "FromIPAddresses", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "from-ip-addresses", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ToIPAddresses": api.Field{Name: "ToIPAddresses", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "to-ip-addresses", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"FromSecurityGroups": api.Field{Name: "FromSecurityGroups", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "from-security-groups", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ToSecurityGroups": api.Field{Name: "ToSecurityGroups", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "to-security-groups", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapSgpolicy)
}
