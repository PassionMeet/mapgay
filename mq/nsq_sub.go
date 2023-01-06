package mq

import (
	"context"
	"encoding/json"
	"time"

	"github.com/cmfunc/jipeng/db"
	"github.com/cmfunc/jipeng/model"
	"github.com/nsqio/go-nsq"
	"go.mongodb.org/mongo-driver/bson"
)

var consumer *nsq.Consumer

type userGeo_matchMan_MessageHandler struct{}

func (h *userGeo_matchMan_MessageHandler) HandleMessage(message *nsq.Message) error {
	msgBody := model.UploadGeoRequest{}
	err := json.Unmarshal(message.Body, &msgBody)
	if err != nil {
		return err
	}

	ctx := context.Background()
	// 保存到mongo
	document := bson.D{
		{
			Key: "location",
			Value: bson.D{
				{Key: "type", Value: "Point"},
				{Key: "coordinates", Value: bson.A{msgBody.Longitude, msgBody.Latitude}},
			},
		},
		{
			Key:   "openid",
			Value: msgBody.Openid,
		},
		{
			Key:   "upload_ts",
			Value: time.Now().UnixMilli(), //TODO 替换为前端记录的事件
		},
	}
	// 用户所有轨迹
	err = db.InsertUsersGeo(ctx, document)
	if err != nil {
		return err
	}
	// 用户当前最新轨迹（每个用户只有当前最新的记录）
	err = db.InsertUserCurrentGeo(ctx, msgBody.Openid, msgBody.Longitude, msgBody.Latitude)
	return err
}

func InitConsumer() {
	nsqConf := nsq.NewConfig()
	var err error
	consumer, err = nsq.NewConsumer("user_geo", "match_man", nsqConf)
	if err != nil {
		panic(err)
	}
	consumer.AddConcurrentHandlers(&userGeo_matchMan_MessageHandler{}, 10)
	err = consumer.ConnectToNSQD("127.0.0.1:4150")
	if err != nil {
		panic(err)
	}

}
