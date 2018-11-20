// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: descriptor_aol.proto

package halproto

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import encoding_binary "encoding/binary"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DescrAolRequest struct {
	DescrAolHandle uint64 `protobuf:"fixed64,1,opt,name=descr_aol_handle,json=descrAolHandle,proto3" json:"descr_aol_handle,omitempty"`
}

func (m *DescrAolRequest) Reset()                    { *m = DescrAolRequest{} }
func (m *DescrAolRequest) String() string            { return proto.CompactTextString(m) }
func (*DescrAolRequest) ProtoMessage()               {}
func (*DescrAolRequest) Descriptor() ([]byte, []int) { return fileDescriptorDescriptorAol, []int{0} }

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
func (*DescrAolRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptorDescriptorAol, []int{1} }

func (m *DescrAolRequestMsg) GetRequest() []*DescrAolRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type DescrAolSpec struct {
	ApiStatus          ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,proto3,enum=types.ApiStatus" json:"api_status,omitempty"`
	DescrAolHandle     uint64    `protobuf:"fixed64,2,opt,name=descr_aol_handle,json=descrAolHandle,proto3" json:"descr_aol_handle,omitempty"`
	Address1           uint64    `protobuf:"fixed64,3,opt,name=Address1,proto3" json:"Address1,omitempty"`
	Offset1            uint32    `protobuf:"fixed32,4,opt,name=Offset1,proto3" json:"Offset1,omitempty"`
	Length1            uint32    `protobuf:"fixed32,5,opt,name=Length1,proto3" json:"Length1,omitempty"`
	Address2           uint64    `protobuf:"fixed64,6,opt,name=Address2,proto3" json:"Address2,omitempty"`
	Offset2            uint32    `protobuf:"fixed32,7,opt,name=Offset2,proto3" json:"Offset2,omitempty"`
	Length2            uint32    `protobuf:"fixed32,8,opt,name=Length2,proto3" json:"Length2,omitempty"`
	Address3           uint64    `protobuf:"fixed64,9,opt,name=Address3,proto3" json:"Address3,omitempty"`
	Offset3            uint32    `protobuf:"fixed32,10,opt,name=Offset3,proto3" json:"Offset3,omitempty"`
	Length3            uint32    `protobuf:"fixed32,11,opt,name=Length3,proto3" json:"Length3,omitempty"`
	NextDescrAolHandle uint64    `protobuf:"fixed64,12,opt,name=next_descr_aol_handle,json=nextDescrAolHandle,proto3" json:"next_descr_aol_handle,omitempty"`
}

func (m *DescrAolSpec) Reset()                    { *m = DescrAolSpec{} }
func (m *DescrAolSpec) String() string            { return proto.CompactTextString(m) }
func (*DescrAolSpec) ProtoMessage()               {}
func (*DescrAolSpec) Descriptor() ([]byte, []int) { return fileDescriptorDescriptorAol, []int{2} }

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
func (*DescrAolResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptorDescriptorAol, []int{3} }

func (m *DescrAolResponseMsg) GetResponse() []*DescrAolSpec {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*DescrAolRequest)(nil), "descraol.DescrAolRequest")
	proto.RegisterType((*DescrAolRequestMsg)(nil), "descraol.DescrAolRequestMsg")
	proto.RegisterType((*DescrAolSpec)(nil), "descraol.DescrAolSpec")
	proto.RegisterType((*DescrAolResponseMsg)(nil), "descraol.DescrAolResponseMsg")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DescrAol service

type DescrAolClient interface {
	DescrAolGet(ctx context.Context, in *DescrAolRequestMsg, opts ...grpc.CallOption) (*DescrAolResponseMsg, error)
}

type descrAolClient struct {
	cc *grpc.ClientConn
}

func NewDescrAolClient(cc *grpc.ClientConn) DescrAolClient {
	return &descrAolClient{cc}
}

