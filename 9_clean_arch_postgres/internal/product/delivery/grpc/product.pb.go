// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: internal/product/delivery/grpc/product.proto

//protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  internal/product/delivery/grpc/product.proto

package grpc

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title string `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Price int64  `protobuf:"varint,3,opt,name=Price,proto3" json:"Price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_product_delivery_grpc_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_internal_product_delivery_grpc_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_internal_product_delivery_grpc_product_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Product) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Product) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ProductIdValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value uint64 `protobuf:"varint,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *ProductIdValue) Reset() {
	*x = ProductIdValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_product_delivery_grpc_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductIdValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductIdValue) ProtoMessage() {}

func (x *ProductIdValue) ProtoReflect() protoreflect.Message {
	mi := &file_internal_product_delivery_grpc_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductIdValue.ProtoReflect.Descriptor instead.
func (*ProductIdValue) Descriptor() ([]byte, []int) {
	return file_internal_product_delivery_grpc_product_proto_rawDescGZIP(), []int{1}
}

func (x *ProductIdValue) GetValue() uint64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type ArrayProducts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []*Product `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty"`
}

func (x *ArrayProducts) Reset() {
	*x = ArrayProducts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_product_delivery_grpc_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ArrayProducts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ArrayProducts) ProtoMessage() {}

func (x *ArrayProducts) ProtoReflect() protoreflect.Message {
	mi := &file_internal_product_delivery_grpc_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ArrayProducts.ProtoReflect.Descriptor instead.
func (*ArrayProducts) Descriptor() ([]byte, []int) {
	return file_internal_product_delivery_grpc_product_proto_rawDescGZIP(), []int{2}
}

func (x *ArrayProducts) GetValue() []*Product {
	if x != nil {
		return x.Value
	}
	return nil
}

type UpdateInfoProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      *ProductIdValue `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Product *Product        `protobuf:"bytes,2,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *UpdateInfoProduct) Reset() {
	*x = UpdateInfoProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_product_delivery_grpc_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoProduct) ProtoMessage() {}

func (x *UpdateInfoProduct) ProtoReflect() protoreflect.Message {
	mi := &file_internal_product_delivery_grpc_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoProduct.ProtoReflect.Descriptor instead.
func (*UpdateInfoProduct) Descriptor() ([]byte, []int) {
	return file_internal_product_delivery_grpc_product_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateInfoProduct) GetId() *ProductIdValue {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *UpdateInfoProduct) GetProduct() *Product {
	if x != nil {
		return x.Product
	}
	return nil
}

var File_internal_product_delivery_grpc_product_proto protoreflect.FileDescriptor

var file_internal_product_delivery_grpc_product_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a,
	0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x22, 0x26, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x0d,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x2f, 0x0a,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x7a,
	0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x12, 0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x02, 0x69, 0x64, 0x12, 0x33, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x32, 0xfd, 0x02, 0x0a, 0x0e, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x00,
	0x12, 0x47, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x07, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00,
	0x12, 0x48, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x49, 0x64, 0x12, 0x20,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x07, 0x5a, 0x05, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_product_delivery_grpc_product_proto_rawDescOnce sync.Once
	file_internal_product_delivery_grpc_product_proto_rawDescData = file_internal_product_delivery_grpc_product_proto_rawDesc
)

func file_internal_product_delivery_grpc_product_proto_rawDescGZIP() []byte {
	file_internal_product_delivery_grpc_product_proto_rawDescOnce.Do(func() {
		file_internal_product_delivery_grpc_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_product_delivery_grpc_product_proto_rawDescData)
	})
	return file_internal_product_delivery_grpc_product_proto_rawDescData
}

var file_internal_product_delivery_grpc_product_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_product_delivery_grpc_product_proto_goTypes = []interface{}{
	(*Product)(nil),           // 0: protobuf_session.Product
	(*ProductIdValue)(nil),    // 1: protobuf_session.ProductIdValue
	(*ArrayProducts)(nil),     // 2: protobuf_session.ArrayProducts
	(*UpdateInfoProduct)(nil), // 3: protobuf_session.UpdateInfoProduct
	(*emptypb.Empty)(nil),     // 4: google.protobuf.Empty
}
var file_internal_product_delivery_grpc_product_proto_depIdxs = []int32{
	0, // 0: protobuf_session.ArrayProducts.Value:type_name -> protobuf_session.Product
	1, // 1: protobuf_session.UpdateInfoProduct.id:type_name -> protobuf_session.ProductIdValue
	0, // 2: protobuf_session.UpdateInfoProduct.product:type_name -> protobuf_session.Product
	4, // 3: protobuf_session.ProductService.List:input_type -> google.protobuf.Empty
	0, // 4: protobuf_session.ProductService.Create:input_type -> protobuf_session.Product
	1, // 5: protobuf_session.ProductService.GetById:input_type -> protobuf_session.ProductIdValue
	3, // 6: protobuf_session.ProductService.UpdateById:input_type -> protobuf_session.UpdateInfoProduct
	1, // 7: protobuf_session.ProductService.DeleteById:input_type -> protobuf_session.ProductIdValue
	2, // 8: protobuf_session.ProductService.List:output_type -> protobuf_session.ArrayProducts
	1, // 9: protobuf_session.ProductService.Create:output_type -> protobuf_session.ProductIdValue
	0, // 10: protobuf_session.ProductService.GetById:output_type -> protobuf_session.Product
	4, // 11: protobuf_session.ProductService.UpdateById:output_type -> google.protobuf.Empty
	4, // 12: protobuf_session.ProductService.DeleteById:output_type -> google.protobuf.Empty
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_internal_product_delivery_grpc_product_proto_init() }
func file_internal_product_delivery_grpc_product_proto_init() {
	if File_internal_product_delivery_grpc_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_product_delivery_grpc_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_internal_product_delivery_grpc_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductIdValue); i {
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
		file_internal_product_delivery_grpc_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ArrayProducts); i {
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
		file_internal_product_delivery_grpc_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoProduct); i {
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
			RawDescriptor: file_internal_product_delivery_grpc_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_product_delivery_grpc_product_proto_goTypes,
		DependencyIndexes: file_internal_product_delivery_grpc_product_proto_depIdxs,
		MessageInfos:      file_internal_product_delivery_grpc_product_proto_msgTypes,
	}.Build()
	File_internal_product_delivery_grpc_product_proto = out.File
	file_internal_product_delivery_grpc_product_proto_rawDesc = nil
	file_internal_product_delivery_grpc_product_proto_goTypes = nil
	file_internal_product_delivery_grpc_product_proto_depIdxs = nil
}
