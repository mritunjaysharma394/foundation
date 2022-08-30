// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: schema/orchestration/event.proto

package orchestration

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

type Event_ReadyTriState int32

const (
	Event_UNKNOWN   Event_ReadyTriState = 0
	Event_NOT_READY Event_ReadyTriState = 1
	Event_READY     Event_ReadyTriState = 2
)

// Enum value maps for Event_ReadyTriState.
var (
	Event_ReadyTriState_name = map[int32]string{
		0: "UNKNOWN",
		1: "NOT_READY",
		2: "READY",
	}
	Event_ReadyTriState_value = map[string]int32{
		"UNKNOWN":   0,
		"NOT_READY": 1,
		"READY":     2,
	}
)

func (x Event_ReadyTriState) Enum() *Event_ReadyTriState {
	p := new(Event_ReadyTriState)
	*p = x
	return p
}

func (x Event_ReadyTriState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_ReadyTriState) Descriptor() protoreflect.EnumDescriptor {
	return file_schema_orchestration_event_proto_enumTypes[0].Descriptor()
}

func (Event_ReadyTriState) Type() protoreflect.EnumType {
	return &file_schema_orchestration_event_proto_enumTypes[0]
}

func (x Event_ReadyTriState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_ReadyTriState.Descriptor instead.
func (Event_ReadyTriState) EnumDescriptor() ([]byte, []int) {
	return file_schema_orchestration_event_proto_rawDescGZIP(), []int{0, 0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Opaque value that uniquely identifies the resource.
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// A resource identifier that is implementation specified.
	Kind string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	// A human-readable label that describes the resource.
	Category string `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	Scope    string `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`
	// READY after the resource is ready.
	Ready          Event_ReadyTriState `protobuf:"varint,5,opt,name=ready,proto3,enum=foundation.schema.orchestration.Event_ReadyTriState" json:"ready,omitempty"`
	AlreadyExisted bool                `protobuf:"varint,6,opt,name=already_existed,json=alreadyExisted,proto3" json:"already_existed,omitempty"`
	// JSON-serialized implementation-specific metadata.
	ImplMetadata []byte              `protobuf:"bytes,7,opt,name=impl_metadata,json=implMetadata,proto3" json:"impl_metadata,omitempty"`
	WaitStatus   []*Event_WaitStatus `protobuf:"bytes,8,rep,name=wait_status,json=waitStatus,proto3" json:"wait_status,omitempty"`
	WaitDetails  string              `protobuf:"bytes,9,opt,name=wait_details,json=waitDetails,proto3" json:"wait_details,omitempty"`
	// Something like `kubectl -n foobar describe pod quux`
	// XXX move to a runtime/ specific type.
	RuntimeSpecificHelp string `protobuf:"bytes,10,opt,name=runtime_specific_help,json=runtimeSpecificHelp,proto3" json:"runtime_specific_help,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_orchestration_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_schema_orchestration_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_schema_orchestration_event_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetResourceId() string {
	if x != nil {
		return x.ResourceId
	}
	return ""
}

func (x *Event) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *Event) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *Event) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Event) GetReady() Event_ReadyTriState {
	if x != nil {
		return x.Ready
	}
	return Event_UNKNOWN
}

func (x *Event) GetAlreadyExisted() bool {
	if x != nil {
		return x.AlreadyExisted
	}
	return false
}

func (x *Event) GetImplMetadata() []byte {
	if x != nil {
		return x.ImplMetadata
	}
	return nil
}

func (x *Event) GetWaitStatus() []*Event_WaitStatus {
	if x != nil {
		return x.WaitStatus
	}
	return nil
}

func (x *Event) GetWaitDetails() string {
	if x != nil {
		return x.WaitDetails
	}
	return ""
}

func (x *Event) GetRuntimeSpecificHelp() string {
	if x != nil {
		return x.RuntimeSpecificHelp
	}
	return ""
}

