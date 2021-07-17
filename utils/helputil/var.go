package helputil

import (
	"strconv"
)

func Interface2Int(value interface{}) (n int) {
	switch v := value.(type) {
	case string:
		n, _ = strconv.Atoi(v)
	case int:
		n = v
	case int32:
		n = int(v)
	case int64:
		n = int(v)
	case uint32:
		n = int(v)
	case uint64:
		n = int(v)
	}

	return
}

func Interface2Int64(value interface{}) (n int64) {
	switch v := value.(type) {
	case string:
		n, _ = strconv.ParseInt(v, 10, 64)
	case int:
		n = int64(v)
	case int32:
		n = int64(v)
	case int64:
		n = v
	case uint32:
		n = int64(v)
	case uint64:
		n = int64(v)
	}

	return
}

func Interface2Uint64(value interface{}) (n uint64) {
	switch v := value.(type) {
	case string:
		n, _ = strconv.ParseUint(v, 10, 64)
	case int:
		n = uint64(v)
	case int32:
		n = uint64(v)
	case int64:
		n = uint64(v)
	case uint32:
		n = uint64(v)
	case uint64:
		n = v
	}

	return
}
