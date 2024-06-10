package models

import (
	"gorm.io/gorm"
)

type DefuntiModel struct {
	gorm.Model
	InumazioneID      uint                 `json:"inumazione"        gorm:"column:inumazione_id;type:int;not null;unique"`
	Nome              string               `json:"nome"              gorm:"column:nome;type:varchar(50);not null"`
	Cognome           string               `json:"cognome"           gorm:"column:cognome;type:varchar(50);not null"`
	Sesso             string               `json:"sesso"             gorm:"column:sesso;type:varchar(6);not null"`
	NotaIdentitaria   *string              `json:"notaIdentitaria"   gorm:"column:nota_identitaria;type:varchar(255)"`
	DataNascita       string               `json:"dataNascita"       gorm:"column:data_nascita;type:varchar(20);not null"`
	DataOraMorte      string               `json:"dataOraMorte"      gorm:"column:data_ora_morte;type:varchar(20);not null"`
	LuogoNascita      string               `json:"luogoNascita"      gorm:"column:luogo_nascita;type:varchar(20);not null"`
	MalattiaInfettiva bool                 `json:"malattiaInfettiva" gorm:"column:malattia_infettiva;type:bool;not null;default:false"`
	DataOraSepoltura  string               `json:"dataOraSepoltura"  gorm:"column:data_ora_sepoltura;type:varchar(50);not null"`
	Contratto         ContrattiModel       `                         gorm:"foreignKey:DefuntoID;constraint:OnDelete:CASCADE;"`
	AventiDiritti     []AventiDirittiModel `                         gorm:"foreignKey:DefuntoID;constraint:OnDelete:CASCADE;"`
}

func (DefuntiModel) TableName() string {
	return "defunti"
}
