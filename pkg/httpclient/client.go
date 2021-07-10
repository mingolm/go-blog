package httpclient

import "github.com/mingolm/go-recharge/utils/httputil"

func DHC() *httputil.HTTPClient {
	return httputil.NewHTTPClient(&httputil.HTTPClientConfig{})
}

func DHCP() *httputil.HTTPClient {
	ins := ProxyInstance()
	return httputil.NewHTTPClient(&httputil.HTTPClientConfig{}).Proxy(ins.GetProxy())
}
