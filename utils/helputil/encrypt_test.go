package helputil

import "testing"

func TestEncryptPassword(t *testing.T) {
	if EncryptPassword("ynBvLO.kEI$=?cxdySvxTZ.8hCUWDT1L") != "911af73850d50813ba5f8182a4eccfc15fccce89" {
		t.Errorf("password check failed")
	}
}
