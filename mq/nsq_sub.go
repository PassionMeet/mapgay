package mq

import (
	"context"
	"encoding/json"

	"github.com/cmfunc/jipeng/cache"
	"github.com/cmfunc/jipeng/model"
	"github.com/nsqio/go-nsq"
)

var consumer *nsq.Consumer

type userGeo_matchMan_MessageHandler struct{}

func (h *userGeo_matchMan_MessageHandler) HandleMessage(message *nsq.Message) error {
	msgBody := model.UploadGeoRequest{}
	err := json.Unmarshal(message.Body, &msgBody)
	if err != nil {
		return err
	}

	// 保存至redis
	_, err = cache.AddGeoPool(context.TODO(), &msgBody)

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
	err = consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		panic(err)
	}

}
