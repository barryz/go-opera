// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package proto

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type IDResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDResponse) Reset()         { *m = IDResponse{} }
func (m *IDResponse) String() string { return proto.CompactTextString(m) }
func (*IDResponse) ProtoMessage()    {}
func (*IDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{0}
}

func (m *IDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDResponse.Unmarshal(m, b)
}
func (m *IDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDResponse.Marshal(b, m, deterministic)
}
func (m *IDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDResponse.Merge(m, src)
}
func (m *IDResponse) XXX_Size() int {
	return xxx_messageInfo_IDResponse.Size(m)
}
func (m *IDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IDResponse proto.InternalMessageInfo

func (m *IDResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StakeResponse struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakeResponse) Reset()         { *m = StakeResponse{} }
func (m *StakeResponse) String() string { return proto.CompactTextString(m) }
func (*StakeResponse) ProtoMessage()    {}
func (*StakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{1}
}

func (m *StakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakeResponse.Unmarshal(m, b)
}
func (m *StakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakeResponse.Marshal(b, m, deterministic)
}
func (m *StakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeResponse.Merge(m, src)
}
func (m *StakeResponse) XXX_Size() int {
	return xxx_messageInfo_StakeResponse.Size(m)
}
func (m *StakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StakeResponse proto.InternalMessageInfo

func (m *StakeResponse) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type InternalTxRequest struct {
	Amount               uint64   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Receiver             string   `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalTxRequest) Reset()         { *m = InternalTxRequest{} }
func (m *InternalTxRequest) String() string { return proto.CompactTextString(m) }
func (*InternalTxRequest) ProtoMessage()    {}
func (*InternalTxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{2}
}

func (m *InternalTxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalTxRequest.Unmarshal(m, b)
}
func (m *InternalTxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalTxRequest.Marshal(b, m, deterministic)
}
func (m *InternalTxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTxRequest.Merge(m, src)
}
func (m *InternalTxRequest) XXX_Size() int {
	return xxx_messageInfo_InternalTxRequest.Size(m)
}
func (m *InternalTxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTxRequest proto.InternalMessageInfo

func (m *InternalTxRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InternalTxRequest) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

type InternalTxResponse struct {
	Value                float32  `protobuf:"fixed32,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalTxResponse) Reset()         { *m = InternalTxResponse{} }
func (m *InternalTxResponse) String() string { return proto.CompactTextString(m) }
func (*InternalTxResponse) ProtoMessage()    {}
func (*InternalTxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{3}
}

func (m *InternalTxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalTxResponse.Unmarshal(m, b)
}
func (m *InternalTxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalTxResponse.Marshal(b, m, deterministic)
}
func (m *InternalTxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTxResponse.Merge(m, src)
}
func (m *InternalTxResponse) XXX_Size() int {
	return xxx_messageInfo_InternalTxResponse.Size(m)
}
func (m *InternalTxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTxResponse proto.InternalMessageInfo

func (m *InternalTxResponse) GetValue() float32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type ToServer struct {
	// Types that are valid to be assigned to Event:
	//	*ToServer_Tx_
	//	*ToServer_Answer_
	Event                isToServer_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ToServer) Reset()         { *m = ToServer{} }
func (m *ToServer) String() string { return proto.CompactTextString(m) }
func (*ToServer) ProtoMessage()    {}
func (*ToServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{4}
}

func (m *ToServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer.Unmarshal(m, b)
}
func (m *ToServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer.Marshal(b, m, deterministic)
}
func (m *ToServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer.Merge(m, src)
}
func (m *ToServer) XXX_Size() int {
	return xxx_messageInfo_ToServer.Size(m)
}
func (m *ToServer) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer proto.InternalMessageInfo

type isToServer_Event interface {
	isToServer_Event()
}

type ToServer_Tx_ struct {
	Tx *ToServer_Tx `protobuf:"bytes,1,opt,name=tx,proto3,oneof"`
}

type ToServer_Answer_ struct {
	Answer *ToServer_Answer `protobuf:"bytes,2,opt,name=answer,proto3,oneof"`
}

func (*ToServer_Tx_) isToServer_Event() {}

func (*ToServer_Answer_) isToServer_Event() {}

func (m *ToServer) GetEvent() isToServer_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ToServer) GetTx() *ToServer_Tx {
	if x, ok := m.GetEvent().(*ToServer_Tx_); ok {
		return x.Tx
	}
	return nil
}

func (m *ToServer) GetAnswer() *ToServer_Answer {
	if x, ok := m.GetEvent().(*ToServer_Answer_); ok {
		return x.Answer
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ToServer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ToServer_Tx_)(nil),
		(*ToServer_Answer_)(nil),
	}
}

type ToServer_Tx struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToServer_Tx) Reset()         { *m = ToServer_Tx{} }
func (m *ToServer_Tx) String() string { return proto.CompactTextString(m) }
func (*ToServer_Tx) ProtoMessage()    {}
func (*ToServer_Tx) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{4, 0}
}

