// Copyright (c) Abstract Machines
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: channels/v1/channels.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChannelsService_Authorize_FullMethodName                    = "/channels.v1.ChannelsService/Authorize"
	ChannelsService_RemoveThingConnections_FullMethodName       = "/channels.v1.ChannelsService/RemoveThingConnections"
	ChannelsService_UnsetParentGroupFromChannels_FullMethodName = "/channels.v1.ChannelsService/UnsetParentGroupFromChannels"
)

// ChannelsServiceClient is the client API for ChannelsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChannelsServiceClient interface {
	Authorize(ctx context.Context, in *AuthzReq, opts ...grpc.CallOption) (*AuthzRes, error)
	RemoveThingConnections(ctx context.Context, in *RemoveThingConnectionsReq, opts ...grpc.CallOption) (*RemoveThingConnectionsRes, error)
	UnsetParentGroupFromChannels(ctx context.Context, in *UnsetParentGroupFromChannelsReq, opts ...grpc.CallOption) (*UnsetParentGroupFromChannelsRes, error)
}

type channelsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChannelsServiceClient(cc grpc.ClientConnInterface) ChannelsServiceClient {
	return &channelsServiceClient{cc}
}

func (c *channelsServiceClient) Authorize(ctx context.Context, in *AuthzReq, opts ...grpc.CallOption) (*AuthzRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthzRes)
	err := c.cc.Invoke(ctx, ChannelsService_Authorize_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsServiceClient) RemoveThingConnections(ctx context.Context, in *RemoveThingConnectionsReq, opts ...grpc.CallOption) (*RemoveThingConnectionsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveThingConnectionsRes)
	err := c.cc.Invoke(ctx, ChannelsService_RemoveThingConnections_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *channelsServiceClient) UnsetParentGroupFromChannels(ctx context.Context, in *UnsetParentGroupFromChannelsReq, opts ...grpc.CallOption) (*UnsetParentGroupFromChannelsRes, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UnsetParentGroupFromChannelsRes)
	err := c.cc.Invoke(ctx, ChannelsService_UnsetParentGroupFromChannels_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChannelsServiceServer is the server API for ChannelsService service.
// All implementations must embed UnimplementedChannelsServiceServer
// for forward compatibility.
type ChannelsServiceServer interface {
	Authorize(context.Context, *AuthzReq) (*AuthzRes, error)
	RemoveThingConnections(context.Context, *RemoveThingConnectionsReq) (*RemoveThingConnectionsRes, error)
	UnsetParentGroupFromChannels(context.Context, *UnsetParentGroupFromChannelsReq) (*UnsetParentGroupFromChannelsRes, error)
	mustEmbedUnimplementedChannelsServiceServer()
}

// UnimplementedChannelsServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChannelsServiceServer struct{}

func (UnimplementedChannelsServiceServer) Authorize(context.Context, *AuthzReq) (*AuthzRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorize not implemented")
}
func (UnimplementedChannelsServiceServer) RemoveThingConnections(context.Context, *RemoveThingConnectionsReq) (*RemoveThingConnectionsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveThingConnections not implemented")
}
func (UnimplementedChannelsServiceServer) UnsetParentGroupFromChannels(context.Context, *UnsetParentGroupFromChannelsReq) (*UnsetParentGroupFromChannelsRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnsetParentGroupFromChannels not implemented")
}
func (UnimplementedChannelsServiceServer) mustEmbedUnimplementedChannelsServiceServer() {}
func (UnimplementedChannelsServiceServer) testEmbeddedByValue()                         {}

// UnsafeChannelsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChannelsServiceServer will
// result in compilation errors.
type UnsafeChannelsServiceServer interface {
	mustEmbedUnimplementedChannelsServiceServer()
}

func RegisterChannelsServiceServer(s grpc.ServiceRegistrar, srv ChannelsServiceServer) {
	// If the following call pancis, it indicates UnimplementedChannelsServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChannelsService_ServiceDesc, srv)
}

func _ChannelsService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthzReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelsServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelsService_Authorize_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelsServiceServer).Authorize(ctx, req.(*AuthzReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelsService_RemoveThingConnections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveThingConnectionsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelsServiceServer).RemoveThingConnections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelsService_RemoveThingConnections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelsServiceServer).RemoveThingConnections(ctx, req.(*RemoveThingConnectionsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChannelsService_UnsetParentGroupFromChannels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsetParentGroupFromChannelsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChannelsServiceServer).UnsetParentGroupFromChannels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChannelsService_UnsetParentGroupFromChannels_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChannelsServiceServer).UnsetParentGroupFromChannels(ctx, req.(*UnsetParentGroupFromChannelsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChannelsService_ServiceDesc is the grpc.ServiceDesc for ChannelsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChannelsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "channels.v1.ChannelsService",
	HandlerType: (*ChannelsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authorize",
			Handler:    _ChannelsService_Authorize_Handler,
		},
		{
			MethodName: "RemoveThingConnections",
			Handler:    _ChannelsService_RemoveThingConnections_Handler,
		},
		{
			MethodName: "UnsetParentGroupFromChannels",
			Handler:    _ChannelsService_UnsetParentGroupFromChannels_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "channels/v1/channels.proto",
}
