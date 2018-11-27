// Code generated by protoc-gen-go. DO NOT EDIT.
// source: descriptor_aol.proto

package halproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DescrAolRequest struct {
	DescrAolHandle uint64 `protobuf:"fixed64,1,opt,name=descr_aol_handle,json=descrAolHandle" json:"descr_aol_handle,omitempty"`
}

func (m *DescrAolRequest) Reset()                    { *m = DescrAolRequest{} }
func (m *DescrAolRequest) String() string            { return proto.CompactTextString(m) }
func (*DescrAolRequest) ProtoMessage()               {}
func (*DescrAolRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *DescrAolRequest) GetDescrAolHandle() uint64 {
	if m != nil {
		return m.DescrAolHandle
	}
	return 0
}

type DescrAolRequestMsg struct {
	Request []*DescrAolRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *DescrAolRequestMsg) Reset()                    { *m = DescrAolRequestMsg{} }
func (m *DescrAolRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*DescrAolRequestMsg) ProtoMessage()               {}
func (*DescrAolRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{1} }

func (m *DescrAolRequestMsg) GetRequest() []*DescrAolRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type DescrAolSpec struct {
	ApiStatus          ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	DescrAolHandle     uint64    `protobuf:"fixed64,2,opt,name=descr_aol_handle,json=descrAolHandle" json:"descr_aol_handle,omitempty"`
	Address1           uint64    `protobuf:"fixed64,3,opt,name=Address1" json:"Address1,omitempty"`
	Offset1            uint32    `protobuf:"fixed32,4,opt,name=Offset1" json:"Offset1,omitempty"`
	Length1            uint32    `protobuf:"fixed32,5,opt,name=Length1" json:"Length1,omitempty"`
	Address2           uint64    `protobuf:"fixed64,6,opt,name=Address2" json:"Address2,omitempty"`
	Offset2            uint32    `protobuf:"fixed32,7,opt,name=Offset2" json:"Offset2,omitempty"`
	Length2            uint32    `protobuf:"fixed32,8,opt,name=Length2" json:"Length2,omitempty"`
	Address3           uint64    `protobuf:"fixed64,9,opt,name=Address3" json:"Address3,omitempty"`
	Offset3            uint32    `protobuf:"fixed32,10,opt,name=Offset3" json:"Offset3,omitempty"`
	Length3            uint32    `protobuf:"fixed32,11,opt,name=Length3" json:"Length3,omitempty"`
	NextDescrAolHandle uint64    `protobuf:"fixed64,12,opt,name=next_descr_aol_handle,json=nextDescrAolHandle" json:"next_descr_aol_handle,omitempty"`
}

func (m *DescrAolSpec) Reset()                    { *m = DescrAolSpec{} }
func (m *DescrAolSpec) String() string            { return proto.CompactTextString(m) }
func (*DescrAolSpec) ProtoMessage()               {}
func (*DescrAolSpec) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{2} }

func (m *DescrAolSpec) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *DescrAolSpec) GetDescrAolHandle() uint64 {
	if m != nil {
		return m.DescrAolHandle
	}
	return 0
}

func (m *DescrAolSpec) GetAddress1() uint64 {
	if m != nil {
		return m.Address1
	}
	return 0
}

func (m *DescrAolSpec) GetOffset1() uint32 {
	if m != nil {
		return m.Offset1
	}
	return 0
}

func (m *DescrAolSpec) GetLength1() uint32 {
	if m != nil {
		return m.Length1
	}
	return 0
}

func (m *DescrAolSpec) GetAddress2() uint64 {
	if m != nil {
		return m.Address2
	}
	return 0
}

func (m *DescrAolSpec) GetOffset2() uint32 {
	if m != nil {
		return m.Offset2
	}
	return 0
}

func (m *DescrAolSpec) GetLength2() uint32 {
	if m != nil {
		return m.Length2
	}
	return 0
}

func (m *DescrAolSpec) GetAddress3() uint64 {
	if m != nil {
		return m.Address3
	}
	return 0
}

func (m *DescrAolSpec) GetOffset3() uint32 {
	if m != nil {
		return m.Offset3
	}
	return 0
}

func (m *DescrAolSpec) GetLength3() uint32 {
	if m != nil {
		return m.Length3
	}
	return 0
}

func (m *DescrAolSpec) GetNextDescrAolHandle() uint64 {
	if m != nil {
		return m.NextDescrAolHandle
	}
	return 0
}

