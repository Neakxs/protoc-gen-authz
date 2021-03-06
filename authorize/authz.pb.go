// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: authorize/authz.proto

package authorize

import (
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

type FileRule_Overloads_Type_Primitive int32

const (
	FileRule_Overloads_Type_PRIMITIVE_UNSPECIFIED FileRule_Overloads_Type_Primitive = 0
	FileRule_Overloads_Type_BOOL                  FileRule_Overloads_Type_Primitive = 1
	FileRule_Overloads_Type_INT                   FileRule_Overloads_Type_Primitive = 2
	FileRule_Overloads_Type_UINT                  FileRule_Overloads_Type_Primitive = 3
	FileRule_Overloads_Type_DOUBLE                FileRule_Overloads_Type_Primitive = 4
	FileRule_Overloads_Type_BYTES                 FileRule_Overloads_Type_Primitive = 5
	FileRule_Overloads_Type_STRING                FileRule_Overloads_Type_Primitive = 6
	FileRule_Overloads_Type_DURATION              FileRule_Overloads_Type_Primitive = 7
	FileRule_Overloads_Type_TIMESTAMP             FileRule_Overloads_Type_Primitive = 8
	FileRule_Overloads_Type_ERROR                 FileRule_Overloads_Type_Primitive = 9
	FileRule_Overloads_Type_DYN                   FileRule_Overloads_Type_Primitive = 10
	FileRule_Overloads_Type_ANY                   FileRule_Overloads_Type_Primitive = 11
)

// Enum value maps for FileRule_Overloads_Type_Primitive.
var (
	FileRule_Overloads_Type_Primitive_name = map[int32]string{
		0:  "PRIMITIVE_UNSPECIFIED",
		1:  "BOOL",
		2:  "INT",
		3:  "UINT",
		4:  "DOUBLE",
		5:  "BYTES",
		6:  "STRING",
		7:  "DURATION",
		8:  "TIMESTAMP",
		9:  "ERROR",
		10: "DYN",
		11: "ANY",
	}
	FileRule_Overloads_Type_Primitive_value = map[string]int32{
		"PRIMITIVE_UNSPECIFIED": 0,
		"BOOL":                  1,
		"INT":                   2,
		"UINT":                  3,
		"DOUBLE":                4,
		"BYTES":                 5,
		"STRING":                6,
		"DURATION":              7,
		"TIMESTAMP":             8,
		"ERROR":                 9,
		"DYN":                   10,
		"ANY":                   11,
	}
)

func (x FileRule_Overloads_Type_Primitive) Enum() *FileRule_Overloads_Type_Primitive {
	p := new(FileRule_Overloads_Type_Primitive)
	*p = x
	return p
}

func (x FileRule_Overloads_Type_Primitive) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FileRule_Overloads_Type_Primitive) Descriptor() protoreflect.EnumDescriptor {
	return file_authorize_authz_proto_enumTypes[0].Descriptor()
}

func (FileRule_Overloads_Type_Primitive) Type() protoreflect.EnumType {
	return &file_authorize_authz_proto_enumTypes[0]
}

func (x FileRule_Overloads_Type_Primitive) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FileRule_Overloads_Type_Primitive.Descriptor instead.
func (FileRule_Overloads_Type_Primitive) EnumDescriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{0, 1, 0, 0}
}

type FileRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Globals   *FileRule_Globals      `protobuf:"bytes,1,opt,name=globals,proto3" json:"globals,omitempty"`
	Rules     map[string]*MethodRule `protobuf:"bytes,2,rep,name=rules,proto3" json:"rules,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Overloads *FileRule_Overloads    `protobuf:"bytes,3,opt,name=overloads,proto3" json:"overloads,omitempty"`
}

func (x *FileRule) Reset() {
	*x = FileRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorize_authz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRule) ProtoMessage() {}

func (x *FileRule) ProtoReflect() protoreflect.Message {
	mi := &file_authorize_authz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRule.ProtoReflect.Descriptor instead.
func (*FileRule) Descriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{0}
}

func (x *FileRule) GetGlobals() *FileRule_Globals {
	if x != nil {
		return x.Globals
	}
	return nil
}

func (x *FileRule) GetRules() map[string]*MethodRule {
	if x != nil {
		return x.Rules
	}
	return nil
}

func (x *FileRule) GetOverloads() *FileRule_Overloads {
	if x != nil {
		return x.Overloads
	}
	return nil
}

type MethodRule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expr string `protobuf:"bytes,1,opt,name=expr,proto3" json:"expr,omitempty"`
}

func (x *MethodRule) Reset() {
	*x = MethodRule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorize_authz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MethodRule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MethodRule) ProtoMessage() {}

func (x *MethodRule) ProtoReflect() protoreflect.Message {
	mi := &file_authorize_authz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MethodRule.ProtoReflect.Descriptor instead.
func (*MethodRule) Descriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{1}
}

func (x *MethodRule) GetExpr() string {
	if x != nil {
		return x.Expr
	}
	return ""
}

type FileRule_Globals struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Functions map[string]string `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Constants map[string]string `protobuf:"bytes,2,rep,name=constants,proto3" json:"constants,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FileRule_Globals) Reset() {
	*x = FileRule_Globals{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorize_authz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRule_Globals) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRule_Globals) ProtoMessage() {}

func (x *FileRule_Globals) ProtoReflect() protoreflect.Message {
	mi := &file_authorize_authz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRule_Globals.ProtoReflect.Descriptor instead.
func (*FileRule_Globals) Descriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{0, 0}
}

func (x *FileRule_Globals) GetFunctions() map[string]string {
	if x != nil {
		return x.Functions
	}
	return nil
}

func (x *FileRule_Globals) GetConstants() map[string]string {
	if x != nil {
		return x.Constants
	}
	return nil
}

type FileRule_Overloads struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Functions map[string]*FileRule_Overloads_Function `protobuf:"bytes,1,rep,name=functions,proto3" json:"functions,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Variables map[string]*FileRule_Overloads_Type     `protobuf:"bytes,2,rep,name=variables,proto3" json:"variables,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *FileRule_Overloads) Reset() {
	*x = FileRule_Overloads{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorize_authz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRule_Overloads) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRule_Overloads) ProtoMessage() {}

func (x *FileRule_Overloads) ProtoReflect() protoreflect.Message {
	mi := &file_authorize_authz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRule_Overloads.ProtoReflect.Descriptor instead.
func (*FileRule_Overloads) Descriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{0, 1}
}

func (x *FileRule_Overloads) GetFunctions() map[string]*FileRule_Overloads_Function {
	if x != nil {
		return x.Functions
	}
	return nil
}

func (x *FileRule_Overloads) GetVariables() map[string]*FileRule_Overloads_Type {
	if x != nil {
		return x.Variables
	}
	return nil
}

type FileRule_Overloads_Type struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Type:
	//	*FileRule_Overloads_Type_Primitive_
	//	*FileRule_Overloads_Type_Object
	//	*FileRule_Overloads_Type_Array_
	//	*FileRule_Overloads_Type_Map_
	Type isFileRule_Overloads_Type_Type `protobuf_oneof:"type"`
}

func (x *FileRule_Overloads_Type) Reset() {
	*x = FileRule_Overloads_Type{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorize_authz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRule_Overloads_Type) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRule_Overloads_Type) ProtoMessage() {}

