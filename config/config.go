package config

import (
	"fmt"

	"github.com/plaid/go-envvar/envvar"
)

// Application is the main configuration structure for the application.
type Application struct {
	DB    Database `envvar:"POSTGRES_"`
	HTTP  Server   `envvar:"SERVER_"`
	Cache Cache    `envvar:"REDIS_"`
}

// Database is the database configuration.
type Database struct {
	Host     string `envvar:"HOST"`
	Port     string `envvar:"PORT"`
	User     string `envvar:"USER"`
	Name     string `envvar:"DB"`
	Password string `envvar:"PASSWORD"`
}

// Server is the HTTP server configuration.
type Server struct {
	Host      string `envvar:"HOST"`
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

// PSQLConn returns a psql connection string.
func (dbCfg *Database) PSQLConn() string {
	// return a psql connection string
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		dbCfg.User, dbCfg.Password, dbCfg.Host,
		dbCfg.Port, dbCfg.Name)
}

// Addr returns the configured server address.
func (srvCfg *Server) Addr() string {
	return fmt.Sprintf("%s:%s", srvCfg.Host, srvCfg.Port)
}
