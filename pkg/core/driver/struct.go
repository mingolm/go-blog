package driver

// 四方取消订单
type OrderCancelOutput struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Desc    string      `json:"desc"`
	Result  interface{} `json:"result"`
}

// 四方查询订单状态
type OrderStatusOutput struct {
	Success bool              `json:"success"`
	Code    int               `json:"code"`
	Desc    string            `json:"desc"`
	Result  OrderStatusResult `json:"result"`
}

// 四方查询订单状态，详细 result
type OrderStatusResult struct {
	OrderID    string     `json:"ORDER_ID"`
	OrderAMT   float64    `json:"ORDER_AMT"`
	BusCode    int32      `json:"BUS_CODE"`
	PageUrl    string     `json:"PAGE_URL"`
	BGUrl      string     `json:"BG_URL"`
	State      OrderState `json:"STATE"`
	Sign       string     `json:"SIGN"`
	CreateTime int64      `json:"CREATE_TIME"`
}

type OrderState int

const (
	OrderStateWaiting       OrderState = 1
	OrderStateWaitingQRCode OrderState = 2
	OrderStateWaitingNotice OrderState = 3
	OrderStateSuccess       OrderState = 4
)
