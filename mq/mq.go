package mq

import (
	"github.com/nats-io/nats.go"
	"github.com/yuexclusive/utils/config"
)

func Publish(subj string, bytes []byte) error {
	nc, err := nats.Connect(config.MustGet().NatsAddress)
	if err != nil {
		return err
	}
	defer nc.Close()
	return nc.Publish(subj, bytes)
	// return nc.Flush()
}

func Subscribe(subj string, handler nats.MsgHandler) error {
	nc, err := nats.Connect(config.MustGet().NatsAddress)
	if err != nil {
		return err
	}
	// defer nc.Close()
	nc.Subscribe(subj, handler)
	return nil
}
