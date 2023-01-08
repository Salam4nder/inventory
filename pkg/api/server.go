package api

import (
	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/pkg/inventory"
)

// Server is the main structure of the API.
type Server struct {
	config  *config.Application
	service *inventory.Service
}
