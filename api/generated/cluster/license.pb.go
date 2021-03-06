// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: license.proto

package cluster

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

// Feature represents each individual feature on the system.
type Feature struct {
	//
	FeatureKey string `protobuf:"bytes,1,opt,name=FeatureKey,json=feature-key,omitempty,proto3" json:"feature-key,omitempty"`
	//
	License string `protobuf:"bytes,2,opt,name=License,json=licence,omitempty,proto3" json:"licence,omitempty"`
}

func (m *Feature) Reset()                    { *m = Feature{} }
func (m *Feature) String() string            { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()               {}
func (*Feature) Descriptor() ([]byte, []int) { return fileDescriptorLicense, []int{0} }

func (m *Feature) GetFeatureKey() string {
	if m != nil {
		return m.FeatureKey
	}
	return ""
}

func (m *Feature) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

// FeatureStatus is the current operational status of a feature
type FeatureStatus struct {
	//
	FeatureKey string `protobuf:"bytes,1,opt,name=FeatureKey,json=feature-key,omitempty,proto3" json:"feature-key,omitempty"`
	//
	Value string `protobuf:"bytes,2,opt,name=Value,json=value,omitempty,proto3" json:"value,omitempty"`
	//
	Expiry string `protobuf:"bytes,3,opt,name=Expiry,json=expiry,omitempty,proto3" json:"expiry,omitempty"`
}

func (m *FeatureStatus) Reset()                    { *m = FeatureStatus{} }
func (m *FeatureStatus) String() string            { return proto.CompactTextString(m) }
func (*FeatureStatus) ProtoMessage()               {}
func (*FeatureStatus) Descriptor() ([]byte, []int) { return fileDescriptorLicense, []int{1} }

func (m *FeatureStatus) GetFeatureKey() string {
	if m != nil {
		return m.FeatureKey
	}
	return ""
}

func (m *FeatureStatus) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *FeatureStatus) GetExpiry() string {
	if m != nil {
		return m.Expiry
	}
	return ""
}

//
type License struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,json=,inline,embedded=T" json:",inline"`
	//
	api.ObjectMeta `protobuf:"bytes,2,opt,name=O,json=meta,omitempty,embedded=O" json:"meta,omitempty"`
	//
	Spec LicenseSpec `protobuf:"bytes,3,opt,name=Spec,json=spec,omitempty" json:"spec,omitempty"`
	//
	Status LicenseStatus `protobuf:"bytes,4,opt,name=Status,json=status,omitempty" json:"status,omitempty"`
}

func (m *License) Reset()                    { *m = License{} }
func (m *License) String() string            { return proto.CompactTextString(m) }
func (*License) ProtoMessage()               {}
func (*License) Descriptor() ([]byte, []int) { return fileDescriptorLicense, []int{2} }

func (m *License) GetSpec() LicenseSpec {
	if m != nil {
		return m.Spec
	}
	return LicenseSpec{}
}

func (m *License) GetStatus() LicenseStatus {
	if m != nil {
		return m.Status
	}
	return LicenseStatus{}
}

//
type LicenseSpec struct {
	// List of Feature licences applied
	Features []Feature `protobuf:"bytes,1,rep,name=Features,json=features,omitempty" json:"features,omitempty"`
}

func (m *LicenseSpec) Reset()                    { *m = LicenseSpec{} }
func (m *LicenseSpec) String() string            { return proto.CompactTextString(m) }
func (*LicenseSpec) ProtoMessage()               {}
func (*LicenseSpec) Descriptor() ([]byte, []int) { return fileDescriptorLicense, []int{3} }

func (m *LicenseSpec) GetFeatures() []Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

//
type LicenseStatus struct {
	// Status of current Licenced features
	Features []FeatureStatus `protobuf:"bytes,1,rep,name=Features,json=features,omitempty" json:"features,omitempty"`
	// Licenses that are not understood by the current running version of software.
	Unknown []string `protobuf:"bytes,2,rep,name=Unknown,json=unknown,omitempty" json:"unknown,omitempty"`
}

func (m *LicenseStatus) Reset()                    { *m = LicenseStatus{} }
func (m *LicenseStatus) String() string            { return proto.CompactTextString(m) }
func (*LicenseStatus) ProtoMessage()               {}
func (*LicenseStatus) Descriptor() ([]byte, []int) { return fileDescriptorLicense, []int{4} }

