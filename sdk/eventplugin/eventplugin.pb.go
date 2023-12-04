// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: sdk/eventplugin/eventplugin.proto

package eventplugin

import (
	_ "github.com/hashicorp/vault/sdk/logical"
	pb "github.com/hashicorp/vault/sdk/plugin/pb"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type InitializeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitializeRequest) Reset() {
	*x = InitializeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeRequest) ProtoMessage() {}

func (x *InitializeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeRequest.ProtoReflect.Descriptor instead.
func (*InitializeRequest) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{0}
}

type InitializeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitializeResponse) Reset() {
	*x = InitializeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InitializeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitializeResponse) ProtoMessage() {}

func (x *InitializeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitializeResponse.ProtoReflect.Descriptor instead.
func (*InitializeResponse) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{1}
}

type SubscriptionEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// only a single subscription ID will be used in a stream, so the value MAY be omitted in subsequent messages,
	// but is REQUIRED for the first message. This is included for future-proofing having a plugin handle multiple
	// subscriptions per plugin instance.
	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	EventJson      string `protobuf:"bytes,2,opt,name=eventJson,proto3" json:"eventJson,omitempty"`
}

func (x *SubscriptionEvent) Reset() {
	*x = SubscriptionEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionEvent) ProtoMessage() {}

func (x *SubscriptionEvent) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionEvent.ProtoReflect.Descriptor instead.
func (*SubscriptionEvent) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{2}
}

func (x *SubscriptionEvent) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *SubscriptionEvent) GetEventJson() string {
	if x != nil {
		return x.EventJson
	}
	return ""
}

type SubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// parameters to set up the request, e.g., queue name, API key
	Config *structpb.Struct `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
	// globally unique subscription ID, generated by Vault
	SubscriptionId string `protobuf:"bytes,2,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
	// whether to verify the connection immediately.
	VerifyConnection bool `protobuf:"varint,3,opt,name=verify_connection,json=verifyConnection,proto3" json:"verify_connection,omitempty"`
}

func (x *SubscribeRequest) Reset() {
	*x = SubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeRequest) ProtoMessage() {}

func (x *SubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeRequest.ProtoReflect.Descriptor instead.
func (*SubscribeRequest) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{3}
}

func (x *SubscribeRequest) GetConfig() *structpb.Struct {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *SubscribeRequest) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

func (x *SubscribeRequest) GetVerifyConnection() bool {
	if x != nil {
		return x.VerifyConnection
	}
	return false
}

type UnsubscribeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubscriptionId string `protobuf:"bytes,1,opt,name=subscription_id,json=subscriptionId,proto3" json:"subscription_id,omitempty"`
}

func (x *UnsubscribeRequest) Reset() {
	*x = UnsubscribeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeRequest) ProtoMessage() {}

func (x *UnsubscribeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeRequest.ProtoReflect.Descriptor instead.
func (*UnsubscribeRequest) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{4}
}

func (x *UnsubscribeRequest) GetSubscriptionId() string {
	if x != nil {
		return x.SubscriptionId
	}
	return ""
}

type UnsubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnsubscribeResponse) Reset() {
	*x = UnsubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnsubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnsubscribeResponse) ProtoMessage() {}

func (x *UnsubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnsubscribeResponse.ProtoReflect.Descriptor instead.
func (*UnsubscribeResponse) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{5}
}

type SubscribeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Err *pb.ProtoError `protobuf:"bytes,1,opt,name=err,proto3" json:"err,omitempty"`
	// If true, Vault can retry sending events with this subscription on error.
	// If false, the subscription is canceled on error.
	Retry bool `protobuf:"varint,2,opt,name=retry,proto3" json:"retry,omitempty"`
}

func (x *SubscribeResponse) Reset() {
	*x = SubscribeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscribeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscribeResponse) ProtoMessage() {}

func (x *SubscribeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscribeResponse.ProtoReflect.Descriptor instead.
func (*SubscribeResponse) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{6}
}

func (x *SubscribeResponse) GetErr() *pb.ProtoError {
	if x != nil {
		return x.Err
	}
	return nil
}

func (x *SubscribeResponse) GetRetry() bool {
	if x != nil {
		return x.Retry
	}
	return false
}

type SendSubscriptionEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SendSubscriptionEventsResponse) Reset() {
	*x = SendSubscriptionEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendSubscriptionEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendSubscriptionEventsResponse) ProtoMessage() {}

