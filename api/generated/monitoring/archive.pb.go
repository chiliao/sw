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
	// ui-hint: Event
	ArchiveRequestSpec_Event ArchiveRequestSpec_LogType = 0
	// ui-hint: AuditEvent
	ArchiveRequestSpec_AuditEvent ArchiveRequestSpec_LogType = 1
)

var ArchiveRequestSpec_LogType_name = map[int32]string{
	0: "Event",
	1: "AuditEvent",
}
var ArchiveRequestSpec_LogType_value = map[string]int32{
	"Event":      0,
	"AuditEvent": 1,
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

// ArchiveQuery is to archive audit logs and events based on time window, field values
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
	// OR of tenants within the scope of which archive needs to be performed. If not specified, it will be set to tenant
	// of the logged in user. Also users in non default tenant can archive logs in their tenant scope only.
	Tenants []string `protobuf:"bytes,6,rep,name=Tenants,json=tenants,omitempty" json:"tenants,omitempty"`
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

func (m *ArchiveQuery) GetTenants() []string {
	if m != nil {
		return m.Tenants
	}
	return nil
}

// ArchiveRequest is to asynchronously archive audit logs and events
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
	if len(m.Tenants) > 0 {
		for _, s := range m.Tenants {
			dAtA[i] = 0x32
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
	if len(m.Tenants) > 0 {
		for _, s := range m.Tenants {
			l = len(s)
			n += 1 + l + sovArchive(uint64(l))
		}
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenants", wireType)
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
			m.Tenants = append(m.Tenants, string(dAtA[iNdEx:postIndex]))
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
	// 826 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x55, 0xdd, 0x6e, 0xdc, 0x44,
	0x18, 0x8d, 0xb3, 0x7f, 0xec, 0xe4, 0xa7, 0xee, 0x34, 0x14, 0x27, 0x42, 0xeb, 0x68, 0x0b, 0x52,
	0x2a, 0x35, 0x36, 0x2a, 0x52, 0xa5, 0x72, 0x57, 0x47, 0x89, 0xf8, 0x09, 0x0a, 0xdd, 0x5d, 0xee,
	0xb8, 0x99, 0xb5, 0xbf, 0x3a, 0x83, 0xec, 0x19, 0xd7, 0x33, 0x0e, 0x44, 0x88, 0x4b, 0xfa, 0x30,
	0x7d, 0x04, 0x78, 0x81, 0x4a, 0x08, 0xa9, 0xe2, 0x01, 0x2c, 0xb4, 0x57, 0xe0, 0xa7, 0x40, 0x33,
	0xe3, 0x80, 0xe3, 0xdd, 0x56, 0xb9, 0xec, 0x4d, 0x32, 0xdf, 0x99, 0x73, 0xce, 0x9e, 0xf9, 0xbe,
	0x99, 0x5d, 0xb4, 0x45, 0xf2, 0xf0, 0x9c, 0x5e, 0x80, 0x97, 0xe5, 0x5c, 0x72, 0x8c, 0x52, 0xce,
	0xa8, 0xe4, 0x39, 0x65, 0xf1, 0xde, 0x87, 0x31, 0xe7, 0x71, 0x02, 0x3e, 0xc9, 0xa8, 0x4f, 0x18,
	0xe3, 0x92, 0x48, 0xca, 0x99, 0x30, 0xcc, 0xbd, 0xe3, 0x98, 0xca, 0xf3, 0x62, 0xee, 0x85, 0x3c,
	0xf5, 0x33, 0x60, 0x82, 0xb0, 0x88, 0xfb, 0xe2, 0x07, 0xff, 0x02, 0x18, 0x0d, 0xc1, 0x2f, 0x24,
	0x4d, 0x84, 0x92, 0xc6, 0xc0, 0x9a, 0x6a, 0x9f, 0xb2, 0x30, 0x29, 0x22, 0xb8, 0xb2, 0x39, 0x6c,
	0xd8, 0xc4, 0x3c, 0xe6, 0xbe, 0x86, 0xe7, 0xc5, 0x33, 0x5d, 0xe9, 0x42, 0xaf, 0x6a, 0xfa, 0xc7,
	0x6f, 0xf8, 0x54, 0x95, 0x31, 0x05, 0x49, 0x6a, 0xda, 0x27, 0x6f, 0xa1, 0x25, 0x64, 0x0e, 0x89,
	0xf0, 0x05, 0x24, 0x10, 0x4a, 0x9e, 0xdf, 0x40, 0xf1, 0x8c, 0x42, 0x12, 0x2d, 0x29, 0x36, 0x05,
	0xa8, 0xde, 0x99, 0x6a, 0xfc, 0x7b, 0x07, 0x6d, 0x3e, 0x31, 0xad, 0x7c, 0x5a, 0x40, 0x7e, 0x89,
	0x4f, 0x51, 0x6f, 0x06, 0x3f, 0x4a, 0xe1, 0x58, 0xfb, 0x9d, 0x83, 0x8d, 0x87, 0x1f, 0x78, 0x35,
	0x5d, 0x81, 0x13, 0x78, 0x5e, 0xd0, 0x1c, 0x52, 0x60, 0x32, 0xb8, 0x53, 0x95, 0xee, 0x2d, 0xa9,
	0x98, 0x0f, 0x78, 0x4a, 0x25, 0xa4, 0x99, 0xbc, 0x9c, 0xb4, 0x01, 0xfc, 0x39, 0xea, 0x9f, 0xe8,
	0x14, 0xce, 0xfa, 0xbe, 0x75, 0xb0, 0xf1, 0xd0, 0xf6, 0x4c, 0x28, 0x6f, 0x5a, 0x87, 0x0a, 0x76,
	0xaa, 0xd2, 0xb5, 0x0d, 0xd8, 0x30, 0x5a, 0x42, 0x94, 0xd3, 0xa9, 0xee, 0x80, 0xd3, 0xa9, 0x9d,
	0x4c, 0x43, 0x5a, 0x4e, 0x06, 0x6c, 0x3a, 0xb5, 0x11, 0xfc, 0x14, 0x0d, 0xa7, 0x92, 0xe4, 0x72,
	0x46, 0x53, 0x70, 0xba, 0xda, 0x6c, 0xdb, 0x23, 0x19, 0xf5, 0x14, 0x20, 0x24, 0x49, 0xb3, 0xc0,
	0xa9, 0x4a, 0x77, 0x47, 0x28, 0xd2, 0xa1, 0xa4, 0x29, 0x34, 0xec, 0x56, 0xa2, 0xf8, 0x2b, 0x34,
	0x38, 0x66, 0x91, 0x36, 0xec, 0xad, 0x34, 0xbc, 0x5b, 0x95, 0x2e, 0x06, 0x16, 0xb5, 0xed, 0x56,
	0x60, 0xf8, 0x31, 0x1a, 0xcc, 0x80, 0x11, 0x26, 0x85, 0xd3, 0xdf, 0xef, 0x1c, 0x0c, 0x83, 0xf7,
	0xab, 0xd2, 0xbd, 0x2d, 0x0d, 0xd4, 0xd0, 0x2e, 0x43, 0xe3, 0x7f, 0xd6, 0xd1, 0x76, 0x3d, 0x4d,
	0x35, 0x2b, 0x10, 0x12, 0x3f, 0x42, 0xd6, 0xcc, 0xb1, 0x74, 0xa8, 0x2d, 0x13, 0xea, 0x32, 0x83,
	0xaf, 0x41, 0x92, 0xe0, 0xce, 0xab, 0xd2, 0x5d, 0x7b, 0x5d, 0xba, 0x56, 0x55, 0xba, 0x83, 0x07,
	0x94, 0x25, 0x94, 0xc1, 0xe4, 0x6a, 0x81, 0x4f, 0x90, 0x75, 0x56, 0x0f, 0xed, 0x96, 0xd6, 0x9d,
	0xcd, 0xbf, 0x87, 0x50, 0x6a, 0xe5, 0x5e, 0x43, 0xb9, 0xad, 0xae, 0x6f, 0x23, 0x55, 0xab, 0xc6,
	0x33, 0xd4, 0x9d, 0x66, 0x10, 0xd6, 0x53, 0x1b, 0x79, 0xff, 0x3f, 0x54, 0xef, 0x7a, 0x52, 0xc5,
	0x0a, 0xee, 0x2a, 0x67, 0xe5, 0x2a, 0x32, 0x08, 0x9b, 0xae, 0xd7, 0x6b, 0xfc, 0x1d, 0xea, 0x4f,
	0x25, 0x91, 0x85, 0xa8, 0x07, 0xb8, 0xff, 0x16, 0x5f, 0xcd, 0x0b, 0x9c, 0xda, 0xd9, 0x16, 0xba,
	0x6e, 0xde, 0x90, 0x36, 0xf2, 0xd9, 0xbd, 0x3f, 0x7f, 0xd9, 0x75, 0xd1, 0x86, 0xff, 0xd3, 0x99,
	0x67, 0x26, 0xf1, 0x33, 0xb6, 0xeb, 0xef, 0x9b, 0xc3, 0xdc, 0xb8, 0x8a, 0xf1, 0xaf, 0x16, 0xc2,
	0xcb, 0x27, 0xc0, 0x47, 0xa8, 0xab, 0x3a, 0xac, 0x5b, 0x3e, 0x0c, 0xfc, 0x97, 0x2f, 0x76, 0xef,
	0x4d, 0x65, 0x7e, 0xcc, 0x8a, 0xf4, 0x60, 0x99, 0xed, 0x9d, 0xf2, 0x58, 0xb1, 0xef, 0x57, 0xa5,
	0xdb, 0x95, 0x97, 0x19, 0x4c, 0xf4, 0x5f, 0xfc, 0x18, 0xf5, 0xf4, 0x6b, 0xac, 0x07, 0xe0, 0xac,
	0x38, 0x9d, 0xde, 0x0f, 0x86, 0x55, 0xe9, 0xf6, 0x9e, 0xab, 0xe5, 0xc4, 0xfc, 0x1b, 0x7f, 0x84,
	0x06, 0xb5, 0x29, 0x1e, 0xa2, 0xde, 0xf1, 0x05, 0x30, 0x69, 0xaf, 0xe1, 0x6d, 0x84, 0x9e, 0x14,
	0x11, 0x95, 0xa6, 0xb6, 0xc6, 0x7f, 0xac, 0xa3, 0x9d, 0x55, 0x6d, 0xc2, 0xe7, 0xff, 0x35, 0xd6,
	0x1c, 0xe0, 0xe4, 0xe5, 0x8b, 0x5d, 0xef, 0x0d, 0x07, 0xd0, 0xbc, 0xab, 0x3c, 0x5f, 0xf2, 0xb9,
	0x01, 0xee, 0xdf, 0xac, 0xc9, 0xf8, 0x11, 0xea, 0x4f, 0x80, 0x08, 0xce, 0xf4, 0x21, 0x87, 0xe6,
	0xf9, 0xe6, 0x1a, 0x69, 0xea, 0xda, 0x08, 0x3e, 0x44, 0x9d, 0x6f, 0x27, 0x5f, 0xe8, 0xfb, 0x34,
	0x0c, 0x6e, 0x57, 0xa5, 0xbb, 0x55, 0xe4, 0xb4, 0xa1, 0xb8, 0x5e, 0x8e, 0x23, 0x64, 0xb7, 0x13,
	0xe2, 0x2d, 0x34, 0x9c, 0x86, 0xe7, 0x10, 0x15, 0x09, 0x44, 0xf6, 0x1a, 0xde, 0x40, 0x83, 0x49,
	0xc1, 0x18, 0x65, 0xb1, 0x6d, 0xa9, 0xbd, 0x23, 0x9e, 0x66, 0x09, 0x48, 0x88, 0xec, 0x75, 0x8c,
	0x50, 0xff, 0x84, 0x50, 0xc5, 0xeb, 0x28, 0x9e, 0x7a, 0xd1, 0x67, 0x85, 0xb4, 0xbb, 0x78, 0x13,
	0xbd, 0x77, 0x44, 0x58, 0x08, 0x6a, 0xab, 0x37, 0xfe, 0xcd, 0x42, 0x3b, 0xa6, 0x7c, 0xb7, 0x9e,
	0xdf, 0x8d, 0xae, 0x72, 0x60, 0xbf, 0x5a, 0x8c, 0xac, 0xd7, 0x8b, 0x91, 0xf5, 0xd7, 0x62, 0x64,
	0xfd, 0xbd, 0x18, 0xad, 0x7d, 0x63, 0xcd, 0xfb, 0xfa, 0xf7, 0xe1, 0xd3, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x59, 0xca, 0x3c, 0x2e, 0x69, 0x07, 0x00, 0x00,
}
