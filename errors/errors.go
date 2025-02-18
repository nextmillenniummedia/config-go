package errors

import "errors"

var (
	ErrorRequired             = errors.New("required")
	ErrorParseBoolean         = errors.New("parsing of boolean value")
	ErrorParseNotAllowed      = errors.New("not allowed param")
	ErrorEmptyParams          = errors.New("empty params")
	ErrorUintShouldBePositive = errors.New("uint should be positive")
	ErrorEnumValues           = errors.New("enum has empty configuration")
	ErrorEnumNotValidValue    = errors.New("enum has not valid value")
)
