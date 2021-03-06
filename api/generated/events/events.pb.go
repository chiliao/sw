// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: events.proto

/*
	Package events is a generated protocol buffer package.

	Service name

	It is generated from these files:
		events.proto
		svc_events.proto

	It has these top-level messages:
		Event
		EventAttributes
		EventList
		EventSource
		GetEventRequest
*/
package events

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/pensando/grpc-gateway/third_party/googleapis/google/api"
import _ "github.com/pensando/sw/venice/utils/apigen/annotations"
import _ "github.com/gogo/protobuf/gogoproto"
import api "github.com/pensando/sw/api"
import _ "github.com/pensando/sw/events/generated/eventattrs"

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

// Event is a system notification of a fault, condition or configuration
// that should be user visible. These objects are created internally by
// Event client and persisted in EventDB.
type Event struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,json=,inline,embedded=T" json:",inline"`
	// ObjectMeta.Name will be the UUID for an event object.
	// TODO: Should there be a predefined list of labels or keep it free form ?
	api.ObjectMeta `protobuf:"bytes,2,opt,name=O,json=meta,omitempty,embedded=O" json:"meta,omitempty"`
	// Attributes contains the attributes of an event.
	EventAttributes `protobuf:"bytes,3,opt,name=Attributes,json=,inline,embedded=Attributes" json:",inline"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptorEvents, []int{0} }

// EventAttributes contains all the event attributes
type EventAttributes struct {
	// Severity represents the criticality level of an event
	Severity string `protobuf:"bytes,1,opt,name=Severity,json=severity,omitempty,proto3" json:"severity,omitempty"`
	// Type represents the type of an event. e.g. NICAdmittedEvent, NodeJoined
	Type string `protobuf:"bytes,2,opt,name=Type,json=type,omitempty,proto3" json:"type,omitempty"`
	// Message represents the human readable description of an event
	Message string `protobuf:"bytes,3,opt,name=Message,json=message,omitempty,proto3" json:"message,omitempty"`
	// Category represents the category of an event. e.g. Cluster/Network/Datapath
	Category string `protobuf:"bytes,4,opt,name=Category,json=category,omitempty,proto3" json:"category,omitempty"`
	// ObjectRef is the reference to the object associated with an event
	ObjectRef *api.ObjectRef `protobuf:"bytes,5,opt,name=ObjectRef,json=object-ref,omitempty" json:"object-ref,omitempty"`
	// Source is the component and host/node which generated an event
	Source *EventSource `protobuf:"bytes,6,opt,name=Source,json=source,omitempty" json:"source,omitempty"`
	// Number of occurrence of this event in the active interval
	Count uint32 `protobuf:"varint,7,opt,name=Count,json=count,omitempty,proto3" json:"count,omitempty"`
}

func (m *EventAttributes) Reset()                    { *m = EventAttributes{} }
func (m *EventAttributes) String() string            { return proto.CompactTextString(m) }
func (*EventAttributes) ProtoMessage()               {}
func (*EventAttributes) Descriptor() ([]byte, []int) { return fileDescriptorEvents, []int{1} }

func (m *EventAttributes) GetSeverity() string {
	if m != nil {
		return m.Severity
	}
	return ""
}

func (m *EventAttributes) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *EventAttributes) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *EventAttributes) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *EventAttributes) GetObjectRef() *api.ObjectRef {
	if m != nil {
		return m.ObjectRef
	}
	return nil
}

func (m *EventAttributes) GetSource() *EventSource {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *EventAttributes) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

// list of events
type EventList struct {
	//
	api.TypeMeta `protobuf:"bytes,1,opt,name=T,embedded=T" json:"T"`
	//
	api.ListMeta `protobuf:"bytes,2,opt,name=ListMeta,embedded=ListMeta" json:"ListMeta"`
	//
	Items []*Event `protobuf:"bytes,3,rep,name=Items,json=items,omitempty" json:"items,omitempty"`
}

func (m *EventList) Reset()                    { *m = EventList{} }
func (m *EventList) String() string            { return proto.CompactTextString(m) }
func (*EventList) ProtoMessage()               {}
func (*EventList) Descriptor() ([]byte, []int) { return fileDescriptorEvents, []int{2} }

func (m *EventList) GetItems() []*Event {
	if m != nil {
		return m.Items
	}
	return nil
}

// EventSource has info about the component and
// host/node that generated the event
type EventSource struct {
	// Component from which the event is generated.
	Component string `protobuf:"bytes,1,opt,name=Component,json=component,omitempty,proto3" json:"component,omitempty"`
	// Name of the venice or workload node which is generating the event.
	NodeName string `protobuf:"bytes,2,opt,name=NodeName,json=node-name,omitempty,proto3" json:"node-name,omitempty"`
}

func (m *EventSource) Reset()                    { *m = EventSource{} }
func (m *EventSource) String() string            { return proto.CompactTextString(m) }
func (*EventSource) ProtoMessage()               {}
func (*EventSource) Descriptor() ([]byte, []int) { return fileDescriptorEvents, []int{3} }

func (m *EventSource) GetComponent() string {
	if m != nil {
		return m.Component
	}
	return ""
}

func (m *EventSource) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "events.Event")
	proto.RegisterType((*EventAttributes)(nil), "events.EventAttributes")
	proto.RegisterType((*EventList)(nil), "events.EventList")
	proto.RegisterType((*EventSource)(nil), "events.EventSource")
}
func (m *Event) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Event) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintEvents(dAtA, i, uint64(m.TypeMeta.Size()))
	n1, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x12
	i++
	i = encodeVarintEvents(dAtA, i, uint64(m.ObjectMeta.Size()))
	n2, err := m.ObjectMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x1a
	i++
	i = encodeVarintEvents(dAtA, i, uint64(m.EventAttributes.Size()))
	n3, err := m.EventAttributes.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *EventAttributes) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAttributes) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Severity) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Severity)))
		i += copy(dAtA[i:], m.Severity)
	}
	if len(m.Type) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Type)))
		i += copy(dAtA[i:], m.Type)
	}
	if len(m.Message) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Message)))
		i += copy(dAtA[i:], m.Message)
	}
	if len(m.Category) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Category)))
		i += copy(dAtA[i:], m.Category)
	}
	if m.ObjectRef != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.ObjectRef.Size()))
		n4, err := m.ObjectRef.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n4
	}
	if m.Source != nil {
		dAtA[i] = 0x32
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.Source.Size()))
		n5, err := m.Source.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	if m.Count != 0 {
		dAtA[i] = 0x38
		i++
		i = encodeVarintEvents(dAtA, i, uint64(m.Count))
	}
	return i, nil
}

func (m *EventList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintEvents(dAtA, i, uint64(m.TypeMeta.Size()))
	n6, err := m.TypeMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	dAtA[i] = 0x12
	i++
	i = encodeVarintEvents(dAtA, i, uint64(m.ListMeta.Size()))
	n7, err := m.ListMeta.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintEvents(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *EventSource) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventSource) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Component) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Component)))
		i += copy(dAtA[i:], m.Component)
	}
	if len(m.NodeName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEvents(dAtA, i, uint64(len(m.NodeName)))
		i += copy(dAtA[i:], m.NodeName)
	}
	return i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Event) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.ObjectMeta.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.EventAttributes.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventAttributes) Size() (n int) {
	var l int
	_ = l
	l = len(m.Severity)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Message)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.Category)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ObjectRef != nil {
		l = m.ObjectRef.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Source != nil {
		l = m.Source.Size()
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovEvents(uint64(m.Count))
	}
	return n
}

func (m *EventList) Size() (n int) {
	var l int
	_ = l
	l = m.TypeMeta.Size()
	n += 1 + l + sovEvents(uint64(l))
	l = m.ListMeta.Size()
	n += 1 + l + sovEvents(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	return n
}

func (m *EventSource) Size() (n int) {
	var l int
	_ = l
	l = len(m.Component)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	l = len(m.NodeName)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Event) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: Event: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Event: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
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
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
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
				return fmt.Errorf("proto: wrong wireType = %d for field EventAttributes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EventAttributes.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventAttributes) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAttributes: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAttributes: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Severity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Severity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Message = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Category", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Category = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectRef", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ObjectRef == nil {
				m.ObjectRef = &api.ObjectRef{}
			}
			if err := m.ObjectRef.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Source == nil {
				m.Source = &EventSource{}
			}
			if err := m.Source.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= (uint32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TypeMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
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
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, &Event{})
			if err := m.Items[len(m.Items)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventSource) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventSource: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventSource: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Component", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Component = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NodeName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEvents
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
				next, err := skipEvents(dAtA[start:])
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
	ErrInvalidLengthEvents = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("events.proto", fileDescriptorEvents) }

var fileDescriptorEvents = []byte{
	// 636 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xd1, 0x4e, 0x13, 0x4d,
	0x14, 0x66, 0x7e, 0x68, 0x69, 0x87, 0x1f, 0x8a, 0x53, 0x0c, 0x2b, 0xd1, 0x2e, 0x69, 0x42, 0x82,
	0x09, 0x74, 0x8d, 0x44, 0x8d, 0x37, 0x1a, 0x97, 0x40, 0x42, 0x14, 0xd0, 0xc2, 0xad, 0x17, 0xdb,
	0xed, 0xe9, 0x3a, 0xa6, 0x3b, 0xb3, 0xd9, 0x99, 0xc5, 0xf4, 0x05, 0xbc, 0x34, 0xf1, 0x35, 0xbc,
	0xf1, 0x35, 0xb8, 0x24, 0x3e, 0xc0, 0xc6, 0xf4, 0xca, 0xec, 0x53, 0x98, 0x99, 0xdd, 0xc5, 0xa1,
	0xb4, 0xdc, 0xed, 0xf9, 0xe6, 0xfb, 0xce, 0x7c, 0xdf, 0x9c, 0x93, 0xc5, 0xff, 0xc3, 0x05, 0x30,
	0x29, 0x3a, 0x51, 0xcc, 0x25, 0x27, 0xd5, 0xbc, 0xda, 0x78, 0x18, 0x70, 0x1e, 0x0c, 0xc1, 0xf1,
	0x22, 0xea, 0x78, 0x8c, 0x71, 0xe9, 0x49, 0xca, 0x59, 0xc1, 0xda, 0x38, 0x08, 0xa8, 0xfc, 0x94,
	0xf4, 0x3a, 0x3e, 0x0f, 0x9d, 0x08, 0x98, 0xf0, 0x58, 0x9f, 0x3b, 0xe2, 0x8b, 0x73, 0x01, 0x8c,
	0xfa, 0xe0, 0x24, 0x92, 0x0e, 0x85, 0x92, 0x06, 0xc0, 0x4c, 0xb5, 0x43, 0x99, 0x3f, 0x4c, 0xfa,
	0x50, 0xb6, 0xd9, 0x35, 0xda, 0x04, 0x3c, 0xe0, 0x8e, 0x86, 0x7b, 0xc9, 0x40, 0x57, 0xba, 0xd0,
	0x5f, 0x05, 0x7d, 0x6b, 0xc6, 0xad, 0xca, 0x63, 0x08, 0xd2, 0x2b, 0x68, 0xcf, 0x66, 0xd0, 0xf2,
	0x64, 0xf9, 0x1d, 0xc2, 0xf1, 0xa4, 0x8c, 0x69, 0x2f, 0x91, 0xa5, 0x99, 0xf6, 0x2f, 0x84, 0x2b,
	0x07, 0x8a, 0x42, 0x9e, 0x63, 0x74, 0x6e, 0xa1, 0x4d, 0xb4, 0xbd, 0xf4, 0x74, 0xb9, 0xe3, 0x45,
	0xb4, 0x73, 0x3e, 0x8a, 0xe0, 0x18, 0xa4, 0xe7, 0x36, 0x2f, 0x53, 0x7b, 0xee, 0x2a, 0xb5, 0x51,
	0x96, 0xda, 0x8b, 0x3b, 0x94, 0x0d, 0x29, 0x83, 0x6e, 0xf9, 0x41, 0x0e, 0x31, 0x3a, 0xb5, 0xfe,
	0xd3, 0xba, 0x86, 0xd6, 0x9d, 0xf6, 0x3e, 0x83, 0x2f, 0xb5, 0x72, 0xc3, 0x50, 0xae, 0x28, 0xb3,
	0x3b, 0x3c, 0xa4, 0x12, 0xc2, 0x48, 0x8e, 0xba, 0x13, 0x35, 0x39, 0xc2, 0xf8, 0xcd, 0xb5, 0x3b,
	0x6b, 0x5e, 0x37, 0x5c, 0xef, 0x14, 0x63, 0xd2, 0x16, 0xff, 0x1d, 0xdf, 0x6d, 0xa9, 0xfd, 0x7d,
	0x01, 0x37, 0x26, 0x14, 0xe4, 0x23, 0xae, 0x9d, 0xc1, 0x05, 0xc4, 0x54, 0x8e, 0x74, 0xca, 0xba,
	0xfb, 0xe2, 0xc7, 0xd7, 0x07, 0x8f, 0xce, 0x64, 0x7c, 0xc0, 0x92, 0x70, 0x3b, 0xbf, 0x48, 0x3d,
	0x8f, 0xe8, 0x94, 0xbc, 0xc7, 0x59, 0x6a, 0x13, 0x51, 0x14, 0x86, 0xf3, 0x29, 0x18, 0x79, 0x82,
	0x17, 0xd4, 0x7b, 0xe9, 0x87, 0xa8, 0xbb, 0x44, 0xe5, 0x95, 0xa3, 0x08, 0xcc, 0xbc, 0x37, 0x6b,
	0xf2, 0x12, 0x2f, 0x1e, 0x83, 0x10, 0x5e, 0x00, 0x3a, 0x6c, 0xdd, 0xbd, 0x9f, 0xa5, 0xf6, 0xbd,
	0x30, 0x87, 0x0c, 0xdd, 0x6d, 0x48, 0x65, 0xd9, 0xf7, 0x24, 0x04, 0x3c, 0x1e, 0x59, 0x0b, 0x77,
	0x67, 0x29, 0x79, 0x3a, 0x8b, 0x5f, 0x14, 0x66, 0x96, 0xdb, 0x18, 0xf9, 0x80, 0xeb, 0xf9, 0x0c,
	0xbb, 0x30, 0xb0, 0x2a, 0x7a, 0x10, 0x2b, 0xc6, 0x64, 0xbb, 0x30, 0x70, 0xad, 0x2c, 0xb5, 0xd7,
	0xb8, 0x2e, 0x77, 0x63, 0x18, 0x18, 0x0d, 0xa7, 0xa2, 0xe4, 0x2d, 0xae, 0x9e, 0xf1, 0x24, 0xf6,
	0xc1, 0xaa, 0xea, 0x7e, 0xcd, 0x1b, 0x83, 0xcd, 0x8f, 0xdc, 0xb5, 0x2c, 0xb5, 0x57, 0x85, 0xfe,
	0x36, 0x1a, 0xde, 0x42, 0xc8, 0x1e, 0xae, 0xec, 0xf3, 0x84, 0x49, 0x6b, 0x71, 0x13, 0x6d, 0x2f,
	0xbb, 0xcd, 0x2c, 0xb5, 0x1b, 0xbe, 0x02, 0x0c, 0xd5, 0x24, 0xd0, 0xfe, 0x89, 0x70, 0x5d, 0x5f,
	0xf6, 0x8e, 0x0a, 0x49, 0xb6, 0x66, 0x2e, 0x7b, 0xad, 0xdc, 0xac, 0x2e, 0x3a, 0x27, 0x7b, 0xb8,
	0xa6, 0xe8, 0xea, 0xa0, 0x58, 0xf1, 0x9c, 0x5d, 0x82, 0x06, 0xfb, 0x9a, 0x48, 0x0e, 0x71, 0xe5,
	0x48, 0x42, 0xa8, 0x76, 0x78, 0x5e, 0x2b, 0xcc, 0xa8, 0xee, 0xfa, 0x65, 0xbe, 0xb5, 0x0d, 0x65,
	0x4c, 0x98, 0x8e, 0x27, 0x80, 0xf6, 0x37, 0x84, 0x97, 0x8c, 0xe7, 0x21, 0xaf, 0x71, 0x7d, 0x9f,
	0x87, 0x11, 0x67, 0xc0, 0x64, 0xb1, 0xc2, 0xeb, 0x59, 0x6a, 0x37, 0xfd, 0x12, 0x34, 0x9a, 0x4d,
	0x03, 0xc9, 0x2b, 0x5c, 0x3b, 0xe1, 0x7d, 0x38, 0xf1, 0xc2, 0x72, 0x4f, 0xb5, 0x9e, 0xf1, 0x3e,
	0xec, 0x32, 0x2f, 0x34, 0x1f, 0x7d, 0x1a, 0xe8, 0xae, 0x5e, 0x8e, 0x5b, 0xe8, 0x6a, 0xdc, 0x42,
	0xbf, 0xc7, 0x2d, 0xf4, 0x67, 0xdc, 0x9a, 0x7b, 0x8f, 0x7a, 0x55, 0xfd, 0x1b, 0xd9, 0xfb, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x52, 0x7d, 0x53, 0x43, 0x50, 0x05, 0x00, 0x00,
}
