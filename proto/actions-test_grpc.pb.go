// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// ActionsClient is the client API for Actions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActionsClient interface {
	// Echoes back the request
	ActionEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error)
}

type actionsClient struct {
	cc grpc.ClientConnInterface
}

func NewActionsClient(cc grpc.ClientConnInterface) ActionsClient {
	return &actionsClient{cc}
}

func (c *actionsClient) ActionEcho(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoReply, error) {
	out := new(EchoReply)
	err := c.cc.Invoke(ctx, "/actions.Actions/ActionEcho", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActionsServer is the server API for Actions service.
// All implementations must embed UnimplementedActionsServer
// for forward compatibility
type ActionsServer interface {
	// Echoes back the request
	ActionEcho(context.Context, *EchoRequest) (*EchoReply, error)
	mustEmbedUnimplementedActionsServer()
}

// UnimplementedActionsServer must be embedded to have forward compatible implementations.
type UnimplementedActionsServer struct {
}

func (UnimplementedActionsServer) ActionEcho(context.Context, *EchoRequest) (*EchoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionEcho not implemented")
}
func (UnimplementedActionsServer) mustEmbedUnimplementedActionsServer() {}

// UnsafeActionsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActionsServer will
// result in compilation errors.
type UnsafeActionsServer interface {
	mustEmbedUnimplementedActionsServer()
}

func RegisterActionsServer(s grpc.ServiceRegistrar, srv ActionsServer) {
	s.RegisterService(&Actions_ServiceDesc, srv)
}

func _Actions_ActionEcho_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActionsServer).ActionEcho(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/actions.Actions/ActionEcho",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActionsServer).ActionEcho(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Actions_ServiceDesc is the grpc.ServiceDesc for Actions service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Actions_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "actions.Actions",
	HandlerType: (*ActionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ActionEcho",
			Handler:    _Actions_ActionEcho_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/actions-test.proto",
}