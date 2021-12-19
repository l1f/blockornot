package config

type EnvType string

const (
	EnvProd EnvType = "PROD"
	EnvDev  EnvType = "DEV"
)

type Config struct {
	Application struct {
		Env  EnvType
		Port int
	}
	Middlewares Middlewares `toml:"Options"`
}

type Middlewares struct {
	Cors struct {
		TrustedOrigins []string
	}
}
