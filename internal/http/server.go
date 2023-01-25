package http

import (
	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/internal/cache"
	"github.com/Salam4nder/inventory/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server is the main structure of the API.
type Server struct {
	config  config.Server
	service domain.Service
	cache   cache.Provider
	logger  *zap.Logger
}

// New creates a new instance of the API server.
func New(
	cfg config.Server,
	srvc domain.Service,
	log *zap.Logger) *Server {
	return &Server{
		config:  cfg,
		service: srvc,
		logger:  log,
	}
}

// Start starts the server.
func (s *Server) Start() error {
	router := gin.Default()

	// router.GET("/health", s.health)
	jwtRout := router.Group("/api").Use(JWTAuth(
		s.config.JWTSecret))
	{
		jwtRout.GET("/item", s.readItems)
		jwtRout.GET("/item/:uuid", s.readItem)
		jwtRout.GET("/item/filter", s.readItemsBy)
		jwtRout.POST("/item", s.createItem)
		jwtRout.PUT("/item/:uuid", s.updateItem)
		jwtRout.DELETE("/item/:uuid", s.deleteItem)
	}

	return router.Run(":" + s.config.Port)
}
