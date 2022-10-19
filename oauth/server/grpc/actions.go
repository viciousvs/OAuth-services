package grpc

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/handler/signUp"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
)

func (s *Server) SignUp(ctx context.Context, request *oauthPb.SignUpRequest) (*oauthPb.SignUpResponse, error) {
	sUp, err := signUp.Handle(ctx, request)
	if err != nil {
		return nil, err
	}

	return &oauthPb.SignUpResponse{
		Uuid: sUp.Uuid,
		Tokens: &oauthPb.Tokens{
			AtUuid:       sUp.Tokens.AtUuid,
			AccessToken:  sUp.Tokens.AccessToken,
			AtExpires:    sUp.Tokens.AtExpires,
			RtUuid:       sUp.Tokens.RtUuid,
			RefreshToken: sUp.Tokens.RefreshToken,
			RtExpires:    sUp.Tokens.RtExpires,
		},
	}, nil
}

func (s *Server) SignIn(ctx context.Context, request *oauthPb.SignInRequest) (*oauthPb.SignInResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Access(ctx context.Context, request *oauthPb.AccessRequest) (*oauthPb.AccessResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Refresh(ctx context.Context, request *oauthPb.RefreshRequest) (*oauthPb.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedOAuthServiceServer() {
	//TODO implement me
	panic("implement me")
}
