package core

import (
	"github.com/go-redis/redis/v8"
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
	})
	return redisCacheClient
}
