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
// source: mapped/gateway/edgecontrol.proto

package gateway

import (
	types "go.mapped.dev/pb/cloud/types"
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

type WritePropsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The property value(s) to write
	Values []*MappingKeyAndValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	// The timeout for each request to each subtended device in milliseconds
	// NOTE: UGAgent makes every attempt to parallelize these, but some devices don't support
	//       writing multiple properties at once, which means write requests must be serialized
	TimeoutPerDeviceInMs uint32 `protobuf:"varint,2,opt,name=timeout_per_device_in_ms,json=timeoutPerDeviceInMs,proto3" json:"timeout_per_device_in_ms,omitempty"`
	// Simulate the command by only validating the input and logging
	// do not send actual commands over the wire
	Simulate bool `protobuf:"varint,100,opt,name=simulate,proto3" json:"simulate,omitempty"`
}

func (x *WritePropsRequest) Reset() {
	*x = WritePropsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapped_gateway_edgecontrol_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePropsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePropsRequest) ProtoMessage() {}

func (x *WritePropsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mapped_gateway_edgecontrol_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePropsRequest.ProtoReflect.Descriptor instead.
func (*WritePropsRequest) Descriptor() ([]byte, []int) {
	return file_mapped_gateway_edgecontrol_proto_rawDescGZIP(), []int{0}
}

func (x *WritePropsRequest) GetValues() []*MappingKeyAndValue {
	if x != nil {
		return x.Values
	}
	return nil
}

func (x *WritePropsRequest) GetTimeoutPerDeviceInMs() uint32 {
	if x != nil {
		return x.TimeoutPerDeviceInMs
	}
	return 0
}

func (x *WritePropsRequest) GetSimulate() bool {
	if x != nil {
		return x.Simulate
	}
	return false
}

type WritePropsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The unique ID of this request to be provided to Mapped Support
	// for troubleshooting in the case of unexpected errors
	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// The result of each property write
	Results []*MappingKeyAndResult `protobuf:"bytes,2,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *WritePropsResponse) Reset() {
	*x = WritePropsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapped_gateway_edgecontrol_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WritePropsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WritePropsResponse) ProtoMessage() {}

func (x *WritePropsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mapped_gateway_edgecontrol_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WritePropsResponse.ProtoReflect.Descriptor instead.
func (*WritePropsResponse) Descriptor() ([]byte, []int) {
	return file_mapped_gateway_edgecontrol_proto_rawDescGZIP(), []int{1}
}

func (x *WritePropsResponse) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *WritePropsResponse) GetResults() []*MappingKeyAndResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type MappingKeyAndValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mappingKey where the value originated from or is to be written to
	MappingKey string `protobuf:"bytes,1,opt,name=mapping_key,json=mappingKey,proto3" json:"mapping_key,omitempty"`
	// The value read or to be written
	// NOTE: The value *type* must NOT change type between a read and a write
	Value *types.TypedValue `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *MappingKeyAndValue) Reset() {
	*x = MappingKeyAndValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapped_gateway_edgecontrol_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MappingKeyAndValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MappingKeyAndValue) ProtoMessage() {}

