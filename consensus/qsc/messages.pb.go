// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package qsc

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

type Message struct {
	Node                 int64    `protobuf:"varint,1,opt,name=node,proto3" json:"node,omitempty"`
	Value                *any.Any `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetNode() int64 {
	if m != nil {
		return m.Node
	}
	return 0
}

func (m *Message) GetValue() *any.Any {
	if m != nil {
		return m.Value
	}
	return nil
}

type MessageSet struct {
	Messages             map[int64]*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	TimeStep             uint64             `protobuf:"varint,2,opt,name=timeStep,proto3" json:"timeStep,omitempty"`
	Node                 int64              `protobuf:"varint,3,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MessageSet) Reset()         { *m = MessageSet{} }
func (m *MessageSet) String() string { return proto.CompactTextString(m) }
func (*MessageSet) ProtoMessage()    {}
func (*MessageSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *MessageSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageSet.Unmarshal(m, b)
}
func (m *MessageSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageSet.Marshal(b, m, deterministic)
}
func (m *MessageSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageSet.Merge(m, src)
}
func (m *MessageSet) XXX_Size() int {
	return xxx_messageInfo_MessageSet.Size(m)
}
func (m *MessageSet) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageSet.DiscardUnknown(m)
}

var xxx_messageInfo_MessageSet proto.InternalMessageInfo

func (m *MessageSet) GetMessages() map[int64]*Message {
	if m != nil {
		return m.Messages
	}
	return nil
}

func (m *MessageSet) GetTimeStep() uint64 {
	if m != nil {
		return m.TimeStep
	}
	return 0
}

func (m *MessageSet) GetNode() int64 {
	if m != nil {
		return m.Node
	}
	return 0
}

type View struct {
	Received             map[int64]*Message `protobuf:"bytes,1,rep,name=received,proto3" json:"received,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Broadcasted          map[int64]*Message `protobuf:"bytes,2,rep,name=broadcasted,proto3" json:"broadcasted,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *View) Reset()         { *m = View{} }
func (m *View) String() string { return proto.CompactTextString(m) }
func (*View) ProtoMessage()    {}
func (*View) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *View) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_View.Unmarshal(m, b)
}
func (m *View) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_View.Marshal(b, m, deterministic)
}
func (m *View) XXX_Merge(src proto.Message) {
	xxx_messageInfo_View.Merge(m, src)
}
func (m *View) XXX_Size() int {
	return xxx_messageInfo_View.Size(m)
}
func (m *View) XXX_DiscardUnknown() {
	xxx_messageInfo_View.DiscardUnknown(m)
}

var xxx_messageInfo_View proto.InternalMessageInfo

func (m *View) GetReceived() map[int64]*Message {
	if m != nil {
		return m.Received
	}
	return nil
}

func (m *View) GetBroadcasted() map[int64]*Message {
	if m != nil {
		return m.Broadcasted
	}
	return nil
}

