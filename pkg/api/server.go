package api

import (
	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/pkg/service"
)

// Server is the main structure of the API.
type Server struct {
	config  *config.Config
	service *service.Service
}
