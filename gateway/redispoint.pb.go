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
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: mapped/gateway/redispoint.proto

package gateway

import (
	types "go.mapped.dev/pb/cloud/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RedisPoint is used for edge present value streams in Redis
type RedisPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The present value of this point
	PresentValue *types.TypedValue `protobuf:"bytes,1,opt,name=present_value,json=presentValue,proto3" json:"present_value,omitempty"`
	// The timestamp of when the point was last updated in the Mapped UG
	LastUpdate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=last_update,json=lastUpdate,proto3" json:"last_update,omitempty"`
	// The timestamp of when network errors started for this point (or really, the deivce that owns this point)
	// A zero value (seconds==0; nanos==0) means there are no network issues
	NetworkErrorStart *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=network_error_start,json=networkErrorStart,proto3" json:"network_error_start,omitempty"`
	// A count of how many consecutive network errors have been observed since `network_error_start`
	// A zero value (==0) here means there are there were no network errors in the last update to this point
	ConsecutiveNetworkErrorCount uint32 `protobuf:"varint,4,opt,name=consecutive_network_error_count,json=consecutiveNetworkErrorCount,proto3" json:"consecutive_network_error_count,omitempty"`
}

func (x *RedisPoint) Reset() {
	*x = RedisPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapped_gateway_redispoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RedisPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RedisPoint) ProtoMessage() {}

func (x *RedisPoint) ProtoReflect() protoreflect.Message {
	mi := &file_mapped_gateway_redispoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RedisPoint.ProtoReflect.Descriptor instead.
func (*RedisPoint) Descriptor() ([]byte, []int) {
	return file_mapped_gateway_redispoint_proto_rawDescGZIP(), []int{0}
}

func (x *RedisPoint) GetPresentValue() *types.TypedValue {
	if x != nil {
		return x.PresentValue
	}
	return nil
}

func (x *RedisPoint) GetLastUpdate() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdate
	}
	return nil
}

func (x *RedisPoint) GetNetworkErrorStart() *timestamppb.Timestamp {
	if x != nil {
		return x.NetworkErrorStart
	}
	return nil
}

func (x *RedisPoint) GetConsecutiveNetworkErrorCount() uint32 {
	if x != nil {
		return x.ConsecutiveNetworkErrorCount
	}
	return 0
}

var File_mapped_gateway_redispoint_proto protoreflect.FileDescriptor

var file_mapped_gateway_redispoint_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x72, 0x65, 0x64, 0x69, 0x73, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x24, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa1, 0x02, 0x0a, 0x0a, 0x52, 0x65, 0x64,
	0x69, 0x73, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x0d, 0x70, 0x72, 0x65, 0x73, 0x65,
	0x6e, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0c,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x3b, 0x0a, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c,
	0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x4a, 0x0a, 0x13, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x72, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x11, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x45, 0x0a, 0x1f, 0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x76, 0x65, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x1c,
	0x63, 0x6f, 0x6e, 0x73, 0x65, 0x63, 0x75, 0x74, 0x69, 0x76, 0x65, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x49, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x6f, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e,
	0x64, 0x65, 0x76, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3b, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0xaa, 0x02, 0x0e, 0x4d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mapped_gateway_redispoint_proto_rawDescOnce sync.Once
	file_mapped_gateway_redispoint_proto_rawDescData = file_mapped_gateway_redispoint_proto_rawDesc
)

func file_mapped_gateway_redispoint_proto_rawDescGZIP() []byte {
	file_mapped_gateway_redispoint_proto_rawDescOnce.Do(func() {
		file_mapped_gateway_redispoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_mapped_gateway_redispoint_proto_rawDescData)
	})
	return file_mapped_gateway_redispoint_proto_rawDescData
}

var file_mapped_gateway_redispoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mapped_gateway_redispoint_proto_goTypes = []interface{}{
	(*RedisPoint)(nil),            // 0: mapped.gateway.RedisPoint
	(*types.TypedValue)(nil),      // 1: mapped.cloud.types.TypedValue
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_mapped_gateway_redispoint_proto_depIdxs = []int32{
	1, // 0: mapped.gateway.RedisPoint.present_value:type_name -> mapped.cloud.types.TypedValue
	2, // 1: mapped.gateway.RedisPoint.last_update:type_name -> google.protobuf.Timestamp
	2, // 2: mapped.gateway.RedisPoint.network_error_start:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_mapped_gateway_redispoint_proto_init() }
func file_mapped_gateway_redispoint_proto_init() {
	if File_mapped_gateway_redispoint_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mapped_gateway_redispoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RedisPoint); i {
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
			RawDescriptor: file_mapped_gateway_redispoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mapped_gateway_redispoint_proto_goTypes,
		DependencyIndexes: file_mapped_gateway_redispoint_proto_depIdxs,
		MessageInfos:      file_mapped_gateway_redispoint_proto_msgTypes,
	}.Build()
	File_mapped_gateway_redispoint_proto = out.File
	file_mapped_gateway_redispoint_proto_rawDesc = nil
	file_mapped_gateway_redispoint_proto_goTypes = nil
	file_mapped_gateway_redispoint_proto_depIdxs = nil
}