type RequestMessageSet struct {
	TimeStep             uint64   `protobuf:"varint,1,opt,name=timeStep,proto3" json:"timeStep,omitempty"`
	Nodes                []int64  `protobuf:"varint,2,rep,packed,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestMessageSet) Reset()         { *m = RequestMessageSet{} }
func (m *RequestMessageSet) String() string { return proto.CompactTextString(m) }
func (*RequestMessageSet) ProtoMessage()    {}
func (*RequestMessageSet) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *RequestMessageSet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestMessageSet.Unmarshal(m, b)
}
func (m *RequestMessageSet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestMessageSet.Marshal(b, m, deterministic)
}
func (m *RequestMessageSet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestMessageSet.Merge(m, src)
}
func (m *RequestMessageSet) XXX_Size() int {
	return xxx_messageInfo_RequestMessageSet.Size(m)
}
func (m *RequestMessageSet) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestMessageSet.DiscardUnknown(m)
}

var xxx_messageInfo_RequestMessageSet proto.InternalMessageInfo

func (m *RequestMessageSet) GetTimeStep() uint64 {
	if m != nil {
		return m.TimeStep
	}
	return 0
}

func (m *RequestMessageSet) GetNodes() []int64 {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Epoch struct {
	Random               int64    `protobuf:"varint,1,opt,name=random,proto3" json:"random,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Epoch) Reset()         { *m = Epoch{} }
func (m *Epoch) String() string { return proto.CompactTextString(m) }
func (*Epoch) ProtoMessage()    {}
func (*Epoch) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *Epoch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Epoch.Unmarshal(m, b)
}
func (m *Epoch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Epoch.Marshal(b, m, deterministic)
}
func (m *Epoch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Epoch.Merge(m, src)
}
func (m *Epoch) XXX_Size() int {
	return xxx_messageInfo_Epoch.Size(m)
}
func (m *Epoch) XXX_DiscardUnknown() {
	xxx_messageInfo_Epoch.DiscardUnknown(m)
}

var xxx_messageInfo_Epoch proto.InternalMessageInfo

func (m *Epoch) GetRandom() int64 {
	if m != nil {
		return m.Random
	}
	return 0
}

