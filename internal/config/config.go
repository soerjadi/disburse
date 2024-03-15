package config

import (
	"net/http"
	"os"

	"github.com/soerjadi/brick/internal/pkg/customhttp"
	"gopkg.in/gcfg.v1"
)

func Init() (*Config, error) {
	cfg = &Config{}

	configFilePath := "../../files/config.ini"

	config, err := os.ReadFile(configFilePath)
	if err != nil {
		return cfg, err
	}

	err = gcfg.ReadStringInto(cfg, string(config))
	if err != nil {
		return cfg, err
	}

	cfg.HttpClient = customhttp.NewHttpHelper(&http.Client{}, &customhttp.HttpConfig{})

	return cfg, nil
}

// GetConfig returns config object
func GetConfig() *Config {
	return cfg
}
