// Code generated by protoc-gen-go.
// source: contacts.proto
// DO NOT EDIT!

package dialog

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/wrappers"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/gogo/protobuf/types"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Phone for import
type PhoneToImport struct {
	// / phone number for import in international format
	PhoneNumber int64 `protobuf:"varint,1,opt,name=phone_number,json=phoneNumber" json:"phone_number,omitempty"`
	// / optional name for contact
	Name *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *PhoneToImport) Reset()                    { *m = PhoneToImport{} }
func (m *PhoneToImport) String() string            { return proto.CompactTextString(m) }
func (*PhoneToImport) ProtoMessage()               {}
func (*PhoneToImport) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *PhoneToImport) GetPhoneNumber() int64 {
	if m != nil {
		return m.PhoneNumber
	}
	return 0
}

func (m *PhoneToImport) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

// Email for import
type EmailToImport struct {
	// / email for importing
	Email string `protobuf:"bytes,1,opt,name=email" json:"email,omitempty"`
	// / optional name for contact
	Name *google_protobuf.StringValue `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *EmailToImport) Reset()                    { *m = EmailToImport{} }
func (m *EmailToImport) String() string            { return proto.CompactTextString(m) }
func (*EmailToImport) ProtoMessage()               {}
func (*EmailToImport) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *EmailToImport) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *EmailToImport) GetName() *google_protobuf.StringValue {
	if m != nil {
		return m.Name
	}
	return nil
}

// Importing phones and emails for building contact list
// Maximum amount of items for import per method call equals to 100.
type RequestImportContacts struct {
	Phones []*PhoneToImport `protobuf:"bytes,1,rep,name=phones" json:"phones,omitempty"`
	Emails []*EmailToImport `protobuf:"bytes,2,rep,name=emails" json:"emails,omitempty"`
	// / Optimizations drops some info from response to decrease traffic and latency
	Optimizations []UpdateOptimization `protobuf:"varint,3,rep,packed,name=optimizations,enum=dialog.UpdateOptimization" json:"optimizations,omitempty"`
}

func (m *RequestImportContacts) Reset()                    { *m = RequestImportContacts{} }
func (m *RequestImportContacts) String() string            { return proto.CompactTextString(m) }
func (*RequestImportContacts) ProtoMessage()               {}
func (*RequestImportContacts) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *RequestImportContacts) GetPhones() []*PhoneToImport {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *RequestImportContacts) GetEmails() []*EmailToImport {
	if m != nil {
		return m.Emails
	}
	return nil
}

func (m *RequestImportContacts) GetOptimizations() []UpdateOptimization {
	if m != nil {
		return m.Optimizations
	}
	return nil
}

type ResponseImportContacts struct {
	// / Registered contacts
	Users []*User `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	Seq   int32   `protobuf:"varint,2,opt,name=seq" json:"seq,omitempty"`
	// / Server state related to current client, used by server only
	State []byte `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	// / Optimizations drops some info from response to decrease traffic and latency
	UserPeers []*UserOutPeer `protobuf:"bytes,4,rep,name=user_peers,json=userPeers" json:"user_peers,omitempty"`
}

func (m *ResponseImportContacts) Reset()                    { *m = ResponseImportContacts{} }
func (m *ResponseImportContacts) String() string            { return proto.CompactTextString(m) }
func (*ResponseImportContacts) ProtoMessage()               {}
func (*ResponseImportContacts) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{3} }

func (m *ResponseImportContacts) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ResponseImportContacts) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *ResponseImportContacts) GetState() []byte {
	if m != nil {
		return m.State
	}
	return nil
}

func (m *ResponseImportContacts) GetUserPeers() []*UserOutPeer {
	if m != nil {
		return m.UserPeers
	}
	return nil
}

// Importing phones and emails for building contact list
// Import evaluated lazily, response does not contain any info
// Maximum amount of items for import per method call equals to 100.
type RequestDeferredImportContacts struct {
	Phones []*PhoneToImport `protobuf:"bytes,1,rep,name=phones" json:"phones,omitempty"`
	Emails []*EmailToImport `protobuf:"bytes,2,rep,name=emails" json:"emails,omitempty"`
}

func (m *RequestDeferredImportContacts) Reset()                    { *m = RequestDeferredImportContacts{} }
func (m *RequestDeferredImportContacts) String() string            { return proto.CompactTextString(m) }
func (*RequestDeferredImportContacts) ProtoMessage()               {}
func (*RequestDeferredImportContacts) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{4} }

func (m *RequestDeferredImportContacts) GetPhones() []*PhoneToImport {
	if m != nil {
		return m.Phones
	}
	return nil
}

func (m *RequestDeferredImportContacts) GetEmails() []*EmailToImport {
	if m != nil {
		return m.Emails
	}
	return nil
}

type ResponseDeferredImportContacts struct {
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
}

func (m *ResponseDeferredImportContacts) Reset()                    { *m = ResponseDeferredImportContacts{} }
func (m *ResponseDeferredImportContacts) String() string            { return proto.CompactTextString(m) }
func (*ResponseDeferredImportContacts) ProtoMessage()               {}
func (*ResponseDeferredImportContacts) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{5} }

func (m *ResponseDeferredImportContacts) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

// Getting current contact list
// SHA256 hash of list of a comma-separated list of contact UIDs in ascending
// order may be passed in contactsHash parameter.
// If the contact list was not changed, isNotChanged will be true.
type RequestGetContacts struct {
	ContactsHash  string               `protobuf:"bytes,1,opt,name=contacts_hash,json=contactsHash" json:"contacts_hash,omitempty"`
	Optimizations []UpdateOptimization `protobuf:"varint,2,rep,packed,name=optimizations,enum=dialog.UpdateOptimization" json:"optimizations,omitempty"`
}

func (m *RequestGetContacts) Reset()                    { *m = RequestGetContacts{} }
func (m *RequestGetContacts) String() string            { return proto.CompactTextString(m) }
func (*RequestGetContacts) ProtoMessage()               {}
func (*RequestGetContacts) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{6} }

func (m *RequestGetContacts) GetContactsHash() string {
	if m != nil {
		return m.ContactsHash
	}
	return ""
}

func (m *RequestGetContacts) GetOptimizations() []UpdateOptimization {
	if m != nil {
		return m.Optimizations
	}
	return nil
}

type ResponseGetContacts struct {
	Users        []*User        `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	IsNotChanged bool           `protobuf:"varint,2,opt,name=is_not_changed,json=isNotChanged" json:"is_not_changed,omitempty"`
	UserPeers    []*UserOutPeer `protobuf:"bytes,3,rep,name=user_peers,json=userPeers" json:"user_peers,omitempty"`
}

