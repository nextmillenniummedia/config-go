package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrimEscape(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()
	result := TrimEscape("'a'")
	assert.Equal("a", result)
}
