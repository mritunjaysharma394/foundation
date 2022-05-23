// Copyright 2022 Namespace Labs Inc; All rights reserved.
// Licensed under the EARLY ACCESS SOFTWARE LICENSE AGREEMENT
// available at http://github.com/namespacelabs/foundation

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: universe/db/postgres/database.proto

package postgres

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "namespacelabs.dev/foundation/std/proto"
	types "namespacelabs.dev/foundation/std/types"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port    uint32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_database_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_database_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_database_proto_rawDescGZIP(), []int{0}
}

func (x *Endpoint) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Endpoint) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	SchemaFile *types.Resource `protobuf:"bytes,2,opt,name=schema_file,json=schemaFile,proto3" json:"schema_file,omitempty"`
	HostedAt   *Endpoint       `protobuf:"bytes,3,opt,name=hosted_at,json=hostedAt,proto3" json:"hosted_at,omitempty"`
}

func (x *Database) Reset() {
	*x = Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_database_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Database) ProtoMessage() {}

func (x *Database) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_database_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Database.ProtoReflect.Descriptor instead.
func (*Database) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_database_proto_rawDescGZIP(), []int{1}
}

func (x *Database) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Database) GetSchemaFile() *types.Resource {
	if x != nil {
		return x.SchemaFile
	}
	return nil
}

func (x *Database) GetHostedAt() *Endpoint {
	if x != nil {
		return x.HostedAt
	}
	return nil
}

type WireDatabaseArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WireDatabaseArgs) Reset() {
	*x = WireDatabaseArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_database_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WireDatabaseArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WireDatabaseArgs) ProtoMessage() {}

func (x *WireDatabaseArgs) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_database_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WireDatabaseArgs.ProtoReflect.Descriptor instead.
func (*WireDatabaseArgs) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_database_proto_rawDescGZIP(), []int{2}
}

// This type represents a fully formed database. This is an internal type, used
// for internal configuration.
type InstantiatedDatabase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PackageName string                            `protobuf:"bytes,1,opt,name=package_name,json=packageName,proto3" json:"package_name,omitempty"`
	Credentials *InstantiatedDatabase_Credentials `protobuf:"bytes,2,opt,name=credentials,proto3" json:"credentials,omitempty"`
	Database    []*Database                       `protobuf:"bytes,3,rep,name=database,proto3" json:"database,omitempty"`
}

func (x *InstantiatedDatabase) Reset() {
	*x = InstantiatedDatabase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_database_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantiatedDatabase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantiatedDatabase) ProtoMessage() {}

func (x *InstantiatedDatabase) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_database_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantiatedDatabase.ProtoReflect.Descriptor instead.
func (*InstantiatedDatabase) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_database_proto_rawDescGZIP(), []int{3}
}

func (x *InstantiatedDatabase) GetPackageName() string {
	if x != nil {
		return x.PackageName
	}
	return ""
}

func (x *InstantiatedDatabase) GetCredentials() *InstantiatedDatabase_Credentials {
	if x != nil {
		return x.Credentials
	}
	return nil
}

func (x *InstantiatedDatabase) GetDatabase() []*Database {
	if x != nil {
		return x.Database
	}
	return nil
}

type InstantiatedDatabases struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Instantiated []*InstantiatedDatabase `protobuf:"bytes,1,rep,name=instantiated,proto3" json:"instantiated,omitempty"`
}

func (x *InstantiatedDatabases) Reset() {
	*x = InstantiatedDatabases{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_database_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantiatedDatabases) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantiatedDatabases) ProtoMessage() {}

func (x *InstantiatedDatabases) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_database_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantiatedDatabases.ProtoReflect.Descriptor instead.
func (*InstantiatedDatabases) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_database_proto_rawDescGZIP(), []int{4}
}

func (x *InstantiatedDatabases) GetInstantiated() []*InstantiatedDatabase {
	if x != nil {
		return x.Instantiated
	}
	return nil
}

type InstantiatedDatabase_Credentials struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SecretName         string `protobuf:"bytes,1,opt,name=secret_name,json=secretName,proto3" json:"secret_name,omitempty"`
	SecretMountPath    string `protobuf:"bytes,2,opt,name=secret_mount_path,json=secretMountPath,proto3" json:"secret_mount_path,omitempty"`
	SecretResourceName string `protobuf:"bytes,3,opt,name=secret_resource_name,json=secretResourceName,proto3" json:"secret_resource_name,omitempty"`
}

func (x *InstantiatedDatabase_Credentials) Reset() {
	*x = InstantiatedDatabase_Credentials{}
	if protoimpl.UnsafeEnabled {
		mi := &file_universe_db_postgres_database_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InstantiatedDatabase_Credentials) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InstantiatedDatabase_Credentials) ProtoMessage() {}