func (m *Epoch) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type History struct {
	Epochs               []*Epoch `protobuf:"bytes,1,rep,name=epochs,proto3" json:"epochs,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *History) Reset()         { *m = History{} }
func (m *History) String() string { return proto.CompactTextString(m) }
func (*History) ProtoMessage()    {}
func (*History) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{5}
}

func (m *History) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_History.Unmarshal(m, b)
}
func (m *History) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_History.Marshal(b, m, deterministic)
}
func (m *History) XXX_Merge(src proto.Message) {
	xxx_messageInfo_History.Merge(m, src)
}
func (m *History) XXX_Size() int {
	return xxx_messageInfo_History.Size(m)
}
func (m *History) XXX_DiscardUnknown() {
	xxx_messageInfo_History.DiscardUnknown(m)
}

var xxx_messageInfo_History proto.InternalMessageInfo

func (m *History) GetEpochs() []*Epoch {
	if m != nil {
		return m.Epochs
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "qsc.Message")
	proto.RegisterType((*MessageSet)(nil), "qsc.MessageSet")
	proto.RegisterMapType((map[int64]*Message)(nil), "qsc.MessageSet.MessagesEntry")
	proto.RegisterType((*View)(nil), "qsc.View")
	proto.RegisterMapType((map[int64]*Message)(nil), "qsc.View.BroadcastedEntry")
	proto.RegisterMapType((map[int64]*Message)(nil), "qsc.View.ReceivedEntry")
	proto.RegisterType((*RequestMessageSet)(nil), "qsc.RequestMessageSet")
	proto.RegisterType((*Epoch)(nil), "qsc.Epoch")
	proto.RegisterType((*History)(nil), "qsc.History")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x49, 0xd2, 0x7f, 0x4c, 0xab, 0xd4, 0xa5, 0x68, 0x0d, 0x08, 0x92, 0x93, 0x08, 0xa6,
	0x60, 0x2f, 0x2a, 0x5e, 0x14, 0x0a, 0x16, 0xf4, 0xb2, 0x05, 0xef, 0xdb, 0x64, 0x6c, 0x83, 0x6d,
	0xb6, 0xcd, 0x6e, 0x2b, 0x79, 0x0d, 0x1f, 0xc9, 0x27, 0x33, 0xbb, 0xd9, 0x34, 0x69, 0x8f, 0x7a,
	0x9b, 0xd9, 0x99, 0xef, 0xdb, 0xdf, 0xec, 0x2c, 0x1c, 0x2f, 0x51, 0x08, 0x36, 0x43, 0xe1, 0xaf,
	0x12, 0x2e, 0x39, 0x71, 0xd6, 0x22, 0x70, 0xcf, 0x67, 0x9c, 0xcf, 0x16, 0x38, 0xd0, 0x47, 0xd3,
	0xcd, 0xc7, 0x80, 0xc5, 0x69, 0x5e, 0xf7, 0xc6, 0xd0, 0x7c, 0xcb, 0x15, 0x84, 0x40, 0x2d, 0xe6,
	0x21, 0xf6, 0xad, 0x4b, 0xeb, 0xca, 0xa1, 0x3a, 0x26, 0xd7, 0x50, 0xdf, 0xb2, 0xc5, 0x06, 0xfb,
	0x76, 0x76, 0xd8, 0xbe, 0xed, 0xf9, 0xb9, 0x93, 0x5f, 0x38, 0xf9, 0x4f, 0x71, 0x4a, 0xf3, 0x16,
	0xef, 0xc7, 0x02, 0x30, 0x5e, 0x13, 0x94, 0xe4, 0x1e, 0x5a, 0x05, 0x4b, 0x66, 0xe9, 0x64, 0xea,
	0x0b, 0x3f, 0x83, 0xf1, 0xcb, 0x96, 0x22, 0x14, 0xa3, 0x58, 0x26, 0x29, 0xdd, 0xb5, 0x13, 0x17,
	0x5a, 0x32, 0x5a, 0xe2, 0x44, 0xe2, 0x4a, 0x5f, 0x5c, 0xa3, 0xbb, 0x7c, 0x47, 0xe9, 0x94, 0x94,
	0xee, 0x18, 0x8e, 0xf6, 0xac, 0x48, 0x17, 0x9c, 0x4f, 0x4c, 0xcd, 0x24, 0x2a, 0x24, 0xde, 0xfe,
	0x20, 0x9d, 0x2a, 0x8a, 0x19, 0xe0, 0xc1, 0xbe, 0xb3, 0xbc, 0x6f, 0x1b, 0x6a, 0xef, 0x11, 0x7e,
	0x91, 0x21, 0xb4, 0x12, 0x0c, 0x30, 0xda, 0x62, 0x68, 0xf0, 0xcf, 0xb4, 0x46, 0x15, 0x7d, 0x6a,
	0x2a, 0x06, 0xbc, 0x68, 0x24, 0x8f, 0xd0, 0x9e, 0x26, 0x9c, 0x85, 0x01, 0x13, 0x32, 0xd3, 0xd9,
	0x5a, 0xe7, 0x96, 0xba, 0xe7, 0xb2, 0x98, 0x4b, 0xab, 0xed, 0x6a, 0x8c, 0x3d, 0xe3, 0xbf, 0x8f,
	0xe1, 0xbe, 0x42, 0xf7, 0xf0, 0xae, 0x7f, 0x3c, 0xca, 0x08, 0x4e, 0x28, 0xae, 0x37, 0x28, 0x64,
	0x65, 0xbf, 0xd5, 0x25, 0x59, 0x07, 0x4b, 0xea, 0x41, 0x5d, 0x2d, 0x46, 0xe8, 0x17, 0x70, 0x68,
	0x9e, 0x78, 0x43, 0xa8, 0x8f, 0x56, 0x3c, 0x98, 0x93, 0x53, 0x68, 0x24, 0x2c, 0x0e, 0xf9, 0xd2,
	0xc0, 0x98, 0x4c, 0xed, 0x76, 0xce, 0xc4, 0x5c, 0xe3, 0x74, 0xa8, 0x8e, 0xbd, 0x1b, 0x68, 0xbe,
	0x44, 0x42, 0xf2, 0x44, 0xe1, 0x36, 0x50, 0xe9, 0x8b, 0xff, 0x04, 0x9a, 0x57, 0x5b, 0x52, 0x53,
	0x99, 0x36, 0xf4, 0xcf, 0x1c, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x6b, 0x1e, 0x0a, 0x5a, 0x08,
	0x03, 0x00, 0x00,
}
