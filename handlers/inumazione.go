package handlers

import (
	"errors"
	"root/models"
	// "root/models"

	"gorm.io/gorm"
)

type InumazioniRequest struct {
  gorm.Model
	SettoreID       *uint   `json:"settore"`
	CoordinataX     *int    `json:"x"`
	CoordinataY     *int    `json:"y"`
	NumeroCippo     *int    `json:"numeroCippo"`
	ParcelNumber    *int    `json:"parcelNumber"`
	StatoInumazione *string `json:"statoInumazione"`
	Occupato        *bool   `json:"occupied"`
	Tipologia       *string `json:"tipologia"`
}

func ValidateInumazioniRequest(db *gorm.DB, req *InumazioniRequest) error {

	if req.SettoreID == nil {
		return errors.New("SettoreID deve essere obbligatorio")
	}
	if *req.SettoreID < 1 {
		return errors.New("CimiteroID non valido")
	}
	var settore models.SettoriModel
  var err error
	if err = db.Where("id = ?", req.SettoreID).First(&settore).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Settore non trovato")
		}
	   return errors.New("Errore Interno al Server: " + err.Error())
	}
	if req.CoordinataX == nil {
		return errors.New("X deve essere obbligatorio")
	}
	if *req.CoordinataX < 0 {
		return errors.New("X non può essere un intero negativo")
	}
	if req.CoordinataY == nil {
		return errors.New("Y deve essere obbligatorio")
	}
	if *req.CoordinataY < 0 {
		return errors.New("Y non può essere un intero negativo")
	}
	if req.Occupato == nil {
		return errors.New("Occupato deve essere obbligatorio")
	}
	if req.Tipologia == nil {
		return errors.New("Tipologia deve essere obbligatorio")
	}
	if req.Tipologia != nil && !checkValue(tipologieInumazione, *req.Tipologia) {
		return errors.New("Tipologia inserita non valida")
	}
	if req.StatoInumazione == nil {
		return errors.New("StatoInumazione deve essere obbligatorio")
	}
	if req.StatoInumazione != nil && !checkValue(statoInumazione, *req.StatoInumazione) {
		return errors.New("Stato Inumazione inserito non valida")
	}

	return nil
}


func checkValue(list []string, value string) bool {
	for _, x := range list {
		if x == value {
			return true
		}
	}
	return false
}

var tipologieInumazione = []string{"Tomba", "Loculo", "Ossario", "Loculo in Cappella"}
var statoInumazione = []string{"Necessita manutenzione", "In assestamento", "Buono"}