func (x *SendSubscriptionEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendSubscriptionEventsResponse.ProtoReflect.Descriptor instead.
func (*SendSubscriptionEventsResponse) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{7}
}

type TypeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TypeRequest) Reset() {
	*x = TypeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeRequest) ProtoMessage() {}

func (x *TypeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeRequest.ProtoReflect.Descriptor instead.
func (*TypeRequest) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{8}
}

type TypeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// plugin type, e.g., "aws-sqs"
	PluginType string `protobuf:"bytes,1,opt,name=plugin_type,json=pluginType,proto3" json:"plugin_type,omitempty"`
	// plugin version, e.g., "v1.0.0"
	PluginVersion string `protobuf:"bytes,2,opt,name=plugin_version,json=pluginVersion,proto3" json:"plugin_version,omitempty"`
}

func (x *TypeResponse) Reset() {
	*x = TypeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeResponse) ProtoMessage() {}

func (x *TypeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeResponse.ProtoReflect.Descriptor instead.
func (*TypeResponse) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{9}
}

func (x *TypeResponse) GetPluginType() string {
	if x != nil {
		return x.PluginType
	}
	return ""
}

func (x *TypeResponse) GetPluginVersion() string {
	if x != nil {
		return x.PluginVersion
	}
	return ""
}

type CloseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseRequest) Reset() {
	*x = CloseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseRequest) ProtoMessage() {}

func (x *CloseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseRequest.ProtoReflect.Descriptor instead.
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{10}
}

type CloseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CloseResponse) Reset() {
	*x = CloseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CloseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CloseResponse) ProtoMessage() {}

func (x *CloseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sdk_eventplugin_eventplugin_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CloseResponse.ProtoReflect.Descriptor instead.
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return file_sdk_eventplugin_eventplugin_proto_rawDescGZIP(), []int{11}
}

var File_sdk_eventplugin_eventplugin_proto protoreflect.FileDescriptor

var file_sdk_eventplugin_eventplugin_proto_rawDesc = []byte{
	0x0a, 0x21, 0x73, 0x64, 0x6b, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1b, 0x73, 0x64, 0x6b, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x70, 0x62,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x73, 0x64, 0x6b, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x2f, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x14, 0x0a, 0x12,
	0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x5a, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4a, 0x73, 0x6f, 0x6e, 0x22, 0x99,
	0x01, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x06, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2b, 0x0a,
	0x11, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x76, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x3d, 0x0a, 0x12, 0x55, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x27, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x55, 0x6e, 0x73,
	0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x4b, 0x0a, 0x11, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x03, 0x65, 0x72, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x22, 0x20, 0x0a,
	0x1e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x0d, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56,
	0x0a, 0x0c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x25, 0x0a, 0x0e, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x94, 0x04, 0x0a, 0x1b, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x53, 0x0a, 0x0a, 0x49, 0x6e, 0x69, 0x74, 0x69,
	0x61, 0x6c, 0x69, 0x7a, 0x65, 0x12, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x6c, 0x69, 0x7a,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61,
	0x6c, 0x69, 0x7a, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x09,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x20, 0x2e, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63,
	0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62,
	0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6d,
	0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x21, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x1a, 0x2e, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x28, 0x01, 0x12, 0x56, 0x0a,
	0x0b, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x12, 0x22, 0x2e, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e,
	0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x23, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x6e, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1b, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x79, 0x70, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x43, 0x6c, 0x6f, 0x73,
	0x65, 0x12, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2c,
	0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73,
	0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2f, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x2f, 0x73, 0x64, 0x6b,
	0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sdk_eventplugin_eventplugin_proto_rawDescOnce sync.Once
	file_sdk_eventplugin_eventplugin_proto_rawDescData = file_sdk_eventplugin_eventplugin_proto_rawDesc
)

func file_sdk_eventplugin_eventplugin_proto_rawDescGZIP() []byte {
	file_sdk_eventplugin_eventplugin_proto_rawDescOnce.Do(func() {
		file_sdk_eventplugin_eventplugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_sdk_eventplugin_eventplugin_proto_rawDescData)
	})
	return file_sdk_eventplugin_eventplugin_proto_rawDescData
}

