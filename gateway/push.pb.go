// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push.proto

package dialog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/types"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Registering push token on server
type RequestRegisterGooglePush struct {
	ProjectId            int64    `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestRegisterGooglePush) Reset()         { *m = RequestRegisterGooglePush{} }
func (m *RequestRegisterGooglePush) String() string { return proto.CompactTextString(m) }
func (*RequestRegisterGooglePush) ProtoMessage()    {}
func (*RequestRegisterGooglePush) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_f7e7c8687bcbf997, []int{0}
}
func (m *RequestRegisterGooglePush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestRegisterGooglePush.Unmarshal(m, b)
}
func (m *RequestRegisterGooglePush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestRegisterGooglePush.Marshal(b, m, deterministic)
}
func (dst *RequestRegisterGooglePush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRegisterGooglePush.Merge(dst, src)
}
func (m *RequestRegisterGooglePush) XXX_Size() int {
	return xxx_messageInfo_RequestRegisterGooglePush.Size(m)
}
func (m *RequestRegisterGooglePush) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRegisterGooglePush.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRegisterGooglePush proto.InternalMessageInfo

func (m *RequestRegisterGooglePush) GetProjectId() int64 {
	if m != nil {
		return m.ProjectId
	}
	return 0
}

func (m *RequestRegisterGooglePush) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Unregistering Google Push
type RequestUnregisterGooglePush struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUnregisterGooglePush) Reset()         { *m = RequestUnregisterGooglePush{} }
func (m *RequestUnregisterGooglePush) String() string { return proto.CompactTextString(m) }
func (*RequestUnregisterGooglePush) ProtoMessage()    {}
func (*RequestUnregisterGooglePush) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_f7e7c8687bcbf997, []int{1}
}
func (m *RequestUnregisterGooglePush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUnregisterGooglePush.Unmarshal(m, b)
}
func (m *RequestUnregisterGooglePush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUnregisterGooglePush.Marshal(b, m, deterministic)
}
func (dst *RequestUnregisterGooglePush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUnregisterGooglePush.Merge(dst, src)
}
func (m *RequestUnregisterGooglePush) XXX_Size() int {
	return xxx_messageInfo_RequestUnregisterGooglePush.Size(m)
}
func (m *RequestUnregisterGooglePush) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUnregisterGooglePush.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUnregisterGooglePush proto.InternalMessageInfo

func (m *RequestUnregisterGooglePush) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Registering apple push on server
type RequestRegisterApplePush struct {
	ApnsKey              int32    `protobuf:"varint,1,opt,name=apns_key,json=apnsKey,proto3" json:"apns_key,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestRegisterApplePush) Reset()         { *m = RequestRegisterApplePush{} }
func (m *RequestRegisterApplePush) String() string { return proto.CompactTextString(m) }
func (*RequestRegisterApplePush) ProtoMessage()    {}
func (*RequestRegisterApplePush) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_f7e7c8687bcbf997, []int{2}
}
func (m *RequestRegisterApplePush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestRegisterApplePush.Unmarshal(m, b)
}
func (m *RequestRegisterApplePush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestRegisterApplePush.Marshal(b, m, deterministic)
}
func (dst *RequestRegisterApplePush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRegisterApplePush.Merge(dst, src)
}
func (m *RequestRegisterApplePush) XXX_Size() int {
	return xxx_messageInfo_RequestRegisterApplePush.Size(m)
}
func (m *RequestRegisterApplePush) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRegisterApplePush.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRegisterApplePush proto.InternalMessageInfo

func (m *RequestRegisterApplePush) GetApnsKey() int32 {
	if m != nil {
		return m.ApnsKey
	}
	return 0
}

func (m *RequestRegisterApplePush) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Unregistering Apple Push
type RequestUnregisterApplePush struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUnregisterApplePush) Reset()         { *m = RequestUnregisterApplePush{} }
func (m *RequestUnregisterApplePush) String() string { return proto.CompactTextString(m) }
func (*RequestUnregisterApplePush) ProtoMessage()    {}
func (*RequestUnregisterApplePush) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_f7e7c8687bcbf997, []int{3}
}
func (m *RequestUnregisterApplePush) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUnregisterApplePush.Unmarshal(m, b)
}
func (m *RequestUnregisterApplePush) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUnregisterApplePush.Marshal(b, m, deterministic)
}
func (dst *RequestUnregisterApplePush) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUnregisterApplePush.Merge(dst, src)
}
func (m *RequestUnregisterApplePush) XXX_Size() int {
	return xxx_messageInfo_RequestUnregisterApplePush.Size(m)
}
func (m *RequestUnregisterApplePush) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUnregisterApplePush.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUnregisterApplePush proto.InternalMessageInfo

