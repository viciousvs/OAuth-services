package authenticateRefreshToken

import (
	"encoding/json"
	"errors"
	"github.com/viciousvs/OAuth-services/gateway/service/grpc/oauth"
	"github.com/viciousvs/OAuth-services/gateway/utils/httpUtils"
	"io"
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
		if errors.Is(err, io.EOF) {
			httpUtils.NewErrorResponse(w, http.StatusBadRequest, "body couldn't be empty")
			return
		}
		httpUtils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if !refToken.ValidateToken() {
		httpUtils.NewErrorResponse(w, http.StatusBadRequest, "invalid refresh token")
		return
	}

	tok, err := oauth.NewOAuthService(h.oAuthServiceAddr).AuthenticateRefreshToken(r.Context(), refToken.TokenString)
	if err != nil {
		httpUtils.NewErrorResponse(w, http.StatusUnauthorized, "refresh token doesn't exist or not valid")
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
		httpUtils.NewErrorResponse(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(200)
	w.Write(b)
}
