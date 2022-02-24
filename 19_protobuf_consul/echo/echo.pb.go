// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: echo/echo.proto

//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  echo/echo.proto

package echo

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Input struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Input) Reset() {
	*x = Input{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Input) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Input) ProtoMessage() {}

func (x *Input) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Input.ProtoReflect.Descriptor instead.
func (*Input) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{0}
}

func (x *Input) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Output) Reset() {
	*x = Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_echo_echo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Output) ProtoMessage() {}

func (x *Output) ProtoReflect() protoreflect.Message {
	mi := &file_echo_echo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Output.ProtoReflect.Descriptor instead.
func (*Output) Descriptor() ([]byte, []int) {
	return file_echo_echo_proto_rawDescGZIP(), []int{1}
}

func (x *Output) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_echo_echo_proto protoreflect.FileDescriptor

var file_echo_echo_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x63, 0x68, 0x6f, 0x2f, 0x65, 0x63, 0x68, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x65, 0x63, 0x68, 0x6f, 0x22, 0x21, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x06, 0x4f, 0x75,
	0x74, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x2a,
	0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x22, 0x0a, 0x03, 0x53, 0x61, 0x79, 0x12, 0x0b, 0x2e,
	0x65, 0x63, 0x68, 0x6f, 0x2e, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x0c, 0x2e, 0x65, 0x63, 0x68,
	0x6f, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x65,
	0x63, 0x68, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_echo_echo_proto_rawDescOnce sync.Once
	file_echo_echo_proto_rawDescData = file_echo_echo_proto_rawDesc
)

func file_echo_echo_proto_rawDescGZIP() []byte {
	file_echo_echo_proto_rawDescOnce.Do(func() {
		file_echo_echo_proto_rawDescData = protoimpl.X.CompressGZIP(file_echo_echo_proto_rawDescData)
	})
	return file_echo_echo_proto_rawDescData
}

var file_echo_echo_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_echo_echo_proto_goTypes = []interface{}{
	(*Input)(nil),  // 0: echo.Input
	(*Output)(nil), // 1: echo.Output
}
var file_echo_echo_proto_depIdxs = []int32{
	0, // 0: echo.Echo.Say:input_type -> echo.Input
	1, // 1: echo.Echo.Say:output_type -> echo.Output
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_echo_echo_proto_init() }
func file_echo_echo_proto_init() {
	if File_echo_echo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_echo_echo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Input); i {
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
		file_echo_echo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Output); i {
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
			RawDescriptor: file_echo_echo_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_echo_echo_proto_goTypes,
		DependencyIndexes: file_echo_echo_proto_depIdxs,
		MessageInfos:      file_echo_echo_proto_msgTypes,
	}.Build()
	File_echo_echo_proto = out.File
	file_echo_echo_proto_rawDesc = nil
	file_echo_echo_proto_goTypes = nil
	file_echo_echo_proto_depIdxs = nil
}
