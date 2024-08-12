package params

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestItem struct {
	value  string
	params Params
	err    error
}

func TestParseRequired(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"required", Params{Required: true}, nil},
		// {"required=1", Params{Required: true}, nil},
		// {"required=0", Params{Required: false}, nil},
		// {"required,doc='Any doc text'", Params{Required: true}, nil},
		// {"required=a", Params{Required: false}, errors.ErrorParseBoolean},
	}

	for _, test := range tests {
		params, err := parseParams(test.value)
		assert.Equal(test.params.Required, params.Required, test)
		assert.Equal(test.err, err, test)
	}
}
