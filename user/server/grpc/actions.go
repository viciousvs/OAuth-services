package grpc

import (
	"context"
	userPb "github.com/viciousvs/OAuth-services/proto/userService"
	"github.com/viciousvs/OAuth-services/user/handler/create"
	"github.com/viciousvs/OAuth-services/user/handler/getByLogin"
	onlyByLogin "github.com/viciousvs/OAuth-services/user/handler/getOnlyByLogin"
)

func (s *Server) Create(ctx context.Context, request *userPb.CreateRequest) (*userPb.UUID, error) {
	h := create.NewHandler(s.repo)
	uuid, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &userPb.UUID{Value: uuid}, nil
}

func (s *Server) GetByLogin(ctx context.Context, request *userPb.GetByLoginRequest) (*userPb.User, error) {
	h := getByLogin.NewHandler(s.repo)
	u, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &userPb.User{Uuid: u.UUID, Login: u.Login, PasswordHash: u.PasswordHash}, nil
}
func (s *Server) GetOnlyByLogin(ctx context.Context, request *userPb.LoginRequest) (*userPb.User, error) {
	h := onlyByLogin.NewHandler(s.repo)
	u, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &userPb.User{Uuid: u.UUID, Login: u.Login, PasswordHash: u.PasswordHash}, nil
}

func (s *Server) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("IMPLEMENT ME")
}
