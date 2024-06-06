package models

import (
	"gorm.io/gorm"
)

type DefuntiModel struct {
	gorm.Model
	InumazioneID      uint                 `json:"inumazione"        gorm:"column:inumazione_id;type:int;not null;unique"`
	Nome              string               `json:"nome"              gorm:"type:varchar(50)"`
	Cognome           string               `json:"cognome"           gorm:"type:varchar(50)"`
	Sesso             string               `json:"sesso"             gorm:"type:varchar(6)"`
	NotaIdentitaria   string               `json:"notaIdentitaria"   gorm:"type:varchar(255)"` //can be null
	DataNascita       string               `json:"dataNascita"       gorm:"type:varchar(50)"`
	DataOraMorte      string               `json:"dataOraMorte"      gorm:"type:varchar(50)"`
	LuogoNascita      string               `json:"luogoNascita"      gorm:"type:varchar(50)"`
	MalattiaInfettiva bool                 `json:"malattiaInfettiva" gorm:"type:bool;default:false"`
	DataOraSepoltura  string               `json:"dataOraSepoltura"  gorm:"type:varchar(50)"`
	Contratto         ContrattiModel       `                         gorm:"foreignKey:DefuntoID;constraint:OnDelete:CASCADE;"`
	AventiDiritti     []AventiDirittiModel `                         gorm:"foreignKey:DefuntoID;constraint:OnDelete:CASCADE;"`
}

func (DefuntiModel) TableName() string {
	return "defunti"
}

var sessoList = []string{"M", "F", "Altro"}
