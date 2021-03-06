// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.0
// source: proto/database.proto

package protobuf

import (
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	any "github.com/golang/protobuf/ptypes/any"
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

type OnDelete int32

const (
	OnDelete_CASCADE     OnDelete = 0
	OnDelete_SET_NULL    OnDelete = 1
	OnDelete_SET_DEFAULT OnDelete = 2
	OnDelete_NOTING      OnDelete = 3
)

// Enum value maps for OnDelete.
var (
	OnDelete_name = map[int32]string{
		0: "CASCADE",
		1: "SET_NULL",
		2: "SET_DEFAULT",
		3: "NOTING",
	}
	OnDelete_value = map[string]int32{
		"CASCADE":     0,
		"SET_NULL":    1,
		"SET_DEFAULT": 2,
		"NOTING":      3,
	}
)

func (x OnDelete) Enum() *OnDelete {
	p := new(OnDelete)
	*p = x
	return p
}

func (x OnDelete) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (OnDelete) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_database_proto_enumTypes[0].Descriptor()
}

func (OnDelete) Type() protoreflect.EnumType {
	return &file_proto_database_proto_enumTypes[0]
}

func (x OnDelete) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use OnDelete.Descriptor instead.
func (OnDelete) EnumDescriptor() ([]byte, []int) {
	return file_proto_database_proto_rawDescGZIP(), []int{0}
}

var file_proto_database_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         53001,
		Name:          "protobuf.is_table",
		Tag:           "varint,53001,opt,name=is_table",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         53002,
		Name:          "protobuf.table_name",
		Tag:           "bytes,53002,opt,name=table_name",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         53004,
		Name:          "protobuf.junction_table",
		Tag:           "varint,53004,opt,name=junction_table",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         53005,
		Name:          "protobuf.unique_together",
		Tag:           "bytes,53005,rep,name=unique_together",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         53006,
		Name:          "protobuf.get_by_fields",
		Tag:           "bytes,53006,rep,name=get_by_fields",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         53007,
		Name:          "protobuf.update_by_fields",
		Tag:           "bytes,53007,rep,name=update_by_fields",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.MessageOptions)(nil),
		ExtensionType: ([]string)(nil),
		Field:         53008,
		Name:          "protobuf.delete_by_fields",
		Tag:           "bytes,53008,rep,name=delete_by_fields",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54001,
		Name:          "protobuf.nullable",
		Tag:           "varint,54001,opt,name=nullable",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         54002,
		Name:          "protobuf.db_tags",
		Tag:           "bytes,54002,opt,name=db_tags",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         54003,
		Name:          "protobuf.json_tags",
		Tag:           "bytes,54003,opt,name=json_tags",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         54004,
		Name:          "protobuf.custom_type",
		Tag:           "bytes,54004,opt,name=custom_type",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         54005,
		Name:          "protobuf.custom_name",
		Tag:           "bytes,54005,opt,name=custom_name",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*any.Any)(nil),
		Field:         54006,
		Name:          "protobuf.default_value",
		Tag:           "bytes,54006,opt,name=default_value",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         54007,
		Name:          "protobuf.foreign_table",
		Tag:           "bytes,54007,opt,name=foreign_table",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         54008,
		Name:          "protobuf.foreign_field",
		Tag:           "bytes,54008,opt,name=foreign_field",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*OnDelete)(nil),
		Field:         54009,
		Name:          "protobuf.foreign_on_delete",
		Tag:           "varint,54009,opt,name=foreign_on_delete,enum=protobuf.OnDelete",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54010,
		Name:          "protobuf.unique",
		Tag:           "varint,54010,opt,name=unique",
		Filename:      "proto/database.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         54011,
		Name:          "protobuf.auto",
		Tag:           "varint,54011,opt,name=auto",
		Filename:      "proto/database.proto",
	},
}

// Extension fields to descriptor.MessageOptions.
var (
	// optional bool is_table = 53001;
	E_IsTable = &file_proto_database_proto_extTypes[0]
	// optional string table_name = 53002;
	E_TableName = &file_proto_database_proto_extTypes[1]
	// optional bool junction_table = 53004;
	E_JunctionTable = &file_proto_database_proto_extTypes[2]
	// repeated string unique_together = 53005;
	E_UniqueTogether = &file_proto_database_proto_extTypes[3]
	// repeated string get_by_fields = 53006;
	E_GetByFields = &file_proto_database_proto_extTypes[4]
	// repeated string update_by_fields = 53007;
	E_UpdateByFields = &file_proto_database_proto_extTypes[5]
	// repeated string delete_by_fields = 53008;
	E_DeleteByFields = &file_proto_database_proto_extTypes[6]
)

