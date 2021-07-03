package app

import (
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"net/http"
)

func NewApp() *App {
	return &App{
		core.Instance(),
	}
}

type App struct {
	*core.Service
}

func (s *App) Routers() router.Routers {
	return []router.Router{
		{
			Path:    "/index",
			Handler: s.index,
			Method:  "GET",
		},
	}
}

func (s *App) Middlewares() []middleware.Middleware {
	return []middleware.Middleware{middleware.Authentication}
}

type Order struct {
	ID string
}

func (s *App) index(req *http.Request) (resp response.Response, err error) {
	return response.Html("index"), nil
}
