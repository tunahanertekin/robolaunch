// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package api

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

// LaunchClient is the client API for Launch service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LaunchClient interface {
	ListLaunch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Launch_ListLaunchClient, error)
	CreateLaunch(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*LaunchState, error)
	OperateLaunch(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*LaunchState, error)
}

type launchClient struct {
	cc grpc.ClientConnInterface
}

func NewLaunchClient(cc grpc.ClientConnInterface) LaunchClient {
	return &launchClient{cc}
}

func (c *launchClient) ListLaunch(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Launch_ListLaunchClient, error) {
	stream, err := c.cc.NewStream(ctx, &Launch_ServiceDesc.Streams[0], "/auth.Launch/ListLaunch", opts...)
	if err != nil {
		return nil, err
	}
	x := &launchListLaunchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Launch_ListLaunchClient interface {
	Recv() (*LaunchList, error)
	grpc.ClientStream
}

type launchListLaunchClient struct {
	grpc.ClientStream
}

func (x *launchListLaunchClient) Recv() (*LaunchList, error) {
	m := new(LaunchList)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *launchClient) CreateLaunch(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*LaunchState, error) {
	out := new(LaunchState)
	err := c.cc.Invoke(ctx, "/auth.Launch/CreateLaunch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *launchClient) OperateLaunch(ctx context.Context, in *OperateRequest, opts ...grpc.CallOption) (*LaunchState, error) {
	out := new(LaunchState)
	err := c.cc.Invoke(ctx, "/auth.Launch/OperateLaunch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LaunchServer is the server API for Launch service.
// All implementations must embed UnimplementedLaunchServer
// for forward compatibility
type LaunchServer interface {
	ListLaunch(*Empty, Launch_ListLaunchServer) error
	CreateLaunch(context.Context, *CreateRequest) (*LaunchState, error)
	OperateLaunch(context.Context, *OperateRequest) (*LaunchState, error)
	mustEmbedUnimplementedLaunchServer()
}

// UnimplementedLaunchServer must be embedded to have forward compatible implementations.
type UnimplementedLaunchServer struct {
}

func (UnimplementedLaunchServer) ListLaunch(*Empty, Launch_ListLaunchServer) error {
	return status.Errorf(codes.Unimplemented, "method ListLaunch not implemented")
}
func (UnimplementedLaunchServer) CreateLaunch(context.Context, *CreateRequest) (*LaunchState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateLaunch not implemented")
}
func (UnimplementedLaunchServer) OperateLaunch(context.Context, *OperateRequest) (*LaunchState, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OperateLaunch not implemented")
}
func (UnimplementedLaunchServer) mustEmbedUnimplementedLaunchServer() {}

// UnsafeLaunchServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LaunchServer will
// result in compilation errors.
type UnsafeLaunchServer interface {
	mustEmbedUnimplementedLaunchServer()
}

func RegisterLaunchServer(s grpc.ServiceRegistrar, srv LaunchServer) {
	s.RegisterService(&Launch_ServiceDesc, srv)
}

func _Launch_ListLaunch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LaunchServer).ListLaunch(m, &launchListLaunchServer{stream})
}

type Launch_ListLaunchServer interface {
	Send(*LaunchList) error
	grpc.ServerStream
}

type launchListLaunchServer struct {
	grpc.ServerStream
}

func (x *launchListLaunchServer) Send(m *LaunchList) error {
	return x.ServerStream.SendMsg(m)
}

func _Launch_CreateLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaunchServer).CreateLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Launch/CreateLaunch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaunchServer).CreateLaunch(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Launch_OperateLaunch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OperateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LaunchServer).OperateLaunch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Launch/OperateLaunch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LaunchServer).OperateLaunch(ctx, req.(*OperateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Launch_ServiceDesc is the grpc.ServiceDesc for Launch service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Launch_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Launch",
	HandlerType: (*LaunchServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateLaunch",
			Handler:    _Launch_CreateLaunch_Handler,
		},
		{
			MethodName: "OperateLaunch",
			Handler:    _Launch_OperateLaunch_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListLaunch",
			Handler:       _Launch_ListLaunch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "launch.proto",
}