func (m *LicenseStatus) GetFeatures() []FeatureStatus {
	if m != nil {
		return m.Features
	}
	return nil
}

func (m *LicenseStatus) GetUnknown() []string {
	if m != nil {
		return m.Unknown
	}
	return nil
}

func init() {
	proto.RegisterType((*Feature)(nil), "cluster.Feature")
	proto.RegisterType((*FeatureStatus)(nil), "cluster.FeatureStatus")
	proto.RegisterType((*License)(nil), "cluster.License")
	proto.RegisterType((*LicenseSpec)(nil), "cluster.LicenseSpec")
	proto.RegisterType((*LicenseStatus)(nil), "cluster.LicenseStatus")
}
func (m *Feature) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Feature) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FeatureKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLicense(dAtA, i, uint64(len(m.FeatureKey)))
		i += copy(dAtA[i:], m.FeatureKey)
	}
	if len(m.License) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLicense(dAtA, i, uint64(len(m.License)))
		i += copy(dAtA[i:], m.License)
	}
	return i, nil
}

func (m *FeatureStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeatureStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.FeatureKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintLicense(dAtA, i, uint64(len(m.FeatureKey)))
		i += copy(dAtA[i:], m.FeatureKey)
	}
	if len(m.Value) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintLicense(dAtA, i, uint64(len(m.Value)))
		i += copy(dAtA[i:], m.Value)
	}
	if len(m.Expiry) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintLicense(dAtA, i, uint64(len(m.Expiry)))
		i += copy(dAtA[i:], m.Expiry)
	}
	return i, nil
}

