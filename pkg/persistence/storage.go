package persistence

import (
	"database/sql"

	"github.com/Salam4nder/inventory/config"
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
// config.Database instance.
func New(dbCfg *config.Database) (*Database, error) {
	db, err := sql.Open("postgres", dbCfg.Connection())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Database{DB: db, Config: *dbCfg}, nil
	//TODO: defer db.Close() in entry point
}
