package handlers

import (
	"errors"
	"root/models"

	"gorm.io/gorm"
)

type SettoriRequest struct {
	gorm.Model
	CimiteroID        *uint           `json:"cimitero"          gorm:"column:cimitero_id;type:int;not null"`
	Nome              *string         `json:"nome"              gorm:"column:nome;type:varchar(50);not null"`
	Righe             *int            `json:"righe"             gorm:"column:righe;type:int;not null"`
	Colonne           *int            `json:"colonne"           gorm:"column:colonne;type:int;not null"`
	InumazioniSettore *int            `json:"inumazioniSettore" gorm:"column:inumazioni_settore;type:int;not null"`
	Snapshot          *string         `json:"snapshot"          gorm:"column:snapshot;type:varchar(255);not null"` //image
	PuntiCanvas       *string         `json:"puntiCanvas"       gorm:"column:punti_canvas;type:jsonb;not null"`
	Color             *string         `json:"color"             gorm:"column:color;type:varchar(50);not null"`
}

func ValidateSettoriRequest(db *gorm.DB, req *SettoriRequest) error {
	if req.CimiteroID == nil {
		return errors.New("Il cimitero deve essere obbligatorio")
	}
	var cimitero models.CimiteriModel
	if err := db.Where("id = ?", req.CimiteroID).First(&cimitero).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("Cimitero non trovato")
		}
		return errors.New("Errore Interno al Server: " + err.Error())
	}
	if req.Righe == nil {
		return errors.New("Le righe del settore devono essere specificate")
	}
	if req.Colonne == nil {
		return errors.New("Le colonne del settore devono essere specificate")
	}
	if req.InumazioniSettore == nil {
		return errors.New("Le inumazioni del settore devono essere specificate")
	}
	if req.PuntiCanvas == nil {
		return errors.New("I punti del canvas del settore devono essere obbligatori")
	}
	if req.Color == nil {
		return errors.New("Errore con il colore del settore")
	}
	return nil
}
