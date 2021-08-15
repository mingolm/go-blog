package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/errutil"
	"github.com/mingolm/go-recharge/utils/helputil"
)

type ArticleCache interface {
	Get(ctx context.Context, id uint64) (row *model.Article, err error)
	GetList(ctx context.Context) (rows []*model.Article, err error)
	GetTotals(ctx context.Context) (output *GetTotalsOutput, err error)
	RefreshAll(ctx context.Context) (err error)
}

func NewArticleCache() ArticleCache {
	return &articleCache{
		Service:             core.Instance(),
		NormalTotalCacheKey: "article:normal-total", // 普通文章总数
		UpTotalCacheKey:     "article:up-total",     // 置顶文章总数
		HideTotalCacheKey:   "article:hide-total",   // 隐藏文章总数

		DetailCacheKey: "article:detail", // 归档列表
		ListCacheKey:   "article:list",   // 归档列表
	}
}

type articleCache struct {
	*core.Service

	NormalTotalCacheKey string
	UpTotalCacheKey     string
	HideTotalCacheKey   string

	DetailCacheKey string
	ListCacheKey   string
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

func (c *articleCache) GetTotals(ctx context.Context) (output *GetTotalsOutput, err error) {
	result, err := c.RedisCache.MGet(ctx, c.NormalTotalCacheKey, c.UpTotalCacheKey, c.HideTotalCacheKey).Result()
	if err != nil {
		return nil, errutil.DBError(err)
	}

	return &GetTotalsOutput{
		TotalNormal: helputil.Interface2Uint64(result[0]),
		TotalUp:     helputil.Interface2Uint64(result[1]),
		TotalHide:   helputil.Interface2Uint64(result[2]),
	}, nil
}

func (c *articleCache) RefreshAll(ctx context.Context) (err error) {
	c.RedisCache.Del(ctx, c.ListCacheKey, c.NormalTotalCacheKey, c.UpTotalCacheKey, c.HideTotalCacheKey)

	output, err := c.ArticleRepo.GetTotals(ctx)
	if err != nil {
		return err
	}
	if err = c.RedisCache.MSet(ctx, c.NormalTotalCacheKey, output.TotalNormal,
		c.UpTotalCacheKey, output.TotalUp,
		c.HideTotalCacheKey, output.TotalHide,
	).Err(); err != nil {
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
