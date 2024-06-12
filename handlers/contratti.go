package handlers

import (
	"fmt"
	"root/models"
	"root/validators"

	"gorm.io/gorm"
)

type ContrattiRequest struct {
	gorm.Model
	DefuntoID       *uint   `json:"defunto"`
	InizioContratto *string `json:"inizio"`
	FineContratto   *string `json:"fine"`
	StatoContratto  *string `json:"statoContratto"`
	TipoContratto   *string `json:"tipoContratto"`
}

func ValidateContrattiRequest(db *gorm.DB, instance *ContrattiRequest) error {
	if instance.DefuntoID == nil {
		return fmt.Errorf(
			"Errore nel Contratto : DefuntoID deve essere obbligatorio")
	}
	if *instance.DefuntoID < 1 {
		return fmt.Errorf("Errore nel Contratto : DefuntoID non valido")
	}
	var defunto models.DefuntiModel
	var err error
	if err = db.Where("id = ?", instance.DefuntoID).First(&defunto).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("Errore nel Contratto : Defunto non trovato")
		}
		return fmt.Errorf("Errore nel Contratto: Errore Interno al Server")
	}
	if instance.InizioContratto == nil {
		return fmt.Errorf(
			"Errore nel Contratto : Inizio del contratto deve essere obbligatorio")
	}
	validateData1, err1 := validators.ValidateDate(*instance.InizioContratto)
	if !validateData1 {
		return fmt.Errorf("Errore al contratto: %v", err1)
	}
	*instance.InizioContratto = validators.ConvertStringToDate(instance.InizioContratto)
	if instance.FineContratto == nil {
		return fmt.Errorf(
			"Errore nel Contratto : Fine del contratto deve essere obbligatorio")
	}
	validateData2, err2 := validators.ValidateDate(*instance.FineContratto)
	if !validateData2 {
		return fmt.Errorf("Errore al contratto : %v", err2)
	}
	*instance.FineContratto = validators.ConvertStringToDate(instance.FineContratto)
	if instance.StatoContratto == nil {
		return fmt.Errorf(
			"Errore nel Contratto : Stato del contratto deve essere obbligatorio")
	}
	if instance.TipoContratto == nil {
		return fmt.Errorf(
			"Errore nel Contratto : Tipo del contratto deve essere obbligatorio")
	}
	if !validateStatoContratto(*instance.StatoContratto) {
		return fmt.Errorf("Errore nel Contratto : Stato del contratto non valido")
	}
	if !validateTipoContratto(*instance.TipoContratto) {
		return fmt.Errorf("Errore nel Contratto : Tipo del contratto non valido")
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
