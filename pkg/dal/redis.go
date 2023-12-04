package dal

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func initRedis(opts *redis.Options) {

	RedisClient = redis.NewClient(opts)

	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return
}
