// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/testprotos/news/news.proto

package news

import (
	protoreflect "github.com/thoohv5/protobuf-go/reflect/protoreflect"
	protoimpl "github.com/thoohv5/protobuf-go/runtime/protoimpl"
	anypb "github.com/thoohv5/protobuf-go/types/known/anypb"
	timestamppb "github.com/thoohv5/protobuf-go/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

type Article_Status int32

const (
	Article_DRAFT     Article_Status = 0
	Article_PUBLISHED Article_Status = 1
	Article_REVOKED   Article_Status = 2
)

// Enum value maps for Article_Status.
var (
	Article_Status_name = map[int32]string{
		0: "DRAFT",
		1: "PUBLISHED",
		2: "REVOKED",
	}
	Article_Status_value = map[string]int32{
		"DRAFT":     0,
		"PUBLISHED": 1,
		"REVOKED":   2,
	}
)

func (x Article_Status) Enum() *Article_Status {
	p := new(Article_Status)
	*p = x
	return p
}

func (x Article_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Article_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_testprotos_news_news_proto_enumTypes[0].Descriptor()
}

func (Article_Status) Type() protoreflect.EnumType {
	return &file_internal_testprotos_news_news_proto_enumTypes[0]
}

func (x Article_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Article_Status.Descriptor instead.
func (Article_Status) EnumDescriptor() ([]byte, []int) {
	return file_internal_testprotos_news_news_proto_rawDescGZIP(), []int{0, 0}
}

type Article struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author      string                 `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
	Date        *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Title       string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content     string                 `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	Status      Article_Status         `protobuf:"varint,8,opt,name=status,proto3,enum=google.golang.org.Article_Status" json:"status,omitempty"`
	Tags        []string               `protobuf:"bytes,7,rep,name=tags,proto3" json:"tags,omitempty"`
	Attachments []*anypb.Any           `protobuf:"bytes,6,rep,name=attachments,proto3" json:"attachments,omitempty"`
}

func (x *Article) Reset() {
	*x = Article{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_news_news_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Article) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Article) ProtoMessage() {}

func (x *Article) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_news_news_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Article.ProtoReflect.Descriptor instead.
func (*Article) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_news_news_proto_rawDescGZIP(), []int{0}
}

func (x *Article) GetAuthor() string {
	if x != nil {
		return x.Author
	}
	return ""
}

func (x *Article) GetDate() *timestamppb.Timestamp {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *Article) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Article) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Article) GetStatus() Article_Status {
	if x != nil {
		return x.Status
	}
	return Article_DRAFT
}

func (x *Article) GetTags() []string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Article) GetAttachments() []*anypb.Any {
	if x != nil {
		return x.Attachments
	}
	return nil
}

type BinaryAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *BinaryAttachment) Reset() {
	*x = BinaryAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_news_news_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryAttachment) ProtoMessage() {}

func (x *BinaryAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_news_news_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryAttachment.ProtoReflect.Descriptor instead.
func (*BinaryAttachment) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_news_news_proto_rawDescGZIP(), []int{1}
}

func (x *BinaryAttachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BinaryAttachment) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type KeyValueAttachment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Data map[string]string `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *KeyValueAttachment) Reset() {
	*x = KeyValueAttachment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_testprotos_news_news_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyValueAttachment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyValueAttachment) ProtoMessage() {}

func (x *KeyValueAttachment) ProtoReflect() protoreflect.Message {
	mi := &file_internal_testprotos_news_news_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyValueAttachment.ProtoReflect.Descriptor instead.
func (*KeyValueAttachment) Descriptor() ([]byte, []int) {
	return file_internal_testprotos_news_news_proto_rawDescGZIP(), []int{2}
}

func (x *KeyValueAttachment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KeyValueAttachment) GetData() map[string]string {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_internal_testprotos_news_news_proto protoreflect.FileDescriptor

var file_internal_testprotos_news_news_proto_rawDesc = []byte{
	0x0a, 0x23, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67, 0x6f,
	0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb9, 0x02, 0x0a, 0x07, 0x41, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x41, 0x72, 0x74,
	0x69, 0x63, 0x6c, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x36, 0x0a, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63,
	0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x0b, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22,
	0x2f, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x44, 0x52, 0x41,
	0x46, 0x54, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x42, 0x4c, 0x49, 0x53, 0x48, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x56, 0x4f, 0x4b, 0x45, 0x44, 0x10, 0x02,
	0x22, 0x3a, 0x0a, 0x10, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xa6, 0x01, 0x0a,
	0x12, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x67,
	0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2e, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x41, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a, 0x09,
	0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6e, 0x65, 0x77, 0x73, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_testprotos_news_news_proto_rawDescOnce sync.Once
	file_internal_testprotos_news_news_proto_rawDescData = file_internal_testprotos_news_news_proto_rawDesc
)

func file_internal_testprotos_news_news_proto_rawDescGZIP() []byte {
	file_internal_testprotos_news_news_proto_rawDescOnce.Do(func() {
		file_internal_testprotos_news_news_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_testprotos_news_news_proto_rawDescData)
	})
	return file_internal_testprotos_news_news_proto_rawDescData
}

var file_internal_testprotos_news_news_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_testprotos_news_news_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_testprotos_news_news_proto_goTypes = []interface{}{
	(Article_Status)(0),           // 0: google.golang.org.Article.Status
	(*Article)(nil),               // 1: google.golang.org.Article
	(*BinaryAttachment)(nil),      // 2: google.golang.org.BinaryAttachment
	(*KeyValueAttachment)(nil),    // 3: google.golang.org.KeyValueAttachment
	nil,                           // 4: google.golang.org.KeyValueAttachment.DataEntry
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
	(*anypb.Any)(nil),             // 6: google.protobuf.Any
}
var file_internal_testprotos_news_news_proto_depIdxs = []int32{
	5, // 0: google.golang.org.Article.date:type_name -> google.protobuf.Timestamp
	0, // 1: google.golang.org.Article.status:type_name -> google.golang.org.Article.Status
	6, // 2: google.golang.org.Article.attachments:type_name -> google.protobuf.Any
	4, // 3: google.golang.org.KeyValueAttachment.data:type_name -> google.golang.org.KeyValueAttachment.DataEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_testprotos_news_news_proto_init() }
func file_internal_testprotos_news_news_proto_init() {
	if File_internal_testprotos_news_news_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_testprotos_news_news_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Article); i {
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
		file_internal_testprotos_news_news_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryAttachment); i {
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
		file_internal_testprotos_news_news_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyValueAttachment); i {
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
			RawDescriptor: file_internal_testprotos_news_news_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_testprotos_news_news_proto_goTypes,
		DependencyIndexes: file_internal_testprotos_news_news_proto_depIdxs,
		EnumInfos:         file_internal_testprotos_news_news_proto_enumTypes,
		MessageInfos:      file_internal_testprotos_news_news_proto_msgTypes,
	}.Build()
	File_internal_testprotos_news_news_proto = out.File
	file_internal_testprotos_news_news_proto_rawDesc = nil
	file_internal_testprotos_news_news_proto_goTypes = nil
	file_internal_testprotos_news_news_proto_depIdxs = nil
}
