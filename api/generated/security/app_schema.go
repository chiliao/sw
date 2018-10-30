// Code generated by protoc-gen-grpc-pensando DO NOT EDIT.

/*
Package securityApiServer is a auto generated package.
Input file: app.proto
*/
package security

import (
	"reflect"

	"github.com/pensando/sw/api"
	"github.com/pensando/sw/venice/utils/runtime"
)

var typesMapApp = map[string]*api.Struct{

	"security.ALG": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(ALG{}) },
		Fields: map[string]api.Field{
			"DNS": api.Field{Name: "DNS", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "dns", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.DNS"},

			"SIP": api.Field{Name: "SIP", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "sip", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SIP"},

			"SunRPC": api.Field{Name: "SunRPC", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "sunrpc", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.SunRPC"},

			"FTP": api.Field{Name: "FTP", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "ftp", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.FTP"},

			"MSRPC": api.Field{Name: "MSRPC", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "msrpc", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.MSRPC"},

			"TFTP": api.Field{Name: "TFTP", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "tftp", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.TFTP"},

			"RSTP": api.Field{Name: "RSTP", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "rstp", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.RSTP"},
		},
	},
	"security.App": &api.Struct{
		Kind: "App", APIGroup: "security", Scopes: []string{"Cluster"}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(App{}) },
		Fields: map[string]api.Field{
			"TypeMeta": api.Field{Name: "TypeMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "", Pointer: false, Slice: false, Map: false, Inline: true, FromInline: false, KeyType: "", Type: "api.TypeMeta"},

			"ObjectMeta": api.Field{Name: "ObjectMeta", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "meta", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "api.ObjectMeta"},

			"Spec": api.Field{Name: "Spec", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "spec", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.AppSpec"},

			"Status": api.Field{Name: "Status", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "status", Pointer: false, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.AppStatus"},

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
			"api-version":                 api.CLIInfo{Path: "APIVersion", Skip: false, Insert: "", Help: ""},
			"attached-policies":           api.CLIInfo{Path: "Status.AttachedPolicies", Skip: false, Insert: "", Help: ""},
			"drop-multi-question-packets": api.CLIInfo{Path: "Spec.ALG.DNS.DropMultiQuestionPackets", Skip: false, Insert: "", Help: ""},
			"generation-id":               api.CLIInfo{Path: "GenerationID", Skip: false, Insert: "", Help: ""},
			"kind":                        api.CLIInfo{Path: "Kind", Skip: false, Insert: "", Help: ""},
			"labels":                      api.CLIInfo{Path: "Labels", Skip: false, Insert: "", Help: ""},
			"map-entry-timeout":           api.CLIInfo{Path: "Spec.ALG.SunRPC.MapEntryTimeout", Skip: false, Insert: "", Help: ""},
			"max-call-duration":           api.CLIInfo{Path: "Spec.ALG.SIP.MaxCallDuration", Skip: false, Insert: "", Help: ""},
			"name":                        api.CLIInfo{Path: "Name", Skip: false, Insert: "", Help: ""},
			"namespace":                   api.CLIInfo{Path: "Namespace", Skip: false, Insert: "", Help: ""},
			"program-id":                  api.CLIInfo{Path: "Spec.ALG.SunRPC.ProgramID", Skip: false, Insert: "", Help: ""},
			"protocol":                    api.CLIInfo{Path: "Spec.Protocol", Skip: false, Insert: "", Help: ""},
			"resource-version":            api.CLIInfo{Path: "ResourceVersion", Skip: false, Insert: "", Help: ""},
			"self-link":                   api.CLIInfo{Path: "SelfLink", Skip: false, Insert: "", Help: ""},
			"tenant":                      api.CLIInfo{Path: "Tenant", Skip: false, Insert: "", Help: ""},
			"uuid":                        api.CLIInfo{Path: "UUID", Skip: false, Insert: "", Help: ""},
		},
	},
	"security.AppSpec": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(AppSpec{}) },
		Fields: map[string]api.Field{
			"Protocol": api.Field{Name: "Protocol", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "protocol", Pointer: true, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"ALG": api.Field{Name: "ALG", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "alg", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "security.ALG"},
		},
	},
	"security.AppStatus": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(AppStatus{}) },
		Fields: map[string]api.Field{
			"AttachedPolicies": api.Field{Name: "AttachedPolicies", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "attached-policies", Pointer: false, Slice: true, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"security.DNS": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(DNS{}) },
		Fields: map[string]api.Field{
			"DropMultiQuestionPackets": api.Field{Name: "DropMultiQuestionPackets", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "drop-multi-question-packets", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_BOOL"},
		},
	},
	"security.FTP": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(FTP{}) },
		Fields: map[string]api.Field{},
	},
	"security.MSRPC": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(MSRPC{}) },
		Fields: map[string]api.Field{},
	},
	"security.RSTP": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(RSTP{}) },
		Fields: map[string]api.Field{},
	},
	"security.SIP": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SIP{}) },
		Fields: map[string]api.Field{
			"MaxCallDuration": api.Field{Name: "MaxCallDuration", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "max-call-duration", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_UINT32"},
		},
	},
	"security.SunRPC": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(SunRPC{}) },
		Fields: map[string]api.Field{
			"ProgramID": api.Field{Name: "ProgramID", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "program-id", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},

			"MapEntryTimeout": api.Field{Name: "MapEntryTimeout", CLITag: api.CLIInfo{Path: "", Skip: false, Insert: "", Help: ""}, JSONTag: "map-entry-timeout", Pointer: true, Slice: false, Map: false, Inline: false, FromInline: false, KeyType: "", Type: "TYPE_STRING"},
		},
	},
	"security.TFTP": &api.Struct{
		Kind: "", APIGroup: "", Scopes: []string{}, GetTypeFn: func() reflect.Type { return reflect.TypeOf(TFTP{}) },
		Fields: map[string]api.Field{},
	},
}

func init() {
	schema := runtime.GetDefaultScheme()
	schema.AddSchema(typesMapApp)
}
