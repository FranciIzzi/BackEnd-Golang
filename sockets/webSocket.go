package sockets

import (
	"net/http"
	"root/validators"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		return origin == "http://localhost:8000"
	},
}

func WebSocketHandler(c *gin.Context) {
	reqToken := c.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato token non valido"})
		return
	}
	tokenString := splitToken[1]
	_, err := validators.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token non valido"})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Impossibile stabilire la connessione WebSocket"},
		)
		return
	}

	defer conn.Close()

	for {
		err := conn.WriteMessage(websocket.TextMessage, []byte("Messaggio periodico"))
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Impossibile mandare messaggi dalla connesione WebSocket"},
			)
		}
		time.Sleep(10 * time.Second)
		// Legge il messaggio dal client
		// messageType, p, err := conn.ReadMessage()
		// if err != nil {
		// 	log.Printf("errore di lettura: %v", err)
		// 	break
		// }
		// // Echo del messaggio
		// if err := conn.WriteMessage(messageType, p); err != nil {
		// 	log.Printf("errore di scrittura: %v", err)
		// 	break
		// }
	}
}
