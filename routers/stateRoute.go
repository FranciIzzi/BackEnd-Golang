package routes

import (
	controller "root/controllers"

	"github.com/gin-gonic/gin"
)

func StateRoute(router *gin.Engine) {
	router.POST("impact/v1/jwt/auth/login/", controller.Login)
	router.POST("impact/v1/api/rest/auth/logout/", controller.Logout)
	router.POST("impact/v1/auth/rest/api/jwt/refresh/", controller.RefreshToken)
}

