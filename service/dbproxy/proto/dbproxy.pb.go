// Code generated by protoc-gen-go. DO NOT EDIT.
// source: dbproxy.proto

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

type ReqUploadFile struct {
	FileSha1             string   `protobuf:"bytes,1,opt,name=fileSha1,proto3" json:"fileSha1,omitempty"`
	FileName             string   `protobuf:"bytes,2,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileSize             int64    `protobuf:"varint,3,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	FileAddr             string   `protobuf:"bytes,4,opt,name=fileAddr,proto3" json:"fileAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUploadFile) Reset()         { *m = ReqUploadFile{} }
func (m *ReqUploadFile) String() string { return proto.CompactTextString(m) }
func (*ReqUploadFile) ProtoMessage()    {}
func (*ReqUploadFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{0}
}

func (m *ReqUploadFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUploadFile.Unmarshal(m, b)
}
func (m *ReqUploadFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUploadFile.Marshal(b, m, deterministic)
}
func (m *ReqUploadFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUploadFile.Merge(m, src)
}
func (m *ReqUploadFile) XXX_Size() int {
	return xxx_messageInfo_ReqUploadFile.Size(m)
}
func (m *ReqUploadFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUploadFile.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUploadFile proto.InternalMessageInfo

func (m *ReqUploadFile) GetFileSha1() string {
	if m != nil {
		return m.FileSha1
	}
	return ""
}

func (m *ReqUploadFile) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *ReqUploadFile) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *ReqUploadFile) GetFileAddr() string {
	if m != nil {
		return m.FileAddr
	}
	return ""
}

type RespUploadFile struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUploadFile) Reset()         { *m = RespUploadFile{} }
func (m *RespUploadFile) String() string { return proto.CompactTextString(m) }
func (*RespUploadFile) ProtoMessage()    {}
func (*RespUploadFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{1}
}

func (m *RespUploadFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUploadFile.Unmarshal(m, b)
}
func (m *RespUploadFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUploadFile.Marshal(b, m, deterministic)
}
func (m *RespUploadFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUploadFile.Merge(m, src)
}
func (m *RespUploadFile) XXX_Size() int {
	return xxx_messageInfo_RespUploadFile.Size(m)
}
func (m *RespUploadFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUploadFile.DiscardUnknown(m)
}

var xxx_messageInfo_RespUploadFile proto.InternalMessageInfo

