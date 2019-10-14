package configutil

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
)

type Config struct {
	ConsulAddress string
	NatsAddress   string
	HostIP        string
	Port          string
	APIPort       string
	ElasticURL    string
	ConnStr       string
	RedisAddress  string
}

func client() (*api.Client, error) {
	client, err := api.NewClient(&api.Config{Scheme: "http", Address: fmt.Sprintf("%s:%s", "172.16.210.250", "8500")})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func Set(c *Config) error {
	client, err := client()
	if err != nil {
		return err
	}
	bytes, err := json.Marshal(c)
	if err != nil {
		return err
	}
	kvPair := &api.KVPair{Key: "future", Value: bytes}
	client.KV().Put(kvPair, nil)
	return nil
}

func MustGet() *Config {
	config, err := Get()
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
	return config
}

func Get() (*Config, error) {
	client, err := client()
	if err != nil {
		return nil, err
	}
	kvPair, _, err := client.KV().Get("future", nil)
	if err != nil {
		return nil, err
	}

	if kvPair == nil {
		//172.16.210.250
		Set(&Config{
			ConsulAddress: "172.16.210.250:8500",
			NatsAddress:   "172.16.210.250:4222",
			HostIP:        "172.16.210.250",
			Port:          "9090",
			APIPort:       "9000",
			ElasticURL:    fmt.Sprintf("http://%s:9200", "172.16.210.250"),
			ConnStr:       fmt.Sprintf("root:123456@tcp(%s:3306)/future?charset=utf8&parseTime=True&loc=Local", "172.16.210.250"),
			RedisAddress:  fmt.Sprintf("%s:6379", "172.16.210.250"),
		})
	}

	kvPair, _, err = client.KV().Get("future", nil)
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(kvPair.Value, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
