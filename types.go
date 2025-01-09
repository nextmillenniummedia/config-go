package configgo

type Setting struct {
	Title  string // Config title
	Prefix string // Prefix of names for all environment variables
}

type IEnv interface {
	Get(name string) (value string, exist bool)
}
