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
		{"required", Params{Required: true}, nil},
		{"required=1", Params{Required: true}, nil},
		{"required=0", Params{Required: false}, nil},
		{"required,doc='Any doc text'", Params{Required: true}, nil},
	}

	for _, test := range tests {
		params, err := ParseParams(test.value)
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

func TestParseDefault(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"default=1", Params{Default: "1"}, nil},
	}

	for _, test := range tests {
		params, err := ParseParams(test.value)
		assert.Equal(test.params.Default, params.Default, test)
		assert.Nil(err, test)
	}
}

func TestParseAllowedParams(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"iw=1,doc='Test'", Params{}, errors.ErrorParseNotAllowed},
	}

	for _, test := range tests {
		_, err := ParseParams(test.value)
		assert.NotNil(err, test)
	}
}

func TestParseEnum(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	tests := []TestItem{
		{"enum=a", Params{Enum: []string{"a"}}, nil},
		{"enum=a|b", Params{Enum: []string{"a", "b"}}, nil},
	}

	testsErrors := []TestItem{
		{"enum=", Params{Enum: []string{}}, errors.ErrorEnumValues},
	}

	for _, test := range tests {
		params, _ := ParseParams(test.value)
		assert.Equal(test.params.Enum, params.Enum, test)
	}

	for _, test := range testsErrors {
		_, err := ParseParams(test.value)
		assert.Equal(err, test.err)
	}
}
