package rawdata

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateRawdata(t *testing.T) {
	lib := NewLib(&http.Client{}, "http://127.0.0.1:9080")
	err := lib.CreateRawData("serviceid", map[string]interface{}{
		"aaa": "bb",
	})
	assert.Nil(t, err)
}
