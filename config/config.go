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
		Port            string        `env:"SERVER_PORT" env-default:"8080"`
		ShutdownTimeout time.Duration `env:"SERVER_SHUTDOWN_TIMEOUT" env-default:"5s"`
		WS
	}

	WS struct {
		ReadBufferSize  int           `env:"SERVER_WS_READ_BUF_SIZE" env-default:"1024"`
		WriteBufferSize int           `env:"SERVER_WS_WRITE_BUF_SIZE" env-default:"1024"`
		PongWait        time.Duration `env:"SERVER_WS_PONG_WAIT" env-default:"60s"`
		PingPeriod      time.Duration
	}
)

// New loads ENV-vars and returns new Config instance.
func New() (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("load env: %w", err)
	}
	// set ping period depending on wait pong period
	cfg.Server.WS.PingPeriod = cfg.Server.WS.PongWait * 9 / 10
	return cfg, nil
}
