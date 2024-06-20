package models

import (
	"gorm.io/gorm"
)

type InumazioniModel struct {
	gorm.Model
	SettoreID       uint         `json:"settore"         gorm:"column:settore_id;type:int;not null"`
	CoordinataX     int          `json:"x"               gorm:"column:coordinata_x;type:int;not null;"`
	CoordinataY     int          `json:"y"               gorm:"column:coordinata_y;type:int;not null;"`
	NumeroCippo     int          `json:"numeroCippo"     gorm:"column:numero_cippo;type:int;not null;"`
	ParcelNumber    int          `json:"parcelNumber"    gorm:"column:parcel_number;type:int;not null;"`
	StatoInumazione string       `json:"statoInumazione" gorm:"column:stato_inumazione;type:varchar(35);not null"`
	Occupato        bool         `json:"occupied"        gorm:"column:occupato;type:bool;not null;default:false"`
	Tipologia       string       `json:"tipologia"       gorm:"column:tipologia;type:varchar(25);not null;"`
	Foto            *string      `json:"foto"            gorm:"column:foto;type:varchar(255);"`
	Defunto         DefuntiModel `                       gorm:"foreignKey:InumazioneID;constraint:OnDelete:CASCADE;"`
}

func (InumazioniModel) TableName() string {
	return "inumazioni"
}
