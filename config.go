package main

import (
	"fmt"
	"reflect"
)

type Config struct {
	Host    string `config:"format=url,require"`
	Enabled bool   `config:"format=boolean,default=true"`
}

func main() {
	conf := Config{
		Host:    "ads",
		Enabled: false,
	}

	typed := reflect.TypeOf(conf)
	for i := 0; i < typed.NumField(); i++ {
		tag := typed.Field(i).Tag.Get("config")
		fmt.Println(tag)
	}
}
