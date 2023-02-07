package http

import (
	"github.com/Salam4nder/inventory/pkg/auth"

	"github.com/gin-gonic/gin"
)

// JWTAuth is a middleware that checks if the
// request has a valid JWT token.
func JWTAuth(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(401, gin.H{"error": " JWT token required"})
			c.Abort()
			return
		}

		if err := auth.ValidateJWT(token, secret); err != nil {
			c.JSON(401, gin.H{"error": "JWT token is invalid"})
			c.Abort()
			return
		}

		c.Next()
	}
}
