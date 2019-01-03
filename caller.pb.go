// Code generated by protoc-gen-go. DO NOT EDIT.
// source: caller.proto

package errors

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Caller of error.
type Caller struct {
	// Source file of the error.
	File string `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	// Source function of the error.
	Function string `protobuf:"bytes,2,opt,name=function,proto3" json:"function,omitempty"`
	// Source line of the error.
	Line int64 `protobuf:"zigzag64,3,opt,name=line,proto3" json:"line,omitempty"`
	// Stack of the error.
	Stacks               []string `protobuf:"bytes,4,rep,name=stacks,proto3" json:"stacks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Caller) Reset()         { *m = Caller{} }
func (m *Caller) String() string { return proto.CompactTextString(m) }
func (*Caller) ProtoMessage()    {}
func (*Caller) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d14de5595b84d78, []int{0}
}

func (m *Caller) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Caller.Unmarshal(m, b)
}
func (m *Caller) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Caller.Marshal(b, m, deterministic)
}
func (m *Caller) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Caller.Merge(m, src)
}
func (m *Caller) XXX_Size() int {
	return xxx_messageInfo_Caller.Size(m)
}
func (m *Caller) XXX_DiscardUnknown() {
	xxx_messageInfo_Caller.DiscardUnknown(m)
}

var xxx_messageInfo_Caller proto.InternalMessageInfo

func (m *Caller) GetFile() string {
	if m != nil {
		return m.File
	}
	return ""
}

func (m *Caller) GetFunction() string {
	if m != nil {
		return m.Function
	}
	return ""
}

func (m *Caller) GetLine() int64 {
	if m != nil {
		return m.Line
	}
	return 0
}

func (m *Caller) GetStacks() []string {
	if m != nil {
		return m.Stacks
	}
	return nil
}

func init() {
	proto.RegisterType((*Caller)(nil), "errors.Caller")
}

func init() { proto.RegisterFile("caller.proto", fileDescriptor_0d14de5595b84d78) }

var fileDescriptor_0d14de5595b84d78 = []byte{
	// 126 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0x4e, 0xcc, 0xc9,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x2a, 0xca, 0x2f, 0x2a,
	0x56, 0x4a, 0xe1, 0x62, 0x73, 0x06, 0x8b, 0x0b, 0x09, 0x71, 0xb1, 0xa4, 0x65, 0xe6, 0xa4, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x42, 0x52, 0x5c, 0x1c, 0x69, 0xa5, 0x79, 0xc9,
	0x25, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x60, 0x71, 0x38, 0x1f, 0xa4, 0x3e, 0x27, 0x33, 0x2f, 0x55,
	0x82, 0x59, 0x81, 0x51, 0x43, 0x28, 0x08, 0xcc, 0x16, 0x12, 0xe3, 0x62, 0x2b, 0x2e, 0x49, 0x4c,
	0xce, 0x2e, 0x96, 0x60, 0x51, 0x60, 0xd6, 0xe0, 0x0c, 0x82, 0xf2, 0x92, 0xd8, 0xc0, 0x96, 0x1a,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xf8, 0xef, 0xe4, 0x84, 0x00, 0x00, 0x00,
}