package main

import (
	"github.com/joho/godotenv"
	"github.com/viciousvs/OAuth-services/user/config"
	"github.com/viciousvs/OAuth-services/user/storage/postgres"
	"log"
)

func init() {
	if err := godotenv.Load("user.env"); err != nil {
		log.Printf("cannot load .env file=> %v", err)
		log.Println("used default values for config")
	}
}
func main() {
	pgCfg := config.MakePostgresConfig()
	db := postgres.NewPostgresDB(pgCfg)
	defer db.Close()

}
