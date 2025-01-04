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
		field := value.Field(i)
		if !field.CanSet() {
			continue
		}
		tag := typed.Field(i).Tag.Get("config")
		params, err := params.ParseParams(tag)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		envName := GetFieldName(typed.Field(i).Name, c.settings.Prefix, params)
		env, _ := c.envGetter.Get(envName)

		switch field.Kind() {
		case reflect.String:
			setString(&field, env, &errs)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			setInt(&field, env, &errs)
		case reflect.Float32, reflect.Float64:
			setFloat(&field, env, &errs)
		case reflect.Bool:
			setBool(&field, env, &errs)
		case reflect.Slice:
			elemKind := field.Type().Elem().Kind()
			envs := strings.Split(env, ",")
			slice := reflect.MakeSlice(field.Type(), len(envs), len(envs))
			switch elemKind {
			case reflect.String:
				setSliceString(&slice, envs, &errs)
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				setSliceInt(&slice, envs, &errs)
			case reflect.Float32, reflect.Float64:
				setSliceFloat(&slice, envs, &errs)
			}
			field.Set(slice)
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

func setString(field *reflect.Value, env string, errs *[]error) {
	field.SetString(env)
}

func setInt(field *reflect.Value, env string, errs *[]error) {
	envInt, err := strconv.Atoi(env)
	if err != nil {
		*errs = append(*errs, err)
		return
	}
	field.SetInt(int64(envInt))
}

func setFloat(field *reflect.Value, env string, errs *[]error) {
	envFloat, err := strconv.ParseFloat(env, 64)
	if err != nil {
		*errs = append(*errs, err)
		return
	}
	field.SetFloat(envFloat)
}

func setBool(field *reflect.Value, env string, errs *[]error) {
	envBool, err := utils.ParseBoolean(env, false)
	if err != nil {
		*errs = append(*errs, err)
		return
	}
	field.SetBool(envBool)
}

func setSliceString(slice *reflect.Value, envs []string, errs *[]error) {
	for j, env := range envs {
		slice.Index(j).SetString(env)
	}
}

func setSliceInt(slice *reflect.Value, envs []string, errs *[]error) {
	for j, env := range envs {
		envInt, err := strconv.Atoi(env)
		if err != nil {
			*errs = append(*errs, err)
			return
		}
		slice.Index(j).SetInt(int64(envInt))
	}
}

func setSliceFloat(slice *reflect.Value, envs []string, errs *[]error) {
	for j, env := range envs {
		envFloat, err := strconv.ParseFloat(env, 64)
		if err != nil {
			*errs = append(*errs, err)
			return
		}
		slice.Index(j).SetFloat(envFloat)
	}
}
