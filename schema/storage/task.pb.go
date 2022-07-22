// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/storage/task.proto

package storage

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Next ID: 21
type StoredTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ParentId           string               `protobuf:"bytes,12,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	AnchorId           string               `protobuf:"bytes,13,opt,name=anchor_id,json=anchorId,proto3" json:"anchor_id,omitempty"`
	SpanId             string               `protobuf:"bytes,20,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Name               string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	HumanReadableLabel string               `protobuf:"bytes,3,opt,name=human_readable_label,json=humanReadableLabel,proto3" json:"human_readable_label,omitempty"`
	Category           string               `protobuf:"bytes,19,opt,name=category,proto3" json:"category,omitempty"`
	CreatedTs          int64                `protobuf:"varint,4,opt,name=created_ts,json=createdTs,proto3" json:"created_ts,omitempty"`                   // Unix time in nanoseconds.
	CompletedTs        int64                `protobuf:"varint,5,opt,name=completed_ts,json=completedTs,proto3" json:"completed_ts,omitempty"`             // Unix time in nanoseconds.
	RelStartedTs       int64                `protobuf:"varint,17,opt,name=rel_started_ts,json=relStartedTs,proto3" json:"rel_started_ts,omitempty"`       // started_ts = created_ts + rel_started_ts
	RelCompletedTs     int64                `protobuf:"varint,18,opt,name=rel_completed_ts,json=relCompletedTs,proto3" json:"rel_completed_ts,omitempty"` // completed_ts = created_ts + rel_completed_ts
	Scope              []string             `protobuf:"bytes,7,rep,name=scope,proto3" json:"scope,omitempty"`
	Argument           []*StoredTask_Value  `protobuf:"bytes,8,rep,name=argument,proto3" json:"argument,omitempty"`
	Result             []*StoredTask_Value  `protobuf:"bytes,9,rep,name=result,proto3" json:"result,omitempty"`
	Output             []*StoredTask_Output `protobuf:"bytes,10,rep,name=output,proto3" json:"output,omitempty"`
	Cached             bool                 `protobuf:"varint,11,opt,name=cached,proto3" json:"cached,omitempty"`
	ErrorCode          int32                `protobuf:"varint,15,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	ErrorMessage       string               `protobuf:"bytes,6,opt,name=error_message,json=errorMessage,proto3" json:"error_message,omitempty"` // When completed_ts is set.
	ErrorDetails       []*anypb.Any         `protobuf:"bytes,16,rep,name=error_details,json=errorDetails,proto3" json:"error_details,omitempty"`
	LogLevel           int32                `protobuf:"varint,14,opt,name=log_level,json=logLevel,proto3" json:"log_level,omitempty"`
}

func (x *StoredTask) Reset() {
	*x = StoredTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredTask) ProtoMessage() {}

func (x *StoredTask) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredTask.ProtoReflect.Descriptor instead.
func (*StoredTask) Descriptor() ([]byte, []int) {
	return file_schema_storage_task_proto_rawDescGZIP(), []int{0}
}

func (x *StoredTask) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoredTask) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *StoredTask) GetAnchorId() string {
	if x != nil {
		return x.AnchorId
	}
	return ""
}

func (x *StoredTask) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *StoredTask) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoredTask) GetHumanReadableLabel() string {
	if x != nil {
		return x.HumanReadableLabel
	}
	return ""
}

func (x *StoredTask) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *StoredTask) GetCreatedTs() int64 {
	if x != nil {
		return x.CreatedTs
	}
	return 0
}

func (x *StoredTask) GetCompletedTs() int64 {
	if x != nil {
		return x.CompletedTs
	}
	return 0
}

func (x *StoredTask) GetRelStartedTs() int64 {
	if x != nil {
		return x.RelStartedTs
	}
	return 0
}

func (x *StoredTask) GetRelCompletedTs() int64 {
	if x != nil {
		return x.RelCompletedTs
	}
	return 0
}

func (x *StoredTask) GetScope() []string {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *StoredTask) GetArgument() []*StoredTask_Value {
	if x != nil {
		return x.Argument
	}
	return nil
}

func (x *StoredTask) GetResult() []*StoredTask_Value {
	if x != nil {
		return x.Result
	}
	return nil
}

func (x *StoredTask) GetOutput() []*StoredTask_Output {
	if x != nil {
		return x.Output
	}
	return nil
}

func (x *StoredTask) GetCached() bool {
	if x != nil {
		return x.Cached
	}
	return false
}

func (x *StoredTask) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *StoredTask) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *StoredTask) GetErrorDetails() []*anypb.Any {
	if x != nil {
		return x.ErrorDetails
	}
	return nil
}

func (x *StoredTask) GetLogLevel() int32 {
	if x != nil {
		return x.LogLevel
	}
	return 0
}

type StoredTask_Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key       string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	JsonValue string `protobuf:"bytes,2,opt,name=json_value,json=jsonValue,proto3" json:"json_value,omitempty"`
}

func (x *StoredTask_Value) Reset() {
	*x = StoredTask_Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredTask_Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredTask_Value) ProtoMessage() {}

func (x *StoredTask_Value) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredTask_Value.ProtoReflect.Descriptor instead.
func (*StoredTask_Value) Descriptor() ([]byte, []int) {
	return file_schema_storage_task_proto_rawDescGZIP(), []int{0, 0}
}

func (x *StoredTask_Value) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *StoredTask_Value) GetJsonValue() string {
	if x != nil {
		return x.JsonValue
	}
	return ""
}

type StoredTask_Output struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ContentType string `protobuf:"bytes,3,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *StoredTask_Output) Reset() {
	*x = StoredTask_Output{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_storage_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoredTask_Output) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoredTask_Output) ProtoMessage() {}

