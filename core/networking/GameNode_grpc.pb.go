// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.1
// source: GameNode.proto

// Version 4

package networking

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	LobbySearchService_RequestActiveLobbies_FullMethodName = "/networking.LobbySearchService/RequestActiveLobbies"
	LobbySearchService_RequestLobbyJoin_FullMethodName     = "/networking.LobbySearchService/RequestLobbyJoin"
	LobbySearchService_RegisterOwner_FullMethodName        = "/networking.LobbySearchService/RegisterOwner"
)

// LobbySearchServiceClient is the client API for LobbySearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LobbySearchServiceClient interface {
	RequestActiveLobbies(ctx context.Context, in *LobbyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[LobbyProposal], error)
	RequestLobbyJoin(ctx context.Context, in *JoinLobbyRequest, opts ...grpc.CallOption) (*JoinLobbyResponse, error)
	RegisterOwner(ctx context.Context, in *RegisterOwnerRequest, opts ...grpc.CallOption) (*RegisterOwnerResponse, error)
}

type lobbySearchServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLobbySearchServiceClient(cc grpc.ClientConnInterface) LobbySearchServiceClient {
	return &lobbySearchServiceClient{cc}
}

func (c *lobbySearchServiceClient) RequestActiveLobbies(ctx context.Context, in *LobbyRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[LobbyProposal], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &LobbySearchService_ServiceDesc.Streams[0], LobbySearchService_RequestActiveLobbies_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LobbyRequest, LobbyProposal]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LobbySearchService_RequestActiveLobbiesClient = grpc.ServerStreamingClient[LobbyProposal]

func (c *lobbySearchServiceClient) RequestLobbyJoin(ctx context.Context, in *JoinLobbyRequest, opts ...grpc.CallOption) (*JoinLobbyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinLobbyResponse)
	err := c.cc.Invoke(ctx, LobbySearchService_RequestLobbyJoin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lobbySearchServiceClient) RegisterOwner(ctx context.Context, in *RegisterOwnerRequest, opts ...grpc.CallOption) (*RegisterOwnerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterOwnerResponse)
	err := c.cc.Invoke(ctx, LobbySearchService_RegisterOwner_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LobbySearchServiceServer is the server API for LobbySearchService service.
// All implementations must embed UnimplementedLobbySearchServiceServer
// for forward compatibility.
type LobbySearchServiceServer interface {
	RequestActiveLobbies(*LobbyRequest, grpc.ServerStreamingServer[LobbyProposal]) error
	RequestLobbyJoin(context.Context, *JoinLobbyRequest) (*JoinLobbyResponse, error)
	RegisterOwner(context.Context, *RegisterOwnerRequest) (*RegisterOwnerResponse, error)
	mustEmbedUnimplementedLobbySearchServiceServer()
}

// UnimplementedLobbySearchServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLobbySearchServiceServer struct{}

func (UnimplementedLobbySearchServiceServer) RequestActiveLobbies(*LobbyRequest, grpc.ServerStreamingServer[LobbyProposal]) error {
	return status.Errorf(codes.Unimplemented, "method RequestActiveLobbies not implemented")
}
func (UnimplementedLobbySearchServiceServer) RequestLobbyJoin(context.Context, *JoinLobbyRequest) (*JoinLobbyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLobbyJoin not implemented")
}
func (UnimplementedLobbySearchServiceServer) RegisterOwner(context.Context, *RegisterOwnerRequest) (*RegisterOwnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterOwner not implemented")
}
func (UnimplementedLobbySearchServiceServer) mustEmbedUnimplementedLobbySearchServiceServer() {}
func (UnimplementedLobbySearchServiceServer) testEmbeddedByValue()                            {}

// UnsafeLobbySearchServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LobbySearchServiceServer will
// result in compilation errors.
type UnsafeLobbySearchServiceServer interface {
	mustEmbedUnimplementedLobbySearchServiceServer()
}

func RegisterLobbySearchServiceServer(s grpc.ServiceRegistrar, srv LobbySearchServiceServer) {
	// If the following call pancis, it indicates UnimplementedLobbySearchServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LobbySearchService_ServiceDesc, srv)
}

func _LobbySearchService_RequestActiveLobbies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LobbyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LobbySearchServiceServer).RequestActiveLobbies(m, &grpc.GenericServerStream[LobbyRequest, LobbyProposal]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type LobbySearchService_RequestActiveLobbiesServer = grpc.ServerStreamingServer[LobbyProposal]

func _LobbySearchService_RequestLobbyJoin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinLobbyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LobbySearchServiceServer).RequestLobbyJoin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LobbySearchService_RequestLobbyJoin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LobbySearchServiceServer).RequestLobbyJoin(ctx, req.(*JoinLobbyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LobbySearchService_RegisterOwner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterOwnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LobbySearchServiceServer).RegisterOwner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LobbySearchService_RegisterOwner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LobbySearchServiceServer).RegisterOwner(ctx, req.(*RegisterOwnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LobbySearchService_ServiceDesc is the grpc.ServiceDesc for LobbySearchService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LobbySearchService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "networking.LobbySearchService",
	HandlerType: (*LobbySearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestLobbyJoin",
			Handler:    _LobbySearchService_RequestLobbyJoin_Handler,
		},
		{
			MethodName: "RegisterOwner",
			Handler:    _LobbySearchService_RegisterOwner_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestActiveLobbies",
			Handler:       _LobbySearchService_RequestActiveLobbies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "GameNode.proto",
}
