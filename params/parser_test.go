package params

import (
	"testing"

	"github.com/be-true/config-go/errors"
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
		{"required=1", Params{Required: true}, nil},
		{"required=0", Params{Required: false}, nil},
		{"required,doc='Any doc text'", Params{Required: true}, nil},
	}

	for _, test := range tests {
		params, err := parseParams(test.value)
		assert.Equal(test.params.Required, params.Required, test)
		assert.Nil(err, test)
	}
}

func TestParseRequiredErrors(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"required=a", Params{Required: false}, errors.ErrorParseBoolean},
	}

	for _, test := range tests {
		params, err := parseParams(test.value)
		assert.Nil(params, test)
		assert.Equal(test.err, err, test)
	}
}

func TestParseFormat(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"format=url", Params{Format: "url"}, nil},
		{"format=Url", Params{Format: "url"}, nil},
		{"format=Url", Params{Format: "url"}, nil},
		{"format='url'", Params{Format: "url"}, nil},
		{"format=url,doc='Any doc text'", Params{Format: "url"}, nil},
	}

	for _, test := range tests {
		params, err := parseParams(test.value)
		assert.Equal(test.params.Format, params.Format, test)
		assert.Nil(err, test)
	}
}

func TestParseDoc(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"format=url,doc=Word", Params{Doc: "Word"}, nil},
		{"format=url,doc='With space'", Params{Doc: "With space"}, nil},
	}

	for _, test := range tests {
		params, err := parseParams(test.value)
		assert.Equal(test.params.Doc, params.Doc, test)
		assert.Nil(err, test)
	}
}