func (m *RespUploadFile) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUploadFile) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqFileMeta struct {
	FileSha1             string   `protobuf:"bytes,1,opt,name=fileSha1,proto3" json:"fileSha1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqFileMeta) Reset()         { *m = ReqFileMeta{} }
func (m *ReqFileMeta) String() string { return proto.CompactTextString(m) }
func (*ReqFileMeta) ProtoMessage()    {}
func (*ReqFileMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{2}
}

func (m *ReqFileMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqFileMeta.Unmarshal(m, b)
}
func (m *ReqFileMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqFileMeta.Marshal(b, m, deterministic)
}
func (m *ReqFileMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqFileMeta.Merge(m, src)
}
func (m *ReqFileMeta) XXX_Size() int {
	return xxx_messageInfo_ReqFileMeta.Size(m)
}
func (m *ReqFileMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqFileMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ReqFileMeta proto.InternalMessageInfo

func (m *ReqFileMeta) GetFileSha1() string {
	if m != nil {
		return m.FileSha1
	}
	return ""
}

type RespFileMeta struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	FileSha1             string   `protobuf:"bytes,3,opt,name=fileSha1,proto3" json:"fileSha1,omitempty"`
	FileName             string   `protobuf:"bytes,4,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileSize             int64    `protobuf:"varint,5,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	FileAddr             string   `protobuf:"bytes,6,opt,name=fileAddr,proto3" json:"fileAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespFileMeta) Reset()         { *m = RespFileMeta{} }
func (m *RespFileMeta) String() string { return proto.CompactTextString(m) }
func (*RespFileMeta) ProtoMessage()    {}
func (*RespFileMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{3}
}

func (m *RespFileMeta) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespFileMeta.Unmarshal(m, b)
}
func (m *RespFileMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespFileMeta.Marshal(b, m, deterministic)
}
func (m *RespFileMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespFileMeta.Merge(m, src)
}
func (m *RespFileMeta) XXX_Size() int {
	return xxx_messageInfo_RespFileMeta.Size(m)
}
func (m *RespFileMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_RespFileMeta.DiscardUnknown(m)
}

var xxx_messageInfo_RespFileMeta proto.InternalMessageInfo

func (m *RespFileMeta) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespFileMeta) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespFileMeta) GetFileSha1() string {
	if m != nil {
		return m.FileSha1
	}
	return ""
}

func (m *RespFileMeta) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *RespFileMeta) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

func (m *RespFileMeta) GetFileAddr() string {
	if m != nil {
		return m.FileAddr
	}
	return ""
}

type ReqUpdateFileLocation struct {
	FileSha1             string   `protobuf:"bytes,1,opt,name=fileSha1,proto3" json:"fileSha1,omitempty"`
	FileAddr             string   `protobuf:"bytes,2,opt,name=fileAddr,proto3" json:"fileAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUpdateFileLocation) Reset()         { *m = ReqUpdateFileLocation{} }
func (m *ReqUpdateFileLocation) String() string { return proto.CompactTextString(m) }
func (*ReqUpdateFileLocation) ProtoMessage()    {}
func (*ReqUpdateFileLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{4}
}

func (m *ReqUpdateFileLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUpdateFileLocation.Unmarshal(m, b)
}
func (m *ReqUpdateFileLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUpdateFileLocation.Marshal(b, m, deterministic)
}
func (m *ReqUpdateFileLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUpdateFileLocation.Merge(m, src)
}
func (m *ReqUpdateFileLocation) XXX_Size() int {
	return xxx_messageInfo_ReqUpdateFileLocation.Size(m)
}
func (m *ReqUpdateFileLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUpdateFileLocation.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUpdateFileLocation proto.InternalMessageInfo

func (m *ReqUpdateFileLocation) GetFileSha1() string {
	if m != nil {
		return m.FileSha1
	}
	return ""
}

func (m *ReqUpdateFileLocation) GetFileAddr() string {
	if m != nil {
		return m.FileAddr
	}
	return ""
}

type RespUpdateFileLocation struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUpdateFileLocation) Reset()         { *m = RespUpdateFileLocation{} }
func (m *RespUpdateFileLocation) String() string { return proto.CompactTextString(m) }
func (*RespUpdateFileLocation) ProtoMessage()    {}
func (*RespUpdateFileLocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{5}
}

func (m *RespUpdateFileLocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUpdateFileLocation.Unmarshal(m, b)
}
func (m *RespUpdateFileLocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUpdateFileLocation.Marshal(b, m, deterministic)
}
func (m *RespUpdateFileLocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUpdateFileLocation.Merge(m, src)
}
func (m *RespUpdateFileLocation) XXX_Size() int {
	return xxx_messageInfo_RespUpdateFileLocation.Size(m)
}
func (m *RespUpdateFileLocation) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUpdateFileLocation.DiscardUnknown(m)
}

var xxx_messageInfo_RespUpdateFileLocation proto.InternalMessageInfo

