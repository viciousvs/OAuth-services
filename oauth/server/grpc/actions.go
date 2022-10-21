package grpc

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/handler/access"
	"github.com/viciousvs/OAuth-services/oauth/handler/refresh"
	"github.com/viciousvs/OAuth-services/oauth/handler/signIn"
	"github.com/viciousvs/OAuth-services/oauth/handler/signUp"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
)

func (s *Server) SignUp(ctx context.Context, request *oauthPb.SignUpRequest) (*oauthPb.SignUpResponse, error) {
	//TODO implement me
	sUp, err := signUp.Handle(ctx, request)
	if err != nil {
		return nil, err
	}

	return &oauthPb.SignUpResponse{
		Uuid: sUp.UUID,
		Tokens: &oauthPb.Tokens{
			AccessToken:  sUp.Tokens.AccessToken,
			AtExpires:    sUp.Tokens.AtExpires,
			RefreshToken: sUp.Tokens.RefreshToken,
			RtExpires:    sUp.Tokens.RtExpires,
		},
	}, nil
}

func (s *Server) SignIn(ctx context.Context, request *oauthPb.SignInRequest) (*oauthPb.SignInResponse, error) {
	//TODO implement me
	sIn, err := signIn.Handle(ctx, request)
	if err != nil || sIn == nil {
		return nil, err
	}
	return &oauthPb.SignInResponse{
		Uuid: sIn.UUID,
		Tokens: &oauthPb.Tokens{
			AccessToken:  sIn.Tokens.AccessToken,
			AtExpires:    sIn.Tokens.AtExpires,
			RefreshToken: sIn.Tokens.RefreshToken,
			RtExpires:    sIn.Tokens.RtExpires,
		},
	}, nil
}

func (s *Server) Access(ctx context.Context, request *oauthPb.AccessRequest) (*oauthPb.AccessResponse, error) {
	//TODO implement me
	uuid, err := access.Handle(ctx, request)
	if err != nil {
		return nil, err
	}
	return &oauthPb.AccessResponse{
		UserUuid: uuid,
	}, nil
}

func (s *Server) Refresh(ctx context.Context, request *oauthPb.RefreshRequest) (*oauthPb.Tokens, error) {
	//TODO implement me
	tokens, err := refresh.Handle(ctx, request)
	if err != nil || tokens == nil {
		return nil, err
	}
	return &oauthPb.Tokens{
		AccessToken:  tokens.AccessToken,
		AtExpires:    tokens.AtExpires,
		RefreshToken: tokens.RefreshToken,
		RtExpires:    tokens.RtExpires,
	}, err
}

func (s *Server) mustEmbedUnimplementedOAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}
