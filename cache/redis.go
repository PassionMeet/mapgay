package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func Init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "123456",
		DB:       0,
	})
	pong, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	println(pong)
}
