// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/escrow.proto

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

type EscrowStatus int32

const (
	EscrowStatus_Pending  EscrowStatus = 0
	EscrowStatus_Approved EscrowStatus = 1
	EscrowStatus_Rejected EscrowStatus = 2
	EscrowStatus_Expired  EscrowStatus = 3
)

var EscrowStatus_name = map[int32]string{
	0: "Pending",
	1: "Approved",
	2: "Rejected",
	3: "Expired",
}

var EscrowStatus_value = map[string]int32{
	"Pending":  0,
	"Approved": 1,
	"Rejected": 2,
	"Expired":  3,
}

func (x EscrowStatus) String() string {
	return proto.EnumName(EscrowStatus_name, int32(x))
}

func (EscrowStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4ffdfca00fa52ba, []int{0}
}

type EscrowApproval int32

const (
	EscrowApproval_Approve EscrowApproval = 0
	EscrowApproval_Reject  EscrowApproval = 1
	EscrowApproval_Expire  EscrowApproval = 2
)

var EscrowApproval_name = map[int32]string{
	0: "Approve",
	1: "Reject",
	2: "Expire",
}

var EscrowApproval_value = map[string]int32{
	"Approve": 0,
	"Reject":  1,
	"Expire":  2,
}

func (x EscrowApproval) String() string {
	return proto.EnumName(EscrowApproval_name, int32(x))
}

func (EscrowApproval) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c4ffdfca00fa52ba, []int{1}
}

