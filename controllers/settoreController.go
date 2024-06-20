package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
	"root/handlers"
	"root/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateSettori(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		real_body := c.PostForm("settore")
		var input handlers.SettoriRequest
		if err := json.Unmarshal([]byte(real_body), &input); err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"error": "Errore nel parsing JSON: " + err.Error(),
				})
			return
		}
		if err := handlers.ValidateSettoriRequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var settore models.SettoriModel
		settore = models.SettoriModel{
			CimiteroID:        *input.CimiteroID,
			Nome:              *input.Nome,
			Righe:             *input.Righe,
			Colonne:           *input.Colonne,
			InumazioniSettore: *input.InumazioniSettore,
			PuntiCanvas:       *input.PuntiCanvas,
			Color:             *input.Color,
		}
		file, err := c.FormFile("snapshot")
		if err == nil {
			baseDir := filepath.Join("media", "settori", "snapshots")
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
			settore.Snapshot = filePath
		}
		if err := db.Create(&settore).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Settore creato con successo"})
		return
	}
}

func GetSettori(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var settori []models.SettoriModel
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

		if cimitero := c.Query("cimitero"); cimitero != "" {
			cimiteroID, err := strconv.Atoi(cimitero)
			if err == nil {
				if cimiteroID <= 0 {
					c.JSON(
						http.StatusBadRequest,
						gin.H{"error": "CimiteroID deve essere un numero positivo"},
					)
					return
				}
				query = query.Where("cimitero_id = ?", cimiteroID)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "CimiteroID deve essere un numero"})
				return
			}
		}
		if nome := c.Query("nome"); nome != "" {
			query = query.Where("nome = ?", nome)
		}
		if err := query.Find(&settori).Error; err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Errore nel recupero dei settori"},
			)
			return
		}
		c.JSON(http.StatusOK, settori)
		return
	}
}

func DeleteSettori(db *gorm.DB) gin.HandlerFunc {
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

		var settore models.SettoriModel
		if err := db.Where("id = ?", id).First(&settore).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Settore non trovato"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		if err := db.Delete(&settore).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Settore eliminato con successo"})
	}
}

func UpdateSettori(db *gorm.DB) gin.HandlerFunc {
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
		var input handlers.SettoriRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := handlers.ValidateSettoriRequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var settore models.SettoriModel
		if err := db.Where("id = ?", id).First(&settore).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Settore non trovato"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		if input.ID != settore.ID {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "ID del settore non corrispondente"},
			)
			return
		}
		settore.Nome = *input.Nome
		settore.Righe = *input.Righe
		settore.Colonne = *input.Colonne
		settore.Color = *input.Color
		settore.CimiteroID = *input.CimiteroID
		settore.Snapshot = *input.Snapshot

		if err := db.Save(&settore).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Settore aggiornato con successo"})
		return
	}
}
