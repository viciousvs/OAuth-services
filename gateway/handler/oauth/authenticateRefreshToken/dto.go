package authenticateRefreshToken

type refreshToken struct {
	TokenString string `json:"refresh_token"`
}

type tokens struct {
	AccessTokenExp int64  `json:"access_token_exp"`
	AccessToken    string `json:"access_token"`

	RefreshTokenExp int64  `json:"refresh_token_exp"`
	RefreshToken    string `json:"refresh_token"`
}

func (t refreshToken) ValidateToken() bool {
	return t.TokenString != ""
}
