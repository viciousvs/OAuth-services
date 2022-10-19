package access

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/utils/customErrors"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func Handle(ctx context.Context, request *oauthPb.AccessRequest) (string, error) {
	if request == nil {
		return "", customErrors.ErrNilRequest
	}
	at := request.GetAccessToken()

	conn, err := grpc.Dial(os.Getenv("TOKEN_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := tokenPb.NewTokenServiceClient(conn)
	resp, err := client.ValidateAccess(ctx, &tokenPb.ValidateAccessRequest{AccessToken: at})
	if err != nil {
		return "", err
	}
	return resp.GetUserUuid(), nil
}
