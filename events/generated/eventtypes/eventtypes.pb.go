// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: eventtypes.proto

/*
Package eventtypes is a generated protocol buffer package.

It is generated from these files:
	eventtypes.proto

It has these top-level messages:
*/
package eventtypes

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import gogoproto "github.com/gogo/protobuf/gogoproto"
import google_protobuf "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
import eventattrs "github.com/pensando/sw/events/generated/eventattrs"

import strconv "strconv"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// goproto_enum_prefix from public import gogo.proto
var E_GoprotoEnumPrefix = gogoproto.E_GoprotoEnumPrefix

// goproto_enum_stringer from public import gogo.proto
var E_GoprotoEnumStringer = gogoproto.E_GoprotoEnumStringer

// enum_stringer from public import gogo.proto
var E_EnumStringer = gogoproto.E_EnumStringer

// enum_customname from public import gogo.proto
var E_EnumCustomname = gogoproto.E_EnumCustomname

// enumdecl from public import gogo.proto
var E_Enumdecl = gogoproto.E_Enumdecl

// enumvalue_customname from public import gogo.proto
var E_EnumvalueCustomname = gogoproto.E_EnumvalueCustomname

// goproto_getters_all from public import gogo.proto
var E_GoprotoGettersAll = gogoproto.E_GoprotoGettersAll

// goproto_enum_prefix_all from public import gogo.proto
var E_GoprotoEnumPrefixAll = gogoproto.E_GoprotoEnumPrefixAll

// goproto_stringer_all from public import gogo.proto
var E_GoprotoStringerAll = gogoproto.E_GoprotoStringerAll

// verbose_equal_all from public import gogo.proto
var E_VerboseEqualAll = gogoproto.E_VerboseEqualAll

// face_all from public import gogo.proto
var E_FaceAll = gogoproto.E_FaceAll

// gostring_all from public import gogo.proto
var E_GostringAll = gogoproto.E_GostringAll

// populate_all from public import gogo.proto
var E_PopulateAll = gogoproto.E_PopulateAll

// stringer_all from public import gogo.proto
var E_StringerAll = gogoproto.E_StringerAll

// onlyone_all from public import gogo.proto
var E_OnlyoneAll = gogoproto.E_OnlyoneAll

// equal_all from public import gogo.proto
var E_EqualAll = gogoproto.E_EqualAll

// description_all from public import gogo.proto
var E_DescriptionAll = gogoproto.E_DescriptionAll

// testgen_all from public import gogo.proto
var E_TestgenAll = gogoproto.E_TestgenAll

// benchgen_all from public import gogo.proto
var E_BenchgenAll = gogoproto.E_BenchgenAll

// marshaler_all from public import gogo.proto
var E_MarshalerAll = gogoproto.E_MarshalerAll

// unmarshaler_all from public import gogo.proto
var E_UnmarshalerAll = gogoproto.E_UnmarshalerAll

// stable_marshaler_all from public import gogo.proto
var E_StableMarshalerAll = gogoproto.E_StableMarshalerAll

// sizer_all from public import gogo.proto
var E_SizerAll = gogoproto.E_SizerAll

// goproto_enum_stringer_all from public import gogo.proto
var E_GoprotoEnumStringerAll = gogoproto.E_GoprotoEnumStringerAll

// enum_stringer_all from public import gogo.proto
var E_EnumStringerAll = gogoproto.E_EnumStringerAll

// unsafe_marshaler_all from public import gogo.proto
var E_UnsafeMarshalerAll = gogoproto.E_UnsafeMarshalerAll

// unsafe_unmarshaler_all from public import gogo.proto
var E_UnsafeUnmarshalerAll = gogoproto.E_UnsafeUnmarshalerAll

// goproto_extensions_map_all from public import gogo.proto
var E_GoprotoExtensionsMapAll = gogoproto.E_GoprotoExtensionsMapAll

// goproto_unrecognized_all from public import gogo.proto
var E_GoprotoUnrecognizedAll = gogoproto.E_GoprotoUnrecognizedAll

// gogoproto_import from public import gogo.proto
var E_GogoprotoImport = gogoproto.E_GogoprotoImport

// protosizer_all from public import gogo.proto
var E_ProtosizerAll = gogoproto.E_ProtosizerAll

// compare_all from public import gogo.proto
var E_CompareAll = gogoproto.E_CompareAll

// typedecl_all from public import gogo.proto
var E_TypedeclAll = gogoproto.E_TypedeclAll

// enumdecl_all from public import gogo.proto
var E_EnumdeclAll = gogoproto.E_EnumdeclAll

// goproto_registration from public import gogo.proto
var E_GoprotoRegistration = gogoproto.E_GoprotoRegistration

// goproto_getters from public import gogo.proto
var E_GoprotoGetters = gogoproto.E_GoprotoGetters

// goproto_stringer from public import gogo.proto
var E_GoprotoStringer = gogoproto.E_GoprotoStringer