var file_sdk_eventplugin_eventplugin_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_sdk_eventplugin_eventplugin_proto_goTypes = []interface{}{
	(*InitializeRequest)(nil),              // 0: eventplugin.v1.InitializeRequest
	(*InitializeResponse)(nil),             // 1: eventplugin.v1.InitializeResponse
	(*SubscriptionEvent)(nil),              // 2: eventplugin.v1.SubscriptionEvent
	(*SubscribeRequest)(nil),               // 3: eventplugin.v1.SubscribeRequest
	(*UnsubscribeRequest)(nil),             // 4: eventplugin.v1.UnsubscribeRequest
	(*UnsubscribeResponse)(nil),            // 5: eventplugin.v1.UnsubscribeResponse
	(*SubscribeResponse)(nil),              // 6: eventplugin.v1.SubscribeResponse
	(*SendSubscriptionEventsResponse)(nil), // 7: eventplugin.v1.SendSubscriptionEventsResponse
	(*TypeRequest)(nil),                    // 8: eventplugin.v1.TypeRequest
	(*TypeResponse)(nil),                   // 9: eventplugin.v1.TypeResponse
	(*CloseRequest)(nil),                   // 10: eventplugin.v1.CloseRequest
	(*CloseResponse)(nil),                  // 11: eventplugin.v1.CloseResponse
	(*structpb.Struct)(nil),                // 12: google.protobuf.Struct
	(*pb.ProtoError)(nil),                  // 13: pb.ProtoError
}
var file_sdk_eventplugin_eventplugin_proto_depIdxs = []int32{
	12, // 0: eventplugin.v1.SubscribeRequest.config:type_name -> google.protobuf.Struct
	13, // 1: eventplugin.v1.SubscribeResponse.err:type_name -> pb.ProtoError
	0,  // 2: eventplugin.v1.EventSubscribePluginService.Initialize:input_type -> eventplugin.v1.InitializeRequest
	3,  // 3: eventplugin.v1.EventSubscribePluginService.Subscribe:input_type -> eventplugin.v1.SubscribeRequest
	2,  // 4: eventplugin.v1.EventSubscribePluginService.SendSubscriptionEvents:input_type -> eventplugin.v1.SubscriptionEvent
	4,  // 5: eventplugin.v1.EventSubscribePluginService.Unsubscribe:input_type -> eventplugin.v1.UnsubscribeRequest
	8,  // 6: eventplugin.v1.EventSubscribePluginService.Type:input_type -> eventplugin.v1.TypeRequest
	10, // 7: eventplugin.v1.EventSubscribePluginService.Close:input_type -> eventplugin.v1.CloseRequest
	1,  // 8: eventplugin.v1.EventSubscribePluginService.Initialize:output_type -> eventplugin.v1.InitializeResponse
	6,  // 9: eventplugin.v1.EventSubscribePluginService.Subscribe:output_type -> eventplugin.v1.SubscribeResponse
	7,  // 10: eventplugin.v1.EventSubscribePluginService.SendSubscriptionEvents:output_type -> eventplugin.v1.SendSubscriptionEventsResponse
	5,  // 11: eventplugin.v1.EventSubscribePluginService.Unsubscribe:output_type -> eventplugin.v1.UnsubscribeResponse
	9,  // 12: eventplugin.v1.EventSubscribePluginService.Type:output_type -> eventplugin.v1.TypeResponse
	11, // 13: eventplugin.v1.EventSubscribePluginService.Close:output_type -> eventplugin.v1.CloseResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_sdk_eventplugin_eventplugin_proto_init() }
func file_sdk_eventplugin_eventplugin_proto_init() {
	if File_sdk_eventplugin_eventplugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sdk_eventplugin_eventplugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeRequest); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InitializeResponse); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscriptionEvent); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeRequest); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeRequest); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnsubscribeResponse); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubscribeResponse); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendSubscriptionEventsResponse); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeRequest); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeResponse); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseRequest); i {
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
		file_sdk_eventplugin_eventplugin_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CloseResponse); i {
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
			RawDescriptor: file_sdk_eventplugin_eventplugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sdk_eventplugin_eventplugin_proto_goTypes,
		DependencyIndexes: file_sdk_eventplugin_eventplugin_proto_depIdxs,
		MessageInfos:      file_sdk_eventplugin_eventplugin_proto_msgTypes,
	}.Build()
	File_sdk_eventplugin_eventplugin_proto = out.File
	file_sdk_eventplugin_eventplugin_proto_rawDesc = nil
	file_sdk_eventplugin_eventplugin_proto_goTypes = nil
	file_sdk_eventplugin_eventplugin_proto_depIdxs = nil
}