// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1--rc1
// source: pkg/sdb-protobuf/log.proto

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

type Op int32

const (
	Op_OP_SET Op = 0
	Op_OP_DEL Op = 1
)

// Enum value maps for Op.
var (
	Op_name = map[int32]string{
		0: "OP_SET",
		1: "OP_DEL",
	}
	Op_value = map[string]int32{
		"OP_SET": 0,
		"OP_DEL": 1,
	}
)

func (x Op) Enum() *Op {
	p := new(Op)
	*p = x
	return p
}

func (x Op) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Op) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_sdb_protobuf_log_proto_enumTypes[0].Descriptor()
}

func (Op) Type() protoreflect.EnumType {
	return &file_pkg_sdb_protobuf_log_proto_enumTypes[0]
}

func (x Op) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Op.Descriptor instead.
func (Op) EnumDescriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_log_proto_rawDescGZIP(), []int{0}
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogEntries []*LogEntry `protobuf:"bytes,1,rep,name=logEntries,proto3" json:"logEntries,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_log_proto_rawDescGZIP(), []int{0}
}

func (x *Log) GetLogEntries() []*LogEntry {
	if x != nil {
		return x.LogEntries
	}
	return nil
}

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Op    Op     `protobuf:"varint,1,opt,name=op,proto3,enum=proto.Op" json:"op,omitempty"`
	Key   []byte `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_sdb_protobuf_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_sdb_protobuf_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_pkg_sdb_protobuf_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogEntry) GetOp() Op {
	if x != nil {
		return x.Op
	}
	return Op_OP_SET
}

func (x *LogEntry) GetKey() []byte {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *LogEntry) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_pkg_sdb_protobuf_log_proto protoreflect.FileDescriptor

var file_pkg_sdb_protobuf_log_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x6b, 0x67, 0x2f, 0x73, 0x64, 0x62, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x2f, 0x0a, 0x0a, 0x6c, 0x6f,
	0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0a, 0x6c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x22, 0x4d, 0x0a, 0x08, 0x4c,
	0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x19, 0x0a, 0x02, 0x6f, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x09, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4f, 0x70, 0x52, 0x02,
	0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x1c, 0x0a, 0x02, 0x4f, 0x70,
	0x12, 0x0a, 0x0a, 0x06, 0x4f, 0x50, 0x5f, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x4f, 0x50, 0x5f, 0x44, 0x45, 0x4c, 0x10, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_sdb_protobuf_log_proto_rawDescOnce sync.Once
	file_pkg_sdb_protobuf_log_proto_rawDescData = file_pkg_sdb_protobuf_log_proto_rawDesc
)

func file_pkg_sdb_protobuf_log_proto_rawDescGZIP() []byte {
	file_pkg_sdb_protobuf_log_proto_rawDescOnce.Do(func() {
		file_pkg_sdb_protobuf_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_sdb_protobuf_log_proto_rawDescData)
	})
	return file_pkg_sdb_protobuf_log_proto_rawDescData
}

var file_pkg_sdb_protobuf_log_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_sdb_protobuf_log_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_pkg_sdb_protobuf_log_proto_goTypes = []interface{}{
	(Op)(0),          // 0: proto.Op
	(*Log)(nil),      // 1: proto.Log
	(*LogEntry)(nil), // 2: proto.LogEntry
}
var file_pkg_sdb_protobuf_log_proto_depIdxs = []int32{
	2, // 0: proto.Log.logEntries:type_name -> proto.LogEntry
	0, // 1: proto.LogEntry.op:type_name -> proto.Op
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_sdb_protobuf_log_proto_init() }
func file_pkg_sdb_protobuf_log_proto_init() {
	if File_pkg_sdb_protobuf_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_sdb_protobuf_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_pkg_sdb_protobuf_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
			RawDescriptor: file_pkg_sdb_protobuf_log_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_sdb_protobuf_log_proto_goTypes,
		DependencyIndexes: file_pkg_sdb_protobuf_log_proto_depIdxs,
		EnumInfos:         file_pkg_sdb_protobuf_log_proto_enumTypes,
		MessageInfos:      file_pkg_sdb_protobuf_log_proto_msgTypes,
	}.Build()
	File_pkg_sdb_protobuf_log_proto = out.File
	file_pkg_sdb_protobuf_log_proto_rawDesc = nil
	file_pkg_sdb_protobuf_log_proto_goTypes = nil
	file_pkg_sdb_protobuf_log_proto_depIdxs = nil
}