func (m *ResponseGetContacts) Reset()                    { *m = ResponseGetContacts{} }
func (m *ResponseGetContacts) String() string            { return proto.CompactTextString(m) }
func (*ResponseGetContacts) ProtoMessage()               {}
func (*ResponseGetContacts) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{7} }

func (m *ResponseGetContacts) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ResponseGetContacts) GetIsNotChanged() bool {
	if m != nil {
		return m.IsNotChanged
	}
	return false
}

func (m *ResponseGetContacts) GetUserPeers() []*UserOutPeer {
	if m != nil {
		return m.UserPeers
	}
	return nil
}

// Removing contact from contact list
type RequestRemoveContact struct {
	Uid        int32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	AccessHash int64 `protobuf:"varint,2,opt,name=access_hash,json=accessHash" json:"access_hash,omitempty"`
}

func (m *RequestRemoveContact) Reset()                    { *m = RequestRemoveContact{} }
func (m *RequestRemoveContact) String() string            { return proto.CompactTextString(m) }
func (*RequestRemoveContact) ProtoMessage()               {}
func (*RequestRemoveContact) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{8} }

func (m *RequestRemoveContact) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RequestRemoveContact) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

// Adding contact to contact list
type RequestAddContact struct {
	Uid        int32 `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	AccessHash int64 `protobuf:"varint,2,opt,name=access_hash,json=accessHash" json:"access_hash,omitempty"`
}

func (m *RequestAddContact) Reset()                    { *m = RequestAddContact{} }
func (m *RequestAddContact) String() string            { return proto.CompactTextString(m) }
func (*RequestAddContact) ProtoMessage()               {}
func (*RequestAddContact) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{9} }

func (m *RequestAddContact) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *RequestAddContact) GetAccessHash() int64 {
	if m != nil {
		return m.AccessHash
	}
	return 0
}

// Searching contacts by user's query
type RequestSearchContacts struct {
	Request       string               `protobuf:"bytes,1,opt,name=request" json:"request,omitempty"`
	Optimizations []UpdateOptimization `protobuf:"varint,2,rep,packed,name=optimizations,enum=dialog.UpdateOptimization" json:"optimizations,omitempty"`
}

func (m *RequestSearchContacts) Reset()                    { *m = RequestSearchContacts{} }
func (m *RequestSearchContacts) String() string            { return proto.CompactTextString(m) }
func (*RequestSearchContacts) ProtoMessage()               {}
func (*RequestSearchContacts) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{10} }

func (m *RequestSearchContacts) GetRequest() string {
	if m != nil {
		return m.Request
	}
	return ""
}

func (m *RequestSearchContacts) GetOptimizations() []UpdateOptimization {
	if m != nil {
		return m.Optimizations
	}
	return nil
}

type ResponseSearchContacts struct {
	Users     []*User        `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
	UserPeers []*UserOutPeer `protobuf:"bytes,2,rep,name=user_peers,json=userPeers" json:"user_peers,omitempty"`
}

