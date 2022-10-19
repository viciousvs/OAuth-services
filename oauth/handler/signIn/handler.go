package signIn

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/utils/customErrors"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	userPb "github.com/viciousvs/OAuth-services/proto/userService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
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
	uconn, err := grpc.Dial(os.Getenv("USER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer uconn.Close()
	uclient := userPb.NewUserServiceClient(uconn)
	u, err := uclient.GetOnlyByLogin(ctx, &userPb.LoginRequest{Login: uIn.Login})
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, customErrors.ErrNilUserResponse
	}
	//////
	hconn, err := grpc.Dial(os.Getenv("HASHER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer hconn.Close()
	hclient := hasherPb.NewHasherServiceClient(hconn)
	ok, err := hclient.CompareHashAndPassword(ctx, &hasherPb.CompareRequest{PasswordHash: u.GetPasswordHash(), Password: uIn.Password})
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
	tconn, err := grpc.Dial(os.Getenv("token_service_addr"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer tconn.Close()
	tclient := tokenPb.NewTokenServiceClient(tconn)
	tokens, err := tclient.GenerateTokens(ctx, &tokenPb.GenerateTokensRequest{Uuid: u.GetUuid()})
	if err != nil {
		return nil, err
	}
	if tokens == nil {
		return nil, customErrors.ErrNilTokensResponse
	}

	return &signInDTO{
		uuid: u.GetUuid(),
		tokens: tokensDTO{
			AtUuid:      tokens.GetAtUuid(),
			AccessToken: tokens.GetAccessToken(),
			AtExpires:   tokens.GetAtExpires(),

			RtUuid:       tokens.GetRtUuid(),
			RefreshToken: tokens.GetRefreshToken(),
			RtExpires:    tokens.GetRtExpires(),
		},
	}, nil
}
