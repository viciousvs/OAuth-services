package oauth

import (
	"github.com/gorilla/mux"
	"github.com/viciousvs/OAuth-services/gateway/handler/home"
	authenticateRefresh "github.com/viciousvs/OAuth-services/gateway/handler/oauth/authenticateRefreshToken"
	signIn "github.com/viciousvs/OAuth-services/gateway/handler/oauth/signIn"
	signUp "github.com/viciousvs/OAuth-services/gateway/handler/oauth/signUp"
	"github.com/viciousvs/OAuth-services/gateway/middlewares"
	"net/http"
)

type Router struct {
	*mux.Router
	oAuthServiceAddr string
}

func NewRouter(oAuthServiceAddr string) *Router {
	return &Router{oAuthServiceAddr: oAuthServiceAddr}
}
func (r *Router) InitRoutes() *Router {
	mux := mux.NewRouter()

	middleware := middlewares.NewMiddleware(r.oAuthServiceAddr)
	blog := mux.PathPrefix("/blog").Subrouter()
	blog.Use(middleware.EnsureAuth)
	blog.HandleFunc("/", home.Home)
	rh := authenticateRefresh.NewHandler(r.oAuthServiceAddr)
	mux.HandleFunc("/refresh", rh.Handle).Methods(http.MethodGet, http.MethodPost)

	sinH := signIn.NewHandler(r.oAuthServiceAddr)
	mux.HandleFunc("/sign-in", sinH.Handle).Methods(http.MethodGet, http.MethodPost)

	supH := signUp.NewHandler(r.oAuthServiceAddr)
	mux.HandleFunc("/sign-up", supH.Handle).Methods(http.MethodGet, http.MethodPost)
	return &Router{Router: mux, oAuthServiceAddr: r.oAuthServiceAddr}
}
