package envs

import "os"

type EnvGetter struct{}

func NewEnvGetter() IEnvGetter {
	return &EnvGetter{}
}

func (e *EnvGetter) Get(name string) (value string, exist bool) {
	value, exist = os.LookupEnv(name)
	return
}
