package envs_test

import (
	"testing"

	"github.com/be-true/config-go/envs"
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
	processor := envs.InitConfigEnvs(settings).SetEnvGetter(envGetter)
	err := processor.Process(&config)

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
	processor := envs.InitConfigEnvs(settings).SetEnvGetter(envGetter)
	err := processor.Process(&config)

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
	config := Config{
		Int:   1,
		Int8:  8,
		Int16: 16,
		Int32: 32,
		Int64: 64,
	}
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
	processor := envs.InitConfigEnvs(settings).SetEnvGetter(envGetter)
	err := processor.Process(&config)

	assert.Nil(err)
	assert.Equal(1, config.Int)
	assert.Equal(int8(8), config.Int8)
	assert.Equal(int16(16), config.Int16)
	assert.Equal(int32(32), config.Int32)
	assert.Equal(int64(64), config.Int64)
}
