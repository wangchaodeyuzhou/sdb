// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1--rc1
// source: pkg/sdb-protobuf/bloom_filter.proto

package protobuf

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

type BFCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte  `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	N   uint32  `protobuf:"varint,2,opt,name=n,proto3" json:"n,omitempty"`
	P   float64 `protobuf:"fixed64,3,opt,name=p,proto3" json:"p,omitempty"`
}

func (x *BFCreateRequest) Reset() {
	*x = BFCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFCreateRequest) ProtoMessage() {}

func (x *BFCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFCreateRequest.ProtoReflect.Descriptor instead.
func (*BFCreateRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP(), []int{0}
}

func (x *BFCreateRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *BFCreateRequest) GetN() uint32 {
	if x != nil {
		return x.N
	}
	return 0
}

func (x *BFCreateRequest) GetP() float64 {
	if x != nil {
		return x.P
	}
	return 0
}

type BFCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BFCreateResponse) Reset() {
	*x = BFCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFCreateResponse) ProtoMessage() {}

func (x *BFCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFCreateResponse.ProtoReflect.Descriptor instead.
func (*BFCreateResponse) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP(), []int{1}
}

func (x *BFCreateResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type BFDelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key []byte `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *BFDelRequest) Reset() {
	*x = BFDelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFDelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFDelRequest) ProtoMessage() {}

func (x *BFDelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFDelRequest.ProtoReflect.Descriptor instead.
func (*BFDelRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP(), []int{2}
}

func (x *BFDelRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

type BFDelResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BFDelResponse) Reset() {
	*x = BFDelResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFDelResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFDelResponse) ProtoMessage() {}

func (x *BFDelResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFDelResponse.ProtoReflect.Descriptor instead.
func (*BFDelResponse) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP(), []int{3}
}

func (x *BFDelResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type BFAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *BFAddRequest) Reset() {
	*x = BFAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFAddRequest) ProtoMessage() {}

func (x *BFAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFAddRequest.ProtoReflect.Descriptor instead.
func (*BFAddRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP(), []int{4}
}

func (x *BFAddRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *BFAddRequest) GetValues() [][]byte {
	if x != nil {
		return x.Values
	}
	return nil
}

type BFAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *BFAddResponse) Reset() {
	*x = BFAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFAddResponse) ProtoMessage() {}

func (x *BFAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFAddResponse.ProtoReflect.Descriptor instead.
func (*BFAddResponse) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP(), []int{5}
}

func (x *BFAddResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type BFExistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values [][]byte `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *BFExistRequest) Reset() {
	*x = BFExistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFExistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFExistRequest) ProtoMessage() {}

func (x *BFExistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFExistRequest.ProtoReflect.Descriptor instead.
func (*BFExistRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP(), []int{6}
}

func (x *BFExistRequest) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *BFExistRequest) GetValues() [][]byte {
	if x != nil {
		return x.Values
	}
	return nil
}

type BFExistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists []bool `protobuf:"varint,1,rep,packed,name=exists,proto3" json:"exists,omitempty"`
}

func (x *BFExistResponse) Reset() {
	*x = BFExistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BFExistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BFExistResponse) ProtoMessage() {}

func (x *BFExistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BFExistResponse.ProtoReflect.Descriptor instead.
func (*BFExistResponse) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP(), []int{7}
}

func (x *BFExistResponse) GetExists() []bool {
	if x != nil {
		return x.Exists
	}
	return nil
}

var File_pkg_sdb_protobuf_bloom_filter_proto protoreflect.FileDescriptor

var file_pkg_sdb_protobuf_bloom_filter_proto_rawDesc = []byte{
	0x0a, 0x23, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x64, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x62, 0x6c, 0x6f, 0x6f, 0x6d, 0x5f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0f,
	0x42, 0x46, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x0c, 0x0a, 0x01, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x01, 0x6e, 0x12,
	0x0c, 0x0a, 0x01, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x70, 0x22, 0x2c, 0x0a,
	0x10, 0x42, 0x46, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x20, 0x0a, 0x0c, 0x42,
	0x46, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x29, 0x0a,
	0x0d, 0x42, 0x46, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x38, 0x0a, 0x0c, 0x42, 0x46, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x22, 0x29, 0x0a, 0x0d, 0x42, 0x46, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3a, 0x0a,
	0x0e, 0x42, 0x46, 0x45, 0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0c, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x29, 0x0a, 0x0f, 0x42, 0x46, 0x45,
	0x78, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sdb_protobuf_bloom_filter_proto_rawDescOnce sync.Once
	file_pkg_sdb_protobuf_bloom_filter_proto_rawDescData = file_pkg_sdb_protobuf_bloom_filter_proto_rawDesc
)

func file_pkg_sdb_protobuf_bloom_filter_proto_rawDescGZIP() []byte {
	file_pkg_sdb_protobuf_bloom_filter_proto_rawDescOnce.Do(func() {
		file_pkg_sdb_protobuf_bloom_filter_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sdb_protobuf_bloom_filter_proto_rawDescData)
	})
	return file_pkg_sdb_protobuf_bloom_filter_proto_rawDescData
}

var file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_pkg_sdb_protobuf_bloom_filter_proto_goTypes = []interface{}{
	(*BFCreateRequest)(nil),  // 0: proto.BFCreateRequest
	(*BFCreateResponse)(nil), // 1: proto.BFCreateResponse
	(*BFDelRequest)(nil),     // 2: proto.BFDelRequest
	(*BFDelResponse)(nil),    // 3: proto.BFDelResponse
	(*BFAddRequest)(nil),     // 4: proto.BFAddRequest
	(*BFAddResponse)(nil),    // 5: proto.BFAddResponse
	(*BFExistRequest)(nil),   // 6: proto.BFExistRequest
	(*BFExistResponse)(nil),  // 7: proto.BFExistResponse
}
var file_pkg_sdb_protobuf_bloom_filter_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_sdb_protobuf_bloom_filter_proto_init() }
func file_pkg_sdb_protobuf_bloom_filter_proto_init() {
	if File_pkg_sdb_protobuf_bloom_filter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFCreateRequest); i {
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
		file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFCreateResponse); i {
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
		file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFDelRequest); i {
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
		file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFDelResponse); i {
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
		file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFAddRequest); i {
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
		file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFAddResponse); i {
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
		file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFExistRequest); i {
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
		file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BFExistResponse); i {
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
			RawDescriptor: file_pkg_sdb_protobuf_bloom_filter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_sdb_protobuf_bloom_filter_proto_goTypes,
		DependencyIndexes: file_pkg_sdb_protobuf_bloom_filter_proto_depIdxs,
		MessageInfos:      file_pkg_sdb_protobuf_bloom_filter_proto_msgTypes,
	}.Build()
	File_pkg_sdb_protobuf_bloom_filter_proto = out.File
	file_pkg_sdb_protobuf_bloom_filter_proto_rawDesc = nil
	file_pkg_sdb_protobuf_bloom_filter_proto_goTypes = nil
	file_pkg_sdb_protobuf_bloom_filter_proto_depIdxs = nil
}