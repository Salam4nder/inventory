package persistence

import (
	"database/sql"

	"github.com/Salam4nder/inventory/config"
)

const (
	// PostgresDriver is a driver name for PostgreSQL.
	PostgresDriver = "postgres"
	// MySQLDriver is a driver name for MySQL.
	MySQLDriver = "mysql"
)

// Storage is a persistence layer interface.
type Storage interface{}

// Database is a persistence layer that implements
// the Storage interface.
type Database struct {
	DB     *sql.DB
	Config config.Database
}

// New creates a new Database instance from the given
// config.Database instance and driver string.
// Drivers are provided in this package as constants.
func New(
	dbCfg *config.Database, driver string) (*Database, error) {
	db, err := sql.Open(driver, dbCfg.Connection())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{DB: db, Config: *dbCfg}, nil
	//TODO: defer db.Close() in entry point
}
