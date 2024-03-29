// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

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

type ReqSignup struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignup) Reset()         { *m = ReqSignup{} }
func (m *ReqSignup) String() string { return proto.CompactTextString(m) }
func (*ReqSignup) ProtoMessage()    {}
func (*ReqSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *ReqSignup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignup.Unmarshal(m, b)
}
func (m *ReqSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignup.Marshal(b, m, deterministic)
}
func (m *ReqSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignup.Merge(m, src)
}
func (m *ReqSignup) XXX_Size() int {
	return xxx_messageInfo_ReqSignup.Size(m)
}
func (m *ReqSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignup.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignup proto.InternalMessageInfo

func (m *ReqSignup) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignup) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignup struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignup) Reset()         { *m = RespSignup{} }
func (m *RespSignup) String() string { return proto.CompactTextString(m) }
func (*RespSignup) ProtoMessage()    {}
func (*RespSignup) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *RespSignup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignup.Unmarshal(m, b)
}
func (m *RespSignup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignup.Marshal(b, m, deterministic)
}
func (m *RespSignup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignup.Merge(m, src)
}
func (m *RespSignup) XXX_Size() int {
	return xxx_messageInfo_RespSignup.Size(m)
}
func (m *RespSignup) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignup.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignup proto.InternalMessageInfo

func (m *RespSignup) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignup) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqSignin struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignin) Reset()         { *m = ReqSignin{} }
func (m *ReqSignin) String() string { return proto.CompactTextString(m) }
func (*ReqSignin) ProtoMessage()    {}
func (*ReqSignin) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *ReqSignin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignin.Unmarshal(m, b)
}
func (m *ReqSignin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignin.Marshal(b, m, deterministic)
}
func (m *ReqSignin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignin.Merge(m, src)
}
func (m *ReqSignin) XXX_Size() int {
	return xxx_messageInfo_ReqSignin.Size(m)
}
func (m *ReqSignin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignin proto.InternalMessageInfo

func (m *ReqSignin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type RespSignin struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignin) Reset()         { *m = RespSignin{} }
func (m *RespSignin) String() string { return proto.CompactTextString(m) }
func (*RespSignin) ProtoMessage()    {}
func (*RespSignin) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *RespSignin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignin.Unmarshal(m, b)
}
func (m *RespSignin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignin.Marshal(b, m, deterministic)
}
func (m *RespSignin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignin.Merge(m, src)
}
func (m *RespSignin) XXX_Size() int {
	return xxx_messageInfo_RespSignin.Size(m)
}
func (m *RespSignin) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignin.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignin proto.InternalMessageInfo

