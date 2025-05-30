// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: protos/tasks.proto

package proto

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

const (
	MoveTasksService_GetMoveTaskGroup_FullMethodName    = "/spqr.MoveTasksService/GetMoveTaskGroup"
	MoveTasksService_WriteMoveTaskGroup_FullMethodName  = "/spqr.MoveTasksService/WriteMoveTaskGroup"
	MoveTasksService_RemoveMoveTaskGroup_FullMethodName = "/spqr.MoveTasksService/RemoveMoveTaskGroup"
	MoveTasksService_RetryMoveTaskGroup_FullMethodName  = "/spqr.MoveTasksService/RetryMoveTaskGroup"
)

// MoveTasksServiceClient is the client API for MoveTasksService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MoveTasksServiceClient interface {
	GetMoveTaskGroup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMoveTaskGroupReply, error)
	WriteMoveTaskGroup(ctx context.Context, in *WriteMoveTaskGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveMoveTaskGroup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RetryMoveTaskGroup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type moveTasksServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMoveTasksServiceClient(cc grpc.ClientConnInterface) MoveTasksServiceClient {
	return &moveTasksServiceClient{cc}
}

func (c *moveTasksServiceClient) GetMoveTaskGroup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetMoveTaskGroupReply, error) {
	out := new(GetMoveTaskGroupReply)
	err := c.cc.Invoke(ctx, MoveTasksService_GetMoveTaskGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moveTasksServiceClient) WriteMoveTaskGroup(ctx context.Context, in *WriteMoveTaskGroupRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MoveTasksService_WriteMoveTaskGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moveTasksServiceClient) RemoveMoveTaskGroup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MoveTasksService_RemoveMoveTaskGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *moveTasksServiceClient) RetryMoveTaskGroup(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, MoveTasksService_RetryMoveTaskGroup_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MoveTasksServiceServer is the server API for MoveTasksService service.
// All implementations must embed UnimplementedMoveTasksServiceServer
// for forward compatibility
type MoveTasksServiceServer interface {
	GetMoveTaskGroup(context.Context, *emptypb.Empty) (*GetMoveTaskGroupReply, error)
	WriteMoveTaskGroup(context.Context, *WriteMoveTaskGroupRequest) (*emptypb.Empty, error)
	RemoveMoveTaskGroup(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	RetryMoveTaskGroup(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedMoveTasksServiceServer()
}

// UnimplementedMoveTasksServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMoveTasksServiceServer struct {
}

func (UnimplementedMoveTasksServiceServer) GetMoveTaskGroup(context.Context, *emptypb.Empty) (*GetMoveTaskGroupReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMoveTaskGroup not implemented")
}
func (UnimplementedMoveTasksServiceServer) WriteMoveTaskGroup(context.Context, *WriteMoveTaskGroupRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteMoveTaskGroup not implemented")
}
func (UnimplementedMoveTasksServiceServer) RemoveMoveTaskGroup(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMoveTaskGroup not implemented")
}
func (UnimplementedMoveTasksServiceServer) RetryMoveTaskGroup(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetryMoveTaskGroup not implemented")
}
func (UnimplementedMoveTasksServiceServer) mustEmbedUnimplementedMoveTasksServiceServer() {}

// UnsafeMoveTasksServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MoveTasksServiceServer will
// result in compilation errors.
type UnsafeMoveTasksServiceServer interface {
	mustEmbedUnimplementedMoveTasksServiceServer()
}

func RegisterMoveTasksServiceServer(s grpc.ServiceRegistrar, srv MoveTasksServiceServer) {
	s.RegisterService(&MoveTasksService_ServiceDesc, srv)
}

func _MoveTasksService_GetMoveTaskGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoveTasksServiceServer).GetMoveTaskGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MoveTasksService_GetMoveTaskGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoveTasksServiceServer).GetMoveTaskGroup(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoveTasksService_WriteMoveTaskGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteMoveTaskGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoveTasksServiceServer).WriteMoveTaskGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MoveTasksService_WriteMoveTaskGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoveTasksServiceServer).WriteMoveTaskGroup(ctx, req.(*WriteMoveTaskGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoveTasksService_RemoveMoveTaskGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoveTasksServiceServer).RemoveMoveTaskGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MoveTasksService_RemoveMoveTaskGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoveTasksServiceServer).RemoveMoveTaskGroup(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MoveTasksService_RetryMoveTaskGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MoveTasksServiceServer).RetryMoveTaskGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MoveTasksService_RetryMoveTaskGroup_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MoveTasksServiceServer).RetryMoveTaskGroup(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MoveTasksService_ServiceDesc is the grpc.ServiceDesc for MoveTasksService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MoveTasksService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spqr.MoveTasksService",
	HandlerType: (*MoveTasksServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMoveTaskGroup",
			Handler:    _MoveTasksService_GetMoveTaskGroup_Handler,
		},
		{
			MethodName: "WriteMoveTaskGroup",
			Handler:    _MoveTasksService_WriteMoveTaskGroup_Handler,
		},
		{
			MethodName: "RemoveMoveTaskGroup",
			Handler:    _MoveTasksService_RemoveMoveTaskGroup_Handler,
		},
		{
			MethodName: "RetryMoveTaskGroup",
			Handler:    _MoveTasksService_RetryMoveTaskGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/tasks.proto",
}

const (
	BalancerTaskService_GetBalancerTask_FullMethodName    = "/spqr.BalancerTaskService/GetBalancerTask"
	BalancerTaskService_WriteBalancerTask_FullMethodName  = "/spqr.BalancerTaskService/WriteBalancerTask"
	BalancerTaskService_RemoveBalancerTask_FullMethodName = "/spqr.BalancerTaskService/RemoveBalancerTask"
)

// BalancerTaskServiceClient is the client API for BalancerTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BalancerTaskServiceClient interface {
	GetBalancerTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetBalancerTaskReply, error)
	WriteBalancerTask(ctx context.Context, in *WriteBalancerTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveBalancerTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type balancerTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBalancerTaskServiceClient(cc grpc.ClientConnInterface) BalancerTaskServiceClient {
	return &balancerTaskServiceClient{cc}
}

func (c *balancerTaskServiceClient) GetBalancerTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetBalancerTaskReply, error) {
	out := new(GetBalancerTaskReply)
	err := c.cc.Invoke(ctx, BalancerTaskService_GetBalancerTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerTaskServiceClient) WriteBalancerTask(ctx context.Context, in *WriteBalancerTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BalancerTaskService_WriteBalancerTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *balancerTaskServiceClient) RemoveBalancerTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, BalancerTaskService_RemoveBalancerTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BalancerTaskServiceServer is the server API for BalancerTaskService service.
// All implementations must embed UnimplementedBalancerTaskServiceServer
// for forward compatibility
type BalancerTaskServiceServer interface {
	GetBalancerTask(context.Context, *emptypb.Empty) (*GetBalancerTaskReply, error)
	WriteBalancerTask(context.Context, *WriteBalancerTaskRequest) (*emptypb.Empty, error)
	RemoveBalancerTask(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedBalancerTaskServiceServer()
}

// UnimplementedBalancerTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBalancerTaskServiceServer struct {
}

func (UnimplementedBalancerTaskServiceServer) GetBalancerTask(context.Context, *emptypb.Empty) (*GetBalancerTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalancerTask not implemented")
}
func (UnimplementedBalancerTaskServiceServer) WriteBalancerTask(context.Context, *WriteBalancerTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteBalancerTask not implemented")
}
func (UnimplementedBalancerTaskServiceServer) RemoveBalancerTask(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBalancerTask not implemented")
}
func (UnimplementedBalancerTaskServiceServer) mustEmbedUnimplementedBalancerTaskServiceServer() {}

// UnsafeBalancerTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BalancerTaskServiceServer will
// result in compilation errors.
type UnsafeBalancerTaskServiceServer interface {
	mustEmbedUnimplementedBalancerTaskServiceServer()
}

func RegisterBalancerTaskServiceServer(s grpc.ServiceRegistrar, srv BalancerTaskServiceServer) {
	s.RegisterService(&BalancerTaskService_ServiceDesc, srv)
}

func _BalancerTaskService_GetBalancerTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerTaskServiceServer).GetBalancerTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BalancerTaskService_GetBalancerTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerTaskServiceServer).GetBalancerTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalancerTaskService_WriteBalancerTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteBalancerTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerTaskServiceServer).WriteBalancerTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BalancerTaskService_WriteBalancerTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerTaskServiceServer).WriteBalancerTask(ctx, req.(*WriteBalancerTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BalancerTaskService_RemoveBalancerTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BalancerTaskServiceServer).RemoveBalancerTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BalancerTaskService_RemoveBalancerTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BalancerTaskServiceServer).RemoveBalancerTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// BalancerTaskService_ServiceDesc is the grpc.ServiceDesc for BalancerTaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BalancerTaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spqr.BalancerTaskService",
	HandlerType: (*BalancerTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBalancerTask",
			Handler:    _BalancerTaskService_GetBalancerTask_Handler,
		},
		{
			MethodName: "WriteBalancerTask",
			Handler:    _BalancerTaskService_WriteBalancerTask_Handler,
		},
		{
			MethodName: "RemoveBalancerTask",
			Handler:    _BalancerTaskService_RemoveBalancerTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/tasks.proto",
}

