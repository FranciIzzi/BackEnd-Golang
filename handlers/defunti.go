package handlers

import (
	"errors"
	"log"
	"root/models"
	"root/validators"

	"gorm.io/gorm"
)

type DefuntiRequest struct {
	gorm.Model
	InumazioneID      *uint   `json:"inumazione"        `
	Nome              *string `json:"nome"              `
	Cognome           *string `json:"cognome"           `
	Sesso             *string `json:"sesso"             `
	NotaIdentitaria   *string `json:"notaIdentitaria"   `
	DataNascita       *string `json:"dataNascita"       `
	DataOraMorte      *string `json:"dataOraMorte"      `
	LuogoNascita      *string `json:"luogoNascita"      `
	MalattiaInfettiva *bool   `json:"malattiaInfettiva" `
	DataOraSepoltura  *string `json:"dataOraSepoltura"  `
}

func ValidateDefuntiRequest(db *gorm.DB, req *DefuntiRequest) error {
	if req.InumazioneID == nil {
		return errors.New("L'inumazione deve essere specificata")
	}
	if *req.InumazioneID < 1 {
		return errors.New("Inumazione non valida, non può essere negativa")
	}
	var inumazione models.InumazioniModel
	var err error
	if err = db.Where("id = ?", req.InumazioneID).First(&inumazione).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Inumazione non trovata nel database")
		}
		return errors.New("Errore Interno al Server")
	}

	if req.Nome == nil {
		return errors.New("Nome deve essere obbligatorio")
	}
	if req.Cognome == nil {
		return errors.New("Cognome deve essere obbligatorio")
	}
	if req.Sesso == nil {
		return errors.New("Sesso deve essere obbligatorio")
	}
	if req.Sesso != nil && !validateSesso(*req.Sesso) {
		return errors.New("Sesso non valido")
	}
	if req.LuogoNascita == nil {
		return errors.New("Luogo di nascita deve essere obbligatorio")
	}
	if req.MalattiaInfettiva == nil {
		req.MalattiaInfettiva = new(bool)
		*req.MalattiaInfettiva = false
	}
	if req.DataNascita == nil {
		return errors.New("Data di nascita deve essere obbligatorio")
	}
	validateData, err := validators.ValidateDate(*req.DataNascita)
	if !validateData {
		return err
	}
	*req.DataNascita = validators.ConvertStringToDate(req.DataNascita)
	if req.DataOraMorte == nil {
		return errors.New("Data di morte deve essere obbligatorio")
	}
	log.Print("ecco la data :" + *req.DataOraMorte)
	validateData1, err1 := validators.ValidateDateTime(*req.DataOraMorte)
	if !validateData1 {
		return err1
	}
	*req.DataOraMorte = validators.ConvertStringToDateTime(req.DataOraMorte)
	if req.DataOraSepoltura == nil {
		return errors.New("Data di sepoltura deve essere obbligatorio")
	}
	validateData2, err2 := validators.ValidateDateTime(*req.DataOraSepoltura)
	if !validateData2 {
		return err2
	}
	*req.DataOraSepoltura = validators.ConvertStringToDateTime(req.DataOraSepoltura)
	return nil
}

var sessoList = []string{"M", "F", "Altro"}

func validateSesso(value string) bool {
	for _, x := range sessoList {
		if x == value {
			return true
		}
	}
	return false
}
