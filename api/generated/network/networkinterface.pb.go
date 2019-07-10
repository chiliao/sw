// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: networkinterface.proto

package network

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pensando/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/pensando/sw/api"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
type NetworkInterfaceStatus_IFStatus int32

const (
	//
	NetworkInterfaceStatus_UP NetworkInterfaceStatus_IFStatus = 0
	//
	NetworkInterfaceStatus_DOWN NetworkInterfaceStatus_IFStatus = 1
)

var NetworkInterfaceStatus_IFStatus_name = map[int32]string{
	0: "UP",
	1: "DOWN",
}
var NetworkInterfaceStatus_IFStatus_value = map[string]int32{
	"UP":   0,
	"DOWN": 1,
}

func (x NetworkInterfaceStatus_IFStatus) String() string {
	return proto.EnumName(NetworkInterfaceStatus_IFStatus_name, int32(x))
}
func (NetworkInterfaceStatus_IFStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorNetworkinterface, []int{3, 0}
}

//
type NetworkInterfaceStatus_IFType int32

const (
	//
	NetworkInterfaceStatus_NONE NetworkInterfaceStatus_IFType = 0
	//
	NetworkInterfaceStatus_HOST_PF NetworkInterfaceStatus_IFType = 1
	//
	NetworkInterfaceStatus_UPLINK_ETH NetworkInterfaceStatus_IFType = 3
	//
	NetworkInterfaceStatus_UPLINK_MGMT NetworkInterfaceStatus_IFType = 4
)

var NetworkInterfaceStatus_IFType_name = map[int32]string{
	0: "NONE",
	1: "HOST_PF",
	3: "UPLINK_ETH",
	4: "UPLINK_MGMT",
}
var NetworkInterfaceStatus_IFType_value = map[string]int32{
	"NONE":        0,
	"HOST_PF":     1,
	"UPLINK_ETH":  3,
	"UPLINK_MGMT": 4,
}

func (x NetworkInterfaceStatus_IFType) String() string {
	return proto.EnumName(NetworkInterfaceStatus_IFType_name, int32(x))
}
func (NetworkInterfaceStatus_IFType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorNetworkinterface, []int{3, 1}
}

//
type NetworkInterface struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,json=,inline,embedded=T" json:",inline"`
	// Object name is Serial-Number of the SmartNIC
	api.ObjectMeta `protobuf:"bytes,2,opt,name=O,json=meta,omitempty,embedded=O" json:"meta,omitempty"`
	// NetworkInterfaceSpec contains the configuration of the network adapter.
	Spec NetworkInterfaceSpec `protobuf:"bytes,3,opt,name=Spec,json=spec,omitempty" json:"spec,omitempty"`
	// NetworkInterfaceStatus contains the current state of the network adapter.
	Status NetworkInterfaceStatus `protobuf:"bytes,4,opt,name=Status,json=status,omitempty" json:"status,omitempty"`
}

func (m *NetworkInterface) Reset()                    { *m = NetworkInterface{} }
func (m *NetworkInterface) String() string            { return proto.CompactTextString(m) }
func (*NetworkInterface) ProtoMessage()               {}
func (*NetworkInterface) Descriptor() ([]byte, []int) { return fileDescriptorNetworkinterface, []int{0} }

func (m *NetworkInterface) GetSpec() NetworkInterfaceSpec {
	if m != nil {
		return m.Spec
	}
	return NetworkInterfaceSpec{}
}

func (m *NetworkInterface) GetStatus() NetworkInterfaceStatus {
	if m != nil {
		return m.Status
	}
	return NetworkInterfaceStatus{}
}

// NetworkInterfaceHostStatus is populated for PF and VF
type NetworkInterfaceHostStatus struct {
	// interface name seen by the host driver
	HostIfName string `protobuf:"bytes,1,opt,name=HostIfName,json=host-ifname,omitempty,proto3" json:"host-ifname,omitempty"`
}

