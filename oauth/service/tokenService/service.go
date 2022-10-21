package tokenService

import (
	"context"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func ValidateAccessToken(ctx context.Context, accessToken string) (*tokenPb.ValidateAccessResponse, error) {
	conn, err := grpc.Dial(os.Getenv("TOKEN_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := tokenPb.NewTokenServiceClient(conn)
	return client.ValidateAccess(ctx, &tokenPb.ValidateAccessRequest{AccessToken: accessToken})
}

func Refresh(ctx context.Context, refreshToken string) (*tokenPb.Tokens, error) {
	conn, err := grpc.Dial(os.Getenv("TOKEN_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := tokenPb.NewTokenServiceClient(conn)
	return client.Refresh(ctx, &tokenPb.RefreshRequest{RefreshToken: refreshToken})
}

func GenerateTokens(ctx context.Context, uuid string) (*tokenPb.Tokens, error) {
	conn, err := grpc.Dial(os.Getenv("TOKEN_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := tokenPb.NewTokenServiceClient(conn)
	return client.GenerateTokens(ctx, &tokenPb.GenerateTokensRequest{Uuid: uuid})
}
