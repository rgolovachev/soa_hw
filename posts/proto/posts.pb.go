// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.3
// source: proto/posts.proto

package postspb

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

type CreatePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *CreatePostReq) Reset() {
	*x = CreatePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostReq) ProtoMessage() {}

func (x *CreatePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostReq.ProtoReflect.Descriptor instead.
func (*CreatePostReq) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePostReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreatePostReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type CreatePostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *CreatePostResp) Reset() {
	*x = CreatePostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePostResp) ProtoMessage() {}

func (x *CreatePostResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePostResp.ProtoReflect.Descriptor instead.
func (*CreatePostResp) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePostResp) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type UpdatePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Text     string `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	PostId   string `protobuf:"bytes,3,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *UpdatePostReq) Reset() {
	*x = UpdatePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostReq) ProtoMessage() {}

func (x *UpdatePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostReq.ProtoReflect.Descriptor instead.
func (*UpdatePostReq) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatePostReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdatePostReq) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *UpdatePostReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type UpdatePostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePostResp) Reset() {
	*x = UpdatePostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePostResp) ProtoMessage() {}

func (x *UpdatePostResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePostResp.ProtoReflect.Descriptor instead.
func (*UpdatePostResp) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{3}
}

type DeletePostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PostId   string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *DeletePostReq) Reset() {
	*x = DeletePostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostReq) ProtoMessage() {}

func (x *DeletePostReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostReq.ProtoReflect.Descriptor instead.
func (*DeletePostReq) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{4}
}

func (x *DeletePostReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *DeletePostReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type DeletePostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePostResp) Reset() {
	*x = DeletePostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePostResp) ProtoMessage() {}

func (x *DeletePostResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePostResp.ProtoReflect.Descriptor instead.
func (*DeletePostResp) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{5}
}

type GetPostReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	PostId   string `protobuf:"bytes,2,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *GetPostReq) Reset() {
	*x = GetPostReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostReq) ProtoMessage() {}

func (x *GetPostReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostReq.ProtoReflect.Descriptor instead.
func (*GetPostReq) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{6}
}

func (x *GetPostReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetPostReq) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type GetPostResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *GetPostResp) Reset() {
	*x = GetPostResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResp) ProtoMessage() {}

func (x *GetPostResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResp.ProtoReflect.Descriptor instead.
func (*GetPostResp) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{7}
}

func (x *GetPostResp) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type GetAllPostsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	From     uint64 `protobuf:"varint,2,opt,name=from,proto3" json:"from,omitempty"`
	Count    uint64 `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllPostsReq) Reset() {
	*x = GetAllPostsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPostsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPostsReq) ProtoMessage() {}

func (x *GetAllPostsReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPostsReq.ProtoReflect.Descriptor instead.
func (*GetAllPostsReq) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllPostsReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetAllPostsReq) GetFrom() uint64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *GetAllPostsReq) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetAllPostsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostIds []string `protobuf:"bytes,1,rep,name=post_ids,json=postIds,proto3" json:"post_ids,omitempty"`
	Texts   []string `protobuf:"bytes,2,rep,name=texts,proto3" json:"texts,omitempty"`
}

func (x *GetAllPostsResp) Reset() {
	*x = GetAllPostsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_posts_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllPostsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllPostsResp) ProtoMessage() {}

func (x *GetAllPostsResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_posts_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllPostsResp.ProtoReflect.Descriptor instead.
func (*GetAllPostsResp) Descriptor() ([]byte, []int) {
	return file_proto_posts_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllPostsResp) GetPostIds() []string {
	if x != nil {
		return x.PostIds
	}
	return nil
}

func (x *GetAllPostsResp) GetTexts() []string {
	if x != nil {
		return x.Texts
	}
	return nil
}

var File_proto_posts_proto protoreflect.FileDescriptor

var file_proto_posts_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x22, 0x3f, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x22, 0x29, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x58, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64,
	0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x44, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x41, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x21, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x56, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x49, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x78, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x78, 0x74, 0x73, 0x32, 0xb9, 0x02, 0x0a,
	0x0c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a,
	0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x6f,
	0x73, 0x74, 0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x12,
	0x11, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41,
	0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x12, 0x15, 0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x16,
	0x2e, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x6f, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x6f, 0x73, 0x74,
	0x73, 0x2f, 0x3b, 0x70, 0x6f, 0x73, 0x74, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_posts_proto_rawDescOnce sync.Once
	file_proto_posts_proto_rawDescData = file_proto_posts_proto_rawDesc
)

func file_proto_posts_proto_rawDescGZIP() []byte {
	file_proto_posts_proto_rawDescOnce.Do(func() {
		file_proto_posts_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_posts_proto_rawDescData)
	})
	return file_proto_posts_proto_rawDescData
}

var file_proto_posts_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_posts_proto_goTypes = []interface{}{
	(*CreatePostReq)(nil),   // 0: posts.CreatePostReq
	(*CreatePostResp)(nil),  // 1: posts.CreatePostResp
	(*UpdatePostReq)(nil),   // 2: posts.UpdatePostReq
	(*UpdatePostResp)(nil),  // 3: posts.UpdatePostResp
	(*DeletePostReq)(nil),   // 4: posts.DeletePostReq
	(*DeletePostResp)(nil),  // 5: posts.DeletePostResp
	(*GetPostReq)(nil),      // 6: posts.GetPostReq
	(*GetPostResp)(nil),     // 7: posts.GetPostResp
	(*GetAllPostsReq)(nil),  // 8: posts.GetAllPostsReq
	(*GetAllPostsResp)(nil), // 9: posts.GetAllPostsResp
}
var file_proto_posts_proto_depIdxs = []int32{
	0, // 0: posts.PostsService.CreatePost:input_type -> posts.CreatePostReq
	2, // 1: posts.PostsService.UpdatePost:input_type -> posts.UpdatePostReq
	4, // 2: posts.PostsService.DeletePost:input_type -> posts.DeletePostReq
	6, // 3: posts.PostsService.GetPost:input_type -> posts.GetPostReq
	8, // 4: posts.PostsService.GetAllPosts:input_type -> posts.GetAllPostsReq
	1, // 5: posts.PostsService.CreatePost:output_type -> posts.CreatePostResp
	3, // 6: posts.PostsService.UpdatePost:output_type -> posts.UpdatePostResp
	5, // 7: posts.PostsService.DeletePost:output_type -> posts.DeletePostResp
	7, // 8: posts.PostsService.GetPost:output_type -> posts.GetPostResp
	9, // 9: posts.PostsService.GetAllPosts:output_type -> posts.GetAllPostsResp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_posts_proto_init() }
func file_proto_posts_proto_init() {
	if File_proto_posts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_posts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostReq); i {
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
		file_proto_posts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePostResp); i {
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
		file_proto_posts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostReq); i {
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
		file_proto_posts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePostResp); i {
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
		file_proto_posts_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostReq); i {
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
		file_proto_posts_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePostResp); i {
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
		file_proto_posts_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostReq); i {
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
		file_proto_posts_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostResp); i {
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
		file_proto_posts_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPostsReq); i {
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
		file_proto_posts_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllPostsResp); i {
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
			RawDescriptor: file_proto_posts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_posts_proto_goTypes,
		DependencyIndexes: file_proto_posts_proto_depIdxs,
		MessageInfos:      file_proto_posts_proto_msgTypes,
	}.Build()
	File_proto_posts_proto = out.File
	file_proto_posts_proto_rawDesc = nil
	file_proto_posts_proto_goTypes = nil
	file_proto_posts_proto_depIdxs = nil
}