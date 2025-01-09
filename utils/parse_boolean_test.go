package utils

import (
	"testing"

	"github.com/nextmillenniummedia/config-go/errors"
	"github.com/stretchr/testify/assert"
)

type TestItem struct {
	value  string
	result bool
}

func TestParseBoolean(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"", false},
		{"0", false},
		{"1", true},
		{"false", false},
		{"true", true},
		{"f", false},
		{"t", true},
		{"FaLse", false},
		{"TruE", true},
	}

	for _, test := range tests {
		result, err := ParseBoolean(test.value, false)
		assert.Equal(test.result, result, test)
		assert.Equal(nil, err, test)
	}
}

func TestParseBooleanEmptyValueAsTrue(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	result, err := ParseBoolean("", true)
	assert.Equal(true, result)
	assert.Nil(err)
}

func TestOnError(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	_, err := ParseBoolean("any text", false)
	assert.Equal(errors.ErrorParseBoolean, err)
}
