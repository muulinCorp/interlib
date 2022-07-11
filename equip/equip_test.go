package rawdata

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateMux(t *testing.T) {
	lib := NewLib(&http.Client{}, "http://localhost:8080")
	var err error
	for i := 0; i < 500; i++ {
		err = lib.TestMux()
	}

	assert.NotNil(t, err)
}

func Test_CreateGin(t *testing.T) {
	lib := NewLib(&http.Client{}, "http://localhost:8080")
	var err error
	for i := 0; i < 500; i++ {
		err = lib.Test()
	}

	assert.NotNil(t, err)
}
