package configgo_test

import (
	"testing"
	"time"

	. "github.com/nextmillenniummedia/config-go"
	"github.com/stretchr/testify/assert"
)

func TestConfigString(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Text string `config:"format=url"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "STRING",
	}
	env := newEnvsMock(map[string]string{
		"STRING_TEXT": "http://domain.com",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal("http://domain.com", config.Text)
}

func TestConfigBool(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Value bool `config:""`
	}
	config := Config{}
	settings := Setting{
		Prefix: "BOOL",
	}
	env := newEnvsMock(map[string]string{
		"BOOL_VALUE": "true",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(true, config.Value)
}

func TestConfigInt(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Int   int   `config:"field=default"`
		Int8  int8  `config:"field=8"`
		Int16 int16 `config:"field=16"`
		Int32 int32 `config:"field=32"`
		Int64 int64 `config:"field=64"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "INT",
	}
	env := newEnvsMock(map[string]string{
		"INT_DEFAULT": "1",
		"INT_8":       "8",
		"INT_16":      "16",
		"INT_32":      "32",
		"INT_64":      "64",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(1, config.Int)
	assert.Equal(int8(8), config.Int8)
	assert.Equal(int16(16), config.Int16)
	assert.Equal(int32(32), config.Int32)
	assert.Equal(int64(64), config.Int64)
}

func TestConfigUint(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Uint   uint   `config:"field=default"`
		Uint8  uint8  `config:"field=8"`
		Uint16 uint16 `config:"field=16"`
		Uint32 uint32 `config:"field=32"`
		Uint64 uint64 `config:"field=64"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "UINT",
	}
	env := newEnvsMock(map[string]string{
		"UINT_DEFAULT": "1",
		"UINT_8":       "8",
		"UINT_16":      "16",
		"UINT_32":      "32",
		"UINT_64":      "64",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(uint(1), config.Uint)
	assert.Equal(uint8(8), config.Uint8)
	assert.Equal(uint16(16), config.Uint16)
	assert.Equal(uint32(32), config.Uint32)
	assert.Equal(uint64(64), config.Uint64)
}

func TestConfigIntEmpty(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Int int `config:"field=default"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "INT",
	}
	env := newEnvsMock(map[string]string{
		"INT_DEFAULT": "",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(0, config.Int)
}

func TestConfigIntDefault(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		A int `config:"default=1"`
		B int `config:"default=2"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "INT",
	}
	env := newEnvsMock(map[string]string{
		"INT_A": "",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(1, config.A)
	assert.Equal(2, config.B)
}

func TestConfigFloat(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Float32 float32 `config:"field=32"`
		Float64 float64 `config:"field=64"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "FLOAT",
	}
	env := newEnvsMock(map[string]string{
		"FLOAT_32": "32.5",
		"FLOAT_64": "64.5",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(float32(32.5), config.Float32)
	assert.Equal(float64(64.5), config.Float64)
}

func TestConfigTimeDuration(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Duration time.Duration `config:"field=duration_ms,format=ms"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "TIME",
	}
	env := newEnvsMock(map[string]string{
		"TIME_DURATION_MS": "1000",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(time.Duration(time.Millisecond*1000), config.Duration)
}

func TestConfigSliceInt(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Int   []int   `config:"field=default"`
		Int8  []int8  `config:"field=8"`
		Int16 []int16 `config:"field=16"`
		Int32 []int32 `config:"field=32"`
		Int64 []int64 `config:"field=64"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "INT",
	}
	env := newEnvsMock(map[string]string{
		"INT_DEFAULT": "1,1",
		"INT_8":       "8,8",
		"INT_16":      "16,16",
		"INT_32":      "32,32",
		"INT_64":      "64,64",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal([]int{1, 1}, config.Int)
	assert.Equal([]int8{8, 8}, config.Int8)
	assert.Equal([]int16{16, 16}, config.Int16)
	assert.Equal([]int32{32, 32}, config.Int32)
	assert.Equal([]int64{64, 64}, config.Int64)
}

func TestConfigSliceUint(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Uint   []uint   `config:"field=default"`
		Uint8  []uint8  `config:"field=8"`
		Uint16 []uint16 `config:"field=16"`
		Uint32 []uint32 `config:"field=32"`
		Uint64 []uint64 `config:"field=64"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "UINT",
	}
	env := newEnvsMock(map[string]string{
		"UINT_DEFAULT": "1,1",
		"UINT_8":       "8,8",
		"UINT_16":      "16,16",
		"UINT_32":      "32,32",
		"UINT_64":      "64,64",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal([]uint{1, 1}, config.Uint)
	assert.Equal([]uint8{8, 8}, config.Uint8)
	assert.Equal([]uint16{16, 16}, config.Uint16)
	assert.Equal([]uint32{32, 32}, config.Uint32)
	assert.Equal([]uint64{64, 64}, config.Uint64)
}

func TestConfigUintNegativeError(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Uint8 uint8 `config:"field=8"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "UINT",
	}
	env := newEnvsMock(map[string]string{
		"UINT_8": "-8",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.NotNil(err)
	assert.Equal("UINT_8: uint should be positive\n", err.Error())
}

func TestConfigSliceUintNegativeError(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Uint8 []uint8 `config:"field=8"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "UINT",
	}
	env := newEnvsMock(map[string]string{
		"UINT_8": "8,-8",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.NotNil(err)
	assert.Equal("UINT_8: uint should be positive\n", err.Error())
}

func TestConfigSliceString(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Text []string `config:""`
	}
	config := Config{}
	settings := Setting{
		Prefix: "STRING",
	}
	env := newEnvsMock(map[string]string{
		"STRING_TEXT": "a,b",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal([]string{"a", "b"}, config.Text)
}

func TestConfigSliceFloat(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Float32 []float32 `config:"field=32"`
		Float64 []float64 `config:"field=64"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "FLOAT",
	}
	env := newEnvsMock(map[string]string{
		"FLOAT_32": "32.5,132.5",
		"FLOAT_64": "64.5,164.5",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal([]float32{32.5, 132.5}, config.Float32)
	assert.Equal([]float64{64.5, 164.5}, config.Float64)
}

func TestConfigSliceSplitter(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Int []int `config:"field=default,splitter=|"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "INT",
	}
	env := newEnvsMock(map[string]string{
		"INT_DEFAULT": "1|2",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal([]int{1, 2}, config.Int)
}

func TestConfigRequired(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Field string `config:"required"`
	}
	config := Config{}
	settings := Setting{
		Title:  "Any config",
		Prefix: "REQUIRED",
	}
	env := newEnvsMock(map[string]string{})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Equal("Any config\nREQUIRED_FIELD: required\n", err.Error())
	assert.Equal("", config.Field)
}

func TestConfigRequiredWithoutTitle(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		A string `config:"required"`
		B string `config:""`
	}
	config := Config{}
	settings := Setting{
		Prefix: "REQUIRED",
	}
	env := newEnvsMock(map[string]string{})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Equal("REQUIRED_A: required\n", err.Error())
	assert.Equal("", config.A)
}

func TestFormatUrlSuccess(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Host  string   `config:"format=url"`
		Hosts []string `config:"format=url"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "URL",
	}
	env := newEnvsMock(map[string]string{
		"URL_HOST":  "https://domain.com/",
		"URL_HOSTS": "https://domain1.com/,https://domain2.com",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal("https://domain.com", config.Host)
	assert.Equal([]string{"https://domain1.com", "https://domain2.com"}, config.Hosts)
}

func TestFormatUrlError(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Host string `config:"format=url"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "URL",
	}
	env := newEnvsMock(map[string]string{
		"URL_HOST": "domain.com/",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.NotNil(err)
	assert.Equal("URL_HOST: parse \"domain.com/\": invalid URI for request\n", err.Error())
}

func TestFormatUrlErrorSlice(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Hosts []string `config:"format=url"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "URL",
	}
	env := newEnvsMock(map[string]string{
		"URL_HOSTS": "domain1.com/,https://domain2.com",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.NotNil(err)
	assert.Equal("URL_HOSTS: parse \"domain1.com/\": invalid URI for request\n", err.Error())
}

func TestEnumError(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Level string `config:"enum=info|error"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "ENUM",
	}
	env := newEnvsMock(map[string]string{
		"ENUM_LEVEL": "verbose",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.NotNil(err)
	assert.Equal("ENUM_LEVEL: enum has not valid value - 'verbose' not contained in the enum list\n", err.Error())
}

func TestEnumSliceError(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Envs []string `config:"enum=prod|dev"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "ENUM",
	}
	env := newEnvsMock(map[string]string{
		"ENUM_ENVS": "qa,local,prod",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.NotNil(err)
	assert.Equal("ENUM_ENVS: enum has not valid value - 'qa', 'local' not contained in the enum list\n", err.Error())
}

func TestEnumSliceSuccess(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Envs []string `config:"enum=prod|dev"`
	}
	config := Config{}
	settings := Setting{
		Prefix: "ENUM",
	}
	env := newEnvsMock(map[string]string{
		"ENUM_ENVS": "prod",
	})
	processor := InitConfig(&config, settings).SetEnv(env)
	err := processor.Process()

	assert.Nil(err)
}

func newEnvsMock(values map[string]string) IEnv {
	return &envsMock{values}
}

type envsMock struct {
	values map[string]string
}

func (e *envsMock) Get(name string) (value string, exist bool) {
	value, exist = e.values[name]
	return
}
