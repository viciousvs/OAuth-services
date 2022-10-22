package oauth

import (
	"github.com/gorilla/mux"
	blogH "github.com/viciousvs/OAuth-services/gateway/handler/blog"
	home "github.com/viciousvs/OAuth-services/gateway/handler/home"
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
	router := mux.NewRouter()

	router.HandleFunc("/", home.Home)

	middleware := middlewares.NewMiddleware(r.oAuthServiceAddr)
	blog := router.PathPrefix("/blog").Subrouter()
	blog.Use(middleware.EnsureAuth)
	blog.HandleFunc("/", blogH.Home).Methods(http.MethodGet)

	rh := authenticateRefresh.NewHandler(r.oAuthServiceAddr)
	router.Handle("/refresh", middlewares.JsonMiddleware(rh.Handle)).Methods(http.MethodPost)

	sinH := signIn.NewHandler(r.oAuthServiceAddr)
	router.Handle("/sign-in", middlewares.JsonMiddleware(sinH.Handle)).Methods(http.MethodPost)

	supH := signUp.NewHandler(r.oAuthServiceAddr)
	router.Handle("/sign-up", middlewares.JsonMiddleware(supH.Handle)).Methods(http.MethodPost)

	return &Router{Router: router, oAuthServiceAddr: r.oAuthServiceAddr}
}
