// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: archive.proto

package monitoring

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pensando/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/pensando/sw/api"
import labels "github.com/pensando/sw/api/labels"
import fields "github.com/pensando/sw/api/fields"
import search "github.com/pensando/sw/api/generated/search"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

//
type ArchiveRequestSpec_LogType int32

const (
	//
	ArchiveRequestSpec_Event ArchiveRequestSpec_LogType = 0
	//
	ArchiveRequestSpec_AuditEvent ArchiveRequestSpec_LogType = 1
	//
	ArchiveRequestSpec_FwLog ArchiveRequestSpec_LogType = 2
)

var ArchiveRequestSpec_LogType_name = map[int32]string{
	0: "Event",
	1: "AuditEvent",
	2: "FwLog",
}
var ArchiveRequestSpec_LogType_value = map[string]int32{
	"Event":      0,
	"AuditEvent": 1,
	"FwLog":      2,
}

func (ArchiveRequestSpec_LogType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorArchive, []int{2, 0}
}

//
type ArchiveRequestStatus_ArchiveJobStatus int32

const (
	//
	ArchiveRequestStatus_Scheduled ArchiveRequestStatus_ArchiveJobStatus = 0
	//
	ArchiveRequestStatus_Running ArchiveRequestStatus_ArchiveJobStatus = 1
	//
	ArchiveRequestStatus_Completed ArchiveRequestStatus_ArchiveJobStatus = 2
	//
	ArchiveRequestStatus_Failed ArchiveRequestStatus_ArchiveJobStatus = 3
	//
	ArchiveRequestStatus_TimeOut ArchiveRequestStatus_ArchiveJobStatus = 4
	//
	ArchiveRequestStatus_Canceled ArchiveRequestStatus_ArchiveJobStatus = 5
)

var ArchiveRequestStatus_ArchiveJobStatus_name = map[int32]string{
	0: "Scheduled",
	1: "Running",
	2: "Completed",
	3: "Failed",
	4: "TimeOut",
	5: "Canceled",
}
var ArchiveRequestStatus_ArchiveJobStatus_value = map[string]int32{
	"Scheduled": 0,
	"Running":   1,
	"Completed": 2,
	"Failed":    3,
	"TimeOut":   4,
	"Canceled":  5,
}

func (ArchiveRequestStatus_ArchiveJobStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorArchive, []int{3, 0}
}

// ArchiveQuery is to archive audit logs, events and firewall logs based on time window, field values
type ArchiveQuery struct {
	// OR of Text-requirements to be matched, Exclude is not supported for Text search
	Texts []*search.TextRequirement `protobuf:"bytes,1,rep,name=Texts,json=texts,omitempty" json:"texts,omitempty"`
	// Field Selector is AND of field.Requirements
	Fields *fields.Selector `protobuf:"bytes,2,opt,name=Fields,json=fields,omitempty" json:"fields,omitempty"`
	// Label Selector is AND of label.Requirememts
	Labels *labels.Selector `protobuf:"bytes,3,opt,name=Labels,json=labels,omitempty" json:"labels,omitempty"`
	// StartTime selects all logs with timestamp greater than the StartTime, example 2018-10-18T00:12:00Z
	StartTime *api.Timestamp `protobuf:"bytes,4,opt,name=StartTime,json=start-time,omitempty" json:"start-time,omitempty"`
	// EndTime selects all logs with timestamp less than the EndTime, example 2018-09-18T00:12:00Z
	EndTime *api.Timestamp `protobuf:"bytes,5,opt,name=EndTime,json=end-time,omitempty" json:"end-time,omitempty"`
}

func (m *ArchiveQuery) Reset()                    { *m = ArchiveQuery{} }
func (m *ArchiveQuery) String() string            { return proto.CompactTextString(m) }
func (*ArchiveQuery) ProtoMessage()               {}
func (*ArchiveQuery) Descriptor() ([]byte, []int) { return fileDescriptorArchive, []int{0} }

func (m *ArchiveQuery) GetTexts() []*search.TextRequirement {
	if m != nil {
		return m.Texts
	}
	return nil
}

func (m *ArchiveQuery) GetFields() *fields.Selector {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *ArchiveQuery) GetLabels() *labels.Selector {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *ArchiveQuery) GetStartTime() *api.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *ArchiveQuery) GetEndTime() *api.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

