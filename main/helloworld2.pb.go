// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0-devel
// 	protoc        v3.12.0
// source: helloworld2.proto

package main

import (
	proto "github.com/golang/protobuf/proto"
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

type UserInfoSetOperationOperation int32

const (
	UserInfoSetOperation_set   UserInfoSetOperationOperation = 0
	UserInfoSetOperation_clear UserInfoSetOperationOperation = 1
)

// Enum value maps for UserInfoSetOperationOperation.
var (
	UserInfoSetOperationOperation_name = map[int32]string{
		0: "set",
		1: "clear",
	}
	UserInfoSetOperationOperation_value = map[string]int32{
		"set":   0,
		"clear": 1,
	}
)

func (x UserInfoSetOperationOperation) Enum() *UserInfoSetOperationOperation {
	p := new(UserInfoSetOperationOperation)
	*p = x
	return p
}

func (x UserInfoSetOperationOperation) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInfoSetOperationOperation) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld2_proto_enumTypes[0].Descriptor()
}

func (UserInfoSetOperationOperation) Type() protoreflect.EnumType {
	return &file_helloworld2_proto_enumTypes[0]
}

func (x UserInfoSetOperationOperation) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserInfoSetOperationOperation.Descriptor instead.
func (UserInfoSetOperationOperation) EnumDescriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{2, 0}
}

type UserInfoSetOperationField int32

const (
	UserInfoSetOperation_nickname    UserInfoSetOperationField = 0
	UserInfoSetOperation_description UserInfoSetOperationField = 1
)

// Enum value maps for UserInfoSetOperationField.
var (
	UserInfoSetOperationField_name = map[int32]string{
		0: "nickname",
		1: "description",
	}
	UserInfoSetOperationField_value = map[string]int32{
		"nickname":    0,
		"description": 1,
	}
)

func (x UserInfoSetOperationField) Enum() *UserInfoSetOperationField {
	p := new(UserInfoSetOperationField)
	*p = x
	return p
}

func (x UserInfoSetOperationField) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserInfoSetOperationField) Descriptor() protoreflect.EnumDescriptor {
	return file_helloworld2_proto_enumTypes[1].Descriptor()
}

func (UserInfoSetOperationField) Type() protoreflect.EnumType {
	return &file_helloworld2_proto_enumTypes[1]
}

func (x UserInfoSetOperationField) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserInfoSetOperationField.Descriptor instead.
func (UserInfoSetOperationField) EnumDescriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{2, 1}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserInfoSetOperation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserInfoSetOperation) Reset() {
	*x = UserInfoSetOperation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfoSetOperation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoSetOperation) ProtoMessage() {}

func (x *UserInfoSetOperation) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoSetOperation.ProtoReflect.Descriptor instead.
func (*UserInfoSetOperation) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{2}
}

type LoginArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginArgs) Reset() {
	*x = LoginArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginArgs) ProtoMessage() {}

func (x *LoginArgs) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginArgs.ProtoReflect.Descriptor instead.
func (*LoginArgs) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{3}
}

func (x *LoginArgs) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginArgs) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GetUserInfoArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetUserInfoArgs) Reset() {
	*x = GetUserInfoArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoArgs) ProtoMessage() {}

func (x *GetUserInfoArgs) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoArgs.ProtoReflect.Descriptor instead.
func (*GetUserInfoArgs) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserInfoArgs) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SetUserInfoArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token                 string                  `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserInfoSetOperations []*UserInfoSetOperation `protobuf:"bytes,2,rep,name=userInfoSetOperations,proto3" json:"userInfoSetOperations,omitempty"`
}

func (x *SetUserInfoArgs) Reset() {
	*x = SetUserInfoArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserInfoArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserInfoArgs) ProtoMessage() {}

func (x *SetUserInfoArgs) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserInfoArgs.ProtoReflect.Descriptor instead.
func (*SetUserInfoArgs) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{5}
}

func (x *SetUserInfoArgs) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SetUserInfoArgs) GetUserInfoSetOperations() []*UserInfoSetOperation {
	if x != nil {
		return x.UserInfoSetOperations
	}
	return nil
}

type GetOthersInfoArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOthersInfoArgs) Reset() {
	*x = GetOthersInfoArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOthersInfoArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOthersInfoArgs) ProtoMessage() {}

func (x *GetOthersInfoArgs) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOthersInfoArgs.ProtoReflect.Descriptor instead.
func (*GetOthersInfoArgs) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{6}
}

func (x *GetOthersInfoArgs) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LogoutArgs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *LogoutArgs) Reset() {
	*x = LogoutArgs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutArgs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutArgs) ProtoMessage() {}

func (x *LogoutArgs) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutArgs.ProtoReflect.Descriptor instead.
func (*LogoutArgs) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{7}
}

func (x *LogoutArgs) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{8}
}

type Err struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code uint32 `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *Err) Reset() {
	*x = Err{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Err) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Err) ProtoMessage() {}

