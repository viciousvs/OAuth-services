package middlewares

import (
	"context"
	"github.com/viciousvs/OAuth-services/gateway/service/grpc/oauth"
	"github.com/viciousvs/OAuth-services/gateway/utils/httpUtils"
	"net/http"
	"strings"
)

const authorizationHeader = "Authorization"

type Middleware struct {
	oAuthServiceAddr string
}

func NewMiddleware(oAuthServiceAddr string) Middleware {
	return Middleware{oAuthServiceAddr}
}
func (m Middleware) EnsureAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get(authorizationHeader)
		if header == "" {
			httpUtils.NewErrorResponse(w, http.StatusUnauthorized, "empty auth header")
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			httpUtils.NewErrorResponse(w, http.StatusUnauthorized, "invalid auth header")
			return
		}

		if len(headerParts[1]) == 0 {
			httpUtils.NewErrorResponse(w, http.StatusUnauthorized, "token is empty")
			return
		}
		uuid, err := oauth.NewOAuthService(m.oAuthServiceAddr).AuthenticateAccessToken(r.Context(), headerParts[1])
		if err != nil {
			httpUtils.NewErrorResponse(w, http.StatusForbidden, "access token not found")
			return
		}

		ctx := context.WithValue(r.Context(), "userUUID", uuid)
		rWithUUID := r.WithContext(ctx)
		next.ServeHTTP(w, rWithUUID)
	}

	return http.HandlerFunc(fn)
}

func JsonMiddleware(handler http.HandlerFunc) http.Handler {
	return contentTypeHandler(handler, "application/json")
}
func isContentType(h http.Header, contentType string) bool {
	ct := h.Get("Content-Type")
	if i := strings.IndexRune(ct, ';'); i != -1 {
		ct = ct[0:i]
	}
	return ct == contentType
}
func contentTypeHandler(h http.Handler, contentTypes ...string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !(r.Method == "PUT" || r.Method == "POST" || r.Method == "PATCH") {
			h.ServeHTTP(w, r)
			return
		}

		for _, ct := range contentTypes {
			if isContentType(r.Header, ct) {
				h.ServeHTTP(w, r)
				return
			}
		}
		httpUtils.NewErrorResponse(w, http.StatusUnsupportedMediaType, "Unsupported content type")
	})
}
