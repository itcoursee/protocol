// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/protocol/lecture/rpc.proto

package lecture

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_com_itcoursee_protocol_lecture_rpc_proto protoreflect.FileDescriptor

var file_com_itcoursee_protocol_lecture_rpc_proto_rawDesc = []byte{
	0x0a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2d, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x84, 0x07, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x7f, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x73,
	0x70, 0x22, 0x28, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x22, 0x1d, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x8a, 0x01, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x73, 0x70, 0x22, 0x33, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12, 0x22, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64,
	0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x62,
	0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0xa1, 0x01, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x2f, 0x2e, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x73,
	0x42, 0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x73, 0x70, 0x22, 0x2f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x29, 0x12, 0x1d, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x62, 0x08, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0xad, 0x01, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x2e, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x2e, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x52, 0x73, 0x70, 0x22, 0x3e, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x38, 0x12, 0x2c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x7b,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x2f, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63,
	0x6b, 0x73, 0x62, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x8d, 0x01, 0x0a,
	0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x2d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x27, 0x1a, 0x22, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x8a, 0x01, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x2a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x24, 0x2a, 0x22, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f,
	0x7b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x7d, 0x2f, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x51, 0x0a, 0x1e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x2d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x3b, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_com_itcoursee_protocol_lecture_rpc_proto_goTypes = []interface{}{
	(*AddReq)(nil),          // 0: com.itcoursee.protocol.lecture.AddReq
	(*GetReq)(nil),          // 1: com.itcoursee.protocol.lecture.GetReq
	(*GetsByCourseReq)(nil), // 2: com.itcoursee.protocol.lecture.GetsByCourseReq
	(*GetPlaybackReq)(nil),  // 3: com.itcoursee.protocol.lecture.GetPlaybackReq
	(*UpdateReq)(nil),       // 4: com.itcoursee.protocol.lecture.UpdateReq
	(*DeleteReq)(nil),       // 5: com.itcoursee.protocol.lecture.DeleteReq
	(*AddRsp)(nil),          // 6: com.itcoursee.protocol.lecture.AddRsp
	(*GetRsp)(nil),          // 7: com.itcoursee.protocol.lecture.GetRsp
	(*GetsByCourseRsp)(nil), // 8: com.itcoursee.protocol.lecture.GetsByCourseRsp
	(*GetPlaybackRsp)(nil),  // 9: com.itcoursee.protocol.lecture.GetPlaybackRsp
	(*UpdateRsp)(nil),       // 10: com.itcoursee.protocol.lecture.UpdateRsp
	(*DeleteRsp)(nil),       // 11: com.itcoursee.protocol.lecture.DeleteRsp
}
var file_com_itcoursee_protocol_lecture_rpc_proto_depIdxs = []int32{
	0,  // 0: com.itcoursee.protocol.lecture.Rpc.Add:input_type -> com.itcoursee.protocol.lecture.AddReq
	1,  // 1: com.itcoursee.protocol.lecture.Rpc.Get:input_type -> com.itcoursee.protocol.lecture.GetReq
	2,  // 2: com.itcoursee.protocol.lecture.Rpc.GetsByCourse:input_type -> com.itcoursee.protocol.lecture.GetsByCourseReq
	3,  // 3: com.itcoursee.protocol.lecture.Rpc.GetPlayback:input_type -> com.itcoursee.protocol.lecture.GetPlaybackReq
	4,  // 4: com.itcoursee.protocol.lecture.Rpc.Update:input_type -> com.itcoursee.protocol.lecture.UpdateReq
	5,  // 5: com.itcoursee.protocol.lecture.Rpc.Delete:input_type -> com.itcoursee.protocol.lecture.DeleteReq
	6,  // 6: com.itcoursee.protocol.lecture.Rpc.Add:output_type -> com.itcoursee.protocol.lecture.AddRsp
	7,  // 7: com.itcoursee.protocol.lecture.Rpc.Get:output_type -> com.itcoursee.protocol.lecture.GetRsp
	8,  // 8: com.itcoursee.protocol.lecture.Rpc.GetsByCourse:output_type -> com.itcoursee.protocol.lecture.GetsByCourseRsp
	9,  // 9: com.itcoursee.protocol.lecture.Rpc.GetPlayback:output_type -> com.itcoursee.protocol.lecture.GetPlaybackRsp
	10, // 10: com.itcoursee.protocol.lecture.Rpc.Update:output_type -> com.itcoursee.protocol.lecture.UpdateRsp
	11, // 11: com.itcoursee.protocol.lecture.Rpc.Delete:output_type -> com.itcoursee.protocol.lecture.DeleteRsp
	6,  // [6:12] is the sub-list for method output_type
	0,  // [0:6] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_protocol_lecture_rpc_proto_init() }
func file_com_itcoursee_protocol_lecture_rpc_proto_init() {
	if File_com_itcoursee_protocol_lecture_rpc_proto != nil {
		return
	}
	file_com_itcoursee_protocol_lecture_add_proto_init()
	file_com_itcoursee_protocol_lecture_course_proto_init()
	file_com_itcoursee_protocol_lecture_delete_proto_init()
	file_com_itcoursee_protocol_lecture_get_proto_init()
	file_com_itcoursee_protocol_lecture_playback_proto_init()
	file_com_itcoursee_protocol_lecture_update_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_protocol_lecture_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_itcoursee_protocol_lecture_rpc_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_protocol_lecture_rpc_proto_depIdxs,
	}.Build()
	File_com_itcoursee_protocol_lecture_rpc_proto = out.File
	file_com_itcoursee_protocol_lecture_rpc_proto_rawDesc = nil
	file_com_itcoursee_protocol_lecture_rpc_proto_goTypes = nil
	file_com_itcoursee_protocol_lecture_rpc_proto_depIdxs = nil
}
