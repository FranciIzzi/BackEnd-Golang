package handlers

import (
	"errors"
	"root/models"
	"root/validators"

	"gorm.io/gorm"
)

type ContrattiRequest struct {
	gorm.Model
	DefuntoID       *uint   `json:"defunto"`
	InizioContratto *string `json:"inizio"`
	FineContratto   *string `json:"fine"`
	StatoContratto *string `json:"statoContratto"`
	TipoContratto  *string `json:"tipoContratto"`
}

func ValidateContrattiRequest(db *gorm.DB, req *ContrattiRequest) error {
	if req.DefuntoID == nil {
		return errors.New("DefuntoID deve essere obbligatorio")
	}
	if *req.DefuntoID < 1 {
		return errors.New("DefuntoID non valido")
	}
	var defunto models.DefuntiModel
	var err error
	if err = db.Where("id = ?", req.DefuntoID).First(&defunto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Defunto non trovato")
		}
		return errors.New("Errore Interno al Server")
	}
	if req.InizioContratto == nil {
		return errors.New("Inizio del contratto deve essere obbligatorio")
	}
	validateData1, err1 := validators.ValidateDate(*req.InizioContratto)
	if !validateData1 {
		return err1
	}
	if req.FineContratto == nil {
		return errors.New("Fine del contratto deve essere obbligatorio")
	}
	validateData2, err2 := validators.ValidateDate(*req.FineContratto)
	if !validateData2 {
		return err2
	}
	if req.StatoContratto == nil {
		return errors.New("Stato del contratto deve essere obbligatorio")
	}
	if req.TipoContratto == nil {
		return errors.New("Tipo del contratto deve essere obbligatorio")
	}
	if !validateStatoContratto(*req.StatoContratto) {
		return errors.New("Stato del contratto non valido")
	}
	if !validateTipoContratto(*req.TipoContratto) {
		return errors.New("Tipo del contratto non valido")
	}
	return nil
}

func validateStatoContratto(stato string) bool {
	for _, v := range statoContratto {
		if v == stato {
			return true
		}
	}
	return false
}

func validateTipoContratto(tipo string) bool {
	for _, v := range tipoContratto {
		if v == tipo {
			return true
		}
	}
	return false
}

var statoContratto = []string{
	"Regolare",
	"In scadenza",
	"Scaduto",
}
var tipoContratto = []string{"Type1", "Type2", "Type3", "Type4"}
