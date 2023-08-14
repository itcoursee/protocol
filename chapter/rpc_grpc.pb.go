// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.18.1
// source: com/itcoursee/protocol/chapter/rpc.proto

package chapter

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

// RpcClient is the client API for Rpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RpcClient interface {
	// 创建
	Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRsp, error)
	// 获取
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error)
	// 获得课程所有
	GetsByCourse(ctx context.Context, in *GetsByCourseReq, opts ...grpc.CallOption) (*GetsByCourseRsp, error)
	// 修改
	Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRsp, error)
	// 删除
	Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRsp, error)
}

type rpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRpcClient(cc grpc.ClientConnInterface) RpcClient {
	return &rpcClient{cc}
}

func (c *rpcClient) Add(ctx context.Context, in *AddReq, opts ...grpc.CallOption) (*AddRsp, error) {
	out := new(AddRsp)
	err := c.cc.Invoke(ctx, "/com.itcoursee.protocol.chapter.Rpc/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetRsp, error) {
	out := new(GetRsp)
	err := c.cc.Invoke(ctx, "/com.itcoursee.protocol.chapter.Rpc/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) GetsByCourse(ctx context.Context, in *GetsByCourseReq, opts ...grpc.CallOption) (*GetsByCourseRsp, error) {
	out := new(GetsByCourseRsp)
	err := c.cc.Invoke(ctx, "/com.itcoursee.protocol.chapter.Rpc/GetsByCourse", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Update(ctx context.Context, in *UpdateReq, opts ...grpc.CallOption) (*UpdateRsp, error) {
	out := new(UpdateRsp)
	err := c.cc.Invoke(ctx, "/com.itcoursee.protocol.chapter.Rpc/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rpcClient) Delete(ctx context.Context, in *DeleteReq, opts ...grpc.CallOption) (*DeleteRsp, error) {
	out := new(DeleteRsp)
	err := c.cc.Invoke(ctx, "/com.itcoursee.protocol.chapter.Rpc/Delete", in, out, opts...)
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
	Add(context.Context, *AddReq) (*AddRsp, error)
	// 获取
	Get(context.Context, *GetReq) (*GetRsp, error)
	// 获得课程所有
	GetsByCourse(context.Context, *GetsByCourseReq) (*GetsByCourseRsp, error)
	// 修改
	Update(context.Context, *UpdateReq) (*UpdateRsp, error)
	// 删除
	Delete(context.Context, *DeleteReq) (*DeleteRsp, error)
	mustEmbedUnimplementedRpcServer()
}

// UnimplementedRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRpcServer struct {
}

func (UnimplementedRpcServer) Add(context.Context, *AddReq) (*AddRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedRpcServer) Get(context.Context, *GetReq) (*GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedRpcServer) GetsByCourse(context.Context, *GetsByCourseReq) (*GetsByCourseRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetsByCourse not implemented")
}
func (UnimplementedRpcServer) Update(context.Context, *UpdateReq) (*UpdateRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedRpcServer) Delete(context.Context, *DeleteReq) (*DeleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
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
	in := new(AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.itcoursee.protocol.chapter.Rpc/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Add(ctx, req.(*AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.itcoursee.protocol.chapter.Rpc/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_GetsByCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetsByCourseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).GetsByCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.itcoursee.protocol.chapter.Rpc/GetsByCourse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).GetsByCourse(ctx, req.(*GetsByCourseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.itcoursee.protocol.chapter.Rpc/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Update(ctx, req.(*UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rpc_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RpcServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/com.itcoursee.protocol.chapter.Rpc/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RpcServer).Delete(ctx, req.(*DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Rpc_ServiceDesc is the grpc.ServiceDesc for Rpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Rpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.itcoursee.protocol.chapter.Rpc",
	HandlerType: (*RpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Rpc_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Rpc_Get_Handler,
		},
		{
			MethodName: "GetsByCourse",
			Handler:    _Rpc_GetsByCourse_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Rpc_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Rpc_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/itcoursee/protocol/chapter/rpc.proto",
}
