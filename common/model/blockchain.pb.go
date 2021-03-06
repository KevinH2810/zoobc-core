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
// source: model/blockchain.proto

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

type ChainStatus struct {
	// Integer indicating chaintype
	ChainType            int32    `protobuf:"varint,1,opt,name=ChainType,proto3" json:"ChainType,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	LastBlock            *Block   `protobuf:"bytes,3,opt,name=LastBlock,proto3" json:"LastBlock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainStatus) Reset()         { *m = ChainStatus{} }
func (m *ChainStatus) String() string { return proto.CompactTextString(m) }
func (*ChainStatus) ProtoMessage()    {}
func (*ChainStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2ffcbd1121992b1, []int{0}
}

func (m *ChainStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainStatus.Unmarshal(m, b)
}
func (m *ChainStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainStatus.Marshal(b, m, deterministic)
}
func (m *ChainStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainStatus.Merge(m, src)
}
func (m *ChainStatus) XXX_Size() int {
	return xxx_messageInfo_ChainStatus.Size(m)
}
func (m *ChainStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ChainStatus proto.InternalMessageInfo

func (m *ChainStatus) GetChainType() int32 {
	if m != nil {
		return m.ChainType
	}
	return 0
}

func (m *ChainStatus) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ChainStatus) GetLastBlock() *Block {
	if m != nil {
		return m.LastBlock
	}
	return nil
}

type GetCumulativeDifficultyResponse struct {
	CumulativeDifficulty string   `protobuf:"bytes,1,opt,name=CumulativeDifficulty,proto3" json:"CumulativeDifficulty,omitempty"`
	Height               uint32   `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCumulativeDifficultyResponse) Reset()         { *m = GetCumulativeDifficultyResponse{} }
func (m *GetCumulativeDifficultyResponse) String() string { return proto.CompactTextString(m) }
func (*GetCumulativeDifficultyResponse) ProtoMessage()    {}
func (*GetCumulativeDifficultyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2ffcbd1121992b1, []int{1}
}

func (m *GetCumulativeDifficultyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCumulativeDifficultyResponse.Unmarshal(m, b)
}
func (m *GetCumulativeDifficultyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCumulativeDifficultyResponse.Marshal(b, m, deterministic)
}
func (m *GetCumulativeDifficultyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCumulativeDifficultyResponse.Merge(m, src)
}
func (m *GetCumulativeDifficultyResponse) XXX_Size() int {
	return xxx_messageInfo_GetCumulativeDifficultyResponse.Size(m)
}
func (m *GetCumulativeDifficultyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCumulativeDifficultyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCumulativeDifficultyResponse proto.InternalMessageInfo

func (m *GetCumulativeDifficultyResponse) GetCumulativeDifficulty() string {
	if m != nil {
		return m.CumulativeDifficulty
	}
	return ""
}

func (m *GetCumulativeDifficultyResponse) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

type GetCumulativeDifficultyRequest struct {
	// Integer indicating chaintype
	ChainType            int32    `protobuf:"varint,1,opt,name=ChainType,proto3" json:"ChainType,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCumulativeDifficultyRequest) Reset()         { *m = GetCumulativeDifficultyRequest{} }
func (m *GetCumulativeDifficultyRequest) String() string { return proto.CompactTextString(m) }
func (*GetCumulativeDifficultyRequest) ProtoMessage()    {}
func (*GetCumulativeDifficultyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2ffcbd1121992b1, []int{2}
}

func (m *GetCumulativeDifficultyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCumulativeDifficultyRequest.Unmarshal(m, b)
}
func (m *GetCumulativeDifficultyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCumulativeDifficultyRequest.Marshal(b, m, deterministic)
}
func (m *GetCumulativeDifficultyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCumulativeDifficultyRequest.Merge(m, src)
}
func (m *GetCumulativeDifficultyRequest) XXX_Size() int {
	return xxx_messageInfo_GetCumulativeDifficultyRequest.Size(m)
}
func (m *GetCumulativeDifficultyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCumulativeDifficultyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCumulativeDifficultyRequest proto.InternalMessageInfo

func (m *GetCumulativeDifficultyRequest) GetChainType() int32 {
	if m != nil {
		return m.ChainType
	}
	return 0
}

type GetCommonMilestoneBlockIdsRequest struct {
	// Integer indicating chaintype
	ChainType            int32    `protobuf:"varint,1,opt,name=ChainType,proto3" json:"ChainType,omitempty"`
	LastBlockID          int64    `protobuf:"varint,2,opt,name=LastBlockID,proto3" json:"LastBlockID,omitempty"`
	LastMilestoneBlockID int64    `protobuf:"varint,3,opt,name=LastMilestoneBlockID,proto3" json:"LastMilestoneBlockID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCommonMilestoneBlockIdsRequest) Reset()         { *m = GetCommonMilestoneBlockIdsRequest{} }
func (m *GetCommonMilestoneBlockIdsRequest) String() string { return proto.CompactTextString(m) }
func (*GetCommonMilestoneBlockIdsRequest) ProtoMessage()    {}
func (*GetCommonMilestoneBlockIdsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2ffcbd1121992b1, []int{3}
}

func (m *GetCommonMilestoneBlockIdsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommonMilestoneBlockIdsRequest.Unmarshal(m, b)
}
func (m *GetCommonMilestoneBlockIdsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommonMilestoneBlockIdsRequest.Marshal(b, m, deterministic)
}
func (m *GetCommonMilestoneBlockIdsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommonMilestoneBlockIdsRequest.Merge(m, src)
}
func (m *GetCommonMilestoneBlockIdsRequest) XXX_Size() int {
	return xxx_messageInfo_GetCommonMilestoneBlockIdsRequest.Size(m)
}
func (m *GetCommonMilestoneBlockIdsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommonMilestoneBlockIdsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommonMilestoneBlockIdsRequest proto.InternalMessageInfo

func (m *GetCommonMilestoneBlockIdsRequest) GetChainType() int32 {
	if m != nil {
		return m.ChainType
	}
	return 0
}

func (m *GetCommonMilestoneBlockIdsRequest) GetLastBlockID() int64 {
	if m != nil {
		return m.LastBlockID
	}
	return 0
}

func (m *GetCommonMilestoneBlockIdsRequest) GetLastMilestoneBlockID() int64 {
	if m != nil {
		return m.LastMilestoneBlockID
	}
	return 0
}

type GetCommonMilestoneBlockIdsResponse struct {
	BlockIds             []int64  `protobuf:"varint,1,rep,packed,name=BlockIds,proto3" json:"BlockIds,omitempty"`
	Last                 bool     `protobuf:"varint,2,opt,name=Last,proto3" json:"Last,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCommonMilestoneBlockIdsResponse) Reset()         { *m = GetCommonMilestoneBlockIdsResponse{} }
func (m *GetCommonMilestoneBlockIdsResponse) String() string { return proto.CompactTextString(m) }
func (*GetCommonMilestoneBlockIdsResponse) ProtoMessage()    {}
func (*GetCommonMilestoneBlockIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2ffcbd1121992b1, []int{4}
}

func (m *GetCommonMilestoneBlockIdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCommonMilestoneBlockIdsResponse.Unmarshal(m, b)
}
func (m *GetCommonMilestoneBlockIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCommonMilestoneBlockIdsResponse.Marshal(b, m, deterministic)
}
func (m *GetCommonMilestoneBlockIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCommonMilestoneBlockIdsResponse.Merge(m, src)
}
func (m *GetCommonMilestoneBlockIdsResponse) XXX_Size() int {
	return xxx_messageInfo_GetCommonMilestoneBlockIdsResponse.Size(m)
}
func (m *GetCommonMilestoneBlockIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCommonMilestoneBlockIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCommonMilestoneBlockIdsResponse proto.InternalMessageInfo

func (m *GetCommonMilestoneBlockIdsResponse) GetBlockIds() []int64 {
	if m != nil {
		return m.BlockIds
	}
	return nil
}

func (m *GetCommonMilestoneBlockIdsResponse) GetLast() bool {
	if m != nil {
		return m.Last
	}
	return false
}

func init() {
	proto.RegisterType((*ChainStatus)(nil), "model.ChainStatus")
	proto.RegisterType((*GetCumulativeDifficultyResponse)(nil), "model.GetCumulativeDifficultyResponse")
	proto.RegisterType((*GetCumulativeDifficultyRequest)(nil), "model.GetCumulativeDifficultyRequest")
	proto.RegisterType((*GetCommonMilestoneBlockIdsRequest)(nil), "model.GetCommonMilestoneBlockIdsRequest")
	proto.RegisterType((*GetCommonMilestoneBlockIdsResponse)(nil), "model.GetCommonMilestoneBlockIdsResponse")
}

func init() {
	proto.RegisterFile("model/blockchain.proto", fileDescriptor_c2ffcbd1121992b1)
}

var fileDescriptor_c2ffcbd1121992b1 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0x51, 0x4b, 0x32, 0x41,
	0x14, 0x65, 0xf4, 0x53, 0xf4, 0xfa, 0xf5, 0xd0, 0x20, 0xb2, 0x44, 0xd8, 0xb6, 0xf4, 0xb0, 0x08,
	0xed, 0x86, 0x41, 0x8f, 0x3d, 0xa8, 0x50, 0x42, 0xbd, 0x4c, 0x3d, 0x44, 0x6f, 0xbb, 0xe3, 0xa8,
	0x43, 0xbb, 0x7b, 0xcd, 0x99, 0x0d, 0xec, 0xcf, 0xf4, 0x57, 0x63, 0xaf, 0xa6, 0x12, 0x2a, 0xbd,
	0x0c, 0x73, 0xcf, 0xb9, 0x9c, 0x73, 0xee, 0x9d, 0x81, 0x56, 0x8a, 0x23, 0x95, 0x84, 0x71, 0x82,
	0xf2, 0x4d, 0x4e, 0x23, 0x9d, 0x05, 0xb3, 0x39, 0x5a, 0xe4, 0x15, 0xc2, 0x4f, 0x8e, 0xb7, 0xe8,
	0x25, 0xe3, 0x21, 0x34, 0xfa, 0x45, 0xe3, 0x93, 0x8d, 0x6c, 0x6e, 0xf8, 0x29, 0xd4, 0xa9, 0x7c,
	0x5e, 0xcc, 0x94, 0xc3, 0x5c, 0xe6, 0x57, 0xc4, 0x06, 0xe0, 0x2d, 0xa8, 0xde, 0x2b, 0x3d, 0x99,
	0x5a, 0xa7, 0xe4, 0x32, 0xff, 0x48, 0xac, 0x2a, 0xde, 0x81, 0xfa, 0x43, 0x64, 0x6c, 0xaf, 0xd0,
	0x75, 0xca, 0x2e, 0xf3, 0x1b, 0xdd, 0xff, 0x01, 0x79, 0x05, 0x84, 0x89, 0x0d, 0xed, 0xa5, 0x70,
	0x76, 0xa7, 0x6c, 0x3f, 0x4f, 0xf3, 0x24, 0xb2, 0xfa, 0x43, 0x0d, 0xf4, 0x78, 0xac, 0x65, 0x9e,
	0xd8, 0x85, 0x50, 0x66, 0x86, 0x99, 0x51, 0xbc, 0x0b, 0xcd, 0x5d, 0x3c, 0xe5, 0xa9, 0x8b, 0x9d,
	0xdc, 0xbe, 0x68, 0xde, 0x2d, 0xb4, 0xf7, 0xda, 0xbd, 0xe7, 0xca, 0xd8, 0xc3, 0x23, 0x7b, 0x5f,
	0x0c, 0xce, 0x0b, 0x01, 0x4c, 0x53, 0xcc, 0x1e, 0x75, 0xa2, 0x8c, 0xc5, 0x4c, 0xd1, 0x28, 0xc3,
	0x91, 0xf9, 0x93, 0x06, 0xbf, 0x80, 0xc6, 0x7a, 0xfe, 0xe1, 0x80, 0x02, 0x96, 0x7b, 0xa5, 0x2b,
	0x26, 0xb6, 0x61, 0x7e, 0x03, 0xcd, 0xa2, 0xfc, 0xe5, 0x31, 0xa0, 0x7d, 0x2e, 0xdb, 0x77, 0xf2,
	0xde, 0x0b, 0x78, 0x87, 0x02, 0xae, 0x76, 0xda, 0x86, 0xda, 0x0f, 0xe6, 0x30, 0xb7, 0xbc, 0x52,
	0x5c, 0x63, 0x9c, 0xc3, 0xbf, 0x42, 0x9d, 0xc2, 0xd5, 0x04, 0xdd, 0x7b, 0x9d, 0x57, 0x7f, 0xa2,
	0xed, 0x34, 0x8f, 0x03, 0x89, 0x69, 0xf8, 0x89, 0x18, 0xcb, 0xe5, 0x79, 0x29, 0x71, 0xae, 0x42,
	0x49, 0x96, 0x21, 0xbd, 0x73, 0x5c, 0xa5, 0xef, 0x74, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x34,
	0x29, 0x28, 0xfc, 0x82, 0x02, 0x00, 0x00,
}
