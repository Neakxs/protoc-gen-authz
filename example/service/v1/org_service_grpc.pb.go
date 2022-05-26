// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: example/service/v1/org_service.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrgServiceClient is the client API for OrgService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrgServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Pong(ctx context.Context, in *PongRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type orgServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrgServiceClient(cc grpc.ClientConnInterface) OrgServiceClient {
	return &orgServiceClient{cc}
}

func (c *orgServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.v1.OrgService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orgServiceClient) Pong(ctx context.Context, in *PongRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/service.v1.OrgService/Pong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrgServiceServer is the server API for OrgService service.
// All implementations must embed UnimplementedOrgServiceServer
// for forward compatibility
type OrgServiceServer interface {
	Ping(context.Context, *PingRequest) (*emptypb.Empty, error)
	Pong(context.Context, *PongRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedOrgServiceServer()
}

// UnimplementedOrgServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrgServiceServer struct {
}

func (UnimplementedOrgServiceServer) Ping(context.Context, *PingRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedOrgServiceServer) Pong(context.Context, *PongRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pong not implemented")
}
func (UnimplementedOrgServiceServer) mustEmbedUnimplementedOrgServiceServer() {}

// UnsafeOrgServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrgServiceServer will
// result in compilation errors.
type UnsafeOrgServiceServer interface {
	mustEmbedUnimplementedOrgServiceServer()
}

func RegisterOrgServiceServer(s grpc.ServiceRegistrar, srv OrgServiceServer) {
	s.RegisterService(&OrgService_ServiceDesc, srv)
}

func _OrgService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.v1.OrgService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrgService_Pong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PongRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrgServiceServer).Pong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.v1.OrgService/Pong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrgServiceServer).Pong(ctx, req.(*PongRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrgService_ServiceDesc is the grpc.ServiceDesc for OrgService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrgService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.v1.OrgService",
	HandlerType: (*OrgServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _OrgService_Ping_Handler,
		},
		{
			MethodName: "Pong",
			Handler:    _OrgService_Pong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "example/service/v1/org_service.proto",
}
