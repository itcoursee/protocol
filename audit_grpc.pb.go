// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.9
// source: com/itcoursee/protocol/audit.proto

package protocol

import (
	context "context"
	audit "github.com/itcoursee/core/audit"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	Audit_Paging_FullMethodName    = "/com.itcoursee.protocol.Audit/Paging"
	Audit_Approval_FullMethodName  = "/com.itcoursee.protocol.Audit/Approval"
	Audit_Rejection_FullMethodName = "/com.itcoursee.protocol.Audit/Rejection"
)

// AuditClient is the client API for Audit service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuditClient interface {
	// 分页
	Paging(ctx context.Context, in *audit.PagingReq, opts ...grpc.CallOption) (*audit.PagingRsp, error)
	// 同意
	Approval(ctx context.Context, in *audit.ApprovalReq, opts ...grpc.CallOption) (*audit.ApprovalRsp, error)
	// 拒绝
	Rejection(ctx context.Context, in *audit.RejectionReq, opts ...grpc.CallOption) (*audit.RejectionRsp, error)
}

type auditClient struct {
	cc grpc.ClientConnInterface
}

func NewAuditClient(cc grpc.ClientConnInterface) AuditClient {
	return &auditClient{cc}
}

func (c *auditClient) Paging(ctx context.Context, in *audit.PagingReq, opts ...grpc.CallOption) (*audit.PagingRsp, error) {
	out := new(audit.PagingRsp)
	err := c.cc.Invoke(ctx, Audit_Paging_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditClient) Approval(ctx context.Context, in *audit.ApprovalReq, opts ...grpc.CallOption) (*audit.ApprovalRsp, error) {
	out := new(audit.ApprovalRsp)
	err := c.cc.Invoke(ctx, Audit_Approval_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *auditClient) Rejection(ctx context.Context, in *audit.RejectionReq, opts ...grpc.CallOption) (*audit.RejectionRsp, error) {
	out := new(audit.RejectionRsp)
	err := c.cc.Invoke(ctx, Audit_Rejection_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuditServer is the server API for Audit service.
// All implementations must embed UnimplementedAuditServer
// for forward compatibility
type AuditServer interface {
	// 分页
	Paging(context.Context, *audit.PagingReq) (*audit.PagingRsp, error)
	// 同意
	Approval(context.Context, *audit.ApprovalReq) (*audit.ApprovalRsp, error)
	// 拒绝
	Rejection(context.Context, *audit.RejectionReq) (*audit.RejectionRsp, error)
	mustEmbedUnimplementedAuditServer()
}

// UnimplementedAuditServer must be embedded to have forward compatible implementations.
type UnimplementedAuditServer struct {
}

func (UnimplementedAuditServer) Paging(context.Context, *audit.PagingReq) (*audit.PagingRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Paging not implemented")
}
func (UnimplementedAuditServer) Approval(context.Context, *audit.ApprovalReq) (*audit.ApprovalRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Approval not implemented")
}
func (UnimplementedAuditServer) Rejection(context.Context, *audit.RejectionReq) (*audit.RejectionRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rejection not implemented")
}
func (UnimplementedAuditServer) mustEmbedUnimplementedAuditServer() {}

// UnsafeAuditServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuditServer will
// result in compilation errors.
type UnsafeAuditServer interface {
	mustEmbedUnimplementedAuditServer()
}

func RegisterAuditServer(s grpc.ServiceRegistrar, srv AuditServer) {
	s.RegisterService(&Audit_ServiceDesc, srv)
}

func _Audit_Paging_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(audit.PagingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServer).Paging(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Audit_Paging_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServer).Paging(ctx, req.(*audit.PagingReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Audit_Approval_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(audit.ApprovalReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServer).Approval(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Audit_Approval_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServer).Approval(ctx, req.(*audit.ApprovalReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Audit_Rejection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(audit.RejectionReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuditServer).Rejection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Audit_Rejection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuditServer).Rejection(ctx, req.(*audit.RejectionReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Audit_ServiceDesc is the grpc.ServiceDesc for Audit service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Audit_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "com.itcoursee.protocol.Audit",
	HandlerType: (*AuditServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Paging",
			Handler:    _Audit_Paging_Handler,
		},
		{
			MethodName: "Approval",
			Handler:    _Audit_Approval_Handler,
		},
		{
			MethodName: "Rejection",
			Handler:    _Audit_Rejection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "com/itcoursee/protocol/audit.proto",
}
