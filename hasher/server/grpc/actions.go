package grpc

import (
	"context"
	"github.com/viciousvs/OAuth-services/hasher/model/hasher/handler/compare"
	"github.com/viciousvs/OAuth-services/hasher/model/hasher/handler/generate"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
)

func (s *Server) GenerateHash(ctx context.Context, req *hasherPb.Password) (*hasherPb.Hash, error) {
	h := generate.NewHandler(s.repo)
	hash, err := h.Handle(req)
	if err != nil {
		return nil, err
	}
	return &hasherPb.Hash{PasswordHash: hash}, nil
}

func (s *Server) CompareHashAndPassword(ctx context.Context, request *hasherPb.CompareRequest) (*hasherPb.CompareResponse, error) {
	h := compare.NewHandler(s.repo)

}

func (s *Server) mustEmbedUnimplementedHasherServiceServer() {
	//TODO implement me
	panic("implement me")
}
