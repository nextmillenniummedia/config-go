package errors

import "errors"

var (
	ErrorConfig       = errors.New("config")
	ErrorRequired     = errors.New("required")
	ErrorParseBoolean = errors.New("parsing of boolean value")
	ErrorEmptyParams  = errors.New("empty params")
)
