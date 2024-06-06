package models

import (
	"gorm.io/gorm"
)

type AventiDirittiModel struct {
	gorm.Model
	DefuntoID     uint   `json:"defunto"       gorm:"column:defunto_id;type:int;not null;unique"`
	Nome          string `json:"nome"          gorm:"type:varchar(50)"`
	Cognome       string `json:"cognome"       gorm:"type:varchar(50)"`
	CodiceFiscale string `json:"codiceFiscale" gorm:"type:varchar(16)"` //can be null
	Email         string `json:"email"         gorm:"type:varchar(50)"` //can be null
	Telefono      int    `json:"telefono"      gorm:"type:int"`         //can be null
}

func (AventiDirittiModel) TableName() string {
	return "aventi_diritti"
}
