// Code generated by protoc-gen-go. DO NOT EDIT.
// source: noderesolver.proto

package noderesolver

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
	grpc "google.golang.org/grpc"
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

//* Represents a request with a list of BaseSPIFFEIDs.
type ResolveRequest struct {
	//* A list of BaseSPIFFE Ids.
	BaseSpiffeIdList     []string `protobuf:"bytes,1,rep,name=baseSpiffeIdList,proto3" json:"baseSpiffeIdList,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResolveRequest) Reset()         { *m = ResolveRequest{} }
func (m *ResolveRequest) String() string { return proto.CompactTextString(m) }
func (*ResolveRequest) ProtoMessage()    {}
func (*ResolveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94c791929f88f3b, []int{0}
}

func (m *ResolveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveRequest.Unmarshal(m, b)
}
func (m *ResolveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveRequest.Marshal(b, m, deterministic)
}
func (m *ResolveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveRequest.Merge(m, src)
}
func (m *ResolveRequest) XXX_Size() int {
	return xxx_messageInfo_ResolveRequest.Size(m)
}
func (m *ResolveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveRequest proto.InternalMessageInfo

func (m *ResolveRequest) GetBaseSpiffeIdList() []string {
	if m != nil {
		return m.BaseSpiffeIdList
	}
	return nil
}

//* Represents a response with a map of SPIFFE ID to a list of Selectors.
type ResolveResponse struct {
	//* Map[SPIFFE_ID] => Selectors.
	Map                  map[string]*common.Selectors `protobuf:"bytes,1,rep,name=map,proto3" json:"map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *ResolveResponse) Reset()         { *m = ResolveResponse{} }
func (m *ResolveResponse) String() string { return proto.CompactTextString(m) }
func (*ResolveResponse) ProtoMessage()    {}
func (*ResolveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b94c791929f88f3b, []int{1}
}

func (m *ResolveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveResponse.Unmarshal(m, b)
}
func (m *ResolveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveResponse.Marshal(b, m, deterministic)
}
func (m *ResolveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveResponse.Merge(m, src)
}
func (m *ResolveResponse) XXX_Size() int {
	return xxx_messageInfo_ResolveResponse.Size(m)
}
func (m *ResolveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveResponse proto.InternalMessageInfo

func (m *ResolveResponse) GetMap() map[string]*common.Selectors {
	if m != nil {
		return m.Map
	}
	return nil
}

func init() {
	proto.RegisterType((*ResolveRequest)(nil), "spire.server.noderesolver.ResolveRequest")
	proto.RegisterType((*ResolveResponse)(nil), "spire.server.noderesolver.ResolveResponse")
	proto.RegisterMapType((map[string]*common.Selectors)(nil), "spire.server.noderesolver.ResolveResponse.MapEntry")
}

func init() { proto.RegisterFile("noderesolver.proto", fileDescriptor_b94c791929f88f3b) }

var fileDescriptor_b94c791929f88f3b = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x4f, 0xfa, 0x40,
	0x10, 0xc7, 0x53, 0xc8, 0xef, 0xa7, 0x0c, 0x3e, 0xc8, 0x5e, 0x84, 0x9e, 0x1a, 0x12, 0x0d, 0x92,
	0xd8, 0x26, 0x70, 0xf0, 0x11, 0x4f, 0x1a, 0x62, 0x48, 0x7c, 0xa5, 0xdc, 0x38, 0x59, 0x60, 0x8a,
	0x8d, 0x65, 0x77, 0xdd, 0xdd, 0x92, 0xf0, 0x27, 0x79, 0xf7, 0x0f, 0x34, 0xdd, 0x5d, 0x08, 0xc5,
	0xa8, 0x9c, 0x76, 0x32, 0xf3, 0xf9, 0xce, 0x6b, 0x07, 0x08, 0x65, 0x13, 0x14, 0x28, 0x59, 0x3a,
	0x47, 0xe1, 0x73, 0xc1, 0x14, 0x23, 0x0d, 0xc9, 0x13, 0x81, 0xbe, 0x44, 0x91, 0xfb, 0xd6, 0x01,
	0xd7, 0xd3, 0xa1, 0x60, 0xcc, 0x66, 0x33, 0x46, 0x03, 0x9e, 0x66, 0xd3, 0x64, 0xf9, 0x18, 0xb1,
	0xdb, 0x28, 0x10, 0xe6, 0x31, 0xa1, 0xe6, 0x35, 0x1c, 0x84, 0x26, 0x51, 0x88, 0xef, 0x19, 0x4a,
	0x45, 0xda, 0x50, 0x1b, 0x45, 0x12, 0x07, 0x3c, 0x89, 0x63, 0xec, 0x4f, 0xee, 0x13, 0xa9, 0xea,
	0x8e, 0x57, 0x6e, 0x55, 0xc2, 0x6f, 0xfe, 0xe6, 0x87, 0x03, 0x87, 0x2b, 0xb9, 0xe4, 0x8c, 0x4a,
	0x24, 0x3d, 0x28, 0xcf, 0x22, 0xae, 0x25, 0xd5, 0x4e, 0xd7, 0xff, 0xb1, 0x6f, 0x7f, 0x43, 0xe8,
	0x3f, 0x44, 0xbc, 0x47, 0x95, 0x58, 0x84, 0xb9, 0xde, 0x7d, 0x82, 0xdd, 0xa5, 0x83, 0xd4, 0xa0,
	0xfc, 0x86, 0x8b, 0xba, 0xe3, 0x39, 0xad, 0x4a, 0x98, 0x9b, 0xe4, 0x0c, 0xfe, 0xcd, 0xa3, 0x34,
	0xc3, 0x7a, 0xc9, 0x73, 0x5a, 0xd5, 0xce, 0x91, 0x2d, 0x63, 0x47, 0x1b, 0x60, 0x8a, 0x63, 0xc5,
	0x84, 0x0c, 0x0d, 0x75, 0x55, 0xba, 0x70, 0x3a, 0x9f, 0x25, 0xd8, 0x7b, 0x64, 0x13, 0xb4, 0x65,
	0x05, 0x79, 0x81, 0x1d, 0x6b, 0x93, 0xd3, 0x6d, 0xda, 0xd4, 0xeb, 0x71, 0xdb, 0xdb, 0x4f, 0x44,
	0x86, 0x50, 0xb9, 0x65, 0x34, 0x4e, 0xa6, 0x99, 0x40, 0x72, 0x5c, 0xec, 0xd1, 0x7e, 0xd0, 0x2a,
	0xbe, 0xcc, 0x7f, 0xf2, 0x17, 0x66, 0x73, 0xc7, 0xb0, 0x7f, 0x87, 0xea, 0x59, 0x87, 0xfb, 0x34,
	0x66, 0xab, 0x19, 0x8a, 0xc2, 0x02, 0xb3, 0x39, 0xc3, 0xaf, 0xa8, 0xa9, 0x73, 0x73, 0x39, 0x3c,
	0x9f, 0x26, 0xea, 0x35, 0x1b, 0xe5, 0x74, 0x20, 0xf5, 0xef, 0x07, 0xe6, 0x9e, 0xf4, 0x05, 0x59,
	0xdb, 0xac, 0x23, 0x58, 0x5f, 0xc7, 0xe8, 0xbf, 0x06, 0xba, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff,
	0x16, 0xc7, 0x5c, 0x81, 0xd0, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NodeResolverClient is the client API for NodeResolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeResolverClient interface {
	//* Retrieves a list of properties reflecting the current state of a particular node(s).
	Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error)
	//* Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	//* Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type nodeResolverClient struct {
	cc *grpc.ClientConn
}

func NewNodeResolverClient(cc *grpc.ClientConn) NodeResolverClient {
	return &nodeResolverClient{cc}
}

func (c *nodeResolverClient) Resolve(ctx context.Context, in *ResolveRequest, opts ...grpc.CallOption) (*ResolveResponse, error) {
	out := new(ResolveResponse)
	err := c.cc.Invoke(ctx, "/spire.server.noderesolver.NodeResolver/Resolve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeResolverClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.server.noderesolver.NodeResolver/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeResolverClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.server.noderesolver.NodeResolver/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeResolverServer is the server API for NodeResolver service.
type NodeResolverServer interface {
	//* Retrieves a list of properties reflecting the current state of a particular node(s).
	Resolve(context.Context, *ResolveRequest) (*ResolveResponse, error)
	//* Responsible for configuration of the plugin.
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	//* Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

func RegisterNodeResolverServer(s *grpc.Server, srv NodeResolverServer) {
	s.RegisterService(&_NodeResolver_serviceDesc, srv)
}

func _NodeResolver_Resolve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResolveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).Resolve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.noderesolver.NodeResolver/Resolve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).Resolve(ctx, req.(*ResolveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeResolver_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.noderesolver.NodeResolver/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeResolver_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeResolverServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.server.noderesolver.NodeResolver/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeResolverServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeResolver_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.server.noderesolver.NodeResolver",
	HandlerType: (*NodeResolverServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Resolve",
			Handler:    _NodeResolver_Resolve_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _NodeResolver_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeResolver_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "noderesolver.proto",
}
