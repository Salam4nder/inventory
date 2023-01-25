package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Salam4nder/inventory/config"
	"github.com/Salam4nder/inventory/internal/domain"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server is the main structure of the API.
type Server struct {
	http    *http.Server
	config  config.Server
	service domain.Service
	logger  *zap.Logger
}

// New creates a new instance of the API server.
func New(
	cfg config.Server,
	srvc domain.Service,
	log *zap.Logger) *Server {
	srv := &http.Server{
		Addr: cfg.Addr(),
	}
	return &Server{
		http:    srv,
		config:  cfg,
		service: srvc,
		logger:  log,
	}
}

// Start starts the server.
func (s *Server) Start() {
	router := gin.Default()

	ctx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt)
	defer stop()

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

	s.http.Handler = router

	go func() {
		if err := s.http.ListenAndServe(); err != nil &&
			err != http.ErrServerClosed {
			s.logger.Fatal("listen: ", zap.Error(err))
		}
	}()

	<-ctx.Done()

	stop()
	s.logger.Info("shutting down the server...")

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	if err := s.http.Shutdown(ctx); err != nil {
		s.logger.Fatal("server shutdown: ", zap.Error(err))
	}
}
