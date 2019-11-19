package redisutil

import (
	"github.com/liuyuexclusive/utils/appconfigutil"
	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis"
)

// Open 打开redis
func Open(f func(client *redis.Client) error) error {
	client := redis.NewClient(&redis.Options{
		Addr:     appconfigutil.MustGet().RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	if err := f(client); err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
