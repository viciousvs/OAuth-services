package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/viciousvs/OAuth-services/gateway/config"
	oauthRoute "github.com/viciousvs/OAuth-services/gateway/route/oauth"
	http "github.com/viciousvs/OAuth-services/gateway/server/http"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	// loads values from gateway.env into the system
	if err := godotenv.Load("gateway.env"); err != nil {
		log.Print("No gateway.env file found")
	}
}
func main() {
	cfg := config.NewConfig()
	r := oauthRoute.NewRouter(os.Getenv("OAUTH_SERVER_ADDR")).InitRoutes()
	srv := new(http.Server)
	go func() {
		log.Printf("server started, port:%s", cfg.ServerConfig.Addr)
		if err := srv.Run(cfg.ServerConfig, r); err != nil {
			log.Fatalf("Cannot run server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Printf("Server shutting down")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
}
