package repo

import (
	"context"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/errutil"
	"gorm.io/gorm"
)

type Order interface {
	GetWaitingList(ctx context.Context, id uint64, limit int) (rows []*model.Order, err error)
	Create(ctx context.Context, row *model.Order) (err error)
	SetSuccessByOrderID(ctx context.Context, orderID string) (err error)
	SetFailedByOrderID(ctx context.Context, orderID string) (err error)
}

type OrderConfig struct {
	DB *gorm.DB
}

func NewOrderRepo(config *OrderConfig) Order {
	return &order{
		config,
	}
}

type order struct {
	*OrderConfig
}

func (r *order) db(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx)
}

func (r *order) GetWaitingList(ctx context.Context, id uint64, limit int) (rows []*model.Order, err error) {
	rows = make([]*model.Order, 0)
	err = r.db(ctx).Where("id>? and status=?", id, model.OrderStatusWaiting).Order("id asc").Limit(limit).Find(&rows).Error
	if err != nil {
		return nil, errutil.DBError(err)
	}
	return rows, nil
}
func (r *order) Create(ctx context.Context, row *model.Order) (err error) {
	if err := r.db(ctx).Create(row).Error; err != nil {
		return errutil.DBError(err)
	}
	return nil
}
func (r *order) SetSuccessByOrderID(ctx context.Context, orderID string) (err error) {
	if err := r.db(ctx).Where("order_id=?", orderID).Update("status", model.OrderStatusSuccess).Error; err != nil {
		return errutil.DBError(err)
	}
	return nil
}
func (r *order) SetFailedByOrderID(ctx context.Context, orderID string) (err error) {
	if err := r.db(ctx).Where("order_id=?", orderID).Update("status", model.OrderStatusFailed).Error; err != nil {
		return errutil.DBError(err)
	}
	return nil
}
