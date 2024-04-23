package routes

import (
	"github.com/gin-gonic/gin"
	controller "root/controllers"
	middlewares "root/middlewares"
)

func UserRoute(router *gin.Engine) {
	userGroup := router.Group("impact/v1/api/rest/user")
	userGroup.Use(middlewares.AuthMiddleware())

	userGroup.GET("/", controller.GetUsers)
	userGroup.POST("/", controller.CreateUser)
	userGroup.DELETE("/:id", controller.DeleteUser)
	userGroup.PUT("/:id", controller.UpdateUser)
}

