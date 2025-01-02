package envs

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/be-true/config-go/params"
)

type ConfigEnvs struct {
	envGetter IEnvGetter
	settings  ConfigSettingEnvs
}

func InitConfigEnvs(settings ConfigSettingEnvs) *ConfigEnvs {
	return &ConfigEnvs{
		envGetter: NewEnvGetter(),
		settings:  settings,
	}
}

func (c *ConfigEnvs) SetEnvGetter(getter IEnvGetter) *ConfigEnvs {
	c.envGetter = getter
	return c
}

func (c *ConfigEnvs) Process(conf any) (err error) {
	value := reflect.ValueOf(conf).Elem()
	typed := value.Type()
	for i := 0; i < value.NumField(); i++ {
		tag := typed.Field(i).Tag.Get("config")
		params, err := params.ParseParams(tag)
		if err != nil {
			continue
		}
		fieldName := GetFieldName(typed.Field(i).Name, c.settings.Prefix, params)
		env, _ := c.envGetter.Get(fieldName)
		field := value.Field(i)
		if field.CanSet() && field.Kind() == reflect.String {
			field.SetString(env)
		}
	}
	return nil
}

func GetFieldName(structName string, prefix string, params *params.Params) string {
	fieldName := structName
	if len(params.Field) > 0 {
		fieldName = params.Field
	}
	fieldName = strings.ToUpper(fieldName)
	return fmt.Sprintf("%s_%s", prefix, fieldName)
}
