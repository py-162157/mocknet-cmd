// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: rpctest.proto

package rpctest

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

type MessageType int32

const (
	// message of command type
	MessageType_emunet_creation MessageType = 0
	MessageType_emunet_deletion MessageType = 1
	MessageType_emunet_update   MessageType = 2
	MessageType_emunet_test     MessageType = 3
	// message of request type
	MessageType_get_emunet_info MessageType = 4
	// message of response type
	MessageType_command_response MessageType = 5
	MessageType_request_response MessageType = 6
)

// Enum value maps for MessageType.
var (
	MessageType_name = map[int32]string{
		0: "emunet_creation",
		1: "emunet_deletion",
		2: "emunet_update",
		3: "emunet_test",
		4: "get_emunet_info",
		5: "command_response",
		6: "request_response",
	}
	MessageType_value = map[string]int32{
		"emunet_creation":  0,
		"emunet_deletion":  1,
		"emunet_update":    2,
		"emunet_test":      3,
		"get_emunet_info":  4,
		"command_response": 5,
		"request_response": 6,
	}
)

func (x MessageType) Enum() *MessageType {
	p := new(MessageType)
	*p = x
	return p
}

func (x MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_rpctest_proto_enumTypes[0].Descriptor()
}

func (MessageType) Type() protoreflect.EnumType {
	return &file_rpctest_proto_enumTypes[0]
}

func (x MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MessageType.Descriptor instead.
func (MessageType) EnumDescriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{0}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=rpctest.MessageType" json:"type,omitempty"`
	Command  *Command    `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Request  *Request    `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	Response *Response   `protobuf:"bytes,4,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[0]
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
	return file_rpctest_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() MessageType {
	if x != nil {
		return x.Type
	}
	return MessageType_emunet_creation
}

func (x *Message) GetCommand() *Command {
	if x != nil {
		return x.Command
	}
	return nil
}

func (x *Message) GetRequest() *Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Message) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

type Command struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmunetCreation *EmunetCreation `protobuf:"bytes,1,opt,name=emunet_creation,json=emunetCreation,proto3" json:"emunet_creation,omitempty"`
	EmunetDeletion *EmunetDeletion `protobuf:"bytes,2,opt,name=emunet_deletion,json=emunetDeletion,proto3" json:"emunet_deletion,omitempty"`
	EmunetUpdate   *EmunetUpdate   `protobuf:"bytes,3,opt,name=emunet_update,json=emunetUpdate,proto3" json:"emunet_update,omitempty"`
	EmunetTest     *EmunetTest     `protobuf:"bytes,4,opt,name=emunet_test,json=emunetTest,proto3" json:"emunet_test,omitempty"`
}

func (x *Command) Reset() {
	*x = Command{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Command) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Command) ProtoMessage() {}

func (x *Command) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Command.ProtoReflect.Descriptor instead.
func (*Command) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{1}
}

func (x *Command) GetEmunetCreation() *EmunetCreation {
	if x != nil {
		return x.EmunetCreation
	}
	return nil
}

func (x *Command) GetEmunetDeletion() *EmunetDeletion {
	if x != nil {
		return x.EmunetDeletion
	}
	return nil
}

func (x *Command) GetEmunetUpdate() *EmunetUpdate {
	if x != nil {
		return x.EmunetUpdate
	}
	return nil
}

func (x *Command) GetEmunetTest() *EmunetTest {
	if x != nil {
		return x.EmunetTest
	}
	return nil
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GetEmunetInfo *GetEmunetInfo `protobuf:"bytes,1,opt,name=get_emunet_info,json=getEmunetInfo,proto3" json:"get_emunet_info,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{2}
}

func (x *Request) GetGetEmunetInfo() *GetEmunetInfo {
	if x != nil {
		return x.GetEmunetInfo
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result           bool             `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
	ErrorDescription string           `protobuf:"bytes,2,opt,name=error_description,json=errorDescription,proto3" json:"error_description,omitempty"`
	CommandResponse  *CommandResponse `protobuf:"bytes,3,opt,name=command_response,json=commandResponse,proto3" json:"command_response,omitempty"`
	RequestResponse  *RequestResponse `protobuf:"bytes,4,opt,name=request_response,json=requestResponse,proto3" json:"request_response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[3]
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
	return file_rpctest_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *Response) GetErrorDescription() string {
	if x != nil {
		return x.ErrorDescription
	}
	return ""
}

func (x *Response) GetCommandResponse() *CommandResponse {
	if x != nil {
		return x.CommandResponse
	}
	return nil
}

func (x *Response) GetRequestResponse() *RequestResponse {
	if x != nil {
		return x.RequestResponse
	}
	return nil
}

// message of command type
type EmunetCreation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emunet *Emunet `protobuf:"bytes,1,opt,name=emunet,proto3" json:"emunet,omitempty"`
}

