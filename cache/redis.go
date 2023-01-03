package cache

import (
	"context"

	"github.com/cmfunc/jipeng/model"
	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func AddGeoPool(ctx context.Context, userLocation *model.UploadGeoRequest) (int64, error) {
	usergeo := &redis.GeoLocation{
		Name:      userLocation.Openid,
		Longitude: userLocation.Longitude,
		Latitude:  userLocation.Latitude,
	}
	return redisClient.GeoAdd(ctx, "user_geo_pool", usergeo).Result()
}

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
