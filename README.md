# Config system for go lang apps

## TODO

- Default value
- tests for errors
- Enum
- Url formatter
- Formatter and type matching

## Example of usage

- If you described config like below
```go
type Config struct {
	Env      string   `config:"enum=local|dev|qa|stage|prod"`
	Port     int      `config:"format=port,default=3000"`
	Hosts    []string `config:"format=url,require,splitter=|,doc='You doc info'"`
	Enabled  bool     `config:"format=boolean,default=true"`
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
APPLICATION_HOSTS=domain1.com:5555|domain2.com:5555
APPLICATION_ENABLED=true
```

- Config should be equal:
```go
Config{
    Env: "prod",
    Port: 3000,
    Hosts: []{"domain1.com:5555", "domain2.com:5555"},
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
| error    | env               | format | doc          | example                      |
|----------|-------------------|--------|--------------|------------------------------|
| required | APPLICATION_HOSTS | url    | You doc info | domain1:port1\|domain2:port2 |
```
