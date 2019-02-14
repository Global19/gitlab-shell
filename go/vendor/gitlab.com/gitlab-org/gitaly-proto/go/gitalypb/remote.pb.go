// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote.proto

package gitalypb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type AddRemoteRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Name       string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Url        string      `protobuf:"bytes,3,opt,name=url" json:"url,omitempty"`
	// If any, the remote is configured as a mirror with those mappings
	MirrorRefmaps []string `protobuf:"bytes,5,rep,name=mirror_refmaps,json=mirrorRefmaps" json:"mirror_refmaps,omitempty"`
}

func (m *AddRemoteRequest) Reset()                    { *m = AddRemoteRequest{} }
func (m *AddRemoteRequest) String() string            { return proto.CompactTextString(m) }
func (*AddRemoteRequest) ProtoMessage()               {}
func (*AddRemoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *AddRemoteRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *AddRemoteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AddRemoteRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *AddRemoteRequest) GetMirrorRefmaps() []string {
	if m != nil {
		return m.MirrorRefmaps
	}
	return nil
}

type AddRemoteResponse struct {
}

func (m *AddRemoteResponse) Reset()                    { *m = AddRemoteResponse{} }
func (m *AddRemoteResponse) String() string            { return proto.CompactTextString(m) }
func (*AddRemoteResponse) ProtoMessage()               {}
func (*AddRemoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

type RemoveRemoteRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Name       string      `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *RemoveRemoteRequest) Reset()                    { *m = RemoveRemoteRequest{} }
func (m *RemoveRemoteRequest) String() string            { return proto.CompactTextString(m) }
func (*RemoveRemoteRequest) ProtoMessage()               {}
func (*RemoveRemoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{2} }

func (m *RemoveRemoteRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *RemoveRemoteRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RemoveRemoteResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *RemoveRemoteResponse) Reset()                    { *m = RemoveRemoteResponse{} }
func (m *RemoveRemoteResponse) String() string            { return proto.CompactTextString(m) }
func (*RemoveRemoteResponse) ProtoMessage()               {}
func (*RemoveRemoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{3} }

func (m *RemoveRemoteResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type FetchInternalRemoteRequest struct {
	Repository       *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	RemoteRepository *Repository `protobuf:"bytes,2,opt,name=remote_repository,json=remoteRepository" json:"remote_repository,omitempty"`
}

func (m *FetchInternalRemoteRequest) Reset()                    { *m = FetchInternalRemoteRequest{} }
func (m *FetchInternalRemoteRequest) String() string            { return proto.CompactTextString(m) }
func (*FetchInternalRemoteRequest) ProtoMessage()               {}
func (*FetchInternalRemoteRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{4} }

func (m *FetchInternalRemoteRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *FetchInternalRemoteRequest) GetRemoteRepository() *Repository {
	if m != nil {
		return m.RemoteRepository
	}
	return nil
}

type FetchInternalRemoteResponse struct {
	Result bool `protobuf:"varint,1,opt,name=result" json:"result,omitempty"`
}

func (m *FetchInternalRemoteResponse) Reset()                    { *m = FetchInternalRemoteResponse{} }
func (m *FetchInternalRemoteResponse) String() string            { return proto.CompactTextString(m) }
func (*FetchInternalRemoteResponse) ProtoMessage()               {}
func (*FetchInternalRemoteResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{5} }

func (m *FetchInternalRemoteResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

type UpdateRemoteMirrorRequest struct {
	Repository           *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	RefName              string      `protobuf:"bytes,2,opt,name=ref_name,json=refName" json:"ref_name,omitempty"`
	OnlyBranchesMatching [][]byte    `protobuf:"bytes,3,rep,name=only_branches_matching,json=onlyBranchesMatching,proto3" json:"only_branches_matching,omitempty"`
	SshKey               string      `protobuf:"bytes,4,opt,name=ssh_key,json=sshKey" json:"ssh_key,omitempty"`
	KnownHosts           string      `protobuf:"bytes,5,opt,name=known_hosts,json=knownHosts" json:"known_hosts,omitempty"`
}

func (m *UpdateRemoteMirrorRequest) Reset()                    { *m = UpdateRemoteMirrorRequest{} }
func (m *UpdateRemoteMirrorRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRemoteMirrorRequest) ProtoMessage()               {}
func (*UpdateRemoteMirrorRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{6} }

func (m *UpdateRemoteMirrorRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *UpdateRemoteMirrorRequest) GetRefName() string {
	if m != nil {
		return m.RefName
	}
	return ""
}

func (m *UpdateRemoteMirrorRequest) GetOnlyBranchesMatching() [][]byte {
	if m != nil {
		return m.OnlyBranchesMatching
	}
	return nil
}

func (m *UpdateRemoteMirrorRequest) GetSshKey() string {
	if m != nil {
		return m.SshKey
	}
	return ""
}

func (m *UpdateRemoteMirrorRequest) GetKnownHosts() string {
	if m != nil {
		return m.KnownHosts
	}
	return ""
}

type UpdateRemoteMirrorResponse struct {
}

func (m *UpdateRemoteMirrorResponse) Reset()                    { *m = UpdateRemoteMirrorResponse{} }
func (m *UpdateRemoteMirrorResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateRemoteMirrorResponse) ProtoMessage()               {}
func (*UpdateRemoteMirrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{7} }

type FindRemoteRepositoryRequest struct {
	Remote string `protobuf:"bytes,1,opt,name=remote" json:"remote,omitempty"`
}

func (m *FindRemoteRepositoryRequest) Reset()                    { *m = FindRemoteRepositoryRequest{} }
func (m *FindRemoteRepositoryRequest) String() string            { return proto.CompactTextString(m) }
func (*FindRemoteRepositoryRequest) ProtoMessage()               {}
func (*FindRemoteRepositoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{8} }

func (m *FindRemoteRepositoryRequest) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

// This migth throw a GRPC Unavailable code, to signal the request failure
// is transient.
type FindRemoteRepositoryResponse struct {
	Exists bool `protobuf:"varint,1,opt,name=exists" json:"exists,omitempty"`
}

func (m *FindRemoteRepositoryResponse) Reset()                    { *m = FindRemoteRepositoryResponse{} }
func (m *FindRemoteRepositoryResponse) String() string            { return proto.CompactTextString(m) }
func (*FindRemoteRepositoryResponse) ProtoMessage()               {}
func (*FindRemoteRepositoryResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{9} }

func (m *FindRemoteRepositoryResponse) GetExists() bool {
	if m != nil {
		return m.Exists
	}
	return false
}

type FindRemoteRootRefRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
	Remote     string      `protobuf:"bytes,2,opt,name=remote" json:"remote,omitempty"`
}

func (m *FindRemoteRootRefRequest) Reset()                    { *m = FindRemoteRootRefRequest{} }
func (m *FindRemoteRootRefRequest) String() string            { return proto.CompactTextString(m) }
func (*FindRemoteRootRefRequest) ProtoMessage()               {}
func (*FindRemoteRootRefRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{10} }

func (m *FindRemoteRootRefRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *FindRemoteRootRefRequest) GetRemote() string {
	if m != nil {
		return m.Remote
	}
	return ""
}

type FindRemoteRootRefResponse struct {
	Ref string `protobuf:"bytes,1,opt,name=ref" json:"ref,omitempty"`
}

func (m *FindRemoteRootRefResponse) Reset()                    { *m = FindRemoteRootRefResponse{} }
func (m *FindRemoteRootRefResponse) String() string            { return proto.CompactTextString(m) }
func (*FindRemoteRootRefResponse) ProtoMessage()               {}
func (*FindRemoteRootRefResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{11} }

func (m *FindRemoteRootRefResponse) GetRef() string {
	if m != nil {
		return m.Ref
	}
	return ""
}

type ListRemotesRequest struct {
	Repository *Repository `protobuf:"bytes,1,opt,name=repository" json:"repository,omitempty"`
}

func (m *ListRemotesRequest) Reset()                    { *m = ListRemotesRequest{} }
func (m *ListRemotesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListRemotesRequest) ProtoMessage()               {}
func (*ListRemotesRequest) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{12} }

func (m *ListRemotesRequest) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

type ListRemotesResponse struct {
	Remotes []*ListRemotesResponse_Remote `protobuf:"bytes,1,rep,name=remotes" json:"remotes,omitempty"`
}

func (m *ListRemotesResponse) Reset()                    { *m = ListRemotesResponse{} }
func (m *ListRemotesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListRemotesResponse) ProtoMessage()               {}
func (*ListRemotesResponse) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13} }

func (m *ListRemotesResponse) GetRemotes() []*ListRemotesResponse_Remote {
	if m != nil {
		return m.Remotes
	}
	return nil
}

type ListRemotesResponse_Remote struct {
	Name     string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	FetchUrl string `protobuf:"bytes,2,opt,name=fetch_url,json=fetchUrl" json:"fetch_url,omitempty"`
	PushUrl  string `protobuf:"bytes,3,opt,name=push_url,json=pushUrl" json:"push_url,omitempty"`
}

func (m *ListRemotesResponse_Remote) Reset()                    { *m = ListRemotesResponse_Remote{} }
func (m *ListRemotesResponse_Remote) String() string            { return proto.CompactTextString(m) }
func (*ListRemotesResponse_Remote) ProtoMessage()               {}
func (*ListRemotesResponse_Remote) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{13, 0} }

func (m *ListRemotesResponse_Remote) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ListRemotesResponse_Remote) GetFetchUrl() string {
	if m != nil {
		return m.FetchUrl
	}
	return ""
}

func (m *ListRemotesResponse_Remote) GetPushUrl() string {
	if m != nil {
		return m.PushUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*AddRemoteRequest)(nil), "gitaly.AddRemoteRequest")
	proto.RegisterType((*AddRemoteResponse)(nil), "gitaly.AddRemoteResponse")
	proto.RegisterType((*RemoveRemoteRequest)(nil), "gitaly.RemoveRemoteRequest")
	proto.RegisterType((*RemoveRemoteResponse)(nil), "gitaly.RemoveRemoteResponse")
	proto.RegisterType((*FetchInternalRemoteRequest)(nil), "gitaly.FetchInternalRemoteRequest")
	proto.RegisterType((*FetchInternalRemoteResponse)(nil), "gitaly.FetchInternalRemoteResponse")
	proto.RegisterType((*UpdateRemoteMirrorRequest)(nil), "gitaly.UpdateRemoteMirrorRequest")
	proto.RegisterType((*UpdateRemoteMirrorResponse)(nil), "gitaly.UpdateRemoteMirrorResponse")
	proto.RegisterType((*FindRemoteRepositoryRequest)(nil), "gitaly.FindRemoteRepositoryRequest")
	proto.RegisterType((*FindRemoteRepositoryResponse)(nil), "gitaly.FindRemoteRepositoryResponse")
	proto.RegisterType((*FindRemoteRootRefRequest)(nil), "gitaly.FindRemoteRootRefRequest")
	proto.RegisterType((*FindRemoteRootRefResponse)(nil), "gitaly.FindRemoteRootRefResponse")
	proto.RegisterType((*ListRemotesRequest)(nil), "gitaly.ListRemotesRequest")
	proto.RegisterType((*ListRemotesResponse)(nil), "gitaly.ListRemotesResponse")
	proto.RegisterType((*ListRemotesResponse_Remote)(nil), "gitaly.ListRemotesResponse.Remote")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RemoteService service

type RemoteServiceClient interface {
	AddRemote(ctx context.Context, in *AddRemoteRequest, opts ...grpc.CallOption) (*AddRemoteResponse, error)
	FetchInternalRemote(ctx context.Context, in *FetchInternalRemoteRequest, opts ...grpc.CallOption) (*FetchInternalRemoteResponse, error)
	RemoveRemote(ctx context.Context, in *RemoveRemoteRequest, opts ...grpc.CallOption) (*RemoveRemoteResponse, error)
	UpdateRemoteMirror(ctx context.Context, opts ...grpc.CallOption) (RemoteService_UpdateRemoteMirrorClient, error)
	FindRemoteRepository(ctx context.Context, in *FindRemoteRepositoryRequest, opts ...grpc.CallOption) (*FindRemoteRepositoryResponse, error)
	FindRemoteRootRef(ctx context.Context, in *FindRemoteRootRefRequest, opts ...grpc.CallOption) (*FindRemoteRootRefResponse, error)
	ListRemotes(ctx context.Context, in *ListRemotesRequest, opts ...grpc.CallOption) (RemoteService_ListRemotesClient, error)
}

type remoteServiceClient struct {
	cc *grpc.ClientConn
}

func NewRemoteServiceClient(cc *grpc.ClientConn) RemoteServiceClient {
	return &remoteServiceClient{cc}
}

func (c *remoteServiceClient) AddRemote(ctx context.Context, in *AddRemoteRequest, opts ...grpc.CallOption) (*AddRemoteResponse, error) {
	out := new(AddRemoteResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/AddRemote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) FetchInternalRemote(ctx context.Context, in *FetchInternalRemoteRequest, opts ...grpc.CallOption) (*FetchInternalRemoteResponse, error) {
	out := new(FetchInternalRemoteResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/FetchInternalRemote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) RemoveRemote(ctx context.Context, in *RemoveRemoteRequest, opts ...grpc.CallOption) (*RemoveRemoteResponse, error) {
	out := new(RemoveRemoteResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/RemoveRemote", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) UpdateRemoteMirror(ctx context.Context, opts ...grpc.CallOption) (RemoteService_UpdateRemoteMirrorClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RemoteService_serviceDesc.Streams[0], c.cc, "/gitaly.RemoteService/UpdateRemoteMirror", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteServiceUpdateRemoteMirrorClient{stream}
	return x, nil
}

type RemoteService_UpdateRemoteMirrorClient interface {
	Send(*UpdateRemoteMirrorRequest) error
	CloseAndRecv() (*UpdateRemoteMirrorResponse, error)
	grpc.ClientStream
}

type remoteServiceUpdateRemoteMirrorClient struct {
	grpc.ClientStream
}

func (x *remoteServiceUpdateRemoteMirrorClient) Send(m *UpdateRemoteMirrorRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *remoteServiceUpdateRemoteMirrorClient) CloseAndRecv() (*UpdateRemoteMirrorResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateRemoteMirrorResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *remoteServiceClient) FindRemoteRepository(ctx context.Context, in *FindRemoteRepositoryRequest, opts ...grpc.CallOption) (*FindRemoteRepositoryResponse, error) {
	out := new(FindRemoteRepositoryResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/FindRemoteRepository", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) FindRemoteRootRef(ctx context.Context, in *FindRemoteRootRefRequest, opts ...grpc.CallOption) (*FindRemoteRootRefResponse, error) {
	out := new(FindRemoteRootRefResponse)
	err := grpc.Invoke(ctx, "/gitaly.RemoteService/FindRemoteRootRef", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteServiceClient) ListRemotes(ctx context.Context, in *ListRemotesRequest, opts ...grpc.CallOption) (RemoteService_ListRemotesClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RemoteService_serviceDesc.Streams[1], c.cc, "/gitaly.RemoteService/ListRemotes", opts...)
	if err != nil {
		return nil, err
	}
	x := &remoteServiceListRemotesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RemoteService_ListRemotesClient interface {
	Recv() (*ListRemotesResponse, error)
	grpc.ClientStream
}

type remoteServiceListRemotesClient struct {
	grpc.ClientStream
}

func (x *remoteServiceListRemotesClient) Recv() (*ListRemotesResponse, error) {
	m := new(ListRemotesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RemoteService service

type RemoteServiceServer interface {
	AddRemote(context.Context, *AddRemoteRequest) (*AddRemoteResponse, error)
	FetchInternalRemote(context.Context, *FetchInternalRemoteRequest) (*FetchInternalRemoteResponse, error)
	RemoveRemote(context.Context, *RemoveRemoteRequest) (*RemoveRemoteResponse, error)
	UpdateRemoteMirror(RemoteService_UpdateRemoteMirrorServer) error
	FindRemoteRepository(context.Context, *FindRemoteRepositoryRequest) (*FindRemoteRepositoryResponse, error)
	FindRemoteRootRef(context.Context, *FindRemoteRootRefRequest) (*FindRemoteRootRefResponse, error)
	ListRemotes(*ListRemotesRequest, RemoteService_ListRemotesServer) error
}

func RegisterRemoteServiceServer(s *grpc.Server, srv RemoteServiceServer) {
	s.RegisterService(&_RemoteService_serviceDesc, srv)
}

func _RemoteService_AddRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).AddRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/AddRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).AddRemote(ctx, req.(*AddRemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_FetchInternalRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchInternalRemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).FetchInternalRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/FetchInternalRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).FetchInternalRemote(ctx, req.(*FetchInternalRemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_RemoveRemote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRemoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).RemoveRemote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/RemoveRemote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).RemoveRemote(ctx, req.(*RemoveRemoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_UpdateRemoteMirror_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RemoteServiceServer).UpdateRemoteMirror(&remoteServiceUpdateRemoteMirrorServer{stream})
}

type RemoteService_UpdateRemoteMirrorServer interface {
	SendAndClose(*UpdateRemoteMirrorResponse) error
	Recv() (*UpdateRemoteMirrorRequest, error)
	grpc.ServerStream
}

type remoteServiceUpdateRemoteMirrorServer struct {
	grpc.ServerStream
}

func (x *remoteServiceUpdateRemoteMirrorServer) SendAndClose(m *UpdateRemoteMirrorResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *remoteServiceUpdateRemoteMirrorServer) Recv() (*UpdateRemoteMirrorRequest, error) {
	m := new(UpdateRemoteMirrorRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RemoteService_FindRemoteRepository_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRemoteRepositoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).FindRemoteRepository(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/FindRemoteRepository",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).FindRemoteRepository(ctx, req.(*FindRemoteRepositoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_FindRemoteRootRef_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindRemoteRootRefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteServiceServer).FindRemoteRootRef(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gitaly.RemoteService/FindRemoteRootRef",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteServiceServer).FindRemoteRootRef(ctx, req.(*FindRemoteRootRefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteService_ListRemotes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListRemotesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RemoteServiceServer).ListRemotes(m, &remoteServiceListRemotesServer{stream})
}

type RemoteService_ListRemotesServer interface {
	Send(*ListRemotesResponse) error
	grpc.ServerStream
}

type remoteServiceListRemotesServer struct {
	grpc.ServerStream
}

func (x *remoteServiceListRemotesServer) Send(m *ListRemotesResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _RemoteService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gitaly.RemoteService",
	HandlerType: (*RemoteServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddRemote",
			Handler:    _RemoteService_AddRemote_Handler,
		},
		{
			MethodName: "FetchInternalRemote",
			Handler:    _RemoteService_FetchInternalRemote_Handler,
		},
		{
			MethodName: "RemoveRemote",
			Handler:    _RemoteService_RemoveRemote_Handler,
		},
		{
			MethodName: "FindRemoteRepository",
			Handler:    _RemoteService_FindRemoteRepository_Handler,
		},
		{
			MethodName: "FindRemoteRootRef",
			Handler:    _RemoteService_FindRemoteRootRef_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateRemoteMirror",
			Handler:       _RemoteService_UpdateRemoteMirror_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ListRemotes",
			Handler:       _RemoteService_ListRemotes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "remote.proto",
}

func init() { proto.RegisterFile("remote.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 672 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcb, 0x6e, 0xd3, 0x4c,
	0x14, 0xae, 0xeb, 0x34, 0x97, 0x93, 0xf4, 0x57, 0x3a, 0xa9, 0xfa, 0x3b, 0x4e, 0x25, 0xd2, 0x01,
	0xa4, 0x6c, 0x88, 0x50, 0xb8, 0xac, 0x90, 0x10, 0x5d, 0xa0, 0xd2, 0x52, 0x16, 0x86, 0x6e, 0x90,
	0x90, 0x71, 0x93, 0x71, 0x6d, 0xd5, 0xf1, 0x98, 0x99, 0x49, 0x21, 0x8f, 0xc1, 0x5b, 0xc0, 0x2b,
	0xf1, 0x14, 0x3c, 0x02, 0x1a, 0xcf, 0xd8, 0x71, 0xa8, 0x13, 0x24, 0x22, 0x76, 0x9e, 0x73, 0x9b,
	0xef, 0x3b, 0xf3, 0x9d, 0x63, 0x68, 0x31, 0x32, 0xa5, 0x82, 0x0c, 0x13, 0x46, 0x05, 0x45, 0xd5,
	0xab, 0x50, 0x78, 0xd1, 0xdc, 0x6e, 0xf1, 0xc0, 0x63, 0x64, 0xa2, 0xac, 0xf8, 0x9b, 0x01, 0xed,
	0x17, 0x93, 0x89, 0x93, 0x46, 0x3a, 0xe4, 0xd3, 0x8c, 0x70, 0x81, 0x46, 0x00, 0x8c, 0x24, 0x94,
	0x87, 0x82, 0xb2, 0xb9, 0x65, 0xf4, 0x8d, 0x41, 0x73, 0x84, 0x86, 0x2a, 0x7f, 0xe8, 0xe4, 0x1e,
	0xa7, 0x10, 0x85, 0x10, 0x54, 0x62, 0x6f, 0x4a, 0xac, 0xed, 0xbe, 0x31, 0x68, 0x38, 0xe9, 0x37,
	0x6a, 0x83, 0x39, 0x63, 0x91, 0x65, 0xa6, 0x26, 0xf9, 0x89, 0xee, 0xc3, 0x7f, 0xd3, 0x90, 0x31,
	0xca, 0x5c, 0x46, 0xfc, 0xa9, 0x97, 0x70, 0x6b, 0xa7, 0x6f, 0x0e, 0x1a, 0xce, 0xae, 0xb2, 0x3a,
	0xca, 0x78, 0x5a, 0xa9, 0x57, 0xda, 0x3b, 0x99, 0x51, 0x87, 0xe2, 0x0e, 0xec, 0x15, 0x90, 0xf2,
	0x84, 0xc6, 0x9c, 0xe0, 0x0f, 0xd0, 0x91, 0x96, 0x1b, 0xf2, 0x4f, 0x18, 0xe0, 0x21, 0xec, 0x2f,
	0x97, 0x57, 0xd7, 0xa2, 0x03, 0xa8, 0x32, 0xc2, 0x67, 0x91, 0x48, 0x6b, 0xd7, 0x1d, 0x7d, 0xc2,
	0x5f, 0x0d, 0xb0, 0x5f, 0x12, 0x31, 0x0e, 0x5e, 0xc5, 0x82, 0xb0, 0xd8, 0x8b, 0x36, 0x87, 0xf5,
	0x1c, 0xf6, 0xd4, 0x3b, 0xba, 0x85, 0xd4, 0xed, 0x95, 0xa9, 0x6d, 0xa6, 0x6f, 0xcc, 0x2c, 0xf8,
	0x09, 0xf4, 0x4a, 0x21, 0xfd, 0x81, 0xca, 0x0f, 0x03, 0xba, 0x17, 0xc9, 0xc4, 0x13, 0x9a, 0xfb,
	0xb9, 0x7e, 0xa1, 0xbf, 0x67, 0xd2, 0x85, 0x3a, 0x23, 0xbe, 0x5b, 0x68, 0x72, 0x8d, 0x11, 0xff,
	0x8d, 0x54, 0xca, 0x63, 0x38, 0xa0, 0x71, 0x34, 0x77, 0x2f, 0x99, 0x17, 0x8f, 0x03, 0xc2, 0xdd,
	0xa9, 0x27, 0xc6, 0x41, 0x18, 0x5f, 0x59, 0x66, 0xdf, 0x1c, 0xb4, 0x9c, 0x7d, 0xe9, 0x3d, 0xd6,
	0xce, 0x73, 0xed, 0x43, 0xff, 0x43, 0x8d, 0xf3, 0xc0, 0xbd, 0x26, 0x73, 0xab, 0x92, 0xd6, 0xab,
	0x72, 0x1e, 0x9c, 0x91, 0x39, 0xba, 0x03, 0xcd, 0xeb, 0x98, 0x7e, 0x8e, 0xdd, 0x80, 0x72, 0x21,
	0x35, 0x26, 0x9d, 0x90, 0x9a, 0x4e, 0xa4, 0x05, 0x1f, 0x82, 0x5d, 0xc6, 0x4d, 0x8b, 0x4a, 0x76,
	0x2c, 0x8c, 0x73, 0xa9, 0xe5, 0x64, 0x34, 0xf7, 0xb4, 0x63, 0xd2, 0x95, 0xf2, 0x6e, 0x38, 0xfa,
	0x84, 0x9f, 0xc2, 0x61, 0x79, 0xda, 0xa2, 0xd3, 0xe4, 0x4b, 0x28, 0x01, 0xe9, 0x4e, 0xab, 0x13,
	0xf6, 0xc1, 0x2a, 0xe4, 0x51, 0x2a, 0x1c, 0xe2, 0x6f, 0xd2, 0xe7, 0x05, 0xbe, 0xed, 0x25, 0x7c,
	0x0f, 0xa0, 0x5b, 0x72, 0x8f, 0x06, 0xd7, 0x06, 0x93, 0x11, 0x5f, 0x33, 0x92, 0x9f, 0xf8, 0x04,
	0xd0, 0xeb, 0x90, 0x0b, 0x15, 0xce, 0x37, 0x00, 0x84, 0xbf, 0x1b, 0xd0, 0x59, 0x2a, 0xa5, 0xef,
	0x7c, 0x06, 0x35, 0x05, 0x4d, 0x76, 0xc4, 0x1c, 0x34, 0x47, 0x38, 0x2b, 0x54, 0x12, 0x3d, 0xd4,
	0xb8, 0xb3, 0x14, 0xfb, 0x1d, 0x54, 0x95, 0x29, 0x9f, 0x5c, 0xa3, 0xb0, 0x7b, 0x7a, 0xd0, 0xf0,
	0xa5, 0xea, 0x5d, 0xb9, 0x81, 0x54, 0x1f, 0xea, 0xa9, 0xe1, 0x82, 0x45, 0x52, 0x89, 0xc9, 0x8c,
	0x2b, 0x9f, 0xda, 0x4e, 0x35, 0x79, 0xbe, 0x60, 0xd1, 0xe8, 0x67, 0x05, 0x76, 0x55, 0xd9, 0xb7,
	0x84, 0xdd, 0x84, 0x63, 0x82, 0x8e, 0xa1, 0x91, 0xef, 0x1d, 0x64, 0x65, 0x08, 0x7f, 0x5f, 0x9a,
	0x76, 0xb7, 0xc4, 0xa3, 0xf5, 0xb4, 0x85, 0x3e, 0x42, 0xa7, 0x64, 0x06, 0x51, 0xce, 0x77, 0xf5,
	0xce, 0xb0, 0xef, 0xae, 0x8d, 0xc9, 0x6f, 0x38, 0x83, 0x56, 0x71, 0x53, 0xa1, 0xde, 0xe2, 0x4d,
	0x6e, 0xad, 0x47, 0xfb, 0xb0, 0xdc, 0x99, 0x17, 0x73, 0x01, 0xdd, 0x1e, 0x0f, 0x74, 0x94, 0x65,
	0xad, 0x5c, 0x0b, 0x36, 0x5e, 0x17, 0x92, 0x95, 0x1f, 0x18, 0x68, 0x0c, 0xfb, 0x65, 0xa3, 0x82,
	0x16, 0x64, 0x57, 0xcf, 0x9f, 0x7d, 0x6f, 0x7d, 0x50, 0xce, 0xe2, 0x3d, 0xec, 0xdd, 0xd2, 0x3b,
	0xea, 0x97, 0x24, 0x2f, 0x8d, 0x9c, 0x7d, 0xb4, 0x26, 0x22, 0xaf, 0x7d, 0x0a, 0xcd, 0x82, 0x46,
	0x91, 0x5d, 0x2a, 0x5c, 0x55, 0xaf, 0xb7, 0x46, 0xd4, 0x78, 0xeb, 0xa1, 0x71, 0x59, 0x4d, 0x7f,
	0xc5, 0x8f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x4b, 0x62, 0xf6, 0x64, 0xb0, 0x07, 0x00, 0x00,
}
