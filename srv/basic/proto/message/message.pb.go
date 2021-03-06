// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/message/message.proto

package message

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ChangeStatusRequest_Status int32

const (
	ChangeStatusRequest_Unread ChangeStatusRequest_Status = 0
	ChangeStatusRequest_Readed ChangeStatusRequest_Status = 10
	ChangeStatusRequest_Trash  ChangeStatusRequest_Status = 20
)

// Enum value maps for ChangeStatusRequest_Status.
var (
	ChangeStatusRequest_Status_name = map[int32]string{
		0:  "Unread",
		10: "Readed",
		20: "Trash",
	}
	ChangeStatusRequest_Status_value = map[string]int32{
		"Unread": 0,
		"Readed": 10,
		"Trash":  20,
	}
)

func (x ChangeStatusRequest_Status) Enum() *ChangeStatusRequest_Status {
	p := new(ChangeStatusRequest_Status)
	*p = x
	return p
}

func (x ChangeStatusRequest_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChangeStatusRequest_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_message_message_proto_enumTypes[0].Descriptor()
}

func (ChangeStatusRequest_Status) Type() protoreflect.EnumType {
	return &file_proto_message_message_proto_enumTypes[0]
}

func (x ChangeStatusRequest_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChangeStatusRequest_Status.Descriptor instead.
func (ChangeStatusRequest_Status) EnumDescriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{2, 0}
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{0}
}

type SendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From    string   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	ToList  []string `protobuf:"bytes,2,rep,name=to_list,json=toList,proto3" json:"to_list,omitempty"`
	Title   string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content string   `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SendRequest) Reset() {
	*x = SendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRequest) ProtoMessage() {}

func (x *SendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRequest.ProtoReflect.Descriptor instead.
func (*SendRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{1}
}

func (x *SendRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendRequest) GetToList() []string {
	if x != nil {
		return x.ToList
	}
	return nil
}

func (x *SendRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *SendRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ChangeStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status ChangeStatusRequest_Status `protobuf:"varint,1,opt,name=status,proto3,enum=message.ChangeStatusRequest_Status" json:"status,omitempty"`
	To     string                     `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Id     int64                      `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ChangeStatusRequest) Reset() {
	*x = ChangeStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeStatusRequest) ProtoMessage() {}

func (x *ChangeStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeStatusRequest.ProtoReflect.Descriptor instead.
func (*ChangeStatusRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeStatusRequest) GetStatus() ChangeStatusRequest_Status {
	if x != nil {
		return x.Status
	}
	return ChangeStatusRequest_Unread
}

func (x *ChangeStatusRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *ChangeStatusRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To string `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
}

func (x *InitRequest) Reset() {
	*x = InitRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitRequest) ProtoMessage() {}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitRequest.ProtoReflect.Descriptor instead.
func (*InitRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{3}
}

