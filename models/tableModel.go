package models

import (
	"gorm.io/gorm"
)

type FieldModel struct {
	gorm.Model
	Fname   string `json:"fname" gorm:"type:varchar(100);not nulli;"`
	Type    string `json:"type"  gorm:"type:varchar(20);not null;"`
	TableID uint   `gorm:"index;not null"`
}

type TableModel struct {
	gorm.Model
	Tname  string       `json:"table" gorm:"type:varchar(50);not null;unique"`
	Fields []FieldModel `gorm:"constraint:OnDelete:CASCADE;"`
}

// 	Email    string `json:"email" gorm:"type:varchar(100);unique"`
// 	Password string `json:"password" gorm:"size:256"`
// 	Token    string `json:"token" gorm:"size:512"`
// 	Logged   bool   `json:"logged" gorm:"default:false"`
// 	Permessi string `json:"permission" gorm:"type:varchar(20)"`

// `foreignKey:<chiave>: Specifica la chiave esterna per una relazione.
// `references:<colonna>: Indica la colonna referenziata in una chiave esterna.
// `many2many:<nome_tabella>: Specifica una relazione molti-a-molti tramite una tabella join.
// `preload: Indica di caricare anticipatamente la relazione in una query.