// ArchiveRequest is to asynchronously archive audit logs, events and firewall logs
type ArchiveRequest struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,json=,inline,embedded=T" json:",inline"`
	//
	api.ObjectMeta `protobuf:"bytes,2,opt,name=O,json=meta,omitempty,embedded=O" json:"meta,omitempty"`
	//
	Spec ArchiveRequestSpec `protobuf:"bytes,3,opt,name=Spec,json=spec,omitempty" json:"spec,omitempty"`
	//
	Status ArchiveRequestStatus `protobuf:"bytes,4,opt,name=Status,json=status,omitempty" json:"status,omitempty"`
}

func (m *ArchiveRequest) Reset()                    { *m = ArchiveRequest{} }
func (m *ArchiveRequest) String() string            { return proto.CompactTextString(m) }
func (*ArchiveRequest) ProtoMessage()               {}
func (*ArchiveRequest) Descriptor() ([]byte, []int) { return fileDescriptorArchive, []int{1} }

func (m *ArchiveRequest) GetSpec() ArchiveRequestSpec {
	if m != nil {
		return m.Spec
	}
	return ArchiveRequestSpec{}
}

func (m *ArchiveRequest) GetStatus() ArchiveRequestStatus {
	if m != nil {
		return m.Status
	}
	return ArchiveRequestStatus{}
}

//
type ArchiveRequestSpec struct {
	//
	Type string `protobuf:"bytes,1,opt,name=Type,json=type,proto3" json:"type"`
	//
	Query *ArchiveQuery `protobuf:"bytes,2,opt,name=Query,json=query" json:"query"`
}

func (m *ArchiveRequestSpec) Reset()                    { *m = ArchiveRequestSpec{} }
func (m *ArchiveRequestSpec) String() string            { return proto.CompactTextString(m) }
func (*ArchiveRequestSpec) ProtoMessage()               {}
func (*ArchiveRequestSpec) Descriptor() ([]byte, []int) { return fileDescriptorArchive, []int{2} }

func (m *ArchiveRequestSpec) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ArchiveRequestSpec) GetQuery() *ArchiveQuery {
	if m != nil {
		return m.Query
	}
	return nil
}

//
type ArchiveRequestStatus struct {
	//
	Status string `protobuf:"bytes,1,opt,name=Status,json=status,omitempty,proto3" json:"status,omitempty"`
	//
	Reason string `protobuf:"bytes,2,opt,name=Reason,json=reason,omitempty,proto3" json:"reason,omitempty"`
	//
	URI string `protobuf:"bytes,3,opt,name=URI,json=uri,omitempty,proto3" json:"uri,omitempty"`
}

func (m *ArchiveRequestStatus) Reset()                    { *m = ArchiveRequestStatus{} }
func (m *ArchiveRequestStatus) String() string            { return proto.CompactTextString(m) }
func (*ArchiveRequestStatus) ProtoMessage()               {}
func (*ArchiveRequestStatus) Descriptor() ([]byte, []int) { return fileDescriptorArchive, []int{3} }

func (m *ArchiveRequestStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *ArchiveRequestStatus) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *ArchiveRequestStatus) GetURI() string {
	if m != nil {
		return m.URI
	}
	return ""
}

// CancelArchiveRequest is to cancel archive request that is in scheduled or running state
type CancelArchiveRequest struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,json=,inline,embedded=T" json:",inline"`
	//
	api.ObjectMeta `protobuf:"bytes,2,opt,name=O,json=meta,omitempty,embedded=O" json:"meta,omitempty"`
}

func (m *CancelArchiveRequest) Reset()                    { *m = CancelArchiveRequest{} }
func (m *CancelArchiveRequest) String() string            { return proto.CompactTextString(m) }
func (*CancelArchiveRequest) ProtoMessage()               {}
func (*CancelArchiveRequest) Descriptor() ([]byte, []int) { return fileDescriptorArchive, []int{4} }

