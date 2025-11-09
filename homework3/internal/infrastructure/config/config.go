package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
)

type Config struct {
	DbHost string          `env:"DB_HOST"`
	DbPort string          `env:"DB_PORT"`
	DbUser string          `env:"DB_USER"`
	DbPass string          `env:"DB_PASS"`
	DbName string          `env:"DB_NAME"`
	HTTPPort string        `env:"HTTP_PORT" envDefault:"8080"`
}

func Load() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("failed to parse environment variables: ", err.Error())
	}
	if cfg.DbHost == "" || cfg.DbUser == "" || cfg.DbName == "" || cfg.DbPort == "" || cfg.DbPass == "" {
		return nil, fmt.Errorf("DB_HOST, DB_USER, DB_NAME, DB_PORT, DB_PASS are required")
	}
	return cfg, nil
}
