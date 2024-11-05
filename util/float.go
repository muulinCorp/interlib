package util

import (
	"encoding/json"
	"fmt"
	"strconv"
)

func GetFloat64(v interface{}) (*float64, error) {
	var f float64
	var err error
	switch i := v.(type) {
	default:
		return nil, fmt.Errorf("unexpected type %T", v)
	case nil:
		return nil, nil
	case string:
		f, err = strconv.ParseFloat(string(i), 64)
		if err != nil {
			return nil, err
		}
	case uint32:
		f = float64(uint32(i))
	case uint16:
		f = float64(uint16(i))
	case int16:
		f = float64(int16(i))
	case int:
		f = float64(int(i))
	case json.Number:
		jn, _ := v.(json.Number)
		f, err = jn.Float64()
		if err != nil {
			return nil, err
		}

	case float64:
		f = v.(float64)
	}
	return &f, nil
}
