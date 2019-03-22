// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: endpoint.proto

/*
	Package workload is a generated protocol buffer package.

	Service name

	It is generated from these files:
		endpoint.proto
		svc_workload.proto
		workload.proto

	It has these top-level messages:
		Endpoint
		EndpointSpec
		EndpointStatus
		AutoMsgEndpointWatchHelper
		AutoMsgWorkloadWatchHelper
		EndpointList
		WorkloadList
		Workload
		WorkloadIntfSpec
		WorkloadIntfStatus
		WorkloadSpec
		WorkloadStatus
*/
package workload

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pensando/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/pensando/sw/api"
import _ "github.com/pensando/sw/api/labels"
import _ "github.com/pensando/sw/api/generated/security"
import _ "github.com/pensando/sw/api/generated/network"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// Endpoint represents a network endpoint
type Endpoint struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,json=,inline,embedded=T" json:",inline"`
	//
	api.ObjectMeta `protobuf:"bytes,2,opt,name=O,json=meta,omitempty,embedded=O" json:"meta,omitempty"`
	// Spec contains the configuration of the Endpoint.
	Spec EndpointSpec `protobuf:"bytes,3,opt,name=Spec,json=spec,omitempty" json:"spec,omitempty"`
	// Status contains the current state of the Endpoint.
	Status EndpointStatus `protobuf:"bytes,4,opt,name=Status,json=status,omitempty" json:"status,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptorEndpoint, []int{0} }

func (m *Endpoint) GetSpec() EndpointSpec {
	if m != nil {
		return m.Spec
	}
	return EndpointSpec{}
}

func (m *Endpoint) GetStatus() EndpointStatus {
	if m != nil {
		return m.Status
	}
	return EndpointStatus{}
}

// spec part of Endpoint object
type EndpointSpec struct {
}

func (m *EndpointSpec) Reset()                    { *m = EndpointSpec{} }
func (m *EndpointSpec) String() string            { return proto.CompactTextString(m) }
func (*EndpointSpec) ProtoMessage()               {}
func (*EndpointSpec) Descriptor() ([]byte, []int) { return fileDescriptorEndpoint, []int{1} }

// status part of Endpoint object
type EndpointStatus struct {
	// VM or container name
	WorkloadName string `protobuf:"bytes,1,opt,name=WorkloadName,json=workload-name,omitempty,proto3" json:"workload-name,omitempty"`
	// network this endpoint belogs to
	Network string `protobuf:"bytes,2,opt,name=Network,json=network,omitempty,proto3" json:"network,omitempty"`
	// host address of the host where this endpoint exists
	HomingHostAddr string `protobuf:"bytes,3,opt,name=HomingHostAddr,json=homing-host-addr,omitempty,proto3" json:"homing-host-addr,omitempty"`
	// host name of the host where this endpoint exists
	HomingHostName string `protobuf:"bytes,4,opt,name=HomingHostName,json=homing-host-name,omitempty,proto3" json:"homing-host-name,omitempty"`
	// IPv4 address of the endpoint
	IPv4Address string `protobuf:"bytes,5,opt,name=IPv4Address,json=ipv4-address,omitempty,proto3" json:"ipv4-address,omitempty"`
	// IPv4 gateway for the endpoint
	IPv4Gateway string `protobuf:"bytes,6,opt,name=IPv4Gateway,json=ipv4-gateway,omitempty,proto3" json:"ipv4-gateway,omitempty"`
	// IPv6 address for the endpoint
	IPv6Address string `protobuf:"bytes,7,opt,name=IPv6Address,json=ipv6-address,omitempty,proto3" json:"ipv6-address,omitempty"`
	// IPv6 gateway
	IPv6Gateway string `protobuf:"bytes,8,opt,name=IPv6Gateway,json=ipv6-gateway,omitempty,proto3" json:"ipv6-gateway,omitempty"`
	// Mac address of the endpoint
	MacAddress string `protobuf:"bytes,9,opt,name=MacAddress,json=mac-address,omitempty,proto3" json:"mac-address,omitempty"`
	// homing host's UUID
	NodeUUID string `protobuf:"bytes,10,opt,name=NodeUUID,json=node-uuid,omitempty,proto3" json:"node-uuid,omitempty"`
	// endpoint FSM state
	EndpointState string `protobuf:"bytes,11,opt,name=EndpointState,proto3" json:"EndpointState,omitempty"`
	// security groups
	SecurityGroups []string `protobuf:"bytes,12,rep,name=SecurityGroups" json:"SecurityGroups,omitempty"`
	// micro-segment VLAN
	MicroSegmentVlan uint32 `protobuf:"varint,13,opt,name=MicroSegmentVlan,json=micro-segment-vlan,omitempty,proto3" json:"micro-segment-vlan,omitempty"`
	// VM or container attribute/labels
	WorkloadAttributes map[string]string `protobuf:"bytes,14,rep,name=WorkloadAttributes,json=workload-attributes,omitempty" json:"workload-attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *EndpointStatus) Reset()                    { *m = EndpointStatus{} }