func (m *RespUpdateFileLocation) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUpdateFileLocation) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqUploadUserFile struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	FileSha1             string   `protobuf:"bytes,2,opt,name=fileSha1,proto3" json:"fileSha1,omitempty"`
	FileName             string   `protobuf:"bytes,3,opt,name=fileName,proto3" json:"fileName,omitempty"`
	FileSize             int64    `protobuf:"varint,4,opt,name=fileSize,proto3" json:"fileSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUploadUserFile) Reset()         { *m = ReqUploadUserFile{} }
func (m *ReqUploadUserFile) String() string { return proto.CompactTextString(m) }
func (*ReqUploadUserFile) ProtoMessage()    {}
func (*ReqUploadUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{6}
}

func (m *ReqUploadUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUploadUserFile.Unmarshal(m, b)
}
func (m *ReqUploadUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUploadUserFile.Marshal(b, m, deterministic)
}
func (m *ReqUploadUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUploadUserFile.Merge(m, src)
}
func (m *ReqUploadUserFile) XXX_Size() int {
	return xxx_messageInfo_ReqUploadUserFile.Size(m)
}
func (m *ReqUploadUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUploadUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUploadUserFile proto.InternalMessageInfo

func (m *ReqUploadUserFile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUploadUserFile) GetFileSha1() string {
	if m != nil {
		return m.FileSha1
	}
	return ""
}

func (m *ReqUploadUserFile) GetFileName() string {
	if m != nil {
		return m.FileName
	}
	return ""
}

func (m *ReqUploadUserFile) GetFileSize() int64 {
	if m != nil {
		return m.FileSize
	}
	return 0
}

type RespUploadUserFile struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUploadUserFile) Reset()         { *m = RespUploadUserFile{} }
func (m *RespUploadUserFile) String() string { return proto.CompactTextString(m) }
func (*RespUploadUserFile) ProtoMessage()    {}
func (*RespUploadUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{7}
}

func (m *RespUploadUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUploadUserFile.Unmarshal(m, b)
}
func (m *RespUploadUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUploadUserFile.Marshal(b, m, deterministic)
}
func (m *RespUploadUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUploadUserFile.Merge(m, src)
}
func (m *RespUploadUserFile) XXX_Size() int {
	return xxx_messageInfo_RespUploadUserFile.Size(m)
}
func (m *RespUploadUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUploadUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_RespUploadUserFile proto.InternalMessageInfo

func (m *RespUploadUserFile) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUploadUserFile) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqQueryUserFile struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Limit                int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqQueryUserFile) Reset()         { *m = ReqQueryUserFile{} }
func (m *ReqQueryUserFile) String() string { return proto.CompactTextString(m) }
func (*ReqQueryUserFile) ProtoMessage()    {}
func (*ReqQueryUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{8}
}

func (m *ReqQueryUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqQueryUserFile.Unmarshal(m, b)
}
func (m *ReqQueryUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqQueryUserFile.Marshal(b, m, deterministic)
}
func (m *ReqQueryUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqQueryUserFile.Merge(m, src)
}
func (m *ReqQueryUserFile) XXX_Size() int {
	return xxx_messageInfo_ReqQueryUserFile.Size(m)
}
func (m *ReqQueryUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqQueryUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_ReqQueryUserFile proto.InternalMessageInfo

func (m *ReqQueryUserFile) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqQueryUserFile) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type RespQueryUserFile struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	UserFileByte         []byte   `protobuf:"bytes,3,opt,name=userFileByte,proto3" json:"userFileByte,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespQueryUserFile) Reset()         { *m = RespQueryUserFile{} }
func (m *RespQueryUserFile) String() string { return proto.CompactTextString(m) }
func (*RespQueryUserFile) ProtoMessage()    {}
func (*RespQueryUserFile) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{9}
}

func (m *RespQueryUserFile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespQueryUserFile.Unmarshal(m, b)
}
func (m *RespQueryUserFile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespQueryUserFile.Marshal(b, m, deterministic)
}
func (m *RespQueryUserFile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespQueryUserFile.Merge(m, src)
}
func (m *RespQueryUserFile) XXX_Size() int {
	return xxx_messageInfo_RespQueryUserFile.Size(m)
}
func (m *RespQueryUserFile) XXX_DiscardUnknown() {
	xxx_messageInfo_RespQueryUserFile.DiscardUnknown(m)
}

var xxx_messageInfo_RespQueryUserFile proto.InternalMessageInfo

func (m *RespQueryUserFile) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespQueryUserFile) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespQueryUserFile) GetUserFileByte() []byte {
	if m != nil {
		return m.UserFileByte
	}
	return nil
}

type ReqSignupUser struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSignupUser) Reset()         { *m = ReqSignupUser{} }
func (m *ReqSignupUser) String() string { return proto.CompactTextString(m) }
func (*ReqSignupUser) ProtoMessage()    {}
func (*ReqSignupUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{10}
}

