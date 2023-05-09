// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/protocol/tag/add.proto

package tag

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

// 创建标签请求
type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 标签，只能是英文小写字母，SEO使用
	// @gotags: validate:"required,alpha,lowercase,max=16"
	Label string `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty" validate:"required,alpha,lowercase,max=16"`
	// 关键字
	// @gotags: validate:"required"
	Keywords []string `protobuf:"bytes,4,rep,name=keywords,proto3" json:"keywords,omitempty" validate:"required"`
	// 详细描述
	// @gotags: validate:"required,max=1024"
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty" validate:"required,max=1024"`
	// 名称
	// @gotags: validate:"required,max=64"
	Name string `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty" validate:"required,max=64"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_protocol_tag_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_protocol_tag_add_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReq.ProtoReflect.Descriptor instead.
func (*AddReq) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_protocol_tag_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *AddReq) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *AddReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *AddReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 创建标签返回
type AddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddRsp) Reset() {
	*x = AddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_protocol_tag_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRsp) ProtoMessage() {}

func (x *AddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_protocol_tag_add_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRsp.ProtoReflect.Descriptor instead.
func (*AddRsp) Descriptor() ([]byte, []int) {
	return file_com_itcoursee_protocol_tag_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddRsp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_com_itcoursee_protocol_tag_add_proto protoreflect.FileDescriptor

var file_com_itcoursee_protocol_tag_add_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x61, 0x67, 0x2f, 0x61, 0x64, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x74,
	0x61, 0x67, 0x22, 0x70, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05,
	0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x18, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x45,
	0x0a, 0x1a, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x74, 0x61, 0x67, 0x50, 0x01, 0x5a, 0x25,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x61,
	0x67, 0x3b, 0x74, 0x61, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_protocol_tag_add_proto_rawDescOnce sync.Once
	file_com_itcoursee_protocol_tag_add_proto_rawDescData = file_com_itcoursee_protocol_tag_add_proto_rawDesc
)

func file_com_itcoursee_protocol_tag_add_proto_rawDescGZIP() []byte {
	file_com_itcoursee_protocol_tag_add_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_protocol_tag_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_protocol_tag_add_proto_rawDescData)
	})
	return file_com_itcoursee_protocol_tag_add_proto_rawDescData
}

var file_com_itcoursee_protocol_tag_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_protocol_tag_add_proto_goTypes = []interface{}{
	(*AddReq)(nil), // 0: com.itcoursee.protocol.tag.AddReq
	(*AddRsp)(nil), // 1: com.itcoursee.protocol.tag.AddRsp
}
var file_com_itcoursee_protocol_tag_add_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_com_itcoursee_protocol_tag_add_proto_init() }
func file_com_itcoursee_protocol_tag_add_proto_init() {
	if File_com_itcoursee_protocol_tag_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_protocol_tag_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReq); i {
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
		file_com_itcoursee_protocol_tag_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRsp); i {
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
			RawDescriptor: file_com_itcoursee_protocol_tag_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_protocol_tag_add_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_protocol_tag_add_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_protocol_tag_add_proto_msgTypes,
	}.Build()
	File_com_itcoursee_protocol_tag_add_proto = out.File
	file_com_itcoursee_protocol_tag_add_proto_rawDesc = nil
	file_com_itcoursee_protocol_tag_add_proto_goTypes = nil
	file_com_itcoursee_protocol_tag_add_proto_depIdxs = nil
}
