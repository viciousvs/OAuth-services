package main

import (
	"github.com/joho/godotenv"
	"github.com/viciousvs/OAuth-services/user/config"
	"github.com/viciousvs/OAuth-services/user/model/user"
	"github.com/viciousvs/OAuth-services/user/server/grpc"
	"github.com/viciousvs/OAuth-services/user/storage/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	if err := godotenv.Load("user.env"); err != nil {
		log.Printf("cannot load user.env file=> %v", err)
		log.Println("used default values for config")
	}
}
func main() {
	pgCfg := config.MakePostgresConfig()
	db := postgres.NewPostgresDB(pgCfg)
	pgRepo := user.NewPostgresRepo(db)
	defer db.Close()

	sCfg := config.MakeServerConfig()
	srv := grpc.NewServer(pgRepo)

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
