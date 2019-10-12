package redisutil

import (
	"future/util/configutil"

	"github.com/go-redis/redis"
)

// Open 打开redis
func Open(f func(client *redis.Client) error) error {
	client := redis.NewClient(&redis.Options{
		Addr:     configutil.MustGet().RedisAddress,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()
	if err := f(client); err != nil {
		return err
	}
	return nil
}
