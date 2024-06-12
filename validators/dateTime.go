package validators

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

// Data e Time di riferimento 2024-06-10T14:30:00.000Z
func ValidateDateTime(dateTime string) (bool, error) {
	print("errore:\n")
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

// Arriva questo 2024-06-10T14:30:00.000Z
// Data di riferimento da storare questa 2024-06-10
func ValidateDate(date string) (bool, error) {
	deficit := strings.Split(date, "T")
	path := strings.Split(deficit[0], "-")
	if len(path) != 3 {
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

func ConvertStringToDate(string *string) string {
	convert := strings.Split(*string, "T")
	return convert[0]
}
func ConvertStringToDateTime(string *string) string {
	convert := strings.Split(*string, ":")
	res := convert[0] + ":" + convert[1]
	return res
}