func (m *ResponseSearchContacts) Reset()                    { *m = ResponseSearchContacts{} }
func (m *ResponseSearchContacts) String() string            { return proto.CompactTextString(m) }
func (*ResponseSearchContacts) ProtoMessage()               {}
func (*ResponseSearchContacts) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{11} }

func (m *ResponseSearchContacts) GetUsers() []*User {
	if m != nil {
		return m.Users
	}
	return nil
}

func (m *ResponseSearchContacts) GetUserPeers() []*UserOutPeer {
	if m != nil {
		return m.UserPeers
	}
	return nil
}

// Update about contact registration
type UpdateContactRegistered struct {
	Uid      int32      `protobuf:"varint,1,opt,name=uid" json:"uid,omitempty"`
	IsSilent bool       `protobuf:"varint,2,opt,name=is_silent,json=isSilent" json:"is_silent,omitempty"`
	Date     int64      `protobuf:"varint,3,opt,name=date" json:"date,omitempty"`
	Mid      *UUIDValue `protobuf:"bytes,5,opt,name=mid" json:"mid,omitempty"`
}

func (m *UpdateContactRegistered) Reset()                    { *m = UpdateContactRegistered{} }
func (m *UpdateContactRegistered) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactRegistered) ProtoMessage()               {}
func (*UpdateContactRegistered) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{12} }

func (m *UpdateContactRegistered) GetUid() int32 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *UpdateContactRegistered) GetIsSilent() bool {
	if m != nil {
		return m.IsSilent
	}
	return false
}

func (m *UpdateContactRegistered) GetDate() int64 {
	if m != nil {
		return m.Date
	}
	return 0
}

func (m *UpdateContactRegistered) GetMid() *UUIDValue {
	if m != nil {
		return m.Mid
	}
	return nil
}

// Update about contacts added
type UpdateContactsAdded struct {
	// / User ids of the registered contacts
	Uids []int32 `protobuf:"varint,1,rep,packed,name=uids" json:"uids,omitempty"`
	// / Id of the task that finished
	TaskId *google_protobuf.StringValue `protobuf:"bytes,4,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
}

func (m *UpdateContactsAdded) Reset()                    { *m = UpdateContactsAdded{} }
func (m *UpdateContactsAdded) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactsAdded) ProtoMessage()               {}
func (*UpdateContactsAdded) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{13} }

func (m *UpdateContactsAdded) GetUids() []int32 {
	if m != nil {
		return m.Uids
	}
	return nil
}

func (m *UpdateContactsAdded) GetTaskId() *google_protobuf.StringValue {
	if m != nil {
		return m.TaskId
	}
	return nil
}

// Update about suspending task - normally it should be ignored
type UpdateContactsAddTaskSuspended struct {
	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId" json:"task_id,omitempty"`
}

func (m *UpdateContactsAddTaskSuspended) Reset()                    { *m = UpdateContactsAddTaskSuspended{} }
func (m *UpdateContactsAddTaskSuspended) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactsAddTaskSuspended) ProtoMessage()               {}
func (*UpdateContactsAddTaskSuspended) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{14} }

func (m *UpdateContactsAddTaskSuspended) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

// Update about contacts removed
type UpdateContactsRemoved struct {
	Uids []int32 `protobuf:"varint,1,rep,packed,name=uids" json:"uids,omitempty"`
}