func (x *Err) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Err.ProtoReflect.Descriptor instead.
func (*Err) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{9}
}

func (x *Err) GetCode() uint32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Err) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type LoginResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Err   *Err   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *LoginResult) Reset() {
	*x = LoginResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResult) ProtoMessage() {}

func (x *LoginResult) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResult.ProtoReflect.Descriptor instead.
func (*LoginResult) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{10}
}

func (x *LoginResult) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginResult) GetErr() *Err {
	if x != nil {
		return x.Err
	}
	return nil
}

type GetUserInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	Err      *Err      `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetUserInfoResult) Reset() {
	*x = GetUserInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResult) ProtoMessage() {}

func (x *GetUserInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResult.ProtoReflect.Descriptor instead.
func (*GetUserInfoResult) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{11}
}

func (x *GetUserInfoResult) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *GetUserInfoResult) GetErr() *Err {
	if x != nil {
		return x.Err
	}
	return nil
}

type SetUserInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Empty *Empty `protobuf:"bytes,2,opt,name=empty,proto3" json:"empty,omitempty"` //想不出来这个数据段可以填什么，直接在server.go的UserInfos里面改应该就可以了
	Err   *Err   `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SetUserInfoResult) Reset() {
	*x = SetUserInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetUserInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetUserInfoResult) ProtoMessage() {}

func (x *SetUserInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetUserInfoResult.ProtoReflect.Descriptor instead.
func (*SetUserInfoResult) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{12}
}

func (x *SetUserInfoResult) GetEmpty() *Empty {
	if x != nil {
		return x.Empty
	}
	return nil
}

func (x *SetUserInfoResult) GetErr() *Err {
	if x != nil {
		return x.Err
	}
	return nil
}

type GetOthersInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserInfo *UserInfo `protobuf:"bytes,2,opt,name=userInfo,proto3" json:"userInfo,omitempty"`
	Err      *Err      `protobuf:"bytes,3,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetOthersInfoResult) Reset() {
	*x = GetOthersInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOthersInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOthersInfoResult) ProtoMessage() {}

func (x *GetOthersInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOthersInfoResult.ProtoReflect.Descriptor instead.
func (*GetOthersInfoResult) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{13}
}

func (x *GetOthersInfoResult) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *GetOthersInfoResult) GetErr() *Err {
	if x != nil {
		return x.Err
	}
	return nil
}

type LogoutResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err *Err `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *LogoutResult) Reset() {
	*x = LogoutResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_helloworld2_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogoutResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResult) ProtoMessage() {}

func (x *LogoutResult) ProtoReflect() protoreflect.Message {
	mi := &file_helloworld2_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResult.ProtoReflect.Descriptor instead.
func (*LogoutResult) Descriptor() ([]byte, []int) {
	return file_helloworld2_proto_rawDescGZIP(), []int{14}
}

func (x *LogoutResult) GetErr() *Err {
	if x != nil {
		return x.Err
	}
	return nil
}

var File_helloworld2_proto protoreflect.FileDescriptor

var file_helloworld2_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x32, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x1d, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x52, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5f, 0x0a, 0x14,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x1f, 0x0a, 0x09, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x07, 0x0a, 0x03, 0x73, 0x65, 0x74, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x63, 0x6c,
	0x65, 0x61, 0x72, 0x10, 0x01, 0x22, 0x26, 0x0a, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x0c,
	0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x22, 0x43, 0x0a,
	0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x27, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x41, 0x72, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x79, 0x0a, 0x0f, 0x53,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x72, 0x67, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x50, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x15, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x65, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x68,
	0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x72, 0x67, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x22, 0x0a, 0x0a, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x72, 0x67, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x2b, 0x0a, 0x03, 0x45, 0x72, 0x72, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x40, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x5c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45, 0x72, 0x72,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x53, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x21, 0x0a, 0x05, 0x65, 0x6d,
	0x70, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x05, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x1b, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x5e, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x2a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1b, 0x0a,
	0x03, 0x65, 0x72, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x61, 0x69,
	0x6e, 0x2e, 0x45, 0x72, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x2b, 0x0a, 0x0c, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1b, 0x0a, 0x03, 0x65, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x45,
	0x72, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0xa9, 0x02, 0x0a, 0x07, 0x47, 0x72, 0x65, 0x65,
	0x74, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0f, 0x2e, 0x6d,
	0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x11, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x0b, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x1a, 0x17,
	0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x53, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e,
	0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x41, 0x72, 0x67, 0x73, 0x1a,
	0x17, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x17, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f,
	0x41, 0x72, 0x67, 0x73, 0x1a, 0x19, 0x2e, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x74, 0x68, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x30, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x10, 0x2e, 0x6d, 0x61,
	0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x41, 0x72, 0x67, 0x73, 0x1a, 0x12, 0x2e,
	0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_helloworld2_proto_rawDescOnce sync.Once
	file_helloworld2_proto_rawDescData = file_helloworld2_proto_rawDesc
)

