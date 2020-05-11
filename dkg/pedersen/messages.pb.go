// Code generated by protoc-gen-go. DO NOT EDIT.
// source: messages.proto

package pedersen

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/any"
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

type Start struct {
	// threshold
	T                    uint32   `protobuf:"varint,1,opt,name=t,proto3" json:"t,omitempty"`
	Addresses            [][]byte `protobuf:"bytes,2,rep,name=addresses,proto3" json:"addresses,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Start) Reset()         { *m = Start{} }
func (m *Start) String() string { return proto.CompactTextString(m) }
func (*Start) ProtoMessage()    {}
func (*Start) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{0}
}

func (m *Start) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Start.Unmarshal(m, b)
}
func (m *Start) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Start.Marshal(b, m, deterministic)
}
func (m *Start) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Start.Merge(m, src)
}
func (m *Start) XXX_Size() int {
	return xxx_messageInfo_Start.Size(m)
}
func (m *Start) XXX_DiscardUnknown() {
	xxx_messageInfo_Start.DiscardUnknown(m)
}

var xxx_messageInfo_Start proto.InternalMessageInfo

func (m *Start) GetT() uint32 {
	if m != nil {
		return m.T
	}
	return 0
}

func (m *Start) GetAddresses() [][]byte {
	if m != nil {
		return m.Addresses
	}
	return nil
}

type Deal struct {
	Index                uint32              `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	EncryptedDeal        *Deal_EncryptedDeal `protobuf:"bytes,2,opt,name=encryptedDeal,proto3" json:"encryptedDeal,omitempty"`
	Signature            []byte              `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Deal) Reset()         { *m = Deal{} }
func (m *Deal) String() string { return proto.CompactTextString(m) }
func (*Deal) ProtoMessage()    {}
func (*Deal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1}
}

func (m *Deal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deal.Unmarshal(m, b)
}
func (m *Deal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deal.Marshal(b, m, deterministic)
}
func (m *Deal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal.Merge(m, src)
}
func (m *Deal) XXX_Size() int {
	return xxx_messageInfo_Deal.Size(m)
}
func (m *Deal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal proto.InternalMessageInfo

func (m *Deal) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Deal) GetEncryptedDeal() *Deal_EncryptedDeal {
	if m != nil {
		return m.EncryptedDeal
	}
	return nil
}

func (m *Deal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Deal_EncryptedDeal struct {
	Dhkey                []byte   `protobuf:"bytes,1,opt,name=dhkey,proto3" json:"dhkey,omitempty"`
	Signature            []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	Nonce                []byte   `protobuf:"bytes,3,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Cipher               []byte   `protobuf:"bytes,4,opt,name=cipher,proto3" json:"cipher,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deal_EncryptedDeal) Reset()         { *m = Deal_EncryptedDeal{} }
func (m *Deal_EncryptedDeal) String() string { return proto.CompactTextString(m) }
func (*Deal_EncryptedDeal) ProtoMessage()    {}
func (*Deal_EncryptedDeal) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{1, 0}
}

func (m *Deal_EncryptedDeal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deal_EncryptedDeal.Unmarshal(m, b)
}
func (m *Deal_EncryptedDeal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deal_EncryptedDeal.Marshal(b, m, deterministic)
}
func (m *Deal_EncryptedDeal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deal_EncryptedDeal.Merge(m, src)
}
func (m *Deal_EncryptedDeal) XXX_Size() int {
	return xxx_messageInfo_Deal_EncryptedDeal.Size(m)
}
func (m *Deal_EncryptedDeal) XXX_DiscardUnknown() {
	xxx_messageInfo_Deal_EncryptedDeal.DiscardUnknown(m)
}

var xxx_messageInfo_Deal_EncryptedDeal proto.InternalMessageInfo

func (m *Deal_EncryptedDeal) GetDhkey() []byte {
	if m != nil {
		return m.Dhkey
	}
	return nil
}

func (m *Deal_EncryptedDeal) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Deal_EncryptedDeal) GetNonce() []byte {
	if m != nil {
		return m.Nonce
	}
	return nil
}

func (m *Deal_EncryptedDeal) GetCipher() []byte {
	if m != nil {
		return m.Cipher
	}
	return nil
}

type Response struct {
	// Index of the Dealer for which this response is for
	Index                uint32         `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Response             *Response_Data `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Response) GetResponse() *Response_Data {
	if m != nil {
		return m.Response
	}
	return nil
}

type Response_Data struct {
	SessionID []byte `protobuf:"bytes,1,opt,name=sessionID,proto3" json:"sessionID,omitempty"`
	// Index of the verifier issuing this Response from the new set of nodes
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Status               bool     `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	Signature            []byte   `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response_Data) Reset()         { *m = Response_Data{} }
func (m *Response_Data) String() string { return proto.CompactTextString(m) }
func (*Response_Data) ProtoMessage()    {}
func (*Response_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{2, 0}
}

