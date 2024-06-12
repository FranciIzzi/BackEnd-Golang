package handlers

import (
	"errors"
	"root/models"

	"gorm.io/gorm"
)

type CimiteriRequest struct {
	gorm.Model
	Latitudine          *float64  `json:"latitudine"`
	Longitudine         *float64  `json:"longitudine"`
	Regione             *string   `json:"regione"`
	Provincia           *string   `json:"provincia"`
	Comune              *string   `json:"comune"`
	Settori             *[]string `json:"settori"`
	PostiTotali         *int      `json:"postiTotali"`
	InumazioniPresenti  *int      `json:"inumazioniPresenti"`
	RotazioneEsumazioni *int      `json:"rotazioneEsumazioniAnni"`
}

func ValidateCimiteriRequest(db *gorm.DB, req *CimiteriRequest) error {
	if req.Latitudine == nil {
		return errors.New("Latitudine deve essere obbligatoria")
	}
	if req.Longitudine == nil {
		return errors.New("Longitudine deve essere obbligatoria")
	}
	if req.Regione == nil {
		return errors.New("Regione obbligatoria")
	}
	if req.Provincia == nil {
		return errors.New("Provincia obbligatoria")
	}
	if req.Comune == nil {
		return errors.New("Comune obbligatorio")
	}
	result, err := validateLocation(
		*req.Regione,
		*req.Provincia,
		*req.Comune,
	)
	if err != nil && result == false {
		return err
	}
  if req.Settori == nil {
    return errors.New("La lista settori deve essere obbligatorio")
  }
  if req.Settori != nil && len(*req.Settori) == 0 {
    return errors.New("Il campo settori deve essere composto da almeno un elemento")
  }
  // TODO: validate settori in maniera rigorosa 
	if req.InumazioniPresenti != nil && *req.InumazioniPresenti < 0 {
		return errors.New("Le inumazioni presenti non possono essere negative")
	}
	if req.PostiTotali != nil && *req.PostiTotali < 0 {
		return errors.New("I posti totali non possono essere negativi")
	}
	if req.RotazioneEsumazioni != nil && *req.RotazioneEsumazioni < 0 {
		return errors.New("La rotazione delle esumazioni non può essere negativa")
	}
	var count int64
	db.Model(&models.CimiteriModel{}).
		Where("latitudine = ? AND longitudine = ?", *req.Latitudine, *req.Longitudine).
		Count(&count)
	if count > 0 {
		return errors.New("la combinazione di latitudine e longitudine esiste già")
	}
	return nil
}

func validateLocation(
	regione string,
	provincia string,
	comune string,
) (bool, error) {
	if province, rErr := regioni_province_comuni[regione]; rErr {
		if comuni, pErr := province[provincia]; pErr {
			for _, c := range comuni {
				if c == comune {
					return true, nil
				}
			}
			return false, errors.New("Il comune non esiste nella provincia inserita")
		}
		return false, errors.New("La provincia non esiste nella regione inserita")
	}
	return false, errors.New("Non è stata inserita una regione italiana")
}

var regioni_province_comuni = map[string]map[string][]string{
	"Abruzzo": {
		"L'Aquila": {"L'Aquila", "Avezzano", "Sulmona"},
		"Teramo":   {"Teramo", "Giulianova", "Roseto degli Abruzzi"},
		"Pescara":  {"Pescara", "Montesilvano", "Spoltore"},
		"Chieti":   {"Chieti", "Vasto", "Lanciano","I148 - San Salvo"},
	},
	"Molise": {
		"Campobasso": {"Campobasso", "Termoli", "Bojano"},
		"Isernia":    {"Isernia", "Venafro", "Agnone"},
	},
}
