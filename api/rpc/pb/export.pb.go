// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: export.proto

package pb

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
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

type EptRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HashMark string `protobuf:"bytes,1,opt,name=hash_mark,json=hashMark,proto3" json:"hash_mark,omitempty"`
	Total    int32  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Header   string `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	Data     string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EptRequest) Reset() {
	*x = EptRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_export_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EptRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EptRequest) ProtoMessage() {}

func (x *EptRequest) ProtoReflect() protoreflect.Message {
	mi := &file_export_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EptRequest.ProtoReflect.Descriptor instead.
func (*EptRequest) Descriptor() ([]byte, []int) {
	return file_export_proto_rawDescGZIP(), []int{0}
}

func (x *EptRequest) GetHashMark() string {
	if x != nil {
		return x.HashMark
	}
	return ""
}

func (x *EptRequest) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *EptRequest) GetHeader() string {
	if x != nil {
		return x.Header
	}
	return ""
}

func (x *EptRequest) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type EptReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *EptReply) Reset() {
	*x = EptReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_export_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EptReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EptReply) ProtoMessage() {}

func (x *EptReply) ProtoReflect() protoreflect.Message {
	mi := &file_export_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EptReply.ProtoReflect.Descriptor instead.
func (*EptReply) Descriptor() ([]byte, []int) {
	return file_export_proto_rawDescGZIP(), []int{1}
}

func (x *EptReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *EptReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_export_proto protoreflect.FileDescriptor

var file_export_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x65, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x6b, 0x0a, 0x0a, 0x45, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x61, 0x73, 0x68, 0x4d, 0x61, 0x72, 0x6b, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x30, 0x0a, 0x08, 0x45, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x32, 0x2f, 0x0a, 0x06, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x25, 0x0a, 0x03, 0x45,
	0x70, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x70, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x70, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_export_proto_rawDescOnce sync.Once
	file_export_proto_rawDescData = file_export_proto_rawDesc
)

func file_export_proto_rawDescGZIP() []byte {
	file_export_proto_rawDescOnce.Do(func() {
		file_export_proto_rawDescData = protoimpl.X.CompressGZIP(file_export_proto_rawDescData)
	})
	return file_export_proto_rawDescData
}

var file_export_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_export_proto_goTypes = []interface{}{
	(*EptRequest)(nil), // 0: pb.EptRequest
	(*EptReply)(nil),   // 1: pb.EptReply
}
var file_export_proto_depIdxs = []int32{
	0, // 0: pb.Export.Ept:input_type -> pb.EptRequest
	1, // 1: pb.Export.Ept:output_type -> pb.EptReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_export_proto_init() }
func file_export_proto_init() {
	if File_export_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_export_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EptRequest); i {
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
		file_export_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EptReply); i {
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
			RawDescriptor: file_export_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_export_proto_goTypes,
		DependencyIndexes: file_export_proto_depIdxs,
		MessageInfos:      file_export_proto_msgTypes,
	}.Build()
	File_export_proto = out.File
	file_export_proto_rawDesc = nil
	file_export_proto_goTypes = nil
	file_export_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ExportClient is the client API for Export service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ExportClient interface {
	Ept(ctx context.Context, in *EptRequest, opts ...grpc.CallOption) (*EptReply, error)
}

type exportClient struct {
	cc grpc.ClientConnInterface
}

func NewExportClient(cc grpc.ClientConnInterface) ExportClient {
	return &exportClient{cc}
}

func (c *exportClient) Ept(ctx context.Context, in *EptRequest, opts ...grpc.CallOption) (*EptReply, error) {
	out := new(EptReply)
	err := c.cc.Invoke(ctx, "/pb.Export/Ept", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExportServer is the server API for Export service.
type ExportServer interface {
	Ept(context.Context, *EptRequest) (*EptReply, error)
}

// UnimplementedExportServer can be embedded to have forward compatible implementations.
type UnimplementedExportServer struct {
}

func (*UnimplementedExportServer) Ept(context.Context, *EptRequest) (*EptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ept not implemented")
}

func RegisterExportServer(s *grpc.Server, srv ExportServer) {
	s.RegisterService(&_Export_serviceDesc, srv)
}

func _Export_Ept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExportServer).Ept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Export/Ept",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExportServer).Ept(ctx, req.(*EptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Export_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Export",
	HandlerType: (*ExportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ept",
			Handler:    _Export_Ept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "export.proto",
}