func (m *UpdateContactsRemoved) Reset()                    { *m = UpdateContactsRemoved{} }
func (m *UpdateContactsRemoved) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactsRemoved) ProtoMessage()               {}
func (*UpdateContactsRemoved) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{15} }

func (m *UpdateContactsRemoved) GetUids() []int32 {
	if m != nil {
		return m.Uids
	}
	return nil
}

func init() {
	proto.RegisterType((*PhoneToImport)(nil), "dialog.PhoneToImport")
	proto.RegisterType((*EmailToImport)(nil), "dialog.EmailToImport")
	proto.RegisterType((*RequestImportContacts)(nil), "dialog.RequestImportContacts")
	proto.RegisterType((*ResponseImportContacts)(nil), "dialog.ResponseImportContacts")
	proto.RegisterType((*RequestDeferredImportContacts)(nil), "dialog.RequestDeferredImportContacts")
	proto.RegisterType((*ResponseDeferredImportContacts)(nil), "dialog.ResponseDeferredImportContacts")
	proto.RegisterType((*RequestGetContacts)(nil), "dialog.RequestGetContacts")
	proto.RegisterType((*ResponseGetContacts)(nil), "dialog.ResponseGetContacts")
	proto.RegisterType((*RequestRemoveContact)(nil), "dialog.RequestRemoveContact")
	proto.RegisterType((*RequestAddContact)(nil), "dialog.RequestAddContact")
	proto.RegisterType((*RequestSearchContacts)(nil), "dialog.RequestSearchContacts")
	proto.RegisterType((*ResponseSearchContacts)(nil), "dialog.ResponseSearchContacts")
	proto.RegisterType((*UpdateContactRegistered)(nil), "dialog.UpdateContactRegistered")
	proto.RegisterType((*UpdateContactsAdded)(nil), "dialog.UpdateContactsAdded")
	proto.RegisterType((*UpdateContactsAddTaskSuspended)(nil), "dialog.UpdateContactsAddTaskSuspended")
	proto.RegisterType((*UpdateContactsRemoved)(nil), "dialog.UpdateContactsRemoved")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Contacts service

type ContactsClient interface {
	// / Import contacts and wait while query is not finished
	ImportContacts(ctx context.Context, in *RequestImportContacts, opts ...grpc.CallOption) (*ResponseImportContacts, error)
	// / Same as above, but without waiting response
	DeferredImportContacts(ctx context.Context, in *RequestDeferredImportContacts, opts ...grpc.CallOption) (*ResponseDeferredImportContacts, error)
	GetContacts(ctx context.Context, in *RequestGetContacts, opts ...grpc.CallOption) (*ResponseGetContacts, error)
	RemoveContact(ctx context.Context, in *RequestRemoveContact, opts ...grpc.CallOption) (*ResponseSeq, error)
	AddContact(ctx context.Context, in *RequestAddContact, opts ...grpc.CallOption) (*ResponseSeq, error)
	// / Search contacts by query string
	SearchContacts(ctx context.Context, in *RequestSearchContacts, opts ...grpc.CallOption) (*ResponseSearchContacts, error)
}

type contactsClient struct {
	cc *grpc.ClientConn
}

func NewContactsClient(cc *grpc.ClientConn) ContactsClient {
	return &contactsClient{cc}
}

func (c *contactsClient) ImportContacts(ctx context.Context, in *RequestImportContacts, opts ...grpc.CallOption) (*ResponseImportContacts, error) {
	out := new(ResponseImportContacts)
	err := grpc.Invoke(ctx, "/dialog.Contacts/ImportContacts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) DeferredImportContacts(ctx context.Context, in *RequestDeferredImportContacts, opts ...grpc.CallOption) (*ResponseDeferredImportContacts, error) {
	out := new(ResponseDeferredImportContacts)
	err := grpc.Invoke(ctx, "/dialog.Contacts/DeferredImportContacts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) GetContacts(ctx context.Context, in *RequestGetContacts, opts ...grpc.CallOption) (*ResponseGetContacts, error) {
	out := new(ResponseGetContacts)
	err := grpc.Invoke(ctx, "/dialog.Contacts/GetContacts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) RemoveContact(ctx context.Context, in *RequestRemoveContact, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Contacts/RemoveContact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) AddContact(ctx context.Context, in *RequestAddContact, opts ...grpc.CallOption) (*ResponseSeq, error) {
	out := new(ResponseSeq)
	err := grpc.Invoke(ctx, "/dialog.Contacts/AddContact", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactsClient) SearchContacts(ctx context.Context, in *RequestSearchContacts, opts ...grpc.CallOption) (*ResponseSearchContacts, error) {
	out := new(ResponseSearchContacts)
	err := grpc.Invoke(ctx, "/dialog.Contacts/SearchContacts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Contacts service

type ContactsServer interface {
	// / Import contacts and wait while query is not finished
	ImportContacts(context.Context, *RequestImportContacts) (*ResponseImportContacts, error)
	// / Same as above, but without waiting response
	DeferredImportContacts(context.Context, *RequestDeferredImportContacts) (*ResponseDeferredImportContacts, error)
	GetContacts(context.Context, *RequestGetContacts) (*ResponseGetContacts, error)
	RemoveContact(context.Context, *RequestRemoveContact) (*ResponseSeq, error)
	AddContact(context.Context, *RequestAddContact) (*ResponseSeq, error)
	// / Search contacts by query string
	SearchContacts(context.Context, *RequestSearchContacts) (*ResponseSearchContacts, error)
}

func RegisterContactsServer(s *grpc.Server, srv ContactsServer) {
	s.RegisterService(&_Contacts_serviceDesc, srv)
}

func _Contacts_ImportContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestImportContacts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).ImportContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Contacts/ImportContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).ImportContacts(ctx, req.(*RequestImportContacts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_DeferredImportContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestDeferredImportContacts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).DeferredImportContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Contacts/DeferredImportContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).DeferredImportContacts(ctx, req.(*RequestDeferredImportContacts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_GetContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetContacts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).GetContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Contacts/GetContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).GetContacts(ctx, req.(*RequestGetContacts))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_RemoveContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestRemoveContact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).RemoveContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Contacts/RemoveContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).RemoveContact(ctx, req.(*RequestRemoveContact))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_AddContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAddContact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).AddContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Contacts/AddContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).AddContact(ctx, req.(*RequestAddContact))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts_SearchContacts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSearchContacts)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactsServer).SearchContacts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dialog.Contacts/SearchContacts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactsServer).SearchContacts(ctx, req.(*RequestSearchContacts))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contacts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dialog.Contacts",
	HandlerType: (*ContactsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ImportContacts",
			Handler:    _Contacts_ImportContacts_Handler,
		},
		{
			MethodName: "DeferredImportContacts",
			Handler:    _Contacts_DeferredImportContacts_Handler,
		},
		{
			MethodName: "GetContacts",
			Handler:    _Contacts_GetContacts_Handler,
		},
		{
			MethodName: "RemoveContact",
			Handler:    _Contacts_RemoveContact_Handler,
		},
		{
			MethodName: "AddContact",
			Handler:    _Contacts_AddContact_Handler,
		},
		{
			MethodName: "SearchContacts",
			Handler:    _Contacts_SearchContacts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contacts.proto",
}

