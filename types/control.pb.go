// ******************************************************************************************************
// Copyright 2022 Mapped Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// ******************************************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: mapped/types/control.proto

package types

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

type ControlResult int32

const (
	ControlResult_RESULT_UNSPECIFIED ControlResult = 0
	// Success codes (1-99)
	ControlResult_SUCCESS ControlResult = 1
	// Error codes (100+)
	ControlResult_ERROR_UNKNOWN                   ControlResult = 100
	ControlResult_ERROR_UNKNOWN_MAPPING_KEY       ControlResult = 101
	ControlResult_ERROR_TIMEOUT                   ControlResult = 103
	ControlResult_ERROR_HOST_UNREACHABLE          ControlResult = 104
	ControlResult_ERROR_VALUE_TYPE_CHANGE         ControlResult = 120
	ControlResult_ERROR_VALUE_ARRAY_INDEX_INVALID ControlResult = 121
	ControlResult_ERROR_VALUE_TYPE_INVALID        ControlResult = 122
	ControlResult_ERROR_VALUE_TYPE_UNSUPPORTED    ControlResult = 123
	ControlResult_ERROR_VALUE_NOT_ARRAY           ControlResult = 124
	ControlResult_ERROR_VALUE_OUT_OF_RANGE        ControlResult = 125
	ControlResult_ERROR_VALUE_WRITE_ACCESS_DENIED ControlResult = 126
	ControlResult_ERROR_DEVICE_OTHER              ControlResult = 150
	ControlResult_ERROR_DEVICE_NO_WRITE_SUPPORT   ControlResult = 151
	ControlResult_ERROR_DEVICE_ACCESS_DENIED      ControlResult = 152
)

// Enum value maps for ControlResult.
var (
	ControlResult_name = map[int32]string{
		0:   "RESULT_UNSPECIFIED",
		1:   "SUCCESS",
		100: "ERROR_UNKNOWN",
		101: "ERROR_UNKNOWN_MAPPING_KEY",
		103: "ERROR_TIMEOUT",
		104: "ERROR_HOST_UNREACHABLE",
		120: "ERROR_VALUE_TYPE_CHANGE",
		121: "ERROR_VALUE_ARRAY_INDEX_INVALID",
		122: "ERROR_VALUE_TYPE_INVALID",
		123: "ERROR_VALUE_TYPE_UNSUPPORTED",
		124: "ERROR_VALUE_NOT_ARRAY",
		125: "ERROR_VALUE_OUT_OF_RANGE",
		126: "ERROR_VALUE_WRITE_ACCESS_DENIED",
		150: "ERROR_DEVICE_OTHER",
		151: "ERROR_DEVICE_NO_WRITE_SUPPORT",
		152: "ERROR_DEVICE_ACCESS_DENIED",
	}
	ControlResult_value = map[string]int32{
		"RESULT_UNSPECIFIED":              0,
		"SUCCESS":                         1,
		"ERROR_UNKNOWN":                   100,
		"ERROR_UNKNOWN_MAPPING_KEY":       101,
		"ERROR_TIMEOUT":                   103,
		"ERROR_HOST_UNREACHABLE":          104,
		"ERROR_VALUE_TYPE_CHANGE":         120,
		"ERROR_VALUE_ARRAY_INDEX_INVALID": 121,
		"ERROR_VALUE_TYPE_INVALID":        122,
		"ERROR_VALUE_TYPE_UNSUPPORTED":    123,
		"ERROR_VALUE_NOT_ARRAY":           124,
		"ERROR_VALUE_OUT_OF_RANGE":        125,
		"ERROR_VALUE_WRITE_ACCESS_DENIED": 126,
		"ERROR_DEVICE_OTHER":              150,
		"ERROR_DEVICE_NO_WRITE_SUPPORT":   151,
		"ERROR_DEVICE_ACCESS_DENIED":      152,
	}
)

func (x ControlResult) Enum() *ControlResult {
	p := new(ControlResult)
	*p = x
	return p
}

func (x ControlResult) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControlResult) Descriptor() protoreflect.EnumDescriptor {
	return file_mapped_types_control_proto_enumTypes[0].Descriptor()
}

