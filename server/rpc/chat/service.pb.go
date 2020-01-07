// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat/service.proto

package chat

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("chat/service.proto", fileDescriptor_6e7a84c7f0c34e37) }

var fileDescriptor_6e7a84c7f0c34e37 = []byte{
	// 167 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x48, 0x2c,
	0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x01, 0x89, 0x49, 0x49, 0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0xea, 0x83, 0xc5, 0x92, 0x4a, 0xd3,
	0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0x21, 0x4a, 0xa4, 0xe4, 0xd0, 0x25, 0xcb, 0x8b, 0x12, 0x0b,
	0x0a, 0x52, 0x8b, 0x8a, 0x21, 0xf2, 0x46, 0x53, 0x18, 0xb9, 0xb8, 0x9d, 0x33, 0x12, 0x4b, 0x82,
	0x21, 0x06, 0x0b, 0xd9, 0x71, 0xb1, 0x04, 0xa7, 0xe6, 0xa5, 0x08, 0xc9, 0xe8, 0x41, 0x34, 0xea,
	0xc1, 0x34, 0xea, 0x05, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0x87, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x89,
	0x61, 0xc8, 0xba, 0x82, 0xec, 0x54, 0x62, 0x10, 0x72, 0xe5, 0xe2, 0x0c, 0x2e, 0x4d, 0x2a, 0x4e,
	0x2e, 0xca, 0x4c, 0x4a, 0x15, 0xc2, 0xa1, 0x4c, 0x0a, 0xaf, 0xe1, 0x4a, 0x0c, 0x06, 0x8c, 0x4e,
	0x6c, 0x51, 0x60, 0xbf, 0x25, 0xb1, 0x81, 0x55, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a,
	0x78, 0x3c, 0x1c, 0xfe, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	// Send sends message to the server
	Send(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error)
	// Subscribe is streaming method to get echo messages from the server
	Subscribe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ChatService_SubscribeClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) Send(ctx context.Context, in *wrappers.StringValue, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/chat.ChatService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Subscribe(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (ChatService_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/chat.ChatService/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_SubscribeClient interface {
	Recv() (*wrappers.StringValue, error)
	grpc.ClientStream
}

type chatServiceSubscribeClient struct {
	grpc.ClientStream
}

func (x *chatServiceSubscribeClient) Recv() (*wrappers.StringValue, error) {
	m := new(wrappers.StringValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	// Send sends message to the server
	Send(context.Context, *wrappers.StringValue) (*empty.Empty, error)
	// Subscribe is streaming method to get echo messages from the server
	Subscribe(*empty.Empty, ChatService_SubscribeServer) error
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) Send(ctx context.Context, req *wrappers.StringValue) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedChatServiceServer) Subscribe(req *empty.Empty, srv ChatService_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrappers.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/chat.ChatService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).Send(ctx, req.(*wrappers.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).Subscribe(m, &chatServiceSubscribeServer{stream})
}

type ChatService_SubscribeServer interface {
	Send(*wrappers.StringValue) error
	grpc.ServerStream
}

type chatServiceSubscribeServer struct {
	grpc.ServerStream
}

func (x *chatServiceSubscribeServer) Send(m *wrappers.StringValue) error {
	return x.ServerStream.SendMsg(m)
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _ChatService_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _ChatService_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "chat/service.proto",
}
