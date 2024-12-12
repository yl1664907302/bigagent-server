// 定义 config 数据结构

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.3
// source: config.proto

package grpc_config

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

// 网络信息配置
type NetworkInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Protocol string `protobuf:"bytes,1,opt,name=protocol,proto3" json:"protocol,omitempty"` // 网络协议，如http或https
	Host     string `protobuf:"bytes,2,opt,name=host,proto3" json:"host,omitempty"`         // 主机地址
	Port     int64  `protobuf:"varint,3,opt,name=port,proto3" json:"port,omitempty"`        // 端口号
	Path     string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`         // 请求路径
}

func (x *NetworkInfo) Reset() {
	*x = NetworkInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkInfo) ProtoMessage() {}

func (x *NetworkInfo) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkInfo.ProtoReflect.Descriptor instead.
func (*NetworkInfo) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{0}
}

func (x *NetworkInfo) GetProtocol() string {
	if x != nil {
		return x.Protocol
	}
	return ""
}

func (x *NetworkInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *NetworkInfo) GetPort() int64 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *NetworkInfo) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

// Agent 配置
type AgentConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Serct       string       `protobuf:"bytes,1,opt,name=serct,proto3" json:"serct,omitempty"`
	AuthName    string       `protobuf:"bytes,2,opt,name=auth_name,json=authName,proto3" json:"auth_name,omitempty"`
	DataName    string       `protobuf:"bytes,3,opt,name=data_name,json=dataName,proto3" json:"data_name,omitempty"`          // 认证名称
	SlotName    string       `protobuf:"bytes,4,opt,name=solt_name,json=slotName,proto3" json:"slot_name,omitempty"`          // 认证名称
	NetworkInfo *NetworkInfo `protobuf:"bytes,5,opt,name=network_info,json=networkInfo,proto3" json:"network_info,omitempty"` // 网络信息
	// 下面为前端传入的各种认证数据
	Token string `protobuf:"bytes,6,opt,name=token,proto3" json:"token,omitempty"` // Token
}

func (x *AgentConfig) Reset() {
	*x = AgentConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentConfig) ProtoMessage() {}

func (x *AgentConfig) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentConfig.ProtoReflect.Descriptor instead.
func (*AgentConfig) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{1}
}

func (x *AgentConfig) GetSerct() string {
	if x != nil {
		return x.Serct
	}
	return ""
}

func (x *AgentConfig) GetAuthName() string {
	if x != nil {
		return x.AuthName
	}
	return ""
}

func (x *AgentConfig) GetDataName() string {
	if x != nil {
		return x.DataName
	}
	return ""
}

func (x *AgentConfig) GetSlotName() string {
	if x != nil {
		return x.SlotName
	}
	return ""
}

func (x *AgentConfig) GetNetworkInfo() *NetworkInfo {
	if x != nil {
		return x.NetworkInfo
	}
	return nil
}

func (x *AgentConfig) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_config_proto_rawDescGZIP(), []int{2}
}

func (x *ResponseMessage) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ResponseMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_config_proto protoreflect.FileDescriptor

var file_config_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x65, 0x0a, 0x0b, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f,
	0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61,
	0x74, 0x68, 0x22, 0xcd, 0x01, 0x0a, 0x0b, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x65, 0x72, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x73, 0x65, 0x72, 0x63, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74,
	0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x6c, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x6c, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x3b, 0x0a, 0x0c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x3f, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x32, 0x5f, 0x0a, 0x12, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0f, 0x50, 0x75, 0x73,
	0x68, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x18, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x41, 0x67, 0x65, 0x6e, 0x74,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x42, 0x04, 0x5a, 0x02, 0x2e, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_config_proto_rawDescOnce sync.Once
	file_config_proto_rawDescData = file_config_proto_rawDesc
)

func file_config_proto_rawDescGZIP() []byte {
	file_config_proto_rawDescOnce.Do(func() {
		file_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_config_proto_rawDescData)
	})
	return file_config_proto_rawDescData
}

var file_config_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_config_proto_goTypes = []interface{}{
	(*NetworkInfo)(nil),     // 0: grpc_config.NetworkInfo
	(*AgentConfig)(nil),     // 1: grpc_config.AgentConfig
	(*ResponseMessage)(nil), // 2: grpc_config.ResponseMessage
}
var file_config_proto_depIdxs = []int32{
	0, // 0: grpc_config.AgentConfig.network_info:type_name -> grpc_config.NetworkInfo
	1, // 1: grpc_config.AgentConfigService.PushAgentConfig:input_type -> grpc_config.AgentConfig
	2, // 2: grpc_config.AgentConfigService.PushAgentConfig:output_type -> grpc_config.ResponseMessage
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_config_proto_init() }
func file_config_proto_init() {
	if File_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkInfo); i {
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
		file_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentConfig); i {
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
		file_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseMessage); i {
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
			RawDescriptor: file_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_config_proto_goTypes,
		DependencyIndexes: file_config_proto_depIdxs,
		MessageInfos:      file_config_proto_msgTypes,
	}.Build()
	File_config_proto = out.File
	file_config_proto_rawDesc = nil
	file_config_proto_goTypes = nil
	file_config_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AgentConfigServiceClient is the client API for AgentConfigService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AgentConfigServiceClient interface {
	// 示例方法：获取 Agent 配置
	PushAgentConfig(ctx context.Context, in *AgentConfig, opts ...grpc.CallOption) (*ResponseMessage, error)
}

type agentConfigServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAgentConfigServiceClient(cc grpc.ClientConnInterface) AgentConfigServiceClient {
	return &agentConfigServiceClient{cc}
}

func (c *agentConfigServiceClient) PushAgentConfig(ctx context.Context, in *AgentConfig, opts ...grpc.CallOption) (*ResponseMessage, error) {
	out := new(ResponseMessage)
	err := c.cc.Invoke(ctx, "/grpc_config.AgentConfigService/PushAgentConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AgentConfigServiceServer is the server API for AgentConfigService service.
type AgentConfigServiceServer interface {
	// 示例方法：获取 Agent 配置
	PushAgentConfig(context.Context, *AgentConfig) (*ResponseMessage, error)
}

// UnimplementedAgentConfigServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAgentConfigServiceServer struct {
}

func (*UnimplementedAgentConfigServiceServer) PushAgentConfig(context.Context, *AgentConfig) (*ResponseMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushAgentConfig not implemented")
}

func RegisterAgentConfigServiceServer(s *grpc.Server, srv AgentConfigServiceServer) {
	s.RegisterService(&_AgentConfigService_serviceDesc, srv)
}

func _AgentConfigService_PushAgentConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgentConfig)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentConfigServiceServer).PushAgentConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc_config.AgentConfigService/PushAgentConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentConfigServiceServer).PushAgentConfig(ctx, req.(*AgentConfig))
	}
	return interceptor(ctx, in, info, handler)
}

var _AgentConfigService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_config.AgentConfigService",
	HandlerType: (*AgentConfigServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushAgentConfig",
			Handler:    _AgentConfigService_PushAgentConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "config.proto",
}
