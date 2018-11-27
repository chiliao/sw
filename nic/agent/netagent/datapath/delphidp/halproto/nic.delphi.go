// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nic.proto

package halproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// NIC boots in classic mode first and may later transition to flow mode
type DeviceMode int32

const (
	DeviceMode_DEVICE_MODE_NONE             DeviceMode = 0
	DeviceMode_DEVICE_MODE_MANAGED_SWITCH   DeviceMode = 1
	DeviceMode_DEVICE_MODE_MANAGED_HOST_PIN DeviceMode = 2
	DeviceMode_DEVICE_MODE_STANDALONE       DeviceMode = 3
)

var DeviceMode_name = map[int32]string{
	0: "DEVICE_MODE_NONE",
	1: "DEVICE_MODE_MANAGED_SWITCH",
	2: "DEVICE_MODE_MANAGED_HOST_PIN",
	3: "DEVICE_MODE_STANDALONE",
}
var DeviceMode_value = map[string]int32{
	"DEVICE_MODE_NONE":             0,
	"DEVICE_MODE_MANAGED_SWITCH":   1,
	"DEVICE_MODE_MANAGED_HOST_PIN": 2,
	"DEVICE_MODE_STANDALONE":       3,
}

func (x DeviceMode) String() string {
	return proto.EnumName(DeviceMode_name, int32(x))
}
func (DeviceMode) EnumDescriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

// Global config object for NIC
type DeviceSpec struct {
	DeviceMode          DeviceMode `protobuf:"varint,1,opt,name=device_mode,json=deviceMode,enum=device.DeviceMode" json:"device_mode,omitempty"`
	AllowDynamicPinning bool       `protobuf:"varint,2,opt,name=allow_dynamic_pinning,json=allowDynamicPinning" json:"allow_dynamic_pinning,omitempty"`
	LocalMacAddress     uint64     `protobuf:"fixed64,3,opt,name=local_mac_address,json=localMacAddress" json:"local_mac_address,omitempty"`
}

func (m *DeviceSpec) Reset()                    { *m = DeviceSpec{} }
func (m *DeviceSpec) String() string            { return proto.CompactTextString(m) }
func (*DeviceSpec) ProtoMessage()               {}
func (*DeviceSpec) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{0} }

func (m *DeviceSpec) GetDeviceMode() DeviceMode {
	if m != nil {
		return m.DeviceMode
	}
	return DeviceMode_DEVICE_MODE_NONE
}

func (m *DeviceSpec) GetAllowDynamicPinning() bool {
	if m != nil {
		return m.AllowDynamicPinning
	}
	return false
}

func (m *DeviceSpec) GetLocalMacAddress() uint64 {
	if m != nil {
		return m.LocalMacAddress
	}
	return 0
}

type DeviceRequest struct {
	Device *DeviceSpec `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
}

func (m *DeviceRequest) Reset()                    { *m = DeviceRequest{} }
func (m *DeviceRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceRequest) ProtoMessage()               {}
func (*DeviceRequest) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{1} }

func (m *DeviceRequest) GetDevice() *DeviceSpec {
	if m != nil {
		return m.Device
	}
	return nil
}

type DeviceRequestMsg struct {
	Request *DeviceRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *DeviceRequestMsg) Reset()                    { *m = DeviceRequestMsg{} }
func (m *DeviceRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*DeviceRequestMsg) ProtoMessage()               {}
func (*DeviceRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{2} }

func (m *DeviceRequestMsg) GetRequest() *DeviceRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type DeviceResponse struct {
	ApiStatus ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
}

func (m *DeviceResponse) Reset()                    { *m = DeviceResponse{} }
func (m *DeviceResponse) String() string            { return proto.CompactTextString(m) }
func (*DeviceResponse) ProtoMessage()               {}
func (*DeviceResponse) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{3} }

func (m *DeviceResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

type DeviceResponseMsg struct {
	Response *DeviceResponse `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *DeviceResponseMsg) Reset()                    { *m = DeviceResponseMsg{} }
func (m *DeviceResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*DeviceResponseMsg) ProtoMessage()               {}
func (*DeviceResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{4} }

func (m *DeviceResponseMsg) GetResponse() *DeviceResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

type DeviceGetRequest struct {
}

func (m *DeviceGetRequest) Reset()                    { *m = DeviceGetRequest{} }
func (m *DeviceGetRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceGetRequest) ProtoMessage()               {}
func (*DeviceGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{5} }

