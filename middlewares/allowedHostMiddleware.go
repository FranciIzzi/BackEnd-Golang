package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var allowedHosts = []string{
	"http://localhost:8000",
	"http://0.0.0.0:8000",
	"http://127.0.0.1:8000",
	"http://192.168.1.9:8000",
	"http://192.168.1.18:8000",
}

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
