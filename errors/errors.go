package errors

import "errors"

var (
	ErrorConfig       = errors.New("error config")
	ErrorRequired     = errors.New("error required")
	ErrorParseBoolean = errors.New("error parsing of boolean value")
	ErrorEmptyParams  = errors.New("empty params")
)
