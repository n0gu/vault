// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package eventplugin

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// EventSubscribePluginServiceClient is the client API for EventSubscribePluginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EventSubscribePluginServiceClient interface {
	Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*InitializeResponse, error)
	// Start a new subscription.
	Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error)
	// Used by Vault to send events to a subscription. Only one stream is allowed per subscription.
	// Ending this stream will stop the events, but won't necessarily mean tha the subscription information
	// (e.g., config data like API keys) should be deleted.
	SendSubscriptionEvents(ctx context.Context, opts ...grpc.CallOption) (EventSubscribePluginService_SendSubscriptionEventsClient, error)
	// Cause the subscription to be deleted and any related information about it cleaned up.
	Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error)
	Type(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*TypeResponse, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
}

type eventSubscribePluginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEventSubscribePluginServiceClient(cc grpc.ClientConnInterface) EventSubscribePluginServiceClient {
	return &eventSubscribePluginServiceClient{cc}
}

func (c *eventSubscribePluginServiceClient) Initialize(ctx context.Context, in *InitializeRequest, opts ...grpc.CallOption) (*InitializeResponse, error) {
	out := new(InitializeResponse)
	err := c.cc.Invoke(ctx, "/eventplugin.v1.EventSubscribePluginService/Initialize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSubscribePluginServiceClient) Subscribe(ctx context.Context, in *SubscribeRequest, opts ...grpc.CallOption) (*SubscribeResponse, error) {
	out := new(SubscribeResponse)
	err := c.cc.Invoke(ctx, "/eventplugin.v1.EventSubscribePluginService/Subscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSubscribePluginServiceClient) SendSubscriptionEvents(ctx context.Context, opts ...grpc.CallOption) (EventSubscribePluginService_SendSubscriptionEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &EventSubscribePluginService_ServiceDesc.Streams[0], "/eventplugin.v1.EventSubscribePluginService/SendSubscriptionEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventSubscribePluginServiceSendSubscriptionEventsClient{stream}
	return x, nil
}

type EventSubscribePluginService_SendSubscriptionEventsClient interface {
	Send(*SubscriptionEvent) error
	CloseAndRecv() (*SendSubscriptionEventsResponse, error)
	grpc.ClientStream
}

type eventSubscribePluginServiceSendSubscriptionEventsClient struct {
	grpc.ClientStream
}

func (x *eventSubscribePluginServiceSendSubscriptionEventsClient) Send(m *SubscriptionEvent) error {
	return x.ClientStream.SendMsg(m)
}

func (x *eventSubscribePluginServiceSendSubscriptionEventsClient) CloseAndRecv() (*SendSubscriptionEventsResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SendSubscriptionEventsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventSubscribePluginServiceClient) Unsubscribe(ctx context.Context, in *UnsubscribeRequest, opts ...grpc.CallOption) (*UnsubscribeResponse, error) {
	out := new(UnsubscribeResponse)
	err := c.cc.Invoke(ctx, "/eventplugin.v1.EventSubscribePluginService/Unsubscribe", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSubscribePluginServiceClient) Type(ctx context.Context, in *TypeRequest, opts ...grpc.CallOption) (*TypeResponse, error) {
	out := new(TypeResponse)
	err := c.cc.Invoke(ctx, "/eventplugin.v1.EventSubscribePluginService/Type", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventSubscribePluginServiceClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/eventplugin.v1.EventSubscribePluginService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventSubscribePluginServiceServer is the server API for EventSubscribePluginService service.
// All implementations must embed UnimplementedEventSubscribePluginServiceServer
// for forward compatibility
type EventSubscribePluginServiceServer interface {
	Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error)
	// Start a new subscription.
	Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error)
	// Used by Vault to send events to a subscription. Only one stream is allowed per subscription.
	// Ending this stream will stop the events, but won't necessarily mean tha the subscription information
	// (e.g., config data like API keys) should be deleted.
	SendSubscriptionEvents(EventSubscribePluginService_SendSubscriptionEventsServer) error
	// Cause the subscription to be deleted and any related information about it cleaned up.
	Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error)
	Type(context.Context, *TypeRequest) (*TypeResponse, error)
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
	mustEmbedUnimplementedEventSubscribePluginServiceServer()
}

// UnimplementedEventSubscribePluginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEventSubscribePluginServiceServer struct {
}

func (UnimplementedEventSubscribePluginServiceServer) Initialize(context.Context, *InitializeRequest) (*InitializeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Initialize not implemented")
}
func (UnimplementedEventSubscribePluginServiceServer) Subscribe(context.Context, *SubscribeRequest) (*SubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedEventSubscribePluginServiceServer) SendSubscriptionEvents(EventSubscribePluginService_SendSubscriptionEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SendSubscriptionEvents not implemented")
}
func (UnimplementedEventSubscribePluginServiceServer) Unsubscribe(context.Context, *UnsubscribeRequest) (*UnsubscribeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsubscribe not implemented")
}
func (UnimplementedEventSubscribePluginServiceServer) Type(context.Context, *TypeRequest) (*TypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Type not implemented")
}
func (UnimplementedEventSubscribePluginServiceServer) Close(context.Context, *CloseRequest) (*CloseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedEventSubscribePluginServiceServer) mustEmbedUnimplementedEventSubscribePluginServiceServer() {
}

// UnsafeEventSubscribePluginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EventSubscribePluginServiceServer will
// result in compilation errors.
type UnsafeEventSubscribePluginServiceServer interface {
	mustEmbedUnimplementedEventSubscribePluginServiceServer()
}

func RegisterEventSubscribePluginServiceServer(s grpc.ServiceRegistrar, srv EventSubscribePluginServiceServer) {
	s.RegisterService(&EventSubscribePluginService_ServiceDesc, srv)
}

func _EventSubscribePluginService_Initialize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitializeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscribePluginServiceServer).Initialize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventplugin.v1.EventSubscribePluginService/Initialize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscribePluginServiceServer).Initialize(ctx, req.(*InitializeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSubscribePluginService_Subscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscribePluginServiceServer).Subscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventplugin.v1.EventSubscribePluginService/Subscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscribePluginServiceServer).Subscribe(ctx, req.(*SubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSubscribePluginService_SendSubscriptionEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EventSubscribePluginServiceServer).SendSubscriptionEvents(&eventSubscribePluginServiceSendSubscriptionEventsServer{stream})
}

