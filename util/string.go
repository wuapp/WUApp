package util

import (
	"fmt"
	"strconv"
)

func ToString(v interface{},quotes string) string {
	switch v := v.(type) {
	case string:
		return quotes + v + quotes
	case bool:
		if v {
			return "true"
		} else {
			return "false"
		}
	case int:
		return strconv.Itoa(v)
	case int8:
		return strconv.Itoa(int(v))
	case int16:
		return strconv.Itoa(int(v))
	case int32:
		return strconv.Itoa(int(v))
	case int64:
		return strconv.FormatInt(v, 10)
	case float32:
		return strconv.FormatFloat(float64(v), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(v), 'f', -1, 64)
	case []byte:
		return quotes+string(v)+quotes
	case complex64:
		//todo
	case complex128:
		//todo
	case fmt.Stringer:
		return v.String()
	case struct{}:
		return "[sruct]"
	case interface{}:
		return "[interface]"

	}
	return ""
}
