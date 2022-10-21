package hasherService

import (
	"context"
	hasherPb "github.com/viciousvs/OAuth-services/proto/hasherService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func CompareHashAndService(ctx context.Context, passwordHash, password string) (*hasherPb.CompareResponse, error) {
	hconn, err := grpc.Dial(os.Getenv("HASHER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer hconn.Close()
	hclient := hasherPb.NewHasherServiceClient(hconn)
	return hclient.CompareHashAndPassword(ctx, &hasherPb.CompareRequest{PasswordHash: passwordHash, Password: password})
}

func GenerateHash(ctx context.Context, password string) (*hasherPb.Hash, error) {
	hconn, err := grpc.Dial(os.Getenv("HASHER_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer hconn.Close()
	hclient := hasherPb.NewHasherServiceClient(hconn)
	return hclient.GenerateHash(ctx, &hasherPb.Password{Value: password})
}