func init() { proto.RegisterFile("contacts.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 1043 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x05, 0x25, 0xcb, 0x8f, 0x2b, 0xc9, 0xa8, 0x29, 0x3f, 0x64, 0x55, 0x76, 0x14, 0x36, 0x0f,
	0xc1, 0x45, 0x24, 0x57, 0xd9, 0x09, 0x05, 0xdc, 0x24, 0x2e, 0x12, 0x6f, 0xe2, 0x80, 0x4a, 0x8a,
	0xee, 0x84, 0x11, 0x79, 0x2d, 0x0d, 0xc2, 0x97, 0x39, 0xa4, 0x0b, 0x14, 0x41, 0x81, 0x1a, 0x5d,
	0x65, 0xd3, 0x45, 0xb7, 0x5d, 0xf4, 0x1b, 0xba, 0x09, 0xba, 0xe8, 0x57, 0xf4, 0x0f, 0x8c, 0xae,
	0xb2, 0xe8, 0x37, 0x14, 0x9c, 0x21, 0x25, 0x72, 0x28, 0xd7, 0x0e, 0x9a, 0xa2, 0x59, 0x89, 0x9c,
	0x73, 0xef, 0x9c, 0x73, 0xe7, 0x1e, 0xde, 0x11, 0xac, 0x1a, 0xae, 0x13, 0x10, 0x23, 0x60, 0x1d,
	0xcf, 0x77, 0x03, 0x57, 0x5d, 0x34, 0x29, 0xb1, 0xdc, 0x71, 0x63, 0x77, 0xec, 0xba, 0x63, 0x0b,
	0xbb, 0x7c, 0x75, 0x14, 0x9e, 0x74, 0xbf, 0xf1, 0x89, 0xe7, 0xa1, 0x1f, 0xc7, 0x35, 0x9a, 0x31,
	0x4e, 0x3c, 0xda, 0x25, 0x8e, 0xe3, 0x06, 0x24, 0xa0, 0xae, 0x93, 0xa0, 0x6b, 0x26, 0x9e, 0x50,
	0x87, 0xa6, 0x97, 0x6a, 0x36, 0x65, 0x06, 0x5a, 0x16, 0x71, 0xd0, 0x0d, 0x93, 0xc5, 0xb2, 0x87,
	0xb3, 0x2d, 0xcb, 0x21, 0x9b, 0xbd, 0x6c, 0x30, 0x83, 0x58, 0xc4, 0x1b, 0x75, 0xe3, 0x5f, 0xb1,
	0xac, 0x9d, 0x2b, 0x50, 0x7d, 0x36, 0x71, 0x1d, 0x7c, 0xee, 0x1e, 0xd9, 0x9e, 0xeb, 0x07, 0xea,
	0x3e, 0x54, 0xbc, 0x68, 0x61, 0xe8, 0x84, 0xf6, 0x08, 0xfd, 0xba, 0xd2, 0x52, 0xda, 0xc5, 0x87,
	0xd5, 0xd7, 0x6f, 0xf7, 0x57, 0x60, 0xe9, 0x8c, 0x32, 0x3a, 0xb2, 0x50, 0x2f, 0xf3, 0x90, 0xa7,
	0x3c, 0x42, 0x3d, 0x80, 0x05, 0x87, 0xd8, 0x58, 0x2f, 0xb4, 0x94, 0x76, 0xb9, 0xd7, 0xec, 0x88,
	0x4a, 0x3a, 0x49, 0xa5, 0x9d, 0x41, 0xe0, 0x53, 0x67, 0xfc, 0x15, 0xb1, 0x42, 0x94, 0xf7, 0xe1,
	0x89, 0x5a, 0x08, 0xd5, 0x2f, 0x6d, 0x42, 0xad, 0xa9, 0x86, 0x4f, 0xa0, 0x84, 0xd1, 0x02, 0x27,
	0x5f, 0x91, 0x93, 0x04, 0xf6, 0xef, 0x69, 0xcf, 0x0b, 0xb0, 0xa1, 0xe3, 0x69, 0x88, 0x2c, 0x10,
	0xbc, 0x8f, 0xe2, 0xd6, 0xa9, 0x9f, 0xc3, 0x22, 0x2f, 0x90, 0xd5, 0x95, 0x56, 0xb1, 0x5d, 0xee,
	0x6d, 0x74, 0x44, 0x17, 0x3b, 0x99, 0xa3, 0x4a, 0x76, 0x35, 0x5c, 0xdb, 0x23, 0x46, 0xa0, 0xc7,
	0x39, 0x51, 0x36, 0x57, 0xc8, 0xea, 0x85, 0x6c, 0x76, 0xa6, 0xc8, 0x5c, 0xb6, 0xc8, 0x51, 0x8f,
	0xa1, 0xea, 0x7a, 0x01, 0xb5, 0xe9, 0xb7, 0xc2, 0x01, 0xf5, 0x62, 0xab, 0xd8, 0x5e, 0xed, 0x35,
	0x92, 0x4d, 0x5e, 0x78, 0x26, 0x09, 0xf0, 0x38, 0x15, 0x22, 0x57, 0x97, 0xcd, 0xef, 0x37, 0x2f,
	0x0e, 0xb6, 0x61, 0x8b, 0xda, 0x1d, 0xd3, 0x1a, 0x77, 0xc6, 0xbe, 0x67, 0x74, 0x1e, 0xfb, 0x9e,
	0x11, 0x17, 0xae, 0xbd, 0x51, 0x60, 0x53, 0x47, 0xe6, 0xb9, 0x0e, 0x43, 0xe9, 0x14, 0x34, 0x28,
	0x71, 0x07, 0xc5, 0x87, 0x50, 0x99, 0x2a, 0x60, 0xe8, 0xeb, 0x02, 0x52, 0x3f, 0x82, 0x22, 0xc3,
	0x53, 0xde, 0x83, 0x92, 0x1e, 0x3d, 0xaa, 0xeb, 0x50, 0x62, 0x01, 0x09, 0xb0, 0x5e, 0x6c, 0x29,
	0xed, 0x8a, 0x2e, 0x5e, 0xd4, 0x1e, 0x40, 0x94, 0x30, 0xe4, 0xfe, 0xac, 0x2f, 0xf0, 0x0d, 0x6b,
	0xe9, 0x0d, 0x8f, 0xc3, 0xe0, 0x19, 0xa2, 0xaf, 0xaf, 0x44, 0x61, 0xd1, 0x13, 0xeb, 0xef, 0x5c,
	0x1c, 0x34, 0xa0, 0x9e, 0x17, 0x2e, 0xc4, 0x6a, 0xbf, 0x29, 0xb0, 0x13, 0x57, 0x71, 0x88, 0x27,
	0xe8, 0xfb, 0x68, 0x7e, 0x38, 0x6d, 0xbc, 0xe2, 0xd4, 0xbf, 0x86, 0xdd, 0xa4, 0x8e, 0x4b, 0xb4,
	0x6f, 0xc1, 0x52, 0x40, 0xd8, 0xcb, 0x21, 0x35, 0xc5, 0x47, 0xa0, 0x2f, 0x46, 0xaf, 0x47, 0xe6,
	0x55, 0xa7, 0xf2, 0x46, 0x01, 0x35, 0x66, 0x79, 0x8c, 0xb3, 0xed, 0x7a, 0x50, 0x4d, 0x06, 0xd3,
	0x70, 0x42, 0xd8, 0x64, 0xfe, 0x97, 0x55, 0x49, 0x62, 0x9e, 0x10, 0x36, 0xc9, 0x3b, 0xb1, 0xf0,
	0x9f, 0x3a, 0xf1, 0x57, 0x05, 0x6a, 0x49, 0x19, 0x69, 0xe9, 0xd7, 0xb1, 0xe1, 0x2d, 0x58, 0xa5,
	0x6c, 0xe8, 0xb8, 0xc1, 0xd0, 0x98, 0x10, 0x67, 0x8c, 0x26, 0x77, 0xe4, 0xb2, 0x5e, 0xa1, 0xec,
	0xa9, 0x1b, 0x3c, 0x12, 0x6b, 0x92, 0x09, 0x8b, 0xef, 0xc3, 0x84, 0x3f, 0x28, 0xb0, 0x1e, 0x17,
	0xa0, 0xa3, 0xed, 0x9e, 0x61, 0x2c, 0x5b, 0xbd, 0x01, 0xc5, 0x30, 0xee, 0x5d, 0x49, 0x3e, 0x96,
	0x08, 0x51, 0xef, 0x41, 0x99, 0x18, 0x06, 0xb2, 0xb8, 0x1f, 0x05, 0x3e, 0x66, 0x2b, 0xaf, 0xdf,
	0xee, 0x2f, 0xc3, 0xa2, 0x19, 0xe9, 0xf5, 0x75, 0x10, 0x01, 0x51, 0x33, 0xae, 0x38, 0xbb, 0xef,
	0x15, 0x58, 0x8b, 0x9f, 0x1f, 0x98, 0xe6, 0xff, 0xa3, 0xe1, 0x17, 0x65, 0x3a, 0x4e, 0x07, 0x48,
	0x7c, 0x63, 0x32, 0xed, 0xe0, 0x5d, 0x58, 0xf2, 0x05, 0x30, 0xdf, 0x76, 0x09, 0xaa, 0x7e, 0xf1,
	0xce, 0x8e, 0x7b, 0x37, 0x8b, 0xfd, 0x98, 0x1a, 0x76, 0x92, 0xc6, 0xeb, 0xb8, 0x2c, 0xeb, 0x9f,
	0xc2, 0xfb, 0xf0, 0xcf, 0xef, 0x0a, 0x6c, 0x89, 0xaa, 0x62, 0x25, 0x3a, 0x8e, 0x29, 0x0b, 0xd0,
	0x47, 0xf3, 0xea, 0xf6, 0xed, 0xc1, 0x0a, 0x65, 0x43, 0x46, 0x2d, 0x74, 0x02, 0x61, 0x78, 0x39,
	0x6c, 0x99, 0xb2, 0x01, 0x87, 0xd5, 0x9b, 0xb0, 0x60, 0x26, 0x53, 0x39, 0x77, 0x9d, 0x73, 0x48,
	0xed, 0x41, 0xd1, 0xa6, 0x66, 0xbd, 0xc4, 0xef, 0xd3, 0xb5, 0x69, 0x5d, 0x2f, 0x8e, 0x0e, 0xe7,
	0x5e, 0xa2, 0x51, 0xb0, 0xf6, 0x1d, 0xd4, 0x32, 0xf2, 0xd9, 0x03, 0xd3, 0x44, 0x33, 0x62, 0x0b,
	0xa9, 0x29, 0x0e, 0xb3, 0x24, 0x8f, 0x48, 0x0e, 0xa9, 0x87, 0xb3, 0x01, 0xb7, 0x70, 0xfd, 0x1b,
	0x7c, 0x3a, 0x66, 0xc5, 0x34, 0xd4, 0x9e, 0xc0, 0x6e, 0x8e, 0xff, 0x39, 0x61, 0x2f, 0x07, 0x21,
	0xf3, 0xd0, 0x89, 0xa4, 0xdc, 0x91, 0x06, 0xe9, 0x65, 0x3b, 0xf5, 0x61, 0x23, 0xbb, 0x93, 0xf8,
	0x9e, 0xaf, 0x53, 0x4b, 0xef, 0xaf, 0x12, 0x2c, 0x4f, 0x9d, 0xf4, 0x0a, 0x56, 0xa5, 0x59, 0xbe,
	0x93, 0x9c, 0xe5, 0xdc, 0x7f, 0x1b, 0x8d, 0xdd, 0x19, 0x3c, 0xef, 0x1e, 0xd6, 0x3e, 0x3d, 0xff,
	0xe3, 0xcf, 0x9f, 0x0a, 0xb7, 0xb5, 0x56, 0xf7, 0xec, 0xb3, 0x6e, 0x64, 0xa1, 0x6e, 0x02, 0x75,
	0xb3, 0x91, 0x7d, 0x65, 0x4f, 0xfd, 0x59, 0x81, 0xcd, 0x4b, 0xae, 0x94, 0xdb, 0x92, 0x8c, 0xf9,
	0x61, 0x8d, 0x3b, 0xb2, 0x9c, 0xf9, 0x71, 0xda, 0x7d, 0x2e, 0xeb, 0x9e, 0xd6, 0xce, 0xcb, 0x9a,
	0x9f, 0x11, 0xc9, 0xf3, 0xa0, 0x9c, 0x9e, 0xed, 0x0d, 0x49, 0x52, 0x0a, 0x6b, 0x7c, 0x2c, 0xeb,
	0x48, 0x81, 0x5a, 0x9b, 0x93, 0x6b, 0xda, 0x4e, 0x9e, 0x3c, 0x15, 0x16, 0x31, 0x3a, 0x50, 0xcd,
	0x4e, 0xe6, 0xa6, 0xc4, 0x99, 0x41, 0x1b, 0x35, 0x99, 0x75, 0x80, 0xa7, 0xda, 0x1e, 0x67, 0xbb,
	0xa5, 0xdd, 0xc8, 0xb3, 0x65, 0xb2, 0x23, 0x3e, 0x04, 0x48, 0x8d, 0xe0, 0x6d, 0x89, 0x6c, 0x06,
	0xcd, 0x67, 0xba, 0xcb, 0x99, 0x6e, 0x6a, 0xcd, 0x3c, 0xd3, 0x2c, 0x35, 0xa2, 0x79, 0x05, 0xab,
	0xd2, 0x04, 0x93, 0x5d, 0x96, 0x85, 0xf3, 0x2e, 0xcb, 0xe2, 0xff, 0xe4, 0xb2, 0x6c, 0x64, 0x5f,
	0xd9, 0x7b, 0xb8, 0x7d, 0x71, 0xb0, 0x09, 0xeb, 0xe9, 0xa9, 0xc6, 0xd0, 0x3f, 0xa3, 0x06, 0xb2,
	0xd1, 0x22, 0xff, 0x7c, 0xef, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x17, 0x2d, 0x1d, 0x09,
	0x0d, 0x00, 0x00,
}
