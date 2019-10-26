package consulutil

import (
	"github.com/liuyuexclusive/utils/configutil"
	"github.com/sirupsen/logrus"

	"github.com/hashicorp/consul/api"
)

func Client() (*api.Client, error) {
	client, err := api.NewClient(&api.Config{Scheme: "http", Address: configutil.MustGet().ConsulAddress})
	return client, err
}

func GetValue(key string) ([]byte, error) {
	client, err := Client()
	if err != nil {
		return nil, err
	}

	kvPair, _, err := client.KV().Get(key, nil)
	if err != nil {
		return nil, err
	}

	return kvPair.Value, nil
}

func SetValue(key string, bytes []byte) error {
	client, err := Client()

	if err != nil {
		logrus.Error(err)
		return err
	}

	kvPair := &api.KVPair{Key: key, Value: bytes}

	_, err = client.KV().Put(kvPair, nil)

	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
