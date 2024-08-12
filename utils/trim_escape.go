package utils

import "strings"

func TrimEscape(text string) string {
	return strings.Trim(text, "'")
}
