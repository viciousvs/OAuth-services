package redisRepo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/viciousvs/OAuth-services/token/config"
	"log"
	"sync"
	"time"
)

var redisDB *RedisDB
var initOnce sync.Once

type RedisDB struct {
	*redis.Client
}

func NewRedisDB(config config.RedisConfig) *RedisDB {
	initOnce.Do(func() {
		var err error
		redisDB, err = getRedisDB(config)
		if err != nil {
			log.Fatalf("cannot connect to postgres db: %v", err)
		}
	})

	return redisDB
}

func getRedisDB(cfg config.RedisConfig) (*RedisDB, error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	opt := &redis.Options{Addr: cfg.Addr, Password: cfg.Password, DB: cfg.DB}
	cl := redis.NewClient(opt)

	err := cl.Ping(ctx).Err()
	if err != nil {
		return nil, err
	}
	return &RedisDB{cl}, nil
}