func init() {
	proto.RegisterType((*ArchiveQuery)(nil), "monitoring.ArchiveQuery")
	proto.RegisterType((*ArchiveRequest)(nil), "monitoring.ArchiveRequest")
	proto.RegisterType((*ArchiveRequestSpec)(nil), "monitoring.ArchiveRequestSpec")
	proto.RegisterType((*ArchiveRequestStatus)(nil), "monitoring.ArchiveRequestStatus")
	proto.RegisterType((*CancelArchiveRequest)(nil), "monitoring.CancelArchiveRequest")
	proto.RegisterEnum("monitoring.ArchiveRequestSpec_LogType", ArchiveRequestSpec_LogType_name, ArchiveRequestSpec_LogType_value)
	proto.RegisterEnum("monitoring.ArchiveRequestStatus_ArchiveJobStatus", ArchiveRequestStatus_ArchiveJobStatus_name, ArchiveRequestStatus_ArchiveJobStatus_value)
}
func (m *ArchiveQuery) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArchiveQuery) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Texts) > 0 {
		for _, msg := range m.Texts {
			dAtA[i] = 0xa
			i++
			i = encodeVarintArchive(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.Fields != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintArchive(dAtA, i, uint64(m.Fields.Size()))
		n1, err := m.Fields.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Labels != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintArchive(dAtA, i, uint64(m.Labels.Size()))
		n2, err := m.Labels.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.StartTime != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintArchive(dAtA, i, uint64(m.StartTime.Size()))
		n3, err := m.StartTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.EndTime != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintArchive(dAtA, i, uint64(m.EndTime.Size()))
		n4, err := m.EndTime.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	return i, nil
}

func (m *ArchiveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArchiveRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintArchive(dAtA, i, uint64(m.TypeMeta.Size()))
	n5, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x12
	i++
	i = encodeVarintArchive(dAtA, i, uint64(m.ObjectMeta.Size()))
	n6, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x1a
	i++
	i = encodeVarintArchive(dAtA, i, uint64(m.Spec.Size()))
	n7, err := m.Spec.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	dAtA[i] = 0x22
	i++
	i = encodeVarintArchive(dAtA, i, uint64(m.Status.Size()))
	n8, err := m.Status.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n8
	return i, nil
}

func (m *ArchiveRequestSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArchiveRequestSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Type) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintArchive(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if m.Query != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintArchive(dAtA, i, uint64(m.Query.Size()))
		n9, err := m.Query.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n9
	}
	return i, nil
}

func (m *ArchiveRequestStatus) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ArchiveRequestStatus) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Status) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintArchive(dAtA, i, uint64(len(m.Status)))
		i += copy(dAtA[i:], m.Status)
	}
	if len(m.Reason) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintArchive(dAtA, i, uint64(len(m.Reason)))
		i += copy(dAtA[i:], m.Reason)
	}
	if len(m.URI) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintArchive(dAtA, i, uint64(len(m.URI)))
		i += copy(dAtA[i:], m.URI)
	}
	return i, nil
}

func (m *CancelArchiveRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CancelArchiveRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintArchive(dAtA, i, uint64(m.TypeMeta.Size()))
	n10, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n10
	dAtA[i] = 0x12
	i++
	i = encodeVarintArchive(dAtA, i, uint64(m.ObjectMeta.Size()))
	n11, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n11
	return i, nil
}

func encodeVarintArchive(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ArchiveQuery) Size() (n int) {
	var l int
	_ = l
	if len(m.Texts) > 0 {
		for _, e := range m.Texts {
			l = e.Size()
			n += 1 + l + sovArchive(uint64(l))
		}
	}
	if m.Fields != nil {
		l = m.Fields.Size()
		n += 1 + l + sovArchive(uint64(l))
	}
	if m.Labels != nil {
		l = m.Labels.Size()
		n += 1 + l + sovArchive(uint64(l))
	}
	if m.StartTime != nil {
		l = m.StartTime.Size()
		n += 1 + l + sovArchive(uint64(l))
	}
	if m.EndTime != nil {
		l = m.EndTime.Size()
		n += 1 + l + sovArchive(uint64(l))
	}
	return n
}

func (m *ArchiveRequest) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovArchive(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovArchive(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovArchive(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovArchive(uint64(l))
	return n
}

func (m *ArchiveRequestSpec) Size() (n int) {
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovArchive(uint64(l))
	}
	if m.Query != nil {
		l = m.Query.Size()
		n += 1 + l + sovArchive(uint64(l))
	}
	return n
}

func (m *ArchiveRequestStatus) Size() (n int) {
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovArchive(uint64(l))
	}
	l = len(m.Reason)
	if l > 0 {
		n += 1 + l + sovArchive(uint64(l))
	}
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovArchive(uint64(l))
	}
	return n
}

func (m *CancelArchiveRequest) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovArchive(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovArchive(uint64(l))
	return n
}

