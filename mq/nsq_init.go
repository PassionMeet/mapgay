package mq

import (
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	InitConsumer()
	InitProducer()
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
