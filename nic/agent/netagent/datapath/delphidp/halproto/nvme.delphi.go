// Code generated by protoc-gen-go. DO NOT EDIT.
// source: nvme.proto

package halproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// **********************   NVME Feature Enable *****************************//
// NvmeEnable Request object
type NvmeEnableRequest struct {
	MaxNs           uint32 `protobuf:"varint,1,opt,name=max_ns,json=maxNs" json:"max_ns,omitempty"`
	MaxSess         uint32 `protobuf:"varint,2,opt,name=max_sess,json=maxSess" json:"max_sess,omitempty"`
	MaxCmdContext   uint32 `protobuf:"varint,3,opt,name=max_cmd_context,json=maxCmdContext" json:"max_cmd_context,omitempty"`
	TxMaxPduContext uint32 `protobuf:"varint,4,opt,name=tx_max_pdu_context,json=txMaxPduContext" json:"tx_max_pdu_context,omitempty"`
	RxMaxPduContext uint32 `protobuf:"varint,5,opt,name=rx_max_pdu_context,json=rxMaxPduContext" json:"rx_max_pdu_context,omitempty"`
}

func (m *NvmeEnableRequest) Reset()                    { *m = NvmeEnableRequest{} }
func (m *NvmeEnableRequest) String() string            { return proto.CompactTextString(m) }
func (*NvmeEnableRequest) ProtoMessage()               {}
func (*NvmeEnableRequest) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{0} }

func (m *NvmeEnableRequest) GetMaxNs() uint32 {
	if m != nil {
		return m.MaxNs
	}
	return 0
}

func (m *NvmeEnableRequest) GetMaxSess() uint32 {
	if m != nil {
		return m.MaxSess
	}
	return 0
}

func (m *NvmeEnableRequest) GetMaxCmdContext() uint32 {
	if m != nil {
		return m.MaxCmdContext
	}
	return 0
}

func (m *NvmeEnableRequest) GetTxMaxPduContext() uint32 {
	if m != nil {
		return m.TxMaxPduContext
	}
	return 0
}

func (m *NvmeEnableRequest) GetRxMaxPduContext() uint32 {
	if m != nil {
		return m.RxMaxPduContext
	}
	return 0
}

// NvmeEnableRequestMsg is batched request
type NvmeEnableRequestMsg struct {
	Request []*NvmeEnableRequest `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *NvmeEnableRequestMsg) Reset()                    { *m = NvmeEnableRequestMsg{} }
func (m *NvmeEnableRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeEnableRequestMsg) ProtoMessage()               {}
func (*NvmeEnableRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{1} }

func (m *NvmeEnableRequestMsg) GetRequest() []*NvmeEnableRequest {
	if m != nil {
		return m.Request
	}
	return nil
}

// NvmeEnableResponse response to one NvmeEnableSpec
type NvmeEnableResponse struct {
	ApiStatus          ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	CmdContextRingBase uint64    `protobuf:"varint,2,opt,name=cmd_context_ring_base,json=cmdContextRingBase" json:"cmd_context_ring_base,omitempty"`
	CmdContextPageBase uint64    `protobuf:"varint,3,opt,name=cmd_context_page_base,json=cmdContextPageBase" json:"cmd_context_page_base,omitempty"`
}

func (m *NvmeEnableResponse) Reset()                    { *m = NvmeEnableResponse{} }
func (m *NvmeEnableResponse) String() string            { return proto.CompactTextString(m) }
func (*NvmeEnableResponse) ProtoMessage()               {}
func (*NvmeEnableResponse) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{2} }

func (m *NvmeEnableResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *NvmeEnableResponse) GetCmdContextRingBase() uint64 {
	if m != nil {
		return m.CmdContextRingBase
	}
	return 0
}

func (m *NvmeEnableResponse) GetCmdContextPageBase() uint64 {
	if m != nil {
		return m.CmdContextPageBase
	}
	return 0
}

// NvmeEnableResponseMsg is batched Response
type NvmeEnableResponseMsg struct {
	Response []*NvmeEnableResponse `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *NvmeEnableResponseMsg) Reset()                    { *m = NvmeEnableResponseMsg{} }
func (m *NvmeEnableResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeEnableResponseMsg) ProtoMessage()               {}
func (*NvmeEnableResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{3} }

func (m *NvmeEnableResponseMsg) GetResponse() []*NvmeEnableResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

// **********************   Submission Queue  *****************************//
// NvmeSq object
type NvmeSqSpec struct {
	SqNum     uint32 `protobuf:"varint,1,opt,name=sq_num,json=sqNum" json:"sq_num,omitempty"`
	HwLifId   uint32 `protobuf:"varint,2,opt,name=hw_lif_id,json=hwLifId" json:"hw_lif_id,omitempty"`
	SqWqeSize uint32 `protobuf:"varint,3,opt,name=sq_wqe_size,json=sqWqeSize" json:"sq_wqe_size,omitempty"`
	NumSqWqes uint32 `protobuf:"varint,4,opt,name=num_sq_wqes,json=numSqWqes" json:"num_sq_wqes,omitempty"`
	BaseAddr  uint64 `protobuf:"varint,5,opt,name=base_addr,json=baseAddr" json:"base_addr,omitempty"`
	CqNum     uint32 `protobuf:"varint,6,opt,name=cq_num,json=cqNum" json:"cq_num,omitempty"`
}

