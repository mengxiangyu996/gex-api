package dal

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	Address  string
	Database int
	Password string
}

func initRedis(opts *redis.Options) {

	Redis = redis.NewClient(opts)

	if _, err := Redis.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}
}
