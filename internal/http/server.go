package http

import (
	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/internal/service"

	"github.com/gin-gonic/gin"
)

// Server is the main structure of the API.
type Server struct {
	config  config.Server
	service service.Repository
}

// New creates a new instance of the API server.
func New(cfg config.Server, srvc service.Repository) *Server {
	return &Server{
		config:  cfg,
		service: srvc,
	}
}

// Start starts the server.
func (s *Server) Start() error {
	router := gin.Default()

	// router.GET("/health", s.health)
	jwtRout := router.Group("/api").Use(JWTAuth(
		s.config.JWTSecret))
	{
		jwtRout.GET("/item/:id", s.readItem)
		jwtRout.GET("/item", s.readItems)
		jwtRout.POST("/item", s.createItem)
		jwtRout.PUT("/item/:id", s.updateItem)
		jwtRout.DELETE("/item/:id", s.deleteItem)
	}

	return router.Run(":" + s.config.Port)
}
