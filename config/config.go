// Package config provides struct with app-level params.
package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		Server
	}

	Server struct {
		Port            string        `env:"APP_PORT" env-default:"8080"`
		ShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT" env-default:"5s"`
	}
)

// New loads ENV-vars and returns new Config instance.
func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("load env: %w", err)
	}
	return cfg, nil
}
