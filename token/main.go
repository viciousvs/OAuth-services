package main

import (
	"github.com/joho/godotenv"
	"github.com/viciousvs/OAuth-services/token/config"
	"github.com/viciousvs/OAuth-services/token/server/grpc"
	"github.com/viciousvs/OAuth-services/token/storage/redisRepo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	if err := godotenv.Load("token.env"); err != nil {
		log.Printf("cannot load .env file=> %v", err)
		log.Println("used default values for config")
	}
}
func main() {
	redisCfg := config.NewRedisConfig()
	db := redisRepo.NewRedisDB(redisCfg)
	defer db.Close()

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
