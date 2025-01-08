package envs

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/be-true/config-go/errors"
	"github.com/be-true/config-go/params"
	"github.com/be-true/config-go/utils"
)

type ConfigItem struct {
	field     *reflect.Value
	fieldName string
	fieldType *reflect.StructField
	params    *params.Params
	prefix    string
	envGetter IEnvGetter
	errs      []error
}

func NewConfigItem(field *reflect.Value, fieldName string, fieldType *reflect.StructField, params *params.Params, prefix string) *ConfigItem {
	return &ConfigItem{
		field:     field,
		fieldName: fieldName,
		fieldType: fieldType,
		prefix:    prefix,
		params:    params,
		errs:      make([]error, 0),
	}
}

func (ci *ConfigItem) SetEnvGetter(getter IEnvGetter) *ConfigItem {
	ci.envGetter = getter
	return ci
}

func (ci *ConfigItem) HasError() bool {
	return len(ci.errs) > 0
}

func (ci *ConfigItem) Process() {
	envName := ci.getEnvName()
	env, has := ci.envGetter.Get(envName)
	
	if !has && ci.params.Required {
		ci.appendError(errors.ErrorConfig)
		return
	}

	switch ci.field.Kind() {
	case reflect.String:
		ci.setString(env)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		ci.setInt(env)
	case reflect.Float32, reflect.Float64:
		ci.setFloat(env)
	case reflect.Bool:
		ci.setBool(env)
	case reflect.Slice:
		elemKind := ci.field.Type().Elem().Kind()
		envs := strings.Split(env, ci.params.Splitter)
		slice := reflect.MakeSlice(ci.field.Type(), len(envs), len(envs))
		switch elemKind {
		case reflect.String:
			ci.setSliceString(&slice, envs)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			ci.setSliceInt(&slice, envs)
		case reflect.Float32, reflect.Float64:
			ci.setSliceFloat(&slice, envs)
		}
		ci.field.Set(slice)
	}
}

func (ci *ConfigItem) setString(env string) {
	ci.field.SetString(env)
}

func (ci *ConfigItem) setInt(env string) {
	envInt, err := strconv.Atoi(env)
	if err != nil {
		ci.appendError(err)
		return
	}
	ci.field.SetInt(int64(envInt))
}

func (ci *ConfigItem) setFloat(env string) {
	envFloat, err := strconv.ParseFloat(env, 64)
	if err != nil {
		ci.appendError(err)
		return
	}
	ci.field.SetFloat(envFloat)
}

func (ci *ConfigItem) setBool(env string) {
	envBool, err := utils.ParseBoolean(env, false)
	if err != nil {
		ci.appendError(err)
		return
	}
	ci.field.SetBool(envBool)
}

func (ci *ConfigItem) setSliceString(slice *reflect.Value, envs []string) {
	for j, env := range envs {
		slice.Index(j).SetString(env)
	}
}

func (ci *ConfigItem) setSliceInt(slice *reflect.Value, envs []string) {
	for j, env := range envs {
		envInt, err := strconv.Atoi(env)
		if err != nil {
			ci.appendError(err)
			return
		}
		slice.Index(j).SetInt(int64(envInt))
	}
}

func (ci *ConfigItem) setSliceFloat(slice *reflect.Value, envs []string) {
	for j, env := range envs {
		envFloat, err := strconv.ParseFloat(env, 64)
		if err != nil {
			ci.appendError(err)
			return
		}
		slice.Index(j).SetFloat(envFloat)
	}
}

func (ci *ConfigItem) getEnvName() string {
	fieldName := ci.fieldName
	if len(ci.params.Field) > 0 {
		fieldName = ci.params.Field
	}
	fieldName = strings.ToUpper(fieldName)
	return fmt.Sprintf("%s_%s", ci.prefix, fieldName)
}

func (ci *ConfigItem) appendError(err error) {
	ci.errs = append(ci.errs, err)
}
