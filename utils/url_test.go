package utils_test

import (
	"testing"

	"github.com/nextmillenniummedia/config-go/utils"
	"github.com/stretchr/testify/assert"
)

func TestUrlValidate(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type TestCase struct {
		value   string
		message string
	}
	testsFail := []TestCase{
		{value: "", message: "empty"},
		{value: "asd", message: "random"},
	}
	for _, test := range testsFail {
		err := utils.UrlValidate(test.value)
		assert.NotNil(err, test.message)
	}

	testsOk := []TestCase{
		{value: "http://localhost", message: "localhost"},
		{value: "http://127.0.0.1", message: "ip"},
		{value: "http://127.0.0.1:3000", message: "ip:port"},
		{value: "https://domain.com", message: "https"},
		{value: "mongodb://user:password@192.168.1.39/name?authSource=admin", message: "mongo"},
	}
	for _, test := range testsOk {
		err := utils.UrlValidate(test.value)
		assert.Nil(err, test.message)
	}
}

func TestUrlClearLastSlash(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	result := utils.UrlClearLastSlash("http://domain.com/")
	assert.Equal("http://domain.com", result)
}

func TestUrlClearLastSlashDoubleSlash(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	result := utils.UrlClearLastSlash("http://domain.com//")
	assert.Equal("http://domain.com", result)
}
