package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UsersGeoCollDoc_Location struct {
	Type        string     `bson:"type"`
	Coordinates [2]float64 `bson:"coordinates"`
}

type UsersGeoCollDoc struct {
	ID       primitive.ObjectID        `bson:"id"`
	Openid   string                    `bson:"openid"`
	Location *UsersGeoCollDoc_Location `bson:"location"`
	UploadTs int64                     `bson:"upload_ts"`
}

// InsertUsersGeo 保存用户的geo
func InsertUsersGeo(ctx context.Context, document interface{}) error {
	userGeoColl := mongoCli.Database("jipeng").Collection("users_geo")
	_, err := userGeoColl.InsertOne(ctx, document)
	return err
}

// InsertUserCurrentGeo 保存用户的geo
func InsertUserCurrentGeo(ctx context.Context, openid string, longitude, latitude float64) error {
	userGeoColl := mongoCli.Database("jipeng").Collection("user_current_geo")
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{Key: "openid", Value: openid}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "location.type", Value: "Point"},
		{Key: "location.coordinates", Value: []float64{longitude, latitude}},
		{Key: "upload_ts", Value: time.Now().UnixMilli()},
	}}}
	_, err := userGeoColl.UpdateOne(ctx, filter, update, opts)
	return err
}

// SearchUsersByGeo 通过经纬度和范围获取用户openid
func SearchUsersByGeo(ctx context.Context, filter interface{}) ([]*UsersGeoCollDoc, error) {
	userGeoColl := mongoCli.Database("jipeng").Collection("user_current_geo")
	cur, err := userGeoColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	documents := make([]*UsersGeoCollDoc, 0)
	return documents, cur.All(ctx, &documents)
}
