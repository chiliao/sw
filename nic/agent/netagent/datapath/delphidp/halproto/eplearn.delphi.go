// Code generated by protoc-gen-go. DO NOT EDIT.
// source: eplearn.proto

package halproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EpLearnType int32

const (
	EpLearnType_LEARN_TYPE_NONE   EpLearnType = 0
	EpLearnType_LEARN_TYPE_LOCAL  EpLearnType = 1
	EpLearnType_LEARN_TYPE_GLOBAL EpLearnType = 2
	EpLearnType_LEARN_TYPE_ANY    EpLearnType = 3
)

var EpLearnType_name = map[int32]string{
	0: "LEARN_TYPE_NONE",
	1: "LEARN_TYPE_LOCAL",
	2: "LEARN_TYPE_GLOBAL",
	3: "LEARN_TYPE_ANY",
}
var EpLearnType_value = map[string]int32{
	"LEARN_TYPE_NONE":   0,
	"LEARN_TYPE_LOCAL":  1,
	"LEARN_TYPE_GLOBAL": 2,
	"LEARN_TYPE_ANY":    3,
}

func (x EpLearnType) String() string {
	return proto.EnumName(EpLearnType_name, int32(x))
}
func (EpLearnType) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

type DhcpTransactionState int32

const (
	DhcpTransactionState_STATE_NONE       DhcpTransactionState = 0
	DhcpTransactionState_STATE_SELECTING  DhcpTransactionState = 1
	DhcpTransactionState_STATE_REQUESTING DhcpTransactionState = 2
	DhcpTransactionState_STATE_BOUND      DhcpTransactionState = 3
	DhcpTransactionState_STATE_RENEWING   DhcpTransactionState = 4
	DhcpTransactionState_STATE_REBINDING  DhcpTransactionState = 5
)

var DhcpTransactionState_name = map[int32]string{
	0: "STATE_NONE",
	1: "STATE_SELECTING",
	2: "STATE_REQUESTING",
	3: "STATE_BOUND",
	4: "STATE_RENEWING",
	5: "STATE_REBINDING",
}
var DhcpTransactionState_value = map[string]int32{
	"STATE_NONE":       0,
	"STATE_SELECTING":  1,
	"STATE_REQUESTING": 2,
	"STATE_BOUND":      3,
	"STATE_RENEWING":   4,
	"STATE_REBINDING":  5,
}

func (x DhcpTransactionState) String() string {
	return proto.EnumName(DhcpTransactionState_name, int32(x))
}
func (DhcpTransactionState) EnumDescriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

// DhcpSpec Cfg
type EplearnDhcpCfg struct {
	TrustedServers []*IPAddress `protobuf:"bytes,1,rep,name=trusted_servers,json=trustedServers" json:"trusted_servers,omitempty"`
}

func (m *EplearnDhcpCfg) Reset()                    { *m = EplearnDhcpCfg{} }
func (m *EplearnDhcpCfg) String() string            { return proto.CompactTextString(m) }
func (*EplearnDhcpCfg) ProtoMessage()               {}
func (*EplearnDhcpCfg) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *EplearnDhcpCfg) GetTrustedServers() []*IPAddress {
	if m != nil {
		return m.TrustedServers
	}
	return nil
}

// Arp/Ndp Cfg
type EplearnArpCfg struct {
	EntryTimeout uint32 `protobuf:"varint,1,opt,name=entry_timeout,json=entryTimeout" json:"entry_timeout,omitempty"`
	ProbeEnabled bool   `protobuf:"varint,2,opt,name=probe_enabled,json=probeEnabled" json:"probe_enabled,omitempty"`
}

func (m *EplearnArpCfg) Reset()                    { *m = EplearnArpCfg{} }
func (m *EplearnArpCfg) String() string            { return proto.CompactTextString(m) }
func (*EplearnArpCfg) ProtoMessage()               {}
func (*EplearnArpCfg) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *EplearnArpCfg) GetEntryTimeout() uint32 {
	if m != nil {
		return m.EntryTimeout
	}
	return 0
}

func (m *EplearnArpCfg) GetProbeEnabled() bool {
	if m != nil {
		return m.ProbeEnabled
	}
	return false
}

