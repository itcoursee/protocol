// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: com/itcoursee/protocol/file.proto

package protocol

import (
	context "context"
	file "github.com/itcoursee/core/file"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	File_Initiate_FullMethodName = "/com.itcoursee.protocol.File/Initiate"
	File_Paging_FullMethodName   = "/com.itcoursee.protocol.File/Paging"
	File_Complete_FullMethodName = "/com.itcoursee.protocol.File/Complete"
	File_Update_FullMethodName   = "/com.itcoursee.protocol.File/Update"
	File_Delete_FullMethodName   = "/com.itcoursee.protocol.File/Delete"
)

// FileClient is the client API for File service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FileClient interface {
	// 初始化
	Initiate(ctx context.Context, in *file.InitiateReq, opts ...grpc.CallOption) (*file.InitiateRsp, error)
	// 分页
	Paging(ctx context.Context, in *file.PagingReq, opts ...grpc.CallOption) (*file.PagingRsp, error)
	// 完成
	Complete(ctx context.Context, in *file.CompleteReq, opts ...grpc.CallOption) (*file.CompleteRsp, error)
	// 修改
	Update(ctx context.Context, in *file.UpdateReq, opts ...grpc.CallOption) (*file.UpdateRsp, error)
	// 删除
	Delete(ctx context.Context, in *file.DeleteReq, opts ...grpc.CallOption) (*file.DeleteRsp, error)
}

type fileClient struct {
	cc grpc.ClientConnInterface
}

func NewFileClient(cc grpc.ClientConnInterface) FileClient {
	return &fileClient{cc}
}

func (c *fileClient) Initiate(ctx context.Context, in *file.InitiateReq, opts ...grpc.CallOption) (*file.InitiateRsp, error) {
	out := new(file.InitiateRsp)
	err := c.cc.Invoke(ctx, File_Initiate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Paging(ctx context.Context, in *file.PagingReq, opts ...grpc.CallOption) (*file.PagingRsp, error) {
	out := new(file.PagingRsp)
	err := c.cc.Invoke(ctx, File_Paging_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Complete(ctx context.Context, in *file.CompleteReq, opts ...grpc.CallOption) (*file.CompleteRsp, error) {
	out := new(file.CompleteRsp)
	err := c.cc.Invoke(ctx, File_Complete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Update(ctx context.Context, in *file.UpdateReq, opts ...grpc.CallOption) (*file.UpdateRsp, error) {
	out := new(file.UpdateRsp)
	err := c.cc.Invoke(ctx, File_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fileClient) Delete(ctx context.Context, in *file.DeleteReq, opts ...grpc.CallOption) (*file.DeleteRsp, error) {
	out := new(file.DeleteRsp)
	err := c.cc.Invoke(ctx, File_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FileServer is the server API for File service.
// All implementations must embed UnimplementedFileServer
// for forward compatibility
type FileServer interface {
	// 初始化
	Initiate(context.Context, *file.InitiateReq) (*file.InitiateRsp, error)
	// 分页
	Paging(context.Context, *file.PagingReq) (*file.PagingRsp, error)
	// 完成
	Complete(context.Context, *file.CompleteReq) (*file.CompleteRsp, error)
	// 修改
	Update(context.Context, *file.UpdateReq) (*file.UpdateRsp, error)
	// 删除
	Delete(context.Context, *file.DeleteReq) (*file.DeleteRsp, error)
	mustEmbedUnimplementedFileServer()
}

// UnimplementedFileServer must be embedded to have forward compatible implementations.
type UnimplementedFileServer struct {
}

func (UnimplementedFileServer) Initiate(context.Context, *file.InitiateReq) (*file.InitiateRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initiate not implemented")
}
func (UnimplementedFileServer) Paging(context.Context, *file.PagingReq) (*file.PagingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Paging not implemented")
}
func (UnimplementedFileServer) Complete(context.Context, *file.CompleteReq) (*file.CompleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Complete not implemented")
}
func (UnimplementedFileServer) Update(context.Context, *file.UpdateReq) (*file.UpdateRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedFileServer) Delete(context.Context, *file.DeleteReq) (*file.DeleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedFileServer) mustEmbedUnimplementedFileServer() {}

// UnsafeFileServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FileServer will
// result in compilation errors.
type UnsafeFileServer interface {
	mustEmbedUnimplementedFileServer()
}

func RegisterFileServer(s grpc.ServiceRegistrar, srv FileServer) {
	s.RegisterService(&File_ServiceDesc, srv)
}

func _File_Initiate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(file.InitiateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Initiate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Initiate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Initiate(ctx, req.(*file.InitiateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Paging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(file.PagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Paging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Paging_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Paging(ctx, req.(*file.PagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Complete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(file.CompleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Complete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Complete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Complete(ctx, req.(*file.CompleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(file.UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Update(ctx, req.(*file.UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _File_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(file.DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FileServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: File_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FileServer).Delete(ctx, req.(*file.DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// File_ServiceDesc is the grpc.ServiceDesc for File service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var File_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.itcoursee.protocol.File",
	HandlerType: (*FileServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Initiate",
			Handler:    _File_Initiate_Handler,
		},
		{
			MethodName: "Paging",
			Handler:    _File_Paging_Handler,
		},
		{
			MethodName: "Complete",
			Handler:    _File_Complete_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _File_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _File_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/itcoursee/protocol/file.proto",
}