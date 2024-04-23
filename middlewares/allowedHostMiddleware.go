package middlewares

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var allowedHosts = []string{"localhost:8000", "MyHub.it"}

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

		c.Next()	}
}
