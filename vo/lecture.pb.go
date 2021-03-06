// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/vo/lecture.proto

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

// 课程讲次
type Lecture struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @gotags: json:"id,string,omitempty"
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 课程编号
	// @gotags: json:"courseId,string,omitempty"
	CourseId int64 `protobuf:"varint,5,opt,name=course_id,json=courseId,proto3" json:"course_id,omitempty"`
	// 章节编号
	// @gotags: json:"chapterId,string,omitempty"
	ChapterId int64 `protobuf:"varint,6,opt,name=chapter_id,json=chapterId,proto3" json:"chapter_id,omitempty"`
	// 类型
	Type core.Type `protobuf:"varint,7,opt,name=type,proto3,enum=itcoursee.core.Type" json:"type,omitempty"`
	// 名称
	Name string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	// 状态
	Status core.Status `protobuf:"varint,9,opt,name=status,proto3,enum=itcoursee.core.Status" json:"status,omitempty"`
	// 排序
	Sequence int32 `protobuf:"varint,10,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *Lecture) Reset() {
	*x = Lecture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_vo_lecture_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lecture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lecture) ProtoMessage() {}

func (x *Lecture) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_vo_lecture_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Lecture.ProtoReflect.Descriptor instead.
func (*Lecture) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_vo_lecture_proto_rawDescGZIP(), []int{0}
}

func (x *Lecture) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lecture) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

func (x *Lecture) GetChapterId() int64 {
	if x != nil {
		return x.ChapterId
	}
	return 0
}

func (x *Lecture) GetType() core.Type {
	if x != nil {
		return x.Type
	}
	return core.Type_TYPE_UNSPECIFIED
}

func (x *Lecture) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lecture) GetStatus() core.Status {
	if x != nil {
		return x.Status
	}
	return core.Status_STATUS_UNSPECIFIED
}

func (x *Lecture) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

var File_itcoursee_protocol_vo_lecture_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_vo_lecture_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x19, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x74, 0x79, 0x70,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdf, 0x01, 0x0a, 0x07, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x69, 0x74, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65,
	0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x42, 0x42, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x76, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_itcoursee_protocol_vo_lecture_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_vo_lecture_proto_rawDescData = file_itcoursee_protocol_vo_lecture_proto_rawDesc
)

func file_itcoursee_protocol_vo_lecture_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_vo_lecture_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_vo_lecture_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_vo_lecture_proto_rawDescData)
	})
	return file_itcoursee_protocol_vo_lecture_proto_rawDescData
}

var file_itcoursee_protocol_vo_lecture_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_itcoursee_protocol_vo_lecture_proto_goTypes = []interface{}{
	(*Lecture)(nil),  // 0: itcoursee.protocol.vo.Lecture
	(core.Type)(0),   // 1: itcoursee.core.Type
	(core.Status)(0), // 2: itcoursee.core.Status
}
var file_itcoursee_protocol_vo_lecture_proto_depIdxs = []int32{
	1, // 0: itcoursee.protocol.vo.Lecture.type:type_name -> itcoursee.core.Type
	2, // 1: itcoursee.protocol.vo.Lecture.status:type_name -> itcoursee.core.Status
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_vo_lecture_proto_init() }
func file_itcoursee_protocol_vo_lecture_proto_init() {
	if File_itcoursee_protocol_vo_lecture_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_vo_lecture_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Lecture); i {
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
			RawDescriptor: file_itcoursee_protocol_vo_lecture_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_vo_lecture_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_vo_lecture_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_vo_lecture_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_vo_lecture_proto = out.File
	file_itcoursee_protocol_vo_lecture_proto_rawDesc = nil
	file_itcoursee_protocol_vo_lecture_proto_goTypes = nil
	file_itcoursee_protocol_vo_lecture_proto_depIdxs = nil
}
