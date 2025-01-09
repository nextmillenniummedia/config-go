package configgo

import "os"

type env struct{}

func newEnv() IEnv {
	return &env{}
}

func (e *env) Get(name string) (value string, exist bool) {
	value, exist = os.LookupEnv(name)
	return
}
