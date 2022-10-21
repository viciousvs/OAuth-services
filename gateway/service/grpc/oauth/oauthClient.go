package oauth

import (
	"context"
	oauthPb "github.com/viciousvs/OAuth-services/proto/oauthService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type OAuthService struct {
	addr string
}

func NewOAuthService(addr string) *OAuthService {
	return &OAuthService{addr: addr}
}
func (r OAuthService) SignUp(ctx context.Context, login, password string) (*signUpResponse, error) {
	conn, err := grpc.Dial(r.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := oauthPb.NewOAuthServiceClient(conn)
	resp, err := client.SignUp(context.Background(), &oauthPb.SignUpRequest{Login: login, Password: password})
	if err != nil {
		return nil, err
	}
	return &signUpResponse{
		UUID: resp.GetUuid(),
		Tokens: &tokens{
			AccessToken:  resp.GetTokens().GetAccessToken(),
			AccessExp:    resp.GetTokens().GetAtExpires(),
			RefreshToken: resp.GetTokens().GetRefreshToken(),
			RefreshExp:   resp.GetTokens().GetRtExpires(),
		},
	}, nil
}

func (r OAuthService) SignIn(ctx context.Context, login, password string) (*signInResponse, error) {
	conn, err := grpc.Dial(r.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := oauthPb.NewOAuthServiceClient(conn)
	resp, err := client.SignIn(context.Background(), &oauthPb.SignInRequest{Login: login, Password: password})
	if err != nil {
		return nil, err
	}

	return &signInResponse{
		UUID: resp.GetUuid(),
		Tokens: &tokens{
			AccessToken: resp.GetTokens().GetAccessToken(),
			AccessExp:   resp.GetTokens().GetAtExpires(),

			RefreshToken: resp.GetTokens().GetRefreshToken(),
			RefreshExp:   resp.GetTokens().GetRtExpires(),
		},
	}, nil
}

func (r OAuthService) AuthenticateAccessToken(ctx context.Context, accessToken string) (string, error) {
	conn, err := grpc.Dial(r.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	client := oauthPb.NewOAuthServiceClient(conn)
	resp, err := client.Access(context.Background(), &oauthPb.AccessRequest{AccessToken: accessToken})
	if err != nil {
		return "", err
	}
	return resp.GetUserUuid(), nil
}
func (r OAuthService) AuthenticateRefreshToken(ctx context.Context, refreshToken string) (*tokens, error) {
	conn, err := grpc.Dial(r.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := oauthPb.NewOAuthServiceClient(conn)
	resp, err := client.Refresh(context.Background(), &oauthPb.RefreshRequest{RefreshToken: refreshToken})
	if err != nil {
		return nil, err
	}
	return &tokens{
		AccessToken:  resp.GetAccessToken(),
		AccessExp:    resp.GetAtExpires(),
		RefreshToken: resp.GetRefreshToken(),
		RefreshExp:   resp.GetRtExpires(),
	}, nil
}
