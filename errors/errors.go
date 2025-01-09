package errors

import "errors"

var (
	ErrorRequired     = errors.New("required")
	ErrorParseBoolean = errors.New("parsing of boolean value")
	ErrorEmptyParams  = errors.New("empty params")
)
