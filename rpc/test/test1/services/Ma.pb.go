// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: Ma.proto

package services

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

type MaRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age    int32 `protobuf:"varint,1,opt,name=age,proto3" json:"age,omitempty"`
	Height int32 `protobuf:"varint,180,opt,name=height,proto3" json:"height,omitempty"`
}

func (x *MaRequest) Reset() {
	*x = MaRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ma_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaRequest) ProtoMessage() {}

func (x *MaRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Ma_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaRequest.ProtoReflect.Descriptor instead.
func (*MaRequest) Descriptor() ([]byte, []int) {
	return file_Ma_proto_rawDescGZIP(), []int{0}
}

func (x *MaRequest) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *MaRequest) GetHeight() int32 {
	if x != nil {
		return x.Height
	}
	return 0
}

type MaResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name int32 `protobuf:"varint,3333,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MaResponse) Reset() {
	*x = MaResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ma_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaResponse) ProtoMessage() {}

func (x *MaResponse) ProtoReflect() protoreflect.Message {
	mi := &file_Ma_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaResponse.ProtoReflect.Descriptor instead.
func (*MaResponse) Descriptor() ([]byte, []int) {
	return file_Ma_proto_rawDescGZIP(), []int{1}
}

func (x *MaResponse) GetName() int32 {
	if x != nil {
		return x.Name
	}
	return 0
}

type MaResponse2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name2 int32 `protobuf:"varint,3333,opt,name=name2,proto3" json:"name2,omitempty"`
}

func (x *MaResponse2) Reset() {
	*x = MaResponse2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Ma_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MaResponse2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MaResponse2) ProtoMessage() {}

func (x *MaResponse2) ProtoReflect() protoreflect.Message {
	mi := &file_Ma_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MaResponse2.ProtoReflect.Descriptor instead.
func (*MaResponse2) Descriptor() ([]byte, []int) {
	return file_Ma_proto_rawDescGZIP(), []int{2}
}

func (x *MaResponse2) GetName2() int32 {
	if x != nil {
		return x.Name2
	}
	return 0
}

var File_Ma_proto protoreflect.FileDescriptor

var file_Ma_proto_rawDesc = []byte{
	0x0a, 0x08, 0x4d, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x09, 0x4d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x61, 0x67, 0x65, 0x12, 0x17, 0x0a, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0xb4, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x21, 0x0a, 0x0a,
	0x4d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x13, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x85, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x24, 0x0a, 0x0b, 0x4d, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x12, 0x15,
	0x0a, 0x05, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x18, 0x85, 0x1a, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0x32, 0x41, 0x0a, 0x09, 0x4d, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x4d, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x3b, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Ma_proto_rawDescOnce sync.Once
	file_Ma_proto_rawDescData = file_Ma_proto_rawDesc
)

func file_Ma_proto_rawDescGZIP() []byte {
	file_Ma_proto_rawDescOnce.Do(func() {
		file_Ma_proto_rawDescData = protoimpl.X.CompressGZIP(file_Ma_proto_rawDescData)
	})
	return file_Ma_proto_rawDescData
}

var file_Ma_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_Ma_proto_goTypes = []interface{}{
	(*MaRequest)(nil),   // 0: services.MaRequest
	(*MaResponse)(nil),  // 1: services.MaResponse
	(*MaResponse2)(nil), // 2: services.MaResponse2
}
var file_Ma_proto_depIdxs = []int32{
	0, // 0: services.MaService.GetName:input_type -> services.MaRequest
	1, // 1: services.MaService.GetName:output_type -> services.MaResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Ma_proto_init() }
func file_Ma_proto_init() {
	if File_Ma_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Ma_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaRequest); i {
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
		file_Ma_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaResponse); i {
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
		file_Ma_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MaResponse2); i {
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
			RawDescriptor: file_Ma_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Ma_proto_goTypes,
		DependencyIndexes: file_Ma_proto_depIdxs,
		MessageInfos:      file_Ma_proto_msgTypes,
	}.Build()
	File_Ma_proto = out.File
	file_Ma_proto_rawDesc = nil
	file_Ma_proto_goTypes = nil
	file_Ma_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the rpc-middle package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MaServiceClient is the client API for MaService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MaServiceClient interface {
	GetName(ctx context.Context, in *MaRequest, opts ...grpc.CallOption) (*MaResponse, error)
}

type maServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMaServiceClient(cc grpc.ClientConnInterface) MaServiceClient {
	return &maServiceClient{cc}
}

func (c *maServiceClient) GetName(ctx context.Context, in *MaRequest, opts ...grpc.CallOption) (*MaResponse, error) {
	out := new(MaResponse)
	err := c.cc.Invoke(ctx, "/services.MaService/GetName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MaServiceServer is the server API for MaService service.
type MaServiceServer interface {
	GetName(context.Context, *MaRequest) (*MaResponse, error)
}

// UnimplementedMaServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMaServiceServer struct {
}

func (*UnimplementedMaServiceServer) GetName(context.Context, *MaRequest) (*MaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetName not implemented")
}

func RegisterMaServiceServer(s *grpc.Server, srv MaServiceServer) {
	s.RegisterService(&_MaService_serviceDesc, srv)
}

func _MaService_GetName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MaServiceServer).GetName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.MaService/GetName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MaServiceServer).GetName(ctx, req.(*MaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "services.MaService",
	HandlerType: (*MaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetName",
			Handler:    _MaService_GetName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "Ma.proto",
}