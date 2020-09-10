// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/nodeRegistration.proto

package service

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	model "github.com/zoobc/zoobc-core/common/model"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

func init() {
	proto.RegisterFile("service/nodeRegistration.proto", fileDescriptor_d783036dc420a87c)
}

var fileDescriptor_d783036dc420a87c = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x8d, 0x87, 0x0a, 0x8b, 0x17, 0x47, 0x41, 0x88, 0x45, 0xb0, 0x0a, 0xb6, 0x55, 0xb3,
	0xfd, 0x73, 0x12, 0x4f, 0x16, 0xa4, 0x07, 0x51, 0x4a, 0xbd, 0x79, 0x6b, 0xd2, 0x21, 0x2e, 0x34,
	0x3b, 0x31, 0xbb, 0x2d, 0xb4, 0x47, 0x4f, 0x9e, 0x3c, 0xe8, 0x9b, 0xe9, 0x2b, 0xf8, 0x20, 0xd2,
	0xcd, 0xd6, 0x3f, 0x69, 0x5a, 0xea, 0x25, 0x87, 0xfd, 0xbe, 0xd9, 0xf9, 0xcd, 0xb7, 0x19, 0xb6,
	0xaf, 0x30, 0x19, 0x89, 0x00, 0xb9, 0xa4, 0x3e, 0x76, 0x31, 0x14, 0x4a, 0x27, 0x3d, 0x2d, 0x48,
	0x7a, 0x71, 0x42, 0x9a, 0x60, 0xc3, 0xea, 0xee, 0x56, 0x44, 0x7d, 0x1c, 0x70, 0x8c, 0x62, 0x3d,
	0x4e, 0x35, 0xb7, 0x98, 0x1e, 0xe5, 0x57, 0xba, 0xc5, 0x90, 0x28, 0x1c, 0x20, 0xef, 0xc5, 0x82,
	0xf7, 0xa4, 0x24, 0x6d, 0x44, 0x95, 0xaa, 0x8d, 0xe7, 0x02, 0xdb, 0xbd, 0xcd, 0x14, 0xde, 0xa5,
	0xad, 0xe0, 0xd5, 0x61, 0x3b, 0x6d, 0xd4, 0x59, 0x59, 0x41, 0xc9, 0x33, 0x1d, 0xbd, 0x3c, 0xb1,
	0x8b, 0x8f, 0x43, 0x54, 0xda, 0x3d, 0x5c, 0xea, 0x51, 0x31, 0x49, 0x85, 0xa5, 0xfa, 0xd3, 0xc7,
	0xe7, 0xdb, 0xfa, 0x09, 0x54, 0xf8, 0xa8, 0x3e, 0xc7, 0xcf, 0x73, 0x7b, 0xbf, 0x38, 0x6c, 0x3b,
	0x47, 0x80, 0x83, 0xc5, 0xfd, 0x66, 0x48, 0xa5, 0x65, 0x16, 0x4b, 0x54, 0x33, 0x44, 0x55, 0x28,
	0xaf, 0x4a, 0x04, 0xef, 0x0e, 0x3b, 0xca, 0x23, 0x6d, 0x8d, 0xa7, 0x47, 0x9d, 0xa1, 0x3f, 0x10,
	0xc1, 0x35, 0x8e, 0x15, 0x34, 0x96, 0x24, 0x92, 0x35, 0xcf, 0x90, 0x9b, 0xff, 0xaa, 0xb1, 0x33,
	0x5c, 0x9a, 0x19, 0x2e, 0xe0, 0x7c, 0xe5, 0x54, 0xe7, 0x58, 0x27, 0x6c, 0xaf, 0x8d, 0xba, 0x83,
	0xb2, 0x2f, 0x64, 0x38, 0xff, 0x08, 0x95, 0x1f, 0xac, 0x45, 0x9e, 0xd9, 0x04, 0xd5, 0x55, 0xac,
	0x16, 0x7c, 0xad, 0xec, 0xd4, 0x1c, 0x18, 0x31, 0x68, 0xa3, 0xbe, 0xf9, 0x8b, 0x04, 0x9b, 0xf6,
	0x9e, 0xab, 0xe9, 0x8f, 0xef, 0xfe, 0x7a, 0xed, 0x8c, 0xf1, 0xfb, 0x32, 0x6e, 0x52, 0xa8, 0xc0,
	0xf1, 0xa2, 0x14, 0x32, 0x85, 0xad, 0xd3, 0xfb, 0x6a, 0x28, 0xf4, 0xc3, 0xd0, 0xf7, 0x02, 0x8a,
	0xf8, 0x84, 0xc8, 0x0f, 0xd2, 0xef, 0x59, 0x40, 0x09, 0xf2, 0x80, 0xa2, 0x88, 0x24, 0xb7, 0x7b,
	0xe8, 0x17, 0xcc, 0xfe, 0x34, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x32, 0xf8, 0x2b, 0x47, 0xb9,
	0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeRegistrationServiceClient is the client API for NodeRegistrationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeRegistrationServiceClient interface {
	GetNodeRegistrations(ctx context.Context, in *model.GetNodeRegistrationsRequest, opts ...grpc.CallOption) (*model.GetNodeRegistrationsResponse, error)
	GetNodeRegistration(ctx context.Context, in *model.GetNodeRegistrationRequest, opts ...grpc.CallOption) (*model.GetNodeRegistrationResponse, error)
	GetNodeRegistrationsByNodePublicKeys(ctx context.Context, in *model.GetNodeRegistrationsByNodePublicKeysRequest, opts ...grpc.CallOption) (*model.GetNodeRegistrationsByNodePublicKeysResponse, error)
	GetPendingNodeRegistrations(ctx context.Context, opts ...grpc.CallOption) (NodeRegistrationService_GetPendingNodeRegistrationsClient, error)
	// GetMyNodePublicKey return the public key of the node
	GetMyNodePublicKey(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*model.GetMyNodePublicKeyResponse, error)
}

type nodeRegistrationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeRegistrationServiceClient(cc grpc.ClientConnInterface) NodeRegistrationServiceClient {
	return &nodeRegistrationServiceClient{cc}
}

func (c *nodeRegistrationServiceClient) GetNodeRegistrations(ctx context.Context, in *model.GetNodeRegistrationsRequest, opts ...grpc.CallOption) (*model.GetNodeRegistrationsResponse, error) {
	out := new(model.GetNodeRegistrationsResponse)
	err := c.cc.Invoke(ctx, "/service.NodeRegistrationService/GetNodeRegistrations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeRegistrationServiceClient) GetNodeRegistration(ctx context.Context, in *model.GetNodeRegistrationRequest, opts ...grpc.CallOption) (*model.GetNodeRegistrationResponse, error) {
	out := new(model.GetNodeRegistrationResponse)
	err := c.cc.Invoke(ctx, "/service.NodeRegistrationService/GetNodeRegistration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeRegistrationServiceClient) GetNodeRegistrationsByNodePublicKeys(ctx context.Context, in *model.GetNodeRegistrationsByNodePublicKeysRequest, opts ...grpc.CallOption) (*model.GetNodeRegistrationsByNodePublicKeysResponse, error) {
	out := new(model.GetNodeRegistrationsByNodePublicKeysResponse)
	err := c.cc.Invoke(ctx, "/service.NodeRegistrationService/GetNodeRegistrationsByNodePublicKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeRegistrationServiceClient) GetPendingNodeRegistrations(ctx context.Context, opts ...grpc.CallOption) (NodeRegistrationService_GetPendingNodeRegistrationsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeRegistrationService_serviceDesc.Streams[0], "/service.NodeRegistrationService/GetPendingNodeRegistrations", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeRegistrationServiceGetPendingNodeRegistrationsClient{stream}
	return x, nil
}

type NodeRegistrationService_GetPendingNodeRegistrationsClient interface {
	Send(*model.GetPendingNodeRegistrationsRequest) error
	Recv() (*model.GetPendingNodeRegistrationsResponse, error)
	grpc.ClientStream
}

type nodeRegistrationServiceGetPendingNodeRegistrationsClient struct {
	grpc.ClientStream
}

func (x *nodeRegistrationServiceGetPendingNodeRegistrationsClient) Send(m *model.GetPendingNodeRegistrationsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeRegistrationServiceGetPendingNodeRegistrationsClient) Recv() (*model.GetPendingNodeRegistrationsResponse, error) {
	m := new(model.GetPendingNodeRegistrationsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeRegistrationServiceClient) GetMyNodePublicKey(ctx context.Context, in *model.Empty, opts ...grpc.CallOption) (*model.GetMyNodePublicKeyResponse, error) {
	out := new(model.GetMyNodePublicKeyResponse)
	err := c.cc.Invoke(ctx, "/service.NodeRegistrationService/GetMyNodePublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeRegistrationServiceServer is the server API for NodeRegistrationService service.
type NodeRegistrationServiceServer interface {
	GetNodeRegistrations(context.Context, *model.GetNodeRegistrationsRequest) (*model.GetNodeRegistrationsResponse, error)
	GetNodeRegistration(context.Context, *model.GetNodeRegistrationRequest) (*model.GetNodeRegistrationResponse, error)
	GetNodeRegistrationsByNodePublicKeys(context.Context, *model.GetNodeRegistrationsByNodePublicKeysRequest) (*model.GetNodeRegistrationsByNodePublicKeysResponse, error)
	GetPendingNodeRegistrations(NodeRegistrationService_GetPendingNodeRegistrationsServer) error
	// GetMyNodePublicKey return the public key of the node
	GetMyNodePublicKey(context.Context, *model.Empty) (*model.GetMyNodePublicKeyResponse, error)
}

// UnimplementedNodeRegistrationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedNodeRegistrationServiceServer struct {
}

func (*UnimplementedNodeRegistrationServiceServer) GetNodeRegistrations(ctx context.Context, req *model.GetNodeRegistrationsRequest) (*model.GetNodeRegistrationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeRegistrations not implemented")
}
func (*UnimplementedNodeRegistrationServiceServer) GetNodeRegistration(ctx context.Context, req *model.GetNodeRegistrationRequest) (*model.GetNodeRegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeRegistration not implemented")
}
func (*UnimplementedNodeRegistrationServiceServer) GetNodeRegistrationsByNodePublicKeys(ctx context.Context, req *model.GetNodeRegistrationsByNodePublicKeysRequest) (*model.GetNodeRegistrationsByNodePublicKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNodeRegistrationsByNodePublicKeys not implemented")
}
func (*UnimplementedNodeRegistrationServiceServer) GetPendingNodeRegistrations(srv NodeRegistrationService_GetPendingNodeRegistrationsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetPendingNodeRegistrations not implemented")
}
func (*UnimplementedNodeRegistrationServiceServer) GetMyNodePublicKey(ctx context.Context, req *model.Empty) (*model.GetMyNodePublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyNodePublicKey not implemented")
}

func RegisterNodeRegistrationServiceServer(s *grpc.Server, srv NodeRegistrationServiceServer) {
	s.RegisterService(&_NodeRegistrationService_serviceDesc, srv)
}

func _NodeRegistrationService_GetNodeRegistrations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetNodeRegistrationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeRegistrationServiceServer).GetNodeRegistrations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NodeRegistrationService/GetNodeRegistrations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeRegistrationServiceServer).GetNodeRegistrations(ctx, req.(*model.GetNodeRegistrationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeRegistrationService_GetNodeRegistration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetNodeRegistrationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeRegistrationServiceServer).GetNodeRegistration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NodeRegistrationService/GetNodeRegistration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeRegistrationServiceServer).GetNodeRegistration(ctx, req.(*model.GetNodeRegistrationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeRegistrationService_GetNodeRegistrationsByNodePublicKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.GetNodeRegistrationsByNodePublicKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeRegistrationServiceServer).GetNodeRegistrationsByNodePublicKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NodeRegistrationService/GetNodeRegistrationsByNodePublicKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeRegistrationServiceServer).GetNodeRegistrationsByNodePublicKeys(ctx, req.(*model.GetNodeRegistrationsByNodePublicKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeRegistrationService_GetPendingNodeRegistrations_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeRegistrationServiceServer).GetPendingNodeRegistrations(&nodeRegistrationServiceGetPendingNodeRegistrationsServer{stream})
}

type NodeRegistrationService_GetPendingNodeRegistrationsServer interface {
	Send(*model.GetPendingNodeRegistrationsResponse) error
	Recv() (*model.GetPendingNodeRegistrationsRequest, error)
	grpc.ServerStream
}

type nodeRegistrationServiceGetPendingNodeRegistrationsServer struct {
	grpc.ServerStream
}

func (x *nodeRegistrationServiceGetPendingNodeRegistrationsServer) Send(m *model.GetPendingNodeRegistrationsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeRegistrationServiceGetPendingNodeRegistrationsServer) Recv() (*model.GetPendingNodeRegistrationsRequest, error) {
	m := new(model.GetPendingNodeRegistrationsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NodeRegistrationService_GetMyNodePublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(model.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeRegistrationServiceServer).GetMyNodePublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.NodeRegistrationService/GetMyNodePublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeRegistrationServiceServer).GetMyNodePublicKey(ctx, req.(*model.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeRegistrationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.NodeRegistrationService",
	HandlerType: (*NodeRegistrationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNodeRegistrations",
			Handler:    _NodeRegistrationService_GetNodeRegistrations_Handler,
		},
		{
			MethodName: "GetNodeRegistration",
			Handler:    _NodeRegistrationService_GetNodeRegistration_Handler,
		},
		{
			MethodName: "GetNodeRegistrationsByNodePublicKeys",
			Handler:    _NodeRegistrationService_GetNodeRegistrationsByNodePublicKeys_Handler,
		},
		{
			MethodName: "GetMyNodePublicKey",
			Handler:    _NodeRegistrationService_GetMyNodePublicKey_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetPendingNodeRegistrations",
			Handler:       _NodeRegistrationService_GetPendingNodeRegistrations_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "service/nodeRegistration.proto",
}
