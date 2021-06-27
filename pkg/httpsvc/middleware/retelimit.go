package middleware

import (
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"net/http"
)

var ReteLimit = func(next Handle) Handle {
	return func(r *http.Request) (response.Response, error) {
		return next(r)
	}
}
