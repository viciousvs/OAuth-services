package validateAccess

import (
	"context"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
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

func (h handler) Handle(ctx context.Context, request *tokenPb.ValidateAccessRequest) (string, error) {
	if request == nil {
		return "nil", customErrors.ErrNilRequest
	}
	accessToken := request.GetAccessToken()
	//the tokenService is validated inside the
	uuid, aID, err := h.jwt.ValidateAccessToken(accessToken)
	if err != nil {
		return "", err
	}
	if err := h.repo.TokensExists(ctx, aID); err != nil {
		return "", err
	}
	return uuid, nil
}
