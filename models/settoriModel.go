package models

import (
	"gorm.io/gorm"
)

type SettoriModel struct {
	gorm.Model
	CimiteroID        uint              `json:"cimitero"          gorm:"column:cimitero_id;type:int;not null"`
	Nome              string            `json:"nome"              gorm:"column:nome;type:varchar(50);not null"`
	Righe             int               `json:"righe"             gorm:"column:righe;type:int;not null"`
	Colonne           int               `json:"colonne"           gorm:"column:colonne;type:int;not null"`
	InumazioniSettore int               `json:"inumazioniSettore" gorm:"column:inumazioni_settore;type:int;not null"`
	Snapshot          string            `json:"snapshot"          gorm:"column:snapshot;type:varchar(255);not null"` //image
	PuntiCanvas       []Canvas          `json:"puntiCanvas"       gorm:"column:punti_canvas;type:jsonb;not null"`
	Color             string            `json:"color"             gorm:"column:color;type:varchar(50);not null"`
	Inumazioni        []InumazioniModel `json:"inumazioni"        gorm:"foreignKey:SettoreID;constraint:OnDelete:CASCADE;"`
}

type Canvas struct {
	X int
	Y int
}

func (SettoriModel) TableName() string {
	return "settori"
}
