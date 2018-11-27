// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tls_proxy_cb2.proto

package halproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// TlsProxyCbKeyHandle is used to operate on a tls_proxy_cb either by its key or handle
type TlsProxyCbKeyHandle struct {
	// Types that are valid to be assigned to KeyOrHandle:
	//	*TlsProxyCbKeyHandle_TlsProxyCbId
	//	*TlsProxyCbKeyHandle_TlsProxyCbHandle
	KeyOrHandle isTlsProxyCbKeyHandle_KeyOrHandle `protobuf_oneof:"key_or_handle"`
}

func (m *TlsProxyCbKeyHandle) Reset()                    { *m = TlsProxyCbKeyHandle{} }
func (m *TlsProxyCbKeyHandle) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbKeyHandle) ProtoMessage()               {}
func (*TlsProxyCbKeyHandle) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{0} }

type isTlsProxyCbKeyHandle_KeyOrHandle interface{ isTlsProxyCbKeyHandle_KeyOrHandle() }

type TlsProxyCbKeyHandle_TlsProxyCbId struct {
	TlsProxyCbId uint32 `protobuf:"varint,1,opt,name=tls_proxy_cb_id,json=tlsProxyCbId,oneof"`
}
type TlsProxyCbKeyHandle_TlsProxyCbHandle struct {
	TlsProxyCbHandle uint64 `protobuf:"fixed64,2,opt,name=tls_proxy_cb_handle,json=tlsProxyCbHandle,oneof"`
}

func (*TlsProxyCbKeyHandle_TlsProxyCbId) isTlsProxyCbKeyHandle_KeyOrHandle()     {}
func (*TlsProxyCbKeyHandle_TlsProxyCbHandle) isTlsProxyCbKeyHandle_KeyOrHandle() {}

func (m *TlsProxyCbKeyHandle) GetKeyOrHandle() isTlsProxyCbKeyHandle_KeyOrHandle {
	if m != nil {
		return m.KeyOrHandle
	}
	return nil
}

func (m *TlsProxyCbKeyHandle) GetTlsProxyCbId() uint32 {
	if x, ok := m.GetKeyOrHandle().(*TlsProxyCbKeyHandle_TlsProxyCbId); ok {
		return x.TlsProxyCbId
	}
	return 0
}

func (m *TlsProxyCbKeyHandle) GetTlsProxyCbHandle() uint64 {
	if x, ok := m.GetKeyOrHandle().(*TlsProxyCbKeyHandle_TlsProxyCbHandle); ok {
		return x.TlsProxyCbHandle
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TlsProxyCbKeyHandle) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TlsProxyCbKeyHandle_OneofMarshaler, _TlsProxyCbKeyHandle_OneofUnmarshaler, _TlsProxyCbKeyHandle_OneofSizer, []interface{}{
		(*TlsProxyCbKeyHandle_TlsProxyCbId)(nil),
		(*TlsProxyCbKeyHandle_TlsProxyCbHandle)(nil),
	}
}

func _TlsProxyCbKeyHandle_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TlsProxyCbKeyHandle)
	// key_or_handle
	switch x := m.KeyOrHandle.(type) {
	case *TlsProxyCbKeyHandle_TlsProxyCbId:
		b.EncodeVarint(1<<3 | proto.WireVarint)
		b.EncodeVarint(uint64(x.TlsProxyCbId))
	case *TlsProxyCbKeyHandle_TlsProxyCbHandle:
		b.EncodeVarint(2<<3 | proto.WireFixed64)
		b.EncodeFixed64(uint64(x.TlsProxyCbHandle))
	case nil:
	default:
		return fmt.Errorf("TlsProxyCbKeyHandle.KeyOrHandle has unexpected type %T", x)
	}
	return nil
}

