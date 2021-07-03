package driver

type Driver interface {
	// 创建 h5 订单
	CreateOrderForH5(sourceID, orderID string, orderAmt float64, busCode int32) (err error)
	// 创建扫码订单
	CreateOrderForQRCode(sourceID, orderID string, orderAmt float64, busCode int32) (err error)
	// 取消订单
	CancelOrder(sourceID, orderID string) (output *OrderCancelOutput, err error)
	// 查询订单状态
	GetOrderStatus(sourceID, orderID string) (output *OrderStatusOutput, err error)
}

type OrderCancelOutput struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Desc    string      `json:"desc"`
	Result  interface{} `json:"result"`
}

type OrderStatusOutput struct {
	State   int32  `json:"state"`
	Message string `json:"msg"`
}
