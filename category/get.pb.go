// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.14.0
// source: itcoursee/protocol/category/get.proto

package category

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

// 获取分类请求
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
		mi := &file_itcoursee_protocol_category_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReq) ProtoMessage() {}

func (x *GetReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_category_get_proto_msgTypes[0]
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
	return file_itcoursee_protocol_category_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// 获取分类返回
type GetRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 分类
	Category *vo.Category `protobuf:"bytes,10,opt,name=category,proto3" json:"category,omitempty"`
}

func (x *GetRsp) Reset() {
	*x = GetRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_category_get_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRsp) ProtoMessage() {}

func (x *GetRsp) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_category_get_proto_msgTypes[1]
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
	return file_itcoursee_protocol_category_get_proto_rawDescGZIP(), []int{1}
}

func (x *GetRsp) GetCategory() *vo.Category {
	if x != nil {
		return x.Category
	}
	return nil
}

// 按父分类获取子分类列表请求
type GetsChildrenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 父分类编号
	// @gotags: json:"id,string" query:"id" param:"id" form:"id"
	ParentId int64 `protobuf:"varint,1,opt,name=parent_id,json=parentId,proto3" json:"id,string" query:"id" param:"id" form:"id"`
}

func (x *GetsChildrenReq) Reset() {
	*x = GetsChildrenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_category_get_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsChildrenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsChildrenReq) ProtoMessage() {}

func (x *GetsChildrenReq) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_category_get_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsChildrenReq.ProtoReflect.Descriptor instead.
func (*GetsChildrenReq) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_category_get_proto_rawDescGZIP(), []int{2}
}

func (x *GetsChildrenReq) GetParentId() int64 {
	if x != nil {
		return x.ParentId
	}
	return 0
}

// 按父分类获取子分类列表响应
type GetsChildrenRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// 列表
	Categories []*vo.Category `protobuf:"bytes,10,rep,name=categories,proto3" json:"categories,omitempty"`
}

func (x *GetsChildrenRsp) Reset() {
	*x = GetsChildrenRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_itcoursee_protocol_category_get_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetsChildrenRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetsChildrenRsp) ProtoMessage() {}

func (x *GetsChildrenRsp) ProtoReflect() protoreflect.Message {
	mi := &file_itcoursee_protocol_category_get_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetsChildrenRsp.ProtoReflect.Descriptor instead.
func (*GetsChildrenRsp) Descriptor() ([]byte, []int) {
	return file_itcoursee_protocol_category_get_proto_rawDescGZIP(), []int{3}
}

func (x *GetsChildrenRsp) GetCategories() []*vo.Category {
	if x != nil {
		return x.Categories
	}
	return nil
}

var File_itcoursee_protocol_category_get_proto protoreflect.FileDescriptor

var file_itcoursee_protocol_category_get_proto_rawDesc = []byte{
	0x0a, 0x25, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x1a, 0x24, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x76, 0x6f, 0x2f, 0x63, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x18, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x45, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x52, 0x73, 0x70, 0x12, 0x3b,
	0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x22, 0x2e, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x73, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x73, 0x43, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x3f,
	0x0a, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x0a, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x76, 0x6f, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x52, 0x0a, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x42,
	0x54, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x50, 0x01, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x69, 0x74, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x3b, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_itcoursee_protocol_category_get_proto_rawDescOnce sync.Once
	file_itcoursee_protocol_category_get_proto_rawDescData = file_itcoursee_protocol_category_get_proto_rawDesc
)

func file_itcoursee_protocol_category_get_proto_rawDescGZIP() []byte {
	file_itcoursee_protocol_category_get_proto_rawDescOnce.Do(func() {
		file_itcoursee_protocol_category_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_itcoursee_protocol_category_get_proto_rawDescData)
	})
	return file_itcoursee_protocol_category_get_proto_rawDescData
}

var file_itcoursee_protocol_category_get_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_itcoursee_protocol_category_get_proto_goTypes = []interface{}{
	(*GetReq)(nil),          // 0: itcoursee.protocol.category.GetReq
	(*GetRsp)(nil),          // 1: itcoursee.protocol.category.GetRsp
	(*GetsChildrenReq)(nil), // 2: itcoursee.protocol.category.GetsChildrenReq
	(*GetsChildrenRsp)(nil), // 3: itcoursee.protocol.category.GetsChildrenRsp
	(*vo.Category)(nil),     // 4: itcoursee.protocol.vo.Category
}
var file_itcoursee_protocol_category_get_proto_depIdxs = []int32{
	4, // 0: itcoursee.protocol.category.GetRsp.category:type_name -> itcoursee.protocol.vo.Category
	4, // 1: itcoursee.protocol.category.GetsChildrenRsp.categories:type_name -> itcoursee.protocol.vo.Category
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_itcoursee_protocol_category_get_proto_init() }
func file_itcoursee_protocol_category_get_proto_init() {
	if File_itcoursee_protocol_category_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_itcoursee_protocol_category_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_itcoursee_protocol_category_get_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_itcoursee_protocol_category_get_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsChildrenReq); i {
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
		file_itcoursee_protocol_category_get_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetsChildrenRsp); i {
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
			RawDescriptor: file_itcoursee_protocol_category_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_itcoursee_protocol_category_get_proto_goTypes,
		DependencyIndexes: file_itcoursee_protocol_category_get_proto_depIdxs,
		MessageInfos:      file_itcoursee_protocol_category_get_proto_msgTypes,
	}.Build()
	File_itcoursee_protocol_category_get_proto = out.File
	file_itcoursee_protocol_category_get_proto_rawDesc = nil
	file_itcoursee_protocol_category_get_proto_goTypes = nil
	file_itcoursee_protocol_category_get_proto_depIdxs = nil
}
