package controllers

import (
	"net/http"
	"root/config"
	"root/models"
	"root/validators"
	"strings"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BindJSON(&loginReq); err != nil {
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
	if email != loginReq.Email {
		c.JSON(http.StatusForbidden, gin.H{"error": "Not Allowed to do this request"})
		return
	}

	var user models.User
	access, refresh, err := config.CreateToken(int(user.ID), user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Impossibile generare il token"})
		return
	}
	user.Token = refresh
	user.Logged = true
	config.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"access": access, "refresh": refresh})
	return
}

func Logout(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refreshToken"`
	}
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Richiesta non valida"})
		return
	}

	// user, err := models.GetUserByRefreshToken(request.RefreshToken)
	// if err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Token non valido"})
	// 	return
	// }
	var user models.User
	// Invalida il refresh token rimuovendolo o settandolo a un valore vuoto
	err := config.DB.Model(&user).Update("refresh_token", "").Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Errore nell'invalidare il token"})
		return
	}

	// Risposta di successo
	c.JSON(http.StatusOK, gin.H{"message": "Logout effettuato con successo"})
	return
}

func RefreshToken(c *gin.Context) {
	token := c.Param("refresh")
	email, err := validators.ValidateToken(token)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err})
		return
	}

	var user models.User
	result := config.DB.Where("email = ?", email).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	access, refresh, err := config.CreateToken(int(user.ID), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create tokens"})
		return
	}

	result = config.DB.Model(&user).Update("token", refresh)
	if result.Error != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "Could not update user refresh token"},
		)
		return
	}

	c.JSON(http.StatusOK, gin.H{"access": access, "refresh": refresh})
}
