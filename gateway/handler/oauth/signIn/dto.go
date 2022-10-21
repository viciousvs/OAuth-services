package authenticateRefreshToken

type signInResponse struct {
	UUID           string `json:"uuid"`
	AccessTokenExp int64  `json:"access_token_exp"`
	AccessToken    string `json:"access_token"`

	RefreshTokenExp int64  `json:"refresh_token_exp"`
	RefreshToken    string `json:"refresh_token"`
}

type inpUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
