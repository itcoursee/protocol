// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/vo/course.proto

package vo

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

// 课程
type Course struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @gotags: json:"id,string"
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 标签，SEO使用
	Label string `protobuf:"bytes,5,opt,name=label,proto3" json:"label,omitempty"`
	// 关键字，SEO使用
	Keywords []string `protobuf:"bytes,6,rep,name=keywords,proto3" json:"keywords,omitempty"`
	// 详细描述，SEO使用
	Description string `protobuf:"bytes,7,opt,name=description,proto3" json:"description,omitempty"`
	// 名称
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	// 类型
	Type core.Type `protobuf:"varint,11,opt,name=type,proto3,enum=itcoursee.core.Type" json:"type,omitempty"`
	// 封面地址
	Cover string `protobuf:"bytes,12,opt,name=cover,proto3" json:"cover,omitempty"`
	// 方向
	// @gotags: json:"directionId,string,omitempty"
	DirectionId int64 `protobuf:"varint,13,opt,name=direction_id,json=directionId,proto3" json:"direction_id,omitempty"`
	// 分类
	// @gotags: json:"categoryId,string,omitempty"
	CategoryId int64 `protobuf:"varint,14,opt,name=category_id,json=categoryId,proto3" json:"category_id,omitempty"`
	// 阶段
	Stage core.Stage `protobuf:"varint,15,opt,name=stage,proto3,enum=itcoursee.core.Stage" json:"stage,omitempty"`
	// 创建用户
	// @gotags: json:"ownerId,string,omitempty"
	OwnerId int64 `protobuf:"varint,16,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	// 简介
	Profile string `protobuf:"bytes,20,opt,name=profile,proto3" json:"profile,omitempty"`
	// 须知
	Notice string `protobuf:"bytes,21,opt,name=notice,proto3" json:"notice,omitempty"`
	// 价格
	Price *core.Price `protobuf:"bytes,22,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_vo_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_vo_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Course.ProtoReflect.Descriptor instead.
func (*Course) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_vo_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Course) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Course) GetKeywords() []string {
	if x != nil {
		return x.Keywords
	}
	return nil
}

func (x *Course) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Course) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Course) GetType() core.Type {
	if x != nil {
		return x.Type
	}
	return core.Type_TYPE_UNSPECIFIED
}

func (x *Course) GetCover() string {
	if x != nil {
		return x.Cover
	}
	return ""
}

func (x *Course) GetDirectionId() int64 {
	if x != nil {
		return x.DirectionId
	}
	return 0
}

func (x *Course) GetCategoryId() int64 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *Course) GetStage() core.Stage {
	if x != nil {
		return x.Stage
	}
	return core.Stage_STAGE_UNSPECIFIED
}

func (x *Course) GetOwnerId() int64 {
	if x != nil {
		return x.OwnerId
	}
	return 0
}

func (x *Course) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *Course) GetNotice() string {
	if x != nil {
		return x.Notice
	}
	return ""
}

func (x *Course) GetPrice() *core.Price {
	if x != nil {
		return x.Price
	}
	return nil
}

var File_itcoursee_protocol_vo_course_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_vo_course_proto_rawDesc = []byte{
	0x0a, 0x22, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x1a, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1a, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xab,
	0x03, 0x0a, 0x06, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x08, 0x6b, 0x65, 0x79, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x14, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69,
	0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x67, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x67, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61,
	0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63,
	0x65, 0x18, 0x15, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x69, 0x63, 0x65, 0x12,
	0x2b, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x16, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x4b, 0x0a, 0x22,
	0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_itcoursee_protocol_vo_course_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_vo_course_proto_rawDescData = file_itcoursee_protocol_vo_course_proto_rawDesc
)

func file_itcoursee_protocol_vo_course_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_vo_course_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_vo_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_vo_course_proto_rawDescData)
	})
	return file_itcoursee_protocol_vo_course_proto_rawDescData
}

var file_itcoursee_protocol_vo_course_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_itcoursee_protocol_vo_course_proto_goTypes = []interface{}{
	(*Course)(nil),     // 0: itcoursee.protocol.vo.Course
	(core.Type)(0),     // 1: itcoursee.core.Type
	(core.Stage)(0),    // 2: itcoursee.core.Stage
	(*core.Price)(nil), // 3: itcoursee.core.Price
}
var file_itcoursee_protocol_vo_course_proto_depIdxs = []int32{
	1, // 0: itcoursee.protocol.vo.Course.type:type_name -> itcoursee.core.Type
	2, // 1: itcoursee.protocol.vo.Course.stage:type_name -> itcoursee.core.Stage
	3, // 2: itcoursee.protocol.vo.Course.price:type_name -> itcoursee.core.Price
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_vo_course_proto_init() }
func file_itcoursee_protocol_vo_course_proto_init() {
	if File_itcoursee_protocol_vo_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_vo_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Course); i {
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
			RawDescriptor: file_itcoursee_protocol_vo_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_vo_course_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_vo_course_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_vo_course_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_vo_course_proto = out.File
	file_itcoursee_protocol_vo_course_proto_rawDesc = nil
	file_itcoursee_protocol_vo_course_proto_goTypes = nil
	file_itcoursee_protocol_vo_course_proto_depIdxs = nil
}
