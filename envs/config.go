package envs

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/be-true/config-go/params"
	"github.com/be-true/config-go/utils"
)

type ConfigEnvs struct {
	envGetter IEnvGetter
	settings  SettingEnvs
}

func InitConfigEnvs(settings SettingEnvs) *ConfigEnvs {
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
	errs := make([]error, 0)
	value := reflect.ValueOf(conf).Elem()
	typed := value.Type()
	for i := 0; i < value.NumField(); i++ {
		tag := typed.Field(i).Tag.Get("config")
		params, err := params.ParseParams(tag)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		envName := GetFieldName(typed.Field(i).Name, c.settings.Prefix, params)
		env, _ := c.envGetter.Get(envName)
		field := value.Field(i)
		if !field.CanSet() {
			continue
		}
		switch field.Kind() {
		case reflect.String:
			field.SetString(env)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			envInt, err := strconv.Atoi(env)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			field.SetInt(int64(envInt))
		case reflect.Bool:
			envBool, err := utils.ParseBoolean(env, false)
			if err != nil {
				errs = append(errs, err)
				continue
			}
			field.SetBool(envBool)
		}
	}
	return errors.Join(errs...)
}

func GetFieldName(structName string, prefix string, params *params.Params) string {
	fieldName := structName
	if len(params.Field) > 0 {
		fieldName = params.Field
	}
	fieldName = strings.ToUpper(fieldName)
	return fmt.Sprintf("%s_%s", prefix, fieldName)
}
