// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/protocol/upload/rpc.proto

package course

import (
	upload "github.com/itcoursee/protocol/upload"
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

var File_com_itcoursee_protocol_upload_rpc_proto protoreflect.FileDescriptor

var file_com_itcoursee_protocol_upload_rpc_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f,
	0x72, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x61, 0x62, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2c, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x75, 0x72, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xdf, 0x03, 0x0a, 0x03, 0x52, 0x70, 0x63, 0x12, 0x77,
	0x0a, 0x08, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52,
	0x73, 0x70, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x22, 0x08, 0x2f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x73, 0x3a, 0x01, 0x2a, 0x12, 0x6f, 0x0a, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x25,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x55, 0x72, 0x6c, 0x52, 0x73, 0x70, 0x22, 0x1a, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x14, 0x22, 0x12, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x12, 0x7c, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x70,
	0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x1a, 0x2a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64,
	0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x18, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x12, 0x1a, 0x0d, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2f, 0x7b,
	0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x70, 0x0a, 0x05, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x12,
	0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e,
	0x41, 0x62, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2e, 0x41, 0x62, 0x6f, 0x72, 0x74, 0x52, 0x73,
	0x70, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x2a, 0x0d, 0x2f, 0x75, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x4e, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_com_itcoursee_protocol_upload_rpc_proto_goTypes = []interface{}{
	(*upload.InitiateReq)(nil), // 0: com.itcoursee.protocol.upload.InitiateReq
	(*upload.UrlReq)(nil),      // 1: com.itcoursee.protocol.upload.UrlReq
	(*upload.CompleteReq)(nil), // 2: com.itcoursee.protocol.upload.CompleteReq
	(*upload.AbortReq)(nil),    // 3: com.itcoursee.protocol.upload.AbortReq
	(*upload.InitiateRsp)(nil), // 4: com.itcoursee.protocol.upload.InitiateRsp
	(*upload.UrlRsp)(nil),      // 5: com.itcoursee.protocol.upload.UrlRsp
	(*upload.CompleteRsp)(nil), // 6: com.itcoursee.protocol.upload.CompleteRsp
	(*upload.AbortRsp)(nil),    // 7: com.itcoursee.protocol.upload.AbortRsp
}
var file_com_itcoursee_protocol_upload_rpc_proto_depIdxs = []int32{
	0, // 0: com.itcoursee.protocol.upload.Rpc.Initiate:input_type -> com.itcoursee.protocol.upload.InitiateReq
	1, // 1: com.itcoursee.protocol.upload.Rpc.Url:input_type -> com.itcoursee.protocol.upload.UrlReq
	2, // 2: com.itcoursee.protocol.upload.Rpc.Complete:input_type -> com.itcoursee.protocol.upload.CompleteReq
	3, // 3: com.itcoursee.protocol.upload.Rpc.Abort:input_type -> com.itcoursee.protocol.upload.AbortReq
	4, // 4: com.itcoursee.protocol.upload.Rpc.Initiate:output_type -> com.itcoursee.protocol.upload.InitiateRsp
	5, // 5: com.itcoursee.protocol.upload.Rpc.Url:output_type -> com.itcoursee.protocol.upload.UrlRsp
	6, // 6: com.itcoursee.protocol.upload.Rpc.Complete:output_type -> com.itcoursee.protocol.upload.CompleteRsp
	7, // 7: com.itcoursee.protocol.upload.Rpc.Abort:output_type -> com.itcoursee.protocol.upload.AbortRsp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_protocol_upload_rpc_proto_init() }
func file_com_itcoursee_protocol_upload_rpc_proto_init() {
	if File_com_itcoursee_protocol_upload_rpc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_protocol_upload_rpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_itcoursee_protocol_upload_rpc_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_protocol_upload_rpc_proto_depIdxs,
	}.Build()
	File_com_itcoursee_protocol_upload_rpc_proto = out.File
	file_com_itcoursee_protocol_upload_rpc_proto_rawDesc = nil
	file_com_itcoursee_protocol_upload_rpc_proto_goTypes = nil
	file_com_itcoursee_protocol_upload_rpc_proto_depIdxs = nil
}
