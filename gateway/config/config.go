package config

import (
	"os"
	"time"
)

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); !exists {
		return value
	}
	return defaultValue
}

type Config struct {
	ServerConfig ServerConfig
}
type ServerConfig struct {
	Addr           string
	MaxHeaderbytes int
	ReadTimeout,
	WriteTimeout time.Duration
}

func NewConfig() *Config {
	return &Config{
		ServerConfig: ServerConfig{
			Addr:           getEnv("SERVER_ADDRESS", "localhost:8081"),
			MaxHeaderbytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}

}
