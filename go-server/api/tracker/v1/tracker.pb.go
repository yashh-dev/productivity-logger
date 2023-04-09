// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.19.1
// source: tracker/v1/tracker.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CreateBlockRequestBody *CreateBlockRequestBody `protobuf:"bytes,1,opt,name=create_block_request_body,json=createBlockRequestBody,proto3" json:"create_block_request_body,omitempty"`
	Status                 string                  `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Code                   int32                   `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Execution              string                  `protobuf:"bytes,3,opt,name=Execution,proto3" json:"Execution,omitempty"`
}

func (x *CreateBlock) Reset() {
	*x = CreateBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_v1_tracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlock) ProtoMessage() {}

func (x *CreateBlock) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_v1_tracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlock.ProtoReflect.Descriptor instead.
func (*CreateBlock) Descriptor() ([]byte, []int) {
	return file_tracker_v1_tracker_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBlock) GetCreateBlockRequestBody() *CreateBlockRequestBody {
	if x != nil {
		return x.CreateBlockRequestBody
	}
	return nil
}

func (x *CreateBlock) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CreateBlock) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreateBlock) GetExecution() string {
	if x != nil {
		return x.Execution
	}
	return ""
}

type BlockResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks    *BlockItems `protobuf:"bytes,1,opt,name=blocks,proto3" json:"blocks,omitempty"`
	Status    string      `protobuf:"bytes,4,opt,name=Status,proto3" json:"Status,omitempty"`
	Code      int32       `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Execution string      `protobuf:"bytes,3,opt,name=Execution,proto3" json:"Execution,omitempty"`
}

func (x *BlockResp) Reset() {
	*x = BlockResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_v1_tracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResp) ProtoMessage() {}

func (x *BlockResp) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_v1_tracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResp.ProtoReflect.Descriptor instead.
func (*BlockResp) Descriptor() ([]byte, []int) {
	return file_tracker_v1_tracker_proto_rawDescGZIP(), []int{1}
}

func (x *BlockResp) GetBlocks() *BlockItems {
	if x != nil {
		return x.Blocks
	}
	return nil
}

func (x *BlockResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *BlockResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *BlockResp) GetExecution() string {
	if x != nil {
		return x.Execution
	}
	return ""
}

type BlockItems struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BlockId   string `protobuf:"bytes,1,opt,name=BlockId,proto3" json:"BlockId,omitempty"`
	StartTime string `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	StopTime  string `protobuf:"bytes,3,opt,name=StopTime,proto3" json:"StopTime,omitempty"`
	Duration  string `protobuf:"bytes,4,opt,name=Duration,proto3" json:"Duration,omitempty"`
	Label     string `protobuf:"bytes,5,opt,name=Label,proto3" json:"Label,omitempty"`
}

func (x *BlockItems) Reset() {
	*x = BlockItems{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_v1_tracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockItems) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockItems) ProtoMessage() {}

func (x *BlockItems) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_v1_tracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockItems.ProtoReflect.Descriptor instead.
func (*BlockItems) Descriptor() ([]byte, []int) {
	return file_tracker_v1_tracker_proto_rawDescGZIP(), []int{2}
}

func (x *BlockItems) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

func (x *BlockItems) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *BlockItems) GetStopTime() string {
	if x != nil {
		return x.StopTime
	}
	return ""
}

func (x *BlockItems) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *BlockItems) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

type DeletBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	BlockId string `protobuf:"bytes,2,opt,name=BlockId,proto3" json:"BlockId,omitempty"`
}

func (x *DeletBlock) Reset() {
	*x = DeletBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_v1_tracker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletBlock) ProtoMessage() {}

func (x *DeletBlock) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_v1_tracker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletBlock.ProtoReflect.Descriptor instead.
func (*DeletBlock) Descriptor() ([]byte, []int) {
	return file_tracker_v1_tracker_proto_rawDescGZIP(), []int{3}
}

func (x *DeletBlock) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *DeletBlock) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

type UpdateBlock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	BlockId string `protobuf:"bytes,2,opt,name=BlockId,proto3" json:"BlockId,omitempty"`
}

func (x *UpdateBlock) Reset() {
	*x = UpdateBlock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_v1_tracker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBlock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBlock) ProtoMessage() {}

