package errors

import "errors"

var (
	ErrorParseBoolean = errors.New("error parsing boolean value")
	ErrorEmptyParams  = errors.New("empty params")
)
