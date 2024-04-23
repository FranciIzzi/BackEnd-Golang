package validators

import (
	"errors"
	"fmt"
	"root/config"
	"root/models"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var jwtKey = []byte("go-secret-8nao4^v267!oztl%io6-bneu#27!111qu(nim$&er3r&0n55t4")

type CustomClaims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

func ValidateToken(tokenString string) (string, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	if claims.Role != "h3ad5h0t" {
		return "", errors.New("invalid role")
	}
  
  bool,_:=ValidateEmail(claims.Email)
	if !bool {
		return "", errors.New("invalid email")
	}

	var user models.User
	result := config.DB.Where("id = ? AND email = ?", claims.ID, claims.Email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return "", errors.New("user not found")
	} else if result.Error != nil {
		return "", result.Error
	}

	return claims.Email, nil
}
