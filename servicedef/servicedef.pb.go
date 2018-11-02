// Code generated by protoc-gen-go. DO NOT EDIT.
// source: servicedef.proto

package servicedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type SignInRequest struct {
	Username             string   `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SignInRequest) Reset()         { *m = SignInRequest{} }
func (m *SignInRequest) String() string { return proto.CompactTextString(m) }
func (*SignInRequest) ProtoMessage()    {}
func (*SignInRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_servicedef_ef9e33f5c5ec687e, []int{0}
}
func (m *SignInRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SignInRequest.Unmarshal(m, b)
}
func (m *SignInRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SignInRequest.Marshal(b, m, deterministic)
}
func (dst *SignInRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignInRequest.Merge(dst, src)
}
func (m *SignInRequest) XXX_Size() int {
	return xxx_messageInfo_SignInRequest.Size(m)
}
func (m *SignInRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SignInRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SignInRequest proto.InternalMessageInfo

func (m *SignInRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *SignInRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*SignInRequest)(nil), "servicedef.SignInRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthApiClient is the client API for AuthApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthApiClient interface {
	// signin, establishing a session with cookie
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// signout, deleting any existing session/cookie
	SignOut(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error)
}

type authApiClient struct {
	cc *grpc.ClientConn
}

func NewAuthApiClient(cc *grpc.ClientConn) AuthApiClient {
	return &authApiClient{cc}
}

func (c *authApiClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/servicedef.AuthApi/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authApiClient) SignOut(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/servicedef.AuthApi/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthApiServer is the server API for AuthApi service.
type AuthApiServer interface {
	// signin, establishing a session with cookie
	SignIn(context.Context, *SignInRequest) (*empty.Empty, error)
	// signout, deleting any existing session/cookie
	SignOut(context.Context, *empty.Empty) (*empty.Empty, error)
}

func RegisterAuthApiServer(s *grpc.Server, srv AuthApiServer) {
	s.RegisterService(&_AuthApi_serviceDesc, srv)
}

func _AuthApi_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthApiServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicedef.AuthApi/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthApiServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthApi_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthApiServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/servicedef.AuthApi/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthApiServer).SignOut(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "servicedef.AuthApi",
	HandlerType: (*AuthApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIn",
			Handler:    _AuthApi_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthApi_SignOut_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "servicedef.proto",
}

func init() { proto.RegisterFile("servicedef.proto", fileDescriptor_servicedef_ef9e33f5c5ec687e) }

var fileDescriptor_servicedef_ef9e33f5c5ec687e = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x4d, 0x49, 0x4d, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x48, 0xc9, 0xa4, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x27, 0x16, 0x64, 0xea, 0x27, 0xe6, 0xe5,
	0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x54, 0x4a, 0x49, 0x43, 0x65, 0xc1, 0xbc,
	0xa4, 0xd2, 0x34, 0xfd, 0xd4, 0xdc, 0x82, 0x92, 0x4a, 0x88, 0xa4, 0x92, 0x3b, 0x17, 0x6f, 0x70,
	0x66, 0x7a, 0x9e, 0x67, 0x5e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x14, 0x17, 0x47,
	0x69, 0x71, 0x6a, 0x51, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c,
	0x0f, 0x92, 0x2b, 0x48, 0x2c, 0x2e, 0x2e, 0xcf, 0x2f, 0x4a, 0x91, 0x60, 0x82, 0xc8, 0xc1, 0xf8,
	0x46, 0xfb, 0x18, 0xb9, 0xd8, 0x1d, 0x4b, 0x4b, 0x32, 0x1c, 0x0b, 0x32, 0x85, 0xa2, 0xb8, 0xd8,
	0x20, 0x86, 0x0a, 0x49, 0xea, 0x21, 0x39, 0x1c, 0xc5, 0x22, 0x29, 0x31, 0x3d, 0x88, 0xbb, 0xf4,
	0x60, 0xee, 0xd2, 0x73, 0x05, 0xb9, 0x4b, 0x49, 0xb6, 0xe9, 0xf2, 0x93, 0xc9, 0x4c, 0xe2, 0x4a,
	0x42, 0xfa, 0x65, 0x86, 0xfa, 0x89, 0xa5, 0x25, 0x19, 0x20, 0x5f, 0x15, 0x67, 0xa6, 0xe7, 0x65,
	0xe6, 0x59, 0x31, 0x6a, 0x09, 0x85, 0x73, 0xb1, 0x83, 0xcc, 0xf1, 0x2f, 0x2d, 0x11, 0xc2, 0x61,
	0x02, 0x4e, 0x93, 0x65, 0xc0, 0x26, 0x8b, 0x29, 0x89, 0x80, 0x4d, 0x06, 0x85, 0x55, 0x69, 0x49,
	0x06, 0xd8, 0xe8, 0xfc, 0xd2, 0x92, 0x24, 0x36, 0xb0, 0x6a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0xcb, 0xac, 0xc0, 0x6b, 0x01, 0x00, 0x00,
}