func (m *EndpointStatus) String() string            { return proto.CompactTextString(m) }
func (*EndpointStatus) ProtoMessage()               {}
func (*EndpointStatus) Descriptor() ([]byte, []int) { return fileDescriptorEndpoint, []int{2} }

func (m *EndpointStatus) GetWorkloadName() string {
	if m != nil {
		return m.WorkloadName
	}
	return ""
}

func (m *EndpointStatus) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *EndpointStatus) GetHomingHostAddr() string {
	if m != nil {
		return m.HomingHostAddr
	}
	return ""
}

func (m *EndpointStatus) GetHomingHostName() string {
	if m != nil {
		return m.HomingHostName
	}
	return ""
}

func (m *EndpointStatus) GetIPv4Address() string {
	if m != nil {
		return m.IPv4Address
	}
	return ""
}

func (m *EndpointStatus) GetIPv4Gateway() string {
	if m != nil {
		return m.IPv4Gateway
	}
	return ""
}

func (m *EndpointStatus) GetIPv6Address() string {
	if m != nil {
		return m.IPv6Address
	}
	return ""
}

func (m *EndpointStatus) GetIPv6Gateway() string {
	if m != nil {
		return m.IPv6Gateway
	}
	return ""
}

func (m *EndpointStatus) GetMacAddress() string {
	if m != nil {
		return m.MacAddress
	}
	return ""
}

func (m *EndpointStatus) GetNodeUUID() string {
	if m != nil {
		return m.NodeUUID
	}
	return ""
}

func (m *EndpointStatus) GetEndpointState() string {
	if m != nil {
		return m.EndpointState
	}
	return ""
}

func (m *EndpointStatus) GetSecurityGroups() []string {
	if m != nil {
		return m.SecurityGroups
	}
	return nil
}

func (m *EndpointStatus) GetMicroSegmentVlan() uint32 {
	if m != nil {
		return m.MicroSegmentVlan
	}
	return 0
}

