package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	IP          string
	Port        string
	Address     string
	Name        string
	ETCDAddress []string
	NatsAddress string

	ElasticURL      string
	ConnStr         string
	RedisAddress    string
	JaegerAddress   string
	LogPath         string
	LogLevel        string
	TLS             ConfigTLS
	AuthServiceName string
}

type ConfigTLS struct {
	CertFile           string
	KeyFile            string
	CACertFile         string
	ServerNameOverride string
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
