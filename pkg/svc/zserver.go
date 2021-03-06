package svc

import (
	"context"
	//"contacts1/pkg/pb"
	"github.com/gshirazi/contacts1/pkg/pb"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/gorm"
)

// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~ A BRIEF DEVELOPMENT GUIDE ~~~~~~~~~~~~~~~~~~~~~~~~~
//
// TODO: Extend the Contacts1 service by defining new RPCs and
// and message types in the pb/service.proto file. These RPCs and messages
// compose the API for your service. After modifying the proto schema in
// pb/service.proto, call "make protobuf" to regenerate the protobuf files.
//
// TODO: Create an implementation of the Contacts1 server
// interface. This interface is generated by the protobuf compiler and exists
// inside the pb/service.pb.go file. The "server" struct already provides an
// implementation of Contacts1 server interface service, but only
// for the GetVersion function. You will need to implement any new RPCs you
// add to your protobuf schema.
//
// TODO: Update the GetVersion function when newer versions of your service
// become available. Feel free to change GetVersion to better-suit how your
// versioning system, or get rid of it entirely. GetVersion helps make up a
// simple "starter" example that allows an end-to-end example. It is not
// required.
//
// TODO: Oh yeah, delete this guide when you no longer need it.
//
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~ FAREWELL AND GOOD LUCK ~~~~~~~~~~~~~~~~~~~~~~~~~~
// ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

const (
	// version is the current version of the service
	version = "0.0.4"
)

// Default implementation of the Contacts1 server interface
// type server struct{ db *gorm.DB }

type defaultContactsServer struct {
	db *gorm.DB
}

type contactsServer struct {
	*pb.Contacts1DefaultServer
}

// NewBasicServer returns an instance of the default server interface'
//func NewBasicServer(database *gorm.DB) (pb.Contacts1Server, error) {
//	return &server{database}, nil
//}

// NewContactServer returns an instance of the enhanced Contacts1 server
func NewContactsServer(database *gorm.DB) (pb.Contacts1Server, error) {
	// return &defaultContactsServer{database}, nil
	return &contactsServer{&pb.Contacts1DefaultServer{database}}, nil
}
// GetVersion returns the current version of the service
func (defaultContactsServer) GetVersion(context.Context, *empty.Empty) (*pb.VersionResponse, error) {
	return &pb.VersionResponse{Version: version}, nil
}

// Reverse returns the first name in reverse
func (defaultContactsServer) Reverse(ctx context.Context, req *pb.ReverseRequest) (*pb.ReverseResponse, error) {
	r := []rune(req.Payload.FirstName)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return &pb.ReverseResponse{Result: string(r)}, nil
}

func (server defaultContactsServer) Create(ctx context.Context, req *pb.CreateContactRequest) (*pb.CreateContactResponse, error) {
	newContact := req.GetPayload()
	contactResponse, err := pb.DefaultCreateContact(ctx, newContact, server.db)
	if err != nil {
		return nil, err
	}
	return &pb.CreateContactResponse{Result: contactResponse}, nil
}

func (server defaultContactsServer) Read(ctx context.Context, req *pb.ReadContactRequest) (*pb.ReadContactResponse, error) {
	newContact := pb.Contact{}
	id := req.GetId()
	newContact.Id = id
	contactResponse, err := pb.DefaultReadContact(ctx, &newContact, server.db)
	if err != nil {
		return nil, err
	}
	return &pb.ReadContactResponse{Result: contactResponse}, nil
}

func (server defaultContactsServer) Update(ctx context.Context, req *pb.UpdateContactRequest) (*pb.UpdateContactResponse, error) {
	newContact := req.GetPayload()
	contactResponse, err := pb.DefaultStrictUpdateContact(ctx, newContact, server.db)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateContactResponse{Result: contactResponse}, nil
}

func (server defaultContactsServer) Delete(ctx context.Context, req *pb.DeleteContactRequest) (*pb.DeleteContactResponse, error) {
	newContact := pb.Contact{}
	id := req.GetId()
	newContact.Id = id
	err := pb.DefaultDeleteContact(ctx, &newContact, server.db)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteContactResponse{}, nil
}

