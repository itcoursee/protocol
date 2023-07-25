// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/protocol/vo/course.proto

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
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// 名称
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// 类型
	Type core.Type `protobuf:"varint,4,opt,name=type,proto3,enum=com.itcoursee.core.Type" json:"type,omitempty"`
	// 观众
	Audience core.Audience `protobuf:"varint,5,opt,name=audience,proto3,enum=com.itcoursee.core.Audience" json:"audience,omitempty"`
	// 状态
	Status core.Status `protobuf:"varint,6,opt,name=status,proto3,enum=com.itcoursee.core.Status" json:"status,omitempty"`
	// 简介
	Profile string `protobuf:"bytes,7,opt,name=profile,proto3" json:"profile,omitempty"`
	// 价格
	Price *core.Price `protobuf:"bytes,8,opt,name=price,proto3" json:"price,omitempty"`
	// 封面地址列表
	// ! 通过字段筛选返回
	// @gotags: str:"-"
	Covers map[int64]string `protobuf:"bytes,10,rep,name=covers,proto3" json:"covers,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3" str:"-"`
	// 标签
	// ! 通过字段筛选返回
	// @gotags: str:"-"
	Tags []*Tag `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty" str:"-"`
	// 创建用户
	Owner int64 `protobuf:"varint,12,opt,name=owner,proto3" json:"owner,omitempty"`
	// 分类
	// ! 通过字段筛选返回
	Category int64 `protobuf:"varint,13,opt,name=category,proto3" json:"category,omitempty"`
	// 方向
	// ! 通过字段筛选返回
	Direction int64 `protobuf:"varint,14,opt,name=direction,proto3" json:"direction,omitempty"`
}

func (x *Course) Reset() {
	*x = Course{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_protocol_vo_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Course) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Course) ProtoMessage() {}

func (x *Course) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_protocol_vo_course_proto_msgTypes[0]
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
	return file_com_itcoursee_protocol_vo_course_proto_rawDescGZIP(), []int{0}
}

func (x *Course) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
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
	return core.Type(0)
}

func (x *Course) GetAudience() core.Audience {
	if x != nil {
		return x.Audience
	}
	return core.Audience(0)
}

func (x *Course) GetStatus() core.Status {
	if x != nil {
		return x.Status
	}
	return core.Status(0)
}

func (x *Course) GetProfile() string {
	if x != nil {
		return x.Profile
	}
	return ""
}

func (x *Course) GetPrice() *core.Price {
	if x != nil {
		return x.Price
	}
	return nil
}

func (x *Course) GetCovers() map[int64]string {
	if x != nil {
		return x.Covers
	}
	return nil
}

func (x *Course) GetTags() []*Tag {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Course) GetOwner() int64 {
	if x != nil {
		return x.Owner
	}
	return 0
}

func (x *Course) GetCategory() int64 {
	if x != nil {
		return x.Category
	}
	return 0
}

func (x *Course) GetDirection() int64 {
	if x != nil {
		return x.Direction
	}
	return 0
}

var File_com_itcoursee_protocol_vo_course_proto protoreflect.FileDescriptor

var file_com_itcoursee_protocol_vo_course_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x76, 0x6f, 0x1a, 0x21, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x6f, 0x2f, 0x74, 0x61, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x99, 0x04, 0x0a, 0x06,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x08, 0x61, 0x75, 0x64, 0x69,
	0x65, 0x6e, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x63, 0x6f, 0x6d,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e,
	0x41, 0x75, 0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x08, 0x61, 0x75, 0x64, 0x69, 0x65, 0x6e,
	0x63, 0x65, 0x12, 0x32, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x12, 0x2f, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x2e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63,
	0x65, 0x12, 0x45, 0x0a, 0x06, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2d, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x06, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x12, 0x32, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e,
	0x76, 0x6f, 0x2e, 0x54, 0x61, 0x67, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0d,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0e, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x39, 0x0a, 0x0b,
	0x43, 0x6f, 0x76, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x4b, 0x0a, 0x22, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x50, 0x01, 0x5a,
	0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76,
	0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_protocol_vo_course_proto_rawDescOnce sync.Once
	file_com_itcoursee_protocol_vo_course_proto_rawDescData = file_com_itcoursee_protocol_vo_course_proto_rawDesc
)

func file_com_itcoursee_protocol_vo_course_proto_rawDescGZIP() []byte {
	file_com_itcoursee_protocol_vo_course_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_protocol_vo_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_protocol_vo_course_proto_rawDescData)
	})
	return file_com_itcoursee_protocol_vo_course_proto_rawDescData
}

var file_com_itcoursee_protocol_vo_course_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_com_itcoursee_protocol_vo_course_proto_goTypes = []interface{}{
	(*Course)(nil),     // 0: com.itcoursee.protocol.vo.Course
	nil,                // 1: com.itcoursee.protocol.vo.Course.CoversEntry
	(core.Type)(0),     // 2: com.itcoursee.core.Type
	(core.Audience)(0), // 3: com.itcoursee.core.Audience
	(core.Status)(0),   // 4: com.itcoursee.core.Status
	(*core.Price)(nil), // 5: com.itcoursee.core.Price
	(*Tag)(nil),        // 6: com.itcoursee.protocol.vo.Tag
}
var file_com_itcoursee_protocol_vo_course_proto_depIdxs = []int32{
	2, // 0: com.itcoursee.protocol.vo.Course.type:type_name -> com.itcoursee.core.Type
	3, // 1: com.itcoursee.protocol.vo.Course.audience:type_name -> com.itcoursee.core.Audience
	4, // 2: com.itcoursee.protocol.vo.Course.status:type_name -> com.itcoursee.core.Status
	5, // 3: com.itcoursee.protocol.vo.Course.price:type_name -> com.itcoursee.core.Price
	1, // 4: com.itcoursee.protocol.vo.Course.covers:type_name -> com.itcoursee.protocol.vo.Course.CoversEntry
	6, // 5: com.itcoursee.protocol.vo.Course.tags:type_name -> com.itcoursee.protocol.vo.Tag
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_com_itcoursee_protocol_vo_course_proto_init() }
func file_com_itcoursee_protocol_vo_course_proto_init() {
	if File_com_itcoursee_protocol_vo_course_proto != nil {
		return
	}
	file_com_itcoursee_protocol_vo_tag_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_protocol_vo_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_com_itcoursee_protocol_vo_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_protocol_vo_course_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_protocol_vo_course_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_protocol_vo_course_proto_msgTypes,
	}.Build()
	File_com_itcoursee_protocol_vo_course_proto = out.File
	file_com_itcoursee_protocol_vo_course_proto_rawDesc = nil
	file_com_itcoursee_protocol_vo_course_proto_goTypes = nil
	file_com_itcoursee_protocol_vo_course_proto_depIdxs = nil
}
