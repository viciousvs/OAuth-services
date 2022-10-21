package signUp

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/service/hasherService"
	"github.com/viciousvs/OAuth-services/oauth/service/tokenService"
	"github.com/viciousvs/OAuth-services/oauth/service/userService"
	"github.com/viciousvs/OAuth-services/oauth/utils/customErrors"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
)

func Handle(ctx context.Context, request *oauthPb.SignUpRequest) (*signUpDTO, error) {
	if request == nil {
		return nil, customErrors.ErrNilRequest
	}
	userIN := userSupDTO{
		request.GetLogin(),
		request.GetPassword(),
	}
	//TODO need valid request
	//hash input password
	hash, err := hasherService.GenerateHash(ctx, userIN.Password)
	if err != nil {
		return nil, err
	}
	if hash == nil {
		return nil, customErrors.ErrNilHashResponse
	}

	//valid
	///
	userCr := userCreateDTO{
		Login:        userIN.Login,
		PasswordHash: hash.GetPasswordHash(),
	}
	//TODO
	uuid, err := userService.Create(ctx, userCr.Login, userCr.PasswordHash)
	if err != nil {
		return nil, err
	}
	if uuid == nil {
		return nil, customErrors.ErrNilUUIDResponse
	}

	////
	tokens, err := tokenService.GenerateTokens(ctx, uuid.GetValue())
	if err != nil {
		return nil, err
	}
	if tokens == nil {
		return nil, customErrors.ErrNilTokensResponse
	}

	return &signUpDTO{
		UUID: uuid.GetValue(),
		Tokens: tokensDTO{
			AccessToken: tokens.GetAccessToken(),
			AtExpires:   tokens.GetAtExpires(),

			RefreshToken: tokens.GetRefreshToken(),
			RtExpires:    tokens.GetRtExpires(),
		},
	}, nil
}
