// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/vo/object.proto

package vo

import (
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

// 对象
type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 对象键
	// @inject_tag: validate:"required"
	Key string `protobuf:"bytes,10,opt,name=key,proto3" json:"key,omitempty"`
	// 上传后返回的ETag头信息
	// @inject_tag: validate:"required"
	Etag string `protobuf:"bytes,20,opt,name=etag,proto3" json:"etag,omitempty"`
	// 大小
	// @inject_tag: validate:"required"
	Size int64 `protobuf:"varint,30,opt,name=size,proto3" json:"size,omitempty"`
	// 分片编号
	// @inject_tag: validate:"required"
	Part int32 `protobuf:"varint,40,opt,name=part,proto3" json:"part,omitempty"`
	// 版本
	Version string `protobuf:"bytes,50,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_vo_object_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_vo_object_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_vo_object_proto_rawDescGZIP(), []int{0}
}

func (x *Object) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Object) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *Object) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Object) GetPart() int32 {
	if x != nil {
		return x.Part
	}
	return 0
}

func (x *Object) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

var File_itcoursee_protocol_vo_object_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_vo_object_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x22, 0x70, 0x0a, 0x06, 0x4f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x74, 0x61, 0x67, 0x18,
	0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x72, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70,
	0x61, 0x72, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x32,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x42, 0x0a,
	0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itcoursee_protocol_vo_object_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_vo_object_proto_rawDescData = file_itcoursee_protocol_vo_object_proto_rawDesc
)

func file_itcoursee_protocol_vo_object_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_vo_object_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_vo_object_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_vo_object_proto_rawDescData)
	})
	return file_itcoursee_protocol_vo_object_proto_rawDescData
}

var file_itcoursee_protocol_vo_object_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_itcoursee_protocol_vo_object_proto_goTypes = []interface{}{
	(*Object)(nil), // 0: itcoursee.protocol.vo.Object
}
var file_itcoursee_protocol_vo_object_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_vo_object_proto_init() }
func file_itcoursee_protocol_vo_object_proto_init() {
	if File_itcoursee_protocol_vo_object_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_vo_object_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
			RawDescriptor: file_itcoursee_protocol_vo_object_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_vo_object_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_vo_object_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_vo_object_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_vo_object_proto = out.File
	file_itcoursee_protocol_vo_object_proto_rawDesc = nil
	file_itcoursee_protocol_vo_object_proto_goTypes = nil
	file_itcoursee_protocol_vo_object_proto_depIdxs = nil
}
