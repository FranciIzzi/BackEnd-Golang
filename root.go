package main

import (
	"log"
	"os"
  "root/sockets"
	"root/config"
	"root/middlewares"
	"root/models"
	routes "root/routers"
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

	//In Go, nil rappresenta l'assenza di un valore o un "null"
	if err != nil {
		log.Fatal("Impossibile connettersi al database:", err)
	}
	// log.Fatal(...) combina log.Print(...) per stampare il messaggio
	// specificato e poi chiama os.Exit(1) per terminare il programma
	// con uno stato di uscita 1, che è un codice di errore generico
	// che indica che il programma è terminato a causa di un errore
	// Migrazione dei modelli al database
  db.AutoMigrate(&models.User{})

	r := gin.Default()
	// set up dei middleware che sono applicati in tutte le route
	r.Use(cors.New(cors.Config{
    AllowOrigins:     []string{"https://localhost:8000"},           // PER DEFINIRE IN MANIERA STATICA CHI PUÒ MANDARE REQUEST
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"}, // DEFINISCE I METODI CONSENTITI
		AllowHeaders:     []string{"Origin"},                       //LA LISTA DEGLI HEADERS ACCETTATI IN INPUT
		ExposeHeaders:    []string{"Content-Length"},               //LA LISTA DEGLI HEADERS DA MANDARE INDIETRO AMMESSI
		AllowCredentials: true,                                     // PERMETTE DI INTEGRARE CREDENTIALS NELL'HEADER
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "https://github.com"
		// }, SERVE SOLO PER IMPLEMENTARE UN CHECK DINAMICO SULLE ORIGIN
		MaxAge: 3 * time.Hour, // TEMPO MASSIMO DEDICATI ALLE RICHIESTE "PREFLIGHT"
	}))
	r.Use(middlewares.AllowedHostsMiddleware())

	routes.UserRoute(r)
	routes.StateRoute(r)
  r.GET("/ws",sockets.WebSocketHandler)

	port := os.Getenv("PORT_backend")
	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Erore nell'avvio del server:", err)
	}
}
