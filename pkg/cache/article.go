package cache

import (
	"context"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/utils/errutil"
	"github.com/mingolm/go-recharge/utils/helputil"
)

func NewArticleCache() *Article {
	return &Article{
		Service:             core.Instance(),
		NormalTotalCacheKey: "mingo:article-normal-total",
		UpTotalCacheKey:     "mingo:article-up-total",
		HideTotalCacheKey:   "mingo:article-hide-total",
	}
}

type Article struct {
	*core.Service

	NormalTotalCacheKey string
	UpTotalCacheKey     string
	HideTotalCacheKey   string
}

type GetTotalsOutput struct {
	TotalNormal int64 `json:"total_normal"`
	TotalUp     int64 `json:"total_up"`
	TotalHide   int64 `json:"total_hide"`
}

func (c *Article) GetTotals(ctx context.Context) (output *GetTotalsOutput, err error) {
	result, err := c.RedisCache.MGet(ctx, c.NormalTotalCacheKey, c.UpTotalCacheKey, c.HideTotalCacheKey).Result()
	if err != nil {
		return nil, errutil.DBError(err)
	}

	return &GetTotalsOutput{
		TotalNormal: helputil.Interface2Int64(result[0]),
		TotalUp:     helputil.Interface2Int64(result[1]),
		TotalHide:   helputil.Interface2Int64(result[2]),
	}, nil
}

func (c *Article) Refresh(ctx context.Context) (err error) {
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