func (m *ReqSignupUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSignupUser.Unmarshal(m, b)
}
func (m *ReqSignupUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSignupUser.Marshal(b, m, deterministic)
}
func (m *ReqSignupUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSignupUser.Merge(m, src)
}
func (m *ReqSignupUser) XXX_Size() int {
	return xxx_messageInfo_ReqSignupUser.Size(m)
}
func (m *ReqSignupUser) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSignupUser.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSignupUser proto.InternalMessageInfo

func (m *ReqSignupUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSignupUser) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type RespSignupUser struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSignupUser) Reset()         { *m = RespSignupUser{} }
func (m *RespSignupUser) String() string { return proto.CompactTextString(m) }
func (*RespSignupUser) ProtoMessage()    {}
func (*RespSignupUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{11}
}

func (m *RespSignupUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSignupUser.Unmarshal(m, b)
}
func (m *RespSignupUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSignupUser.Marshal(b, m, deterministic)
}
func (m *RespSignupUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSignupUser.Merge(m, src)
}
func (m *RespSignupUser) XXX_Size() int {
	return xxx_messageInfo_RespSignupUser.Size(m)
}
func (m *RespSignupUser) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSignupUser.DiscardUnknown(m)
}

var xxx_messageInfo_RespSignupUser proto.InternalMessageInfo

func (m *RespSignupUser) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSignupUser) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type ReqSigninUser struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Encpwd               string   `protobuf:"bytes,2,opt,name=encpwd,proto3" json:"encpwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSigninUser) Reset()         { *m = ReqSigninUser{} }
func (m *ReqSigninUser) String() string { return proto.CompactTextString(m) }
func (*ReqSigninUser) ProtoMessage()    {}
func (*ReqSigninUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{12}
}

func (m *ReqSigninUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSigninUser.Unmarshal(m, b)
}
func (m *ReqSigninUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSigninUser.Marshal(b, m, deterministic)
}
func (m *ReqSigninUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSigninUser.Merge(m, src)
}
func (m *ReqSigninUser) XXX_Size() int {
	return xxx_messageInfo_ReqSigninUser.Size(m)
}
func (m *ReqSigninUser) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSigninUser.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSigninUser proto.InternalMessageInfo

func (m *ReqSigninUser) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqSigninUser) GetEncpwd() string {
	if m != nil {
		return m.Encpwd
	}
	return ""
}

type RespSigninUser struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespSigninUser) Reset()         { *m = RespSigninUser{} }
func (m *RespSigninUser) String() string { return proto.CompactTextString(m) }
func (*RespSigninUser) ProtoMessage()    {}
func (*RespSigninUser) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{13}
}

func (m *RespSigninUser) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespSigninUser.Unmarshal(m, b)
}
func (m *RespSigninUser) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespSigninUser.Marshal(b, m, deterministic)
}
func (m *RespSigninUser) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespSigninUser.Merge(m, src)
}
func (m *RespSigninUser) XXX_Size() int {
	return xxx_messageInfo_RespSigninUser.Size(m)
}
func (m *RespSigninUser) XXX_DiscardUnknown() {
	xxx_messageInfo_RespSigninUser.DiscardUnknown(m)
}

var xxx_messageInfo_RespSigninUser proto.InternalMessageInfo

func (m *RespSigninUser) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespSigninUser) GetMessage() string {
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
	return fileDescriptor_353d095adafb115d, []int{14}
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
	EmailValidated       bool     `protobuf:"varint,6,opt,name=emailValidated,proto3" json:"emailValidated,omitempty"`
	PhoneValidated       bool     `protobuf:"varint,7,opt,name=phoneValidated,proto3" json:"phoneValidated,omitempty"`
	SignupAt             string   `protobuf:"bytes,8,opt,name=signupAt,proto3" json:"signupAt,omitempty"`
	LastActive           string   `protobuf:"bytes,9,opt,name=lastActive,proto3" json:"lastActive,omitempty"`
	Status               int32    `protobuf:"varint,10,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUserInfo) Reset()         { *m = RespUserInfo{} }
