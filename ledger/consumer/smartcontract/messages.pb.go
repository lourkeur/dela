// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package smartcontract

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

// TransactionProto is the message that represents a transaction.
type TransactionProto struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionProto) Reset()         { *m = TransactionProto{} }
func (m *TransactionProto) String() string { return proto.CompactTextString(m) }
func (*TransactionProto) ProtoMessage()    {}
func (*TransactionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *TransactionProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionProto.Unmarshal(m, b)
}
func (m *TransactionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionProto.Marshal(b, m, deterministic)
}
func (m *TransactionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionProto.Merge(m, src)
}
func (m *TransactionProto) XXX_Size() int {
	return xxx_messageInfo_TransactionProto.Size(m)
}
func (m *TransactionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionProto.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionProto proto.InternalMessageInfo

func (m *TransactionProto) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*TransactionProto)(nil), "smartcontract.TransactionProto")
}

func init() {
	proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5)
}

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 95 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0x4d, 0x2d, 0x2e,
	0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2d, 0xce, 0x4d, 0x2c,
	0x2a, 0x49, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0x51, 0xd2, 0xe0, 0x12, 0x08, 0x29, 0x4a,
	0xcc, 0x2b, 0x4e, 0x4c, 0x2e, 0xc9, 0xcc, 0xcf, 0x0b, 0x00, 0x2b, 0x11, 0xe1, 0x62, 0x2d, 0x4b,
	0xcc, 0x29, 0x4d, 0x95, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x92, 0xd8, 0xc0, 0xfa,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0xeb, 0x40, 0x1f, 0x51, 0x00, 0x00, 0x00,
}
