package routes

import (
	"github.com/gin-gonic/gin"
	controller "root/controllers"
)

func ChangePwRoute(router *gin.Engine) {
	router.POST("/impact/v1/api/rest/change/password/", controller.ChangePassword)
	router.POST("/impact/v1/api/rest/password/forget/", controller.ForgetPassword)
  router.GET("/impact/v1/api/rest/restore/password/set/new", controller.PasswordPostForget)

}