func (x *StoredTask_Output) ProtoReflect() protoreflect.Message {
	mi := &file_schema_storage_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoredTask_Output.ProtoReflect.Descriptor instead.
func (*StoredTask_Output) Descriptor() ([]byte, []int) {
	return file_schema_storage_task_proto_rawDescGZIP(), []int{0, 1}
}

func (x *StoredTask_Output) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *StoredTask_Output) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StoredTask_Output) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

var File_schema_storage_task_proto protoreflect.FileDescriptor

var file_schema_storage_task_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8c, 0x07, 0x0a, 0x0a, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x61, 0x6e, 0x63, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x70,
	0x61, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x14, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x70, 0x61,
	0x6e, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x68, 0x75, 0x6d, 0x61, 0x6e,
	0x5f, 0x72, 0x65, 0x61, 0x64, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x68, 0x75, 0x6d, 0x61, 0x6e, 0x52, 0x65, 0x61, 0x64,
	0x61, 0x62, 0x6c, 0x65, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x54, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x6c, 0x5f, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x73, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0c, 0x72, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x65, 0x64, 0x54, 0x73, 0x12, 0x28, 0x0a,
	0x10, 0x72, 0x65, 0x6c, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x74,
	0x73, 0x18, 0x12, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x72, 0x65, 0x6c, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x54, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x47, 0x0a,
	0x08, 0x61, 0x72, 0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2b, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68,
	0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x72,
	0x67, 0x75, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x2e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x44, 0x0a, 0x06, 0x6f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x66, 0x6f,
	0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x2e, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x63, 0x61, 0x63, 0x68, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x39, 0x0a,
	0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x10,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x5f,
	0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x6c, 0x6f, 0x67,
	0x4c, 0x65, 0x76, 0x65, 0x6c, 0x1a, 0x38, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x1d, 0x0a, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a,
	0x4f, 0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x42, 0x2d, 0x5a, 0x2b, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62,
	0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_storage_task_proto_rawDescOnce sync.Once
	file_schema_storage_task_proto_rawDescData = file_schema_storage_task_proto_rawDesc
)

func file_schema_storage_task_proto_rawDescGZIP() []byte {
	file_schema_storage_task_proto_rawDescOnce.Do(func() {
		file_schema_storage_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_storage_task_proto_rawDescData)
	})
	return file_schema_storage_task_proto_rawDescData
}

var file_schema_storage_task_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_schema_storage_task_proto_goTypes = []interface{}{
	(*StoredTask)(nil),        // 0: foundation.schema.storage.StoredTask
	(*StoredTask_Value)(nil),  // 1: foundation.schema.storage.StoredTask.Value
	(*StoredTask_Output)(nil), // 2: foundation.schema.storage.StoredTask.Output
	(*anypb.Any)(nil),         // 3: google.protobuf.Any
}
var file_schema_storage_task_proto_depIdxs = []int32{
	1, // 0: foundation.schema.storage.StoredTask.argument:type_name -> foundation.schema.storage.StoredTask.Value
	1, // 1: foundation.schema.storage.StoredTask.result:type_name -> foundation.schema.storage.StoredTask.Value
	2, // 2: foundation.schema.storage.StoredTask.output:type_name -> foundation.schema.storage.StoredTask.Output
	3, // 3: foundation.schema.storage.StoredTask.error_details:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_schema_storage_task_proto_init() }
func file_schema_storage_task_proto_init() {
	if File_schema_storage_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_storage_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredTask); i {
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
		file_schema_storage_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredTask_Value); i {
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
		file_schema_storage_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoredTask_Output); i {
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
			RawDescriptor: file_schema_storage_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_storage_task_proto_goTypes,
		DependencyIndexes: file_schema_storage_task_proto_depIdxs,
		MessageInfos:      file_schema_storage_task_proto_msgTypes,
	}.Build()
	File_schema_storage_task_proto = out.File
	file_schema_storage_task_proto_rawDesc = nil
	file_schema_storage_task_proto_goTypes = nil
	file_schema_storage_task_proto_depIdxs = nil
}
