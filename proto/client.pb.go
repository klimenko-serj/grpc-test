// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: client.proto

package grpc_test

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Header struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header     string `protobuf:"bytes,2,opt,name=header,proto3" json:"header,omitempty"`
}

func (x *Header) Reset() {
	*x = Header{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Header) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Header) ProtoMessage() {}

func (x *Header) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Header.ProtoReflect.Descriptor instead.
func (*Header) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{0}
}

func (x *Header) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Header) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

type Body struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Body []byte `protobuf:"bytes,1,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Body) Reset() {
	*x = Body{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Body) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Body) ProtoMessage() {}

func (x *Body) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Body.ProtoReflect.Descriptor instead.
func (*Body) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{1}
}

func (x *Body) GetBody() []byte {
	if x != nil {
		return x.Body
	}
	return nil
}

type FinishMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error        bool   `protobuf:"varint,1,opt,name=error,proto3" json:"error,omitempty"`
	ErrorMessage string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *FinishMessage) Reset() {
	*x = FinishMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FinishMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FinishMessage) ProtoMessage() {}

func (x *FinishMessage) ProtoReflect() protoreflect.Message {
	mi := &file_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FinishMessage.ProtoReflect.Descriptor instead.
func (*FinishMessage) Descriptor() ([]byte, []int) {
	return file_client_proto_rawDescGZIP(), []int{2}
}

func (x *FinishMessage) GetError() bool {
	if x != nil {
		return x.Error
	}
	return false
}

func (x *FinishMessage) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_client_proto protoreflect.FileDescriptor

var file_client_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x06, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x22, 0x1a, 0x0a, 0x04, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x22, 0x49, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32,
	0xb5, 0x01, 0x0a, 0x09, 0x55, 0x72, 0x6c, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x12, 0x37, 0x0a,
	0x0a, 0x53, 0x65, 0x6e, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x1a, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x33, 0x0a, 0x08, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6f,
	0x64, 0x79, 0x12, 0x0f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x42,
	0x6f, 0x64, 0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x06, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x2e, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_proto_rawDescOnce sync.Once
	file_client_proto_rawDescData = file_client_proto_rawDesc
)

func file_client_proto_rawDescGZIP() []byte {
	file_client_proto_rawDescOnce.Do(func() {
		file_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_proto_rawDescData)
	})
	return file_client_proto_rawDescData
}

var file_client_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_client_proto_goTypes = []interface{}{
	(*Header)(nil),        // 0: grpc_test.Header
	(*Body)(nil),          // 1: grpc_test.Body
	(*FinishMessage)(nil), // 2: grpc_test.FinishMessage
	(*empty.Empty)(nil),   // 3: google.protobuf.Empty
}
var file_client_proto_depIdxs = []int32{
	0, // 0: grpc_test.UrlClient.SendHeader:input_type -> grpc_test.Header
	1, // 1: grpc_test.UrlClient.SendBody:input_type -> grpc_test.Body
	2, // 2: grpc_test.UrlClient.Finish:input_type -> grpc_test.FinishMessage
	3, // 3: grpc_test.UrlClient.SendHeader:output_type -> google.protobuf.Empty
	3, // 4: grpc_test.UrlClient.SendBody:output_type -> google.protobuf.Empty
	3, // 5: grpc_test.UrlClient.Finish:output_type -> google.protobuf.Empty
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_client_proto_init() }
func file_client_proto_init() {
	if File_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Header); i {
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
		file_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Body); i {
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
		file_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FinishMessage); i {
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
			RawDescriptor: file_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_proto_goTypes,
		DependencyIndexes: file_client_proto_depIdxs,
		MessageInfos:      file_client_proto_msgTypes,
	}.Build()
	File_client_proto = out.File
	file_client_proto_rawDesc = nil
	file_client_proto_goTypes = nil
	file_client_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UrlClientClient is the client API for UrlClient service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UrlClientClient interface {
	SendHeader(ctx context.Context, in *Header, opts ...grpc.CallOption) (*empty.Empty, error)
	SendBody(ctx context.Context, in *Body, opts ...grpc.CallOption) (*empty.Empty, error)
	Finish(ctx context.Context, in *FinishMessage, opts ...grpc.CallOption) (*empty.Empty, error)
}

type urlClientClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlClientClient(cc grpc.ClientConnInterface) UrlClientClient {
	return &urlClientClient{cc}
}

func (c *urlClientClient) SendHeader(ctx context.Context, in *Header, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc_test.UrlClient/SendHeader", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlClientClient) SendBody(ctx context.Context, in *Body, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc_test.UrlClient/SendBody", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlClientClient) Finish(ctx context.Context, in *FinishMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/grpc_test.UrlClient/Finish", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlClientServer is the server API for UrlClient service.
type UrlClientServer interface {
	SendHeader(context.Context, *Header) (*empty.Empty, error)
	SendBody(context.Context, *Body) (*empty.Empty, error)
	Finish(context.Context, *FinishMessage) (*empty.Empty, error)
}

// UnimplementedUrlClientServer can be embedded to have forward compatible implementations.
type UnimplementedUrlClientServer struct {
}

func (*UnimplementedUrlClientServer) SendHeader(context.Context, *Header) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeader not implemented")
}
func (*UnimplementedUrlClientServer) SendBody(context.Context, *Body) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendBody not implemented")
}
func (*UnimplementedUrlClientServer) Finish(context.Context, *FinishMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Finish not implemented")
}

func RegisterUrlClientServer(s *grpc.Server, srv UrlClientServer) {
	s.RegisterService(&_UrlClient_serviceDesc, srv)
}

func _UrlClient_SendHeader_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Header)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlClientServer).SendHeader(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_test.UrlClient/SendHeader",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlClientServer).SendHeader(ctx, req.(*Header))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlClient_SendBody_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Body)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlClientServer).SendBody(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_test.UrlClient/SendBody",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlClientServer).SendBody(ctx, req.(*Body))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlClient_Finish_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlClientServer).Finish(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_test.UrlClient/Finish",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlClientServer).Finish(ctx, req.(*FinishMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _UrlClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_test.UrlClient",
	HandlerType: (*UrlClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendHeader",
			Handler:    _UrlClient_SendHeader_Handler,
		},
		{
			MethodName: "SendBody",
			Handler:    _UrlClient_SendBody_Handler,
		},
		{
			MethodName: "Finish",
			Handler:    _UrlClient_Finish_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client.proto",
}
