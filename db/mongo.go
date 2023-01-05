package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCli *mongo.Client

func Disconnect(ctx context.Context) {
	mongoCli.Disconnect(ctx)
}

// mongo中存储用户的位置信息
func InitMongo() {
	uri := "mongodb://localhost:27017"
	var err error
	mongoCli, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}