// Extension fields to descriptor.FieldOptions.
var (
	// optional bool nullable = 54001;
	E_Nullable = &file_proto_database_proto_extTypes[7]
	// optional string db_tags = 54002;
	E_DbTags = &file_proto_database_proto_extTypes[8]
	// optional string json_tags = 54003;
	E_JsonTags = &file_proto_database_proto_extTypes[9]
	// optional string custom_type = 54004;
	E_CustomType = &file_proto_database_proto_extTypes[10]
	// optional string custom_name = 54005;
	E_CustomName = &file_proto_database_proto_extTypes[11]
	// optional google.protobuf.Any default_value = 54006;
	E_DefaultValue = &file_proto_database_proto_extTypes[12]
	// optional string foreign_table = 54007;
	E_ForeignTable = &file_proto_database_proto_extTypes[13]
	// optional string foreign_field = 54008;
	E_ForeignField = &file_proto_database_proto_extTypes[14]
	// optional protobuf.OnDelete foreign_on_delete = 54009;
	E_ForeignOnDelete = &file_proto_database_proto_extTypes[15]
	// optional bool unique = 54010;
	E_Unique = &file_proto_database_proto_extTypes[16]
	// optional bool auto = 54011;
	E_Auto = &file_proto_database_proto_extTypes[17]
)

var File_proto_database_proto protoreflect.FileDescriptor

var file_proto_database_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x42, 0x0a,
	0x08, 0x4f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x43, 0x41, 0x53,
	0x43, 0x41, 0x44, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x45, 0x54, 0x5f, 0x4e, 0x55,
	0x4c, 0x4c, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x54, 0x5f, 0x44, 0x45, 0x46, 0x41,
	0x55, 0x4c, 0x54, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x4e, 0x4f, 0x54, 0x49, 0x4e, 0x47, 0x10,
	0x03, 0x3a, 0x3c, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x89,
	0x9e, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x3a,
	0x40, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8a,
	0x9e, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x3a, 0x48, 0x0a, 0x0e, 0x6a, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8c, 0x9e, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x6a, 0x75,
	0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x3a, 0x4a, 0x0a, 0x0f, 0x75,
	0x6e, 0x69, 0x71, 0x75, 0x65, 0x5f, 0x74, 0x6f, 0x67, 0x65, 0x74, 0x68, 0x65, 0x72, 0x12, 0x1f,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x8d, 0x9e, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65, 0x54,
	0x6f, 0x67, 0x65, 0x74, 0x68, 0x65, 0x72, 0x3a, 0x45, 0x0a, 0x0d, 0x67, 0x65, 0x74, 0x5f, 0x62,
	0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x8e, 0x9e, 0x03, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0b, 0x67, 0x65, 0x74, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x4b,
	0x0a, 0x10, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x73, 0x12, 0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x8f, 0x9e, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x4b, 0x0a, 0x10, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x12,
	0x1f, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x90, 0x9e, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x79, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x73, 0x3a, 0x3b, 0x0a, 0x08, 0x6e, 0x75, 0x6c, 0x6c,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf1, 0xa5, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6e, 0x75, 0x6c,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x3a, 0x38, 0x0a, 0x07, 0x64, 0x62, 0x5f, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xf2, 0xa5, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x62, 0x54, 0x61, 0x67, 0x73, 0x3a,
	0x3c, 0x0a, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x74, 0x61, 0x67, 0x73, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf3, 0xa5, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6a, 0x73, 0x6f, 0x6e, 0x54, 0x61, 0x67, 0x73, 0x3a, 0x40, 0x0a,
	0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf4, 0xa5, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x3a,
	0x40, 0x0a, 0x0b, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf5, 0xa5,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x3a, 0x5a, 0x0a, 0x0d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xf6, 0xa5, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x44, 0x0a,
	0x0d, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf7, 0xa5,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x3a, 0x44, 0x0a, 0x0d, 0x66, 0x6f, 0x72, 0x65, 0x69, 0x67, 0x6e, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xf8, 0xa5, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x66, 0x6f, 0x72,
	0x65, 0x69, 0x67, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x3a, 0x5f, 0x0a, 0x11, 0x66, 0x6f, 0x72,
	0x65, 0x69, 0x67, 0x6e, 0x5f, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1d,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf9, 0xa5,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x0f, 0x66, 0x6f, 0x72, 0x65, 0x69,
	0x67, 0x6e, 0x4f, 0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x3a, 0x37, 0x0a, 0x06, 0x75, 0x6e,
	0x69, 0x71, 0x75, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xfa, 0xa5, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x75, 0x6e, 0x69,
	0x71, 0x75, 0x65, 0x3a, 0x33, 0x0a, 0x04, 0x61, 0x75, 0x74, 0x6f, 0x12, 0x1d, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xfb, 0xa5, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x04, 0x61, 0x75, 0x74, 0x6f, 0x42, 0x45, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2d, 0x74, 0x69, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x42, 0x09, 0x74, 0x69, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e,
	0x67, 0x2d, 0x74, 0x69, 0x72, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_database_proto_rawDescOnce sync.Once
	file_proto_database_proto_rawDescData = file_proto_database_proto_rawDesc
)