func _TlsProxyCbKeyHandle_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TlsProxyCbKeyHandle)
	switch tag {
	case 1: // key_or_handle.tls_proxy_cb_id
		if wire != proto.WireVarint {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeVarint()
		m.KeyOrHandle = &TlsProxyCbKeyHandle_TlsProxyCbId{uint32(x)}
		return true, err
	case 2: // key_or_handle.tls_proxy_cb_handle
		if wire != proto.WireFixed64 {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeFixed64()
		m.KeyOrHandle = &TlsProxyCbKeyHandle_TlsProxyCbHandle{x}
		return true, err
	default:
		return false, nil
	}
}

func _TlsProxyCbKeyHandle_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TlsProxyCbKeyHandle)
	// key_or_handle
	switch x := m.KeyOrHandle.(type) {
	case *TlsProxyCbKeyHandle_TlsProxyCbId:
		n += proto.SizeVarint(1<<3 | proto.WireVarint)
		n += proto.SizeVarint(uint64(x.TlsProxyCbId))
	case *TlsProxyCbKeyHandle_TlsProxyCbHandle:
		n += proto.SizeVarint(2<<3 | proto.WireFixed64)
		n += 8
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// TlsProxyCbSpec captures all the tls_proxy_cb level configuration
type TlsProxyCbSpec struct {
	KeyOrHandle             *TlsProxyCbKeyHandle `protobuf:"bytes,1,opt,name=key_or_handle,json=keyOrHandle" json:"key_or_handle,omitempty"`
	NicDecHead              uint32               `protobuf:"varint,2,opt,name=nic_dec_head,json=nicDecHead" json:"nic_dec_head,omitempty"`
	NicDecTail              uint32               `protobuf:"varint,3,opt,name=nic_dec_tail,json=nicDecTail" json:"nic_dec_tail,omitempty"`
	Command                 uint32               `protobuf:"varint,4,opt,name=command" json:"command,omitempty"`
	DebugDol                uint32               `protobuf:"varint,5,opt,name=debug_dol,json=debugDol" json:"debug_dol,omitempty"`
	SerqPi                  uint32               `protobuf:"varint,6,opt,name=serq_pi,json=serqPi" json:"serq_pi,omitempty"`
	SerqCi                  uint32               `protobuf:"varint,7,opt,name=serq_ci,json=serqCi" json:"serq_ci,omitempty"`
	BsqPi                   uint32               `protobuf:"varint,8,opt,name=bsq_pi,json=bsqPi" json:"bsq_pi,omitempty"`
	BsqCi                   uint32               `protobuf:"varint,9,opt,name=bsq_ci,json=bsqCi" json:"bsq_ci,omitempty"`
	CryptoKeyIdx            uint32               `protobuf:"varint,10,opt,name=crypto_key_idx,json=cryptoKeyIdx" json:"crypto_key_idx,omitempty"`
	SerqBase                uint32               `protobuf:"varint,11,opt,name=serq_base,json=serqBase" json:"serq_base,omitempty"`
	SesqBase                uint32               `protobuf:"varint,12,opt,name=sesq_base,json=sesqBase" json:"sesq_base,omitempty"`
	TnmdrAlloc              uint64               `protobuf:"fixed64,13,opt,name=tnmdr_alloc,json=tnmdrAlloc" json:"tnmdr_alloc,omitempty"`
	TnmprAlloc              uint64               `protobuf:"fixed64,14,opt,name=tnmpr_alloc,json=tnmprAlloc" json:"tnmpr_alloc,omitempty"`
	RnmdrFree               uint64               `protobuf:"fixed64,15,opt,name=rnmdr_free,json=rnmdrFree" json:"rnmdr_free,omitempty"`
	RnmprFree               uint64               `protobuf:"fixed64,16,opt,name=rnmpr_free,json=rnmprFree" json:"rnmpr_free,omitempty"`
	EncRequests             uint64               `protobuf:"fixed64,17,opt,name=enc_requests,json=encRequests" json:"enc_requests,omitempty"`
	EncCompletions          uint64               `protobuf:"fixed64,18,opt,name=enc_completions,json=encCompletions" json:"enc_completions,omitempty"`
	EncFailures             uint64               `protobuf:"fixed64,19,opt,name=enc_failures,json=encFailures" json:"enc_failures,omitempty"`
	DecRequests             uint64               `protobuf:"fixed64,20,opt,name=dec_requests,json=decRequests" json:"dec_requests,omitempty"`
	DecCompletions          uint64               `protobuf:"fixed64,21,opt,name=dec_completions,json=decCompletions" json:"dec_completions,omitempty"`
	DecFailures             uint64               `protobuf:"fixed64,22,opt,name=dec_failures,json=decFailures" json:"dec_failures,omitempty"`
	Salt                    uint32               `protobuf:"varint,23,opt,name=salt" json:"salt,omitempty"`
	ExplicitIv              uint64               `protobuf:"varint,24,opt,name=explicit_iv,json=explicitIv" json:"explicit_iv,omitempty"`
	PreDebugStage0_7Thread  uint32               `protobuf:"varint,25,opt,name=pre_debug_stage0_7_thread,json=preDebugStage07Thread" json:"pre_debug_stage0_7_thread,omitempty"`
	PostDebugStage0_7Thread uint32               `protobuf:"varint,26,opt,name=post_debug_stage0_7_thread,json=postDebugStage07Thread" json:"post_debug_stage0_7_thread,omitempty"`
	IsDecryptFlow           bool                 `protobuf:"varint,27,opt,name=is_decrypt_flow,json=isDecryptFlow" json:"is_decrypt_flow,omitempty"`
	OtherFid                uint32               `protobuf:"varint,28,opt,name=other_fid,json=otherFid" json:"other_fid,omitempty"`
	L7ProxyType             AppRedirType         `protobuf:"varint,29,opt,name=l7_proxy_type,json=l7ProxyType,enum=types.AppRedirType" json:"l7_proxy_type,omitempty"`
	CryptoHmacKeyIdx        uint32               `protobuf:"varint,30,opt,name=crypto_hmac_key_idx,json=cryptoHmacKeyIdx" json:"crypto_hmac_key_idx,omitempty"`
	MacRequests             uint64               `protobuf:"fixed64,31,opt,name=mac_requests,json=macRequests" json:"mac_requests,omitempty"`
	MacCompletions          uint64               `protobuf:"fixed64,32,opt,name=mac_completions,json=macCompletions" json:"mac_completions,omitempty"`
	MacFailures             uint64               `protobuf:"fixed64,33,opt,name=mac_failures,json=macFailures" json:"mac_failures,omitempty"`
	CpuId                   uint32               `protobuf:"varint,34,opt,name=cpu_id,json=cpuId" json:"cpu_id,omitempty"`
}

func (m *TlsProxyCbSpec) Reset()                    { *m = TlsProxyCbSpec{} }
func (m *TlsProxyCbSpec) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbSpec) ProtoMessage()               {}
func (*TlsProxyCbSpec) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{1} }