func (x *FileRule_Overloads_Type) ProtoReflect() protoreflect.Message {
	mi := &file_authorize_authz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRule_Overloads_Type.ProtoReflect.Descriptor instead.
func (*FileRule_Overloads_Type) Descriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (m *FileRule_Overloads_Type) GetType() isFileRule_Overloads_Type_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (x *FileRule_Overloads_Type) GetPrimitive() FileRule_Overloads_Type_Primitive {
	if x, ok := x.GetType().(*FileRule_Overloads_Type_Primitive_); ok {
		return x.Primitive
	}
	return FileRule_Overloads_Type_PRIMITIVE_UNSPECIFIED
}

func (x *FileRule_Overloads_Type) GetObject() string {
	if x, ok := x.GetType().(*FileRule_Overloads_Type_Object); ok {
		return x.Object
	}
	return ""
}

func (x *FileRule_Overloads_Type) GetArray() *FileRule_Overloads_Type_Array {
	if x, ok := x.GetType().(*FileRule_Overloads_Type_Array_); ok {
		return x.Array
	}
	return nil
}

func (x *FileRule_Overloads_Type) GetMap() *FileRule_Overloads_Type_Map {
	if x, ok := x.GetType().(*FileRule_Overloads_Type_Map_); ok {
		return x.Map
	}
	return nil
}

type isFileRule_Overloads_Type_Type interface {
	isFileRule_Overloads_Type_Type()
}

type FileRule_Overloads_Type_Primitive_ struct {
	Primitive FileRule_Overloads_Type_Primitive `protobuf:"varint,1,opt,name=primitive,proto3,enum=authorize.FileRule_Overloads_Type_Primitive,oneof"`
}

type FileRule_Overloads_Type_Object struct {
	Object string `protobuf:"bytes,2,opt,name=object,proto3,oneof"`
}

type FileRule_Overloads_Type_Array_ struct {
	Array *FileRule_Overloads_Type_Array `protobuf:"bytes,3,opt,name=array,proto3,oneof"`
}

type FileRule_Overloads_Type_Map_ struct {
	Map *FileRule_Overloads_Type_Map `protobuf:"bytes,4,opt,name=map,proto3,oneof"`
}

func (*FileRule_Overloads_Type_Primitive_) isFileRule_Overloads_Type_Type() {}

func (*FileRule_Overloads_Type_Object) isFileRule_Overloads_Type_Type() {}

func (*FileRule_Overloads_Type_Array_) isFileRule_Overloads_Type_Type() {}

func (*FileRule_Overloads_Type_Map_) isFileRule_Overloads_Type_Type() {}

type FileRule_Overloads_Function struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Args   []*FileRule_Overloads_Type `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	Result *FileRule_Overloads_Type   `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *FileRule_Overloads_Function) Reset() {
	*x = FileRule_Overloads_Function{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorize_authz_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRule_Overloads_Function) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRule_Overloads_Function) ProtoMessage() {}

func (x *FileRule_Overloads_Function) ProtoReflect() protoreflect.Message {
	mi := &file_authorize_authz_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRule_Overloads_Function.ProtoReflect.Descriptor instead.
func (*FileRule_Overloads_Function) Descriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *FileRule_Overloads_Function) GetArgs() []*FileRule_Overloads_Type {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *FileRule_Overloads_Function) GetResult() *FileRule_Overloads_Type {
	if x != nil {
		return x.Result
	}
	return nil
}

