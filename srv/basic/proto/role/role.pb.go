// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: proto/role/role.proto

package role

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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_role_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_role_proto_msgTypes[0]
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
	return file_proto_role_role_proto_rawDescGZIP(), []int{0}
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_role_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_role_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_proto_role_role_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Access []string `protobuf:"bytes,2,rep,name=access,proto3" json:"access,omitempty"`
	Avatar string   `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_role_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_role_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_proto_role_role_proto_rawDescGZIP(), []int{2}
}

func (x *GetResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetResponse) GetAccess() []string {
	if x != nil {
		return x.Access
	}
	return nil
}

func (x *GetResponse) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type RoleAddOrUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *RoleAddOrUpdateRequest) Reset() {
	*x = RoleAddOrUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_role_role_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoleAddOrUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoleAddOrUpdateRequest) ProtoMessage() {}

func (x *RoleAddOrUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_role_role_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoleAddOrUpdateRequest.ProtoReflect.Descriptor instead.
func (*RoleAddOrUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_role_role_proto_rawDescGZIP(), []int{3}
}

func (x *RoleAddOrUpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *RoleAddOrUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_proto_role_role_proto protoreflect.FileDescriptor

var file_proto_role_role_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c, 0x65, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x22, 0x0a, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x51, 0x0a, 0x0b, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x3c,
	0x0a, 0x16, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x73, 0x0a, 0x04,
	0x52, 0x6f, 0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x10, 0x2e, 0x72, 0x6f,
	0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x1c, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x41, 0x64, 0x64,
	0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x72, 0x6f, 0x6c, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x6f, 0x6c,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_role_role_proto_rawDescOnce sync.Once
	file_proto_role_role_proto_rawDescData = file_proto_role_role_proto_rawDesc
)

func file_proto_role_role_proto_rawDescGZIP() []byte {
	file_proto_role_role_proto_rawDescOnce.Do(func() {
		file_proto_role_role_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_role_role_proto_rawDescData)
	})
	return file_proto_role_role_proto_rawDescData
}

var file_proto_role_role_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_role_role_proto_goTypes = []interface{}{
	(*Response)(nil),               // 0: role.Response
	(*GetRequest)(nil),             // 1: role.GetRequest
	(*GetResponse)(nil),            // 2: role.GetResponse
	(*RoleAddOrUpdateRequest)(nil), // 3: role.RoleAddOrUpdateRequest
}
var file_proto_role_role_proto_depIdxs = []int32{
	1, // 0: role.Role.Get:input_type -> role.GetRequest
	3, // 1: role.Role.AddOrUpdate:input_type -> role.RoleAddOrUpdateRequest
	2, // 2: role.Role.Get:output_type -> role.GetResponse
	0, // 3: role.Role.AddOrUpdate:output_type -> role.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_role_role_proto_init() }
func file_proto_role_role_proto_init() {
	if File_proto_role_role_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_role_role_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_role_role_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_proto_role_role_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_proto_role_role_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoleAddOrUpdateRequest); i {
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
			RawDescriptor: file_proto_role_role_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_role_role_proto_goTypes,
		DependencyIndexes: file_proto_role_role_proto_depIdxs,
		MessageInfos:      file_proto_role_role_proto_msgTypes,
	}.Build()
	File_proto_role_role_proto = out.File
	file_proto_role_role_proto_rawDesc = nil
	file_proto_role_role_proto_goTypes = nil
	file_proto_role_role_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// RoleClient is the client API for Role service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RoleClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	AddOrUpdate(ctx context.Context, in *RoleAddOrUpdateRequest, opts ...grpc.CallOption) (*Response, error)
}

type roleClient struct {
	cc grpc.ClientConnInterface
}

func NewRoleClient(cc grpc.ClientConnInterface) RoleClient {
	return &roleClient{cc}
}

func (c *roleClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/role.Role/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *roleClient) AddOrUpdate(ctx context.Context, in *RoleAddOrUpdateRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/role.Role/AddOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RoleServer is the server API for Role service.
type RoleServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	AddOrUpdate(context.Context, *RoleAddOrUpdateRequest) (*Response, error)
}

// UnimplementedRoleServer can be embedded to have forward compatible implementations.
type UnimplementedRoleServer struct {
}

func (*UnimplementedRoleServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedRoleServer) AddOrUpdate(context.Context, *RoleAddOrUpdateRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdate not implemented")
}

func RegisterRoleServer(s *grpc.Server, srv RoleServer) {
	s.RegisterService(&_Role_serviceDesc, srv)
}

func _Role_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.Role/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Role_AddOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleAddOrUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RoleServer).AddOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/role.Role/AddOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RoleServer).AddOrUpdate(ctx, req.(*RoleAddOrUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Role_serviceDesc = grpc.ServiceDesc{
	ServiceName: "role.Role",
	HandlerType: (*RoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Role_Get_Handler,
		},
		{
			MethodName: "AddOrUpdate",
			Handler:    _Role_AddOrUpdate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/role/role.proto",
}
