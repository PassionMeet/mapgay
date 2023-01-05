package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UsersGeoCollDoc_Location struct {
	Type        string     `bson:"type"`
	Coordinates [2]float64 `bson:"coordinates"`
}

type UsersGeoCollDoc struct {
	ID       primitive.ObjectID        `bson:"id"`
	Openid   string                    `bson:"openid"`
	Location *UsersGeoCollDoc_Location `bson:"location"`
}

// InsertUsersGeo 保存用户的geo
func InsertUsersGeo(ctx context.Context, document interface{}) error {
	userGeoColl := mongoCli.Database("jipeng").Collection("users_geo")
	_, err := userGeoColl.InsertOne(ctx, document)
	return err
}

// SearchUsersByGeo 通过经纬度和范围获取用户openid
func SearchUsersByGeo(ctx context.Context) {

}
