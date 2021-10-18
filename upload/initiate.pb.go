// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/upload/initiate.proto

package upload

import (
	vo "github.com/itcoursee/protocol/vo"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 初始化分块上传请求
type InitiateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @inject_tag: json:"id,string" query:"id" param:"id" form:"id" validate:"required"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *InitiateReq) Reset() {
	*x = InitiateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_upload_initiate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateReq) ProtoMessage() {}

func (x *InitiateReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_upload_initiate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateReq.ProtoReflect.Descriptor instead.
func (*InitiateReq) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_upload_initiate_proto_rawDescGZIP(), []int{0}
}

func (x *InitiateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 初始化分块上传返回
type InitiateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分片上传编号
	// @inject_tag: json:",omitempty"
	Multipart *vo.Multipart `protobuf:"bytes,10,opt,name=multipart,proto3" json:"multipart,omitempty"`
}

func (x *InitiateRsp) Reset() {
	*x = InitiateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_upload_initiate_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitiateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitiateRsp) ProtoMessage() {}

func (x *InitiateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_upload_initiate_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitiateRsp.ProtoReflect.Descriptor instead.
func (*InitiateRsp) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_upload_initiate_proto_rawDescGZIP(), []int{1}
}

func (x *InitiateRsp) GetMultipart() *vo.Multipart {
	if x != nil {
		return x.Multipart
	}
	return nil
}

var File_itcoursee_protocol_upload_initiate_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_upload_initiate_proto_rawDesc = []byte{
	0x0a, 0x28, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x2f, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75,
	0x70, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x25, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x6d, 0x75, 0x6c,
	0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1d, 0x0a, 0x0b,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4d, 0x0a, 0x0b, 0x49,
	0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x12, 0x3e, 0x0a, 0x09, 0x6d, 0x75,
	0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x52,
	0x09, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x61, 0x72, 0x74, 0x42, 0x4e, 0x0a, 0x1d, 0x63, 0x6f,
	0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x01, 0x5a, 0x2b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x75, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x3b, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_itcoursee_protocol_upload_initiate_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_upload_initiate_proto_rawDescData = file_itcoursee_protocol_upload_initiate_proto_rawDesc
)

func file_itcoursee_protocol_upload_initiate_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_upload_initiate_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_upload_initiate_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_upload_initiate_proto_rawDescData)
	})
	return file_itcoursee_protocol_upload_initiate_proto_rawDescData
}

var file_itcoursee_protocol_upload_initiate_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_itcoursee_protocol_upload_initiate_proto_goTypes = []interface{}{
	(*InitiateReq)(nil),  // 0: itcoursee.protocol.upload.InitiateReq
	(*InitiateRsp)(nil),  // 1: itcoursee.protocol.upload.InitiateRsp
	(*vo.Multipart)(nil), // 2: itcoursee.protocol.vo.Multipart
}
var file_itcoursee_protocol_upload_initiate_proto_depIdxs = []int32{
	2, // 0: itcoursee.protocol.upload.InitiateRsp.multipart:type_name -> itcoursee.protocol.vo.Multipart
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_upload_initiate_proto_init() }
func file_itcoursee_protocol_upload_initiate_proto_init() {
	if File_itcoursee_protocol_upload_initiate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_upload_initiate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_itcoursee_protocol_upload_initiate_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitiateRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_itcoursee_protocol_upload_initiate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_upload_initiate_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_upload_initiate_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_upload_initiate_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_upload_initiate_proto = out.File
	file_itcoursee_protocol_upload_initiate_proto_rawDesc = nil
	file_itcoursee_protocol_upload_initiate_proto_goTypes = nil
	file_itcoursee_protocol_upload_initiate_proto_depIdxs = nil
}
