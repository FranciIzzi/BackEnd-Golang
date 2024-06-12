package models

import (
	"gorm.io/gorm"
)

type ContrattiModel struct {
	gorm.Model
	DefuntoID       uint   `json:"defunto"        gorm:"column:defunto_id;type:int;not null"`
	InizioContratto string `json:"inizio"         gorm:"column:inizio_contratto;type:varchar(25);not null;"`
	FineContratto   string `json:"fine"           gorm:"column:fine_contratto;type:varchar(25);not null;"`
	File            *string `json:"file"           gorm:"column:file;type:varchar(255);"`
	StatoContratto  string `json:"statoContratto" gorm:"column:stato_contratto;type:varchar(25);not null;"`
	TipoContratto   string `json:"tipoContratto"  gorm:"column:tipo_contratto;type:varchar(25);not null;"`
}

func (ContrattiModel) TableName() string {
	return "contratti"
}
