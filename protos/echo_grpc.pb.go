// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.2
// source: echo.proto

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EchoServiceClient is the client API for EchoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EchoServiceClient interface {
	GetUnaryEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

type echoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEchoServiceClient(cc grpc.ClientConnInterface) EchoServiceClient {
	return &echoServiceClient{cc}
}

func (c *echoServiceClient) GetUnaryEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	err := c.cc.Invoke(ctx, "/EchoService/GetUnaryEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EchoServiceServer is the server API for EchoService service.
// All implementations must embed UnimplementedEchoServiceServer
// for forward compatibility
type EchoServiceServer interface {
	GetUnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error)
	mustEmbedUnimplementedEchoServiceServer()
}

// UnimplementedEchoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEchoServiceServer struct {
}

func (UnimplementedEchoServiceServer) GetUnaryEcho(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUnaryEcho not implemented")
}
func (UnimplementedEchoServiceServer) mustEmbedUnimplementedEchoServiceServer() {}

// UnsafeEchoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EchoServiceServer will
// result in compilation errors.
type UnsafeEchoServiceServer interface {
	mustEmbedUnimplementedEchoServiceServer()
}

func RegisterEchoServiceServer(s grpc.ServiceRegistrar, srv EchoServiceServer) {
	s.RegisterService(&EchoService_ServiceDesc, srv)
}

func _EchoService_GetUnaryEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServiceServer).GetUnaryEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EchoService/GetUnaryEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServiceServer).GetUnaryEcho(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EchoService_ServiceDesc is the grpc.ServiceDesc for EchoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EchoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EchoService",
	HandlerType: (*EchoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUnaryEcho",
			Handler:    _EchoService_GetUnaryEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}