func (m *NetworkInterfaceHostStatus) Reset()         { *m = NetworkInterfaceHostStatus{} }
func (m *NetworkInterfaceHostStatus) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceHostStatus) ProtoMessage()    {}
func (*NetworkInterfaceHostStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorNetworkinterface, []int{1}
}

func (m *NetworkInterfaceHostStatus) GetHostIfName() string {
	if m != nil {
		return m.HostIfName
	}
	return ""
}

// NetworkInterfaceSpec
type NetworkInterfaceSpec struct {
}

func (m *NetworkInterfaceSpec) Reset()         { *m = NetworkInterfaceSpec{} }
func (m *NetworkInterfaceSpec) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceSpec) ProtoMessage()    {}
func (*NetworkInterfaceSpec) Descriptor() ([]byte, []int) {
	return fileDescriptorNetworkinterface, []int{2}
}

// NetworkInterfaceStatus
type NetworkInterfaceStatus struct {
	//
	SmartNIC string `protobuf:"bytes,1,opt,name=SmartNIC,json=smart-nic,omitempty,proto3" json:"smart-nic,omitempty"`
	//
	Type string `protobuf:"bytes,2,opt,name=Type,json=type,omitempty,proto3" json:"type,omitempty"`
	//
	OperStatus string `protobuf:"bytes,3,opt,name=OperStatus,json=oper-status,omitempty,proto3" json:"oper-status,omitempty"`
	//
	PrimaryMac string `protobuf:"bytes,4,opt,name=PrimaryMac,json=primary-mac,omitempty,proto3" json:"primary-mac,omitempty"`
	//
	IFHostStatus *NetworkInterfaceHostStatus `protobuf:"bytes,5,opt,name=IFHostStatus,json=if-host-status,omitempty" json:"if-host-status,omitempty"`
	//
	IFUplinkStatus *NetworkInterfaceUplinkStatus `protobuf:"bytes,6,opt,name=IFUplinkStatus,json=if-uplink-status,omitempty" json:"if-uplink-status,omitempty"`
}

func (m *NetworkInterfaceStatus) Reset()         { *m = NetworkInterfaceStatus{} }
func (m *NetworkInterfaceStatus) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceStatus) ProtoMessage()    {}
func (*NetworkInterfaceStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorNetworkinterface, []int{3}
}

func (m *NetworkInterfaceStatus) GetSmartNIC() string {
	if m != nil {
		return m.SmartNIC
	}
	return ""
}

func (m *NetworkInterfaceStatus) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *NetworkInterfaceStatus) GetOperStatus() string {
	if m != nil {
		return m.OperStatus
	}
	return ""
}

func (m *NetworkInterfaceStatus) GetPrimaryMac() string {
	if m != nil {
		return m.PrimaryMac
	}
	return ""
}

func (m *NetworkInterfaceStatus) GetIFHostStatus() *NetworkInterfaceHostStatus {
	if m != nil {
		return m.IFHostStatus
	}
	return nil
}

func (m *NetworkInterfaceStatus) GetIFUplinkStatus() *NetworkInterfaceUplinkStatus {
	if m != nil {
		return m.IFUplinkStatus
	}
	return nil
}

//
type NetworkInterfaceUplinkStatus struct {
	// LinkSpeed auto-negotiated
	LinkSpeed string `protobuf:"bytes,2,opt,name=LinkSpeed,json=link-speed,omitempty,proto3" json:"link-speed,omitempty"`
}

func (m *NetworkInterfaceUplinkStatus) Reset()         { *m = NetworkInterfaceUplinkStatus{} }
func (m *NetworkInterfaceUplinkStatus) String() string { return proto.CompactTextString(m) }
func (*NetworkInterfaceUplinkStatus) ProtoMessage()    {}
func (*NetworkInterfaceUplinkStatus) Descriptor() ([]byte, []int) {
	return fileDescriptorNetworkinterface, []int{4}
}

func (m *NetworkInterfaceUplinkStatus) GetLinkSpeed() string {
	if m != nil {
		return m.LinkSpeed
	}
	return ""
}