func (x *InitRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	To     string                  `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	Unread []*InitResponse_Message `protobuf:"bytes,2,rep,name=Unread,proto3" json:"Unread,omitempty"`
	Readed []*InitResponse_Message `protobuf:"bytes,3,rep,name=Readed,proto3" json:"Readed,omitempty"`
	Trash  []*InitResponse_Message `protobuf:"bytes,4,rep,name=Trash,proto3" json:"Trash,omitempty"`
}

func (x *InitResponse) Reset() {
	*x = InitResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse) ProtoMessage() {}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse.ProtoReflect.Descriptor instead.
func (*InitResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{4}
}

func (x *InitResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *InitResponse) GetUnread() []*InitResponse_Message {
	if x != nil {
		return x.Unread
	}
	return nil
}

func (x *InitResponse) GetReaded() []*InitResponse_Message {
	if x != nil {
		return x.Readed
	}
	return nil
}

func (x *InitResponse) GetTrash() []*InitResponse_Message {
	if x != nil {
		return x.Trash
	}
	return nil
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From    string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{6}
}

func (x *GetResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *GetResponse) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetResponse) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type InitResponse_Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	From  string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	Title string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
}

func (x *InitResponse_Message) Reset() {
	*x = InitResponse_Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_message_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitResponse_Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitResponse_Message) ProtoMessage() {}

func (x *InitResponse_Message) ProtoReflect() protoreflect.Message {
	mi := &file_proto_message_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitResponse_Message.ProtoReflect.Descriptor instead.
func (*InitResponse_Message) Descriptor() ([]byte, []int) {
	return file_proto_message_message_proto_rawDescGZIP(), []int{4, 0}
}

func (x *InitResponse_Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *InitResponse_Message) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *InitResponse_Message) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_proto_message_message_proto protoreflect.FileDescriptor

var file_proto_message_message_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x6a, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x9f,
	0x01, 0x0a, 0x13, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x2b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0a, 0x0a,
	0x06, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x65, 0x61,
	0x64, 0x65, 0x64, 0x10, 0x0a, 0x12, 0x09, 0x0a, 0x05, 0x54, 0x72, 0x61, 0x73, 0x68, 0x10, 0x14,
	0x22, 0x1d, 0x0a, 0x0b, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x22,
	0x86, 0x02, 0x0a, 0x0c, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x35, 0x0a, 0x06, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x06, 0x55, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x12, 0x35, 0x0a, 0x06, 0x52, 0x65, 0x61, 0x64, 0x65,
	0x64, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x06, 0x52, 0x65, 0x61, 0x64, 0x65, 0x64, 0x12, 0x33,
	0x0a, 0x05, 0x54, 0x72, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x05, 0x54, 0x72,
	0x61, 0x73, 0x68, 0x1a, 0x43, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0xea, 0x01, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x14, 0x2e,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0c, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x49,
	0x6e, 0x69, 0x74, 0x12, 0x14, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e,
	0x69, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x32, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x13, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_message_message_proto_rawDescOnce sync.Once
	file_proto_message_message_proto_rawDescData = file_proto_message_message_proto_rawDesc
)

func file_proto_message_message_proto_rawDescGZIP() []byte {
	file_proto_message_message_proto_rawDescOnce.Do(func() {
		file_proto_message_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_message_message_proto_rawDescData)
	})
	return file_proto_message_message_proto_rawDescData
}

var file_proto_message_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_message_message_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_message_message_proto_goTypes = []interface{}{
	(ChangeStatusRequest_Status)(0), // 0: message.ChangeStatusRequest.Status
	(*Response)(nil),                // 1: message.Response
	(*SendRequest)(nil),             // 2: message.SendRequest
	(*ChangeStatusRequest)(nil),     // 3: message.ChangeStatusRequest
	(*InitRequest)(nil),             // 4: message.InitRequest
	(*InitResponse)(nil),            // 5: message.InitResponse
	(*GetRequest)(nil),              // 6: message.GetRequest
	(*GetResponse)(nil),             // 7: message.GetResponse
	(*InitResponse_Message)(nil),    // 8: message.InitResponse.Message
}
var file_proto_message_message_proto_depIdxs = []int32{
	0, // 0: message.ChangeStatusRequest.status:type_name -> message.ChangeStatusRequest.Status
	8, // 1: message.InitResponse.Unread:type_name -> message.InitResponse.Message
	8, // 2: message.InitResponse.Readed:type_name -> message.InitResponse.Message
	8, // 3: message.InitResponse.Trash:type_name -> message.InitResponse.Message
	2, // 4: message.Message.Send:input_type -> message.SendRequest
	3, // 5: message.Message.ChangeStatus:input_type -> message.ChangeStatusRequest
	4, // 6: message.Message.Init:input_type -> message.InitRequest
	6, // 7: message.Message.Get:input_type -> message.GetRequest
	1, // 8: message.Message.Send:output_type -> message.Response
	1, // 9: message.Message.ChangeStatus:output_type -> message.Response
	5, // 10: message.Message.Init:output_type -> message.InitResponse
	7, // 11: message.Message.Get:output_type -> message.GetResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_message_message_proto_init() }
func file_proto_message_message_proto_init() {
	if File_proto_message_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_message_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proto_message_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRequest); i {
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
		file_proto_message_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeStatusRequest); i {
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
		file_proto_message_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitRequest); i {
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
		file_proto_message_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse); i {
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
		file_proto_message_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_proto_message_message_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_proto_message_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitResponse_Message); i {
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
			RawDescriptor: file_proto_message_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_message_message_proto_goTypes,
		DependencyIndexes: file_proto_message_message_proto_depIdxs,
		EnumInfos:         file_proto_message_message_proto_enumTypes,
		MessageInfos:      file_proto_message_message_proto_msgTypes,
	}.Build()
	File_proto_message_message_proto = out.File
	file_proto_message_message_proto_rawDesc = nil
	file_proto_message_message_proto_goTypes = nil
	file_proto_message_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageClient is the client API for Message service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageClient interface {
	Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*Response, error)
	ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*Response, error)
	Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type messageClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageClient(cc grpc.ClientConnInterface) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) Send(ctx context.Context, in *SendRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/message.Message/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) ChangeStatus(ctx context.Context, in *ChangeStatusRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/message.Message/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Init(ctx context.Context, in *InitRequest, opts ...grpc.CallOption) (*InitResponse, error) {
	out := new(InitResponse)
	err := c.cc.Invoke(ctx, "/message.Message/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/message.Message/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessageServer is the server API for Message service.
type MessageServer interface {
	Send(context.Context, *SendRequest) (*Response, error)
	ChangeStatus(context.Context, *ChangeStatusRequest) (*Response, error)
	Init(context.Context, *InitRequest) (*InitResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

// UnimplementedMessageServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServer struct {
}

func (*UnimplementedMessageServer) Send(context.Context, *SendRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedMessageServer) ChangeStatus(context.Context, *ChangeStatusRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (*UnimplementedMessageServer) Init(context.Context, *InitRequest) (*InitResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (*UnimplementedMessageServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Message/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Send(ctx, req.(*SendRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Message/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).ChangeStatus(ctx, req.(*ChangeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Message/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Init(ctx, req.(*InitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Message_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Message/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Message_Send_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _Message_ChangeStatus_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Message_Init_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Message_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/message/message.proto",
}
