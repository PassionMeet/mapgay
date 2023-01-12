package mq

import (
	"fmt"

	"github.com/cmfunc/jipeng/conf"
	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func PubUserGeo(body []byte) error {
	return producer.Publish("user_geo", body)
}

func InitProducer(cfg *conf.Server) {
	nsqConf := nsq.NewConfig()
	var err error
	uri := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	producer, err = nsq.NewProducer(uri, nsqConf)
	if err != nil {
		panic(err)
	}
}
