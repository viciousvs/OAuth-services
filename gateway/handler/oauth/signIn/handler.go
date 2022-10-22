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
	var inUser inpUser
	err := json.NewDecoder(r.Body).Decode(&inUser)
	if err != nil {
		if errors.Is(err, io.EOF) {
			httpUtils.NewErrorResponse(w, http.StatusBadRequest, "body couldn't be empty")
			return
		}
		httpUtils.NewErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	if msg, ok := inUser.Validate(); !ok {
		httpUtils.NewErrorResponse(w, http.StatusBadRequest, msg)
		return
	}

	sin, err := oauth.NewOAuthService(h.oAuthServiceAddr).SignIn(r.Context(), inUser.Login, inUser.Password)
	if err != nil {
		httpUtils.NewErrorResponse(w, http.StatusConflict, "user not found")
		return
	}
	resp := signInResponse{
		UUID:            sin.UUID,
		AccessToken:     sin.Tokens.AccessToken,
		AccessTokenExp:  sin.Tokens.AccessExp,
		RefreshToken:    sin.Tokens.RefreshToken,
		RefreshTokenExp: sin.Tokens.RefreshExp,
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
