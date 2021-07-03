package driver

import (
	"github.com/mingolm/go-recharge/configs"
	"testing"
)

func TestGenerateSign(t *testing.T) {
	td := NewThirdDriver(&ThirdConfig{
		Key:              configs.DefaultConfigs.PAYKey,
		PageUrl:          configs.DefaultConfigs.PAYPageUrl,
		BGUrl:            configs.DefaultConfigs.PAYBGUrl,
		H5RemoteAddr:     configs.DefaultConfigs.PAYH5RemoteAddr,
		QRCodeRemoteAddr: configs.DefaultConfigs.PAYQRCodeRemoteAddr,
	}).(*thirdDriver)
	sign := td.generateSign("shop888", "20180912154311shop201809131545", 100.00, 3201)
	t.Log(sign)
}
