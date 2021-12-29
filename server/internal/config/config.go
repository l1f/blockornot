package config

type EnvType string

const (
	EnvProd EnvType = "PROD"
	EnvDev  EnvType = "DEV"
)

type Config struct {
	Application application `toml:"App"`
	Middlewares middlewares `toml:"Options"`
	Twitter     twitter     `toml:"twitter"`
}

type application struct {
	Env    EnvType
	Port   int
	Secret string
}

type middlewares struct {
	Cors struct {
		TrustedOrigins []string
	}
}

type twitter struct {
	ConsumerKey    string
	ConsumerSecret string
}
