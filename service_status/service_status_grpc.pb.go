// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: service_status.proto

package service_status

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

// ServiceStatusClient is the client API for ServiceStatus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceStatusClient interface {
	GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Status, error)
}

type serviceStatusClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceStatusClient(cc grpc.ClientConnInterface) ServiceStatusClient {
	return &serviceStatusClient{cc}
}

func (c *serviceStatusClient) GetStatus(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/service_status.ServiceStatus/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceStatusServer is the server API for ServiceStatus service.
// All implementations must embed UnimplementedServiceStatusServer
// for forward compatibility
type ServiceStatusServer interface {
	GetStatus(context.Context, *emptypb.Empty) (*Status, error)
	mustEmbedUnimplementedServiceStatusServer()
}

// UnimplementedServiceStatusServer must be embedded to have forward compatible implementations.
type UnimplementedServiceStatusServer struct {
}

func (UnimplementedServiceStatusServer) GetStatus(context.Context, *emptypb.Empty) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedServiceStatusServer) mustEmbedUnimplementedServiceStatusServer() {}

// UnsafeServiceStatusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceStatusServer will
// result in compilation errors.
type UnsafeServiceStatusServer interface {
	mustEmbedUnimplementedServiceStatusServer()
}

func RegisterServiceStatusServer(s grpc.ServiceRegistrar, srv ServiceStatusServer) {
	s.RegisterService(&ServiceStatus_ServiceDesc, srv)
}

func _ServiceStatus_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceStatusServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service_status.ServiceStatus/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceStatusServer).GetStatus(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ServiceStatus_ServiceDesc is the grpc.ServiceDesc for ServiceStatus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServiceStatus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service_status.ServiceStatus",
	HandlerType: (*ServiceStatusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _ServiceStatus_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_status.proto",
}
