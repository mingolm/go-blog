package core

import (
	"context"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/mingolm/go-recharge/configs"
	gmysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	gmLogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func mustNewGormDB(dsn string) *gorm.DB {
	db, err := newGormDB(dsn)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func newGormDB(dsn string) (*gorm.DB, error) {
	parsedDSN, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, fmt.Errorf("database parse dsn fail")
	}

	gormConfig := &gorm.Config{}
	if !configs.IsProd() {
		gormConfig.Logger = gmLogger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			gmLogger.Config{
				SlowThreshold: time.Second,
				LogLevel:      gmLogger.Info,
				Colorful:      true,
			},
		)
	}

	db, err := gorm.Open(gmysql.Open(dsn), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("gorm open failed: %w", err)
	}

	if rawDb, err := db.DB(); err != nil {
		return nil, fmt.Errorf("gorm get DB failed: %w", err)
	} else {
		rawDb.SetConnMaxLifetime(time.Hour)
		rawDb.SetMaxIdleConns(32)
		rawDb.SetMaxOpenConns(128)

		pingCtx, cancel := context.WithTimeout(context.Background(), time.Second*5)
		err := rawDb.PingContext(pingCtx)
		cancel()

		if err != nil {
			return nil, fmt.Errorf("ping mysql fail: dbname=%s", parsedDSN.DBName)
		}
	}

	return db, nil
}
