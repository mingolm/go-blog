package errutil

import (
	"database/sql"
	"errors"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

func DBError(err error) error {
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, gorm.ErrRecordNotFound) || errors.Is(err, redis.Nil) {
		return ErrNotFound
	}
	return ErrInternal.Msg(err.Error())
}