// verbose_equal from public import gogo.proto
var E_VerboseEqual = gogoproto.E_VerboseEqual

// face from public import gogo.proto
var E_Face = gogoproto.E_Face

// gostring from public import gogo.proto
var E_Gostring = gogoproto.E_Gostring

// populate from public import gogo.proto
var E_Populate = gogoproto.E_Populate

// stringer from public import gogo.proto
var E_Stringer = gogoproto.E_Stringer

// onlyone from public import gogo.proto
var E_Onlyone = gogoproto.E_Onlyone

// equal from public import gogo.proto
var E_Equal = gogoproto.E_Equal

// description from public import gogo.proto
var E_Description = gogoproto.E_Description

// testgen from public import gogo.proto
var E_Testgen = gogoproto.E_Testgen

// benchgen from public import gogo.proto
var E_Benchgen = gogoproto.E_Benchgen

// marshaler from public import gogo.proto
var E_Marshaler = gogoproto.E_Marshaler

// unmarshaler from public import gogo.proto
var E_Unmarshaler = gogoproto.E_Unmarshaler

// stable_marshaler from public import gogo.proto
var E_StableMarshaler = gogoproto.E_StableMarshaler

// sizer from public import gogo.proto
var E_Sizer = gogoproto.E_Sizer

// unsafe_marshaler from public import gogo.proto
var E_UnsafeMarshaler = gogoproto.E_UnsafeMarshaler

// unsafe_unmarshaler from public import gogo.proto
var E_UnsafeUnmarshaler = gogoproto.E_UnsafeUnmarshaler

// goproto_extensions_map from public import gogo.proto
var E_GoprotoExtensionsMap = gogoproto.E_GoprotoExtensionsMap

// goproto_unrecognized from public import gogo.proto
var E_GoprotoUnrecognized = gogoproto.E_GoprotoUnrecognized

// protosizer from public import gogo.proto
var E_Protosizer = gogoproto.E_Protosizer

// compare from public import gogo.proto
var E_Compare = gogoproto.E_Compare

// typedecl from public import gogo.proto
var E_Typedecl = gogoproto.E_Typedecl

// nullable from public import gogo.proto
var E_Nullable = gogoproto.E_Nullable

// embed from public import gogo.proto
var E_Embed = gogoproto.E_Embed

// customtype from public import gogo.proto
var E_Customtype = gogoproto.E_Customtype

// customname from public import gogo.proto
var E_Customname = gogoproto.E_Customname

// jsontag from public import gogo.proto
var E_Jsontag = gogoproto.E_Jsontag

// moretags from public import gogo.proto
var E_Moretags = gogoproto.E_Moretags

// casttype from public import gogo.proto
var E_Casttype = gogoproto.E_Casttype

// castkey from public import gogo.proto
var E_Castkey = gogoproto.E_Castkey

// castvalue from public import gogo.proto
var E_Castvalue = gogoproto.E_Castvalue

// stdtime from public import gogo.proto
var E_Stdtime = gogoproto.E_Stdtime

// stdduration from public import gogo.proto
var E_Stdduration = gogoproto.E_Stdduration

//
type EventType int32

const (
	// ----------------------------- Cluster events -------------------------- //
	ELECTION_STARTED EventType = 0
	//
	ELECTION_CANCELLED EventType = 1
	//
	ELECTION_NOTIFICATION_FAILED EventType = 2
	//
	ELECTION_STOPPED EventType = 3
	//
	LEADER_ELECTED EventType = 4
	//
	LEADER_LOST EventType = 5
	//
	LEADER_CHANGED EventType = 6
	// ------------------------------- Node events ----------------------------- //
	NODE_JOINED EventType = 7
	//
	NODE_DISJOINED EventType = 8
	//
	NODE_HEALTHY EventType = 9
	//
	NODE_UNREACHABLE EventType = 10
	// ------------------------------- Quorum events ----------------------------- //
	QUORUM_MEMBER_ADD EventType = 15
	//
	QUORUM_MEMBER_REMOVE EventType = 16
	//
	QUORUM_MEMBER_HEALTHY EventType = 17
	//
	QUORUM_MEMBER_UNHEALTHY EventType = 18
	//
	UNSUPPORTED_QUORUM_SIZE EventType = 19
	//
	QUORUM_UNHEALTHY EventType = 20
	// ------------------------------- Diagnostic events ----------------------------- //
	MODULE_CREATION_FAILED EventType = 24
	// -------------------------------- Host/NIC events -------------------------- //
	NIC_ADMITTED EventType = 100
	//
	NIC_REJECTED EventType = 101
	//
	NIC_UNREACHABLE EventType = 102
	//
	NIC_HEALTHY EventType = 103
	//
	NIC_UNHEALTHY EventType = 104
	//
	HOST_SMART_NIC_SPEC_CONFLICT EventType = 105
	//
	NIC_DEADMITTED EventType = 106
	//
	NIC_DECOMMISSIONED EventType = 107
	// ----------------------------- API Gateway events ---------------------- //
	AUTO_GENERATED_TLS_CERT EventType = 200
	// --------------------------- Auth/Audit events ------------------------- //
	LOGIN_FAILED EventType = 300
	//
	AUDITING_FAILED EventType = 301
	// --------------------------- HAL/Link events --------------------------- //
	LINK_UP EventType = 400
	//
	LINK_DOWN EventType = 401
	// ------------------------------ System events -------------------------- //
	SERVICE_STARTED EventType = 500
	//
	SERVICE_STOPPED EventType = 501
	//
	NAPLES_SERVICE_STOPPED EventType = 502
	//
	SERVICE_PENDING EventType = 503
	//
	SERVICE_RUNNING EventType = 504
	//
	SERVICE_UNRESPONSIVE EventType = 505
	//
	SYSTEM_COLDBOOT EventType = 506
	// ------------------------------ Rollout events -------------------------- //
	ROLLOUT_STARTED EventType = 701
	//
	ROLLOUT_SUCCESS EventType = 702
	//
	ROLLOUT_FAILED EventType = 703
	//
	ROLLOUT_SUSPENDED EventType = 704
)

