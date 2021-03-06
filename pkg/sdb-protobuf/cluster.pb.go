// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1--rc1
// source: pkg/sdb-protobuf/cluster.proto

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

type Node struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Leader  bool   `protobuf:"varint,3,opt,name=leader,proto3" json:"leader,omitempty"`
}

func (x *Node) Reset() {
	*x = Node{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Node) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Node) ProtoMessage() {}

func (x *Node) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Node.ProtoReflect.Descriptor instead.
func (*Node) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Node) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Node) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Node) GetLeader() bool {
	if x != nil {
		return x.Leader
	}
	return false
}

type CInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CInfoRequest) Reset() {
	*x = CInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_cluster_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CInfoRequest) ProtoMessage() {}

func (x *CInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_cluster_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CInfoRequest.ProtoReflect.Descriptor instead.
func (*CInfoRequest) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_cluster_proto_rawDescGZIP(), []int{1}
}

type CInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *CInfoResponse) Reset() {
	*x = CInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_cluster_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CInfoResponse) ProtoMessage() {}

func (x *CInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_cluster_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CInfoResponse.ProtoReflect.Descriptor instead.
func (*CInfoResponse) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_cluster_proto_rawDescGZIP(), []int{2}
}

func (x *CInfoResponse) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_pkg_sdb_protobuf_cluster_proto protoreflect.FileDescriptor

var file_pkg_sdb_protobuf_cluster_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x64, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x32, 0x0a, 0x0d, 0x43, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05,
	0x6e, 0x6f, 0x64, 0x65, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sdb_protobuf_cluster_proto_rawDescOnce sync.Once
	file_pkg_sdb_protobuf_cluster_proto_rawDescData = file_pkg_sdb_protobuf_cluster_proto_rawDesc
)

func file_pkg_sdb_protobuf_cluster_proto_rawDescGZIP() []byte {
	file_pkg_sdb_protobuf_cluster_proto_rawDescOnce.Do(func() {
		file_pkg_sdb_protobuf_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sdb_protobuf_cluster_proto_rawDescData)
	})
	return file_pkg_sdb_protobuf_cluster_proto_rawDescData
}

var file_pkg_sdb_protobuf_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pkg_sdb_protobuf_cluster_proto_goTypes = []interface{}{
	(*Node)(nil),          // 0: proto.Node
	(*CInfoRequest)(nil),  // 1: proto.CInfoRequest
	(*CInfoResponse)(nil), // 2: proto.CInfoResponse
}
var file_pkg_sdb_protobuf_cluster_proto_depIdxs = []int32{
	0, // 0: proto.CInfoResponse.nodes:type_name -> proto.Node
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_sdb_protobuf_cluster_proto_init() }
func file_pkg_sdb_protobuf_cluster_proto_init() {
	if File_pkg_sdb_protobuf_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sdb_protobuf_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Node); i {
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
		file_pkg_sdb_protobuf_cluster_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CInfoRequest); i {
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
		file_pkg_sdb_protobuf_cluster_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CInfoResponse); i {
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
			RawDescriptor: file_pkg_sdb_protobuf_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_sdb_protobuf_cluster_proto_goTypes,
		DependencyIndexes: file_pkg_sdb_protobuf_cluster_proto_depIdxs,
		MessageInfos:      file_pkg_sdb_protobuf_cluster_proto_msgTypes,
	}.Build()
	File_pkg_sdb_protobuf_cluster_proto = out.File
	file_pkg_sdb_protobuf_cluster_proto_rawDesc = nil
	file_pkg_sdb_protobuf_cluster_proto_goTypes = nil
	file_pkg_sdb_protobuf_cluster_proto_depIdxs = nil
}