func sovArchive(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozArchive(x uint64) (n int) {
	return sovArchive(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ArchiveQuery) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArchive
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
			return fmt.Errorf("proto: ArchiveQuery: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArchiveQuery: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Texts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Texts = append(m.Texts, &search.TextRequirement{})
			if err := m.Texts[len(m.Texts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fields", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Fields == nil {
				m.Fields = &fields.Selector{}
			}
			if err := m.Fields.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Labels", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Labels == nil {
				m.Labels = &labels.Selector{}
			}
			if err := m.Labels.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartTime == nil {
				m.StartTime = &api.Timestamp{}
			}
			if err := m.StartTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndTime == nil {
				m.EndTime = &api.Timestamp{}
			}
			if err := m.EndTime.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArchive
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
func (m *ArchiveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArchive
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
			return fmt.Errorf("proto: ArchiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArchiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
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
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
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
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
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
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
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
			skippy, err := skipArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArchive
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
func (m *ArchiveRequestSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArchive
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
			return fmt.Errorf("proto: ArchiveRequestSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArchiveRequestSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Query", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Query == nil {
				m.Query = &ArchiveQuery{}
			}
			if err := m.Query.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArchive
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
func (m *ArchiveRequestStatus) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArchive
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
			return fmt.Errorf("proto: ArchiveRequestStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ArchiveRequestStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Reason = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArchive
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
func (m *CancelArchiveRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowArchive
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
			return fmt.Errorf("proto: CancelArchiveRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CancelArchiveRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
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
					return ErrIntOverflowArchive
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
				return ErrInvalidLengthArchive
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipArchive(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthArchive
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
func skipArchive(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowArchive
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
					return 0, ErrIntOverflowArchive
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
					return 0, ErrIntOverflowArchive
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
				return 0, ErrInvalidLengthArchive
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowArchive
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
				next, err := skipArchive(dAtA[start:])
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
	ErrInvalidLengthArchive = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowArchive   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("archive.proto", fileDescriptorArchive) }

var fileDescriptorArchive = []byte{
	// 810 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x94, 0x4d, 0x6f, 0xdc, 0x44,
	0x18, 0xc7, 0xe3, 0x7d, 0x0b, 0x3b, 0x79, 0xa9, 0x99, 0x46, 0xc5, 0x89, 0xd0, 0x3a, 0xda, 0x0a,
	0x29, 0x95, 0x1a, 0x1b, 0x15, 0xa9, 0x12, 0xdc, 0xea, 0x28, 0x2b, 0x5e, 0x82, 0x42, 0xbd, 0xcb,
	0x8d, 0xcb, 0xac, 0xfd, 0xd4, 0x3b, 0xc8, 0x9e, 0x71, 0x3d, 0xe3, 0x94, 0x15, 0xe2, 0x48, 0xcf,
	0x7c, 0x8e, 0x7e, 0x05, 0xce, 0x48, 0xbd, 0x20, 0x55, 0x7c, 0x00, 0x0b, 0xed, 0x09, 0xfc, 0x29,
	0xd0, 0x8c, 0x1d, 0x70, 0xbc, 0xdb, 0x2a, 0x47, 0x2e, 0x89, 0x9f, 0xff, 0xfc, 0x9f, 0xdf, 0x3e,
	0xf3, 0x1f, 0x8f, 0xd1, 0x1e, 0xc9, 0x82, 0x05, 0xbd, 0x02, 0x27, 0xcd, 0xb8, 0xe4, 0x18, 0x25,
	0x9c, 0x51, 0xc9, 0x33, 0xca, 0xa2, 0xa3, 0x0f, 0x23, 0xce, 0xa3, 0x18, 0x5c, 0x92, 0x52, 0x97,
	0x30, 0xc6, 0x25, 0x91, 0x94, 0x33, 0x51, 0x39, 0x8f, 0xce, 0x23, 0x2a, 0x17, 0xf9, 0xdc, 0x09,
	0x78, 0xe2, 0xa6, 0xc0, 0x04, 0x61, 0x21, 0x77, 0xc5, 0x0b, 0xf7, 0x0a, 0x18, 0x0d, 0xc0, 0xcd,
	0x25, 0x8d, 0x85, 0x6a, 0x8d, 0x80, 0x35, 0xbb, 0x5d, 0xca, 0x82, 0x38, 0x0f, 0xe1, 0x1a, 0x73,
	0xda, 0xc0, 0x44, 0x3c, 0xe2, 0xae, 0x96, 0xe7, 0xf9, 0x33, 0x5d, 0xe9, 0x42, 0x3f, 0xd5, 0xf6,
	0x8f, 0xde, 0xf2, 0xab, 0x6a, 0xc6, 0x04, 0x24, 0xa9, 0x6d, 0x1f, 0xbf, 0xc3, 0x16, 0x93, 0x39,
	0xc4, 0xc2, 0x15, 0x10, 0x43, 0x20, 0x79, 0x76, 0x8b, 0x8e, 0x67, 0x14, 0xe2, 0x70, 0xad, 0x63,
	0x57, 0x80, 0xca, 0xae, 0xaa, 0xc6, 0xbf, 0x74, 0xd1, 0xee, 0x93, 0x2a, 0xca, 0xa7, 0x39, 0x64,
	0x4b, 0x7c, 0x81, 0xfa, 0x33, 0xf8, 0x41, 0x0a, 0xcb, 0x38, 0xee, 0x9e, 0xec, 0x3c, 0xfa, 0xc0,
	0xa9, 0xed, 0x4a, 0xf4, 0xe1, 0x79, 0x4e, 0x33, 0x48, 0x80, 0x49, 0xef, 0x6e, 0x59, 0xd8, 0x77,
	0xa4, 0x72, 0x3e, 0xe4, 0x09, 0x95, 0x90, 0xa4, 0x72, 0xe9, 0xb7, 0x05, 0xfc, 0x39, 0x1a, 0x4c,
	0xf4, 0x14, 0x56, 0xe7, 0xd8, 0x38, 0xd9, 0x79, 0x64, 0x3a, 0xd5, 0x50, 0xce, 0xb4, 0x1e, 0xca,
	0x3b, 0x28, 0x0b, 0xdb, 0xac, 0xc4, 0x06, 0x68, 0x4d, 0x51, 0xa4, 0x0b, 0x9d, 0x80, 0xd5, 0xad,
	0x49, 0x55, 0x20, 0x2d, 0x52, 0x25, 0x36, 0x49, 0x6d, 0x05, 0x3f, 0x45, 0xc3, 0xa9, 0x24, 0x99,
	0x9c, 0xd1, 0x04, 0xac, 0x9e, 0x86, 0xed, 0x3b, 0x24, 0xa5, 0x8e, 0x12, 0x84, 0x24, 0x49, 0xea,
	0x59, 0x65, 0x61, 0x1f, 0x08, 0x65, 0x3a, 0x95, 0x34, 0x81, 0x06, 0x6e, 0xa3, 0x8a, 0xbf, 0x42,
	0xdb, 0xe7, 0x2c, 0xd4, 0xc0, 0xfe, 0x46, 0xe0, 0xbd, 0xb2, 0xb0, 0x31, 0xb0, 0xb0, 0x8d, 0xdb,
	0xa0, 0x8d, 0xff, 0xee, 0xa0, 0xfd, 0xfa, 0x48, 0x54, 0xe0, 0x20, 0x24, 0x7e, 0x8c, 0x8c, 0x99,
	0x65, 0x68, 0xf2, 0x5e, 0x45, 0x5e, 0xa6, 0xf0, 0x35, 0x48, 0xe2, 0xdd, 0x7d, 0x5d, 0xd8, 0x5b,
	0x6f, 0x0a, 0xdb, 0x28, 0x0b, 0x7b, 0xfb, 0x21, 0x65, 0x31, 0x65, 0xe0, 0x5f, 0x3f, 0xe0, 0x09,
	0x32, 0x2e, 0xeb, 0xe4, 0xef, 0xe8, 0xbe, 0xcb, 0xf9, 0xf7, 0x10, 0x48, 0xdd, 0x79, 0xd4, 0xe8,
	0xdc, 0x57, 0xef, 0x60, 0x63, 0xac, 0x56, 0x8d, 0x67, 0xa8, 0x37, 0x4d, 0x21, 0xa8, 0xa3, 0x1f,
	0x39, 0xff, 0xdd, 0x36, 0xe7, 0xe6, 0xa4, 0xca, 0xe5, 0xdd, 0x53, 0x64, 0x45, 0x15, 0x29, 0x04,
	0x4d, 0xea, 0xcd, 0x1a, 0x7f, 0x87, 0x06, 0x53, 0x49, 0x64, 0x2e, 0xea, 0x53, 0x38, 0x7e, 0x07,
	0x57, 0xfb, 0x3c, 0xab, 0x26, 0x9b, 0x42, 0xd7, 0xcd, 0x63, 0x6e, 0x2b, 0x9f, 0xdd, 0xff, 0xe3,
	0xe7, 0x43, 0x1b, 0xed, 0xb8, 0x3f, 0x5e, 0x3a, 0x33, 0x60, 0x84, 0xc9, 0x9f, 0xb0, 0x59, 0x7f,
	0x34, 0x4e, 0xb3, 0x8a, 0x2a, 0xc6, 0xbf, 0x19, 0x08, 0xaf, 0xef, 0x00, 0x9f, 0xa1, 0x9e, 0x4a,
	0x58, 0x47, 0x3e, 0xf4, 0xdc, 0x57, 0x2f, 0x0f, 0xef, 0x4f, 0x65, 0x76, 0xce, 0xf2, 0xe4, 0x64,
	0xdd, 0xed, 0x5c, 0xf0, 0x48, 0xb9, 0x1f, 0x94, 0x85, 0xdd, 0x93, 0xcb, 0x14, 0x7c, 0xfd, 0x17,
	0x7f, 0x8a, 0xfa, 0xfa, 0x4a, 0xd5, 0x07, 0x60, 0x6d, 0xd8, 0x9d, 0x5e, 0xf7, 0x86, 0x65, 0x61,
	0xf7, 0x9f, 0xab, 0x47, 0xbf, 0xfa, 0x37, 0x76, 0xd1, 0x76, 0x0d, 0xc5, 0x43, 0xd4, 0x3f, 0xbf,
	0x02, 0x26, 0xcd, 0x2d, 0xbc, 0x8f, 0xd0, 0x93, 0x3c, 0xa4, 0xb2, 0xaa, 0x0d, 0xb5, 0x34, 0x79,
	0x71, 0xc1, 0x23, 0xb3, 0x33, 0xfe, 0xbd, 0x83, 0x0e, 0x36, 0x25, 0x86, 0x17, 0xff, 0x66, 0x5c,
	0xed, 0x65, 0xf2, 0xea, 0xe5, 0xa1, 0xf3, 0x96, 0xbd, 0x68, 0xdf, 0xf5, 0x68, 0x5f, 0xf2, 0x79,
	0x25, 0x3c, 0xb8, 0x5d, 0xde, 0xf8, 0x31, 0x1a, 0xf8, 0x40, 0x04, 0x67, 0x7a, 0xbf, 0xc3, 0xea,
	0x3a, 0x66, 0x5a, 0x69, 0xf6, 0xb5, 0x15, 0x7c, 0x8a, 0xba, 0xdf, 0xfa, 0x5f, 0xe8, 0x57, 0x6b,
	0xe8, 0xbd, 0x5f, 0x16, 0xf6, 0x5e, 0x9e, 0xd1, 0x46, 0xc7, 0xcd, 0x72, 0x1c, 0x22, 0xb3, 0x3d,
	0x21, 0xde, 0x43, 0xc3, 0x69, 0xb0, 0x80, 0x30, 0x8f, 0x21, 0x34, 0xb7, 0xf0, 0x0e, 0xda, 0xf6,
	0x73, 0xc6, 0x28, 0x8b, 0x4c, 0x43, 0xad, 0x9d, 0xf1, 0x24, 0x8d, 0x41, 0x42, 0x68, 0x76, 0x30,
	0x42, 0x83, 0x09, 0xa1, 0xca, 0xd7, 0x55, 0x3e, 0x75, 0x43, 0x2f, 0x73, 0x69, 0xf6, 0xf0, 0x2e,
	0x7a, 0xef, 0x8c, 0xb0, 0x00, 0xd4, 0x52, 0x7f, 0xfc, 0xab, 0x81, 0x0e, 0xaa, 0xf2, 0xff, 0x75,
	0x13, 0x6f, 0xf5, 0x56, 0x7b, 0xe6, 0xeb, 0xd5, 0xc8, 0x78, 0xb3, 0x1a, 0x19, 0x7f, 0xae, 0x46,
	0xc6, 0x5f, 0xab, 0xd1, 0xd6, 0x37, 0xc6, 0x7c, 0xa0, 0xbf, 0xf7, 0x9f, 0xfc, 0x13, 0x00, 0x00,
	0xff, 0xff, 0xb8, 0x3e, 0xb0, 0xfb, 0x39, 0x07, 0x00, 0x00,
}
