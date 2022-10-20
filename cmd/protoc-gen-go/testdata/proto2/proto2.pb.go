// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmd/protoc-gen-go/testdata/proto2/proto2.proto

package proto2

import (
	protoreflect "github.com/thoohv5/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/thoohv5/protobuf-go/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	I32 *int32   `protobuf:"varint,1,opt,name=i32" json:"i32,omitempty"`
	M   *Message `protobuf:"bytes,2,opt,name=m" json:"m,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetI32() int32 {
	if x != nil && x.I32 != nil {
		return *x.I32
	}
	return 0
}

func (x *Message) GetM() *Message {
	if x != nil {
		return x.M
	}
	return nil
}

var File_cmd_protoc_gen_go_testdata_proto2_proto2_proto protoreflect.FileDescriptor

var file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e,
	0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x32, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x15, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x22, 0x49, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x33, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x69, 0x33, 0x32, 0x12, 0x2c, 0x0a, 0x01, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x01, 0x6d, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c,
	0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x63, 0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d,
	0x67, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x32,
}

var (
	file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDescOnce sync.Once
	file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDescData = file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDesc
)

func file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDescGZIP() []byte {
	file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDescOnce.Do(func() {
		file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDescData = protoimpl.X.CompressGZIP(file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDescData)
	})
	return file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDescData
}

var file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_goTypes = []interface{}{
	(*Message)(nil), // 0: goproto.protoc.proto2.Message
}
var file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_depIdxs = []int32{
	0, // 0: goproto.protoc.proto2.Message.m:type_name -> goproto.protoc.proto2.Message
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_init() }
func file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_init() {
	if File_cmd_protoc_gen_go_testdata_proto2_proto2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_goTypes,
		DependencyIndexes: file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_depIdxs,
		MessageInfos:      file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_msgTypes,
	}.Build()
	File_cmd_protoc_gen_go_testdata_proto2_proto2_proto = out.File
	file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_rawDesc = nil
	file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_goTypes = nil
	file_cmd_protoc_gen_go_testdata_proto2_proto2_proto_depIdxs = nil
}
