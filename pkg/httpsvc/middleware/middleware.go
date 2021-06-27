package middleware

import (
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"net/http"
)

type Middleware func(next Handle) Handle

type Handle func(r *http.Request) (response.Response, error)