type EventSubscribePluginService_SendSubscriptionEventsServer interface {
	SendAndClose(*SendSubscriptionEventsResponse) error
	Recv() (*SubscriptionEvent, error)
	grpc.ServerStream
}

type eventSubscribePluginServiceSendSubscriptionEventsServer struct {
	grpc.ServerStream
}

func (x *eventSubscribePluginServiceSendSubscriptionEventsServer) SendAndClose(m *SendSubscriptionEventsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *eventSubscribePluginServiceSendSubscriptionEventsServer) Recv() (*SubscriptionEvent, error) {
	m := new(SubscriptionEvent)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _EventSubscribePluginService_Unsubscribe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnsubscribeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscribePluginServiceServer).Unsubscribe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventplugin.v1.EventSubscribePluginService/Unsubscribe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscribePluginServiceServer).Unsubscribe(ctx, req.(*UnsubscribeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSubscribePluginService_Type_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscribePluginServiceServer).Type(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventplugin.v1.EventSubscribePluginService/Type",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscribePluginServiceServer).Type(ctx, req.(*TypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventSubscribePluginService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventSubscribePluginServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventplugin.v1.EventSubscribePluginService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventSubscribePluginServiceServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EventSubscribePluginService_ServiceDesc is the grpc.ServiceDesc for EventSubscribePluginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EventSubscribePluginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "eventplugin.v1.EventSubscribePluginService",
	HandlerType: (*EventSubscribePluginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Initialize",
			Handler:    _EventSubscribePluginService_Initialize_Handler,
		},
		{
			MethodName: "Subscribe",
			Handler:    _EventSubscribePluginService_Subscribe_Handler,
		},
		{
			MethodName: "Unsubscribe",
			Handler:    _EventSubscribePluginService_Unsubscribe_Handler,
		},
		{
			MethodName: "Type",
			Handler:    _EventSubscribePluginService_Type_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _EventSubscribePluginService_Close_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendSubscriptionEvents",
			Handler:       _EventSubscribePluginService_SendSubscriptionEvents_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "sdk/eventplugin/eventplugin.proto",
}