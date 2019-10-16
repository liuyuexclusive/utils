package configutil

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/hashicorp/consul/api"
)

type Config struct {
	ConsulAddress string
	NatsAddress   string
	HostIP        string
	APIPort       string
	ElasticURL    string
	ConnStr       string
	RedisAddress  string
}

type Consul struct {
	Address string
}

func Client() (*api.Client, error) {
	consulbytes, err := ioutil.ReadFile("consul.json")
	if err != nil {
		log.Fatal(err)
	}

	var consul Consul

	err = json.Unmarshal(consulbytes, &consul)
	if err != nil {
		log.Fatal(err)
	}

	client, err := api.NewClient(&api.Config{Scheme: "http", Address: consul.Address})
	return client, err
}

func MustGet() *Config {
	config, err := Get()
	if err != nil {
		log.Fatal(err)
	}
	return config
}

func Get() (*Config, error) {
	client, err := Client()
	if err != nil {
		return nil, err
	}

	kvPair, _, err := client.KV().Get("future", nil)
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