func (m *RespUserInfo) String() string { return proto.CompactTextString(m) }
func (*RespUserInfo) ProtoMessage()    {}
func (*RespUserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{15}
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

func (m *RespUserInfo) GetEmailValidated() bool {
	if m != nil {
		return m.EmailValidated
	}
	return false
}

func (m *RespUserInfo) GetPhoneValidated() bool {
	if m != nil {
		return m.PhoneValidated
	}
	return false
}

func (m *RespUserInfo) GetSignupAt() string {
	if m != nil {
		return m.SignupAt
	}
	return ""
}

func (m *RespUserInfo) GetLastActive() string {
	if m != nil {
		return m.LastActive
	}
	return ""
}

func (m *RespUserInfo) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

type ReqToken struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqToken) Reset()         { *m = ReqToken{} }
func (m *ReqToken) String() string { return proto.CompactTextString(m) }
func (*ReqToken) ProtoMessage()    {}
func (*ReqToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{16}
}

func (m *ReqToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqToken.Unmarshal(m, b)
}
func (m *ReqToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqToken.Marshal(b, m, deterministic)
}
func (m *ReqToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqToken.Merge(m, src)
}
func (m *ReqToken) XXX_Size() int {
	return xxx_messageInfo_ReqToken.Size(m)
}
func (m *ReqToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqToken.DiscardUnknown(m)
}

var xxx_messageInfo_ReqToken proto.InternalMessageInfo

func (m *ReqToken) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type RespToken struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespToken) Reset()         { *m = RespToken{} }
func (m *RespToken) String() string { return proto.CompactTextString(m) }
func (*RespToken) ProtoMessage()    {}
func (*RespToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{17}
}

func (m *RespToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespToken.Unmarshal(m, b)
}
func (m *RespToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespToken.Marshal(b, m, deterministic)
}
func (m *RespToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespToken.Merge(m, src)
}
func (m *RespToken) XXX_Size() int {
	return xxx_messageInfo_RespToken.Size(m)
}
func (m *RespToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RespToken.DiscardUnknown(m)
}

var xxx_messageInfo_RespToken proto.InternalMessageInfo

func (m *RespToken) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespToken) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *RespToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type ReqUpdateToken struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqUpdateToken) Reset()         { *m = ReqUpdateToken{} }
func (m *ReqUpdateToken) String() string { return proto.CompactTextString(m) }
func (*ReqUpdateToken) ProtoMessage()    {}
func (*ReqUpdateToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{18}
}

func (m *ReqUpdateToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqUpdateToken.Unmarshal(m, b)
}
func (m *ReqUpdateToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqUpdateToken.Marshal(b, m, deterministic)
}
func (m *ReqUpdateToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqUpdateToken.Merge(m, src)
}
func (m *ReqUpdateToken) XXX_Size() int {
	return xxx_messageInfo_ReqUpdateToken.Size(m)
}
func (m *ReqUpdateToken) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqUpdateToken.DiscardUnknown(m)
}

var xxx_messageInfo_ReqUpdateToken proto.InternalMessageInfo

func (m *ReqUpdateToken) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *ReqUpdateToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type RespUpdateToken struct {
	Code                 int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RespUpdateToken) Reset()         { *m = RespUpdateToken{} }
func (m *RespUpdateToken) String() string { return proto.CompactTextString(m) }
func (*RespUpdateToken) ProtoMessage()    {}
func (*RespUpdateToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_353d095adafb115d, []int{19}
}

func (m *RespUpdateToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RespUpdateToken.Unmarshal(m, b)
}
func (m *RespUpdateToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RespUpdateToken.Marshal(b, m, deterministic)
}
func (m *RespUpdateToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RespUpdateToken.Merge(m, src)
}
func (m *RespUpdateToken) XXX_Size() int {
	return xxx_messageInfo_RespUpdateToken.Size(m)
}
func (m *RespUpdateToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RespUpdateToken.DiscardUnknown(m)
}

