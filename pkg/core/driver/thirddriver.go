package driver

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func NewThirdDriver(config *ThirdConfig) Driver {
	return &thirdDriver{
		config,
	}
}

type ThirdConfig struct {
	Key              string // 商户密钥
	PageUrl          string // 支付后页面回跳地址
	BGUrl            string // 支付结果后台通知地址
	H5RemoteAddr     string // h5支付接口地址
	QRCodeRemoteAddr string // 二维码支付接口地址
}

type thirdDriver struct {
	*ThirdConfig
}

func (t *thirdDriver) CreateOrderForH5(sourceID, orderID string, orderAmt float64, busCode int32) (err error) {
	sign := t.generateSign(sourceID, orderID, orderAmt, busCode)
	_, err = dhc().Post(t.H5RemoteAddr, map[string]interface{}{
		"ORDER_AMT": orderAmt,
		"ORDER_ID":  orderID,
		"USER_ID":   sourceID,
		"BUS_CODE":  busCode,
		"PAGE_URL":  t.PageUrl,
		"BG_URL":    t.BGUrl,
		"SIGN":      sign,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *thirdDriver) CreateOrderForQRCode(sourceID, orderID string, orderAmt float64, busCode int32) (err error) {
	sign := t.generateSign(sourceID, orderID, orderAmt, busCode)
	_, err = dhc().Post(t.H5RemoteAddr, map[string]interface{}{
		"ORDER_AMT": orderAmt,
		"ORDER_ID":  orderID,
		"USER_ID":   sourceID,
		"BUS_CODE":  busCode,
		"PAGE_URL":  t.PageUrl,
		"BG_URL":    t.BGUrl,
		"SIGN":      sign,
	})
	if err != nil {
		return err
	}
	return nil
}

func (t *thirdDriver) CancelOrder(sourceID, orderID string) (output *OrderCancelOutput, err error) {
	return
}

func (t *thirdDriver) GetOrderStatus(sourceID, orderID string) (output *OrderStatusOutput, err error) {
	return
}

func (t *thirdDriver) generateSign(sourceID, orderId string, orderAmt float64, busCode int32) (sign string) {
	// sign1
	m := md5.New()
	m.Write([]byte(orderId))
	m.Write([]byte(strconv.FormatFloat(orderAmt, 'f', 2, 64)))
	m.Write([]byte(sourceID))
	m.Write([]byte(strconv.FormatInt(int64(busCode), 10)))
	sign1 := m.Sum(nil)

	// sign2
	m2 := md5.New()
	m2.Write([]byte(hex.EncodeToString(sign1)))
	m2.Write([]byte(t.Key))
	sign2 := m2.Sum(nil)

	sign = hex.EncodeToString(sign2)
	return sign[8:24]
}
