package utils

import (
	"strings"

	"github.com/nextmillenniummedia/config-go/errors"
)

func ParseBoolean(value string, emptyValue bool) (bool, error) {
	value = strings.ToLower(value)
	switch value {
	case "":
		return emptyValue, nil
	case "false", "f", "0":
		return false, nil
	case "true", "t", "1":
		return true, nil
	}
	return false, errors.ErrorParseBoolean
}
