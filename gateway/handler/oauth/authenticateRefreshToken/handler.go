package authenticateRefreshToken

import (
	"encoding/json"
	"github.com/viciousvs/OAuth-services/gateway/service/grpc/oauth"
	"github.com/viciousvs/OAuth-services/gateway/utils/httpUtils"
	"net/http"
)

type handler struct {
	oAuthServiceAddr string
}

func NewHandler(oAuthServiceAddr string) *handler {
	return &handler{oAuthServiceAddr}
}
func (h handler) Handle(w http.ResponseWriter, r *http.Request) {
	//TODO validations
	var refToken refreshToken
	err := json.NewDecoder(r.Body).Decode(&refToken)
	if err != nil {
		httpUtils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	tok, err := oauth.NewOAuthService(h.oAuthServiceAddr).AuthenticateRefreshToken(r.Context(), refToken.TokenString)
	if err != nil {
		httpUtils.NewErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}
	resp := tokens{
		AccessToken:     tok.AccessToken,
		AccessTokenExp:  tok.AccessExp,
		RefreshToken:    tok.RefreshToken,
		RefreshTokenExp: tok.RefreshExp,
	}
	b, err := json.Marshal(resp)
	if err != nil {
		httpUtils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(b)
}