const (
	RedistributeTaskService_GetRedistributeTask_FullMethodName    = "/spqr.RedistributeTaskService/GetRedistributeTask"
	RedistributeTaskService_WriteRedistributeTask_FullMethodName  = "/spqr.RedistributeTaskService/WriteRedistributeTask"
	RedistributeTaskService_RemoveRedistributeTask_FullMethodName = "/spqr.RedistributeTaskService/RemoveRedistributeTask"
)

// RedistributeTaskServiceClient is the client API for RedistributeTaskService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RedistributeTaskServiceClient interface {
	GetRedistributeTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRedistributeTaskReply, error)
	WriteRedistributeTask(ctx context.Context, in *WriteRedistributeTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RemoveRedistributeTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type redistributeTaskServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRedistributeTaskServiceClient(cc grpc.ClientConnInterface) RedistributeTaskServiceClient {
	return &redistributeTaskServiceClient{cc}
}

func (c *redistributeTaskServiceClient) GetRedistributeTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetRedistributeTaskReply, error) {
	out := new(GetRedistributeTaskReply)
	err := c.cc.Invoke(ctx, RedistributeTaskService_GetRedistributeTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redistributeTaskServiceClient) WriteRedistributeTask(ctx context.Context, in *WriteRedistributeTaskRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RedistributeTaskService_WriteRedistributeTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redistributeTaskServiceClient) RemoveRedistributeTask(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, RedistributeTaskService_RemoveRedistributeTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedistributeTaskServiceServer is the server API for RedistributeTaskService service.
// All implementations must embed UnimplementedRedistributeTaskServiceServer
// for forward compatibility
type RedistributeTaskServiceServer interface {
	GetRedistributeTask(context.Context, *emptypb.Empty) (*GetRedistributeTaskReply, error)
	WriteRedistributeTask(context.Context, *WriteRedistributeTaskRequest) (*emptypb.Empty, error)
	RemoveRedistributeTask(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedRedistributeTaskServiceServer()
}

// UnimplementedRedistributeTaskServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRedistributeTaskServiceServer struct {
}

func (UnimplementedRedistributeTaskServiceServer) GetRedistributeTask(context.Context, *emptypb.Empty) (*GetRedistributeTaskReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRedistributeTask not implemented")
}
func (UnimplementedRedistributeTaskServiceServer) WriteRedistributeTask(context.Context, *WriteRedistributeTaskRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteRedistributeTask not implemented")
}
func (UnimplementedRedistributeTaskServiceServer) RemoveRedistributeTask(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRedistributeTask not implemented")
}
func (UnimplementedRedistributeTaskServiceServer) mustEmbedUnimplementedRedistributeTaskServiceServer() {
}

// UnsafeRedistributeTaskServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedistributeTaskServiceServer will
// result in compilation errors.
type UnsafeRedistributeTaskServiceServer interface {
	mustEmbedUnimplementedRedistributeTaskServiceServer()
}

func RegisterRedistributeTaskServiceServer(s grpc.ServiceRegistrar, srv RedistributeTaskServiceServer) {
	s.RegisterService(&RedistributeTaskService_ServiceDesc, srv)
}

func _RedistributeTaskService_GetRedistributeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedistributeTaskServiceServer).GetRedistributeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedistributeTaskService_GetRedistributeTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedistributeTaskServiceServer).GetRedistributeTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedistributeTaskService_WriteRedistributeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRedistributeTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedistributeTaskServiceServer).WriteRedistributeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedistributeTaskService_WriteRedistributeTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedistributeTaskServiceServer).WriteRedistributeTask(ctx, req.(*WriteRedistributeTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedistributeTaskService_RemoveRedistributeTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedistributeTaskServiceServer).RemoveRedistributeTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedistributeTaskService_RemoveRedistributeTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedistributeTaskServiceServer).RemoveRedistributeTask(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// RedistributeTaskService_ServiceDesc is the grpc.ServiceDesc for RedistributeTaskService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedistributeTaskService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spqr.RedistributeTaskService",
	HandlerType: (*RedistributeTaskServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRedistributeTask",
			Handler:    _RedistributeTaskService_GetRedistributeTask_Handler,
		},
		{
			MethodName: "WriteRedistributeTask",
			Handler:    _RedistributeTaskService_WriteRedistributeTask_Handler,
		},
		{
			MethodName: "RemoveRedistributeTask",
			Handler:    _RedistributeTaskService_RemoveRedistributeTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/tasks.proto",
}