func file_proto_database_proto_rawDescGZIP() []byte {
	file_proto_database_proto_rawDescOnce.Do(func() {
		file_proto_database_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_database_proto_rawDescData)
	})
	return file_proto_database_proto_rawDescData
}

var file_proto_database_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_database_proto_goTypes = []interface{}{
	(OnDelete)(0),                     // 0: protobuf.OnDelete
	(*descriptor.MessageOptions)(nil), // 1: google.protobuf.MessageOptions
	(*descriptor.FieldOptions)(nil),   // 2: google.protobuf.FieldOptions
	(*any.Any)(nil),                   // 3: google.protobuf.Any
}
var file_proto_database_proto_depIdxs = []int32{
	1,  // 0: protobuf.is_table:extendee -> google.protobuf.MessageOptions
	1,  // 1: protobuf.table_name:extendee -> google.protobuf.MessageOptions
	1,  // 2: protobuf.junction_table:extendee -> google.protobuf.MessageOptions
	1,  // 3: protobuf.unique_together:extendee -> google.protobuf.MessageOptions
	1,  // 4: protobuf.get_by_fields:extendee -> google.protobuf.MessageOptions
	1,  // 5: protobuf.update_by_fields:extendee -> google.protobuf.MessageOptions
	1,  // 6: protobuf.delete_by_fields:extendee -> google.protobuf.MessageOptions
	2,  // 7: protobuf.nullable:extendee -> google.protobuf.FieldOptions
	2,  // 8: protobuf.db_tags:extendee -> google.protobuf.FieldOptions
	2,  // 9: protobuf.json_tags:extendee -> google.protobuf.FieldOptions
	2,  // 10: protobuf.custom_type:extendee -> google.protobuf.FieldOptions
	2,  // 11: protobuf.custom_name:extendee -> google.protobuf.FieldOptions
	2,  // 12: protobuf.default_value:extendee -> google.protobuf.FieldOptions
	2,  // 13: protobuf.foreign_table:extendee -> google.protobuf.FieldOptions
	2,  // 14: protobuf.foreign_field:extendee -> google.protobuf.FieldOptions
	2,  // 15: protobuf.foreign_on_delete:extendee -> google.protobuf.FieldOptions
	2,  // 16: protobuf.unique:extendee -> google.protobuf.FieldOptions
	2,  // 17: protobuf.auto:extendee -> google.protobuf.FieldOptions
	3,  // 18: protobuf.default_value:type_name -> google.protobuf.Any
	0,  // 19: protobuf.foreign_on_delete:type_name -> protobuf.OnDelete
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	18, // [18:20] is the sub-list for extension type_name
	0,  // [0:18] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_proto_database_proto_init() }
func file_proto_database_proto_init() {
	if File_proto_database_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_database_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 18,
			NumServices:   0,
		},
		GoTypes:           file_proto_database_proto_goTypes,
		DependencyIndexes: file_proto_database_proto_depIdxs,
		EnumInfos:         file_proto_database_proto_enumTypes,
		ExtensionInfos:    file_proto_database_proto_extTypes,
	}.Build()
	File_proto_database_proto = out.File
	file_proto_database_proto_rawDesc = nil
	file_proto_database_proto_goTypes = nil
	file_proto_database_proto_depIdxs = nil
}
