package models

import (
	"gorm.io/gorm"
)

type AventiDirittiModel struct {
	gorm.Model
	DefuntoID     uint    `json:"defunto"       gorm:"column:defunto_id;type:int;not null"`
	Nome          string  `json:"nome"          gorm:"column:nome;type:varchar(50);not null"`
	Cognome       string  `json:"cognome"       gorm:"column:cognome;type:varchar(50);not null"`
	CodiceFiscale *string `json:"codiceFiscale" gorm:"column:codice_fiscale;type:varchar(16)"`
	Email         *string `json:"email"         gorm:"column:email;type:varchar(50)"`
	Telefono      *int    `json:"telefono"      gorm:"column:telefono;type:int"`
}

func (AventiDirittiModel) TableName() string {
	return "aventi_diritti"
}
