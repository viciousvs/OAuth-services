package refresh

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/utils/customErrors"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
	tokenPb "github.com/viciousvs/OAuth-services/proto/tokenService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func Handle(ctx context.Context, request *oauthPb.RefreshRequest) (*tokensDTO, error) {
	if request == nil {
		return nil, customErrors.ErrNilRequest
	}
	rt := request.GetRefreshToken()

	conn, err := grpc.Dial(os.Getenv("TOKEN_SERVICE_ADDR"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := tokenPb.NewTokenServiceClient(conn)
	tokens, err := client.Refresh(ctx, &tokenPb.RefreshRequest{RefreshToken: rt})
	if err != nil {
		return nil, err
	}
	dto := &tokensDTO{
		AtUuid:      tokens.GetAtUuid(),
		AccessToken: tokens.GetAccessToken(),
		AtExpires:   tokens.GetRtExpires(),

		RtUuid:       tokens.GetRtUuid(),
		RefreshToken: tokens.GetRefreshToken(),
		RtExpires:    tokens.GetRtExpires(),
	}
	return dto, nil
}