func (m *Response_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response_Data.Unmarshal(m, b)
}
func (m *Response_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response_Data.Marshal(b, m, deterministic)
}
func (m *Response_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response_Data.Merge(m, src)
}
func (m *Response_Data) XXX_Size() int {
	return xxx_messageInfo_Response_Data.Size(m)
}
func (m *Response_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_Response_Data.DiscardUnknown(m)
}

var xxx_messageInfo_Response_Data proto.InternalMessageInfo

func (m *Response_Data) GetSessionID() []byte {
	if m != nil {
		return m.SessionID
	}
	return nil
}

func (m *Response_Data) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Response_Data) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func (m *Response_Data) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type StartDone struct {
	PubKey               []byte   `protobuf:"bytes,1,opt,name=pubKey,proto3" json:"pubKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StartDone) Reset()         { *m = StartDone{} }
func (m *StartDone) String() string { return proto.CompactTextString(m) }
func (*StartDone) ProtoMessage()    {}
func (*StartDone) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{3}
}

func (m *StartDone) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartDone.Unmarshal(m, b)
}
func (m *StartDone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartDone.Marshal(b, m, deterministic)
}
func (m *StartDone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartDone.Merge(m, src)
}
func (m *StartDone) XXX_Size() int {
	return xxx_messageInfo_StartDone.Size(m)
}
func (m *StartDone) XXX_DiscardUnknown() {
	xxx_messageInfo_StartDone.DiscardUnknown(m)
}

var xxx_messageInfo_StartDone proto.InternalMessageInfo

func (m *StartDone) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type DecryptRequest struct {
	K                    []byte   `protobuf:"bytes,1,opt,name=K,proto3" json:"K,omitempty"`
	C                    []byte   `protobuf:"bytes,2,opt,name=C,proto3" json:"C,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptRequest) Reset()         { *m = DecryptRequest{} }
func (m *DecryptRequest) String() string { return proto.CompactTextString(m) }
func (*DecryptRequest) ProtoMessage()    {}
func (*DecryptRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{4}
}

func (m *DecryptRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptRequest.Unmarshal(m, b)
}
func (m *DecryptRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptRequest.Marshal(b, m, deterministic)
}
func (m *DecryptRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptRequest.Merge(m, src)
}
func (m *DecryptRequest) XXX_Size() int {
	return xxx_messageInfo_DecryptRequest.Size(m)
}
func (m *DecryptRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptRequest proto.InternalMessageInfo

func (m *DecryptRequest) GetK() []byte {
	if m != nil {
		return m.K
	}
	return nil
}

func (m *DecryptRequest) GetC() []byte {
	if m != nil {
		return m.C
	}
	return nil
}

type DecryptReply struct {
	V                    []byte   `protobuf:"bytes,1,opt,name=V,proto3" json:"V,omitempty"`
	I                    int64    `protobuf:"varint,2,opt,name=I,proto3" json:"I,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DecryptReply) Reset()         { *m = DecryptReply{} }
func (m *DecryptReply) String() string { return proto.CompactTextString(m) }
func (*DecryptReply) ProtoMessage()    {}
func (*DecryptReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc296cbfe5ffcd5, []int{5}
}

func (m *DecryptReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DecryptReply.Unmarshal(m, b)
}
func (m *DecryptReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DecryptReply.Marshal(b, m, deterministic)
}
func (m *DecryptReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DecryptReply.Merge(m, src)
}
func (m *DecryptReply) XXX_Size() int {
	return xxx_messageInfo_DecryptReply.Size(m)
}
func (m *DecryptReply) XXX_DiscardUnknown() {
	xxx_messageInfo_DecryptReply.DiscardUnknown(m)
}

var xxx_messageInfo_DecryptReply proto.InternalMessageInfo

func (m *DecryptReply) GetV() []byte {
	if m != nil {
		return m.V
	}
	return nil
}

func (m *DecryptReply) GetI() int64 {
	if m != nil {
		return m.I
	}
	return 0
}

func init() {
	proto.RegisterType((*Start)(nil), "pedersen.Start")
	proto.RegisterType((*Deal)(nil), "pedersen.Deal")
	proto.RegisterType((*Deal_EncryptedDeal)(nil), "pedersen.Deal.EncryptedDeal")
	proto.RegisterType((*Response)(nil), "pedersen.Response")
	proto.RegisterType((*Response_Data)(nil), "pedersen.Response.Data")
	proto.RegisterType((*StartDone)(nil), "pedersen.StartDone")
	proto.RegisterType((*DecryptRequest)(nil), "pedersen.DecryptRequest")
	proto.RegisterType((*DecryptReply)(nil), "pedersen.DecryptReply")
}

func init() { proto.RegisterFile("messages.proto", fileDescriptor_4dc296cbfe5ffcd5) }

var fileDescriptor_4dc296cbfe5ffcd5 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xc1, 0x8e, 0x9b, 0x30,
	0x14, 0x94, 0x09, 0x1b, 0xb1, 0xaf, 0x64, 0x0f, 0xa8, 0xda, 0xd2, 0x68, 0x0f, 0x88, 0x5e, 0xa2,
	0xaa, 0x62, 0xa5, 0xe6, 0x0f, 0xba, 0xf4, 0x10, 0xe5, 0xe6, 0x4a, 0x7b, 0x77, 0xc2, 0x2b, 0x41,
	0x4d, 0x6d, 0xc7, 0x36, 0x52, 0xf9, 0xc2, 0x7e, 0x4f, 0xff, 0xa0, 0xb2, 0x31, 0x10, 0x22, 0xed,
	0x71, 0x86, 0x99, 0xe7, 0x99, 0xc7, 0x83, 0x87, 0xdf, 0xa8, 0x35, 0xab, 0x51, 0x17, 0x52, 0x09,
	0x23, 0x92, 0x48, 0x62, 0x85, 0x4a, 0x23, 0x5f, 0x7f, 0xac, 0x85, 0xa8, 0xcf, 0xf8, 0xec, 0xf8,
	0x43, 0xfb, 0xf3, 0x99, 0xf1, 0xae, 0x17, 0xe5, 0x5b, 0xb8, 0xfb, 0x61, 0x98, 0x32, 0x49, 0x0c,
	0xc4, 0xa4, 0x24, 0x23, 0x9b, 0x15, 0x25, 0x26, 0x79, 0x82, 0x7b, 0x56, 0x55, 0x0a, 0xb5, 0x46,
	0x9d, 0x06, 0xd9, 0x62, 0x13, 0xd3, 0x89, 0xc8, 0xff, 0x11, 0x08, 0x4b, 0x64, 0xe7, 0xe4, 0x3d,
	0xdc, 0x35, 0xbc, 0xc2, 0x3f, 0xde, 0xd8, 0x83, 0xe4, 0x1b, 0xac, 0x90, 0x1f, 0x55, 0x27, 0x0d,
	0x56, 0x56, 0x96, 0x06, 0x19, 0xd9, 0xbc, 0xfb, 0xfa, 0x54, 0x0c, 0x81, 0x0a, 0xcb, 0x16, 0xdf,
	0xaf, 0x35, 0x74, 0x6e, 0xb1, 0x01, 0x74, 0x53, 0x73, 0x66, 0x5a, 0x85, 0xe9, 0x22, 0x23, 0x36,
	0xc0, 0x48, 0xac, 0x2f, 0xb0, 0x9a, 0xb9, 0x6d, 0x90, 0xea, 0xf4, 0x0b, 0x3b, 0x17, 0x24, 0xa6,
	0x3d, 0x98, 0x0f, 0x09, 0x6e, 0x86, 0x58, 0x0f, 0x17, 0xfc, 0x38, 0x8c, 0xef, 0x41, 0xf2, 0x08,
	0xcb, 0x63, 0x23, 0x4f, 0xa8, 0xd2, 0xd0, 0xd1, 0x1e, 0xe5, 0x7f, 0x09, 0x44, 0x14, 0xb5, 0x14,
	0x5c, 0xe3, 0x1b, 0xbd, 0xb7, 0x10, 0x29, 0xaf, 0xf0, 0x95, 0x3f, 0x4c, 0x95, 0x07, 0x6f, 0x51,
	0x32, 0xc3, 0xe8, 0x28, 0x5c, 0x4b, 0x08, 0x2d, 0xe3, 0xb2, 0xa2, 0xd6, 0x8d, 0xe0, 0xbb, 0xd2,
	0xb7, 0x98, 0x88, 0xe9, 0xc1, 0xe0, 0xfa, 0xc1, 0x47, 0x58, 0x6a, 0xc3, 0x4c, 0xab, 0x5d, 0x85,
	0x88, 0x7a, 0x34, 0xef, 0x1d, 0xde, 0xf4, 0xce, 0x3f, 0xc1, 0xbd, 0xfb, 0xe5, 0xa5, 0xe0, 0xae,
	0xae, 0x6c, 0x0f, 0xfb, 0x71, 0x73, 0x1e, 0xe5, 0x5f, 0xe0, 0xa1, 0x44, 0xb7, 0x61, 0x8a, 0x97,
	0x16, 0xb5, 0x3b, 0x90, 0xbd, 0x17, 0x91, 0xbd, 0x45, 0x2f, 0x7e, 0xa5, 0xe4, 0x25, 0xff, 0x0c,
	0xf1, 0xa8, 0x96, 0xe7, 0xce, 0x7e, 0x7d, 0x1d, 0xb4, 0xaf, 0x16, 0xed, 0x9c, 0x76, 0x41, 0xc9,
	0xee, 0xb0, 0x74, 0x87, 0xb7, 0xfd, 0x1f, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x77, 0x1d, 0x92, 0xaf,
	0x02, 0x00, 0x00,
}
