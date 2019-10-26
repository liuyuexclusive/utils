package configutil

import (
	"encoding/json"
	"io/ioutil"

	"github.com/sirupsen/logrus"
)

type Config struct {
	ETCDAddress   string
	ConsulAddress string
	NatsAddress   string
	HostIP        string
	APIPort       string
	ElasticURL    string
	ConnStr       string
	RedisAddress  string
}

func MustGet() *Config {
	config, err := Get()
	if err != nil {
		logrus.Fatal(err)
		return nil
	}
	return config
}

func Get() (*Config, error) {
	byets, err := ioutil.ReadFile("appconfig.json")
	if err != nil {
		return nil, err
	}

	var config Config

	err = json.Unmarshal(byets, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
