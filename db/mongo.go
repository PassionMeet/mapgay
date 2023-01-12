package db

import (
	"context"
	"fmt"

	"github.com/cmfunc/jipeng/conf"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCli *mongo.Client

func Disconnect(ctx context.Context) {
	mongoCli.Disconnect(ctx)
}

// mongo中存储用户的位置信息
func InitMongo(cfg *conf.MongoDB) {
	uri := fmt.Sprintf("mongodb://%s:%d", cfg.Host, cfg.Port)
	var err error
	mongoCli, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
}