func (m *ToServer_Tx) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer_Tx.Unmarshal(m, b)
}
func (m *ToServer_Tx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer_Tx.Marshal(b, m, deterministic)
}
func (m *ToServer_Tx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer_Tx.Merge(m, src)
}
func (m *ToServer_Tx) XXX_Size() int {
	return xxx_messageInfo_ToServer_Tx.Size(m)
}
func (m *ToServer_Tx) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer_Tx.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer_Tx proto.InternalMessageInfo

func (m *ToServer_Tx) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ToServer_Answer struct {
	Uid []byte `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	// Types that are valid to be assigned to Payload:
	//	*ToServer_Answer_Data
	//	*ToServer_Answer_Error
	Payload              isToServer_Answer_Payload `protobuf_oneof:"payload"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ToServer_Answer) Reset()         { *m = ToServer_Answer{} }
func (m *ToServer_Answer) String() string { return proto.CompactTextString(m) }
func (*ToServer_Answer) ProtoMessage()    {}
func (*ToServer_Answer) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{4, 1}
}

func (m *ToServer_Answer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToServer_Answer.Unmarshal(m, b)
}
func (m *ToServer_Answer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToServer_Answer.Marshal(b, m, deterministic)
}
func (m *ToServer_Answer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToServer_Answer.Merge(m, src)
}
func (m *ToServer_Answer) XXX_Size() int {
	return xxx_messageInfo_ToServer_Answer.Size(m)
}
func (m *ToServer_Answer) XXX_DiscardUnknown() {
	xxx_messageInfo_ToServer_Answer.DiscardUnknown(m)
}

var xxx_messageInfo_ToServer_Answer proto.InternalMessageInfo

func (m *ToServer_Answer) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

type isToServer_Answer_Payload interface {
	isToServer_Answer_Payload()
}

type ToServer_Answer_Data struct {
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3,oneof"`
}

type ToServer_Answer_Error struct {
	Error string `protobuf:"bytes,3,opt,name=error,proto3,oneof"`
}

func (*ToServer_Answer_Data) isToServer_Answer_Payload() {}

func (*ToServer_Answer_Error) isToServer_Answer_Payload() {}

func (m *ToServer_Answer) GetPayload() isToServer_Answer_Payload {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ToServer_Answer) GetData() []byte {
	if x, ok := m.GetPayload().(*ToServer_Answer_Data); ok {
		return x.Data
	}
	return nil
}

func (m *ToServer_Answer) GetError() string {
	if x, ok := m.GetPayload().(*ToServer_Answer_Error); ok {
		return x.Error
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ToServer_Answer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ToServer_Answer_Data)(nil),
		(*ToServer_Answer_Error)(nil),
	}
}

