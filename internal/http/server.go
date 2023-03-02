package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Salam4nder/inventory/internal/cache"
	"github.com/Salam4nder/inventory/internal/config"
	"github.com/Salam4nder/inventory/internal/persistence"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server is the main structure of the API.
type Server struct {
	http    *http.Server
	config  config.Server
	storage persistence.Storage
	cache   cache.Service
	logger  *zap.Logger
}

// Health is the structure for the health check endpoint.
type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

// New creates a new instance of the API server.
func New(
	cfg config.Server,
	store persistence.Storage,
	cache cache.Service,
	log *zap.Logger) *Server {
	srv := &http.Server{
		Addr: cfg.Addr(),
	}

	return &Server{
		http:    srv,
		config:  cfg,
		storage: store,
		cache:   cache,
		logger:  log,
	}
}

// Start starts the server.
func (s *Server) Start() {
	ctx, stop := signal.NotifyContext(
		context.Background(), os.Interrupt)
	defer stop()

	s.initEndpoints()

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

func (s *Server) initEndpoints() {
	router := gin.Default()
	router.GET("/auth", s.newJWT)
	router.GET("/health", s.health)

	authRoute := router.Group("/api").
		Use(jwtValidator(s.config.JWTSecret))
	{
		authRoute.GET("/item", s.readItems)
		authRoute.GET("/item/:uuid", s.readItem)
		authRoute.GET("/item/filter", s.readItemsBy)
		authRoute.POST("/item", s.createItem)
		authRoute.PUT("/item/:uuid", s.updateItem)
		authRoute.DELETE("/item/:uuid", s.deleteItem)
	}

	s.http.Handler = router
}