var xxx_messageInfo_RespUpdateToken proto.InternalMessageInfo

func (m *RespUpdateToken) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *RespUpdateToken) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*ReqUploadFile)(nil), "proto.ReqUploadFile")
	proto.RegisterType((*RespUploadFile)(nil), "proto.RespUploadFile")
	proto.RegisterType((*ReqFileMeta)(nil), "proto.ReqFileMeta")
	proto.RegisterType((*RespFileMeta)(nil), "proto.RespFileMeta")
	proto.RegisterType((*ReqUpdateFileLocation)(nil), "proto.ReqUpdateFileLocation")
	proto.RegisterType((*RespUpdateFileLocation)(nil), "proto.RespUpdateFileLocation")
	proto.RegisterType((*ReqUploadUserFile)(nil), "proto.ReqUploadUserFile")
	proto.RegisterType((*RespUploadUserFile)(nil), "proto.RespUploadUserFile")
	proto.RegisterType((*ReqQueryUserFile)(nil), "proto.ReqQueryUserFile")
	proto.RegisterType((*RespQueryUserFile)(nil), "proto.RespQueryUserFile")
	proto.RegisterType((*ReqSignupUser)(nil), "proto.ReqSignupUser")
	proto.RegisterType((*RespSignupUser)(nil), "proto.RespSignupUser")
	proto.RegisterType((*ReqSigninUser)(nil), "proto.ReqSigninUser")
	proto.RegisterType((*RespSigninUser)(nil), "proto.RespSigninUser")
	proto.RegisterType((*ReqUserInfo)(nil), "proto.ReqUserInfo")
	proto.RegisterType((*RespUserInfo)(nil), "proto.RespUserInfo")
	proto.RegisterType((*ReqToken)(nil), "proto.ReqToken")
	proto.RegisterType((*RespToken)(nil), "proto.RespToken")
	proto.RegisterType((*ReqUpdateToken)(nil), "proto.ReqUpdateToken")
	proto.RegisterType((*RespUpdateToken)(nil), "proto.RespUpdateToken")
}

func init() { proto.RegisterFile("dbproxy.proto", fileDescriptor_353d095adafb115d) }

