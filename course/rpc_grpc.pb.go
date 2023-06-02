// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: com/itcoursee/protocol/upload/rpc.proto

package course

import (
	context "context"
	upload "github.com/itcoursee/protocol/upload"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	// 创建
	Add(ctx context.Context, in *upload.InitiateReq, opts ...grpc.CallOption) (*upload.InitiateRsp, error)
	// 完成
	Complete(ctx context.Context, in *upload.CompleteReq, opts ...grpc.CallOption) (*upload.CompleteRsp, error)
	// 取消
	Abort(ctx context.Context, in *upload.AbortReq, opts ...grpc.CallOption) (*upload.AbortRsp, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) Add(ctx context.Context, in *upload.InitiateReq, opts ...grpc.CallOption) (*upload.InitiateRsp, error) {
	out := new(upload.InitiateRsp)
	err := c.cc.Invoke(ctx, "/com.itcoursee.protocol.upload.Rpc/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Complete(ctx context.Context, in *upload.CompleteReq, opts ...grpc.CallOption) (*upload.CompleteRsp, error) {
	out := new(upload.CompleteRsp)
	err := c.cc.Invoke(ctx, "/com.itcoursee.protocol.upload.Rpc/Complete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Abort(ctx context.Context, in *upload.AbortReq, opts ...grpc.CallOption) (*upload.AbortRsp, error) {
	out := new(upload.AbortRsp)
	err := c.cc.Invoke(ctx, "/com.itcoursee.protocol.upload.Rpc/Abort", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RpcServer is the server API for Rpc service.
// All implementations must embed UnimplementedRpcServer
// for forward compatibility
type RpcServer interface {
	// 创建
	Add(context.Context, *upload.InitiateReq) (*upload.InitiateRsp, error)
	// 完成
	Complete(context.Context, *upload.CompleteReq) (*upload.CompleteRsp, error)
	// 取消
	Abort(context.Context, *upload.AbortReq) (*upload.AbortRsp, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) Add(context.Context, *upload.InitiateReq) (*upload.InitiateRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedRpcServer) Complete(context.Context, *upload.CompleteReq) (*upload.CompleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complete not implemented")
}
func (UnimplementedRpcServer) Abort(context.Context, *upload.AbortReq) (*upload.AbortRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Abort not implemented")
}
func (UnimplementedRpcServer) mustEmbedUnimplementedRpcServer() {}

// UnsafeRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RpcServer will
// result in compilation errors.
type UnsafeRpcServer interface {
	mustEmbedUnimplementedRpcServer()
}

func RegisterRpcServer(s grpc.ServiceRegistrar, srv RpcServer) {
	s.RegisterService(&Rpc_ServiceDesc, srv)
}

func _Rpc_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(upload.InitiateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.itcoursee.protocol.upload.Rpc/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Add(ctx, req.(*upload.InitiateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(upload.CompleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.itcoursee.protocol.upload.Rpc/Complete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Complete(ctx, req.(*upload.CompleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Abort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(upload.AbortReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Abort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.itcoursee.protocol.upload.Rpc/Abort",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Abort(ctx, req.(*upload.AbortReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.itcoursee.protocol.upload.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Rpc_Add_Handler,
		},
		{
			MethodName: "Complete",
			Handler:    _Rpc_Complete_Handler,
		},
		{
			MethodName: "Abort",
			Handler:    _Rpc_Abort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/itcoursee/protocol/upload/rpc.proto",
}
