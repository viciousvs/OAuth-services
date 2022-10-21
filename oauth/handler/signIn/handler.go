package signIn

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/service/hasherService"
	"github.com/viciousvs/OAuth-services/oauth/service/tokenService"
	"github.com/viciousvs/OAuth-services/oauth/service/userService"
	"github.com/viciousvs/OAuth-services/oauth/utils/customErrors"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
)

func Handle(ctx context.Context, request *oauthPb.SignInRequest) (*signInDTO, error) {
	if request == nil {
		return nil, customErrors.ErrNilRequest
	}
	uIn := userInDTO{
		Login:    request.GetLogin(),
		Password: request.GetPassword(),
	}
	if err := uIn.validate(); err != nil {
		return nil, customErrors.ErrInvalidData
	}

	u, err := userService.GetByLogin(ctx, uIn.Login, uIn.Password)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, customErrors.ErrNilUserResponse
	}

	//////
	ok, err := hasherService.CompareHashAndService(ctx, u.GetPasswordHash(), uIn.Password)
	if err != nil {
		return nil, err
	}
	if ok == nil {
		return nil, customErrors.ErrNilHashResponse
	}
	if !ok.GetValue() {
		return nil, customErrors.ErrInvalidData
	}

	///////////

	tokens, err := tokenService.GenerateTokens(ctx, u.Uuid)
	if err != nil {
		return nil, err
	}
	if tokens == nil {
		return nil, customErrors.ErrNilTokensResponse
	}

	return &signInDTO{
		UUID: u.GetUuid(),
		Tokens: tokensDTO{
			AccessToken: tokens.GetAccessToken(),
			AtExpires:   tokens.GetAtExpires(),

			RefreshToken: tokens.GetRefreshToken(),
			RtExpires:    tokens.GetRtExpires(),
		},
	}, nil
}
