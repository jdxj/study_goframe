// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: user/v1/user.proto

package v1

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

const (
	H_HH_FullMethodName = "/H/HH"
)

// HClient is the client API for H service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HClient interface {
	HH(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error)
}

type hClient struct {
	cc grpc.ClientConnInterface
}

func NewHClient(cc grpc.ClientConnInterface) HClient {
	return &hClient{cc}
}

func (c *hClient) HH(ctx context.Context, in *Hello, opts ...grpc.CallOption) (*Hello, error) {
	out := new(Hello)
	err := c.cc.Invoke(ctx, H_HH_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HServer is the server API for H service.
// All implementations must embed UnimplementedHServer
// for forward compatibility
type HServer interface {
	HH(context.Context, *Hello) (*Hello, error)
	mustEmbedUnimplementedHServer()
}

// UnimplementedHServer must be embedded to have forward compatible implementations.
type UnimplementedHServer struct {
}

func (UnimplementedHServer) HH(context.Context, *Hello) (*Hello, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HH not implemented")
}
func (UnimplementedHServer) mustEmbedUnimplementedHServer() {}

// UnsafeHServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HServer will
// result in compilation errors.
type UnsafeHServer interface {
	mustEmbedUnimplementedHServer()
}

func RegisterHServer(s grpc.ServiceRegistrar, srv HServer) {
	s.RegisterService(&H_ServiceDesc, srv)
}

func _H_HH_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Hello)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HServer).HH(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: H_HH_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HServer).HH(ctx, req.(*Hello))
	}
	return interceptor(ctx, in, info, handler)
}

// H_ServiceDesc is the grpc.ServiceDesc for H service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var H_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "H",
	HandlerType: (*HServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HH",
			Handler:    _H_HH_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user.proto",
}
