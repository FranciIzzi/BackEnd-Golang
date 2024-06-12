package controllers

import (
	// "bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"root/handlers"
	"root/models"
	"strconv"
	"time"

	// "io"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateInumazioni(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		real_body := c.PostForm("inumazione")
    log.Println(real_body)
		var input handlers.InumazioniRequest
		if err := json.Unmarshal([]byte(real_body), &input); err != nil {
			c.JSON(http.StatusBadRequest, 
        gin.H{
          "error": "Errore nel parsing JSON: " + err.Error(),
          "body": real_body,
        })
			return
		}
		if err := handlers.ValidateInumazioniRequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var inumazione models.InumazioniModel
		inumazione = models.InumazioniModel{
			CimiteroID:      *input.CimiteroID,
			Settore:         *input.Settore,
			CoordinataX:     *input.CoordinataX,
			CoordinataY:     *input.CoordinataY,
			NumeroCippo:     *input.NumeroCippo,
			ParcelNumber:    *input.ParcelNumber,
			StatoInumazione: *input.StatoInumazione,
			Occupato:        *input.Occupato,
			Tipologia:       *input.Tipologia,
		}

		file, err := c.FormFile("foto")
		if err == nil {
			baseDir := filepath.Join("media", "inumazioni", "foto")
			originalFilePath := filepath.Join(baseDir, filepath.Base(file.Filename))
			filePath := originalFilePath
			if _, err := os.Stat(filePath); !os.IsNotExist(err) {
				extension := filepath.Ext(file.Filename)
				fileNameWithoutExt := file.Filename[:len(file.Filename)-len(extension)]
				newFileName := fileNameWithoutExt + "_" + time.Now().
					Format("20060102-150405") +
					extension
				filePath = filepath.Join(baseDir, newFileName)
			}

			if err := c.SaveUploadedFile(file, filePath); err != nil {
				c.JSON(
					http.StatusInternalServerError,
					gin.H{"error": "Impossibile salvare il file: " + err.Error()},
				)
				return
			}
			inumazione.Foto = &filePath
		}
		if err := db.Create(&inumazione).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Inumazione creata con successo", "id": inumazione.ID})
		return
	}
}

func GetInumazioni(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var inumazioni []models.InumazioniModel
		query := db

		if idStr := c.Query("id"); idStr != "" {
			log.Println("ID:", idStr)
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

		if cimitero := c.Query("cimitero"); cimitero != "" {
			query = query.Where("cimitero_id = ?", cimitero)
		}
		if occupied := c.Query("occupied"); occupied != "" {
			occupiedBool, err := strconv.ParseBool(occupied)
			if err == nil {
				query = query.Where("occupato = ?", occupiedBool)
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "Occupato ha un valore non valido"})
			return
		}
		if tipologia := c.Query("tipologia"); tipologia != "" {
			query = query.Where("tipologia = ?", tipologia)
		}
		if err := query.Find(&inumazioni).Error; err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Errore nel recupero delle inumazioni"},
			)
			return
		}
		c.JSON(http.StatusOK, inumazioni)
		return
	}
}

func DeleteInumazioni(db *gorm.DB) gin.HandlerFunc {
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

		var inumazione models.InumazioniModel
		if err := db.Where("id = ?", id).First(&inumazione).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Inumazione non trovata"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		if err := db.Delete(&inumazione).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Inumazione eliminata con successo"})
	}
}

func UpdateInumazioni(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}
