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
	//
	CLOCK_SYNC_FAILED EventType = 11
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
	// ------------------------ Config Snapshot/Restore events ------------------------ //
	CONFIG_RESTORED EventType = 30
	//
	CONFIG_RESTORE_ABORTED EventType = 31
	//
	CONFIG_RESTORE_FAILED EventType = 32
	// -------------------------------- Host/DSC events -------------------------- //
	DSC_ADMITTED EventType = 100
	//
	DSC_REJECTED EventType = 101
	//
	DSC_UNREACHABLE EventType = 102
	//
	DSC_HEALTHY EventType = 103
	//
	DSC_UNHEALTHY EventType = 104
	//
	HOST_DSC_SPEC_CONFLICT EventType = 105
	//
	DSC_DEADMITTED EventType = 106
	//
	DSC_DECOMMISSIONED EventType = 107
	// ----------------------------- API Gateway events ---------------------- //
	AUTO_GENERATED_TLS_CERT EventType = 200
	// --------------------------- Auth/Audit events ------------------------- //
	LOGIN_FAILED EventType = 300
	//
	AUDITING_FAILED EventType = 301
	//
	PASSWORD_CHANGED EventType = 302
	//
	PASSWORD_RESET EventType = 303
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
	//
	SYSTEM_RESOURCE_USAGE EventType = 507
	//
	NAPLES_FATAL_INTERRUPT EventType = 508
	//
	NAPLES_CATTRIP_INTERRUPT EventType = 509
	//
	NAPLES_OVER_TEMP EventType = 510
	//
	NAPLES_OVER_TEMP_EXIT EventType = 511
	//
	NAPLES_PANIC_EVENT EventType = 512
	//
	NAPLES_POST_DIAG_FAILURE_EVENT EventType = 513
	//
	NAPLES_INFO_PCIEHEALTH_EVENT EventType = 514
	//
	NAPLES_WARN_PCIEHEALTH_EVENT EventType = 515
	//
	NAPLES_ERR_PCIEHEALTH_EVENT EventType = 516
	// ------------------------------ Resource events -------------------------- //
	DISK_THRESHOLD_EXCEEDED EventType = 601
	// ------------------------------ Rollout events -------------------------- //
	ROLLOUT_STARTED EventType = 701
	//
	ROLLOUT_SUCCESS EventType = 702
	//
	ROLLOUT_FAILED EventType = 703
	//
	ROLLOUT_SUSPENDED EventType = 704
	// ------------------------------ Config events -------------------------- //
	CONFIG_FAIL EventType = 801
	// ------------------------------ Session Limit events -------------------------- //
	TCP_HALF_OPEN_SESSION_LIMIT_APPROACH EventType = 901
	//
	TCP_HALF_OPEN_SESSION_LIMIT_REACHED EventType = 902
	//
	UDP_ACTIVE_SESSION_LIMIT_APPROACH EventType = 903
	//
	UDP_ACTIVE_SESSION_LIMIT_REACHED EventType = 904
	//
	ICMP_ACTIVE_SESSION_LIMIT_APPROACH EventType = 905
	//
	ICMP_ACTIVE_SESSION_LIMIT_REACHED EventType = 906
	//
	OTHER_ACTIVE_SESSION_LIMIT_APPROACH EventType = 907
	//
	OTHER_ACTIVE_SESSION_LIMIT_REACHED EventType = 908
	// ------------------------- Orchestration events ------------------------- //
	ORCH_CONNECTION_ERROR EventType = 1001
	//
	ORCH_LOGIN_FAILURE EventType = 1002
	//
	ORCH_CONFIG_PUSH_FAILURE EventType = 1003
	//
	ORCH_INVALID_ACTION EventType = 1004
	//
	MIGRATION_FAILED EventType = 1005
	//
	MIGRATION_TIMED_OUT EventType = 1006
	//
	ORCH_ALREADY_MANAGED EventType = 1007
	//
	ORCH_UNSUPPORTED_VERSION EventType = 1008
	//
	ORCH_DSC_NOT_ADMITTED EventType = 1009
	//
	ORCH_DSC_MODE_INCOMPATIBLE EventType = 1010
	// ------------------------- Controller events ------------------------- //
	COLLECTOR_REACHABLE EventType = 1101
	//
	COLLECTOR_UNREACHABLE EventType = 1102
	// ------------------------- Control Plane events ------------------------- //
	BGP_SESSION_ESTABLISHED EventType = 1201
	//
	BGP_SESSION_DOWN EventType = 1202
	// ----------------------------- Flowlog events -------------------------- //
	FLOWLOGS_DROPPED EventType = 1301
	//
	FLOWLOGS_REPORTING_ERROR EventType = 1302
	//
	FLOWLOGS_RATE_LIMITED EventType = 1303
)

