package refresh

import (
	"context"
	"errors"
	"fmt"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"github.com/viciousvs/OAuth-services/token/model/token"
	"github.com/viciousvs/OAuth-services/token/model/token/repository"
	"github.com/viciousvs/OAuth-services/token/service/tokenJWT"
	"github.com/viciousvs/OAuth-services/token/utils/customErrors"
)

type handler struct {
	repo repository.Repository
	jwt  tokenJWT.JWT
}

func NewHandler(repo repository.Repository, jwt tokenJWT.JWT) *handler {
	return &handler{repo: repo, jwt: jwt}
}

func (h handler) Handle(ctx context.Context, request *tokenPb.RefreshRequest) (*token.Tokens, error) {
	if request == nil {
		return nil, customErrors.ErrNilRequest
	}
	refreshToken := request.GetRefreshToken()
	//the token is validated inside the
	userUUID, aID, rID, err := h.jwt.Refresh(refreshToken)
	if err != nil {
		return nil, err
	}
	if err := h.repo.TokensExists(ctx, aID, rID); err != nil {
		if errors.Is(err, fmt.Errorf("not found")) {
			return nil, err
		}
		return nil, err
	}
	if err := h.repo.DeleteTokens(ctx, aID, rID); err != nil {
		return nil, err
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