func (x *UpdateBlock) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_v1_tracker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBlock.ProtoReflect.Descriptor instead.
func (*UpdateBlock) Descriptor() ([]byte, []int) {
	return file_tracker_v1_tracker_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateBlock) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateBlock) GetBlockId() string {
	if x != nil {
		return x.BlockId
	}
	return ""
}

type GetListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
}

func (x *GetListReq) Reset() {
	*x = GetListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_v1_tracker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetListReq) ProtoMessage() {}

func (x *GetListReq) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_v1_tracker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetListReq.ProtoReflect.Descriptor instead.
func (*GetListReq) Descriptor() ([]byte, []int) {
	return file_tracker_v1_tracker_proto_rawDescGZIP(), []int{5}
}

func (x *GetListReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CreateBlockRequestBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string                           `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	StartTime string                           `protobuf:"bytes,2,opt,name=StartTime,proto3" json:"StartTime,omitempty"`
	Lapse     map[int32]*timestamppb.Timestamp `protobuf:"bytes,5,rep,name=Lapse,proto3" json:"Lapse,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Duration  string                           `protobuf:"bytes,7,opt,name=Duration,proto3" json:"Duration,omitempty"`
	StopTime  string                           `protobuf:"bytes,3,opt,name=StopTime,proto3" json:"StopTime,omitempty"`
	Label     string                           `protobuf:"bytes,4,opt,name=Label,proto3" json:"Label,omitempty"`
}

func (x *CreateBlockRequestBody) Reset() {
	*x = CreateBlockRequestBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_v1_tracker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBlockRequestBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBlockRequestBody) ProtoMessage() {}

func (x *CreateBlockRequestBody) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_v1_tracker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBlockRequestBody.ProtoReflect.Descriptor instead.
func (*CreateBlockRequestBody) Descriptor() ([]byte, []int) {
	return file_tracker_v1_tracker_proto_rawDescGZIP(), []int{6}
}

func (x *CreateBlockRequestBody) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *CreateBlockRequestBody) GetStartTime() string {
	if x != nil {
		return x.StartTime
	}
	return ""
}

func (x *CreateBlockRequestBody) GetLapse() map[int32]*timestamppb.Timestamp {
	if x != nil {
		return x.Lapse
	}
	return nil
}

func (x *CreateBlockRequestBody) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *CreateBlockRequestBody) GetStopTime() string {
	if x != nil {
		return x.StopTime
	}
	return ""
}

func (x *CreateBlockRequestBody) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

var File_tracker_v1_tracker_proto protoreflect.FileDescriptor

var file_tracker_v1_tracker_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65,
	0x72, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xc1, 0x01, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x68, 0x0a, 0x19, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f,
	0x64, 0x79, 0x52, 0x16, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0x90, 0x01, 0x0a, 0x09, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x39, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x49, 0x74, 0x65, 0x6d, 0x73, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x3e, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x3f, 0x0a, 0x0b,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x22, 0x24, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x22, 0xc2, 0x02, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f, 0x64, 0x79, 0x12, 0x16,
	0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54,
	0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x12, 0x4e, 0x0a, 0x05, 0x4c, 0x61, 0x70, 0x73, 0x65, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x6f,
	0x64, 0x79, 0x2e, 0x4c, 0x61, 0x70, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x4c,
	0x61, 0x70, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x53, 0x74, 0x6f, 0x70, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x1a, 0x54, 0x0a, 0x0a, 0x4c, 0x61, 0x70, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x30, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xf9, 0x03, 0x0a, 0x07, 0x54, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x12, 0x88, 0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x75, 0x6e, 0x63, 0x12, 0x22, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f,
	0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x20, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x29, 0x3a, 0x19, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x62, 0x6f, 0x64,
	0x79, 0x22, 0x0c, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x7e, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x75, 0x6e,
	0x63, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e, 0x76,
	0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c,
	0x64, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x27, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21, 0x2a, 0x1f,
	0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2f, 0x7b, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x7d, 0x2f, 0x7b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x64, 0x7d, 0x12,
	0x70, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x75,
	0x6e, 0x63, 0x12, 0x22, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f,
	0x72, 0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x3a, 0x01, 0x2a, 0x1a, 0x0c, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x71, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x46, 0x75,
	0x6e, 0x63, 0x12, 0x21, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x2e,
	0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x20, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2e, 0x76, 0x31, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12,
	0x13, 0x2f, 0x67, 0x65, 0x74, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x7b, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x7d, 0x42, 0x22, 0x50, 0x01, 0x5a, 0x1e, 0x67, 0x6f, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72,
	0x6c, 0x64, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracker_v1_tracker_proto_rawDescOnce sync.Once
	file_tracker_v1_tracker_proto_rawDescData = file_tracker_v1_tracker_proto_rawDesc
)