func (m *RequestUnregisterApplePush) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Registration of a new Apple's PushKit tokens
type RequestRegisterApplePushKit struct {
	ApnsKey              int32    `protobuf:"varint,1,opt,name=apns_key,json=apnsKey,proto3" json:"apns_key,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestRegisterApplePushKit) Reset()         { *m = RequestRegisterApplePushKit{} }
func (m *RequestRegisterApplePushKit) String() string { return proto.CompactTextString(m) }
func (*RequestRegisterApplePushKit) ProtoMessage()    {}
func (*RequestRegisterApplePushKit) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_f7e7c8687bcbf997, []int{4}
}
func (m *RequestRegisterApplePushKit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestRegisterApplePushKit.Unmarshal(m, b)
}
func (m *RequestRegisterApplePushKit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestRegisterApplePushKit.Marshal(b, m, deterministic)
}
func (dst *RequestRegisterApplePushKit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRegisterApplePushKit.Merge(dst, src)
}
func (m *RequestRegisterApplePushKit) XXX_Size() int {
	return xxx_messageInfo_RequestRegisterApplePushKit.Size(m)
}
func (m *RequestRegisterApplePushKit) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRegisterApplePushKit.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRegisterApplePushKit proto.InternalMessageInfo

func (m *RequestRegisterApplePushKit) GetApnsKey() int32 {
	if m != nil {
		return m.ApnsKey
	}
	return 0
}

func (m *RequestRegisterApplePushKit) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Unregistering Apple Push Kit token
type RequestUnregisterApplePushKit struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUnregisterApplePushKit) Reset()         { *m = RequestUnregisterApplePushKit{} }
func (m *RequestUnregisterApplePushKit) String() string { return proto.CompactTextString(m) }
func (*RequestUnregisterApplePushKit) ProtoMessage()    {}
func (*RequestUnregisterApplePushKit) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_f7e7c8687bcbf997, []int{5}
}
func (m *RequestUnregisterApplePushKit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUnregisterApplePushKit.Unmarshal(m, b)
}
func (m *RequestUnregisterApplePushKit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUnregisterApplePushKit.Marshal(b, m, deterministic)
}
func (dst *RequestUnregisterApplePushKit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUnregisterApplePushKit.Merge(dst, src)
}
func (m *RequestUnregisterApplePushKit) XXX_Size() int {
	return xxx_messageInfo_RequestUnregisterApplePushKit.Size(m)
}
func (m *RequestUnregisterApplePushKit) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUnregisterApplePushKit.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUnregisterApplePushKit proto.InternalMessageInfo

func (m *RequestUnregisterApplePushKit) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Registering Apple Push Token
type RequestRegisterApplePushToken struct {
	BundleId             string   `protobuf:"bytes,1,opt,name=bundle_id,json=bundleId,proto3" json:"bundle_id,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestRegisterApplePushToken) Reset()         { *m = RequestRegisterApplePushToken{} }
func (m *RequestRegisterApplePushToken) String() string { return proto.CompactTextString(m) }
func (*RequestRegisterApplePushToken) ProtoMessage()    {}
func (*RequestRegisterApplePushToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_f7e7c8687bcbf997, []int{6}
}
func (m *RequestRegisterApplePushToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestRegisterApplePushToken.Unmarshal(m, b)
}
func (m *RequestRegisterApplePushToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestRegisterApplePushToken.Marshal(b, m, deterministic)
}
func (dst *RequestRegisterApplePushToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestRegisterApplePushToken.Merge(dst, src)
}
func (m *RequestRegisterApplePushToken) XXX_Size() int {
	return xxx_messageInfo_RequestRegisterApplePushToken.Size(m)
}
func (m *RequestRegisterApplePushToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestRegisterApplePushToken.DiscardUnknown(m)
}

var xxx_messageInfo_RequestRegisterApplePushToken proto.InternalMessageInfo

func (m *RequestRegisterApplePushToken) GetBundleId() string {
	if m != nil {
		return m.BundleId
	}
	return ""
}

func (m *RequestRegisterApplePushToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Unregister Apple Push token
type RequestUnregisterApplePushToken struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestUnregisterApplePushToken) Reset()         { *m = RequestUnregisterApplePushToken{} }
func (m *RequestUnregisterApplePushToken) String() string { return proto.CompactTextString(m) }
func (*RequestUnregisterApplePushToken) ProtoMessage()    {}
func (*RequestUnregisterApplePushToken) Descriptor() ([]byte, []int) {
	return fileDescriptor_push_f7e7c8687bcbf997, []int{7}
}
func (m *RequestUnregisterApplePushToken) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestUnregisterApplePushToken.Unmarshal(m, b)
}
func (m *RequestUnregisterApplePushToken) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestUnregisterApplePushToken.Marshal(b, m, deterministic)
}
func (dst *RequestUnregisterApplePushToken) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestUnregisterApplePushToken.Merge(dst, src)
}
func (m *RequestUnregisterApplePushToken) XXX_Size() int {
	return xxx_messageInfo_RequestUnregisterApplePushToken.Size(m)
}
func (m *RequestUnregisterApplePushToken) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestUnregisterApplePushToken.DiscardUnknown(m)
}