type DeviceGetRequestMsg struct {
	Request *DeviceGetRequest `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
}

func (m *DeviceGetRequestMsg) Reset()                    { *m = DeviceGetRequestMsg{} }
func (m *DeviceGetRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*DeviceGetRequestMsg) ProtoMessage()               {}
func (*DeviceGetRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{6} }

func (m *DeviceGetRequestMsg) GetRequest() *DeviceGetRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

type DeviceGetResponse struct {
	ApiStatus ApiStatus   `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	Device    *DeviceSpec `protobuf:"bytes,2,opt,name=device" json:"device,omitempty"`
}

func (m *DeviceGetResponse) Reset()                    { *m = DeviceGetResponse{} }
func (m *DeviceGetResponse) String() string            { return proto.CompactTextString(m) }
func (*DeviceGetResponse) ProtoMessage()               {}
func (*DeviceGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{7} }

func (m *DeviceGetResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *DeviceGetResponse) GetDevice() *DeviceSpec {
	if m != nil {
		return m.Device
	}
	return nil
}

type DeviceGetResponseMsg struct {
	Response *DeviceGetResponse `protobuf:"bytes,1,opt,name=response" json:"response,omitempty"`
}

func (m *DeviceGetResponseMsg) Reset()                    { *m = DeviceGetResponseMsg{} }
func (m *DeviceGetResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*DeviceGetResponseMsg) ProtoMessage()               {}
func (*DeviceGetResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor24, []int{8} }

func (m *DeviceGetResponseMsg) GetResponse() *DeviceGetResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceSpec)(nil), "halproto.DeviceSpec")
	proto.RegisterType((*DeviceRequest)(nil), "halproto.DeviceRequest")
	proto.RegisterType((*DeviceRequestMsg)(nil), "halproto.DeviceRequestMsg")
	proto.RegisterType((*DeviceResponse)(nil), "halproto.DeviceResponse")
	proto.RegisterType((*DeviceResponseMsg)(nil), "halproto.DeviceResponseMsg")
	proto.RegisterType((*DeviceGetRequest)(nil), "halproto.DeviceGetRequest")
	proto.RegisterType((*DeviceGetRequestMsg)(nil), "halproto.DeviceGetRequestMsg")
	proto.RegisterType((*DeviceGetResponse)(nil), "halproto.DeviceGetResponse")
	proto.RegisterType((*DeviceGetResponseMsg)(nil), "halproto.DeviceGetResponseMsg")
	proto.RegisterEnum("halproto.DeviceMode", DeviceMode_name, DeviceMode_value)
}

func init() { proto.RegisterFile("nic.proto", fileDescriptor24) }

