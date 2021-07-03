package repo

import (
	"context"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/errutil"
	"github.com/mingolm/go-recharge/utils/helputil"
	"gorm.io/gorm"
)

type User interface {
	GetForLogin(ctx context.Context, username, password string) (row *model.User, err error)
	GetByID(ctx context.Context, id uint64) (row *model.User, err error)
	Create(ctx context.Context, row *model.User) (err error)
	DeleteByID(ctx context.Context, id uint64) (err error)
}

type UserConfig struct {
	DB *gorm.DB
}

func NewUserRepo(config *UserConfig) User {
	return &user{
		config,
	}
}

type user struct {
	*UserConfig
}

func (r *user) db(ctx context.Context) *gorm.DB {
	return r.DB.WithContext(ctx)
}

func (r *user) GetForLogin(ctx context.Context, username, password string) (row *model.User, err error) {
	row = &model.User{}
	err = r.db(ctx).Where("username=? and password=?", username, helputil.EncryptPassword(password)).First(row).Error
	if err != nil {
		return nil, errutil.DBError(err)
	}
	return row, nil
}

func (r *user) GetByID(ctx context.Context, id uint64) (row *model.User, err error) {
	row = &model.User{}
	err = r.db(ctx).Where("id=?", id).First(row).Error
	if err != nil {
		return nil, errutil.DBError(err)
	}
	return row, nil
}

func (r *user) Create(ctx context.Context, row *model.User) (err error) {
	row.Password = helputil.EncryptPassword(row.Password)
	if err := r.db(ctx).Create(row).Error; err != nil {
		return errutil.DBError(err)
	}
	return nil
}

func (r *user) DeleteByID(ctx context.Context, id uint64) (err error) {
	err = r.db(ctx).Where("id=?", id).Delete(&model.User{}).Error
	if err != nil {
		return errutil.DBError(err)
	}
	return nil
}
