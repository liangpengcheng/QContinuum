// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	C2SLogin
	S2CLogin
	C2STokenLogin
	S2CTokenLogin
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 登录结果
type LoginResult int32

const (
	LoginResult_Success       LoginResult = 0
	LoginResult_UserNameError LoginResult = -1
	LoginResult_PasswordError LoginResult = -2
	LoginResult_TokenError    LoginResult = -3
	LoginResult_UnknownError  LoginResult = -4
)

var LoginResult_name = map[int32]string{
	0:  "Success",
	-1: "UserNameError",
	-2: "PasswordError",
	-3: "TokenError",
	-4: "UnknownError",
}
var LoginResult_value = map[string]int32{
	"Success":       0,
	"UserNameError": -1,
	"PasswordError": -2,
	"TokenError":    -3,
	"UnknownError":  -4,
}

func (x LoginResult) String() string {
	return proto.EnumName(LoginResult_name, int32(x))
}
func (LoginResult) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type C2SLogin_MsgID int32

const (
	C2SLogin_Zero C2SLogin_MsgID = 0
	C2SLogin_ID   C2SLogin_MsgID = 1
)

var C2SLogin_MsgID_name = map[int32]string{
	0: "Zero",
	1: "ID",
}
var C2SLogin_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   1,
}

func (x C2SLogin_MsgID) String() string {
	return proto.EnumName(C2SLogin_MsgID_name, int32(x))
}
func (C2SLogin_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

type S2CLogin_MsgID int32

const (
	S2CLogin_Zero S2CLogin_MsgID = 0
	S2CLogin_ID   S2CLogin_MsgID = 2
)

var S2CLogin_MsgID_name = map[int32]string{
	0: "Zero",
	2: "ID",
}
var S2CLogin_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   2,
}

func (x S2CLogin_MsgID) String() string {
	return proto.EnumName(S2CLogin_MsgID_name, int32(x))
}
func (S2CLogin_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1, 0} }

type C2STokenLogin_MsgID int32

const (
	C2STokenLogin_Zero C2STokenLogin_MsgID = 0
	C2STokenLogin_ID   C2STokenLogin_MsgID = 3
)

var C2STokenLogin_MsgID_name = map[int32]string{
	0: "Zero",
	3: "ID",
}
var C2STokenLogin_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   3,
}

func (x C2STokenLogin_MsgID) String() string {
	return proto.EnumName(C2STokenLogin_MsgID_name, int32(x))
}
func (C2STokenLogin_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type S2CTokenLogin_MsgID int32

const (
	S2CTokenLogin_Zero S2CTokenLogin_MsgID = 0
	S2CTokenLogin_ID   S2CTokenLogin_MsgID = 4
)

var S2CTokenLogin_MsgID_name = map[int32]string{
	0: "Zero",
	4: "ID",
}
var S2CTokenLogin_MsgID_value = map[string]int32{
	"Zero": 0,
	"ID":   4,
}

func (x S2CTokenLogin_MsgID) String() string {
	return proto.EnumName(S2CTokenLogin_MsgID_name, int32(x))
}
func (S2CTokenLogin_MsgID) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

// 登陆消息，客户端发送给登录服务器
type C2SLogin struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
}

func (m *C2SLogin) Reset()                    { *m = C2SLogin{} }
func (m *C2SLogin) String() string            { return proto.CompactTextString(m) }
func (*C2SLogin) ProtoMessage()               {}
func (*C2SLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *C2SLogin) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *C2SLogin) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

// 登录结果消息，服务器发送给客户端
type S2CLogin struct {
	Result           LoginResult `protobuf:"varint,1,opt,name=result,enum=LoginResult" json:"result,omitempty"`
	Userid           int64       `protobuf:"varint,2,opt,name=userid" json:"userid,omitempty"`
	Token            string      `protobuf:"bytes,3,opt,name=token" json:"token,omitempty"`
	ConnectionServer string      `protobuf:"bytes,4,opt,name=connectionServer" json:"connectionServer,omitempty"`
}

func (m *S2CLogin) Reset()                    { *m = S2CLogin{} }
func (m *S2CLogin) String() string            { return proto.CompactTextString(m) }
func (*S2CLogin) ProtoMessage()               {}
func (*S2CLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *S2CLogin) GetResult() LoginResult {
	if m != nil {
		return m.Result
	}
	return LoginResult_Success
}

func (m *S2CLogin) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *S2CLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *S2CLogin) GetConnectionServer() string {
	if m != nil {
		return m.ConnectionServer
	}
	return ""
}