var fileDescriptor_353d095adafb115d = []byte{
	// 723 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x75, 0xfe, 0xb8, 0x4d, 0xa6, 0x6d, 0xda, 0xdf, 0xfe, 0xda, 0x62, 0x2c, 0x40, 0xd5, 0x1e,
	0xaa, 0x22, 0xa1, 0x4a, 0x80, 0x38, 0x21, 0x15, 0x35, 0x54, 0x45, 0x48, 0x40, 0xc1, 0xa1, 0xdc,
	0xdd, 0x78, 0xdb, 0x5a, 0x38, 0xb6, 0x63, 0x3b, 0x85, 0x70, 0xe1, 0xc0, 0x81, 0xaf, 0xc2, 0x81,
	0x0f, 0x89, 0xf6, 0x9f, 0x77, 0xd7, 0x41, 0x26, 0xee, 0x29, 0x99, 0x37, 0xfb, 0x66, 0x9f, 0xdf,
	0xee, 0xcc, 0xc2, 0x46, 0x70, 0x91, 0x66, 0xc9, 0xd7, 0xf9, 0x61, 0x9a, 0x25, 0x45, 0x82, 0x6c,
	0xf6, 0x83, 0xbf, 0xc3, 0x86, 0x47, 0xa6, 0xe7, 0x69, 0x94, 0xf8, 0xc1, 0x69, 0x18, 0x11, 0xe4,
	0x42, 0xef, 0x32, 0x8c, 0xc8, 0xe8, 0xda, 0x7f, 0xec, 0xb4, 0xf6, 0x5a, 0x07, 0x7d, 0xaf, 0x8c,
	0x65, 0xee, 0x9d, 0x3f, 0x21, 0x4e, 0x5b, 0xe5, 0x68, 0x5c, 0xf2, 0xc2, 0x6f, 0xc4, 0xe9, 0xec,
	0xb5, 0x0e, 0x3a, 0x5e, 0x19, 0xcb, 0xdc, 0x71, 0x10, 0x64, 0x4e, 0x57, 0xf1, 0x68, 0x8c, 0x8f,
	0x60, 0xe0, 0x91, 0x3c, 0xd5, 0x14, 0x20, 0xe8, 0x8e, 0x93, 0x80, 0xb0, 0xdd, 0x6d, 0x8f, 0xfd,
	0x47, 0x0e, 0xac, 0x4e, 0x48, 0x9e, 0xfb, 0x57, 0x72, 0x63, 0x19, 0xe2, 0x87, 0xb0, 0xe6, 0x91,
	0x29, 0x25, 0xbe, 0x25, 0x85, 0x5f, 0x27, 0x1f, 0xff, 0x6e, 0xc1, 0x3a, 0xdd, 0xab, 0x5c, 0xdc,
	0x68, 0x27, 0xa3, 0x74, 0xa7, 0xc6, 0x99, 0x6e, 0x8d, 0x33, 0x76, 0x8d, 0x33, 0x2b, 0x15, 0x67,
	0xce, 0x60, 0x87, 0x1d, 0x4d, 0xe0, 0x17, 0x84, 0x4a, 0x7e, 0x93, 0x8c, 0xfd, 0x22, 0x4c, 0xe2,
	0x65, 0x8e, 0x88, 0x15, 0x6c, 0x57, 0x0a, 0x9e, 0xc2, 0x2e, 0xb7, 0x7a, 0xa1, 0x62, 0x33, 0xcb,
	0x7f, 0xb4, 0xe0, 0xbf, 0xf2, 0xd2, 0x9c, 0xe7, 0x24, 0x93, 0x17, 0x67, 0x96, 0x93, 0x2c, 0xa6,
	0x16, 0x08, 0x55, 0x32, 0x36, 0x14, 0xb7, 0x6b, 0xac, 0xeb, 0xd4, 0x58, 0xd7, 0x35, 0xad, 0xc3,
	0x43, 0x40, 0xea, 0xe2, 0x94, 0x2a, 0x9a, 0x7d, 0xc9, 0x09, 0x6c, 0x79, 0x64, 0xfa, 0x61, 0x46,
	0xb2, 0xf9, 0x52, 0xdf, 0xb1, 0x0d, 0x76, 0x14, 0x4e, 0xc2, 0x82, 0xd5, 0xb1, 0x3d, 0x1e, 0x60,
	0x42, 0xed, 0xc8, 0x53, 0xb3, 0x4c, 0xb3, 0xbb, 0x85, 0x61, 0x7d, 0x26, 0x98, 0xc3, 0x79, 0xc1,
	0x8d, 0x58, 0xf7, 0x0c, 0x0c, 0xbf, 0x64, 0xad, 0x3a, 0x0a, 0xaf, 0xe2, 0x59, 0x4a, 0xb7, 0xa9,
	0x55, 0xba, 0x0b, 0x2b, 0xa9, 0x9f, 0xe7, 0x5f, 0x02, 0xb1, 0x93, 0x88, 0x64, 0xbb, 0x69, 0x55,
	0x9a, 0x39, 0xa6, 0x44, 0x84, 0xf1, 0x32, 0x22, 0x48, 0x3c, 0x4e, 0x95, 0x08, 0x1e, 0xe9, 0x22,
	0x44, 0x95, 0xdb, 0xf4, 0x3c, 0x25, 0xbe, 0x8e, 0x2f, 0x93, 0x3a, 0x09, 0xf8, 0x57, 0x9b, 0xf7,
	0x7c, 0xb9, 0xb8, 0x71, 0xcf, 0x97, 0xa5, 0x3b, 0x8b, 0x97, 0x81, 0x4c, 0xfc, 0x30, 0x12, 0x0d,
	0xcf, 0x03, 0x8a, 0xa6, 0xd7, 0x49, 0xcc, 0x5b, 0xbd, 0xef, 0xf1, 0x00, 0xed, 0xc3, 0x80, 0xa5,
	0x3f, 0xf9, 0x51, 0x48, 0xbb, 0x2f, 0x60, 0xdd, 0xde, 0xf3, 0x2a, 0x28, 0x5d, 0xc7, 0x08, 0x6a,
	0xdd, 0x2a, 0x5f, 0x67, 0xa2, 0x54, 0x57, 0xce, 0x8e, 0xf0, 0xb8, 0x70, 0x7a, 0x5c, 0x97, 0x8c,
	0xd1, 0x03, 0x80, 0xc8, 0xcf, 0x8b, 0xe3, 0x71, 0x11, 0xde, 0x10, 0xa7, 0xcf, 0xb2, 0x1a, 0x42,
	0x4f, 0x25, 0x2f, 0xfc, 0x62, 0x96, 0x3b, 0xc0, 0x3c, 0x10, 0x11, 0xde, 0x87, 0x9e, 0x47, 0xa6,
	0x1f, 0x93, 0xcf, 0x24, 0xae, 0xb5, 0xf4, 0x0c, 0xfa, 0xd4, 0x51, 0xbe, 0xb0, 0x99, 0x9d, 0xdb,
	0x60, 0x17, 0x94, 0x26, 0xbc, 0xe4, 0x01, 0x1e, 0xd2, 0xeb, 0x20, 0x06, 0xdd, 0x3f, 0xb7, 0x57,
	0x35, 0xda, 0x7a, 0x8d, 0x17, 0xb0, 0xa9, 0x66, 0xdb, 0x2d, 0xa4, 0x3d, 0xf9, 0x69, 0xc3, 0xe0,
	0x64, 0xf8, 0x9e, 0xbe, 0x90, 0x23, 0x92, 0xdd, 0x84, 0x63, 0x82, 0x9e, 0x03, 0x68, 0xcf, 0xd2,
	0x36, 0x7f, 0x38, 0x0f, 0x8d, 0xe7, 0xd2, 0xdd, 0x29, 0x51, 0xfd, 0x0d, 0xc3, 0x16, 0x7a, 0x06,
	0x3d, 0xf5, 0xce, 0x28, 0xaa, 0xc4, 0xdc, 0xff, 0x35, 0xa2, 0x04, 0xb1, 0x85, 0x46, 0x80, 0xfe,
	0x32, 0x9f, 0xef, 0xe9, 0x7b, 0x57, 0xb3, 0xee, 0x7d, 0x43, 0x43, 0x35, 0x8d, 0x2d, 0xf4, 0x0a,
	0x06, 0x95, 0x31, 0xe9, 0x54, 0x3f, 0x46, 0x66, 0xdc, 0xbb, 0x0b, 0x1f, 0x24, 0x53, 0xd8, 0x42,
	0x27, 0xb0, 0x61, 0x4e, 0xb9, 0x3b, 0xaa, 0x8e, 0x91, 0x70, 0x1d, 0xad, 0x8c, 0x91, 0xc1, 0x16,
	0xf5, 0x55, 0x9b, 0x3f, 0x9a, 0xaf, 0x0a, 0x35, 0x7c, 0x55, 0xb0, 0x22, 0x8b, 0xb9, 0x51, 0x21,
	0x73, 0x74, 0x81, 0xcc, 0x61, 0x7e, 0x28, 0x6a, 0x10, 0x68, 0x16, 0x08, 0xcc, 0x38, 0x14, 0x09,
	0x62, 0x0b, 0x3d, 0x02, 0x9b, 0x5f, 0xa9, 0x4d, 0xc5, 0x61, 0x80, 0xbb, 0xa5, 0x11, 0x18, 0x82,
	0x2d, 0x74, 0x04, 0x6b, 0xfa, 0x35, 0xdc, 0xa9, 0x9e, 0x1d, 0x67, 0xee, 0x2e, 0x1c, 0x9a, 0xe0,
	0x5f, 0xac, 0xb0, 0xc4, 0xd3, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x6f, 0x06, 0xab, 0xb1,
	0x09, 0x00, 0x00,
}