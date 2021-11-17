// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.5.1
// source: api/group/v1/group.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Group *GroupEntity `protobuf:"bytes,1,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *CreateGroupRequest) Reset() {
	*x = CreateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_group_v1_group_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateGroupRequest) ProtoMessage() {}

func (x *CreateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_group_v1_group_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateGroupRequest.ProtoReflect.Descriptor instead.
func (*CreateGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_group_v1_group_proto_rawDescGZIP(), []int{0}
}

func (x *CreateGroupRequest) GetGroup() *GroupEntity {
	if x != nil {
		return x.Group
	}
	return nil
}

type UpdateGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Group *GroupEntity `protobuf:"bytes,2,opt,name=group,proto3" json:"group,omitempty"`
}

func (x *UpdateGroupRequest) Reset() {
	*x = UpdateGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_group_v1_group_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateGroupRequest) ProtoMessage() {}

func (x *UpdateGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_group_v1_group_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateGroupRequest.ProtoReflect.Descriptor instead.
func (*UpdateGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_group_v1_group_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateGroupRequest) GetGroup() *GroupEntity {
	if x != nil {
		return x.Group
	}
	return nil
}

type DeleteGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteGroupRequest) Reset() {
	*x = DeleteGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_group_v1_group_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteGroupRequest) ProtoMessage() {}

func (x *DeleteGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_group_v1_group_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteGroupRequest.ProtoReflect.Descriptor instead.
func (*DeleteGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_group_v1_group_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetGroupRequest) Reset() {
	*x = GetGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_group_v1_group_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetGroupRequest) ProtoMessage() {}

func (x *GetGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_group_v1_group_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetGroupRequest.ProtoReflect.Descriptor instead.
func (*GetGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_group_v1_group_proto_rawDescGZIP(), []int{3}
}

func (x *GetGroupRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ListGroupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListGroupRequest) Reset() {
	*x = ListGroupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_group_v1_group_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListGroupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListGroupRequest) ProtoMessage() {}

func (x *ListGroupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_group_v1_group_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListGroupRequest.ProtoReflect.Descriptor instead.
func (*ListGroupRequest) Descriptor() ([]byte, []int) {
	return file_api_group_v1_group_proto_rawDescGZIP(), []int{4}
}

//------------------------------------------------------------------------------
// Basic data types
//------------------------------------------------------------------------------
type GroupEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc       string            `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
	Parent     string            `protobuf:"bytes,3,opt,name=parent,proto3" json:"parent,omitempty"`
	Ext        map[string]string `protobuf:"bytes,4,rep,name=ext,proto3" json:"ext,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XId        string            `protobuf:"bytes,5,opt,name=_id,json=Id,proto3" json:"_id,omitempty"`
	XCreatedAt uint64            `protobuf:"varint,6,opt,name=_createdAt,json=CreatedAt,proto3" json:"_createdAt,omitempty"`
	XUpdatedAt uint64            `protobuf:"varint,7,opt,name=_updatedAt,json=UpdatedAt,proto3" json:"_updatedAt,omitempty"`
}

func (x *GroupEntity) Reset() {
	*x = GroupEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_group_v1_group_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GroupEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GroupEntity) ProtoMessage() {}

func (x *GroupEntity) ProtoReflect() protoreflect.Message {
	mi := &file_api_group_v1_group_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GroupEntity.ProtoReflect.Descriptor instead.
func (*GroupEntity) Descriptor() ([]byte, []int) {
	return file_api_group_v1_group_proto_rawDescGZIP(), []int{5}
}

func (x *GroupEntity) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GroupEntity) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *GroupEntity) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *GroupEntity) GetExt() map[string]string {
	if x != nil {
		return x.Ext
	}
	return nil
}

func (x *GroupEntity) GetXId() string {
	if x != nil {
		return x.XId
	}
	return ""
}

func (x *GroupEntity) GetXCreatedAt() uint64 {
	if x != nil {
		return x.XCreatedAt
	}
	return 0
}

func (x *GroupEntity) GetXUpdatedAt() uint64 {
	if x != nil {
		return x.XUpdatedAt
	}
	return 0
}

type CommonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CommonResponse) Reset() {
	*x = CommonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_group_v1_group_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommonResponse) ProtoMessage() {}

