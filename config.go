package configgo

import (
	"fmt"
)

type Config struct {
	config   any
	settings Setting
	items    []*configItem
}

func newConfig(inst any, settings Setting, items []*configItem) *Config {
	config := &Config{
		config:   inst,
		settings: settings,
		items:    items,
	}
	config.SetEnv(newEnv())
	return config
}

func (c *Config) Process() (err error) {
	for _, item := range c.items {
		item.Process()
	}
	if c.hasErrors() {
		return fmt.Errorf("%s", c.GetErrorsMessage())
	}
	return nil
}

func (c *Config) SetEnv(env IEnv) *Config {
	for _, item := range c.items {
		item.SetEnv(env)
	}
	return c
}

func (c Config) GetErrorsMessage() string {
	result := ""
	if len(c.settings.Title) > 0 {
		result = result + c.settings.Title + "\n"
	}
	for _, item := range c.items {
		err := item.GetErrorsMessage()
		if err == "" {
			continue
		}
		result = result + item.GetErrorsMessage() + "\n"
	}
	return result
}

func (c Config) hasErrors() bool {
	for _, item := range c.items {
		if item.HasError() {
			return true
		}
	}
	return false
}