type ToClient struct {
	// Types that are valid to be assigned to Event:
	//	*ToClient_Block_
	//	*ToClient_Query_
	//	*ToClient_Restore_
	Event                isToClient_Event `protobuf_oneof:"event"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ToClient) Reset()         { *m = ToClient{} }
func (m *ToClient) String() string { return proto.CompactTextString(m) }
func (*ToClient) ProtoMessage()    {}
func (*ToClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{5}
}

func (m *ToClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient.Unmarshal(m, b)
}
func (m *ToClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient.Marshal(b, m, deterministic)
}
func (m *ToClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient.Merge(m, src)
}
func (m *ToClient) XXX_Size() int {
	return xxx_messageInfo_ToClient.Size(m)
}
func (m *ToClient) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient proto.InternalMessageInfo

type isToClient_Event interface {
	isToClient_Event()
}

type ToClient_Block_ struct {
	Block *ToClient_Block `protobuf:"bytes,1,opt,name=block,proto3,oneof"`
}

type ToClient_Query_ struct {
	Query *ToClient_Query `protobuf:"bytes,2,opt,name=query,proto3,oneof"`
}

type ToClient_Restore_ struct {
	Restore *ToClient_Restore `protobuf:"bytes,3,opt,name=restore,proto3,oneof"`
}

func (*ToClient_Block_) isToClient_Event() {}

func (*ToClient_Query_) isToClient_Event() {}

func (*ToClient_Restore_) isToClient_Event() {}

func (m *ToClient) GetEvent() isToClient_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ToClient) GetBlock() *ToClient_Block {
	if x, ok := m.GetEvent().(*ToClient_Block_); ok {
		return x.Block
	}
	return nil
}

func (m *ToClient) GetQuery() *ToClient_Query {
	if x, ok := m.GetEvent().(*ToClient_Query_); ok {
		return x.Query
	}
	return nil
}

func (m *ToClient) GetRestore() *ToClient_Restore {
	if x, ok := m.GetEvent().(*ToClient_Restore_); ok {
		return x.Restore
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ToClient) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ToClient_Block_)(nil),
		(*ToClient_Query_)(nil),
		(*ToClient_Restore_)(nil),
	}
}

type ToClient_Block struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Block) Reset()         { *m = ToClient_Block{} }
func (m *ToClient_Block) String() string { return proto.CompactTextString(m) }
func (*ToClient_Block) ProtoMessage()    {}
func (*ToClient_Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{5, 0}
}

func (m *ToClient_Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Block.Unmarshal(m, b)
}
func (m *ToClient_Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Block.Marshal(b, m, deterministic)
}
func (m *ToClient_Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Block.Merge(m, src)
}
func (m *ToClient_Block) XXX_Size() int {
	return xxx_messageInfo_ToClient_Block.Size(m)
}
func (m *ToClient_Block) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Block.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Block proto.InternalMessageInfo

func (m *ToClient_Block) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Block) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ToClient_Query struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Index                int64    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Query) Reset()         { *m = ToClient_Query{} }
func (m *ToClient_Query) String() string { return proto.CompactTextString(m) }
func (*ToClient_Query) ProtoMessage()    {}
func (*ToClient_Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{5, 1}
}

func (m *ToClient_Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Query.Unmarshal(m, b)
}
func (m *ToClient_Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Query.Marshal(b, m, deterministic)
}
func (m *ToClient_Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Query.Merge(m, src)
}
func (m *ToClient_Query) XXX_Size() int {
	return xxx_messageInfo_ToClient_Query.Size(m)
}
func (m *ToClient_Query) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Query.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Query proto.InternalMessageInfo

func (m *ToClient_Query) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Query) GetIndex() int64 {
	if m != nil {
		return m.Index
	}
	return 0
}

type ToClient_Restore struct {
	Uid                  []byte   `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ToClient_Restore) Reset()         { *m = ToClient_Restore{} }
func (m *ToClient_Restore) String() string { return proto.CompactTextString(m) }
func (*ToClient_Restore) ProtoMessage()    {}
func (*ToClient_Restore) Descriptor() ([]byte, []int) {
	return fileDescriptor_bedfbfc9b54e5600, []int{5, 2}
}

func (m *ToClient_Restore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ToClient_Restore.Unmarshal(m, b)
}
func (m *ToClient_Restore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ToClient_Restore.Marshal(b, m, deterministic)
}
func (m *ToClient_Restore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ToClient_Restore.Merge(m, src)
}
func (m *ToClient_Restore) XXX_Size() int {
	return xxx_messageInfo_ToClient_Restore.Size(m)
}
func (m *ToClient_Restore) XXX_DiscardUnknown() {
	xxx_messageInfo_ToClient_Restore.DiscardUnknown(m)
}

var xxx_messageInfo_ToClient_Restore proto.InternalMessageInfo

func (m *ToClient_Restore) GetUid() []byte {
	if m != nil {
		return m.Uid
	}
	return nil
}

