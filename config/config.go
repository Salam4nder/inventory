package config

import (
	"fmt"

	"github.com/plaid/go-envvar/envvar"
)

// Application is the main configuration structure for the application.
type Application struct {
	DB    Database `envvar:"DB_"`
	HTTP  Server   `envvar:"SERVER_"`
	Cache Cache    `envvar:"REDIS_"`
}

// Database is the database configuration.
type Database struct {
	Host     string `envvar:"HOST"`
	Port     string `envvar:"PORT"`
	User     string `envvar:"USER"`
	Name     string `envvar:"NAME"`
	Password string `envvar:"PASSWORD"`
}

// Server is the HTTP server configuration.
type Server struct {
	Port      string `envvar:"PORT"`
	JWTSecret string `envvar:"JWT"`
	LogFile   string `envvar:"LOG_FILE"`
}

// Cache is the cache configuration.
type Cache struct {
	Host     string `envvar:"HOST"`
	Port     string `envvar:"PORT"`
	Password string `envvar:"PASSWORD"`
}

// New parses the environment variables and returns a new Config.
// It returns an error if any env variables are unset.
func New() (*Application, error) {
	var appCfg Application

	if err := envvar.Parse(&appCfg); err != nil {
		return nil, err
	}

	return &appCfg, nil
}

// Connection returns the database connection string.
func (dbCfg *Database) Connection() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Name, dbCfg.Password)
	// todo enable sslmode?
}
