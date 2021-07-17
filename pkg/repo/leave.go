package repo

import (
	"context"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/errutil"
	"gorm.io/gorm"
)

type Leave interface {
	GetList(ctx context.Context, offset, limit int) (rows []*model.Leave, err error)
}

func NewLeave(config *LeaveConfig) Leave {
	return &leave{
		config,
	}
}

type LeaveConfig struct {
	DB *gorm.DB
}

type leave struct {
	*LeaveConfig
}

func (r *leave) db(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx)
}

func (r *leave) GetList(ctx context.Context, offset, limit int) (rows []*model.Leave, err error) {
	rows = make([]*model.Leave, 0)
	err = r.db(ctx).Where("status=?", model.Leave{}).Order("id desc").Offset(offset).Limit(limit).Find(&rows).Error
	if err != nil {
		return nil, errutil.DBError(err)
	}
	return rows, nil
}
