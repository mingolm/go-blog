package cache

import (
	"context"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/utils/errutil"
	"github.com/mingolm/go-recharge/utils/helputil"
	"time"
)

type ArticleCache interface {
	GetList(ctx context.Context) (output []*ArticleInfo, err error)
	GetTotals(ctx context.Context) (output *GetTotalsOutput, err error)
	Refresh(ctx context.Context) (err error)
}

func NewArticleCache() ArticleCache {
	return &articleCache{
		Service:             core.Instance(),
		NormalTotalCacheKey: "mingo:article-normal-total", // 普通文章总数
		UpTotalCacheKey:     "mingo:article-up-total",     // 置顶文章总数
		HideTotalCacheKey:   "mingo:article-hide-total",   // 隐藏文章总数

		ListCacheKey: "mingo:article-list", // 归档列表
	}
}

type articleCache struct {
	*core.Service

	NormalTotalCacheKey string
	UpTotalCacheKey     string
	HideTotalCacheKey   string

	ListCacheKey string
}

type ArticleInfo struct {
	CreatedAt time.Time
	Title     string
}

func (c *articleCache) GetList(ctx context.Context) (output []*ArticleInfo, err error) {
	zs, err := c.RedisCache.ZRevRangeWithScores(ctx, c.ListCacheKey, 0, -1).Result()
	if err != nil {
		return nil, errutil.DBError(err)
	}
	for _, v := range zs {
		output = append(output, &ArticleInfo{
			CreatedAt: time.Unix(int64(v.Score), 0),
			Title:     v.Member.(string),
		})
	}

	return output, nil
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

func (c *articleCache) Refresh(ctx context.Context) (err error) {
	output, err := c.ArticleRepo.GetTotals(ctx)
	if err != nil {
		return errutil.DBError(err)
	}
	if err = c.RedisCache.MSet(ctx,
		c.NormalTotalCacheKey, output.TotalNormal,
		c.UpTotalCacheKey, output.TotalUp,
		c.HideTotalCacheKey, output.TotalHide,
	).Err(); err != nil {
		return errutil.DBError(err)
	}
	return nil
}
