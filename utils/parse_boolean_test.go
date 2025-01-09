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
		result, err := ParseBoolean(test.value)
		assert.Equal(test.result, result, test)
		assert.Equal(nil, err, test)
	}
}

func TestParseBooleanEmptyValueAsTrue(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	result, err := ParseBoolean("")
	assert.Equal(false, result)
	assert.Equal(errors.ErrorParseBoolean, err)
}

func TestOnError(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	_, err := ParseBoolean("any text")
	assert.Equal(errors.ErrorParseBoolean, err)
}
