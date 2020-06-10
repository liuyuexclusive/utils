package appconfig

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type Config struct {
	Name          string `yaml:"Name"`
	ETCDAddress   string `yaml:"ETCDAddress"`
	NatsAddress   string `yaml:"NatsAddress"`
	HostIP        string `yaml:"HostIP"`
	APIPort       string `yaml:"APIPort"`
	ElasticURL    string `yaml:"ElasticURL"`
	ConnStr       string `yaml:"ConnStr"`
	RedisAddress  string `yaml:"RedisAddress"`
	JaegerAddress string `yaml:"JaegerAddress"`
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
	bytes, err := ioutil.ReadFile("appconfig.yml")
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(bytes, &config)

	if err != nil {
		return nil, fmt.Errorf("读取appconfig.yml失败：%w", err)
	}
	return &config, nil
}