// Arp/Ndp Cfg
type EplearnDataPacketCfg struct {
	Enabled bool `protobuf:"varint,1,opt,name=enabled" json:"enabled,omitempty"`
}

func (m *EplearnDataPacketCfg) Reset()                    { *m = EplearnDataPacketCfg{} }
func (m *EplearnDataPacketCfg) String() string            { return proto.CompactTextString(m) }
func (*EplearnDataPacketCfg) ProtoMessage()               {}
func (*EplearnDataPacketCfg) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *EplearnDataPacketCfg) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

type EplearnCfg struct {
	LearnType           EpLearnType           `protobuf:"varint,1,opt,name=learn_type,json=learnType,enum=eplearn.EpLearnType" json:"learn_type,omitempty"`
	DropOnStaticMisatch bool                  `protobuf:"varint,2,opt,name=drop_on_static_misatch,json=dropOnStaticMisatch" json:"drop_on_static_misatch,omitempty"`
	Arp                 *EplearnArpCfg        `protobuf:"bytes,3,opt,name=arp" json:"arp,omitempty"`
	Dhcp                *EplearnDhcpCfg       `protobuf:"bytes,4,opt,name=dhcp" json:"dhcp,omitempty"`
	Dpkt                *EplearnDataPacketCfg `protobuf:"bytes,5,opt,name=dpkt" json:"dpkt,omitempty"`
}

func (m *EplearnCfg) Reset()                    { *m = EplearnCfg{} }
func (m *EplearnCfg) String() string            { return proto.CompactTextString(m) }
func (*EplearnCfg) ProtoMessage()               {}
func (*EplearnCfg) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *EplearnCfg) GetLearnType() EpLearnType {
	if m != nil {
		return m.LearnType
	}
	return EpLearnType_LEARN_TYPE_NONE
}

func (m *EplearnCfg) GetDropOnStaticMisatch() bool {
	if m != nil {
		return m.DropOnStaticMisatch
	}
	return false
}

func (m *EplearnCfg) GetArp() *EplearnArpCfg {
	if m != nil {
		return m.Arp
	}
	return nil
}

func (m *EplearnCfg) GetDhcp() *EplearnDhcpCfg {
	if m != nil {
		return m.Dhcp
	}
	return nil
}

func (m *EplearnCfg) GetDpkt() *EplearnDataPacketCfg {
	if m != nil {
		return m.Dpkt
	}
	return nil
}

type ArpStats struct {
	NumIpLearnt uint32 `protobuf:"varint,1,opt,name=num_ip_learnt,json=numIpLearnt" json:"num_ip_learnt,omitempty"`
	NumSpoofs   uint32 `protobuf:"varint,2,opt,name=num_spoofs,json=numSpoofs" json:"num_spoofs,omitempty"`
	NumDaiFails uint32 `protobuf:"varint,3,opt,name=num_dai_fails,json=numDaiFails" json:"num_dai_fails,omitempty"`
}

func (m *ArpStats) Reset()                    { *m = ArpStats{} }
func (m *ArpStats) String() string            { return proto.CompactTextString(m) }
func (*ArpStats) ProtoMessage()               {}
func (*ArpStats) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *ArpStats) GetNumIpLearnt() uint32 {
	if m != nil {
		return m.NumIpLearnt
	}
	return 0
}

func (m *ArpStats) GetNumSpoofs() uint32 {
	if m != nil {
		return m.NumSpoofs
	}
	return 0
}

func (m *ArpStats) GetNumDaiFails() uint32 {
	if m != nil {
		return m.NumDaiFails
	}
	return 0
}

type DhcpStats struct {
	NumIpLearnt uint32 `protobuf:"varint,1,opt,name=num_ip_learnt,json=numIpLearnt" json:"num_ip_learnt,omitempty"`
	NumNwFails  uint32 `protobuf:"varint,2,opt,name=num_nw_fails,json=numNwFails" json:"num_nw_fails,omitempty"`
}

func (m *DhcpStats) Reset()                    { *m = DhcpStats{} }
func (m *DhcpStats) String() string            { return proto.CompactTextString(m) }
func (*DhcpStats) ProtoMessage()               {}
func (*DhcpStats) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *DhcpStats) GetNumIpLearnt() uint32 {
	if m != nil {
		return m.NumIpLearnt
	}
	return 0
}

