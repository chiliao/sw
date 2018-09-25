// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: common.proto

/*
	Package iotamodel is a generated protocol buffer package.

	It is generated from these files:
		common.proto
		req_types.proto
		resp_types.proto
		cfg_svc.proto
		topo_svc.proto

	It has these top-level messages:
		ConfigInfo
		ConfigObject
		ConfigTopologyInfo
		App
		Command
		Node
		IotaAPIResponse
		ClusterHealthResponse
		NodeStatus
		InstantiateAppResponse
		InstantiateTopoResponse
		AddNodeResponse
		GeneratedConfigResponse
		ConfigPushResponse
		TriggerAppResponse
*/
package iotamodel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ConfigObject_CfgMethodType int32

const (
	ConfigObject_NONE          ConfigObject_CfgMethodType = 0
	ConfigObject_CREATE        ConfigObject_CfgMethodType = 1
	ConfigObject_UPDATE        ConfigObject_CfgMethodType = 2
	ConfigObject_DELETE        ConfigObject_CfgMethodType = 3
	ConfigObject_CREATE_DELETE ConfigObject_CfgMethodType = 4
)

var ConfigObject_CfgMethodType_name = map[int32]string{
	0: "NONE",
	1: "CREATE",
	2: "UPDATE",
	3: "DELETE",
	4: "CREATE_DELETE",
}
var ConfigObject_CfgMethodType_value = map[string]int32{
	"NONE":          0,
	"CREATE":        1,
	"UPDATE":        2,
	"DELETE":        3,
	"CREATE_DELETE": 4,
}

func (x ConfigObject_CfgMethodType) String() string {
	return proto.EnumName(ConfigObject_CfgMethodType_name, int32(x))
}
func (ConfigObject_CfgMethodType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorCommon, []int{1, 0}
}

type ConfigInfo struct {
	// Configs captures all the user visible objects that need to be pushed.
	Configs []*ConfigObject `protobuf:"bytes,1,rep,name=Configs" json:"configs,omitempty"`
}

func (m *ConfigInfo) Reset()                    { *m = ConfigInfo{} }
func (m *ConfigInfo) String() string            { return proto.CompactTextString(m) }
func (*ConfigInfo) ProtoMessage()               {}
func (*ConfigInfo) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{0} }

func (m *ConfigInfo) GetConfigs() []*ConfigObject {
	if m != nil {
		return m.Configs
	}
	return nil
}

type ConfigObject struct {
	// Method indicates the REST method to call on the object. CREATE_DELETE is used to create delete loop test cases
	Method string `protobuf:"bytes,1,opt,name=Method,proto3" json:"method,omitempty"`
	// ConfigEntrypoint captures where an individual config is expected to be pushed
	ConfigEntrypoint string `protobuf:"bytes,2,opt,name=ConfigEntrypoint,proto3" json:"config-entrypoint,omitempty"`
	// Config captures the Marshaled JSON object
	Config string `protobuf:"bytes,3,opt,name=Config,proto3" json:"config,omitempty"`
}

func (m *ConfigObject) Reset()                    { *m = ConfigObject{} }
func (m *ConfigObject) String() string            { return proto.CompactTextString(m) }
func (*ConfigObject) ProtoMessage()               {}
func (*ConfigObject) Descriptor() ([]byte, []int) { return fileDescriptorCommon, []int{1} }

func (m *ConfigObject) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *ConfigObject) GetConfigEntrypoint() string {
	if m != nil {
		return m.ConfigEntrypoint
	}
	return ""
}

func (m *ConfigObject) GetConfig() string {
	if m != nil {
		return m.Config
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigInfo)(nil), "iotamodel.ConfigInfo")
	proto.RegisterType((*ConfigObject)(nil), "iotamodel.ConfigObject")
	proto.RegisterEnum("iotamodel.ConfigObject_CfgMethodType", ConfigObject_CfgMethodType_name, ConfigObject_CfgMethodType_value)
}
func (m *ConfigInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for _, msg := range m.Configs {
			dAtA[i] = 0xa
			i++
			i = encodeVarintCommon(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ConfigObject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigObject) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Method) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Method)))
		i += copy(dAtA[i:], m.Method)
	}
	if len(m.ConfigEntrypoint) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.ConfigEntrypoint)))
		i += copy(dAtA[i:], m.ConfigEntrypoint)
	}
	if len(m.Config) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCommon(dAtA, i, uint64(len(m.Config)))
		i += copy(dAtA[i:], m.Config)
	}
	return i, nil
}

func encodeVarintCommon(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ConfigInfo) Size() (n int) {
	var l int
	_ = l
	if len(m.Configs) > 0 {
		for _, e := range m.Configs {
			l = e.Size()
			n += 1 + l + sovCommon(uint64(l))
		}
	}
	return n
}

func (m *ConfigObject) Size() (n int) {
	var l int
	_ = l
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.ConfigEntrypoint)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	l = len(m.Config)
	if l > 0 {
		n += 1 + l + sovCommon(uint64(l))
	}
	return n
}

