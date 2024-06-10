package validators

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

//  Data e Time di riferimento 2024-06-10T14:30
func ValidateDateTime(dateTime string) (bool, error) {
	path := strings.Split(dateTime, "T")
	if len(path) != 2 {
		return false, errors.New("Formato della data non valida")
	}
	date := strings.Split(path[0], "-")
	if len(date) != 3 {
		return false, errors.New("Formato non valido, la data non è valida")
	}
	year, err := strconv.Atoi(date[0])
	if err != nil || len(date[0]) != 4 || year < 1 {
		return false, errors.New("L'anno inserito non è valido")
	}
	month, err := strconv.Atoi(date[1])
	if err != nil || month < 1 || month > 12 {
		return false, errors.New("Il mese inserito non è valido")
	}
	day, err := strconv.Atoi(date[2])
	if err != nil || day < 1 || day > daysInMonth(month, year) {
		return false, errors.New("Il giorno inserito non è valido")
	}
	timeParts := strings.Split(path[1], ":")
	if len(timeParts) != 2 {
		return false, errors.New("formato non valido, l'orario non è valido")
	}
	hour, err := strconv.Atoi(timeParts[0])
	if err != nil || hour < 0 || hour > 23 {
		return false, errors.New("l'ora inserita non è valida")
	}
	minute, err := strconv.Atoi(timeParts[1])
	if err != nil || minute < 0 || minute > 59 {
		return false, errors.New("il minuto inserito non è valido")
	}

	return true, nil
}

// Data di riferimento 2024-06-10
func ValidateDate(date string) (bool, error) {
	path := strings.Split(date, "-")
	if len(date) != 3 {
		return false, errors.New("Formato non valido, la data non è valida")
	}
	year, err := strconv.Atoi(path[0])
	if err != nil || len(path[0]) != 4 || year < 1 {
		return false, errors.New("L'anno inserito non è valido")
	}
	month, err := strconv.Atoi(path[1])
	if err != nil || month < 1 || month > 12 {
		return false, errors.New("Il mese inserito non è valido")
	}
	day, err := strconv.Atoi(path[2])
	if err != nil || day < 1 || day > daysInMonth(month, year) {
		return false, errors.New("Il giorno inserito non è valido")
	}
	return true, nil
}

func daysInMonth(month, year int) int {
	return time.Date(year, time.Month(month+1), 0, 0, 0, 0, 0, time.UTC).Day()
}
