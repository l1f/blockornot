package config

import (
	"errors"
	"github.com/pelletier/go-toml/v2"
	"io/fs"
	"io/ioutil"
	"os"
)

var (
	ErrorWrongPermission = errors.New("insufficient permission to access the configuration")
	ErrorFileNotFound    = errors.New("the config file could not be found under the specified path")
)

func ReadFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		switch {
		case errors.Is(err, os.ErrNotExist):
			return nil, ErrorFileNotFound
		case errors.Is(err, os.ErrPermission):
			return nil, ErrorWrongPermission
		default:
			return nil, err
		}
	}

	var config Config
	err = toml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func GenerateNewConfigFile(path string) error {
	emptyConfig := Config{
		Application: struct {
			Env    EnvType
			Port   int
			Secret string
		}{
			Env:    EnvProd,
			Port:   8080,
			Secret: "",
		},
		Middlewares: middlewares{
			Cors: struct {
				TrustedOrigins []string
			}{},
		},
		Twitter: twitter{
			ConsumerKey:    "",
			ConsumerSecret: "",
		},
	}
	data, err := toml.Marshal(emptyConfig)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path, data, fs.FileMode(0640))
	if err != nil {
		return err
	}

	return nil
}
