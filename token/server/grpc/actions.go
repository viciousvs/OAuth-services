package grpc

import (
	"context"
	userPb "github.com/viciousvs/OAuth-services/proto/userService"
)

func (s *Server) Create(ctx context.Context, request *userPb.CreateRequest) (*userPb.UUID, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) GetByLogin(ctx context.Context, request *userPb.GetByLoginRequest) (*userPb.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
