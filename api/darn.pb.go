// Code generated by protoc-gen-go. DO NOT EDIT.
// source: darn.proto

/*
Package darn is a generated protocol buffer package.

It is generated from these files:
	darn.proto

It has these top-level messages:
	Data
	Fix
*/
package darn

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
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

type Data struct {
	Plugin   string `protobuf:"bytes,1,opt,name=plugin" json:"plugin,omitempty"`
	Trigger  string `protobuf:"bytes,2,opt,name=trigger" json:"trigger,omitempty"`
	Error    string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
	Platform string `protobuf:"bytes,4,opt,name=platform" json:"platform,omitempty"`
}

func (m *Data) Reset()                    { *m = Data{} }
func (m *Data) String() string            { return proto.CompactTextString(m) }
func (*Data) ProtoMessage()               {}
func (*Data) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Data) GetPlugin() string {
	if m != nil {
		return m.Plugin
	}
	return ""
}

func (m *Data) GetTrigger() string {
	if m != nil {
		return m.Trigger
	}
	return ""
}

func (m *Data) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *Data) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

type Fix struct {
	Fix string `protobuf:"bytes,1,opt,name=fix" json:"fix,omitempty"`
}

func (m *Fix) Reset()                    { *m = Fix{} }
func (m *Fix) String() string            { return proto.CompactTextString(m) }
func (*Fix) ProtoMessage()               {}
func (*Fix) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Fix) GetFix() string {
	if m != nil {
		return m.Fix
	}
	return ""
}

func init() {
	proto.RegisterType((*Data)(nil), "darn.Data")
	proto.RegisterType((*Fix)(nil), "darn.Fix")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ErrorIn service

type ErrorInClient interface {
	Submit(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Fix, error)
}

type errorInClient struct {
	cc *grpc.ClientConn
}

func NewErrorInClient(cc *grpc.ClientConn) ErrorInClient {
	return &errorInClient{cc}
}

func (c *errorInClient) Submit(ctx context.Context, in *Data, opts ...grpc.CallOption) (*Fix, error) {
	out := new(Fix)
	err := grpc.Invoke(ctx, "/darn.ErrorIn/Submit", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ErrorIn service

type ErrorInServer interface {
	Submit(context.Context, *Data) (*Fix, error)
}

func RegisterErrorInServer(s *grpc.Server, srv ErrorInServer) {
	s.RegisterService(&_ErrorIn_serviceDesc, srv)
}

func _ErrorIn_Submit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Data)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ErrorInServer).Submit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/darn.ErrorIn/Submit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ErrorInServer).Submit(ctx, req.(*Data))
	}
	return interceptor(ctx, in, info, handler)
}

var _ErrorIn_serviceDesc = grpc.ServiceDesc{
	ServiceName: "darn.ErrorIn",
	HandlerType: (*ErrorInServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Submit",
			Handler:    _ErrorIn_Submit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "darn.proto",
}

func init() { proto.RegisterFile("darn.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x49, 0x2c, 0xca,
	0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0xa5, 0x64, 0xd2, 0xf3, 0xf3, 0xd3,
	0x73, 0x52, 0xf5, 0x13, 0x0b, 0x32, 0xf5, 0x13, 0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3,
	0xf3, 0x8a, 0x21, 0x6a, 0x94, 0xb2, 0xb8, 0x58, 0x5c, 0x12, 0x4b, 0x12, 0x85, 0xc4, 0xb8, 0xd8,
	0x0a, 0x72, 0x4a, 0xd3, 0x33, 0xf3, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21,
	0x09, 0x2e, 0xf6, 0x92, 0xa2, 0xcc, 0xf4, 0xf4, 0xd4, 0x22, 0x09, 0x26, 0xb0, 0x04, 0x8c, 0x2b,
	0x24, 0xc2, 0xc5, 0x9a, 0x5a, 0x54, 0x94, 0x5f, 0x24, 0xc1, 0x0c, 0x16, 0x87, 0x70, 0x84, 0xa4,
	0xb8, 0x38, 0x0a, 0x72, 0x12, 0x4b, 0xd2, 0xf2, 0x8b, 0x72, 0x25, 0x58, 0xc0, 0x12, 0x70, 0xbe,
	0x92, 0x38, 0x17, 0xb3, 0x5b, 0x66, 0x85, 0x90, 0x00, 0x17, 0x73, 0x5a, 0x66, 0x05, 0xd4, 0x1e,
	0x10, 0xd3, 0xc8, 0x91, 0x8b, 0xdd, 0x15, 0xa4, 0xdb, 0x33, 0x4f, 0xc8, 0x8c, 0x8b, 0x2d, 0xb8,
	0x34, 0x29, 0x37, 0xb3, 0x44, 0x88, 0x4b, 0x0f, 0xec, 0x15, 0x90, 0xeb, 0xa4, 0x38, 0x21, 0x6c,
	0xb7, 0xcc, 0x0a, 0x25, 0xd1, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0xf1, 0x2b, 0x71, 0xe9, 0x97, 0x19,
	0xea, 0x17, 0x83, 0x95, 0x5a, 0x31, 0x6a, 0x25, 0xb1, 0x81, 0xbd, 0x63, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x89, 0x5b, 0x34, 0x68, 0x00, 0x01, 0x00, 0x00,
}