func (m *NvmeSqSpec) Reset()                    { *m = NvmeSqSpec{} }
func (m *NvmeSqSpec) String() string            { return proto.CompactTextString(m) }
func (*NvmeSqSpec) ProtoMessage()               {}
func (*NvmeSqSpec) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{4} }

func (m *NvmeSqSpec) GetSqNum() uint32 {
	if m != nil {
		return m.SqNum
	}
	return 0
}

func (m *NvmeSqSpec) GetHwLifId() uint32 {
	if m != nil {
		return m.HwLifId
	}
	return 0
}

func (m *NvmeSqSpec) GetSqWqeSize() uint32 {
	if m != nil {
		return m.SqWqeSize
	}
	return 0
}

func (m *NvmeSqSpec) GetNumSqWqes() uint32 {
	if m != nil {
		return m.NumSqWqes
	}
	return 0
}

func (m *NvmeSqSpec) GetBaseAddr() uint64 {
	if m != nil {
		return m.BaseAddr
	}
	return 0
}

func (m *NvmeSqSpec) GetCqNum() uint32 {
	if m != nil {
		return m.CqNum
	}
	return 0
}

// NvmeSqRequestMsg is batched request
type NvmeSqRequestMsg struct {
	Request []*NvmeSqSpec `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *NvmeSqRequestMsg) Reset()                    { *m = NvmeSqRequestMsg{} }
func (m *NvmeSqRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeSqRequestMsg) ProtoMessage()               {}
func (*NvmeSqRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{5} }

func (m *NvmeSqRequestMsg) GetRequest() []*NvmeSqSpec {
	if m != nil {
		return m.Request
	}
	return nil
}

// NvmeSqResponse response to one NvmeSqSpec
type NvmeSqResponse struct {
	ApiStatus ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
}

func (m *NvmeSqResponse) Reset()                    { *m = NvmeSqResponse{} }
func (m *NvmeSqResponse) String() string            { return proto.CompactTextString(m) }
func (*NvmeSqResponse) ProtoMessage()               {}
func (*NvmeSqResponse) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{6} }

func (m *NvmeSqResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

// NvmeSqResponseMsg is response to NvmeSqRequestMsg
type NvmeSqResponseMsg struct {
	Response []*NvmeSqResponse `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *NvmeSqResponseMsg) Reset()                    { *m = NvmeSqResponseMsg{} }
func (m *NvmeSqResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeSqResponseMsg) ProtoMessage()               {}
func (*NvmeSqResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{7} }

func (m *NvmeSqResponseMsg) GetResponse() []*NvmeSqResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

// **********************   Completion Queue  *****************************//
// NvmeCq object
type NvmeCqSpec struct {
	CqNum     uint32 `protobuf:"varint,1,opt,name=cq_num,json=cqNum" json:"cq_num,omitempty"`
	HwLifId   uint32 `protobuf:"varint,2,opt,name=hw_lif_id,json=hwLifId" json:"hw_lif_id,omitempty"`
	CqWqeSize uint32 `protobuf:"varint,3,opt,name=cq_wqe_size,json=cqWqeSize" json:"cq_wqe_size,omitempty"`
	NumCqWqes uint32 `protobuf:"varint,4,opt,name=num_cq_wqes,json=numCqWqes" json:"num_cq_wqes,omitempty"`
	BaseAddr  uint64 `protobuf:"varint,5,opt,name=base_addr,json=baseAddr" json:"base_addr,omitempty"`
	IntNum    uint32 `protobuf:"varint,6,opt,name=int_num,json=intNum" json:"int_num,omitempty"`
}

func (m *NvmeCqSpec) Reset()                    { *m = NvmeCqSpec{} }
func (m *NvmeCqSpec) String() string            { return proto.CompactTextString(m) }
func (*NvmeCqSpec) ProtoMessage()               {}
func (*NvmeCqSpec) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{8} }

func (m *NvmeCqSpec) GetCqNum() uint32 {
	if m != nil {
		return m.CqNum
	}
	return 0
}

func (m *NvmeCqSpec) GetHwLifId() uint32 {
	if m != nil {
		return m.HwLifId
	}
	return 0
}

func (m *NvmeCqSpec) GetCqWqeSize() uint32 {
	if m != nil {
		return m.CqWqeSize
	}
	return 0
}

func (m *NvmeCqSpec) GetNumCqWqes() uint32 {
	if m != nil {
		return m.NumCqWqes
	}
	return 0
}

func (m *NvmeCqSpec) GetBaseAddr() uint64 {
	if m != nil {
		return m.BaseAddr
	}
	return 0
}

func (m *NvmeCqSpec) GetIntNum() uint32 {
	if m != nil {
		return m.IntNum
	}
	return 0
}