func (m *DhcpStats) GetNumNwFails() uint32 {
	if m != nil {
		return m.NumNwFails
	}
	return 0
}

type L2EplearnStats struct {
	L2SegmentKeyHandle *L2SegmentKeyHandle `protobuf:"bytes,1,opt,name=l2Segment_key_handle,json=l2SegmentKeyHandle" json:"l2Segment_key_handle,omitempty"`
	ArpStats           *ArpStats           `protobuf:"bytes,2,opt,name=arp_stats,json=arpStats" json:"arp_stats,omitempty"`
	DhcpStats          *DhcpStats          `protobuf:"bytes,3,opt,name=dhcp_stats,json=dhcpStats" json:"dhcp_stats,omitempty"`
}

func (m *L2EplearnStats) Reset()                    { *m = L2EplearnStats{} }
func (m *L2EplearnStats) String() string            { return proto.CompactTextString(m) }
func (*L2EplearnStats) ProtoMessage()               {}
func (*L2EplearnStats) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *L2EplearnStats) GetL2SegmentKeyHandle() *L2SegmentKeyHandle {
	if m != nil {
		return m.L2SegmentKeyHandle
	}
	return nil
}

func (m *L2EplearnStats) GetArpStats() *ArpStats {
	if m != nil {
		return m.ArpStats
	}
	return nil
}

func (m *L2EplearnStats) GetDhcpStats() *DhcpStats {
	if m != nil {
		return m.DhcpStats
	}
	return nil
}

type DhcpStatus struct {
	State         DhcpTransactionState `protobuf:"varint,1,opt,name=state,enum=eplearn.DhcpTransactionState" json:"state,omitempty"`
	Xid           uint32               `protobuf:"varint,2,opt,name=xid" json:"xid,omitempty"`
	RenewalTime   uint32               `protobuf:"varint,3,opt,name=renewal_time,json=renewalTime" json:"renewal_time,omitempty"`
	RebindingTime uint32               `protobuf:"varint,4,opt,name=rebinding_time,json=rebindingTime" json:"rebinding_time,omitempty"`
	IpAddr        *IPAddress           `protobuf:"bytes,5,opt,name=ip_addr,json=ipAddr" json:"ip_addr,omitempty"`
	SubnetMask    *IPAddress           `protobuf:"bytes,6,opt,name=subnet_mask,json=subnetMask" json:"subnet_mask,omitempty"`
	GatewayIp     *IPAddress           `protobuf:"bytes,7,opt,name=gateway_ip,json=gatewayIp" json:"gateway_ip,omitempty"`
}

func (m *DhcpStatus) Reset()                    { *m = DhcpStatus{} }
func (m *DhcpStatus) String() string            { return proto.CompactTextString(m) }
func (*DhcpStatus) ProtoMessage()               {}
func (*DhcpStatus) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

func (m *DhcpStatus) GetState() DhcpTransactionState {
	if m != nil {
		return m.State
	}
	return DhcpTransactionState_STATE_NONE
}

func (m *DhcpStatus) GetXid() uint32 {
	if m != nil {
		return m.Xid
	}
	return 0
}

func (m *DhcpStatus) GetRenewalTime() uint32 {
	if m != nil {
		return m.RenewalTime
	}
	return 0
}

func (m *DhcpStatus) GetRebindingTime() uint32 {
	if m != nil {
		return m.RebindingTime
	}
	return 0
}

func (m *DhcpStatus) GetIpAddr() *IPAddress {
	if m != nil {
		return m.IpAddr
	}
	return nil
}

func (m *DhcpStatus) GetSubnetMask() *IPAddress {
	if m != nil {
		return m.SubnetMask
	}
	return nil
}

func (m *DhcpStatus) GetGatewayIp() *IPAddress {
	if m != nil {
		return m.GatewayIp
	}
	return nil
}

type ArpStatus struct {
	EntryActive  bool   `protobuf:"varint,1,opt,name=entry_active,json=entryActive" json:"entry_active,omitempty"`
	EntryTimeout uint32 `protobuf:"varint,2,opt,name=entry_timeout,json=entryTimeout" json:"entry_timeout,omitempty"`
}