func init() {
	proto.RegisterType((*NetworkInterface)(nil), "network.NetworkInterface")
	proto.RegisterType((*NetworkInterfaceHostStatus)(nil), "network.NetworkInterfaceHostStatus")
	proto.RegisterType((*NetworkInterfaceSpec)(nil), "network.NetworkInterfaceSpec")
	proto.RegisterType((*NetworkInterfaceStatus)(nil), "network.NetworkInterfaceStatus")
	proto.RegisterType((*NetworkInterfaceUplinkStatus)(nil), "network.NetworkInterfaceUplinkStatus")
	proto.RegisterEnum("network.NetworkInterfaceStatus_IFStatus", NetworkInterfaceStatus_IFStatus_name, NetworkInterfaceStatus_IFStatus_value)
	proto.RegisterEnum("network.NetworkInterfaceStatus_IFType", NetworkInterfaceStatus_IFType_name, NetworkInterfaceStatus_IFType_value)
}
func (m *NetworkInterface) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkInterface) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintNetworkinterface(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintNetworkinterface(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintNetworkinterface(dAtA, i, uint64(m.Spec.Size()))
	n3, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintNetworkinterface(dAtA, i, uint64(m.Status.Size()))
	n4, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *NetworkInterfaceHostStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkInterfaceHostStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.HostIfName) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkinterface(dAtA, i, uint64(len(m.HostIfName)))
		i += copy(dAtA[i:], m.HostIfName)
	}
	return i, nil
}

func (m *NetworkInterfaceSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkInterfaceSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *NetworkInterfaceStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkInterfaceStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.SmartNIC) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintNetworkinterface(dAtA, i, uint64(len(m.SmartNIC)))
		i += copy(dAtA[i:], m.SmartNIC)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNetworkinterface(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.OperStatus) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintNetworkinterface(dAtA, i, uint64(len(m.OperStatus)))
		i += copy(dAtA[i:], m.OperStatus)
	}
	if len(m.PrimaryMac) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintNetworkinterface(dAtA, i, uint64(len(m.PrimaryMac)))
		i += copy(dAtA[i:], m.PrimaryMac)
	}
	if m.IFHostStatus != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintNetworkinterface(dAtA, i, uint64(m.IFHostStatus.Size()))
		n5, err := m.IFHostStatus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.IFUplinkStatus != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintNetworkinterface(dAtA, i, uint64(m.IFUplinkStatus.Size()))
		n6, err := m.IFUplinkStatus.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n6
	}
	return i, nil
}

func (m *NetworkInterfaceUplinkStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NetworkInterfaceUplinkStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.LinkSpeed) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintNetworkinterface(dAtA, i, uint64(len(m.LinkSpeed)))
		i += copy(dAtA[i:], m.LinkSpeed)
	}
	return i, nil
}

func encodeVarintNetworkinterface(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *NetworkInterface) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovNetworkinterface(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovNetworkinterface(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovNetworkinterface(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovNetworkinterface(uint64(l))
	return n
}

func (m *NetworkInterfaceHostStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.HostIfName)
	if l > 0 {
		n += 1 + l + sovNetworkinterface(uint64(l))
	}
	return n
}

func (m *NetworkInterfaceSpec) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *NetworkInterfaceStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.SmartNIC)
	if l > 0 {
		n += 1 + l + sovNetworkinterface(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovNetworkinterface(uint64(l))
	}
	l = len(m.OperStatus)
	if l > 0 {
		n += 1 + l + sovNetworkinterface(uint64(l))
	}
	l = len(m.PrimaryMac)
	if l > 0 {
		n += 1 + l + sovNetworkinterface(uint64(l))
	}
	if m.IFHostStatus != nil {
		l = m.IFHostStatus.Size()
		n += 1 + l + sovNetworkinterface(uint64(l))
	}
	if m.IFUplinkStatus != nil {
		l = m.IFUplinkStatus.Size()
		n += 1 + l + sovNetworkinterface(uint64(l))
	}
	return n
}

func (m *NetworkInterfaceUplinkStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.LinkSpeed)
	if l > 0 {
		n += 1 + l + sovNetworkinterface(uint64(l))
	}
	return n
}

