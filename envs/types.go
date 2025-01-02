package envs

type ConfigSettingEnvs struct {
	Title  string // Config title
	Prefix string // Prefix of names for all environment variables
}

type IEnvGetter interface {
	Get(name string) (value string, exist bool)
}
