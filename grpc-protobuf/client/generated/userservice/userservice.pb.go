// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.14.0
// source: userservice.proto

package userservice

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

// User represents a user with an id, name, and email
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_userservice_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Request argument for AddUser RPC
type AddUserArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AddUserArg) Reset() {
	*x = AddUserArg{}
	mi := &file_userservice_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddUserArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserArg) ProtoMessage() {}

func (x *AddUserArg) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserArg.ProtoReflect.Descriptor instead.
func (*AddUserArg) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{1}
}

func (x *AddUserArg) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *AddUserArg) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Response returned by AddUser RPC
type AddUserRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AddUserRet) Reset() {
	*x = AddUserRet{}
	mi := &file_userservice_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddUserRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddUserRet) ProtoMessage() {}

func (x *AddUserRet) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddUserRet.ProtoReflect.Descriptor instead.
func (*AddUserRet) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{2}
}

func (x *AddUserRet) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Request argument for GetUser RPC
type GetUserArg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetUserArg) Reset() {
	*x = GetUserArg{}
	mi := &file_userservice_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserArg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserArg) ProtoMessage() {}

func (x *GetUserArg) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserArg.ProtoReflect.Descriptor instead.
func (*GetUserArg) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserArg) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response returned by GetUser RPC
type GetUserRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserRet) Reset() {
	*x = GetUserRet{}
	mi := &file_userservice_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserRet) ProtoMessage() {}

func (x *GetUserRet) ProtoReflect() protoreflect.Message {
	mi := &file_userservice_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserRet.ProtoReflect.Descriptor instead.
func (*GetUserRet) Descriptor() ([]byte, []int) {
	return file_userservice_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserRet) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_userservice_proto protoreflect.FileDescriptor

var file_userservice_proto_rawDesc = []byte{
	0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x40, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x36, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x67,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x41, 0x72, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x74, 0x12, 0x25, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x87, 0x01, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x67, 0x1a,
	0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x74, 0x12, 0x3b, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x41, 0x72, 0x67, 0x1a, 0x17, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x74, 0x42, 0x17, 0x5a, 0x15, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x64, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userservice_proto_rawDescOnce sync.Once
	file_userservice_proto_rawDescData = file_userservice_proto_rawDesc
)

func file_userservice_proto_rawDescGZIP() []byte {
	file_userservice_proto_rawDescOnce.Do(func() {
		file_userservice_proto_rawDescData = protoimpl.X.CompressGZIP(file_userservice_proto_rawDescData)
	})
	return file_userservice_proto_rawDescData
}

var file_userservice_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_userservice_proto_goTypes = []any{
	(*User)(nil),       // 0: userservice.User
	(*AddUserArg)(nil), // 1: userservice.AddUserArg
	(*AddUserRet)(nil), // 2: userservice.AddUserRet
	(*GetUserArg)(nil), // 3: userservice.GetUserArg
	(*GetUserRet)(nil), // 4: userservice.GetUserRet
}
var file_userservice_proto_depIdxs = []int32{
	0, // 0: userservice.GetUserRet.user:type_name -> userservice.User
	1, // 1: userservice.UserService.AddUser:input_type -> userservice.AddUserArg
	3, // 2: userservice.UserService.GetUser:input_type -> userservice.GetUserArg
	2, // 3: userservice.UserService.AddUser:output_type -> userservice.AddUserRet
	4, // 4: userservice.UserService.GetUser:output_type -> userservice.GetUserRet
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_userservice_proto_init() }
func file_userservice_proto_init() {
	if File_userservice_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userservice_proto_goTypes,
		DependencyIndexes: file_userservice_proto_depIdxs,
		MessageInfos:      file_userservice_proto_msgTypes,
	}.Build()
	File_userservice_proto = out.File
	file_userservice_proto_rawDesc = nil
	file_userservice_proto_goTypes = nil
	file_userservice_proto_depIdxs = nil
}
