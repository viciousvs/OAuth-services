package oauth

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	authenticateAccess "github.com/viciousvs/OAuth-services/gateway/handler/oauth/authenticateAccessToken"
	authenticateRefresh "github.com/viciousvs/OAuth-services/gateway/handler/oauth/authenticateRefreshToken"
	signIn "github.com/viciousvs/OAuth-services/gateway/handler/oauth/signIn"
	signUp "github.com/viciousvs/OAuth-services/gateway/handler/oauth/signUp"
	"github.com/viciousvs/OAuth-services/gateway/middlewares"
	"net/http"
	"os"
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
	router.Use(func(next http.Handler) http.Handler { return handlers.LoggingHandler(os.Stdout, next) })

	middleware := middlewares.NewMiddleware(r.oAuthServiceAddr)
	//blog := router.PathPrefix("/blog").Subrouter()
	//blog.Use(middleware.EnsureAuth)
	//blog.HandleFunc("", blogH.Home).Methods(http.MethodGet)
	apiv1 := router.PathPrefix("/v1").Subrouter()
	apiv1.Use(middlewares.JsonMiddleware)
	oauthRouter := apiv1.PathPrefix("/oauth").Subrouter()

	oauthRouter.Handle("/authentication", middleware.EnsureAuth(authenticateAccess.Handle)).Methods(http.MethodGet)

	rh := authenticateRefresh.NewHandler(r.oAuthServiceAddr)
	oauthRouter.HandleFunc("/refresh", rh.Handle).Methods(http.MethodPost)

	sinH := signIn.NewHandler(r.oAuthServiceAddr)
	oauthRouter.HandleFunc("/sign-in", sinH.Handle).Methods(http.MethodPost)

	supH := signUp.NewHandler(r.oAuthServiceAddr)
	oauthRouter.HandleFunc("/sign-up", supH.Handle).Methods(http.MethodPost)

	return &Router{Router: router, oAuthServiceAddr: r.oAuthServiceAddr}
}
