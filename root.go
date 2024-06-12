package main

import (
	"log"
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
	db.AutoMigrate(&models.InumazioniModel{})
	db.AutoMigrate(&models.DefuntiModel{})
	db.AutoMigrate(&models.ContrattiModel{})
	db.AutoMigrate(&models.AventiDirittiModel{})

	r := gin.Default()
  r.Static("/media", "media")
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{
			"Content-Type",
			"Authorization",
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
	// r.Use(middlewares.AllowedHostsMiddleware())

	routes.UserRoute(r)
	routes.StateRoute(r)
	routes.CimiteriRoute(db, r)
	routes.InumazioniRoute(db, r)
	routes.DefuntiRoute(db, r)
	routes.ServiceRoute(db, r)
	routes.ContrattiRoute(db, r)
	routes.AventiDirittiRoute(db, r)
	r.GET("/ws", sockets.WebSocketHandler)

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