func sovNetworkinterface(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozNetworkinterface(x uint64) (n int) {
	return sovNetworkinterface(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NetworkInterface) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkinterface
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
			return fmt.Errorf("proto: NetworkInterface: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkInterface: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
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
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
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
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
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
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
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
			skippy, err := skipNetworkinterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkinterface
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
func (m *NetworkInterfaceHostStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkinterface
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
			return fmt.Errorf("proto: NetworkInterfaceHostStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkInterfaceHostStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostIfName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostIfName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkinterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkinterface
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
func (m *NetworkInterfaceSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkinterface
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
			return fmt.Errorf("proto: NetworkInterfaceSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkInterfaceSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkinterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkinterface
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
func (m *NetworkInterfaceStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkinterface
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
			return fmt.Errorf("proto: NetworkInterfaceStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkInterfaceStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartNIC", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SmartNIC = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrimaryMac", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrimaryMac = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IFHostStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IFHostStatus == nil {
				m.IFHostStatus = &NetworkInterfaceHostStatus{}
			}
			if err := m.IFHostStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IFUplinkStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IFUplinkStatus == nil {
				m.IFUplinkStatus = &NetworkInterfaceUplinkStatus{}
			}
			if err := m.IFUplinkStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkinterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkinterface
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
func (m *NetworkInterfaceUplinkStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNetworkinterface
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
			return fmt.Errorf("proto: NetworkInterfaceUplinkStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NetworkInterfaceUplinkStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LinkSpeed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNetworkinterface
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
				return ErrInvalidLengthNetworkinterface
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LinkSpeed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNetworkinterface(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthNetworkinterface
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
func skipNetworkinterface(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNetworkinterface
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
					return 0, ErrIntOverflowNetworkinterface
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
					return 0, ErrIntOverflowNetworkinterface
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
				return 0, ErrInvalidLengthNetworkinterface
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowNetworkinterface
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
				next, err := skipNetworkinterface(dAtA[start:])
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
	ErrInvalidLengthNetworkinterface = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNetworkinterface   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("networkinterface.proto", fileDescriptorNetworkinterface) }

var fileDescriptorNetworkinterface = []byte{
	// 721 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdd, 0x4e, 0xdb, 0x48,
	0x18, 0xc5, 0x10, 0x02, 0x19, 0x58, 0xf0, 0x0e, 0x10, 0x9c, 0x28, 0x1b, 0xaf, 0xb2, 0x62, 0x15,
	0x24, 0x12, 0x4b, 0xbb, 0xd2, 0x5e, 0xac, 0x56, 0xab, 0xe2, 0x92, 0x94, 0xa8, 0xe4, 0x47, 0x24,
	0xa8, 0x12, 0xbd, 0x40, 0x13, 0x67, 0x12, 0xa6, 0xc4, 0xe3, 0xa9, 0x3d, 0x29, 0xca, 0x03, 0x94,
	0x87, 0xe1, 0x49, 0xb8, 0x44, 0xbd, 0xaf, 0x55, 0x71, 0x99, 0xa7, 0xa8, 0x66, 0xec, 0x08, 0x37,
	0x71, 0xe8, 0x9d, 0xcf, 0x99, 0xef, 0x3b, 0xe7, 0xfb, 0x93, 0x41, 0x9a, 0x62, 0x7e, 0xeb, 0xb8,
	0x37, 0x84, 0x72, 0xec, 0xf6, 0x91, 0x85, 0xcb, 0xcc, 0x75, 0xb8, 0x03, 0xd7, 0x42, 0x3e, 0x9b,
	0x1b, 0x38, 0xce, 0x60, 0x88, 0x0d, 0xc4, 0x88, 0x81, 0x28, 0x75, 0x38, 0xe2, 0xc4, 0xa1, 0x5e,
	0x10, 0x96, 0xad, 0x0c, 0x08, 0xbf, 0x1e, 0x75, 0xcb, 0x96, 0x63, 0x1b, 0x0c, 0x53, 0x0f, 0xd1,
	0x9e, 0x63, 0x78, 0xb7, 0xc6, 0x27, 0x4c, 0x89, 0x85, 0x8d, 0x11, 0x27, 0x43, 0x4f, 0xa4, 0x0e,
	0x30, 0x8d, 0x66, 0x1b, 0x84, 0x5a, 0xc3, 0x51, 0x0f, 0x4f, 0x65, 0x4a, 0x11, 0x99, 0x81, 0x33,
	0x70, 0x0c, 0x49, 0x77, 0x47, 0x7d, 0x89, 0x24, 0x90, 0x5f, 0x61, 0xf8, 0xc1, 0x02, 0x57, 0x51,
	0xa3, 0x8d, 0x39, 0x0a, 0xc2, 0x0a, 0x5f, 0x97, 0x81, 0xda, 0x08, 0xda, 0xa8, 0x4d, 0xdb, 0x83,
	0xff, 0x00, 0xa5, 0xa3, 0x29, 0xbf, 0x2b, 0xc5, 0x8d, 0xbf, 0x7e, 0x29, 0x23, 0x46, 0xca, 0x9d,
	0x31, 0xc3, 0x75, 0xcc, 0x91, 0xb9, 0xf3, 0xe0, 0xeb, 0x4b, 0x8f, 0xbe, 0xae, 0x4c, 0x7c, 0x7d,
	0xed, 0x88, 0xd0, 0x21, 0xa1, 0xf8, 0x7c, 0xfa, 0x01, 0xab, 0x40, 0x69, 0x6a, 0xcb, 0x32, 0x6f,
	0x5b, 0xe6, 0x35, 0xbb, 0x1f, 0xb0, 0xc5, 0x65, 0x66, 0x36, 0x92, 0xb9, 0x25, 0x0a, 0x38, 0x72,
	0x6c, 0xc2, 0xb1, 0xcd, 0xf8, 0xf8, 0x7c, 0x06, 0xc3, 0x36, 0x48, 0xb4, 0x19, 0xb6, 0xb4, 0x15,
	0x29, 0xf5, 0x5b, 0x39, 0x9c, 0x73, 0x79, 0xb6, 0x50, 0x11, 0x64, 0xa6, 0x85, 0xb0, 0x10, 0xf5,
	0x18, 0xb6, 0xa2, 0xa2, 0x3f, 0x62, 0xf8, 0x1e, 0x24, 0xdb, 0x1c, 0xf1, 0x91, 0xa7, 0x25, 0xa4,
	0xac, 0xbe, 0x58, 0x56, 0x86, 0x99, 0x5a, 0x28, 0xac, 0x7a, 0x12, 0x47, 0xa4, 0xe7, 0x98, 0x7f,
	0xf7, 0xbf, 0x7c, 0xce, 0xec, 0xc0, 0x5f, 0x67, 0x0f, 0xc5, 0x2b, 0x74, 0x41, 0x76, 0x56, 0xfe,
	0xd4, 0xf1, 0x78, 0x60, 0x01, 0x4f, 0x00, 0x10, 0xa8, 0xd6, 0x6f, 0x20, 0x1b, 0xcb, 0x89, 0xa7,
	0xcc, 0xcc, 0xc4, 0xd7, 0xf7, 0xae, 0x1d, 0x8f, 0x97, 0x48, 0x9f, 0x22, 0x1b, 0x47, 0x7c, 0xe3,
	0xe9, 0x42, 0x1a, 0xec, 0xc6, 0x4d, 0xa6, 0xf0, 0xb8, 0x0a, 0xd2, 0xf1, 0xbd, 0xc1, 0xff, 0xc1,
	0x7a, 0xdb, 0x46, 0x2e, 0x6f, 0xd4, 0x5e, 0x87, 0xb6, 0xfb, 0x13, 0x5f, 0xdf, 0xf1, 0x04, 0x57,
	0xa2, 0x24, 0x3a, 0xc7, 0x38, 0x12, 0x5e, 0x82, 0x84, 0xb8, 0x09, 0xb9, 0xec, 0x94, 0xf9, 0xdf,
	0xfd, 0x5d, 0xe6, 0xcf, 0x36, 0x77, 0x2b, 0x74, 0x64, 0x17, 0xe3, 0x1d, 0xcb, 0xb5, 0xaa, 0x48,
	0x38, 0x14, 0x8b, 0xe2, 0x63, 0x16, 0xed, 0x6a, 0x06, 0x43, 0x06, 0x40, 0x93, 0x61, 0x37, 0x5c,
	0xd6, 0x8a, 0x74, 0xa8, 0xdc, 0xdf, 0x65, 0x8a, 0x3f, 0x75, 0x08, 0x3e, 0x84, 0xc7, 0x9e, 0xc3,
	0xb0, 0x5b, 0x9a, 0x5b, 0x5c, 0x3c, 0x0d, 0x2f, 0x01, 0x68, 0xb9, 0xc4, 0x46, 0xee, 0xb8, 0x8e,
	0x2c, 0x79, 0x1e, 0x29, 0xd3, 0xb8, 0xbf, 0xcb, 0xc0, 0x8a, 0x78, 0x6e, 0xba, 0xc5, 0x3a, 0xb2,
	0x8e, 0x7b, 0x3d, 0xb7, 0x78, 0x28, 0xb5, 0x59, 0x10, 0x5b, 0xb2, 0x51, 0x74, 0x4e, 0xf1, 0x34,
	0xfc, 0x08, 0x36, 0x6b, 0xd5, 0xe7, 0x95, 0x6b, 0xab, 0xf2, 0xf8, 0xfe, 0x58, 0x78, 0x7c, 0xcf,
	0xa1, 0x66, 0x6e, 0xe2, 0xeb, 0x1a, 0xe9, 0x97, 0xe4, 0xd6, 0xe7, 0x7a, 0x59, 0xf8, 0x02, 0xc7,
	0x60, 0xab, 0x56, 0xbd, 0x60, 0x43, 0x42, 0x6f, 0x42, 0xd3, 0xa4, 0x34, 0x3d, 0x58, 0x68, 0x1a,
	0x0d, 0x36, 0xf3, 0x13, 0x5f, 0xcf, 0x92, 0x7e, 0x69, 0x24, 0xc9, 0x79, 0xe3, 0x17, 0xde, 0x0a,
	0x39, 0xb0, 0x3e, 0x5d, 0x03, 0x4c, 0x82, 0xe5, 0x8b, 0x96, 0xba, 0x04, 0xd7, 0x41, 0xe2, 0xa4,
	0xf9, 0xae, 0xa1, 0x2a, 0x85, 0x57, 0x20, 0x19, 0x9c, 0x81, 0xe0, 0x1a, 0xcd, 0x46, 0x45, 0x5d,
	0x82, 0x1b, 0x60, 0xed, 0xb4, 0xd9, 0xee, 0x5c, 0xb5, 0xaa, 0xaa, 0x02, 0xb7, 0x00, 0xb8, 0x68,
	0x9d, 0xd5, 0x1a, 0x6f, 0xaf, 0x2a, 0x9d, 0x53, 0x75, 0x05, 0x6e, 0x83, 0x8d, 0x10, 0xd7, 0xdf,
	0xd4, 0x3b, 0x6a, 0xa2, 0x80, 0x40, 0xee, 0xa5, 0xda, 0xe1, 0x31, 0x48, 0x9d, 0x09, 0xc4, 0x30,
	0xee, 0x85, 0xc7, 0xa9, 0x4d, 0x7c, 0x7d, 0x37, 0xa8, 0x56, 0xb0, 0x91, 0x46, 0x62, 0x59, 0x73,
	0xf3, 0xe1, 0x29, 0xaf, 0x3c, 0x3e, 0xe5, 0x95, 0x6f, 0x4f, 0x79, 0xa5, 0xa5, 0x74, 0x93, 0xf2,
	0x47, 0xf9, 0xf7, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x05, 0x45, 0xad, 0x06, 0x06, 0x00,
	0x00,
}
