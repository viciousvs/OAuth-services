package generate

import (
	"context"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"github.com/viciousvs/OAuth-services/token/model/token"
	"github.com/viciousvs/OAuth-services/token/model/token/repository"
	"github.com/viciousvs/OAuth-services/token/service/tokenJWT"
	"github.com/viciousvs/OAuth-services/token/utils/customErrors"
	"github.com/viciousvs/OAuth-services/token/utils/uuid"
)

type handler struct {
	repo repository.Repository
	jwt  tokenJWT.JWT
}

func NewHandler(repo repository.Repository, jwt tokenJWT.JWT) *handler {
	return &handler{repo: repo, jwt: jwt}
}

func (h handler) Handle(ctx context.Context, request *tokenPb.GenerateTokensRequest) (*token.Tokens, error) {
	if request == nil {
		return nil, customErrors.ErrNilRequest
	}
	userUUID := request.GetUuid()
	if !uuid.IsValidUUID(userUUID) {
		return nil, customErrors.ErrInvalidUUID
	}

	tokens, err := h.jwt.Generate(userUUID)
	if err != nil {
		return nil, err
	}

	if err := h.repo.SetTokens(ctx, tokens); err != nil {
		return nil, err
	}
	return tokens, nil
}
