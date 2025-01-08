package envs_test

import (
	"testing"

	"github.com/be-true/config-go/envs"
	"github.com/be-true/config-go/errors"
	"github.com/stretchr/testify/assert"
)

func TestConfigString(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Text string `config:"format=url,require"`
	}
	config := Config{}
	settings := envs.SettingEnvs{
		Prefix: "STRING",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"STRING_TEXT": "domain.com",
	})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal("domain.com", config.Text)
}

func TestConfigBool(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Value bool `config:""`
	}
	config := Config{}
	settings := envs.SettingEnvs{
		Prefix: "BOOL",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"BOOL_VALUE": "true",
	})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(true, config.Value)
}

func TestConfigFloat(t *testing.T) {
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
	settings := envs.SettingEnvs{
		Prefix: "INT",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"INT_DEFAULT": "1",
		"INT_8":       "8",
		"INT_16":      "16",
		"INT_32":      "32",
		"INT_64":      "64",
	})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(1, config.Int)
	assert.Equal(int8(8), config.Int8)
	assert.Equal(int16(16), config.Int16)
	assert.Equal(int32(32), config.Int32)
	assert.Equal(int64(64), config.Int64)
}

func TestConfigInt(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Float32 float32 `config:"field=32"`
		Float64 float64 `config:"field=64"`
	}
	config := Config{}
	settings := envs.SettingEnvs{
		Prefix: "FLOAT",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"FLOAT_32": "32.5",
		"FLOAT_64": "64.5",
	})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal(float32(32.5), config.Float32)
	assert.Equal(float64(64.5), config.Float64)
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
	settings := envs.SettingEnvs{
		Prefix: "INT",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"INT_DEFAULT": "1,1",
		"INT_8":       "8,8",
		"INT_16":      "16,16",
		"INT_32":      "32,32",
		"INT_64":      "64,64",
	})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal([]int{1, 1}, config.Int)
	assert.Equal([]int8{8, 8}, config.Int8)
	assert.Equal([]int16{16, 16}, config.Int16)
	assert.Equal([]int32{32, 32}, config.Int32)
	assert.Equal([]int64{64, 64}, config.Int64)
}

func TestConfigSliceString(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		Text []string `config:""`
	}
	config := Config{}
	settings := envs.SettingEnvs{
		Prefix: "STRING",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"STRING_TEXT": "a,b",
	})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
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
	settings := envs.SettingEnvs{
		Prefix: "FLOAT",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"FLOAT_32": "32.5,132.5",
		"FLOAT_64": "64.5,164.5",
	})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
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
	settings := envs.SettingEnvs{
		Prefix: "INT",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{
		"INT_DEFAULT": "1|2",
	})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
	err := processor.Process()

	assert.Nil(err)
	assert.Equal([]int{1, 2}, config.Int)
}

func TestConfigRequired(t *testing.T) {
	assert := assert.New(t)
	t.Parallel()

	type Config struct {
		RequireField string `config:"require"`
	}
	config := Config{}
	settings := envs.SettingEnvs{
		Prefix: "REQUIRE",
	}
	envGetter := envs.NewEnvsGetterMock(map[string]string{})
	processor := envs.InitConfigEnvs(&config, settings).SetEnvGetter(envGetter)
	err := processor.Process()

	assert.Equal(errors.ErrorConfig, err)
	assert.Equal("", config.RequireField)
}
