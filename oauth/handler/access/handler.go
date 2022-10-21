package access

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/service/tokenService"
	"github.com/viciousvs/OAuth-services/oauth/utils/customErrors"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
)

func Handle(ctx context.Context, request *oauthPb.AccessRequest) (string, error) {
	if request == nil {
		return "", customErrors.ErrNilRequest
	}
	at := request.GetAccessToken()

	resp, err := tokenService.ValidateAccessToken(ctx, at)

	if err != nil {
		return "", err
	}
	return resp.GetUserUuid(), nil
}