var xxx_messageInfo_RequestUnregisterApplePushToken proto.InternalMessageInfo

func (m *RequestUnregisterApplePushToken) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*RequestRegisterGooglePush)(nil), "dialog.RequestRegisterGooglePush")
	proto.RegisterType((*RequestUnregisterGooglePush)(nil), "dialog.RequestUnregisterGooglePush")
	proto.RegisterType((*RequestRegisterApplePush)(nil), "dialog.RequestRegisterApplePush")
	proto.RegisterType((*RequestUnregisterApplePush)(nil), "dialog.RequestUnregisterApplePush")
	proto.RegisterType((*RequestRegisterApplePushKit)(nil), "dialog.RequestRegisterApplePushKit")
	proto.RegisterType((*RequestUnregisterApplePushKit)(nil), "dialog.RequestUnregisterApplePushKit")
	proto.RegisterType((*RequestRegisterApplePushToken)(nil), "dialog.RequestRegisterApplePushToken")
	proto.RegisterType((*RequestUnregisterApplePushToken)(nil), "dialog.RequestUnregisterApplePushToken")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PushClient is the client API for Push service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PushClient interface {
	RegisterGooglePush(ctx context.Context, in *RequestRegisterGooglePush, opts ...grpc.CallOption) (*ResponseVoid, error)
	UnregisterGooglePush(ctx context.Context, in *RequestUnregisterGooglePush, opts ...grpc.CallOption) (*ResponseVoid, error)
	RegisterApplePush(ctx context.Context, in *RequestRegisterApplePush, opts ...grpc.CallOption) (*ResponseVoid, error)
	UnregisterApplePush(ctx context.Context, in *RequestUnregisterApplePush, opts ...grpc.CallOption) (*ResponseVoid, error)
	RegisterApplePushKit(ctx context.Context, in *RequestRegisterApplePushKit, opts ...grpc.CallOption) (*ResponseVoid, error)
	UnregisterApplePushKit(ctx context.Context, in *RequestUnregisterApplePushKit, opts ...grpc.CallOption) (*ResponseVoid, error)
	RegisterApplePushToken(ctx context.Context, in *RequestRegisterApplePushToken, opts ...grpc.CallOption) (*ResponseVoid, error)
	UnregisterApplePushToken(ctx context.Context, in *RequestUnregisterApplePushToken, opts ...grpc.CallOption) (*ResponseVoid, error)
}

type pushClient struct {
	cc *grpc.ClientConn
}

func NewPushClient(cc *grpc.ClientConn) PushClient {
	return &pushClient{cc}
}