func sovCommon(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCommon(x uint64) (n int) {
	return sovCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConfigInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Configs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Configs = append(m.Configs, &ConfigObject{})
			if err := m.Configs[len(m.Configs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigObject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigObject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigObject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConfigEntrypoint", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConfigEntrypoint = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Config", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Config = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCommon
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCommon
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCommon
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCommon
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCommon
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCommon(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCommon = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCommon   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("common.proto", fileDescriptorCommon) }

var fileDescriptorCommon = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xcf, 0x6e, 0xd4, 0x30,
	0x10, 0xc6, 0x9b, 0x6c, 0x15, 0xa8, 0x69, 0xa5, 0xd4, 0x02, 0x91, 0x16, 0x94, 0x56, 0x0b, 0x87,
	0x45, 0x6a, 0x63, 0x09, 0x8e, 0x9c, 0xd8, 0xad, 0x05, 0x48, 0xd0, 0xae, 0xda, 0xc0, 0x15, 0xe5,
	0x8f, 0xe3, 0x1a, 0xc5, 0x9e, 0x68, 0xe3, 0x80, 0xf2, 0x02, 0x7d, 0x06, 0x9e, 0xa1, 0x4f, 0xc2,
	0x91, 0x23, 0xa7, 0x0a, 0x2d, 0xb7, 0x3e, 0x05, 0x8a, 0x9d, 0x5d, 0x82, 0xb8, 0xcd, 0x37, 0xfe,
	0xcd, 0xf7, 0xcd, 0xc8, 0x68, 0x3b, 0x03, 0x29, 0x41, 0x45, 0xd5, 0x02, 0x34, 0xe0, 0x2d, 0x01,
	0x3a, 0x91, 0x90, 0xb3, 0x72, 0xff, 0x31, 0x07, 0xe0, 0x25, 0x23, 0x49, 0x25, 0x48, 0xa2, 0x14,
	0xe8, 0x44, 0x0b, 0x50, 0xb5, 0x05, 0xf7, 0x8f, 0xb9, 0xd0, 0x97, 0x4d, 0x1a, 0x65, 0x20, 0x09,
	0x07, 0x0e, 0xc4, 0xb4, 0xd3, 0xa6, 0x30, 0xca, 0x08, 0x53, 0xf5, 0x38, 0x1d, 0xe0, 0x15, 0x53,
	0x75, 0xa2, 0x72, 0x20, 0xf5, 0x57, 0xf2, 0x85, 0x29, 0x91, 0x31, 0xd2, 0x68, 0x51, 0xd6, 0x5d,
	0x12, 0x67, 0x6a, 0x18, 0x46, 0x84, 0xca, 0xca, 0x26, 0x67, 0x7d, 0xea, 0xf8, 0x23, 0x42, 0x33,
	0x50, 0x85, 0xe0, 0x6f, 0x55, 0x01, 0xf8, 0x0d, 0xba, 0x63, 0x55, 0x1d, 0x38, 0x87, 0xa3, 0xc9,
	0xbd, 0xe7, 0x0f, 0xa3, 0xf5, 0xfa, 0x91, 0x7d, 0x39, 0x4b, 0x3f, 0xb3, 0x4c, 0x4f, 0x1f, 0xdc,
	0xde, 0x1c, 0xec, 0x66, 0x96, 0x3d, 0x02, 0x29, 0x34, 0x93, 0x95, 0x6e, 0xcf, 0x57, 0xe3, 0xe3,
	0x9f, 0x2e, 0xda, 0x1e, 0x0e, 0xe0, 0x0b, 0xe4, 0xbd, 0x67, 0xfa, 0x12, 0xf2, 0xc0, 0x39, 0x74,
	0x26, 0x5b, 0xd3, 0x97, 0xd7, 0x57, 0x7b, 0x4f, 0x2e, 0xf4, 0x82, 0xaa, 0x46, 0x4e, 0x86, 0x64,
	0x34, 0x2b, 0xb8, 0x25, 0xe3, 0xb6, 0x62, 0xcf, 0x6e, 0x6f, 0x0e, 0x7c, 0x69, 0xe4, 0x20, 0xa6,
	0xb7, 0xc2, 0x35, 0xf2, 0xed, 0x28, 0x55, 0x7a, 0xd1, 0x56, 0x20, 0x94, 0x0e, 0x5c, 0x63, 0xff,
	0xfa, 0xfa, 0x6a, 0xef, 0xe9, 0xca, 0x3e, 0x86, 0x0a, 0x4a, 0xe0, 0x6d, 0x77, 0x63, 0xf4, 0x17,
	0x5d, 0xf9, 0x3f, 0xb2, 0x77, 0x1c, 0xb3, 0xf5, 0xcb, 0x20, 0xea, 0xbf, 0x00, 0x7c, 0x84, 0x3c,
	0xdb, 0x0b, 0x46, 0x26, 0xea, 0x7e, 0xb7, 0xa2, 0xb5, 0x18, 0xae, 0x68, 0x99, 0xf1, 0x1c, 0xed,
	0xfc, 0x73, 0x10, 0xbe, 0x8b, 0x36, 0x4f, 0xcf, 0x4e, 0xa9, 0xbf, 0x81, 0x11, 0xf2, 0x66, 0xe7,
	0xf4, 0x55, 0x4c, 0x7d, 0xa7, 0xab, 0x3f, 0xcc, 0x4f, 0xba, 0xda, 0xed, 0xea, 0x13, 0xfa, 0x8e,
	0xc6, 0xd4, 0x1f, 0xe1, 0x5d, 0xb4, 0x63, 0x99, 0x4f, 0x7d, 0x6b, 0x73, 0xea, 0x7f, 0x5f, 0x86,
	0xce, 0x8f, 0x65, 0xe8, 0xfc, 0x5a, 0x86, 0xce, 0xb7, 0xdf, 0xe1, 0xc6, 0xdc, 0x4d, 0x3d, 0xf3,
	0x9b, 0x2f, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x9a, 0xb5, 0x58, 0x7c, 0x02, 0x00, 0x00,
}
