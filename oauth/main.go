package main

import (
	"github.com/joho/godotenv"
	"github.com/viciousvs/OAuth-services/oauth/config"
	"github.com/viciousvs/OAuth-services/oauth/server/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	if err := godotenv.Load("oauth.env"); err != nil {
		log.Printf("cannot load oauth.env file=> %v", err)
		log.Fatalf("used default values for config")
	}
}
func main() {
	sCfg := config.MakeServerConfig()
	srv := grpc.NewServer()

	go func() {
		log.Printf("GRPC server has been started, addr:%s", sCfg.Addr)
		if err := srv.Run(sCfg); err != nil {
			log.Fatalf("cannot run GRPC server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