var fileDescriptor24 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0xcd, 0x6e, 0xd3, 0x40,
	0x18, 0xac, 0x1b, 0x29, 0x24, 0x5f, 0xa0, 0x38, 0xdb, 0xa6, 0x72, 0x43, 0x05, 0x91, 0x4f, 0x51,
	0x0e, 0x41, 0x4a, 0xc5, 0x05, 0x4e, 0x6e, 0x6c, 0x92, 0x48, 0xd8, 0xa9, 0xec, 0xf0, 0x23, 0x2e,
	0xab, 0xc5, 0x5e, 0x85, 0x95, 0x1c, 0xdb, 0x78, 0xdd, 0xa2, 0xde, 0xb8, 0xf1, 0x1c, 0x3c, 0x18,
	0x3c, 0x04, 0x4f, 0x80, 0xbc, 0x6b, 0xe7, 0xc7, 0x4d, 0xc5, 0x01, 0x4e, 0xc9, 0xce, 0x4c, 0x26,
	0xf3, 0xcd, 0x7e, 0x36, 0x34, 0x23, 0xe6, 0x0f, 0x93, 0x34, 0xce, 0x62, 0x54, 0x0f, 0xe8, 0x0d,
	0xf3, 0x69, 0xb7, 0x95, 0xdd, 0x26, 0x94, 0x4b, 0x50, 0xff, 0xa1, 0x00, 0x98, 0x02, 0xf7, 0x12,
	0xea, 0xa3, 0x0b, 0x68, 0x49, 0x15, 0x5e, 0xc5, 0x01, 0xd5, 0x94, 0x9e, 0xd2, 0x3f, 0x1a, 0xa1,
	0xa1, 0xc4, 0x86, 0x52, 0x68, 0xc7, 0x01, 0x75, 0x21, 0x58, 0x7f, 0x47, 0x23, 0xe8, 0x90, 0x30,
	0x8c, 0xbf, 0xe2, 0xe0, 0x36, 0x22, 0x2b, 0xe6, 0xe3, 0x84, 0x45, 0x11, 0x8b, 0x96, 0xda, 0x61,
	0x4f, 0xe9, 0x37, 0xdc, 0x63, 0x41, 0x9a, 0x92, 0xbb, 0x92, 0x14, 0x1a, 0x40, 0x3b, 0x8c, 0x7d,
	0x12, 0xe2, 0x15, 0xf1, 0x31, 0x09, 0x82, 0x94, 0x72, 0xae, 0xd5, 0x7a, 0x4a, 0xbf, 0xee, 0x3e,
	0x16, 0x84, 0x4d, 0x7c, 0x43, 0xc2, 0xfa, 0x2b, 0x78, 0x24, 0xff, 0xd9, 0xa5, 0x5f, 0xae, 0x29,
	0xcf, 0xd0, 0x00, 0x8a, 0x59, 0x44, 0xc0, 0x56, 0x35, 0x60, 0x3e, 0x89, 0x5b, 0x28, 0xf4, 0x31,
	0xa8, 0x3b, 0x3f, 0xb6, 0xf9, 0x12, 0x3d, 0x87, 0x07, 0xa9, 0x3c, 0x15, 0x06, 0x9d, 0x5d, 0x83,
	0x42, 0xea, 0x96, 0x2a, 0xfd, 0x03, 0x1c, 0x95, 0x0c, 0x4f, 0xe2, 0x88, 0x53, 0xf4, 0x1a, 0x80,
	0x24, 0x0c, 0xf3, 0x8c, 0x64, 0xd7, 0xbc, 0xe8, 0x49, 0x1d, 0xca, 0x66, 0x8d, 0x84, 0x79, 0x02,
	0xbf, 0xec, 0xfc, 0xfe, 0xf5, 0xac, 0x7d, 0x43, 0x23, 0xe6, 0xd3, 0x97, 0x1b, 0xb9, 0xdb, 0x24,
	0xa5, 0x42, 0x9f, 0x40, 0x7b, 0xd7, 0x39, 0xcf, 0x37, 0x82, 0x46, 0x5a, 0x1c, 0x8b, 0x80, 0xa7,
	0xd5, 0x80, 0x92, 0x75, 0xd7, 0x3a, 0x1d, 0x95, 0x73, 0x4e, 0x68, 0x56, 0xe4, 0xd7, 0x67, 0x70,
	0x5c, 0xc5, 0xa4, 0x7d, 0x65, 0x7c, 0x6d, 0xd7, 0x7d, 0xa3, 0xde, 0x34, 0xf0, 0x5d, 0x29, 0x83,
	0x0a, 0xf6, 0xff, 0xb6, 0xb0, 0x75, 0xa1, 0x87, 0x7f, 0xbd, 0x50, 0x1b, 0x4e, 0xee, 0x04, 0xc9,
	0xa7, 0x7a, 0x71, 0xa7, 0xb4, 0xb3, 0x3d, 0x63, 0x55, 0x7b, 0x1b, 0x7c, 0x5b, 0x3f, 0x00, 0x62,
	0x97, 0x4f, 0x40, 0x35, 0xad, 0x77, 0xb3, 0xb1, 0x85, 0xed, 0xb9, 0x69, 0x61, 0x67, 0xee, 0x58,
	0xea, 0x01, 0x7a, 0x0a, 0xdd, 0x6d, 0xd4, 0x36, 0x1c, 0x63, 0x62, 0x99, 0xd8, 0x7b, 0x3f, 0x5b,
	0x8c, 0xa7, 0xaa, 0x82, 0x7a, 0x70, 0xbe, 0x8f, 0x9f, 0xce, 0xbd, 0x05, 0xbe, 0x9a, 0x39, 0xea,
	0x21, 0xea, 0xc2, 0xe9, 0xb6, 0xc2, 0x5b, 0x18, 0x8e, 0x69, 0xbc, 0xc9, 0xdd, 0x6b, 0xa3, 0x9f,
	0x0a, 0xd4, 0x1c, 0xe6, 0x23, 0x0b, 0x1e, 0xca, 0x24, 0xe3, 0x94, 0x92, 0x8c, 0x22, 0x6d, 0xef,
	0x56, 0xda, 0x7c, 0xd9, 0x3d, 0xdb, 0xbf, 0x0e, 0x36, 0x5f, 0xea, 0x07, 0x1b, 0x9b, 0xb7, 0x49,
	0xf0, 0x0f, 0x36, 0x53, 0x68, 0xae, 0x7b, 0x43, 0x4f, 0xee, 0xdb, 0x90, 0xdc, 0xe6, 0xfc, 0xde,
	0x9e, 0x85, 0xd3, 0x25, 0x7c, 0x6c, 0x7c, 0x26, 0xa1, 0x78, 0xdf, 0x7c, 0xaa, 0x8b, 0x8f, 0x8b,
	0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x08, 0x3d, 0xda, 0xcc, 0x98, 0x04, 0x00, 0x00,
}
