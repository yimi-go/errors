// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: error_enum/test.proto

package error_enum

import (
	_ "github.com/yimi-go/errors"
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

type Enum int32

const (
	Enum_ZERO Enum = 0
	Enum_ONE  Enum = 1
	Enum_TWO  Enum = 2
)

// Enum value maps for Enum.
var (
	Enum_name = map[int32]string{
		0: "ZERO",
		1: "ONE",
		2: "TWO",
	}
	Enum_value = map[string]int32{
		"ZERO": 0,
		"ONE":  1,
		"TWO":  2,
	}
)

func (x Enum) Enum() *Enum {
	p := new(Enum)
	*p = x
	return p
}

func (x Enum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Enum) Descriptor() protoreflect.EnumDescriptor {
	return file_error_enum_test_proto_enumTypes[0].Descriptor()
}

func (Enum) Type() protoreflect.EnumType {
	return &file_error_enum_test_proto_enumTypes[0]
}

func (x Enum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Enum.Descriptor instead.
func (Enum) EnumDescriptor() ([]byte, []int) {
	return file_error_enum_test_proto_rawDescGZIP(), []int{0}
}

var File_error_enum_test_proto protoreflect.FileDescriptor

var file_error_enum_test_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x65,
	0x6e, 0x75, 0x6d, 0x1a, 0x1b, 0x79, 0x69, 0x6d, 0x69, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x2f, 0x76, 0x31, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2a, 0x3e, 0x0a, 0x04, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x0d, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f,
	0x10, 0x00, 0x1a, 0x03, 0xb0, 0x45, 0x01, 0x12, 0x0d, 0x0a, 0x03, 0x4f, 0x4e, 0x45, 0x10, 0x01,
	0x1a, 0x04, 0xb8, 0x45, 0x90, 0x03, 0x12, 0x10, 0x0a, 0x03, 0x54, 0x57, 0x4f, 0x10, 0x02, 0x1a,
	0x07, 0xb0, 0x45, 0x04, 0xb8, 0x45, 0xac, 0x02, 0x1a, 0x06, 0xa0, 0x45, 0x01, 0xa8, 0x45, 0x0c,
	0x42, 0x55, 0x5a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79,
	0x69, 0x6d, 0x69, 0x2d, 0x67, 0x6f, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x63, 0x6d,
	0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x79, 0x69, 0x6d,
	0x69, 0x2d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74,
	0x61, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_enum_test_proto_rawDescOnce sync.Once
	file_error_enum_test_proto_rawDescData = file_error_enum_test_proto_rawDesc
)

func file_error_enum_test_proto_rawDescGZIP() []byte {
	file_error_enum_test_proto_rawDescOnce.Do(func() {
		file_error_enum_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_enum_test_proto_rawDescData)
	})
	return file_error_enum_test_proto_rawDescData
}

var file_error_enum_test_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_enum_test_proto_goTypes = []interface{}{
	(Enum)(0), // 0: error_enum.Enum
}
var file_error_enum_test_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_enum_test_proto_init() }
func file_error_enum_test_proto_init() {
	if File_error_enum_test_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_enum_test_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_enum_test_proto_goTypes,
		DependencyIndexes: file_error_enum_test_proto_depIdxs,
		EnumInfos:         file_error_enum_test_proto_enumTypes,
	}.Build()
	File_error_enum_test_proto = out.File
	file_error_enum_test_proto_rawDesc = nil
	file_error_enum_test_proto_goTypes = nil
	file_error_enum_test_proto_depIdxs = nil
}
