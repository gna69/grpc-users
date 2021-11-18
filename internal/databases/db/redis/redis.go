package redis

import (
	"fmt"
	"github.com/gna69/grpc-users/config"
	"github.com/go-redis/cache/v8"
	"github.com/go-redis/redis/v8"
	"time"
)

func New(config config.RedisConfig) *cache.Cache {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"cache": fmt.Sprintf("%s:%d", config.Host, config.Port),
		},
	})

	appCache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})

	return appCache
}
