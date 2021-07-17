package core

import (
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"sync"
)

var redisCacheClient redis.Cmdable
var redisCacheOnce sync.Once

func NewRedisCache(addr, password string) redis.Cmdable {
	redisCacheOnce.Do(func() {
		redisCacheClient = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
		})

		zap.S().Infow("redis: new client success",
			"addr", addr,
		)
	})
	return redisCacheClient
}
