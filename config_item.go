package configgo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/nextmillenniummedia/config-go/errors"
	"github.com/nextmillenniummedia/config-go/params"
	"github.com/nextmillenniummedia/config-go/utils"
)

type configItem struct {
	field     *reflect.Value
	fieldName string
	fieldType *reflect.StructField
	params    *params.Params
	prefix    string
	env       IEnv
	errs      []error
}

func newConfigItem(
	field *reflect.Value,
	fieldName string,
	fieldType *reflect.StructField,
	params *params.Params,
	prefix string,
) *configItem {
	return &configItem{
		field:     field,
		fieldName: fieldName,
		fieldType: fieldType,
		prefix:    prefix,
		params:    params,
		errs:      make([]error, 0),
	}
}

func (ci *configItem) SetEnv(env IEnv) *configItem {
	ci.env = env
	return ci
}

func (ci *configItem) HasError() bool {
	return len(ci.errs) > 0
}

func (ci *configItem) Process() {
	ci.clear()

	envName := ci.getEnvName()
	env, has := ci.env.Get(envName)

	if !has && ci.params.Required {
		ci.appendError(errors.ErrorRequired)
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

func (ci configItem) GetErrorsMessage() string {
	result := ci.getEnvName() + ": "
	errors := make([]string, 0)
	for _, err := range ci.errs {
		errors = append(errors, err.Error())
	}
	result = result + strings.Join(errors, ", ")
	return result
}

func (ci *configItem) clear() {
	ci.errs = []error{}
}

func (ci *configItem) setString(env string) {
	ci.field.SetString(env)
}

func (ci *configItem) setInt(env string) {
	envInt, err := strconv.Atoi(env)
	if err != nil {
		ci.appendError(err)
		return
	}
	ci.field.SetInt(int64(envInt))
}

func (ci *configItem) setFloat(env string) {
	envFloat, err := strconv.ParseFloat(env, 64)
	if err != nil {
		ci.appendError(err)
		return
	}
	ci.field.SetFloat(envFloat)
}

func (ci *configItem) setBool(env string) {
	envBool, err := utils.ParseBoolean(env, false)
	if err != nil {
		ci.appendError(err)
		return
	}
	ci.field.SetBool(envBool)
}

func (ci *configItem) setSliceString(slice *reflect.Value, envs []string) {
	for j, env := range envs {
		slice.Index(j).SetString(env)
	}
}

func (ci *configItem) setSliceInt(slice *reflect.Value, envs []string) {
	for j, env := range envs {
		envInt, err := strconv.Atoi(env)
		if err != nil {
			ci.appendError(err)
			return
		}
		slice.Index(j).SetInt(int64(envInt))
	}
}

func (ci *configItem) setSliceFloat(slice *reflect.Value, envs []string) {
	for j, env := range envs {
		envFloat, err := strconv.ParseFloat(env, 64)
		if err != nil {
			ci.appendError(err)
			return
		}
		slice.Index(j).SetFloat(envFloat)
	}
}

func (ci *configItem) getEnvName() string {
	fieldName := ci.fieldName
	if len(ci.params.Field) > 0 {
		fieldName = ci.params.Field
	}
	fieldName = strings.ToUpper(fieldName)
	return fmt.Sprintf("%s_%s", ci.prefix, fieldName)
}

func (ci *configItem) appendError(err error) {
	ci.errs = append(ci.errs, err)
}
