package config

import (
	"github.com/caarlos0/env"
)

// Config is the configuration for the server
type Config struct {
	Port           string `env:"SERVER_PORT" envDefault:":8080"`
	BotServiceAddr string `env:"BOT_SERVICE_ADDR" envDefault:":11111"`
}

// NewConfig generates a new configuration
func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
