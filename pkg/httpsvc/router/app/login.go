package app

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
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
			Handler: s.LoginTemplate,
			Method:  "GET",
			Middlewares: []middleware.Middleware{
				middleware.Authentication,
			},
		},
		{
			Path:    "/login",
			Handler: s.Login,
			Method:  "POST",
		},
	}
}

func (s *Login) Middlewares() []middleware.Middleware {
	return []middleware.Middleware{}
}

func (s *Login) LoginTemplate(req *http.Request) (resp response.Response, err error) {
	return response.Html("index", "123"), nil
}

func (s *Login) Login(req *http.Request) (resp response.Response, err error) {
	username := req.FormValue("username")
	password := req.FormValue("password")
	if username == "" || password == "" {
		return response.Error(fmt.Errorf("login: username or password is empty")), nil
	}
	return response.Redirect("index", 302), nil
}
