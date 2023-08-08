// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.18.1
// source: com/itcoursee/protocol/vo/lecture.proto

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
	Id int64 `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	// 课程编号
	Course int64 `protobuf:"varint,5,opt,name=course,proto3" json:"course,omitempty"`
	// 章节编号
	Chapter int64 `protobuf:"varint,6,opt,name=chapter,proto3" json:"chapter,omitempty"`
	// 类型
	Type core.Type `protobuf:"varint,7,opt,name=type,proto3,enum=com.itcoursee.core.Type" json:"type,omitempty"`
	// 名称
	Name string `protobuf:"bytes,8,opt,name=name,proto3" json:"name,omitempty"`
	// 状态
	Status core.LectureStatus `protobuf:"varint,9,opt,name=status,proto3,enum=com.itcoursee.core.LectureStatus" json:"status,omitempty"`
	// 排序
	Sequence int32 `protobuf:"varint,10,opt,name=sequence,proto3" json:"sequence,omitempty"`
}

func (x *Lecture) Reset() {
	*x = Lecture{}
	if protoimpl.UnsafeEnabled {
		mi := &file_com_itcoursee_protocol_vo_lecture_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Lecture) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Lecture) ProtoMessage() {}

func (x *Lecture) ProtoReflect() protoreflect.Message {
	mi := &file_com_itcoursee_protocol_vo_lecture_proto_msgTypes[0]
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
	return file_com_itcoursee_protocol_vo_lecture_proto_rawDescGZIP(), []int{0}
}

func (x *Lecture) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Lecture) GetCourse() int64 {
	if x != nil {
		return x.Course
	}
	return 0
}

func (x *Lecture) GetChapter() int64 {
	if x != nil {
		return x.Chapter
	}
	return 0
}

func (x *Lecture) GetType() core.Type {
	if x != nil {
		return x.Type
	}
	return core.Type(0)
}

func (x *Lecture) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Lecture) GetStatus() core.LectureStatus {
	if x != nil {
		return x.Status
	}
	return core.LectureStatus(0)
}

func (x *Lecture) GetSequence() int32 {
	if x != nil {
		return x.Sequence
	}
	return 0
}

var File_com_itcoursee_protocol_vo_lecture_proto protoreflect.FileDescriptor

var file_com_itcoursee_protocol_vo_lecture_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x6c, 0x65, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2e, 0x76, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a,
	0x07, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x63, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x63,
	0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x73, 0x65, 0x71, 0x75, 0x65,
	0x6e, 0x63, 0x65, 0x42, 0x42, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f,
	0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x2f, 0x76, 0x6f, 0x3b, 0x76, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_com_itcoursee_protocol_vo_lecture_proto_rawDescOnce sync.Once
	file_com_itcoursee_protocol_vo_lecture_proto_rawDescData = file_com_itcoursee_protocol_vo_lecture_proto_rawDesc
)

func file_com_itcoursee_protocol_vo_lecture_proto_rawDescGZIP() []byte {
	file_com_itcoursee_protocol_vo_lecture_proto_rawDescOnce.Do(func() {
		file_com_itcoursee_protocol_vo_lecture_proto_rawDescData = protoimpl.X.CompressGZIP(file_com_itcoursee_protocol_vo_lecture_proto_rawDescData)
	})
	return file_com_itcoursee_protocol_vo_lecture_proto_rawDescData
}

var file_com_itcoursee_protocol_vo_lecture_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_com_itcoursee_protocol_vo_lecture_proto_goTypes = []interface{}{
	(*Lecture)(nil),         // 0: com.itcoursee.protocol.vo.Lecture
	(core.Type)(0),          // 1: com.itcoursee.core.Type
	(core.LectureStatus)(0), // 2: com.itcoursee.core.LectureStatus
}
var file_com_itcoursee_protocol_vo_lecture_proto_depIdxs = []int32{
	1, // 0: com.itcoursee.protocol.vo.Lecture.type:type_name -> com.itcoursee.core.Type
	2, // 1: com.itcoursee.protocol.vo.Lecture.status:type_name -> com.itcoursee.core.LectureStatus
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_com_itcoursee_protocol_vo_lecture_proto_init() }
func file_com_itcoursee_protocol_vo_lecture_proto_init() {
	if File_com_itcoursee_protocol_vo_lecture_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_com_itcoursee_protocol_vo_lecture_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
			RawDescriptor: file_com_itcoursee_protocol_vo_lecture_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_com_itcoursee_protocol_vo_lecture_proto_goTypes,
		DependencyIndexes: file_com_itcoursee_protocol_vo_lecture_proto_depIdxs,
		MessageInfos:      file_com_itcoursee_protocol_vo_lecture_proto_msgTypes,
	}.Build()
	File_com_itcoursee_protocol_vo_lecture_proto = out.File
	file_com_itcoursee_protocol_vo_lecture_proto_rawDesc = nil
	file_com_itcoursee_protocol_vo_lecture_proto_goTypes = nil
	file_com_itcoursee_protocol_vo_lecture_proto_depIdxs = nil
}
