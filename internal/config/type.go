package config

import (
	"time"

	"github.com/soerjadi/brick/internal/pkg/customhttp"
)

var cfg *Config

type Config struct {
	Server     Server
	Database   DatabaseConfig
	Bank       Bank
	HttpClient customhttp.HttpHelper
}

type Server struct {
	Port        string
	WaitTimeout int
}

type DatabaseConfig struct {
	Name     string
	Driver   string
	Host     string
	Port     string
	User     string
	Password string
	SSL      string
}

type Bank struct {
	URL string
}

func (c Config) WaitTimeout() time.Duration {
	return time.Duration(c.Server.WaitTimeout)
}
