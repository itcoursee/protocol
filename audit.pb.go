// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.9
// source: com/itcoursee/protocol/audit.proto

package protocol

import (
	audit "github.com/itcoursee/core/audit"
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

var File_com_itcoursee_protocol_audit_proto protoreflect.FileDescriptor

var file_com_itcoursee_protocol_audit_proto_rawDesc = []byte{
	0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x63, 0x6f, 0x6d, 0x2f,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61,
	0x75, 0x64, 0x69, 0x74, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32, 0x94, 0x01, 0x0a, 0x05, 0x41, 0x75, 0x64, 0x69, 0x74, 0x12, 0x8a, 0x01, 0x0a, 0x0c,
	0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x12, 0x29, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x61, 0x75, 0x64, 0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x52, 0x65, 0x71, 0x1a, 0x29, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x64,
	0x69, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x52,
	0x73, 0x70, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x62, 0x06, 0x61, 0x75, 0x64, 0x69,
	0x74, 0x73, 0x12, 0x14, 0x2f, 0x72, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x74, 0x73, 0x42, 0x42, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x50, 0x01, 0x5a, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_com_itcoursee_protocol_audit_proto_goTypes = []interface{}{
	(*audit.GetsByReviewReq)(nil), // 0: com.itcoursee.core.audit.GetsByReviewReq
	(*audit.GetsByReviewRsp)(nil), // 1: com.itcoursee.core.audit.GetsByReviewRsp
}
var file_com_itcoursee_protocol_audit_proto_depIdxs = []int32{
	0, // 0: com.itcoursee.protocol.Audit.GetsByReview:input_type -> com.itcoursee.core.audit.GetsByReviewReq
	1, // 1: com.itcoursee.protocol.Audit.GetsByReview:output_type -> com.itcoursee.core.audit.GetsByReviewRsp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_protocol_audit_proto_init() }
func file_com_itcoursee_protocol_audit_proto_init() {
	if File_com_itcoursee_protocol_audit_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_com_itcoursee_protocol_audit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_com_itcoursee_protocol_audit_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_protocol_audit_proto_depIdxs,
	}.Build()
	File_com_itcoursee_protocol_audit_proto = out.File
	file_com_itcoursee_protocol_audit_proto_rawDesc = nil
	file_com_itcoursee_protocol_audit_proto_goTypes = nil
	file_com_itcoursee_protocol_audit_proto_depIdxs = nil
}
