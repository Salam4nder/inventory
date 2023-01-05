package config

import (
	"fmt"

	"github.com/plaid/go-envvar/envvar"
)

// Configuration is the main interface for the application configuration.
type Configuration interface {
	NewConfig() (*Config, error)
}

// Config is the main configuration structure for the application.
// It implements the Configuration interface.
type Config struct {
	Db Database `envvar:"DB_"`
}

// Database is the database configuration.
type Database struct {
	Host     string `envvar:"HOST"`
	Port     string `envvar:"PORT"`
	User     string `envvar:"USER"`
	Name     string `envvar:"NAME"`
	Password string `envvar:"PASSWORD"`
}

// NewConfig parses the environment variables and returns a new Config.
// It returns an error if any env variables are unset.
func NewConfig() (*Config, error) {
	var cfg Config

	if err := envvar.Parse(&cfg); err != nil {
		return nil, err
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) validate() error {
	if cfg.Db.Host == "" {
		return fmt.Errorf("DB_HOST is required")
	}

	if cfg.Db.Port == "" {
		return fmt.Errorf("DB_PORT is required")
	}

	if cfg.Db.User == "" {
		return fmt.Errorf("DB_USER is required")
	}

	if cfg.Db.Name == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	if cfg.Db.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	return nil
}
