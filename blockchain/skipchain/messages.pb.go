// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package skipchain

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

// ConodeProto is the message that contains the address and the public key of a
// conode.
type ConodeProto struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	PublicKey            *any.Any `protobuf:"bytes,2,opt,name=publicKey,proto3" json:"publicKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConodeProto) Reset()         { *m = ConodeProto{} }
func (m *ConodeProto) String() string { return proto.CompactTextString(m) }
func (*ConodeProto) ProtoMessage()    {}
func (*ConodeProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *ConodeProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConodeProto.Unmarshal(m, b)
}
func (m *ConodeProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConodeProto.Marshal(b, m, deterministic)
}
func (m *ConodeProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConodeProto.Merge(m, src)
}
func (m *ConodeProto) XXX_Size() int {
	return xxx_messageInfo_ConodeProto.Size(m)
}
func (m *ConodeProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ConodeProto.DiscardUnknown(m)
}

var xxx_messageInfo_ConodeProto proto.InternalMessageInfo

func (m *ConodeProto) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ConodeProto) GetPublicKey() *any.Any {
	if m != nil {
		return m.PublicKey
	}
	return nil
}

type Roster struct {
	Conodes              []*ConodeProto `protobuf:"bytes,1,rep,name=conodes,proto3" json:"conodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Roster) Reset()         { *m = Roster{} }
func (m *Roster) String() string { return proto.CompactTextString(m) }
func (*Roster) ProtoMessage()    {}
func (*Roster) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *Roster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Roster.Unmarshal(m, b)
}
func (m *Roster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Roster.Marshal(b, m, deterministic)
}
func (m *Roster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Roster.Merge(m, src)
}
func (m *Roster) XXX_Size() int {
	return xxx_messageInfo_Roster.Size(m)
}
func (m *Roster) XXX_DiscardUnknown() {
	xxx_messageInfo_Roster.DiscardUnknown(m)
}

var xxx_messageInfo_Roster proto.InternalMessageInfo

func (m *Roster) GetConodes() []*ConodeProto {
	if m != nil {
		return m.Conodes
	}
	return nil
}

// BlockProto is the message that contains the minimal data to instantiate a
// block. It is not sufficient to prove the block integrity as the chain from
// the genesis block is missing.
type BlockProto struct {
	Index                uint64   `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	GenesisID            []byte   `protobuf:"bytes,5,opt,name=genesisID,proto3" json:"genesisID,omitempty"`
	Backlink             []byte   `protobuf:"bytes,6,opt,name=backlink,proto3" json:"backlink,omitempty"`
	Roster               *Roster  `protobuf:"bytes,7,opt,name=roster,proto3" json:"roster,omitempty"`
	Payload              *any.Any `protobuf:"bytes,8,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockProto) Reset()         { *m = BlockProto{} }
func (m *BlockProto) String() string { return proto.CompactTextString(m) }
func (*BlockProto) ProtoMessage()    {}
func (*BlockProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *BlockProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockProto.Unmarshal(m, b)
}
func (m *BlockProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockProto.Marshal(b, m, deterministic)
}
func (m *BlockProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockProto.Merge(m, src)
}
func (m *BlockProto) XXX_Size() int {
	return xxx_messageInfo_BlockProto.Size(m)
}
func (m *BlockProto) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockProto.DiscardUnknown(m)
}

var xxx_messageInfo_BlockProto proto.InternalMessageInfo

func (m *BlockProto) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *BlockProto) GetGenesisID() []byte {
	if m != nil {
		return m.GenesisID
	}
	return nil
}

func (m *BlockProto) GetBacklink() []byte {
	if m != nil {
		return m.Backlink
	}
	return nil
}

func (m *BlockProto) GetRoster() *Roster {
	if m != nil {
		return m.Roster
	}
	return nil
}

func (m *BlockProto) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

// VerifiableBlockProto is the message that contains a verifiable block. It
// contains everything necessary to prove a block is valid as long as the
// receiver has the genesis block.
type VerifiableBlockProto struct {
	Block                *BlockProto `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Chain                *any.Any    `protobuf:"bytes,2,opt,name=chain,proto3" json:"chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *VerifiableBlockProto) Reset()         { *m = VerifiableBlockProto{} }
func (m *VerifiableBlockProto) String() string { return proto.CompactTextString(m) }
func (*VerifiableBlockProto) ProtoMessage()    {}
func (*VerifiableBlockProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *VerifiableBlockProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifiableBlockProto.Unmarshal(m, b)
}
func (m *VerifiableBlockProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifiableBlockProto.Marshal(b, m, deterministic)
}
func (m *VerifiableBlockProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifiableBlockProto.Merge(m, src)
}
func (m *VerifiableBlockProto) XXX_Size() int {
	return xxx_messageInfo_VerifiableBlockProto.Size(m)
}
func (m *VerifiableBlockProto) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifiableBlockProto.DiscardUnknown(m)
}

var xxx_messageInfo_VerifiableBlockProto proto.InternalMessageInfo

func (m *VerifiableBlockProto) GetBlock() *BlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *VerifiableBlockProto) GetChain() *any.Any {
	if m != nil {
		return m.Chain
	}
	return nil
}

