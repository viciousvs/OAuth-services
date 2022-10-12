package grpc

import (
	"context"
	"github.com/viciousvs/OAuth-services/user/model/user/handler/create"
	"github.com/viciousvs/OAuth-services/user/model/user/handler/getUUID"

	userPb "github.com/viciousvs/OAuth-services/proto/userService"
)

//Create
func (s *Server) Create(ctx context.Context, request *userPb.CreateRequest) (*userPb.CreateResponse, error) {
	//TODO create Errors
	h := create.NewHandler(s.repo)
	uuid, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &userPb.CreateResponse{Uuid: uuid}, err
}

//GetUUID
func (s *Server) GetUUID(ctx context.Context, request *userPb.GetUUIDRequest) (*userPb.GetUUIDResponse, error) {
	//TODO implement me
	h := getUUID.NewHandler(s.repo)
	uuid, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err

	}
	return &userPb.GetUUIDResponse{Uuid: uuid}, nil
}

//mustEmbedUnimplementedUserServiceServer
func (s *Server) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}