func (m *TlsProxyCbSpec) GetKeyOrHandle() *TlsProxyCbKeyHandle {
	if m != nil {
		return m.KeyOrHandle
	}
	return nil
}

func (m *TlsProxyCbSpec) GetNicDecHead() uint32 {
	if m != nil {
		return m.NicDecHead
	}
	return 0
}

func (m *TlsProxyCbSpec) GetNicDecTail() uint32 {
	if m != nil {
		return m.NicDecTail
	}
	return 0
}

func (m *TlsProxyCbSpec) GetCommand() uint32 {
	if m != nil {
		return m.Command
	}
	return 0
}

func (m *TlsProxyCbSpec) GetDebugDol() uint32 {
	if m != nil {
		return m.DebugDol
	}
	return 0
}

func (m *TlsProxyCbSpec) GetSerqPi() uint32 {
	if m != nil {
		return m.SerqPi
	}
	return 0
}

func (m *TlsProxyCbSpec) GetSerqCi() uint32 {
	if m != nil {
		return m.SerqCi
	}
	return 0
}

func (m *TlsProxyCbSpec) GetBsqPi() uint32 {
	if m != nil {
		return m.BsqPi
	}
	return 0
}

func (m *TlsProxyCbSpec) GetBsqCi() uint32 {
	if m != nil {
		return m.BsqCi
	}
	return 0
}

func (m *TlsProxyCbSpec) GetCryptoKeyIdx() uint32 {
	if m != nil {
		return m.CryptoKeyIdx
	}
	return 0
}

func (m *TlsProxyCbSpec) GetSerqBase() uint32 {
	if m != nil {
		return m.SerqBase
	}
	return 0
}

func (m *TlsProxyCbSpec) GetSesqBase() uint32 {
	if m != nil {
		return m.SesqBase
	}
	return 0
}

func (m *TlsProxyCbSpec) GetTnmdrAlloc() uint64 {
	if m != nil {
		return m.TnmdrAlloc
	}
	return 0
}

func (m *TlsProxyCbSpec) GetTnmprAlloc() uint64 {
	if m != nil {
		return m.TnmprAlloc
	}
	return 0
}

func (m *TlsProxyCbSpec) GetRnmdrFree() uint64 {
	if m != nil {
		return m.RnmdrFree
	}
	return 0
}

func (m *TlsProxyCbSpec) GetRnmprFree() uint64 {
	if m != nil {
		return m.RnmprFree
	}
	return 0
}

func (m *TlsProxyCbSpec) GetEncRequests() uint64 {
	if m != nil {
		return m.EncRequests
	}
	return 0
}

func (m *TlsProxyCbSpec) GetEncCompletions() uint64 {
	if m != nil {
		return m.EncCompletions
	}
	return 0
}

func (m *TlsProxyCbSpec) GetEncFailures() uint64 {
	if m != nil {
		return m.EncFailures
	}
	return 0
}

func (m *TlsProxyCbSpec) GetDecRequests() uint64 {
	if m != nil {
		return m.DecRequests
	}
	return 0
}

func (m *TlsProxyCbSpec) GetDecCompletions() uint64 {
	if m != nil {
		return m.DecCompletions
	}
	return 0
}

func (m *TlsProxyCbSpec) GetDecFailures() uint64 {
	if m != nil {
		return m.DecFailures
	}
	return 0
}

func (m *TlsProxyCbSpec) GetSalt() uint32 {
	if m != nil {
		return m.Salt
	}
	return 0
}

func (m *TlsProxyCbSpec) GetExplicitIv() uint64 {
	if m != nil {
		return m.ExplicitIv
	}
	return 0
}

func (m *TlsProxyCbSpec) GetPreDebugStage0_7Thread() uint32 {
	if m != nil {
		return m.PreDebugStage0_7Thread
	}
	return 0
}

func (m *TlsProxyCbSpec) GetPostDebugStage0_7Thread() uint32 {
	if m != nil {
		return m.PostDebugStage0_7Thread
	}
	return 0
}

func (m *TlsProxyCbSpec) GetIsDecryptFlow() bool {
	if m != nil {
		return m.IsDecryptFlow
	}
	return false
}

func (m *TlsProxyCbSpec) GetOtherFid() uint32 {
	if m != nil {
		return m.OtherFid
	}
	return 0
}

func (m *TlsProxyCbSpec) GetL7ProxyType() AppRedirType {
	if m != nil {
		return m.L7ProxyType
	}
	return AppRedirType_APP_REDIR_TYPE_NONE
}

func (m *TlsProxyCbSpec) GetCryptoHmacKeyIdx() uint32 {
	if m != nil {
		return m.CryptoHmacKeyIdx
	}
	return 0
}