func (m *License) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *License) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintLicense(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintLicense(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintLicense(dAtA, i, uint64(m.Spec.Size()))
	n3, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x22
	i++
	i = encodeVarintLicense(dAtA, i, uint64(m.Status.Size()))
	n4, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func (m *LicenseSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LicenseSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Features) > 0 {
		for _, msg := range m.Features {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLicense(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *LicenseStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LicenseStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Features) > 0 {
		for _, msg := range m.Features {
			dAtA[i] = 0xa
			i++
			i = encodeVarintLicense(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Unknown) > 0 {
		for _, s := range m.Unknown {
			dAtA[i] = 0x12
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
	return i, nil
}

func encodeVarintLicense(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Feature) Size() (n int) {
	var l int
	_ = l
	l = len(m.FeatureKey)
	if l > 0 {
		n += 1 + l + sovLicense(uint64(l))
	}
	l = len(m.License)
	if l > 0 {
		n += 1 + l + sovLicense(uint64(l))
	}
	return n
}

func (m *FeatureStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.FeatureKey)
	if l > 0 {
		n += 1 + l + sovLicense(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovLicense(uint64(l))
	}
	l = len(m.Expiry)
	if l > 0 {
		n += 1 + l + sovLicense(uint64(l))
	}
	return n
}

func (m *License) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovLicense(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovLicense(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovLicense(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovLicense(uint64(l))
	return n
}

func (m *LicenseSpec) Size() (n int) {
	var l int
	_ = l
	if len(m.Features) > 0 {
		for _, e := range m.Features {
			l = e.Size()
			n += 1 + l + sovLicense(uint64(l))
		}
	}
	return n
}

func (m *LicenseStatus) Size() (n int) {
	var l int
	_ = l
	if len(m.Features) > 0 {
		for _, e := range m.Features {
			l = e.Size()
			n += 1 + l + sovLicense(uint64(l))
		}
	}
	if len(m.Unknown) > 0 {
		for _, s := range m.Unknown {
			l = len(s)
			n += 1 + l + sovLicense(uint64(l))
		}
	}
	return n
}

func sovLicense(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozLicense(x uint64) (n int) {
	return sovLicense(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Feature) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLicense
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
			return fmt.Errorf("proto: Feature: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Feature: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeatureKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeatureKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field License", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.License = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLicense(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLicense
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
func (m *FeatureStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLicense
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
			return fmt.Errorf("proto: FeatureStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeatureStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeatureKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FeatureKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Expiry = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLicense(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLicense
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
func (m *License) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLicense
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
			return fmt.Errorf("proto: License: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: License: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
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
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
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
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
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
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
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
			skippy, err := skipLicense(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLicense
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
func (m *LicenseSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLicense
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
			return fmt.Errorf("proto: LicenseSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LicenseSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Features", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Features = append(m.Features, Feature{})
			if err := m.Features[len(m.Features)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLicense(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLicense
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
func (m *LicenseStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLicense
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
			return fmt.Errorf("proto: LicenseStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LicenseStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Features", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Features = append(m.Features, FeatureStatus{})
			if err := m.Features[len(m.Features)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Unknown", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLicense
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
				return ErrInvalidLengthLicense
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Unknown = append(m.Unknown, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLicense(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthLicense
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
func skipLicense(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLicense
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
					return 0, ErrIntOverflowLicense
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
					return 0, ErrIntOverflowLicense
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
				return 0, ErrInvalidLengthLicense
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowLicense
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
				next, err := skipLicense(dAtA[start:])
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
	ErrInvalidLengthLicense = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLicense   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("license.proto", fileDescriptorLicense) }

var fileDescriptorLicense = []byte{
	// 544 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcb, 0x4e, 0xdb, 0x4c,
	0x14, 0x66, 0x12, 0xfe, 0x84, 0x7f, 0xa2, 0x40, 0x3a, 0x40, 0x64, 0xa2, 0x2a, 0x8e, 0x22, 0x55,
	0x62, 0x01, 0xb6, 0x14, 0x24, 0xa4, 0x76, 0x19, 0x15, 0x16, 0xbd, 0x88, 0x2a, 0xd0, 0xaa, 0xdb,
	0x89, 0x39, 0xb8, 0x53, 0x9c, 0x19, 0x2b, 0x33, 0x86, 0xe6, 0x01, 0xba, 0xe9, 0x73, 0xf4, 0x25,
	0xba, 0xed, 0x8a, 0x25, 0xea, 0x03, 0x58, 0x55, 0x56, 0x95, 0x9f, 0xa2, 0xf2, 0x78, 0x82, 0x86,
	0x98, 0x2e, 0x2a, 0x75, 0x37, 0xe7, 0x9b, 0xf9, 0x2e, 0x3e, 0xc7, 0x07, 0x37, 0x23, 0x16, 0x00,
	0x97, 0xe0, 0xc5, 0x53, 0xa1, 0x04, 0xa9, 0x07, 0x51, 0x22, 0x15, 0x4c, 0x3b, 0x8f, 0x43, 0x21,
	0xc2, 0x08, 0x7c, 0x1a, 0x33, 0x9f, 0x72, 0x2e, 0x14, 0x55, 0x4c, 0x70, 0x59, 0x3c, 0xeb, 0x1c,
	0x85, 0x4c, 0x7d, 0x48, 0xc6, 0x5e, 0x20, 0x26, 0x7e, 0x0c, 0x5c, 0x52, 0x7e, 0x2e, 0x7c, 0x79,
	0xed, 0x5f, 0x01, 0x67, 0x01, 0xf8, 0x89, 0x62, 0x91, 0xcc, 0xa9, 0x21, 0x70, 0x9b, 0xed, 0x33,
	0x1e, 0x44, 0xc9, 0x39, 0x2c, 0x64, 0xf6, 0x2d, 0x99, 0x50, 0x84, 0xc2, 0xd7, 0xf0, 0x38, 0xb9,
	0xd0, 0x95, 0x2e, 0xf4, 0xc9, 0x3c, 0x7f, 0xf2, 0x07, 0xd7, 0x3c, 0xe3, 0x04, 0x14, 0x2d, 0x9e,
	0xf5, 0xbf, 0x20, 0x5c, 0x3f, 0x06, 0xaa, 0x92, 0x29, 0x90, 0xe7, 0x18, 0x9b, 0xe3, 0x4b, 0x98,
	0x39, 0xa8, 0x87, 0x76, 0xff, 0x1f, 0xee, 0x64, 0xa9, 0xbb, 0x7d, 0x51, 0xa0, 0xfb, 0x97, 0x30,
	0xdb, 0x13, 0x13, 0xa6, 0x60, 0x12, 0xab, 0xd9, 0xe8, 0x61, 0x98, 0x3c, 0xc5, 0xf5, 0x57, 0x45,
	0x9b, 0x9c, 0x8a, 0x96, 0xd8, 0xce, 0x52, 0xf7, 0x91, 0xee, 0x5c, 0x00, 0x16, 0xbd, 0x0c, 0xf5,
	0xbf, 0x23, 0xdc, 0x34, 0x09, 0x4e, 0x15, 0x55, 0x89, 0xfc, 0x47, 0x91, 0x0e, 0xf0, 0x7f, 0xef,
	0x68, 0x94, 0x2c, 0x02, 0x6d, 0x66, 0xa9, 0xbb, 0x71, 0x95, 0x03, 0x16, 0x75, 0x19, 0x20, 0x87,
	0xb8, 0x76, 0xf4, 0x29, 0x66, 0xd3, 0x99, 0x53, 0xd5, 0xac, 0xad, 0x2c, 0x75, 0x5b, 0xa0, 0x11,
	0x8b, 0x56, 0x42, 0xfa, 0xdf, 0x2a, 0x77, 0x0d, 0x20, 0x87, 0x18, 0x9d, 0xe9, 0xd4, 0x8d, 0x41,
	0xd3, 0xa3, 0x31, 0xf3, 0xce, 0x66, 0x31, 0xbc, 0x06, 0x45, 0x87, 0x9b, 0x37, 0xa9, 0xbb, 0x72,
	0x9b, 0xba, 0x28, 0x4b, 0xdd, 0xfa, 0x1e, 0xe3, 0x11, 0xe3, 0x30, 0x5a, 0x1c, 0xc8, 0x31, 0x46,
	0x27, 0x3a, 0x6c, 0x63, 0xb0, 0xa1, 0x79, 0x27, 0xe3, 0x8f, 0x10, 0x28, 0xcd, 0xec, 0x58, 0xcc,
	0xf5, 0x7c, 0x92, 0x56, 0x9a, 0xa5, 0x9a, 0xbc, 0xc0, 0xab, 0xa7, 0x31, 0x04, 0xfa, 0x0b, 0x1a,
	0x83, 0x2d, 0xcf, 0xfc, 0xb0, 0x9e, 0xc9, 0x97, 0xdf, 0x0d, 0xdb, 0xb9, 0x5e, 0xae, 0x25, 0x63,
	0x08, 0x6c, 0xad, 0xfb, 0x35, 0x19, 0xe1, 0x5a, 0x31, 0x14, 0x67, 0x55, 0xab, 0xb5, 0x4b, 0x6a,
	0xfa, 0x76, 0xe8, 0x18, 0xbd, 0x96, 0xd4, 0xb5, 0xdd, 0xab, 0x65, 0xe4, 0xd9, 0xfa, 0x8f, 0xcf,
	0x3b, 0xb8, 0xb3, 0x66, 0xd6, 0x4a, 0xf6, 0x29, 0x6e, 0x58, 0xd1, 0xc8, 0x08, 0xaf, 0x99, 0xe9,
	0x4b, 0x07, 0xf5, 0xaa, 0xbb, 0x8d, 0x41, 0xeb, 0xce, 0xd4, 0x5c, 0x14, 0xed, 0xc8, 0x52, 0x97,
	0x98, 0xd1, 0xdb, 0x86, 0x0f, 0x60, 0xfd, 0xaf, 0x08, 0x37, 0xef, 0x05, 0x26, 0xef, 0x4b, 0x2e,
	0xed, 0x65, 0x17, 0xf3, 0x69, 0x7f, 0xe9, 0x95, 0xaf, 0xc2, 0x5b, 0x7e, 0xc9, 0xc5, 0x35, 0x77,
	0x2a, 0xbd, 0xea, 0x62, 0x15, 0x92, 0x02, 0xb2, 0x57, 0xa1, 0x04, 0x0d, 0x5b, 0x37, 0xf3, 0x2e,
	0xba, 0x9d, 0x77, 0xd1, 0xcf, 0x79, 0x17, 0xfd, 0x9a, 0x77, 0x57, 0xde, 0xa0, 0x71, 0x4d, 0xaf,
	0xec, 0xc1, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xbc, 0xf7, 0x73, 0xea, 0x87, 0x04, 0x00, 0x00,
}
