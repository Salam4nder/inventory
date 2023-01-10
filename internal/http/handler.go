package http

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) readItem(c *gin.Context) {
	uuid := c.Param("id")

	ctx, cancel := context.WithTimeout(
		context.Background(), 5*time.Second)
	defer cancel()

	item, err := s.service.Read(ctx, uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (s *Server) readItems(c *gin.Context) {
	//TODO: implement
}

func (s *Server) createItem(c *gin.Context) {
	//TODO: implement
}

func (s *Server) updateItem(c *gin.Context) {
	//TODO: implement
}

func (s *Server) deleteItem(c *gin.Context) {
	//TODO: implement
}
