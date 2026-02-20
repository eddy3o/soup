package server

import (
	"os"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis() *Redis {
	addr := os.Getenv("REDIS_ADDR")
	if addr == "" {
		addr = "localhost:6379"
	}
	rdb := redis.NewClient(&redis.Options{Addr: addr})
	return &Redis{Client: rdb}
}
