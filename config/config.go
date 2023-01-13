package config

import (
	"fmt"

	"github.com/plaid/go-envvar/envvar"
)

// Application is the main configuration structure for the application.
type Application struct {
	DB   Database `envvar:"DB_"`
	HTTP Server   `envvar:"SERVER_"`
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
}

// New parses the environment variables and returns a new Config.
// It returns an error if any env variables are unset.
func New() (*Application, error) {
	var appCfg Application

	if err := envvar.Parse(&appCfg); err != nil {
		return nil, err
	}

	if err := appCfg.validate(); err != nil {
		return nil, err
	}

	return &appCfg, nil
}

// Connection returns the database connection string.
func (dbCfg *Database) Connection() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Name, dbCfg.Password)
	//todo enable sslmode?
}

// Envvar does this for us. Keeping in case
// we use a different library.
func (appCfg *Application) validate() error {
	if appCfg.DB.Host == "" {
		return fmt.Errorf("DB_HOST is required")
	}

	if appCfg.DB.Port == "" {
		return fmt.Errorf("DB_PORT is required")
	}

	if appCfg.DB.User == "" {
		return fmt.Errorf("DB_USER is required")
	}

	if appCfg.DB.Name == "" {
		return fmt.Errorf("DB_NAME is required")
	}

	if appCfg.DB.Password == "" {
		return fmt.Errorf("DB_PASSWORD is required")
	}

	if appCfg.HTTP.Port == "" {
		return fmt.Errorf("SERVER_PORT is required")
	}

	if appCfg.HTTP.JWTSecret == "" {
		return fmt.Errorf("SERVER_JWT is required")
	}

	return nil
}