type Event_WaitStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string     `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Opaque      *anypb.Any `protobuf:"bytes,2,opt,name=opaque,proto3" json:"opaque,omitempty"`
}

func (x *Event_WaitStatus) Reset() {
	*x = Event_WaitStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_schema_orchestration_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_WaitStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_WaitStatus) ProtoMessage() {}

func (x *Event_WaitStatus) ProtoReflect() protoreflect.Message {
	mi := &file_schema_orchestration_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_WaitStatus.ProtoReflect.Descriptor instead.
func (*Event_WaitStatus) Descriptor() ([]byte, []int) {
	return file_schema_orchestration_event_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Event_WaitStatus) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Event_WaitStatus) GetOpaque() *anypb.Any {
	if x != nil {
		return x.Opaque
	}
	return nil
}

var File_schema_orchestration_event_proto protoreflect.FileDescriptor

var file_schema_orchestration_event_proto_rawDesc = []byte{
	0x0a, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc9,
	0x04, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x63, 0x6f,
	0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12,
	0x4a, 0x0a, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x34,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x72, 0x69, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x72, 0x65, 0x61, 0x64, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x61,
	0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x5f, 0x65, 0x78, 0x69, 0x73, 0x74, 0x65, 0x64, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x61, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6d, 0x70, 0x6c, 0x5f, 0x6d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x69, 0x6d, 0x70,
	0x6c, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x52, 0x0a, 0x0b, 0x77, 0x61, 0x69,
	0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x57, 0x61, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x0a, 0x77, 0x61, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x21, 0x0a,
	0x0c, 0x77, 0x61, 0x69, 0x74, 0x5f, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x61, 0x69, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x12, 0x32, 0x0a, 0x15, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x63, 0x5f, 0x68, 0x65, 0x6c, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x72, 0x75, 0x6e, 0x74, 0x69, 0x6d, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63,
	0x48, 0x65, 0x6c, 0x70, 0x1a, 0x5c, 0x0a, 0x0a, 0x57, 0x61, 0x69, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x0a, 0x06, 0x6f, 0x70, 0x61, 0x71, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x06, 0x6f, 0x70, 0x61, 0x71,
	0x75, 0x65, 0x22, 0x36, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x72, 0x69, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00,
	0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x01, 0x12,
	0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x42, 0x33, 0x5a, 0x31, 0x6e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c, 0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x6d,
	0x61, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_schema_orchestration_event_proto_rawDescOnce sync.Once
	file_schema_orchestration_event_proto_rawDescData = file_schema_orchestration_event_proto_rawDesc
)

func file_schema_orchestration_event_proto_rawDescGZIP() []byte {
	file_schema_orchestration_event_proto_rawDescOnce.Do(func() {
		file_schema_orchestration_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_schema_orchestration_event_proto_rawDescData)
	})
	return file_schema_orchestration_event_proto_rawDescData
}

var file_schema_orchestration_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_schema_orchestration_event_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_schema_orchestration_event_proto_goTypes = []interface{}{
	(Event_ReadyTriState)(0), // 0: foundation.schema.orchestration.Event.ReadyTriState
	(*Event)(nil),            // 1: foundation.schema.orchestration.Event
	(*Event_WaitStatus)(nil), // 2: foundation.schema.orchestration.Event.WaitStatus
	(*anypb.Any)(nil),        // 3: google.protobuf.Any
}
var file_schema_orchestration_event_proto_depIdxs = []int32{
	0, // 0: foundation.schema.orchestration.Event.ready:type_name -> foundation.schema.orchestration.Event.ReadyTriState
	2, // 1: foundation.schema.orchestration.Event.wait_status:type_name -> foundation.schema.orchestration.Event.WaitStatus
	3, // 2: foundation.schema.orchestration.Event.WaitStatus.opaque:type_name -> google.protobuf.Any
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_schema_orchestration_event_proto_init() }
func file_schema_orchestration_event_proto_init() {
	if File_schema_orchestration_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_schema_orchestration_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_schema_orchestration_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_WaitStatus); i {
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
			RawDescriptor: file_schema_orchestration_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_schema_orchestration_event_proto_goTypes,
		DependencyIndexes: file_schema_orchestration_event_proto_depIdxs,
		EnumInfos:         file_schema_orchestration_event_proto_enumTypes,
		MessageInfos:      file_schema_orchestration_event_proto_msgTypes,
	}.Build()
	File_schema_orchestration_event_proto = out.File
	file_schema_orchestration_event_proto_rawDesc = nil
	file_schema_orchestration_event_proto_goTypes = nil
	file_schema_orchestration_event_proto_depIdxs = nil
}