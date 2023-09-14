// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.3
// source: protocol_plugin.proto

package protocol_plugin

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

// ProtocolPluginServiceClient is the client API for ProtocolPluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProtocolPluginServiceClient interface {
	PluginDeviceConfig(ctx context.Context, in *PluginDeviceConfigRequest, opts ...grpc.CallOption) (*PluginDeviceConfigReply, error)
	PluginDeviceConfigList(ctx context.Context, in *PluginDeviceConfigListRequest, opts ...grpc.CallOption) (*PluginDeviceConfigListReply, error)
}

type protocolPluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProtocolPluginServiceClient(cc grpc.ClientConnInterface) ProtocolPluginServiceClient {
	return &protocolPluginServiceClient{cc}
}

func (c *protocolPluginServiceClient) PluginDeviceConfig(ctx context.Context, in *PluginDeviceConfigRequest, opts ...grpc.CallOption) (*PluginDeviceConfigReply, error) {
	out := new(PluginDeviceConfigReply)
	err := c.cc.Invoke(ctx, "/protocol_plugin.ProtocolPluginService/PluginDeviceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *protocolPluginServiceClient) PluginDeviceConfigList(ctx context.Context, in *PluginDeviceConfigListRequest, opts ...grpc.CallOption) (*PluginDeviceConfigListReply, error) {
	out := new(PluginDeviceConfigListReply)
	err := c.cc.Invoke(ctx, "/protocol_plugin.ProtocolPluginService/PluginDeviceConfigList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProtocolPluginServiceServer is the server API for ProtocolPluginService service.
// All implementations must embed UnimplementedProtocolPluginServiceServer
// for forward compatibility
type ProtocolPluginServiceServer interface {
	PluginDeviceConfig(context.Context, *PluginDeviceConfigRequest) (*PluginDeviceConfigReply, error)
	PluginDeviceConfigList(context.Context, *PluginDeviceConfigListRequest) (*PluginDeviceConfigListReply, error)
	mustEmbedUnimplementedProtocolPluginServiceServer()
}

// UnimplementedProtocolPluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProtocolPluginServiceServer struct {
}

func (UnimplementedProtocolPluginServiceServer) PluginDeviceConfig(context.Context, *PluginDeviceConfigRequest) (*PluginDeviceConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginDeviceConfig not implemented")
}
func (UnimplementedProtocolPluginServiceServer) PluginDeviceConfigList(context.Context, *PluginDeviceConfigListRequest) (*PluginDeviceConfigListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PluginDeviceConfigList not implemented")
}
func (UnimplementedProtocolPluginServiceServer) mustEmbedUnimplementedProtocolPluginServiceServer() {}

// UnsafeProtocolPluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProtocolPluginServiceServer will
// result in compilation errors.
type UnsafeProtocolPluginServiceServer interface {
	mustEmbedUnimplementedProtocolPluginServiceServer()
}

func RegisterProtocolPluginServiceServer(s grpc.ServiceRegistrar, srv ProtocolPluginServiceServer) {
	s.RegisterService(&ProtocolPluginService_ServiceDesc, srv)
}

func _ProtocolPluginService_PluginDeviceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginDeviceConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolPluginServiceServer).PluginDeviceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol_plugin.ProtocolPluginService/PluginDeviceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolPluginServiceServer).PluginDeviceConfig(ctx, req.(*PluginDeviceConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProtocolPluginService_PluginDeviceConfigList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PluginDeviceConfigListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProtocolPluginServiceServer).PluginDeviceConfigList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protocol_plugin.ProtocolPluginService/PluginDeviceConfigList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProtocolPluginServiceServer).PluginDeviceConfigList(ctx, req.(*PluginDeviceConfigListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProtocolPluginService_ServiceDesc is the grpc.ServiceDesc for ProtocolPluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProtocolPluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protocol_plugin.ProtocolPluginService",
	HandlerType: (*ProtocolPluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PluginDeviceConfig",
			Handler:    _ProtocolPluginService_PluginDeviceConfig_Handler,
		},
		{
			MethodName: "PluginDeviceConfigList",
			Handler:    _ProtocolPluginService_PluginDeviceConfigList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protocol_plugin.proto",
}
