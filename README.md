# Config system for go lang apps

### Install

```bash
go get github.com/nextmillenniummedia/config-go
```

## Allowed types

- `int`, `[]int`
- `uint`, `[]uint`
- `string`, `[]string` with format: `url`
- `float`, `[]int`
- `bool`
- `time.Duration` with format: `ns`, `us`, `ms`, `s`, `m`, `h` 

## Allowed params in config notation

- `required` - This field is required
- `field` - Field name for env
- `splitter` - The char used to split the env value 
- `enum` - Enum of allowed value
- `format` - Format of value for other types
- `doc` - Addition info for errors or documentation

## Example of usage

- If you described config like below
```go
type Config struct {
	Env      string   `config:"enum=local|dev|qa|stage|prod"`
	Port     int      `config:"default=3000"`
	Hosts    []string `config:"format=url,required,splitter=|,doc='You doc info'"`
	Enabled  bool     `config:"default=true"`
}
```

- Than you call parsing envs variable to this config
```go
inst := Config{}
settings := envs.SettingsEnv{
    Prefix: "APPLICATION",
}
config.InitConfigEnvs(settings)
err := config.Process(&inst)
```

In system has variables:
```bash
APPLICATION_ENV=prod
APPLICATION_PORT=3000
APPLICATION_HOSTS=http://domain1.com:5555/|http://domain2.com:5555
APPLICATION_ENABLED=true
```

- Config should be equal:
```go
Config{
    Env: "prod",
    Port: 3000,
    Hosts: []{"http://domain1.com:5555", "http://domain2.com:5555"},
    Enabled: true,
}
```

- If in config has errors, than you may call method for visualize errors in terminal

```go
if err != nil {
    log.Fatal(err)
}
```

- Errors will printed in terminal as md table
```bash
| error    | env               | format | doc          | example                                    |
|----------|-------------------|--------|--------------|--------------------------------------------|
| required | APPLICATION_HOSTS | url    | You doc info | http://domain1:port1|http://domain2:port2  |
```

## TODO

- CamelToUnderscore converter for name fields
- ip format
- pointer
- tests for errors
- formatter and type matching
- export config to md table
