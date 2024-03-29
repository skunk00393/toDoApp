// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: listItem.proto

package toDoItem

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

type NewItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ListId int32     `protobuf:"varint,1,opt,name=listId,proto3" json:"listId,omitempty"`
	Item   *ListItem `protobuf:"bytes,2,opt,name=item,proto3" json:"item,omitempty"`
}

func (x *NewItemRequest) Reset() {
	*x = NewItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewItemRequest) ProtoMessage() {}

func (x *NewItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewItemRequest.ProtoReflect.Descriptor instead.
func (*NewItemRequest) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{0}
}

func (x *NewItemRequest) GetListId() int32 {
	if x != nil {
		return x.ListId
	}
	return 0
}

func (x *NewItemRequest) GetItem() *ListItem {
	if x != nil {
		return x.Item
	}
	return nil
}

type NewItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *NewItemResponse) Reset() {
	*x = NewItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewItemResponse) ProtoMessage() {}

func (x *NewItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewItemResponse.ProtoReflect.Descriptor instead.
func (*NewItemResponse) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{1}
}

func (x *NewItemResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type ListItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *ListItem) Reset() {
	*x = ListItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListItem) ProtoMessage() {}

func (x *ListItem) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListItem.ProtoReflect.Descriptor instead.
func (*ListItem) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{2}
}

func (x *ListItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListItem) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	NewName string `protobuf:"bytes,2,opt,name=newName,proto3" json:"newName,omitempty"`
}

func (x *UpdateItemRequest) Reset() {
	*x = UpdateItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemRequest) ProtoMessage() {}

func (x *UpdateItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemRequest.ProtoReflect.Descriptor instead.
func (*UpdateItemRequest) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateItemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateItemRequest) GetNewName() string {
	if x != nil {
		return x.NewName
	}
	return ""
}

type UpdateItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *UpdateItemResponse) Reset() {
	*x = UpdateItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateItemResponse) ProtoMessage() {}

func (x *UpdateItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateItemResponse.ProtoReflect.Descriptor instead.
func (*UpdateItemResponse) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateItemResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type DeleteItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteItemRequest) Reset() {
	*x = DeleteItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemRequest) ProtoMessage() {}

func (x *DeleteItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemRequest.ProtoReflect.Descriptor instead.
func (*DeleteItemRequest) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteItemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *DeleteItemResponse) Reset() {
	*x = DeleteItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteItemResponse) ProtoMessage() {}

func (x *DeleteItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteItemResponse.ProtoReflect.Descriptor instead.
func (*DeleteItemResponse) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteItemResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type MarkItemRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MarkItemRequest) Reset() {
	*x = MarkItemRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkItemRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkItemRequest) ProtoMessage() {}

func (x *MarkItemRequest) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkItemRequest.ProtoReflect.Descriptor instead.
func (*MarkItemRequest) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{7}
}

func (x *MarkItemRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type MarkItemResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,2,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *MarkItemResponse) Reset() {
	*x = MarkItemResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_listItem_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MarkItemResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MarkItemResponse) ProtoMessage() {}

func (x *MarkItemResponse) ProtoReflect() protoreflect.Message {
	mi := &file_listItem_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MarkItemResponse.ProtoReflect.Descriptor instead.
func (*MarkItemResponse) Descriptor() ([]byte, []int) {
	return file_listItem_proto_rawDescGZIP(), []int{8}
}

func (x *MarkItemResponse) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

var File_listItem_proto protoreflect.FileDescriptor

var file_listItem_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x47, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6c, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x04, 0x69, 0x74,
	0x65, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x22, 0x2d, 0x0a, 0x0f, 0x4e, 0x65, 0x77,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x40, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x30, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64,
	0x22, 0x30, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x4d, 0x61, 0x72, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2e, 0x0a, 0x10, 0x4d, 0x61, 0x72, 0x6b, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xe8, 0x01, 0x0a, 0x0b, 0x69, 0x74, 0x65, 0x6d, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x65, 0x77, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0f, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65, 0x6d,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x4e, 0x65, 0x77, 0x49, 0x74, 0x65,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x12, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x31, 0x0a,
	0x08, 0x4d, 0x61, 0x72, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x2e, 0x4d, 0x61, 0x72, 0x6b,
	0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x4d, 0x61,
	0x72, 0x6b, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x74, 0x6f, 0x44, 0x6f, 0x49, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_listItem_proto_rawDescOnce sync.Once
	file_listItem_proto_rawDescData = file_listItem_proto_rawDesc
)

func file_listItem_proto_rawDescGZIP() []byte {
	file_listItem_proto_rawDescOnce.Do(func() {
		file_listItem_proto_rawDescData = protoimpl.X.CompressGZIP(file_listItem_proto_rawDescData)
	})
	return file_listItem_proto_rawDescData
}

var file_listItem_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_listItem_proto_goTypes = []interface{}{
	(*NewItemRequest)(nil),     // 0: NewItemRequest
	(*NewItemResponse)(nil),    // 1: NewItemResponse
	(*ListItem)(nil),           // 2: ListItem
	(*UpdateItemRequest)(nil),  // 3: UpdateItemRequest
	(*UpdateItemResponse)(nil), // 4: UpdateItemResponse
	(*DeleteItemRequest)(nil),  // 5: DeleteItemRequest
	(*DeleteItemResponse)(nil), // 6: DeleteItemResponse
	(*MarkItemRequest)(nil),    // 7: MarkItemRequest
	(*MarkItemResponse)(nil),   // 8: MarkItemResponse
}
var file_listItem_proto_depIdxs = []int32{
	2, // 0: NewItemRequest.item:type_name -> ListItem
	0, // 1: itemService.CreateNewItem:input_type -> NewItemRequest
	3, // 2: itemService.UpdateItem:input_type -> UpdateItemRequest
	5, // 3: itemService.DeleteItem:input_type -> DeleteItemRequest
	7, // 4: itemService.MarkItem:input_type -> MarkItemRequest
	1, // 5: itemService.CreateNewItem:output_type -> NewItemResponse
	4, // 6: itemService.UpdateItem:output_type -> UpdateItemResponse
	6, // 7: itemService.DeleteItem:output_type -> DeleteItemResponse
	8, // 8: itemService.MarkItem:output_type -> MarkItemResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_listItem_proto_init() }
func file_listItem_proto_init() {
	if File_listItem_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_listItem_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewItemRequest); i {
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
		file_listItem_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewItemResponse); i {
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
		file_listItem_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListItem); i {
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
		file_listItem_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemRequest); i {
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
		file_listItem_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateItemResponse); i {
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
		file_listItem_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemRequest); i {
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
		file_listItem_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteItemResponse); i {
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
		file_listItem_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkItemRequest); i {
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
		file_listItem_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MarkItemResponse); i {
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
			RawDescriptor: file_listItem_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_listItem_proto_goTypes,
		DependencyIndexes: file_listItem_proto_depIdxs,
		MessageInfos:      file_listItem_proto_msgTypes,
	}.Build()
	File_listItem_proto = out.File
	file_listItem_proto_rawDesc = nil
	file_listItem_proto_goTypes = nil
	file_listItem_proto_depIdxs = nil
}
