package helputil

import "testing"

func TestEncryptPassword(t *testing.T) {
	t.Logf("password: %s, password_len:%d", EncryptPassword("1"), len(EncryptPassword("1")))
	t.Logf("password: %s, password_len:%d", EncryptPassword("ynBvLO.kEIcxdySvxTZ.8hCUWDT1L"), len(EncryptPassword("1")))
	t.Logf("password: %s, password_len:%d", EncryptPassword("ynBvLO.kEI$=?cxdySvxTZ.8hCUWDT1L"), len(EncryptPassword("1")))
}
