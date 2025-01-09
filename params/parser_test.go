package params

import (
	"testing"

	"github.com/nextmillenniummedia/config-go/errors"
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
		{"require", Params{Require: true}, nil},
		{"require=1", Params{Require: true}, nil},
		{"require=0", Params{Require: false}, nil},
		{"require,doc='Any doc text'", Params{Require: true}, nil},
	}

	for _, test := range tests {
		params, err := ParseParams(test.value)
		assert.Equal(test.params.Require, params.Require, test)
		assert.Nil(err, test)
	}
}

func TestParseRequiredErrors(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"require=a", Params{Require: false}, errors.ErrorParseBoolean},
	}

	for _, test := range tests {
		params, err := ParseParams(test.value)
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
		params, err := ParseParams(test.value)
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
		params, err := ParseParams(test.value)
		assert.Equal(test.params.Doc, params.Doc, test)
		assert.Nil(err, test)
	}
}

func TestParseField(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"field=PORT", Params{Field: "PORT"}, nil},
		{"doc='With space'", Params{Doc: "With space"}, nil},
	}

	for _, test := range tests {
		params, err := ParseParams(test.value)
		assert.Equal(test.params.Field, params.Field, test)
		assert.Nil(err, test)
	}
}

func TestParseSplitter(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"splitter=|", Params{Splitter: "|"}, nil},
		{"", Params{Splitter: ","}, nil},
	}

	for _, test := range tests {
		params, err := ParseParams(test.value)
		assert.Equal(test.params.Splitter, params.Splitter, test)
		assert.Nil(err, test)
	}
}
