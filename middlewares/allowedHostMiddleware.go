package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var allowedHosts = []string{"localhost:8000", "gesthub.it"}

func AllowedHostsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		host := c.Request.Host
		bool := false
		for _, allowed := range allowedHosts {
			if host == allowed {
				bool = true
				break
			}
		}

		if !bool {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Access denided"})
			return
		}

		c.Next()
	}
}
