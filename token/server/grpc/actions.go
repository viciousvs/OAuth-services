package grpc

import (
	"context"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
)

func (s *Server) GenerateTokens(ctx context.Context, request *tokenPb.GenerateTokensRequest) (*tokenPb.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) VerifyAccess(ctx context.Context, request *tokenPb.VerifyAccessRequest) (*tokenPb.VerifyAccessResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Refresh(ctx context.Context, request *tokenPb.RefreshRequest) (*tokenPb.Tokens, error) {
	//TODO implement me
	panic("implement me")
}

func (s *Server) mustEmbedUnimplementedTokenServiceServer() {
	//TODO implement me
	panic("implement me")
}
