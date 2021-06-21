package app

import (
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"net/http"
)

type App struct {
}

func (s *App) Routers() router.Routers {
	return []router.Router{
		{
			Path:    "/login",
			Handler: s.Callback,
			Method:  "POST",
		},
	}
}

type Order struct {
	ID string
}

func (s *App) Callback(req *http.Request) (resp response.Response, err error) {
	orderID := req.FormValue("order_id")
	fmt.Printf("orderID: %s \n", orderID)
	return response.Data(&Order{
		ID: orderID,
	}), nil
}
