package controllers

import (
	"net/http"
	"root/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func GetRicerca(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var listDefunti []models.DefuntiModel

		query := db.Model(&models.DefuntiModel{}).
			Preload("Contratti").
			Preload("AventiDiritti").
			Joins("join inumazioni on defunti.inumazione_id = inumazioni.id")

		for key, fields := range querySettings {
			for _, field := range fields {
				if value := c.Query(field); value != "" {
					// TODO: aggiungere qui dei filtri particolari facendo un check su value
					query = query.Where(key+"."+field+" = ?", value)
				}
			}
		}

		if err := query.Find(&listDefunti).Error; err != nil {
			c.JSON(500, gin.H{"error": "Errore nel recuperare i dati"})
			return
		}

		c.JSON(http.StatusOK, listDefunti)
	}
}

var querySettings = map[string][]string{
	"defunti": {
		"nome",
		"cognome",
		"sesso",
		"data_nascita",
		"data_ora_morte",
		"luogo_nascita",
		"malattia_infettiva",
		"data_ora_sepoltura",
	},
	"inumazioni": {
		"cimitero_id",
		"settore",
		"coordinata_x",
		"coordinata_y",
		"numero_cippo",
		"parcel_number",
		"stato_inumazione",
		"occupato",
		"tipologia",
	},
	"contratti": {
		"tipo_contratto",
		"inizio_contratto",
		"fine_contratto",
		"stato_contratto",
	},
	"aventi_diritti": {
		"nome",
		"cognome",
		"codice_fiscale",
		"email",
		"telefono",
	},
}