var EventType_name = map[int32]string{
	0:   "ELECTION_STARTED",
	1:   "ELECTION_CANCELLED",
	2:   "ELECTION_NOTIFICATION_FAILED",
	3:   "ELECTION_STOPPED",
	4:   "LEADER_ELECTED",
	5:   "LEADER_LOST",
	6:   "LEADER_CHANGED",
	7:   "NODE_JOINED",
	8:   "NODE_DISJOINED",
	9:   "NODE_HEALTHY",
	10:  "NODE_UNREACHABLE",
	15:  "QUORUM_MEMBER_ADD",
	16:  "QUORUM_MEMBER_REMOVE",
	17:  "QUORUM_MEMBER_HEALTHY",
	18:  "QUORUM_MEMBER_UNHEALTHY",
	19:  "UNSUPPORTED_QUORUM_SIZE",
	20:  "QUORUM_UNHEALTHY",
	24:  "MODULE_CREATION_FAILED",
	100: "NIC_ADMITTED",
	101: "NIC_REJECTED",
	102: "NIC_UNREACHABLE",
	103: "NIC_HEALTHY",
	104: "NIC_UNHEALTHY",
	105: "HOST_SMART_NIC_SPEC_CONFLICT",
	106: "NIC_DEADMITTED",
	107: "NIC_DECOMMISSIONED",
	200: "AUTO_GENERATED_TLS_CERT",
	300: "LOGIN_FAILED",
	301: "AUDITING_FAILED",
	400: "LINK_UP",
	401: "LINK_DOWN",
	500: "SERVICE_STARTED",
	501: "SERVICE_STOPPED",
	502: "NAPLES_SERVICE_STOPPED",
	503: "SERVICE_PENDING",
	504: "SERVICE_RUNNING",
	505: "SERVICE_UNRESPONSIVE",
	506: "SYSTEM_COLDBOOT",
	701: "ROLLOUT_STARTED",
	702: "ROLLOUT_SUCCESS",
	703: "ROLLOUT_FAILED",
	704: "ROLLOUT_SUSPENDED",
}
var EventType_value = map[string]int32{
	"ELECTION_STARTED":             0,
	"ELECTION_CANCELLED":           1,
	"ELECTION_NOTIFICATION_FAILED": 2,
	"ELECTION_STOPPED":             3,
	"LEADER_ELECTED":               4,
	"LEADER_LOST":                  5,
	"LEADER_CHANGED":               6,
	"NODE_JOINED":                  7,
	"NODE_DISJOINED":               8,
	"NODE_HEALTHY":                 9,
	"NODE_UNREACHABLE":             10,
	"QUORUM_MEMBER_ADD":            15,
	"QUORUM_MEMBER_REMOVE":         16,
	"QUORUM_MEMBER_HEALTHY":        17,
	"QUORUM_MEMBER_UNHEALTHY":      18,
	"UNSUPPORTED_QUORUM_SIZE":      19,
	"QUORUM_UNHEALTHY":             20,
	"MODULE_CREATION_FAILED":       24,
	"NIC_ADMITTED":                 100,
	"NIC_REJECTED":                 101,
	"NIC_UNREACHABLE":              102,
	"NIC_HEALTHY":                  103,
	"NIC_UNHEALTHY":                104,
	"HOST_SMART_NIC_SPEC_CONFLICT": 105,
	"NIC_DEADMITTED":               106,
	"NIC_DECOMMISSIONED":           107,
	"AUTO_GENERATED_TLS_CERT":      200,
	"LOGIN_FAILED":                 300,
	"AUDITING_FAILED":              301,
	"LINK_UP":                      400,
	"LINK_DOWN":                    401,
	"SERVICE_STARTED":              500,
	"SERVICE_STOPPED":              501,
	"NAPLES_SERVICE_STOPPED":       502,
	"SERVICE_PENDING":              503,
	"SERVICE_RUNNING":              504,
	"SERVICE_UNRESPONSIVE":         505,
	"SYSTEM_COLDBOOT":              506,
	"ROLLOUT_STARTED":              701,
	"ROLLOUT_SUCCESS":              702,
	"ROLLOUT_FAILED":               703,
	"ROLLOUT_SUSPENDED":            704,
}