func (ControlResult) Type() protoreflect.EnumType {
	return &file_mapped_types_control_proto_enumTypes[0]
}

func (x ControlResult) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControlResult.Descriptor instead.
func (ControlResult) EnumDescriptor() ([]byte, []int) {
	return file_mapped_types_control_proto_rawDescGZIP(), []int{0}
}

var File_mapped_types_control_proto protoreflect.FileDescriptor

var file_mapped_types_control_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6d, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2a, 0xd3, 0x03, 0x0a, 0x0d, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x12,
	0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49,
	0x45, 0x44, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10,
	0x01, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x64, 0x12, 0x1d, 0x0a, 0x19, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x4d, 0x41, 0x50, 0x50, 0x49, 0x4e, 0x47, 0x5f, 0x4b, 0x45,
	0x59, 0x10, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x54, 0x49, 0x4d,
	0x45, 0x4f, 0x55, 0x54, 0x10, 0x67, 0x12, 0x1a, 0x0a, 0x16, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f,
	0x48, 0x4f, 0x53, 0x54, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x4c, 0x45,
	0x10, 0x68, 0x12, 0x1b, 0x0a, 0x17, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x78, 0x12,
	0x23, 0x0a, 0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x41,
	0x52, 0x52, 0x41, 0x59, 0x5f, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x10, 0x79, 0x12, 0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44,
	0x10, 0x7a, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x55,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54,
	0x45, 0x44, 0x10, 0x7b, 0x12, 0x19, 0x0a, 0x15, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x52, 0x52, 0x41, 0x59, 0x10, 0x7c, 0x12,
	0x1c, 0x0a, 0x18, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x4f,
	0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x7d, 0x12, 0x23, 0x0a,
	0x1f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x57, 0x52, 0x49,
	0x54, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44,
	0x10, 0x7e, 0x12, 0x17, 0x0a, 0x12, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x56, 0x49,
	0x43, 0x45, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x96, 0x01, 0x12, 0x22, 0x0a, 0x1d, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x4e, 0x4f, 0x5f, 0x57,
	0x52, 0x49, 0x54, 0x45, 0x5f, 0x53, 0x55, 0x50, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x97, 0x01, 0x12,
	0x1f, 0x0a, 0x1a, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x5f, 0x44, 0x45, 0x56, 0x49, 0x43, 0x45, 0x5f,
	0x41, 0x43, 0x43, 0x45, 0x53, 0x53, 0x5f, 0x44, 0x45, 0x4e, 0x49, 0x45, 0x44, 0x10, 0x98, 0x01,
	0x42, 0x41, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e, 0x74,
	0x79, 0x70, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x1c, 0x67, 0x6f, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x62, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x3b, 0x74,
	0x79, 0x70, 0x65, 0x73, 0xaa, 0x02, 0x0c, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mapped_types_control_proto_rawDescOnce sync.Once
	file_mapped_types_control_proto_rawDescData = file_mapped_types_control_proto_rawDesc
)

func file_mapped_types_control_proto_rawDescGZIP() []byte {
	file_mapped_types_control_proto_rawDescOnce.Do(func() {
		file_mapped_types_control_proto_rawDescData = protoimpl.X.CompressGZIP(file_mapped_types_control_proto_rawDescData)
	})
	return file_mapped_types_control_proto_rawDescData
}

var file_mapped_types_control_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_mapped_types_control_proto_goTypes = []interface{}{
	(ControlResult)(0), // 0: mapped.types.ControlResult
}
var file_mapped_types_control_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mapped_types_control_proto_init() }
func file_mapped_types_control_proto_init() {
	if File_mapped_types_control_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mapped_types_control_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mapped_types_control_proto_goTypes,
		DependencyIndexes: file_mapped_types_control_proto_depIdxs,
		EnumInfos:         file_mapped_types_control_proto_enumTypes,
	}.Build()
	File_mapped_types_control_proto = out.File
	file_mapped_types_control_proto_rawDesc = nil
	file_mapped_types_control_proto_goTypes = nil
	file_mapped_types_control_proto_depIdxs = nil
}
