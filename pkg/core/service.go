package core

import (
	"github.com/mingolm/go-recharge/configs"
	"github.com/mingolm/go-recharge/pkg/repo"
	"go.uber.org/zap"
	"sync"
)

type Service struct {
	UserRepo repo.User

	Logger *zap.SugaredLogger
}

var serviceOnce sync.Once
var service *Service

func Instance() *Service {
	serviceOnce.Do(func() {
		service = &Service{
			UserRepo: repo.NewUserRepo(&repo.UserConfig{
				DB: mustNewGormDB(configs.DefaultConfigs.DatabaseDsn),
			}),
			Logger: zap.S(),
		}
	})
	return service
}
