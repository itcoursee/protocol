// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: com/itcoursee/protocol/lecture.proto

package protocol

import (
	context "context"
	lecture "github.com/itcoursee/core/lecture"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Lecture_Add_FullMethodName          = "/com.itcoursee.protocol.Lecture/Add"
	Lecture_Get_FullMethodName          = "/com.itcoursee.protocol.Lecture/Get"
	Lecture_GetsByCourse_FullMethodName = "/com.itcoursee.protocol.Lecture/GetsByCourse"
	Lecture_GetPlayback_FullMethodName  = "/com.itcoursee.protocol.Lecture/GetPlayback"
	Lecture_Update_FullMethodName       = "/com.itcoursee.protocol.Lecture/Update"
	Lecture_Resource_FullMethodName     = "/com.itcoursee.protocol.Lecture/Resource"
	Lecture_Delete_FullMethodName       = "/com.itcoursee.protocol.Lecture/Delete"
)

// LectureClient is the client API for Lecture service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LectureClient interface {
	// 创建
	Add(ctx context.Context, in *lecture.AddReq, opts ...grpc.CallOption) (*lecture.AddRsp, error)
	// 获取
	Get(ctx context.Context, in *lecture.GetReq, opts ...grpc.CallOption) (*lecture.GetRsp, error)
	// 按课程获得所有的讲次
	GetsByCourse(ctx context.Context, in *lecture.GetsByCourseReq, opts ...grpc.CallOption) (*lecture.GetsByCourseRsp, error)
	// 获得讲次的播放地址
	GetPlayback(ctx context.Context, in *lecture.GetPlaybackReq, opts ...grpc.CallOption) (*lecture.GetPlaybackRsp, error)
	// 修改
	Update(ctx context.Context, in *lecture.UpdateReq, opts ...grpc.CallOption) (*lecture.UpdateRsp, error)
	// 视频关联
	Resource(ctx context.Context, in *lecture.ResourceReq, opts ...grpc.CallOption) (*lecture.ResourceRsp, error)
	// 删除
	Delete(ctx context.Context, in *lecture.DeleteReq, opts ...grpc.CallOption) (*lecture.DeleteRsp, error)
}

type lectureClient struct {
	cc grpc.ClientConnInterface
}

func NewLectureClient(cc grpc.ClientConnInterface) LectureClient {
	return &lectureClient{cc}
}

