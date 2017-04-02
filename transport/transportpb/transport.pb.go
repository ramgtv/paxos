// Code generated by protoc-gen-go.
// source: transport/transportpb/transport.proto
// DO NOT EDIT!

/*
Package transportpb is a generated protocol buffer package.

It is generated from these files:
	transport/transportpb/transport.proto

It has these top-level messages:
	Empty
*/
package transportpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import paxospb "github.com/nvanbenschoten/paxos/paxos/paxospb"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Empty is an empty message. It is identical to google.protobuf.Empty, but
// permits future modifications because it is custom.
type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*Empty)(nil), "transportpb.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PaxosTransport service

type PaxosTransportClient interface {
	DeliverMessage(ctx context.Context, opts ...grpc.CallOption) (PaxosTransport_DeliverMessageClient, error)
}

type paxosTransportClient struct {
	cc *grpc.ClientConn
}

func NewPaxosTransportClient(cc *grpc.ClientConn) PaxosTransportClient {
	return &paxosTransportClient{cc}
}

func (c *paxosTransportClient) DeliverMessage(ctx context.Context, opts ...grpc.CallOption) (PaxosTransport_DeliverMessageClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PaxosTransport_serviceDesc.Streams[0], c.cc, "/transportpb.PaxosTransport/DeliverMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &paxosTransportDeliverMessageClient{stream}
	return x, nil
}

type PaxosTransport_DeliverMessageClient interface {
	Send(*paxospb.Message) error
	CloseAndRecv() (*Empty, error)
	grpc.ClientStream
}

type paxosTransportDeliverMessageClient struct {
	grpc.ClientStream
}

func (x *paxosTransportDeliverMessageClient) Send(m *paxospb.Message) error {
	return x.ClientStream.SendMsg(m)
}

func (x *paxosTransportDeliverMessageClient) CloseAndRecv() (*Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for PaxosTransport service

type PaxosTransportServer interface {
	DeliverMessage(PaxosTransport_DeliverMessageServer) error
}

func RegisterPaxosTransportServer(s *grpc.Server, srv PaxosTransportServer) {
	s.RegisterService(&_PaxosTransport_serviceDesc, srv)
}

func _PaxosTransport_DeliverMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PaxosTransportServer).DeliverMessage(&paxosTransportDeliverMessageServer{stream})
}

type PaxosTransport_DeliverMessageServer interface {
	SendAndClose(*Empty) error
	Recv() (*paxospb.Message, error)
	grpc.ServerStream
}

type paxosTransportDeliverMessageServer struct {
	grpc.ServerStream
}

func (x *paxosTransportDeliverMessageServer) SendAndClose(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *paxosTransportDeliverMessageServer) Recv() (*paxospb.Message, error) {
	m := new(paxospb.Message)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _PaxosTransport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transportpb.PaxosTransport",
	HandlerType: (*PaxosTransportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DeliverMessage",
			Handler:       _PaxosTransport_DeliverMessage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "transport/transportpb/transport.proto",
}

// Client API for ClientService service

type ClientServiceClient interface {
	ClientUpdate(ctx context.Context, in *paxospb.ClientUpdate, opts ...grpc.CallOption) (*paxospb.GloballyOrderedUpdate, error)
}

type clientServiceClient struct {
	cc *grpc.ClientConn
}

func NewClientServiceClient(cc *grpc.ClientConn) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) ClientUpdate(ctx context.Context, in *paxospb.ClientUpdate, opts ...grpc.CallOption) (*paxospb.GloballyOrderedUpdate, error) {
	out := new(paxospb.GloballyOrderedUpdate)
	err := grpc.Invoke(ctx, "/transportpb.ClientService/ClientUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClientService service

type ClientServiceServer interface {
	ClientUpdate(context.Context, *paxospb.ClientUpdate) (*paxospb.GloballyOrderedUpdate, error)
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_ClientUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(paxospb.ClientUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).ClientUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/transportpb.ClientService/ClientUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).ClientUpdate(ctx, req.(*paxospb.ClientUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "transportpb.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientUpdate",
			Handler:    _ClientService_ClientUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "transport/transportpb/transport.proto",
}

func init() { proto.RegisterFile("transport/transportpb/transport.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 211 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xeb, 0x41, 0x85, 0xa8, 0x45, 0x02, 0x5e, 0x7a, 0xf0, 0x50, 0x10, 0x3c, 0xa5, 0x50,
	0x4f, 0x7a, 0x55, 0xe9, 0x45, 0xd9, 0x60, 0x1b, 0xec, 0x9a, 0xb4, 0x8f, 0x36, 0x90, 0x26, 0x21,
	0x79, 0x2b, 0xeb, 0x7f, 0x3f, 0xd6, 0x74, 0x5d, 0x2e, 0xef, 0x7d, 0xfc, 0x3e, 0xf8, 0xf1, 0x91,
	0x37, 0x74, 0x5c, 0x7b, 0x6b, 0x1c, 0x16, 0x4b, 0xb2, 0xe2, 0x9a, 0x99, 0x75, 0x06, 0x0d, 0x7d,
	0x88, 0xca, 0xec, 0xb3, 0x95, 0xd8, 0x1d, 0x04, 0xab, 0x4d, 0x5f, 0xe8, 0x81, 0x6b, 0x01, 0xda,
	0xd7, 0x9d, 0x41, 0xd0, 0x85, 0xe5, 0x47, 0xe3, 0xe3, 0x6b, 0x45, 0xf8, 0xc1, 0x93, 0xdf, 0x93,
	0xdb, 0xdf, 0xde, 0xe2, 0x58, 0xfe, 0x91, 0x74, 0x7d, 0xe6, 0xdb, 0x8b, 0x97, 0x7e, 0x91, 0xf4,
	0x07, 0x94, 0x1c, 0xc0, 0xfd, 0x83, 0xf7, 0xbc, 0x05, 0xfa, 0xcc, 0x66, 0x05, 0x9b, 0x49, 0x46,
	0x59, 0xb4, 0x83, 0x4d, 0xa6, 0x3c, 0x79, 0xbf, 0x29, 0xf7, 0xe4, 0xe9, 0x5b, 0x49, 0xd0, 0xb8,
	0x01, 0x37, 0xc8, 0x1a, 0x68, 0x45, 0x1e, 0x03, 0xd8, 0xd9, 0x86, 0x23, 0xd0, 0x97, 0x45, 0x15,
	0xe3, 0xec, 0x75, 0xc1, 0x95, 0x32, 0x82, 0x2b, 0x35, 0xae, 0x5c, 0x03, 0x0e, 0x9a, 0xd0, 0xe7,
	0x89, 0xb8, 0x9b, 0x76, 0x7f, 0x9c, 0x02, 0x00, 0x00, 0xff, 0xff, 0xeb, 0xa0, 0xd9, 0xa6, 0x28,
	0x01, 0x00, 0x00,
}