package cache

import (
	"context"

	"github.com/cmfunc/jipeng/model"
	"github.com/go-redis/redis/v8"
)

func AddGeoPool(ctx context.Context, userLocation *model.UploadGeoRequest) (int64, error) {
	usergeo := &redis.GeoLocation{
		Name:      userLocation.Openid,
		Longitude: userLocation.Longitude,
		Latitude:  userLocation.Latitude,
	}
	return redisClient.GeoAdd(ctx, "user_geo_pool", usergeo).Result()
}
