package http

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/Salam4nder/inventory/internal/persistence"
	"github.com/Salam4nder/inventory/pkg/auth"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (s *Server) readItem(c *gin.Context) {
	uuid, found := c.Params.Get("uuid")
	if !found {
		c.JSON(http.StatusBadRequest, "uuid not found")
		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(), 5*time.Second)
	defer cancel()

	cachedItem, err := s.cache.Get(ctx, uuid)
	if err == nil {
		c.JSON(http.StatusOK, cachedItem)
		return
	}

	item, err := s.storage.Read(ctx, uuid)
	if err != nil {
		if errors.Is(err, persistence.ErrNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, item)
}

func (s *Server) readItems(c *gin.Context) {
	ctx, cancel := context.WithTimeout(
		c.Request.Context(), 5*time.Second)
	defer cancel()

	items, err := s.storage.ReadAll(ctx)
	if err != nil {
		if errors.Is(err, persistence.ErrNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, items)
}

func (s *Server) readItemsBy(c *gin.Context) {
	var filter persistence.ItemFilter

	if err := c.ShouldBindJSON(&filter); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(), 5*time.Second)
	defer cancel()

	items, err := s.storage.ReadBy(ctx, filter)
	if err != nil {
		if errors.Is(err, persistence.ErrNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, items)
}

func (s *Server) createItem(c *gin.Context) {
	var createRequest CreateItemRequest

	if err := c.ShouldBindJSON(&createRequest); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	item := createRequest.ToPersistenceItem()

	ctx, cancel := context.WithTimeout(
		c.Request.Context(), 5*time.Second)
	defer cancel()

	uuid, err := s.storage.Create(ctx, item)
	if err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := s.cache.Set(
		ctx,
		uuid.String(),
		item,
		time.Minute*20); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
	}

	c.JSON(http.StatusOK, uuid)
}

func (s *Server) updateItem(c *gin.Context) {
	var updateRequest UpdateItemRequest

	if err := c.ShouldBindJSON(updateRequest); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusBadRequest, err)
		return
	}

	item := updateRequest.ToPersistenceItem()

	ctx, cancel := context.WithTimeout(
		c.Request.Context(), 5*time.Second)
	defer cancel()

	updatedItem, err := s.storage.Update(
		ctx, item)
	if err != nil {
		if errors.Is(err, persistence.ErrNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	if err := s.cache.Set(
		ctx,
		updatedItem.ID.String(),
		updatedItem,
		time.Minute*20); err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
	}

	c.JSON(http.StatusOK, updatedItem)
}

func (s *Server) deleteItem(c *gin.Context) {
	uuid, found := c.Params.Get("uuid")
	if !found {
		c.JSON(http.StatusBadRequest, "uuid not found")
		return
	}

	ctx, cancel := context.WithTimeout(
		c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := s.storage.Delete(ctx, uuid); err != nil {
		if errors.Is(err, persistence.ErrNotFound) {
			c.JSON(http.StatusNotFound, err)
			return
		}
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	if err := s.cache.Delete(ctx, uuid); err != nil {
		s.logger.Info(err.Error(), zap.Error(err))
	}

	c.JSON(http.StatusOK,
		gin.H{"deleted": uuid})
}

// health check
func (s *Server) health(c *gin.Context) {
	var health []Health

	dbStatus := "Healthy"
	cacheStatus := "Healthy"
	serviceStatus := "Healthy"

	ctx, cancel := context.WithTimeout(
		c.Request.Context(), 5*time.Second)
	defer cancel()

	err := s.storage.Ping(ctx)
	if err != nil {
		dbStatus = "Unhealthy"
	}

	err = s.cache.Ping(ctx)
	if err != nil {
		cacheStatus = "Unhealthy"
	}

	health = append(health, Health{
		Service: "Database",
		Status:  dbStatus,
		Time:    time.Now().Local().Format(time.DateTime),
	})
	health = append(health, Health{
		Service: "Cache",
		Status:  cacheStatus,
		Time:    time.Now().Local().Format(time.DateTime),
	})
	health = append(health, Health{
		Service: "Service",
		Status:  serviceStatus,
		Time:    time.Now().Local().Format(time.DateTime),
	})

	c.JSON(http.StatusOK, health)
}

// Temporary, will add auth to the endpoint that creates
// a new JWT token.
func (s *Server) newJWT(c *gin.Context) {
	token, err := auth.NewJWT(
		s.config.JWTSecret, s.config.AccessTokenExpiration)
	if err != nil {
		s.logger.Error(err.Error(), zap.Error(err))
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
