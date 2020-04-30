// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package darc

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

type Expression struct {
	Matches              []string `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Expression) Reset()         { *m = Expression{} }
func (m *Expression) String() string { return proto.CompactTextString(m) }
func (*Expression) ProtoMessage()    {}
func (*Expression) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Expression) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Expression.Unmarshal(m, b)
}
func (m *Expression) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Expression.Marshal(b, m, deterministic)
}
func (m *Expression) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Expression.Merge(m, src)
}
func (m *Expression) XXX_Size() int {
	return xxx_messageInfo_Expression.Size(m)
}
func (m *Expression) XXX_DiscardUnknown() {
	xxx_messageInfo_Expression.DiscardUnknown(m)
}

var xxx_messageInfo_Expression proto.InternalMessageInfo

func (m *Expression) GetMatches() []string {
	if m != nil {
		return m.Matches
	}
	return nil
}

type AccessControlProto struct {
	Rules                map[string]*Expression `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *AccessControlProto) Reset()         { *m = AccessControlProto{} }
func (m *AccessControlProto) String() string { return proto.CompactTextString(m) }
func (*AccessControlProto) ProtoMessage()    {}
func (*AccessControlProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *AccessControlProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccessControlProto.Unmarshal(m, b)
}
func (m *AccessControlProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccessControlProto.Marshal(b, m, deterministic)
}
func (m *AccessControlProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccessControlProto.Merge(m, src)
}
func (m *AccessControlProto) XXX_Size() int {
	return xxx_messageInfo_AccessControlProto.Size(m)
}
func (m *AccessControlProto) XXX_DiscardUnknown() {
	xxx_messageInfo_AccessControlProto.DiscardUnknown(m)
}

var xxx_messageInfo_AccessControlProto proto.InternalMessageInfo

func (m *AccessControlProto) GetRules() map[string]*Expression {
	if m != nil {
		return m.Rules
	}
	return nil
}

type ActionProto struct {
	Key                  []byte              `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Access               *AccessControlProto `protobuf:"bytes,2,opt,name=access,proto3" json:"access,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ActionProto) Reset()         { *m = ActionProto{} }
func (m *ActionProto) String() string { return proto.CompactTextString(m) }
func (*ActionProto) ProtoMessage()    {}
func (*ActionProto) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *ActionProto) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ActionProto.Unmarshal(m, b)
}
func (m *ActionProto) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ActionProto.Marshal(b, m, deterministic)
}
func (m *ActionProto) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ActionProto.Merge(m, src)
}
func (m *ActionProto) XXX_Size() int {
	return xxx_messageInfo_ActionProto.Size(m)
}
func (m *ActionProto) XXX_DiscardUnknown() {
	xxx_messageInfo_ActionProto.DiscardUnknown(m)
}

var xxx_messageInfo_ActionProto proto.InternalMessageInfo

func (m *ActionProto) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *ActionProto) GetAccess() *AccessControlProto {
	if m != nil {
		return m.Access
	}
	return nil
}

func init() {
	proto.RegisterType((*Expression)(nil), "darc.Expression")
	proto.RegisterType((*AccessControlProto)(nil), "darc.AccessControlProto")
	proto.RegisterMapType((map[string]*Expression)(nil), "darc.AccessControlProto.RulesEntry")
	proto.RegisterType((*ActionProto)(nil), "darc.ActionProto")
}

func init() {
	proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5)
}

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x4f, 0x4b, 0x87, 0x30,
	0x18, 0x80, 0xd9, 0xef, 0x97, 0x86, 0xaf, 0x11, 0xb2, 0xd3, 0xe8, 0x24, 0x06, 0xe2, 0x69, 0x84,
	0x5d, 0xaa, 0x9b, 0x84, 0x97, 0x4e, 0xb5, 0x6f, 0xb0, 0xd6, 0x28, 0x49, 0x37, 0xd9, 0x3b, 0x23,
	0x3f, 0x4b, 0x5f, 0x36, 0x74, 0x86, 0x41, 0x74, 0xdb, 0x9f, 0xe7, 0x7d, 0x78, 0x78, 0xe1, 0x7c,
	0xd0, 0x88, 0xf2, 0x55, 0x23, 0x1f, 0x9d, 0xf5, 0x96, 0x9e, 0xbc, 0x48, 0xa7, 0x8a, 0x12, 0xa0,
	0xfd, 0x1c, 0x9d, 0x46, 0xec, 0xac, 0xa1, 0x0c, 0x4e, 0x07, 0xe9, 0xd5, 0x9b, 0x46, 0x46, 0xf2,
	0x63, 0x95, 0x88, 0x9f, 0x6b, 0xf1, 0x45, 0x80, 0x36, 0x4a, 0x69, 0xc4, 0x7b, 0x6b, 0xbc, 0xb3,
	0xfd, 0xe3, 0x2a, 0xb9, 0x85, 0xc8, 0x4d, 0xfd, 0x86, 0xa7, 0xf5, 0x25, 0x5f, 0xa4, 0xfc, 0x2f,
	0xc8, 0xc5, 0x42, 0xb5, 0xc6, 0xbb, 0x59, 0x84, 0x89, 0x8b, 0x07, 0x80, 0xfd, 0x91, 0x66, 0x70,
	0x7c, 0xd7, 0x33, 0x23, 0x39, 0xa9, 0x12, 0xb1, 0x1c, 0x69, 0x09, 0xd1, 0x87, 0xec, 0x27, 0xcd,
	0x0e, 0x39, 0xa9, 0xd2, 0x3a, 0x0b, 0xea, 0x3d, 0x56, 0x84, 0xef, 0xbb, 0xc3, 0x0d, 0x29, 0x9e,
	0x20, 0x6d, 0x94, 0xef, 0xac, 0x09, 0x55, 0xbf, 0x64, 0x67, 0x41, 0x76, 0x05, 0xb1, 0x5c, 0xa3,
	0x36, 0x1b, 0xfb, 0x2f, 0x54, 0x6c, 0xdc, 0x73, 0xbc, 0x6e, 0xe9, 0xfa, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0x76, 0x85, 0xae, 0x37, 0x01, 0x00, 0x00,
}
