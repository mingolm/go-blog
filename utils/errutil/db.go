package errutil

import (
	"database/sql"
	"errors"
	"gorm.io/gorm"
)

func DBError(err error) error {
	if errors.Is(err, sql.ErrNoRows) || errors.Is(err, gorm.ErrRecordNotFound) {
		return ErrNotFound
	}
	return ErrInternal.Msg(err.Error())
}
