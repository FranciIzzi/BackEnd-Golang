package middlewares

import (
	"net/http"
	"strings"
	"root/validators"	
  "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		reqToken := c.GetHeader("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		if len(splitToken) != 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato token non valido"})
			return
		}

		reqToken = splitToken[1]
		email, err := validators.ValidateToken(reqToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token non valido"})
			return
		}

		authorized := false
		allowed := []string{"francesco@dedpartners.com", "cristiano@dedpartners.com", "mattia@dedpartners.com"}
		for _, allowedEmail := range allowed {
			if email == allowedEmail {
				authorized = true
				break
			}
		}

		if !authorized {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Accesso non autorizzato"})
			return
		}

		c.Next()
	}
}
