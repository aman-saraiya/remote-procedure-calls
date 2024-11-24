/*
 * This file implements the gRPC handlers for UserService
 */

package userservice

import (
	"context"
	"fmt"
	"log"

	gen_usersvc "github.com/aman-saraiya/remote-procedure-calls/grpc-protobuf/client/generated/userservice"
)

// Define the UserService gRPC service struct
type UserServiceServer struct {
	gen_usersvc.UnimplementedUserServiceServer
	Users map[int32]*gen_usersvc.User // Simple in-memory store for users
}

// Function to create UserServiceServer instance
func NewUserServiceServer() *UserServiceServer {
	return &UserServiceServer{
		Users: make(map[int32]*gen_usersvc.User),
	}
}

// Handler for AddUser RPC
func (s *UserServiceServer) AddUser(ctx context.Context, req *gen_usersvc.AddUserArg) (*gen_usersvc.AddUserRet, error) {
	log.Println("Handling AddUser RPC with argument ", req)
	userID := int32(len(s.Users)+1)

	// Create a new user
	user := &gen_usersvc.User{
		Id:    userID,
		Email: req.GetEmail(),
		Name:  req.GetName(),
	}

	s.Users[userID] = user

	return &gen_usersvc.AddUserRet{
		Id: userID,
	}, nil
}

// Handler for GetUser RPC
func (s *UserServiceServer) GetUser(ctx context.Context, req *gen_usersvc.GetUserArg) (*gen_usersvc.GetUserRet, error) {
	log.Println("Handling GetUser RPC with argument ", req)
	userID := req.GetId()

	user, exists := s.Users[userID]
	if !exists {
		return nil, fmt.Errorf("User with ID %d not found", userID)
	}

	return &gen_usersvc.GetUserRet{
		User: user,
	}, nil
}
