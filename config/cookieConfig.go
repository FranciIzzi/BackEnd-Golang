package config

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SetCookieHandler(w http.ResponseWriter, r *http.Request) {
	cookie := &http.Cookie{
		Name:     "session_id",
		Value:    "your_secure_session_id",
		Path:     "/",
		Domain:   "example.com",
		Expires:  time.Now().Add(24 * time.Hour),
		MaxAge:   86400,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	http.SetCookie(w, cookie)
	w.Write([]byte("Cookie impostato con successo"))
}

func deleteCookieHandler(c *gin.Context) {
    c.SetCookie("user", "", -1, "/", "localhost", false, true)
    c.String(http.StatusOK, "Cookie has been deleted")
}

func CookieMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
    cookie, err := c.Request.Cookie("session_id")
    if err != nil {
      if err == http.ErrNoCookie {
        c.AbortWithStatus(http.StatusUnauthorized)
        return
      }
      c.AbortWithStatus(http.StatusBadRequest)
      return
    }
    if cookie.Domain != "example.com" {
      c.AbortWithStatus(http.StatusUnauthorized)
      return
    }
    // TODO: validate cookie value
		c.Next()
	}
}