func (m *TlsProxyCbSpec) GetMacRequests() uint64 {
	if m != nil {
		return m.MacRequests
	}
	return 0
}

func (m *TlsProxyCbSpec) GetMacCompletions() uint64 {
	if m != nil {
		return m.MacCompletions
	}
	return 0
}

func (m *TlsProxyCbSpec) GetMacFailures() uint64 {
	if m != nil {
		return m.MacFailures
	}
	return 0
}

func (m *TlsProxyCbSpec) GetCpuId() uint32 {
	if m != nil {
		return m.CpuId
	}
	return 0
}

// TlsProxyCbRequestMsg is batched add or modify tls_proxy_cb request
type TlsProxyCbRequestMsg struct {
	Request []*TlsProxyCbSpec `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *TlsProxyCbRequestMsg) Reset()                    { *m = TlsProxyCbRequestMsg{} }
func (m *TlsProxyCbRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbRequestMsg) ProtoMessage()               {}
func (*TlsProxyCbRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{2} }

func (m *TlsProxyCbRequestMsg) GetRequest() []*TlsProxyCbSpec {
	if m != nil {
		return m.Request
	}
	return nil
}

// TlsProxyCbStatus is the operational status of a given tls_proxy_cb
type TlsProxyCbStatus struct {
	TlsProxyCbHandle uint64 `protobuf:"fixed64,1,opt,name=tls_proxy_cb_handle,json=tlsProxyCbHandle" json:"tls_proxy_cb_handle,omitempty"`
}

func (m *TlsProxyCbStatus) Reset()                    { *m = TlsProxyCbStatus{} }
func (m *TlsProxyCbStatus) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbStatus) ProtoMessage()               {}
func (*TlsProxyCbStatus) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{3} }

func (m *TlsProxyCbStatus) GetTlsProxyCbHandle() uint64 {
	if m != nil {
		return m.TlsProxyCbHandle
	}
	return 0
}

// TlsProxyCbResponse is response to TlsProxyCbSpec
type TlsProxyCbResponse struct {
	ApiStatus        ApiStatus         `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	TlsProxyCbStatus *TlsProxyCbStatus `protobuf:"bytes,2,opt,name=tls_proxy_cb_status,json=tlsProxyCbStatus" json:"tls_proxy_cb_status,omitempty"`
}

func (m *TlsProxyCbResponse) Reset()                    { *m = TlsProxyCbResponse{} }
func (m *TlsProxyCbResponse) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbResponse) ProtoMessage()               {}
func (*TlsProxyCbResponse) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{4} }

func (m *TlsProxyCbResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *TlsProxyCbResponse) GetTlsProxyCbStatus() *TlsProxyCbStatus {
	if m != nil {
		return m.TlsProxyCbStatus
	}
	return nil
}

// TlsProxyCbResponseMsg is batched response to TlsProxyCbRequestMsg
type TlsProxyCbResponseMsg struct {
	Response []*TlsProxyCbResponse `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *TlsProxyCbResponseMsg) Reset()                    { *m = TlsProxyCbResponseMsg{} }
func (m *TlsProxyCbResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbResponseMsg) ProtoMessage()               {}
func (*TlsProxyCbResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{5} }

func (m *TlsProxyCbResponseMsg) GetResponse() []*TlsProxyCbResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

// TlsProxyCbDeleteRequest is used to delete a tls_proxy_cb
type TlsProxyCbDeleteRequest struct {
	KeyOrHandle *TlsProxyCbKeyHandle `protobuf:"bytes,1,opt,name=key_or_handle,json=keyOrHandle" json:"key_or_handle,omitempty"`
}

func (m *TlsProxyCbDeleteRequest) Reset()                    { *m = TlsProxyCbDeleteRequest{} }
func (m *TlsProxyCbDeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbDeleteRequest) ProtoMessage()               {}
func (*TlsProxyCbDeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{6} }

func (m *TlsProxyCbDeleteRequest) GetKeyOrHandle() *TlsProxyCbKeyHandle {
	if m != nil {
		return m.KeyOrHandle
	}
	return nil
}

// TlsProxyCbDeleteRequestMsg is used to delete a batch of tls_proxy_cbs
type TlsProxyCbDeleteRequestMsg struct {
	Request []*TlsProxyCbDeleteRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *TlsProxyCbDeleteRequestMsg) Reset()                    { *m = TlsProxyCbDeleteRequestMsg{} }
func (m *TlsProxyCbDeleteRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbDeleteRequestMsg) ProtoMessage()               {}
func (*TlsProxyCbDeleteRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{7} }

func (m *TlsProxyCbDeleteRequestMsg) GetRequest() []*TlsProxyCbDeleteRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// TlsProxyCbDeleteResponseMsg is batched response to TlsProxyCbDeleteRequestMsg
type TlsProxyCbDeleteResponseMsg struct {
	ApiStatus []ApiStatus `protobuf:"varint,1,rep,packed,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
}

func (m *TlsProxyCbDeleteResponseMsg) Reset()                    { *m = TlsProxyCbDeleteResponseMsg{} }
func (m *TlsProxyCbDeleteResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbDeleteResponseMsg) ProtoMessage()               {}
func (*TlsProxyCbDeleteResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{8} }

