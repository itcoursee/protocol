// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/course/update.proto

package course

import (
	core "github.com/itcoursee/core"
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

// 更新课程请求
type UpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @inject_tag: json:"id,string" query:"id" param:"id" form:"id" validate:"required"
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 标签，SEO使用
	// @inject_tag: validate:"required,alpha,lowercase,max=16"
	Label string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	// 关键字，SEO使用
	// @inject_tag: validate:"required"
	Keywords []string `protobuf:"bytes,6,rep,name=keywords,proto3" json:"keywords,omitempty"`
	// 详细描述，SEO使用
	// @inject_tag: validate:"required,max=1024"
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// 名称
	// @inject_tag: validate:"required,max=64"
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// 类型
	// @inject_tag: validate:"required,oneof=10 20 30"
	Type int32 `protobuf:"varint,11,opt,name=type,proto3" json:"type,omitempty"`
	// 方向
	DirectionId int64 `protobuf:"varint,12,opt,name=direction_id,json=directionId,proto3" json:"direction_id,omitempty"`
	// 分类
	// @inject_tag: json:"categoryId,string"
	CategoryId int64 `protobuf:"varint,13,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	// 阶段
	Stage int32 `protobuf:"varint,14,opt,name=stage,proto3" json:"stage,omitempty"`
	// 创建用户
	// @inject_tag: json:"creatorId,string"
	CreatorId int64 `protobuf:"varint,15,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
	// 简介
	// @inject_tag: validate:"required,max=1024"
	Profile string `protobuf:"bytes,20,opt,name=profile,proto3" json:"profile,omitempty"`
	// 须知
	// @inject_tag: validate:"required,max=1024"
	Notice string `protobuf:"bytes,21,opt,name=notice,proto3" json:"notice,omitempty"`
	// 价格
	Price *core.Price `protobuf:"bytes,22,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *UpdateReq) Reset() {
	*x = UpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_course_update_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateReq) ProtoMessage() {}

func (x *UpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_course_update_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateReq.ProtoReflect.Descriptor instead.
func (*UpdateReq) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_course_update_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateReq) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *UpdateReq) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *UpdateReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateReq) GetType() int32 {
	if x != nil {
		return x.Type
	}
	return 0
}

func (x *UpdateReq) GetDirectionId() int64 {
	if x != nil {
		return x.DirectionId
	}
	return 0
}

func (x *UpdateReq) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *UpdateReq) GetStage() int32 {
	if x != nil {
		return x.Stage
	}
	return 0
}

func (x *UpdateReq) GetCreatorId() int64 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

func (x *UpdateReq) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *UpdateReq) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

func (x *UpdateReq) GetPrice() *core.Price {
	if x != nil {
		return x.Price
	}
	return nil
}

// 更新课程返回
type UpdateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateRsp) Reset() {
	*x = UpdateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_course_update_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRsp) ProtoMessage() {}

func (x *UpdateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_course_update_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRsp.ProtoReflect.Descriptor instead.
func (*UpdateRsp) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_course_update_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateRsp) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_itcoursee_protocol_course_update_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_course_update_proto_rawDesc = []byte{
	0x0a, 0x26, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x1a, 0x1a, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xef, 0x02, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18,
	0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x72,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f,
	0x74, 0x69, 0x63, 0x65, 0x12, 0x2b, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x22, 0x1b, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x42, 0x4e,
	0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50,
	0x01, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itcoursee_protocol_course_update_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_course_update_proto_rawDescData = file_itcoursee_protocol_course_update_proto_rawDesc
)

func file_itcoursee_protocol_course_update_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_course_update_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_course_update_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_course_update_proto_rawDescData)
	})
	return file_itcoursee_protocol_course_update_proto_rawDescData
}

var file_itcoursee_protocol_course_update_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_itcoursee_protocol_course_update_proto_goTypes = []interface{}{
	(*UpdateReq)(nil),  // 0: itcoursee.protocol.course.UpdateReq
	(*UpdateRsp)(nil),  // 1: itcoursee.protocol.course.UpdateRsp
	(*core.Price)(nil), // 2: itcoursee.core.Price
}
var file_itcoursee_protocol_course_update_proto_depIdxs = []int32{
	2, // 0: itcoursee.protocol.course.UpdateReq.price:type_name -> itcoursee.core.Price
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_course_update_proto_init() }
func file_itcoursee_protocol_course_update_proto_init() {
	if File_itcoursee_protocol_course_update_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_course_update_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateReq); i {
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
		file_itcoursee_protocol_course_update_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRsp); i {
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
			RawDescriptor: file_itcoursee_protocol_course_update_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_course_update_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_course_update_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_course_update_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_course_update_proto = out.File
	file_itcoursee_protocol_course_update_proto_rawDesc = nil
	file_itcoursee_protocol_course_update_proto_goTypes = nil
	file_itcoursee_protocol_course_update_proto_depIdxs = nil
}