type Escrow struct {
	ID               int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	SenderAddress    string `protobuf:"bytes,2,opt,name=SenderAddress,proto3" json:"SenderAddress,omitempty"`
	RecipientAddress string `protobuf:"bytes,3,opt,name=RecipientAddress,proto3" json:"RecipientAddress,omitempty"`
	ApproverAddress  string `protobuf:"bytes,4,opt,name=ApproverAddress,proto3" json:"ApproverAddress,omitempty"`
	Amount           int64  `protobuf:"varint,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Commission       int64  `protobuf:"varint,6,opt,name=Commission,proto3" json:"Commission,omitempty"`
	// Timeout is BlockHeight gap
	Timeout              uint64       `protobuf:"varint,7,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
	Status               EscrowStatus `protobuf:"varint,8,opt,name=Status,proto3,enum=model.EscrowStatus" json:"Status,omitempty"`
	BlockHeight          uint32       `protobuf:"varint,9,opt,name=BlockHeight,proto3" json:"BlockHeight,omitempty"`
	Latest               bool         `protobuf:"varint,10,opt,name=Latest,proto3" json:"Latest,omitempty"`
	Instruction          string       `protobuf:"bytes,11,opt,name=Instruction,proto3" json:"Instruction,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Escrow) Reset()         { *m = Escrow{} }
func (m *Escrow) String() string { return proto.CompactTextString(m) }
func (*Escrow) ProtoMessage()    {}
func (*Escrow) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ffdfca00fa52ba, []int{0}
}

func (m *Escrow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Escrow.Unmarshal(m, b)
}
func (m *Escrow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Escrow.Marshal(b, m, deterministic)
}
func (m *Escrow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Escrow.Merge(m, src)
}
func (m *Escrow) XXX_Size() int {
	return xxx_messageInfo_Escrow.Size(m)
}
func (m *Escrow) XXX_DiscardUnknown() {
	xxx_messageInfo_Escrow.DiscardUnknown(m)
}

var xxx_messageInfo_Escrow proto.InternalMessageInfo

func (m *Escrow) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Escrow) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *Escrow) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *Escrow) GetApproverAddress() string {
	if m != nil {
		return m.ApproverAddress
	}
	return ""
}

func (m *Escrow) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Escrow) GetCommission() int64 {
	if m != nil {
		return m.Commission
	}
	return 0
}

func (m *Escrow) GetTimeout() uint64 {
	if m != nil {
		return m.Timeout
	}
	return 0
}

func (m *Escrow) GetStatus() EscrowStatus {
	if m != nil {
		return m.Status
	}
	return EscrowStatus_Pending
}

func (m *Escrow) GetBlockHeight() uint32 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *Escrow) GetLatest() bool {
	if m != nil {
		return m.Latest
	}
	return false
}

func (m *Escrow) GetInstruction() string {
	if m != nil {
		return m.Instruction
	}
	return ""
}

// GetEscrowTransactionsRequest message for get escrow transactions
type GetEscrowTransactionsRequest struct {
	ApproverAddress      string         `protobuf:"bytes,1,opt,name=ApproverAddress,proto3" json:"ApproverAddress,omitempty"`
	SenderAddress        string         `protobuf:"bytes,2,opt,name=SenderAddress,proto3" json:"SenderAddress,omitempty"`
	RecipientAddress     string         `protobuf:"bytes,3,opt,name=RecipientAddress,proto3" json:"RecipientAddress,omitempty"`
	ID                   int64          `protobuf:"varint,4,opt,name=ID,proto3" json:"ID,omitempty"`
	Statuses             []EscrowStatus `protobuf:"varint,5,rep,packed,name=Statuses,proto3,enum=model.EscrowStatus" json:"Statuses,omitempty"`
	BlockHeightStart     uint32         `protobuf:"varint,6,opt,name=BlockHeightStart,proto3" json:"BlockHeightStart,omitempty"`
	BlockHeightEnd       uint32         `protobuf:"varint,7,opt,name=BlockHeightEnd,proto3" json:"BlockHeightEnd,omitempty"`
	Latest               bool           `protobuf:"varint,8,opt,name=Latest,proto3" json:"Latest,omitempty"`
	Pagination           *Pagination    `protobuf:"bytes,9,opt,name=Pagination,proto3" json:"Pagination,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *GetEscrowTransactionsRequest) Reset()         { *m = GetEscrowTransactionsRequest{} }
func (m *GetEscrowTransactionsRequest) String() string { return proto.CompactTextString(m) }
func (*GetEscrowTransactionsRequest) ProtoMessage()    {}
func (*GetEscrowTransactionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ffdfca00fa52ba, []int{1}
}

func (m *GetEscrowTransactionsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEscrowTransactionsRequest.Unmarshal(m, b)
}
func (m *GetEscrowTransactionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEscrowTransactionsRequest.Marshal(b, m, deterministic)
}
func (m *GetEscrowTransactionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEscrowTransactionsRequest.Merge(m, src)
}
func (m *GetEscrowTransactionsRequest) XXX_Size() int {
	return xxx_messageInfo_GetEscrowTransactionsRequest.Size(m)
}
func (m *GetEscrowTransactionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEscrowTransactionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEscrowTransactionsRequest proto.InternalMessageInfo

func (m *GetEscrowTransactionsRequest) GetApproverAddress() string {
	if m != nil {
		return m.ApproverAddress
	}
	return ""
}

func (m *GetEscrowTransactionsRequest) GetSenderAddress() string {
	if m != nil {
		return m.SenderAddress
	}
	return ""
}

func (m *GetEscrowTransactionsRequest) GetRecipientAddress() string {
	if m != nil {
		return m.RecipientAddress
	}
	return ""
}

func (m *GetEscrowTransactionsRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *GetEscrowTransactionsRequest) GetStatuses() []EscrowStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (m *GetEscrowTransactionsRequest) GetBlockHeightStart() uint32 {
	if m != nil {
		return m.BlockHeightStart
	}
	return 0
}

func (m *GetEscrowTransactionsRequest) GetBlockHeightEnd() uint32 {
	if m != nil {
		return m.BlockHeightEnd
	}
	return 0
}

func (m *GetEscrowTransactionsRequest) GetLatest() bool {
	if m != nil {
		return m.Latest
	}
	return false
}

func (m *GetEscrowTransactionsRequest) GetPagination() *Pagination {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// GetEscrowTransactionsResponse returns fields of GetEscrowTransactionsRequest
type GetEscrowTransactionsResponse struct {
	// Number of transactions in total
	Total uint64 `protobuf:"varint,1,opt,name=Total,proto3" json:"Total,omitempty"`
	// Transaction transactions returned
	Escrows              []*Escrow `protobuf:"bytes,2,rep,name=Escrows,proto3" json:"Escrows,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GetEscrowTransactionsResponse) Reset()         { *m = GetEscrowTransactionsResponse{} }
func (m *GetEscrowTransactionsResponse) String() string { return proto.CompactTextString(m) }
func (*GetEscrowTransactionsResponse) ProtoMessage()    {}
func (*GetEscrowTransactionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ffdfca00fa52ba, []int{2}
}

func (m *GetEscrowTransactionsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEscrowTransactionsResponse.Unmarshal(m, b)
}
func (m *GetEscrowTransactionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEscrowTransactionsResponse.Marshal(b, m, deterministic)
}
func (m *GetEscrowTransactionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEscrowTransactionsResponse.Merge(m, src)
}
func (m *GetEscrowTransactionsResponse) XXX_Size() int {
	return xxx_messageInfo_GetEscrowTransactionsResponse.Size(m)
}
func (m *GetEscrowTransactionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEscrowTransactionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetEscrowTransactionsResponse proto.InternalMessageInfo

func (m *GetEscrowTransactionsResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *GetEscrowTransactionsResponse) GetEscrows() []*Escrow {
	if m != nil {
		return m.Escrows
	}
	return nil
}

// GetEscrowTransactionRequest represents GetEscrowTransaction parameter
type GetEscrowTransactionRequest struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetEscrowTransactionRequest) Reset()         { *m = GetEscrowTransactionRequest{} }
func (m *GetEscrowTransactionRequest) String() string { return proto.CompactTextString(m) }
func (*GetEscrowTransactionRequest) ProtoMessage()    {}
func (*GetEscrowTransactionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4ffdfca00fa52ba, []int{3}
}

func (m *GetEscrowTransactionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetEscrowTransactionRequest.Unmarshal(m, b)
}
func (m *GetEscrowTransactionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetEscrowTransactionRequest.Marshal(b, m, deterministic)
}
func (m *GetEscrowTransactionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetEscrowTransactionRequest.Merge(m, src)
}
func (m *GetEscrowTransactionRequest) XXX_Size() int {
	return xxx_messageInfo_GetEscrowTransactionRequest.Size(m)
}
func (m *GetEscrowTransactionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetEscrowTransactionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetEscrowTransactionRequest proto.InternalMessageInfo

func (m *GetEscrowTransactionRequest) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func init() {
	proto.RegisterEnum("model.EscrowStatus", EscrowStatus_name, EscrowStatus_value)
	proto.RegisterEnum("model.EscrowApproval", EscrowApproval_name, EscrowApproval_value)
	proto.RegisterType((*Escrow)(nil), "model.Escrow")
	proto.RegisterType((*GetEscrowTransactionsRequest)(nil), "model.GetEscrowTransactionsRequest")
	proto.RegisterType((*GetEscrowTransactionsResponse)(nil), "model.GetEscrowTransactionsResponse")
	proto.RegisterType((*GetEscrowTransactionRequest)(nil), "model.GetEscrowTransactionRequest")
}

func init() {
	proto.RegisterFile("model/escrow.proto", fileDescriptor_c4ffdfca00fa52ba)
}

var fileDescriptor_c4ffdfca00fa52ba = []byte{
	// 536 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x94, 0x5f, 0x8f, 0xd2, 0x40,
	0x14, 0xc5, 0x77, 0x0a, 0x14, 0xf6, 0xb2, 0x60, 0x1d, 0x93, 0xcd, 0x64, 0x5d, 0x93, 0x86, 0x18,
	0x6d, 0x30, 0x52, 0x17, 0xe3, 0x07, 0x00, 0x21, 0x4a, 0xe2, 0xc3, 0x66, 0xe0, 0xc9, 0xb7, 0xd2,
	0x4e, 0xd8, 0x51, 0x3a, 0x53, 0x3b, 0x53, 0x35, 0x7e, 0x78, 0x35, 0x9d, 0x29, 0x6c, 0xd9, 0x65,
	0xdf, 0xf6, 0x85, 0x30, 0xe7, 0x9c, 0x3b, 0x7f, 0x7e, 0xf7, 0x02, 0xe0, 0x54, 0x26, 0x6c, 0x1b,
	0x32, 0x15, 0xe7, 0xf2, 0xd7, 0x28, 0xcb, 0xa5, 0x96, 0xb8, 0x65, 0xb4, 0x8b, 0x73, 0x6b, 0x65,
	0xd1, 0x86, 0x8b, 0x48, 0x73, 0x29, 0xac, 0x3d, 0xf8, 0xe7, 0x80, 0x3b, 0x37, 0x79, 0x8c, 0xc1,
	0x59, 0xcc, 0x08, 0xf2, 0x51, 0xd0, 0x98, 0x3a, 0xef, 0x10, 0x75, 0x16, 0x33, 0xfc, 0x12, 0x7a,
	0x4b, 0x26, 0x12, 0x96, 0x4f, 0x92, 0x24, 0x67, 0x4a, 0x11, 0xc7, 0x47, 0xc1, 0x29, 0x3d, 0x14,
	0xf1, 0x10, 0x3c, 0xca, 0x62, 0x9e, 0x71, 0x26, 0xf4, 0x2e, 0xd8, 0x30, 0xc1, 0x7b, 0x3a, 0x0e,
	0xe0, 0xc9, 0x24, 0xcb, 0x72, 0xf9, 0xf3, 0x76, 0xcf, 0xa6, 0x89, 0xde, 0x95, 0xf1, 0x05, 0xb8,
	0x93, 0x54, 0x16, 0x42, 0x93, 0xd6, 0xfe, 0x4e, 0x95, 0x82, 0x07, 0x00, 0x1f, 0x65, 0x9a, 0x72,
	0xa5, 0xb8, 0x14, 0xc4, 0xdd, 0xfb, 0x35, 0x15, 0x5f, 0x42, 0x7b, 0xc5, 0x53, 0x26, 0x0b, 0x4d,
	0xda, 0x3e, 0x0a, 0x9a, 0x26, 0xb0, 0x93, 0xf0, 0x1b, 0x70, 0x97, 0x3a, 0xd2, 0x85, 0x22, 0x1d,
	0x1f, 0x05, 0xfd, 0xf1, 0xb3, 0x91, 0x21, 0x34, 0xb2, 0x30, 0xac, 0x45, 0xab, 0x08, 0xf6, 0xa1,
	0x3b, 0xdd, 0xca, 0xf8, 0xfb, 0x67, 0xc6, 0x37, 0x37, 0x9a, 0x9c, 0xfa, 0x28, 0xe8, 0xd1, 0xba,
	0x84, 0xcf, 0xc1, 0xfd, 0x12, 0x69, 0xa6, 0x34, 0x01, 0x1f, 0x05, 0x1d, 0x5a, 0xad, 0xca, 0xca,
	0x85, 0x50, 0x3a, 0x2f, 0xe2, 0x12, 0x3a, 0xe9, 0x9a, 0xa7, 0xd6, 0xa5, 0xc1, 0x5f, 0x07, 0x2e,
	0x3f, 0x31, 0x6d, 0xcf, 0x5d, 0xe5, 0x91, 0x50, 0x91, 0x31, 0x14, 0x65, 0x3f, 0x8a, 0x72, 0x8b,
	0x23, 0xc4, 0xd0, 0x71, 0x62, 0x8f, 0xdf, 0x2d, 0x3b, 0x13, 0xcd, 0x83, 0x99, 0x08, 0xa1, 0x63,
	0xb1, 0x30, 0x45, 0x5a, 0x7e, 0xe3, 0x21, 0x76, 0xfb, 0x50, 0x79, 0x60, 0x0d, 0xd5, 0x52, 0x47,
	0xb9, 0x36, 0x2d, 0xeb, 0xd1, 0x7b, 0x3a, 0x7e, 0x05, 0xfd, 0x9a, 0x36, 0x17, 0x89, 0xe9, 0x5d,
	0x8f, 0xde, 0x51, 0x6b, 0xbc, 0x3b, 0x07, 0xbc, 0xaf, 0x00, 0xae, 0xf7, 0x33, 0x6e, 0x1a, 0xd5,
	0x1d, 0x3f, 0xad, 0xae, 0x77, 0x6b, 0xd0, 0x5a, 0x68, 0xb0, 0x86, 0x17, 0x0f, 0xf0, 0x57, 0x99,
	0x14, 0x8a, 0x61, 0x02, 0xad, 0x95, 0xd4, 0xd1, 0xd6, 0x60, 0xb7, 0x63, 0x64, 0x05, 0xfc, 0x1a,
	0xda, 0xb6, 0xae, 0x44, 0xdd, 0x08, 0xba, 0xe3, 0xde, 0x01, 0x09, 0xba, 0x73, 0x07, 0x57, 0xf0,
	0xfc, 0xd8, 0x19, 0xbb, 0x16, 0x1f, 0xf9, 0xe9, 0x0d, 0x67, 0x70, 0x56, 0xe7, 0x89, 0xbb, 0xd0,
	0xbe, 0x66, 0x22, 0xe1, 0x62, 0xe3, 0x9d, 0xe0, 0x33, 0xe8, 0x54, 0xcd, 0x4f, 0x3c, 0x54, 0xae,
	0x28, 0xfb, 0xc6, 0x62, 0xcd, 0x12, 0xcf, 0x29, 0x83, 0xf3, 0xdf, 0x19, 0xcf, 0x59, 0xe2, 0x35,
	0x86, 0x1f, 0xa0, 0x6f, 0x77, 0xb1, 0xf1, 0x68, 0x5b, 0xda, 0x55, 0xa9, 0x77, 0x82, 0x01, 0x5c,
	0x5b, 0xe9, 0xa1, 0xf2, 0xbb, 0xad, 0xf3, 0x9c, 0xe9, 0xf0, 0x6b, 0xb0, 0xe1, 0xfa, 0xa6, 0x58,
	0x8f, 0x62, 0x99, 0x86, 0x7f, 0xa4, 0x5c, 0xc7, 0xf6, 0xf3, 0x6d, 0x2c, 0x73, 0x16, 0xc6, 0x32,
	0x4d, 0xa5, 0x08, 0xcd, 0x5b, 0xd7, 0xae, 0xf9, 0x27, 0x79, 0xff, 0x3f, 0x00, 0x00, 0xff, 0xff,
	0x3f, 0x31, 0xa3, 0x19, 0x7e, 0x04, 0x00, 0x00,
}
