package util

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
)

func StrAppend(strs ...string) string {
	var buffer bytes.Buffer
	for _, str := range strs {
		buffer.WriteString(str)
	}
	return buffer.String()
}

func MD5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return hex.EncodeToString(w.Sum(nil))
}

func IsStrInList(input string, target ...string) bool {
	for _, paramName := range target {
		if input == paramName {
			return true
		}
	}
	return false
}

func RemoveDuplicates(fields []string) []string {
	// Create a map to store unique strings
	var result []string
	uniqueFields := make(map[string]bool)
	for _, field := range fields {
		if _, ok := uniqueFields[field]; ok {
			continue
		}
		uniqueFields[field] = true
		result = append(result, field)
	}
	return result
}