// PropagateGenesis is the message containing a genesis to be transmitted to the
// participants.
type PropagateGenesis struct {
	Genesis              *BlockProto `protobuf:"bytes,1,opt,name=genesis,proto3" json:"genesis,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PropagateGenesis) Reset()         { *m = PropagateGenesis{} }
func (m *PropagateGenesis) String() string { return proto.CompactTextString(m) }
func (*PropagateGenesis) ProtoMessage()    {}
func (*PropagateGenesis) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *PropagateGenesis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PropagateGenesis.Unmarshal(m, b)
}
func (m *PropagateGenesis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PropagateGenesis.Marshal(b, m, deterministic)
}
func (m *PropagateGenesis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PropagateGenesis.Merge(m, src)
}
func (m *PropagateGenesis) XXX_Size() int {
	return xxx_messageInfo_PropagateGenesis.Size(m)
}
func (m *PropagateGenesis) XXX_DiscardUnknown() {
	xxx_messageInfo_PropagateGenesis.DiscardUnknown(m)
}

var xxx_messageInfo_PropagateGenesis proto.InternalMessageInfo

func (m *PropagateGenesis) GetGenesis() *BlockProto {
	if m != nil {
		return m.Genesis
	}
	return nil
}

func init() {
	proto.RegisterType((*ConodeProto)(nil), "skipchain.ConodeProto")
	proto.RegisterType((*Roster)(nil), "skipchain.Roster")
	proto.RegisterType((*BlockProto)(nil), "skipchain.BlockProto")
	proto.RegisterType((*VerifiableBlockProto)(nil), "skipchain.VerifiableBlockProto")
	proto.RegisterType((*PropagateGenesis)(nil), "skipchain.PropagateGenesis")
}

func init() {
	proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5)
}

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x15, 0x20, 0x49, 0x7b, 0x41, 0x08, 0xac, 0x82, 0x4c, 0xc5, 0x50, 0x65, 0x2a, 0x20,
	0xb9, 0x28, 0x6c, 0x6c, 0x50, 0x24, 0x84, 0x58, 0x2a, 0x0f, 0x2c, 0x4c, 0x4e, 0x72, 0x0d, 0x56,
	0x82, 0x1d, 0xc5, 0xad, 0x44, 0x9e, 0x8c, 0xd7, 0x43, 0xb5, 0x9b, 0x36, 0x53, 0xc7, 0xcb, 0x7d,
	0xf9, 0xef, 0xfb, 0x0d, 0x67, 0x3f, 0x68, 0x8c, 0x28, 0xd0, 0xb0, 0xba, 0xd1, 0x2b, 0x4d, 0x86,
	0xa6, 0x94, 0x75, 0xf6, 0x2d, 0xa4, 0x1a, 0x5f, 0x17, 0x5a, 0x17, 0x15, 0xce, 0xec, 0x22, 0x5d,
	0x2f, 0x67, 0x42, 0xb5, 0x8e, 0x8a, 0xbf, 0x20, 0x9a, 0x6b, 0xa5, 0x73, 0x5c, 0xd8, 0x9f, 0x28,
	0x84, 0x22, 0xcf, 0x1b, 0x34, 0x86, 0x7a, 0x13, 0x6f, 0x7a, 0xca, 0xbb, 0x91, 0x24, 0x30, 0xac,
	0xd7, 0x69, 0x25, 0xb3, 0x0f, 0x6c, 0xe9, 0xd1, 0xc4, 0x9b, 0x46, 0xc9, 0x88, 0xb9, 0x5c, 0xd6,
	0xe5, 0xb2, 0x67, 0xd5, 0xf2, 0x3d, 0x16, 0x3f, 0x41, 0xc0, 0xb5, 0x59, 0x61, 0x43, 0x1e, 0x20,
	0xcc, 0xec, 0x99, 0x4d, 0xee, 0xf1, 0x34, 0x4a, 0xae, 0xd8, 0x4e, 0x8f, 0xf5, 0x04, 0x78, 0x87,
	0xc5, 0x7f, 0x1e, 0xc0, 0x4b, 0xa5, 0xb3, 0xd2, 0x89, 0x8d, 0xc0, 0x97, 0x2a, 0xc7, 0x5f, 0xab,
	0x75, 0xc2, 0xdd, 0x40, 0x6e, 0x60, 0x58, 0xa0, 0x42, 0x23, 0xcd, 0xfb, 0x2b, 0xf5, 0xad, 0xf0,
	0xfe, 0x03, 0x19, 0xc3, 0x20, 0x15, 0x59, 0x59, 0x49, 0x55, 0xd2, 0xc0, 0x2e, 0x77, 0x33, 0xb9,
	0x85, 0xa0, 0xb1, 0x6a, 0x34, 0xb4, 0x5d, 0x2e, 0x7a, 0x3e, 0xce, 0x99, 0x6f, 0x01, 0xc2, 0x20,
	0xac, 0x45, 0x5b, 0x69, 0x91, 0xd3, 0xc1, 0x81, 0xde, 0x1d, 0x14, 0x6b, 0x18, 0x7d, 0x62, 0x23,
	0x97, 0x52, 0xa4, 0x15, 0xf6, 0x2a, 0xdc, 0x83, 0x9f, 0x6e, 0x26, 0x5b, 0x21, 0x4a, 0x2e, 0x7b,
	0x17, 0xf7, 0x14, 0x77, 0x0c, 0xb9, 0x03, 0xdf, 0xae, 0x0e, 0x3e, 0xb5, 0x43, 0xe2, 0x39, 0x9c,
	0x2f, 0x1a, 0x5d, 0x8b, 0x42, 0xac, 0xf0, 0xcd, 0xb5, 0x27, 0x33, 0x08, 0xb7, 0x0f, 0x71, 0xf8,
	0x5c, 0x47, 0xa5, 0x81, 0x4d, 0x7e, 0xfc, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x17, 0xec, 0x54, 0xa7,
	0x47, 0x02, 0x00, 0x00,
}