func (m *ArpStatus) Reset()                    { *m = ArpStatus{} }
func (m *ArpStatus) String() string            { return proto.CompactTextString(m) }
func (*ArpStatus) ProtoMessage()               {}
func (*ArpStatus) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *ArpStatus) GetEntryActive() bool {
	if m != nil {
		return m.EntryActive
	}
	return false
}

func (m *ArpStatus) GetEntryTimeout() uint32 {
	if m != nil {
		return m.EntryTimeout
	}
	return 0
}

type EplearnStats struct {
	NumDhcpTrans  uint32 `protobuf:"varint,1,opt,name=num_dhcp_trans,json=numDhcpTrans" json:"num_dhcp_trans,omitempty"`
	NumArpTrans   uint32 `protobuf:"varint,2,opt,name=num_arp_trans,json=numArpTrans" json:"num_arp_trans,omitempty"`
	NumDhcpErrors uint32 `protobuf:"varint,3,opt,name=num_dhcp_errors,json=numDhcpErrors" json:"num_dhcp_errors,omitempty"`
	NumArpErrors  uint32 `protobuf:"varint,4,opt,name=num_arp_errors,json=numArpErrors" json:"num_arp_errors,omitempty"`
}

func (m *EplearnStats) Reset()                    { *m = EplearnStats{} }
func (m *EplearnStats) String() string            { return proto.CompactTextString(m) }
func (*EplearnStats) ProtoMessage()               {}
func (*EplearnStats) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *EplearnStats) GetNumDhcpTrans() uint32 {
	if m != nil {
		return m.NumDhcpTrans
	}
	return 0
}

func (m *EplearnStats) GetNumArpTrans() uint32 {
	if m != nil {
		return m.NumArpTrans
	}
	return 0
}

func (m *EplearnStats) GetNumDhcpErrors() uint32 {
	if m != nil {
		return m.NumDhcpErrors
	}
	return 0
}

func (m *EplearnStats) GetNumArpErrors() uint32 {
	if m != nil {
		return m.NumArpErrors
	}
	return 0
}

type EplearnStatus struct {
	DhcpStatus *DhcpStatus `protobuf:"bytes,1,opt,name=dhcp_status,json=dhcpStatus" json:"dhcp_status,omitempty"`
	ArpStatus  *ArpStatus  `protobuf:"bytes,2,opt,name=arp_status,json=arpStatus" json:"arp_status,omitempty"`
}

func (m *EplearnStatus) Reset()                    { *m = EplearnStatus{} }
func (m *EplearnStatus) String() string            { return proto.CompactTextString(m) }
func (*EplearnStatus) ProtoMessage()               {}
func (*EplearnStatus) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *EplearnStatus) GetDhcpStatus() *DhcpStatus {
	if m != nil {
		return m.DhcpStatus
	}
	return nil
}

func (m *EplearnStatus) GetArpStatus() *ArpStatus {
	if m != nil {
		return m.ArpStatus
	}
	return nil
}

func init() {
	proto.RegisterType((*EplearnDhcpCfg)(nil), "halproto.EplearnDhcpCfg")
	proto.RegisterType((*EplearnArpCfg)(nil), "halproto.EplearnArpCfg")
	proto.RegisterType((*EplearnDataPacketCfg)(nil), "halproto.EplearnDataPacketCfg")
	proto.RegisterType((*EplearnCfg)(nil), "halproto.EplearnCfg")
	proto.RegisterType((*ArpStats)(nil), "halproto.ArpStats")
	proto.RegisterType((*DhcpStats)(nil), "halproto.DhcpStats")
	proto.RegisterType((*L2EplearnStats)(nil), "halproto.L2EplearnStats")
	proto.RegisterType((*DhcpStatus)(nil), "halproto.DhcpStatus")
	proto.RegisterType((*ArpStatus)(nil), "halproto.ArpStatus")
	proto.RegisterType((*EplearnStats)(nil), "halproto.EplearnStats")
	proto.RegisterType((*EplearnStatus)(nil), "halproto.EplearnStatus")
	proto.RegisterEnum("halproto.EpLearnType", EpLearnType_name, EpLearnType_value)
	proto.RegisterEnum("halproto.DhcpTransactionState", DhcpTransactionState_name, DhcpTransactionState_value)
}

