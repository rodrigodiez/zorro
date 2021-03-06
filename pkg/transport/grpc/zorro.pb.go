// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zorro.proto

package grpc

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type MaskRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaskRequest) Reset()         { *m = MaskRequest{} }
func (m *MaskRequest) String() string { return proto.CompactTextString(m) }
func (*MaskRequest) ProtoMessage()    {}
func (*MaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_zorro_c1c888b9d84938b8, []int{0}
}
func (m *MaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaskRequest.Unmarshal(m, b)
}
func (m *MaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaskRequest.Marshal(b, m, deterministic)
}
func (dst *MaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaskRequest.Merge(dst, src)
}
func (m *MaskRequest) XXX_Size() int {
	return xxx_messageInfo_MaskRequest.Size(m)
}
func (m *MaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MaskRequest proto.InternalMessageInfo

func (m *MaskRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type MaskResponse struct {
	Value                string   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaskResponse) Reset()         { *m = MaskResponse{} }
func (m *MaskResponse) String() string { return proto.CompactTextString(m) }
func (*MaskResponse) ProtoMessage()    {}
func (*MaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_zorro_c1c888b9d84938b8, []int{1}
}
func (m *MaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaskResponse.Unmarshal(m, b)
}
func (m *MaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaskResponse.Marshal(b, m, deterministic)
}
func (dst *MaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaskResponse.Merge(dst, src)
}
func (m *MaskResponse) XXX_Size() int {
	return xxx_messageInfo_MaskResponse.Size(m)
}
func (m *MaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MaskResponse proto.InternalMessageInfo

func (m *MaskResponse) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UnmaskRequest struct {
	Value                string   `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnmaskRequest) Reset()         { *m = UnmaskRequest{} }
func (m *UnmaskRequest) String() string { return proto.CompactTextString(m) }
func (*UnmaskRequest) ProtoMessage()    {}
func (*UnmaskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_zorro_c1c888b9d84938b8, []int{2}
}
func (m *UnmaskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnmaskRequest.Unmarshal(m, b)
}
func (m *UnmaskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnmaskRequest.Marshal(b, m, deterministic)
}
func (dst *UnmaskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnmaskRequest.Merge(dst, src)
}
func (m *UnmaskRequest) XXX_Size() int {
	return xxx_messageInfo_UnmaskRequest.Size(m)
}
func (m *UnmaskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnmaskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnmaskRequest proto.InternalMessageInfo

func (m *UnmaskRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type UnmaskResponse struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnmaskResponse) Reset()         { *m = UnmaskResponse{} }
func (m *UnmaskResponse) String() string { return proto.CompactTextString(m) }
func (*UnmaskResponse) ProtoMessage()    {}
func (*UnmaskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_zorro_c1c888b9d84938b8, []int{3}
}
func (m *UnmaskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnmaskResponse.Unmarshal(m, b)
}
func (m *UnmaskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnmaskResponse.Marshal(b, m, deterministic)
}
func (dst *UnmaskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnmaskResponse.Merge(dst, src)
}
func (m *UnmaskResponse) XXX_Size() int {
	return xxx_messageInfo_UnmaskResponse.Size(m)
}
func (m *UnmaskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnmaskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnmaskResponse proto.InternalMessageInfo

func (m *UnmaskResponse) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*MaskRequest)(nil), "grpc.MaskRequest")
	proto.RegisterType((*MaskResponse)(nil), "grpc.MaskResponse")
	proto.RegisterType((*UnmaskRequest)(nil), "grpc.UnmaskRequest")
	proto.RegisterType((*UnmaskResponse)(nil), "grpc.UnmaskResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ZorroClient is the client API for Zorro service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ZorroClient interface {
	Mask(ctx context.Context, in *MaskRequest, opts ...grpc.CallOption) (*MaskResponse, error)
	Unmask(ctx context.Context, in *UnmaskRequest, opts ...grpc.CallOption) (*UnmaskResponse, error)
}

type zorroClient struct {
	cc *grpc.ClientConn
}

func NewZorroClient(cc *grpc.ClientConn) ZorroClient {
	return &zorroClient{cc}
}

func (c *zorroClient) Mask(ctx context.Context, in *MaskRequest, opts ...grpc.CallOption) (*MaskResponse, error) {
	out := new(MaskResponse)
	err := c.cc.Invoke(ctx, "/grpc.Zorro/Mask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *zorroClient) Unmask(ctx context.Context, in *UnmaskRequest, opts ...grpc.CallOption) (*UnmaskResponse, error) {
	out := new(UnmaskResponse)
	err := c.cc.Invoke(ctx, "/grpc.Zorro/Unmask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ZorroServer is the server API for Zorro service.
type ZorroServer interface {
	Mask(context.Context, *MaskRequest) (*MaskResponse, error)
	Unmask(context.Context, *UnmaskRequest) (*UnmaskResponse, error)
}

func RegisterZorroServer(s *grpc.Server, srv ZorroServer) {
	s.RegisterService(&_Zorro_serviceDesc, srv)
}

func _Zorro_Mask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZorroServer).Mask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Zorro/Mask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZorroServer).Mask(ctx, req.(*MaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Zorro_Unmask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnmaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ZorroServer).Unmask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Zorro/Unmask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ZorroServer).Unmask(ctx, req.(*UnmaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Zorro_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Zorro",
	HandlerType: (*ZorroServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mask",
			Handler:    _Zorro_Mask_Handler,
		},
		{
			MethodName: "Unmask",
			Handler:    _Zorro_Unmask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "zorro.proto",
}

func init() { proto.RegisterFile("zorro.proto", fileDescriptor_zorro_c1c888b9d84938b8) }

var fileDescriptor_zorro_c1c888b9d84938b8 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xae, 0xca, 0x2f, 0x2a,
	0xca, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2f, 0x2a, 0x48, 0x56, 0x92, 0xe7,
	0xe2, 0xf6, 0x4d, 0x2c, 0xce, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62,
	0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x54, 0xb8, 0x78,
	0x20, 0x0a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a,
	0x53, 0xa1, 0x6a, 0x20, 0x1c, 0x25, 0x55, 0x2e, 0xde, 0xd0, 0xbc, 0x5c, 0x24, 0x83, 0xb0, 0x2b,
	0x53, 0xe2, 0xe2, 0x83, 0x29, 0x83, 0x1a, 0x87, 0x61, 0xa1, 0x51, 0x3e, 0x17, 0x6b, 0x14, 0xc8,
	0x99, 0x42, 0xfa, 0x5c, 0x2c, 0x20, 0x9b, 0x85, 0x04, 0xf5, 0x40, 0x2e, 0xd5, 0x43, 0x72, 0xa6,
	0x94, 0x10, 0xb2, 0x10, 0xc4, 0x24, 0x25, 0x06, 0x21, 0x53, 0x2e, 0x36, 0x88, 0xe9, 0x42, 0xc2,
	0x10, 0x79, 0x14, 0x27, 0x49, 0x89, 0xa0, 0x0a, 0xc2, 0xb4, 0x25, 0xb1, 0x81, 0xc3, 0xc3, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x5d, 0xcd, 0xf8, 0x8f, 0x1e, 0x01, 0x00, 0x00,
}