// NvmeCqRequestMsg is batched request
type NvmeCqRequestMsg struct {
	Request []*NvmeCqSpec `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *NvmeCqRequestMsg) Reset()                    { *m = NvmeCqRequestMsg{} }
func (m *NvmeCqRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeCqRequestMsg) ProtoMessage()               {}
func (*NvmeCqRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{9} }

func (m *NvmeCqRequestMsg) GetRequest() []*NvmeCqSpec {
	if m != nil {
		return m.Request
	}
	return nil
}

// NvmeCqResponse response to one NvmeCqSpec
type NvmeCqResponse struct {
	ApiStatus     ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	CqIntrTblAddr uint32    `protobuf:"varint,2,opt,name=cq_intr_tbl_addr,json=cqIntrTblAddr" json:"cq_intr_tbl_addr,omitempty"`
}

func (m *NvmeCqResponse) Reset()                    { *m = NvmeCqResponse{} }
func (m *NvmeCqResponse) String() string            { return proto.CompactTextString(m) }
func (*NvmeCqResponse) ProtoMessage()               {}
func (*NvmeCqResponse) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{10} }

func (m *NvmeCqResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *NvmeCqResponse) GetCqIntrTblAddr() uint32 {
	if m != nil {
		return m.CqIntrTblAddr
	}
	return 0
}

// NvmeCqResponseMsg is response to NvmeCqRequestMsg
type NvmeCqResponseMsg struct {
	Response []*NvmeCqResponse `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *NvmeCqResponseMsg) Reset()                    { *m = NvmeCqResponseMsg{} }
func (m *NvmeCqResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeCqResponseMsg) ProtoMessage()               {}
func (*NvmeCqResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{11} }

func (m *NvmeCqResponseMsg) GetResponse() []*NvmeCqResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

// **********************   NameSpace *****************************//
// NvmeNs object
type NvmeNsSpec struct {
	Nsid        uint32 `protobuf:"varint,1,opt,name=nsid" json:"nsid,omitempty"`
	HwLifId     uint32 `protobuf:"varint,2,opt,name=hw_lif_id,json=hwLifId" json:"hw_lif_id,omitempty"`
	BackendNsid uint32 `protobuf:"varint,3,opt,name=backend_nsid,json=backendNsid" json:"backend_nsid,omitempty"`
	Size        uint32 `protobuf:"varint,4,opt,name=size" json:"size,omitempty"`
	LbaSize     uint32 `protobuf:"varint,5,opt,name=lba_size,json=lbaSize" json:"lba_size,omitempty"`
	MaxSess     uint32 `protobuf:"varint,6,opt,name=max_sess,json=maxSess" json:"max_sess,omitempty"`
	KeyIndex    uint32 `protobuf:"varint,7,opt,name=key_index,json=keyIndex" json:"key_index,omitempty"`
	SecKeyIndex uint32 `protobuf:"varint,8,opt,name=sec_key_index,json=secKeyIndex" json:"sec_key_index,omitempty"`
}

func (m *NvmeNsSpec) Reset()                    { *m = NvmeNsSpec{} }
func (m *NvmeNsSpec) String() string            { return proto.CompactTextString(m) }
func (*NvmeNsSpec) ProtoMessage()               {}
func (*NvmeNsSpec) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{12} }

func (m *NvmeNsSpec) GetNsid() uint32 {
	if m != nil {
		return m.Nsid
	}
	return 0
}

func (m *NvmeNsSpec) GetHwLifId() uint32 {
	if m != nil {
		return m.HwLifId
	}
	return 0
}

func (m *NvmeNsSpec) GetBackendNsid() uint32 {
	if m != nil {
		return m.BackendNsid
	}
	return 0
}

func (m *NvmeNsSpec) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *NvmeNsSpec) GetLbaSize() uint32 {
	if m != nil {
		return m.LbaSize
	}
	return 0
}

func (m *NvmeNsSpec) GetMaxSess() uint32 {
	if m != nil {
		return m.MaxSess
	}
	return 0
}

func (m *NvmeNsSpec) GetKeyIndex() uint32 {
	if m != nil {
		return m.KeyIndex
	}
	return 0
}

func (m *NvmeNsSpec) GetSecKeyIndex() uint32 {
	if m != nil {
		return m.SecKeyIndex
	}
	return 0
}

// NvmeNsRequestMsg is batched request used to create/update of Nvme QPs
type NvmeNsRequestMsg struct {
	Request []*NvmeNsSpec `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *NvmeNsRequestMsg) Reset()                    { *m = NvmeNsRequestMsg{} }
func (m *NvmeNsRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeNsRequestMsg) ProtoMessage()               {}
func (*NvmeNsRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{13} }

func (m *NvmeNsRequestMsg) GetRequest() []*NvmeNsSpec {
	if m != nil {
		return m.Request
	}
	return nil
}

// NvmeNsResponse response to one NvmeNsSpec
type NvmeNsResponse struct {
	ApiStatus ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	NscbAddr  uint64    `protobuf:"varint,2,opt,name=nscb_addr,json=nscbAddr" json:"nscb_addr,omitempty"`
}

func (m *NvmeNsResponse) Reset()                    { *m = NvmeNsResponse{} }
func (m *NvmeNsResponse) String() string            { return proto.CompactTextString(m) }
func (*NvmeNsResponse) ProtoMessage()               {}
func (*NvmeNsResponse) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{14} }

func (m *NvmeNsResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *NvmeNsResponse) GetNscbAddr() uint64 {
	if m != nil {
		return m.NscbAddr
	}
	return 0
}

// NvmeNsResponseMsg is response to NvmeNsRequestMsg
type NvmeNsResponseMsg struct {
	Response []*NvmeNsResponse `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *NvmeNsResponseMsg) Reset()                    { *m = NvmeNsResponseMsg{} }
func (m *NvmeNsResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeNsResponseMsg) ProtoMessage()               {}
func (*NvmeNsResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{15} }

func (m *NvmeNsResponseMsg) GetResponse() []*NvmeNsResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

