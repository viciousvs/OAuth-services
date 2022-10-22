package authenticateRefreshToken

import "strings"

const (
	invalidLogPass  = "invalid login and password"
	invalidLogin    = "invalid login"
	invalidPassword = "invalid password"
)

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

func (d inpUser) invalidateLogin() bool {
	return !(len(d.Login) >= 5 && !strings.Contains(d.Login, " "))
}
func (d inpUser) invalidatePassword() bool {
	return !(len(d.Password) >= 8 && !strings.Contains(d.Password, " "))
}
func (d inpUser) Validate() (string, bool) {
	if d.invalidateLogin() && d.invalidatePassword() {
		return invalidLogPass, false
	}
	if d.invalidateLogin() {
		return invalidLogin, false
	}
	if d.invalidatePassword() {
		return invalidPassword, false
	}
	return "", true
}
