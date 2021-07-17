package core

import (
	"github.com/go-redis/redis/v8"
	"github.com/mingolm/go-recharge/configs"
	"github.com/mingolm/go-recharge/pkg/core/driver"
	"github.com/mingolm/go-recharge/pkg/repo"
	"go.uber.org/zap"
	"sync"
)

type Service struct {
	RedisCache redis.Cmdable

	UserRepo    repo.User
	OrderRepo   repo.Order
	ArticleRepo repo.Article
	LeaveRepo   repo.Leave

	OrderDriver *driver.OrderDriver // 话单商
	ThirdDriver *driver.ThirdDriver // 四方

	Logger *zap.SugaredLogger
}

var serviceOnce sync.Once
var service *Service

func Instance() *Service {
	serviceOnce.Do(func() {
		service = &Service{
			RedisCache: NewRedisCache(configs.DefaultConfigs.RedisAddr, configs.DefaultConfigs.RedisPassword),

			UserRepo: repo.NewUserRepo(&repo.UserConfig{
				DB: mustNewGormDB(configs.DefaultConfigs.DatabaseDsn),
			}),
			OrderRepo: repo.NewOrderRepo(&repo.OrderConfig{
				DB: mustNewGormDB(configs.DefaultConfigs.DatabaseDsn),
			}),
			ArticleRepo: repo.NewArticleRepo(&repo.ArticleConfig{
				DB: mustNewGormDB(configs.DefaultConfigs.DatabaseDsn),
			}),
			OrderDriver: driver.NewOrderDriver(&driver.OrderDriverConfig{}),
			ThirdDriver: driver.NewThirdDriver(&driver.ThirdConfig{
				Key:              configs.DefaultConfigs.PAYKey,
				PageUrl:          configs.DefaultConfigs.PAYPageUrl,
				BGUrl:            configs.DefaultConfigs.PAYBGUrl,
				H5RemoteAddr:     configs.DefaultConfigs.PAYH5RemoteAddr,
				QRCodeRemoteAddr: configs.DefaultConfigs.PAYQRCodeRemoteAddr,
				CancelRemoteAddr: configs.DefaultConfigs.PAYCancelRemoteAddr,
				StatusRemoteAddr: configs.DefaultConfigs.PAYStatusRemoteAddr,
			}),
			Logger: zap.S(),
		}
	})
	return service
}