func init() { proto.RegisterFile("eplearn.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 910 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x6e, 0x1a, 0x47,
	0x14, 0xc7, 0xbb, 0x80, 0x6d, 0x38, 0xcb, 0x57, 0xc6, 0x34, 0x45, 0x91, 0x2c, 0xd1, 0xed, 0x87,
	0x68, 0x2a, 0xd1, 0x1a, 0xf7, 0xa6, 0x97, 0x6b, 0x7b, 0x9b, 0xa0, 0x10, 0x70, 0x16, 0xa2, 0xca,
	0xbd, 0x59, 0x0d, 0xec, 0xd8, 0xac, 0x80, 0xdd, 0xe9, 0xcc, 0x6c, 0x1c, 0xee, 0x7a, 0xdf, 0xf7,
	0xe8, 0x55, 0x5f, 0xa1, 0xef, 0x56, 0xcd, 0xd7, 0x86, 0x84, 0x5a, 0xea, 0x15, 0xb3, 0xff, 0xf9,
	0xcd, 0x39, 0x33, 0xe7, 0xfc, 0x67, 0x80, 0x06, 0xa1, 0x1b, 0x82, 0x59, 0x3a, 0xa0, 0x2c, 0x13,
	0x19, 0x3a, 0x31, 0x9f, 0xcf, 0x5c, 0xb1, 0xa3, 0x84, 0x6b, 0xf5, 0x59, 0x75, 0xbd, 0xd2, 0x23,
	0xef, 0x15, 0x34, 0x03, 0x4d, 0x5c, 0xaf, 0x96, 0xf4, 0xea, 0xee, 0x1e, 0xfd, 0x0c, 0x2d, 0xc1,
	0x72, 0x2e, 0x48, 0x1c, 0x71, 0xc2, 0xde, 0x11, 0xc6, 0xbb, 0x4e, 0xaf, 0xdc, 0x77, 0x87, 0xed,
	0x81, 0x0e, 0x31, 0xba, 0xf1, 0xe3, 0x98, 0x11, 0xce, 0xc3, 0xa6, 0x01, 0x67, 0x9a, 0xf3, 0x6e,
	0xa1, 0x61, 0x82, 0xf9, 0x4c, 0xc5, 0xfa, 0x0a, 0x1a, 0x24, 0x15, 0x6c, 0x17, 0x89, 0x64, 0x4b,
	0xb2, 0x5c, 0x74, 0x9d, 0x9e, 0xd3, 0x6f, 0x84, 0x75, 0x25, 0xce, 0xb5, 0x26, 0x21, 0xca, 0xb2,
	0x05, 0x89, 0x48, 0x8a, 0x17, 0x1b, 0x12, 0x77, 0x4b, 0x3d, 0xa7, 0x5f, 0x0d, 0xeb, 0x4a, 0x0c,
	0xb4, 0xe6, 0xfd, 0x08, 0x1d, 0xbb, 0x4f, 0x2c, 0xf0, 0x0d, 0x5e, 0xae, 0x89, 0x90, 0x19, 0xba,
	0x70, 0x62, 0x97, 0x39, 0x6a, 0x99, 0xfd, 0xf4, 0xfe, 0x28, 0x01, 0x98, 0x25, 0x12, 0xbc, 0x00,
	0x50, 0xe3, 0x48, 0x1e, 0x42, 0xb1, 0xcd, 0x61, 0x67, 0x60, 0x8b, 0x15, 0xd0, 0xb1, 0xfc, 0x9d,
	0xef, 0x28, 0x09, 0x6b, 0x1b, 0x3b, 0x44, 0x17, 0xf0, 0x34, 0x66, 0x19, 0x8d, 0xb2, 0x34, 0xe2,
	0x02, 0x8b, 0x64, 0x19, 0x6d, 0x13, 0x8e, 0xc5, 0x72, 0x65, 0xf6, 0x78, 0x2a, 0x67, 0xa7, 0xe9,
	0x4c, 0xcd, 0xbd, 0xd6, 0x53, 0xa8, 0x0f, 0x65, 0xcc, 0x68, 0xb7, 0xdc, 0x73, 0xfa, 0xee, 0xf0,
	0xe9, 0x5e, 0x8a, 0xbd, 0xca, 0x84, 0x12, 0x41, 0xdf, 0x43, 0x25, 0x5e, 0x2d, 0x69, 0xb7, 0xa2,
	0xd0, 0x2f, 0x3e, 0x45, 0x4d, 0x47, 0x42, 0x05, 0xa1, 0x73, 0xa8, 0xc4, 0x74, 0x2d, 0xba, 0x47,
	0x0a, 0x3e, 0x3b, 0x80, 0xf7, 0xcb, 0x12, 0x2a, 0xd4, 0xfb, 0x1d, 0xaa, 0x3e, 0xa3, 0x72, 0x77,
	0x1c, 0x79, 0xd0, 0x48, 0xf3, 0x6d, 0x94, 0xd0, 0x48, 0xe1, 0xb6, 0x15, 0x6e, 0x9a, 0x6f, 0x47,
	0xfa, 0xf0, 0x02, 0x9d, 0x01, 0x48, 0x86, 0xd3, 0x2c, 0xbb, 0xe3, 0xea, 0x88, 0x8d, 0xb0, 0x96,
	0xe6, 0xdb, 0x99, 0x12, 0x6c, 0x88, 0x18, 0x27, 0xd1, 0x1d, 0x4e, 0x36, 0x5c, 0x1d, 0x51, 0x87,
	0xb8, 0xc6, 0xc9, 0x2f, 0x52, 0xf2, 0xde, 0x40, 0x4d, 0x6e, 0xfb, 0xff, 0xe7, 0xec, 0x41, 0x5d,
	0x32, 0xe9, 0x83, 0x89, 0xa9, 0xb3, 0xca, 0x7d, 0x4c, 0x1e, 0x74, 0xc8, 0x7f, 0x1c, 0x68, 0x8e,
	0x87, 0xe6, 0x98, 0x3a, 0xf0, 0x4b, 0xe8, 0x6c, 0x86, 0x33, 0x72, 0xbf, 0x25, 0xa9, 0x88, 0xd6,
	0x64, 0x17, 0xad, 0x70, 0x1a, 0x6f, 0x74, 0x5b, 0x65, 0xcd, 0xd7, 0xab, 0xc1, 0xd8, 0xce, 0xbf,
	0x22, 0xbb, 0x97, 0x6a, 0x36, 0x44, 0x9b, 0x03, 0x0d, 0x0d, 0xa0, 0x86, 0x19, 0x55, 0xdd, 0xd5,
	0xb9, 0xdd, 0xe1, 0x93, 0xa2, 0xb4, 0xb6, 0x78, 0x61, 0x15, 0xdb, 0x32, 0x9e, 0x03, 0xc8, 0x6e,
	0x98, 0x05, 0xba, 0xc7, 0xa8, 0x58, 0x50, 0x1c, 0x3d, 0xac, 0xc5, 0x76, 0xe8, 0xfd, 0x5d, 0x02,
	0xb0, 0x13, 0x39, 0x47, 0x17, 0x70, 0x24, 0x17, 0x5b, 0x0f, 0x9e, 0x7d, 0xb4, 0x78, 0xce, 0x70,
	0xca, 0xf1, 0x52, 0x24, 0x99, 0x3a, 0x29, 0x09, 0x35, 0x8b, 0xda, 0x50, 0x7e, 0x9f, 0xc4, 0xa6,
	0x38, 0x72, 0x88, 0xbe, 0x84, 0x3a, 0x23, 0x29, 0x79, 0xc0, 0x1b, 0x75, 0xb9, 0x6c, 0x2f, 0x8c,
	0x26, 0xef, 0x16, 0xfa, 0x06, 0x9a, 0x8c, 0x2c, 0x92, 0x34, 0x4e, 0xd2, 0x7b, 0x0d, 0x55, 0x14,
	0xd4, 0x28, 0x54, 0x85, 0x7d, 0x07, 0x27, 0x09, 0x8d, 0x70, 0x1c, 0x33, 0xe3, 0xad, 0xc3, 0x8b,
	0x7e, 0x9c, 0x50, 0x39, 0x44, 0xe7, 0xe0, 0xf2, 0x7c, 0x91, 0x12, 0x11, 0x6d, 0x31, 0x5f, 0x77,
	0x8f, 0x1f, 0xc1, 0x41, 0x43, 0xaf, 0x31, 0x5f, 0xa3, 0x1f, 0x00, 0xee, 0xb1, 0x20, 0x0f, 0x78,
	0x17, 0x25, 0xb4, 0x7b, 0xf2, 0xc8, 0x8a, 0x9a, 0x61, 0x46, 0xd4, 0x9b, 0x41, 0xcd, 0xd4, 0x3d,
	0xe7, 0xf2, 0x94, 0xfa, 0x01, 0x91, 0x35, 0x79, 0x47, 0xcc, 0x1d, 0x77, 0x95, 0xe6, 0x2b, 0xe9,
	0xf0, 0x8d, 0x29, 0x1d, 0xbe, 0x31, 0xde, 0x5f, 0x0e, 0xd4, 0x3f, 0x72, 0xd0, 0xd7, 0xd0, 0x54,
	0x5e, 0x96, 0xbd, 0x14, 0xb2, 0xe8, 0xf6, 0x69, 0x92, 0x66, 0xb6, 0x8d, 0xb0, 0x06, 0x96, 0x0e,
	0xd1, 0x50, 0xa9, 0x30, 0xb0, 0xcf, 0x0c, 0xf3, 0x2d, 0xb4, 0x8a, 0x48, 0x84, 0xb1, 0x8c, 0xd9,
	0x7b, 0xd1, 0x30, 0xa1, 0x02, 0x25, 0xda, 0x8c, 0x32, 0x96, 0xc1, 0x2a, 0x45, 0x46, 0x9f, 0x19,
	0xca, 0x7b, 0x5f, 0x3c, 0xa1, 0xa6, 0x02, 0x3f, 0x81, 0x5b, 0x18, 0x2e, 0xe7, 0xc6, 0xe1, 0xa7,
	0x07, 0x8e, 0xcb, 0x79, 0x08, 0xf1, 0x07, 0x93, 0x9d, 0x03, 0x58, 0x5b, 0xe7, 0xd6, 0xd7, 0xe8,
	0x53, 0x5f, 0xe7, 0x3c, 0xac, 0x61, 0x3b, 0x7c, 0xbe, 0x04, 0x77, 0xef, 0x15, 0x44, 0xa7, 0xd0,
	0x1a, 0x07, 0x7e, 0x38, 0x89, 0xe6, 0xb7, 0x37, 0x41, 0x34, 0x99, 0x4e, 0x82, 0xf6, 0x67, 0xa8,
	0x03, 0xed, 0x3d, 0x71, 0x3c, 0xbd, 0xf2, 0xc7, 0x6d, 0x07, 0x7d, 0x0e, 0x4f, 0xf6, 0xd4, 0x17,
	0xe3, 0xe9, 0xa5, 0x3f, 0x6e, 0x97, 0x10, 0x82, 0xe6, 0x9e, 0xec, 0x4f, 0x6e, 0xdb, 0xe5, 0xe7,
	0x7f, 0x3a, 0xd0, 0xf9, 0x2f, 0x9f, 0xa3, 0x26, 0xc0, 0x6c, 0xee, 0xcf, 0x8b, 0x4c, 0xa7, 0xd0,
	0xd2, 0xdf, 0xb3, 0x60, 0x1c, 0x5c, 0xcd, 0x47, 0x93, 0x17, 0x6d, 0x47, 0xa6, 0xd7, 0x62, 0x18,
	0xbc, 0x79, 0x1b, 0xcc, 0x94, 0x5a, 0x42, 0x2d, 0x70, 0xb5, 0x7a, 0x39, 0x7d, 0x3b, 0xb9, 0x6e,
	0x97, 0x65, 0x62, 0x8b, 0x4d, 0x82, 0x5f, 0x25, 0x54, 0xf9, 0x10, 0x2f, 0x0c, 0x2e, 0x47, 0x93,
	0x6b, 0x29, 0x1e, 0x5d, 0xc2, 0x6f, 0xd5, 0x15, 0xde, 0xa8, 0x3f, 0xc2, 0xc5, 0xb1, 0xfa, 0xb9,
	0xf8, 0x37, 0x00, 0x00, 0xff, 0xff, 0x45, 0xbd, 0xda, 0x7f, 0x40, 0x07, 0x00, 0x00,
}