// token登录
type C2STokenLogin struct {
	Userid int64  `protobuf:"varint,1,opt,name=userid" json:"userid,omitempty"`
	Token  string `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
}

func (m *C2STokenLogin) Reset()                    { *m = C2STokenLogin{} }
func (m *C2STokenLogin) String() string            { return proto.CompactTextString(m) }
func (*C2STokenLogin) ProtoMessage()               {}
func (*C2STokenLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *C2STokenLogin) GetUserid() int64 {
	if m != nil {
		return m.Userid
	}
	return 0
}

func (m *C2STokenLogin) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// token登陆结果
type S2CTokenLogin struct {
	Resut LoginResult `protobuf:"varint,1,opt,name=resut,enum=LoginResult" json:"resut,omitempty"`
}

func (m *S2CTokenLogin) Reset()                    { *m = S2CTokenLogin{} }
func (m *S2CTokenLogin) String() string            { return proto.CompactTextString(m) }
func (*S2CTokenLogin) ProtoMessage()               {}
func (*S2CTokenLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *S2CTokenLogin) GetResut() LoginResult {
	if m != nil {
		return m.Resut
	}
	return LoginResult_Success
}

func init() {
	proto.RegisterType((*C2SLogin)(nil), "C2SLogin")
	proto.RegisterType((*S2CLogin)(nil), "S2CLogin")
	proto.RegisterType((*C2STokenLogin)(nil), "C2STokenLogin")
	proto.RegisterType((*S2CTokenLogin)(nil), "S2CTokenLogin")
	proto.RegisterEnum("LoginResult", LoginResult_name, LoginResult_value)
	proto.RegisterEnum("C2SLogin_MsgID", C2SLogin_MsgID_name, C2SLogin_MsgID_value)
	proto.RegisterEnum("S2CLogin_MsgID", S2CLogin_MsgID_name, S2CLogin_MsgID_value)
	proto.RegisterEnum("C2STokenLogin_MsgID", C2STokenLogin_MsgID_name, C2STokenLogin_MsgID_value)
	proto.RegisterEnum("S2CTokenLogin_MsgID", S2CTokenLogin_MsgID_name, S2CTokenLogin_MsgID_value)
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x92, 0xcd, 0x4e, 0xfa, 0x40,
	0x14, 0xc5, 0x99, 0x02, 0xfd, 0xf7, 0x7f, 0xa1, 0xa6, 0x99, 0x18, 0x2d, 0xac, 0x4c, 0xe3, 0xc2,
	0xb0, 0x60, 0x51, 0x1f, 0xa1, 0xb8, 0x20, 0x51, 0x62, 0x3a, 0x92, 0x18, 0x13, 0x17, 0x58, 0x6f,
	0x48, 0x83, 0xcc, 0x90, 0x3b, 0x05, 0xde, 0xc0, 0x37, 0xf1, 0x2d, 0xfd, 0x4a, 0x67, 0xca, 0x47,
	0x22, 0x74, 0x77, 0xee, 0xaf, 0x73, 0xce, 0x99, 0x0f, 0xf0, 0x35, 0xd2, 0x2a, 0xcf, 0xb0, 0xbf,
	0x20, 0x55, 0xa8, 0xe8, 0x19, 0xbc, 0x24, 0x16, 0xb7, 0x6a, 0x9a, 0x4b, 0xde, 0x05, 0x6f, 0xa9,
	0x91, 0xe4, 0x64, 0x8e, 0x21, 0xbb, 0x60, 0x57, 0xff, 0xd3, 0xad, 0x2e, 0xd9, 0x62, 0xa2, 0xf5,
	0x5a, 0xd1, 0x6b, 0xe8, 0x58, 0xb6, 0xd1, 0x51, 0x07, 0x9a, 0x77, 0x7a, 0x3a, 0x1c, 0x70, 0x0f,
	0x1a, 0x4f, 0x48, 0x2a, 0xa8, 0x71, 0x17, 0x9c, 0xe1, 0x20, 0x60, 0xd1, 0x07, 0x03, 0x4f, 0xc4,
	0x89, 0xf5, 0xbf, 0x04, 0x97, 0x50, 0x2f, 0xdf, 0x0a, 0xe3, 0x7e, 0x12, 0xb7, 0xfb, 0x66, 0x9e,
	0x9a, 0x59, 0x5a, 0x31, 0x7e, 0x06, 0x6e, 0x99, 0x9a, 0xdb, 0x9c, 0x7a, 0x5a, 0x29, 0x7e, 0x0a,
	0xcd, 0x42, 0xcd, 0x50, 0x86, 0x75, 0x13, 0x6f, 0x05, 0xef, 0x41, 0x90, 0x29, 0x29, 0x31, 0x2b,
	0x72, 0x25, 0x05, 0xd2, 0x0a, 0x29, 0x6c, 0x98, 0x1f, 0xfe, 0xcc, 0x8f, 0xf7, 0x74, 0xa2, 0x47,
	0xf0, 0x93, 0x58, 0x3c, 0x94, 0x96, 0xb6, 0xeb, 0xae, 0x05, 0x3b, 0xdc, 0xc2, 0xd9, 0x6b, 0x71,
	0xdc, 0xb9, 0x1e, 0x8d, 0xc0, 0x17, 0x71, 0xb2, 0xe7, 0x1c, 0x41, 0xb3, 0xdc, 0xe9, 0xe1, 0x43,
	0xb0, 0xe8, 0xb8, 0x5f, 0xa3, 0xf7, 0xce, 0xa0, 0xb5, 0xb7, 0x82, 0xb7, 0xe0, 0x9f, 0x58, 0x66,
	0x19, 0x6a, 0x1d, 0xd4, 0x78, 0x17, 0xfc, 0xb1, 0x46, 0x1a, 0x4d, 0xe6, 0x78, 0x43, 0xa4, 0x28,
	0xf8, 0xd9, 0x7c, 0xac, 0x64, 0xf7, 0xd5, 0x8d, 0x59, 0xf6, 0xbd, 0x63, 0xe7, 0x00, 0xa6, 0xa1,
	0x05, 0x5f, 0x3b, 0xd0, 0x81, 0xf6, 0x58, 0xce, 0xa4, 0x5a, 0x57, 0xe8, 0x73, 0x8b, 0x5e, 0x5c,
	0xf3, 0x80, 0xae, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x42, 0x39, 0x8c, 0x71, 0x51, 0x02, 0x00,
	0x00,
}