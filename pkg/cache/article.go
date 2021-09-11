package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/errutil"
	"strconv"
)

type ArticleCache interface {
	Get(ctx context.Context, id uint64) (row *model.Article, err error)
	GetList(ctx context.Context) (rows []*model.Article, err error)
	GetTotal(ctx context.Context) (total int64, err error)
	RefreshAll(ctx context.Context) (err error)
}

func NewArticleCache() ArticleCache {
	return &articleCache{
		Service:        core.Instance(),
		TotalCacheKey:  "article:total",
		DetailCacheKey: "article:detail",
		ListCacheKey:   "article:list",
	}
}

type articleCache struct {
	*core.Service
	TotalCacheKey  string // 文章总数
	DetailCacheKey string // 归档列表
	ListCacheKey   string // 归档列表
}

func (c *articleCache) Get(ctx context.Context, id uint64) (row *model.Article, err error) {
	result, err := c.RedisCache.Get(ctx, fmt.Sprintf("%s:%d", c.DetailCacheKey, id)).Result()
	if err != nil {
		return nil, errutil.DBError(err)
	}
	row = &model.Article{}
	if err = json.Unmarshal([]byte(result), row); err != nil {
		return nil, errutil.InternalError(err)
	}
	return row, nil
}

func (c *articleCache) GetList(ctx context.Context) (rows []*model.Article, err error) {
	result, err := c.RedisCache.LRange(ctx, c.ListCacheKey, 0, -1).Result()
	if err != nil {
		return nil, errutil.DBError(err)
	}
	for _, v := range result {
		row := &model.Article{}
		if err = json.Unmarshal([]byte(v), row); err != nil {
			return nil, errutil.InternalError(err)
		}
		rows = append(rows, row)
	}

	return rows, nil
}

type GetTotalsOutput struct {
	TotalNormal uint64 `json:"total_normal"`
	TotalUp     uint64 `json:"total_up"`
	TotalHide   uint64 `json:"total_hide"`
}

func (c *articleCache) GetTotal(ctx context.Context) (total int64, err error) {
	result, err := c.RedisCache.Get(ctx, c.TotalCacheKey).Result()
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return 0, nil
		}
		return 0, errutil.DBError(err)
	}

	total, err = strconv.ParseInt(result, 10, 64)
	if err != nil {
		return 0, errutil.InternalError(err)
	}

	return total, nil
}

func (c *articleCache) RefreshAll(ctx context.Context) (err error) {
	c.RedisCache.Del(ctx, c.ListCacheKey, c.TotalCacheKey)

	total, err := c.ArticleRepo.GetTotal(ctx)
	if err != nil {
		return err
	}
	if err = c.RedisCache.Set(ctx, c.TotalCacheKey, total, 0).Err(); err != nil {
		return errutil.InternalError(err)
	}

	rows, err := c.ArticleRepo.GetAllList(ctx)
	if err != nil {
		return err
	}

	msets := make([]interface{}, 0, len(rows)*2)

	for _, row := range rows {
		bs, err := json.Marshal(row)
		if err != nil {
			return errutil.DBError(err)
		}
		if err = c.RedisCache.RPush(ctx, c.ListCacheKey, string(bs)).Err(); err != nil {
			return errutil.DBError(err)
		}
		msets = append(msets, fmt.Sprintf("%s:%d", c.DetailCacheKey, row.ID), string(bs))
	}

	if err = c.RedisCache.MSet(ctx, msets...).Err(); err != nil {
		return errutil.DBError(err)
	}

	return nil
}
