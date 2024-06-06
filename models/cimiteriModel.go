package models

import (
	"gorm.io/gorm"
)

type CimiteriModel struct {
	gorm.Model
	Latitudine          float64           `json:"latitudine"              gorm:"column:latitudine;type:float;not null;uniqueIndex:idx_lat_long"`
	Longitudine         float64           `json:"longitudine"             gorm:"column:longitudine;type:float;not null;uniqueIndex:idx_lat_long"`
	Regione             string            `json:"regione"                 gorm:"column:regione;type:varchar(50);not null"`
	Provincia           string            `json:"provincia"               gorm:"column:provincia;type:varchar(50);not null"`
	Comune              string            `json:"comune"                  gorm:"column:comune;type:varchar(80);not null"`
	Settori             []string          `json:"settori"                 gorm:"column:settori;type:jsonb;not null"`
	PostiTotali         int               `json:"postiTotali"             gorm:"column:posti_totali;type:int;not null"`
	InumazioniPresenti  int               `json:"inumazioniPresenti"      gorm:"column:inumazioni_presenti;type:int;not null"`
	RotazioneEsumazioni int               `json:"rotazioneEsumazioniAnni" gorm:"column:rotazione_esumazioni;type:int;not null"`
	Inumazioni          []InumazioniModel `                               gorm:"foreignKey:CimiteroID;constraint:OnDelete:CASCADE;"`
}

func (CimiteriModel) TableName() string {
	return "cimiteri"
}
