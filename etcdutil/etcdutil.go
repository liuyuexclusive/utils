package etcdutil

import (
	"context"
	"fmt"
	"time"

	"github.com/liuyuexclusive/utils/configutil"

	"go.etcd.io/etcd/clientv3"
)

func kv() (clientv3.KV, error) {
	config := clientv3.Config{
		Endpoints:   []string{configutil.MustGet().ETCDAddress},
		DialTimeout: 10 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		return nil, err
	}
	defer client.Close()

	kv := clientv3.NewKV(client)

	return kv, nil
}

func test() {
	kv, err := kv()

	if err != nil {
		panic(err)
	}

	kv.Put(context.TODO(), "test", "aaa")

	fmt.Println(kv.Get(context.TODO(), "test"))
}