// **********************   NameSpace *****************************//
// NvmeSess object
type NvmeSessSpec struct {
	HwLifId      uint32        `protobuf:"varint,1,opt,name=hw_lif_id,json=hwLifId" json:"hw_lif_id,omitempty"`
	Nsid         uint32        `protobuf:"varint,2,opt,name=nsid" json:"nsid,omitempty"`
	FlowKey      *FlowKey      `protobuf:"bytes,3,opt,name=flow_key,json=flowKey" json:"flow_key,omitempty"`
	VrfKeyHandle *VrfKeyHandle `protobuf:"bytes,4,opt,name=vrf_key_handle,json=vrfKeyHandle" json:"vrf_key_handle,omitempty"`
}

func (m *NvmeSessSpec) Reset()                    { *m = NvmeSessSpec{} }
func (m *NvmeSessSpec) String() string            { return proto.CompactTextString(m) }
func (*NvmeSessSpec) ProtoMessage()               {}
func (*NvmeSessSpec) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{16} }

func (m *NvmeSessSpec) GetHwLifId() uint32 {
	if m != nil {
		return m.HwLifId
	}
	return 0
}

func (m *NvmeSessSpec) GetNsid() uint32 {
	if m != nil {
		return m.Nsid
	}
	return 0
}

func (m *NvmeSessSpec) GetFlowKey() *FlowKey {
	if m != nil {
		return m.FlowKey
	}
	return nil
}

func (m *NvmeSessSpec) GetVrfKeyHandle() *VrfKeyHandle {
	if m != nil {
		return m.VrfKeyHandle
	}
	return nil
}

// NvmeNsRequestMsg is batched request used to create/update of Nvme QPs
type NvmeSessRequestMsg struct {
	Request []*NvmeSessSpec `protobuf:"bytes,1,rep,name=request" json:"request,omitempty"`
}

func (m *NvmeSessRequestMsg) Reset()                    { *m = NvmeSessRequestMsg{} }
func (m *NvmeSessRequestMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeSessRequestMsg) ProtoMessage()               {}
func (*NvmeSessRequestMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{17} }

func (m *NvmeSessRequestMsg) GetRequest() []*NvmeSessSpec {
	if m != nil {
		return m.Request
	}
	return nil
}

// NvmeSessResponse response to one NvmeSessSpec
type NvmeSessResponse struct {
	ApiStatus         ApiStatus `protobuf:"varint,1,opt,name=api_status,json=apiStatus,enum=types.ApiStatus" json:"api_status,omitempty"`
	SessId            uint32    `protobuf:"varint,2,opt,name=sess_id,json=sessId" json:"sess_id,omitempty"`
	TxsessprodcbAddr  uint64    `protobuf:"varint,3,opt,name=txsessprodcb_addr,json=txsessprodcbAddr" json:"txsessprodcb_addr,omitempty"`
	RxsessprodcbAddr  uint64    `protobuf:"varint,4,opt,name=rxsessprodcb_addr,json=rxsessprodcbAddr" json:"rxsessprodcb_addr,omitempty"`
	TxXtsqBase        uint64    `protobuf:"varint,5,opt,name=tx_xtsq_base,json=txXtsqBase" json:"tx_xtsq_base,omitempty"`
	TxXtsqNumEntries  uint64    `protobuf:"varint,6,opt,name=tx_xtsq_num_entries,json=txXtsqNumEntries" json:"tx_xtsq_num_entries,omitempty"`
	TxDgstqBase       uint64    `protobuf:"varint,7,opt,name=tx_dgstq_base,json=txDgstqBase" json:"tx_dgstq_base,omitempty"`
	TxDgstqNumEntries uint64    `protobuf:"varint,8,opt,name=tx_dgstq_num_entries,json=txDgstqNumEntries" json:"tx_dgstq_num_entries,omitempty"`
	TxSesqBase        uint64    `protobuf:"varint,9,opt,name=tx_sesq_base,json=txSesqBase" json:"tx_sesq_base,omitempty"`
	TxSesqNumEntries  uint64    `protobuf:"varint,10,opt,name=tx_sesq_num_entries,json=txSesqNumEntries" json:"tx_sesq_num_entries,omitempty"`
	RxXtsqBase        uint64    `protobuf:"varint,11,opt,name=rx_xtsq_base,json=rxXtsqBase" json:"rx_xtsq_base,omitempty"`
	RxXtsqNumEntries  uint64    `protobuf:"varint,12,opt,name=rx_xtsq_num_entries,json=rxXtsqNumEntries" json:"rx_xtsq_num_entries,omitempty"`
	RxDgstqBase       uint64    `protobuf:"varint,13,opt,name=rx_dgstq_base,json=rxDgstqBase" json:"rx_dgstq_base,omitempty"`
	RxDgstqNumEntries uint64    `protobuf:"varint,14,opt,name=rx_dgstq_num_entries,json=rxDgstqNumEntries" json:"rx_dgstq_num_entries,omitempty"`
	RxSerqBase        uint64    `protobuf:"varint,15,opt,name=rx_serq_base,json=rxSerqBase" json:"rx_serq_base,omitempty"`
	RxSerqNumEntries  uint64    `protobuf:"varint,16,opt,name=rx_serq_num_entries,json=rxSerqNumEntries" json:"rx_serq_num_entries,omitempty"`
}

func (m *NvmeSessResponse) Reset()                    { *m = NvmeSessResponse{} }
func (m *NvmeSessResponse) String() string            { return proto.CompactTextString(m) }
func (*NvmeSessResponse) ProtoMessage()               {}
func (*NvmeSessResponse) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{18} }

