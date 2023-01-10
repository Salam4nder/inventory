package api

import (
	"errors"

	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/pkg/inventory"
)

// Server is the main structure of the API.
type Server struct {
	config  config.Application
	service inventory.Service
}

// New creates a new instance of the API server.
func New(cfg config.Application, srvc inventory.Service) *Server {
	return &Server{
		config:  cfg,
		service: srvc,
	}
}

// Start starts the server.
func (s *Server) Start() error {
	return errors.New("not implemented")
}
