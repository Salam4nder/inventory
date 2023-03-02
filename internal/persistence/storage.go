package persistence

import (
	"context"
	"database/sql"

	"github.com/Salam4nder/inventory/internal/config"

	"github.com/google/uuid"
	// --nolint:staticcheck.
	_ "github.com/lib/pq"
)

const (
	// PostgresDriver is a driver name for PostgreSQL.
	PostgresDriver = "postgres"
	// MySQLDriver is a driver name for MySQL.
	MySQLDriver = "mysql"
)

// Storage is a persistence layer interface
// with basic CRUD operations.
type Storage interface {
	Create(ctx context.Context, item Item) (
		uuid.UUID, error)
	Read(ctx context.Context, uuid string) (
		Item, error)
	ReadAll(ctx context.Context) ([]Item, error)
	ReadBy(ctx context.Context, filter ItemFilter) (
		[]Item, error)
	Update(ctx context.Context, item Item) (
		Item, error)
	Delete(ctx context.Context, uuid string) error
	Ping(ctx context.Context) error
}

// SQLDatabase implements the Storage interface.
type SQLDatabase struct {
	DB     *sql.DB
	Config config.Database
}

// Ping checks the connection to the database.
func (s *SQLDatabase) Ping(ctx context.Context) error {
	return s.DB.PingContext(ctx)
}

// New creates a new SQLDatabase instance from the given
// config.Database and driver string.
// Drivers are provided in this package as constants.
func New(
	dbCfg config.Database, driver string) (*SQLDatabase, error) {
	db, err := sql.Open(driver, dbCfg.PSQLConn())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &SQLDatabase{DB: db, Config: dbCfg}, nil
}
