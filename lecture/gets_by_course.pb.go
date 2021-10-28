// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/lecture/gets_by_course.proto

package lecture

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

// 按课程获得所有讲次请求
type GetsByCourseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程编号
	// @gotags: json:"courseId,string" query:"courseId" param:"courseId" form:"courseId" validate:"required"
	CourseId int64 `protobuf:"varint,1,opt,name=course_id,json=courseId,proto3" json:"courseId,string" query:"courseId" param:"courseId" form:"courseId" validate:"required"`
}

func (x *GetsByCourseReq) Reset() {
	*x = GetsByCourseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_lecture_gets_by_course_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByCourseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByCourseReq) ProtoMessage() {}

func (x *GetsByCourseReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_lecture_gets_by_course_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByCourseReq.ProtoReflect.Descriptor instead.
func (*GetsByCourseReq) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescGZIP(), []int{0}
}

func (x *GetsByCourseReq) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

// 按课程获得所有讲次响应
type GetsByCourseRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据
	Lectures []*vo.Lecture `protobuf:"bytes,10,rep,name=lectures,proto3" json:"lectures,omitempty"`
}

func (x *GetsByCourseRsp) Reset() {
	*x = GetsByCourseRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_lecture_gets_by_course_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByCourseRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByCourseRsp) ProtoMessage() {}

func (x *GetsByCourseRsp) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_lecture_gets_by_course_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByCourseRsp.ProtoReflect.Descriptor instead.
func (*GetsByCourseRsp) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescGZIP(), []int{1}
}

func (x *GetsByCourseRsp) GetLectures() []*vo.Lecture {
	if x != nil {
		return x.Lectures
	}
	return nil
}

var File_itcoursee_protocol_lecture_gets_by_course_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_lecture_gets_by_course_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x67, 0x65, 0x74,
	0x73, 0x5f, 0x62, 0x79, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1a, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x1a, 0x23, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x49, 0x64, 0x22, 0x4d, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x73, 0x42, 0x79, 0x43, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x52, 0x73, 0x70, 0x12, 0x3a, 0x0a, 0x08, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e,
	0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x08, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x42, 0x51, 0x0a, 0x1e, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x3b, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescData = file_itcoursee_protocol_lecture_gets_by_course_proto_rawDesc
)

func file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescData)
	})
	return file_itcoursee_protocol_lecture_gets_by_course_proto_rawDescData
}

var file_itcoursee_protocol_lecture_gets_by_course_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_itcoursee_protocol_lecture_gets_by_course_proto_goTypes = []interface{}{
	(*GetsByCourseReq)(nil), // 0: itcoursee.protocol.lecture.GetsByCourseReq
	(*GetsByCourseRsp)(nil), // 1: itcoursee.protocol.lecture.GetsByCourseRsp
	(*vo.Lecture)(nil),      // 2: itcoursee.protocol.vo.Lecture
}
var file_itcoursee_protocol_lecture_gets_by_course_proto_depIdxs = []int32{
	2, // 0: itcoursee.protocol.lecture.GetsByCourseRsp.lectures:type_name -> itcoursee.protocol.vo.Lecture
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_lecture_gets_by_course_proto_init() }
func file_itcoursee_protocol_lecture_gets_by_course_proto_init() {
	if File_itcoursee_protocol_lecture_gets_by_course_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_lecture_gets_by_course_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByCourseReq); i {
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
		file_itcoursee_protocol_lecture_gets_by_course_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByCourseRsp); i {
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
			RawDescriptor: file_itcoursee_protocol_lecture_gets_by_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_lecture_gets_by_course_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_lecture_gets_by_course_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_lecture_gets_by_course_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_lecture_gets_by_course_proto = out.File
	file_itcoursee_protocol_lecture_gets_by_course_proto_rawDesc = nil
	file_itcoursee_protocol_lecture_gets_by_course_proto_goTypes = nil
	file_itcoursee_protocol_lecture_gets_by_course_proto_depIdxs = nil
}