func file_tracker_v1_tracker_proto_rawDescGZIP() []byte {
	file_tracker_v1_tracker_proto_rawDescOnce.Do(func() {
		file_tracker_v1_tracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracker_v1_tracker_proto_rawDescData)
	})
	return file_tracker_v1_tracker_proto_rawDescData
}

var file_tracker_v1_tracker_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tracker_v1_tracker_proto_goTypes = []interface{}{
	(*CreateBlock)(nil),            // 0: helloworld.v1.tracker.CreateBlock
	(*BlockResp)(nil),              // 1: helloworld.v1.tracker.BlockResp
	(*BlockItems)(nil),             // 2: helloworld.v1.tracker.BlockItems
	(*DeletBlock)(nil),             // 3: helloworld.v1.tracker.DeletBlock
	(*UpdateBlock)(nil),            // 4: helloworld.v1.tracker.UpdateBlock
	(*GetListReq)(nil),             // 5: helloworld.v1.tracker.GetListReq
	(*CreateBlockRequestBody)(nil), // 6: helloworld.v1.tracker.CreateBlockRequestBody
	nil,                            // 7: helloworld.v1.tracker.CreateBlockRequestBody.LapseEntry
	(*timestamppb.Timestamp)(nil),  // 8: google.protobuf.Timestamp
}
var file_tracker_v1_tracker_proto_depIdxs = []int32{
	6, // 0: helloworld.v1.tracker.CreateBlock.create_block_request_body:type_name -> helloworld.v1.tracker.CreateBlockRequestBody
	2, // 1: helloworld.v1.tracker.BlockResp.blocks:type_name -> helloworld.v1.tracker.BlockItems
	7, // 2: helloworld.v1.tracker.CreateBlockRequestBody.Lapse:type_name -> helloworld.v1.tracker.CreateBlockRequestBody.LapseEntry
	8, // 3: helloworld.v1.tracker.CreateBlockRequestBody.LapseEntry.value:type_name -> google.protobuf.Timestamp
	0, // 4: helloworld.v1.tracker.Tracker.CreateBlockFunc:input_type -> helloworld.v1.tracker.CreateBlock
	3, // 5: helloworld.v1.tracker.Tracker.DeletBlockFunc:input_type -> helloworld.v1.tracker.DeletBlock
	4, // 6: helloworld.v1.tracker.Tracker.UpdateBlockFunc:input_type -> helloworld.v1.tracker.UpdateBlock
	5, // 7: helloworld.v1.tracker.Tracker.ListBlockFunc:input_type -> helloworld.v1.tracker.GetListReq
	1, // 8: helloworld.v1.tracker.Tracker.CreateBlockFunc:output_type -> helloworld.v1.tracker.BlockResp
	1, // 9: helloworld.v1.tracker.Tracker.DeletBlockFunc:output_type -> helloworld.v1.tracker.BlockResp
	1, // 10: helloworld.v1.tracker.Tracker.UpdateBlockFunc:output_type -> helloworld.v1.tracker.BlockResp
	1, // 11: helloworld.v1.tracker.Tracker.ListBlockFunc:output_type -> helloworld.v1.tracker.BlockResp
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_tracker_v1_tracker_proto_init() }
func file_tracker_v1_tracker_proto_init() {
	if File_tracker_v1_tracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracker_v1_tracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlock); i {
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
		file_tracker_v1_tracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResp); i {
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
		file_tracker_v1_tracker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockItems); i {
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
		file_tracker_v1_tracker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletBlock); i {
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
		file_tracker_v1_tracker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBlock); i {
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
		file_tracker_v1_tracker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetListReq); i {
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
		file_tracker_v1_tracker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBlockRequestBody); i {
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
			RawDescriptor: file_tracker_v1_tracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tracker_v1_tracker_proto_goTypes,
		DependencyIndexes: file_tracker_v1_tracker_proto_depIdxs,
		MessageInfos:      file_tracker_v1_tracker_proto_msgTypes,
	}.Build()
	File_tracker_v1_tracker_proto = out.File
	file_tracker_v1_tracker_proto_rawDesc = nil
	file_tracker_v1_tracker_proto_goTypes = nil
	file_tracker_v1_tracker_proto_depIdxs = nil
}
