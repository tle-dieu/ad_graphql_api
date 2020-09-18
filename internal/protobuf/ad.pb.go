// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: ad.proto

package protobuf

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

type Ad struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ref     string      `protobuf:"bytes,1,opt,name=ref,proto3" json:"ref,omitempty"`
	Brand   string      `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model   string      `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Price   int32       `protobuf:"varint,4,opt,name=price,proto3" json:"price,omitempty"`
	Options *Ad_Options `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *Ad) Reset() {
	*x = Ad{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ad) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ad) ProtoMessage() {}

func (x *Ad) ProtoReflect() protoreflect.Message {
	mi := &file_ad_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ad.ProtoReflect.Descriptor instead.
func (*Ad) Descriptor() ([]byte, []int) {
	return file_ad_proto_rawDescGZIP(), []int{0}
}

func (x *Ad) GetRef() string {
	if x != nil {
		return x.Ref
	}
	return ""
}

func (x *Ad) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Ad) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Ad) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Ad) GetOptions() *Ad_Options {
	if x != nil {
		return x.Options
	}
	return nil
}

type Ads struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ads []*Ad `protobuf:"bytes,1,rep,name=ads,proto3" json:"ads,omitempty"`
}

func (x *Ads) Reset() {
	*x = Ads{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ads) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ads) ProtoMessage() {}

func (x *Ads) ProtoReflect() protoreflect.Message {
	mi := &file_ad_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ads.ProtoReflect.Descriptor instead.
func (*Ads) Descriptor() ([]byte, []int) {
	return file_ad_proto_rawDescGZIP(), []int{1}
}

func (x *Ads) GetAds() []*Ad {
	if x != nil {
		return x.Ads
	}
	return nil
}

type Ad_Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bluetooth bool `protobuf:"varint,1,opt,name=bluetooth,proto3" json:"bluetooth,omitempty"`
	Gps       bool `protobuf:"varint,2,opt,name=gps,proto3" json:"gps,omitempty"`
}

func (x *Ad_Options) Reset() {
	*x = Ad_Options{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ad_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ad_Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ad_Options) ProtoMessage() {}

func (x *Ad_Options) ProtoReflect() protoreflect.Message {
	mi := &file_ad_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ad_Options.ProtoReflect.Descriptor instead.
func (*Ad_Options) Descriptor() ([]byte, []int) {
	return file_ad_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Ad_Options) GetBluetooth() bool {
	if x != nil {
		return x.Bluetooth
	}
	return false
}

func (x *Ad_Options) GetGps() bool {
	if x != nil {
		return x.Gps
	}
	return false
}

var File_ad_proto protoreflect.FileDescriptor

var file_ad_proto_rawDesc = []byte{
	0x0a, 0x08, 0x61, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xba, 0x01, 0x0a, 0x02, 0x41,
	0x64, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x66, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x72, 0x65, 0x66, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x25, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x41, 0x64, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x39, 0x0a, 0x07,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6c, 0x75, 0x65, 0x74,
	0x6f, 0x6f, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6c, 0x75, 0x65,
	0x74, 0x6f, 0x6f, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x03, 0x67, 0x70, 0x73, 0x22, 0x1c, 0x0a, 0x03, 0x41, 0x64, 0x73, 0x12, 0x15,
	0x0a, 0x03, 0x61, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x03, 0x2e, 0x41, 0x64,
	0x52, 0x03, 0x61, 0x64, 0x73, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ad_proto_rawDescOnce sync.Once
	file_ad_proto_rawDescData = file_ad_proto_rawDesc
)

func file_ad_proto_rawDescGZIP() []byte {
	file_ad_proto_rawDescOnce.Do(func() {
		file_ad_proto_rawDescData = protoimpl.X.CompressGZIP(file_ad_proto_rawDescData)
	})
	return file_ad_proto_rawDescData
}

var file_ad_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ad_proto_goTypes = []interface{}{
	(*Ad)(nil),         // 0: Ad
	(*Ads)(nil),        // 1: Ads
	(*Ad_Options)(nil), // 2: Ad.Options
}
var file_ad_proto_depIdxs = []int32{
	2, // 0: Ad.options:type_name -> Ad.Options
	0, // 1: Ads.ads:type_name -> Ad
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_ad_proto_init() }
func file_ad_proto_init() {
	if File_ad_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ad_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ad); i {
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
		file_ad_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ads); i {
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
		file_ad_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ad_Options); i {
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
			RawDescriptor: file_ad_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ad_proto_goTypes,
		DependencyIndexes: file_ad_proto_depIdxs,
		MessageInfos:      file_ad_proto_msgTypes,
	}.Build()
	File_ad_proto = out.File
	file_ad_proto_rawDesc = nil
	file_ad_proto_goTypes = nil
	file_ad_proto_depIdxs = nil
}