func (x *EmunetCreation) Reset() {
	*x = EmunetCreation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmunetCreation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmunetCreation) ProtoMessage() {}

func (x *EmunetCreation) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmunetCreation.ProtoReflect.Descriptor instead.
func (*EmunetCreation) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{4}
}

func (x *EmunetCreation) GetEmunet() *Emunet {
	if x != nil {
		return x.Emunet
	}
	return nil
}

type EmunetDeletion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmunetDeletion) Reset() {
	*x = EmunetDeletion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmunetDeletion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmunetDeletion) ProtoMessage() {}

func (x *EmunetDeletion) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmunetDeletion.ProtoReflect.Descriptor instead.
func (*EmunetDeletion) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{5}
}

type EmunetUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emunet *Emunet `protobuf:"bytes,1,opt,name=emunet,proto3" json:"emunet,omitempty"`
}

func (x *EmunetUpdate) Reset() {
	*x = EmunetUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmunetUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmunetUpdate) ProtoMessage() {}

func (x *EmunetUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmunetUpdate.ProtoReflect.Descriptor instead.
func (*EmunetUpdate) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{6}
}

func (x *EmunetUpdate) GetEmunet() *Emunet {
	if x != nil {
		return x.Emunet
	}
	return nil
}

type EmunetTest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmunetTest) Reset() {
	*x = EmunetTest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmunetTest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmunetTest) ProtoMessage() {}

func (x *EmunetTest) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmunetTest.ProtoReflect.Descriptor instead.
func (*EmunetTest) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{7}
}

// message of request type
type GetEmunetInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEmunetInfo) Reset() {
	*x = GetEmunetInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEmunetInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEmunetInfo) ProtoMessage() {}

func (x *GetEmunetInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEmunetInfo.ProtoReflect.Descriptor instead.
func (*GetEmunetInfo) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{8}
}

// message of response type
type CommandResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CommandResponse) Reset() {
	*x = CommandResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommandResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommandResponse) ProtoMessage() {}

func (x *CommandResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommandResponse.ProtoReflect.Descriptor instead.
func (*CommandResponse) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{9}
}

type RequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Emunet *Emunet `protobuf:"bytes,1,opt,name=emunet,proto3" json:"emunet,omitempty"`
}

func (x *RequestResponse) Reset() {
	*x = RequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestResponse) ProtoMessage() {}

func (x *RequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestResponse.ProtoReflect.Descriptor instead.
func (*RequestResponse) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{10}
}

func (x *RequestResponse) GetEmunet() *Emunet {
	if x != nil {
		return x.Emunet
	}
	return nil
}