func (x *CommonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_group_v1_group_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommonResponse.ProtoReflect.Descriptor instead.
func (*CommonResponse) Descriptor() ([]byte, []int) {
	return file_api_group_v1_group_proto_rawDescGZIP(), []int{6}
}

func (x *CommonResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_api_group_v1_group_proto protoreflect.FileDescriptor

var file_api_group_v1_group_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x22, 0x55, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x2f, 0x0a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x05, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65,
	0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x8a, 0x02, 0x0a, 0x0b, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x34, 0x0a, 0x03, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x45, 0x78, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x03, 0x65, 0x78, 0x74, 0x12, 0x0f, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x5f, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x5f, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x36, 0x0a, 0x08, 0x45, 0x78, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x28,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xf5, 0x03, 0x0a, 0x05, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x64, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75,
	0x70, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x22, 0x06, 0x2f, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x3a, 0x05, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x69, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x1a,
	0x0b, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x3a, 0x05, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x62, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x2a, 0x0b, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x5c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x47, 0x72,
	0x6f, 0x75, 0x70, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x59, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f,
	0x75, 0x70, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x42, 0x28, 0x0a, 0x0c, 0x61, 0x70, 0x69, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x76, 0x31,
	0x50, 0x01, 0x5a, 0x16, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x72, 0x6f, 0x75, 0x70, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_group_v1_group_proto_rawDescOnce sync.Once
	file_api_group_v1_group_proto_rawDescData = file_api_group_v1_group_proto_rawDesc
)

func file_api_group_v1_group_proto_rawDescGZIP() []byte {
	file_api_group_v1_group_proto_rawDescOnce.Do(func() {
		file_api_group_v1_group_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_group_v1_group_proto_rawDescData)
	})
	return file_api_group_v1_group_proto_rawDescData
}

var file_api_group_v1_group_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_group_v1_group_proto_goTypes = []interface{}{
	(*CreateGroupRequest)(nil), // 0: api.group.v1.CreateGroupRequest
	(*UpdateGroupRequest)(nil), // 1: api.group.v1.UpdateGroupRequest
	(*DeleteGroupRequest)(nil), // 2: api.group.v1.DeleteGroupRequest
	(*GetGroupRequest)(nil),    // 3: api.group.v1.GetGroupRequest
	(*ListGroupRequest)(nil),   // 4: api.group.v1.ListGroupRequest
	(*GroupEntity)(nil),        // 5: api.group.v1.GroupEntity
	(*CommonResponse)(nil),     // 6: api.group.v1.CommonResponse
	nil,                        // 7: api.group.v1.GroupEntity.ExtEntry
}
var file_api_group_v1_group_proto_depIdxs = []int32{
	5, // 0: api.group.v1.CreateGroupRequest.group:type_name -> api.group.v1.GroupEntity
	5, // 1: api.group.v1.UpdateGroupRequest.group:type_name -> api.group.v1.GroupEntity
	7, // 2: api.group.v1.GroupEntity.ext:type_name -> api.group.v1.GroupEntity.ExtEntry
	0, // 3: api.group.v1.Group.CreateGroup:input_type -> api.group.v1.CreateGroupRequest
	1, // 4: api.group.v1.Group.UpdateGroup:input_type -> api.group.v1.UpdateGroupRequest
	2, // 5: api.group.v1.Group.DeleteGroup:input_type -> api.group.v1.DeleteGroupRequest
	3, // 6: api.group.v1.Group.GetGroup:input_type -> api.group.v1.GetGroupRequest
	4, // 7: api.group.v1.Group.ListGroup:input_type -> api.group.v1.ListGroupRequest
	6, // 8: api.group.v1.Group.CreateGroup:output_type -> api.group.v1.CommonResponse
	6, // 9: api.group.v1.Group.UpdateGroup:output_type -> api.group.v1.CommonResponse
	6, // 10: api.group.v1.Group.DeleteGroup:output_type -> api.group.v1.CommonResponse
	6, // 11: api.group.v1.Group.GetGroup:output_type -> api.group.v1.CommonResponse
	6, // 12: api.group.v1.Group.ListGroup:output_type -> api.group.v1.CommonResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_group_v1_group_proto_init() }
func file_api_group_v1_group_proto_init() {
	if File_api_group_v1_group_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_group_v1_group_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateGroupRequest); i {
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
		file_api_group_v1_group_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateGroupRequest); i {
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
		file_api_group_v1_group_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteGroupRequest); i {
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
		file_api_group_v1_group_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetGroupRequest); i {
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
		file_api_group_v1_group_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListGroupRequest); i {
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
		file_api_group_v1_group_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GroupEntity); i {
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
		file_api_group_v1_group_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommonResponse); i {
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
			RawDescriptor: file_api_group_v1_group_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_group_v1_group_proto_goTypes,
		DependencyIndexes: file_api_group_v1_group_proto_depIdxs,
		MessageInfos:      file_api_group_v1_group_proto_msgTypes,
	}.Build()
	File_api_group_v1_group_proto = out.File
	file_api_group_v1_group_proto_rawDesc = nil
	file_api_group_v1_group_proto_goTypes = nil
	file_api_group_v1_group_proto_depIdxs = nil
}
