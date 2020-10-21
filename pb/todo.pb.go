// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/todo.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type CreateTodoRequest struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoRequest) Reset()         { *m = CreateTodoRequest{} }
func (m *CreateTodoRequest) String() string { return proto.CompactTextString(m) }
func (*CreateTodoRequest) ProtoMessage()    {}
func (*CreateTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f52905b89a1b23f, []int{0}
}

func (m *CreateTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoRequest.Unmarshal(m, b)
}
func (m *CreateTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoRequest.Marshal(b, m, deterministic)
}
func (m *CreateTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoRequest.Merge(m, src)
}
func (m *CreateTodoRequest) XXX_Size() int {
	return xxx_messageInfo_CreateTodoRequest.Size(m)
}
func (m *CreateTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoRequest proto.InternalMessageInfo

func (m *CreateTodoRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

type CreateTodoResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateTodoResponse) Reset()         { *m = CreateTodoResponse{} }
func (m *CreateTodoResponse) String() string { return proto.CompactTextString(m) }
func (*CreateTodoResponse) ProtoMessage()    {}
func (*CreateTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f52905b89a1b23f, []int{1}
}

func (m *CreateTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateTodoResponse.Unmarshal(m, b)
}
func (m *CreateTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateTodoResponse.Marshal(b, m, deterministic)
}
func (m *CreateTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateTodoResponse.Merge(m, src)
}
func (m *CreateTodoResponse) XXX_Size() int {
	return xxx_messageInfo_CreateTodoResponse.Size(m)
}
func (m *CreateTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateTodoResponse proto.InternalMessageInfo

func (m *CreateTodoResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListTodoRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoRequest) Reset()         { *m = ListTodoRequest{} }
func (m *ListTodoRequest) String() string { return proto.CompactTextString(m) }
func (*ListTodoRequest) ProtoMessage()    {}
func (*ListTodoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f52905b89a1b23f, []int{2}
}

func (m *ListTodoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoRequest.Unmarshal(m, b)
}
func (m *ListTodoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoRequest.Marshal(b, m, deterministic)
}
func (m *ListTodoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoRequest.Merge(m, src)
}
func (m *ListTodoRequest) XXX_Size() int {
	return xxx_messageInfo_ListTodoRequest.Size(m)
}
func (m *ListTodoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoRequest proto.InternalMessageInfo

type ListTodoResponse struct {
	Todos                []*Todo  `protobuf:"bytes,1,rep,name=todos,proto3" json:"todos,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListTodoResponse) Reset()         { *m = ListTodoResponse{} }
func (m *ListTodoResponse) String() string { return proto.CompactTextString(m) }
func (*ListTodoResponse) ProtoMessage()    {}
func (*ListTodoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f52905b89a1b23f, []int{3}
}

func (m *ListTodoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListTodoResponse.Unmarshal(m, b)
}
func (m *ListTodoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListTodoResponse.Marshal(b, m, deterministic)
}
func (m *ListTodoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListTodoResponse.Merge(m, src)
}
func (m *ListTodoResponse) XXX_Size() int {
	return xxx_messageInfo_ListTodoResponse.Size(m)
}
func (m *ListTodoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListTodoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListTodoResponse proto.InternalMessageInfo

func (m *ListTodoResponse) GetTodos() []*Todo {
	if m != nil {
		return m.Todos
	}
	return nil
}

type Todo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Todo) Reset()         { *m = Todo{} }
func (m *Todo) String() string { return proto.CompactTextString(m) }
func (*Todo) ProtoMessage()    {}
func (*Todo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f52905b89a1b23f, []int{4}
}

func (m *Todo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Todo.Unmarshal(m, b)
}
func (m *Todo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Todo.Marshal(b, m, deterministic)
}
func (m *Todo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Todo.Merge(m, src)
}
func (m *Todo) XXX_Size() int {
	return xxx_messageInfo_Todo.Size(m)
}
func (m *Todo) XXX_DiscardUnknown() {
	xxx_messageInfo_Todo.DiscardUnknown(m)
}

var xxx_messageInfo_Todo proto.InternalMessageInfo

func (m *Todo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Todo) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateTodoRequest)(nil), "pb.CreateTodoRequest")
	proto.RegisterType((*CreateTodoResponse)(nil), "pb.CreateTodoResponse")
	proto.RegisterType((*ListTodoRequest)(nil), "pb.ListTodoRequest")
	proto.RegisterType((*ListTodoResponse)(nil), "pb.ListTodoResponse")
	proto.RegisterType((*Todo)(nil), "pb.Todo")
}

func init() { proto.RegisterFile("pb/todo.proto", fileDescriptor_1f52905b89a1b23f) }

var fileDescriptor_1f52905b89a1b23f = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x4d, 0xdc, 0x5d, 0x74, 0xc4, 0xd5, 0x8e, 0x55, 0x4a, 0x0f, 0x52, 0x82, 0x87, 0x0a,
	0x52, 0xa1, 0x3d, 0x79, 0xf6, 0xea, 0xa9, 0xf8, 0x02, 0x86, 0xe4, 0x10, 0x10, 0x27, 0x36, 0xe3,
	0xc1, 0x9b, 0x8f, 0x2e, 0x69, 0x2a, 0xad, 0xdd, 0xe3, 0xfc, 0xfc, 0xf3, 0xf1, 0xcd, 0xc0, 0xb9,
	0xd7, 0x8f, 0x4c, 0x86, 0x1a, 0x3f, 0x10, 0x13, 0x4a, 0xaf, 0xd5, 0x3d, 0x64, 0xcf, 0x83, 0x7d,
	0x63, 0xfb, 0x4a, 0x86, 0x7a, 0xfb, 0xf9, 0x65, 0x03, 0x63, 0x0e, 0x5b, 0x76, 0xfc, 0x6e, 0x0b,
	0x51, 0x89, 0xfa, 0xb4, 0x4f, 0x83, 0xba, 0x03, 0x5c, 0x56, 0x83, 0xa7, 0x8f, 0x60, 0x71, 0x0f,
	0xd2, 0x99, 0xa9, 0x28, 0x9d, 0x51, 0x19, 0x5c, 0xbc, 0xb8, 0xc0, 0x0b, 0x9c, 0x6a, 0xe1, 0x72,
	0x8e, 0xa6, 0xb5, 0x5b, 0xd8, 0x46, 0x93, 0x50, 0x88, 0xea, 0xb8, 0x3e, 0x6b, 0x4f, 0x1a, 0xaf,
	0x9b, 0xb1, 0x90, 0x62, 0xf5, 0x00, 0x9b, 0x38, 0xae, 0xf1, 0xb3, 0x9a, 0x5c, 0xa8, 0xb5, 0x3f,
	0x02, 0xf6, 0x09, 0xef, 0x29, 0x38, 0xa6, 0xe1, 0x1b, 0x9f, 0x60, 0x97, 0x6c, 0xf1, 0x3a, 0xb2,
	0x0f, 0x8e, 0x2c, 0x6f, 0xd6, 0x71, 0x32, 0x53, 0x47, 0xd8, 0xc1, 0x26, 0xfa, 0xe2, 0x55, 0x6c,
	0xac, 0x8e, 0x29, 0xf3, 0xff, 0xe1, 0xdf, 0x92, 0xde, 0x8d, 0x3f, 0xed, 0x7e, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x2f, 0x31, 0x6e, 0x19, 0x64, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TodoRepositoryClient is the client API for TodoRepository service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TodoRepositoryClient interface {
	Create(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error)
	List(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (*ListTodoResponse, error)
}

type todoRepositoryClient struct {
	cc *grpc.ClientConn
}

func NewTodoRepositoryClient(cc *grpc.ClientConn) TodoRepositoryClient {
	return &todoRepositoryClient{cc}
}

func (c *todoRepositoryClient) Create(ctx context.Context, in *CreateTodoRequest, opts ...grpc.CallOption) (*CreateTodoResponse, error) {
	out := new(CreateTodoResponse)
	err := c.cc.Invoke(ctx, "/pb.TodoRepository/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *todoRepositoryClient) List(ctx context.Context, in *ListTodoRequest, opts ...grpc.CallOption) (*ListTodoResponse, error) {
	out := new(ListTodoResponse)
	err := c.cc.Invoke(ctx, "/pb.TodoRepository/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TodoRepositoryServer is the server API for TodoRepository service.
type TodoRepositoryServer interface {
	Create(context.Context, *CreateTodoRequest) (*CreateTodoResponse, error)
	List(context.Context, *ListTodoRequest) (*ListTodoResponse, error)
}

// UnimplementedTodoRepositoryServer can be embedded to have forward compatible implementations.
type UnimplementedTodoRepositoryServer struct {
}

func (*UnimplementedTodoRepositoryServer) Create(ctx context.Context, req *CreateTodoRequest) (*CreateTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedTodoRepositoryServer) List(ctx context.Context, req *ListTodoRequest) (*ListTodoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterTodoRepositoryServer(s *grpc.Server, srv TodoRepositoryServer) {
	s.RegisterService(&_TodoRepository_serviceDesc, srv)
}

func _TodoRepository_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRepositoryServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoRepository/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRepositoryServer).Create(ctx, req.(*CreateTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TodoRepository_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTodoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TodoRepositoryServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TodoRepository/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TodoRepositoryServer).List(ctx, req.(*ListTodoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TodoRepository_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TodoRepository",
	HandlerType: (*TodoRepositoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TodoRepository_Create_Handler,
		},
		{
			MethodName: "List",
			Handler:    _TodoRepository_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/todo.proto",
}
