package signUp

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
	hconn, err := grpc.Dial(os.Getenv("HASHER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer hconn.Close()
	hclient := hasherPb.NewHasherServiceClient(hconn)
	hash, err := hclient.GenerateHash(ctx, &hasherPb.Password{Value: userIN.Password})
	if err != nil {
		return nil, err
	}
	//valid
	if hash == nil {
		return nil, customErrors.ErrNilHashResponse
	}

	///
	userCr := userCreateDTO{
		Login:        userIN.Login,
		PasswordHash: hash.GetPasswordHash(),
	}
	uconn, err := grpc.Dial(os.Getenv("USER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer uconn.Close()
	uclient := userPb.NewUserServiceClient(uconn)
	uuid, err := uclient.Create(ctx, &userPb.CreateRequest{
		Login:        userCr.Login,
		PasswordHash: userCr.PasswordHash,
	})
	if uuid == nil {
		return nil, customErrors.ErrNilUUIDResponse
	}

	////
	tconn, err := grpc.Dial(os.Getenv("token_service_addr"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer tconn.Close()
	tclient := tokenPb.NewTokenServiceClient(tconn)
	tokens, err := tclient.GenerateTokens(ctx, &tokenPb.GenerateTokensRequest{Uuid: uuid.GetValue()})
	if err != nil {
		return nil, err
	}
	if tokens == nil {
		return nil, customErrors.ErrNilTokensResponse
	}

	return &signUpDTO{
		Uuid: uuid.GetValue(),
		Tokens: tokensDTO{
			AtUuid:      tokens.GetAtUuid(),
			AccessToken: tokens.GetAccessToken(),
			AtExpires:   tokens.GetAtExpires(),

			RtUuid:       tokens.GetRtUuid(),
			RefreshToken: tokens.GetRefreshToken(),
			RtExpires:    tokens.GetRtExpires(),
		},
	}, nil
}