func (c *pushClient) RegisterGooglePush(ctx context.Context, in *RequestRegisterGooglePush, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.Push/RegisterGooglePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) UnregisterGooglePush(ctx context.Context, in *RequestUnregisterGooglePush, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.Push/UnregisterGooglePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) RegisterApplePush(ctx context.Context, in *RequestRegisterApplePush, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.Push/RegisterApplePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) UnregisterApplePush(ctx context.Context, in *RequestUnregisterApplePush, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.Push/UnregisterApplePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) RegisterApplePushKit(ctx context.Context, in *RequestRegisterApplePushKit, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.Push/RegisterApplePushKit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) UnregisterApplePushKit(ctx context.Context, in *RequestUnregisterApplePushKit, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.Push/UnregisterApplePushKit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) RegisterApplePushToken(ctx context.Context, in *RequestRegisterApplePushToken, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.Push/RegisterApplePushToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pushClient) UnregisterApplePushToken(ctx context.Context, in *RequestUnregisterApplePushToken, opts ...grpc.CallOption) (*ResponseVoid, error) {
	out := new(ResponseVoid)
	err := c.cc.Invoke(ctx, "/dialog.Push/UnregisterApplePushToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PushServer is the server API for Push service.
type PushServer interface {
	RegisterGooglePush(context.Context, *RequestRegisterGooglePush) (*ResponseVoid, error)
	UnregisterGooglePush(context.Context, *RequestUnregisterGooglePush) (*ResponseVoid, error)
	RegisterApplePush(context.Context, *RequestRegisterApplePush) (*ResponseVoid, error)
	UnregisterApplePush(context.Context, *RequestUnregisterApplePush) (*ResponseVoid, error)
	RegisterApplePushKit(context.Context, *RequestRegisterApplePushKit) (*ResponseVoid, error)
	UnregisterApplePushKit(context.Context, *RequestUnregisterApplePushKit) (*ResponseVoid, error)
	RegisterApplePushToken(context.Context, *RequestRegisterApplePushToken) (*ResponseVoid, error)
	UnregisterApplePushToken(context.Context, *RequestUnregisterApplePushToken) (*ResponseVoid, error)
}

func RegisterPushServer(s *grpc.Server, srv PushServer) {
	s.RegisterService(&_Push_serviceDesc, srv)
}

func _Push_RegisterGooglePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRegisterGooglePush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).RegisterGooglePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Push/RegisterGooglePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).RegisterGooglePush(ctx, req.(*RequestRegisterGooglePush))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_UnregisterGooglePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUnregisterGooglePush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).UnregisterGooglePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Push/UnregisterGooglePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).UnregisterGooglePush(ctx, req.(*RequestUnregisterGooglePush))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_RegisterApplePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRegisterApplePush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).RegisterApplePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Push/RegisterApplePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).RegisterApplePush(ctx, req.(*RequestRegisterApplePush))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_UnregisterApplePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUnregisterApplePush)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).UnregisterApplePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Push/UnregisterApplePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).UnregisterApplePush(ctx, req.(*RequestUnregisterApplePush))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_RegisterApplePushKit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRegisterApplePushKit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).RegisterApplePushKit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Push/RegisterApplePushKit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).RegisterApplePushKit(ctx, req.(*RequestRegisterApplePushKit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_UnregisterApplePushKit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUnregisterApplePushKit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).UnregisterApplePushKit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Push/UnregisterApplePushKit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).UnregisterApplePushKit(ctx, req.(*RequestUnregisterApplePushKit))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_RegisterApplePushToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRegisterApplePushToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).RegisterApplePushToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Push/RegisterApplePushToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).RegisterApplePushToken(ctx, req.(*RequestRegisterApplePushToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Push_UnregisterApplePushToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestUnregisterApplePushToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PushServer).UnregisterApplePushToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Push/UnregisterApplePushToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PushServer).UnregisterApplePushToken(ctx, req.(*RequestUnregisterApplePushToken))
	}
	return interceptor(ctx, in, info, handler)
}

var _Push_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialog.Push",
	HandlerType: (*PushServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterGooglePush",
			Handler:    _Push_RegisterGooglePush_Handler,
		},
		{
			MethodName: "UnregisterGooglePush",
			Handler:    _Push_UnregisterGooglePush_Handler,
		},
		{
			MethodName: "RegisterApplePush",
			Handler:    _Push_RegisterApplePush_Handler,
		},
		{
			MethodName: "UnregisterApplePush",
			Handler:    _Push_UnregisterApplePush_Handler,
		},
		{
			MethodName: "RegisterApplePushKit",
			Handler:    _Push_RegisterApplePushKit_Handler,
		},
		{
			MethodName: "UnregisterApplePushKit",
			Handler:    _Push_UnregisterApplePushKit_Handler,
		},
		{
			MethodName: "RegisterApplePushToken",
			Handler:    _Push_RegisterApplePushToken_Handler,
		},
		{
			MethodName: "UnregisterApplePushToken",
			Handler:    _Push_UnregisterApplePushToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "push.proto",
}