func (m *RespSignin) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *RespSignin) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqUserInfo struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserInfo) Reset()         { *m = ReqUserInfo{} }
func (m *ReqUserInfo) String() string { return proto.CompactTextString(m) }
func (*ReqUserInfo) ProtoMessage()    {}
func (*ReqUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *ReqUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserInfo.Unmarshal(m, b)
}
func (m *ReqUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserInfo.Marshal(b, m, deterministic)
}
func (m *ReqUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserInfo.Merge(m, src)
}
func (m *ReqUserInfo) XXX_Size() int {
	return xxx_messageInfo_ReqUserInfo.Size(m)
}
func (m *ReqUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserInfo proto.InternalMessageInfo

func (m *ReqUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RespUserInfo struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string   `protobuf:"bytes,5,opt,name=phone,proto3" json:"phone,omitempty"`
	SignupAt             string   `protobuf:"bytes,6,opt,name=signupAt,proto3" json:"signupAt,omitempty"`
	LastActveAt          string   `protobuf:"bytes,7,opt,name=lastActveAt,proto3" json:"lastActveAt,omitempty"`
	Status               int32    `protobuf:"varint,8,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUserInfo) Reset()         { *m = RespUserInfo{} }
func (m *RespUserInfo) String() string { return proto.CompactTextString(m) }
func (*RespUserInfo) ProtoMessage()    {}
func (*RespUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{5}
}

func (m *RespUserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUserInfo.Unmarshal(m, b)
}
func (m *RespUserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUserInfo.Marshal(b, m, deterministic)
}
func (m *RespUserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUserInfo.Merge(m, src)
}
func (m *RespUserInfo) XXX_Size() int {
	return xxx_messageInfo_RespUserInfo.Size(m)
}
func (m *RespUserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RespUserInfo proto.InternalMessageInfo

func (m *RespUserInfo) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserInfo) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserInfo) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *RespUserInfo) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RespUserInfo) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RespUserInfo) GetSignupAt() string {
	if m != nil {
		return m.SignupAt
	}
	return ""
}

func (m *RespUserInfo) GetLastActveAt() string {
	if m != nil {
		return m.LastActveAt
	}
	return ""
}

func (m *RespUserInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReqUserFile struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUserFile) Reset()         { *m = ReqUserFile{} }
func (m *ReqUserFile) String() string { return proto.CompactTextString(m) }
func (*ReqUserFile) ProtoMessage()    {}
func (*ReqUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{6}
}

func (m *ReqUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUserFile.Unmarshal(m, b)
}
func (m *ReqUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUserFile.Marshal(b, m, deterministic)
}
func (m *ReqUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUserFile.Merge(m, src)
}
func (m *ReqUserFile) XXX_Size() int {
	return xxx_messageInfo_ReqUserFile.Size(m)
}
func (m *ReqUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUserFile proto.InternalMessageInfo

func (m *ReqUserFile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUserFile) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type RespUserFile struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FileData             string   `protobuf:"bytes,3,opt,name=fileData,proto3" json:"fileData,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUserFile) Reset()         { *m = RespUserFile{} }
func (m *RespUserFile) String() string { return proto.CompactTextString(m) }
func (*RespUserFile) ProtoMessage()    {}
func (*RespUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{7}
}

func (m *RespUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUserFile.Unmarshal(m, b)
}
func (m *RespUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUserFile.Marshal(b, m, deterministic)
}
func (m *RespUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUserFile.Merge(m, src)
}
func (m *RespUserFile) XXX_Size() int {
	return xxx_messageInfo_RespUserFile.Size(m)
}
func (m *RespUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_RespUserFile proto.InternalMessageInfo

func (m *RespUserFile) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUserFile) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespUserFile) GetFileData() string {
	if m != nil {
		return m.FileData
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqSignup)(nil), "proto.ReqSignup")
	proto.RegisterType((*RespSignup)(nil), "proto.RespSignup")
	proto.RegisterType((*ReqSignin)(nil), "proto.ReqSignin")
	proto.RegisterType((*RespSignin)(nil), "proto.RespSignin")
	proto.RegisterType((*ReqUserInfo)(nil), "proto.ReqUserInfo")
	proto.RegisterType((*RespUserInfo)(nil), "proto.RespUserInfo")
	proto.RegisterType((*ReqUserFile)(nil), "proto.ReqUserFile")
	proto.RegisterType((*RespUserFile)(nil), "proto.RespUserFile")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 380 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x52, 0xcf, 0x6f, 0xda, 0x30,
	0x18, 0x4d, 0x06, 0x09, 0xf0, 0xb1, 0xc3, 0xe6, 0xa1, 0xc9, 0xe2, 0x84, 0x7c, 0xda, 0x2e, 0x4c,
	0xda, 0xb4, 0x1d, 0x76, 0xa9, 0x50, 0xab, 0x4a, 0xbd, 0x55, 0x41, 0x95, 0x7a, 0x75, 0xe1, 0x83,
	0x5a, 0x4d, 0x9c, 0x10, 0x1b, 0xfa, 0xbf, 0xf6, 0xd0, 0xbf, 0xa5, 0xb2, 0x1d, 0x27, 0x01, 0x2a,
	0x24, 0x7a, 0x82, 0xf7, 0x3e, 0xbf, 0x97, 0xf7, 0xfd, 0x00, 0xd8, 0x2a, 0x2c, 0xa7, 0x45, 0x99,
	0xeb, 0x9c, 0x44, 0xf6, 0x87, 0x5d, 0xc2, 0x20, 0xc1, 0xcd, 0x5c, 0xac, 0xe5, 0xb6, 0x20, 0x63,
	0xe8, 0x9b, 0x17, 0x92, 0x67, 0x48, 0xc3, 0x49, 0xf8, 0x63, 0x90, 0xd4, 0xd8, 0xd4, 0x0a, 0xae,
	0xd4, 0x73, 0x5e, 0x2e, 0xe9, 0x27, 0x57, 0xf3, 0x98, 0xfd, 0x07, 0x48, 0x50, 0x15, 0x95, 0x0b,
	0x81, 0xee, 0x22, 0x5f, 0x3a, 0x87, 0x28, 0xb1, 0xff, 0x09, 0x85, 0x5e, 0x86, 0x4a, 0xf1, 0x35,
	0x56, 0x62, 0x0f, 0x5b, 0x01, 0x84, 0xfc, 0x70, 0x80, 0xdb, 0x26, 0x80, 0x90, 0xef, 0x06, 0x18,
	0x41, 0xa4, 0xf3, 0x27, 0x94, 0x95, 0xd4, 0x81, 0x76, 0xac, 0xce, 0x7e, 0xac, 0x9f, 0x30, 0x4c,
	0x70, 0x73, 0xa7, 0xb0, 0xbc, 0x91, 0xab, 0xfc, 0x54, 0x30, 0xf6, 0x12, 0xc2, 0x67, 0xf3, 0xf5,
	0xfa, 0xf1, 0x59, 0x03, 0xd8, 0xb3, 0xee, 0x1c, 0xf4, 0x3c, 0x82, 0x08, 0x33, 0x2e, 0x52, 0xda,
	0x75, 0xa9, 0x2d, 0x30, 0x6c, 0xf1, 0x98, 0x4b, 0xa4, 0x91, 0x63, 0x2d, 0x30, 0x3e, 0xca, 0x2e,
	0x60, 0xa6, 0x69, 0xec, 0x7c, 0x3c, 0x26, 0x13, 0x18, 0xa6, 0x5c, 0xe9, 0xd9, 0x42, 0xef, 0x70,
	0xa6, 0x69, 0xcf, 0x96, 0xdb, 0x14, 0xf9, 0x0e, 0xb1, 0xd2, 0x5c, 0x6f, 0x15, 0xed, 0xdb, 0xd4,
	0x15, 0x62, 0x17, 0xf5, 0x1c, 0xae, 0x45, 0x8a, 0x27, 0x17, 0x34, 0x82, 0x28, 0x15, 0x99, 0xd0,
	0xb6, 0xc1, 0x28, 0x71, 0x80, 0xdd, 0x37, 0xc3, 0xb1, 0x0e, 0x67, 0x0f, 0x67, 0x25, 0x52, 0xbc,
	0xe2, 0x9a, 0xfb, 0xe1, 0x78, 0xfc, 0xfb, 0x35, 0x84, 0xa1, 0xb1, 0x9d, 0x63, 0xb9, 0x13, 0x0b,
	0x24, 0xbf, 0x20, 0xae, 0x2e, 0xf0, 0x8b, 0xbb, 0xf1, 0x69, 0x7d, 0xd9, 0xe3, 0xaf, 0x35, 0xe3,
	0xcf, 0x94, 0x05, 0x5e, 0x20, 0xe4, 0xa1, 0x40, 0xc8, 0x23, 0x81, 0x90, 0x2c, 0x20, 0x7f, 0xa1,
	0xdf, 0x2c, 0xb9, 0x91, 0x78, 0x6e, 0xfc, 0xad, 0x25, 0xf2, 0x24, 0x0b, 0xc8, 0x3f, 0x18, 0xf8,
	0xf6, 0xd5, 0xa1, 0xce, 0x90, 0x47, 0x3a, 0x43, 0xb2, 0xe0, 0x21, 0xb6, 0xec, 0x9f, 0xb7, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x30, 0x05, 0x7b, 0xd1, 0xb7, 0x03, 0x00, 0x00,
}
