package mq

import (
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func PubUserGeo(body []byte) error {
	return producer.Publish("user_geo", body)
}

func InitProducer() {
	nsqConf := nsq.NewConfig()
	var err error
	producer, err = nsq.NewProducer("127.0.0.1:4150", nsqConf)
	if err != nil {
		panic(err)
	}
}
