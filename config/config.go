package config

import (
	"os"
)

type Config struct {
	Http              Http
	RequestsTimeLimit int
	RateLimit         int
}

type Http struct {
	Host string
	Port string
}

// New get new instance of Config
func New() *Config {
	return &Config{
		Http: Http{
			Host: os.Getenv("APP_HTTP_HOST"),
			Port: os.Getenv("APP_HTTP_PORT"),
		},
		RequestsTimeLimit: 60,
		RateLimit:         5,
	}
}
