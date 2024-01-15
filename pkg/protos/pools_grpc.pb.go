// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: protos/pools.proto

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

const (
	PoolService_ListPools_FullMethodName = "/spqr.PoolService/ListPools"
)

// PoolServiceClient is the client API for PoolService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoolServiceClient interface {
	ListPools(ctx context.Context, in *ListPoolsRequest, opts ...grpc.CallOption) (*ListPoolsResponse, error)
}

type poolServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPoolServiceClient(cc grpc.ClientConnInterface) PoolServiceClient {
	return &poolServiceClient{cc}
}

func (c *poolServiceClient) ListPools(ctx context.Context, in *ListPoolsRequest, opts ...grpc.CallOption) (*ListPoolsResponse, error) {
	out := new(ListPoolsResponse)
	err := c.cc.Invoke(ctx, PoolService_ListPools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoolServiceServer is the server API for PoolService service.
// All implementations must embed UnimplementedPoolServiceServer
// for forward compatibility
type PoolServiceServer interface {
	ListPools(context.Context, *ListPoolsRequest) (*ListPoolsResponse, error)
	mustEmbedUnimplementedPoolServiceServer()
}

// UnimplementedPoolServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPoolServiceServer struct {
}

func (UnimplementedPoolServiceServer) ListPools(context.Context, *ListPoolsRequest) (*ListPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPools not implemented")
}
func (UnimplementedPoolServiceServer) mustEmbedUnimplementedPoolServiceServer() {}

// UnsafePoolServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoolServiceServer will
// result in compilation errors.
type UnsafePoolServiceServer interface {
	mustEmbedUnimplementedPoolServiceServer()
}

func RegisterPoolServiceServer(s grpc.ServiceRegistrar, srv PoolServiceServer) {
	s.RegisterService(&PoolService_ServiceDesc, srv)
}

func _PoolService_ListPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoolsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolServiceServer).ListPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PoolService_ListPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolServiceServer).ListPools(ctx, req.(*ListPoolsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PoolService_ServiceDesc is the grpc.ServiceDesc for PoolService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoolService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spqr.PoolService",
	HandlerType: (*PoolServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListPools",
			Handler:    _PoolService_ListPools_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/pools.proto",
}