func (EventType) EnumDescriptor() ([]byte, []int) { return fileDescriptorEventtypes, []int{0} }

var E_Severity = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*eventattrs.Severity)(nil),
	Field:         10005,
	Name:          "eventtypes.severity",
	Tag:           "varint,10005,opt,name=severity,enum=eventattrs.Severity",
	Filename:      "eventtypes.proto",
}

var E_Category = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*eventattrs.Category)(nil),
	Field:         10006,
	Name:          "eventtypes.category",
	Tag:           "varint,10006,opt,name=category,enum=eventattrs.Category",
	Filename:      "eventtypes.proto",
}

var E_Desc = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         10007,
	Name:          "eventtypes.desc",
	Tag:           "bytes,10007,opt,name=desc",
	Filename:      "eventtypes.proto",
}

func init() {
	proto.RegisterEnum("eventtypes.EventType", EventType_name, EventType_value)
	proto.RegisterExtension(E_Severity)
	proto.RegisterExtension(E_Category)
	proto.RegisterExtension(E_Desc)
}
func (x EventType) String() string {
	s, ok := EventType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

func init() { proto.RegisterFile("eventtypes.proto", fileDescriptorEventtypes) }

var fileDescriptorEventtypes = []byte{
	// 1604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x96, 0x4f, 0x93, 0xdc, 0x46,
	0x15, 0xc0, 0x77, 0x76, 0xd6, 0x49, 0xdc, 0x36, 0xbb, 0xb2, 0xe2, 0x78, 0x1d, 0x65, 0x99, 0x34,
	0x04, 0x8c, 0xc9, 0x92, 0x5d, 0x48, 0xd6, 0x50, 0x4e, 0x41, 0x62, 0x59, 0x6a, 0xcf, 0x2a, 0x68,
	0xa4, 0x89, 0xa4, 0x71, 0xca, 0x70, 0x10, 0x1a, 0xe9, 0xcd, 0x8c, 0x6c, 0x8d, 0x7a, 0xa2, 0x6e,
	0xed, 0xd6, 0xf2, 0x09, 0x38, 0x51, 0x70, 0x00, 0x3e, 0x00, 0x3e, 0xf8, 0x00, 0x55, 0x3e, 0xa7,
	0x8a, 0x3f, 0xc7, 0x1c, 0x39, 0xf0, 0x01, 0x28, 0x9f, 0xb8, 0x6e, 0x15, 0xff, 0x4f, 0x54, 0xf7,
	0x48, 0x5a, 0xed, 0x7a, 0x5d, 0x95, 0xe3, 0x4c, 0xf5, 0xef, 0xd7, 0xaf, 0x5f, 0xbf, 0xf7, 0xd4,
	0x48, 0x81, 0x03, 0xc8, 0x39, 0x3f, 0x5a, 0x00, 0xdb, 0x59, 0x14, 0x94, 0x53, 0x15, 0x9d, 0xfc,
	0xa3, 0xa1, 0x29, 0x9d, 0xd2, 0xe5, 0xff, 0x1a, 0x9e, 0x52, 0x3a, 0xcd, 0x60, 0x57, 0xfe, 0x1a,
	0x97, 0x93, 0xdd, 0x04, 0x58, 0x5c, 0xa4, 0x0b, 0x4e, 0x8b, 0x6a, 0x85, 0x12, 0x71, 0x5e, 0xa4,
	0xe3, 0x92, 0xd7, 0xae, 0xb7, 0xff, 0xfa, 0x1a, 0x42, 0x44, 0xe8, 0x02, 0xa1, 0x53, 0xf7, 0x91,
	0x42, 0x6c, 0x62, 0x04, 0x96, 0xeb, 0x84, 0x7e, 0xa0, 0x7b, 0x01, 0x31, 0x95, 0x15, 0xed, 0xdd,
	0x27, 0xc7, 0x6b, 0x2b, 0x4f, 0x8f, 0xd7, 0x56, 0x3e, 0x3b, 0x5e, 0xbb, 0x61, 0x43, 0x94, 0x40,
	0x81, 0x21, 0x83, 0x98, 0xa7, 0x34, 0xc7, 0x8c, 0x47, 0x05, 0x87, 0x04, 0xa7, 0x39, 0xe6, 0x33,
	0xc0, 0x71, 0x56, 0x32, 0x0e, 0x85, 0xfa, 0x03, 0xa4, 0x36, 0x26, 0x43, 0x77, 0x0c, 0x62, 0xdb,
	0xc4, 0x54, 0x3a, 0xda, 0xd7, 0x9f, 0x1c, 0xaf, 0x75, 0x2a, 0xd7, 0xeb, 0x67, 0x5d, 0x71, 0x94,
	0xc7, 0x90, 0x65, 0x90, 0xa8, 0x0f, 0xd0, 0x56, 0x83, 0x3b, 0x6e, 0x60, 0xdd, 0xb3, 0x0c, 0x5d,
	0xfe, 0xb8, 0xa7, 0x5b, 0x42, 0xb4, 0xaa, 0x7d, 0xaf, 0x25, 0xda, 0xbe, 0x17, 0xa5, 0x19, 0x24,
	0x98, 0x53, 0xcc, 0x20, 0x4f, 0x70, 0x76, 0xc6, 0x9b, 0x53, 0x9e, 0x4e, 0xd2, 0x38, 0x12, 0x3f,
	0xd4, 0xdb, 0xa7, 0xce, 0xe8, 0x0e, 0x87, 0xc4, 0x54, 0xba, 0xda, 0x5b, 0x2d, 0xdd, 0xe6, 0xf3,
	0x67, 0xa4, 0x8b, 0x05, 0x24, 0xea, 0x07, 0x68, 0xdd, 0x26, 0xba, 0x49, 0xbc, 0x50, 0x1a, 0x88,
	0xa9, 0xac, 0x69, 0x6f, 0xb7, 0x92, 0xd3, 0x6b, 0x83, 0x90, 0xe0, 0x09, 0x2d, 0x4e, 0x25, 0x85,
	0xa0, 0x4b, 0x15, 0x6f, 0xbb, 0x7e, 0xa0, 0x5c, 0xd0, 0xf6, 0x5a, 0xf0, 0x4d, 0x87, 0x26, 0x80,
	0x33, 0xca, 0x78, 0x15, 0x3f, 0x9b, 0xa5, 0x0b, 0x9c, 0x94, 0x45, 0x9a, 0x4f, 0xa5, 0xa5, 0x0e,
	0xa7, 0x15, 0x86, 0xb1, 0xaf, 0x3b, 0x7d, 0x62, 0x2a, 0x2f, 0x9d, 0x1b, 0x46, 0x3c, 0x8b, 0xf2,
	0xe9, 0xc9, 0xd5, 0x34, 0xfc, 0x1e, 0xba, 0xe4, 0xb8, 0x26, 0x09, 0x3f, 0x72, 0x2d, 0x87, 0x98,
	0xca, 0xcb, 0xf2, 0xf0, 0x35, 0xbc, 0x29, 0xc3, 0x78, 0x48, 0xd3, 0x5c, 0x24, 0xb4, 0x15, 0xfc,
	0x87, 0x68, 0x5d, 0x52, 0xa6, 0xe5, 0x57, 0xe0, 0x2b, 0xda, 0x76, 0x2b, 0x6b, 0x6f, 0x4a, 0x30,
	0x49, 0x59, 0xc5, 0x4e, 0x0a, 0x3a, 0x3f, 0x25, 0x78, 0x07, 0x5d, 0x96, 0x82, 0x7d, 0xa2, 0xdb,
	0xc1, 0xfe, 0x03, 0xe5, 0xa2, 0xf6, 0x46, 0x6b, 0xdf, 0x0d, 0x89, 0xa7, 0x0c, 0xcf, 0x20, 0xca,
	0xf8, 0xec, 0x48, 0xbd, 0x85, 0x14, 0xb9, 0x7c, 0xe4, 0x78, 0x44, 0x37, 0xf6, 0xf5, 0xbb, 0x36,
	0x51, 0x90, 0xf6, 0xe6, 0x93, 0xe3, 0xb5, 0xd5, 0x0a, 0x79, 0xb5, 0x46, 0xca, 0xbc, 0x80, 0x28,
	0x9e, 0x45, 0xe3, 0x0c, 0xd4, 0xdb, 0xe8, 0xca, 0xc7, 0x23, 0xd7, 0x1b, 0x0d, 0xc2, 0x01, 0x19,
	0xdc, 0x25, 0x5e, 0xa8, 0x9b, 0xa6, 0xb2, 0xa1, 0x7d, 0xb5, 0xb5, 0xd5, 0x35, 0x3d, 0x49, 0x20,
	0xc1, 0x73, 0x98, 0x8f, 0xa1, 0x10, 0x45, 0xf3, 0x69, 0x49, 0x8b, 0x72, 0xae, 0xde, 0x41, 0x57,
	0x4f, 0xa3, 0x1e, 0x19, 0xb8, 0xf7, 0x89, 0xa2, 0x68, 0x37, 0x5a, 0xb4, 0xe6, 0xc1, 0x9c, 0x1e,
	0x9c, 0xf0, 0xf2, 0x9c, 0x95, 0xc1, 0x40, 0xaf, 0x9d, 0x36, 0xd4, 0x67, 0xbd, 0xa2, 0xdd, 0x6c,
	0x29, 0xb6, 0x3e, 0x96, 0xcb, 0x6b, 0x43, 0xca, 0x70, 0x4e, 0x0f, 0x9b, 0x83, 0xf7, 0xd1, 0xe6,
	0x69, 0xc9, 0xc8, 0xa9, 0x35, 0xaa, 0xbc, 0xe7, 0xfa, 0xfc, 0xbd, 0x73, 0x35, 0x65, 0x5e, 0x8b,
	0x5c, 0xb4, 0x39, 0x72, 0xfc, 0xd1, 0x70, 0xe8, 0x8a, 0x46, 0x0e, 0x2b, 0xa9, 0x6f, 0xfd, 0x88,
	0x28, 0xaf, 0xca, 0xa6, 0xae, 0x45, 0x37, 0x2a, 0x11, 0x4b, 0x7f, 0x2a, 0xf3, 0x39, 0x86, 0x8c,
	0x1e, 0x62, 0x56, 0x2e, 0x16, 0x54, 0xf6, 0xf6, 0x3c, 0xcd, 0xd3, 0x79, 0x39, 0x57, 0x07, 0x48,
	0xa9, 0x24, 0x27, 0x21, 0x5d, 0x95, 0x9d, 0x58, 0x9b, 0xb6, 0x2b, 0x53, 0x42, 0x41, 0x44, 0xc3,
	0xf1, 0x2c, 0x3a, 0x00, 0x0c, 0x39, 0x2d, 0xa7, 0xb3, 0xfa, 0x80, 0x55, 0xbc, 0x4c, 0x75, 0xd0,
	0xb5, 0x81, 0x6b, 0x8e, 0x6c, 0x12, 0x1a, 0x1e, 0x39, 0xd5, 0xde, 0xd7, 0x65, 0x78, 0x75, 0x65,
	0xdd, 0x18, 0xd0, 0xa4, 0xcc, 0x00, 0xc7, 0x05, 0xc8, 0x16, 0x96, 0x7d, 0x95, 0xa4, 0xd1, 0x34,
	0xa7, 0x8c, 0xa7, 0x31, 0xc3, 0x13, 0xd9, 0xfe, 0xea, 0x6d, 0x74, 0xd9, 0xb1, 0x8c, 0x50, 0x37,
	0x07, 0x56, 0x20, 0x9a, 0x33, 0xd1, 0xbe, 0xd1, 0x4a, 0xfa, 0x1b, 0x8e, 0x65, 0xe0, 0x28, 0x99,
	0xa7, 0x9c, 0x2f, 0x47, 0x45, 0xbb, 0x36, 0xfb, 0x4b, 0xd4, 0x23, 0x1f, 0x2d, 0xfb, 0x1a, 0xb4,
	0x5b, 0xad, 0x00, 0xbe, 0xa9, 0x63, 0x7f, 0x1e, 0x15, 0xbc, 0x36, 0x30, 0x26, 0xa2, 0x28, 0xe0,
	0xd3, 0x12, 0x18, 0xc7, 0x87, 0x11, 0xc3, 0x05, 0x3c, 0x94, 0x2d, 0xaf, 0xbe, 0x87, 0x36, 0x84,
	0xa8, 0x5d, 0xb4, 0x13, 0xad, 0xd7, 0xca, 0x90, 0x2a, 0x24, 0x67, 0x6a, 0x76, 0x1b, 0x5d, 0x12,
	0x50, 0x9d, 0xd2, 0xa9, 0xa6, 0xb5, 0xe2, 0x5e, 0xaf, 0x80, 0xfa, 0x56, 0x77, 0xd1, 0x97, 0x96,
	0x3b, 0xd4, 0xcb, 0x67, 0xda, 0x56, 0xcb, 0xaf, 0x34, 0xfe, 0x1a, 0xf8, 0x79, 0x07, 0x6d, 0xed,
	0xbb, 0x7e, 0x10, 0xfa, 0x03, 0xdd, 0x0b, 0x42, 0x01, 0xfb, 0x43, 0x62, 0x84, 0x86, 0xeb, 0xdc,
	0xb3, 0x2d, 0x23, 0x50, 0x52, 0x2d, 0x6b, 0x1d, 0xf6, 0x27, 0xc1, 0x0c, 0x30, 0x3b, 0x62, 0x1c,
	0xe6, 0x78, 0x16, 0x31, 0x9c, 0x00, 0x5f, 0x4e, 0xb3, 0x08, 0xc7, 0x34, 0x9f, 0x64, 0x69, 0xcc,
	0xf1, 0x18, 0xf8, 0x21, 0xc0, 0x72, 0xac, 0x34, 0x99, 0x61, 0x0b, 0x88, 0x9b, 0x31, 0xcb, 0x30,
	0x9d, 0xe0, 0x24, 0x9d, 0x4c, 0xa0, 0x80, 0x9c, 0xe3, 0x7d, 0x31, 0xd6, 0xe8, 0x58, 0x24, 0x89,
	0xa9, 0x77, 0x90, 0x38, 0x53, 0x68, 0x92, 0xe6, 0xa6, 0x1e, 0x6a, 0xdf, 0x6a, 0x9d, 0x18, 0x0b,
	0x5b, 0x02, 0xef, 0x34, 0x97, 0xf5, 0xdc, 0x28, 0xe9, 0x23, 0x75, 0x69, 0x30, 0xdc, 0xc1, 0xc0,
	0xf2, 0x7d, 0xcb, 0x15, 0xf3, 0xe8, 0x91, 0xb6, 0xdb, 0xb2, 0xbc, 0xb5, 0xb4, 0xc4, 0x74, 0x5e,
	0x5d, 0xd8, 0x79, 0xa2, 0x09, 0xda, 0xd4, 0x47, 0x81, 0x1b, 0xf6, 0x89, 0x43, 0x3c, 0x5d, 0x74,
	0x49, 0x60, 0xfb, 0xa1, 0x41, 0xbc, 0x40, 0xf9, 0xbc, 0xa3, 0xed, 0xb7, 0xd2, 0xf2, 0x7d, 0xbd,
	0xe4, 0x14, 0x4f, 0x21, 0x87, 0x22, 0x12, 0x31, 0xc5, 0x50, 0x54, 0x1f, 0x95, 0xaa, 0x65, 0xc4,
	0x98, 0x2e, 0x59, 0x35, 0xf5, 0xf5, 0xa1, 0x85, 0xfb, 0x11, 0x87, 0xc3, 0xe8, 0x08, 0x07, 0xb6,
	0xaf, 0x7e, 0x1b, 0x5d, 0xb6, 0xdd, 0xbe, 0xd5, 0x14, 0xf8, 0xef, 0x56, 0xb5, 0x2f, 0xb7, 0x62,
	0xbd, 0x32, 0x62, 0x50, 0xe0, 0x8c, 0x4e, 0xd3, 0xbc, 0x2e, 0xe6, 0x0f, 0xd0, 0x86, 0x3e, 0x32,
	0xad, 0xc0, 0x72, 0xfa, 0x35, 0xf4, 0xfb, 0x55, 0x39, 0x45, 0xea, 0x9b, 0xde, 0xfa, 0xa4, 0x48,
	0xb9, 0xd8, 0x95, 0x4e, 0xb0, 0x5e, 0x26, 0x29, 0x97, 0x1f, 0xf2, 0x9a, 0xdf, 0x46, 0x2f, 0xdb,
	0x96, 0xf3, 0xc3, 0x70, 0x34, 0x54, 0x7e, 0xd1, 0x6d, 0x36, 0xeb, 0x88, 0xcd, 0x86, 0xb4, 0xe0,
	0x22, 0xe6, 0x2c, 0xcd, 0x1f, 0x41, 0x82, 0xcb, 0x85, 0xba, 0x87, 0x2e, 0xca, 0xc5, 0xa6, 0xfb,
	0x89, 0xa3, 0xfc, 0xb2, 0xab, 0x7d, 0xad, 0x3a, 0xb8, 0x58, 0x7e, 0x5d, 0x2e, 0x17, 0x6b, 0xc5,
	0xb7, 0x9e, 0x97, 0x4c, 0xa0, 0x09, 0x3d, 0xcc, 0xd5, 0xef, 0xa0, 0x0d, 0x9f, 0x78, 0xf7, 0x2d,
	0x83, 0x34, 0x8f, 0x85, 0x7f, 0x74, 0xe5, 0x50, 0xef, 0x3e, 0x3d, 0x5e, 0x5b, 0x15, 0x43, 0xdd,
	0x87, 0xe2, 0x20, 0x8d, 0xa1, 0x7e, 0x25, 0x9c, 0x46, 0x96, 0xdf, 0xde, 0x7f, 0x2e, 0x91, 0xce,
	0xf3, 0xc8, 0xf2, 0xa3, 0xfb, 0x21, 0xba, 0xe6, 0xe8, 0x43, 0x9b, 0xf8, 0xe1, 0x59, 0xf2, 0x5f,
	0x5d, 0x39, 0xd6, 0x57, 0x2b, 0xf2, 0x9a, 0x13, 0x2d, 0x32, 0x60, 0x98, 0x9d, 0x11, 0xb4, 0xf6,
	0x1c, 0x12, 0xc7, 0xb4, 0x9c, 0xbe, 0xf2, 0xef, 0x17, 0x84, 0xb9, 0x80, 0x3c, 0x49, 0xf3, 0x69,
	0x1b, 0xf1, 0x46, 0x8e, 0x23, 0x90, 0xff, 0xbc, 0x00, 0x29, 0xca, 0x3c, 0x17, 0xc8, 0x8f, 0xd1,
	0xd5, 0x1a, 0x11, 0xcd, 0xef, 0x0f, 0x5d, 0xc7, 0xb7, 0xee, 0x13, 0xe5, 0xbf, 0x5d, 0xed, 0x4e,
	0x2b, 0xc8, 0xbd, 0x9a, 0x13, 0xfd, 0xcf, 0x16, 0x34, 0x67, 0xe9, 0x01, 0xe0, 0xa4, 0x04, 0x31,
	0x91, 0xb2, 0x28, 0x7e, 0x24, 0x6e, 0xb3, 0xea, 0xbe, 0x02, 0x18, 0x2d, 0x8b, 0x18, 0x98, 0xba,
	0x87, 0x36, 0xfc, 0x07, 0x7e, 0x40, 0x06, 0xa1, 0xe1, 0xda, 0xe6, 0x5d, 0xd7, 0x0d, 0x94, 0xff,
	0x75, 0xe5, 0x58, 0xa9, 0xd3, 0xa6, 0xfa, 0x4b, 0x26, 0xa6, 0x59, 0x82, 0xc7, 0x94, 0x56, 0xb3,
	0xc8, 0x73, 0x6d, 0xdb, 0x1d, 0x05, 0xcd, 0xfd, 0xfc, 0xe1, 0x42, 0x53, 0x0a, 0x5d, 0x51, 0x0a,
	0x1e, 0xcd, 0x32, 0x5a, 0x72, 0x9c, 0xe6, 0x29, 0x4f, 0x45, 0x5d, 0xab, 0x77, 0x5a, 0xd0, 0xc8,
	0x30, 0x88, 0xef, 0x2b, 0x7f, 0xbc, 0xd0, 0x3c, 0x2f, 0x04, 0xd4, 0xab, 0x21, 0x56, 0xc6, 0x31,
	0x30, 0x36, 0x29, 0xb3, 0xec, 0x08, 0xc7, 0x74, 0xbe, 0xc8, 0x40, 0x18, 0x76, 0xd1, 0x7a, 0x6d,
	0xa8, 0x0a, 0xf7, 0x4f, 0x17, 0xe4, 0x44, 0x5b, 0xad, 0x04, 0xeb, 0xb5, 0xa0, 0x2a, 0xd5, 0x5b,
	0xe8, 0xca, 0xc9, 0x96, 0xbe, 0xb8, 0x23, 0x62, 0x2a, 0x7f, 0x7e, 0x41, 0xa4, 0xac, 0x64, 0xe2,
	0x96, 0x20, 0xd1, 0x5e, 0xff, 0xd9, 0x6f, 0x7b, 0x2b, 0x4f, 0x1e, 0xf7, 0x56, 0x9e, 0x3e, 0xee,
	0x75, 0x3e, 0x7b, 0xdc, 0xbb, 0xd8, 0x3c, 0x64, 0xdf, 0x0f, 0xd0, 0x2b, 0x0c, 0x0e, 0xa0, 0x48,
	0xf9, 0x91, 0xfa, 0x95, 0x9d, 0xe5, 0xbb, 0x78, 0xa7, 0x7e, 0x17, 0xef, 0x90, 0xbc, 0x9c, 0xdf,
	0x8f, 0xb2, 0x12, 0xdc, 0x85, 0x1c, 0x55, 0xd7, 0x7f, 0xe5, 0xe0, 0xce, 0xcd, 0xf5, 0x77, 0xaf,
	0xee, 0xc8, 0x97, 0xb5, 0x78, 0x24, 0xb3, 0x1d, 0xbf, 0xe2, 0xbd, 0xc6, 0x24, 0xac, 0xa2, 0xdd,
	0xa7, 0xb4, 0xf8, 0x42, 0xd6, 0x5f, 0x9f, 0x63, 0x35, 0x2a, 0xde, 0x6b, 0x4c, 0xef, 0x7f, 0x17,
	0xad, 0x89, 0x87, 0xfa, 0x17, 0x31, 0xfe, 0x46, 0x18, 0x2f, 0x7a, 0x72, 0xfd, 0x5d, 0xe5, 0xf3,
	0x67, 0xbd, 0xce, 0x5f, 0x9e, 0xf5, 0x3a, 0x7f, 0x7b, 0xd6, 0xeb, 0xfc, 0xfd, 0x59, 0x6f, 0x65,
	0xb8, 0x32, 0x7e, 0x49, 0xb2, 0xef, 0xfd, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x79, 0xdd, 0xb9,
	0x35, 0x0c, 0x00, 0x00,
}
