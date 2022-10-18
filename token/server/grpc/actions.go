package grpc

import (
	"context"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"github.com/viciousvs/OAuth-services/token/handler/generate"
	"github.com/viciousvs/OAuth-services/token/handler/refresh"
	"github.com/viciousvs/OAuth-services/token/handler/validateAccess"
)

func (s *Server) GenerateTokens(ctx context.Context, request *tokenPb.GenerateTokensRequest) (*tokenPb.Tokens, error) {
	//TODO implement me
	h := generate.NewHandler(s.repo, s.service)
	tokens, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &tokenPb.Tokens{AccessToken: tokens.AccessToken.TokenString, RefreshToken: tokens.RefreshToken.TokenString}, nil
}

func (s *Server) ValidateAccess(ctx context.Context, request *tokenPb.ValidateAccessRequest) (*tokenPb.ValidateAccessResponse, error) {
	//TODO implement me
	h := validateAccess.NewHandler(s.repo, s.service)
	uuid, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &tokenPb.ValidateAccessResponse{UserUuid: uuid}, nil
}

func (s *Server) Refresh(ctx context.Context, request *tokenPb.RefreshRequest) (*tokenPb.Tokens, error) {
	//TODO implement me
	h := refresh.NewHandler(s.repo, s.service)
	tokens, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &tokenPb.Tokens{AccessToken: tokens.AccessToken.TokenString, RefreshToken: tokens.RefreshToken.TokenString}, nil
}

func (s *Server) mustEmbedUnimplementedTokenServiceServer() {
	//TODO implement me
	panic("implement me")
}
