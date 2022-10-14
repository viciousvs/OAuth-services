// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.12.4
// source: token.proto

package __

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

//base
type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AtUuid       string `protobuf:"bytes,1,opt,name=at_uuid,json=atUuid,proto3" json:"at_uuid,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	AtExpires    string `protobuf:"bytes,3,opt,name=at_expires,json=atExpires,proto3" json:"at_expires,omitempty"`
	RtUuid       string `protobuf:"bytes,4,opt,name=rt_uuid,json=rtUuid,proto3" json:"rt_uuid,omitempty"`
	RefreshToken string `protobuf:"bytes,5,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	RtExpires    string `protobuf:"bytes,6,opt,name=rt_expires,json=rtExpires,proto3" json:"rt_expires,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

func (x *Tokens) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{0}
}

func (x *Tokens) GetAtUuid() string {
	if x != nil {
		return x.AtUuid
	}
	return ""
}

func (x *Tokens) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *Tokens) GetAtExpires() string {
	if x != nil {
		return x.AtExpires
	}
	return ""
}

func (x *Tokens) GetRtUuid() string {
	if x != nil {
		return x.RtUuid
	}
	return ""
}

func (x *Tokens) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *Tokens) GetRtExpires() string {
	if x != nil {
		return x.RtExpires
	}
	return ""
}

//Request
type GenerateTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *GenerateTokensRequest) Reset() {
	*x = GenerateTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokensRequest) ProtoMessage() {}

func (x *GenerateTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokensRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokensRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokensRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type VerifyAccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *VerifyAccessRequest) Reset() {
	*x = VerifyAccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAccessRequest) ProtoMessage() {}

func (x *VerifyAccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAccessRequest.ProtoReflect.Descriptor instead.
func (*VerifyAccessRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{2}
}

func (x *VerifyAccessRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type RefreshRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshRequest) Reset() {
	*x = RefreshRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshRequest) ProtoMessage() {}

func (x *RefreshRequest) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshRequest.ProtoReflect.Descriptor instead.
func (*RefreshRequest) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

//Response
type GenerateTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens *Tokens `protobuf:"bytes,1,opt,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *GenerateTokensResponse) Reset() {
	*x = GenerateTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokensResponse) ProtoMessage() {}

func (x *GenerateTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokensResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokensResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateTokensResponse) GetTokens() *Tokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

type VerifyAccessResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Header    map[string]string `protobuf:"bytes,1,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Claims    map[string]string `protobuf:"bytes,2,rep,name=claims,proto3" json:"claims,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Signature string            `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *VerifyAccessResponse) Reset() {
	*x = VerifyAccessResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyAccessResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyAccessResponse) ProtoMessage() {}

func (x *VerifyAccessResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyAccessResponse.ProtoReflect.Descriptor instead.
func (*VerifyAccessResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{5}
}

func (x *VerifyAccessResponse) GetHeader() map[string]string {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *VerifyAccessResponse) GetClaims() map[string]string {
	if x != nil {
		return x.Claims
	}
	return nil
}

func (x *VerifyAccessResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

type RefreshResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tokens *Tokens `protobuf:"bytes,1,opt,name=tokens,proto3" json:"tokens,omitempty"`
}

func (x *RefreshResponse) Reset() {
	*x = RefreshResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_token_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshResponse) ProtoMessage() {}

func (x *RefreshResponse) ProtoReflect() protoreflect.Message {
	mi := &file_token_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshResponse.ProtoReflect.Descriptor instead.
func (*RefreshResponse) Descriptor() ([]byte, []int) {
	return file_token_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshResponse) GetTokens() *Tokens {
	if x != nil {
		return x.Tokens
	}
	return nil
}

var File_token_proto protoreflect.FileDescriptor

var file_token_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0xc0, 0x01, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x17, 0x0a, 0x07, 0x61, 0x74, 0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x61,
	0x74, 0x5f, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x61, 0x74, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x74,
	0x5f, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x74, 0x55,
	0x75, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x74, 0x5f, 0x65,
	0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x74,
	0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x15, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x75, 0x75, 0x69, 0x64, 0x22, 0x38, 0x0a, 0x13, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x35,
	0x0a, 0x0e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x25, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x06,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x22, 0xac, 0x02, 0x0a, 0x14, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x3f, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x3f, 0x0a, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x27, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x63, 0x6c, 0x61, 0x69, 0x6d,
	0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x1a,
	0x39, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x43, 0x6c,
	0x61, 0x69, 0x6d, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x38, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x06, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x32,
	0xe6, 0x01, 0x0a, 0x0c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4f, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x49, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x12, 0x1a, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x07,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_token_proto_rawDescOnce sync.Once
	file_token_proto_rawDescData = file_token_proto_rawDesc
)

func file_token_proto_rawDescGZIP() []byte {
	file_token_proto_rawDescOnce.Do(func() {
		file_token_proto_rawDescData = protoimpl.X.CompressGZIP(file_token_proto_rawDescData)
	})
	return file_token_proto_rawDescData
}

var file_token_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_token_proto_goTypes = []interface{}{
	(*Tokens)(nil),                 // 0: token.Tokens
	(*GenerateTokensRequest)(nil),  // 1: token.GenerateTokensRequest
	(*VerifyAccessRequest)(nil),    // 2: token.VerifyAccessRequest
	(*RefreshRequest)(nil),         // 3: token.RefreshRequest
	(*GenerateTokensResponse)(nil), // 4: token.GenerateTokensResponse
	(*VerifyAccessResponse)(nil),   // 5: token.VerifyAccessResponse
	(*RefreshResponse)(nil),        // 6: token.RefreshResponse
	nil,                            // 7: token.VerifyAccessResponse.HeaderEntry
	nil,                            // 8: token.VerifyAccessResponse.ClaimsEntry
}
var file_token_proto_depIdxs = []int32{
	0, // 0: token.GenerateTokensResponse.tokens:type_name -> token.Tokens
	7, // 1: token.VerifyAccessResponse.header:type_name -> token.VerifyAccessResponse.HeaderEntry
	8, // 2: token.VerifyAccessResponse.claims:type_name -> token.VerifyAccessResponse.ClaimsEntry
	0, // 3: token.RefreshResponse.tokens:type_name -> token.Tokens
	1, // 4: token.TokenService.GenerateTokens:input_type -> token.GenerateTokensRequest
	2, // 5: token.TokenService.VerifyAccess:input_type -> token.VerifyAccessRequest
	3, // 6: token.TokenService.Refresh:input_type -> token.RefreshRequest
	4, // 7: token.TokenService.GenerateTokens:output_type -> token.GenerateTokensResponse
	5, // 8: token.TokenService.VerifyAccess:output_type -> token.VerifyAccessResponse
	6, // 9: token.TokenService.Refresh:output_type -> token.RefreshResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_token_proto_init() }
func file_token_proto_init() {
	if File_token_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_token_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
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
		file_token_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokensRequest); i {
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
		file_token_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAccessRequest); i {
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
		file_token_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshRequest); i {
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
		file_token_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokensResponse); i {
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
		file_token_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyAccessResponse); i {
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
		file_token_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshResponse); i {
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
			RawDescriptor: file_token_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_token_proto_goTypes,
		DependencyIndexes: file_token_proto_depIdxs,
		MessageInfos:      file_token_proto_msgTypes,
	}.Build()
	File_token_proto = out.File
	file_token_proto_rawDesc = nil
	file_token_proto_goTypes = nil
	file_token_proto_depIdxs = nil
}