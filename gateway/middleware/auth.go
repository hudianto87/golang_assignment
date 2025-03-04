package middleware

import (
	"golangassignment/gateway/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username, pass, ok := c.Request.BasicAuth()

		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			c.Abort()
			return
		}

		isvalid := (username == config.AuthBasicUsername) && (pass == config.AuthBasicPassword)
		if !isvalid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization"})
			c.Abort()
			return
		}

		c.Next()
	}
}