type FileRule_Overloads_Type_Array struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type *FileRule_Overloads_Type `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *FileRule_Overloads_Type_Array) Reset() {
	*x = FileRule_Overloads_Type_Array{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorize_authz_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRule_Overloads_Type_Array) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRule_Overloads_Type_Array) ProtoMessage() {}

func (x *FileRule_Overloads_Type_Array) ProtoReflect() protoreflect.Message {
	mi := &file_authorize_authz_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRule_Overloads_Type_Array.ProtoReflect.Descriptor instead.
func (*FileRule_Overloads_Type_Array) Descriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{0, 1, 0, 0}
}

func (x *FileRule_Overloads_Type_Array) GetType() *FileRule_Overloads_Type {
	if x != nil {
		return x.Type
	}
	return nil
}

type FileRule_Overloads_Type_Map struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   *FileRule_Overloads_Type `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value *FileRule_Overloads_Type `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *FileRule_Overloads_Type_Map) Reset() {
	*x = FileRule_Overloads_Type_Map{}
	if protoimpl.UnsafeEnabled {
		mi := &file_authorize_authz_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileRule_Overloads_Type_Map) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileRule_Overloads_Type_Map) ProtoMessage() {}

func (x *FileRule_Overloads_Type_Map) ProtoReflect() protoreflect.Message {
	mi := &file_authorize_authz_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileRule_Overloads_Type_Map.ProtoReflect.Descriptor instead.
func (*FileRule_Overloads_Type_Map) Descriptor() ([]byte, []int) {
	return file_authorize_authz_proto_rawDescGZIP(), []int{0, 1, 0, 1}
}

func (x *FileRule_Overloads_Type_Map) GetKey() *FileRule_Overloads_Type {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *FileRule_Overloads_Type_Map) GetValue() *FileRule_Overloads_Type {
	if x != nil {
		return x.Value
	}
	return nil
}

var file_authorize_authz_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*FileRule)(nil),
		Field:         1145,
		Name:          "authorize.file",
		Tag:           "bytes,1145,opt,name=file",
		Filename:      "authorize/authz.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*MethodRule)(nil),
		Field:         1145,
		Name:          "authorize.method",
		Tag:           "bytes,1145,opt,name=method",
		Filename:      "authorize/authz.proto",
	},
}

// Extension fields to descriptorpb.FileOptions.
var (
	// optional authorize.FileRule file = 1145;
	E_File = &file_authorize_authz_proto_extTypes[0]
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional authorize.MethodRule method = 1145;
	E_Method = &file_authorize_authz_proto_extTypes[1]
)

var File_authorize_authz_proto protoreflect.FileDescriptor

var file_authorize_authz_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x68,
	0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe1, 0x0c, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c,
	0x65, 0x12, 0x35, 0x0a, 0x07, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x52,
	0x07, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x12, 0x34, 0x0a, 0x05, 0x72, 0x75, 0x6c, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x52, 0x75, 0x6c,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x3b,
	0x0a, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1d, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x52, 0x09, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x1a, 0x99, 0x02, 0x0a, 0x07,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x12, 0x48, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e,
	0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x48, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x47, 0x6c, 0x6f, 0x62, 0x61, 0x6c,
	0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x09, 0x63, 0x6f, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x1a, 0x3c, 0x0a, 0x0e, 0x46,
	0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a,
	0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3c, 0x0a, 0x0e, 0x43, 0x6f, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x74, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xbd, 0x08, 0x0a, 0x09, 0x4f, 0x76, 0x65, 0x72,
	0x6c, 0x6f, 0x61, 0x64, 0x73, 0x12, 0x4a, 0x0a, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x76,
	0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x09, 0x66, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x12, 0x4a, 0x0a, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x2e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x09, 0x76, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x73, 0x1a, 0xcf, 0x04,
	0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x4c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6d, 0x69, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2c, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f,
	0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x50, 0x72,
	0x69, 0x6d, 0x69, 0x74, 0x69, 0x76, 0x65, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6d, 0x69,
	0x74, 0x69, 0x76, 0x65, 0x12, 0x18, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x40,
	0x0a, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x41, 0x72, 0x72, 0x61, 0x79, 0x48, 0x00, 0x52, 0x05, 0x61, 0x72, 0x72, 0x61, 0x79,
	0x12, 0x3a, 0x0a, 0x03, 0x6d, 0x61, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75,
	0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x54, 0x79, 0x70,
	0x65, 0x2e, 0x4d, 0x61, 0x70, 0x48, 0x00, 0x52, 0x03, 0x6d, 0x61, 0x70, 0x1a, 0x3f, 0x0a, 0x05,
	0x41, 0x72, 0x72, 0x61, 0x79, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a, 0x75, 0x0a,
	0x03, 0x4d, 0x61, 0x70, 0x12, 0x34, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f,
	0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x09, 0x50, 0x72, 0x69, 0x6d, 0x69, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x19, 0x0a, 0x15, 0x50, 0x52, 0x49, 0x4d, 0x49, 0x54, 0x49, 0x56, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x49, 0x4e, 0x54, 0x10, 0x02,
	0x12, 0x08, 0x0a, 0x04, 0x55, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x44, 0x4f,
	0x55, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x09, 0x0a, 0x05, 0x42, 0x59, 0x54, 0x45, 0x53, 0x10,
	0x05, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x54, 0x52, 0x49, 0x4e, 0x47, 0x10, 0x06, 0x12, 0x0c, 0x0a,
	0x08, 0x44, 0x55, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x54,
	0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x10, 0x08, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x09, 0x12, 0x07, 0x0a, 0x03, 0x44, 0x59, 0x4e, 0x10, 0x0a, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x4e, 0x59, 0x10, 0x0b, 0x42, 0x06, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x1a,
	0x7e, 0x0a, 0x08, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x36, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f,
	0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61,
	0x64, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x1a,
	0x64, 0x0a, 0x0e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x76, 0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64,
	0x73, 0x2e, 0x46, 0x75, 0x6e, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x60, 0x0a, 0x0e, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x2e, 0x4f, 0x76,
	0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x4f, 0x0a, 0x0a, 0x52, 0x75, 0x6c, 0x65, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x20, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x78, 0x70, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x65, 0x78, 0x70, 0x72, 0x3a, 0x46, 0x0a, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xf9, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x69, 0x7a, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x3a, 0x4e, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1e, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xf9, 0x08, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x2e,
	0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x4e, 0x65, 0x61, 0x6b, 0x78, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x7a, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69,
	0x7a, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_authorize_authz_proto_rawDescOnce sync.Once
	file_authorize_authz_proto_rawDescData = file_authorize_authz_proto_rawDesc
)

func file_authorize_authz_proto_rawDescGZIP() []byte {
	file_authorize_authz_proto_rawDescOnce.Do(func() {
		file_authorize_authz_proto_rawDescData = protoimpl.X.CompressGZIP(file_authorize_authz_proto_rawDescData)
	})
	return file_authorize_authz_proto_rawDescData
}

var file_authorize_authz_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_authorize_authz_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_authorize_authz_proto_goTypes = []interface{}{
	(FileRule_Overloads_Type_Primitive)(0), // 0: authorize.FileRule.Overloads.Type.Primitive
	(*FileRule)(nil),                       // 1: authorize.FileRule
	(*MethodRule)(nil),                     // 2: authorize.MethodRule
	(*FileRule_Globals)(nil),               // 3: authorize.FileRule.Globals
	(*FileRule_Overloads)(nil),             // 4: authorize.FileRule.Overloads
	nil,                                    // 5: authorize.FileRule.RulesEntry
	nil,                                    // 6: authorize.FileRule.Globals.FunctionsEntry
	nil,                                    // 7: authorize.FileRule.Globals.ConstantsEntry
	(*FileRule_Overloads_Type)(nil),        // 8: authorize.FileRule.Overloads.Type
	(*FileRule_Overloads_Function)(nil),    // 9: authorize.FileRule.Overloads.Function
	nil,                                    // 10: authorize.FileRule.Overloads.FunctionsEntry
	nil,                                    // 11: authorize.FileRule.Overloads.VariablesEntry
	(*FileRule_Overloads_Type_Array)(nil),  // 12: authorize.FileRule.Overloads.Type.Array
	(*FileRule_Overloads_Type_Map)(nil),    // 13: authorize.FileRule.Overloads.Type.Map
	(*descriptorpb.FileOptions)(nil),       // 14: google.protobuf.FileOptions
	(*descriptorpb.MethodOptions)(nil),     // 15: google.protobuf.MethodOptions
}
var file_authorize_authz_proto_depIdxs = []int32{
	3,  // 0: authorize.FileRule.globals:type_name -> authorize.FileRule.Globals
	5,  // 1: authorize.FileRule.rules:type_name -> authorize.FileRule.RulesEntry
	4,  // 2: authorize.FileRule.overloads:type_name -> authorize.FileRule.Overloads
	6,  // 3: authorize.FileRule.Globals.functions:type_name -> authorize.FileRule.Globals.FunctionsEntry
	7,  // 4: authorize.FileRule.Globals.constants:type_name -> authorize.FileRule.Globals.ConstantsEntry
	10, // 5: authorize.FileRule.Overloads.functions:type_name -> authorize.FileRule.Overloads.FunctionsEntry
	11, // 6: authorize.FileRule.Overloads.variables:type_name -> authorize.FileRule.Overloads.VariablesEntry
	2,  // 7: authorize.FileRule.RulesEntry.value:type_name -> authorize.MethodRule
	0,  // 8: authorize.FileRule.Overloads.Type.primitive:type_name -> authorize.FileRule.Overloads.Type.Primitive
	12, // 9: authorize.FileRule.Overloads.Type.array:type_name -> authorize.FileRule.Overloads.Type.Array
	13, // 10: authorize.FileRule.Overloads.Type.map:type_name -> authorize.FileRule.Overloads.Type.Map
	8,  // 11: authorize.FileRule.Overloads.Function.args:type_name -> authorize.FileRule.Overloads.Type
	8,  // 12: authorize.FileRule.Overloads.Function.result:type_name -> authorize.FileRule.Overloads.Type
	9,  // 13: authorize.FileRule.Overloads.FunctionsEntry.value:type_name -> authorize.FileRule.Overloads.Function
	8,  // 14: authorize.FileRule.Overloads.VariablesEntry.value:type_name -> authorize.FileRule.Overloads.Type
	8,  // 15: authorize.FileRule.Overloads.Type.Array.type:type_name -> authorize.FileRule.Overloads.Type
	8,  // 16: authorize.FileRule.Overloads.Type.Map.key:type_name -> authorize.FileRule.Overloads.Type
	8,  // 17: authorize.FileRule.Overloads.Type.Map.value:type_name -> authorize.FileRule.Overloads.Type
	14, // 18: authorize.file:extendee -> google.protobuf.FileOptions
	15, // 19: authorize.method:extendee -> google.protobuf.MethodOptions
	1,  // 20: authorize.file:type_name -> authorize.FileRule
	2,  // 21: authorize.method:type_name -> authorize.MethodRule
	22, // [22:22] is the sub-list for method output_type
	22, // [22:22] is the sub-list for method input_type
	20, // [20:22] is the sub-list for extension type_name
	18, // [18:20] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_authorize_authz_proto_init() }
func file_authorize_authz_proto_init() {
	if File_authorize_authz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_authorize_authz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRule); i {
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
		file_authorize_authz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MethodRule); i {
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
		file_authorize_authz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRule_Globals); i {
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
		file_authorize_authz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRule_Overloads); i {
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
		file_authorize_authz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRule_Overloads_Type); i {
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
		file_authorize_authz_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRule_Overloads_Function); i {
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
		file_authorize_authz_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRule_Overloads_Type_Array); i {
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
		file_authorize_authz_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileRule_Overloads_Type_Map); i {
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
	file_authorize_authz_proto_msgTypes[7].OneofWrappers = []interface{}{
		(*FileRule_Overloads_Type_Primitive_)(nil),
		(*FileRule_Overloads_Type_Object)(nil),
		(*FileRule_Overloads_Type_Array_)(nil),
		(*FileRule_Overloads_Type_Map_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_authorize_authz_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   13,
			NumExtensions: 2,
			NumServices:   0,
		},
		GoTypes:           file_authorize_authz_proto_goTypes,
		DependencyIndexes: file_authorize_authz_proto_depIdxs,
		EnumInfos:         file_authorize_authz_proto_enumTypes,
		MessageInfos:      file_authorize_authz_proto_msgTypes,
		ExtensionInfos:    file_authorize_authz_proto_extTypes,
	}.Build()
	File_authorize_authz_proto = out.File
	file_authorize_authz_proto_rawDesc = nil
	file_authorize_authz_proto_goTypes = nil
	file_authorize_authz_proto_depIdxs = nil
}
