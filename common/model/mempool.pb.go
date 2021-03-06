// ZooBC Copyright (C) 2020 Quasisoft Limited - Hong Kong
// This file is part of ZooBC <https://github.com/zoobc/zoobc-core>
//
// ZooBC is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ZooBC is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ZooBC.  If not, see <http://www.gnu.org/licenses/>.
//
// Additional Permission Under GNU GPL Version 3 section 7.
// As the special exception permitted under Section 7b, c and e,
// in respect with the Author’s copyright, please refer to this section:
//
// 1. You are free to convey this Program according to GNU GPL Version 3,
//     as long as you respect and comply with the Author’s copyright by
//     showing in its user interface an Appropriate Notice that the derivate
//     program and its source code are “powered by ZooBC”.
//     This is an acknowledgement for the copyright holder, ZooBC,
//     as the implementation of appreciation of the exclusive right of the
//     creator and to avoid any circumvention on the rights under trademark
//     law for use of some trade names, trademarks, or service marks.
//
// 2. Complying to the GNU GPL Version 3, you may distribute
//     the program without any permission from the Author.
//     However a prior notification to the authors will be appreciated.
//
// ZooBC is architected by Roberto Capodieci & Barton Johnston
//             contact us at roberto.capodieci[at]blockchainzoo.com
//             and barton.johnston[at]blockchainzoo.com
//
// Core developers that contributed to the current implementation of the
// software are:
//             Ahmad Ali Abdilah ahmad.abdilah[at]blockchainzoo.com
//             Allan Bintoro allan.bintoro[at]blockchainzoo.com
//             Andy Herman
//             Gede Sukra
//             Ketut Ariasa
//             Nawi Kartini nawi.kartini[at]blockchainzoo.com
//             Stefano Galassi stefano.galassi[at]blockchainzoo.com
//
// IMPORTANT: The above copyright notice and this permission notice
// shall be included in all copies or substantial portions of the Software.
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
	ID                      int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BlockHeight             uint32   `protobuf:"varint,2,opt,name=BlockHeight,proto3" json:"BlockHeight,omitempty"`
	FeePerByte              int64    `protobuf:"varint,3,opt,name=FeePerByte,proto3" json:"FeePerByte,omitempty"`
	ArrivalTimestamp        int64    `protobuf:"varint,4,opt,name=ArrivalTimestamp,proto3" json:"ArrivalTimestamp,omitempty"`
	TransactionBytes        []byte   `protobuf:"bytes,5,opt,name=TransactionBytes,proto3" json:"TransactionBytes,omitempty"`
	SenderAccountAddress    []byte   `protobuf:"bytes,6,opt,name=SenderAccountAddress,proto3" json:"SenderAccountAddress,omitempty"`
	RecipientAccountAddress []byte   `protobuf:"bytes,7,opt,name=RecipientAccountAddress,proto3" json:"RecipientAccountAddress,omitempty"`
	Escrow                  *Escrow  `protobuf:"bytes,8,opt,name=Escrow,proto3" json:"Escrow,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
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

func (m *MempoolTransaction) GetSenderAccountAddress() []byte {
	if m != nil {
		return m.SenderAccountAddress
	}
	return nil
}

func (m *MempoolTransaction) GetRecipientAccountAddress() []byte {
	if m != nil {
		return m.RecipientAccountAddress
	}
	return nil
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
	Address              []byte      `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
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

