package config

import (
	"log"

	"github.com/spf13/viper"
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
	LogPath       string `yaml:"LogPath"`
	LogLevel      string `yaml:"LogLevel"`
}

func MustGet() *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	var config Config

	viper.Unmarshal(&config)

	return &config
}