type Emunet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pods  []*Pod  `protobuf:"bytes,1,rep,name=pods,proto3" json:"pods,omitempty"`
	Links []*Link `protobuf:"bytes,2,rep,name=links,proto3" json:"links,omitempty"`
	Type  string  `protobuf:"bytes,3,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *Emunet) Reset() {
	*x = Emunet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Emunet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Emunet) ProtoMessage() {}

func (x *Emunet) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Emunet.ProtoReflect.Descriptor instead.
func (*Emunet) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{11}
}

func (x *Emunet) GetPods() []*Pod {
	if x != nil {
		return x.Pods
	}
	return nil
}

func (x *Emunet) GetLinks() []*Link {
	if x != nil {
		return x.Links
	}
	return nil
}

func (x *Emunet) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type Link struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Node1    *Pod   `protobuf:"bytes,2,opt,name=node1,proto3" json:"node1,omitempty"`
	Node2    *Pod   `protobuf:"bytes,3,opt,name=node2,proto3" json:"node2,omitempty"`
	Node1Inf string `protobuf:"bytes,4,opt,name=node1_inf,json=node1Inf,proto3" json:"node1_inf,omitempty"`
	Node2Inf string `protobuf:"bytes,5,opt,name=node2_inf,json=node2Inf,proto3" json:"node2_inf,omitempty"`
}

func (x *Link) Reset() {
	*x = Link{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Link) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Link) ProtoMessage() {}

func (x *Link) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Link.ProtoReflect.Descriptor instead.
func (*Link) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{12}
}

func (x *Link) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Link) GetNode1() *Pod {
	if x != nil {
		return x.Node1
	}
	return nil
}

func (x *Link) GetNode2() *Pod {
	if x != nil {
		return x.Node2
	}
	return nil
}

func (x *Link) GetNode1Inf() string {
	if x != nil {
		return x.Node1Inf
	}
	return ""
}

func (x *Link) GetNode2Inf() string {
	if x != nil {
		return x.Node2Inf
	}
	return ""
}

type Pod struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Pod) Reset() {
	*x = Pod{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rpctest_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pod) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pod) ProtoMessage() {}

func (x *Pod) ProtoReflect() protoreflect.Message {
	mi := &file_rpctest_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pod.ProtoReflect.Descriptor instead.
func (*Pod) Descriptor() ([]byte, []int) {
	return file_rpctest_proto_rawDescGZIP(), []int{13}
}

func (x *Pod) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_rpctest_proto protoreflect.FileDescriptor

var file_rpctest_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x22, 0xba, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x14, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x2a,
	0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x12, 0x2a, 0x0a, 0x07, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x72, 0x70,
	0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x07, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2d, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xff, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x12, 0x40, 0x0a, 0x0f, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72, 0x70, 0x63,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0f, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x72,
	0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x0d, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x5f,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x72,
	0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x0c, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x34, 0x0a, 0x0b, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x52, 0x0a, 0x65, 0x6d, 0x75,
	0x6e, 0x65, 0x74, 0x54, 0x65, 0x73, 0x74, 0x22, 0x49, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x72, 0x70,
	0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x67, 0x65, 0x74, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x22, 0xd9, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x43, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18,
	0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x39,
	0x0a, 0x0e, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x06, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x75, 0x6e, 0x65,
	0x74, 0x52, 0x06, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x22, 0x10, 0x0a, 0x0e, 0x45, 0x6d, 0x75,
	0x6e, 0x65, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x0c, 0x45,
	0x6d, 0x75, 0x6e, 0x65, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x65,
	0x6d, 0x75, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70,
	0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x52, 0x06, 0x65, 0x6d,
	0x75, 0x6e, 0x65, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x54, 0x65,
	0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x11, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3a, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x65, 0x6d, 0x75,
	0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x72, 0x70, 0x63, 0x74,
	0x65, 0x73, 0x74, 0x2e, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x52, 0x06, 0x65, 0x6d, 0x75, 0x6e,
	0x65, 0x74, 0x22, 0x63, 0x0a, 0x06, 0x45, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x12, 0x20, 0x0a, 0x04,
	0x70, 0x6f, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x70, 0x63,
	0x74, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x04, 0x70, 0x6f, 0x64, 0x73, 0x12, 0x23,
	0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x05, 0x6c, 0x69,
	0x6e, 0x6b, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x9c, 0x01, 0x0a, 0x04, 0x4c, 0x69, 0x6e, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x31, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x50, 0x6f,
	0x64, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x31, 0x12, 0x22, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65,
	0x32, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x50, 0x6f, 0x64, 0x52, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x32, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x6f, 0x64, 0x65, 0x31, 0x5f, 0x69, 0x6e, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x31, 0x49, 0x6e, 0x66, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64,
	0x65, 0x32, 0x5f, 0x69, 0x6e, 0x66, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f,
	0x64, 0x65, 0x32, 0x49, 0x6e, 0x66, 0x22, 0x19, 0x0a, 0x03, 0x50, 0x6f, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x2a, 0x9c, 0x01, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x13, 0x0a, 0x0f, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x5f, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74,
	0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x65,
	0x6d, 0x75, 0x6e, 0x65, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x0f,
	0x0a, 0x0b, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x10, 0x03, 0x12,
	0x13, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x5f, 0x65, 0x6d, 0x75, 0x6e, 0x65, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x05, 0x12, 0x14, 0x0a, 0x10, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x10, 0x06,
	0x32, 0x38, 0x0a, 0x07, 0x4d, 0x6f, 0x63, 0x6b, 0x6e, 0x65, 0x74, 0x12, 0x2d, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65, 0x73, 0x74,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x74, 0x65,
	0x73, 0x74, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x72,
	0x70, 0x63, 0x74, 0x65, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rpctest_proto_rawDescOnce sync.Once
	file_rpctest_proto_rawDescData = file_rpctest_proto_rawDesc
)

func file_rpctest_proto_rawDescGZIP() []byte {
	file_rpctest_proto_rawDescOnce.Do(func() {
		file_rpctest_proto_rawDescData = protoimpl.X.CompressGZIP(file_rpctest_proto_rawDescData)
	})
	return file_rpctest_proto_rawDescData
}

var file_rpctest_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_rpctest_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_rpctest_proto_goTypes = []interface{}{
	(MessageType)(0),        // 0: rpctest.MessageType
	(*Message)(nil),         // 1: rpctest.Message
	(*Command)(nil),         // 2: rpctest.Command
	(*Request)(nil),         // 3: rpctest.Request
	(*Response)(nil),        // 4: rpctest.Response
	(*EmunetCreation)(nil),  // 5: rpctest.EmunetCreation
	(*EmunetDeletion)(nil),  // 6: rpctest.EmunetDeletion
	(*EmunetUpdate)(nil),    // 7: rpctest.EmunetUpdate
	(*EmunetTest)(nil),      // 8: rpctest.EmunetTest
	(*GetEmunetInfo)(nil),   // 9: rpctest.GetEmunetInfo
	(*CommandResponse)(nil), // 10: rpctest.CommandResponse
	(*RequestResponse)(nil), // 11: rpctest.RequestResponse
	(*Emunet)(nil),          // 12: rpctest.Emunet
	(*Link)(nil),            // 13: rpctest.Link
	(*Pod)(nil),             // 14: rpctest.Pod
}
var file_rpctest_proto_depIdxs = []int32{
	0,  // 0: rpctest.Message.type:type_name -> rpctest.MessageType
	2,  // 1: rpctest.Message.command:type_name -> rpctest.Command
	3,  // 2: rpctest.Message.request:type_name -> rpctest.Request
	4,  // 3: rpctest.Message.response:type_name -> rpctest.Response
	5,  // 4: rpctest.Command.emunet_creation:type_name -> rpctest.EmunetCreation
	6,  // 5: rpctest.Command.emunet_deletion:type_name -> rpctest.EmunetDeletion
	7,  // 6: rpctest.Command.emunet_update:type_name -> rpctest.EmunetUpdate
	8,  // 7: rpctest.Command.emunet_test:type_name -> rpctest.EmunetTest
	9,  // 8: rpctest.Request.get_emunet_info:type_name -> rpctest.GetEmunetInfo
	10, // 9: rpctest.Response.command_response:type_name -> rpctest.CommandResponse
	11, // 10: rpctest.Response.request_response:type_name -> rpctest.RequestResponse
	12, // 11: rpctest.EmunetCreation.emunet:type_name -> rpctest.Emunet
	12, // 12: rpctest.EmunetUpdate.emunet:type_name -> rpctest.Emunet
	12, // 13: rpctest.RequestResponse.emunet:type_name -> rpctest.Emunet
	14, // 14: rpctest.Emunet.pods:type_name -> rpctest.Pod
	13, // 15: rpctest.Emunet.links:type_name -> rpctest.Link
	14, // 16: rpctest.Link.node1:type_name -> rpctest.Pod
	14, // 17: rpctest.Link.node2:type_name -> rpctest.Pod
	1,  // 18: rpctest.Mocknet.process:input_type -> rpctest.Message
	1,  // 19: rpctest.Mocknet.process:output_type -> rpctest.Message
	19, // [19:20] is the sub-list for method output_type
	18, // [18:19] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_rpctest_proto_init() }
func file_rpctest_proto_init() {
	if File_rpctest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rpctest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpctest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Command); i {
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
		file_rpctest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_rpctest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_rpctest_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmunetCreation); i {
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
		file_rpctest_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmunetDeletion); i {
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
		file_rpctest_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmunetUpdate); i {
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
		file_rpctest_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmunetTest); i {
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
		file_rpctest_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEmunetInfo); i {
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
		file_rpctest_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommandResponse); i {
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
		file_rpctest_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestResponse); i {
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
		file_rpctest_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Emunet); i {
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
		file_rpctest_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Link); i {
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
		file_rpctest_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pod); i {
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
			RawDescriptor: file_rpctest_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rpctest_proto_goTypes,
		DependencyIndexes: file_rpctest_proto_depIdxs,
		EnumInfos:         file_rpctest_proto_enumTypes,
		MessageInfos:      file_rpctest_proto_msgTypes,
	}.Build()
	File_rpctest_proto = out.File
	file_rpctest_proto_rawDesc = nil
	file_rpctest_proto_goTypes = nil
	file_rpctest_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MocknetClient is the client API for Mocknet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MocknetClient interface {
	Process(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type mocknetClient struct {
	cc grpc.ClientConnInterface
}

func NewMocknetClient(cc grpc.ClientConnInterface) MocknetClient {
	return &mocknetClient{cc}
}

func (c *mocknetClient) Process(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/rpctest.Mocknet/process", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MocknetServer is the server API for Mocknet service.
type MocknetServer interface {
	Process(context.Context, *Message) (*Message, error)
}

// UnimplementedMocknetServer can be embedded to have forward compatible implementations.
type UnimplementedMocknetServer struct {
}

func (*UnimplementedMocknetServer) Process(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Process not implemented")
}

func RegisterMocknetServer(s *grpc.Server, srv MocknetServer) {
	s.RegisterService(&_Mocknet_serviceDesc, srv)
}

func _Mocknet_Process_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MocknetServer).Process(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpctest.Mocknet/Process",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MocknetServer).Process(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

var _Mocknet_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpctest.Mocknet",
	HandlerType: (*MocknetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "process",
			Handler:    _Mocknet_Process_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpctest.proto",
}
