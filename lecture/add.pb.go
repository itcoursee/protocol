// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/lecture/add.proto

package lecture

import (
	core "github.com/itcoursee/core"
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

// 创建讲次请求
type AddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程编号
	// @gotags: json:"courseId,string" query:"courseId" param:"courseId" form:"courseId" validate:"required"
	CourseId int64 `protobuf:"varint,5,opt,name=course_id,json=courseId,proto3" json:"courseId,string" query:"courseId" param:"courseId" form:"courseId" validate:"required"`
	// 章节编号
	// @gotags: json:"chapterId,string"
	ChapterId int64 `protobuf:"varint,6,opt,name=chapter_id,json=chapterId,proto3" json:"chapterId,string"`
	// 类型
	// @gotags: validate:"required,oneof=10 20 30"
	Type core.Type `protobuf:"varint,7,opt,name=type,proto3,enum=itcoursee.core.Type" json:"type,omitempty" validate:"required,oneof=10 20 30"`
	// 讲次名称
	// @gotags: validate:"required,max=255"
	Name string `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty" validate:"required,max=255"`
	// 排序
	Sequence int32 `protobuf:"varint,11,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *AddReq) Reset() {
	*x = AddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_lecture_add_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReq) ProtoMessage() {}

func (x *AddReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_lecture_add_proto_msgTypes[0]
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
	return file_itcoursee_protocol_lecture_add_proto_rawDescGZIP(), []int{0}
}

func (x *AddReq) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *AddReq) GetChapterId() int64 {
	if x != nil {
		return x.ChapterId
	}
	return 0
}

func (x *AddReq) GetType() core.Type {
	if x != nil {
		return x.Type
	}
	return core.Type_TYPE_UNSPECIFIED
}

func (x *AddReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AddReq) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

// 创建讲次返回
type AddRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 讲次
	Lecture *vo.Lecture `protobuf:"bytes,5,opt,name=lecture,proto3" json:"lecture,omitempty"`
}

func (x *AddRsp) Reset() {
	*x = AddRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_lecture_add_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRsp) ProtoMessage() {}

func (x *AddRsp) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_lecture_add_proto_msgTypes[1]
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
	return file_itcoursee_protocol_lecture_add_proto_rawDescGZIP(), []int{1}
}

func (x *AddRsp) GetLecture() *vo.Lecture {
	if x != nil {
		return x.Lecture
	}
	return nil
}

var File_itcoursee_protocol_lecture_add_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_lecture_add_proto_rawDesc = []byte{
	0x0a, 0x24, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x64, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x1a, 0x19, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f,
	0x72, 0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x9e, 0x01, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x22, 0x42, 0x0a, 0x06, 0x41, 0x64, 0x64, 0x52, 0x73, 0x70, 0x12, 0x38, 0x0a,
	0x07, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x07,
	0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x42, 0x51, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x3b, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_itcoursee_protocol_lecture_add_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_lecture_add_proto_rawDescData = file_itcoursee_protocol_lecture_add_proto_rawDesc
)

func file_itcoursee_protocol_lecture_add_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_lecture_add_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_lecture_add_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_lecture_add_proto_rawDescData)
	})
	return file_itcoursee_protocol_lecture_add_proto_rawDescData
}

var file_itcoursee_protocol_lecture_add_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_itcoursee_protocol_lecture_add_proto_goTypes = []interface{}{
	(*AddReq)(nil),     // 0: itcoursee.protocol.lecture.AddReq
	(*AddRsp)(nil),     // 1: itcoursee.protocol.lecture.AddRsp
	(core.Type)(0),     // 2: itcoursee.core.Type
	(*vo.Lecture)(nil), // 3: itcoursee.protocol.vo.Lecture
}
var file_itcoursee_protocol_lecture_add_proto_depIdxs = []int32{
	2, // 0: itcoursee.protocol.lecture.AddReq.type:type_name -> itcoursee.core.Type
	3, // 1: itcoursee.protocol.lecture.AddRsp.lecture:type_name -> itcoursee.protocol.vo.Lecture
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_lecture_add_proto_init() }
func file_itcoursee_protocol_lecture_add_proto_init() {
	if File_itcoursee_protocol_lecture_add_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_lecture_add_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_itcoursee_protocol_lecture_add_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_itcoursee_protocol_lecture_add_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_lecture_add_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_lecture_add_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_lecture_add_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_lecture_add_proto = out.File
	file_itcoursee_protocol_lecture_add_proto_rawDesc = nil
	file_itcoursee_protocol_lecture_add_proto_goTypes = nil
	file_itcoursee_protocol_lecture_add_proto_depIdxs = nil
}
