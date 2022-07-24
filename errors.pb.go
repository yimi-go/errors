// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: yimi/errors/v1/errors.proto

package errors

import (
	code "google.golang.org/genproto/googleapis/rpc/code"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code          code.Code         `protobuf:"varint,1,opt,name=code,proto3,enum=google.rpc.Code" json:"code,omitempty"`
	Reason        string            `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Message       string            `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Metadata      map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ForceHttpCode int32             `protobuf:"varint,5,opt,name=force_http_code,json=forceHttpCode,proto3" json:"force_http_code,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_yimi_errors_v1_errors_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_yimi_errors_v1_errors_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_yimi_errors_v1_errors_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetCode() code.Code {
	if x != nil {
		return x.Code
	}
	return code.Code(0)
}

func (x *Status) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *Status) GetForceHttpCode() int32 {
	if x != nil {
		return x.ForceHttpCode
	}
	return 0
}

var file_yimi_errors_v1_errors_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1108,
		Name:          "yimi.errors.v1.error_define",
		Tag:           "varint,1108,opt,name=error_define",
		Filename:      "yimi/errors/v1/errors.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*code.Code)(nil),
		Field:         1109,
		Name:          "yimi.errors.v1.default_code",
		Tag:           "varint,1109,opt,name=default_code,enum=google.rpc.Code",
		Filename:      "yimi/errors/v1/errors.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*code.Code)(nil),
		Field:         1110,
		Name:          "yimi.errors.v1.code",
		Tag:           "varint,1110,opt,name=code,enum=google.rpc.Code",
		Filename:      "yimi/errors/v1/errors.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*int32)(nil),
		Field:         1111,
		Name:          "yimi.errors.v1.force_http_code",
		Tag:           "varint,1111,opt,name=force_http_code",
		Filename:      "yimi/errors/v1/errors.proto",
	},
}

// Extension fields to descriptorpb.EnumOptions.
var (
	// optional bool error_define = 1108;
	E_ErrorDefine = &file_yimi_errors_v1_errors_proto_extTypes[0]
	// optional google.rpc.Code default_code = 1109;
	E_DefaultCode = &file_yimi_errors_v1_errors_proto_extTypes[1]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional google.rpc.Code code = 1110;
	E_Code = &file_yimi_errors_v1_errors_proto_extTypes[2]
	// optional int32 force_http_code = 1111;
	E_ForceHttpCode = &file_yimi_errors_v1_errors_proto_extTypes[3]
)

var File_yimi_errors_v1_errors_proto protoreflect.FileDescriptor

var file_yimi_errors_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x79, 0x69, 0x6d, 0x69, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x79,
	0x69, 0x6d, 0x69, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x15, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x02, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x24, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64,
	0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x79, 0x69,
	0x6d, 0x69, 0x2e, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x48, 0x74, 0x74, 0x70, 0x43,
	0x6f, 0x64, 0x65, 0x1a, 0x3b, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x3a, 0x40, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65,
	0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd4,
	0x08, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x66, 0x69,
	0x6e, 0x65, 0x3a, 0x52, 0x0a, 0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x63, 0x6f,
	0x64, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xd5, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x0b, 0x64, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x3a, 0x48, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x21,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd6, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x3a, 0x4a, 0x0a, 0x0f, 0x66, 0x6f, 0x72, 0x63, 0x65, 0x5f, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd7, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x66,
	0x6f, 0x72, 0x63, 0x65, 0x48, 0x74, 0x74, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x42, 0x1b, 0x5a, 0x19,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x69, 0x6d, 0x69, 0x2d,
	0x67, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_yimi_errors_v1_errors_proto_rawDescOnce sync.Once
	file_yimi_errors_v1_errors_proto_rawDescData = file_yimi_errors_v1_errors_proto_rawDesc
)

func file_yimi_errors_v1_errors_proto_rawDescGZIP() []byte {
	file_yimi_errors_v1_errors_proto_rawDescOnce.Do(func() {
		file_yimi_errors_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_yimi_errors_v1_errors_proto_rawDescData)
	})
	return file_yimi_errors_v1_errors_proto_rawDescData
}

var file_yimi_errors_v1_errors_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_yimi_errors_v1_errors_proto_goTypes = []interface{}{
	(*Status)(nil),                        // 0: yimi.errors.v1.Status
	nil,                                   // 1: yimi.errors.v1.Status.MetadataEntry
	(code.Code)(0),                        // 2: google.rpc.Code
	(*descriptorpb.EnumOptions)(nil),      // 3: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 4: google.protobuf.EnumValueOptions
}
var file_yimi_errors_v1_errors_proto_depIdxs = []int32{
	2, // 0: yimi.errors.v1.Status.code:type_name -> google.rpc.Code
	1, // 1: yimi.errors.v1.Status.metadata:type_name -> yimi.errors.v1.Status.MetadataEntry
	3, // 2: yimi.errors.v1.error_define:extendee -> google.protobuf.EnumOptions
	3, // 3: yimi.errors.v1.default_code:extendee -> google.protobuf.EnumOptions
	4, // 4: yimi.errors.v1.code:extendee -> google.protobuf.EnumValueOptions
	4, // 5: yimi.errors.v1.force_http_code:extendee -> google.protobuf.EnumValueOptions
	2, // 6: yimi.errors.v1.default_code:type_name -> google.rpc.Code
	2, // 7: yimi.errors.v1.code:type_name -> google.rpc.Code
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	6, // [6:8] is the sub-list for extension type_name
	2, // [2:6] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_yimi_errors_v1_errors_proto_init() }
func file_yimi_errors_v1_errors_proto_init() {
	if File_yimi_errors_v1_errors_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_yimi_errors_v1_errors_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_yimi_errors_v1_errors_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 4,
			NumServices:   0,
		},
		GoTypes:           file_yimi_errors_v1_errors_proto_goTypes,
		DependencyIndexes: file_yimi_errors_v1_errors_proto_depIdxs,
		MessageInfos:      file_yimi_errors_v1_errors_proto_msgTypes,
		ExtensionInfos:    file_yimi_errors_v1_errors_proto_extTypes,
	}.Build()
	File_yimi_errors_v1_errors_proto = out.File
	file_yimi_errors_v1_errors_proto_rawDesc = nil
	file_yimi_errors_v1_errors_proto_goTypes = nil
	file_yimi_errors_v1_errors_proto_depIdxs = nil
}