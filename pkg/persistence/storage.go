package persistence

import (
	"context"
	"database/sql"

	"github.com/Salam4nder/inventory/config"
)

const (
	// PostgresDriver is a driver name for PostgreSQL.
	PostgresDriver = "postgres"
	// MySQLDriver is a driver name for MySQL.
	MySQLDriver = "mysql"
)

// Repository is a persistence layer interface.
type Repository interface{}

// Database is an abstraction over the database/sql.DB.
// Mostly used for testing purposes.
type Database interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	QueryRowContext(
		ctx context.Context, query string, args ...interface{}) *sql.Row
    QueryContext(
        ctx context.Context, query string, args ...any) (*sql.Rows, error)
	ExecContext(
		ctx context.Context, query string, args ...any) (sql.Result, error)
}

// Storage is a persistence layer that implements
// the Storage interface.
type Storage struct {
	DB     Database
	Config config.Database
}

// New creates a new Database instance from the given
// config.Database and driver string.
// Drivers are provided in this package as constants.
func New(
	dbCfg *config.Database, driver string) (*Storage, error) {
	db, err := sql.Open(driver, dbCfg.Connection())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{DB: db, Config: *dbCfg}, nil
	//TODO: defer db.Close() in entry point
}
