package etcd

import (
	"time"

	"github.com/liuyuexclusive/utils/appconfig"
	"github.com/sirupsen/logrus"

	"go.etcd.io/etcd/clientv3"
)

type KV interface {
	clientv3.KV
}

// Open 操作etcd kv
func Open(fn func(kv KV) error) error {
	config := clientv3.Config{
		Endpoints:   []string{appconfig.MustGet().ETCDAddress},
		DialTimeout: 10 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	defer client.Close()

	kv := clientv3.NewKV(client)

	err = fn(kv)

	if err != nil {
		logrus.Error(err)
		return err
	}

	return nil
}
