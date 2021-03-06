package repo

import (
	"context"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/errutil"
	"gorm.io/gorm"
)

type Article interface {
	Get(ctx context.Context, id uint64) (row *model.Article, err error)
	GetList(ctx context.Context, offset, limit int) (rows []*model.Article, err error)
	GetListByType(ctx context.Context, t model.ArticleType, offset, limit int) (rows []*model.Article, err error)
	Create(ctx context.Context, row *model.Article) (err error)
	Update(ctx context.Context, id uint64, row *model.Article) (err error)
	UpdateStatus(ctx context.Context, id uint64, status model.ArticleStatus) (err error)
	Delete(ctx context.Context, id uint64) (err error)

	// internal
	GetTotal(ctx context.Context) (total int64, err error)
	GetAllList(ctx context.Context) (rows []*model.Article, err error)
}

func NewArticleRepo(config *ArticleConfig) Article {
	return &article{
		config,
	}
}

type ArticleConfig struct {
	DB *gorm.DB
}

type article struct {
	*ArticleConfig
}

func (r *article) db(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx).Model(&model.Article{})
}

func (r *article) Get(ctx context.Context, id uint64) (row *model.Article, err error) {
	row = &model.Article{}
	err = r.db(ctx).Where("id=?", id).First(row).Error
	if err != nil {
		return nil, errutil.DBError(err)
	}
	return row, nil
}

func (r *article) GetList(ctx context.Context, offset, limit int) (rows []*model.Article, err error) {
	rows = make([]*model.Article, 0)
	err = r.db(ctx).Where("status=?", model.ArticleStatusNormal).Order("created_at desc").Offset(offset).Limit(limit).Find(&rows).Error
	if err != nil {
		return nil, errutil.DBError(err)
	}
	return rows, nil
}

func (r *article) GetListByType(ctx context.Context, t model.ArticleType, offset, limit int) (rows []*model.Article, err error) {
	rows = make([]*model.Article, 0)
	err = r.db(ctx).Where("type=?", t).Order("id desc").Offset(offset).Limit(limit).Find(&rows).Error
	if err != nil {
		return nil, errutil.DBError(err)
	}
	return rows, nil
}

func (r *article) Create(ctx context.Context, row *model.Article) (err error) {
	err = r.db(ctx).Create(row).Error
	if err != nil {
		return errutil.DBError(err)
	}
	return nil
}

func (r *article) Update(ctx context.Context, id uint64, row *model.Article) (err error) {
	err = r.db(ctx).Where("id=?", id).Updates(row).Error
	if err != nil {
		return errutil.DBError(err)
	}
	return nil
}

func (r *article) UpdateStatus(ctx context.Context, id uint64, status model.ArticleStatus) (err error) {
	err = r.db(ctx).Where("id=?", id).UpdateColumn("status", status).Error
	if err != nil {
		return errutil.DBError(err)
	}
	return nil
}

func (r *article) Delete(ctx context.Context, id uint64) (err error) {
	err = r.db(ctx).Where("id=?", id).Delete(&model.Article{}).Error
	if err != nil {
		return errutil.DBError(err)
	}
	return nil
}

type GetTotalsOutput struct {
	TotalNormal int64
	TotalUp     int64
	TotalHide   int64
}

func (r *article) GetTotal(ctx context.Context) (total int64, err error) {
	err = r.db(ctx).Where("status=?", model.ArticleStatusNormal).Count(&total).Error
	if err != nil {
		return 0, errutil.DBError(err)
	}
	return total, nil
}

func (r *article) GetAllList(ctx context.Context) (rows []*model.Article, err error) {
	rows = make([]*model.Article, 0)
	err = r.db(ctx).Where("status=?", model.ArticleStatusNormal).Order("created_at desc").Find(&rows).Error
	if err != nil {
		return nil, errutil.DBError(err)
	}
	return rows, nil
}
