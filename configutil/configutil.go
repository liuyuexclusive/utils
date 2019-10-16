package configutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

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
	bytes, err := ioutil.ReadFile("init.json")
	if err != nil {
		return nil, err
	}
	var config Config
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	client, err := api.NewClient(&api.Config{Scheme: "http", Address: config.ConsulAddress})
	if err != nil {
		return nil, err
	}
	return client, nil
}

func set(c *Config) error {
	fmt.Println(1111, c)
	fmt.Println(c == nil)
	var bytes []byte
	var err error
	if c == nil {
		bytes, err = ioutil.ReadFile("init.json")
		if err != nil {
			return err
		}
	} else {
		bytes, err = json.Marshal(c)
		if err != nil {
			return err
		}
	}

	kvPair := &api.KVPair{Key: "future", Value: bytes}

	client, err := client()
	if err != nil {
		return err
	}

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
		set(nil)
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