func (m *EndpointStatus) GetWorkloadAttributes() map[string]string {
	if m != nil {
		return m.WorkloadAttributes
	}
	return nil
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "workload.Endpoint")
	proto.RegisterType((*EndpointSpec)(nil), "workload.EndpointSpec")
	proto.RegisterType((*EndpointStatus)(nil), "workload.EndpointStatus")
}
func (m *Endpoint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Endpoint) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintEndpoint(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintEndpoint(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintEndpoint(dAtA, i, uint64(m.Spec.Size()))
	n3, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintEndpoint(dAtA, i, uint64(m.Status.Size()))
	n4, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *EndpointSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndpointSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *EndpointStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EndpointStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.WorkloadName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.WorkloadName)))
		i += copy(dAtA[i:], m.WorkloadName)
	}
	if len(m.Network) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.Network)))
		i += copy(dAtA[i:], m.Network)
	}
	if len(m.HomingHostAddr) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.HomingHostAddr)))
		i += copy(dAtA[i:], m.HomingHostAddr)
	}
	if len(m.HomingHostName) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.HomingHostName)))
		i += copy(dAtA[i:], m.HomingHostName)
	}
	if len(m.IPv4Address) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.IPv4Address)))
		i += copy(dAtA[i:], m.IPv4Address)
	}
	if len(m.IPv4Gateway) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.IPv4Gateway)))
		i += copy(dAtA[i:], m.IPv4Gateway)
	}
	if len(m.IPv6Address) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.IPv6Address)))
		i += copy(dAtA[i:], m.IPv6Address)
	}
	if len(m.IPv6Gateway) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.IPv6Gateway)))
		i += copy(dAtA[i:], m.IPv6Gateway)
	}
	if len(m.MacAddress) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.MacAddress)))
		i += copy(dAtA[i:], m.MacAddress)
	}
	if len(m.NodeUUID) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.NodeUUID)))
		i += copy(dAtA[i:], m.NodeUUID)
	}
	if len(m.EndpointState) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(len(m.EndpointState)))
		i += copy(dAtA[i:], m.EndpointState)
	}
	if len(m.SecurityGroups) > 0 {
		for _, s := range m.SecurityGroups {
			dAtA[i] = 0x62
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.MicroSegmentVlan != 0 {
		dAtA[i] = 0x68
		i++
		i = encodeVarintEndpoint(dAtA, i, uint64(m.MicroSegmentVlan))
	}
	if len(m.WorkloadAttributes) > 0 {
		for k, _ := range m.WorkloadAttributes {
			dAtA[i] = 0x72
			i++
			v := m.WorkloadAttributes[k]
			mapSize := 1 + len(k) + sovEndpoint(uint64(len(k))) + 1 + len(v) + sovEndpoint(uint64(len(v)))
			i = encodeVarintEndpoint(dAtA, i, uint64(mapSize))
			dAtA[i] = 0xa
			i++
			i = encodeVarintEndpoint(dAtA, i, uint64(len(k)))
			i += copy(dAtA[i:], k)
			dAtA[i] = 0x12
			i++
			i = encodeVarintEndpoint(dAtA, i, uint64(len(v)))
			i += copy(dAtA[i:], v)
		}
	}
	return i, nil
}

func encodeVarintEndpoint(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Endpoint) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovEndpoint(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovEndpoint(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovEndpoint(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovEndpoint(uint64(l))
	return n
}

func (m *EndpointSpec) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *EndpointStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.WorkloadName)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.Network)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.HomingHostAddr)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.HomingHostName)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.IPv4Address)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.IPv4Gateway)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.IPv6Address)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.IPv6Gateway)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.MacAddress)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.NodeUUID)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	l = len(m.EndpointState)
	if l > 0 {
		n += 1 + l + sovEndpoint(uint64(l))
	}
	if len(m.SecurityGroups) > 0 {
		for _, s := range m.SecurityGroups {
			l = len(s)
			n += 1 + l + sovEndpoint(uint64(l))
		}
	}
	if m.MicroSegmentVlan != 0 {
		n += 1 + sovEndpoint(uint64(m.MicroSegmentVlan))
	}
	if len(m.WorkloadAttributes) > 0 {
		for k, v := range m.WorkloadAttributes {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovEndpoint(uint64(len(k))) + 1 + len(v) + sovEndpoint(uint64(len(v)))
			n += mapEntrySize + 1 + sovEndpoint(uint64(mapEntrySize))
		}
	}
	return n
}

