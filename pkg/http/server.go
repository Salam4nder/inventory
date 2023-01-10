package http

import (
	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/pkg/inventory"

	"github.com/gin-gonic/gin"
)

// Server is the main structure of the API.
type Server struct {
	Config  config.Server
	Service inventory.Service
}

// New creates a new instance of the API server.
func New(cfg config.Server, srvc inventory.Service) *Server {
	return &Server{
		Config:  cfg,
		Service: srvc,
	}
}

// Start starts the server.
func (s *Server) Start() error {
	router := gin.Default()
	// router.GET("/health", s.health)
	router.GET("/item/:id", s.getItem)
	router.GET("/item", s.getItems)
	router.POST("/item", s.createItem)
	router.PUT("/item/:id", s.updateItem)
	router.DELETE("/item/:id", s.deleteItem)

	return router.Run(":" + s.Config.Port)
}
