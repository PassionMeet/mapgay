package cache

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
)

type UserGeo struct {
	Openid    string
	Latitude  float64 `json:"latitude"`  //纬度
	Longitude float64 `json:"longitude"` //经度
}

const UserGeoKeyPrefix = "user::geo" //user::geo::[openid]

// add user's geo into geohash cache
func AddUserGeo(ctx context.Context, geo *UserGeo) error {
	location := &redis.GeoLocation{
		Name:      geo.Openid,
		Longitude: geo.Longitude,
		Latitude:  geo.Latitude,
		Dist:      0.0,
		GeoHash:   0,
	}
	_, err := redisClient.GeoAdd(ctx, UserGeoKeyPrefix, location).Result()
	if err != nil {
		return err
	}
	return nil
}

type GeoFilter struct {
	Openid string
}

type CacheUserGeo struct {
	Longitude, Latitude, Dist float64
}

func GetUsersByGeo(ctx context.Context, filter *GeoFilter) (usergeos map[string]*CacheUserGeo, err error) {
	query := &redis.GeoRadiusQuery{
		Radius:      1000,
		Unit:        "m",
		WithCoord:   false,
		WithDist:    false,
		WithGeoHash: false,
		Count:       0,
		Sort:        "ASC",
		Store:       "",
		StoreDist:   "",
	}
	users, err := redisClient.GeoRadiusByMember(ctx, UserGeoKeyPrefix, filter.Openid, query).Result()
	if err != nil {
		return nil, errors.Wrapf(err, "GetUsersByGeo filter:%+v", filter)
	}
	usergeos = make(map[string]*CacheUserGeo, 0)
	for _, user := range users {
		usergeos[user.Name] = &CacheUserGeo{
			Longitude: user.Longitude,
			Latitude:  user.Latitude,
			Dist:      user.Dist,
		}
	}
	return usergeos, nil
}
