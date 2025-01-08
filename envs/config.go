package envs

import (
	"reflect"

	"github.com/be-true/config-go/errors"
	"github.com/be-true/config-go/params"
)

type ConfigEnvs struct {
	config   any
	settings SettingEnvs
	items    []*ConfigItem
}

func InitConfigEnvs(config any, settings SettingEnvs) *ConfigEnvs {
	value := reflect.ValueOf(config).Elem()
	typed := value.Type()
	items := make([]*ConfigItem, 0)
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
		items = append(items, NewConfigItem(
			&field,
			fieldName,
			&fieldType,
			params,
			settings.Prefix,
		))
	}

	return &ConfigEnvs{
		config:   config,
		settings: settings,
		items:    items,
	}
}

func (c *ConfigEnvs) Process() (err error) {
	for _, item := range c.items {
		item.Process()
	}
	if c.hasErrors() {
		return errors.ErrorConfig
	}
	return nil
}

func (c *ConfigEnvs) SetEnvGetter(getter IEnvGetter) *ConfigEnvs {
	for _, item := range c.items {
		item.SetEnvGetter(getter)
	}
	return c
}

func (c ConfigEnvs) hasErrors() bool {
	for _, item := range c.items {
		if item.HasError() {
			return true
		}
	}
	return false
}
