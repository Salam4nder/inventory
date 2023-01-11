package persistence

import (
	"context"
	"database/sql"

	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/internal/entity"

	"github.com/google/uuid"
)

const (
	// PostgresDriver is a driver name for PostgreSQL.
	PostgresDriver = "postgres"
	// MySQLDriver is a driver name for MySQL.
	MySQLDriver = "mysql"
)

// Repository is a persistence layer interface.
type Repository interface {
	Create(ctx context.Context, item entity.Item) (
		uuid.UUID, error)
	Read(ctx context.Context, uuid string) (
		*entity.Item, error)
	ReadBy(ctx context.Context, filter entity.ItemFilter) (
		[]*entity.Item, error)
	Update(ctx context.Context, item *entity.Item) (
		*entity.Item, error)
	Delete(ctx context.Context, uuid string) error
}

// Storage is a persistence layer that implements
// the Repository interface.
type Storage struct {
	DB     *sql.DB
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
