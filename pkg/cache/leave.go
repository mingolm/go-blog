package cache

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/utils/errutil"
	"strconv"
)

type LeaveCache interface {
	GetTotal(ctx context.Context) (total uint64, err error)
}

func NewLeaveCache() LeaveCache {
	return &leaveCache{
		Service:       core.Instance(),
		totalCacheKey: "mingo:leave-total",
	}
}

type leaveCache struct {
	*core.Service
	totalCacheKey string
}

func (c *leaveCache) GetTotal(ctx context.Context) (total uint64, err error) {
	val, err := c.RedisCache.Get(ctx, c.totalCacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, errutil.DBError(err)
	}

	total, _ = strconv.ParseUint(val, 10, 64)
	return total, nil
}
