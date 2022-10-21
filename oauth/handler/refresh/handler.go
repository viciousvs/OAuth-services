package refresh

import (
	"context"
	"github.com/viciousvs/OAuth-services/oauth/service/tokenService"
	"github.com/viciousvs/OAuth-services/oauth/utils/customErrors"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
)

func Handle(ctx context.Context, request *oauthPb.RefreshRequest) (*tokensDTO, error) {
	if request == nil {
		return nil, customErrors.ErrNilRequest
	}
	rt := request.GetRefreshToken()

	tokens, err := tokenService.Refresh(ctx, rt)
	if err != nil {
		return nil, err
	}
	dto := &tokensDTO{
		AccessToken: tokens.GetAccessToken(),
		AtExpires:   tokens.GetRtExpires(),

		RefreshToken: tokens.GetRefreshToken(),
		RtExpires:    tokens.GetRtExpires(),
	}
	return dto, nil
}
