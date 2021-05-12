package etcd

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/yuexclusive/utils/config"

	etcd "go.etcd.io/etcd/client/v3"
)

type KV interface {
	etcd.KV
}

// Open 操作etcd kv
func Open(fn func(kv KV) error) error {
	config := etcd.Config{
		Endpoints:   []string{config.MustGet().ETCDAddress},
		DialTimeout: 10 * time.Second,
	}
	client, err := etcd.New(config)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer client.Close()

	kv := etcd.NewKV(client)

	err = fn(kv)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
