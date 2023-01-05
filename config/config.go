package config

import (
	"fmt"
	"log"

	env "github.com/Netflix/go-env"
)

// Configuration is the main interface for the application configuration.
type Configuration interface {
	NewConfig() (*Config, error)
}

// Config is the main configuration structure for the application.
// It implements the Configuration interface.
type Config struct {
	db Database
}

// Database is the database configuration
type Database struct {
	Host     string `env:"DB_HOST, defualt:localhost, required:true"`
	Port     string `env:"DB_PORT, default:5432, required:true"`
	User     string `env:"DB_USER, required:true"`
	Password string `env:"DB_PASSWORD, required:true"`
	Name     string `env:"DB_NAME, required:true"`
}

// NewConfig parses the environment variables and returns a new Config.
// It returns an error if any env variables are unsupported or unset
func NewConfig() (*Config, error) {
	cfg := &Config{}

	unSet, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Printf("Unset variables: %v", unSet)
		return nil, fmt.Errorf(
			"error parsing environment variables: %v", err)
	}

	return cfg, nil
}