func file_helloworld2_proto_rawDescGZIP() []byte {
	file_helloworld2_proto_rawDescOnce.Do(func() {
		file_helloworld2_proto_rawDescData = protoimpl.X.CompressGZIP(file_helloworld2_proto_rawDescData)
	})
	return file_helloworld2_proto_rawDescData
}

var file_helloworld2_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_helloworld2_proto_msgTypes = make([]protoimpl.MessageInfo, 15)
var file_helloworld2_proto_goTypes = []interface{}{
	(UserInfoSetOperationOperation)(0), // 0: main.UserInfoSetOperation.operation
	(UserInfoSetOperationField)(0),     // 1: main.UserInfoSetOperation.field
	(*Token)(nil),                      // 2: main.Token
	(*UserInfo)(nil),                   // 3: main.UserInfo
	(*UserInfoSetOperation)(nil),       // 4: main.UserInfoSetOperation
	(*LoginArgs)(nil),                  // 5: main.LoginArgs
	(*GetUserInfoArgs)(nil),            // 6: main.GetUserInfoArgs
	(*SetUserInfoArgs)(nil),            // 7: main.SetUserInfoArgs
	(*GetOthersInfoArgs)(nil),          // 8: main.GetOthersInfoArgs
	(*LogoutArgs)(nil),                 // 9: main.LogoutArgs
	(*Empty)(nil),                      // 10: main.Empty
	(*Err)(nil),                        // 11: main.Err
	(*LoginResult)(nil),                // 12: main.LoginResult
	(*GetUserInfoResult)(nil),          // 13: main.GetUserInfoResult
	(*SetUserInfoResult)(nil),          // 14: main.SetUserInfoResult
	(*GetOthersInfoResult)(nil),        // 15: main.GetOthersInfoResult
	(*LogoutResult)(nil),               // 16: main.LogoutResult
}
var file_helloworld2_proto_depIdxs = []int32{
	4,  // 0: main.SetUserInfoArgs.userInfoSetOperations:type_name -> main.UserInfoSetOperation
	11, // 1: main.LoginResult.err:type_name -> main.Err
	3,  // 2: main.GetUserInfoResult.userInfo:type_name -> main.UserInfo
	11, // 3: main.GetUserInfoResult.err:type_name -> main.Err
	10, // 4: main.SetUserInfoResult.empty:type_name -> main.Empty
	11, // 5: main.SetUserInfoResult.err:type_name -> main.Err
	3,  // 6: main.GetOthersInfoResult.userInfo:type_name -> main.UserInfo
	11, // 7: main.GetOthersInfoResult.err:type_name -> main.Err
	11, // 8: main.LogoutResult.err:type_name -> main.Err
	5,  // 9: main.Greeter.Login:input_type -> main.LoginArgs
	2,  // 10: main.Greeter.GetUserInfo:input_type -> main.Token
	7,  // 11: main.Greeter.SetUserInfo:input_type -> main.SetUserInfoArgs
	8,  // 12: main.Greeter.GetOthersInfo:input_type -> main.GetOthersInfoArgs
	9,  // 13: main.Greeter.Logout:input_type -> main.LogoutArgs
	12, // 14: main.Greeter.Login:output_type -> main.LoginResult
	13, // 15: main.Greeter.GetUserInfo:output_type -> main.GetUserInfoResult
	14, // 16: main.Greeter.SetUserInfo:output_type -> main.SetUserInfoResult
	15, // 17: main.Greeter.GetOthersInfo:output_type -> main.GetOthersInfoResult
	16, // 18: main.Greeter.Logout:output_type -> main.LogoutResult
	14, // [14:19] is the sub-list for method output_type
	9,  // [9:14] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_helloworld2_proto_init() }
func file_helloworld2_proto_init() {
	if File_helloworld2_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_helloworld2_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_helloworld2_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_helloworld2_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfoSetOperation); i {
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
		file_helloworld2_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginArgs); i {
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
		file_helloworld2_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoArgs); i {
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
		file_helloworld2_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserInfoArgs); i {
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
		file_helloworld2_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOthersInfoArgs); i {
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
		file_helloworld2_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutArgs); i {
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
		file_helloworld2_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_helloworld2_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Err); i {
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
		file_helloworld2_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginResult); i {
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
		file_helloworld2_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserInfoResult); i {
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
		file_helloworld2_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetUserInfoResult); i {
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
		file_helloworld2_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOthersInfoResult); i {
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
		file_helloworld2_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogoutResult); i {
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
			RawDescriptor: file_helloworld2_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   15,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_helloworld2_proto_goTypes,
		DependencyIndexes: file_helloworld2_proto_depIdxs,
		EnumInfos:         file_helloworld2_proto_enumTypes,
		MessageInfos:      file_helloworld2_proto_msgTypes,
	}.Build()
	File_helloworld2_proto = out.File
	file_helloworld2_proto_rawDesc = nil
	file_helloworld2_proto_goTypes = nil
	file_helloworld2_proto_depIdxs = nil
}