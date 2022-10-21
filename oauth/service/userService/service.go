package userService

import (
	"context"
	userPb "github.com/viciousvs/OAuth-services/proto/userService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetByLogin(ctx context.Context, login, password string) (*userPb.User, error) {
	conn, err := grpc.Dial(os.Getenv("USER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := userPb.NewUserServiceClient(conn)
	return client.GetOnlyByLogin(ctx, &userPb.LoginRequest{Login: login})
}

func Create(ctx context.Context, login, passwordHash string) (*userPb.UUID, error) {
	conn, err := grpc.Dial(os.Getenv("USER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := userPb.NewUserServiceClient(conn)

	return client.Create(ctx, &userPb.CreateRequest{
		Login:        login,
		PasswordHash: passwordHash,
	})
}
