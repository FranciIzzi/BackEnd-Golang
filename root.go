package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"root/config"
	"root/middlewares"
	"root/models"
	routes "root/routers"
	"root/sockets"
	"time"
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
  db.AutoMigrate(&models.InumazioniModel{})

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8000",
		},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Origin",
		},
		ExposeHeaders: []string{
			"Content-Length",
		},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// }, SERVE SOLO PER IMPLEMENTARE UN CHECK DINAMICO SULLE ORIGIN
		MaxAge: 3 * time.Hour,
	}))
	r.Use(middlewares.AllowedHostsMiddleware())

	routes.UserRoute(r)
	routes.StateRoute(r)
  routes.CimiteriRoute(db,r)
	r.GET("/ws", sockets.WebSocketHandler)

	port := os.Getenv("PORT_BACKEND")
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Erore nell'avvio del server:", err)
	}
}
