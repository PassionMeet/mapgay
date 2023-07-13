package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type UserGeo struct {
	Openid    string
	Latitude  float64 `json:"latitude"`  //纬度
	Longitude float64 `json:"longitude"` //经度
}

const UserGeoKeyPrefix = "user::geo" //user::geo::[openid]

// add user's geo into geohash cache
func AddUserGeo(ctx context.Context, geo *UserGeo) error {
	_, err := redisClient.GeoAdd(ctx, UserGeoKeyPrefix, &redis.GeoLocation{
		Name:      geo.Openid,
		Longitude: geo.Longitude,
		Latitude:  geo.Latitude,
		Dist:      0.0,
		GeoHash:   0,
	}).Result()
	if err != nil {
		return err
	}
	return nil
}
