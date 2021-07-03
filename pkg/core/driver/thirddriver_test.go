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
	})
	sign, err := td.generateSign("20180912154311shop201809131545", 100.00, "shop888", 3201)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(sign)
}