package controllers

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"root/handlers"
	"root/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateAventiDiritti(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("URL:", c.Request.URL.Path)
		bodyBytes, err := io.ReadAll(io.Reader(c.Request.Body))
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Errore nella lettura del body: " + err.Error()},
			)
			return
		}
		c.Request.Body = io.NopCloser(io.Reader(bytes.NewBuffer(bodyBytes)))
		log.Println("Corpo della richiesta:", string(bodyBytes))

		var input []handlers.AventiDirittiRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := handlers.ValidateAventiDirittirequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		for _, instance := range input {
			var aventiDiritti models.AventiDirittiModel
			aventiDiritti = models.AventiDirittiModel{
				DefuntoID:     *instance.DefuntoID,
				Nome:          *instance.Nome,
				Cognome:       *instance.Cognome,
				CodiceFiscale: instance.CodiceFiscale,
				Email:         instance.Email,
				Telefono:      instance.Telefono,
			}
			if err := db.Create(&aventiDiritti).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "Tutti gli aventi diritti creati con successo"})
		return
	}
}

func GetAventiDiritti(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var aventiDiritti []models.AventiDirittiModel
		query := db

		if idStr := c.Query("id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err == nil {
				if id <= 0 {
					c.JSON(
						http.StatusBadRequest,
						gin.H{"error": "ID deve essere un numero positivo"},
					)
					return
				}
				query = query.Where("id = ?", id)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero"})
				return
			}
		}

		if defunto := c.Query("defunto"); defunto != "" {
			defuntoID, err := strconv.Atoi(defunto)
			if err == nil {
				if defuntoID <= 0 {
					c.JSON(
						http.StatusBadRequest,
						gin.H{"error": "DefuntoID deve essere un numero positivo"},
					)
					return
				}
				query = query.Where("defunto_id = ?", defuntoID)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "DefuntoID deve essere un numero"})
				return
			}
		}
		if nome := c.Query("nome"); nome != "" {
			query = query.Where("nome = ?", nome)
		}
		if cognome := c.Query("cognome"); cognome != "" {
			query = query.Where("cognome = ?", cognome)
		}
		if codiceFiscale := c.Query("codiceFiscale"); codiceFiscale != "" {
			query = query.Where("codice_fiscale = ?", codiceFiscale)
		}
		if email := c.Query("email"); email != "" {
			query = query.Where("email = ?", email)
		}
		if telefono := c.Query("telefono"); telefono != "" {
			query = query.Where("telefono = ?", telefono)
		}
		if err := query.Find(&aventiDiritti).Error; err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Errore nel recupero degli aventi diritti"},
			)
			return
		}
		c.JSON(http.StatusOK, aventiDiritti)
		return
	}
}

func DeleteAventiDiritti(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID mancante nella richiesta"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero"})
			return
		}

		if id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero positivo"})
			return
		}

		var aventiDiritti models.AventiDirittiModel
		if err := db.Where("id = ?", id).First(&aventiDiritti).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Aventi Diritti non trovato"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		if err := db.Delete(&aventiDiritti).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Aventi Diritti eliminato con successo"})
	}
}

func UpdateAventiDiritti(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")

		if idStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID mancante nella richiesta"})
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero"})
			return
		}

		if id <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero positivo"})
			return
		}
		var input handlers.AventiDirittiRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		list := []handlers.AventiDirittiRequest{input}
		if err := handlers.ValidateAventiDirittirequest(db, &list); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var aventiDiritti models.AventiDirittiModel

		if err := db.Where("id = ?", id).First(&aventiDiritti).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Aventi Diritti non trovato"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		if input.ID != aventiDiritti.ID {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "ID dell'aventi diritti non corrispondente"},
			)
			return
		}
		aventiDiritti.DefuntoID = *input.DefuntoID
		aventiDiritti.Nome = *input.Nome
		aventiDiritti.Cognome = *input.Cognome
		aventiDiritti.CodiceFiscale = input.CodiceFiscale
		aventiDiritti.Email = *&input.Email
		aventiDiritti.Telefono = input.Telefono

		if err := db.Save(&aventiDiritti).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Aventi Diritti aggiornato con successo"})
		return
	}
}
