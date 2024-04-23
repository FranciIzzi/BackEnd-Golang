package controllers
func Table() {
	return
}

// GET /query-table?table=users&fields=id,username,email

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/lib/pq" // Importa il driver PostgreSQL
// )

// type CreateTableRequest struct {
// 	TableName string `json:"table_name"`
// 	Fields    []struct {
// 		Name string `json:"name"`
// 		Type string `json:"type"`
// 	} `json:"fields"`
// }

// func createTableHandler(c *gin.Context) {
// 	var req CreateTableRequest
// 	if err := c.ShouldBindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//	return
// 	}

// 	// Connessione al database (esempio con PostgreSQL)
// 	db, err := sql.Open("postgres", "host=localhost user=youruser dbname=yourdb sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	// Costruisci la query SQL per creare la tabella
// 	query := fmt.Sprintf("CREATE TABLE %s (", req.TableName)
// 	for i, field := range req.Fields {
// 		query += fmt.Sprintf("%s %s", field.Name, field.Type)
// 		if i < len(req.Fields)-1 {
// 			query += ", "
// 		}
// 	}
// 	query += ");"

// 	// Esegui la query
// 	if _, err := db.Exec(query); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossibile creare la tabella"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"success": "Tabella creata"})
// }

// func main() {
// 	router := gin.Default()
// 	router.POST("/create-table", createTableHandler)
// 	router.Run(":8080")
// }