func (m *ToClient_Restore) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*IDResponse)(nil), "internal.IDResponse")
	proto.RegisterType((*StakeResponse)(nil), "internal.StakeResponse")
	proto.RegisterType((*InternalTxRequest)(nil), "internal.InternalTxRequest")
	proto.RegisterType((*InternalTxResponse)(nil), "internal.InternalTxResponse")
	proto.RegisterType((*ToServer)(nil), "internal.ToServer")
	proto.RegisterType((*ToServer_Tx)(nil), "internal.ToServer.Tx")
	proto.RegisterType((*ToServer_Answer)(nil), "internal.ToServer.Answer")
	proto.RegisterType((*ToClient)(nil), "internal.ToClient")
	proto.RegisterType((*ToClient_Block)(nil), "internal.ToClient.Block")
	proto.RegisterType((*ToClient_Query)(nil), "internal.ToClient.Query")
	proto.RegisterType((*ToClient_Restore)(nil), "internal.ToClient.Restore")
}

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 524 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x6f, 0x6f, 0xd3, 0x30,
	0x10, 0xc6, 0x93, 0x6c, 0x69, 0xb7, 0x5b, 0x41, 0x70, 0x2a, 0xa5, 0x64, 0xbc, 0x40, 0x91, 0x10,
	0x13, 0x12, 0xe9, 0xd4, 0x21, 0x90, 0x78, 0x47, 0x3b, 0xb4, 0x56, 0x02, 0x24, 0xdc, 0x7e, 0x01,
	0xb7, 0x39, 0x4a, 0xb4, 0xd4, 0xee, 0x1c, 0xa7, 0xa4, 0xdf, 0x0e, 0x89, 0x2f, 0xc2, 0x47, 0x41,
	0x76, 0xd2, 0x3f, 0xb0, 0x0e, 0xf1, 0xce, 0x97, 0xfb, 0xdd, 0x3d, 0xbe, 0x8b, 0x1f, 0x80, 0x99,
	0x5a, 0x4c, 0xa3, 0x85, 0x92, 0x5a, 0xe2, 0x51, 0x22, 0x34, 0x29, 0xc1, 0xd3, 0xe0, 0x74, 0x26,
	0xe5, 0x2c, 0xa5, 0x8e, 0xfd, 0x3e, 0xc9, 0xbf, 0x76, 0x68, 0xbe, 0xd0, 0xab, 0x12, 0x0b, 0x9f,
	0x02, 0x0c, 0x2f, 0x19, 0x65, 0x0b, 0x29, 0x32, 0xc2, 0xfb, 0xe0, 0x25, 0x71, 0xdb, 0x7d, 0xe6,
	0x9e, 0x1d, 0x33, 0x2f, 0x89, 0xc3, 0xe7, 0x70, 0x6f, 0xa4, 0xf9, 0x35, 0x6d, 0x80, 0x26, 0xf8,
	0x4b, 0x9e, 0xe6, 0x64, 0x19, 0x8f, 0x95, 0x41, 0x78, 0x05, 0x0f, 0x87, 0x95, 0xda, 0xb8, 0x60,
	0x74, 0x93, 0x53, 0xa6, 0xb1, 0x05, 0x35, 0x3e, 0x97, 0xb9, 0xd0, 0x96, 0x3d, 0x64, 0x55, 0x84,
	0x01, 0x1c, 0x29, 0x9a, 0x52, 0xb2, 0x24, 0xd5, 0xf6, 0xac, 0xd2, 0x26, 0x0e, 0x5f, 0x02, 0xee,
	0x36, 0xfa, 0xa7, 0xe8, 0x2f, 0x17, 0x8e, 0xc6, 0x72, 0x44, 0x6a, 0x49, 0x0a, 0x5f, 0x80, 0xa7,
	0x0b, 0x9b, 0x3f, 0xe9, 0x3e, 0x8a, 0xd6, 0xa3, 0x47, 0xeb, 0x7c, 0x34, 0x2e, 0x06, 0x0e, 0xf3,
	0x74, 0x81, 0x17, 0x50, 0xe3, 0x22, 0xfb, 0x5e, 0x69, 0x9f, 0x74, 0x9f, 0xec, 0x81, 0xdf, 0x5b,
	0x60, 0xe0, 0xb0, 0x0a, 0x0d, 0xda, 0xe0, 0x8d, 0x0b, 0x44, 0x38, 0x8c, 0xb9, 0xe6, 0x56, 0xa5,
	0xc1, 0xec, 0x39, 0x18, 0x41, 0xad, 0xa4, 0xf1, 0x01, 0x1c, 0xe4, 0xd5, 0xee, 0x1a, 0xcc, 0x1c,
	0xb1, 0x59, 0xf1, 0x46, 0xa8, 0x31, 0x70, 0xca, 0x0a, 0x6c, 0x81, 0x4f, 0x4a, 0x49, 0xd5, 0x3e,
	0x30, 0xb3, 0x0f, 0x1c, 0x56, 0x86, 0xbd, 0x63, 0xa8, 0x2f, 0xf8, 0x2a, 0x95, 0x3c, 0xee, 0xd5,
	0xc1, 0xa7, 0x25, 0x09, 0x1d, 0xfe, 0xf0, 0xcc, 0x88, 0xfd, 0x34, 0x21, 0xa1, 0xf1, 0x1c, 0xfc,
	0x49, 0x2a, 0xa7, 0xd7, 0xd5, 0x94, 0xed, 0xdd, 0x8b, 0x97, 0x48, 0xd4, 0x33, 0x79, 0xd3, 0xd2,
	0x82, 0xa6, 0xe2, 0x26, 0x27, 0xb5, 0xaa, 0x46, 0xdd, 0x57, 0xf1, 0xc5, 0xe4, 0x4d, 0x85, 0x05,
	0xf1, 0x0d, 0xd4, 0x15, 0x65, 0x5a, 0x2a, 0xb2, 0xd7, 0x3b, 0xe9, 0x06, 0x7b, 0x6a, 0x58, 0x49,
	0x0c, 0x1c, 0xb6, 0x86, 0x83, 0x57, 0xe0, 0x5b, 0xed, 0x3d, 0x5b, 0xc0, 0xdd, 0x2d, 0x54, 0x5b,
	0xeb, 0x80, 0x6f, 0x85, 0xf7, 0x2e, 0xcd, 0x4f, 0x44, 0x4c, 0x85, 0xe5, 0x0f, 0x58, 0x19, 0x04,
	0x1d, 0xa8, 0x57, 0xaa, 0xff, 0xa7, 0xb0, 0x59, 0x61, 0xf7, 0xa7, 0x0b, 0xf0, 0x89, 0x0b, 0x3e,
	0xa3, 0xb9, 0x59, 0xe2, 0x6b, 0xf0, 0x86, 0x97, 0xd8, 0x8a, 0x4a, 0x4b, 0x44, 0x6b, 0x4b, 0x44,
	0x1f, 0x8c, 0x25, 0x82, 0xe6, 0x76, 0xda, 0xad, 0x29, 0x42, 0x07, 0xdf, 0x81, 0x6f, 0x6d, 0x70,
	0x67, 0xe1, 0xe3, 0x6d, 0xe1, 0x1f, 0x7e, 0x09, 0x1d, 0xec, 0x03, 0x6c, 0x9f, 0x34, 0x9e, 0xee,
	0x28, 0xfc, 0xed, 0x98, 0xe0, 0x8e, 0xee, 0xa1, 0xd3, 0xbd, 0x82, 0xc6, 0x47, 0x3e, 0xfd, 0x46,
	0x59, 0x92, 0x7d, 0x96, 0x31, 0xe1, 0x5b, 0xa8, 0xf7, 0xa5, 0x10, 0x34, 0xd5, 0x88, 0xb7, 0x1f,
	0x70, 0x80, 0xb7, 0xff, 0x5a, 0xe8, 0x9c, 0xb9, 0xe7, 0xee, 0xa4, 0x66, 0x5b, 0x5f, 0xfc, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xf5, 0x9f, 0x3e, 0x79, 0x2a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ManagementClient is the client API for Management service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ManagementClient interface {
	ID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IDResponse, error)
	Stake(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StakeResponse, error)
	InternalTx(ctx context.Context, in *InternalTxRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type managementClient struct {
	cc *grpc.ClientConn
}

func NewManagementClient(cc *grpc.ClientConn) ManagementClient {
	return &managementClient{cc}
}

func (c *managementClient) ID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/internal.Management/ID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) Stake(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StakeResponse, error) {
	out := new(StakeResponse)
	err := c.cc.Invoke(ctx, "/internal.Management/Stake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *managementClient) InternalTx(ctx context.Context, in *InternalTxRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/internal.Management/InternalTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ManagementServer is the server API for Management service.
type ManagementServer interface {
	ID(context.Context, *empty.Empty) (*IDResponse, error)
	Stake(context.Context, *empty.Empty) (*StakeResponse, error)
	InternalTx(context.Context, *InternalTxRequest) (*empty.Empty, error)
}

// UnimplementedManagementServer can be embedded to have forward compatible implementations.
type UnimplementedManagementServer struct {
}

func (*UnimplementedManagementServer) ID(ctx context.Context, req *empty.Empty) (*IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ID not implemented")
}
func (*UnimplementedManagementServer) Stake(ctx context.Context, req *empty.Empty) (*StakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stake not implemented")
}
func (*UnimplementedManagementServer) InternalTx(ctx context.Context, req *InternalTxRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InternalTx not implemented")
}

func RegisterManagementServer(s *grpc.Server, srv ManagementServer) {
	s.RegisterService(&_Management_serviceDesc, srv)
}

func _Management_ID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).ID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Management/ID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).ID(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_Stake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).Stake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Management/Stake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).Stake(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Management_InternalTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalTxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ManagementServer).InternalTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/internal.Management/InternalTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ManagementServer).InternalTx(ctx, req.(*InternalTxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Management_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.Management",
	HandlerType: (*ManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ID",
			Handler:    _Management_ID_Handler,
		},
		{
			MethodName: "Stake",
			Handler:    _Management_Stake_Handler,
		},
		{
			MethodName: "InternalTx",
			Handler:    _Management_InternalTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc.proto",
}

// LachesisNodeClient is the client API for LachesisNode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LachesisNodeClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (LachesisNode_ConnectClient, error)
}

type lachesisNodeClient struct {
	cc *grpc.ClientConn
}

func NewLachesisNodeClient(cc *grpc.ClientConn) LachesisNodeClient {
	return &lachesisNodeClient{cc}
}

func (c *lachesisNodeClient) Connect(ctx context.Context, opts ...grpc.CallOption) (LachesisNode_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LachesisNode_serviceDesc.Streams[0], "/internal.LachesisNode/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &lachesisNodeConnectClient{stream}
	return x, nil
}

type LachesisNode_ConnectClient interface {
	Send(*ToServer) error
	Recv() (*ToClient, error)
	grpc.ClientStream
}

type lachesisNodeConnectClient struct {
	grpc.ClientStream
}

func (x *lachesisNodeConnectClient) Send(m *ToServer) error {
	return x.ClientStream.SendMsg(m)
}

func (x *lachesisNodeConnectClient) Recv() (*ToClient, error) {
	m := new(ToClient)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LachesisNodeServer is the server API for LachesisNode service.
type LachesisNodeServer interface {
	Connect(LachesisNode_ConnectServer) error
}

// UnimplementedLachesisNodeServer can be embedded to have forward compatible implementations.
type UnimplementedLachesisNodeServer struct {
}

func (*UnimplementedLachesisNodeServer) Connect(srv LachesisNode_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterLachesisNodeServer(s *grpc.Server, srv LachesisNodeServer) {
	s.RegisterService(&_LachesisNode_serviceDesc, srv)
}

func _LachesisNode_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LachesisNodeServer).Connect(&lachesisNodeConnectServer{stream})
}

type LachesisNode_ConnectServer interface {
	Send(*ToClient) error
	Recv() (*ToServer, error)
	grpc.ServerStream
}

type lachesisNodeConnectServer struct {
	grpc.ServerStream
}

func (x *lachesisNodeConnectServer) Send(m *ToClient) error {
	return x.ServerStream.SendMsg(m)
}

func (x *lachesisNodeConnectServer) Recv() (*ToServer, error) {
	m := new(ToServer)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LachesisNode_serviceDesc = grpc.ServiceDesc{
	ServiceName: "internal.LachesisNode",
	HandlerType: (*LachesisNodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _LachesisNode_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