func (c *descrAolClient) DescrAolGet(ctx context.Context, in *DescrAolRequestMsg, opts ...grpc.CallOption) (*DescrAolResponseMsg, error) {
	out := new(DescrAolResponseMsg)
	err := grpc.Invoke(ctx, "/descraol.DescrAol/DescrAolGet", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DescrAol service

type DescrAolServer interface {
	DescrAolGet(context.Context, *DescrAolRequestMsg) (*DescrAolResponseMsg, error)
}

func RegisterDescrAolServer(s *grpc.Server, srv DescrAolServer) {
	s.RegisterService(&_DescrAol_serviceDesc, srv)
}

func _DescrAol_DescrAolGet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescrAolRequestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DescrAolServer).DescrAolGet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/descraol.DescrAol/DescrAolGet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DescrAolServer).DescrAolGet(ctx, req.(*DescrAolRequestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

var _DescrAol_serviceDesc = grpc.ServiceDesc{
	ServiceName: "descraol.DescrAol",
	HandlerType: (*DescrAolServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DescrAolGet",
			Handler:    _DescrAol_DescrAolGet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "descriptor_aol.proto",
}

func (m *DescrAolRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DescrAolRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.DescrAolHandle != 0 {
		dAtA[i] = 0x9
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.DescrAolHandle))
		i += 8
	}
	return i, nil
}

func (m *DescrAolRequestMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DescrAolRequestMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Request) > 0 {
		for _, msg := range m.Request {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDescriptorAol(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *DescrAolSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DescrAolSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ApiStatus != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintDescriptorAol(dAtA, i, uint64(m.ApiStatus))
	}
	if m.DescrAolHandle != 0 {
		dAtA[i] = 0x11
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.DescrAolHandle))
		i += 8
	}
	if m.Address1 != 0 {
		dAtA[i] = 0x19
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Address1))
		i += 8
	}
	if m.Offset1 != 0 {
		dAtA[i] = 0x25
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Offset1))
		i += 4
	}
	if m.Length1 != 0 {
		dAtA[i] = 0x2d
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Length1))
		i += 4
	}
	if m.Address2 != 0 {
		dAtA[i] = 0x31
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Address2))
		i += 8
	}
	if m.Offset2 != 0 {
		dAtA[i] = 0x3d
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Offset2))
		i += 4
	}
	if m.Length2 != 0 {
		dAtA[i] = 0x45
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Length2))
		i += 4
	}
	if m.Address3 != 0 {
		dAtA[i] = 0x49
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.Address3))
		i += 8
	}
	if m.Offset3 != 0 {
		dAtA[i] = 0x55
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Offset3))
		i += 4
	}
	if m.Length3 != 0 {
		dAtA[i] = 0x5d
		i++
		encoding_binary.LittleEndian.PutUint32(dAtA[i:], uint32(m.Length3))
		i += 4
	}
	if m.NextDescrAolHandle != 0 {
		dAtA[i] = 0x61
		i++
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(m.NextDescrAolHandle))
		i += 8
	}
	return i, nil
}

