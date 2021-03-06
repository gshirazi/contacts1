// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contacts1/pkg/pb/service.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	contacts1/pkg/pb/service.proto

It has these top-level messages:
	VersionResponse
	Contact
	ReverseRequest
	ReverseResponse
	CreateContactRequest
	CreateContactResponse
	ReadContactRequest
	ReadContactResponse
	UpdateContactRequest
	UpdateContactResponse
	DeleteContactRequest
	DeleteContactResponse
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
import _ "github.com/infobloxopen/protoc-gen-gorm/options"

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

// TODO: Structure your own protobuf messages. Each protocol buffer message is a
// small logical record of information, containing a series of name-value pairs.
type VersionResponse struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *VersionResponse) Reset()                    { *m = VersionResponse{} }
func (m *VersionResponse) String() string            { return proto.CompactTextString(m) }
func (*VersionResponse) ProtoMessage()               {}
func (*VersionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *VersionResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

// Contact represents a particular person in an contacts list.
type Contact struct {
	// The id associated to the person
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// The first name of the person.
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName" json:"first_name,omitempty"`
	// The middle name of the person.
	MiddleName string `protobuf:"bytes,3,opt,name=middle_name,json=middleName" json:"middle_name,omitempty"`
	// The last name of the person.
	LastName string `protobuf:"bytes,4,opt,name=last_name,json=lastName" json:"last_name,omitempty"`
	// The person's primary email.
	Email string `protobuf:"bytes,5,opt,name=email" json:"email,omitempty"`
	// The home address of the person.
	HomeAddress string `protobuf:"bytes,6,opt,name=home_address,json=homeAddress" json:"home_address,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Contact) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Contact) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Contact) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *Contact) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *Contact) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Contact) GetHomeAddress() string {
	if m != nil {
		return m.HomeAddress
	}
	return ""
}

type ReverseRequest struct {
	Payload *Contact `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *ReverseRequest) Reset()                    { *m = ReverseRequest{} }
func (m *ReverseRequest) String() string            { return proto.CompactTextString(m) }
func (*ReverseRequest) ProtoMessage()               {}
func (*ReverseRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReverseRequest) GetPayload() *Contact {
	if m != nil {
		return m.Payload
	}
	return nil
}

type ReverseResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *ReverseResponse) Reset()                    { *m = ReverseResponse{} }
func (m *ReverseResponse) String() string            { return proto.CompactTextString(m) }
func (*ReverseResponse) ProtoMessage()               {}
func (*ReverseResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReverseResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type CreateContactRequest struct {
	Payload *Contact `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *CreateContactRequest) Reset()                    { *m = CreateContactRequest{} }
func (m *CreateContactRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateContactRequest) ProtoMessage()               {}
func (*CreateContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CreateContactRequest) GetPayload() *Contact {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CreateContactResponse struct {
	Result *Contact `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *CreateContactResponse) Reset()                    { *m = CreateContactResponse{} }
func (m *CreateContactResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateContactResponse) ProtoMessage()               {}
func (*CreateContactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CreateContactResponse) GetResult() *Contact {
	if m != nil {
		return m.Result
	}
	return nil
}

type ReadContactRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadContactRequest) Reset()                    { *m = ReadContactRequest{} }
func (m *ReadContactRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadContactRequest) ProtoMessage()               {}
func (*ReadContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ReadContactRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ReadContactResponse struct {
	Result *Contact `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *ReadContactResponse) Reset()                    { *m = ReadContactResponse{} }
func (m *ReadContactResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadContactResponse) ProtoMessage()               {}
func (*ReadContactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *ReadContactResponse) GetResult() *Contact {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateContactRequest struct {
	Payload *Contact `protobuf:"bytes,1,opt,name=payload" json:"payload,omitempty"`
}

func (m *UpdateContactRequest) Reset()                    { *m = UpdateContactRequest{} }
func (m *UpdateContactRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactRequest) ProtoMessage()               {}
func (*UpdateContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *UpdateContactRequest) GetPayload() *Contact {
	if m != nil {
		return m.Payload
	}
	return nil
}

type UpdateContactResponse struct {
	Result *Contact `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *UpdateContactResponse) Reset()                    { *m = UpdateContactResponse{} }
func (m *UpdateContactResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateContactResponse) ProtoMessage()               {}
func (*UpdateContactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *UpdateContactResponse) GetResult() *Contact {
	if m != nil {
		return m.Result
	}
	return nil
}

type DeleteContactRequest struct {
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteContactRequest) Reset()                    { *m = DeleteContactRequest{} }
func (m *DeleteContactRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteContactRequest) ProtoMessage()               {}
func (*DeleteContactRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteContactRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeleteContactResponse struct {
}

func (m *DeleteContactResponse) Reset()                    { *m = DeleteContactResponse{} }
func (m *DeleteContactResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteContactResponse) ProtoMessage()               {}
func (*DeleteContactResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func init() {
	proto.RegisterType((*VersionResponse)(nil), "service.VersionResponse")
	proto.RegisterType((*Contact)(nil), "service.Contact")
	proto.RegisterType((*ReverseRequest)(nil), "service.ReverseRequest")
	proto.RegisterType((*ReverseResponse)(nil), "service.ReverseResponse")
	proto.RegisterType((*CreateContactRequest)(nil), "service.CreateContactRequest")
	proto.RegisterType((*CreateContactResponse)(nil), "service.CreateContactResponse")
	proto.RegisterType((*ReadContactRequest)(nil), "service.ReadContactRequest")
	proto.RegisterType((*ReadContactResponse)(nil), "service.ReadContactResponse")
	proto.RegisterType((*UpdateContactRequest)(nil), "service.UpdateContactRequest")
	proto.RegisterType((*UpdateContactResponse)(nil), "service.UpdateContactResponse")
	proto.RegisterType((*DeleteContactRequest)(nil), "service.DeleteContactRequest")
	proto.RegisterType((*DeleteContactResponse)(nil), "service.DeleteContactResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Contacts1 service

type Contacts1Client interface {
	GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error)
	// Use this method to create a contact information.
	Create(ctx context.Context, in *CreateContactRequest, opts ...grpc.CallOption) (*CreateContactResponse, error)
	// Use this method to read a contact information by identifier.
	Read(ctx context.Context, in *ReadContactRequest, opts ...grpc.CallOption) (*ReadContactResponse, error)
	// Use this method to update a contact information.
	Update(ctx context.Context, in *UpdateContactRequest, opts ...grpc.CallOption) (*UpdateContactResponse, error)
	// Use this method to delete a particular contact.
	Delete(ctx context.Context, in *DeleteContactRequest, opts ...grpc.CallOption) (*DeleteContactResponse, error)
}

type contacts1Client struct {
	cc *grpc.ClientConn
}

func NewContacts1Client(cc *grpc.ClientConn) Contacts1Client {
	return &contacts1Client{cc}
}

func (c *contacts1Client) GetVersion(ctx context.Context, in *google_protobuf.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := grpc.Invoke(ctx, "/service.Contacts1/GetVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacts1Client) Reverse(ctx context.Context, in *ReverseRequest, opts ...grpc.CallOption) (*ReverseResponse, error) {
	out := new(ReverseResponse)
	err := grpc.Invoke(ctx, "/service.Contacts1/Reverse", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacts1Client) Create(ctx context.Context, in *CreateContactRequest, opts ...grpc.CallOption) (*CreateContactResponse, error) {
	out := new(CreateContactResponse)
	err := grpc.Invoke(ctx, "/service.Contacts1/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacts1Client) Read(ctx context.Context, in *ReadContactRequest, opts ...grpc.CallOption) (*ReadContactResponse, error) {
	out := new(ReadContactResponse)
	err := grpc.Invoke(ctx, "/service.Contacts1/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacts1Client) Update(ctx context.Context, in *UpdateContactRequest, opts ...grpc.CallOption) (*UpdateContactResponse, error) {
	out := new(UpdateContactResponse)
	err := grpc.Invoke(ctx, "/service.Contacts1/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contacts1Client) Delete(ctx context.Context, in *DeleteContactRequest, opts ...grpc.CallOption) (*DeleteContactResponse, error) {
	out := new(DeleteContactResponse)
	err := grpc.Invoke(ctx, "/service.Contacts1/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Contacts1 service

type Contacts1Server interface {
	GetVersion(context.Context, *google_protobuf.Empty) (*VersionResponse, error)
	Reverse(context.Context, *ReverseRequest) (*ReverseResponse, error)
	// Use this method to create a contact information.
	Create(context.Context, *CreateContactRequest) (*CreateContactResponse, error)
	// Use this method to read a contact information by identifier.
	Read(context.Context, *ReadContactRequest) (*ReadContactResponse, error)
	// Use this method to update a contact information.
	Update(context.Context, *UpdateContactRequest) (*UpdateContactResponse, error)
	// Use this method to delete a particular contact.
	Delete(context.Context, *DeleteContactRequest) (*DeleteContactResponse, error)
}

func RegisterContacts1Server(s *grpc.Server, srv Contacts1Server) {
	s.RegisterService(&_Contacts1_serviceDesc, srv)
}

func _Contacts1_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(google_protobuf.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Contacts1Server).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Contacts1/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Contacts1Server).GetVersion(ctx, req.(*google_protobuf.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts1_Reverse_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReverseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Contacts1Server).Reverse(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Contacts1/Reverse",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Contacts1Server).Reverse(ctx, req.(*ReverseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Contacts1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Contacts1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Contacts1Server).Create(ctx, req.(*CreateContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts1_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Contacts1Server).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Contacts1/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Contacts1Server).Read(ctx, req.(*ReadContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts1_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Contacts1Server).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Contacts1/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Contacts1Server).Update(ctx, req.(*UpdateContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Contacts1_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(Contacts1Server).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Contacts1/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(Contacts1Server).Delete(ctx, req.(*DeleteContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Contacts1_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Contacts1",
	HandlerType: (*Contacts1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVersion",
			Handler:    _Contacts1_GetVersion_Handler,
		},
		{
			MethodName: "Reverse",
			Handler:    _Contacts1_Reverse_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Contacts1_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Contacts1_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Contacts1_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Contacts1_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contacts1/pkg/pb/service.proto",
}

func init() { proto.RegisterFile("contacts1/pkg/pb/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdf, 0x4e, 0xd4, 0x4e,
	0x14, 0xce, 0x2e, 0xd0, 0xdd, 0x3d, 0xfc, 0x02, 0x64, 0x58, 0x96, 0xa5, 0xfc, 0xf9, 0x49, 0x35,
	0x06, 0x51, 0xda, 0x88, 0x77, 0x60, 0x62, 0xf8, 0x63, 0x4c, 0xbc, 0x30, 0xa6, 0x89, 0x1a, 0xb9,
	0x21, 0xb3, 0xdb, 0xb3, 0x65, 0x62, 0xdb, 0xa9, 0xed, 0x2c, 0x4a, 0x0c, 0x37, 0xbc, 0x02, 0x0f,
	0x63, 0xe0, 0x35, 0x7c, 0x05, 0x1f, 0xc4, 0xec, 0xcc, 0xb4, 0x74, 0x4b, 0x8d, 0x41, 0xef, 0x76,
	0xce, 0x77, 0xe6, 0xfb, 0xce, 0x39, 0xfd, 0xe6, 0x2c, 0xac, 0xf5, 0x79, 0x24, 0x68, 0x5f, 0xa4,
	0x4f, 0x9d, 0xf8, 0x93, 0xef, 0xc4, 0x3d, 0x27, 0xc5, 0xe4, 0x94, 0xf5, 0xd1, 0x8e, 0x13, 0x2e,
	0x38, 0x69, 0xe8, 0xa3, 0xb9, 0xec, 0x73, 0xee, 0x07, 0xe8, 0xc8, 0x70, 0x6f, 0x38, 0x70, 0x30,
	0x8c, 0xc5, 0x99, 0xca, 0x32, 0x57, 0x34, 0x48, 0x63, 0xe6, 0xd0, 0x28, 0xe2, 0x82, 0x0a, 0xc6,
	0xa3, 0x54, 0xa3, 0xbb, 0x3e, 0x13, 0x27, 0xc3, 0x9e, 0xdd, 0xe7, 0xa1, 0x13, 0x9c, 0x0d, 0x84,
	0xe2, 0xe8, 0x6f, 0xf9, 0x18, 0x6d, 0x9d, 0xd2, 0x80, 0x79, 0x54, 0xa0, 0x73, 0xeb, 0x87, 0xbe,
	0xfc, 0xa4, 0x90, 0x9c, 0x7e, 0xa1, 0xbe, 0x8f, 0x89, 0xc3, 0x63, 0x49, 0x5f, 0x21, 0xb5, 0x53,
	0x90, 0x62, 0xd1, 0x80, 0xf7, 0x02, 0xfe, 0x95, 0xc7, 0x18, 0x15, 0x25, 0x7d, 0x9e, 0x84, 0x39,
	0xc5, 0xe8, 0xa0, 0xee, 0x5a, 0x8f, 0x61, 0xf6, 0x3d, 0x26, 0x29, 0xe3, 0x91, 0x8b, 0x69, 0xcc,
	0xa3, 0x14, 0x49, 0x17, 0x1a, 0xa7, 0x2a, 0xd4, 0xad, 0xdd, 0xab, 0x6d, 0xb4, 0xdc, 0xec, 0x68,
	0x7d, 0xaf, 0x41, 0xe3, 0x40, 0x8d, 0x8e, 0xcc, 0x40, 0x9d, 0x79, 0x32, 0x61, 0xc2, 0xad, 0x33,
	0x8f, 0xac, 0x02, 0x0c, 0x58, 0x92, 0x8a, 0xe3, 0x88, 0x86, 0xd8, 0xad, 0xcb, 0x8b, 0x2d, 0x19,
	0x79, 0x43, 0x43, 0x24, 0xff, 0xc3, 0x74, 0xc8, 0x3c, 0x2f, 0x40, 0x85, 0x4f, 0x48, 0x1c, 0x54,
	0x48, 0x26, 0x2c, 0x43, 0x2b, 0xa0, 0xd9, 0xf5, 0x49, 0x09, 0x37, 0x47, 0x01, 0x09, 0xb6, 0x61,
	0x0a, 0x43, 0xca, 0x82, 0xee, 0x94, 0x04, 0xd4, 0x81, 0xac, 0xc3, 0x7f, 0x27, 0x3c, 0xc4, 0x63,
	0xea, 0x79, 0x09, 0xa6, 0x69, 0xd7, 0x90, 0xe0, 0xf4, 0x28, 0xb6, 0xa7, 0x42, 0x3b, 0xc6, 0xf5,
	0xd5, 0x52, 0xbd, 0x59, 0xb3, 0x9e, 0xc3, 0x8c, 0x8b, 0xa3, 0x36, 0xd0, 0xc5, 0xcf, 0x43, 0x4c,
	0x05, 0xd9, 0x84, 0x46, 0x4c, 0xcf, 0x02, 0x4e, 0x55, 0x13, 0xd3, 0xdb, 0x73, 0x76, 0x66, 0x02,
	0xdd, 0xa2, 0x9b, 0x25, 0x58, 0x8f, 0x60, 0x36, 0xbf, 0xad, 0x87, 0xd4, 0x01, 0x23, 0xc1, 0x74,
	0x18, 0x08, 0x3d, 0x23, 0x7d, 0xb2, 0xf6, 0xa1, 0x7d, 0x90, 0x20, 0x15, 0x98, 0x91, 0xfc, 0x85,
	0xdc, 0x1e, 0x2c, 0x94, 0x38, 0xb4, 0xe8, 0xc6, 0x98, 0x68, 0x15, 0x47, 0x56, 0xc6, 0x03, 0x20,
	0x2e, 0x52, 0xaf, 0x54, 0x44, 0xe9, 0x9b, 0x59, 0x2f, 0x60, 0x7e, 0x2c, 0xeb, 0xce, 0x32, 0xfb,
	0xd0, 0x7e, 0x17, 0x7b, 0xff, 0xdc, 0x6d, 0x89, 0xe3, 0xce, 0x65, 0x3c, 0x84, 0xf6, 0x21, 0x06,
	0x78, 0xab, 0x8c, 0x72, 0xbf, 0x8b, 0xb0, 0x50, 0xca, 0x53, 0x52, 0xdb, 0x17, 0x53, 0xd0, 0x3a,
	0xc8, 0x76, 0x02, 0x79, 0x0b, 0xf0, 0x0a, 0x85, 0x7e, 0x16, 0xa4, 0x63, 0xab, 0x77, 0x6e, 0x67,
	0x4b, 0xc0, 0x7e, 0x39, 0x5a, 0x02, 0x66, 0x37, 0x2f, 0xa7, 0xf4, 0x80, 0xac, 0xb9, 0x8b, 0x1f,
	0x3f, 0x2f, 0xeb, 0x40, 0x9a, 0x8e, 0x7e, 0x38, 0xe4, 0x03, 0x34, 0xb4, 0x81, 0xc8, 0x62, 0x7e,
	0x6d, 0xdc, 0x90, 0x05, 0xbe, 0x92, 0xd7, 0xac, 0x25, 0xc9, 0x37, 0x6f, 0x35, 0x9d, 0x44, 0x21,
	0x3b, 0xd9, 0xf0, 0x48, 0x1f, 0x0c, 0x65, 0x15, 0xb2, 0x7a, 0x33, 0x9d, 0x0a, 0xff, 0x99, 0x6b,
	0xbf, 0x83, 0xb5, 0x86, 0x29, 0x35, 0xda, 0x56, 0xcb, 0xc9, 0x76, 0xe3, 0x8d, 0xc8, 0x47, 0x98,
	0x1c, 0xd9, 0x84, 0x2c, 0x17, 0x2a, 0x2c, 0x7b, 0xcb, 0x5c, 0xa9, 0x06, 0x35, 0x7d, 0x47, 0xd2,
	0xcf, 0x91, 0x99, 0x9c, 0xde, 0xf9, 0xc6, 0xbc, 0x73, 0x72, 0x59, 0x03, 0x43, 0x7d, 0xfd, 0x42,
	0x03, 0x55, 0x96, 0x2a, 0x34, 0x50, 0xe9, 0x16, 0xeb, 0xb5, 0x54, 0x38, 0x34, 0x3b, 0x05, 0x05,
	0xdd, 0x80, 0xcd, 0xbc, 0xf3, 0xbc, 0x9b, 0xa3, 0xf5, 0xed, 0x3f, 0xa5, 0x90, 0x00, 0x0c, 0xe5,
	0x93, 0x42, 0x51, 0x55, 0x06, 0x2b, 0x14, 0x55, 0xe9, 0x2b, 0xeb, 0xfe, 0xf5, 0xd5, 0x52, 0x2b,
	0xdf, 0x99, 0x6a, 0x06, 0x9b, 0xa5, 0x19, 0x98, 0x7a, 0x47, 0xed, 0x2f, 0x1c, 0xcd, 0x97, 0xff,
	0x97, 0x76, 0xe3, 0x5e, 0xcf, 0x90, 0xbe, 0x7b, 0xf6, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x06,
	0x34, 0xcb, 0xb5, 0x06, 0x00, 0x00,
}
