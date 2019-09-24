// Code generated by protoc-gen-go. DO NOT EDIT.
// source: upload.proto

package go_micro_service_upload

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type ReqEntry struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqEntry) Reset()         { *m = ReqEntry{} }
func (m *ReqEntry) String() string { return proto.CompactTextString(m) }
func (*ReqEntry) ProtoMessage()    {}
func (*ReqEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b94b655bd2a7e5, []int{0}
}

func (m *ReqEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqEntry.Unmarshal(m, b)
}
func (m *ReqEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqEntry.Marshal(b, m, deterministic)
}
func (m *ReqEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqEntry.Merge(m, src)
}
func (m *ReqEntry) XXX_Size() int {
	return xxx_messageInfo_ReqEntry.Size(m)
}
func (m *ReqEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqEntry.DiscardUnknown(m)
}

var xxx_messageInfo_ReqEntry proto.InternalMessageInfo

type RespEntry struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Entry                string   `protobuf:"bytes,3,opt,name=entry,proto3" json:"entry,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespEntry) Reset()         { *m = RespEntry{} }
func (m *RespEntry) String() string { return proto.CompactTextString(m) }
func (*RespEntry) ProtoMessage()    {}
func (*RespEntry) Descriptor() ([]byte, []int) {
	return fileDescriptor_91b94b655bd2a7e5, []int{1}
}

func (m *RespEntry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespEntry.Unmarshal(m, b)
}
func (m *RespEntry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespEntry.Marshal(b, m, deterministic)
}
func (m *RespEntry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespEntry.Merge(m, src)
}
func (m *RespEntry) XXX_Size() int {
	return xxx_messageInfo_RespEntry.Size(m)
}
func (m *RespEntry) XXX_DiscardUnknown() {
	xxx_messageInfo_RespEntry.DiscardUnknown(m)
}

var xxx_messageInfo_RespEntry proto.InternalMessageInfo

func (m *RespEntry) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespEntry) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespEntry) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqEntry)(nil), "go.micro.service.upload.ReqEntry")
	proto.RegisterType((*RespEntry)(nil), "go.micro.service.upload.RespEntry")
}

func init() { proto.RegisterFile("upload.proto", fileDescriptor_91b94b655bd2a7e5) }

var fileDescriptor_91b94b655bd2a7e5 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x2d, 0xc8, 0xc9,
	0x4f, 0x4c, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4f, 0xcf, 0xd7, 0xcb, 0xcd, 0x4c,
	0x2e, 0xca, 0xd7, 0x2b, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x83, 0x48, 0x2b, 0x71, 0x71,
	0x71, 0x04, 0xa5, 0x16, 0xba, 0xe6, 0x95, 0x14, 0x55, 0x2a, 0xf9, 0x73, 0x71, 0x06, 0xa5, 0x16,
	0x17, 0x80, 0x39, 0x42, 0x42, 0x5c, 0x2c, 0xc9, 0xf9, 0x29, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a,
	0xac, 0x41, 0x60, 0xb6, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71, 0x71, 0x62, 0x7a, 0xaa, 0x04,
	0x93, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc2, 0xc5, 0x9a, 0x0a, 0xd2, 0x26, 0xc1,
	0x0c, 0x16, 0x87, 0x70, 0x8c, 0x52, 0xb9, 0x78, 0x43, 0xc1, 0xd6, 0x04, 0x43, 0x2c, 0x15, 0x0a,
	0xe1, 0xe2, 0x86, 0x08, 0x40, 0xec, 0x50, 0xd4, 0xc3, 0xe1, 0x2c, 0x3d, 0x98, 0x9b, 0xa4, 0x94,
	0xf0, 0x28, 0x81, 0x3a, 0x35, 0x89, 0x0d, 0xec, 0x47, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x27, 0x04, 0x3a, 0x5f, 0xf3, 0x00, 0x00, 0x00,
}