func (m *DescrAolResponseMsg) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DescrAolResponseMsg) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Response) > 0 {
		for _, msg := range m.Response {
			dAtA[i] = 0xa
			i++
			i = encodeVarintDescriptorAol(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintDescriptorAol(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *DescrAolRequest) Size() (n int) {
	var l int
	_ = l
	if m.DescrAolHandle != 0 {
		n += 9
	}
	return n
}

func (m *DescrAolRequestMsg) Size() (n int) {
	var l int
	_ = l
	if len(m.Request) > 0 {
		for _, e := range m.Request {
			l = e.Size()
			n += 1 + l + sovDescriptorAol(uint64(l))
		}
	}
	return n
}

func (m *DescrAolSpec) Size() (n int) {
	var l int
	_ = l
	if m.ApiStatus != 0 {
		n += 1 + sovDescriptorAol(uint64(m.ApiStatus))
	}
	if m.DescrAolHandle != 0 {
		n += 9
	}
	if m.Address1 != 0 {
		n += 9
	}
	if m.Offset1 != 0 {
		n += 5
	}
	if m.Length1 != 0 {
		n += 5
	}
	if m.Address2 != 0 {
		n += 9
	}
	if m.Offset2 != 0 {
		n += 5
	}
	if m.Length2 != 0 {
		n += 5
	}
	if m.Address3 != 0 {
		n += 9
	}
	if m.Offset3 != 0 {
		n += 5
	}
	if m.Length3 != 0 {
		n += 5
	}
	if m.NextDescrAolHandle != 0 {
		n += 9
	}
	return n
}

func (m *DescrAolResponseMsg) Size() (n int) {
	var l int
	_ = l
	if len(m.Response) > 0 {
		for _, e := range m.Response {
			l = e.Size()
			n += 1 + l + sovDescriptorAol(uint64(l))
		}
	}
	return n
}

func sovDescriptorAol(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozDescriptorAol(x uint64) (n int) {
	return sovDescriptorAol(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DescrAolRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDescriptorAol
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
			return fmt.Errorf("proto: DescrAolRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DescrAolRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DescrAolHandle", wireType)
			}
			m.DescrAolHandle = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.DescrAolHandle = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		default:
			iNdEx = preIndex
			skippy, err := skipDescriptorAol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDescriptorAol
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
func (m *DescrAolRequestMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDescriptorAol
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
			return fmt.Errorf("proto: DescrAolRequestMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DescrAolRequestMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Request", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptorAol
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
				return ErrInvalidLengthDescriptorAol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Request = append(m.Request, &DescrAolRequest{})
			if err := m.Request[len(m.Request)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDescriptorAol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDescriptorAol
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
func (m *DescrAolSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDescriptorAol
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
			return fmt.Errorf("proto: DescrAolSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DescrAolSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ApiStatus", wireType)
			}
			m.ApiStatus = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptorAol
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ApiStatus |= (ApiStatus(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field DescrAolHandle", wireType)
			}
			m.DescrAolHandle = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.DescrAolHandle = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 3:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address1", wireType)
			}
			m.Address1 = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Address1 = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 4:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset1", wireType)
			}
			m.Offset1 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Offset1 = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 5:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length1", wireType)
			}
			m.Length1 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Length1 = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 6:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address2", wireType)
			}
			m.Address2 = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Address2 = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 7:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset2", wireType)
			}
			m.Offset2 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Offset2 = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 8:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length2", wireType)
			}
			m.Length2 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Length2 = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 9:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address3", wireType)
			}
			m.Address3 = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.Address3 = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		case 10:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset3", wireType)
			}
			m.Offset3 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Offset3 = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 11:
			if wireType != 5 {
				return fmt.Errorf("proto: wrong wireType = %d for field Length3", wireType)
			}
			m.Length3 = 0
			if (iNdEx + 4) > l {
				return io.ErrUnexpectedEOF
			}
			m.Length3 = uint32(encoding_binary.LittleEndian.Uint32(dAtA[iNdEx:]))
			iNdEx += 4
		case 12:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextDescrAolHandle", wireType)
			}
			m.NextDescrAolHandle = 0
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			m.NextDescrAolHandle = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
		default:
			iNdEx = preIndex
			skippy, err := skipDescriptorAol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDescriptorAol
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
func (m *DescrAolResponseMsg) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDescriptorAol
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
			return fmt.Errorf("proto: DescrAolResponseMsg: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DescrAolResponseMsg: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Response", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDescriptorAol
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
				return ErrInvalidLengthDescriptorAol
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Response = append(m.Response, &DescrAolSpec{})
			if err := m.Response[len(m.Response)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDescriptorAol(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDescriptorAol
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
func skipDescriptorAol(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDescriptorAol
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
					return 0, ErrIntOverflowDescriptorAol
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
					return 0, ErrIntOverflowDescriptorAol
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
				return 0, ErrInvalidLengthDescriptorAol
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDescriptorAol
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
				next, err := skipDescriptorAol(dAtA[start:])
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
	ErrInvalidLengthDescriptorAol = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDescriptorAol   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("descriptor_aol.proto", fileDescriptorDescriptorAol) }

var fileDescriptorDescriptorAol = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x4e, 0x83, 0x40,
	0x14, 0x85, 0xc5, 0x6a, 0xa1, 0x97, 0xa6, 0x36, 0xa3, 0x36, 0x63, 0xa3, 0x4d, 0xc3, 0x8a, 0x15,
	0x86, 0x99, 0xa5, 0xab, 0x9a, 0x26, 0x6a, 0xd2, 0xc6, 0x84, 0x6e, 0x8c, 0x1b, 0x82, 0x65, 0xfa,
	0x93, 0x10, 0x40, 0x66, 0x9a, 0xe8, 0x3b, 0xf8, 0x60, 0x2e, 0x7d, 0x04, 0xd3, 0x27, 0x31, 0x0c,
	0x50, 0xe8, 0x8f, 0x2b, 0xe6, 0xde, 0x73, 0xe6, 0x9b, 0x1b, 0xce, 0x85, 0x0b, 0x9f, 0xf1, 0x69,
	0xb2, 0x8c, 0x45, 0x94, 0xb8, 0x5e, 0x14, 0x58, 0x71, 0x12, 0x89, 0x08, 0x69, 0xb2, 0xeb, 0x45,
	0x41, 0x57, 0x17, 0x9f, 0x31, 0xe3, 0x59, 0xdb, 0xb8, 0x83, 0xb3, 0x61, 0x2a, 0x0c, 0xa2, 0xc0,
	0x61, 0xef, 0x2b, 0xc6, 0x05, 0x32, 0xa1, 0x2d, 0xbd, 0xe9, 0x65, 0x77, 0xe1, 0x85, 0x7e, 0xc0,
	0xb0, 0xd2, 0x57, 0xcc, 0xba, 0xd3, 0xf2, 0x73, 0xeb, 0xa3, 0xec, 0x1a, 0x4f, 0x80, 0x76, 0x2e,
	0x8f, 0xf9, 0x1c, 0x51, 0x50, 0x93, 0xac, 0xc2, 0x4a, 0xbf, 0x66, 0xea, 0xe4, 0xca, 0x2a, 0xde,
	0xb6, 0x76, 0xec, 0x4e, 0xe1, 0x34, 0xbe, 0x6a, 0xd0, 0x2c, 0xc4, 0x49, 0xcc, 0xa6, 0xe8, 0x16,
	0xc0, 0x8b, 0x97, 0x2e, 0x17, 0x9e, 0x58, 0x71, 0xf9, 0x7e, 0x8b, 0xb4, 0xad, 0x6c, 0xf4, 0x41,
	0xbc, 0x9c, 0xc8, 0xbe, 0xd3, 0xf0, 0x8a, 0xe3, 0xc1, 0xb1, 0x8f, 0x0f, 0x8d, 0x8d, 0xba, 0xa0,
	0x0d, 0x7c, 0x3f, 0x61, 0x9c, 0xdb, 0xb8, 0x26, 0x1d, 0x9b, 0x1a, 0x61, 0x50, 0x9f, 0x67, 0x33,
	0xce, 0x84, 0x8d, 0x4f, 0xfa, 0x8a, 0xa9, 0x3a, 0x45, 0x99, 0x2a, 0x23, 0x16, 0xce, 0xc5, 0xc2,
	0xc6, 0xa7, 0x99, 0x92, 0x97, 0x15, 0x1e, 0xc1, 0xf5, 0x2d, 0x1e, 0x29, 0x79, 0x04, 0xab, 0x55,
	0x1e, 0x29, 0x79, 0x04, 0x6b, 0x55, 0x1e, 0xa9, 0xf0, 0x28, 0x6e, 0x6c, 0xf1, 0x68, 0xc9, 0xa3,
	0x18, 0xaa, 0x3c, 0x5a, 0xf2, 0x28, 0xd6, 0xab, 0x3c, 0x8a, 0x6c, 0xb8, 0x0c, 0xd9, 0x87, 0x70,
	0xf7, 0x7e, 0x4f, 0x53, 0xc2, 0x51, 0x2a, 0x0e, 0x77, 0x93, 0x3d, 0x2f, 0xa3, 0xe2, 0x71, 0x14,
	0x72, 0x96, 0x46, 0x4b, 0x40, 0x4b, 0xf2, 0x32, 0xcf, 0xb6, 0xb3, 0x9f, 0x6d, 0x1a, 0x9f, 0xb3,
	0xf1, 0x91, 0x17, 0xd0, 0x0a, 0x05, 0x8d, 0x40, 0x2f, 0xce, 0x0f, 0x4c, 0xa0, 0xeb, 0x7f, 0x17,
	0x63, 0xcc, 0xe7, 0xdd, 0x9b, 0x43, 0xea, 0x66, 0x16, 0xe3, 0xe8, 0xbe, 0xf3, 0xbd, 0xee, 0x29,
	0x3f, 0xeb, 0x9e, 0xf2, 0xbb, 0xee, 0x29, 0xaf, 0xda, 0xc2, 0x0b, 0xe4, 0x4e, 0xbf, 0xd5, 0xe5,
	0x87, 0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xb7, 0x13, 0xa9, 0x89, 0x09, 0x03, 0x00, 0x00,
}