func (m *TlsProxyCbDeleteResponseMsg) GetApiStatus() []ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return nil
}

// TlsProxyCbGetRequest is used to get information about a tls_proxy_cb
type TlsProxyCbGetRequest struct {
	KeyOrHandle *TlsProxyCbKeyHandle `protobuf:"bytes,1,opt,name=key_or_handle,json=keyOrHandle" json:"key_or_handle,omitempty"`
}

func (m *TlsProxyCbGetRequest) Reset()                    { *m = TlsProxyCbGetRequest{} }
func (m *TlsProxyCbGetRequest) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbGetRequest) ProtoMessage()               {}
func (*TlsProxyCbGetRequest) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{9} }

func (m *TlsProxyCbGetRequest) GetKeyOrHandle() *TlsProxyCbKeyHandle {
	if m != nil {
		return m.KeyOrHandle
	}
	return nil
}

// TlsProxyCbGetRequestMsg is batched GET requests for tls_proxy_cbs
type TlsProxyCbGetRequestMsg struct {
	Request []*TlsProxyCbGetRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *TlsProxyCbGetRequestMsg) Reset()                    { *m = TlsProxyCbGetRequestMsg{} }
func (m *TlsProxyCbGetRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbGetRequestMsg) ProtoMessage()               {}
func (*TlsProxyCbGetRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{10} }

func (m *TlsProxyCbGetRequestMsg) GetRequest() []*TlsProxyCbGetRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// TlsProxyCbStats is the statistics object for each tls_proxy_cb
type TlsProxyCbStats struct {
}

func (m *TlsProxyCbStats) Reset()                    { *m = TlsProxyCbStats{} }
func (m *TlsProxyCbStats) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbStats) ProtoMessage()               {}
func (*TlsProxyCbStats) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{11} }

// TlsProxyCbGetResponse captures all the information about a tls_proxy_cb
// only if api_status indicates success, other fields are valid
type TlsProxyCbGetResponse struct {
	ApiStatus ApiStatus         `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	Spec      *TlsProxyCbSpec   `protobuf:"bytes,2,opt,name=spec" json:"spec,omitempty"`
	Status    *TlsProxyCbStatus `protobuf:"bytes,3,opt,name=status" json:"status,omitempty"`
	Stats     *TlsProxyCbStats  `protobuf:"bytes,4,opt,name=stats" json:"stats,omitempty"`
}

func (m *TlsProxyCbGetResponse) Reset()                    { *m = TlsProxyCbGetResponse{} }
func (m *TlsProxyCbGetResponse) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbGetResponse) ProtoMessage()               {}
func (*TlsProxyCbGetResponse) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{12} }

func (m *TlsProxyCbGetResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *TlsProxyCbGetResponse) GetSpec() *TlsProxyCbSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *TlsProxyCbGetResponse) GetStatus() *TlsProxyCbStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *TlsProxyCbGetResponse) GetStats() *TlsProxyCbStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

// TlsProxyCbGetResponseMsg is batched response to TlsProxyCbGetRequestMsg
type TlsProxyCbGetResponseMsg struct {
	Response []*TlsProxyCbGetResponse `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *TlsProxyCbGetResponseMsg) Reset()                    { *m = TlsProxyCbGetResponseMsg{} }
func (m *TlsProxyCbGetResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*TlsProxyCbGetResponseMsg) ProtoMessage()               {}
func (*TlsProxyCbGetResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor40, []int{13} }

func (m *TlsProxyCbGetResponseMsg) GetResponse() []*TlsProxyCbGetResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*TlsProxyCbKeyHandle)(nil), "halproto.TlsProxyCbKeyHandle")
	proto.RegisterType((*TlsProxyCbSpec)(nil), "halproto.TlsProxyCbSpec")
	proto.RegisterType((*TlsProxyCbRequestMsg)(nil), "halproto.TlsProxyCbRequestMsg")
	proto.RegisterType((*TlsProxyCbStatus)(nil), "halproto.TlsProxyCbStatus")
	proto.RegisterType((*TlsProxyCbResponse)(nil), "halproto.TlsProxyCbResponse")
	proto.RegisterType((*TlsProxyCbResponseMsg)(nil), "halproto.TlsProxyCbResponseMsg")
	proto.RegisterType((*TlsProxyCbDeleteRequest)(nil), "halproto.TlsProxyCbDeleteRequest")
	proto.RegisterType((*TlsProxyCbDeleteRequestMsg)(nil), "halproto.TlsProxyCbDeleteRequestMsg")
	proto.RegisterType((*TlsProxyCbDeleteResponseMsg)(nil), "halproto.TlsProxyCbDeleteResponseMsg")
	proto.RegisterType((*TlsProxyCbGetRequest)(nil), "halproto.TlsProxyCbGetRequest")
	proto.RegisterType((*TlsProxyCbGetRequestMsg)(nil), "halproto.TlsProxyCbGetRequestMsg")
	proto.RegisterType((*TlsProxyCbStats)(nil), "halproto.TlsProxyCbStats")
	proto.RegisterType((*TlsProxyCbGetResponse)(nil), "halproto.TlsProxyCbGetResponse")
	proto.RegisterType((*TlsProxyCbGetResponseMsg)(nil), "halproto.TlsProxyCbGetResponseMsg")
}

