package controllers

import (
	"net/http"
	"root/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetRicerca(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var listInumazioni []models.InumazioniModel

		query := db

		// Filtri Inumazioni
		if settore := c.Query("7"); settore != "" {
			query = query.Where("settore = ?", "%"+settore+"%")
		}
		if tipologia := c.Query("24"); tipologia != "" {
			query = query.Where("tipologia = ?", tipologia)
		}
		if occupied := c.Query("12"); occupied != "" {
			if occupiedBool, err := strconv.ParseBool(occupied); err == nil {
				query = query.Where("occupato = ?", occupiedBool)
			}
		}
		if numeroCippo := c.Query("10"); numeroCippo != "" {
			if number, err := strconv.Atoi(numeroCippo); err == nil {
				query = query.Where("numero_cippo = ?", number)
			}
		}
		if parcelNumber := c.Query("11"); parcelNumber != "" {
			if number, err := strconv.Atoi(parcelNumber); err == nil {
				query = query.Where("parcel_number = ?", number)
			}
		}
		if x := c.Query("8"); x != "" {
			if number, err := strconv.Atoi(x); err == nil {
				query = query.Where("coordinata_x = ?", number)
			}
		}
		if y := c.Query("9"); y != "" {
			if number, err := strconv.Atoi(y); err == nil {
				query = query.Where("coordinata_y = ?", number)
			}
		}
		if statoInumazione := c.Query("13"); statoInumazione != "" {
			query = query.Where("stato_inumazione = ?", statoInumazione)
		}
		if cimitero := c.Query("6"); cimitero != "" {
			if number, err := strconv.Atoi(cimitero); err == nil {
				query = query.Where("cimitero_id = ?", number)
			}
		}

		query1 := query.Joins("JOIN defunti ON defunti.inumazione_id = inumazioni.id")

		// Filtri Defunti
		if nome := c.Query("0"); nome != "" {
			likeValue := "%" + nome + "%"
			query1 = query1.Where("defunti.nome LIKE ?", likeValue)
		}
		if cognome := c.Query("1"); cognome != "" {
			likeValue := "%" + cognome + "%"
			query1 = query1.Where("defunti.cognome LIKE ?", likeValue)
		}
		if sesso := c.Query("4"); sesso != "" {
			likeValue := "%" + sesso + "%"
			query1 = query1.Where("defunti.sesso LIKE ?", likeValue)
		}
		if data_nascita := c.Query("2"); data_nascita != "" {
			likeValue := "%" + data_nascita + "%"
			query1 = query1.Where("defunti.data_nascita LIKE ?", likeValue)
		}
		if luogo_nascita := c.Query("3"); luogo_nascita != "" {
			likeValue := "%" + luogo_nascita + "%"
			query1 = query1.Where("defunti.luogo_nascita LIKE ?", likeValue)
		}
		if malattia_infettiva := c.Query("5"); malattia_infettiva != "" {
			if bool, err := strconv.ParseBool(malattia_infettiva); err == nil {
				query1 = query1.Where("defunti.malattia_infettiva = ?", bool)
			}
		}
		if data_ora_sepoltura := c.Query("14"); data_ora_sepoltura != "" {
			likeValue := "%" + data_ora_sepoltura + "%"
			query1 = query1.Where("defunti.data_ora_sepoltura LIKE ?", likeValue)
		}
		if data_ora_morte := c.Query("22"); data_ora_morte != "" {
			likeValue := "%" + data_ora_morte + "%"
			query1 = query1.Where("defunti.data_ora_morte LIKE ?", likeValue)
		}

		query2 := query1.Joins("JOIN contratti ON contratti.defunto_id = defunti.id")

		// Filtri Contratti
		if tipo_contratto := c.Query("16"); tipo_contratto != "" {
			query2 = query2.Where("contratti.tipo_contratto = ?", tipo_contratto)
		}
		if stato_contratto := c.Query("19"); stato_contratto != "" {
			query2 = query2.Where("contratti.stato_contratto = ?", stato_contratto)
		}
		if inizio_contratto := c.Query("17"); inizio_contratto != "" {
			query2 = query2.Where("contratti.inizio_contratto LIKE ?", "%"+inizio_contratto+"%")
		}
		if fine_contratto := c.Query("18"); fine_contratto != "" {
			query2 = query2.Where("contratti.fine_contratto LIKE ?", "%"+fine_contratto+"%")
		}

		query3 := query2.Joins("JOIN aventi_diritti ON aventi_diritti.defunto_id = defunti.id")

		// Filtri Aventi Diritti
		// if nome := c.Query("25"); nome != "" {
		// 	query = query.Joins("JOIN defunti ON defunti.inumazione_id = inumazioni.id").
		// 		Joins("JOIN aventi_diritti ON aventi_diritti.defunto_id = defunti.id").
		// 		Where("aventi_diritti.nome LIKE ?", "%"+nome+"%")
		// }
		// if cognome := c.Query("26"); cognome != "" {
		// 	query = query.Joins("JOIN defunti ON defunti.inumazione_id = inumazioni.id").
		// 		Joins("JOIN aventi_diritti ON aventi_diritti.defunto_id = defunti.id").
		// 		Where("aventi_diritti.cognome LIKE ?", "%"+cognome+"%")
		// }
		if email := c.Query("28"); email != "" {
			query3 = query3.
				Where("aventi_diritti.email LIKE ?", "%"+email+"%")
		}
		if cf := c.Query("27"); cf != "" {
			query3 = query3.Where("aventi_diritti.codice_fiscale LIKE ?", "%"+cf+"%")
		}
		if telefono := c.Query("29"); telefono != "" {
			if number, err := strconv.Atoi(telefono); err == nil {
				query3 = query3.Where("aventi_diritti.telefono LIKE ?", number)
			}
		}

		if err := query.Preload("Defunto").
			Preload("Defunto.Contratto").
			Preload("Defunto.AventiDiritti").
			Find(&listInumazioni).Error; err != nil {
			c.JSON(500, gin.H{"errore": "Errore nel recuperare i dati, riprovare"})
			return
		}
		uniqueInumazioni := removeDuplicates(listInumazioni)
		c.JSON(http.StatusOK, uniqueInumazioni)
		return
	}
}

func removeDuplicates(inumazioni []models.InumazioniModel) []models.InumazioniModel {
	seen := make(map[uint]models.InumazioniModel)
	for _, inumazione := range inumazioni {
		seen[inumazione.ID] = inumazione
	}

	uniqueInumazioni := make([]models.InumazioniModel, 0, len(seen))
	for _, inumazione := range seen {
		uniqueInumazioni = append(uniqueInumazioni, inumazione)
	}
	return uniqueInumazioni
}
