package app

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"net/http"
)

type App struct {
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
	return []middleware.Middleware{}
}

type Order struct {
	ID string
}

func (s *App) index(req *http.Request) (resp response.Response, err error) {
	orderID := req.FormValue("order_id")
	fmt.Printf("orderID: %s \n", orderID)
	return response.Data(&Order{
		ID: orderID,
	}), nil
}
