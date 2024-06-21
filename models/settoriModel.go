package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type SettoriModel struct {
	gorm.Model
	CimiteroID        uint              `json:"cimitero"          gorm:"column:cimitero_id;type:int;not null"`
	Nome              string            `json:"nome"              gorm:"column:nome;type:varchar(50);not null"`
	Righe             int               `json:"righe"             gorm:"column:righe;type:int;not null"`
	Colonne           int               `json:"colonne"           gorm:"column:colonne;type:int;not null"`
	InumazioniSettore int               `json:"inumazioniSettore" gorm:"column:inumazioni_settore;type:int;not null"`
	Snapshot          string            `json:"snapshot"          gorm:"column:snapshot;type:varchar(255);not null"`
	PuntiCanvas       string            `json:"puntiCanvas"       gorm:"column:punti_canvas;type:varchar(255);not null"`
	Color             string            `json:"color"             gorm:"column:color;type:varchar(50);not null"`
	Inumazioni        []InumazioniModel `json:"inumazioni"        gorm:"foreignKey:SettoreID;constraint:OnDelete:CASCADE;"`
}

type Canvas struct {
	X int `json:"X"`
	Y int `json:"Y"`
}

func (SettoriModel) TableName() string {
	return "settori"
}

type JSONMap map[string]Canvas

func (m JSONMap) Value() (driver.Value, error) {
	if m == nil {
		return nil, nil
	}
	ba, err := m.MarshalJSON()
	return string(ba), err
}

func (m *JSONMap) Scan(val interface{}) error {
	var ba []byte
	switch v := val.(type) {
	case []byte:
		ba = v
	case string:
		ba = []byte(v)
	default:
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", val))
	}
	t := map[string]Canvas{}
	err := json.Unmarshal(ba, &t)
	*m = JSONMap(t)
	return err
}

func (m JSONMap) MarshalJSON() ([]byte, error) {
	if m == nil {
		return []byte("null"), nil
	}
	t := (map[string]Canvas)(m)
	return json.Marshal(t)
}

func (m *JSONMap) UnmarshalJSON(b []byte) error {
	t := map[string]Canvas{}
	err := json.Unmarshal(b, &t)
	*m = JSONMap(t)
	return err
}

func (m JSONMap) GormDataType() string {
	return "jsonmap"
}

func (JSONMap) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	switch db.Dialector.Name() {
	case "sqlite":
		return "JSON"
	case "mysql":
		return "JSON"
	case "postgres":
		return "JSONB"
	}
	return ""
}
