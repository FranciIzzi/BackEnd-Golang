package controllers

import (
	"bytes"
	"log"
	"net/http"
	"root/handlers"
	"root/models"
	"strconv"

	"io"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDefunto(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyBytes, err := io.ReadAll(io.Reader(c.Request.Body))
		if err != nil {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "Errore nella lettura della richiesta: " + err.Error()},
			)
			return
		}
		bodyString := string(bodyBytes)
		log.Println("Corpo della richiesta:", bodyString)

		c.Request.Body = io.NopCloser(io.Reader(bytes.NewBuffer(bodyBytes)))

		var input handlers.DefuntiRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := handlers.ValidateDefuntiRequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var defunto models.DefuntiModel
		defunto = models.DefuntiModel{
			InumazioneID:      *input.InumazioneID,
			Nome:              *input.Nome,
			Cognome:           *input.Cognome,
			Sesso:             *input.Sesso,
			NotaIdentitaria:   input.NotaIdentitaria,
			DataNascita:       *input.DataNascita,
			DataOraMorte:      *input.DataOraMorte,
			LuogoNascita:      *input.LuogoNascita,
			MalattiaInfettiva: *input.MalattiaInfettiva,
			DataOraSepoltura:  *input.DataOraSepoltura,
		}
		if err := db.Create(&defunto).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Defunto creato con successo", "id": defunto.ID})
		return
	}

}

func GetDefunti(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var defunti []models.DefuntiModel
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

		if inumazione := c.Query("inumazione"); inumazione != "" {
			inumazioneID, err := strconv.Atoi(inumazione)
			if err == nil {
				if inumazioneID <= 0 {
					c.JSON(
						http.StatusBadRequest,
						gin.H{"error": "InumazioneID deve essere un numero positivo"},
					)
					return
				}
				query = query.Where("inumazione_id = ?", inumazioneID)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "InumazioneID deve essere un numero"})
				return
			}
		}
		if nome := c.Query("nome"); nome != "" {
			query = query.Where("nome = ?", nome)
		}
		if cognome := c.Query("cognome"); cognome != "" {
			query = query.Where("cognome = ?", cognome)
		}
		if sesso := c.Query("sesso"); sesso != "" {
			query = query.Where("sesso = ?", sesso)
		}
		if dataOraMorte := c.Query("dataOraMorte"); dataOraMorte != "" {
			query = query.Where("data_ora_morte = ?", dataOraMorte)
		}
		if luogoNascita := c.Query("luogoNascita"); luogoNascita != "" {
			query = query.Where("luogo_nascita = ?", luogoNascita)
		}
		if malattiaInfettiva := c.Query("malattiaInfettiva"); malattiaInfettiva != "" {
			query = query.Where("malattia_infettiva = ?", malattiaInfettiva)
		}
		if dataNascita := c.Query("dataNascita"); dataNascita != "" {
			query = query.Where("data_nascita = ?", dataNascita)
		}
		if dataOraSepoltura := c.Query("dataOraSepoltura"); dataOraSepoltura != "" {
			query = query.Where("data_ora_sepoltura = ?", dataOraSepoltura)
		}
		if err := query.Find(&defunti).Error; err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{"error": "Errore nel recupero dei defunti"},
			)
			return
		}
		c.JSON(http.StatusOK, defunti)
		return
	}
}

func DeleteDefunto(db *gorm.DB) gin.HandlerFunc {
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

		var defunto models.DefuntiModel
		if err := db.Where("id = ?", id).First(&defunto).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Defunto non trovato"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}
		if err := db.Delete(&defunto).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Defunto eliminato con successo"})
	}
}

func UpdateDefunto(db *gorm.DB) gin.HandlerFunc {
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
		var input handlers.DefuntiRequest
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := handlers.ValidateDefuntiRequest(db, &input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var defunto models.DefuntiModel

		if err := db.Where("id = ?", id).First(&defunto).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Defunto non trovato"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			return
		}

		if input.ID != defunto.ID {
			c.JSON(
				http.StatusBadRequest,
				gin.H{"error": "ID del defunto non corrispondente"},
			)
			return
		}
		defunto.InumazioneID = *input.InumazioneID
		defunto.Nome = *input.Nome
		defunto.Cognome = *input.Cognome
		defunto.Sesso = *input.Sesso
		defunto.NotaIdentitaria = input.NotaIdentitaria
		defunto.DataNascita = *input.DataNascita
		defunto.DataOraMorte = *input.DataOraMorte
		defunto.DataOraSepoltura = *input.DataOraSepoltura
		defunto.LuogoNascita = *input.LuogoNascita
		defunto.MalattiaInfettiva = *input.MalattiaInfettiva

		if err := db.Save(&defunto).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Defunto aggiornato con successo"})
		return
	}
}
