package model

import (
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"net"
	"reflect"
)

type IPv4 struct {
	net.IP
}

// Scan scan value into Jsonb, implements sql.Scanner interface
func (f *IPv4) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		ipByte := make([]byte, 4)
		binary.BigEndian.PutUint32(ipByte, uint32(v))
		f.IP = ipByte
	case uint64:
		ipByte := make([]byte, 4)
		binary.BigEndian.PutUint32(ipByte, uint32(v))
		f.IP = ipByte
	case uint32:
		ipByte := make([]byte, 4)
		binary.BigEndian.PutUint32(ipByte, v)
		f.IP = ipByte
	case []byte:
		f.IP = v
	default:
		return fmt.Errorf("unknown IP field type %s", reflect.TypeOf(value).String())
	}

	return nil
}

func (f *IPv4) FromString(s string) {
	f.IP = net.ParseIP(s).To4()
}

// Value return json value, implement driver.Valuer interface
func (f IPv4) Value() (driver.Value, error) {
	ip := f.IP.To4()
	if ip == nil {
		return int64(0), nil
	}
	return int64(binary.BigEndian.Uint32(ip)), nil
}

func GetIPv4(ip string) IPv4 {
	return IPv4{IP: net.ParseIP(ip).To4()}
}
