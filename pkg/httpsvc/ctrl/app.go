package ctrl

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
		{ // 创建四方订单（h5）
			Path:    "/order",
			Handler: s.order,
			Method:  "POST",
		},
		{ // 创建四方订单（二维码）
			Path:    "/order-for-qrcode",
			Handler: s.orderForQRCode,
			Method:  "POST",
		},
		{ // 取消订单
			Path:    "/cancel-order",
			Handler: s.cancelOrder,
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

type CreateOrderOutput struct {
	OrderID string `json:"order_id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *App) order(req *http.Request) (resp response.Response, err error) {
	orderID := req.FormValue("order_id")
	orderAmt, _ := strconv.ParseFloat(req.FormValue("order_amt"), 10)
	sourceID := req.FormValue("user_id")
	busCodeInt, err := strconv.ParseInt(req.FormValue("bus_code"), 10, 64)
	if err != nil {
		return nil, errutil.ErrInvalidArguments.Msg("bus_code")
	}
	busCode := model.BusCode(busCodeInt)

	row := &model.Order{
		Type:     model.OrderTypeH5,
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

	// 创建四方订单
	output := &CreateOrderOutput{
		OrderID: orderID,
		Success: true,
	}
	err = s.ThirdDriver.CreateOrderForH5(orderID, orderAmt, sourceID, int32(busCode))
	if err != nil {
		s.Logger.Errorw("create order for h5 failed",
			"err", err,
		)
		output.Success = false
		output.Message = err.Error()
	}

	return response.Data(output), nil
}

func (s *App) orderForQRCode(req *http.Request) (resp response.Response, err error) {
	orderID := req.FormValue("order_id")
	orderAmt, _ := strconv.ParseFloat(req.FormValue("order_amt"), 10)
	sourceID := req.FormValue("user_id")
	busCodeInt, err := strconv.ParseInt(req.FormValue("bus_code"), 10, 64)
	if err != nil {
		return nil, errutil.ErrInvalidArguments.Msg("bus_code")
	}
	busCode := model.BusCode(busCodeInt)

	row := &model.Order{
		Type:     model.OrderTypeQRCode,
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

	// 创建四方订单
	output := &CreateOrderOutput{
		OrderID: orderID,
		Success: true,
	}
	err = s.ThirdDriver.CreateOrderForQRCode(orderID, orderAmt, sourceID, int32(busCode))
	if err != nil {
		s.Logger.Errorw("create order for qrcode failed",
			"err", err,
		)
		output.Success = false
		output.Message = err.Error()
	}

	return response.Data(output), nil
}

type CancelOrderOutput struct {
	OrderID string `json:"order_id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func (s *App) cancelOrder(req *http.Request) (resp response.Response, err error) {
	orderID := req.FormValue("order_id")
	sourceID := req.FormValue("user_id")

	// 删除内部订单
	if err := s.OrderRepo.Delete(req.Context(), orderID, sourceID); err != nil {
		return nil, err
	}

	// 删除四方订单
	output := &CancelOrderOutput{
		OrderID: orderID,
		Success: true,
	}
	_, err = s.ThirdDriver.CancelOrder(orderID, sourceID)
	if err != nil {
		s.Logger.Errorw("cancel order failed",
			"err", err,
		)
		output.Success = false
		output.Message = err.Error()
	}

	return response.Data(output), nil
}
