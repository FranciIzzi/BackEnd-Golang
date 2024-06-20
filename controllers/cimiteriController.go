package controllers

import (
	// "bytes"
	"encoding/json"
	// "io"
	// "log"
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

func GetCimiteri(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Query("id")
		comune := c.Query("comune")

		var err error

		if idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero"})
				return
			}

			if id <= 0 {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero positivo"})
				return
			}

			var cimitero models.CimiteriModel
			if err = db.Where("id = ?", id).First(&cimitero).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					c.JSON(http.StatusNotFound, gin.H{"error": "Cimitero non trovato"})
				} else {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				}
				return
			}

			c.JSON(http.StatusOK, cimitero)
			return
		}

		if comune != "" {
			var cimiteri []models.CimiteriModel
			if err = db.Where("comune = ?", comune).Find(&cimiteri).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, cimiteri)
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": "Deve essere passato un ID o un comune"})
		return
	}
}

func DeleteCimitero(db *gorm.DB) gin.HandlerFunc {
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

		var cimitero models.CimiteriModel

		if err := db.Where("id = ?", id).First(&cimitero).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Cimitero non trovato"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		if err := db.Delete(&cimitero).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Cimitero eliminato con successo"})
		return
	}
}

func CreateCimitero(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		real_body := c.PostForm("cimitero")
		var input handlers.CimiteriRequest
		if err := json.Unmarshal([]byte(real_body), &input); err != nil {
			c.JSON(http.StatusBadRequest,
				gin.H{
					"error": "Errore nel parsing JSON: " + err.Error(),
				})
			return
		}
		if err := handlers.ValidateCimiteriRequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
    var cimitero models.CimiteriModel
		cimitero = models.CimiteriModel{
			Latitudine:          *input.Latitudine,
			Longitudine:         *input.Longitudine,
			Regione:             *input.Regione,
			Provincia:           *input.Provincia,
			Comune:              *input.Comune,
			PostiTotali:         *input.PostiTotali,
			InumazioniPresenti:  *input.InumazioniPresenti,
			RotazioneEsumazioni: *input.RotazioneEsumazioni,
		}

		file, err := c.FormFile("planimetria")
		if err == nil {
			baseDir := filepath.Join("media", "cimiteri", "planimetrie")
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
			cimitero.Planimetria = filePath
		}
		if err := db.Create(&cimitero).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Cimitero creato con successo", "cimitero": cimitero.ID})
		return
	}
}

func UpdateCimitero(db *gorm.DB) gin.HandlerFunc {
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

		var input handlers.CimiteriRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := handlers.ValidateCimiteriRequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var cimitero models.CimiteriModel

		if err := db.Where("id = ?", id).First(&cimitero).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Cimitero non trovato"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		if input.ID != cimitero.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID del cimitero non corrispondente"})
			return
		}
		cimitero.Latitudine = *input.Latitudine
		cimitero.Longitudine = *input.Longitudine
		cimitero.Regione = *input.Regione
		cimitero.Provincia = *input.Provincia
		cimitero.Comune = *input.Comune
		cimitero.PostiTotali = *input.PostiTotali
		cimitero.InumazioniPresenti = *input.InumazioniPresenti
		cimitero.RotazioneEsumazioni = *input.RotazioneEsumazioni

		if err := db.Save(&cimitero).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Cimitero aggiornato con successo"})
		return
	}
}
