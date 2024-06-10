package controllers

import (
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

func CreateContratti(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input handlers.ContrattiRequest
		if err := c.ShouldBind(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := handlers.ValidateContrattiRequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var contratto models.ContrattiModel
		contratto = models.ContrattiModel{
      DefuntoID: *input.DefuntoID ,	
      InizioContratto: *input.InizioContratto,	
      FineContratto: *input.FineContratto,
      StatoContratto: *input.StatoContratto,	
      TipoContratto: *input.TipoContratto,
		}
		file, err := c.FormFile("file")
		if err == nil {
			baseDir := filepath.Join("media", "contratti", "file")
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
      contratto.File = &filePath
		}
		if err := db.Create(&contratto).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Inumazione creata con successo"})
		return
	}
}

func GetContratti(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var contratti []models.ContrattiModel
		query := db

		if idStr := c.Query("id"); idStr != "" {
			id, err := strconv.Atoi(idStr)
			if err == nil {
        if id <= 0 {
          c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero positivo"})
          return
        }
				query = query.Where("id = ?", id)
			}
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID deve essere un numero"})
      return
		}

		if defunto := c.Query("defunto"); defunto != "" {
			query = query.Where("defunto_id = ?", defunto)
		}
		if inizioContratto := c.Query("inizio"); inizioContratto != "" {
			query = query.Where("inizio_contratto = ?", inizioContratto)
		}
    if finContratto := c.Query("fine"); finContratto != "" {
      query = query.Where("fine_contratto = ?", finContratto)
    }
    if stato := c.Query("statoContratto"); stato != "" {
      query = query.Where("stato_contratto = ?", stato)
    }
    if tipologia := c.Query("tipoContratto"); tipologia != "" {
      query = query.Where("tipo_contratto = ?", tipologia)
    }
		if err := query.Find(&contratti).Error; err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Errore nel recupero dei contratti"},
			)
			return
		}
		c.JSON(http.StatusOK, contratti)
    return
	}
}

func DeleteContratti(db *gorm.DB) gin.HandlerFunc {
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

		var contratto models.ContrattiModel
		if err := db.Where("id = ?", id).First(&contratto).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Contratto non trovata"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		if err := db.Delete(&contratto).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Contratto eliminata con successo"})
	}
}

func UpdateContratti(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

