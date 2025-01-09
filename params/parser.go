package params

import (
	"strings"

	"github.com/nextmillenniummedia/config-go/errors"
	"github.com/nextmillenniummedia/config-go/utils"
)

func ParseParams(tag string) (*Params, error) {
	paramsMap, err := getParamsMap(tag)
	if err != nil {
		return nil, err
	}
	required, err := utils.ParseBoolean(paramsMap["required"], true)
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
	return &Params{
		Field:    field,
		Splitter: splitter,
		Required: required,
		Format:   format,
		Doc:      doc,
	}, nil
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
		result[name] = utils.TrimEscape(value)
	}
	return result, nil
}