func init() { proto.RegisterFile("push.proto", fileDescriptor_push_f7e7c8687bcbf997) }

var fileDescriptor_push_f7e7c8687bcbf997 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x96, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xe5, 0xc1, 0xba, 0xf6, 0x13, 0x1c, 0x96, 0x95, 0xd1, 0x96, 0xa1, 0x75, 0xde, 0xba,
	0x95, 0x16, 0x92, 0x0d, 0x6e, 0xbb, 0x4c, 0xf4, 0x32, 0xa1, 0x5e, 0x50, 0x05, 0x1c, 0xa9, 0xd2,
	0xc4, 0xa4, 0x66, 0x99, 0x6d, 0xe2, 0xb4, 0x68, 0x48, 0x48, 0x88, 0xc3, 0x24, 0x2a, 0x38, 0xf1,
	0x04, 0x3c, 0x13, 0x6f, 0x30, 0x71, 0xda, 0x53, 0xa0, 0x24, 0x6d, 0xd6, 0x35, 0x31, 0xd1, 0xe8,
	0xb4, 0x53, 0x2b, 0xfb, 0xf3, 0xf7, 0xff, 0xfd, 0x3f, 0xf9, 0x6f, 0x05, 0x40, 0x0c, 0x64, 0x5f,
	0x17, 0x1e, 0xf7, 0xb9, 0x96, 0xb3, 0xa9, 0xe9, 0x72, 0xa7, 0xb2, 0xe6, 0x70, 0xee, 0xb8, 0xc4,
	0x30, 0x05, 0x35, 0x4c, 0xc6, 0xb8, 0x6f, 0xfa, 0x94, 0x33, 0x19, 0x55, 0x55, 0x96, 0x6d, 0xf2,
	0x8e, 0x32, 0x3a, 0xbd, 0xb4, 0x72, 0x4c, 0xa5, 0x45, 0x5c, 0xd7, 0x64, 0x84, 0x0f, 0x26, 0x8b,
	0xf7, 0xa4, 0x65, 0xba, 0xa6, 0xe8, 0x19, 0xe3, 0xdf, 0x68, 0x19, 0x8f, 0x10, 0x94, 0x3b, 0xe4,
	0xc3, 0x80, 0x48, 0xbf, 0x43, 0x1c, 0x2a, 0x7d, 0xe2, 0x1d, 0x86, 0x72, 0x2f, 0x07, 0xb2, 0xaf,
	0x35, 0x01, 0x84, 0xc7, 0xdf, 0x13, 0xcb, 0xef, 0x52, 0xbb, 0x84, 0xaa, 0xa8, 0x7e, 0xab, 0x75,
	0x67, 0x74, 0xbe, 0x9b, 0x87, 0x5c, 0x9f, 0xda, 0x36, 0x61, 0x9d, 0xc2, 0x78, 0xff, 0x85, 0xad,
	0x61, 0x58, 0xf4, 0xf9, 0x11, 0x61, 0xa5, 0x85, 0x2a, 0xaa, 0x17, 0x66, 0xea, 0xa2, 0xad, 0xfd,
	0xb5, 0xb3, 0x83, 0x32, 0xdc, 0xa7, 0xc7, 0xba, 0xed, 0x3a, 0xba, 0xe3, 0x09, 0x4b, 0x3f, 0xf4,
	0x84, 0x35, 0x06, 0xc0, 0x5d, 0x78, 0x30, 0xfe, 0xfb, 0x9a, 0x79, 0x49, 0x9a, 0x58, 0x00, 0xfd,
	0xaf, 0xc0, 0x29, 0x82, 0xd2, 0x8c, 0xdb, 0xe7, 0x42, 0x8c, 0xdb, 0xd7, 0x21, 0x6f, 0x0a, 0x26,
	0xbb, 0x47, 0xe4, 0x24, 0x54, 0x58, 0x6c, 0xdd, 0x1d, 0x9d, 0xef, 0x16, 0x60, 0x69, 0x48, 0x25,
	0xed, 0xb9, 0xa4, 0xb3, 0x14, 0x6c, 0xb7, 0xc9, 0xc9, 0x35, 0x38, 0x7d, 0x0b, 0x95, 0x84, 0xd3,
	0x0b, 0x92, 0xf9, 0x8d, 0x7e, 0x43, 0xf1, 0x28, 0x13, 0x46, 0xdb, 0xd4, 0xbf, 0x51, 0xaf, 0x26,
	0x3c, 0x54, 0x7b, 0x0d, 0x60, 0xe6, 0xb7, 0xfb, 0x03, 0xc5, 0x1a, 0x09, 0xbb, 0xaf, 0x82, 0xf3,
	0x5a, 0x03, 0x0a, 0xbd, 0x01, 0xb3, 0x5d, 0x32, 0xb9, 0xc8, 0x85, 0x59, 0xc7, 0xf9, 0x68, 0xff,
	0x5a, 0x2e, 0xb2, 0x05, 0xeb, 0x6a, 0xcb, 0x11, 0xd0, 0xdc, 0xa6, 0x9f, 0xfe, 0xca, 0xc3, 0xed,
	0xf0, 0xba, 0x7c, 0x02, 0x2d, 0x25, 0xbb, 0x1b, 0x7a, 0xf4, 0x7e, 0xe8, 0xca, 0x78, 0x57, 0x8a,
	0x17, 0x25, 0x52, 0x70, 0x26, 0xc9, 0x1b, 0x4e, 0x6d, 0xdc, 0xfc, 0xfa, 0xfb, 0xcf, 0xcf, 0x85,
	0x1a, 0xae, 0x1a, 0xc3, 0x3d, 0x23, 0x50, 0x35, 0x82, 0x62, 0x23, 0x79, 0x7e, 0x1f, 0x35, 0xb4,
	0x2f, 0x08, 0x8a, 0xa9, 0x61, 0xdd, 0x9c, 0x91, 0x4f, 0x2b, 0x52, 0x00, 0x3c, 0x09, 0x01, 0x76,
	0x30, 0xbe, 0x0c, 0x90, 0xd6, 0x21, 0x40, 0xf8, 0x08, 0xcb, 0xc9, 0x30, 0x57, 0x15, 0xee, 0xe3,
	0x0a, 0x85, 0x76, 0x23, 0xd4, 0xde, 0xc2, 0xeb, 0xe9, 0xe6, 0xe3, 0xe3, 0x81, 0xf0, 0x67, 0x58,
	0x49, 0x4d, 0xaf, 0xd2, 0x79, 0x96, 0xf8, 0xe3, 0x50, 0x7c, 0x1b, 0x6f, 0xa8, 0x8c, 0x5f, 0x92,
	0x0f, 0x46, 0x9f, 0x1a, 0xee, 0xcd, 0x2c, 0xef, 0x6d, 0xea, 0x5f, 0x6d, 0xf4, 0x69, 0x1d, 0x02,
	0x84, 0x53, 0x04, 0xab, 0x8a, 0x50, 0xd7, 0xb2, 0xa7, 0xa0, 0xc6, 0x30, 0x42, 0x8c, 0x47, 0x78,
	0x2b, 0x73, 0x10, 0x53, 0x20, 0x8a, 0xe4, 0xd7, 0xb2, 0xa6, 0x11, 0x96, 0x5d, 0x0d, 0x24, 0xbd,
	0x47, 0x00, 0xf2, 0x1d, 0x41, 0x49, 0x99, 0xf9, 0x9d, 0xec, 0x99, 0xfc, 0x0b, 0x66, 0x2f, 0x84,
	0x69, 0xe2, 0xed, 0xcc, 0xa9, 0x4c, 0x70, 0x5a, 0xe5, 0xb3, 0x83, 0x55, 0x28, 0x4e, 0xbf, 0x20,
	0x92, 0x78, 0x43, 0x6a, 0x11, 0xd9, 0xcb, 0x85, 0x1f, 0x00, 0xcf, 0xfe, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x9b, 0xd1, 0x65, 0x88, 0x73, 0x08, 0x00, 0x00,
}
