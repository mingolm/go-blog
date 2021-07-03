package helputil

import (
	"crypto/sha1"
	"encoding/hex"
)

func EncryptPassword(password string) string {
	h := sha1.New()
	h.Write([]byte(password))
	h.Write([]byte("gbxf7h6FBlYvkcp8mWVGWJV6"))
	return hex.EncodeToString(h.Sum(nil))
}
