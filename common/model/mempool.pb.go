// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/mempool.proto

package model

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// Mempool represent the mempool data structure stored in the database
type MempoolTransaction struct {
	ID                          int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BlockHeight                 uint32   `protobuf:"varint,2,opt,name=BlockHeight,proto3" json:"BlockHeight,omitempty"`
	FeePerByte                  int64    `protobuf:"varint,3,opt,name=FeePerByte,proto3" json:"FeePerByte,omitempty"`
	ArrivalTimestamp            int64    `protobuf:"varint,4,opt,name=ArrivalTimestamp,proto3" json:"ArrivalTimestamp,omitempty"`
	TransactionBytes            []byte   `protobuf:"bytes,5,opt,name=TransactionBytes,proto3" json:"TransactionBytes,omitempty"`
	SenderAccountAddressType    []byte   `protobuf:"bytes,6,opt,name=SenderAccountAddressType,proto3" json:"SenderAccountAddressType,omitempty"`
	SenderAccountAddress        string   `protobuf:"bytes,7,opt,name=SenderAccountAddress,proto3" json:"SenderAccountAddress,omitempty"`
	RecipientAccountAddressType []byte   `protobuf:"bytes,8,opt,name=RecipientAccountAddressType,proto3" json:"RecipientAccountAddressType,omitempty"`
	RecipientAccountAddress     string   `protobuf:"bytes,9,opt,name=RecipientAccountAddress,proto3" json:"RecipientAccountAddress,omitempty"`
	Escrow                      *Escrow  `protobuf:"bytes,10,opt,name=Escrow,proto3" json:"Escrow,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *MempoolTransaction) Reset()         { *m = MempoolTransaction{} }
func (m *MempoolTransaction) String() string { return proto.CompactTextString(m) }
func (*MempoolTransaction) ProtoMessage()    {}
func (*MempoolTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ea31ac6d427b7b, []int{0}
}

func (m *MempoolTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MempoolTransaction.Unmarshal(m, b)
}
func (m *MempoolTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MempoolTransaction.Marshal(b, m, deterministic)
}
func (m *MempoolTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MempoolTransaction.Merge(m, src)
}
func (m *MempoolTransaction) XXX_Size() int {
	return xxx_messageInfo_MempoolTransaction.Size(m)
}
func (m *MempoolTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MempoolTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MempoolTransaction proto.InternalMessageInfo

func (m *MempoolTransaction) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *MempoolTransaction) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MempoolTransaction) GetFeePerByte() int64 {
	if m != nil {
		return m.FeePerByte
	}
	return 0
}

func (m *MempoolTransaction) GetArrivalTimestamp() int64 {
	if m != nil {
		return m.ArrivalTimestamp
	}
	return 0
}

func (m *MempoolTransaction) GetTransactionBytes() []byte {
	if m != nil {
		return m.TransactionBytes
	}
	return nil
}

func (m *MempoolTransaction) GetSenderAccountAddressType() []byte {
	if m != nil {
		return m.SenderAccountAddressType
	}
	return nil
}

func (m *MempoolTransaction) GetSenderAccountAddress() string {
	if m != nil {
		return m.SenderAccountAddress
	}
	return ""
}

func (m *MempoolTransaction) GetRecipientAccountAddressType() []byte {
	if m != nil {
		return m.RecipientAccountAddressType
	}
	return nil
}

func (m *MempoolTransaction) GetRecipientAccountAddress() string {
	if m != nil {
		return m.RecipientAccountAddress
	}
	return ""
}

func (m *MempoolTransaction) GetEscrow() *Escrow {
	if m != nil {
		return m.Escrow
	}
	return nil
}

type GetMempoolTransactionRequest struct {
	// Fetch Mempool Transaction by its TransactionBytes
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetMempoolTransactionRequest) Reset()         { *m = GetMempoolTransactionRequest{} }
func (m *GetMempoolTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*GetMempoolTransactionRequest) ProtoMessage()    {}
func (*GetMempoolTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ea31ac6d427b7b, []int{1}
}

func (m *GetMempoolTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMempoolTransactionRequest.Unmarshal(m, b)
}
func (m *GetMempoolTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMempoolTransactionRequest.Marshal(b, m, deterministic)
}
func (m *GetMempoolTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMempoolTransactionRequest.Merge(m, src)
}
func (m *GetMempoolTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_GetMempoolTransactionRequest.Size(m)
}
func (m *GetMempoolTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMempoolTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMempoolTransactionRequest proto.InternalMessageInfo

func (m *GetMempoolTransactionRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

type GetMempoolTransactionResponse struct {
	Transaction          *MempoolTransaction `protobuf:"bytes,1,opt,name=transaction,proto3" json:"transaction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *GetMempoolTransactionResponse) Reset()         { *m = GetMempoolTransactionResponse{} }
func (m *GetMempoolTransactionResponse) String() string { return proto.CompactTextString(m) }
func (*GetMempoolTransactionResponse) ProtoMessage()    {}
func (*GetMempoolTransactionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ea31ac6d427b7b, []int{2}
}

func (m *GetMempoolTransactionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMempoolTransactionResponse.Unmarshal(m, b)
}
func (m *GetMempoolTransactionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMempoolTransactionResponse.Marshal(b, m, deterministic)
}
func (m *GetMempoolTransactionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMempoolTransactionResponse.Merge(m, src)
}
func (m *GetMempoolTransactionResponse) XXX_Size() int {
	return xxx_messageInfo_GetMempoolTransactionResponse.Size(m)
}
func (m *GetMempoolTransactionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMempoolTransactionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMempoolTransactionResponse proto.InternalMessageInfo

func (m *GetMempoolTransactionResponse) GetTransaction() *MempoolTransaction {
	if m != nil {
		return m.Transaction
	}
	return nil
}

type GetMempoolTransactionsRequest struct {
	// Fetch Mempool transactions from arrival timestamp start from n
	TimestampStart int64 `protobuf:"varint,1,opt,name=timestampStart,proto3" json:"timestampStart,omitempty"`
	// Fetch Mempool transactions to arrival timestamp end until n
	TimestampEnd int64 `protobuf:"varint,2,opt,name=timestampEnd,proto3" json:"timestampEnd,omitempty"`
	// SenderAddress and RecipientAddress
	Address              string      `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Pagination           *Pagination `protobuf:"bytes,4,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetMempoolTransactionsRequest) Reset()         { *m = GetMempoolTransactionsRequest{} }
func (m *GetMempoolTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetMempoolTransactionsRequest) ProtoMessage()    {}
func (*GetMempoolTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ea31ac6d427b7b, []int{3}
}

func (m *GetMempoolTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMempoolTransactionsRequest.Unmarshal(m, b)
}
func (m *GetMempoolTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMempoolTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *GetMempoolTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMempoolTransactionsRequest.Merge(m, src)
}
func (m *GetMempoolTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetMempoolTransactionsRequest.Size(m)
}
func (m *GetMempoolTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMempoolTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetMempoolTransactionsRequest proto.InternalMessageInfo

func (m *GetMempoolTransactionsRequest) GetTimestampStart() int64 {
	if m != nil {
		return m.TimestampStart
	}
	return 0
}

func (m *GetMempoolTransactionsRequest) GetTimestampEnd() int64 {
	if m != nil {
		return m.TimestampEnd
	}
	return 0
}

func (m *GetMempoolTransactionsRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *GetMempoolTransactionsRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type GetMempoolTransactionsResponse struct {
	Total                uint64                `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	MempoolTransactions  []*MempoolTransaction `protobuf:"bytes,2,rep,name=MempoolTransactions,proto3" json:"MempoolTransactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *GetMempoolTransactionsResponse) Reset()         { *m = GetMempoolTransactionsResponse{} }
func (m *GetMempoolTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetMempoolTransactionsResponse) ProtoMessage()    {}
func (*GetMempoolTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_22ea31ac6d427b7b, []int{4}
}

func (m *GetMempoolTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetMempoolTransactionsResponse.Unmarshal(m, b)
}
func (m *GetMempoolTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetMempoolTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *GetMempoolTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetMempoolTransactionsResponse.Merge(m, src)
}
func (m *GetMempoolTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetMempoolTransactionsResponse.Size(m)
}
func (m *GetMempoolTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetMempoolTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetMempoolTransactionsResponse proto.InternalMessageInfo

func (m *GetMempoolTransactionsResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetMempoolTransactionsResponse) GetMempoolTransactions() []*MempoolTransaction {
	if m != nil {
		return m.MempoolTransactions
	}
	return nil
}

func init() {
	proto.RegisterType((*MempoolTransaction)(nil), "model.MempoolTransaction")
	proto.RegisterType((*GetMempoolTransactionRequest)(nil), "model.GetMempoolTransactionRequest")
	proto.RegisterType((*GetMempoolTransactionResponse)(nil), "model.GetMempoolTransactionResponse")
	proto.RegisterType((*GetMempoolTransactionsRequest)(nil), "model.GetMempoolTransactionsRequest")
	proto.RegisterType((*GetMempoolTransactionsResponse)(nil), "model.GetMempoolTransactionsResponse")
}

func init() {
	proto.RegisterFile("model/mempool.proto", fileDescriptor_22ea31ac6d427b7b)
}

var fileDescriptor_22ea31ac6d427b7b = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xcf, 0x6e, 0xd3, 0x40,
	0x10, 0xc6, 0x65, 0xbb, 0x49, 0xe9, 0xb8, 0x45, 0x65, 0x5a, 0xc1, 0x52, 0xfe, 0xc8, 0xb2, 0x04,
	0xb2, 0x22, 0xe1, 0x80, 0xb9, 0x20, 0xb8, 0x90, 0xa8, 0x05, 0x2a, 0x84, 0x54, 0xb9, 0x39, 0x21,
	0x2e, 0xce, 0x7a, 0x94, 0x5a, 0xd8, 0x5e, 0xb3, 0xbb, 0x01, 0x95, 0x2b, 0x2f, 0xc6, 0x8b, 0xf0,
	0x2e, 0xa8, 0x6b, 0x27, 0x35, 0xc4, 0xee, 0xc5, 0x92, 0xbf, 0xdf, 0x7c, 0x33, 0xb3, 0x33, 0xbb,
	0x70, 0x50, 0x88, 0x94, 0xf2, 0x71, 0x41, 0x45, 0x25, 0x44, 0x1e, 0x56, 0x52, 0x68, 0x81, 0x03,
	0x23, 0x1e, 0xdd, 0xad, 0x59, 0x95, 0x2c, 0xb2, 0x32, 0xd1, 0x99, 0x28, 0x6b, 0x7c, 0x84, 0xb5,
	0x4e, 0x8a, 0x4b, 0xf1, 0xa3, 0xd6, 0xfc, 0x3f, 0x0e, 0xe0, 0xa7, 0x3a, 0xc9, 0x4c, 0x26, 0xa5,
	0x4a, 0xf8, 0x95, 0x01, 0x11, 0xec, 0xd3, 0x63, 0x66, 0x79, 0x56, 0xe0, 0x4c, 0xed, 0xe7, 0x56,
	0x6c, 0x9f, 0x1e, 0xa3, 0x07, 0xee, 0x34, 0x17, 0xfc, 0xeb, 0x07, 0xca, 0x16, 0x17, 0x9a, 0xd9,
	0x9e, 0x15, 0xec, 0xc5, 0x6d, 0x09, 0x7d, 0x80, 0x77, 0x44, 0x67, 0x24, 0xa7, 0x97, 0x9a, 0x98,
	0xb3, 0x76, 0xb7, 0x54, 0x0c, 0x61, 0x7f, 0x22, 0x65, 0xf6, 0x3d, 0xc9, 0x67, 0x59, 0x41, 0x4a,
	0x27, 0x45, 0xc5, 0xb6, 0xd6, 0x91, 0x1b, 0x0c, 0x47, 0xb0, 0xdf, 0x6a, 0xec, 0x2a, 0x85, 0x62,
	0x03, 0xcf, 0x0a, 0x76, 0xe3, 0x0d, 0x1d, 0x5f, 0x03, 0x3b, 0xa7, 0x32, 0x25, 0x39, 0xe1, 0x5c,
	0x2c, 0x4b, 0x3d, 0x49, 0x53, 0x49, 0x4a, 0xcd, 0x2e, 0x2b, 0x62, 0x43, 0xe3, 0xe9, 0xe5, 0x18,
	0xc1, 0x61, 0x17, 0x63, 0xdb, 0x9e, 0x15, 0xec, 0xc4, 0x9d, 0x0c, 0xdf, 0xc2, 0x83, 0x98, 0x78,
	0x56, 0x65, 0x54, 0xea, 0x8e, 0x92, 0xb7, 0x4c, 0xc9, 0x9b, 0x42, 0xf0, 0x15, 0xdc, 0xeb, 0xc1,
	0x6c, 0xc7, 0x14, 0xee, 0xc3, 0xf8, 0x04, 0x86, 0x27, 0x66, 0x91, 0x0c, 0x3c, 0x2b, 0x70, 0xa3,
	0xbd, 0xd0, 0x6c, 0x37, 0xac, 0xc5, 0xb8, 0x81, 0x7e, 0x04, 0x0f, 0xdf, 0x93, 0xde, 0xdc, 0x70,
	0x4c, 0xdf, 0x96, 0xa4, 0x74, 0xd7, 0xa2, 0xfd, 0x2f, 0xf0, 0xa8, 0xc7, 0xa3, 0x2a, 0x51, 0x2a,
	0xc2, 0x37, 0xe0, 0xea, 0x6b, 0xd9, 0xb8, 0xdd, 0xe8, 0x7e, 0xd3, 0x40, 0x87, 0xaf, 0x1d, 0xed,
	0xff, 0xb6, 0x7a, 0xd2, 0xab, 0x55, 0x4f, 0x23, 0xb8, 0xad, 0x57, 0xfb, 0x3f, 0xd7, 0x89, 0xd4,
	0xad, 0xfe, 0xfe, 0x23, 0xf8, 0x14, 0x76, 0xd7, 0xca, 0x49, 0x99, 0x9a, 0x5b, 0x59, 0x47, 0xfe,
	0xa3, 0x23, 0x83, 0xed, 0xa4, 0x19, 0xac, 0x63, 0x06, 0xbb, 0xfa, 0xc5, 0x17, 0x00, 0x67, 0xeb,
	0x97, 0x62, 0xae, 0xa2, 0x1b, 0xdd, 0x69, 0xce, 0x72, 0x0d, 0xe2, 0x56, 0x90, 0xff, 0xcb, 0x82,
	0xc7, 0x7d, 0x47, 0x68, 0x46, 0x74, 0x08, 0x83, 0x99, 0xd0, 0x49, 0x6e, 0x5a, 0xdf, 0x8a, 0xeb,
	0x1f, 0xfc, 0x08, 0x07, 0x1d, 0x26, 0x66, 0x7b, 0xce, 0xcd, 0x03, 0xec, 0x72, 0x4d, 0x47, 0x9f,
	0x83, 0x45, 0xa6, 0x2f, 0x96, 0xf3, 0x90, 0x8b, 0x62, 0xfc, 0x53, 0x88, 0x39, 0xaf, 0xbf, 0xcf,
	0xb8, 0x90, 0x34, 0xe6, 0xa2, 0x28, 0x44, 0x39, 0x36, 0x39, 0xe7, 0x43, 0xf3, 0xda, 0x5f, 0xfe,
	0x0d, 0x00, 0x00, 0xff, 0xff, 0x6a, 0x04, 0x34, 0xf6, 0x37, 0x04, 0x00, 0x00,
}
