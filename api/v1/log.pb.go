// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/log.proto

package log_v1

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

type Record struct {
	Value                []byte   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Offset               uint64   `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_19a5c3fde3f7ae80, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Record) GetOffset() uint64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*Record)(nil), "log.v1.Record")
}

func init() {
	proto.RegisterFile("api/v1/log.proto", fileDescriptor_19a5c3fde3f7ae80)
}

var fileDescriptor_19a5c3fde3f7ae80 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0x2c, 0xc8, 0xd4,
	0x2f, 0x33, 0xd4, 0xcf, 0xc9, 0x4f, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x31,
	0xcb, 0x0c, 0x95, 0xcc, 0xb8, 0xd8, 0x82, 0x52, 0x93, 0xf3, 0x8b, 0x52, 0x84, 0x44, 0xb8, 0x58,
	0xcb, 0x12, 0x73, 0x4a, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0x20, 0x1c, 0x21, 0x31,
	0x2e, 0xb6, 0xfc, 0xb4, 0xb4, 0xe2, 0xd4, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x28,
	0xcf, 0x49, 0x3e, 0x4a, 0x36, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f,
	0x31, 0x27, 0xb5, 0x22, 0x2b, 0xb1, 0x5c, 0x1f, 0x64, 0x4d, 0x4e, 0x7e, 0x7a, 0x7c, 0x99, 0x61,
	0x12, 0x1b, 0xd8, 0x1e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd6, 0x69, 0x5a, 0xa0, 0x7b,
	0x00, 0x00, 0x00,
}
