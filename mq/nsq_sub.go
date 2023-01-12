package mq

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cmfunc/jipeng/conf"
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

func InitConsumer(cfg *conf.NSQ_Consumer) {
	nsqConf := nsq.NewConfig()
	var err error
	consumer, err = nsq.NewConsumer(cfg.Topic, cfg.Channel, nsqConf)
	if err != nil {
		panic(err)
	}
	consumer.AddConcurrentHandlers(&userGeo_matchMan_MessageHandler{}, 10)
	uri := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	err = consumer.ConnectToNSQD(uri)
	if err != nil {
		panic(err)
	}

}
