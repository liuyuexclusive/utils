package appconfig

import (
	"os"

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
}

func MustGet() *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("appconfig.yml")
	if _, err := os.Stat("appconfig.yml"); err != nil {
		viper.Set("Name", "test")
		viper.Set("ETCDAddress", "192.168.1.100:2379")
		viper.Set("NatsAddress", "192.168.1.100:4222")
		viper.Set("ElasticURL", "http://192.168.1.100:9200")
		viper.Set("HostIP", "192.168.1.100")
		viper.Set("APIPort", "9000")
		viper.Set("ConnStr", "host=192.168.1.100 port=5432 password=123456 sslmode=disable user=postgres dbname=future")
		viper.Set("RedisAddress", "192.168.1.100:6379")
		viper.Set("JaegerAddress", "192.168.1.100:6831")
		viper.WriteConfig()
	} else {
		viper.ReadInConfig()
	}

	var config Config

	viper.Unmarshal(&config)

	return &config
}
