// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: chat.proto

package chat

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

type ChatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId    uint32 `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	UserId     uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	PreMsgTime uint32 `protobuf:"varint,3,opt,name=pre_msg_time,json=preMsgTime,proto3" json:"pre_msg_time,omitempty"`
}

func (x *ChatRequest) Reset() {
	*x = ChatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRequest) ProtoMessage() {}

func (x *ChatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRequest.ProtoReflect.Descriptor instead.
func (*ChatRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatRequest) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *ChatRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ChatRequest) GetPreMsgTime() uint32 {
	if x != nil {
		return x.PreMsgTime
	}
	return 0
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         uint32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Content    string  `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	CreateTime uint32  `protobuf:"varint,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	FromUserId *uint32 `protobuf:"varint,4,opt,name=from_user_id,json=fromUserId,proto3,oneof" json:"from_user_id,omitempty"`
	ToUserId   *uint32 `protobuf:"varint,5,opt,name=to_user_id,json=toUserId,proto3,oneof" json:"to_user_id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *Message) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetCreateTime() uint32 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *Message) GetFromUserId() uint32 {
	if x != nil && x.FromUserId != nil {
		return *x.FromUserId
	}
	return 0
}

func (x *Message) GetToUserId() uint32 {
	if x != nil && x.ToUserId != nil {
		return *x.ToUserId
	}
	return 0
}

type ChatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32      `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg   string     `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
	MessageList []*Message `protobuf:"bytes,3,rep,name=message_list,json=messageList,proto3" json:"message_list,omitempty"`
}

func (x *ChatResponse) Reset() {
	*x = ChatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatResponse) ProtoMessage() {}

func (x *ChatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatResponse.ProtoReflect.Descriptor instead.
func (*ChatResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ChatResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *ChatResponse) GetMessageList() []*Message {
	if x != nil {
		return x.MessageList
	}
	return nil
}

type ActionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActorId    uint32 `protobuf:"varint,1,opt,name=actor_id,json=actorId,proto3" json:"actor_id,omitempty"`
	UserId     uint32 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ActionType uint32 `protobuf:"varint,3,opt,name=action_type,json=actionType,proto3" json:"action_type,omitempty"` // 1-发送消息
	Content    string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *ActionRequest) Reset() {
	*x = ActionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionRequest) ProtoMessage() {}

func (x *ActionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionRequest.ProtoReflect.Descriptor instead.
func (*ActionRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{3}
}

func (x *ActionRequest) GetActorId() uint32 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *ActionRequest) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ActionRequest) GetActionType() uint32 {
	if x != nil {
		return x.ActionType
	}
	return 0
}

func (x *ActionRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type ActionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`
}

func (x *ActionResponse) Reset() {
	*x = ActionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActionResponse) ProtoMessage() {}

func (x *ActionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActionResponse.ProtoReflect.Descriptor instead.
func (*ActionResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{4}
}

func (x *ActionResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *ActionResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x70,
	0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x22, 0x63, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x70, 0x72, 0x65,
	0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x70, 0x72, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xbe, 0x01, 0x0a, 0x07,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0c, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x0a, 0x66, 0x72, 0x6f, 0x6d,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x21, 0x0a, 0x0a, 0x74, 0x6f, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x01, 0x52,
	0x08, 0x74, 0x6f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x0f, 0x0a, 0x0d,
	0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x0d, 0x0a,
	0x0b, 0x5f, 0x74, 0x6f, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x22, 0x84, 0x01, 0x0a,
	0x0c, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x34, 0x0a,
	0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x7e, 0x0a, 0x0d, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x6e, 0x74, 0x22, 0x50, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x32, 0x85, 0x01, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x15, 0x2e,
	0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x43, 0x68, 0x61, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x2e, 0x72, 0x70, 0x63,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x16, 0x5a,
	0x14, 0x47, 0x75, 0x47, 0x6f, 0x54, 0x69, 0x6b, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x68, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData = file_chat_proto_rawDesc
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_chat_proto_rawDescData)
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_chat_proto_goTypes = []interface{}{
	(*ChatRequest)(nil),    // 0: rpc.chat.ChatRequest
	(*Message)(nil),        // 1: rpc.chat.Message
	(*ChatResponse)(nil),   // 2: rpc.chat.ChatResponse
	(*ActionRequest)(nil),  // 3: rpc.chat.ActionRequest
	(*ActionResponse)(nil), // 4: rpc.chat.ActionResponse
}
var file_chat_proto_depIdxs = []int32{
	1, // 0: rpc.chat.ChatResponse.message_list:type_name -> rpc.chat.Message
	0, // 1: rpc.chat.ChatService.Chat:input_type -> rpc.chat.ChatRequest
	3, // 2: rpc.chat.ChatService.ChatAction:input_type -> rpc.chat.ActionRequest
	2, // 3: rpc.chat.ChatService.Chat:output_type -> rpc.chat.ChatResponse
	4, // 4: rpc.chat.ChatService.ChatAction:output_type -> rpc.chat.ActionResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRequest); i {
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
		file_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_chat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatResponse); i {
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
		file_chat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionRequest); i {
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
		file_chat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActionResponse); i {
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
	file_chat_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_rawDesc = nil
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