var EventType_name = map[int32]string{
	0:    "ELECTION_STARTED",
	1:    "ELECTION_CANCELLED",
	2:    "ELECTION_NOTIFICATION_FAILED",
	3:    "ELECTION_STOPPED",
	4:    "LEADER_ELECTED",
	5:    "LEADER_LOST",
	6:    "LEADER_CHANGED",
	7:    "NODE_JOINED",
	8:    "NODE_DISJOINED",
	9:    "NODE_HEALTHY",
	10:   "NODE_UNREACHABLE",
	11:   "CLOCK_SYNC_FAILED",
	15:   "QUORUM_MEMBER_ADD",
	16:   "QUORUM_MEMBER_REMOVE",
	17:   "QUORUM_MEMBER_HEALTHY",
	18:   "QUORUM_MEMBER_UNHEALTHY",
	19:   "UNSUPPORTED_QUORUM_SIZE",
	20:   "QUORUM_UNHEALTHY",
	24:   "MODULE_CREATION_FAILED",
	30:   "CONFIG_RESTORED",
	31:   "CONFIG_RESTORE_ABORTED",
	32:   "CONFIG_RESTORE_FAILED",
	100:  "DSC_ADMITTED",
	101:  "DSC_REJECTED",
	102:  "DSC_UNREACHABLE",
	103:  "DSC_HEALTHY",
	104:  "DSC_UNHEALTHY",
	105:  "HOST_DSC_SPEC_CONFLICT",
	106:  "DSC_DEADMITTED",
	107:  "DSC_DECOMMISSIONED",
	200:  "AUTO_GENERATED_TLS_CERT",
	300:  "LOGIN_FAILED",
	301:  "AUDITING_FAILED",
	302:  "PASSWORD_CHANGED",
	303:  "PASSWORD_RESET",
	400:  "LINK_UP",
	401:  "LINK_DOWN",
	500:  "SERVICE_STARTED",
	501:  "SERVICE_STOPPED",
	502:  "NAPLES_SERVICE_STOPPED",
	503:  "SERVICE_PENDING",
	504:  "SERVICE_RUNNING",
	505:  "SERVICE_UNRESPONSIVE",
	506:  "SYSTEM_COLDBOOT",
	507:  "SYSTEM_RESOURCE_USAGE",
	508:  "NAPLES_FATAL_INTERRUPT",
	509:  "NAPLES_CATTRIP_INTERRUPT",
	510:  "NAPLES_OVER_TEMP",
	511:  "NAPLES_OVER_TEMP_EXIT",
	512:  "NAPLES_PANIC_EVENT",
	513:  "NAPLES_POST_DIAG_FAILURE_EVENT",
	514:  "NAPLES_INFO_PCIEHEALTH_EVENT",
	515:  "NAPLES_WARN_PCIEHEALTH_EVENT",
	516:  "NAPLES_ERR_PCIEHEALTH_EVENT",
	601:  "DISK_THRESHOLD_EXCEEDED",
	701:  "ROLLOUT_STARTED",
	702:  "ROLLOUT_SUCCESS",
	703:  "ROLLOUT_FAILED",
	704:  "ROLLOUT_SUSPENDED",
	801:  "CONFIG_FAIL",
	901:  "TCP_HALF_OPEN_SESSION_LIMIT_APPROACH",
	902:  "TCP_HALF_OPEN_SESSION_LIMIT_REACHED",
	903:  "UDP_ACTIVE_SESSION_LIMIT_APPROACH",
	904:  "UDP_ACTIVE_SESSION_LIMIT_REACHED",
	905:  "ICMP_ACTIVE_SESSION_LIMIT_APPROACH",
	906:  "ICMP_ACTIVE_SESSION_LIMIT_REACHED",
	907:  "OTHER_ACTIVE_SESSION_LIMIT_APPROACH",
	908:  "OTHER_ACTIVE_SESSION_LIMIT_REACHED",
	1001: "ORCH_CONNECTION_ERROR",
	1002: "ORCH_LOGIN_FAILURE",
	1003: "ORCH_CONFIG_PUSH_FAILURE",
	1004: "ORCH_INVALID_ACTION",
	1005: "MIGRATION_FAILED",
	1006: "MIGRATION_TIMED_OUT",
	1007: "ORCH_ALREADY_MANAGED",
	1008: "ORCH_UNSUPPORTED_VERSION",
	1009: "ORCH_DSC_NOT_ADMITTED",
	1010: "ORCH_DSC_MODE_INCOMPATIBLE",
	1101: "COLLECTOR_REACHABLE",
	1102: "COLLECTOR_UNREACHABLE",
	1201: "BGP_SESSION_ESTABLISHED",
	1202: "BGP_SESSION_DOWN",
	1301: "FLOWLOGS_DROPPED",
	1302: "FLOWLOGS_REPORTING_ERROR",
	1303: "FLOWLOGS_RATE_LIMITED",
}
var EventType_value = map[string]int32{
	"ELECTION_STARTED":                     0,
	"ELECTION_CANCELLED":                   1,
	"ELECTION_NOTIFICATION_FAILED":         2,
	"ELECTION_STOPPED":                     3,
	"LEADER_ELECTED":                       4,
	"LEADER_LOST":                          5,
	"LEADER_CHANGED":                       6,
	"NODE_JOINED":                          7,
	"NODE_DISJOINED":                       8,
	"NODE_HEALTHY":                         9,
	"NODE_UNREACHABLE":                     10,
	"CLOCK_SYNC_FAILED":                    11,
	"QUORUM_MEMBER_ADD":                    15,
	"QUORUM_MEMBER_REMOVE":                 16,
	"QUORUM_MEMBER_HEALTHY":                17,
	"QUORUM_MEMBER_UNHEALTHY":              18,
	"UNSUPPORTED_QUORUM_SIZE":              19,
	"QUORUM_UNHEALTHY":                     20,
	"MODULE_CREATION_FAILED":               24,
	"CONFIG_RESTORED":                      30,
	"CONFIG_RESTORE_ABORTED":               31,
	"CONFIG_RESTORE_FAILED":                32,
	"DSC_ADMITTED":                         100,
	"DSC_REJECTED":                         101,
	"DSC_UNREACHABLE":                      102,
	"DSC_HEALTHY":                          103,
	"DSC_UNHEALTHY":                        104,
	"HOST_DSC_SPEC_CONFLICT":               105,
	"DSC_DEADMITTED":                       106,
	"DSC_DECOMMISSIONED":                   107,
	"AUTO_GENERATED_TLS_CERT":              200,
	"LOGIN_FAILED":                         300,
	"AUDITING_FAILED":                      301,
	"PASSWORD_CHANGED":                     302,
	"PASSWORD_RESET":                       303,
	"LINK_UP":                              400,
	"LINK_DOWN":                            401,
	"SERVICE_STARTED":                      500,
	"SERVICE_STOPPED":                      501,
	"NAPLES_SERVICE_STOPPED":               502,
	"SERVICE_PENDING":                      503,
	"SERVICE_RUNNING":                      504,
	"SERVICE_UNRESPONSIVE":                 505,
	"SYSTEM_COLDBOOT":                      506,
	"SYSTEM_RESOURCE_USAGE":                507,
	"NAPLES_FATAL_INTERRUPT":               508,
	"NAPLES_CATTRIP_INTERRUPT":             509,
	"NAPLES_OVER_TEMP":                     510,
	"NAPLES_OVER_TEMP_EXIT":                511,
	"NAPLES_PANIC_EVENT":                   512,
	"NAPLES_POST_DIAG_FAILURE_EVENT":       513,
	"NAPLES_INFO_PCIEHEALTH_EVENT":         514,
	"NAPLES_WARN_PCIEHEALTH_EVENT":         515,
	"NAPLES_ERR_PCIEHEALTH_EVENT":          516,
	"DISK_THRESHOLD_EXCEEDED":              601,
	"ROLLOUT_STARTED":                      701,
	"ROLLOUT_SUCCESS":                      702,
	"ROLLOUT_FAILED":                       703,
	"ROLLOUT_SUSPENDED":                    704,
	"CONFIG_FAIL":                          801,
	"TCP_HALF_OPEN_SESSION_LIMIT_APPROACH": 901,
	"TCP_HALF_OPEN_SESSION_LIMIT_REACHED":  902,
	"UDP_ACTIVE_SESSION_LIMIT_APPROACH":    903,
	"UDP_ACTIVE_SESSION_LIMIT_REACHED":     904,
	"ICMP_ACTIVE_SESSION_LIMIT_APPROACH":   905,
	"ICMP_ACTIVE_SESSION_LIMIT_REACHED":    906,
	"OTHER_ACTIVE_SESSION_LIMIT_APPROACH":  907,
	"OTHER_ACTIVE_SESSION_LIMIT_REACHED":   908,
	"ORCH_CONNECTION_ERROR":                1001,
	"ORCH_LOGIN_FAILURE":                   1002,
	"ORCH_CONFIG_PUSH_FAILURE":             1003,
	"ORCH_INVALID_ACTION":                  1004,
	"MIGRATION_FAILED":                     1005,
	"MIGRATION_TIMED_OUT":                  1006,
	"ORCH_ALREADY_MANAGED":                 1007,
	"ORCH_UNSUPPORTED_VERSION":             1008,
	"ORCH_DSC_NOT_ADMITTED":                1009,
	"ORCH_DSC_MODE_INCOMPATIBLE":           1010,
	"COLLECTOR_REACHABLE":                  1101,
	"COLLECTOR_UNREACHABLE":                1102,
	"BGP_SESSION_ESTABLISHED":              1201,
	"BGP_SESSION_DOWN":                     1202,
	"FLOWLOGS_DROPPED":                     1301,
	"FLOWLOGS_REPORTING_ERROR":             1302,
	"FLOWLOGS_RATE_LIMITED":                1303,
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

var E_SuppressMm = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.EnumValueOptions)(nil),
	ExtensionType: (*bool)(nil),
	Field:         10008,
	Name:          "eventtypes.suppress_mm",
	Tag:           "varint,10008,opt,name=suppress_mm,json=suppressMm",
	Filename:      "eventtypes.proto",
}