func (m *NvmeSessResponse) GetApiStatus() ApiStatus {
	if m != nil {
		return m.ApiStatus
	}
	return ApiStatus_API_STATUS_OK
}

func (m *NvmeSessResponse) GetSessId() uint32 {
	if m != nil {
		return m.SessId
	}
	return 0
}

func (m *NvmeSessResponse) GetTxsessprodcbAddr() uint64 {
	if m != nil {
		return m.TxsessprodcbAddr
	}
	return 0
}

func (m *NvmeSessResponse) GetRxsessprodcbAddr() uint64 {
	if m != nil {
		return m.RxsessprodcbAddr
	}
	return 0
}

func (m *NvmeSessResponse) GetTxXtsqBase() uint64 {
	if m != nil {
		return m.TxXtsqBase
	}
	return 0
}

func (m *NvmeSessResponse) GetTxXtsqNumEntries() uint64 {
	if m != nil {
		return m.TxXtsqNumEntries
	}
	return 0
}

func (m *NvmeSessResponse) GetTxDgstqBase() uint64 {
	if m != nil {
		return m.TxDgstqBase
	}
	return 0
}

func (m *NvmeSessResponse) GetTxDgstqNumEntries() uint64 {
	if m != nil {
		return m.TxDgstqNumEntries
	}
	return 0
}

func (m *NvmeSessResponse) GetTxSesqBase() uint64 {
	if m != nil {
		return m.TxSesqBase
	}
	return 0
}

func (m *NvmeSessResponse) GetTxSesqNumEntries() uint64 {
	if m != nil {
		return m.TxSesqNumEntries
	}
	return 0
}

func (m *NvmeSessResponse) GetRxXtsqBase() uint64 {
	if m != nil {
		return m.RxXtsqBase
	}
	return 0
}

func (m *NvmeSessResponse) GetRxXtsqNumEntries() uint64 {
	if m != nil {
		return m.RxXtsqNumEntries
	}
	return 0
}

func (m *NvmeSessResponse) GetRxDgstqBase() uint64 {
	if m != nil {
		return m.RxDgstqBase
	}
	return 0
}

func (m *NvmeSessResponse) GetRxDgstqNumEntries() uint64 {
	if m != nil {
		return m.RxDgstqNumEntries
	}
	return 0
}

func (m *NvmeSessResponse) GetRxSerqBase() uint64 {
	if m != nil {
		return m.RxSerqBase
	}
	return 0
}

func (m *NvmeSessResponse) GetRxSerqNumEntries() uint64 {
	if m != nil {
		return m.RxSerqNumEntries
	}
	return 0
}

// NvmeSessResponseMsg is response to NvmeSessRequestMsg
type NvmeSessResponseMsg struct {
	Response []*NvmeSessResponse `protobuf:"bytes,1,rep,name=response" json:"response,omitempty"`
}

func (m *NvmeSessResponseMsg) Reset()                    { *m = NvmeSessResponseMsg{} }
func (m *NvmeSessResponseMsg) String() string            { return proto.CompactTextString(m) }
func (*NvmeSessResponseMsg) ProtoMessage()               {}
func (*NvmeSessResponseMsg) Descriptor() ([]byte, []int) { return fileDescriptor21, []int{19} }

func (m *NvmeSessResponseMsg) GetResponse() []*NvmeSessResponse {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*NvmeEnableRequest)(nil), "halproto.NvmeEnableRequest")
	proto.RegisterType((*NvmeEnableRequestMsg)(nil), "halproto.NvmeEnableRequestMsg")
	proto.RegisterType((*NvmeEnableResponse)(nil), "halproto.NvmeEnableResponse")
	proto.RegisterType((*NvmeEnableResponseMsg)(nil), "halproto.NvmeEnableResponseMsg")
	proto.RegisterType((*NvmeSqSpec)(nil), "halproto.NvmeSqSpec")
	proto.RegisterType((*NvmeSqRequestMsg)(nil), "halproto.NvmeSqRequestMsg")
	proto.RegisterType((*NvmeSqResponse)(nil), "halproto.NvmeSqResponse")
	proto.RegisterType((*NvmeSqResponseMsg)(nil), "halproto.NvmeSqResponseMsg")
	proto.RegisterType((*NvmeCqSpec)(nil), "halproto.NvmeCqSpec")
	proto.RegisterType((*NvmeCqRequestMsg)(nil), "halproto.NvmeCqRequestMsg")
	proto.RegisterType((*NvmeCqResponse)(nil), "halproto.NvmeCqResponse")
	proto.RegisterType((*NvmeCqResponseMsg)(nil), "halproto.NvmeCqResponseMsg")
	proto.RegisterType((*NvmeNsSpec)(nil), "halproto.NvmeNsSpec")
	proto.RegisterType((*NvmeNsRequestMsg)(nil), "halproto.NvmeNsRequestMsg")
	proto.RegisterType((*NvmeNsResponse)(nil), "halproto.NvmeNsResponse")
	proto.RegisterType((*NvmeNsResponseMsg)(nil), "halproto.NvmeNsResponseMsg")
	proto.RegisterType((*NvmeSessSpec)(nil), "halproto.NvmeSessSpec")
	proto.RegisterType((*NvmeSessRequestMsg)(nil), "halproto.NvmeSessRequestMsg")
	proto.RegisterType((*NvmeSessResponse)(nil), "halproto.NvmeSessResponse")
	proto.RegisterType((*NvmeSessResponseMsg)(nil), "halproto.NvmeSessResponseMsg")
}

