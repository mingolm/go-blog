package httpclient

import (
	"github.com/mingolm/go-recharge/configs"
	"sync"
)

var ps *ProxyService
var psOnce sync.Once

func ProxyInstance() *ProxyService {
	psOnce.Do(func() {
		ps = &ProxyService{
			ProxyConfig: &ProxyConfig{
				ProxyIPUrl: configs.DefaultConfigs.ProxyIPUrl,
			},
		}
	})
	return ps
}

type ProxyService struct {
	*ProxyConfig
}

type ProxyConfig struct {
	ProxyIPUrl string
}

func (h *ProxyService) GetProxy() (proxy string) {
	return
}