func (c *lectureClient) Add(ctx context.Context, in *lecture.AddReq, opts ...grpc.CallOption) (*lecture.AddRsp, error) {
	out := new(lecture.AddRsp)
	err := c.cc.Invoke(ctx, Lecture_Add_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lectureClient) Get(ctx context.Context, in *lecture.GetReq, opts ...grpc.CallOption) (*lecture.GetRsp, error) {
	out := new(lecture.GetRsp)
	err := c.cc.Invoke(ctx, Lecture_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lectureClient) GetsByCourse(ctx context.Context, in *lecture.GetsByCourseReq, opts ...grpc.CallOption) (*lecture.GetsByCourseRsp, error) {
	out := new(lecture.GetsByCourseRsp)
	err := c.cc.Invoke(ctx, Lecture_GetsByCourse_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lectureClient) GetPlayback(ctx context.Context, in *lecture.GetPlaybackReq, opts ...grpc.CallOption) (*lecture.GetPlaybackRsp, error) {
	out := new(lecture.GetPlaybackRsp)
	err := c.cc.Invoke(ctx, Lecture_GetPlayback_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lectureClient) Update(ctx context.Context, in *lecture.UpdateReq, opts ...grpc.CallOption) (*lecture.UpdateRsp, error) {
	out := new(lecture.UpdateRsp)
	err := c.cc.Invoke(ctx, Lecture_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lectureClient) Resource(ctx context.Context, in *lecture.ResourceReq, opts ...grpc.CallOption) (*lecture.ResourceRsp, error) {
	out := new(lecture.ResourceRsp)
	err := c.cc.Invoke(ctx, Lecture_Resource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lectureClient) Delete(ctx context.Context, in *lecture.DeleteReq, opts ...grpc.CallOption) (*lecture.DeleteRsp, error) {
	out := new(lecture.DeleteRsp)
	err := c.cc.Invoke(ctx, Lecture_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LectureServer is the server API for Lecture service.
// All implementations must embed UnimplementedLectureServer
// for forward compatibility
type LectureServer interface {
	// 创建
	Add(context.Context, *lecture.AddReq) (*lecture.AddRsp, error)
	// 获取
	Get(context.Context, *lecture.GetReq) (*lecture.GetRsp, error)
	// 按课程获得所有的讲次
	GetsByCourse(context.Context, *lecture.GetsByCourseReq) (*lecture.GetsByCourseRsp, error)
	// 获得讲次的播放地址
	GetPlayback(context.Context, *lecture.GetPlaybackReq) (*lecture.GetPlaybackRsp, error)
	// 修改
	Update(context.Context, *lecture.UpdateReq) (*lecture.UpdateRsp, error)
	// 视频关联
	Resource(context.Context, *lecture.ResourceReq) (*lecture.ResourceRsp, error)
	// 删除
	Delete(context.Context, *lecture.DeleteReq) (*lecture.DeleteRsp, error)
	mustEmbedUnimplementedLectureServer()
}

// UnimplementedLectureServer must be embedded to have forward compatible implementations.
type UnimplementedLectureServer struct {
}

func (UnimplementedLectureServer) Add(context.Context, *lecture.AddReq) (*lecture.AddRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedLectureServer) Get(context.Context, *lecture.GetReq) (*lecture.GetRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLectureServer) GetsByCourse(context.Context, *lecture.GetsByCourseReq) (*lecture.GetsByCourseRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetsByCourse not implemented")
}
func (UnimplementedLectureServer) GetPlayback(context.Context, *lecture.GetPlaybackReq) (*lecture.GetPlaybackRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlayback not implemented")
}
func (UnimplementedLectureServer) Update(context.Context, *lecture.UpdateReq) (*lecture.UpdateRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedLectureServer) Resource(context.Context, *lecture.ResourceReq) (*lecture.ResourceRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Resource not implemented")
}
func (UnimplementedLectureServer) Delete(context.Context, *lecture.DeleteReq) (*lecture.DeleteRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLectureServer) mustEmbedUnimplementedLectureServer() {}

// UnsafeLectureServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LectureServer will
// result in compilation errors.
type UnsafeLectureServer interface {
	mustEmbedUnimplementedLectureServer()
}

func RegisterLectureServer(s grpc.ServiceRegistrar, srv LectureServer) {
	s.RegisterService(&Lecture_ServiceDesc, srv)
}

func _Lecture_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(lecture.AddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LectureServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lecture_Add_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LectureServer).Add(ctx, req.(*lecture.AddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lecture_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(lecture.GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LectureServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lecture_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LectureServer).Get(ctx, req.(*lecture.GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lecture_GetsByCourse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(lecture.GetsByCourseReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LectureServer).GetsByCourse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lecture_GetsByCourse_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LectureServer).GetsByCourse(ctx, req.(*lecture.GetsByCourseReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lecture_GetPlayback_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(lecture.GetPlaybackReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LectureServer).GetPlayback(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lecture_GetPlayback_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LectureServer).GetPlayback(ctx, req.(*lecture.GetPlaybackReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lecture_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(lecture.UpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LectureServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lecture_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LectureServer).Update(ctx, req.(*lecture.UpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lecture_Resource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(lecture.ResourceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LectureServer).Resource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lecture_Resource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LectureServer).Resource(ctx, req.(*lecture.ResourceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Lecture_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(lecture.DeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LectureServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Lecture_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LectureServer).Delete(ctx, req.(*lecture.DeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Lecture_ServiceDesc is the grpc.ServiceDesc for Lecture service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Lecture_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.itcoursee.protocol.Lecture",
	HandlerType: (*LectureServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Lecture_Add_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Lecture_Get_Handler,
		},
		{
			MethodName: "GetsByCourse",
			Handler:    _Lecture_GetsByCourse_Handler,
		},
		{
			MethodName: "GetPlayback",
			Handler:    _Lecture_GetPlayback_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Lecture_Update_Handler,
		},
		{
			MethodName: "Resource",
			Handler:    _Lecture_Resource_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Lecture_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/itcoursee/protocol/lecture.proto",
}