func sovEndpoint(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEndpoint(x uint64) (n int) {
	return sovEndpoint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Endpoint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: Endpoint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Endpoint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TypeMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEndpoint
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
func (m *EndpointSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: EndpointSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndpointSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEndpoint
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
func (m *EndpointStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEndpoint
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
			return fmt.Errorf("proto: EndpointStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EndpointStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkloadName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.WorkloadName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Network", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Network = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomingHostAddr", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomingHostAddr = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomingHostName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomingHostName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPv4Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPv4Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPv4Gateway", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPv4Gateway = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPv6Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPv6Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IPv6Gateway", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IPv6Gateway = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MacAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MacAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeUUID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeUUID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndpointState", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EndpointState = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityGroups", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityGroups = append(m.SecurityGroups, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MicroSegmentVlan", wireType)
			}
			m.MicroSegmentVlan = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MicroSegmentVlan |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WorkloadAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEndpoint
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
				return ErrInvalidLengthEndpoint
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.WorkloadAttributes == nil {
				m.WorkloadAttributes = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowEndpoint
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEndpoint
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthEndpoint
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowEndpoint
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthEndpoint
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipEndpoint(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthEndpoint
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.WorkloadAttributes[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEndpoint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEndpoint
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
func skipEndpoint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
					return 0, ErrIntOverflowEndpoint
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
				return 0, ErrInvalidLengthEndpoint
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEndpoint
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
				next, err := skipEndpoint(dAtA[start:])
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
	ErrInvalidLengthEndpoint = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEndpoint   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("endpoint.proto", fileDescriptorEndpoint) }

var fileDescriptorEndpoint = []byte{
	// 792 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x95, 0x4f, 0x6f, 0xdb, 0x36,
	0x18, 0xc6, 0xa7, 0x24, 0x4d, 0x62, 0x3a, 0xf1, 0x32, 0x76, 0x4d, 0x35, 0x2f, 0xb5, 0x0d, 0xa3,
	0x1d, 0x7c, 0xa8, 0xa5, 0xa2, 0xeb, 0x8c, 0xad, 0xb7, 0x0a, 0xf3, 0x9a, 0x62, 0x8b, 0x53, 0xd8,
	0x6e, 0x77, 0xd9, 0x85, 0x96, 0xde, 0x29, 0x5c, 0x25, 0x52, 0x10, 0x29, 0x07, 0xc6, 0xb0, 0xe3,
	0xbe, 0xc0, 0x80, 0x7d, 0xa7, 0x1c, 0x76, 0x28, 0xf6, 0x01, 0x8c, 0x21, 0xc7, 0x7e, 0x8a, 0x41,
	0x14, 0x15, 0xc8, 0xb1, 0xe4, 0xde, 0xf8, 0xfe, 0x79, 0x7e, 0x7c, 0xf4, 0xd2, 0xa4, 0x51, 0x03,
	0x98, 0x17, 0x71, 0xca, 0xa4, 0x15, 0xc5, 0x5c, 0x72, 0xbc, 0x7f, 0xc9, 0xe3, 0x77, 0x01, 0x27,
	0x5e, 0xf3, 0xc4, 0xe7, 0xdc, 0x0f, 0xc0, 0x26, 0x11, 0xb5, 0x09, 0x63, 0x5c, 0x12, 0x49, 0x39,
	0x13, 0x59, 0x5f, 0x73, 0xe8, 0x53, 0x79, 0x91, 0xcc, 0x2c, 0x97, 0x87, 0x76, 0x04, 0x4c, 0x10,
	0xe6, 0x71, 0x5b, 0x5c, 0xda, 0x73, 0x60, 0xd4, 0x05, 0x3b, 0x91, 0x34, 0x10, 0xa9, 0xd4, 0x07,
	0x56, 0x54, 0xdb, 0x94, 0xb9, 0x41, 0xe2, 0x41, 0x8e, 0xe9, 0x17, 0x30, 0x3e, 0xf7, 0xb9, 0xad,
	0xd2, 0xb3, 0xe4, 0x57, 0x15, 0xa9, 0x40, 0xad, 0x74, 0xfb, 0xa3, 0x8a, 0x5d, 0x53, 0x8f, 0x21,
	0x48, 0xa2, 0xdb, 0x9e, 0x6c, 0x68, 0x0b, 0xc8, 0x0c, 0x02, 0x61, 0x0b, 0x08, 0xc0, 0x95, 0x3c,
	0xd6, 0x8a, 0x6f, 0x36, 0x28, 0x54, 0x47, 0xaa, 0x70, 0x93, 0x98, 0xca, 0x85, 0x1f, 0xf3, 0x24,
	0xd2, 0x32, 0xfb, 0xe3, 0x32, 0x06, 0x32, 0x1d, 0x69, 0x26, 0xe8, 0xfe, 0xb3, 0x85, 0xf6, 0x87,
	0x7a, 0xe2, 0x78, 0x80, 0x8c, 0xa9, 0x69, 0x74, 0x8c, 0x5e, 0xfd, 0xe9, 0xa1, 0x45, 0x22, 0x6a,
	0x4d, 0x17, 0x11, 0x9c, 0x81, 0x24, 0xce, 0xdd, 0xab, 0x65, 0xfb, 0x93, 0xf7, 0xcb, 0xb6, 0xf1,
	0x61, 0xd9, 0xde, 0x7b, 0x4c, 0x59, 0x40, 0x19, 0x8c, 0xf3, 0x05, 0xfe, 0x01, 0x19, 0xe7, 0xe6,
	0x96, 0xd2, 0x7d, 0xaa, 0x74, 0xe7, 0xb3, 0xdf, 0xc0, 0x95, 0x4a, 0xd9, 0x2c, 0x28, 0x1b, 0xe9,
	0x48, 0x1e, 0xf3, 0x90, 0x4a, 0x08, 0x23, 0xb9, 0x18, 0xdf, 0x8a, 0xf1, 0x4f, 0x68, 0x67, 0x12,
	0x81, 0x6b, 0x6e, 0x2b, 0xd4, 0xb1, 0x95, 0x1f, 0xbd, 0x95, 0x3b, 0x4c, 0xab, 0xce, 0x71, 0x4a,
	0x4c, 0x69, 0x22, 0x02, 0xb7, 0x48, 0x5b, 0x8d, 0xf1, 0x14, 0xed, 0x4e, 0x24, 0x91, 0x89, 0x30,
	0x77, 0x14, 0xcf, 0x2c, 0xe1, 0xa9, 0xba, 0x63, 0x6a, 0xe2, 0x91, 0x50, 0x71, 0x81, 0xb9, 0x96,
	0x79, 0x7e, 0xf2, 0xef, 0x9f, 0x5f, 0x98, 0xa8, 0x6e, 0xff, 0x7e, 0x6e, 0x4d, 0x81, 0x11, 0x26,
	0xff, 0xc0, 0xb5, 0xfc, 0x27, 0x2b, 0xba, 0x0d, 0x74, 0x50, 0xf4, 0xda, 0xfd, 0xbb, 0x86, 0x1a,
	0xab, 0x9b, 0xe1, 0x11, 0x3a, 0xf8, 0x59, 0xfb, 0x18, 0x91, 0x10, 0xd4, 0xbc, 0x6b, 0x4e, 0xfb,
	0x2a, 0x1b, 0xd1, 0xfd, 0xdc, 0x63, 0x9f, 0x91, 0x10, 0x0a, 0x4e, 0xaa, 0x0a, 0xf8, 0x3b, 0xb4,
	0x37, 0xca, 0x8e, 0x54, 0x1d, 0x41, 0xcd, 0xb9, 0xf7, 0x61, 0xd9, 0xfe, 0x4c, 0x9f, 0x72, 0x01,
	0xb0, 0x9e, 0xc2, 0x6f, 0x51, 0xe3, 0x94, 0x87, 0x94, 0xf9, 0xa7, 0x5c, 0xc8, 0x17, 0x9e, 0x17,
	0xab, 0xc9, 0xd7, 0x9c, 0xae, 0x36, 0xd3, 0xbc, 0x50, 0xd5, 0xfe, 0x05, 0x17, 0xb2, 0x4f, 0x3c,
	0x2f, 0x2e, 0xe0, 0x36, 0xd4, 0x56, 0xb9, 0xea, 0x23, 0x77, 0xaa, 0xb9, 0xb7, 0xbe, 0x73, 0x43,
	0x0d, 0xff, 0x88, 0xea, 0xaf, 0x5e, 0xcf, 0x9f, 0xa5, 0x4e, 0x41, 0x08, 0xf3, 0x8e, 0x82, 0xb6,
	0x34, 0xf4, 0x98, 0x46, 0xf3, 0x67, 0xca, 0x09, 0x88, 0xe2, 0x11, 0x56, 0xe4, 0x73, 0xd8, 0x4b,
	0x22, 0xe1, 0x92, 0x2c, 0xcc, 0xdd, 0x12, 0x98, 0x9f, 0xd5, 0xd6, 0x60, 0x6b, 0x79, 0x0d, 0x1b,
	0xe4, 0xce, 0xf6, 0xd6, 0x60, 0x83, 0x0a, 0x67, 0x83, 0x4a, 0x67, 0x83, 0xdc, 0xd9, 0x7e, 0x09,
	0xac, 0xdc, 0x59, 0x49, 0x1e, 0x9f, 0x22, 0x74, 0x46, 0xdc, 0xdc, 0x58, 0x4d, 0xb1, 0x1e, 0x68,
	0xd6, 0xbd, 0x90, 0xb8, 0x25, 0xbe, 0xca, 0xd3, 0xd8, 0x41, 0xfb, 0x23, 0xee, 0xc1, 0x9b, 0x37,
	0xaf, 0xbe, 0x37, 0x91, 0xe2, 0x7c, 0xa9, 0x39, 0x77, 0x19, 0xf7, 0xa0, 0x9f, 0x24, 0xd4, 0x2b,
	0x50, 0xca, 0x92, 0xf8, 0x21, 0x3a, 0x2c, 0x5e, 0x07, 0x30, 0xeb, 0x29, 0x68, 0xbc, 0x9a, 0xc4,
	0x5f, 0xa1, 0xc6, 0x44, 0x3f, 0x6e, 0x2f, 0xd3, 0xc7, 0x4d, 0x98, 0x07, 0x9d, 0xed, 0x5e, 0x6d,
	0x7c, 0x2b, 0x8b, 0x7f, 0x41, 0x47, 0x67, 0xd4, 0x8d, 0xf9, 0x04, 0xfc, 0x10, 0x98, 0x7c, 0x1b,
	0x10, 0x66, 0x1e, 0x76, 0x8c, 0xde, 0xa1, 0xf3, 0x50, 0x3b, 0x3b, 0x09, 0xd3, 0x7a, 0x5f, 0x64,
	0x0d, 0xfd, 0x79, 0x40, 0x58, 0xc1, 0xe2, 0xc6, 0x2a, 0xfe, 0xcb, 0x40, 0x38, 0xbf, 0xa9, 0x2f,
	0xa4, 0x8c, 0xe9, 0x2c, 0x91, 0x20, 0xcc, 0x46, 0x67, 0xbb, 0x57, 0x7f, 0xfa, 0xa4, 0xea, 0x31,
	0xb1, 0xd6, 0x25, 0x43, 0x26, 0xe3, 0x85, 0xf3, 0x48, 0x5b, 0x7a, 0x70, 0x73, 0x91, 0xc9, 0x4d,
	0x47, 0xc1, 0xd3, 0xe6, 0x72, 0x73, 0x88, 0xee, 0x57, 0x6c, 0x80, 0x8f, 0xd0, 0xf6, 0x3b, 0x58,
	0x64, 0xef, 0xc9, 0x38, 0x5d, 0xe2, 0xcf, 0xd1, 0x9d, 0x39, 0x09, 0x12, 0xc8, 0x1e, 0x86, 0x71,
	0x16, 0x3c, 0xdf, 0xfa, 0xd6, 0x70, 0x0e, 0xae, 0xae, 0x5b, 0xc6, 0xfb, 0xeb, 0x96, 0xf1, 0xdf,
	0x75, 0xcb, 0x78, 0x6d, 0xcc, 0x76, 0xd5, 0xbf, 0xc1, 0xd7, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff,
	0xf3, 0x60, 0x97, 0x2b, 0x7e, 0x07, 0x00, 0x00,
}
