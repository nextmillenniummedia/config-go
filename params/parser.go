package params

import (
	"slices"
	"strings"

	"github.com/nextmillenniummedia/config-go/errors"
	"github.com/nextmillenniummedia/config-go/utils"
)

var PARAMS = []string{"require", "format", "splitter", "format", "doc", "field", "default", "enum"}

func ParseParams(tag string) (*Params, error) {
	paramsMap, err := getParamsMap(tag)
	if err != nil {
		return nil, err
	}
	require, err := utils.ParseBoolean(getRequireValue(paramsMap))
	if err != nil {
		return nil, err
	}
	splitter := strings.ToLower(paramsMap["splitter"])
	if len(splitter) == 0 {
		splitter = ","
	}
	format := strings.ToLower(paramsMap["format"])
	doc := paramsMap["doc"]
	field := paramsMap["field"]
	defaultValue := paramsMap["default"]
	return &Params{
		Field:    field,
		Splitter: splitter,
		Require:  require,
		Format:   format,
		Default:  defaultValue,
		Doc:      doc,
	}, nil
}

func getRequireValue(params map[string]string) string {
	requireValue, has := params["require"]
	if has && requireValue == "" {
		requireValue = "true"
	}
	if !has {
		requireValue = "false"
	}
	return requireValue
}

func getParamsMap(tag string) (map[string]string, error) {
	result := make(map[string]string)
	params := strings.Split(tag, ",")
	for _, param := range params {
		split := strings.Split(param, "=")
		if len(split) == 0 {
			return nil, errors.ErrorEmptyParams
		}
		value := ""
		if len(split) > 1 {
			value = split[1]
		}
		name := split[0]
		if name == "" {
			continue
		}
		result[name] = utils.TrimEscape(value)
		if !slices.Contains(PARAMS, name) {
			return nil, errors.ErrorParseNotAllowed
		}
	}
	return result, nil
}
