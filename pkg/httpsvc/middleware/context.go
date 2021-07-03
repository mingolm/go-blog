package middleware

import (
	"context"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"net"
	"net/http"
)

var Context = func(next Handle) Handle {
	return func(r *http.Request) (response.Response, error) {
		ctx := context.WithValue(r.Context(), "ip", getIPFromRequest(r))
		return next(r.WithContext(ctx))
	}
}

func getIPFromRequest(r *http.Request) string {
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}
	userIP := net.ParseIP(ip)
	if userIP == nil {
		return ""
	}
	return userIP.String()
}