func init() {
	proto.RegisterEnum("eventtypes.EventType", EventType_name, EventType_value)
	proto.RegisterExtension(E_Severity)
	proto.RegisterExtension(E_Category)
	proto.RegisterExtension(E_Desc)
	proto.RegisterExtension(E_SuppressMm)
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
	// 3012 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x99, 0x4b, 0x74, 0xd5, 0xc6,
	0x7f, 0xc7, 0xfd, 0xb8, 0x60, 0x18, 0x52, 0x2c, 0x04, 0x06, 0x22, 0x88, 0x33, 0x40, 0x4b, 0xd3,
	0x90, 0x18, 0x02, 0xcd, 0x8b, 0x26, 0x0d, 0xb2, 0x34, 0xb6, 0x15, 0x74, 0x25, 0x21, 0xe9, 0x9a,
	0x10, 0x4e, 0x8e, 0x8e, 0x2c, 0x8d, 0xef, 0x15, 0xe8, 0x6a, 0x6e, 0x34, 0x92, 0x5d, 0x77, 0xd5,
	0x36, 0x7d, 0xa4, 0x8f, 0x45, 0xbb, 0x48, 0xd2, 0x6d, 0x4e, 0xb3, 0x60, 0xd1, 0x47, 0xda, 0x65,
	0xce, 0x69, 0x9b, 0x65, 0x36, 0xed, 0xe9, 0xb6, 0xbb, 0x1e, 0x56, 0x7d, 0xb7, 0xee, 0xe9, 0xfb,
	0x7d, 0x66, 0x34, 0xd2, 0xd5, 0xb5, 0x0d, 0xe1, 0xbf, 0xc3, 0x1c, 0x7f, 0x3f, 0xf3, 0x9d, 0xdf,
	0xcc, 0xfc, 0x1e, 0x32, 0x90, 0xf0, 0x16, 0xce, 0x8a, 0x62, 0x67, 0x84, 0xe9, 0xd2, 0x28, 0x27,
	0x05, 0x91, 0xc1, 0xf8, 0x7f, 0x14, 0xd0, 0x27, 0x7d, 0x52, 0xfd, 0xbf, 0x02, 0xfb, 0x84, 0xf4,
	0x53, 0x7c, 0x95, 0xff, 0xb4, 0x51, 0x6e, 0x5e, 0x8d, 0x31, 0x8d, 0xf2, 0x64, 0x54, 0x90, 0x5c,
	0xfc, 0x86, 0x14, 0x16, 0x45, 0x9e, 0x6c, 0x94, 0x45, 0xcd, 0x7a, 0xf9, 0xcb, 0x6b, 0x00, 0x20,
	0x86, 0xf3, 0x19, 0x4e, 0x36, 0x81, 0x84, 0x4c, 0xa4, 0xf9, 0x86, 0x6d, 0x05, 0x9e, 0xaf, 0xba,
	0x3e, 0xd2, 0xa5, 0x29, 0xe5, 0x8d, 0x47, 0xbb, 0x9d, 0xa9, 0xaf, 0x77, 0x3b, 0x53, 0xdf, 0xec,
	0x76, 0x2e, 0x9b, 0x38, 0x8c, 0x71, 0x0e, 0x71, 0x8a, 0xa3, 0x22, 0x21, 0x19, 0xa4, 0x45, 0x98,
	0x17, 0x38, 0x86, 0x49, 0x06, 0x8b, 0x01, 0x86, 0x51, 0x5a, 0xd2, 0x02, 0xe7, 0xdf, 0xee, 0x76,
	0xa6, 0x65, 0x15, 0xc8, 0x0d, 0x4d, 0x53, 0x2d, 0x0d, 0x99, 0x26, 0xd2, 0xa5, 0x69, 0xe5, 0xc7,
	0x1e, 0xed, 0x76, 0xa6, 0x05, 0xef, 0xf9, 0xbd, 0xbc, 0x28, 0xcc, 0x22, 0x9c, 0xa6, 0x38, 0xe6,
	0x88, 0x8f, 0xc0, 0xf9, 0x06, 0x61, 0xd9, 0xbe, 0xb1, 0x62, 0x68, 0x2a, 0xff, 0x61, 0x45, 0x35,
	0x18, 0x6c, 0x46, 0xf9, 0x89, 0x16, 0xec, 0xca, 0x4a, 0x98, 0xa4, 0x38, 0x86, 0x05, 0x81, 0x14,
	0x67, 0x31, 0x4c, 0xf7, 0xb0, 0x33, 0x52, 0x24, 0x9b, 0x49, 0x14, 0xb2, 0x1f, 0x38, 0xfe, 0xdd,
	0x89, 0xfd, 0xda, 0x8e, 0x83, 0x74, 0x69, 0x56, 0xf9, 0xd1, 0x16, 0xf2, 0xcc, 0xfe, 0xfd, 0x92,
	0xd1, 0x48, 0xb8, 0x5b, 0x06, 0xc7, 0x4d, 0xa4, 0xea, 0xc8, 0x0d, 0x38, 0x05, 0xe9, 0x52, 0x47,
	0x59, 0x6a, 0x05, 0x6b, 0xb1, 0x2d, 0xc6, 0x31, 0xdc, 0x24, 0xf9, 0xbe, 0x20, 0x19, 0xe0, 0x98,
	0x60, 0x98, 0xb6, 0xe7, 0x4b, 0x87, 0x94, 0xb7, 0x5a, 0x80, 0x97, 0x2c, 0x12, 0x63, 0x98, 0x12,
	0x5a, 0x88, 0xbd, 0xd0, 0x41, 0x32, 0x82, 0x71, 0x99, 0x27, 0x59, 0x9f, 0x93, 0x6a, 0x5b, 0x7b,
	0xec, 0x68, 0x6b, 0xaa, 0xb5, 0x8a, 0x74, 0xe9, 0xf0, 0x81, 0x76, 0xa2, 0x41, 0x98, 0xf5, 0xc7,
	0x47, 0x36, 0xc1, 0x78, 0x0b, 0x1c, 0xb3, 0x6c, 0x1d, 0x05, 0xef, 0xdb, 0x86, 0x85, 0x74, 0x69,
	0x8e, 0x07, 0xa3, 0x06, 0x9c, 0xe1, 0x76, 0x1e, 0x90, 0x24, 0x63, 0x41, 0xde, 0xb3, 0x91, 0xf7,
	0xc0, 0x71, 0xae, 0xd4, 0x0d, 0x4f, 0x88, 0x8f, 0x28, 0x57, 0x5a, 0x91, 0x7c, 0x91, 0x8b, 0xe3,
	0x84, 0x0a, 0xfd, 0x66, 0x4e, 0x86, 0x6d, 0x88, 0xfc, 0x1a, 0x78, 0x8e, 0x03, 0xd6, 0x90, 0x6a,
	0xfa, 0x6b, 0xf7, 0xa4, 0xa3, 0xca, 0x8b, 0xad, 0xb5, 0xe7, 0xb9, 0x3c, 0xa1, 0x70, 0x80, 0xc3,
	0xb4, 0x18, 0xec, 0xf0, 0x35, 0xdf, 0x06, 0x12, 0x97, 0xf4, 0x2c, 0x17, 0xa9, 0xda, 0x9a, 0xba,
	0x6c, 0x22, 0x09, 0x28, 0x97, 0x1e, 0xed, 0x76, 0x66, 0x84, 0xec, 0x64, 0x2d, 0x2b, 0xb3, 0x1c,
	0x87, 0xd1, 0x20, 0xdc, 0x48, 0x31, 0x97, 0xae, 0x82, 0x13, 0x9a, 0x69, 0x6b, 0xb7, 0x03, 0xef,
	0x9e, 0xa5, 0xd5, 0xd7, 0xe9, 0x98, 0x72, 0xad, 0xe5, 0x18, 0x72, 0xed, 0xe6, 0xf8, 0x4e, 0xed,
	0x64, 0xd1, 0x20, 0x27, 0x59, 0xf2, 0xd3, 0xcc, 0x34, 0x89, 0x1e, 0x8a, 0x3b, 0x74, 0xe2, 0x4e,
	0xcf, 0x76, 0x7b, 0xdd, 0xa0, 0x8b, 0xba, 0xcb, 0xc8, 0x0d, 0x54, 0x5d, 0x97, 0xe6, 0x95, 0xcb,
	0x2d, 0xef, 0xa7, 0xd5, 0x38, 0xc6, 0x31, 0x1c, 0xe2, 0xe1, 0x06, 0xce, 0x19, 0xe9, 0xe3, 0x92,
	0xe4, 0xe5, 0x90, 0xcb, 0x75, 0x70, 0x6a, 0x52, 0xee, 0xa2, 0xae, 0xbd, 0x8e, 0x24, 0x49, 0x79,
	0xb9, 0x45, 0x50, 0x5c, 0x3c, 0x24, 0x5b, 0x63, 0x06, 0x0f, 0x5e, 0x8b, 0xb2, 0x0a, 0x16, 0x26,
	0x29, 0x75, 0x10, 0x4f, 0x28, 0xaf, 0xb4, 0x30, 0xe7, 0xef, 0x70, 0x49, 0x4d, 0x49, 0x28, 0xcc,
	0xc8, 0xf6, 0x44, 0x44, 0x6f, 0x83, 0x33, 0x93, 0xa0, 0x9e, 0x55, 0xa3, 0x64, 0x7e, 0x99, 0xea,
	0xc0, 0x2e, 0x1e, 0x88, 0x2a, 0xb3, 0x36, 0xcc, 0x03, 0x67, 0x7a, 0x96, 0xd7, 0x73, 0x1c, 0x9b,
	0x65, 0x92, 0x40, 0x80, 0x3d, 0xe3, 0x43, 0x24, 0x9d, 0xe4, 0x59, 0xa5, 0x86, 0x5d, 0x16, 0x30,
	0xca, 0x42, 0x9b, 0x50, 0xb8, 0x81, 0x53, 0xb2, 0x0d, 0x69, 0x39, 0x1a, 0x11, 0x9e, 0x5c, 0x86,
	0x49, 0x96, 0x0c, 0xc5, 0x56, 0xef, 0x00, 0x49, 0x80, 0xc6, 0xd6, 0x4e, 0xf1, 0x34, 0x50, 0xd3,
	0xae, 0x08, 0x5a, 0x4c, 0x30, 0x73, 0x55, 0xc0, 0x41, 0xb8, 0x85, 0x21, 0xce, 0x48, 0xd9, 0x1f,
	0xd4, 0x9b, 0x15, 0xbe, 0x29, 0x47, 0xba, 0xe0, 0x74, 0xd7, 0xd6, 0x7b, 0x26, 0x0a, 0x34, 0x17,
	0x4d, 0xe4, 0x97, 0xb3, 0xdc, 0x66, 0x7d, 0x21, 0x2e, 0x77, 0x49, 0x5c, 0xa6, 0x18, 0x46, 0x39,
	0xe6, 0x39, 0x84, 0x3f, 0xe8, 0x38, 0x09, 0xfb, 0x19, 0xa1, 0x45, 0x12, 0x51, 0x71, 0x57, 0xc4,
	0x73, 0x98, 0xd7, 0x6c, 0x6b, 0xc5, 0x58, 0x0d, 0x5c, 0xe4, 0xf9, 0xb6, 0x8b, 0x74, 0x69, 0x71,
	0xf2, 0x48, 0x35, 0x92, 0x6d, 0x26, 0xfd, 0x32, 0xaf, 0x50, 0xdb, 0x21, 0x85, 0x39, 0xa6, 0x05,
	0xc9, 0x05, 0xe0, 0x2e, 0x38, 0x3d, 0x09, 0x08, 0xd4, 0x65, 0x1e, 0x47, 0xe9, 0xc5, 0xc9, 0xa4,
	0x37, 0xc9, 0x11, 0x0c, 0x48, 0x46, 0xb8, 0x45, 0x0e, 0x37, 0x78, 0x20, 0x45, 0x00, 0x17, 0xf6,
	0x80, 0xc5, 0x66, 0xe1, 0xe4, 0x99, 0x7c, 0x1f, 0xb7, 0xb5, 0xd9, 0x77, 0xc1, 0x73, 0xba, 0xa7,
	0x05, 0xaa, 0xde, 0x35, 0x7c, 0xe6, 0x30, 0xe6, 0x2f, 0xbf, 0xde, 0xe9, 0x39, 0xdd, 0xd3, 0x60,
	0x18, 0x0f, 0x93, 0xa2, 0xa8, 0x1e, 0xd2, 0xde, 0xd4, 0x71, 0xab, 0x92, 0xbb, 0xe8, 0xfd, 0x2a,
	0x8b, 0x62, 0x7e, 0xd3, 0xea, 0x0d, 0x5e, 0xac, 0xe5, 0x94, 0x56, 0x46, 0x3e, 0x2e, 0x31, 0x2d,
	0x44, 0xc0, 0x1e, 0xf0, 0xcc, 0x2a, 0xbf, 0x09, 0xe6, 0x19, 0xa1, 0x9d, 0x07, 0x36, 0x95, 0x8b,
	0xad, 0xdd, 0xc8, 0x0c, 0x72, 0x40, 0x1a, 0xb8, 0x0a, 0x8e, 0x31, 0x61, 0x7d, 0x91, 0xfa, 0xca,
	0x62, 0xcb, 0xf8, 0x71, 0x21, 0x6a, 0xdf, 0xe9, 0x1b, 0xe0, 0x87, 0xaa, 0x95, 0x6a, 0xc9, 0x40,
	0x81, 0xad, 0x75, 0xa4, 0x66, 0x9d, 0xb6, 0xe8, 0x93, 0x69, 0x70, 0x7a, 0xcd, 0xf6, 0xfc, 0x80,
	0x49, 0x3d, 0x07, 0x69, 0x01, 0x3b, 0x01, 0xd3, 0xd0, 0x7c, 0x29, 0x51, 0xfa, 0xad, 0xbd, 0xde,
	0xf7, 0x07, 0x18, 0xd2, 0x1d, 0x5a, 0xe0, 0x21, 0x1c, 0x84, 0x14, 0xc6, 0xb8, 0xa8, 0x4a, 0x47,
	0x08, 0x23, 0x92, 0x6d, 0xa6, 0x49, 0x54, 0xc0, 0x0d, 0x5c, 0x6c, 0x63, 0x5c, 0xe5, 0x6e, 0xb6,
	0x1a, 0x1d, 0xe1, 0xa8, 0x29, 0x6b, 0x14, 0x92, 0x4d, 0x18, 0x27, 0x9b, 0x9b, 0x38, 0xc7, 0x59,
	0x01, 0xd7, 0x58, 0xe9, 0x20, 0x1b, 0x2c, 0x4a, 0x54, 0xbe, 0x05, 0xd8, 0x86, 0x02, 0x1d, 0x35,
	0xe7, 0xf4, 0x60, 0x22, 0x3b, 0x70, 0x5a, 0x8c, 0x5f, 0x6d, 0x8e, 0x6a, 0x5f, 0x8a, 0x5e, 0x05,
	0x72, 0x45, 0xd0, 0xec, 0x6e, 0xd7, 0xf0, 0x3c, 0xc3, 0x66, 0x79, 0xfe, 0xa1, 0x72, 0xb5, 0x45,
	0xb9, 0x54, 0x51, 0x22, 0x32, 0x14, 0x27, 0x76, 0x10, 0xe8, 0x01, 0x38, 0xa3, 0xf6, 0x7c, 0x3b,
	0x58, 0x45, 0x16, 0x72, 0x55, 0x96, 0x1c, 0x7c, 0xd3, 0x0b, 0x34, 0xe4, 0xfa, 0xd2, 0x77, 0xd3,
	0x8a, 0xd9, 0x8a, 0xc8, 0x3b, 0x6a, 0x59, 0x10, 0xd8, 0xc7, 0x19, 0xbb, 0x76, 0x38, 0x86, 0x11,
	0xce, 0x45, 0x11, 0x17, 0x99, 0x82, 0x95, 0xc2, 0x92, 0x8a, 0xea, 0xaa, 0x3a, 0x06, 0x5c, 0x0d,
	0x0b, 0xbc, 0x1d, 0xee, 0x40, 0xdf, 0xf4, 0x78, 0xf0, 0xaf, 0x81, 0xe7, 0x4c, 0x7b, 0xd5, 0x68,
	0xde, 0xf4, 0x6f, 0xcf, 0x28, 0x2f, 0xb4, 0xfc, 0x9e, 0xe8, 0x51, 0x9c, 0xc3, 0x94, 0xf4, 0x93,
	0xfa, 0x4a, 0xcb, 0xcb, 0x60, 0x5e, 0xed, 0xe9, 0x86, 0x6f, 0x58, 0xab, 0xb5, 0xe8, 0x77, 0x66,
	0x78, 0xa8, 0xea, 0x63, 0x3e, 0x7f, 0x37, 0x4f, 0x0a, 0xb6, 0x32, 0xd9, 0x84, 0x6a, 0x19, 0x27,
	0x05, 0x6f, 0xa2, 0xda, 0x4f, 0xe2, 0x06, 0x90, 0x1c, 0xd5, 0xf3, 0xee, 0xda, 0xae, 0xde, 0x94,
	0xe3, 0xdf, 0x9d, 0x51, 0xce, 0xb7, 0xb6, 0x26, 0x39, 0x21, 0xa5, 0xdb, 0x24, 0x8f, 0xeb, 0x8a,
	0x2c, 0x5f, 0x05, 0xc7, 0x1b, 0x91, 0x8b, 0x3c, 0xe4, 0x4b, 0xbf, 0x37, 0xa3, 0x28, 0x2d, 0xc9,
	0xf1, 0x46, 0x92, 0x63, 0x8a, 0x0b, 0xf9, 0x2a, 0x98, 0x33, 0x0d, 0xeb, 0x76, 0xd0, 0x73, 0xa4,
	0x5f, 0x9f, 0x55, 0x2e, 0x88, 0x6d, 0x4d, 0xb3, 0x6d, 0x39, 0x24, 0x2f, 0x58, 0x84, 0xd2, 0x24,
	0x7b, 0x88, 0x63, 0x58, 0x8e, 0x44, 0x7d, 0x3f, 0xca, 0x05, 0xba, 0x7d, 0xd7, 0x92, 0x7e, 0x63,
	0x56, 0x79, 0x49, 0xc0, 0x99, 0xe4, 0x2c, 0x97, 0xb0, 0xdf, 0x67, 0x5d, 0x5d, 0x51, 0x52, 0x26,
	0x8f, 0xc9, 0x76, 0xd5, 0x19, 0xbc, 0x0e, 0xe6, 0x3d, 0xe4, 0xae, 0x1b, 0x1a, 0x6a, 0x5a, 0xc3,
	0x7f, 0x9e, 0xe5, 0x25, 0x7a, 0xf6, 0xeb, 0xdd, 0xce, 0x0c, 0x2b, 0xd1, 0x1e, 0xce, 0xb7, 0x92,
	0x08, 0xd7, 0x3d, 0xe1, 0x7e, 0x59, 0xd5, 0x61, 0xfd, 0x4b, 0x25, 0x9b, 0xde, 0x2f, 0x1b, 0xb7,
	0x56, 0x1a, 0x38, 0x6d, 0xa9, 0x8e, 0x89, 0xbc, 0x60, 0xaf, 0xfa, 0x5f, 0x67, 0x79, 0x6d, 0x9d,
	0x11, 0xea, 0xd3, 0x56, 0x38, 0x4a, 0x31, 0x85, 0xf4, 0x00, 0xc8, 0x6b, 0xe3, 0xb5, 0x1d, 0x64,
	0xe9, 0x86, 0xb5, 0x2a, 0xfd, 0xdb, 0xac, 0x72, 0xee, 0x20, 0xcb, 0x23, 0x9c, 0xc5, 0x49, 0xd6,
	0x6f, 0xdb, 0x75, 0x7b, 0x96, 0xc5, 0x24, 0xff, 0xfe, 0x84, 0x5d, 0xe6, 0x65, 0x96, 0x25, 0x59,
	0x9f, 0xaf, 0x74, 0x1f, 0x9c, 0xaa, 0x65, 0x2c, 0x07, 0x79, 0x8e, 0x6d, 0x79, 0xc6, 0x3a, 0x92,
	0xfe, 0x63, 0x56, 0xb9, 0xd5, 0x32, 0xfb, 0xe3, 0xb5, 0x96, 0xa5, 0x21, 0x3a, 0x22, 0x19, 0x4d,
	0xb6, 0x30, 0x8c, 0x4b, 0xcc, 0xb2, 0x62, 0x1a, 0x46, 0x0f, 0xd9, 0x7d, 0x12, 0xcf, 0x3f, 0xc7,
	0x94, 0x94, 0x79, 0x84, 0xa9, 0xfc, 0x16, 0x98, 0xf7, 0xee, 0x79, 0x3e, 0xea, 0x06, 0x9a, 0x6d,
	0xea, 0xcb, 0xb6, 0xed, 0x4b, 0xff, 0x39, 0xcb, 0xb3, 0x5b, 0x1d, 0x42, 0xd9, 0xab, 0x34, 0x11,
	0x49, 0x63, 0xb8, 0x41, 0x48, 0x1d, 0x7c, 0x04, 0x16, 0x84, 0xd2, 0x45, 0x9e, 0xdd, 0x73, 0x99,
	0x3d, 0x4f, 0x5d, 0x45, 0xd2, 0x7f, 0xcd, 0x36, 0x5d, 0x38, 0xd3, 0xbf, 0xe0, 0x4d, 0xae, 0x09,
	0x4b, 0x1a, 0xf6, 0xab, 0x6e, 0x2b, 0xe9, 0x0f, 0x64, 0xd4, 0x1c, 0xc6, 0x8a, 0xea, 0xab, 0x66,
	0x60, 0x58, 0x3e, 0x72, 0xdd, 0x9e, 0xe3, 0x4b, 0xff, 0x5d, 0xdd, 0xa0, 0x7a, 0x7f, 0xe7, 0xc5,
	0x61, 0xb0, 0xd4, 0x15, 0xc2, 0xcd, 0xb0, 0x08, 0x53, 0x98, 0x64, 0x05, 0xce, 0xf3, 0x72, 0x54,
	0xc8, 0xeb, 0xe0, 0xac, 0xc0, 0x68, 0xaa, 0xef, 0xbb, 0x86, 0xd3, 0x02, 0xfd, 0xcf, 0xac, 0xf2,
	0x66, 0x0b, 0x74, 0x45, 0x18, 0xc2, 0x59, 0x44, 0x4a, 0x06, 0x60, 0x2f, 0x9f, 0x4f, 0x2f, 0xa3,
	0xea, 0xe6, 0xf3, 0xc7, 0x57, 0x45, 0x4a, 0x36, 0x80, 0x24, 0xb8, 0xf6, 0x3a, 0x72, 0x03, 0x1f,
	0x75, 0x1d, 0xe9, 0x7f, 0x67, 0x95, 0xeb, 0x2d, 0xde, 0x65, 0xc1, 0x2b, 0xf0, 0x90, 0xd7, 0xaf,
	0x32, 0xe7, 0xbb, 0x0b, 0x37, 0xc8, 0x16, 0x86, 0xc5, 0x20, 0xc7, 0x74, 0x40, 0xd2, 0x78, 0x49,
	0xb6, 0xc0, 0xc2, 0x5e, 0x54, 0x80, 0x3e, 0x30, 0x7c, 0xe9, 0xff, 0x2a, 0xde, 0xd4, 0xd3, 0x79,
	0x55, 0xdf, 0xd2, 0xe2, 0xad, 0x00, 0x59, 0xf0, 0x1c, 0xd5, 0x32, 0xb4, 0x00, 0xad, 0x23, 0xcb,
	0x97, 0x7e, 0xa6, 0xa3, 0xbc, 0xda, 0x32, 0x77, 0x41, 0xc0, 0x46, 0x61, 0x96, 0x44, 0x90, 0x54,
	0x99, 0x7d, 0x94, 0xe3, 0xad, 0x84, 0x94, 0x94, 0x1f, 0xa7, 0xdc, 0x05, 0x8b, 0x35, 0x87, 0x97,
	0x11, 0x43, 0xad, 0x32, 0x53, 0xcf, 0x45, 0x82, 0xf9, 0xb3, 0x9d, 0xe6, 0x2d, 0xf3, 0x93, 0xa8,
	0x99, 0xac, 0x00, 0xb0, 0x16, 0x05, 0x16, 0xac, 0x68, 0x8a, 0x04, 0x77, 0x1f, 0x9c, 0x17, 0x38,
	0xc3, 0x5a, 0xb1, 0x03, 0x47, 0x33, 0x50, 0x55, 0xce, 0x04, 0xec, 0xe7, 0x3a, 0xcd, 0x18, 0xc2,
	0x60, 0xaf, 0x78, 0x07, 0x56, 0xa4, 0x51, 0x94, 0xe0, 0x2a, 0x5d, 0x54, 0xb5, 0x0e, 0xf2, 0xc1,
	0x94, 0xcd, 0x6c, 0x02, 0x7e, 0x57, 0x75, 0xad, 0xfd, 0xf0, 0x4f, 0x3a, 0xca, 0xcd, 0x96, 0xd3,
	0xa5, 0x67, 0x84, 0x6f, 0x87, 0x39, 0x7b, 0x6e, 0xf2, 0x87, 0xe0, 0x9c, 0xc0, 0x23, 0xd7, 0xdd,
	0x4f, 0xff, 0xf9, 0xca, 0xfa, 0xcc, 0x0f, 0x6a, 0x3d, 0xcf, 0x49, 0x2e, 0xab, 0xe0, 0x8c, 0x6e,
	0x78, 0xb7, 0x03, 0x7f, 0xcd, 0x45, 0xde, 0x9a, 0x6d, 0xea, 0x01, 0xfa, 0x40, 0x43, 0x48, 0x47,
	0xba, 0xf4, 0xe7, 0x9d, 0x66, 0xae, 0x38, 0xc4, 0x46, 0x21, 0x3d, 0xa1, 0x0f, 0xc7, 0x27, 0x0d,
	0xf1, 0x4f, 0x45, 0x18, 0xc7, 0x38, 0x96, 0x6f, 0x80, 0x79, 0xd7, 0x36, 0x4d, 0xbb, 0xe7, 0x37,
	0x69, 0xf2, 0x0f, 0x0f, 0x35, 0x05, 0x67, 0x96, 0x65, 0x66, 0x97, 0xa4, 0x29, 0x29, 0x0b, 0x98,
	0x64, 0x49, 0x91, 0xb0, 0xa2, 0x26, 0xdf, 0x6a, 0x89, 0x7a, 0x9a, 0x86, 0x3c, 0x4f, 0xfa, 0xa3,
	0x43, 0x4d, 0xb7, 0xc8, 0x44, 0x8b, 0xb5, 0x88, 0x96, 0x51, 0x84, 0x29, 0xdd, 0x2c, 0xd3, 0x74,
	0x07, 0x46, 0x64, 0x38, 0x4a, 0x71, 0x51, 0x55, 0x8e, 0x9a, 0x20, 0x2a, 0xd6, 0x1f, 0x1f, 0xe2,
	0x95, 0x63, 0x46, 0x00, 0x8e, 0xd7, 0x00, 0x71, 0x05, 0x5e, 0x07, 0x27, 0xc6, 0x4b, 0x7a, 0x2c,
	0x3d, 0x22, 0x5d, 0xfa, 0xf6, 0x09, 0x4e, 0x69, 0x49, 0x59, 0x82, 0xc4, 0xb1, 0x7c, 0x1d, 0x1c,
	0x13, 0xcd, 0x23, 0x5b, 0x46, 0xfa, 0xf2, 0x30, 0xef, 0x7e, 0xd8, 0x59, 0x76, 0xbe, 0xd9, 0xed,
	0x9c, 0x9a, 0xec, 0x19, 0xc5, 0x52, 0x3e, 0xf8, 0x61, 0x5f, 0x73, 0x82, 0x35, 0xd5, 0x5c, 0x09,
	0x6c, 0x07, 0x59, 0x81, 0x87, 0x78, 0xeb, 0x10, 0x98, 0x46, 0xd7, 0xf0, 0x03, 0xd5, 0x71, 0x5c,
	0x5b, 0xd5, 0xd6, 0xa4, 0x5f, 0x98, 0x9b, 0xf8, 0x34, 0xf0, 0x82, 0xaf, 0x39, 0xd0, 0xc3, 0x55,
	0xd7, 0x67, 0x26, 0xc3, 0xa4, 0x80, 0xe1, 0x68, 0x94, 0x93, 0x30, 0x1a, 0xb0, 0x7b, 0x60, 0x83,
	0x4b, 0x4f, 0xa3, 0xf2, 0x46, 0x10, 0xe9, 0xd2, 0x2f, 0xce, 0x29, 0x3f, 0xd2, 0x2a, 0xdc, 0xcf,
	0xef, 0x87, 0xf2, 0x86, 0x10, 0xc7, 0xf2, 0x1d, 0x70, 0xa1, 0xa7, 0x3b, 0x81, 0xaa, 0xf9, 0xc6,
	0x3a, 0x7a, 0x92, 0xc7, 0x5f, 0xda, 0xe3, 0xb1, 0xa7, 0x3f, 0xcd, 0xa3, 0x09, 0xe0, 0x13, 0x91,
	0xb5, 0xc1, 0x4f, 0xf7, 0x18, 0xdc, 0x4f, 0xac, 0x0d, 0x7a, 0xe0, 0xa2, 0xa1, 0x75, 0xbf, 0xcf,
	0xe1, 0x2f, 0xcf, 0xf1, 0x8b, 0x53, 0x3b, 0x5c, 0x64, 0x92, 0xa7, 0x58, 0xb4, 0xc0, 0x85, 0x27,
	0x43, 0x6b, 0x8f, 0xbf, 0x32, 0xd7, 0xd4, 0x5c, 0x3e, 0xba, 0x1c, 0xc0, 0xac, 0x4d, 0xf6, 0xc0,
	0x25, 0xdb, 0x5f, 0x63, 0x63, 0xf0, 0x53, 0x5d, 0xfe, 0xea, 0xdc, 0xe4, 0xc7, 0x01, 0xae, 0x79,
	0x8a, 0x4d, 0x07, 0x5c, 0x7c, 0x0a, 0xb6, 0xf6, 0xf9, 0x6b, 0xd5, 0xf7, 0x8a, 0xda, 0xe7, 0xb9,
	0x83, 0xa8, 0xb5, 0x51, 0x03, 0x2c, 0xd8, 0xae, 0xb6, 0xc6, 0x3a, 0x71, 0x4b, 0x7c, 0x02, 0x42,
	0xae, 0x6b, 0xbb, 0xd2, 0x5f, 0xcd, 0x35, 0xd9, 0xf9, 0x30, 0xcb, 0xce, 0xe3, 0x8f, 0x4a, 0x11,
	0xc9, 0x32, 0x1c, 0x15, 0xec, 0x9f, 0x24, 0x8f, 0x06, 0x98, 0x16, 0x79, 0x58, 0x90, 0x5c, 0xee,
	0x01, 0x99, 0xa3, 0xc6, 0x6d, 0x66, 0xcf, 0x45, 0xd2, 0x5f, 0xcf, 0x29, 0xef, 0xb4, 0x38, 0xd7,
	0x4c, 0xde, 0x62, 0x46, 0x39, 0x8e, 0x71, 0x56, 0x24, 0x61, 0x4a, 0xe1, 0x36, 0x66, 0x65, 0x23,
	0xdb, 0x0a, 0xd3, 0xa4, 0x6a, 0x5f, 0xdb, 0xd4, 0x25, 0xf9, 0x23, 0x70, 0xb6, 0x76, 0xc8, 0x1e,
	0x9c, 0xd3, 0xf3, 0xd6, 0x1a, 0xf8, 0xdf, 0xcc, 0x29, 0x3f, 0x29, 0xe2, 0xc7, 0xe0, 0xd7, 0xc7,
	0x26, 0x47, 0x25, 0x1d, 0x40, 0x4a, 0x86, 0x98, 0x4f, 0x0c, 0xcd, 0x63, 0xa4, 0xfb, 0x5c, 0xfb,
	0xe0, 0x24, 0xc7, 0x1b, 0xd6, 0xba, 0x6a, 0x1a, 0x3a, 0x8f, 0xac, 0x6d, 0x49, 0x7f, 0x3b, 0xd7,
	0xa4, 0x67, 0x46, 0x5e, 0xe2, 0xed, 0xf1, 0x08, 0xe7, 0x9b, 0x24, 0x1f, 0xb2, 0xdc, 0x99, 0xc1,
	0xb0, 0xfa, 0x12, 0x56, 0x0c, 0xc2, 0xa2, 0x9a, 0xfb, 0x8b, 0xf1, 0xa8, 0xce, 0xfa, 0xde, 0xae,
	0xb1, 0xea, 0x4e, 0x4c, 0xd1, 0x7f, 0x37, 0xd7, 0xf4, 0xbd, 0x0c, 0x29, 0x75, 0x93, 0xbe, 0xc8,
	0x10, 0x95, 0x6d, 0xf9, 0x6d, 0x70, 0x72, 0x2c, 0xf2, 0x8d, 0x2e, 0xd2, 0x03, 0xbb, 0xe7, 0x4b,
	0x7f, 0x3f, 0xd7, 0x34, 0x8a, 0x4c, 0x77, 0x72, 0xac, 0xf3, 0x13, 0xe6, 0x85, 0x94, 0x85, 0xdc,
	0x03, 0xa7, 0xf8, 0x2e, 0x54, 0xd3, 0x45, 0xaa, 0x7e, 0x2f, 0xe8, 0xaa, 0x96, 0xca, 0x7a, 0xed,
	0x7f, 0xd8, 0xb3, 0x0d, 0x35, 0x23, 0xc5, 0x00, 0xe7, 0xd0, 0xf1, 0xba, 0x70, 0x18, 0xee, 0xc0,
	0x0d, 0x0c, 0x87, 0x61, 0x16, 0xf6, 0xeb, 0x2f, 0x69, 0x34, 0x1c, 0x62, 0x98, 0x85, 0x43, 0x4c,
	0x47, 0x61, 0x84, 0xe5, 0xdb, 0x22, 0xf6, 0xed, 0xef, 0x17, 0xeb, 0xc8, 0x65, 0x77, 0x4e, 0xfa,
	0xc7, 0xb9, 0x66, 0x16, 0x60, 0x68, 0xd8, 0xcb, 0xc6, 0x5f, 0x2a, 0xda, 0x11, 0x86, 0x5b, 0x38,
	0x67, 0x17, 0x4f, 0xbe, 0x2f, 0xae, 0x1a, 0x9b, 0x9d, 0x2c, 0xdb, 0x1f, 0xcf, 0xc9, 0xbb, 0x73,
	0xca, 0x7b, 0xad, 0x57, 0x70, 0xe3, 0x2e, 0xc9, 0x1f, 0xa6, 0x24, 0x8c, 0xe1, 0x28, 0x27, 0x5b,
	0x89, 0x18, 0x9e, 0x48, 0x06, 0x43, 0x38, 0x60, 0x75, 0x7c, 0x3b, 0x29, 0x06, 0x30, 0x23, 0xe3,
	0x59, 0x5a, 0xf7, 0x34, 0x79, 0x05, 0x28, 0x0d, 0xbc, 0x6b, 0xeb, 0x28, 0x30, 0x2c, 0xcd, 0xee,
	0x3a, 0xaa, 0x6f, 0xb0, 0x29, 0xf8, 0x9f, 0xf6, 0x64, 0x17, 0x36, 0x9c, 0x0d, 0xc5, 0x27, 0xb1,
	0x24, 0x63, 0xf5, 0x23, 0x2c, 0x92, 0x8d, 0x14, 0xb3, 0x56, 0xe5, 0xa4, 0x66, 0x9b, 0x26, 0xd2,
	0x7c, 0xdb, 0x0d, 0xc6, 0x63, 0xf4, 0x9f, 0x1c, 0x69, 0x66, 0x44, 0x36, 0x23, 0x9c, 0xd7, 0x48,
	0x9a, 0xe2, 0x88, 0xed, 0xae, 0x19, 0xa4, 0xab, 0xf1, 0x4e, 0xf7, 0x34, 0xde, 0x73, 0xae, 0x81,
	0x85, 0x31, 0xa7, 0x3d, 0x90, 0xff, 0x69, 0x45, 0xaa, 0xa7, 0x0d, 0x38, 0x26, 0xb1, 0x1b, 0xb4,
	0x9f, 0x26, 0xaf, 0x81, 0x33, 0xcb, 0xab, 0x4e, 0xf3, 0xd4, 0x91, 0xe7, 0xab, 0xcb, 0xa6, 0xe1,
	0xb1, 0x87, 0xfe, 0xfb, 0x47, 0x9b, 0xea, 0xc8, 0x58, 0xca, 0xf2, 0xaa, 0x03, 0xa9, 0x78, 0xe6,
	0x09, 0x85, 0x98, 0x16, 0xe1, 0x46, 0x9a, 0xd0, 0x81, 0xe8, 0x83, 0x6f, 0x02, 0xa9, 0x4d, 0xe2,
	0xc3, 0xcf, 0x1f, 0x1c, 0xe5, 0x05, 0xbd, 0xb6, 0x73, 0x72, 0x0f, 0xa2, 0x99, 0x7b, 0x96, 0x81,
	0xb4, 0x62, 0xda, 0x77, 0x4d, 0x7b, 0xd5, 0x0b, 0x74, 0xb7, 0x9a, 0x41, 0x3e, 0x03, 0x3c, 0x7b,
	0xd5, 0x4d, 0xc6, 0xb9, 0x95, 0x94, 0x6c, 0xa7, 0xa4, 0x4f, 0x61, 0x9c, 0xf3, 0xe9, 0x03, 0x86,
	0x45, 0x3d, 0x9e, 0x8b, 0xcf, 0xd8, 0x67, 0x1b, 0x86, 0x8b, 0xd8, 0x75, 0x62, 0xa3, 0x65, 0x95,
	0x6e, 0x3e, 0x07, 0xcd, 0x4b, 0xe6, 0xed, 0x50, 0xc3, 0x8a, 0x48, 0x99, 0xc6, 0x3c, 0x36, 0x1b,
	0x18, 0xe6, 0x58, 0x5c, 0xb0, 0x3a, 0x3a, 0xec, 0x39, 0x3b, 0x5e, 0x57, 0x7c, 0x68, 0x5f, 0x18,
	0xe3, 0x55, 0x1f, 0x55, 0x59, 0x11, 0xe9, 0xd2, 0x17, 0xa0, 0x39, 0x7d, 0xc6, 0x5e, 0x68, 0xd8,
	0x6c, 0x96, 0x86, 0x29, 0xcb, 0x87, 0x55, 0x84, 0x94, 0xe7, 0x3f, 0xfd, 0xad, 0xc5, 0xa9, 0x47,
	0x5f, 0x2d, 0x4e, 0x7d, 0xfd, 0xd5, 0xe2, 0xf4, 0x37, 0x5f, 0x2d, 0x1e, 0x6d, 0xfe, 0x2a, 0x70,
	0xd3, 0x07, 0x47, 0x28, 0xde, 0xc2, 0x79, 0x52, 0xec, 0xc8, 0x17, 0x96, 0xaa, 0x3f, 0x32, 0x2c,
	0xd5, 0x7f, 0x64, 0x58, 0x42, 0x59, 0x39, 0x5c, 0x0f, 0xd3, 0x12, 0xdb, 0x23, 0x9e, 0x66, 0xce,
	0x7e, 0x66, 0xc1, 0xe9, 0x97, 0x8e, 0x5f, 0x3f, 0xb5, 0xc4, 0xbb, 0x41, 0xd6, 0xb3, 0xd3, 0x25,
	0x4f, 0xe8, 0xdd, 0x86, 0xc4, 0xa8, 0x6c, 0x86, 0xef, 0x93, 0xfc, 0x99, 0xa8, 0x9f, 0x1f, 0x40,
	0xd5, 0x84, 0xde, 0x6d, 0x48, 0x37, 0xdf, 0x00, 0x9d, 0x18, 0xd3, 0xe8, 0x59, 0x88, 0x5f, 0x30,
	0xe2, 0x51, 0x97, 0xff, 0xfe, 0x4d, 0x0d, 0x1c, 0x63, 0x6f, 0x38, 0xc7, 0x94, 0x06, 0xc3, 0xe1,
	0xb3, 0xc8, 0x7f, 0x93, 0xc9, 0x8f, 0xb8, 0xa0, 0x96, 0x75, 0x87, 0xcb, 0xd2, 0x77, 0x8f, 0x17,
	0xa7, 0xff, 0xec, 0xf1, 0xe2, 0xf4, 0x5f, 0x3c, 0x5e, 0x9c, 0xfe, 0xcb, 0xc7, 0x8b, 0x53, 0xce,
	0xd4, 0xc6, 0x61, 0x4e, 0xb8, 0xf1, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5d, 0x30, 0xbb, 0x08,
	0xc7, 0x19, 0x00, 0x00,
}
