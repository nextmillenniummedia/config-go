package main

import (
	"fmt"

	"github.com/be-true/config-go/envs"
)

type Config struct {
	Domain string `config:"format=url,require"`
	Port   int    `config:"require"`
	Debug  bool   `config:"format=boolean,default=true"`
}

func main() {
	config := Config{
		Domain: "ads",
		Debug:  false,
	}
	settings := envs.ConfigSettingEnvs{
		Title:  "Http server",
		Prefix: "HTTP",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"HTTP_DOMAIN": "domain.com",
		"HTTP_PORT":   "1024",
		"HTTP_DEBUG":  "false",
	})
	processor := envs.InitConfigEnvs(settings).SetEnvGetter(envGetter)
	err := processor.Process(&config)

	fmt.Printf("config: %#v, err: %#v", config, err)
}
