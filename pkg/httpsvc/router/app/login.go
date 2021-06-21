package app

import (
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"net/http"
)

type Login struct {
}

func (s *Login) Routers() router.Routers {
	return []router.Router{
		{
			Path:    "/login",
			Handler: s.Login,
			Method:  "GET",
		},
	}
}

func (s *Login) Login(req *http.Request) (resp response.Response, err error) {
	return response.Html("index", "123"), nil
}
