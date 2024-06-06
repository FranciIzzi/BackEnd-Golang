package handlers

import (
	"errors"
	"root/models"

	"gorm.io/gorm"
)

type InumazioniRequest struct {
	gorm.Model
	CimiteroID      *uint   `json:"cimitero"`
	Settore         *string `json:"settore"`
	CoordinataX     *int    `json:"x"`
	CoordinataY     *int    `json:"y"`
	NumeroCippo     *int    `json:"numeroCippo"`
	ParcelNumber    *int    `json:"parcelNumber"`
	StatoInumazione *string `json:"statoInumazione"`
	Occupato        *bool   `json:"occupied"`
	Tipologia       *string `json:"tipologia"`
}

func ValidateInumazioniRequest(db *gorm.DB, req *InumazioniRequest) error {

	if req.CimiteroID == nil {
		return errors.New("CimiteroID deve essere obbligatorio")
	}
	if *req.CimiteroID < 1 {
		return errors.New("CimiteroID non valido")
	}
	var cimitero models.CimiteriModel
	var err error
	if err = db.Where("id = ?", req.CimiteroID).First(&cimitero).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Cimitero non trovato")
		}
		return errors.New("Errore Interno al Server")
	}
	var settoriDetected = cimitero.Settori
	if req.Settore == nil {
		return errors.New("Settore deve essere obbligatorio")
	}
	if req.Settore != nil && !DetectSettore(settoriDetected, *req.Settore) {
		return errors.New("Settore non presente nel cimitero inserito")
	}
	if req.CoordinataX == nil {
		return errors.New("X deve essere obbligatorio")
	}
	if *req.CoordinataX < 1 {
		return errors.New("X deve essere un intero positivo")
	}
	if req.CoordinataY == nil {
		return errors.New("Y deve essere obbligatorio")
	}
	if *req.CoordinataY < 1 {
		return errors.New("Y deve essere un intero positivo")
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

func DetectSettore(settori []string, settore string) bool {
	for _, set := range settori {
		if set == settore {
			return true
		}
	}
	return false
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