func (x *InstantiatedDatabase_Credentials) ProtoReflect() protoreflect.Message {
	mi := &file_universe_db_postgres_database_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InstantiatedDatabase_Credentials.ProtoReflect.Descriptor instead.
func (*InstantiatedDatabase_Credentials) Descriptor() ([]byte, []int) {
	return file_universe_db_postgres_database_proto_rawDescGZIP(), []int{3, 0}
}

func (x *InstantiatedDatabase_Credentials) GetSecretName() string {
	if x != nil {
		return x.SecretName
	}
	return ""
}

func (x *InstantiatedDatabase_Credentials) GetSecretMountPath() string {
	if x != nil {
		return x.SecretMountPath
	}
	return ""
}

func (x *InstantiatedDatabase_Credentials) GetSecretResourceName() string {
	if x != nil {
		return x.SecretResourceName
	}
	return ""
}

var File_universe_db_postgres_database_proto protoreflect.FileDescriptor

var file_universe_db_postgres_database_proto_rawDesc = []byte{
	0x0a, 0x23, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x64, 0x62, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x1a, 0x17, 0x73, 0x74, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x18, 0x73, 0x74, 0x64, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x08, 0x45, 0x6e, 0x64,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x22, 0xad, 0x01, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x45, 0x0a, 0x0b, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x66, 0x6f, 0x75, 0x6e,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x73, 0x74, 0x64, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x04, 0x90, 0xa6, 0x1d, 0x01, 0x52,
	0x0a, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x46, 0x0a, 0x09, 0x68,
	0x6f, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x6e, 0x69, 0x76,
	0x65, 0x72, 0x73, 0x65, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73,
	0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x12, 0x0a, 0x10, 0x57, 0x69, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x62,
	0x61, 0x73, 0x65, 0x41, 0x72, 0x67, 0x73, 0x22, 0xf4, 0x02, 0x0a, 0x14, 0x49, 0x6e, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x63, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61,
	0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x41, 0x2e, 0x66, 0x6f, 0x75, 0x6e, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2e, 0x64,
	0x62, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x52, 0x0b, 0x63, 0x72, 0x65,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12, 0x45, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x66, 0x6f, 0x75,
	0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65,
	0x2e, 0x64, 0x62, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2e, 0x44, 0x61, 0x74,
	0x61, 0x62, 0x61, 0x73, 0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a,
	0x8c, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x61, 0x6c, 0x73, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x2a, 0x0a, 0x11, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x65, 0x63,
	0x72, 0x65, 0x74, 0x4d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x61, 0x74, 0x68, 0x12, 0x30, 0x0a, 0x14,
	0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x73, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x72,
	0x0a, 0x15, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61,
	0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x73, 0x12, 0x59, 0x0a, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x35, 0x2e,
	0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x75, 0x6e, 0x69, 0x76, 0x65,
	0x72, 0x73, 0x65, 0x2e, 0x64, 0x62, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x2e,
	0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74,
	0x65, 0x64, 0x42, 0x33, 0x5a, 0x31, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x6c,
	0x61, 0x62, 0x73, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x75, 0x6e, 0x69, 0x76, 0x65, 0x72, 0x73, 0x65, 0x2f, 0x64, 0x62, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x67, 0x72, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_universe_db_postgres_database_proto_rawDescOnce sync.Once
	file_universe_db_postgres_database_proto_rawDescData = file_universe_db_postgres_database_proto_rawDesc
)

func file_universe_db_postgres_database_proto_rawDescGZIP() []byte {
	file_universe_db_postgres_database_proto_rawDescOnce.Do(func() {
		file_universe_db_postgres_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_universe_db_postgres_database_proto_rawDescData)
	})
	return file_universe_db_postgres_database_proto_rawDescData
}

var file_universe_db_postgres_database_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_universe_db_postgres_database_proto_goTypes = []interface{}{
	(*Endpoint)(nil),                         // 0: foundation.universe.db.postgres.Endpoint
	(*Database)(nil),                         // 1: foundation.universe.db.postgres.Database
	(*WireDatabaseArgs)(nil),                 // 2: foundation.universe.db.postgres.WireDatabaseArgs
	(*InstantiatedDatabase)(nil),             // 3: foundation.universe.db.postgres.InstantiatedDatabase
	(*InstantiatedDatabases)(nil),            // 4: foundation.universe.db.postgres.InstantiatedDatabases
	(*InstantiatedDatabase_Credentials)(nil), // 5: foundation.universe.db.postgres.InstantiatedDatabase.Credentials
	(*types.Resource)(nil),                   // 6: foundation.std.types.Resource
}
var file_universe_db_postgres_database_proto_depIdxs = []int32{
	6, // 0: foundation.universe.db.postgres.Database.schema_file:type_name -> foundation.std.types.Resource
	0, // 1: foundation.universe.db.postgres.Database.hosted_at:type_name -> foundation.universe.db.postgres.Endpoint
	5, // 2: foundation.universe.db.postgres.InstantiatedDatabase.credentials:type_name -> foundation.universe.db.postgres.InstantiatedDatabase.Credentials
	1, // 3: foundation.universe.db.postgres.InstantiatedDatabase.database:type_name -> foundation.universe.db.postgres.Database
	3, // 4: foundation.universe.db.postgres.InstantiatedDatabases.instantiated:type_name -> foundation.universe.db.postgres.InstantiatedDatabase
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_universe_db_postgres_database_proto_init() }
func file_universe_db_postgres_database_proto_init() {
	if File_universe_db_postgres_database_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_universe_db_postgres_database_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
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
		file_universe_db_postgres_database_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Database); i {
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
		file_universe_db_postgres_database_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WireDatabaseArgs); i {
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
		file_universe_db_postgres_database_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantiatedDatabase); i {
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
		file_universe_db_postgres_database_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantiatedDatabases); i {
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
		file_universe_db_postgres_database_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InstantiatedDatabase_Credentials); i {
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
			RawDescriptor: file_universe_db_postgres_database_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_universe_db_postgres_database_proto_goTypes,
		DependencyIndexes: file_universe_db_postgres_database_proto_depIdxs,
		MessageInfos:      file_universe_db_postgres_database_proto_msgTypes,
	}.Build()
	File_universe_db_postgres_database_proto = out.File
	file_universe_db_postgres_database_proto_rawDesc = nil
	file_universe_db_postgres_database_proto_goTypes = nil
	file_universe_db_postgres_database_proto_depIdxs = nil
}
