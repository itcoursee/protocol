// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/courseware/get.proto

package courseware

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

// 获取课件请求
type GetReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 编号
	// @gotags: json:"id,string" query:"id" param:"id" form:"id" validate:"required"
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,string" query:"id" param:"id" form:"id" validate:"required"`
}

func (x *GetReq) Reset() {
	*x = GetReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReq.ProtoReflect.Descriptor instead.
func (*GetReq) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_courseware_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 获取转码返回
type GetRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课件
	Courseware *vo.Courseware `protobuf:"bytes,10,opt,name=courseware,proto3" json:"courseware,omitempty"`
}

func (x *GetRsp) Reset() {
	*x = GetRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRsp) ProtoMessage() {}

func (x *GetRsp) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRsp.ProtoReflect.Descriptor instead.
func (*GetRsp) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_courseware_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetRsp) GetCourseware() *vo.Courseware {
	if x != nil {
		return x.Courseware
	}
	return nil
}

// 获得课程的所有课件请求
type GetsByCourseReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 课程编号
	// @gotags: json:"courseId,string" query:"courseId" param:"courseId" form:"courseId" validate:"required"
	CourseId int64 `protobuf:"varint,10,opt,name=course_id,json=courseId,proto3" json:"courseId,string" query:"courseId" param:"courseId" form:"courseId" validate:"required"`
}

func (x *GetsByCourseReq) Reset() {
	*x = GetsByCourseReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByCourseReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByCourseReq) ProtoMessage() {}

func (x *GetsByCourseReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[2]
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
	return file_itcoursee_protocol_courseware_get_proto_rawDescGZIP(), []int{2}
}

func (x *GetsByCourseReq) GetCourseId() int64 {
	if x != nil {
		return x.CourseId
	}
	return 0
}

// 获得讲次的所 有课件请求
type GetsByLectureReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 讲次编号
	// @gotags: json:"lectureId,string" query:"lectureId" param:"lectureId" form:"lectureId" validate:"required"
	LectureId int64 `protobuf:"varint,20,opt,name=lecture_id,json=lectureId,proto3" json:"lectureId,string" query:"lectureId" param:"lectureId" form:"lectureId" validate:"required"`
}

func (x *GetsByLectureReq) Reset() {
	*x = GetsByLectureReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByLectureReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByLectureReq) ProtoMessage() {}

func (x *GetsByLectureReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByLectureReq.ProtoReflect.Descriptor instead.
func (*GetsByLectureReq) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_courseware_get_proto_rawDescGZIP(), []int{3}
}

func (x *GetsByLectureReq) GetLectureId() int64 {
	if x != nil {
		return x.LectureId
	}
	return 0
}

// 按拥有者获得所有课件响应
type GetsByOwnerRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 数据
	Coursewares []*vo.Courseware `protobuf:"bytes,10,rep,name=coursewares,proto3" json:"coursewares,omitempty"`
}

func (x *GetsByOwnerRsp) Reset() {
	*x = GetsByOwnerRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsByOwnerRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsByOwnerRsp) ProtoMessage() {}

func (x *GetsByOwnerRsp) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_courseware_get_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsByOwnerRsp.ProtoReflect.Descriptor instead.
func (*GetsByOwnerRsp) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_courseware_get_proto_rawDescGZIP(), []int{4}
}

func (x *GetsByOwnerRsp) GetCoursewares() []*vo.Courseware {
	if x != nil {
		return x.Coursewares
	}
	return nil
}

var File_itcoursee_protocol_courseware_get_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_courseware_get_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2f,
	0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x1a, 0x26, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x18, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x52, 0x73, 0x70, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f,
	0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x22, 0x2e, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x49, 0x64, 0x22, 0x31, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x73, 0x42,
	0x79, 0x4c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x6c,
	0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x6c, 0x65, 0x63, 0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x73, 0x42, 0x79, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x52, 0x73, 0x70, 0x12, 0x43, 0x0a, 0x0b,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x77, 0x61, 0x72, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65,
	0x73, 0x42, 0x5a, 0x0a, 0x21, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61,
	0x72, 0x65, 0x3b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x77, 0x61, 0x72, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itcoursee_protocol_courseware_get_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_courseware_get_proto_rawDescData = file_itcoursee_protocol_courseware_get_proto_rawDesc
)

func file_itcoursee_protocol_courseware_get_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_courseware_get_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_courseware_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_courseware_get_proto_rawDescData)
	})
	return file_itcoursee_protocol_courseware_get_proto_rawDescData
}

var file_itcoursee_protocol_courseware_get_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_itcoursee_protocol_courseware_get_proto_goTypes = []interface{}{
	(*GetReq)(nil),           // 0: itcoursee.protocol.courseware.GetReq
	(*GetRsp)(nil),           // 1: itcoursee.protocol.courseware.GetRsp
	(*GetsByCourseReq)(nil),  // 2: itcoursee.protocol.courseware.GetsByCourseReq
	(*GetsByLectureReq)(nil), // 3: itcoursee.protocol.courseware.GetsByLectureReq
	(*GetsByOwnerRsp)(nil),   // 4: itcoursee.protocol.courseware.GetsByOwnerRsp
	(*vo.Courseware)(nil),    // 5: itcoursee.protocol.vo.Courseware
}
var file_itcoursee_protocol_courseware_get_proto_depIdxs = []int32{
	5, // 0: itcoursee.protocol.courseware.GetRsp.courseware:type_name -> itcoursee.protocol.vo.Courseware
	5, // 1: itcoursee.protocol.courseware.GetsByOwnerRsp.coursewares:type_name -> itcoursee.protocol.vo.Courseware
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_courseware_get_proto_init() }
func file_itcoursee_protocol_courseware_get_proto_init() {
	if File_itcoursee_protocol_courseware_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_courseware_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetReq); i {
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
		file_itcoursee_protocol_courseware_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRsp); i {
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
		file_itcoursee_protocol_courseware_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_itcoursee_protocol_courseware_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByLectureReq); i {
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
		file_itcoursee_protocol_courseware_get_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsByOwnerRsp); i {
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
			RawDescriptor: file_itcoursee_protocol_courseware_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_courseware_get_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_courseware_get_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_courseware_get_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_courseware_get_proto = out.File
	file_itcoursee_protocol_courseware_get_proto_rawDesc = nil
	file_itcoursee_protocol_courseware_get_proto_goTypes = nil
	file_itcoursee_protocol_courseware_get_proto_depIdxs = nil
}
