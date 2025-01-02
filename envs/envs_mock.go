package envs

type EnvsGetterMock struct {
	values map[string]string
}

func NewEnvsGetterMock(values map[string]string) IEnvGetter {
	getter := &EnvsGetterMock{
		values: values,
	}
	return getter
}

func (e *EnvsGetterMock) Get(name string) (value string, exist bool) {
	value, exist = e.values[name]
	return
}