func init() { proto.RegisterFile("nvme.proto", fileDescriptor21) }

var fileDescriptor21 = []byte{
	// 1136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x97, 0xdf, 0x6e, 0xe3, 0x44,
	0x14, 0xc6, 0x37, 0xbb, 0xd9, 0xfc, 0x39, 0x69, 0xda, 0x74, 0xb6, 0xdd, 0x66, 0x5b, 0x09, 0x15,
	0x5f, 0xc0, 0x8a, 0x42, 0x97, 0x2d, 0x88, 0x4b, 0xa4, 0xd6, 0x94, 0x25, 0x2a, 0x8d, 0x56, 0x0e,
	0x02, 0xc4, 0x05, 0x96, 0x63, 0x4f, 0x12, 0xd3, 0x78, 0x12, 0xcf, 0x4c, 0x5a, 0x77, 0x5f, 0x07,
	0xee, 0xb8, 0x44, 0x3c, 0x03, 0xcf, 0xc1, 0x2d, 0x4f, 0x81, 0x66, 0x8e, 0x1d, 0x8f, 0xe3, 0x14,
	0xaa, 0x5e, 0x75, 0x3c, 0xf3, 0xcd, 0xe9, 0x39, 0xdf, 0x6f, 0x7c, 0xc6, 0x01, 0x60, 0xd7, 0x11,
	0x3d, 0x9e, 0xf3, 0x99, 0x9c, 0x91, 0xaa, 0x1a, 0xef, 0xb7, 0xe4, 0xed, 0x9c, 0x0a, 0x9c, 0xda,
	0x6f, 0x0b, 0x2a, 0x44, 0x38, 0x63, 0xe9, 0x63, 0xe3, 0x6a, 0x82, 0x23, 0xeb, 0xaf, 0x0a, 0x6c,
	0xf7, 0xaf, 0x23, 0x7a, 0xce, 0xbc, 0xe1, 0x94, 0x3a, 0x34, 0x5e, 0x50, 0x21, 0xc9, 0x2e, 0xd4,
	0x22, 0x2f, 0x71, 0x99, 0xe8, 0x56, 0x0e, 0x2b, 0x2f, 0xdb, 0xce, 0xd3, 0xc8, 0x4b, 0xfa, 0x82,
	0xbc, 0x80, 0x86, 0x9a, 0x56, 0xb1, 0xba, 0x8f, 0xf5, 0x42, 0x3d, 0xf2, 0x92, 0x01, 0x15, 0x82,
	0x7c, 0x00, 0x5b, 0x6a, 0xc9, 0x8f, 0x02, 0xd7, 0x9f, 0x31, 0x49, 0x13, 0xd9, 0x7d, 0xa2, 0x15,
	0xed, 0xc8, 0x4b, 0xec, 0x28, 0xb0, 0x71, 0x92, 0x1c, 0x01, 0x91, 0x89, 0xab, 0xa4, 0xf3, 0x60,
	0xb1, 0x94, 0x56, 0xb5, 0x74, 0x4b, 0x26, 0x97, 0x5e, 0xf2, 0x36, 0x58, 0x18, 0x62, 0x5e, 0x16,
	0x3f, 0x45, 0x31, 0x2f, 0x8a, 0xad, 0x1e, 0xec, 0x94, 0x0a, 0xb9, 0x14, 0x63, 0xf2, 0x1a, 0xea,
	0x1c, 0x9f, 0xba, 0x95, 0xc3, 0x27, 0x2f, 0x5b, 0x27, 0x7b, 0xc7, 0xda, 0xab, 0x92, 0xd8, 0xc9,
	0x74, 0xd6, 0xef, 0x15, 0x20, 0xe6, 0xb2, 0x98, 0xcf, 0x98, 0xa0, 0xe4, 0x15, 0x80, 0x37, 0x0f,
	0x5d, 0x21, 0x3d, 0xb9, 0x40, 0x67, 0x36, 0x4f, 0x3a, 0xc7, 0x68, 0xf3, 0xe9, 0x3c, 0x1c, 0xe8,
	0x79, 0xa7, 0xe9, 0x65, 0x43, 0xf2, 0x1a, 0x76, 0x0d, 0x43, 0x5c, 0x1e, 0xb2, 0xb1, 0x3b, 0xf4,
	0x04, 0xd5, 0xe6, 0x55, 0x1d, 0xe2, 0x2f, 0x7d, 0x71, 0x42, 0x36, 0x3e, 0xf3, 0x04, 0x5d, 0xdd,
	0x32, 0xf7, 0xc6, 0x14, 0xb7, 0x3c, 0x59, 0xdd, 0xf2, 0xd6, 0x1b, 0x53, 0xb5, 0xc5, 0xba, 0x84,
	0xdd, 0x72, 0xb2, 0xaa, 0xf2, 0xcf, 0xa1, 0xc1, 0xd3, 0xc7, 0xb4, 0xf4, 0x6e, 0xb9, 0x74, 0x5c,
	0x77, 0x96, 0x4a, 0xeb, 0x8f, 0x0a, 0x80, 0x12, 0x0c, 0xe2, 0xc1, 0x9c, 0xfa, 0xea, 0x28, 0x88,
	0xd8, 0x65, 0x8b, 0x28, 0x3b, 0x0a, 0x22, 0xee, 0x2f, 0x22, 0xb2, 0x0f, 0xcd, 0xc9, 0x8d, 0x3b,
	0x0d, 0x47, 0x6e, 0x18, 0x64, 0x67, 0x61, 0x72, 0xf3, 0x6d, 0x38, 0xea, 0x05, 0xe4, 0x3d, 0x68,
	0x89, 0xd8, 0xbd, 0x89, 0xa9, 0x2b, 0xc2, 0x77, 0x34, 0x3d, 0x07, 0x4d, 0x11, 0xff, 0x10, 0xd3,
	0x41, 0xf8, 0x8e, 0xaa, 0x75, 0xb6, 0x88, 0x5c, 0xd4, 0x88, 0x14, 0x7e, 0x93, 0x2d, 0xa2, 0x81,
	0x92, 0x08, 0x72, 0x00, 0x4d, 0x55, 0xb2, 0xeb, 0x05, 0x01, 0xd7, 0xb4, 0xab, 0x4e, 0x43, 0x4d,
	0x9c, 0x06, 0x01, 0x57, 0xf9, 0xf8, 0x98, 0x4f, 0x0d, 0xf3, 0xf1, 0x55, 0x3e, 0xd6, 0x97, 0xd0,
	0xc1, 0xa4, 0x0d, 0xf2, 0x1f, 0xad, 0x92, 0xef, 0xe4, 0xe5, 0x63, 0x75, 0x39, 0xf2, 0x53, 0xd8,
	0xcc, 0xf6, 0x3f, 0x90, 0xb6, 0x75, 0x8e, 0x6f, 0x52, 0x1e, 0x42, 0xe5, 0xf0, 0x69, 0x89, 0xc1,
	0x8e, 0x99, 0xc4, 0x1a, 0xff, 0xff, 0x4c, 0xfd, 0xb7, 0x97, 0xfe, 0xfb, 0x05, 0xff, 0xfd, 0xfb,
	0xf8, 0xef, 0x97, 0xfd, 0xf7, 0x57, 0xfd, 0xf7, 0x4b, 0xfe, 0xdb, 0xf7, 0xf0, 0x7f, 0x0f, 0xea,
	0x21, 0x93, 0x06, 0x80, 0x5a, 0xc8, 0xa4, 0x41, 0xc0, 0xbe, 0x2f, 0x01, 0x7b, 0x85, 0xc0, 0x2f,
	0x48, 0xc0, 0x7e, 0x38, 0x01, 0xf2, 0x21, 0x74, 0xfc, 0xd8, 0x0d, 0x99, 0xe4, 0xae, 0x1c, 0x4e,
	0x31, 0x7f, 0xf4, 0xa6, 0xed, 0xc7, 0x3d, 0x26, 0xf9, 0x77, 0xc3, 0xa9, 0x2a, 0x22, 0x43, 0x65,
	0xdf, 0x1f, 0x95, 0xbd, 0x0e, 0xd5, 0x3f, 0x29, 0xaa, 0xbe, 0xd0, 0xa8, 0x08, 0x54, 0x99, 0x08,
	0x83, 0x14, 0x94, 0x1e, 0xff, 0x27, 0xa7, 0xf7, 0x61, 0x63, 0xe8, 0xf9, 0x57, 0x94, 0x05, 0xae,
	0xde, 0x87, 0xa0, 0x5a, 0xe9, 0x5c, 0x5f, 0x6d, 0x27, 0x50, 0xd5, 0x0c, 0x91, 0x91, 0x1e, 0xab,
	0x2e, 0x3c, 0x1d, 0x7a, 0xc8, 0x16, 0x7b, 0x61, 0x7d, 0x3a, 0xf4, 0x06, 0xe9, 0xd2, 0xb2, 0x41,
	0xd7, 0x8a, 0x0d, 0xfa, 0x00, 0x9a, 0x57, 0xf4, 0xd6, 0x0d, 0x59, 0x40, 0x93, 0x6e, 0x5d, 0xaf,
	0x35, 0xae, 0xe8, 0x6d, 0x4f, 0x3d, 0x13, 0x0b, 0xda, 0x82, 0xfa, 0x6e, 0x2e, 0x68, 0x60, 0x2a,
	0x82, 0xfa, 0x17, 0xa9, 0x26, 0xe3, 0xdb, 0x17, 0xf7, 0xe4, 0x8b, 0xa6, 0xe4, 0x7c, 0x7f, 0x46,
	0xbe, 0x6a, 0xff, 0x43, 0xf9, 0x1e, 0x40, 0x93, 0x09, 0x7f, 0x98, 0x83, 0xad, 0x3a, 0x0d, 0x35,
	0x61, 0x32, 0xcd, 0xe3, 0xff, 0x2f, 0xd3, 0x5c, 0x6a, 0x30, 0xfd, 0xb5, 0x02, 0x1b, 0xfa, 0xdd,
	0xa4, 0x02, 0xa9, 0x16, 0x08, 0x56, 0x8a, 0x04, 0x33, 0xe2, 0x8f, 0x0d, 0xe2, 0x47, 0xd0, 0x18,
	0x4d, 0x67, 0x37, 0xca, 0x4c, 0x4d, 0x54, 0x99, 0x92, 0xdd, 0xbe, 0x5f, 0x4f, 0x67, 0x37, 0x17,
	0xf4, 0xd6, 0xa9, 0x8f, 0x70, 0x40, 0xbe, 0x80, 0xcd, 0x6b, 0x3e, 0xd2, 0xc6, 0x4f, 0x3c, 0x16,
	0x4c, 0x91, 0xb4, 0xda, 0x72, 0x35, 0x39, 0xfe, 0x9e, 0x8f, 0x2e, 0xe8, 0xed, 0x37, 0x7a, 0xde,
	0xd9, 0xb8, 0x36, 0x9e, 0xac, 0x33, 0xbc, 0xa0, 0x54, 0x92, 0x06, 0x8e, 0x8f, 0x57, 0x71, 0x10,
	0xa3, 0xd7, 0xa4, 0xf5, 0xe4, 0x40, 0x7e, 0x7b, 0x9a, 0xf6, 0x4c, 0x1d, 0xe4, 0xa1, 0x4c, 0xf6,
	0xa0, 0xae, 0xaa, 0xcb, 0x8f, 0x77, 0x4d, 0x3d, 0xf6, 0x94, 0x0f, 0xdb, 0x32, 0x51, 0xe3, 0x39,
	0x9f, 0x05, 0x19, 0x34, 0xbc, 0xc5, 0x3a, 0xe6, 0x82, 0xee, 0x2a, 0x47, 0xb0, 0xcd, 0x4b, 0xe2,
	0x2a, 0x8a, 0xf9, 0xaa, 0xf8, 0x10, 0x36, 0x64, 0xe2, 0x26, 0x52, 0xc4, 0x78, 0x35, 0x62, 0x8b,
	0x02, 0x99, 0xfc, 0x28, 0x45, 0xac, 0x6f, 0xd1, 0x4f, 0xe0, 0x59, 0xa6, 0x50, 0x9d, 0x8e, 0x32,
	0xc9, 0x43, 0x8a, 0xaf, 0x84, 0xfe, 0xef, 0x4a, 0xd8, 0x5f, 0x44, 0xe7, 0x38, 0xaf, 0x8e, 0xbf,
	0x4c, 0xdc, 0x60, 0x2c, 0x64, 0x1a, 0xb1, 0xae, 0x85, 0x2d, 0x99, 0x7c, 0xa5, 0xe6, 0x74, 0xc8,
	0x57, 0xb0, 0xb3, 0xd4, 0x98, 0x31, 0x1b, 0x5a, 0xba, 0x9d, 0x4a, 0x8d, 0xa0, 0x98, 0xa5, 0xa0,
	0x59, 0x96, 0xcd, 0x2c, 0xcb, 0x01, 0x2d, 0x64, 0xa9, 0x15, 0x66, 0x44, 0xc8, 0xb2, 0x54, 0xc2,
	0x62, 0x40, 0x6e, 0x96, 0xdd, 0xc2, 0x80, 0xbc, 0x50, 0x36, 0x5f, 0x53, 0xf6, 0x46, 0xe6, 0x63,
	0xb9, 0x6c, 0x5e, 0x28, 0xbb, 0x8d, 0x65, 0xf3, 0x62, 0xd9, 0x7c, 0x5d, 0xd9, 0x9b, 0x58, 0x36,
	0x5f, 0x57, 0x36, 0x57, 0x45, 0xf1, 0x34, 0xe6, 0x56, 0x96, 0xe5, 0x80, 0x72, 0x33, 0x4b, 0xad,
	0x30, 0x23, 0x76, 0xb2, 0x2c, 0x95, 0x30, 0x0f, 0x68, 0xf5, 0xe0, 0xd9, 0xea, 0x29, 0x55, 0x67,
	0xfd, 0xa4, 0xf4, 0x66, 0x3f, 0x2f, 0x1e, 0xf6, 0xf2, 0xbb, 0x7d, 0xf2, 0xf7, 0x63, 0xa8, 0xaa,
	0x65, 0xf2, 0x06, 0xfb, 0x36, 0x7e, 0x03, 0x91, 0xfd, 0x3b, 0x3e, 0x08, 0x2f, 0xc5, 0x78, 0xff,
	0xe0, 0xae, 0x2f, 0xa6, 0x4b, 0x31, 0xb6, 0x1e, 0x91, 0xd3, 0xb4, 0x59, 0xc4, 0x36, 0xa7, 0x9e,
	0xa4, 0xe4, 0x79, 0xf1, 0x72, 0x5f, 0x86, 0xd9, 0x5b, 0x77, 0xe9, 0x17, 0x42, 0xd8, 0x6b, 0x42,
	0xd8, 0x77, 0x84, 0xb0, 0xd7, 0x87, 0xe8, 0x8b, 0x72, 0x08, 0xb3, 0x5d, 0x9b, 0x21, 0x0a, 0x6d,
	0xd2, 0x7a, 0x44, 0xde, 0xa4, 0xdf, 0x3f, 0x54, 0x64, 0x41, 0xba, 0xab, 0x76, 0x2e, 0xc3, 0xbc,
	0x58, 0x6f, 0xb4, 0x0e, 0x74, 0x06, 0x3f, 0x35, 0x26, 0xde, 0x54, 0xff, 0xb8, 0x18, 0xd6, 0xf4,
	0x9f, 0xcf, 0xfe, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xde, 0xbb, 0x48, 0x36, 0x9d, 0x0c, 0x00, 0x00,
}
