package main

import (
	"log"

	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/internal/http"
	"github.com/Salam4nder/inventory/internal/persistence"
	"github.com/Salam4nder/inventory/pkg/logger"

	"github.com/stimtech/go-migration"
)

func main() {
	cfg, err := config.New()
	panicOnError(err)

	logger, err := logger.New("")
	panicOnError(err)

	store, err := persistence.New(
		&cfg.DB, persistence.PostgresDriver)
	panicOnError(err)
	defer store.DB.Close()
	logger.Info("Database connection established...")

	migration := migration.New(
		store.DB, logger)
	if err := migration.Migrate(); err != nil {
		panicOnError(err)
	}

	server := http.New(cfg.HTTP, store, logger)
	server.Start()
}

func panicOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
