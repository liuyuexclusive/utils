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
		viper.Set("ETCDAddress", "192.168.0.102:2379")
		viper.Set("NatsAddress", "192.168.0.102:4222")
		viper.Set("ElasticURL", "http://192.168.0.102:9200")
		viper.Set("HostIP", "192.168.0.102")
		viper.Set("APIPort", "9000")
		viper.Set("ConnStr", "root:123456@tcp(192.168.0.102:3306)/future?charset=utf8&parseTime=True&loc=Local")
		viper.Set("RedisAddress", "192.168.0.102:6379")
		viper.Set("JaegerAddress", "192.168.0.102:6831")
		viper.WriteConfig()
	} else {
		viper.ReadInConfig()
	}

	var config Config

	config.Name = viper.GetString("Name")
	config.ETCDAddress = viper.GetString("ETCDAddress")
	config.NatsAddress = viper.GetString("NatsAddress")
	config.ElasticURL = viper.GetString("ElasticURL")
	config.HostIP = viper.GetString("HostIP")
	config.APIPort = viper.GetString("APIPort")
	config.ConnStr = viper.GetString("ConnStr")
	config.RedisAddress = viper.GetString("RedisAddress")
	config.JaegerAddress = viper.GetString("JaegerAddress")

	return &config
}
