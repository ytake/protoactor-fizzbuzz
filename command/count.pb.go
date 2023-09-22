// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v4.24.3
// source: count.proto

package command

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

type Say struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int64  `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Text   string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *Say) Reset() {
	*x = Say{}
	if protoimpl.UnsafeEnabled {
		mi := &file_count_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Say) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Say) ProtoMessage() {}

func (x *Say) ProtoReflect() protoreflect.Message {
	mi := &file_count_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Say.ProtoReflect.Descriptor instead.
func (*Say) Descriptor() ([]byte, []int) {
	return file_count_proto_rawDescGZIP(), []int{0}
}

func (x *Say) GetNumber() int64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Say) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type IllegalStatePanicMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *IllegalStatePanicMessage) Reset() {
	*x = IllegalStatePanicMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_count_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IllegalStatePanicMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IllegalStatePanicMessage) ProtoMessage() {}

func (x *IllegalStatePanicMessage) ProtoReflect() protoreflect.Message {
	mi := &file_count_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IllegalStatePanicMessage.ProtoReflect.Descriptor instead.
func (*IllegalStatePanicMessage) Descriptor() ([]byte, []int) {
	return file_count_proto_rawDescGZIP(), []int{1}
}

func (x *IllegalStatePanicMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_count_proto protoreflect.FileDescriptor

var file_count_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x22, 0x31, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x34, 0x0a, 0x18, 0x49, 0x6c,
	0x6c, 0x65, 0x67, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x74, 0x61, 0x6b, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d,
	0x66, 0x69, 0x7a, 0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_count_proto_rawDescOnce sync.Once
	file_count_proto_rawDescData = file_count_proto_rawDesc
)

func file_count_proto_rawDescGZIP() []byte {
	file_count_proto_rawDescOnce.Do(func() {
		file_count_proto_rawDescData = protoimpl.X.CompressGZIP(file_count_proto_rawDescData)
	})
	return file_count_proto_rawDescData
}

var file_count_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_count_proto_goTypes = []interface{}{
	(*Say)(nil),                      // 0: protobuf.Say
	(*IllegalStatePanicMessage)(nil), // 1: protobuf.IllegalStatePanicMessage
}
var file_count_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_count_proto_init() }
func file_count_proto_init() {
	if File_count_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_count_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Say); i {
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
		file_count_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IllegalStatePanicMessage); i {
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
			RawDescriptor: file_count_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_count_proto_goTypes,
		DependencyIndexes: file_count_proto_depIdxs,
		MessageInfos:      file_count_proto_msgTypes,
	}.Build()
	File_count_proto = out.File
	file_count_proto_rawDesc = nil
	file_count_proto_goTypes = nil
	file_count_proto_depIdxs = nil
}
