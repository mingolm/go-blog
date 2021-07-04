package httpclient

import "sync"

var ps *ProxyService
var psOnce sync.Once

func ProxyInstance() *ProxyService {
	psOnce.Do(func() {
		ps = &ProxyService{}
	})
	return ps
}

type ProxyService struct {
	*ProxyConfig
}

type ProxyConfig struct {
}

func (h *ProxyService) GetProxy() (proxy string) {
	return
}
