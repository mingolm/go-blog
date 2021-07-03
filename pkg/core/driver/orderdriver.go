package driver

func NewOrderDriver(config *OrderDriverConfig) *OrderDriver {
	return &OrderDriver{
		config,
	}
}

type OrderDriverConfig struct {
	CallbackRemoteAddr string // 回调地址
}

type OrderDriver struct {
	*OrderDriverConfig
}

func (d *OrderDriver) Callback() {

}
