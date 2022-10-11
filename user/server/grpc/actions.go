package grpc

import (
	"context"

	userPb "github.com/viciousvs/OAuth-services/proto/userService"
)

func (s *Server) Create(ctx context.Context, request *userPb.CreateRequest) (*userPb.CreateResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetID(ctx context.Context, request *userPb.GetIDRequest) (*userPb.GetIDResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
