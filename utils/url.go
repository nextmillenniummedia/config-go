package utils

import (
	"net/url"
	"strings"
)

func UrlValidate(uri string) error {
	_, err := url.ParseRequestURI(uri)
	return err
}

func UrlClearLastSlash(uri string) string {
	return strings.TrimRight(uri, "/")
}