func (x *MappingKeyAndValue) ProtoReflect() protoreflect.Message {
	mi := &file_mapped_gateway_edgecontrol_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MappingKeyAndValue.ProtoReflect.Descriptor instead.
func (*MappingKeyAndValue) Descriptor() ([]byte, []int) {
	return file_mapped_gateway_edgecontrol_proto_rawDescGZIP(), []int{2}
}

func (x *MappingKeyAndValue) GetMappingKey() string {
	if x != nil {
		return x.MappingKey
	}
	return ""
}

func (x *MappingKeyAndValue) GetValue() *types.TypedValue {
	if x != nil {
		return x.Value
	}
	return nil
}

type MappingKeyAndResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The mappingKey where the this result originated from
	MappingKey string `protobuf:"bytes,1,opt,name=mapping_key,json=mappingKey,proto3" json:"mapping_key,omitempty"`
	// The result code
	Result types.ControlResult `protobuf:"varint,2,opt,name=result,proto3,enum=mapped.cloud.types.ControlResult" json:"result,omitempty"`
	// Error description (if available)
	ErrorDescription string `protobuf:"bytes,3,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	// Protocol-specific error class, for protocols that split class/code - like BACnet
	ProtocolErrorClass int64 `protobuf:"varint,4,opt,name=protocol_error_class,json=protocolErrorClass,proto3" json:"protocol_error_class,omitempty"`
	// Protocol-specific error code
	ProtocolErrorCode int64 `protobuf:"varint,5,opt,name=protocol_error_code,json=protocolErrorCode,proto3" json:"protocol_error_code,omitempty"`
}

func (x *MappingKeyAndResult) Reset() {
	*x = MappingKeyAndResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mapped_gateway_edgecontrol_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MappingKeyAndResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MappingKeyAndResult) ProtoMessage() {}

func (x *MappingKeyAndResult) ProtoReflect() protoreflect.Message {
	mi := &file_mapped_gateway_edgecontrol_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MappingKeyAndResult.ProtoReflect.Descriptor instead.
func (*MappingKeyAndResult) Descriptor() ([]byte, []int) {
	return file_mapped_gateway_edgecontrol_proto_rawDescGZIP(), []int{3}
}

func (x *MappingKeyAndResult) GetMappingKey() string {
	if x != nil {
		return x.MappingKey
	}
	return ""
}

func (x *MappingKeyAndResult) GetResult() types.ControlResult {
	if x != nil {
		return x.Result
	}
	return types.ControlResult(0)
}

func (x *MappingKeyAndResult) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

func (x *MappingKeyAndResult) GetProtocolErrorClass() int64 {
	if x != nil {
		return x.ProtocolErrorClass
	}
	return 0
}

func (x *MappingKeyAndResult) GetProtocolErrorCode() int64 {
	if x != nil {
		return x.ProtocolErrorCode
	}
	return 0
}

var File_mapped_gateway_edgecontrol_proto protoreflect.FileDescriptor

var file_mapped_gateway_edgecontrol_proto_rawDesc = []byte{
	0x0a, 0x20, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2f, 0x65, 0x64, 0x67, 0x65, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x1a, 0x24, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa3, 0x01, 0x0a, 0x11, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x3a, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x41, 0x6e, 0x64, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x36, 0x0a, 0x18,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x14,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x50, 0x65, 0x72, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x49, 0x6e, 0x4d, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x18, 0x64, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x73, 0x69, 0x6d, 0x75, 0x6c, 0x61, 0x74, 0x65,
	0x22, 0x72, 0x0a, 0x12, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x3d, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4b,
	0x65, 0x79, 0x41, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x22, 0x6b, 0x0a, 0x12, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4b,
	0x65, 0x79, 0x41, 0x6e, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61,
	0x70, 0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x34, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6d, 0x61, 0x70,
	0x70, 0x65, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x54, 0x79, 0x70, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x80, 0x02, 0x0a, 0x13, 0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79,
	0x41, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x61, 0x70,
	0x70, 0x69, 0x6e, 0x67, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x6d, 0x61, 0x70,
	0x70, 0x65, 0x64, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x30, 0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x43, 0x6f, 0x64, 0x65, 0x32, 0x6b, 0x0a, 0x12, 0x45, 0x64, 0x67, 0x65, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0a, 0x57, 0x72,
	0x69, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x12, 0x21, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65,
	0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6d, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x49, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x61, 0x70, 0x70, 0x65, 0x64, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x50, 0x01, 0x5a, 0x20, 0x67, 0x6f, 0x2e, 0x6d, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x3b, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0xaa, 0x02, 0x0e, 0x4d, 0x61,
	0x70, 0x70, 0x65, 0x64, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mapped_gateway_edgecontrol_proto_rawDescOnce sync.Once
	file_mapped_gateway_edgecontrol_proto_rawDescData = file_mapped_gateway_edgecontrol_proto_rawDesc
)

func file_mapped_gateway_edgecontrol_proto_rawDescGZIP() []byte {
	file_mapped_gateway_edgecontrol_proto_rawDescOnce.Do(func() {
		file_mapped_gateway_edgecontrol_proto_rawDescData = protoimpl.X.CompressGZIP(file_mapped_gateway_edgecontrol_proto_rawDescData)
	})
	return file_mapped_gateway_edgecontrol_proto_rawDescData
}

var file_mapped_gateway_edgecontrol_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mapped_gateway_edgecontrol_proto_goTypes = []interface{}{
	(*WritePropsRequest)(nil),   // 0: mapped.gateway.WritePropsRequest
	(*WritePropsResponse)(nil),  // 1: mapped.gateway.WritePropsResponse
	(*MappingKeyAndValue)(nil),  // 2: mapped.gateway.MappingKeyAndValue
	(*MappingKeyAndResult)(nil), // 3: mapped.gateway.MappingKeyAndResult
	(*types.TypedValue)(nil),    // 4: mapped.cloud.types.TypedValue
	(types.ControlResult)(0),    // 5: mapped.cloud.types.ControlResult
}
var file_mapped_gateway_edgecontrol_proto_depIdxs = []int32{
	2, // 0: mapped.gateway.WritePropsRequest.values:type_name -> mapped.gateway.MappingKeyAndValue
	3, // 1: mapped.gateway.WritePropsResponse.results:type_name -> mapped.gateway.MappingKeyAndResult
	4, // 2: mapped.gateway.MappingKeyAndValue.value:type_name -> mapped.cloud.types.TypedValue
	5, // 3: mapped.gateway.MappingKeyAndResult.result:type_name -> mapped.cloud.types.ControlResult
	0, // 4: mapped.gateway.EdgeControlService.WriteProps:input_type -> mapped.gateway.WritePropsRequest
	1, // 5: mapped.gateway.EdgeControlService.WriteProps:output_type -> mapped.gateway.WritePropsResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_mapped_gateway_edgecontrol_proto_init() }
func file_mapped_gateway_edgecontrol_proto_init() {
	if File_mapped_gateway_edgecontrol_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mapped_gateway_edgecontrol_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WritePropsRequest); i {
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
		file_mapped_gateway_edgecontrol_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WritePropsResponse); i {
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
		file_mapped_gateway_edgecontrol_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MappingKeyAndValue); i {
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
		file_mapped_gateway_edgecontrol_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MappingKeyAndResult); i {
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
			RawDescriptor: file_mapped_gateway_edgecontrol_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mapped_gateway_edgecontrol_proto_goTypes,
		DependencyIndexes: file_mapped_gateway_edgecontrol_proto_depIdxs,
		MessageInfos:      file_mapped_gateway_edgecontrol_proto_msgTypes,
	}.Build()
	File_mapped_gateway_edgecontrol_proto = out.File
	file_mapped_gateway_edgecontrol_proto_rawDesc = nil
	file_mapped_gateway_edgecontrol_proto_goTypes = nil
	file_mapped_gateway_edgecontrol_proto_depIdxs = nil
}
