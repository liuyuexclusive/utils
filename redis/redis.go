package redis

import (
	"github.com/liuyuexclusive/utils/appconfig"
	"github.com/sirupsen/logrus"

	r "github.com/go-redis/redis"
)

// Open 打开redis
func Open(f func(client *r.Client) error) error {
	client := r.NewClient(&r.Options{
		Addr:     appconfig.MustGet().RedisAddress,
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