func (m *GetMempoolTransactionsRequest) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
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
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xd1, 0x6a, 0xd4, 0x40,
	0x14, 0x86, 0x49, 0xb6, 0xbb, 0x95, 0x93, 0x56, 0xea, 0xb4, 0xe8, 0x58, 0x54, 0x42, 0x40, 0x09,
	0x0b, 0x66, 0x35, 0xde, 0x08, 0x5e, 0xed, 0xd2, 0xaa, 0x45, 0x84, 0x92, 0xee, 0x95, 0x78, 0x33,
	0x3b, 0x39, 0x6c, 0x83, 0xc9, 0x4c, 0x9c, 0x99, 0x55, 0xf4, 0xd6, 0x17, 0xf3, 0x4d, 0x7c, 0x15,
	0xe9, 0x4c, 0x76, 0x1b, 0x77, 0x13, 0x6f, 0x02, 0xf3, 0xfd, 0xe7, 0x3f, 0x9c, 0xf9, 0x73, 0x06,
	0x8e, 0x2b, 0x99, 0x63, 0x39, 0xa9, 0xb0, 0xaa, 0xa5, 0x2c, 0x93, 0x5a, 0x49, 0x23, 0xc9, 0xd0,
	0xc2, 0xd3, 0xfb, 0x4e, 0xab, 0xd9, 0xb2, 0x10, 0xcc, 0x14, 0x52, 0x38, 0xf9, 0x94, 0x38, 0x8e,
	0x9a, 0x2b, 0xf9, 0xdd, 0xb1, 0xe8, 0x8f, 0x0f, 0xe4, 0xa3, 0x6b, 0x32, 0x57, 0x4c, 0x68, 0xc6,
	0x6f, 0x0c, 0x84, 0x80, 0x7f, 0x71, 0x46, 0xbd, 0xd0, 0x8b, 0x07, 0x33, 0xff, 0x85, 0x97, 0xf9,
	0x17, 0x67, 0x24, 0x84, 0x60, 0x56, 0x4a, 0xfe, 0xe5, 0x3d, 0x16, 0xcb, 0x6b, 0x43, 0xfd, 0xd0,
	0x8b, 0x0f, 0xb3, 0x36, 0x22, 0x11, 0xc0, 0x5b, 0xc4, 0x4b, 0x54, 0xb3, 0x1f, 0x06, 0xe9, 0x60,
	0xe3, 0x6e, 0x51, 0x92, 0xc0, 0xd1, 0x54, 0xa9, 0xe2, 0x1b, 0x2b, 0xe7, 0x45, 0x85, 0xda, 0xb0,
	0xaa, 0xa6, 0x7b, 0x9b, 0xca, 0x1d, 0x8d, 0x8c, 0xe1, 0xa8, 0x35, 0xd8, 0x4d, 0x0b, 0x4d, 0x87,
	0xa1, 0x17, 0x1f, 0x64, 0x3b, 0x9c, 0xa4, 0x70, 0x72, 0x85, 0x22, 0x47, 0x35, 0xe5, 0x5c, 0xae,
	0x84, 0x99, 0xe6, 0xb9, 0x42, 0xad, 0xe9, 0xc8, 0xd6, 0x77, 0x6a, 0xe4, 0x35, 0x3c, 0xc8, 0x90,
	0x17, 0x75, 0x81, 0xc2, 0x6c, 0xd9, 0xf6, 0xad, 0xad, 0x4f, 0x26, 0x4f, 0x61, 0x74, 0x6e, 0xa3,
	0xa4, 0x77, 0x42, 0x2f, 0x0e, 0xd2, 0xc3, 0xc4, 0xe6, 0x9b, 0x38, 0x98, 0x35, 0x62, 0x94, 0xc2,
	0xa3, 0x77, 0x68, 0x76, 0x33, 0xce, 0xf0, 0xeb, 0x0a, 0xb5, 0xe9, 0x8a, 0x3a, 0xfa, 0x0c, 0x8f,
	0x7b, 0x3c, 0xba, 0x96, 0x42, 0x23, 0x79, 0x03, 0x81, 0xb9, 0xc5, 0xd6, 0x1d, 0xa4, 0x0f, 0x9b,
	0x01, 0x3a, 0x7c, 0xed, 0xea, 0xe8, 0xb7, 0xd7, 0xd3, 0x5e, 0xaf, 0x67, 0x1a, 0xc3, 0x5d, 0xb3,
	0xfe, 0x03, 0x57, 0x86, 0x29, 0xd3, 0x9a, 0x6f, 0x4b, 0x21, 0xcf, 0xe0, 0x60, 0x43, 0xce, 0x45,
	0x6e, 0xf7, 0xc2, 0x55, 0xfe, 0xc3, 0x09, 0x85, 0x7d, 0xd6, 0x04, 0x3b, 0xb0, 0xc1, 0xae, 0x8f,
	0xe4, 0x25, 0xc0, 0xe5, 0x66, 0x57, 0xed, 0x32, 0x04, 0xe9, 0xbd, 0xe6, 0x2e, 0xb7, 0x42, 0xd6,
	0x2a, 0x8a, 0x7e, 0x79, 0xf0, 0xa4, 0xef, 0x0a, 0x4d, 0x44, 0x27, 0x30, 0x9c, 0x4b, 0xc3, 0x4a,
	0x3b, 0xfa, 0x5e, 0xe6, 0x0e, 0xe4, 0x03, 0x1c, 0x77, 0x98, 0xa8, 0x1f, 0x0e, 0xfe, 0x1f, 0x60,
	0x97, 0x6b, 0x36, 0xfe, 0x14, 0x2f, 0x0b, 0x73, 0xbd, 0x5a, 0x24, 0x5c, 0x56, 0x93, 0x9f, 0x52,
	0x2e, 0xb8, 0xfb, 0x3e, 0xe7, 0x52, 0xe1, 0x84, 0xcb, 0xaa, 0x92, 0x62, 0x62, 0x7b, 0x2e, 0x46,
	0xf6, 0xbd, 0xbd, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x18, 0x48, 0xf4, 0xb9, 0x03, 0x00,
	0x00,
}