type DescrAolResponseMsg struct {
	Response []*DescrAolSpec `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *DescrAolResponseMsg) Reset()                    { *m = DescrAolResponseMsg{} }
func (m *DescrAolResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*DescrAolResponseMsg) ProtoMessage()               {}
func (*DescrAolResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{3} }

func (m *DescrAolResponseMsg) GetResponse() []*DescrAolSpec {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*DescrAolRequest)(nil), "halproto.DescrAolRequest")
	proto.RegisterType((*DescrAolRequestMsg)(nil), "halproto.DescrAolRequestMsg")
	proto.RegisterType((*DescrAolSpec)(nil), "halproto.DescrAolSpec")
	proto.RegisterType((*DescrAolResponseMsg)(nil), "halproto.DescrAolResponseMsg")
}

func init() { proto.RegisterFile("descriptor_aol.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdb, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xad, 0xd3, 0xb5, 0x3b, 0x1d, 0x73, 0xc4, 0x0b, 0x71, 0x28, 0x8c, 0x3e, 0xf5, 0xa9,
	0xd2, 0xe4, 0xd1, 0xa7, 0xc9, 0x40, 0x85, 0x0d, 0xa1, 0x7b, 0x11, 0x5f, 0x4a, 0x5d, 0xb3, 0x0b,
	0x94, 0x36, 0x36, 0x19, 0xe8, 0xff, 0xe0, 0x1f, 0x2d, 0x4d, 0x9b, 0xb5, 0xbb, 0xf8, 0xd4, 0x9c,
	0xf3, 0x7d, 0xf9, 0xe5, 0xd0, 0xef, 0xc0, 0x55, 0xcc, 0xc4, 0x3c, 0x5f, 0x73, 0x99, 0xe5, 0x61,
	0x94, 0x25, 0x1e, 0xcf, 0x33, 0x99, 0x21, 0x4b, 0x75, 0xa3, 0x2c, 0x19, 0xd8, 0xf2, 0x87, 0x33,
	0x51, 0xb6, 0x9d, 0x47, 0xb8, 0x18, 0x17, 0xc2, 0x28, 0x4b, 0x02, 0xf6, 0xb5, 0x61, 0x42, 0x22,
	0x17, 0xfa, 0xca, 0x5b, 0x5c, 0x0e, 0x57, 0x51, 0x1a, 0x27, 0x0c, 0x1b, 0x43, 0xc3, 0x6d, 0x07,
	0xbd, 0xb8, 0xb2, 0xbe, 0xa8, 0xae, 0xf3, 0x0a, 0x68, 0xef, 0xf2, 0x54, 0x2c, 0x11, 0x05, 0x33,
	0x2f, 0x2b, 0x6c, 0x0c, 0x5b, 0xae, 0x4d, 0x6e, 0x3d, 0xfd, 0xb6, 0xb7, 0x67, 0x0f, 0xb4, 0xd3,
	0xf9, 0x6d, 0x41, 0x57, 0x8b, 0x33, 0xce, 0xe6, 0xe8, 0x01, 0x20, 0xe2, 0xeb, 0x50, 0xc8, 0x48,
	0x6e, 0x84, 0x7a, 0xbf, 0x47, 0xfa, 0x5e, 0x39, 0xfa, 0x88, 0xaf, 0x67, 0xaa, 0x1f, 0x74, 0x22,
	0x7d, 0x3c, 0x3a, 0xf6, 0xe9, 0xb1, 0xb1, 0xd1, 0x00, 0xac, 0x51, 0x1c, 0xe7, 0x4c, 0x08, 0x1f,
	0xb7, 0x94, 0x63, 0x5b, 0x23, 0x0c, 0xe6, 0xdb, 0x62, 0x21, 0x98, 0xf4, 0xf1, 0xd9, 0xd0, 0x70,
	0xcd, 0x40, 0x97, 0x85, 0x32, 0x61, 0xe9, 0x52, 0xae, 0x7c, 0x7c, 0x5e, 0x2a, 0x55, 0xd9, 0xe0,
	0x11, 0xdc, 0xde, 0xe1, 0x91, 0x9a, 0x47, 0xb0, 0xd9, 0xe4, 0x91, 0x9a, 0x47, 0xb0, 0xd5, 0xe4,
	0x91, 0x06, 0x8f, 0xe2, 0xce, 0x0e, 0x8f, 0xd6, 0x3c, 0x8a, 0xa1, 0xc9, 0xa3, 0x35, 0x8f, 0x62,
	0xbb, 0xc9, 0xa3, 0xc8, 0x87, 0xeb, 0x94, 0x7d, 0xcb, 0xf0, 0xe0, 0xf7, 0x74, 0x15, 0x1c, 0x15,
	0xe2, 0x78, 0x3f, 0xd9, 0xcb, 0x3a, 0x2a, 0xc1, 0xb3, 0x54, 0xb0, 0x22, 0x5a, 0x02, 0x56, 0x5e,
	0x95, 0x55, 0xb6, 0x37, 0x87, 0xd9, 0x16, 0xf1, 0x05, 0x5b, 0x1f, 0x79, 0x07, 0x4b, 0x2b, 0x68,
	0x02, 0xb6, 0x3e, 0x3f, 0x33, 0x89, 0xee, 0xfe, 0x5d, 0x8c, 0xa9, 0x58, 0x0e, 0xee, 0x8f, 0xa9,
	0xdb, 0x59, 0x9c, 0x93, 0x27, 0xf8, 0xb0, 0x56, 0x51, 0xa2, 0xf6, 0xf8, 0xb3, 0xad, 0x3e, 0xf4,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x7f, 0x6d, 0xf1, 0xaf, 0xfd, 0x02, 0x00, 0x00,
}
