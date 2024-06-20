package routes

import (
	controller "root/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SettoriRoute(db *gorm.DB,router *gin.Engine) {
	userGroup := router.Group("server/settori/api/method")
  userGroup.GET("/get", controller.GetSettori(db))
	userGroup.POST("/post", controller.CreateSettori(db))
	userGroup.DELETE("/delete/:id", controller.DeleteSettori(db))
	userGroup.PUT("/put/:id", controller.UpdateSettori(db))
}
