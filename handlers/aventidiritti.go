package handlers

import (
	"errors"
	"regexp"
	"root/models"
	"strconv"

	"gorm.io/gorm"
)

type AventiDirittiRequest struct {
	gorm.Model
	DefuntoID     *uint   `json:"defunto"       `
	Nome          *string `json:"nome"          `
	Cognome       *string `json:"cognome"       `
	CodiceFiscale *string `json:"codiceFiscale" `
	Email         *string `json:"email"         `
	Telefono      *int    `json:"telefono"      `
}

func ValidateAventiDirittirequest(db *gorm.DB, req *AventiDirittiRequest) error {
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
	if req.Nome == nil {
		return errors.New("Nome deve essere obbligatorio")
	}
	if req.Cognome == nil {
		return errors.New("Cognome deve essere obbligatorio")
	}

	if req.CodiceFiscale != nil {
		if err = validateCodiceFiscale(*req.CodiceFiscale); err != nil {
			return err
		}
	}
	if req.Email != nil {
		if err = validateEmail(*req.Email); err != nil {
			return err
		}
	}
	if req.Telefono != nil {
		if err = validateTelefono(*req.Telefono); err != nil {
			return err
		}
	}
	return nil
}

func validateCodiceFiscale(cf string) error {
	if len(cf) != 16 {
		return errors.New("Codice Fiscale non ha 16 caratteri")
	}
	match, _ := regexp.MatchString(`^[A-Z]{6}[0-9]{2}[A-Z][0-9]{2}[A-Z][0-9]{3}[A-Z]$`, cf)
	if !match {
		return errors.New("Codice Fiscale non valido")
	}
	return nil
}

func validateEmail(email string) error {
	match, _ := regexp.MatchString(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`, email)
	if !match {
		return errors.New("Email non valida")
	}
	return nil
}

func validateTelefono(telefono int) error {
	if len(strconv.Itoa(telefono)) != 10 {
		return errors.New("Telefono non valido")
	}
	return nil
}
