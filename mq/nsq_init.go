package mq

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/cmfunc/jipeng/conf"
)

func Init(cfg *conf.NSQ) {
	InitConsumer(cfg.Consumer)
	InitProducer(cfg.Producer)
	go func() {
		// wait for signal to exit
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
		<-sigChan

		// Gracefully stop the consumer.
		producer.Stop()
		consumer.Stop()
	}()
}
