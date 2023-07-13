package cache

import (
	"context"
	"fmt"

	"github.com/cmfunc/jipeng/conf"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func Init(cfg *conf.Redis) {
	addr:=fmt.Sprintf("%s:%d",cfg.Host,cfg.Port)
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Auth,
		DB:       0,
	})
	pong, err := redisClient.Ping(context.TODO()).Result()
	if err != nil {
		panic(err)
	}
	println(pong)
}
