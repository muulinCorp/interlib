package util

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
)

func EncodeMap(dataMap *map[string]interface{}) (string, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(dataMap)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b.Bytes()), nil
}

func DecodeMap(ecodeMapStr string) (*map[string]interface{}, error) {
	mapByte, err := base64.StdEncoding.DecodeString(ecodeMapStr)
	if err != nil {
		return nil, err
	}
	b := bytes.Buffer{}
	b.Write(mapByte)
	d := gob.NewDecoder(&b)
	result := make(map[string]interface{})
	err = d.Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
