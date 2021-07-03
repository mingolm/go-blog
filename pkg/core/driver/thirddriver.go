package driver

func NewThirdDriver() Driver {
	return &thirdDriver{}
}

type thirdDriver struct {
	pageUrl string // 支付后页面回跳地址
	bgUrl   string // 支付结果后台通知地址
}

func (t *thirdDriver) CreateOrderForH5(sourceID, orderId string, orderAmt float64, busCode int32) (err error) {
	return
}

func (t *thirdDriver) CreateOrderForQRCode(sourceID, orderId string, orderAmt float64, busCode int32) (err error) {
	return
}

func (t *thirdDriver) CancelOrder(sourceID, orderID string) (output *OrderCancelOutput, err error) {
	return
}

func (t *thirdDriver) GetOrderStatus(sourceID, orderID string) (output *OrderStatusOutput, err error) {
	return
}
