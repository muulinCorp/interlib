package device

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetService(t *testing.T) {
	lib := NewLib(&http.Client{}, "http://127.0.0.1:9080")
	res, err := lib.GetChannel("00:00:00:00:00:aa", "")
	assert.Nil(t, err)
	assert.Equal(t, "", res)
}
