package handlers

import (
	"errors"
	"fmt"
	"regexp"
	"root/models"
	"strconv"

	"gorm.io/gorm"
)

type AventiDirittiRequest struct {
	gorm.Model
	DefuntoID     *uint   `json:"defunto"`
	Nome          *string `json:"nome"`
	Cognome       *string `json:"cognome"`
	CodiceFiscale *string `json:"codiceFiscale"`
	Email         *string `json:"email"`
	Telefono      *int    `json:"telefono"`
}

func ValidateAventiDirittirequest(db *gorm.DB, req *[]AventiDirittiRequest) error {
	for index, instance := range *req {
		if instance.DefuntoID == nil {
			return fmt.Errorf(
				"Errore nell'Avente Diritto %d: DefuntoID deve essere obbligatorio",
				index+1,
			)
		}
		if *instance.DefuntoID < 1 {
			return fmt.Errorf("Errore nell'Avente Diritto %d: DefuntoID non valido", index+1)
		}
		var defunto models.DefuntiModel
		var err error
		if err = db.Where("id = ?", instance.DefuntoID).First(&defunto).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return fmt.Errorf("Errore nell'Avente Diritto %d: Defunto non trovato", index+1)
			}
			return fmt.Errorf("Errore Interno al Server nell'Avente Diritto %d", index+1)
		}
		if instance.Nome == nil {
			return fmt.Errorf(
				"Errore nell'Avente Diritto %d: Nome deve essere obbligatorio",
				index+1,
			)
		}
		if instance.Cognome == nil {
			return fmt.Errorf(
				"Errore nell'Avente Diritto %d: Cognome deve essere obbligatorio",
				index+1,
			)
		}

		if instance.CodiceFiscale != nil {
			if err = validateCodiceFiscale(*instance.CodiceFiscale, &index); err != nil {
				return err
			}
		}
		if instance.Email != nil {
			if err = validateEmail(*instance.Email, &index); err != nil {
				return err
			}
		}
		if instance.Telefono != nil {
			if err = validateTelefono(*instance.Telefono, &index); err != nil {
				return err
			}
		}
	}
	return nil
}

func validateCodiceFiscale(cf string, index *int) error {
	if index != nil && *index >= 0 {
		if len(cf) != 16 {
			return fmt.Errorf(
				"Errore nell'Avente Diritto %d: Codice Fiscale non ha 16 caratteri", *index+1)
		}
		match, _ := regexp.MatchString(`^[A-Z]{6}[0-9]{2}[A-Z][0-9]{2}[A-Z][0-9]{3}[A-Z]$`, cf)
		if !match {
			return fmt.Errorf("Errore nell'Avente Diritto %d: Codice Fiscale non valido", *index+1)
		}
	} else {
		if len(cf) != 16 {
			return errors.New("Codice Fiscale non ha 16 caratteri")
		}
		match, _ := regexp.MatchString(`^[A-Z]{6}[0-9]{2}[A-Z][0-9]{2}[A-Z][0-9]{3}[A-Z]$`, cf)
		if !match {
			return errors.New("Codice Fiscale non valido")
		}
	}
	return nil
}

func validateEmail(email string, index *int) error {
	if index != nil && *index >= 0 {
		match, _ := regexp.MatchString(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`, email)
		if !match {
			return fmt.Errorf("Errore nell'Avente Diritto %d: Email non valida", *index+1)
		}
	} else {
		match, _ := regexp.MatchString(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,4}$`, email)
		if !match {
			return errors.New("Email non valida")
		}
	}
	return nil
}

func validateTelefono(telefono int, index *int) error {
	if index != nil && *index >= 0 {
		if len(strconv.Itoa(telefono)) != 10 {
			return fmt.Errorf("Errore nell'Avente Diritto %d: Telefono non valido", *index+1)
		}
	} else {
		if len(strconv.Itoa(telefono)) != 10 {
			return errors.New("Telefono non valido")
		}
	}
	return nil
}
