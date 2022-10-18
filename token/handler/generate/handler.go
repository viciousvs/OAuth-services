package generate

import (
	"context"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"github.com/viciousvs/OAuth-services/token/model/token"
	"github.com/viciousvs/OAuth-services/token/model/token/repository"
	"github.com/viciousvs/OAuth-services/token/model/token/service"
)

type handler struct {
	repo    repository.Repository
	service service.Service
}

func NewHandler(repo repository.Repository, service service.Service) *handler {
	return &handler{repo: repo, service: service}
}

func (h handler) Handle(ctx context.Context, request *tokenPb.GenerateTokensRequest) (*token.Tokens, error) {
	userUUID := request.GetUuid()
	//TODO validation
	tokens, err := h.service.Generate(userUUID)
	if err != nil {
		return nil, err
	}

	if err := h.repo.SetTokens(ctx, tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}
