package controllers

import (
	"net/http"
	"root/config"
  "root/validators"
	"root/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func ChangePassword(c *gin.Context) {
	 reqToken := c.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato token non valido"})
		return
	}
	tokenString := splitToken[1]

	email, err := validators.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token non valido"})
		return
	}

	var changePwReq struct {
		OldPassword string `json:"oldPassword"`
		NewPassword string `json:"newPassword"`
	}
	if err := c.BindJSON(&changePwReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
		return
	}

	var user models.User

	config.DB.Where("email=?", email).First(&user)
	pwDB := user.Password
	bool := validators.CheckPasswordWithHash(changePwReq.OldPassword, pwDB)
	if !bool {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenziali non valide"})
		return
	}

	bin, _ := validators.ValidatePassword(changePwReq.NewPassword, email)
	if !bin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nuova Password not secure"})
	}
	user.Password, _ = validators.HashPassword(changePwReq.NewPassword)
	config.DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"message": "Password aggiornata con successo"})
	return
}


func ForgetPassword(c *gin.Context) {
	var request struct {
		Email string `json:"email"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
		return
	}
	reqToken := c.GetHeader("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	if len(splitToken) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Formato token non valido"})
		return
	}
	tokenString := splitToken[1]

	email, err := validators.ValidateToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token non valido"})
		return
	}
	if email != request.Email {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not Allowed to do this request"})
    return
	}
	var user models.User
	result := config.DB.Where("email = ?", request.Email).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Email non trovata"})
		return
	}

	_, refresh, err := config.CreateToken(int(user.ID), user.Email)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Errore nella generazione del token di reset"},
		)
		return
	}

	if err := validators.SendWelcomeEmail(user.Email, refresh); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore nell'invio dell'email"})
		return
	  }

	c.JSON(http.StatusOK, gin.H{"message": "Email di reset inviata con successo"})
}


func PasswordPostForget(c *gin.Context){
  return
}
