package app

import (
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/httpsvc/middleware"
	"github.com/mingolm/go-recharge/pkg/httpsvc/response"
	"github.com/mingolm/go-recharge/pkg/httpsvc/router"
	"github.com/mingolm/go-recharge/pkg/model"
	"github.com/mingolm/go-recharge/utils/errutil"
	"net/http"
	"strconv"
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
		{
			Path:    "/order",
			Handler: s.order,
			Method:  "POST",
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
	return response.Html("index", struct {
		Username string `json:"username"`
	}{
		Username: "123",
	}), nil
}

func (s *App) order(req *http.Request) (resp response.Response, err error) {
	orderID := req.FormValue("order_id")
	orderAmt, _ := strconv.ParseUint(req.FormValue("order_amt"), 10, 64)
	sourceID := req.FormValue("user_id")
	busCodeInt, err := strconv.ParseInt(req.FormValue("bus_code"), 10, 64)
	if err != nil {
		return nil, errutil.ErrInvalidArguments.Msg("bus_code")
	}
	busCode := model.BusCode(busCodeInt)

	row := &model.Order{
		SourceID: sourceID,
		OrderID:  orderID,
		OrderAmt: orderAmt,
		BusCode:  busCode,
		IP:       model.GetIPv4(req.Context().Value("ip").(string)),
	}
	// 创建内部订单
	if err := s.OrderRepo.Create(req.Context(), row); err != nil {
		return nil, err
	}

	return response.Html("index", nil), nil
}
