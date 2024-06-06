package routes

import (
	controller "root/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InumazioniRoute(db *gorm.DB,router *gin.Engine) {
	userGroup := router.Group("server/inumazioni/api/method")
  userGroup.GET("/get", controller.GetInumazioni(db))
	userGroup.POST("/post", controller.CreateInumazioni(db))
	userGroup.DELETE("/delete/:id", controller.DeleteInumazioni(db))
	userGroup.PUT("/put/:id", controller.UpdateInumazioni(db))
}

