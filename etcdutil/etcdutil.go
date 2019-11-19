package etcdutil

import (
	"time"

	"github.com/liuyuexclusive/utils/appconfigutil"
	"github.com/sirupsen/logrus"

	"go.etcd.io/etcd/clientv3"
)

// Open 操作etcd kv
func Open(fn func(kv clientv3.KV) error) error {
	config := clientv3.Config{
		Endpoints:   []string{appconfigutil.MustGet().ETCDAddress},
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
