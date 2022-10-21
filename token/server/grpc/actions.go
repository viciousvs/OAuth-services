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
	h := generate.NewHandler(s.repo, s.jwt)
	tokens, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &tokenPb.Tokens{
		AtUuid:       tokens.AccessToken.Claims.ID,
		AccessToken:  tokens.AccessToken.TokenString,
		AtExpires:    tokens.AccessToken.Claims.ExpiresAt.Unix(),
		RtUuid:       tokens.RefreshToken.Claims.ID,
		RefreshToken: tokens.RefreshToken.TokenString,
		RtExpires:    tokens.RefreshToken.Claims.ExpiresAt.Unix(),
	}, nil
}

func (s *Server) ValidateAccess(ctx context.Context, request *tokenPb.ValidateAccessRequest) (*tokenPb.ValidateAccessResponse, error) {
	//TODO implement me
	h := validateAccess.NewHandler(s.repo, s.jwt)
	uuid, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &tokenPb.ValidateAccessResponse{UserUuid: uuid}, nil
}

func (s *Server) Refresh(ctx context.Context, request *tokenPb.RefreshRequest) (*tokenPb.Tokens, error) {
	//TODO implement me
	h := refresh.NewHandler(s.repo, s.jwt)
	tokens, err := h.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &tokenPb.Tokens{
		AtUuid:       tokens.AccessToken.Claims.ID,
		AccessToken:  tokens.AccessToken.TokenString,
		AtExpires:    tokens.AccessToken.Claims.ExpiresAt.Unix(),
		RtUuid:       tokens.RefreshToken.Claims.ID,
		RefreshToken: tokens.RefreshToken.TokenString,
		RtExpires:    tokens.RefreshToken.Claims.ExpiresAt.Unix(),
	}, nil
}

func (s *Server) mustEmbedUnimplementedTokenServiceServer() {
	//TODO implement me
	panic("implement me")
}
