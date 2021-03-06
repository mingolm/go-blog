package driver

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/mingolm/go-recharge/pkg/httpclient"
	"github.com/mingolm/go-recharge/utils/errutil"
	"reflect"
	"strconv"
)

/**
 * 四方交互 driver
 */

func NewThirdDriver(config *ThirdConfig) *ThirdDriver {
	return &ThirdDriver{
		config,
	}
}

type ThirdConfig struct {
	Key              string // 商户密钥
	PageUrl          string // 支付后页面回跳地址
	BGUrl            string // 支付结果后台通知地址
	H5RemoteAddr     string // h5支付接口地址
	QRCodeRemoteAddr string // 二维码支付接口地址
	CancelRemoteAddr string // 取消订单接口地址
	StatusRemoteAddr string // 查询订单接口地址
}

type ThirdDriver struct {
	*ThirdConfig
}

// 创建 h5 订单
func (t *ThirdDriver) CreateOrderForH5(orderID string, orderAmt float64, sourceID string, busCode int32) (err error) {
	sign, err := t.generateSign(orderID, orderAmt, sourceID, busCode)
	if err != nil {
		return errutil.ErrInternal.Msg(err.Error())
	}
	_, err = httpclient.DHCP().PostForm(t.H5RemoteAddr, map[string]string{
		"ORDER_ID":  orderID,
		"ORDER_AMT": strconv.FormatFloat(orderAmt, 'f', 2, 10),
		"USER_ID":   sourceID,
		"BUS_CODE":  strconv.Itoa(int(busCode)),
		"PAGE_URL":  t.PageUrl,
		"BG_URL":    t.BGUrl,
		"SIGN":      sign,
	})
	if err != nil {
		return errutil.ErrInternal.Msg(err.Error())
	}
	return nil
}

// 创建扫码订单
func (t *ThirdDriver) CreateOrderForQRCode(orderID string, orderAmt float64, sourceID string, busCode int32) (err error) {
	sign, err := t.generateSign(orderID, orderAmt, sourceID, busCode)
	if err != nil {
		return errutil.ErrInternal.Msg(err.Error())
	}
	_, err = httpclient.DHCP().PostForm(t.H5RemoteAddr, map[string]string{
		"ORDER_ID":  orderID,
		"ORDER_AMT": strconv.FormatFloat(orderAmt, 'f', 2, 10),
		"USER_ID":   sourceID,
		"BUS_CODE":  strconv.Itoa(int(busCode)),
		"PAGE_URL":  t.PageUrl,
		"BG_URL":    t.BGUrl,
		"SIGN":      sign,
	})
	if err != nil {
		return errutil.ErrInternal.Msg(err.Error())
	}
	return nil
}

// 取消订单
func (t *ThirdDriver) CancelOrder(orderID, sourceID string) (output *OrderCancelOutput, err error) {
	sign, err := t.generateSign(orderID, sourceID)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	bs, err := httpclient.DHCP().PostForm(t.CancelRemoteAddr, map[string]string{
		"ORDER_ID": orderID,
		"USER_ID":  sourceID,
		"SIGN":     sign,
	})
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	output = &OrderCancelOutput{}
	if err := json.Unmarshal(bs, output); err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	if !output.Success || output.Code != 200 {
		return nil, errutil.ErrInternal.Msg(output.Desc)
	}
	return output, nil
}

// 查询订单状态
func (t *ThirdDriver) GetOrderStatus(orderID, sourceID string) (output *OrderStatusOutput, err error) {
	sign, err := t.generateSign(orderID, sourceID)
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	bs, err := httpclient.DHCP().PostForm(t.StatusRemoteAddr, map[string]string{
		"ORDER_ID": orderID,
		"USER_ID":  sourceID,
		"SIGN":     sign,
	})
	if err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	output = &OrderStatusOutput{}
	if err := json.Unmarshal(bs, output); err != nil {
		return nil, errutil.ErrInternal.Msg(err.Error())
	}
	if !output.Success || output.Code != 200 {
		return nil, errutil.ErrInternal.Msg(output.Desc)
	}

	return output, nil
}

func (t *ThirdDriver) generateSign(values ...interface{}) (sign string, err error) {
	// sign1
	m := md5.New()
	for _, value := range values {
		switch v := value.(type) {
		case string:
			m.Write([]byte(v))
		case []byte:
			m.Write(v)
		case float64:
			m.Write([]byte(strconv.FormatFloat(v, 'f', 2, 64)))
		case float32:
			m.Write([]byte(strconv.FormatFloat(float64(v), 'f', 2, 64)))
		case int64:
			m.Write([]byte(strconv.FormatInt(v, 10)))
		case int:
			m.Write([]byte(strconv.Itoa(v)))
		default:
			return "", fmt.Errorf("unknown sign field value type %s", reflect.TypeOf(value).String())
		}
	}
	sign1 := m.Sum(nil)

	// sign2
	m2 := md5.New()
	m2.Write([]byte(hex.EncodeToString(sign1)))
	m2.Write([]byte(t.Key))
	sign2 := m2.Sum(nil)

	sign = hex.EncodeToString(sign2)
	return sign[8:24], nil
}
