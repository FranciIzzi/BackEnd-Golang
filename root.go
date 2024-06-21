package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"root/config"

	// "root/middlewares"
	"root/models"
	routes "root/routers"
	"root/sockets"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := config.ConnectDatabase()

	if err != nil {
		log.Fatal("Impossibile connettersi al database:", err)
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.CimiteriModel{})
	db.AutoMigrate(&models.SettoriModel{})
	db.AutoMigrate(&models.InumazioniModel{})
	db.AutoMigrate(&models.DefuntiModel{})
	db.AutoMigrate(&models.ContrattiModel{})
	db.AutoMigrate(&models.AventiDirittiModel{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Authorization",
			"Access-Control-Allow-Origin",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// }, SERVE SOLO PER IMPLEMENTARE UN CHECK DINAMICO SULLE ORIGIN
		MaxAge: 6 * time.Hour,
	}))

	r.Static("/media", "media")

	routes.UserRoute(r)
	routes.StateRoute(r)
	routes.CimiteriRoute(db, r)
	routes.SettoriRoute(db, r)
	routes.InumazioniRoute(db, r)
	routes.DefuntiRoute(db, r)
	routes.ServiceRoute(db, r)
	routes.ContrattiRoute(db, r)
	routes.AventiDirittiRoute(db, r)
	r.GET("/ws", sockets.WebSocketHandler)
	r.GET("/test", func(c *gin.Context) {
		imagePath := c.Query("path")
		file, err := os.Open(imagePath)
		if err != nil {
			c.String(http.StatusInternalServerError, "Failed to open image file")
			return
		}
		defer file.Close()

		c.Writer.Header().Set("Content-Type", "image/png")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		io.Copy(c.Writer, file)
	})

	port := os.Getenv("PORT_backend")
	if port == "" {
		log.Fatal("PORT_BACKEND non è impostato.")
	} else {
		err = r.Run(":" + port)
		if err != nil {
			log.Fatal("Errore nell'avvio del server:", err)
		}
	}
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Erore nell'avvio del server:", err)
	}
}