func init() { proto.RegisterFile("tls_proxy_cb2.proto", fileDescriptor40) }

var fileDescriptor40 = []byte{
	// 1078 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x5b, 0x53, 0xdb, 0x46,
	0x14, 0x8e, 0xb8, 0x18, 0x7c, 0x7c, 0x81, 0x2c, 0x21, 0x2c, 0xa6, 0x04, 0xa3, 0xb6, 0xa9, 0x5f,
	0x42, 0x32, 0xe6, 0x81, 0xb6, 0xd3, 0x17, 0x2e, 0x43, 0xa1, 0x4c, 0xda, 0x8c, 0x42, 0xa7, 0x1d,
	0x66, 0x3a, 0x1a, 0x79, 0x75, 0xc0, 0x3b, 0x91, 0xad, 0x8d, 0x56, 0x4e, 0xf0, 0x5b, 0xff, 0x42,
	0x7f, 0x57, 0xff, 0x43, 0xff, 0x4a, 0x3b, 0x7b, 0xb1, 0x24, 0x1b, 0x83, 0x67, 0x3a, 0x3c, 0x21,
	0x9d, 0xef, 0x3b, 0xdf, 0x39, 0x3a, 0x37, 0x0c, 0x6b, 0x69, 0x24, 0x7d, 0x91, 0xc4, 0xb7, 0x43,
	0x9f, 0x75, 0xda, 0x7b, 0x22, 0x89, 0xd3, 0x98, 0x94, 0x33, 0x63, 0xa3, 0x92, 0x0e, 0x05, 0x4a,
	0x63, 0x77, 0xff, 0x74, 0x60, 0xed, 0x32, 0x92, 0xef, 0x14, 0x72, 0xdc, 0xb9, 0xc0, 0xe1, 0x59,
	0xd0, 0x0f, 0x23, 0x24, 0xdf, 0xc0, 0x4a, 0x51, 0xc6, 0xe7, 0x21, 0x75, 0x9a, 0x4e, 0xab, 0x76,
	0xf6, 0xc4, 0xab, 0xa6, 0x19, 0xff, 0x3c, 0x24, 0xaf, 0xc7, 0xe3, 0xf9, 0x5d, 0xed, 0x4f, 0xe7,
	0x9a, 0x4e, 0xab, 0x74, 0xf6, 0xc4, 0x5b, 0xcd, 0xc9, 0x46, 0xf9, 0x68, 0x05, 0x6a, 0x1f, 0x70,
	0xe8, 0xc7, 0x89, 0xa5, 0xba, 0x7f, 0x97, 0xa1, 0x9e, 0xa7, 0xf0, 0x5e, 0x20, 0x23, 0x47, 0x13,
	0x1c, 0x1d, 0xbb, 0xd2, 0x7e, 0xb1, 0x97, 0x85, 0xda, 0x9b, 0x92, 0xb4, 0x57, 0xf9, 0x80, 0xc3,
	0x5f, 0x12, 0xfb, 0x05, 0x4d, 0xa8, 0xf6, 0x39, 0xf3, 0x43, 0x64, 0x7e, 0x17, 0x83, 0x50, 0x67,
	0x54, 0xf3, 0xa0, 0xcf, 0xd9, 0x09, 0xb2, 0x33, 0x0c, 0xc2, 0x22, 0x23, 0x0d, 0x78, 0x44, 0xe7,
	0x8b, 0x8c, 0xcb, 0x80, 0x47, 0x84, 0xc2, 0x12, 0x8b, 0x7b, 0xbd, 0xa0, 0x1f, 0xd2, 0x05, 0x0d,
	0x8e, 0x5e, 0xc9, 0x16, 0x94, 0x43, 0xec, 0x0c, 0x6e, 0xfc, 0x30, 0x8e, 0xe8, 0xa2, 0xc6, 0x96,
	0xb5, 0xe1, 0x24, 0x8e, 0xc8, 0x06, 0x2c, 0x49, 0x4c, 0x3e, 0xfa, 0x82, 0xd3, 0x92, 0x86, 0x4a,
	0xea, 0xf5, 0x1d, 0xcf, 0x00, 0xc6, 0xe9, 0x52, 0x0e, 0x1c, 0x73, 0xb2, 0x0e, 0xa5, 0x8e, 0xd4,
	0x0e, 0xcb, 0xda, 0xbe, 0xd8, 0x91, 0x8a, 0x6f, 0xcd, 0x8c, 0xd3, 0x72, 0x66, 0x3e, 0xe6, 0xe4,
	0x2b, 0xa8, 0xb3, 0x64, 0x28, 0xd2, 0xd8, 0x57, 0x55, 0xe2, 0xe1, 0x2d, 0x05, 0x0d, 0x57, 0x8d,
	0xf5, 0x02, 0x87, 0xe7, 0xe1, 0xad, 0x4a, 0x51, 0x07, 0xeb, 0x04, 0x12, 0x69, 0xc5, 0xa4, 0xa8,
	0x0c, 0x47, 0x81, 0x44, 0x03, 0x4a, 0x0b, 0x56, 0x47, 0xa0, 0x34, 0xe0, 0x0e, 0x54, 0xd2, 0x7e,
	0x2f, 0x4c, 0xfc, 0x20, 0x8a, 0x62, 0x46, 0x6b, 0xaa, 0x97, 0x1e, 0x68, 0xd3, 0xa1, 0xb2, 0x58,
	0x82, 0x18, 0x11, 0xea, 0x19, 0x41, 0x58, 0xc2, 0x36, 0x40, 0xa2, 0x15, 0xae, 0x13, 0x44, 0xba,
	0xa2, 0xf1, 0xb2, 0xb6, 0x9c, 0x26, 0x88, 0x16, 0x16, 0x16, 0x5e, 0xcd, 0x60, 0x61, 0xe0, 0x5d,
	0xa8, 0x62, 0x9f, 0xf9, 0x09, 0x7e, 0x1c, 0xa0, 0x4c, 0x25, 0x7d, 0xaa, 0x09, 0x15, 0xec, 0x33,
	0xcf, 0x9a, 0xd4, 0x7c, 0x2a, 0x0a, 0x8b, 0x7b, 0x22, 0xc2, 0x94, 0xc7, 0x7d, 0x49, 0x89, 0x66,
	0xd5, 0xb1, 0xcf, 0x8e, 0x73, 0xeb, 0x48, 0xeb, 0x3a, 0xe0, 0xd1, 0x20, 0x41, 0x49, 0xd7, 0x32,
	0xad, 0x53, 0x6b, 0x52, 0x14, 0x35, 0x03, 0x59, 0xb8, 0x67, 0x86, 0x12, 0xe2, 0x58, 0x38, 0x45,
	0x29, 0x86, 0x5b, 0x37, 0xe1, 0x42, 0x9c, 0x0c, 0xa7, 0x88, 0x59, 0xb8, 0xe7, 0x99, 0x56, 0x16,
	0x8e, 0xc0, 0x82, 0x0c, 0xa2, 0x94, 0x6e, 0xe8, 0xaa, 0xeb, 0x67, 0x55, 0x50, 0xbc, 0x15, 0x11,
	0x67, 0x3c, 0xf5, 0xf9, 0x27, 0x4a, 0x9b, 0x4e, 0x6b, 0xc1, 0x83, 0x91, 0xe9, 0xfc, 0x13, 0xf9,
	0x16, 0x36, 0x45, 0x82, 0xbe, 0x99, 0x39, 0x99, 0x06, 0x37, 0xf8, 0xc6, 0x3f, 0xf0, 0xd3, 0x6e,
	0xa2, 0x46, 0x7b, 0x53, 0x2b, 0xad, 0x8b, 0x04, 0x4f, 0x14, 0xfe, 0x5e, 0xc3, 0x07, 0x97, 0x1a,
	0x24, 0xdf, 0x43, 0x43, 0xc4, 0x32, 0xbd, 0xc7, 0xb5, 0xa1, 0x5d, 0x9f, 0x2b, 0xc6, 0x14, 0xdf,
	0x97, 0xb0, 0xc2, 0xa5, 0x5a, 0x10, 0x35, 0x57, 0xfe, 0x75, 0x14, 0x7f, 0xa6, 0x5b, 0x4d, 0xa7,
	0xb5, 0xec, 0xd5, 0xb8, 0x3c, 0x31, 0xd6, 0xd3, 0x28, 0xfe, 0xac, 0xa6, 0x29, 0x4e, 0xbb, 0x98,
	0xf8, 0xd7, 0x3c, 0xa4, 0x5f, 0x98, 0x69, 0xd2, 0x86, 0x53, 0x1e, 0x92, 0x03, 0xa8, 0x45, 0x07,
	0xf6, 0x40, 0xa8, 0xd3, 0x43, 0xb7, 0x9b, 0x4e, 0xab, 0xde, 0x5e, 0xdb, 0x33, 0x77, 0xe8, 0x50,
	0x08, 0x0f, 0x43, 0x9e, 0x5c, 0x0e, 0x05, 0x7a, 0x95, 0xe8, 0x40, 0x6f, 0xb5, 0x7a, 0x21, 0xaf,
	0x60, 0xcd, 0x8e, 0x79, 0xb7, 0x17, 0xb0, 0x6c, 0xd6, 0x5f, 0x68, 0xfd, 0x55, 0x03, 0x9d, 0xf5,
	0x02, 0x66, 0xe7, 0x7d, 0x17, 0xaa, 0x8a, 0x96, 0xb5, 0x71, 0xc7, 0x94, 0xbe, 0x17, 0x8c, 0xb5,
	0x51, 0x51, 0x8a, 0x6d, 0x6c, 0x9a, 0x36, 0xf6, 0x82, 0xc9, 0x36, 0x2a, 0x62, 0xd6, 0xc6, 0xdd,
	0x4c, 0x2b, 0x6b, 0xe3, 0x3a, 0x94, 0x98, 0x18, 0xa8, 0xc3, 0xe8, 0x9a, 0xdd, 0x64, 0x62, 0x70,
	0x1e, 0xba, 0x17, 0xf0, 0x2c, 0x3f, 0x4d, 0x36, 0xf0, 0x5b, 0x79, 0x43, 0xf6, 0x61, 0xc9, 0x66,
	0x46, 0x9d, 0xe6, 0x7c, 0xab, 0xd2, 0xde, 0x9c, 0x7a, 0xcc, 0xd4, 0xf9, 0xf3, 0x46, 0x4c, 0xf7,
	0x10, 0x56, 0x0b, 0x50, 0x1a, 0xa4, 0x03, 0xa9, 0xaa, 0x32, 0xed, 0xe0, 0x3a, 0x3a, 0xc3, 0x3b,
	0xe7, 0xd6, 0xfd, 0xcb, 0x01, 0x52, 0x4c, 0x48, 0x8a, 0xb8, 0x2f, 0x91, 0xbc, 0x06, 0x08, 0x04,
	0x57, 0xe3, 0x90, 0x0e, 0xa4, 0x76, 0xae, 0xb7, 0x57, 0xb3, 0x8e, 0x70, 0x13, 0xcb, 0x2b, 0x07,
	0xa3, 0x47, 0xf2, 0xd3, 0x44, 0x58, 0xeb, 0x39, 0xa7, 0x0f, 0xf3, 0xd6, 0xf4, 0x6f, 0x31, 0x22,
	0x85, 0x9c, 0x8c, 0xc5, 0xf5, 0x60, 0xfd, 0x6e, 0x4a, 0xaa, 0x48, 0xdf, 0xc1, 0x72, 0x62, 0x5f,
	0x6d, 0x95, 0xb6, 0xa7, 0x2a, 0x8f, 0x7c, 0xbc, 0x8c, 0xee, 0xfe, 0x01, 0x1b, 0x39, 0x7e, 0x82,
	0x11, 0xa6, 0x68, 0xab, 0xff, 0x18, 0xff, 0x4d, 0xdc, 0x2b, 0x68, 0xdc, 0x23, 0xaf, 0xf2, 0xfe,
	0x61, 0xb2, 0xb9, 0xee, 0x54, 0xed, 0x31, 0xbf, 0xbc, 0xcb, 0x3f, 0xc3, 0xd6, 0x5d, 0x4e, 0x5e,
	0x94, 0xc9, 0x56, 0xcd, 0xcf, 0x68, 0x95, 0x7b, 0x55, 0x1c, 0xc1, 0x1f, 0x31, 0x7d, 0xcc, 0x3a,
	0x5c, 0x16, 0xcb, 0x9c, 0x6b, 0x9b, 0xe6, 0x4d, 0x14, 0x61, 0x67, 0xaa, 0x70, 0xee, 0x94, 0x57,
	0xe0, 0x29, 0xac, 0x8c, 0x8f, 0x8d, 0x74, 0xff, 0x71, 0x8a, 0x43, 0xa2, 0x9d, 0xfe, 0xef, 0xe8,
	0xbe, 0x82, 0x05, 0x29, 0x90, 0xd9, 0x59, 0x7d, 0x60, 0xef, 0x34, 0x8d, 0xec, 0x43, 0xc9, 0x6a,
	0xcf, 0xcf, 0x1e, 0x6e, 0x4b, 0x25, 0x6f, 0x60, 0x51, 0x3d, 0x49, 0xfd, 0x3b, 0xa1, 0xd2, 0x6e,
	0xdc, 0xeb, 0x23, 0x3d, 0x43, 0x74, 0x7f, 0x07, 0x3a, 0xf5, 0xfb, 0xcc, 0x3c, 0x4d, 0xee, 0x41,
	0xf3, 0xfe, 0x5a, 0x4e, 0xae, 0x42, 0xfb, 0xdf, 0x39, 0x80, 0x9c, 0x43, 0x7e, 0x2b, 0x1e, 0x91,
	0xe3, 0x04, 0x83, 0x14, 0xc9, 0xce, 0x3d, 0x6b, 0x35, 0x6a, 0x66, 0xa3, 0xf9, 0xe0, 0xde, 0xbd,
	0x95, 0x37, 0xee, 0x93, 0x71, 0xe1, 0x5f, 0x45, 0xf8, 0x68, 0xc2, 0xac, 0x28, 0x6c, 0x16, 0x82,
	0x7c, 0x3d, 0x7b, 0xa3, 0x94, 0xfc, 0xcb, 0x07, 0x69, 0xc5, 0x20, 0x57, 0x50, 0x1b, 0x2b, 0x24,
	0x71, 0x67, 0x8c, 0xab, 0x92, 0xff, 0x72, 0x56, 0x1b, 0xb4, 0xf6, 0x11, 0x5c, 0x2d, 0x77, 0x83,
	0x48, 0xff, 0xc2, 0xee, 0x94, 0xf4, 0x9f, 0xfd, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8f, 0xa2,
	0xf9, 0x04, 0x97, 0x0b, 0x00, 0x00,
}
