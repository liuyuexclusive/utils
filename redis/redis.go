package redis

import (
	"github.com/liuyuexclusive/utils/config"
	"github.com/sirupsen/logrus"

	r "github.com/go-redis/redis"
)

type Client struct {
	*r.Client
}

// Open 打开redis
func Open(f func(client *Client) error) error {
	client := r.NewClient(&r.Options{
		Addr:     config.MustGet().RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	newClient := &Client{Client: client}
	if err := f(newClient); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func GetClient() *Client {
	client := r.NewClient(&r.Options{
		Addr:     config.MustGet().RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &Client{Client: client}
}
