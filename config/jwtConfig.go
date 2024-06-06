package config

import (
	"root/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateToken(id int, email string) (string, string, error) {
	access := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"type":  "access",
			"id":    id,
			"email": email,
			"role":  "h3ad5h0t",
			"exp":   time.Now().Add(time.Minute * 30).Unix(),
		})

	refresh := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"type":  "refresh",
			"id":    id,
			"email": email,
			"role":  "h3ad5h0t",
			"exp":   time.Now().Add(time.Hour * 2).Unix(),
		})

	accessString, er := access.SignedString(
		[]byte("go-secret-8nao4^v267!oztl%io6-bneu#27!111qu(nim$&er3r&0n55t4"),
	)
	refreshString, err := refresh.SignedString(
		[]byte("go-secret-8nao4^v267!oztl%io6-bneu#27!111qu(nim$&er3r&0n55t4"),
	)
	if er != nil {
		return "Invalid to create token", "access", er
	}

	if err != nil {
		return "Invalid to create token", "refresh", err
	}

	var user models.User
	result := DB.Where("email = ?", email).First(&user).Update("token", refreshString)

	if result.Error != nil {
		return "error", "User token not updated", result.Error
	}

	return accessString, refreshString, nil
}
