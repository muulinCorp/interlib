package util

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func GetFloat64(v interface{}) float64 {
	switch i := v.(type) {
	default:
		panic(fmt.Sprintf("unexpected type %T", v))
	case nil:
		return 0.0
	case string:
		f, err := strconv.ParseFloat(string(i), 64)
		if err != nil {
			return 0.0
		}
		return f
	case uint32:
		return float64(uint32(i))
	case uint16:
		return float64(uint16(i))
	case int16:
		return float64(int16(i))
	case int:
		return float64(int(i))
	case json.Number:
		jn, _ := v.(json.Number)
		result, err := jn.Float64()
		if err != nil {
			return 0.0
		}
		return result
	case float64:
		return v.(float64)
	}
}
