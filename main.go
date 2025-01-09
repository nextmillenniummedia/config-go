package configgo

import (
	"reflect"

	"github.com/be-true/config-go/params"
)

func InitConfig(config any, settings Setting) *Config {
	value := reflect.ValueOf(config).Elem()
	typed := value.Type()
	items := make([]*configItem, 0)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		fieldType := typed.Field(i)
		fieldName := typed.Field(i).Name
		if !field.CanSet() {
			// TODO[pt]: need check for pointer or error
			continue
		}
		tag := typed.Field(i).Tag.Get("config")
		params, err := params.ParseParams(tag)
		if err != nil {
			panic(err)
		}
		items = append(items, newConfigItem(
			&field,
			fieldName,
			&fieldType,
			params,
			settings.Prefix,
		))
	}

	return newConfig(config, settings, items)
}
