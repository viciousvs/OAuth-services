package grpc

import (
	"context"
	"github.com/viciousvs/OAuth-services/hasher/handler/compare"
	"github.com/viciousvs/OAuth-services/hasher/handler/generate"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
)

func (s *Server) GenerateHash(ctx context.Context, req *hasherPb.Password) (*hasherPb.Hash, error) {
	hash, err := generate.Handle(req)
	if err != nil {
		return nil, err
	}
	return &hasherPb.Hash{PasswordHash: hash}, nil
}

func (s *Server) CompareHashAndPassword(ctx context.Context, request *hasherPb.CompareRequest) (*hasherPb.CompareResponse, error) {
	match, err := compare.Handle(request);
	if err != nil {
		return nil, err
	}
	return &hasherPb.CompareResponse{Value: match}, nil
}

func (s *Server) mustEmbedUnimplementedHasherServiceServer() {
	//TODO implement me
	panic("implement me")
}
