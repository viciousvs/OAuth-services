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
	var inUser inpUser
	err := json.NewDecoder(r.Body).Decode(&inUser)
	if err != nil {
		httpUtils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	sup, err := oauth.NewOAuthService(h.oAuthServiceAddr).SignUp(r.Context(), inUser.Login, inUser.Password)
	if err != nil {
		httpUtils.NewErrorResponse(w, http.StatusForbidden, err.Error())
		return
	}
	resp := signUpResponse{
		UUID:            sup.UUID,
		AccessToken:     sup.Tokens.AccessToken,
		AccessTokenExp:  sup.Tokens.AccessExp,
		RefreshToken:    sup.Tokens.RefreshToken,
		RefreshTokenExp: sup.Tokens.RefreshExp,
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
