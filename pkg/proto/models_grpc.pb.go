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

// MachineServiceClient is the client API for MachineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MachineServiceClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddResult, error)
	AddPassword(ctx context.Context, in *AddPasswordRequest, opts ...grpc.CallOption) (*AddResult, error)
	AddKey(ctx context.Context, in *AddKeyRequest, opts ...grpc.CallOption) (*AddResult, error)
	AddHost(ctx context.Context, in *AddHostRequest, opts ...grpc.CallOption) (*AddResult, error)
	GetMachine(ctx context.Context, in *MachineRequest, opts ...grpc.CallOption) (*Machine, error)
}

type machineServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMachineServiceClient(cc grpc.ClientConnInterface) MachineServiceClient {
	return &machineServiceClient{cc}
}

func (c *machineServiceClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddResult, error) {
	out := new(AddResult)
	err := c.cc.Invoke(ctx, "/proto.MachineService/AddUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) AddPassword(ctx context.Context, in *AddPasswordRequest, opts ...grpc.CallOption) (*AddResult, error) {
	out := new(AddResult)
	err := c.cc.Invoke(ctx, "/proto.MachineService/AddPassword", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) AddKey(ctx context.Context, in *AddKeyRequest, opts ...grpc.CallOption) (*AddResult, error) {
	out := new(AddResult)
	err := c.cc.Invoke(ctx, "/proto.MachineService/AddKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) AddHost(ctx context.Context, in *AddHostRequest, opts ...grpc.CallOption) (*AddResult, error) {
	out := new(AddResult)
	err := c.cc.Invoke(ctx, "/proto.MachineService/AddHost", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *machineServiceClient) GetMachine(ctx context.Context, in *MachineRequest, opts ...grpc.CallOption) (*Machine, error) {
	out := new(Machine)
	err := c.cc.Invoke(ctx, "/proto.MachineService/GetMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MachineServiceServer is the server API for MachineService service.
// All implementations must embed UnimplementedMachineServiceServer
// for forward compatibility
type MachineServiceServer interface {
	AddUser(context.Context, *AddUserRequest) (*AddResult, error)
	AddPassword(context.Context, *AddPasswordRequest) (*AddResult, error)
	AddKey(context.Context, *AddKeyRequest) (*AddResult, error)
	AddHost(context.Context, *AddHostRequest) (*AddResult, error)
	GetMachine(context.Context, *MachineRequest) (*Machine, error)
	mustEmbedUnimplementedMachineServiceServer()
}

// UnimplementedMachineServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMachineServiceServer struct {
}

func (UnimplementedMachineServiceServer) AddUser(context.Context, *AddUserRequest) (*AddResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedMachineServiceServer) AddPassword(context.Context, *AddPasswordRequest) (*AddResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPassword not implemented")
}
func (UnimplementedMachineServiceServer) AddKey(context.Context, *AddKeyRequest) (*AddResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddKey not implemented")
}
func (UnimplementedMachineServiceServer) AddHost(context.Context, *AddHostRequest) (*AddResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddHost not implemented")
}
func (UnimplementedMachineServiceServer) GetMachine(context.Context, *MachineRequest) (*Machine, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMachine not implemented")
}
func (UnimplementedMachineServiceServer) mustEmbedUnimplementedMachineServiceServer() {}

// UnsafeMachineServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MachineServiceServer will
// result in compilation errors.
type UnsafeMachineServiceServer interface {
	mustEmbedUnimplementedMachineServiceServer()
}

func RegisterMachineServiceServer(s grpc.ServiceRegistrar, srv MachineServiceServer) {
	s.RegisterService(&MachineService_ServiceDesc, srv)
}

func _MachineService_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MachineService/AddUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_AddPassword_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPasswordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).AddPassword(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MachineService/AddPassword",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).AddPassword(ctx, req.(*AddPasswordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_AddKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).AddKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MachineService/AddKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).AddKey(ctx, req.(*AddKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_AddHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).AddHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MachineService/AddHost",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).AddHost(ctx, req.(*AddHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MachineService_GetMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MachineServiceServer).GetMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.MachineService/GetMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MachineServiceServer).GetMachine(ctx, req.(*MachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MachineService_ServiceDesc is the grpc.ServiceDesc for MachineService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MachineService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MachineService",
	HandlerType: (*MachineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _MachineService_AddUser_Handler,
		},
		{
			MethodName: "AddPassword",
			Handler:    _MachineService_AddPassword_Handler,
		},
		{
			MethodName: "AddKey",
			Handler:    _MachineService_AddKey_Handler,
		},
		{
			MethodName: "AddHost",
			Handler:    _MachineService_AddHost_Handler,
		},
		{
			MethodName: "GetMachine",
			Handler:    _MachineService_GetMachine_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/models.proto",
}
