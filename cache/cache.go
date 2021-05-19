package cache

import (
	"github.com/yuexclusive/utils/config"

	"github.com/go-redis/redis"
	r "github.com/go-redis/redis"
)

func Client() *redis.Client {
	client := r.NewClient(&r.Options{
		Addr:     config.MustGet().RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return client
}
